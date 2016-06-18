package main

import (
	"math"

	"github.com/luxengine/lux/glm"
	"github.com/vova616/chipmunk/vect"
)

func AngleBetweenPoints(px, py, cx, cy float32) float32 {
	return float32(math.Atan2(float64(cx-px), float64(cy-py)))
}

func AngleBetweenAngles(pa, ca float32) float32 {
	cv := vect.FromAngle(vect.Float(ca))
	pv := vect.FromAngle(vect.Float(pa))

	sin := pv.X*cv.Y - cv.X*pv.Y
	cos := pv.X*cv.X + pv.Y*cv.Y

	return float32(math.Atan2(float64(sin), float64(cos)))
}

func getCursorPos(x, z float32, w, h int, campos glm.Vec3) (float32, float32) {
	x = (x-float32(w)/2)/500*campos.Y + campos.X
	z = (z-float32(h)/2)/500*campos.Y + campos.Z

	return x, z
}
