package env

import lux "github.com/luxengine/lux/render"

//NewCamera create camera #temporary hardcode
func NewCamera() (cam lux.Camera) {
	cam.SetPerspective(0.05, Ratio, 0.1, 200)
	cam.LookAtval(0, 40, 1, 0, 0, 0, 0, 1, 0)
	return
}
