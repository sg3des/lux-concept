package main

import (
	"env"
	"math"

	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/vova616/chipmunk/vect"

	"param"
)

var (
	player *env.Object
	mb1    = 0
)

func newPlayer() {
	player = env.NewMesh("player", "trapeze", &env.Phys{W: 2, H: 2, Mass: 8}, "brown")
	player.SetPosition(0, 1, 0)
	player.Player = &param.Player{RotSpeed: 15, MovSpeed: 4}

	env.PlayerMovement = playerMovement
	// env.PlayerControl = playerControl
	r.Window.SetMouseButtonCallback(mouseCountrol)
	// player.SetCallback(playerMovement)

	box := env.NewMesh("box", "cube", &env.Phys{W: 2, H: 2, Mass: 8}, "brown")
	box.SetPosition(0, 1, -5)

	box1 := env.NewMesh("box1", "cube", &env.Phys{W: 2, H: 2, Mass: 50}, "brown")
	box1.SetPosition(0, 1, -7)

}

func playerMovement(dt float32) {
	//rotate prepare
	cpx, cpy := cursor.Position()
	pp := player.Shape.Body.Position()

	ca := AngleBetweenPoints(float32(pp.X), float32(pp.Y), cpx, cpy)
	pa := player.Shape.Body.Angle()

	cv := vect.FromAngle(vect.Float(ca))
	pv := vect.FromAngle(vect.Float(pa))

	vx := pv.X + (pv.X+cv.X)*vect.Float(dt*float32(player.Player.RotSpeed)*0.1)
	vy := pv.Y + (pv.Y+cv.Y)*vect.Float(dt*float32(player.Player.RotSpeed)*0.1)

	subAngle := AngleBetweenAngles(float32(pa), ca)

	if subAngle > -player.Player.SubAngle && player.Player.SubAngle > -1.5 {
		player.Player.SubAngle -= dt * player.Player.RotSpeed * 0.05
	}

	if subAngle < -player.Player.SubAngle && player.Player.SubAngle < 1.5 {
		player.Player.SubAngle += dt * player.Player.RotSpeed * 0.05
	}

	// rotate
	player.Shape.Body.SetAngle(vect.Float(math.Atan2(float64(vy), float64(vx))))

	//movement
	dist := math.Sqrt(math.Pow(float64(cpx)-float64(pp.X), 2) + math.Pow(float64(cpy)-float64(pp.Y), 2))

	player.Shape.Body.AddVelocity(player.VectorForward(dt * float32(player.Player.MovSpeed) * 0.1 * float32(dist)))
}

func mouseCountrol(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	x, y := player.Position()

	if button == 0 && action == 1 {
		bullet := env.NewMesh("box", "cube", &env.Phys{W: 2, H: 2, Mass: 10}, "brown")
		bullet.SetPosition(x, 1, y)
		bullet.Shape.Body.SetVelocity(player.VectorForward(30))

		// if bullet.Shape.Body.CallbackHandler.CollisionEnter(arbiter) {
		// arbiter.Ignore()
		// }

		// arbiter.

		// fmt.Println(bullet.Shape.Body.BodyActivate())
		// bullet.Shape.Body.CallbackHandler.CollisionEnter(arbiter)
	}
}
