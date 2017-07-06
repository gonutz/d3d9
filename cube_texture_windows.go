package d3d9

import (
	"syscall"
	"unsafe"
)

// CubeTexture and its methods are used to manipulate a cube texture resource.
type CubeTexture struct {
	vtbl *cubeTextureVtbl
}

type cubeTextureVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	GetDevice            uintptr
	SetPrivateData       uintptr
	GetPrivateData       uintptr
	FreePrivateData      uintptr
	SetPriority          uintptr
	GetPriority          uintptr
	PreLoad              uintptr
	GetType              uintptr
	SetLOD               uintptr
	GetLOD               uintptr
	GetLevelCount        uintptr
	SetAutoGenFilterType uintptr
	GetAutoGenFilterType uintptr
	GenerateMipSubLevels uintptr
	GetLevelDesc         uintptr
	GetCubeMapSurface    uintptr
	LockRect             uintptr
	UnlockRect           uintptr
	AddDirtyRect         uintptr
}

func (obj *CubeTexture) baseTexturePointer() uintptr {
	return uintptr(unsafe.Pointer(obj))
}

// AddRef increments the reference count for an interface on an object. This
// method should be called for every new copy of a pointer to an interface on an
// object.
func (obj *CubeTexture) AddRef() uint32 {
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
func (obj *CubeTexture) Release() uint32 {
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
func (obj *CubeTexture) GetDevice() (device *Device, err Error) {
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
func (obj *CubeTexture) SetPrivateData(
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
func (obj *CubeTexture) SetPrivateDataBytes(
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
func (obj *CubeTexture) GetPrivateData(refguid GUID) (data []byte, err Error) {
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
func (obj *CubeTexture) FreePrivateData(refguid GUID) Error {
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
func (obj *CubeTexture) SetPriority(priorityNew uint32) uint32 {
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
func (obj *CubeTexture) GetPriority() uint32 {
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
func (obj *CubeTexture) PreLoad() {
	syscall.Syscall(
		obj.vtbl.PreLoad,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
}

// GetType returns the type of the resource.
func (obj *CubeTexture) GetType() RESOURCETYPE {
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
func (obj *CubeTexture) SetLOD(lodNew uint32) uint32 {
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
func (obj *CubeTexture) GetLOD() uint32 {
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
func (obj *CubeTexture) GetLevelCount() uint32 {
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
func (obj *CubeTexture) SetAutoGenFilterType(typ TEXTUREFILTERTYPE) Error {
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
func (obj *CubeTexture) GetAutoGenFilterType() TEXTUREFILTERTYPE {
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
func (obj *CubeTexture) GenerateMipSubLevels() {
	syscall.Syscall(
		obj.vtbl.GenerateMipSubLevels,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
}

// GetLevelDesc retrieves a description of one face of the specified cube
// texture level.
func (obj *CubeTexture) GetLevelDesc(level uint) (desc SURFACE_DESC, err Error) {
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

// GetCubeMapSurface retrieves a cube texture map surface.
func (obj *CubeTexture) GetCubeMapSurface(
	faceType CUBEMAP_FACES,
	level uint,
) (surface *Surface, err Error) {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.GetCubeMapSurface,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(faceType),
		uintptr(level),
		uintptr(unsafe.Pointer(&surface)),
		0,
		0,
	)
	err = toErr(ret)
	return
}

// LockRect locks a rectangle on a cube texture resource.
func (obj *CubeTexture) LockRect(
	faceType CUBEMAP_FACES,
	level uint,
	rect *RECT,
	flags uint32,
) (lockedRect LOCKED_RECT, err Error) {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.LockRect,
		6,
		uintptr(unsafe.Pointer(obj)),
		uintptr(faceType),
		uintptr(level),
		uintptr(unsafe.Pointer(&lockedRect)),
		uintptr(unsafe.Pointer(rect)),
		uintptr(flags),
	)
	err = toErr(ret)
	return
}

// UnlockRect unlocks a rectangle on a cube texture resource.
func (obj *CubeTexture) UnlockRect(faceType CUBEMAP_FACES, level uint) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.UnlockRect,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(faceType),
		uintptr(level),
	)
	return toErr(ret)
}

// AddDirtyRect adds a dirty region to a cube texture resource.
func (obj *CubeTexture) AddDirtyRect(faceType CUBEMAP_FACES, dirtyRect *RECT) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.AddDirtyRect,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(faceType),
		uintptr(unsafe.Pointer(dirtyRect)),
	)
	return toErr(ret)
}
