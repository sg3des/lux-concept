package env

import (
	"math"

	"github.com/luxengine/lux/gl"
	lux "github.com/luxengine/lux/render"

	"param"
	"phys"
	"phys/vect"
)

var (
	dynamics []*Object
	objects  []*Object
	assman   lux.AssetManager
)

//Object is object
type Object struct {
	Name    string
	Mesh    lux.Mesh
	Texture gl.Texture2D
	Trans   *lux.Transform
	Shadow  bool

	Shape *phys.Shape

	Param  *param.Player
	Parent *Object
}

func initAssets() {
	assman = lux.NewAssetManager("assets/", "models/", "shaders/", "textures/")

	assman.LoadModel("skydome.obj", "skydome")
	assman.LoadModel("plane.obj", "plane")
	assman.LoadModel("square.obj", "square")
	assman.LoadModel("ground.obj", "ground")
	assman.LoadModel("cube.obj", "cube")
	assman.LoadModel("tree.obj", "tree")
	assman.LoadModel("trapeze.obj", "trapeze")
	assman.LoadModel("cursor.obj", "cursor")

	assman.LoadTexture("square.png", "square")
	assman.LoadTexture("grass.png", "grass")
	assman.LoadTexture("skydome.png", "skydome")
	assman.LoadTexture("red.png", "red")
	assman.LoadTexture("green.png", "green")
	assman.LoadTexture("brown.png", "brown")
}

func NewSquare() {
	o := &Object{}
	o.Trans = lux.NewTransform()
	// o.Mesh = lux.NewVUNModel(
	// 	[]uint16{4},
	// 	[]glm.Vec3{glm.Vec3{-1, 0, 1}, glm.Vec3{1, 0, 1}, glm.Vec3{-1, 0, -1}, glm.Vec3{1, 0, -1}},
	// 	[]glm.Vec2{glm.Vec2{0.9, 0.1}, glm.Vec2{0.1, 0.9}, glm.Vec2{0.1, 0.1}, glm.Vec2{0.9, 0.9}},
	// 	[]glm.Vec3{glm.Vec3{0, 1, 0}},
	// )
	o.Trans.Translate(-2, 1, -2)
	o.Mesh = assman.Models["square"]
	o.Texture = assman.Textures["red"]
	o.Shadow = true

	o.Trans.Scale(10)

	objects = append(objects, o)
}

func NewMesh(p param.Object) *Object {
	o := &Object{
		Name:    p.Name,
		Trans:   lux.NewTransform(),
		Mesh:    assman.Models[p.Mesh.Model],
		Texture: assman.Textures[p.Mesh.Texture],
		Shadow:  p.Mesh.Shadow,
	}
	o.Trans.Translate(p.Pos.X, p.Pos.Y, p.Pos.Z)

	if p.PH.W > 0 && p.PH.H > 0 {

		if p.PH.Mass == 0 {
			o.SetStaticShape(p.PH, p.Pos)
		} else {
			o.SetShape(p.PH, p.Pos)
			dynamics = append(dynamics, o)
		}

	}

	objects = append(objects, o)
	return o
}

func (o *Object) SetShape(ph param.Phys, pos param.Pos) {
	p := vect.Vect{X: pos.X, Y: pos.Z}
	o.Shape = phys.NewBox(vect.Vector_Zero, ph.W, ph.H)
	o.Shape.SetElasticity(0.6)

	body := phys.NewBody(ph.Mass, o.Shape.Moment(ph.Mass))
	body.SetMass(ph.Mass)
	body.AddShape(o.Shape)
	body.UserData = o

	o.Shape.Body.SetPosition(p)

	Space.AddBody(body)
}

func (o *Object) SetStaticShape(ph param.Phys, pos param.Pos) {
	p := vect.Vect{X: pos.X, Y: pos.Z}
	o.Shape = phys.NewBox(vect.Vector_Zero, ph.W, ph.H)
	o.Shape.SetElasticity(0.6)

	body := phys.NewBodyStatic()
	body.AddShape(o.Shape)
	body.UserData = o

	o.Shape.Body.SetPosition(p)

	Space.AddBody(body)
}

// Move object from current position
func (o *Object) Move(x, y, z float32) {
	o.Trans.Translate(x, y, z)
}

func (o *Object) Position() (x, y float32) {
	if o.Shape != nil {
		v := o.Shape.Body.Position()
		return v.X, v.Y
	}

	v4 := o.Trans.LocalToWorld.Col(3)
	return v4.X, v4.Z
}

//SetPosition instantly move object to global point
func (o *Object) SetPosition(x, y, z float32) {
	if o.Shape != nil {
		o.Shape.Body.SetPosition(vect.Vect{X: x, Y: z})
	}
	o.Trans.SetTranslate(x, y, z)
}

func (e *Object) VectorForward(scale float32) (float32, float32) {
	y := float64(e.Shape.Body.Angle())

	xa := float32(math.Sin(y)) * scale
	za := float32(math.Cos(y)) * scale

	return xa, za
}