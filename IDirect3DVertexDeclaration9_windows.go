package d3d9

/*
#include "d3d9wrapper.h"

HRESULT IDirect3DVertexDeclaration9GetDeclaration(
		IDirect3DVertexDeclaration9* obj,
		D3DVERTEXELEMENT9* pDecl,
		UINT* pNumElements) {
	return obj->lpVtbl->GetDeclaration(obj, pDecl, pNumElements);
}

HRESULT IDirect3DVertexDeclaration9GetDevice(
		IDirect3DVertexDeclaration9* obj,
		IDirect3DDevice9** ppDevice) {
	return obj->lpVtbl->GetDevice(obj, ppDevice);
}

void IDirect3DVertexDeclaration9Release(IDirect3DVertexDeclaration9* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"

// VertexDeclaration and its methods are used to encapsulate the vertex shader
// declaration.
type VertexDeclaration struct {
	handle *C.IDirect3DVertexDeclaration9
}

// Release has to be called when finished using the object to free its
// associated resources.
func (obj VertexDeclaration) Release() {
	C.IDirect3DVertexDeclaration9Release(obj.handle)
}

// GetDeclaration returns the vertex shader declaration.
func (obj VertexDeclaration) GetDeclaration() (
	pDecl []VERTEXELEMENT,
	err Error,
) {
	// first pass nil for the elements to get the count
	var cpNumElements C.UINT
	err = toErr(C.IDirect3DVertexDeclaration9GetDeclaration(
		obj.handle,
		nil,
		&cpNumElements,
	))
	if err != nil {
		return
	}
	cpDecl := make([]C.D3DVERTEXELEMENT9, cpNumElements)
	err = toErr(C.IDirect3DVertexDeclaration9GetDeclaration(
		obj.handle,
		&cpDecl[0],
		&cpNumElements,
	))
	if err != nil {
		return
	}
	pDecl = make([]VERTEXELEMENT, cpNumElements)
	for i := range pDecl {
		pDecl[i].fromC(&cpDecl[i])
	}
	return
}

// GetDevice returns the current device.
func (obj VertexDeclaration) GetDevice() (ppDevice Device, err Error) {
	var cppDevice *C.IDirect3DDevice9
	err = toErr(C.IDirect3DVertexDeclaration9GetDevice(obj.handle, &cppDevice))
	ppDevice = Device{cppDevice}
	return
}
