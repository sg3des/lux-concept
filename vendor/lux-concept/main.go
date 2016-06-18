package main

import (
	"env"
	"log"
	"param"
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

	initCursor()

	// env.Movement =
	log.SetFlags(log.Lshortfile)
}

func main() {
	loadScene("name")
	newPlayer()

	go env.Physloop()
	r.Loop()
	r.Clean()
}

func loadScene(name string) {
	// env.NewSquare()

	env.NewMesh(param.Object{
		Name: "ground",
		Mesh: param.Mesh{Model: "ground", Texture: "grass"},
		Pos:  param.Pos{Y: -15},
	})

	env.NewLight(&param.Light{
		Shadow:    true,
		Intensity: 18,
		Pos:       param.Pos{X: 2, Y: 15, Z: 2},
	})
	env.NewLight(&param.Light{
		Intensity: 10,
		Pos:       param.Pos{X: -2, Y: 15, Z: -4},
	})

	env.NewMesh(param.Object{
		Name: "box0",
		Mesh: param.Mesh{Model: "cube", Texture: "brown", Shadow: true},
		Pos:  param.Pos{Z: 5},
		PH:   param.Phys{W: 2, H: 2, Mass: 8},
	})

	env.NewMesh(param.Object{
		Name: "box1",
		Mesh: param.Mesh{Model: "cube", Texture: "brown", Shadow: true},
		Pos:  param.Pos{Z: 10},
		PH:   param.Phys{W: 2, H: 2, Mass: 50},
	})

	env.NewMesh(param.Object{
		Name: "tree",
		Mesh: param.Mesh{Model: "tree", Texture: "brown"},
		Pos:  param.Pos{X: 5, Y: -15, Z: 15},
		PH:   param.Phys{W: 4, H: 4},
	})
}

func initCursor() {
	cursor = env.NewMesh(param.Object{
		Name: "cursor",
		Mesh: param.Mesh{Model: "cursor", Texture: "green"},
	})
	env.CameraMovement = cameraMovement
}

func cameraMovement() {
	campos := env.MainCam.Pos
	x, z := r.Window.GetCursorPos()
	w, h := r.Window.GetSize()
	cx, cz := getCursorPos(float32(x), float32(z), w, h, campos)

	cursor.SetPosition(cx, 1, cz)
}
