package d3d9

import (
	"syscall"
	"unsafe"
)

// Resource and its methods are used to query and prepare resources.
type Resource struct {
	vtbl *resourceVtbl
}

type resourceVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	GetDevice       uintptr
	SetPrivateData  uintptr
	GetPrivateData  uintptr
	FreePrivateData uintptr
	SetPriority     uintptr
	GetPriority     uintptr
	PreLoad         uintptr
	GetType         uintptr
}

// AddRef increments the reference count for an interface on an object. This
// method should be called for every new copy of a pointer to an interface on an
// object.
func (obj *Resource) AddRef() uint32 {
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
func (obj *Resource) Release() uint32 {
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
func (obj *Resource) GetDevice() (device *Device, err Error) {
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
func (obj *Resource) SetPrivateData(
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
func (obj *Resource) SetPrivateDataBytes(
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
func (obj *Resource) GetPrivateData(refguid GUID) (data []byte, err Error) {
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
func (obj *Resource) FreePrivateData(refguid GUID) Error {
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
func (obj *Resource) SetPriority(priorityNew uint32) uint32 {
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
func (obj *Resource) GetPriority() uint32 {
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
func (obj *Resource) PreLoad() {
	syscall.Syscall(
		obj.vtbl.PreLoad,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
}

// GetType returns the type of the resource.
func (obj *Resource) GetType() RESOURCETYPE {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetType,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return RESOURCETYPE(ret)
}
