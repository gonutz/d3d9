package d3d9

/*
#include "d3d9wrapper.h"

HRESULT IDirect3DSwapChain9GetBackBuffer(
		IDirect3DSwapChain9* obj,
		UINT BackBuffer,
		D3DBACKBUFFER_TYPE Type,
		IDirect3DSurface9** ppBackBuffer) {
	return obj->lpVtbl->GetBackBuffer(obj, BackBuffer, Type, ppBackBuffer);
}

HRESULT IDirect3DSwapChain9GetDevice(
		IDirect3DSwapChain9* obj,
		IDirect3DDevice9** ppDevice) {
	return obj->lpVtbl->GetDevice(obj, ppDevice);
}

HRESULT IDirect3DSwapChain9GetDisplayMode(
		IDirect3DSwapChain9* obj,
		D3DDISPLAYMODE* pMode) {
	return obj->lpVtbl->GetDisplayMode(obj, pMode);
}

HRESULT IDirect3DSwapChain9GetFrontBufferData(
		IDirect3DSwapChain9* obj,
		IDirect3DSurface9* pDestSurface) {
	return obj->lpVtbl->GetFrontBufferData(obj, pDestSurface);
}

HRESULT IDirect3DSwapChain9GetPresentParameters(
		IDirect3DSwapChain9* obj,
		D3DPRESENT_PARAMETERS* pPresentationParameters) {
	return obj->lpVtbl->GetPresentParameters(obj, pPresentationParameters);
}

HRESULT IDirect3DSwapChain9GetRasterStatus(
		IDirect3DSwapChain9* obj,
		D3DRASTER_STATUS* pRasterStatus) {
	return obj->lpVtbl->GetRasterStatus(obj, pRasterStatus);
}

HRESULT IDirect3DSwapChain9Present(
		IDirect3DSwapChain9* obj,
		RECT* pSourceRect,
		RECT* pDestRect,
		HWND hDestWindowOverride,
		RGNDATA* pDirtyRegion,
		DWORD dwFlags) {
	return obj->lpVtbl->Present(obj, pSourceRect, pDestRect,
		hDestWindowOverride, pDirtyRegion, dwFlags);
}

void IDirect3DSwapChain9Release(IDirect3DSwapChain9* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"
import "unsafe"

// SwapChain and its methods are used to manipulate a swap chain.
type SwapChain struct {
	handle *C.IDirect3DSwapChain9
}

// Release has to be called when finished using the object to free its
// associated resources.
func (obj SwapChain) Release() {
	C.IDirect3DSwapChain9Release(obj.handle)
}

// GetBackBuffer retrieves a back buffer from the swap chain of the device.
// Call Release on the returned surface when finished using it.
func (obj SwapChain) GetBackBuffer(
	BackBuffer uint,
	Type BACKBUFFER_TYPE,
) (
	ppBackBuffer Surface,
	err Error,
) {
	var cppBackBuffer *C.IDirect3DSurface9
	err = toErr(C.IDirect3DSwapChain9GetBackBuffer(
		obj.handle,
		(C.UINT)(BackBuffer),
		(C.D3DBACKBUFFER_TYPE)(Type),
		&cppBackBuffer,
	))
	resource := Resource{
		(*C.IDirect3DResource9)(unsafe.Pointer(cppBackBuffer)),
	}
	ppBackBuffer = Surface{resource, cppBackBuffer}
	return
}

// GetDevice retrieves the device associated with the swap chain.
// Call Release on the returned device when finished using it.
func (obj SwapChain) GetDevice() (ppDevice Device, err Error) {
	var cppDevice *C.IDirect3DDevice9
	err = toErr(C.IDirect3DSwapChain9GetDevice(obj.handle, &cppDevice))
	ppDevice = Device{cppDevice}
	return
}

// GetDisplayMode retrieves the display mode's spatial resolution, color
// resolution, and refresh frequency.
func (obj SwapChain) GetDisplayMode() (pMode DISPLAYMODE, err Error) {
	var cpMode C.D3DDISPLAYMODE
	err = toErr(C.IDirect3DSwapChain9GetDisplayMode(obj.handle, &cpMode))
	pMode.fromC(&cpMode)
	return
}

// GetFrontBufferData generates a copy of the swapchain's front buffer and
// places that copy in a system memory buffer provided by the application.
// Call Release on the returned surface when finished using it.
func (obj SwapChain) GetFrontBufferData(pDestSurface Surface) (err Error) {
	err = toErr(C.IDirect3DSwapChain9GetFrontBufferData(
		obj.handle,
		pDestSurface.handle,
	))
	return
}

// GetPresentParameters retrieves the presentation parameters associated with a
// swap chain.
func (obj SwapChain) GetPresentParameters() (
	pPresentationParameters PRESENT_PARAMETERS,
	err Error,
) {
	var cpPresentationParameters C.D3DPRESENT_PARAMETERS
	err = toErr(C.IDirect3DSwapChain9GetPresentParameters(
		obj.handle,
		&cpPresentationParameters,
	))
	pPresentationParameters.fromC(&cpPresentationParameters)
	return
}

// GetRasterStatus returns information describing the raster of the monitor on
// which the swap chain is presented.
func (obj SwapChain) GetRasterStatus() (
	pRasterStatus RASTER_STATUS,
	err Error,
) {
	var cpRasterStatus C.D3DRASTER_STATUS
	err = toErr(C.IDirect3DSwapChain9GetRasterStatus(
		obj.handle,
		&cpRasterStatus,
	))
	pRasterStatus.fromC(&cpRasterStatus)
	return
}

// Present presents the contents of the next buffer in the sequence of back
// buffers owned by the swap chain.
func (obj SwapChain) Present(
	pSourceRect *RECT,
	pDestRect *RECT,
	hDestWindowOverride unsafe.Pointer,
	pDirtyRegion *RGNDATA,
	dwFlags uint32,
) (
	err Error,
) {
	if pSourceRect == nil && pDestRect == nil && pDirtyRegion == nil {
		err = toErr(C.IDirect3DSwapChain9Present(
			obj.handle,
			nil,
			nil,
			(C.HWND)(hDestWindowOverride),
			nil,
			C.DWORD(dwFlags),
		))
		return
	}

	var csourceRect C.RECT
	var cpSourceRect *C.RECT
	if pSourceRect != nil {
		csourceRect = pSourceRect.toC()
		cpSourceRect = &csourceRect
	}

	var cdestRect C.RECT
	var cpDestRect *C.RECT
	if pDestRect != nil {
		cdestRect = pDestRect.toC()
		cpDestRect = &cdestRect
	}

	var cdirtyRegion C.RGNDATA
	var cpDirtyRegion *C.RGNDATA
	if pDirtyRegion != nil {
		cdirtyRegion = pDirtyRegion.toC()
		cpDirtyRegion = &cdirtyRegion
	}

	err = toErr(C.IDirect3DSwapChain9Present(
		obj.handle,
		cpSourceRect,
		cpDestRect,
		(C.HWND)(hDestWindowOverride),
		cpDirtyRegion,
		C.DWORD(dwFlags),
	))
	return
}
