package d3d9

/*
#include <d3d9.h>

HRESULT IDirect3DTexture9AddDirtyRect(IDirect3DTexture9* obj, RECT* pDirtyRect) {
	return obj->lpVtbl->AddDirtyRect(obj, pDirtyRect);
}

HRESULT IDirect3DTexture9GetLevelDesc(IDirect3DTexture9* obj, UINT Level, D3DSURFACE_DESC* pDesc) {
	return obj->lpVtbl->GetLevelDesc(obj, Level, pDesc);
}

HRESULT IDirect3DTexture9GetSurfaceLevel(IDirect3DTexture9* obj, UINT Level, IDirect3DSurface9** ppSurfaceLevel) {
	return obj->lpVtbl->GetSurfaceLevel(obj, Level, ppSurfaceLevel);
}

HRESULT IDirect3DTexture9LockRect(IDirect3DTexture9* obj, UINT Level, D3DLOCKED_RECT* pLockedRect, RECT* pRect, DWORD Flags) {
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

type Texture struct {
	BaseTexture
	handle *C.IDirect3DTexture9
}

func (obj Texture) Release() {
	C.IDirect3DTexture9Release(obj.handle)
}

func (obj Texture) AddDirtyRect(pDirtyRect *RECT) (err error) {
	if pDirtyRect == nil {
		err = toErr(C.IDirect3DTexture9AddDirtyRect(obj.handle, nil))
	} else {
		c_pDirtyRect := pDirtyRect.toC()
		err = toErr(C.IDirect3DTexture9AddDirtyRect(obj.handle, &c_pDirtyRect))
	}
	return
}

func (obj Texture) GetLevelDesc(Level uint) (pDesc SURFACE_DESC, err error) {
	var c_pDesc C.D3DSURFACE_DESC
	err = toErr(C.IDirect3DTexture9GetLevelDesc(obj.handle, (C.UINT)(Level), &c_pDesc))
	pDesc.fromC(&c_pDesc)
	return
}

func (obj Texture) GetSurfaceLevel(Level uint) (ppSurfaceLevel Surface, err error) {
	var c_ppSurfaceLevel *C.IDirect3DSurface9
	err = toErr(C.IDirect3DTexture9GetSurfaceLevel(obj.handle, (C.UINT)(Level), &c_ppSurfaceLevel))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(c_ppSurfaceLevel))}
	ppSurfaceLevel = Surface{resource, c_ppSurfaceLevel}
	return
}

func (obj Texture) LockRect(Level uint, pRect *RECT, Flags uint32) (pLockedRect LOCKED_RECT, err error) {
	var c_pLockedRect C.D3DLOCKED_RECT
	if pRect == nil {
		err = toErr(C.IDirect3DTexture9LockRect(obj.handle, (C.UINT)(Level), &c_pLockedRect, nil, (C.DWORD)(Flags)))
	} else {
		c_pRect := pRect.toC()
		err = toErr(C.IDirect3DTexture9LockRect(obj.handle, (C.UINT)(Level), &c_pLockedRect, &c_pRect, (C.DWORD)(Flags)))
	}
	pLockedRect.fromC(&c_pLockedRect)
	return
}

func (obj Texture) UnlockRect(Level uint) (err error) {
	err = toErr(C.IDirect3DTexture9UnlockRect(obj.handle, (C.UINT)(Level)))
	return
}
