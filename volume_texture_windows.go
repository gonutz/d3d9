package d3d9

import (
	"syscall"
	"unsafe"
)

// VolumeTexture and its methods are used to manipulate a volume texture
// resource.
type VolumeTexture struct {
	vtbl *volumeTextureVtbl
}

type volumeTextureVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	GetDevice                uintptr
	SetPrivateData           uintptr
	GetPrivateData           uintptr
	FreePrivateData          uintptr
	SetPriority              uintptr
	GetPriority              uintptr
	PreLoad                  uintptr
	D3DRESOURCETYPE, GetType uintptr
	SetLOD                   uintptr
	GetLOD                   uintptr
	GetLevelCount            uintptr
	SetAutoGenFilterType     uintptr
	GetAutoGenFilterType     uintptr
	GenerateMipSubLevels     uintptr
	GetLevelDesc             uintptr
	GetVolumeLevel           uintptr
	LockBox                  uintptr
	UnlockBox                uintptr
	AddDirtyBox              uintptr
}

func (obj *VolumeTexture) baseTexturePointer() uintptr {
	return uintptr(unsafe.Pointer(obj))
}

// AddRef increments the reference count for an interface on an object. This
// method should be called for every new copy of a pointer to an interface on an
// object.
func (obj *VolumeTexture) AddRef() uint32 {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.AddRef,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint32(ret)
}

// Release has to be called when finished using the object to free its
// associated resources.
func (obj *VolumeTexture) Release() uint32 {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint32(ret)
}

// GetDevice retrieves the device associated with a resource.
// Call Release on the returned device when finished using it.
func (obj *VolumeTexture) GetDevice() (device *Device, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetDevice,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&device)),
		0,
	)
	err = toErr(ret)
	return
}

// SetPrivateData associates data with the resource that is intended for use by
// the application, not by Direct3D. Data is passed by value, and multiple sets
// of data can be associated with a single resource.
func (obj *VolumeTexture) SetPrivateData(
	refguid GUID,
	data uintptr,
	sizeOfData uint32,
	flags uint32,
) Error {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.SetPrivateData,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&refguid)),
		data,
		uintptr(sizeOfData),
		uintptr(flags),
		0,
	)
	return toErr(ret)
}

// SetPrivateDataBytes associates data with the resource that is intended for
// use by the application, not by Direct3D. Data is passed by value, and
// multiple sets of data can be associated with a single resource.
func (obj *VolumeTexture) SetPrivateDataBytes(
	refguid GUID,
	data []byte,
	flags uint32,
) Error {
	return obj.SetPrivateData(
		refguid,
		uintptr(unsafe.Pointer(&data[0])),
		uint32(len(data)),
		flags,
	)
}

// GetPrivateData copies the private data associated with the resource to a
// provided buffer.
func (obj *VolumeTexture) GetPrivateData(refguid GUID) (data []byte, err Error) {
	// first get the data size by passing nil as the data pointer
	var sizeInBytes uint
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.GetPrivateData,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&refguid)),
		0,
		uintptr(unsafe.Pointer(&sizeInBytes)),
		0,
		0,
	)
	if err := toErr(ret); err != nil {
		return nil, err
	}
	data = make([]byte, sizeInBytes)
	ret, _, _ = syscall.Syscall6(
		obj.vtbl.GetPrivateData,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&refguid)),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(unsafe.Pointer(&sizeInBytes)),
		0,
		0,
	)
	return data, toErr(ret)
}

// FreePrivateData frees the specified private data associated with this
// resource.
func (obj *VolumeTexture) FreePrivateData(refguid GUID) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.FreePrivateData,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&refguid)),
		0,
	)
	return toErr(ret)
}

// SetPriority assigns the priority of a resource for scheduling purposes.
func (obj *VolumeTexture) SetPriority(priorityNew uint32) uint32 {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.SetPriority,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(priorityNew),
		0,
	)
	return uint32(ret)
}

// GetPriority retrieves the priority for this resource.
func (obj *VolumeTexture) GetPriority() uint32 {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetPriority,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint32(ret)
}

// PreLoad preloads a managed resource.
func (obj *VolumeTexture) PreLoad() {
	syscall.Syscall(
		obj.vtbl.PreLoad,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
}

// GetType returns the type of the resource.
func (obj *VolumeTexture) GetType() RESOURCETYPE {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetType,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return RESOURCETYPE(ret)
}

// SetLOD sets the most detailed level-of-detail for a managed texture.
func (obj *VolumeTexture) SetLOD(lodNew uint32) uint32 {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.SetLOD,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(lodNew),
		0,
	)
	return uint32(ret)
}

// GetLOD returns a value clamped to the maximum level-of-detail set for a
// managed texture (this method is not supported for an unmanaged texture).
func (obj *VolumeTexture) GetLOD() uint32 {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetLOD,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint32(ret)
}

// GetLevelCount returns the number of texture levels in a multilevel texture.
func (obj *VolumeTexture) GetLevelCount() uint32 {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetLevelCount,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint32(ret)
}

// SetAutoGenFilterType sets the filter type that is used for automatically
// generated mipmap sublevels.
func (obj *VolumeTexture) SetAutoGenFilterType(typ TEXTUREFILTERTYPE) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.SetAutoGenFilterType,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(typ),
		0,
	)
	return toErr(ret)
}

// GetAutoGenFilterType returns the filter type that is used for automatically
// generated mipmap sublevels.
func (obj *VolumeTexture) GetAutoGenFilterType() TEXTUREFILTERTYPE {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetAutoGenFilterType,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return TEXTUREFILTERTYPE(ret)
}

// GenerateMipSubLevels generates mip sub levels.
func (obj *VolumeTexture) GenerateMipSubLevels() {
	syscall.Syscall(
		obj.vtbl.GenerateMipSubLevels,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
}

// GetLevelDesc retrieves a level description of a volume texture resource.
func (obj *VolumeTexture) GetLevelDesc(level uint) (desc VOLUME_DESC, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetLevelDesc,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(level),
		uintptr(unsafe.Pointer(&desc)),
	)
	err = toErr(ret)
	return
}

// GetVolumeLevel retrieves the specified volume texture level.
// Call Release on the returned volume when finished using it.
func (obj *VolumeTexture) GetVolumeLevel(level uint) (volume *Volume, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetVolumeLevel,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(level),
		uintptr(unsafe.Pointer(&volume)),
	)
	err = toErr(ret)
	return
}

// LockBox locks a box on a volume texture resource.
func (obj *VolumeTexture) LockBox(
	level uint,
	box *BOX,
	flags uint32,
) (lockedVolume LOCKED_BOX, err Error) {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.LockBox,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(level),
		uintptr(unsafe.Pointer(&lockedVolume)),
		uintptr(unsafe.Pointer(box)),
		uintptr(flags),
		0,
	)
	err = toErr(ret)
	return
}

// UnlockBox unlocks a box on a volume texture resource.
func (obj *VolumeTexture) UnlockBox(level uint) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.UnlockBox,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(level),
		0,
	)
	return toErr(ret)
}

// AddDirtyBox adds a dirty region to a volume texture resource.
func (obj *VolumeTexture) AddDirtyBox(dirtyBox *BOX) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.AddDirtyBox,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(dirtyBox)),
		0,
	)
	return toErr(ret)
}
