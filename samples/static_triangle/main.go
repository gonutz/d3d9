package main

//#include <windows.h>
import "C"
import (
	"github.com/AllenDang/gform"
	"github.com/AllenDang/w32"
	"github.com/gonutz/d3d9"
	"unsafe"
)

func main() {
	// setup gform and create a window
	gform.Init()

	form := gform.NewForm(nil)
	windowHandle := unsafe.Pointer(form.Handle())
	form.Show()
	form.OnClose().Bind(func(arg *gform.EventArg) {
		w32.DestroyWindow(form.Handle())
	})

	// set up Direct3D9
	check(d3d9.Init())
	defer d3d9.Close()

	d3d, err := d3d9.Create(d3d9.SDK_VERSION)
	check(err)
	defer d3d.Release()

	device, _, err := d3d.CreateDevice(
		d3d9.ADAPTER_DEFAULT,
		d3d9.DEVTYPE_HAL,
		windowHandle,
		d3d9.CREATE_HARDWARE_VERTEXPROCESSING,
		d3d9.PRESENT_PARAMETERS{
			Windowed:      true,
			SwapEffect:    d3d9.SWAPEFFECT_DISCARD,
			HDeviceWindow: windowHandle,
		},
	)
	check(err)
	defer device.Release()

	check(device.SetRenderState(d3d9.RS_CULLMODE, uint32(d3d9.CULL_NONE)))

	vs, err := device.CreateVertexShaderFromBytes(vso)
	check(err)
	defer vs.Release()
	check(device.SetVertexShader(vs))

	ps, err := device.CreatePixelShaderFromBytes(pso)
	check(err)
	defer ps.Release()
	check(device.SetPixelShader(ps))

	vertices := []float32{
		-0.5, -0.5,
		0, 0.5,
		0.5, -0.5,
	}
	vb, err := device.CreateVertexBuffer(uint(len(vertices)*4),
		d3d9.USAGE_WRITEONLY, 0, d3d9.POOL_DEFAULT, nil)
	check(err)
	defer vb.Release()
	vb.LockedSetFloats(0, 0, vertices)
	check(device.SetStreamSource(0, vb, 0, 2*4))

	decl, err := device.CreateVertexDeclaration([]d3d9.VERTEXELEMENT{
		{0, 0, d3d9.DECLTYPE_FLOAT2, d3d9.DECLMETHOD_DEFAULT, d3d9.DECLUSAGE_POSITION, 0},
		d3d9.DeclEnd(),
	})
	check(err)
	defer decl.Release()
	check(device.SetVertexDeclaration(decl))

	// create a timer that ticks every 100ms and register a callback for it
	C.SetTimer(C.HWND(windowHandle), 1, 100, nil)
	form.Bind(C.WM_TIMER, func(*gform.EventArg) {
		check(device.Clear(nil, d3d9.CLEAR_TARGET, 0, 0, 0))
		check(device.BeginScene())
		check(device.DrawPrimitive(d3d9.PT_TRIANGLELIST, 0, 1))
		check(device.EndScene())
		check(device.Present(nil, nil, nil, nil))
	})

	gform.RunMainLoop()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
