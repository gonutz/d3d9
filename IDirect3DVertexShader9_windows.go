package d3d9

/*
#include "d3d9wrapper.h"

HRESULT IDirect3DVertexShader9GetDevice(
		IDirect3DVertexShader9* obj,
		IDirect3DDevice9** ppDevice) {
	return obj->lpVtbl->GetDevice(obj, ppDevice);
}

HRESULT IDirect3DVertexShader9GetFunction(
		IDirect3DVertexShader9* obj,
		void* pData,
		UINT* pSizeOfData) {
	return obj->lpVtbl->GetFunction(obj, pData, pSizeOfData);
}

void IDirect3DVertexShader9Release(IDirect3DVertexShader9* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"
import "unsafe"

// VertexShader and its methods are used to encapsulate the functionality of a
// vertex shader.
type VertexShader struct {
	handle *C.IDirect3DVertexShader9
}

// Release has to be called when finished using the object to free its
// associated resources.
func (obj VertexShader) Release() {
	C.IDirect3DVertexShader9Release(obj.handle)
}

// GetDevice retrieves the device.
func (obj VertexShader) GetDevice() (ppDevice Device, err Error) {
	var cppDevice *C.IDirect3DDevice9
	err = toErr(C.IDirect3DVertexShader9GetDevice(obj.handle, &cppDevice))
	ppDevice = Device{cppDevice}
	return
}

// GetFunction returns the shader data.
func (obj VertexShader) GetFunction() (pData []byte, err Error) {
	// first get the needed buffer size, pass nil as the data pointer
	var cpSizeOfDataInBytes C.UINT
	err = toErr(C.IDirect3DVertexShader9GetFunction(
		obj.handle,
		nil,
		&cpSizeOfDataInBytes,
	))
	if err != nil {
		return
	}
	pData = make([]byte, cpSizeOfDataInBytes)
	err = toErr(C.IDirect3DVertexShader9GetFunction(
		obj.handle,
		unsafe.Pointer(&pData[0]),
		&cpSizeOfDataInBytes,
	))
	return
}
