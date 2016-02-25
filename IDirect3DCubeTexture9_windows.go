package d3d9

/*
#include "d3d9wrapper.h"

HRESULT IDirect3DCubeTexture9AddDirtyRect(
		IDirect3DCubeTexture9* obj,
		D3DCUBEMAP_FACES FaceType,
		RECT* pDirtyRect) {
	return obj->lpVtbl->AddDirtyRect(obj, FaceType, pDirtyRect);
}

HRESULT IDirect3DCubeTexture9GetCubeMapSurface(
		IDirect3DCubeTexture9* obj,
		D3DCUBEMAP_FACES FaceType,
		UINT Level,
		IDirect3DSurface9** ppCubeMapSurface) {
	return obj->lpVtbl->GetCubeMapSurface(obj, FaceType, Level,
		ppCubeMapSurface);
}

HRESULT IDirect3DCubeTexture9GetLevelDesc(
		IDirect3DCubeTexture9* obj,
		UINT Level,
		D3DSURFACE_DESC* pDesc) {
	return obj->lpVtbl->GetLevelDesc(obj, Level, pDesc);
}

HRESULT IDirect3DCubeTexture9LockRect(
		IDirect3DCubeTexture9* obj,
		D3DCUBEMAP_FACES FaceType,
		UINT Level,
		D3DLOCKED_RECT* pLockedRect,
		RECT* pRect,
		DWORD Flags) {
	return obj->lpVtbl->LockRect(obj, FaceType, Level, pLockedRect, pRect,
		Flags);
}

HRESULT IDirect3DCubeTexture9UnlockRect(
		IDirect3DCubeTexture9* obj,
		D3DCUBEMAP_FACES FaceType,
		UINT Level) {
	return obj->lpVtbl->UnlockRect(obj, FaceType, Level);
}

void IDirect3DCubeTexture9Release(IDirect3DCubeTexture9* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"
import "unsafe"

// CubeTexture and its methods are used to manipulate a cube texture resource.
type CubeTexture struct {
	BaseTexture
	handle *C.IDirect3DCubeTexture9
}

// Release has to be called when finished using the object to free its
// associated resources.
func (obj CubeTexture) Release() {
	C.IDirect3DCubeTexture9Release(obj.handle)
}

// AddDirtyRect adds a dirty region to a cube texture resource.
func (obj CubeTexture) AddDirtyRect(
	FaceType CUBEMAP_FACES,
	pDirtyRect *RECT,
) (
	err Error,
) {
	if pDirtyRect == nil {
		err = toErr(C.IDirect3DCubeTexture9AddDirtyRect(
			obj.handle,
			(C.D3DCUBEMAP_FACES)(FaceType),
			nil,
		))
	} else {
		cpDirtyRect := pDirtyRect.toC()
		err = toErr(C.IDirect3DCubeTexture9AddDirtyRect(
			obj.handle,
			(C.D3DCUBEMAP_FACES)(FaceType),
			&cpDirtyRect,
		))
	}
	return
}

// GetCubeMapSurface retrieves a cube texture map surface.
func (obj CubeTexture) GetCubeMapSurface(
	FaceType CUBEMAP_FACES,
	Level uint,
) (
	ppCubeMapSurface Surface,
	err Error,
) {
	var cppCubeMapSurface *C.IDirect3DSurface9
	err = toErr(C.IDirect3DCubeTexture9GetCubeMapSurface(
		obj.handle,
		(C.D3DCUBEMAP_FACES)(FaceType),
		(C.UINT)(Level),
		&cppCubeMapSurface,
	))
	resource := Resource{
		(*C.IDirect3DResource9)(unsafe.Pointer(cppCubeMapSurface)),
	}
	ppCubeMapSurface = Surface{resource, cppCubeMapSurface}
	return
}

// GetLevelDesc retrieves a description of one face of the specified cube
// texture level.
func (obj CubeTexture) GetLevelDesc(
	Level uint,
) (
	pDesc SURFACE_DESC,
	err Error,
) {
	var cpDesc C.D3DSURFACE_DESC
	err = toErr(C.IDirect3DCubeTexture9GetLevelDesc(
		obj.handle,
		(C.UINT)(Level),
		&cpDesc,
	))
	pDesc.fromC(&cpDesc)
	return
}

// LockRect locks a rectangle on a cube texture resource.
func (obj CubeTexture) LockRect(
	FaceType CUBEMAP_FACES,
	Level uint,
	pRect *RECT,
	Flags uint32,
) (
	pLockedRect LOCKED_RECT,
	err Error,
) {
	var cpLockedRect C.D3DLOCKED_RECT
	if pRect == nil {
		err = toErr(C.IDirect3DCubeTexture9LockRect(
			obj.handle,
			(C.D3DCUBEMAP_FACES)(FaceType),
			(C.UINT)(Level),
			&cpLockedRect,
			nil,
			(C.DWORD)(Flags),
		))
	} else {
		cpRect := pRect.toC()
		err = toErr(C.IDirect3DCubeTexture9LockRect(
			obj.handle,
			(C.D3DCUBEMAP_FACES)(FaceType),
			(C.UINT)(Level),
			&cpLockedRect,
			&cpRect,
			(C.DWORD)(Flags),
		))
	}
	pLockedRect.fromC(&cpLockedRect)
	return
}

// UnlockRect unlocks a rectangle on a cube texture resource.
func (obj CubeTexture) UnlockRect(
	FaceType CUBEMAP_FACES,
	Level uint,
) (
	err Error,
) {
	err = toErr(C.IDirect3DCubeTexture9UnlockRect(
		obj.handle,
		(C.D3DCUBEMAP_FACES)(FaceType),
		(C.UINT)(Level),
	))
	return
}
