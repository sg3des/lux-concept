package param

type Player struct {
	Name   string
	Object Object

	MovSpeed, RotSpeed, SubAngle float32

	LeftWeapon, RightWeapon *Weapon
}

type Weapon struct {
	BulletObject Object

	Damage       float32
	AttackRate   float32
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
