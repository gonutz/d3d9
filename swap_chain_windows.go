package d3d9

import (
	"syscall"
	"unsafe"
)

// SwapChain and its methods are used to manipulate a swap chain.
type SwapChain struct {
	vtbl *swapChainVtbl
}

type swapChainVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	Present              uintptr
	GetFrontBufferData   uintptr
	GetBackBuffer        uintptr
	GetRasterStatus      uintptr
	GetDisplayMode       uintptr
	GetDevice            uintptr
	GetPresentParameters uintptr
}

// AddRef increments the reference count for an interface on an object. This
// method should be called for every new copy of a pointer to an interface on an
// object.
func (obj *SwapChain) AddRef() uint32 {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.AddRef,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint32(ret)
}

// Release has to be called when finished using the object to free its
// associated resources.
func (obj *SwapChain) Release() uint32 {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint32(ret)
}

// Present presents the contents of the next buffer in the sequence of back
// buffers owned by the swap chain.
func (obj *SwapChain) Present(
	sourceRect *RECT,
	destRect *RECT,
	destWindowOverride HWND,
	dirtyRegion *RGNDATA,
	flags uint32,
) Error {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.Present,
		6,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(sourceRect)),
		uintptr(unsafe.Pointer(destRect)),
		uintptr(destWindowOverride),
		uintptr(unsafe.Pointer(dirtyRegion)),
		uintptr(flags),
	)
	return toErr(ret)
}

// GetFrontBufferData generates a copy of the swapchain's front buffer and
// places that copy in a system memory buffer provided by the application.
// Call Release on the returned surface when finished using it.
func (obj *SwapChain) GetFrontBufferData(destSurface *Surface) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetFrontBufferData,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(destSurface)),
		0,
	)
	return toErr(ret)
}

// GetBackBuffer retrieves a back buffer from the swap chain of the device.
// Call Release on the returned surface when finished using it.
func (obj *SwapChain) GetBackBuffer(
	backBuffer uint,
	typ BACKBUFFER_TYPE,
) (buffer *Surface, err Error) {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.GetBackBuffer,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(backBuffer),
		uintptr(typ),
		uintptr(unsafe.Pointer(&buffer)),
		0,
		0,
	)
	err = toErr(ret)
	return
}

// GetRasterStatus returns information describing the raster of the monitor on
// which the swap chain is presented.
func (obj *SwapChain) GetRasterStatus() (status RASTER_STATUS, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetRasterStatus,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&status)),
		0,
	)
	err = toErr(ret)
	return
}

// GetDisplayMode retrieves the display mode's spatial resolution, color
// resolution, and refresh frequency.
func (obj *SwapChain) GetDisplayMode() (mode DISPLAYMODE, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetDisplayMode,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&mode)),
		0,
	)
	err = toErr(ret)
	return
}

// GetDevice retrieves the device associated with the swap chain.
// Call Release on the returned device when finished using it.
func (obj *SwapChain) GetDevice() (device *Device, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetDevice,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&device)),
		0,
	)
	err = toErr(ret)
	return
}

// GetPresentParameters retrieves the presentation parameters associated with a
// swap chain.
func (obj *SwapChain) GetPresentParameters() (params PRESENT_PARAMETERS, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetPresentParameters,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&params)),
		0,
	)
	err = toErr(ret)
	return
}
