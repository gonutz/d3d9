package d3d9

/*
#include "d3d9wrapper.h"

HRESULT IDirect3DVolume9FreePrivateData(IDirect3DVolume9* obj, GUID* refguid) {
	return obj->lpVtbl->FreePrivateData(obj, refguid);
}

HRESULT IDirect3DVolume9GetContainer(
		IDirect3DVolume9* obj,
		GUID* riid,
		IDirect3DVolumeTexture9** ppContainer) {
	return obj->lpVtbl->GetContainer(obj, riid, (void**)ppContainer);
}

HRESULT IDirect3DVolume9GetDesc(IDirect3DVolume9* obj, D3DVOLUME_DESC* pDesc) {
	return obj->lpVtbl->GetDesc(obj, pDesc);
}

HRESULT IDirect3DVolume9GetDevice(
		IDirect3DVolume9* obj,
		IDirect3DDevice9** ppDevice) {
	return obj->lpVtbl->GetDevice(obj, ppDevice);
}

HRESULT IDirect3DVolume9GetPrivateData(
		IDirect3DVolume9* obj,
		GUID* refguid,
		void* pData,
		DWORD* pSizeOfData) {
	return obj->lpVtbl->GetPrivateData(obj, refguid, pData, pSizeOfData);
}

HRESULT IDirect3DVolume9LockBox(
		IDirect3DVolume9* obj,
		D3DLOCKED_BOX* pLockedVolume,
		D3DBOX* pBox,
		DWORD Flags) {
	return obj->lpVtbl->LockBox(obj, pLockedVolume, pBox, Flags);
}

HRESULT IDirect3DVolume9SetPrivateData(
		IDirect3DVolume9* obj,
		GUID* refguid,
		void* pData,
		DWORD SizeOfData,
		DWORD Flags) {
	return obj->lpVtbl->SetPrivateData(obj, refguid, pData, SizeOfData, Flags);
}

HRESULT IDirect3DVolume9UnlockBox(IDirect3DVolume9* obj) {
	return obj->lpVtbl->UnlockBox(obj);
}

void IDirect3DVolume9Release(IDirect3DVolume9* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"
import "unsafe"

// Volume and its methods are used to manipulate volume resources.
type Volume struct {
	handle *C.IDirect3DVolume9
}

// Release has to be called when finished using the object to free its
// associated resources.
func (obj Volume) Release() {
	C.IDirect3DVolume9Release(obj.handle)
}

// FreePrivateData frees the specified private data associated with this volume.
func (obj Volume) FreePrivateData(refguid GUID) (err Error) {
	crefguid := refguid.toC()
	err = toErr(C.IDirect3DVolume9FreePrivateData(obj.handle, &crefguid))
	return
}

// GetContainer provides access to the parent volume texture object, if this
// surface is a child level of a volume texture.
// Call Release on the returned volume texture when finished using it.
func (obj Volume) GetContainer(
	riid GUID,
) (
	ppContainer VolumeTexture,
	err Error,
) {
	var handle *C.IDirect3DVolumeTexture9
	criid := riid.toC()
	err = toErr(C.IDirect3DVolume9GetContainer(
		obj.handle,
		&criid,
		&handle,
	))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(handle))}
	baseTexture := BaseTexture{
		resource,
		(*C.IDirect3DBaseTexture9)(unsafe.Pointer(handle)),
	}
	ppContainer = VolumeTexture{
		baseTexture,
		(*C.IDirect3DVolumeTexture9)(handle),
	}
	return
}

// GetDesc retrieves a description of the volume.
func (obj Volume) GetDesc() (pDesc VOLUME_DESC, err Error) {
	var cpDesc C.D3DVOLUME_DESC
	err = toErr(C.IDirect3DVolume9GetDesc(obj.handle, &cpDesc))
	pDesc.fromC(&cpDesc)
	return
}

// GetDevice retrieves the device associated with a volume.
func (obj Volume) GetDevice() (ppDevice Device, err Error) {
	var cppDevice *C.IDirect3DDevice9
	err = toErr(C.IDirect3DVolume9GetDevice(obj.handle, &cppDevice))
	ppDevice = Device{cppDevice}
	return
}

// GetPrivateData copies the private data associated with the volume to a
// provided buffer.
func (obj Volume) GetPrivateData(refguid GUID) (pData []byte, err Error) {
	// first get the buffer size by passing nil as the data pointer
	crefguid := refguid.toC()
	var cpSizeOfDataInBytes C.DWORD
	err = toErr(C.IDirect3DVolume9GetPrivateData(
		obj.handle,
		&crefguid,
		nil,
		&cpSizeOfDataInBytes,
	))
	if err != nil {
		return
	}
	pData = make([]byte, cpSizeOfDataInBytes)
	err = toErr(C.IDirect3DVolume9GetPrivateData(
		obj.handle,
		&crefguid,
		unsafe.Pointer(&pData[0]),
		&cpSizeOfDataInBytes,
	))
	return
}

// LockBox locks a box on a volume resource.
func (obj Volume) LockBox(
	pBox *BOX,
	Flags uint32,
) (
	pLockedVolume LOCKED_BOX,
	err Error,
) {
	var cpLockedVolume C.D3DLOCKED_BOX
	if pBox == nil {
		err = toErr(C.IDirect3DVolume9LockBox(
			obj.handle,
			&cpLockedVolume,
			nil,
			(C.DWORD)(Flags),
		))
	} else {
		cpBox := pBox.toC()
		err = toErr(C.IDirect3DVolume9LockBox(
			obj.handle,
			&cpLockedVolume,
			&cpBox,
			(C.DWORD)(Flags),
		))
	}
	pLockedVolume.fromC(&cpLockedVolume)
	return
}

// SetPrivateData associates data with the volume that is intended for use by
// the application, not by Direct3D.
func (obj Volume) SetPrivateData(
	refguid GUID,
	pData unsafe.Pointer,
	SizeOfData uint32,
	Flags uint32,
) (
	err Error,
) {
	crefguid := refguid.toC()
	err = toErr(C.IDirect3DVolume9SetPrivateData(
		obj.handle,
		&crefguid,
		pData,
		(C.DWORD)(SizeOfData),
		(C.DWORD)(Flags),
	))
	return
}

// SetPrivateDataInBytes associates data with the volume that is intended for
// use by the application, not by Direct3D.
func (obj Volume) SetPrivateDataInBytes(
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

// UnlockBox unlocks a box on a volume resource.
func (obj Volume) UnlockBox() (err Error) {
	err = toErr(C.IDirect3DVolume9UnlockBox(obj.handle))
	return
}
