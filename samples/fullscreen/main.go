package main

import (
	"runtime"
	"syscall"

	"github.com/gonutz/d3d9"
	"github.com/gonutz/w32/v2"
)

func main() {
	// This sample starts by setting up a window using the Win32 API, skip down
	// to line 44 for the D3D9 part.
	runtime.LockOSThread()

	const className = "fullscreen_window_class"
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

	windowNamePtr, _ := syscall.UTF16PtrFromString("Fullscreen")
	window := w32.CreateWindow(
		classNamePtr,
		windowNamePtr,
		w32.WS_OVERLAPPEDWINDOW|w32.WS_VISIBLE,
		w32.CW_USEDEFAULT, w32.CW_USEDEFAULT, 640, 480,
		0, 0, 0, nil,
	)

	// Initialize Direct3D9.
	d3d, err := d3d9.Create(d3d9.SDK_VERSION)
	check(err)
	defer d3d.Release()

	// EnumAdapterModes lets us inspect all supported graphics modes for our
	// graphics card. We use the first one at index 0 to go full screen.
	mode, err := d3d.EnumAdapterModes(0, d3d9.FMT_X8R8G8B8, 0)
	check(err)

	device, _, err := d3d.CreateDevice(
		d3d9.ADAPTER_DEFAULT,
		d3d9.DEVTYPE_HAL,
		d3d9.HWND(window),
		// We use software emulation for this sample, as we do not render
		// anything. In reality you would check for hardware support by
		// inspecting the device capabilities and use hardware mode if
		// available.
		d3d9.CREATE_SOFTWARE_VERTEXPROCESSING,
		d3d9.PRESENT_PARAMETERS{
			Windowed:         0,
			HDeviceWindow:    d3d9.HWND(window),
			SwapEffect:       d3d9.SWAPEFFECT_DISCARD,
			BackBufferFormat: d3d9.FMT_X8R8G8B8,
			BackBufferWidth:  mode.Width,
			BackBufferHeight: mode.Height,
		},
	)
	check(err)
	defer device.Release()

	// Fill the screen with magic pink and call Present to show it.
	check(device.Clear(nil, d3d9.CLEAR_TARGET, d3d9.ColorRGB(255, 0, 255), 0, 1))
	check(device.Present(nil, nil, 0, nil))

	// This is the main message loop that runs until the window is closed.
	var msg w32.MSG
	for w32.GetMessage(&msg, 0, 0, 0) != 0 {
		w32.TranslateMessage(&msg)
		w32.DispatchMessage(&msg)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
