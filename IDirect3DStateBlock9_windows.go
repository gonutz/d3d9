package d3d9

/*
#include "d3d9wrapper.h"

HRESULT IDirect3DStateBlock9Apply(IDirect3DStateBlock9* obj) {
	return obj->lpVtbl->Apply(obj);
}

HRESULT IDirect3DStateBlock9Capture(IDirect3DStateBlock9* obj) {
	return obj->lpVtbl->Capture(obj);
}

HRESULT IDirect3DStateBlock9GetDevice(
		IDirect3DStateBlock9* obj,
		IDirect3DDevice9** ppDevice) {
	return obj->lpVtbl->GetDevice(obj, ppDevice);
}

void IDirect3DStateBlock9Release(IDirect3DStateBlock9* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"

// StateBlock and its methods are used to encapsulate render states.
type StateBlock struct {
	handle *C.IDirect3DStateBlock9
}

// Release has to be called when finished using the object to free its
// associated resources.
func (obj StateBlock) Release() {
	C.IDirect3DStateBlock9Release(obj.handle)
}

// Apply applies the state block to the current device state.
func (obj StateBlock) Apply() (err Error) {
	err = toErr(C.IDirect3DStateBlock9Apply(obj.handle))
	return
}

// Capture captures the current value of states that are included in a
// stateblock.
func (obj StateBlock) Capture() (err Error) {
	err = toErr(C.IDirect3DStateBlock9Capture(obj.handle))
	return
}

// GetDevice returns the device.
func (obj StateBlock) GetDevice() (ppDevice Device, err Error) {
	var cppDevice *C.IDirect3DDevice9
	err = toErr(C.IDirect3DStateBlock9GetDevice(obj.handle, &cppDevice))
	ppDevice = Device{cppDevice}
	return
}
