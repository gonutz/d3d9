package d3d9

import (
	"syscall"
	"unsafe"
)

// Surface and its methods are used to query and prepare surfaces.
type Surface struct {
	vtbl *surfaceVtbl
}

type surfaceVtbl struct {
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
	GetContainer    uintptr
	GetDesc         uintptr
	LockRect        uintptr
	UnlockRect      uintptr
	GetDC           uintptr
	ReleaseDC       uintptr
}

// AddRef increments the reference count for an interface on an object. This
// method should be called for every new copy of a pointer to an interface on an
// object.
func (obj *Surface) AddRef() uint32 {
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
func (obj *Surface) Release() uint32 {
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
func (obj *Surface) GetDevice() (device *Device, err Error) {
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
func (obj *Surface) SetPrivateData(
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
func (obj *Surface) SetPrivateDataBytes(
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
func (obj *Surface) GetPrivateData(refguid GUID) (data []byte, err Error) {
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
func (obj *Surface) FreePrivateData(refguid GUID) Error {
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
func (obj *Surface) SetPriority(priorityNew uint32) uint32 {
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
func (obj *Surface) GetPriority() uint32 {
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
func (obj *Surface) PreLoad() {
	syscall.Syscall(
		obj.vtbl.PreLoad,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
}

// GetType returns the type of the resource.
func (obj *Surface) GetType() RESOURCETYPE {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetType,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return RESOURCETYPE(ret)
}

func (obj *Surface) getContainer(riid GUID, container uintptr) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetContainer,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&riid)),
		container,
	)
	return toErr(ret)
}

// GetContainerCubeTexture provides access to the parent cube texture.
func (obj *Surface) GetContainerCubeTexture(riid GUID) (tex *CubeTexture, err Error) {
	err = obj.getContainer(riid, uintptr(unsafe.Pointer(&tex)))
	return
}

// GetContainerTexture provides access to the parent texture.
func (obj *Surface) GetContainerTexture(riid GUID) (tex *Texture, err Error) {
	err = obj.getContainer(riid, uintptr(unsafe.Pointer(&tex)))
	return
}

// GetContainerSwapChain provides access to the parent swap chain if this
// surface is a back buffer.
func (obj Surface) GetContainerSwapChain(riid GUID) (chain *SwapChain, err Error) {
	err = obj.getContainer(riid, uintptr(unsafe.Pointer(&chain)))
	return
}

// GetDesc retrieves a description of the surface.
func (obj *Surface) GetDesc() (desc SURFACE_DESC, err Error) {
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

// LockRect locks a rectangle on a surface.
func (obj *Surface) LockRect(
	rect *RECT,
	flags uint32,
) (lockedRect LOCKED_RECT, err Error) {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.LockRect,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&lockedRect)),
		uintptr(unsafe.Pointer(rect)),
		uintptr(flags),
		0,
		0,
	)
	err = toErr(ret)
	return
}

// UnlockRect unlocks a rectangle on a surface.
func (obj *Surface) UnlockRect() Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.UnlockRect,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return toErr(ret)
}

// GetDC retrieves a device context.
func (obj *Surface) GetDC() (hdc HDC, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetDC,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&hdc)),
		0,
	)
	err = toErr(ret)
	return
}

// ReleaseDC releases a device context handle.
func (obj *Surface) ReleaseDC(hdc HDC) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.ReleaseDC,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(hdc)),
		0,
	)
	return toErr(ret)
}
