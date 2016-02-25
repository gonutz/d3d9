package d3d9

/*
#include "d3d9wrapper.h"

HRESULT IDirect3DPixelShader9GetDevice(
		IDirect3DPixelShader9* obj,
		IDirect3DDevice9** ppDevice) {
	return obj->lpVtbl->GetDevice(obj, ppDevice);
}

HRESULT IDirect3DPixelShader9GetFunction(
		IDirect3DPixelShader9* obj,
		void* pData,
		UINT* pSizeOfData) {
	return obj->lpVtbl->GetFunction(obj, pData, pSizeOfData);
}

void IDirect3DPixelShader9Release(IDirect3DPixelShader9* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"
import "unsafe"

// PixelShader and its methods are used to manipulate a pixel shader resource.
type PixelShader struct {
	handle *C.IDirect3DPixelShader9
}

// Release has to be called when finished using the object to free its
// associated resources.
func (obj PixelShader) Release() {
	C.IDirect3DPixelShader9Release(obj.handle)
}

// GetDevice returns the device.
func (obj PixelShader) GetDevice() (ppDevice Device, err Error) {
	var cppDevice *C.IDirect3DDevice9
	err = toErr(C.IDirect3DPixelShader9GetDevice(obj.handle, &cppDevice))
	ppDevice = Device{cppDevice}
	return
}

// GetFunction returns the shader data.
func (obj PixelShader) GetFunction() (pData []byte, err Error) {
	// first get the needed buffer size, pass nil as the data pointer
	var cpSizeOfDataInBytes C.UINT
	err = toErr(C.IDirect3DPixelShader9GetFunction(
		obj.handle,
		nil,
		&cpSizeOfDataInBytes,
	))
	if err != nil {
		return
	}
	pData = make([]byte, cpSizeOfDataInBytes)
	err = toErr(C.IDirect3DPixelShader9GetFunction(
		obj.handle,
		unsafe.Pointer(&pData[0]),
		&cpSizeOfDataInBytes,
	))
	return
}
