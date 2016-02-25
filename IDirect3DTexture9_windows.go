package d3d9

/*
#include "d3d9wrapper.h"

HRESULT IDirect3DTexture9AddDirtyRect(
		IDirect3DTexture9* obj,
		RECT* pDirtyRect) {
	return obj->lpVtbl->AddDirtyRect(obj, pDirtyRect);
}

HRESULT IDirect3DTexture9GetLevelDesc(
		IDirect3DTexture9* obj,
		UINT Level,
		D3DSURFACE_DESC* pDesc) {
	return obj->lpVtbl->GetLevelDesc(obj, Level, pDesc);
}

HRESULT IDirect3DTexture9GetSurfaceLevel(
		IDirect3DTexture9* obj,
		UINT Level,
		IDirect3DSurface9** ppSurfaceLevel) {
	return obj->lpVtbl->GetSurfaceLevel(obj, Level, ppSurfaceLevel);
}

HRESULT IDirect3DTexture9LockRect(
		IDirect3DTexture9* obj,
		UINT Level,
		D3DLOCKED_RECT* pLockedRect,
		RECT* pRect,
		DWORD Flags) {
	return obj->lpVtbl->LockRect(obj, Level, pLockedRect, pRect, Flags);
}

HRESULT IDirect3DTexture9UnlockRect(IDirect3DTexture9* obj, UINT Level) {
	return obj->lpVtbl->UnlockRect(obj, Level);
}

void IDirect3DTexture9Release(IDirect3DTexture9* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"
import "unsafe"

// Texture and its methods are used to manipulate a texture resource.
type Texture struct {
	BaseTexture
	handle *C.IDirect3DTexture9
}

// Release has to be called when finished using the object to free its
// associated resources.
func (obj Texture) Release() {
	C.IDirect3DTexture9Release(obj.handle)
}

// AddDirtyRect adds a dirty region to a texture resource.
func (obj Texture) AddDirtyRect(pDirtyRect *RECT) (err Error) {
	if pDirtyRect == nil {
		err = toErr(C.IDirect3DTexture9AddDirtyRect(obj.handle, nil))
	} else {
		cpDirtyRect := pDirtyRect.toC()
		err = toErr(C.IDirect3DTexture9AddDirtyRect(obj.handle, &cpDirtyRect))
	}
	return
}

// GetLevelDesc retrieves a level description of a texture resource.
func (obj Texture) GetLevelDesc(Level uint) (pDesc SURFACE_DESC, err Error) {
	var cpDesc C.D3DSURFACE_DESC
	err = toErr(C.IDirect3DTexture9GetLevelDesc(
		obj.handle,
		(C.UINT)(Level),
		&cpDesc,
	))
	pDesc.fromC(&cpDesc)
	return
}

// GetSurfaceLevel retrieves the specified texture surface level.
func (obj Texture) GetSurfaceLevel(
	Level uint,
) (
	ppSurfaceLevel Surface,
	err Error,
) {
	var cppSurfaceLevel *C.IDirect3DSurface9
	err = toErr(C.IDirect3DTexture9GetSurfaceLevel(
		obj.handle,
		(C.UINT)(Level),
		&cppSurfaceLevel,
	))
	resource := Resource{
		(*C.IDirect3DResource9)(unsafe.Pointer(cppSurfaceLevel)),
	}
	ppSurfaceLevel = Surface{resource, cppSurfaceLevel}
	return
}

// LockRect locks a rectangle on a texture resource.
func (obj Texture) LockRect(
	Level uint,
	pRect *RECT,
	Flags uint32,
) (
	pLockedRect LOCKED_RECT,
	err Error,
) {
	var cpLockedRect C.D3DLOCKED_RECT
	if pRect == nil {
		err = toErr(C.IDirect3DTexture9LockRect(
			obj.handle,
			(C.UINT)(Level),
			&cpLockedRect,
			nil,
			(C.DWORD)(Flags),
		))
	} else {
		cpRect := pRect.toC()
		err = toErr(C.IDirect3DTexture9LockRect(
			obj.handle,
			(C.UINT)(Level),
			&cpLockedRect,
			&cpRect,
			(C.DWORD)(Flags),
		))
	}
	pLockedRect.fromC(&cpLockedRect)
	return
}

// UnlockRect unlocks a rectangle on a texture resource.
func (obj Texture) UnlockRect(Level uint) (err Error) {
	err = toErr(C.IDirect3DTexture9UnlockRect(obj.handle, (C.UINT)(Level)))
	return
}
