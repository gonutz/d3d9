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

// VolumeTexture and its methods are used to manipulate a volume texture
// resource.
type VolumeTexture struct {
	BaseTexture
	handle *C.IDirect3DVolumeTexture9
}

// Release has to be called when finished using the object to free its
// associated resources.
func (obj VolumeTexture) Release() {
	C.IDirect3DVolumeTexture9Release(obj.handle)
}

// AddDirtyBox adds a dirty region to a volume texture resource.
func (obj VolumeTexture) AddDirtyBox(pDirtyBox *BOX) (err Error) {
	if pDirtyBox == nil {
		err = toErr(C.IDirect3DVolumeTexture9AddDirtyBox(obj.handle, nil))
	} else {
		cpDirtyBox := pDirtyBox.toC()
		err = toErr(C.IDirect3DVolumeTexture9AddDirtyBox(
			obj.handle,
			&cpDirtyBox,
		))
	}
	return
}

// GetLevelDesc retrieves a level description of a volume texture resource.
func (obj VolumeTexture) GetLevelDesc(
	Level uint,
) (
	pDesc VOLUME_DESC,
	err Error,
) {
	var cpDesc C.D3DVOLUME_DESC
	err = toErr(C.IDirect3DVolumeTexture9GetLevelDesc(
		obj.handle,
		(C.UINT)(Level),
		&cpDesc,
	))
	pDesc.fromC(&cpDesc)
	return
}

// GetVolumeLevel retrieves the specified volume texture level.
// Call Release on the returned volume when finished using it.
func (obj VolumeTexture) GetVolumeLevel(
	Level uint,
) (
	ppVolumeLevel Volume,
	err Error,
) {
	var cppVolumeLevel *C.IDirect3DVolume9
	err = toErr(C.IDirect3DVolumeTexture9GetVolumeLevel(
		obj.handle,
		(C.UINT)(Level),
		&cppVolumeLevel,
	))
	ppVolumeLevel = Volume{cppVolumeLevel}
	return
}

// LockBox locks a box on a volume texture resource.
func (obj VolumeTexture) LockBox(
	Level uint,
	pBox *BOX,
	Flags uint32,
) (
	pLockedVolume LOCKED_BOX,
	err Error,
) {
	var cpLockedVolume C.D3DLOCKED_BOX
	if pBox == nil {
		err = toErr(C.IDirect3DVolumeTexture9LockBox(
			obj.handle,
			(C.UINT)(Level),
			&cpLockedVolume,
			nil,
			(C.DWORD)(Flags),
		))
	} else {
		cpBox := pBox.toC()
		err = toErr(C.IDirect3DVolumeTexture9LockBox(
			obj.handle,
			(C.UINT)(Level),
			&cpLockedVolume,
			&cpBox,
			(C.DWORD)(Flags),
		))
	}
	pLockedVolume.fromC(&cpLockedVolume)
	return
}

// UnlockBox unlocks a box on a volume texture resource.
func (obj VolumeTexture) UnlockBox(Level uint) (err Error) {
	err = toErr(C.IDirect3DVolumeTexture9UnlockBox(obj.handle, (C.UINT)(Level)))
	return
}
