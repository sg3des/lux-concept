package env

import (
	"log"
	"time"

	"github.com/go-gl/glfw/v3.1/glfw"
	lux "github.com/luxengine/lux/render"
)

var (
	WindowWidth  = 1200
	WindowHeight = 800
	Ratio        = float32(WindowWidth) / float32(WindowHeight)

	MainCam lux.Camera //MainCamera
)

//Render is main object contains window, etc
type Render struct {
	Window *glfw.Window
	gbuf   lux.GBuffer

	tonemap *lux.PostProcessFramebuffer
	fxaa    *lux.PostProcessFramebuffer
}

func initWindow() (r *Render, err error) {
	r = &Render{}

	lux.InitGLFW()

	r.Window = lux.CreateWindow(WindowWidth, WindowHeight, "", false)

	r.gbuf, err = lux.NewGBuffer(int32(WindowWidth), int32(WindowHeight))
	if err != nil {
		return
	}

	lux.InitPostProcessSystem()

	//shadows tones etc...
	r.fxaa, err = lux.NewPostProcessFramebuffer(int32(WindowWidth), int32(WindowHeight), lux.PostprocessfragmentshaderFxaa)
	if err != nil {
		return
	}

	r.tonemap, err = lux.NewPostProcessFramebuffer(int32(WindowWidth), int32(WindowHeight), lux.PostProcessFragmentShaderToneMapping)
	if err != nil {
		return
	}

	r.fxaa.SetNext(r.tonemap)

	//create MainCam
	MainCam = NewCamera()

	return
}

//Loop is main loop
func (r *Render) Loop() {
	lastFrame := time.Now()
	var fps int
	for !r.Window.ShouldClose() {
		if r.Window.GetKey(glfw.KeyEscape) == glfw.Press {
			break
		}
		fps++
		thisFrame := time.Now()
		if lastFrame.Add(time.Second).Before(thisFrame) {
			log.Println(fps)
			fps = 0
			lastFrame = thisFrame
		}

		phys2gl()

		// lights
		for _, light := range lights {
			if light.shadowfbo != nil {
				light.shadowfbo.BindForDrawing()

				for object, enabled := range objects {
					if object.Shadow && enabled {
						light.shadowfbo.Render(object.Mesh, object.Trans)
					}
				}
				light.shadowfbo.Unbind()
				light.shadowfbo.LookAt(light.Point.X, light.Point.Y, light.Point.Z, 0, 0, 0)
			}
		}

		r.gbuf.Bind(&MainCam)

		// render objects
		for object, enabled := range objects {
			if enabled {
				r.gbuf.Render(&MainCam, object.Mesh, object.Texture, object.Trans)
			}
		}

		// render shadows
		for _, light := range lights {
			if light.shadowfbo != nil {
				r.gbuf.RenderLight(&MainCam, &light.Point, light.shadowfbo.ShadowMat(), light.shadowfbo.ShadowMap(), 10, 0.2, light.Intensity/100)
			}
		}

		// aggregate
		r.gbuf.Aggregate()

		r.tonemap.PreRender()
		r.tonemap.Render(r.gbuf.AggregateFramebuffer.Out)
		r.tonemap.PostRender()

		r.Window.SwapBuffers()
	}
}

func (r *Render) Clean() {
	assman.Clean()
}
