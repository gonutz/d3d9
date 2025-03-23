package main

import (
	"runtime"
	"syscall"

	"github.com/gonutz/d3d9"
	"github.com/gonutz/w32/v2"
)

func main() {
	runtime.LockOSThread()

	const className = "glowing_window_class"
	classNamePtr, _ := syscall.UTF16PtrFromString(className)
	w32.RegisterClassEx(&w32.WNDCLASSEX{
		Cursor: w32.LoadCursor(0, w32.MakeIntResource(w32.IDC_ARROW)),
		WndProc: syscall.NewCallback(func(window w32.HWND, msg uint32, w, l uintptr) uintptr {
			switch msg {
			case w32.WM_KEYDOWN:
				if w == w32.VK_ESCAPE {
					w32.SendMessage(window, w32.WM_CLOSE, 0, 0)
				}
				return 0
			case w32.WM_DESTROY:
				w32.PostQuitMessage(0)
				return 0
			default:
				return w32.DefWindowProc(window, msg, w, l)
			}
		}),
		ClassName: classNamePtr,
	})

	windowNamePtr, _ := syscall.UTF16PtrFromString("Glow")
	windowHandle := w32.CreateWindow(
		classNamePtr,
		windowNamePtr,
		w32.WS_OVERLAPPEDWINDOW|w32.WS_VISIBLE,
		w32.CW_USEDEFAULT, w32.CW_USEDEFAULT, 640, 480,
		0, 0, 0, nil,
	)

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

	// Create a timer that ticks every 10ms and register a callback for it.
	w32.SetTimer(windowHandle, 1, 10, 0)
	red, dRed := 255, -1

	// This is the main message loop that runs until the window is closed.
	var msg w32.MSG
	for w32.GetMessage(&msg, 0, 0, 0) != 0 {
		w32.TranslateMessage(&msg)

		// Clear the screen to the current color.
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

		w32.DispatchMessage(&msg)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
