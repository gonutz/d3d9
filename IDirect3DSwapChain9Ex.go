package d3d9

/*
#include <d3d9.h>

HRESULT IDirect3DSwapChain9ExGetDisplayModeEx(IDirect3DSwapChain9Ex* obj, D3DDISPLAYMODEEX* pMode, D3DDISPLAYROTATION* pRotation) {
	return obj->lpVtbl->GetDisplayModeEx(obj, pMode, pRotation);
}

HRESULT IDirect3DSwapChain9ExGetLastPresentCount(IDirect3DSwapChain9Ex* obj, UINT* pLastPresentCount) {
	return obj->lpVtbl->GetLastPresentCount(obj, pLastPresentCount);
}

HRESULT IDirect3DSwapChain9ExGetPresentStats(IDirect3DSwapChain9Ex* obj, D3DPRESENTSTATS* pPresentationStatistics) {
	return obj->lpVtbl->GetPresentStats(obj, pPresentationStatistics);
}

void IDirect3DSwapChain9ExRelease(IDirect3DSwapChain9Ex* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"

type SwapChainEx struct {
	SwapChain
	handle *C.IDirect3DSwapChain9Ex
}

func (obj SwapChainEx) Release() {
	C.IDirect3DSwapChain9ExRelease(obj.handle)
}

func (obj SwapChainEx) GetDisplayModeEx() (pMode DISPLAYMODEEX, pRotation DISPLAYROTATION, err error) {
	var c_pMode C.D3DDISPLAYMODEEX
	var c_pRotation C.D3DDISPLAYROTATION
	err = toErr(C.IDirect3DSwapChain9ExGetDisplayModeEx(obj.handle, &c_pMode, &c_pRotation))
	pMode.fromC(&c_pMode)
	pRotation = (DISPLAYROTATION)(c_pRotation)
	return
}

func (obj SwapChainEx) GetLastPresentCount() (pLastPresentCount uint, err error) {
	var c_pLastPresentCount C.UINT
	err = toErr(C.IDirect3DSwapChain9ExGetLastPresentCount(obj.handle, &c_pLastPresentCount))
	pLastPresentCount = (uint)(c_pLastPresentCount)
	return
}

func (obj SwapChainEx) GetPresentStats() (pPresentationStatistics PRESENTSTATS, err error) {
	var c_pPresentationStatistics C.D3DPRESENTSTATS
	err = toErr(C.IDirect3DSwapChain9ExGetPresentStats(obj.handle, &c_pPresentationStatistics))
	pPresentationStatistics.fromC(&c_pPresentationStatistics)
	return
}
