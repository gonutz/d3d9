package d3d9

/*
#include <d3d9.h>

VOID IDirect3DBaseTexture9GenerateMipSubLevels(IDirect3DBaseTexture9* obj) {
	obj->lpVtbl->GenerateMipSubLevels(obj);
}

D3DTEXTUREFILTERTYPE IDirect3DBaseTexture9GetAutoGenFilterType(IDirect3DBaseTexture9* obj) {
	return obj->lpVtbl->GetAutoGenFilterType(obj);
}

DWORD IDirect3DBaseTexture9GetLevelCount(IDirect3DBaseTexture9* obj) {
	return obj->lpVtbl->GetLevelCount(obj);
}

DWORD IDirect3DBaseTexture9GetLOD(IDirect3DBaseTexture9* obj) {
	return obj->lpVtbl->GetLOD(obj);
}

HRESULT IDirect3DBaseTexture9SetAutoGenFilterType(IDirect3DBaseTexture9* obj, D3DTEXTUREFILTERTYPE FilterType) {
	return obj->lpVtbl->SetAutoGenFilterType(obj, FilterType);
}

DWORD IDirect3DBaseTexture9SetLOD(IDirect3DBaseTexture9* obj, DWORD LODNew) {
	return obj->lpVtbl->SetLOD(obj, LODNew);
}

void IDirect3DBaseTexture9Release(IDirect3DBaseTexture9* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"

type BaseTexture struct {
	Resource
	handle *C.IDirect3DBaseTexture9
}

func (obj BaseTexture) Release() {
	C.IDirect3DBaseTexture9Release(obj.handle)
}

func (obj BaseTexture) GenerateMipSubLevels() {
	C.IDirect3DBaseTexture9GenerateMipSubLevels(obj.handle)
}

func (obj BaseTexture) GetAutoGenFilterType() TEXTUREFILTERTYPE {
	return (TEXTUREFILTERTYPE)(C.IDirect3DBaseTexture9GetAutoGenFilterType(obj.handle))
}

func (obj BaseTexture) GetLevelCount() uint32 {
	return (uint32)(C.IDirect3DBaseTexture9GetLevelCount(obj.handle))
}

func (obj BaseTexture) GetLOD() uint32 {
	return (uint32)(C.IDirect3DBaseTexture9GetLOD(obj.handle))
}

func (obj BaseTexture) SetAutoGenFilterType(FilterType TEXTUREFILTERTYPE) (err error) {
	err = toErr(C.IDirect3DBaseTexture9SetAutoGenFilterType(obj.handle, (C.D3DTEXTUREFILTERTYPE)(FilterType)))
	return
}

func (obj BaseTexture) SetLOD(LODNew uint32) uint32 {
	return (uint32)(C.IDirect3DBaseTexture9SetLOD(obj.handle, (C.DWORD)(LODNew)))
}
