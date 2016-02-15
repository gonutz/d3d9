package d3d9

/*
#include "d3d9wrapper.h"

HRESULT IDirect3DVolumeTexture9AddDirtyBox(
		IDirect3DVolumeTexture9* obj,
		D3DBOX* pDirtyBox) {
	return obj->lpVtbl->AddDirtyBox(obj, pDirtyBox);
}

HRESULT IDirect3DVolumeTexture9GetLevelDesc(
		IDirect3DVolumeTexture9* obj,
		UINT Level, D3DVOLUME_DESC* pDesc) {
	return obj->lpVtbl->GetLevelDesc(obj, Level, pDesc);
}

HRESULT IDirect3DVolumeTexture9GetVolumeLevel(
		IDirect3DVolumeTexture9* obj,
		UINT Level,
		IDirect3DVolume9** ppVolumeLevel) {
	return obj->lpVtbl->GetVolumeLevel(obj, Level, ppVolumeLevel);
}

HRESULT IDirect3DVolumeTexture9LockBox(
		IDirect3DVolumeTexture9* obj,
		UINT Level,
		D3DLOCKED_BOX* pLockedVolume,
		D3DBOX* pBox,
		DWORD Flags) {
	return obj->lpVtbl->LockBox(obj, Level, pLockedVolume, pBox, Flags);
}

HRESULT IDirect3DVolumeTexture9UnlockBox(
		IDirect3DVolumeTexture9* obj,
		UINT Level) {
	return obj->lpVtbl->UnlockBox(obj, Level);
}

void IDirect3DVolumeTexture9Release(IDirect3DVolumeTexture9* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"

type VolumeTexture struct {
	BaseTexture
	handle *C.IDirect3DVolumeTexture9
}

func (obj VolumeTexture) Release() {
	C.IDirect3DVolumeTexture9Release(obj.handle)
}

// AddDirtyBox adds a dirty region to a volume texture resource.
func (obj VolumeTexture) AddDirtyBox(pDirtyBox *BOX) (err error) {
	if pDirtyBox == nil {
		err = toErr(C.IDirect3DVolumeTexture9AddDirtyBox(obj.handle, nil))
	} else {
		c_pDirtyBox := pDirtyBox.toC()
		err = toErr(C.IDirect3DVolumeTexture9AddDirtyBox(
			obj.handle,
			&c_pDirtyBox,
		))
	}
	return
}

// GetLevelDesc retrieves a level description of a volume texture resource.
func (obj VolumeTexture) GetLevelDesc(
	Level uint,
) (
	pDesc VOLUME_DESC,
	err error,
) {
	var c_pDesc C.D3DVOLUME_DESC
	err = toErr(C.IDirect3DVolumeTexture9GetLevelDesc(
		obj.handle,
		(C.UINT)(Level),
		&c_pDesc,
	))
	pDesc.fromC(&c_pDesc)
	return
}

// GetVolumeLevel retrieves the specified volume texture level.
// Call Release on the returned volume when finished using it.
func (obj VolumeTexture) GetVolumeLevel(
	Level uint,
) (
	ppVolumeLevel Volume,
	err error,
) {
	var c_ppVolumeLevel *C.IDirect3DVolume9
	err = toErr(C.IDirect3DVolumeTexture9GetVolumeLevel(
		obj.handle,
		(C.UINT)(Level),
		&c_ppVolumeLevel,
	))
	ppVolumeLevel = Volume{c_ppVolumeLevel}
	return
}

// LockBox locks a box on a volume texture resource.
func (obj VolumeTexture) LockBox(
	Level uint,
	pBox *BOX,
	Flags uint32,
) (
	pLockedVolume LOCKED_BOX,
	err error,
) {
	var c_pLockedVolume C.D3DLOCKED_BOX
	if pBox == nil {
		err = toErr(C.IDirect3DVolumeTexture9LockBox(
			obj.handle,
			(C.UINT)(Level),
			&c_pLockedVolume,
			nil,
			(C.DWORD)(Flags),
		))
	} else {
		c_pBox := pBox.toC()
		err = toErr(C.IDirect3DVolumeTexture9LockBox(
			obj.handle,
			(C.UINT)(Level),
			&c_pLockedVolume,
			&c_pBox,
			(C.DWORD)(Flags),
		))
	}
	pLockedVolume.fromC(&c_pLockedVolume)
	return
}

// UnlockBox unlocks a box on a volume texture resource.
func (obj VolumeTexture) UnlockBox(Level uint) (err error) {
	err = toErr(C.IDirect3DVolumeTexture9UnlockBox(obj.handle, (C.UINT)(Level)))
	return
}
