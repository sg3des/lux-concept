package env

import lux "github.com/luxengine/lux/render"

//NewCamera create camera #temporary hardcode
func NewCamera() (cam lux.Camera) {
	cam.SetPerspective(90, float32(int32(WindowWidth))/float32(WindowHeight), 0.1, 100.0)
	cam.LookAtval(0, 50, 5, 0, -200, 0, 0, 100, 0)
	return
}
