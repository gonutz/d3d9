package d3d9

/*
#include <d3d9.h>

HRESULT IDirect3DDevice9ExCheckDeviceState(IDirect3DDevice9Ex* obj, HWND hWindow) {
	return obj->lpVtbl->CheckDeviceState(obj, hWindow);
}

HRESULT IDirect3DDevice9ExComposeRects(IDirect3DDevice9Ex* obj, IDirect3DSurface9* pSource, IDirect3DSurface9* pDestination, IDirect3DVertexBuffer9* pSrcRectDescriptors, UINT NumRects, IDirect3DVertexBuffer9* pDstRectDescriptors, D3DCOMPOSERECTSOP Operation, INT XOffset, INT YOffset) {
	return obj->lpVtbl->ComposeRects(obj, pSource, pDestination, pSrcRectDescriptors, NumRects, pDstRectDescriptors, Operation, XOffset, YOffset);
}

HRESULT IDirect3DDevice9ExCreateDepthStencilSurfaceEx(IDirect3DDevice9Ex* obj, UINT Width, UINT Height, D3DFORMAT Format, D3DMULTISAMPLE_TYPE MultiSample, DWORD MultisampleQuality, BOOL Discard, IDirect3DSurface9** ppSurface, HANDLE* pSharedHandle, DWORD Usage) {
	return obj->lpVtbl->CreateDepthStencilSurfaceEx(obj, Width, Height, Format, MultiSample, MultisampleQuality, Discard, ppSurface, pSharedHandle, Usage);
}

HRESULT IDirect3DDevice9ExCreateOffscreenPlainSurfaceEx(IDirect3DDevice9Ex* obj, UINT Width, UINT Height, D3DFORMAT Format, D3DPOOL Pool, IDirect3DSurface9** ppSurface, HANDLE* pSharedHandle, DWORD Usage) {
	return obj->lpVtbl->CreateOffscreenPlainSurfaceEx(obj, Width, Height, Format, Pool, ppSurface, pSharedHandle, Usage);
}

HRESULT IDirect3DDevice9ExCreateRenderTargetEx(IDirect3DDevice9Ex* obj, UINT Width, UINT Height, D3DFORMAT Format, D3DMULTISAMPLE_TYPE MultiSample, DWORD MultisampleQuality, BOOL Lockable, IDirect3DSurface9** ppSurface, HANDLE* pSharedHandle, DWORD Usage) {
	return obj->lpVtbl->CreateRenderTargetEx(obj, Width, Height, Format, MultiSample, MultisampleQuality, Lockable, ppSurface, pSharedHandle, Usage);
}

HRESULT IDirect3DDevice9ExGetDisplayModeEx(IDirect3DDevice9Ex* obj, UINT iSwapChain, D3DDISPLAYMODEEX* pMode, D3DDISPLAYROTATION* pRotation) {
	return obj->lpVtbl->GetDisplayModeEx(obj, iSwapChain, pMode, pRotation);
}

HRESULT IDirect3DDevice9ExGetGPUThreadPriority(IDirect3DDevice9Ex* obj, INT* pPriority) {
	return obj->lpVtbl->GetGPUThreadPriority(obj, pPriority);
}

HRESULT IDirect3DDevice9ExGetMaximumFrameLatency(IDirect3DDevice9Ex* obj, UINT* pMaxLatency) {
	return obj->lpVtbl->GetMaximumFrameLatency(obj, pMaxLatency);
}

HRESULT IDirect3DDevice9ExPresentEx(IDirect3DDevice9Ex* obj, RECT* pSourceRect, RECT* pDestRect, HWND hDestWindowOverride, RGNDATA* pDirtyRegion, DWORD dwFlags) {
	return obj->lpVtbl->PresentEx(obj, pSourceRect, pDestRect, hDestWindowOverride, pDirtyRegion, dwFlags);
}

HRESULT IDirect3DDevice9ExResetEx(IDirect3DDevice9Ex* obj, D3DPRESENT_PARAMETERS* pPresentationParameters, D3DDISPLAYMODEEX* pFullscreenDisplayMode) {
	return obj->lpVtbl->ResetEx(obj, pPresentationParameters, pFullscreenDisplayMode);
}

HRESULT IDirect3DDevice9ExSetGPUThreadPriority(IDirect3DDevice9Ex* obj, INT pPriority) {
	return obj->lpVtbl->SetGPUThreadPriority(obj, pPriority);
}

HRESULT IDirect3DDevice9ExSetMaximumFrameLatency(IDirect3DDevice9Ex* obj, UINT pMaxLatency) {
	return obj->lpVtbl->SetMaximumFrameLatency(obj, pMaxLatency);
}

HRESULT IDirect3DDevice9ExTestCooperativeLevel(IDirect3DDevice9Ex* obj) {
	return obj->lpVtbl->TestCooperativeLevel(obj);
}

HRESULT IDirect3DDevice9ExWaitForVBlank(IDirect3DDevice9Ex* obj, UINT SwapChainIndex) {
	return obj->lpVtbl->WaitForVBlank(obj, SwapChainIndex);
}

void IDirect3DDevice9ExRelease(IDirect3DDevice9Ex* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"
import "unsafe"

type DeviceEx struct {
	Device
	handle *C.IDirect3DDevice9Ex
}

func (obj DeviceEx) Release() {
	C.IDirect3DDevice9ExRelease(obj.handle)
}

func (obj DeviceEx) CheckDeviceState(hWindow unsafe.Pointer) (err error) {
	err = toErr(C.IDirect3DDevice9ExCheckDeviceState(obj.handle, (C.HWND)(hWindow)))
	return
}

func (obj DeviceEx) ComposeRects(pSource Surface, pDestination Surface, pSrcRectDescriptors VertexBuffer, NumRects uint, pDstRectDescriptors VertexBuffer, Operation COMPOSERECTSOP, XOffset int, YOffset int) (err error) {
	err = toErr(C.IDirect3DDevice9ExComposeRects(obj.handle, pSource.handle, pDestination.handle, pSrcRectDescriptors.handle, (C.UINT)(NumRects), pDstRectDescriptors.handle, (C.D3DCOMPOSERECTSOP)(Operation), (C.INT)(XOffset), (C.INT)(YOffset)))
	return
}

func (obj DeviceEx) CreateDepthStencilSurfaceEx(Width uint, Height uint, Format FORMAT, MultiSample MULTISAMPLE_TYPE, MultisampleQuality uint32, Discard bool, pSharedHandle unsafe.Pointer, Usage uint32) (ppSurface Surface, err error) {
	var c_ppSurface *C.IDirect3DSurface9
	err = toErr(C.IDirect3DDevice9ExCreateDepthStencilSurfaceEx(obj.handle, (C.UINT)(Width), (C.UINT)(Height), (C.D3DFORMAT)(Format), (C.D3DMULTISAMPLE_TYPE)(MultiSample), (C.DWORD)(MultisampleQuality), toBOOL(Discard), &c_ppSurface, (*C.HANDLE)(pSharedHandle), (C.DWORD)(Usage)))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(c_ppSurface))}
	ppSurface = Surface{resource, c_ppSurface}
	return
}

func (obj DeviceEx) CreateOffscreenPlainSurfaceEx(Width uint, Height uint, Format FORMAT, Pool POOL, pSharedHandle unsafe.Pointer, Usage uint32) (ppSurface Surface, err error) {
	var c_ppSurface *C.IDirect3DSurface9
	err = toErr(C.IDirect3DDevice9ExCreateOffscreenPlainSurfaceEx(obj.handle, (C.UINT)(Width), (C.UINT)(Height), (C.D3DFORMAT)(Format), (C.D3DPOOL)(Pool), &c_ppSurface, (*C.HANDLE)(pSharedHandle), (C.DWORD)(Usage)))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(c_ppSurface))}
	ppSurface = Surface{resource, c_ppSurface}
	return
}

func (obj DeviceEx) CreateRenderTargetEx(Width uint, Height uint, Format FORMAT, MultiSample MULTISAMPLE_TYPE, MultisampleQuality uint32, Lockable bool, pSharedHandle unsafe.Pointer, Usage uint32) (ppSurface Surface, err error) {
	var c_ppSurface *C.IDirect3DSurface9
	err = toErr(C.IDirect3DDevice9ExCreateRenderTargetEx(obj.handle, (C.UINT)(Width), (C.UINT)(Height), (C.D3DFORMAT)(Format), (C.D3DMULTISAMPLE_TYPE)(MultiSample), (C.DWORD)(MultisampleQuality), toBOOL(Lockable), &c_ppSurface, (*C.HANDLE)(pSharedHandle), (C.DWORD)(Usage)))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(c_ppSurface))}
	ppSurface = Surface{resource, c_ppSurface}
	return
}

func (obj DeviceEx) GetDisplayModeEx(iSwapChain uint) (pMode DISPLAYMODEEX, pRotation DISPLAYROTATION, err error) {
	var c_pMode C.D3DDISPLAYMODEEX
	var c_pRotation C.D3DDISPLAYROTATION
	err = toErr(C.IDirect3DDevice9ExGetDisplayModeEx(obj.handle, (C.UINT)(iSwapChain), &c_pMode, &c_pRotation))
	pMode.fromC(&c_pMode)
	pRotation = (DISPLAYROTATION)(c_pRotation)
	return
}

func (obj DeviceEx) GetGPUThreadPriority() (pPriority int, err error) {
	var c_pPriority C.INT
	err = toErr(C.IDirect3DDevice9ExGetGPUThreadPriority(obj.handle, &c_pPriority))
	pPriority = (int)(c_pPriority)
	return
}

func (obj DeviceEx) GetMaximumFrameLatency() (pMaxLatency uint, err error) {
	var c_pMaxLatency C.UINT
	err = toErr(C.IDirect3DDevice9ExGetMaximumFrameLatency(obj.handle, &c_pMaxLatency))
	pMaxLatency = (uint)(c_pMaxLatency)
	return
}

func (obj DeviceEx) PresentEx(pSourceRect RECT, pDestRect RECT, hDestWindowOverride unsafe.Pointer, pDirtyRegion RGNDATA, dwFlags uint32) (err error) {
	c_pSourceRect := pSourceRect.toC()
	c_pDestRect := pDestRect.toC()
	c_pDirtyRegion := pDirtyRegion.toC()
	err = toErr(C.IDirect3DDevice9ExPresentEx(obj.handle, &c_pSourceRect, &c_pDestRect, (C.HWND)(hDestWindowOverride), &c_pDirtyRegion, (C.DWORD)(dwFlags)))
	return
}

func (obj DeviceEx) ResetEx(inpPresentationParameters PRESENT_PARAMETERS, inpFullscreenDisplayMode DISPLAYMODEEX) (outpPresentationParameters PRESENT_PARAMETERS, outpFullscreenDisplayMode DISPLAYMODEEX, err error) {
	c_pPresentationParameters := inpPresentationParameters.toC()
	c_pFullscreenDisplayMode := inpFullscreenDisplayMode.toC()
	err = toErr(C.IDirect3DDevice9ExResetEx(obj.handle, &c_pPresentationParameters, &c_pFullscreenDisplayMode))
	outpPresentationParameters.fromC(&c_pPresentationParameters)
	outpFullscreenDisplayMode.fromC(&c_pFullscreenDisplayMode)
	return
}

func (obj DeviceEx) SetGPUThreadPriority(pPriority int) (err error) {
	err = toErr(C.IDirect3DDevice9ExSetGPUThreadPriority(obj.handle, (C.INT)(pPriority)))
	return
}

func (obj DeviceEx) SetMaximumFrameLatency(pMaxLatency uint) (err error) {
	err = toErr(C.IDirect3DDevice9ExSetMaximumFrameLatency(obj.handle, (C.UINT)(pMaxLatency)))
	return
}

func (obj DeviceEx) TestCooperativeLevel() (err error) {
	err = toErr(C.IDirect3DDevice9ExTestCooperativeLevel(obj.handle))
	return
}

func (obj DeviceEx) WaitForVBlank(SwapChainIndex uint) (err error) {
	err = toErr(C.IDirect3DDevice9ExWaitForVBlank(obj.handle, (C.UINT)(SwapChainIndex)))
	return
}
