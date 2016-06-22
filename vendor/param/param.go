package param

import "time"

type Player struct {
	Name   string
	Object Object

	Health, MovSpeed, RotSpeed, SubAngle float32

	LeftWeapon, RightWeapon *Weapon
}

type Weapon struct {
	NextShot time.Time
	Shoot    bool

	BulletObject Object

	X float32

	Damage       float32
	AttackRate   time.Duration
	BulletSpeed  float32
	BulletRotate float32
	Lifetime     float32
}

type Object struct {
	Name string
	Mesh Mesh
	Pos  Pos
	PH   Phys
}

type Mesh struct {
	Model, Texture string
	Shadow         bool
}

type Pos struct {
	X, Y, Z float32
}

type Phys struct {
	W, H, Mass float32
}

type Light struct {
	Name      string
	Shadow    bool
	Intensity float32
	Pos       Pos
}
