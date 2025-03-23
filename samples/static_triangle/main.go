package main

import (
	"runtime"
	"syscall"

	"github.com/gonutz/d3d9"
	"github.com/gonutz/dxc"
	"github.com/gonutz/w32/v2"
)

const vertexShaderCode = `
struct input {
	float4 position : POSITION;
};

struct output {
	float4 position : POSITION;
};

void main(in input IN, out output OUT) {
	OUT.position = IN.position;
}
`

const pixelShaderCode = `
struct input {};

struct output {
	float4 color : COLOR0;
};

void main(in input IN, out output OUT) {
	OUT.color = float4(1, 0, 1, 1);
}
`

func main() {
	runtime.LockOSThread()

	vertexShaderObject, err := dxc.Compile(
		[]byte(vertexShaderCode),
		"main",
		"vs_2_0",
		dxc.WARNINGS_ARE_ERRORS,
		0,
	)
	check(err)

	pixelShaderObject, err := dxc.Compile(
		[]byte(pixelShaderCode),
		"main",
		"ps_2_0",
		dxc.WARNINGS_ARE_ERRORS,
		0,
	)
	check(err)

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

	windowNamePtr, _ := syscall.UTF16PtrFromString("Static Triangle")
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

	check(device.SetRenderState(d3d9.RS_CULLMODE, uint32(d3d9.CULL_NONE)))

	vs, err := device.CreateVertexShaderFromBytes(vertexShaderObject)
	check(err)
	defer vs.Release()
	check(device.SetVertexShader(vs))

	ps, err := device.CreatePixelShaderFromBytes(pixelShaderObject)
	check(err)
	defer ps.Release()
	check(device.SetPixelShader(ps))

	vertices := []float32{
		-0.5, -0.5,
		0, 0.5,
		0.5, -0.5,
	}
	vb, err := device.CreateVertexBuffer(uint(len(vertices)*4),
		d3d9.USAGE_WRITEONLY, 0, d3d9.POOL_DEFAULT, 0)
	check(err)
	defer vb.Release()
	data, err := vb.Lock(0, 0, d3d9.LOCK_DISCARD)
	check(err)
	data.SetFloat32s(0, vertices)
	check(vb.Unlock())
	check(device.SetStreamSource(0, vb, 0, 2*4))

	decl, err := device.CreateVertexDeclaration([]d3d9.VERTEXELEMENT{
		{Type: d3d9.DECLTYPE_FLOAT2, Usage: d3d9.DECLUSAGE_POSITION},
		d3d9.DeclEnd(),
	})
	check(err)
	defer decl.Release()
	check(device.SetVertexDeclaration(decl))

	// Create a timer that ticks every 100ms and register a callback for it.
	w32.SetTimer(windowHandle, 1, 100, 0)
	var msg w32.MSG
	for w32.GetMessage(&msg, 0, 0, 0) != 0 {
		w32.TranslateMessage(&msg)
		check(device.Clear(nil, d3d9.CLEAR_TARGET, 0, 0, 0))
		check(device.BeginScene())
		check(device.DrawPrimitive(d3d9.PT_TRIANGLELIST, 0, 1))
		check(device.EndScene())
		check(device.Present(nil, nil, 0, nil))
		w32.DispatchMessage(&msg)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
