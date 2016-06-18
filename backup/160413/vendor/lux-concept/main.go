package main

import (
	"env"
	"log"
)

var (
	r   *env.Render
	err error

	cursor *env.Object
)

func init() {
	r, err = env.Initialize()
	if err != nil {
		log.Fatal("failed initialize")
	}

	env.CameraMovement = cameraMovement

	// env.Movement =
}

func main() {
	sky := env.NewMesh("sky", "skydome", nil, "skydome")
	sky.Shadow = false

	ground := env.NewMesh("ground", "ground", nil, "square")
	// ground.SetRigidBodyBox(glm.Vec3{999, 0.5, 999})
	ground.SetPosition(0, -2.5, 0)

	// === lighs === //
	lamp := env.NewLight("lamp", true)
	lamp.SetPosition(2, 5, 2)

	// light.Point.Move(0, 2, 1)
	// light.SetPosition(2, 10, 2)
	// light.SetCallback(func(dt float64) {
	// 	// light.Move(3*math.Cos(float32(dt/2)), 0, 3*math.Sin(float32(dt/2)))
	// 	// light.lamp.Move(3*math.Cos(float32(dt/2)), 3, 3*math.Sin(float32(dt/2)))
	// })

	sun := env.NewLight("sun", false)
	sun.SetPosition(-2, 20, -2)

	cursor = env.NewMesh("cursor", "cube", &env.Phys{W: 4, H: 4, Mass: 1000}, "brown")

	newPlayer()

	go env.Physloop()
	r.Loop()
	r.Clean()
}

func cameraMovement() {
	campos := env.MainCam.Pos

	x, z := r.Window.GetCursorPos()
	w, h := r.Window.GetSize()

	cx, cz := getCursorPos(float32(x), float32(z), w, h, campos)

	cursor.SetPosition(cx, 1, cz)
}
