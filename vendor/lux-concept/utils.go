package main

import (
	"env"
	"math"
	"phys/vect"

	"github.com/luxengine/lux/glm"
)

func AngleBetweenPoints(px, py, cx, cy float32) float32 {
	return float32(math.Atan2(float64(cx-px), float64(cy-py)))
}

func AngleBetweenAngles(pa, ca float32) float32 {
	cv := vect.FromAngle(ca)
	pv := vect.FromAngle(pa)

	sin := pv.X*cv.Y - cv.X*pv.Y
	cos := pv.X*cv.X + pv.Y*cv.Y

	return float32(math.Atan2(float64(sin), float64(cos)))
}

func getCursorPos(x, z float32, w, h int, campos glm.Vec3) (float32, float32) {
	x = (x-float32(w)/2)/670*campos.Y + campos.X
	z = (z-float32(h)/2)/670*campos.Y + campos.Z

	return x, z
}

func LookAtTarget(o *env.Object, cpx, cpy, dt float32) {
	pp := o.Shape.Body.Position()

	ca := AngleBetweenPoints(pp.X, pp.Y, cpx, cpy)
	pa := o.Shape.Body.Angle()

	cv := vect.FromAngle(ca)
	pv := vect.FromAngle(pa)

	vx := pv.X + (pv.X+cv.X)*(dt*o.Param.RotSpeed*0.1)
	vy := pv.Y + (pv.Y+cv.Y)*(dt*o.Param.RotSpeed*0.1)

	subAngle := AngleBetweenAngles(pa, ca)

	if subAngle > -o.Param.SubAngle && o.Param.SubAngle > -1.5 {
		o.Param.SubAngle -= dt * localPlayer.Param.RotSpeed * 0.05
	}

	if subAngle < -o.Param.SubAngle && o.Param.SubAngle < 1.5 {
		o.Param.SubAngle += dt * o.Param.RotSpeed * 0.05
	}

	//rotate
	o.Shape.Body.SetAngle(float32(math.Atan2(float64(vy), float64(vx))))
}

func Distance(o1, o2 *env.Object) float32 {
	cpx, cpy := o1.Position()
	ppx, ppy := o2.Position()

	dist := math.Sqrt(math.Pow(float64(cpx)-float64(ppx), 2) + math.Pow(float64(cpy)-float64(ppy), 2))
	return float32(dist)
}

// func Vec2Angle(xa, za float32) float32 {
// 	vect.Vect{xa, za}.
// }
