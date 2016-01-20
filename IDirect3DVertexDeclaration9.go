package d3d9

/*
#include "d3d9wrapper.h"

HRESULT IDirect3DVertexDeclaration9GetDeclaration(IDirect3DVertexDeclaration9* obj, D3DVERTEXELEMENT9* pDecl, UINT* pNumElements) {
	return obj->lpVtbl->GetDeclaration(obj, pDecl, pNumElements);
}

HRESULT IDirect3DVertexDeclaration9GetDevice(IDirect3DVertexDeclaration9* obj, IDirect3DDevice9** ppDevice) {
	return obj->lpVtbl->GetDevice(obj, ppDevice);
}

void IDirect3DVertexDeclaration9Release(IDirect3DVertexDeclaration9* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"

type VertexDeclaration struct {
	handle *C.IDirect3DVertexDeclaration9
}

func (obj VertexDeclaration) Release() {
	C.IDirect3DVertexDeclaration9Release(obj.handle)
}

func (obj VertexDeclaration) GetDeclaration() (pDecl []VERTEXELEMENT, err error) {
	var c_pNumElements C.UINT
	err = toErr(C.IDirect3DVertexDeclaration9GetDeclaration(obj.handle, nil, &c_pNumElements))
	if err != nil {
		return
	}
	c_pDecl := make([]C.D3DVERTEXELEMENT9, c_pNumElements)
	err = toErr(C.IDirect3DVertexDeclaration9GetDeclaration(obj.handle, &c_pDecl[0], &c_pNumElements))
	if err != nil {
		return
	}
	pDecl = make([]VERTEXELEMENT, c_pNumElements)
	for i := range pDecl {
		pDecl[i].fromC(&c_pDecl[i])
	}
	return
}

func (obj VertexDeclaration) GetDevice() (ppDevice Device, err error) {
	var c_ppDevice *C.IDirect3DDevice9
	err = toErr(C.IDirect3DVertexDeclaration9GetDevice(obj.handle, &c_ppDevice))
	ppDevice = Device{c_ppDevice}
	return
}
