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

type PixelShader struct {
	handle *C.IDirect3DPixelShader9
}

func (obj PixelShader) Release() {
	C.IDirect3DPixelShader9Release(obj.handle)
}

func (obj PixelShader) GetDevice() (ppDevice Device, err error) {
	var c_ppDevice *C.IDirect3DDevice9
	err = toErr(C.IDirect3DPixelShader9GetDevice(obj.handle, &c_ppDevice))
	ppDevice = Device{c_ppDevice}
	return
}

// GetFunction returns the shader data.
func (obj PixelShader) GetFunction() (pData []byte, err error) {
	// first get the needed buffer size, pass nil as the data pointer
	var c_pSizeOfDataInBytes C.UINT
	err = toErr(C.IDirect3DPixelShader9GetFunction(
		obj.handle,
		nil,
		&c_pSizeOfDataInBytes,
	))
	if err != nil {
		return
	}
	pData = make([]byte, c_pSizeOfDataInBytes)
	err = toErr(C.IDirect3DPixelShader9GetFunction(
		obj.handle,
		unsafe.Pointer(&pData[0]),
		&c_pSizeOfDataInBytes,
	))
	return
}