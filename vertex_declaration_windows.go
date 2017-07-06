package d3d9

import (
	"syscall"
	"unsafe"
)

// VertexDeclaration and its methods are used to encapsulate the vertex shader
// declaration.
type VertexDeclaration struct {
	vtbl *vertexDeclarationVtbl
}

type vertexDeclarationVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	GetDevice      uintptr
	GetDeclaration uintptr
}

// AddRef increments the reference count for an interface on an object. This
// method should be called for every new copy of a pointer to an interface on an
// object.
func (obj *VertexDeclaration) AddRef() uint32 {
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
func (obj *VertexDeclaration) Release() uint32 {
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
func (obj *VertexDeclaration) GetDevice() (device *Device, err Error) {
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

// GetDeclaration returns the vertex shader declaration.
func (obj *VertexDeclaration) GetDeclaration() (decl []VERTEXELEMENT, err Error) {
	// first pass nil for the elements to get the count
	var elemCount uint
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetDeclaration,
		3,
		uintptr(unsafe.Pointer(obj)),
		0,
		uintptr(unsafe.Pointer(&elemCount)),
	)
	if err := toErr(ret); err != nil {
		return nil, err
	}

	decl = make([]VERTEXELEMENT, elemCount)
	ret, _, _ = syscall.Syscall(
		obj.vtbl.GetDeclaration,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&decl[0])),
		uintptr(unsafe.Pointer(&elemCount)),
	)
	err = toErr(ret)
	return
}
