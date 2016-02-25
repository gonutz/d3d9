package d3d9

/*
#include "d3d9wrapper.h"

HRESULT IDirect3DResource9FreePrivateData(
		IDirect3DResource9* obj,
		GUID* refguid) {
	return obj->lpVtbl->FreePrivateData(obj, refguid);
}

HRESULT IDirect3DResource9GetDevice(
		IDirect3DResource9* obj,
		IDirect3DDevice9** ppDevice) {
	return obj->lpVtbl->GetDevice(obj, ppDevice);
}

DWORD IDirect3DResource9GetPriority(IDirect3DResource9* obj) {
	return obj->lpVtbl->GetPriority(obj);
}

HRESULT IDirect3DResource9GetPrivateData(
		IDirect3DResource9* obj,
		GUID* refguid,
		void* pData,
		DWORD* pSizeOfData) {
	return obj->lpVtbl->GetPrivateData(obj, refguid, pData, pSizeOfData);
}

D3DRESOURCETYPE IDirect3DResource9GetType(IDirect3DResource9* obj) {
	return obj->lpVtbl->GetType(obj);
}

void IDirect3DResource9PreLoad(IDirect3DResource9* obj) {
	obj->lpVtbl->PreLoad(obj);
}

DWORD IDirect3DResource9SetPriority(
		IDirect3DResource9* obj,
		DWORD PriorityNew) {
	return obj->lpVtbl->SetPriority(obj, PriorityNew);
}

HRESULT IDirect3DResource9SetPrivateData(
		IDirect3DResource9* obj,
		GUID* refguid,
		void* pData,
		DWORD SizeOfData,
		DWORD Flags) {
	return obj->lpVtbl->SetPrivateData(obj, refguid, pData, SizeOfData, Flags);
}

void IDirect3DResource9Release(IDirect3DResource9* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"
import "unsafe"

// Resource and its methods are used to query and prepare resources.
type Resource struct {
	handle *C.IDirect3DResource9
}

// Release has to be called when finished using the object to free its
// associated resources.
func (obj Resource) Release() {
	C.IDirect3DResource9Release(obj.handle)
}

// FreePrivateData frees the specified private data associated with this
// resource.
func (obj Resource) FreePrivateData(refguid GUID) (err Error) {
	crefguid := refguid.toC()
	err = toErr(C.IDirect3DResource9FreePrivateData(obj.handle, &crefguid))
	return
}

// GetDevice retrieves the device associated with a resource.
// Call Release on the returned device when finished using it.
func (obj Resource) GetDevice() (ppDevice Device, err Error) {
	var cppDevice *C.IDirect3DDevice9
	err = toErr(C.IDirect3DResource9GetDevice(obj.handle, &cppDevice))
	ppDevice = Device{cppDevice}
	return
}

// GetPriority retrieves the priority for this resource.
func (obj Resource) GetPriority() uint32 {
	return (uint32)(C.IDirect3DResource9GetPriority(obj.handle))
}

// GetPrivateData copies the private data associated with the resource to a
// provided buffer.
func (obj Resource) GetPrivateData(refguid GUID) (data []byte, err Error) {
	// first get the data size by passing nil as the data pointer
	crefguid := refguid.toC()
	var cpSizeOfData C.DWORD
	err = toErr(C.IDirect3DResource9GetPrivateData(
		obj.handle,
		&crefguid,
		nil,
		&cpSizeOfData,
	))
	if err != nil {
		return
	}
	data = make([]byte, cpSizeOfData)
	err = toErr(C.IDirect3DResource9GetPrivateData(
		obj.handle,
		&crefguid,
		unsafe.Pointer(&data[0]),
		&cpSizeOfData,
	))
	return
}

// GetType returns the type of the resource.
func (obj Resource) GetType() RESOURCETYPE {
	return (RESOURCETYPE)(C.IDirect3DResource9GetType(obj.handle))
}

// PreLoad preloads a managed resource.
func (obj Resource) PreLoad() {
	C.IDirect3DResource9PreLoad(obj.handle)
}

// SetPriority assigns the priority of a resource for scheduling purposes.
func (obj Resource) SetPriority(PriorityNew uint32) uint32 {
	return (uint32)(C.IDirect3DResource9SetPriority(
		obj.handle,
		(C.DWORD)(PriorityNew),
	))
}

// SetPrivateData associates data with the resource that is intended for use by
// the application, not by Direct3D. Data is passed by value, and multiple sets
// of data can be associated with a single resource.
func (obj Resource) SetPrivateData(
	refguid GUID,
	pData unsafe.Pointer,
	SizeOfData uint32,
	Flags uint32,
) (
	err Error,
) {
	crefguid := refguid.toC()
	err = toErr(C.IDirect3DResource9SetPrivateData(
		obj.handle,
		&crefguid,
		pData,
		(C.DWORD)(SizeOfData),
		(C.DWORD)(Flags),
	))
	return
}

// SetPrivateDataBytes associates data with the resource that is intended for
// use by the application, not by Direct3D. Data is passed by value, and
// multiple sets of data can be associated with a single resource.
func (obj Resource) SetPrivateDataBytes(
	refguid GUID,
	data []byte,
	Flags uint32,
) (
	err Error,
) {
	return obj.SetPrivateData(
		refguid,
		unsafe.Pointer(&data[0]),
		uint32(len(data)),
		Flags,
	)
}
