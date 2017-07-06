package d3d9

import (
	"syscall"
	"unsafe"
)

// VertexShader and its methods are used to encapsulate the functionality of a
// vertex shader.
type VertexShader struct {
	vtbl *vertexShaderVtbl
}

type vertexShaderVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	GetDevice   uintptr
	GetFunction uintptr
}

// AddRef increments the reference count for an interface on an object. This
// method should be called for every new copy of a pointer to an interface on an
// object.
func (obj *VertexShader) AddRef() uint32 {
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
func (obj *VertexShader) Release() uint32 {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint32(ret)
}

// GetDevice retrieves the associated device.
// Call Release on the returned device when finished using it.
func (obj *VertexShader) GetDevice() (device *Device, err Error) {
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

// GetFunction returns the shader data.
func (obj *VertexShader) GetFunction() (data []byte, err Error) {
	// first get the needed buffer size, pass nil as the data pointer
	var sizeInBytes uint
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetFunction,
		3,
		uintptr(unsafe.Pointer(obj)),
		0,
		uintptr(unsafe.Pointer(&sizeInBytes)),
	)
	if err := toErr(ret); err != nil {
		return nil, err
	}

	data = make([]byte, sizeInBytes)
	ret, _, _ = syscall.Syscall(
		obj.vtbl.GetFunction,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(unsafe.Pointer(&sizeInBytes)),
	)
	err = toErr(ret)
	return
}
