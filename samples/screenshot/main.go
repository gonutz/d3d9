package main

import (
	"image"
	"image/png"
	"os"
	"unsafe"

	"github.com/gonutz/d3d9"
)

func main() {
	d3d, err := d3d9.Create(d3d9.SDK_VERSION)
	check(err)
	defer d3d.Release()

	mode, err := d3d.GetAdapterDisplayMode(d3d9.ADAPTER_DEFAULT)
	check(err)
	if mode.Format != d3d9.FMT_X8R8G8B8 && mode.Format != d3d9.FMT_A8R8G8B8 {
		panic("we assume the mode is ARGB.")
	}

	device, _, err := d3d.CreateDevice(
		d3d9.ADAPTER_DEFAULT,
		d3d9.DEVTYPE_HAL,
		0,
		d3d9.CREATE_SOFTWARE_VERTEXPROCESSING,
		d3d9.PRESENT_PARAMETERS{
			Windowed:         1,
			BackBufferCount:  1,
			BackBufferWidth:  mode.Width,
			BackBufferHeight: mode.Height,
			SwapEffect:       d3d9.SWAPEFFECT_DISCARD,
		},
	)
	check(err)
	defer device.Release()

	surface, err := device.CreateOffscreenPlainSurface(
		uint(mode.Width),
		uint(mode.Height),
		d3d9.FMT_A8R8G8B8,
		d3d9.POOL_SYSTEMMEM,
		0,
	)
	check(err)
	defer surface.Release()

	check(device.GetFrontBufferData(0, surface))

	r, err := surface.LockRect(nil, 0)
	check(err)
	defer surface.UnlockRect()

	// Pitch says how many bytes there are in one line of pixels of the surface.
	// This may include padding bytes that would need to be handled correctly
	// if not 0.
	if r.Pitch != int32(mode.Width*4) {
		panic("We assume that there is no padding in the surface data.")
	}

	// Create an image of the same size as the surface and copy it.
	img := image.NewRGBA(image.Rect(0, 0, int(mode.Width), int(mode.Height)))
	for i := range img.Pix {
		// Do the magic to use r.PBits as a byte pointer.
		img.Pix[i] = *((*byte)(unsafe.Pointer(r.PBits + uintptr(i))))
	}
	// We want RGBA for our Go image but have ARGB in our surface. Swap some
	// bytes to make it fit.
	for i := 0; i < len(img.Pix); i += 4 {
		img.Pix[i+0], img.Pix[i+2] = img.Pix[i+2], img.Pix[i+0]
	}

	f, err := os.Create("screen.png")
	check(err)
	defer f.Close()
	check(png.Encode(f, img))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
