package d3d9

/*
#include <d3d9.h>

HRESULT IDirect3DSurface9GetContainer(IDirect3DSurface9* obj, REFIID riid, void** ppContainer) {
	return obj->lpVtbl->GetContainer(obj, riid, ppContainer);
}

HRESULT IDirect3DSurface9GetDC(IDirect3DSurface9* obj, HDC* phdc) {
	return obj->lpVtbl->GetDC(obj, phdc);
}

HRESULT IDirect3DSurface9GetDesc(IDirect3DSurface9* obj, D3DSURFACE_DESC* pDesc) {
	return obj->lpVtbl->GetDesc(obj, pDesc);
}

HRESULT IDirect3DSurface9LockRect(IDirect3DSurface9* obj, D3DLOCKED_RECT* pLockedRect, RECT* pRect, DWORD Flags) {
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

func (obj Surface) GetContainer(riid GUID) (phdc unsafe.Pointer, err error) {
	c_riid := riid.toC()
	err = toErr(C.IDirect3DSurface9GetContainer(obj.handle, &c_riid, &phdc))
	return
}

func (obj Surface) GetDC() (phdc unsafe.Pointer, err error) {
	var c_phdc C.HDC
	err = toErr(C.IDirect3DSurface9GetDC(obj.handle, &c_phdc))
	phdc = (unsafe.Pointer)(c_phdc)
	return
}

func (obj Surface) GetDesc() (pDesc SURFACE_DESC, err error) {
	var c_pDesc C.D3DSURFACE_DESC
	err = toErr(C.IDirect3DSurface9GetDesc(obj.handle, &c_pDesc))
	pDesc.fromC(&c_pDesc)
	return
}

func (obj Surface) LockRect(pRect *RECT, Flags uint32) (pLockedRect LOCKED_RECT, err error) {
	var c_pLockedRect C.D3DLOCKED_RECT
	if pRect == nil {
		err = toErr(C.IDirect3DSurface9LockRect(obj.handle, &c_pLockedRect, nil, (C.DWORD)(Flags)))
	} else {
		c_pRect := pRect.toC()
		err = toErr(C.IDirect3DSurface9LockRect(obj.handle, &c_pLockedRect, &c_pRect, (C.DWORD)(Flags)))
	}
	pLockedRect.fromC(&c_pLockedRect)
	return
}

func (obj Surface) ReleaseDC(hdc unsafe.Pointer) (err error) {
	err = toErr(C.IDirect3DSurface9ReleaseDC(obj.handle, (C.HDC)(hdc)))
	return
}

func (obj Surface) UnlockRect() (err error) {
	err = toErr(C.IDirect3DSurface9UnlockRect(obj.handle))
	return
}
