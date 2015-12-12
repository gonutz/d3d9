package d3d9

/*
#include <d3d9.h>

HRESULT IDirect3DStateBlock9Apply(IDirect3DStateBlock9* obj) {
	return obj->lpVtbl->Apply(obj);
}

HRESULT IDirect3DStateBlock9Capture(IDirect3DStateBlock9* obj) {
	return obj->lpVtbl->Capture(obj);
}

HRESULT IDirect3DStateBlock9GetDevice(IDirect3DStateBlock9* obj, IDirect3DDevice9** ppDevice) {
	return obj->lpVtbl->GetDevice(obj, ppDevice);
}

void IDirect3DStateBlock9Release(IDirect3DStateBlock9* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"

type StateBlock struct {
	handle *C.IDirect3DStateBlock9
}

func (obj StateBlock) Release() {
	C.IDirect3DStateBlock9Release(obj.handle)
}

func (obj StateBlock) Apply() (err error) {
	err = toErr(C.IDirect3DStateBlock9Apply(obj.handle))
	return
}

func (obj StateBlock) Capture() (err error) {
	err = toErr(C.IDirect3DStateBlock9Capture(obj.handle))
	return
}

func (obj StateBlock) GetDevice() (ppDevice Device, err error) {
	var c_ppDevice *C.IDirect3DDevice9
	err = toErr(C.IDirect3DStateBlock9GetDevice(obj.handle, &c_ppDevice))
	ppDevice = Device{c_ppDevice}
	return
}
