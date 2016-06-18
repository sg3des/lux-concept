package env

import (
	"log"
	"param"

	lux "github.com/luxengine/lux/render"
)

var (
	lights []*Light
)

//Light object
type Light struct {
	Name      string
	Point     lux.PointLight
	Shadow    bool
	Intensity float32

	shadowfbo *lux.ShadowFBO
}

//NewLight create new light
func NewLight(light *param.Light) *Light {
	l := &Light{
		Name:      light.Name,
		Shadow:    light.Shadow,
		Intensity: light.Intensity,
	}
	// l.Point.Move(0, 10, 0)
	l.Point.X = light.Pos.X
	l.Point.Y = light.Pos.Y
	l.Point.Z = light.Pos.Z

	// l.Point.SetColor(255, 255, 255)

	// if light.Shadow {
	var err error
	l.shadowfbo, err = lux.NewShadowFBO(4096, 4096)
	if err != nil {
		log.Fatal(err)
	}
	l.shadowfbo.SetOrtho(-5, 5, -5, 5, 1, 100)
	// }

	lights = append(lights, l)
	return l
}

func (l *Light) Move(x, y, z float32) {
	l.Point.X += x
	l.Point.Y += y
	l.Point.Z += z
}

func (l *Light) SetPosition(x, y, z float32) {
	// l.Point.Move(x, y, z)
	l.Point.X = x
	l.Point.Y = y
	l.Point.Z = z
}
