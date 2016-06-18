package env

import (
	"math"
	"param"

	"github.com/luxengine/lux/gl"
	"github.com/luxengine/lux/glm"
	lux "github.com/luxengine/lux/render"

	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
)

var (
	objects []*Object

	assman lux.AssetManager
)

//Object is object
type Object struct {
	Name    string
	Mesh    lux.Mesh
	Texture gl.Texture2D
	Trans   *lux.Transform
	Shadow  bool

	Shape *chipmunk.Shape

	Callback func(dt float64)

	Player *param.Player
}

func initAssets() {
	assman = lux.NewAssetManager("assets/", "models/", "shaders/", "textures/")

	assman.LoadModel("skydome.obj", "skydome")
	assman.LoadModel("ground.obj", "ground")
	assman.LoadModel("cube.obj", "cube")
	assman.LoadModel("trapeze.obj", "trapeze")

	assman.LoadTexture("square.png", "square")
	assman.LoadTexture("skydome.png", "skydome")
	assman.LoadTexture("red.png", "red")
	assman.LoadTexture("brown.png", "brown")
}

func NewMesh(name string, modelname string, ph *Phys, texturename string) *Object {
	o := &Object{}

	o.Name = name
	o.Trans = lux.NewTransform()
	o.Texture = assman.Textures[texturename]
	o.Shadow = true

	o.Mesh = assman.Models[modelname]

	if ph != nil {
		o.Shape = chipmunk.NewBox(vect.Vector_Zero, vect.Float(ph.W), vect.Float(ph.H))
		o.Shape.SetElasticity(0.95)

		body := chipmunk.NewBody(vect.Float(ph.Mass), o.Shape.Moment(ph.Mass))
		// body.SetPosition(vect.Vect{0, 0})
		// body.SetAngle(vect.Float(0))
		body.SetMass(vect.Float(ph.Mass))

		body.AddShape(o.Shape)
		Space.AddBody(body)
	}

	// fmt.Println(name, size, len(size))
	// if len(size) > 0 {
	// 	o.Box = NewCubezBox(size)
	// }

	objects = append(objects, o)
	return o
}

func (o *Object) SetCallback(f func(dt float64)) {
	o.Callback = f
}

//Move object from current position
func (o *Object) Move(x, y, z float32) {
	// if o.Box != nil {
	// o.Box.Body.Position.Add(&mathz.Vector3{mathz.Real(x), mathz.Real(y), mathz.Real(z)})

	// x = pos[0] + x
	// y = pos[1] + y
	// z = pos[2] + z
	// o.Box.Body.Position = &glm.Vec3{x, y, z}
	// o.Box..SetPositionVec3(&glm.Vec3{x, y, z})
	// }
	//?????????
	o.Trans.Translate(x, y, z)

}

func (o *Object) Position() (x, y float32) {
	if o.Shape != nil {
		v := o.Shape.Body.Position()
		return float32(v.X), float32(v.Y)
	} else {
		x, _, z := glm.Extract3DScale(&o.Trans.LocalToWorld)
		return x, z
	}
}

//SetPosition instantly move object to global point
func (o *Object) SetPosition(x, y, z float32) {
	if o.Shape != nil {
		o.Shape.Body.SetPosition(vect.Vect{X: vect.Float(x), Y: vect.Float(z)})
	}
	o.Trans.SetTranslate(x, y, z)

	// if o.Box != nil {
	// 	o.Box.Body.Position = mathz.Vector3{mathz.Real(x), mathz.Real(y), mathz.Real(z)}

	// }

	// o.Trans.Translate(x, y, z)

}

func (e *Object) VectorForward(scale float32) (float32, float32) {
	// _, y, _, _ := Quat2Rad(e.Node.LocalRotation)

	y := float64(e.Shape.Body.Angle())

	xa := float32(math.Sin(y)) * scale
	za := float32(math.Cos(y)) * scale

	// fmt.Println(y, xa, za)

	return xa, za
}
