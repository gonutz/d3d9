package main

import (
	"runtime"
	"syscall"

	"github.com/gonutz/d3d9"
	"github.com/gonutz/w32"
)

func main() {
	runtime.LockOSThread()

	const className = "fullscreen_window_class"
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
		ClassName: syscall.StringToUTF16Ptr(className),
	})

	window := w32.CreateWindow(
		syscall.StringToUTF16Ptr(className),
		syscall.StringToUTF16Ptr("Fullscreen"),
		w32.WS_OVERLAPPEDWINDOW|w32.WS_VISIBLE,
		w32.CW_USEDEFAULT, w32.CW_USEDEFAULT, 640, 480,
		0, 0, 0, nil,
	)

	d3d, err := d3d9.Create(d3d9.SDK_VERSION)
	check(err)
	defer d3d.Release()

	// use the first found display mode for going full screen
	mode, err := d3d.EnumAdapterModes(0, d3d9.FMT_X8R8G8B8, 0)
	check(err)

	device, _, err := d3d.CreateDevice(
		d3d9.ADAPTER_DEFAULT,
		d3d9.DEVTYPE_HAL,
		d3d9.HWND(window),
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

	check(device.Clear(nil, d3d9.CLEAR_TARGET, d3d9.ColorRGB(255, 0, 255), 0, 1))
	check(device.Present(nil, nil, 0, nil))

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
