package main

import (
	"env"
	"param"
	"phys"
	"time"

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

		Health:   100,
		MovSpeed: 4,
		RotSpeed: 15,

		LeftWeapon: &param.Weapon{
			BulletObject: param.Object{
				Name: "bullet",
				Mesh: param.Mesh{Model: "bullet", Texture: "green"},
				PH:   param.Phys{W: 0.1, H: 0.1, Mass: 0.001},
			},
			X:           -1,
			Damage:      20,
			AttackRate:  1 * time.Millisecond,
			BulletSpeed: 300,
		},
		RightWeapon: &param.Weapon{
			BulletObject: param.Object{
				Name: "bullet",
				Mesh: param.Mesh{Model: "bullet", Texture: "red"},
				PH:   param.Phys{W: 0.1, H: 0.1, Mass: 0.1},
			},
			X:           1,
			Damage:      50,
			AttackRate:  500 * time.Millisecond,
			BulletSpeed: 20,
		},
	}
	localPlayer = env.NewMesh(p.Object)
	localPlayer.Param = p

	env.PlayerMovement = playerMovement
	r.Window.SetMouseButtonCallback(mouseCountrol)
}

func playerMovement(dt float32) {
	cpx, cpz := cursor.Position()
	LookAtTarget(localPlayer, cpx, cpz, dt)

	dist := Distance(cursor, localPlayer)

	localPlayer.Shape.Body.AddVelocity(localPlayer.VectorForward(dt * localPlayer.Param.MovSpeed * 0.08 * dist))

	Fire(localPlayer.Param.LeftWeapon, localPlayer)
	Fire(localPlayer.Param.RightWeapon, localPlayer)
}

func mouseCountrol(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {

	if button == 0 {
		localPlayer.Param.LeftWeapon.Shoot = action == 1
	}

	if button == 1 {
		localPlayer.Param.RightWeapon.Shoot = action == 1
	}
}

func Fire(w *param.Weapon, p *env.Object) {
	if w.Shoot && w.NextShot.Before(time.Now()) {
		vx, vz := localPlayer.VectorSide(1)
		x, z := localPlayer.Position()

		bullet := env.NewMesh(w.BulletObject)
		bullet.SetPosition(x+w.X*vx, 1, z+w.X*vz)
		bullet.SetRotation(localPlayer.Rotation())
		bullet.Shape.Body.SetVelocity(localPlayer.VectorForward(w.BulletSpeed))

		bullet.Parent = localPlayer
		bullet.Shape.Body.CallBackCollision = BulletCollision

		w.NextShot = time.Now().Add(w.AttackRate)
	}
}

func BulletCollision(arb *phys.Arbiter) bool {
	if arb.BodyA.UserData == nil || arb.BodyB.UserData == nil {
		return true
	}

	var bullet *env.Object
	var target *env.Object

	if arb.BodyA.UserData.(*env.Object).Name == "bullet" {
		bullet = arb.BodyA.UserData.(*env.Object)
		target = arb.BodyB.UserData.(*env.Object)
	} else {
		bullet = arb.BodyB.UserData.(*env.Object)
		target = arb.BodyA.UserData.(*env.Object)
	}

	if bullet.Parent == target {
		return false
	}

	bullet.Destroy()

	return true
}
