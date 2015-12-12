package d3d9

/*
#include <d3d9.h>

HRESULT IDirect3DSwapChain9GetBackBuffer(IDirect3DSwapChain9* obj, UINT BackBuffer, D3DBACKBUFFER_TYPE Type, IDirect3DSurface9** ppBackBuffer) {
	return obj->lpVtbl->GetBackBuffer(obj, BackBuffer, Type, ppBackBuffer);
}

HRESULT IDirect3DSwapChain9GetDevice(IDirect3DSwapChain9* obj, IDirect3DDevice9** ppDevice) {
	return obj->lpVtbl->GetDevice(obj, ppDevice);
}

HRESULT IDirect3DSwapChain9GetDisplayMode(IDirect3DSwapChain9* obj, D3DDISPLAYMODE* pMode) {
	return obj->lpVtbl->GetDisplayMode(obj, pMode);
}

HRESULT IDirect3DSwapChain9GetFrontBufferData(IDirect3DSwapChain9* obj, IDirect3DSurface9* pDestSurface) {
	return obj->lpVtbl->GetFrontBufferData(obj, pDestSurface);
}

HRESULT IDirect3DSwapChain9GetPresentParameters(IDirect3DSwapChain9* obj, D3DPRESENT_PARAMETERS* pPresentationParameters) {
	return obj->lpVtbl->GetPresentParameters(obj, pPresentationParameters);
}

HRESULT IDirect3DSwapChain9GetRasterStatus(IDirect3DSwapChain9* obj, D3DRASTER_STATUS* pRasterStatus) {
	return obj->lpVtbl->GetRasterStatus(obj, pRasterStatus);
}

HRESULT IDirect3DSwapChain9Present(IDirect3DSwapChain9* obj, RECT* pSourceRect, RECT* pDestRect, HWND hDestWindowOverride, RGNDATA* pDirtyRegion, DWORD dwFlags) {
	return obj->lpVtbl->Present(obj, pSourceRect, pDestRect, hDestWindowOverride, pDirtyRegion, dwFlags);
}

void IDirect3DSwapChain9Release(IDirect3DSwapChain9* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"
import "unsafe"

type SwapChain struct {
	handle *C.IDirect3DSwapChain9
}

func (obj SwapChain) Release() {
	C.IDirect3DSwapChain9Release(obj.handle)
}

func (obj SwapChain) GetBackBuffer(BackBuffer uint, Type BACKBUFFER_TYPE) (ppBackBuffer Surface, err error) {
	var c_ppBackBuffer *C.IDirect3DSurface9
	err = toErr(C.IDirect3DSwapChain9GetBackBuffer(obj.handle, (C.UINT)(BackBuffer), (C.D3DBACKBUFFER_TYPE)(Type), &c_ppBackBuffer))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(c_ppBackBuffer))}
	ppBackBuffer = Surface{resource, c_ppBackBuffer}
	return
}

func (obj SwapChain) GetDevice() (ppDevice Device, err error) {
	var c_ppDevice *C.IDirect3DDevice9
	err = toErr(C.IDirect3DSwapChain9GetDevice(obj.handle, &c_ppDevice))
	ppDevice = Device{c_ppDevice}
	return
}

func (obj SwapChain) GetDisplayMode() (pMode DISPLAYMODE, err error) {
	var c_pMode C.D3DDISPLAYMODE
	err = toErr(C.IDirect3DSwapChain9GetDisplayMode(obj.handle, &c_pMode))
	pMode.fromC(&c_pMode)
	return
}

func (obj SwapChain) GetFrontBufferData(pDestSurface Surface) (err error) {
	err = toErr(C.IDirect3DSwapChain9GetFrontBufferData(obj.handle, pDestSurface.handle))
	return
}

func (obj SwapChain) GetPresentParameters() (pPresentationParameters PRESENT_PARAMETERS, err error) {
	var c_pPresentationParameters C.D3DPRESENT_PARAMETERS
	err = toErr(C.IDirect3DSwapChain9GetPresentParameters(obj.handle, &c_pPresentationParameters))
	pPresentationParameters.fromC(&c_pPresentationParameters)
	return
}

func (obj SwapChain) GetRasterStatus() (pRasterStatus RASTER_STATUS, err error) {
	var c_pRasterStatus C.D3DRASTER_STATUS
	err = toErr(C.IDirect3DSwapChain9GetRasterStatus(obj.handle, &c_pRasterStatus))
	pRasterStatus.fromC(&c_pRasterStatus)
	return
}

func (obj SwapChain) Present(pSourceRect *RECT, pDestRect *RECT, hDestWindowOverride unsafe.Pointer, pDirtyRegion *RGNDATA, dwFlags uint32) (err error) {
	if pSourceRect == nil && pDestRect == nil && pDirtyRegion == nil {
		err = toErr(C.IDirect3DSwapChain9Present(obj.handle, nil, nil, (C.HWND)(hDestWindowOverride), nil, C.DWORD(dwFlags)))
		return
	}
	// TODO what if some are nil and some are not?
	var c_sourceRect, c_destRect C.RECT
	var c_dirtyRegion C.RGNDATA

	var c_pSourceRect *C.RECT = nil
	var c_pDestRect *C.RECT = nil
	var c_pDirtyRegion *C.RGNDATA = nil

	if pSourceRect != nil {
		c_sourceRect = pSourceRect.toC()
		c_pSourceRect = &c_sourceRect
	}
	if pDestRect != nil {
		c_destRect = pDestRect.toC()
		c_pDestRect = &c_destRect
	}
	if pDirtyRegion != nil {
		c_dirtyRegion = pDirtyRegion.toC()
		c_pDirtyRegion = &c_dirtyRegion
	}

	err = toErr(C.IDirect3DSwapChain9Present(obj.handle, c_pSourceRect, c_pDestRect, (C.HWND)(hDestWindowOverride), c_pDirtyRegion, C.DWORD(dwFlags)))
	return
}
