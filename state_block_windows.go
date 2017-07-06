package d3d9

import (
	"syscall"
	"unsafe"
)

// StateBlock and its methods are used to encapsulate render states.
type StateBlock struct {
	vtbl *stateBlockVtbl
}

type stateBlockVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	GetDevice uintptr
	Capture   uintptr
	Apply     uintptr
}

// AddRef increments the reference count for an interface on an object. This
// method should be called for every new copy of a pointer to an interface on an
// object.
func (obj *StateBlock) AddRef() uint32 {
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
func (obj *StateBlock) Release() uint32 {
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
func (obj *StateBlock) GetDevice() (device *Device, err Error) {
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

// Capture captures the current value of states that are included in a
// stateblock.
func (obj *StateBlock) Capture() Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Capture,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return toErr(ret)
}

// Apply applies the state block to the current device state.
func (obj *StateBlock) Apply() Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Apply,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return toErr(ret)
}
