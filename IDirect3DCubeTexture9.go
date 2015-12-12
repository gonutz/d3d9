package d3d9

/*
#include <d3d9.h>

HRESULT IDirect3DCubeTexture9AddDirtyRect(IDirect3DCubeTexture9* obj, D3DCUBEMAP_FACES FaceType, RECT* pDirtyRect) {
	return obj->lpVtbl->AddDirtyRect(obj, FaceType, pDirtyRect);
}

HRESULT IDirect3DCubeTexture9GetCubeMapSurface(IDirect3DCubeTexture9* obj, D3DCUBEMAP_FACES FaceType, UINT Level, IDirect3DSurface9** ppCubeMapSurface) {
	return obj->lpVtbl->GetCubeMapSurface(obj, FaceType, Level, ppCubeMapSurface);
}

HRESULT IDirect3DCubeTexture9GetLevelDesc(IDirect3DCubeTexture9* obj, UINT Level, D3DSURFACE_DESC* pDesc) {
	return obj->lpVtbl->GetLevelDesc(obj, Level, pDesc);
}

HRESULT IDirect3DCubeTexture9LockRect(IDirect3DCubeTexture9* obj, D3DCUBEMAP_FACES FaceType, UINT Level, D3DLOCKED_RECT* pLockedRect, RECT* pRect, DWORD Flags) {
	return obj->lpVtbl->LockRect(obj, FaceType, Level, pLockedRect, pRect, Flags);
}

HRESULT IDirect3DCubeTexture9UnlockRect(IDirect3DCubeTexture9* obj, D3DCUBEMAP_FACES FaceType, UINT Level) {
	return obj->lpVtbl->UnlockRect(obj, FaceType, Level);
}

void IDirect3DCubeTexture9Release(IDirect3DCubeTexture9* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"
import "unsafe"

type CubeTexture struct {
	BaseTexture
	handle *C.IDirect3DCubeTexture9
}

func (obj CubeTexture) Release() {
	C.IDirect3DCubeTexture9Release(obj.handle)
}

func (obj CubeTexture) AddDirtyRect(FaceType CUBEMAP_FACES, pDirtyRect RECT) (err error) {
	c_pDirtyRect := pDirtyRect.toC()
	err = toErr(C.IDirect3DCubeTexture9AddDirtyRect(obj.handle, (C.D3DCUBEMAP_FACES)(FaceType), &c_pDirtyRect))
	return
}

func (obj CubeTexture) GetCubeMapSurface(FaceType CUBEMAP_FACES, Level uint) (ppCubeMapSurface Surface, err error) {
	var c_ppCubeMapSurface *C.IDirect3DSurface9
	err = toErr(C.IDirect3DCubeTexture9GetCubeMapSurface(obj.handle, (C.D3DCUBEMAP_FACES)(FaceType), (C.UINT)(Level), &c_ppCubeMapSurface))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(c_ppCubeMapSurface))}
	ppCubeMapSurface = Surface{resource, c_ppCubeMapSurface}
	return
}

func (obj CubeTexture) GetLevelDesc(Level uint) (pDesc SURFACE_DESC, err error) {
	var c_pDesc C.D3DSURFACE_DESC
	err = toErr(C.IDirect3DCubeTexture9GetLevelDesc(obj.handle, (C.UINT)(Level), &c_pDesc))
	pDesc.fromC(&c_pDesc)
	return
}

func (obj CubeTexture) LockRect(FaceType CUBEMAP_FACES, Level uint, pRect RECT, Flags uint32) (pLockedRect LOCKED_RECT, err error) {
	var c_pLockedRect C.D3DLOCKED_RECT
	c_pRect := pRect.toC()
	err = toErr(C.IDirect3DCubeTexture9LockRect(obj.handle, (C.D3DCUBEMAP_FACES)(FaceType), (C.UINT)(Level), &c_pLockedRect, &c_pRect, (C.DWORD)(Flags)))
	pLockedRect.fromC(&c_pLockedRect)
	return
}

func (obj CubeTexture) UnlockRect(FaceType CUBEMAP_FACES, Level uint) (err error) {
	err = toErr(C.IDirect3DCubeTexture9UnlockRect(obj.handle, (C.D3DCUBEMAP_FACES)(FaceType), (C.UINT)(Level)))
	return
}
