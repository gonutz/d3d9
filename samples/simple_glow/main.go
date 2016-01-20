package main

import (
	"github.com/gonutz/d3d9"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	// set up SDL2, init nothing as we will be initializing D3D ourselves
	check(sdl.Init(0))
	defer sdl.Quit()

	window, err := sdl.CreateWindow("D3D9 in Go", 100, 100, 640, 480, 0)
	check(err)
	defer window.Destroy()

	// this gets the Windows window handle for the native window
	info, err := window.GetWMInfo()
	check(err)
	windowHandle := info.GetWindowsInfo().Window

	// set up Direct3D9
	check(d3d9.Init())
	defer d3d9.Close()

	d3d, err := d3d9.Create(d3d9.SDK_VERSION)
	check(err)
	defer d3d.Release()

	device, _, err := d3d.CreateDevice(d3d9.ADAPTER_DEFAULT, d3d9.DEVTYPE_HAL,
		windowHandle, d3d9.CREATE_HARDWARE_VERTEXPROCESSING,
		d3d9.PRESENT_PARAMETERS{
			Windowed:      true,
			SwapEffect:    d3d9.SWAPEFFECT_DISCARD,
			HDeviceWindow: windowHandle,
		})
	check(err)
	defer device.Release()

	// run the main event loop
	red, dRed := 255, -1
	running := true
	for running {
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch e.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}

		// simply clear the screen
		check(device.Clear(nil, d3d9.CLEAR_TARGET, d3d9.ColorRGB(uint(red), 92, 128), 0, 0))
		if red == 0 {
			dRed = +1
		}
		if red == 255 {
			dRed = -1
		}
		red += dRed
		check(device.Present(nil, nil, nil, nil))
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
