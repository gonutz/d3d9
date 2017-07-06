package d3d9

import (
	"syscall"
	"unsafe"
)

// Query and its methods are used to perform asynchronous queries on a driver.
type Query struct {
	vtbl *queryVtbl
}

type queryVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	GetDevice   uintptr
	GetType     uintptr
	GetDataSize uintptr
	Issue       uintptr
	GetData     uintptr
}

// AddRef increments the reference count for an interface on an object. This
// method should be called for every new copy of a pointer to an interface on an
// object.
func (obj *Query) AddRef() uint32 {
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
func (obj *Query) Release() uint32 {
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
func (obj *Query) GetDevice() (device *Device, err Error) {
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

// GetType returns the query type.
func (obj *Query) GetType() QUERYTYPE {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetType,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return QUERYTYPE(ret)
}

// GetDataSize returns the number of bytes in the query data.
func (obj *Query) GetDataSize() uint32 {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetDataSize,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint32(ret)
}

// Issue issues a query.
func (obj *Query) Issue(issueFlags uint32) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Issue,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(issueFlags),
		0,
	)
	return toErr(ret)
}

// GetData polls a queried resource to get the query state or a query result.
func (obj *Query) GetData(data []byte, flags uint32) Error {
	var dataPtr uintptr
	if len(data) > 0 {
		dataPtr = uintptr(unsafe.Pointer(&data[0]))
	}
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.GetData,
		4,
		uintptr(unsafe.Pointer(obj)),
		dataPtr,
		uintptr(len(data)),
		uintptr(flags),
		0,
		0,
	)
	return toErr(ret)
}
