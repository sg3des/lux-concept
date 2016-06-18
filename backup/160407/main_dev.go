package main

import (
	"github.com/go-gl/glfw/v3.1/glfw"

	"log"
	"runtime"

	"github.com/luxengine/lux"
	"github.com/luxengine/lux/debug"
	"github.com/luxengine/math"
)

// Default window size
var (
	WindowWidth  = 1200
	WindowHeight = 800

	assman lux.AssetManager

	angle float32
	cam   lux.Camera
	lamp  lux.PointLight

	window *glfw.Window
	gbuf   lux.GBuffer

	tonemap *lux.PostProcessFramebuffer
	fxaa    *lux.PostProcessFramebuffer

	shadowfbo *lux.ShadowFBO

	// w *tornago.World

	skydome       lux.Mesh
	skydomeTransf *lux.Transform

	err error
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	newWindow()

	// === asset manager testing === //
	assman = lux.NewAssetManager("assets/", "models/", "shaders/", "textures/")
	assman.LoadModel("skydome.obj", "skydome")
	assman.LoadModel("ground.obj", "ground")

	assman.LoadTexture("square.png", "square")
	assman.LoadTexture("skydome.png", "skydome")
	assman.LoadTexture("red.png", "red")
	assman.LoadTexture("brown.png", "brown")

	skydome = assman.Models["skydome"]

	skydomeTransf = lux.NewTransform()

	// ==camera== //
	// var cam lux.Camera
	cam.SetPerspective(70.0, float32(int32(WindowWidth))/float32(WindowHeight), 0.1, 100.0)
	cam.LookAtval(-9, 9, 9, 0, 0, 0, 0, 1, 0)

	// === lights === //
	lamp.Move(0, 2, 1)

	// === shadow === //
	shadowfbo, err = lux.NewShadowFBO(4096, 4096)
	if err != nil {
		log.Fatal(err)
	}
	shadowfbo.SetOrtho(-10, 10, -10, 10, 0, 10)

	// // === tornago === //
	// w = tornago.NewWorld(&tornago.NaiveBroadphase{}, tornago.ContactResolver{})
	// boxBody := tornago.NewRigidBody()
	// boxBody.SetPosition3f(0, 0, 0)
	// boxBody.SetVelocity3f(0, 0, 0)
	// boxShape := tornago.NewCollisionBox(glm.Vec3{0.5, 0.5, 0.5})

	// boxBody.SetCollisionShape(boxShape)
	// boxBody.SetAcceleration3f(0, -5, 0)
	// boxBody.SetRestitution(1)
	// boxBody.SetAngularDamping(0.8)
	// boxBody.SetLinearDamping(0.8)
	// boxBody.SetFriction(2)
	// boxBody.SetMass(1)

	// boxBody2 := tornago.NewRigidBody()
	// boxBody2.SetPosition3f(0.25, 2, 0)
	// boxBody2.SetVelocity3f(0, 0, 0)
	// boxShape2 := tornago.NewCollisionBox(glm.Vec3{0.5, 0.5, 0.5})

	// boxBody2.SetCollisionShape(boxShape2)
	// boxBody2.SetAcceleration3f(0, -10, 0)
	// boxBody2.SetRestitution(1)
	// boxBody2.SetAngularDamping(0.9)
	// boxBody2.SetLinearDamping(0.9)
	// boxBody2.SetFriction(5)
	// boxBody2.SetMass(2)

	// b2 := tornago.NewRigidBody()
	// boxShapeGround := tornago.NewCollisionBox(glm.Vec3{999, 1, 999})
	// b2.SetCollisionShape(boxShapeGround)
	// b2.SetPosition3f(0, -2.5, 0)
	// b2.SetMass(0)
	// b2.SetRestitution(0)
	// b2.SetFriction(0)

	// w.AddRigidBody(b2)
	// w.AddRigidBody(boxBody)
	// w.AddRigidBody(boxBody2)

	// boxModel := lux.NewVUNModel(boxShape.Mesh())

	// boxRBTransf := lux.NewTransform()
	// boxRBTransf2 := lux.NewTransform()

	// am := lux.NewAgentManager()
	// am.NewAgent(func() bool {
	// 	var m glm.Mat4
	// 	boxBody.OpenGLMatrix(&m)
	// 	boxRBTransf.SetMatrix((*[16]float32)(&m))
	// 	return true
	// })
	// am.NewAgent(func() bool {
	// 	var m glm.Mat4
	// 	boxBody2.OpenGLMatrix(&m)
	// 	boxRBTransf2.SetMatrix((*[16]float32)(&m))
	// 	return true
	// })

	loop()

	assman.Clean()
}

func newWindow() {
	// === ok thats pretty clean === //
	lux.InitGLFW()

	window = lux.CreateWindow(WindowWidth, WindowHeight, "", false)
	debug.EnableGLDebugLogging()

	gbuf, err = lux.NewGBuffer(int32(WindowWidth), int32(WindowHeight))
	if err != nil {
		log.Fatal(err)
	}

	// post process //
	lux.InitPostProcessSystem()

	fxaa, err = lux.NewPostProcessFramebuffer(int32(WindowWidth), int32(WindowHeight), lux.PostprocessfragmentshaderFxaa)
	if err != nil {
		log.Fatal(err)
	}
	// defer fxaa.Delete()

	tonemap, err = lux.NewPostProcessFramebuffer(int32(WindowWidth), int32(WindowHeight), lux.PostProcessFragmentShaderToneMapping)
	if err != nil {
		log.Fatal(err)
	}
	// defer tonemap.Delete()

	fxaa.SetNext(tonemap)
}

func loop() {
	// ===init of app specific stuff=== //
	angle := 0.0
	previousTime := glfw.GetTime()

	// ===render loop=== //
	for !window.ShouldClose() {
		if window.GetKey(glfw.KeyEscape) == glfw.Press {
			break
		}

		// ==Update== //
		time := glfw.GetTime()
		elapsed := time - previousTime
		previousTime = time
		angle += elapsed

		// === tornago === //
		// i++
		//if i%10 == 0 {
		// w.Step(float32(elapsed))
		// am.Tick()
		//}

		// ==shadow== //
		shadowfbo.BindForDrawing()
		// shadowfbo.Render(sphere, sphereTransf)
		// shadowfbo.Render(ground, groundTransf)
		// shadowfbo.Render(boxModel, boxRBTransf)
		// shadowfbo.Render(boxModel, boxRBTransf2)
		shadowfbo.Unbind()

		// cam.LookAtval(float32(5*math.Cos(angle/2)), 5, float32(5*math.Sin(angle/2)), 0, 0, 0, 0, 1, 0)
		lamp.Move(3*math.Cos(float32(angle/2)), 3, 3*math.Sin(float32(angle/2)))
		shadowfbo.LookAt(lamp.X, lamp.Y, lamp.Z, 0, 0, 0)

		// ==Render== //
		gbuf.Bind(&cam)

		// normal rendering
		gbuf.Render(&cam, skydome, assman.Textures["skydome"], skydomeTransf)
		// gbuf.Render(&cam, boxModel, assman.Textures["brown"], boxRBTransf)
		// gbuf.Render(&cam, boxModel, assman.Textures["brown"], boxRBTransf2)
		// gbuf.Render(&cam, ground, assman.Textures["square"], groundTransf)

		// render lights
		gbuf.RenderLight(&cam, &lamp, shadowfbo.ShadowMat(), shadowfbo.ShadowMap(), 0.5, 0.9, 0.9)

		// aggregate
		gbuf.Aggregate()

		tonemap.PreRender()
		tonemap.Render(gbuf.AggregateFramebuffer.Out)
		tonemap.PostRender()

		// ==Maintenance== //
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
