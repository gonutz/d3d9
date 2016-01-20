package d3d9

/*
#include "d3d9wrapper.h"

HRESULT IDirect3DResource9FreePrivateData(IDirect3DResource9* obj, GUID* refguid) {
	return obj->lpVtbl->FreePrivateData(obj, refguid);
}

HRESULT IDirect3DResource9GetDevice(IDirect3DResource9* obj, IDirect3DDevice9** ppDevice) {
	return obj->lpVtbl->GetDevice(obj, ppDevice);
}

DWORD IDirect3DResource9GetPriority(IDirect3DResource9* obj) {
	return obj->lpVtbl->GetPriority(obj);
}

HRESULT IDirect3DResource9GetPrivateData(IDirect3DResource9* obj, GUID* refguid, void* pData, DWORD* pSizeOfData) {
	return obj->lpVtbl->GetPrivateData(obj, refguid, pData, pSizeOfData);
}

D3DRESOURCETYPE IDirect3DResource9GetType(IDirect3DResource9* obj) {
	return obj->lpVtbl->GetType(obj);
}

void IDirect3DResource9PreLoad(IDirect3DResource9* obj) {
	obj->lpVtbl->PreLoad(obj);
}

DWORD IDirect3DResource9SetPriority(IDirect3DResource9* obj, DWORD PriorityNew) {
	return obj->lpVtbl->SetPriority(obj, PriorityNew);
}

HRESULT IDirect3DResource9SetPrivateData(IDirect3DResource9* obj, GUID* refguid, void* pData, DWORD SizeOfData, DWORD Flags) {
	return obj->lpVtbl->SetPrivateData(obj, refguid, pData, SizeOfData, Flags);
}

void IDirect3DResource9Release(IDirect3DResource9* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"
import "unsafe"

type Resource struct {
	handle *C.IDirect3DResource9
}

func (obj Resource) Release() {
	C.IDirect3DResource9Release(obj.handle)
}

func (obj Resource) FreePrivateData(refguid GUID) (err error) {
	c_refguid := refguid.toC()
	err = toErr(C.IDirect3DResource9FreePrivateData(obj.handle, &c_refguid))
	return
}

func (obj Resource) GetDevice() (ppDevice Device, err error) {
	var c_ppDevice *C.IDirect3DDevice9
	err = toErr(C.IDirect3DResource9GetDevice(obj.handle, &c_ppDevice))
	ppDevice = Device{c_ppDevice}
	return
}

func (obj Resource) GetPriority() uint32 {
	return (uint32)(C.IDirect3DResource9GetPriority(obj.handle))
}

func (obj Resource) GetPrivateData(refguid GUID) (data []byte, err error) {
	c_refguid := refguid.toC()
	var c_pSizeOfData C.DWORD
	err = toErr(C.IDirect3DResource9GetPrivateData(obj.handle, &c_refguid, nil, &c_pSizeOfData))
	if err != nil {
		return
	}
	data = make([]byte, c_pSizeOfData)
	err = toErr(C.IDirect3DResource9GetPrivateData(obj.handle, &c_refguid, unsafe.Pointer(&data[0]), &c_pSizeOfData))
	return
}

func (obj Resource) GetType() RESOURCETYPE {
	return (RESOURCETYPE)(C.IDirect3DResource9GetType(obj.handle))
}

func (obj Resource) PreLoad() {
	C.IDirect3DResource9PreLoad(obj.handle)
}

func (obj Resource) SetPriority(PriorityNew uint32) uint32 {
	return (uint32)(C.IDirect3DResource9SetPriority(obj.handle, (C.DWORD)(PriorityNew)))
}

func (obj Resource) SetPrivateData(refguid GUID, pData unsafe.Pointer, SizeOfData uint32, Flags uint32) (err error) {
	c_refguid := refguid.toC()
	err = toErr(C.IDirect3DResource9SetPrivateData(obj.handle, &c_refguid, pData, (C.DWORD)(SizeOfData), (C.DWORD)(Flags)))
	return
}
