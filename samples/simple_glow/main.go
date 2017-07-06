package main

import (
	"runtime"

	"github.com/AllenDang/gform"
	"github.com/gonutz/d3d9"
	"github.com/gonutz/w32"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	// setup gform and create a window
	gform.Init()
	form := gform.NewForm(nil)
	windowHandle := form.Handle()
	form.Show()
	form.OnClose().Bind(func(*gform.EventArg) {
		w32.DestroyWindow(w32.HWND(form.Handle()))
	})

	d3d, err := d3d9.Create(d3d9.SDK_VERSION)
	check(err)
	defer d3d.Release()

	device, _, err := d3d.CreateDevice(
		d3d9.ADAPTER_DEFAULT,
		d3d9.DEVTYPE_HAL,
		d3d9.HWND(windowHandle),
		d3d9.CREATE_HARDWARE_VERTEXPROCESSING,
		d3d9.PRESENT_PARAMETERS{
			Windowed:      1,
			SwapEffect:    d3d9.SWAPEFFECT_DISCARD,
			HDeviceWindow: d3d9.HWND(windowHandle),
		},
	)
	check(err)
	defer device.Release()

	// create a timer that ticks every 10ms and register a callback for it
	w32.SetTimer(w32.HWND(windowHandle), 1, 10, 0)
	red, dRed := 255, -1
	form.Bind(w32.WM_TIMER, func(*gform.EventArg) {
		// clear the screen to the current color
		check(device.Clear(
			nil,
			d3d9.CLEAR_TARGET,
			d3d9.ColorRGB(uint8(red), 92, 128),
			0,
			0,
		))
		if red == 0 {
			dRed = +1
		}
		if red == 255 {
			dRed = -1
		}
		red += dRed
		check(device.Present(nil, nil, 0, nil))
	})

	gform.RunMainLoop()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
