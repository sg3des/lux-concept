package env

import (
	"github.com/go-gl/glfw/v3.1/glfw"
	// _ "github.com/luxengine/glm" //implicit import

	lux "github.com/luxengine/lux/render"
	// "github.com/luxengine/lux/debug"
)

var (
	WindowWidth  = 1200
	WindowHeight = 800

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

	// debug.EnableGLDebugLogging()

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
	previousTime := glfw.GetTime()

	for !r.Window.ShouldClose() {
		if r.Window.GetKey(glfw.KeyEscape) == glfw.Press {
			break
		}

		// dt ????
		time := glfw.GetTime()
		dt := time - previousTime
		previousTime = time

		phys2gl()

		// PhysFrame(float32(dt))
		// tornago
		// w.Step(float32(dt))
		// am.Tick()

		// update
		for _, object := range objects {
			if object.Callback != nil {
				object.Callback(dt)
			}
		}

		// lights
		for _, light := range lights {
			if light.Callback != nil {
				light.Callback(dt)
			}

			light.shadowfbo.BindForDrawing()
			for _, object := range objects {
				if object.Shadow {
					light.shadowfbo.Render(object.Mesh, object.Trans)
				}
			}
			light.shadowfbo.Unbind()
			light.shadowfbo.LookAt(light.Point.X, light.Point.Y, light.Point.Z, 0, 0, 0)
		}

		r.gbuf.Bind(&MainCam)

		// render objects
		for _, object := range objects {
			r.gbuf.Render(&MainCam, object.Mesh, object.Texture, object.Trans)
		}

		// render shadows
		for _, light := range lights {
			r.gbuf.RenderLight(&MainCam, &light.Point, light.shadowfbo.ShadowMat(), light.shadowfbo.ShadowMap(), 0.5, 0.8, 0.8)
		}

		// aggregate
		r.gbuf.Aggregate()

		r.tonemap.PreRender()
		r.tonemap.Render(r.gbuf.AggregateFramebuffer.Out)
		r.tonemap.PostRender()

		r.Window.SwapBuffers()

		//events
		// glfw.PollEvents()
	}
}

func (r *Render) Clean() {
	assman.Clean()
}
