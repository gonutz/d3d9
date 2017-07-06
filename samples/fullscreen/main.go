package main

import (
	"runtime"

	"github.com/AllenDang/gform"
	"github.com/AllenDang/w32"
	"github.com/gonutz/d3d9"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	gform.Init()
	form := gform.NewForm(nil)
	windowHandle := form.Handle()
	form.Show()
	form.OnClose().Bind(func(arg *gform.EventArg) {
		w32.DestroyWindow(form.Handle())
	})

	d3d, err := d3d9.Create(d3d9.SDK_VERSION)
	check(err)
	defer d3d.Release()

	// use the first found display mode for going full screen
	mode, err := d3d.EnumAdapterModes(0, d3d9.FMT_X8R8G8B8, 0)
	check(err)

	device, _, err := d3d.CreateDevice(
		d3d9.ADAPTER_DEFAULT,
		d3d9.DEVTYPE_HAL,
		d3d9.HWND(windowHandle),
		d3d9.CREATE_SOFTWARE_VERTEXPROCESSING,
		d3d9.PRESENT_PARAMETERS{
			Windowed:         0,
			HDeviceWindow:    d3d9.HWND(windowHandle),
			SwapEffect:       d3d9.SWAPEFFECT_DISCARD,
			BackBufferFormat: d3d9.FMT_X8R8G8B8,
			BackBufferWidth:  mode.Width,
			BackBufferHeight: mode.Height,
		},
	)
	check(err)
	defer device.Release()

	check(device.Clear(nil, d3d9.CLEAR_TARGET, d3d9.ColorRGB(255, 0, 255), 0, 1))
	check(device.Present(nil, nil, 0, nil))

	gform.RunMainLoop()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
