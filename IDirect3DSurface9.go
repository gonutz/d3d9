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

type Surface struct {
	Resource
	handle *C.IDirect3DSurface9
}

func (obj Surface) Release() {
	C.IDirect3DSurface9Release(obj.handle)
}

// GetContainerCubeTexture provides access to the parent cube texture.
func (obj Surface) GetContainerCubeTexture(
	riid GUID,
) (
	cubeTex CubeTexture,
	err error,
) {
	c_riid := riid.toC()
	var c_handle *C.IDirect3DCubeTexture9
	err = toErr(C.IDirect3DSurface9GetContainerCubeTexture(
		obj.handle,
		&c_riid,
		&c_handle,
	))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(c_handle))}
	baseTexture := BaseTexture{
		resource,
		(*C.IDirect3DBaseTexture9)(unsafe.Pointer(c_handle)),
	}
	cubeTex = CubeTexture{baseTexture, c_handle}
	return
}

// GetContainerTexture provides access to the parent texture.
func (obj Surface) GetContainerTexture(
	riid GUID,
) (
	tex Texture,
	err error,
) {
	c_riid := riid.toC()
	var c_handle *C.IDirect3DTexture9
	err = toErr(C.IDirect3DSurface9GetContainerTexture(
		obj.handle,
		&c_riid,
		&c_handle,
	))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(c_handle))}
	baseTexture := BaseTexture{
		resource,
		(*C.IDirect3DBaseTexture9)(unsafe.Pointer(c_handle)),
	}
	tex = Texture{baseTexture, c_handle}
	return
}

// GetContainerSwapChain provides access to the parent swap chain if this
// surface is a back buffer.
func (obj Surface) GetContainerSwapChain(
	riid GUID,
) (
	chain SwapChain,
	err error,
) {
	c_riid := riid.toC()
	var c_handle *C.IDirect3DSwapChain9
	err = toErr(C.IDirect3DSurface9GetContainerSwapChain(
		obj.handle,
		&c_riid,
		&c_handle,
	))
	chain = SwapChain{c_handle}
	return
}

// GetDC retrieves a device context.
func (obj Surface) GetDC() (phdc unsafe.Pointer, err error) {
	var c_phdc C.HDC
	err = toErr(C.IDirect3DSurface9GetDC(obj.handle, &c_phdc))
	phdc = (unsafe.Pointer)(c_phdc)
	return
}

// GetDesc retrieves a description of the surface.
func (obj Surface) GetDesc() (pDesc SURFACE_DESC, err error) {
	var c_pDesc C.D3DSURFACE_DESC
	err = toErr(C.IDirect3DSurface9GetDesc(obj.handle, &c_pDesc))
	pDesc.fromC(&c_pDesc)
	return
}

// LockRect locks a rectangle on a surface.
func (obj Surface) LockRect(
	pRect *RECT,
	Flags uint32,
) (
	pLockedRect LOCKED_RECT,
	err error,
) {
	var c_pLockedRect C.D3DLOCKED_RECT
	if pRect == nil {
		err = toErr(C.IDirect3DSurface9LockRect(
			obj.handle,
			&c_pLockedRect,
			nil,
			(C.DWORD)(Flags),
		))
	} else {
		c_pRect := pRect.toC()
		err = toErr(C.IDirect3DSurface9LockRect(
			obj.handle,
			&c_pLockedRect,
			&c_pRect,
			(C.DWORD)(Flags),
		))
	}
	pLockedRect.fromC(&c_pLockedRect)
	return
}

// ReleaseDC releases a device context handle.
func (obj Surface) ReleaseDC(hdc unsafe.Pointer) (err error) {
	err = toErr(C.IDirect3DSurface9ReleaseDC(obj.handle, (C.HDC)(hdc)))
	return
}

// UnlockRect unlocks a rectangle on a surface.
func (obj Surface) UnlockRect() (err error) {
	err = toErr(C.IDirect3DSurface9UnlockRect(obj.handle))
	return
}
