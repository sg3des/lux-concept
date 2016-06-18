package env

import (
	"log"

	lux "github.com/luxengine/lux/render"
)

var (
	lights []*Light
)

//Light object
type Light struct {
	Name   string
	Point  lux.PointLight
	Shadow bool

	Callback func(dt float64)

	shadowfbo *lux.ShadowFBO
}

// func initLights() {
// 	lights = make(map[string]*Light)
// }

//NewLight create new light
func NewLight(name string, shadow bool) *Light {
	l := &Light{}

	l.Name = name
	l.Shadow = shadow
	l.Point.Move(0, 10, 0)
	// l.Point.CastsShadow(shadow)

	var err error
	l.shadowfbo, err = lux.NewShadowFBO(4096, 4096)
	if err != nil {
		log.Fatal(err)
	}
	l.shadowfbo.SetOrtho(-10, 10, -10, 10, 0, 100)

	lights = append(lights, l)
	return l
}

func (l *Light) SetCallback(f func(dt float64)) {
	l.Callback = f
}

func (l *Light) Move(x, y, z float32) {
	l.Point.X += x
	l.Point.Y += y
	l.Point.Z += z
}

func (l *Light) SetPosition(x, y, z float32) {
	l.Point.Move(x, y, z)
	// l.Point.X = x
	// l.Point.Y = y
	// l.Point.Z = z
}
