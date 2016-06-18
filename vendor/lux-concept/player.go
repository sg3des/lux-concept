package main

import (
	"env"
	"param"
	"phys"

	"github.com/go-gl/glfw/v3.1/glfw"
)

var (
	localPlayer *env.Object
	// arbiters []*chipmunk.Arbiter
)

func newPlayer() {
	p := &param.Player{
		Name: "player0",
		Object: param.Object{
			Name: "player",
			Mesh: param.Mesh{Model: "trapeze", Texture: "brown", Shadow: true},
			PH:   param.Phys{W: 2, H: 2, Mass: 10},
		},

		MovSpeed: 4,
		RotSpeed: 15,

		LeftWeapon: &param.Weapon{
			BulletObject: param.Object{
				Name: "bullet",
				Mesh: param.Mesh{Model: "cube", Texture: "brown", Shadow: true},
				PH:   param.Phys{W: 1, H: 1, Mass: 1},
			},
			BulletSpeed: 30,
		},
		RightWeapon: &param.Weapon{
			BulletObject: param.Object{
				Name: "bullet",
				Mesh: param.Mesh{Model: "cube", Texture: "brown", Shadow: true},
				PH:   param.Phys{W: 1, H: 1, Mass: 1},
			},
			BulletSpeed: 20,
		},
	}
	localPlayer = env.NewMesh(p.Object)
	localPlayer.Param = p

	env.PlayerMovement = playerMovement
	r.Window.SetMouseButtonCallback(mouseCountrol)
}

func playerMovement(dt float32) {
	cpx, cpy := cursor.Position()
	LookAtTarget(localPlayer, cpx, cpy, dt)

	dist := Distance(cursor, localPlayer)

	localPlayer.Shape.Body.AddVelocity(localPlayer.VectorForward(dt * localPlayer.Param.MovSpeed * 0.08 * dist))
}

func mouseCountrol(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {

	if button == 0 && action == 1 {
		Fire(localPlayer.Param.LeftWeapon, localPlayer)
	}

	if button == 1 && action == 1 {
		Fire(localPlayer.Param.RightWeapon, localPlayer)
	}
}

func Fire(w *param.Weapon, p *env.Object) {
	x, y := localPlayer.Position()

	bullet := env.NewMesh(w.BulletObject)
	bullet.SetPosition(x, 1, y)
	bullet.Shape.Body.SetVelocity(localPlayer.VectorForward(w.BulletSpeed))
	// bullet.Shape.Body.SetAngle()

	bullet.Parent = localPlayer
	bullet.Shape.Body.CallBackCollision = collision
}

func collision(arv *phys.Arbiter) bool {
	if arv.BodyA.UserData == nil || arv.BodyB.UserData == nil {
		return true
	}

	a := arv.BodyA.UserData.(*env.Object)
	b := arv.BodyB.UserData.(*env.Object)

	if a.Parent != nil && a.Parent == b {
		return false
	}

	if b.Parent != nil && b.Parent == a {
		return false
	}

	return true
}