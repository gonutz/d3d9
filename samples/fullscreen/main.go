package main

import (
	"github.com/AllenDang/gform"
	"github.com/AllenDang/w32"
	"github.com/gonutz/d3d9"
	"runtime"
	"unsafe"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	gform.Init()

	form := gform.NewForm(nil)
	windowHandle := unsafe.Pointer(form.Handle())
	form.Show()
	form.OnClose().Bind(func(arg *gform.EventArg) {
		w32.DestroyWindow(form.Handle())
	})

	check(d3d9.Init())
	defer d3d9.Close()

	d3d, err := d3d9.Create(d3d9.SDK_VERSION)
	check(err)
	defer d3d.Release()

	device, _, err := d3d.CreateDevice(
		d3d9.ADAPTER_DEFAULT,
		d3d9.DEVTYPE_HAL,
		windowHandle,
		d3d9.CREATE_SOFTWARE_VERTEXPROCESSING,
		d3d9.PRESENT_PARAMETERS{
			Windowed:         false,
			HDeviceWindow:    windowHandle,
			SwapEffect:       d3d9.SWAPEFFECT_DISCARD,
			BackBufferFormat: d3d9.FMT_X8R8G8B8,
			BackBufferWidth:  640,
			BackBufferHeight: 480,
		},
	)
	check(err)
	defer device.Release()

	check(device.Clear(nil, d3d9.CLEAR_TARGET, d3d9.ColorRGB(255, 0, 255), 0, 1))
	check(device.Present(nil, nil, nil, nil))

	gform.RunMainLoop()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
