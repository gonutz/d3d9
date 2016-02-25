package main

import (
	"github.com/gonutz/d3d9"
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	// set up SDL2, init nothing as we will be initializing D3D ourselves
	check(sdl.Init(0))
	defer sdl.Quit()

	window, err := sdl.CreateWindow("D3D9 in Go", 100, 100, 640, 640, 0)
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

	vertices := []float32{
		-0.5, -0.5,
		-0.5, 0.5,
		0.5, -0.5,
		0.5, 0.5,
	}
	vb, err := device.CreateVertexBuffer(
		uint(len(vertices)*4),
		d3d9.USAGE_WRITEONLY,
		0,
		d3d9.POOL_DEFAULT,
		nil,
	)
	check(err)
	defer vb.Release()
	vbMem, err := vb.Lock(0, 0, d3d9.LOCK_DISCARD)
	check(err)
	vbMem.SetFloat32s(0, vertices)
	check(vb.Unlock())

	check(device.SetStreamSource(0, vb, 0, 2*4))

	decl, err := device.CreateVertexDeclaration([]d3d9.VERTEXELEMENT{
		{0, 0, d3d9.DECLTYPE_FLOAT2, d3d9.DECLMETHOD_DEFAULT, d3d9.DECLUSAGE_POSITION, 0},
		d3d9.DeclEnd(),
	})
	check(err)
	defer decl.Release()
	check(device.SetVertexDeclaration(decl))

	vs, err := device.CreateVertexShaderFromBytes(vso)
	check(err)
	defer vs.Release()
	check(device.SetVertexShader(vs))

	ps, err := device.CreatePixelShaderFromBytes(pso)
	check(err)
	defer ps.Release()
	check(device.SetPixelShader(ps))

	// run the main event loop
	rotation := float32(0)
	running := true
	for running {
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch e.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}

		// create a 4x4 model-view-projection matrix that rotates the
		// rectangle about the Z-axis
		s, c := math.Sincos(float64(rotation))
		sin, cos := float32(s), float32(c)
		mvp := [16]float32{
			cos, sin, 0, 0,
			-sin, cos, 0, 0,
			0, 0, 1, 0,
			0, 0, 0, 1,
		}
		rotation += 0.1
		check(device.SetVertexShaderConstantF(0, mvp[:]))
		check(device.Clear(nil, d3d9.CLEAR_TARGET, 255, 1, 0))
		check(device.BeginScene())
		check(device.DrawPrimitive(d3d9.PT_TRIANGLESTRIP, 0, 2))
		check(device.EndScene())
		check(device.Present(nil, nil, nil, nil))
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
