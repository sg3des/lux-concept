package env

import "runtime"

//Initialize is enter point for enable lux engine
func Initialize() (*Render, error) {
	runtime.LockOSThread()

	render, err := initWindow()
	if err != nil {
		return nil, err
	}
	initAssets() //need error handling
	// initLights()
	// initPhysics()
	return render, nil
}
