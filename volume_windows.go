package d3d9

import (
	"syscall"
	"unsafe"
)

// Volume and its methods are used to manipulate volume resources.
type Volume struct {
	vtbl *volumeVtbl
}

type volumeVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	GetDevice       uintptr
	SetPrivateData  uintptr
	GetPrivateData  uintptr
	FreePrivateData uintptr
	GetContainer    uintptr
	GetDesc         uintptr
	LockBox         uintptr
	UnlockBox       uintptr
}

// AddRef increments the reference count for an interface on an object. This
// method should be called for every new copy of a pointer to an interface on an
// object.
func (obj *Volume) AddRef() uint32 {
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
func (obj *Volume) Release() uint32 {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint32(ret)
}

// GetDevice retrieves the device associated with a volume.
func (obj *Volume) GetDevice() (device *Device, err Error) {
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

// SetPrivateData associates data with the volume that is intended for use by
// the application, not by Direct3D.
func (obj *Volume) SetPrivateData(
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

// SetPrivateDataInBytes associates data with the volume that is intended for
// use by the application, not by Direct3D.
func (obj Volume) SetPrivateDataInBytes(
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

// GetPrivateData copies the private data associated with the volume to a
// buffer.
func (obj *Volume) GetPrivateData(refguid GUID) (data []byte, err Error) {
	// first get the buffer size by passing nil as the data pointer
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
	err = toErr(ret)
	return
}

// FreePrivateData frees the specified private data associated with this volume.
func (obj *Volume) FreePrivateData(refguid GUID) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.FreePrivateData,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&refguid)),
		0,
	)
	return toErr(ret)
}

// GetContainer provides access to the parent volume texture object, if this
// surface is a child level of a volume texture.
// Call Release on the returned volume texture when finished using it.
func (obj *Volume) GetContainer(riid GUID) (tex *VolumeTexture, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetContainer,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&riid)),
		uintptr(unsafe.Pointer(tex)),
	)
	err = toErr(ret)
	return
}

// GetDesc retrieves a description of the volume.
func (obj *Volume) GetDesc() (desc VOLUME_DESC, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetDesc,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&desc)),
		0,
	)
	err = toErr(ret)
	return
}

// LockBox locks a box on a volume resource.
func (obj *Volume) LockBox(
	box *BOX,
	flags uint32,
) (locked *LOCKED_BOX, err Error) {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.LockBox,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&locked)),
		uintptr(unsafe.Pointer(&box)),
		uintptr(flags),
		0,
		0,
	)
	err = toErr(ret)
	return
}

// UnlockBox unlocks a box on a volume resource.
func (obj *Volume) UnlockBox() Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.UnlockBox,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return toErr(ret)
}
