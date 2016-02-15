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

type Volume struct {
	handle *C.IDirect3DVolume9
}

func (obj Volume) Release() {
	C.IDirect3DVolume9Release(obj.handle)
}

// FreePrivateData frees the specified private data associated with this volume.
func (obj Volume) FreePrivateData(refguid GUID) (err error) {
	c_refguid := refguid.toC()
	err = toErr(C.IDirect3DVolume9FreePrivateData(obj.handle, &c_refguid))
	return
}

// GetContainer provides access to the parent volume texture object, if this
// surface is a child level of a volume texture.
// Call Release on the returned volume texture when finished using it.
func (obj Volume) GetContainer(
	riid GUID,
) (
	ppContainer VolumeTexture,
	err error,
) {
	var handle *C.IDirect3DVolumeTexture9
	c_riid := riid.toC()
	err = toErr(C.IDirect3DVolume9GetContainer(
		obj.handle,
		&c_riid,
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
func (obj Volume) GetDesc() (pDesc VOLUME_DESC, err error) {
	var c_pDesc C.D3DVOLUME_DESC
	err = toErr(C.IDirect3DVolume9GetDesc(obj.handle, &c_pDesc))
	pDesc.fromC(&c_pDesc)
	return
}

// GetDevice retrieves the device associated with a volume.
func (obj Volume) GetDevice() (ppDevice Device, err error) {
	var c_ppDevice *C.IDirect3DDevice9
	err = toErr(C.IDirect3DVolume9GetDevice(obj.handle, &c_ppDevice))
	ppDevice = Device{c_ppDevice}
	return
}

// GetPrivateData copies the private data associated with the volume to a
// provided buffer.
func (obj Volume) GetPrivateData(refguid GUID) (pData []byte, err error) {
	// first get the buffer size by passing nil as the data pointer
	c_refguid := refguid.toC()
	var c_pSizeOfDataInBytes C.DWORD
	err = toErr(C.IDirect3DVolume9GetPrivateData(
		obj.handle,
		&c_refguid,
		nil,
		&c_pSizeOfDataInBytes,
	))
	if err != nil {
		return
	}
	pData = make([]byte, c_pSizeOfDataInBytes)
	err = toErr(C.IDirect3DVolume9GetPrivateData(
		obj.handle,
		&c_refguid,
		unsafe.Pointer(&pData[0]),
		&c_pSizeOfDataInBytes,
	))
	return
}

// LockBox locks a box on a volume resource.
func (obj Volume) LockBox(
	pBox *BOX,
	Flags uint32,
) (
	pLockedVolume LOCKED_BOX,
	err error,
) {
	var c_pLockedVolume C.D3DLOCKED_BOX
	if pBox == nil {
		err = toErr(C.IDirect3DVolume9LockBox(
			obj.handle,
			&c_pLockedVolume,
			nil,
			(C.DWORD)(Flags),
		))
	} else {
		c_pBox := pBox.toC()
		err = toErr(C.IDirect3DVolume9LockBox(
			obj.handle,
			&c_pLockedVolume,
			&c_pBox,
			(C.DWORD)(Flags),
		))
	}
	pLockedVolume.fromC(&c_pLockedVolume)
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
	err error,
) {
	c_refguid := refguid.toC()
	err = toErr(C.IDirect3DVolume9SetPrivateData(
		obj.handle,
		&c_refguid,
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
	err error,
) {
	return obj.SetPrivateData(
		refguid,
		unsafe.Pointer(&data[0]),
		uint32(len(data)),
		Flags,
	)
}

// UnlockBox unlocks a box on a volume resource.
func (obj Volume) UnlockBox() (err error) {
	err = toErr(C.IDirect3DVolume9UnlockBox(obj.handle))
	return
}
