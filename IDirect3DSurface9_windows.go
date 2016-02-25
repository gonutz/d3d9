package d3d9

/*
#include "d3d9wrapper.h"

HRESULT IDirect3DSurface9GetContainerCubeTexture(
		IDirect3DSurface9* obj,
		REFIID riid,
		IDirect3DCubeTexture9** ppContainer) {
	return obj->lpVtbl->GetContainer(obj, riid, (void**)ppContainer);
}

HRESULT IDirect3DSurface9GetContainerTexture(
		IDirect3DSurface9* obj,
		REFIID riid,
		IDirect3DTexture9** ppContainer) {
	return obj->lpVtbl->GetContainer(obj, riid, (void**)ppContainer);
}

HRESULT IDirect3DSurface9GetContainerSwapChain(
		IDirect3DSurface9* obj,
		REFIID riid,
		IDirect3DSwapChain9** ppContainer) {
	return obj->lpVtbl->GetContainer(obj, riid, (void**)ppContainer);
}

HRESULT IDirect3DSurface9GetDC(IDirect3DSurface9* obj, HDC* phdc) {
	return obj->lpVtbl->GetDC(obj, phdc);
}

HRESULT IDirect3DSurface9GetDesc(
		IDirect3DSurface9* obj,
		D3DSURFACE_DESC* pDesc) {
	return obj->lpVtbl->GetDesc(obj, pDesc);
}

HRESULT IDirect3DSurface9LockRect(
		IDirect3DSurface9* obj,
		D3DLOCKED_RECT* pLockedRect,
		RECT* pRect,
		DWORD Flags) {
	return obj->lpVtbl->LockRect(obj, pLockedRect, pRect, Flags);
}

HRESULT IDirect3DSurface9ReleaseDC(IDirect3DSurface9* obj, HDC hdc) {
	return obj->lpVtbl->ReleaseDC(obj, hdc);
}

HRESULT IDirect3DSurface9UnlockRect(IDirect3DSurface9* obj) {
	return obj->lpVtbl->UnlockRect(obj);
}

void IDirect3DSurface9Release(IDirect3DSurface9* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"
import "unsafe"

// Surface and its methods are used to query and prepare surfaces.
type Surface struct {
	Resource
	handle *C.IDirect3DSurface9
}

// Release has to be called when finished using the object to free its
// associated resources.
func (obj Surface) Release() {
	C.IDirect3DSurface9Release(obj.handle)
}

// GetContainerCubeTexture provides access to the parent cube texture.
func (obj Surface) GetContainerCubeTexture(
	riid GUID,
) (
	cubeTex CubeTexture,
	err Error,
) {
	criid := riid.toC()
	var chandle *C.IDirect3DCubeTexture9
	err = toErr(C.IDirect3DSurface9GetContainerCubeTexture(
		obj.handle,
		&criid,
		&chandle,
	))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(chandle))}
	baseTexture := BaseTexture{
		resource,
		(*C.IDirect3DBaseTexture9)(unsafe.Pointer(chandle)),
	}
	cubeTex = CubeTexture{baseTexture, chandle}
	return
}

// GetContainerTexture provides access to the parent texture.
func (obj Surface) GetContainerTexture(
	riid GUID,
) (
	tex Texture,
	err Error,
) {
	criid := riid.toC()
	var chandle *C.IDirect3DTexture9
	err = toErr(C.IDirect3DSurface9GetContainerTexture(
		obj.handle,
		&criid,
		&chandle,
	))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(chandle))}
	baseTexture := BaseTexture{
		resource,
		(*C.IDirect3DBaseTexture9)(unsafe.Pointer(chandle)),
	}
	tex = Texture{baseTexture, chandle}
	return
}

// GetContainerSwapChain provides access to the parent swap chain if this
// surface is a back buffer.
func (obj Surface) GetContainerSwapChain(
	riid GUID,
) (
	chain SwapChain,
	err Error,
) {
	criid := riid.toC()
	var chandle *C.IDirect3DSwapChain9
	err = toErr(C.IDirect3DSurface9GetContainerSwapChain(
		obj.handle,
		&criid,
		&chandle,
	))
	chain = SwapChain{chandle}
	return
}

// GetDC retrieves a device context.
func (obj Surface) GetDC() (phdc unsafe.Pointer, err Error) {
	var cphdc C.HDC
	err = toErr(C.IDirect3DSurface9GetDC(obj.handle, &cphdc))
	phdc = (unsafe.Pointer)(cphdc)
	return
}

// GetDesc retrieves a description of the surface.
func (obj Surface) GetDesc() (pDesc SURFACE_DESC, err Error) {
	var cpDesc C.D3DSURFACE_DESC
	err = toErr(C.IDirect3DSurface9GetDesc(obj.handle, &cpDesc))
	pDesc.fromC(&cpDesc)
	return
}

// LockRect locks a rectangle on a surface.
func (obj Surface) LockRect(
	pRect *RECT,
	Flags uint32,
) (
	pLockedRect LOCKED_RECT,
	err Error,
) {
	var cpLockedRect C.D3DLOCKED_RECT
	if pRect == nil {
		err = toErr(C.IDirect3DSurface9LockRect(
			obj.handle,
			&cpLockedRect,
			nil,
			(C.DWORD)(Flags),
		))
	} else {
		cpRect := pRect.toC()
		err = toErr(C.IDirect3DSurface9LockRect(
			obj.handle,
			&cpLockedRect,
			&cpRect,
			(C.DWORD)(Flags),
		))
	}
	pLockedRect.fromC(&cpLockedRect)
	return
}

// ReleaseDC releases a device context handle.
func (obj Surface) ReleaseDC(hdc unsafe.Pointer) (err Error) {
	err = toErr(C.IDirect3DSurface9ReleaseDC(obj.handle, (C.HDC)(hdc)))
	return
}

// UnlockRect unlocks a rectangle on a surface.
func (obj Surface) UnlockRect() (err Error) {
	err = toErr(C.IDirect3DSurface9UnlockRect(obj.handle))
	return
}
