package env

import (
	"time"

	"github.com/luxengine/lux/glm"

	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
)

var (
	Space = chipmunk.NewSpace()

	PlayerMovement func(float32)
	PlayerControl  func(float32)

	CameraMovement func()
)

type Phys struct {
	W, H, Mass float32
}

func Physloop() {
	lastFrame := time.Now()
	var dt float32
	for {
		thisFrame := time.Now()
		dt = float32(thisFrame.Sub(lastFrame).Seconds())

		glfw.PollEvents()

		// if Movement != nil {
		CameraMovement()
		PlayerMovement(dt)
		// PlayerControl(dt)

		for _, object := range objects {
			if object.Player != nil && object.Shape != nil {
				milk(object, dt)
			}
		}
		// }

		Space.Step(vect.Float(dt))
		lastFrame = thisFrame
	}
}

func milk(object *Object, dt float32) {
	vel := object.Shape.Body.Velocity()
	vel2 := object.Shape.Body.Velocity()
	vel2.Mult(vect.Float(dt))
	vel.Sub(vel2)
	object.Shape.Body.SetVelocity(float32(vel.X), float32(vel.Y))

	avel := object.Shape.Body.AngularVelocity()
	object.Shape.Body.SetAngularVelocity(avel - avel*dt)
}

func phys2gl() {
	for _, object := range objects {
		if object.Shape == nil {
			continue
		}

		// fmt.Println(object)
		// update position
		pos := object.Shape.Body.Position()
		ang := float32(object.Shape.Body.Angle())
		// object.Trans.SetTranslate(float32(pos.X), 1, float32(pos.Y))

		// update rotation
		// ang := object.Shape.Body.Angle()

		mat := glm.Translate3D(float32(pos.X), 1, float32(pos.Y))

		// rot := glm.HomogRotate3DY(float32(ang))

		if object.Player != nil {

			q := glm.AnglesToQuat(0, ang, object.Player.SubAngle, 1)
			qm4 := q.Mat4()
			mat = mat.Mul4(&qm4)
			// q.Mat4()

			// mat = q.Mat4().Add(&mat)

			// mat.Add(q.Mat4())

			// roll := glm.HomogRotate3DX(object.Player.SubAngle)
			// mat = mat.Mul4(&roll)
			// object.Trans.SetQuatRotate(float32(ang), &glm.Vec3{0, 1, 0})
			// object.Node.LocalRotation = mgl32.AnglesToQuat(0, float32(ang), 0, 1)
			// } else {
			// 	object.Node.LocalRotation = mgl32.AnglesToQuat(0, float32(ang), object.Player.SubAngle, 1)

			// 	object.Shape.Shape().GetAsBox().Width = 2 - ang
		} else {

			q := glm.AnglesToQuat(0, ang, 0, 1)
			qm4 := q.Mat4()
			mat = mat.Mul4(&qm4)
		}

		object.Trans.LocalToWorld = mat

	}
}
