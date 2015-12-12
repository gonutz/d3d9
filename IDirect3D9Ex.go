package d3d9

/*
#include <d3d9.h>

HRESULT IDirect3D9ExCreateDeviceEx(IDirect3D9Ex* obj, UINT Adapter, D3DDEVTYPE DeviceType, HWND hFocusWindow, DWORD BehaviorFlags, D3DPRESENT_PARAMETERS* pPresentationParameters, D3DDISPLAYMODEEX* pFullscreenDisplayMode, IDirect3DDevice9Ex** ppReturnedDeviceInterface) {
	return obj->lpVtbl->CreateDeviceEx(obj, Adapter, DeviceType, hFocusWindow, BehaviorFlags, pPresentationParameters, pFullscreenDisplayMode, ppReturnedDeviceInterface);
}

HRESULT IDirect3D9ExEnumAdapterModesEx(IDirect3D9Ex* obj, UINT Adapter, D3DDISPLAYMODEFILTER* pFilter, UINT Mode, D3DDISPLAYMODEEX* pMode) {
	return obj->lpVtbl->EnumAdapterModesEx(obj, Adapter, pFilter, Mode, pMode);
}

HRESULT IDirect3D9ExGetAdapterDisplayModeEx(IDirect3D9Ex* obj, UINT Adapter, D3DDISPLAYMODEEX* pMode, D3DDISPLAYROTATION* pRotation) {
	return obj->lpVtbl->GetAdapterDisplayModeEx(obj, Adapter, pMode, pRotation);
}

HRESULT IDirect3D9ExGetAdapterLUID(IDirect3D9Ex* obj, UINT Adapter, LUID* pLUID) {
	return obj->lpVtbl->GetAdapterLUID(obj, Adapter, pLUID);
}

UINT IDirect3D9ExGetAdapterModeCountEx(IDirect3D9Ex* obj, UINT Adapter, D3DDISPLAYMODEFILTER* pFilter) {
	return obj->lpVtbl->GetAdapterModeCountEx(obj, Adapter, pFilter);
}

void IDirect3D9ExRelease(IDirect3D9Ex* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"
import "unsafe"

type Direct3DEx struct {
	Direct3D
	handle *C.IDirect3D9Ex
}

func (obj Direct3DEx) Release() {
	C.IDirect3D9ExRelease(obj.handle)
}

func (obj Direct3DEx) CreateDeviceEx(Adapter uint, DeviceType DEVTYPE, hFocusWindow unsafe.Pointer, BehaviorFlags uint32, inpPresentationParameters PRESENT_PARAMETERS, inpFullscreenDisplayMode DISPLAYMODEEX) (ppReturnedDeviceInterface DeviceEx, outpPresentationParameters PRESENT_PARAMETERS, outpFullscreenDisplayMode DISPLAYMODEEX, err error) {
	c_pPresentationParameters := inpPresentationParameters.toC()
	c_pFullscreenDisplayMode := inpFullscreenDisplayMode.toC()
	var c_ppReturnedDeviceInterface *C.IDirect3DDevice9Ex
	err = toErr(C.IDirect3D9ExCreateDeviceEx(obj.handle, (C.UINT)(Adapter), (C.D3DDEVTYPE)(DeviceType), (C.HWND)(hFocusWindow), (C.DWORD)(BehaviorFlags), &c_pPresentationParameters, &c_pFullscreenDisplayMode, &c_ppReturnedDeviceInterface))
	outpPresentationParameters.fromC(&c_pPresentationParameters)
	outpFullscreenDisplayMode.fromC(&c_pFullscreenDisplayMode)
	device := Device{(*C.IDirect3DDevice9)(unsafe.Pointer(c_ppReturnedDeviceInterface))}
	ppReturnedDeviceInterface = DeviceEx{device, c_ppReturnedDeviceInterface}
	return
}

func (obj Direct3DEx) EnumAdapterModesEx(Adapter uint, pFilter DISPLAYMODEFILTER, Mode uint) (pMode DISPLAYMODEEX, err error) {
	c_pFilter := pFilter.toC()
	var c_pMode C.D3DDISPLAYMODEEX
	err = toErr(C.IDirect3D9ExEnumAdapterModesEx(obj.handle, (C.UINT)(Adapter), &c_pFilter, (C.UINT)(Mode), &c_pMode))
	pMode.fromC(&c_pMode)
	return
}

func (obj Direct3DEx) GetAdapterDisplayModeEx(Adapter uint, inpMode DISPLAYMODEEX, inpRotation DISPLAYROTATION) (outpMode DISPLAYMODEEX, outpRotation DISPLAYROTATION, err error) {
	c_pMode := inpMode.toC()
	c_pRotation := (C.D3DDISPLAYROTATION)(inpRotation)
	err = toErr(C.IDirect3D9ExGetAdapterDisplayModeEx(obj.handle, (C.UINT)(Adapter), &c_pMode, &c_pRotation))
	outpMode.fromC(&c_pMode)
	outpRotation = (DISPLAYROTATION)(c_pRotation)
	return
}

func (obj Direct3DEx) GetAdapterLUID(Adapter uint, pLUID LUID) (err error) {
	c_pLUID := pLUID.toC()
	err = toErr(C.IDirect3D9ExGetAdapterLUID(obj.handle, (C.UINT)(Adapter), &c_pLUID))
	return
}

func (obj Direct3DEx) GetAdapterModeCountEx(Adapter uint, pFilter DISPLAYMODEFILTER) {
	c_pFilter := pFilter.toC()
	C.IDirect3D9ExGetAdapterModeCountEx(obj.handle, (C.UINT)(Adapter), &c_pFilter)
}
