package env

import (
	"phys"
	"time"

	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/luxengine/lux/glm"
)

var (
	Space = phys.NewSpace()

	CameraMovement func()
	PlayerMovement func(float32)
)

func Physloop() {
	lastFrame := time.Now()
	var dt float32
	for {
		thisFrame := time.Now()
		dt = float32(thisFrame.Sub(lastFrame).Seconds())

		glfw.PollEvents()

		CameraMovement()
		PlayerMovement(dt)

		// for _, object := range objects {
		// 	if object.Shape != nil {
		// 		milk(object, dt)
		// 	}
		// }

		for object, enabled := range dynamics {
			if object.Param != nil && enabled {
				milk(object, dt)
			}
		}

		Space.Step(dt)
		// log.Println(Space.ContactBuffer)
		lastFrame = thisFrame
	}
}

func milk(object *Object, dt float32) {
	vel := object.Shape.Body.Velocity()
	vel2 := object.Shape.Body.Velocity()
	vel2.Mult(dt)
	vel.Sub(vel2)
	object.Shape.Body.SetVelocity(vel.X, vel.Y)

	avel := object.Shape.Body.AngularVelocity()
	object.Shape.Body.SetAngularVelocity(avel - avel*dt)
}

func phys2gl() {
	for object, enabled := range dynamics {
		if object.Shape == nil && !enabled {
			continue
		}

		y := object.Trans.LocalToWorld.Col(3).Y
		pos := object.Shape.Body.Position()
		ang := object.Shape.Body.Angle()

		mat := glm.Translate3D(pos.X, y, pos.Y)

		if object.Param != nil {
			q := glm.AnglesToQuat(0, ang, object.Param.SubAngle, 1)
			qm4 := q.Mat4()
			mat = mat.Mul4(&qm4)
		} else {
			q := glm.AnglesToQuat(0, ang, 0, 1)
			qm4 := q.Mat4()
			mat = mat.Mul4(&qm4)
		}

		object.Trans.LocalToWorld = mat
	}
}
