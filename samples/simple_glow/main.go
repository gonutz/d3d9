package main

//#include <windows.h>
import "C"
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
	// setup gform and create a window
	gform.Init()

	form := gform.NewForm(nil)
	windowHandle := unsafe.Pointer(form.Handle())
	form.Show()
	form.OnClose().Bind(func(*gform.EventArg) {
		w32.DestroyWindow(form.Handle())
	})

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

	// create a timer that ticks every 10ms and register a callback for it
	C.SetTimer(C.HWND(windowHandle), 1, 10, nil)
	red, dRed := 255, -1
	form.Bind(C.WM_TIMER, func(*gform.EventArg) {
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
		check(device.Present(nil, nil, nil, nil))
	})

	gform.RunMainLoop()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
