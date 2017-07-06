package d3d9

import (
	"syscall"
	"unsafe"
)

// VertexBuffer and its methods are used to manipulate vertex buffer resources.
type VertexBuffer struct {
	vtbl *vertexBufferVtbl
}

type vertexBufferVtbl struct {
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
	Lock            uintptr
	Unlock          uintptr
	GetDesc         uintptr
}

// AddRef increments the reference count for an interface on an object. This
// method should be called for every new copy of a pointer to an interface on an
// object.
func (obj *VertexBuffer) AddRef() uint32 {
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
func (obj *VertexBuffer) Release() uint32 {
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
func (obj *VertexBuffer) GetDevice() (device *Device, err Error) {
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
func (obj *VertexBuffer) SetPrivateData(
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
func (obj *VertexBuffer) SetPrivateDataBytes(
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
func (obj *VertexBuffer) GetPrivateData(refguid GUID) (data []byte, err Error) {
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
func (obj *VertexBuffer) FreePrivateData(refguid GUID) Error {
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
func (obj *VertexBuffer) SetPriority(priorityNew uint32) uint32 {
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
func (obj *VertexBuffer) GetPriority() uint32 {
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
func (obj *VertexBuffer) PreLoad() {
	syscall.Syscall(
		obj.vtbl.PreLoad,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
}

// GetType returns the type of the resource.
func (obj *VertexBuffer) GetType() RESOURCETYPE {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetType,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return RESOURCETYPE(ret)
}

// Lock locks a range of vertex data and obtains a pointer to the vertex buffer
// memory.
func (obj *VertexBuffer) Lock(
	offsetToLock uint,
	sizeToLock uint,
	flags uint32,
) (data VertexBufferMemory, err Error) {
	var dataPtr uintptr
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.Lock,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(offsetToLock),
		uintptr(sizeToLock),
		uintptr(unsafe.Pointer(&dataPtr)),
		uintptr(flags),
		0,
	)
	err = toErr(ret)
	data = VertexBufferMemory{Memory: dataPtr}
	return
}

// Unlock unlocks vertex data.
func (obj *VertexBuffer) Unlock() Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Unlock,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return toErr(ret)
}

// GetDesc retrieves a description of the vertex buffer resource.
func (obj *VertexBuffer) GetDesc() (desc VERTEXBUFFER_DESC, err Error) {
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

// VertexBufferMemory encapsulates a raw memory pointer and provides methods to
// get and set typical slice types.
type VertexBufferMemory struct {
	Memory uintptr
}

// SetFloat32s copies the given slice to the memory, starting at the given
// offset in bytes.
func (m VertexBufferMemory) SetFloat32s(offsetInBytes int, data []float32) {
	if len(data) > 0 {
		dest := m.Memory + uintptr(offsetInBytes)
		for i := range data {
			*((*float32)(unsafe.Pointer(dest))) = data[i]
			dest += 4
		}
	}
}

// SetFloat64s copies the given slice to the memory, starting at the given
// offset in bytes.
func (m VertexBufferMemory) SetFloat64s(offsetInBytes int, data []float64) {
	if len(data) > 0 {
		dest := m.Memory + uintptr(offsetInBytes)
		for i := range data {
			*((*float64)(unsafe.Pointer(dest))) = data[i]
			dest += 8
		}
	}
}

// SetInt32s copies the given slice to the memory, starting at the given
// offset in bytes.
func (m VertexBufferMemory) SetInt32s(offsetInBytes int, data []int32) {
	if len(data) > 0 {
		dest := m.Memory + uintptr(offsetInBytes)
		for i := range data {
			*((*int32)(unsafe.Pointer(dest))) = data[i]
			dest += 4
		}
	}
}

// SetUint32s copies the given slice to the memory, starting at the given
// offset in bytes.
func (m VertexBufferMemory) SetUint32s(offsetInBytes int, data []uint32) {
	if len(data) > 0 {
		dest := m.Memory + uintptr(offsetInBytes)
		for i := range data {
			*((*uint32)(unsafe.Pointer(dest))) = data[i]
			dest += 4
		}
	}
}

// SetInt16s copies the given slice to the memory, starting at the given
// offset in bytes.
func (m VertexBufferMemory) SetInt16s(offsetInBytes int, data []int16) {
	if len(data) > 0 {
		dest := m.Memory + uintptr(offsetInBytes)
		for i := range data {
			*((*int16)(unsafe.Pointer(dest))) = data[i]
			dest += 2
		}
	}
}

// SetUint16s copies the given slice to the memory, starting at the given
// offset in bytes.
func (m VertexBufferMemory) SetUint16s(offsetInBytes int, data []uint16) {
	if len(data) > 0 {
		dest := m.Memory + uintptr(offsetInBytes)
		for i := range data {
			*((*uint16)(unsafe.Pointer(dest))) = data[i]
			dest += 2
		}
	}
}

// SetInt8s copies the given slice to the memory, starting at the given
// offset in bytes.
func (m VertexBufferMemory) SetInt8s(offsetInBytes int, data []int8) {
	if len(data) > 0 {
		dest := m.Memory + uintptr(offsetInBytes)
		for i := range data {
			*((*int8)(unsafe.Pointer(dest))) = data[i]
			dest++
		}
	}
}

// SetUint8s copies the given slice to the memory, starting at the given offset
// in bytes.
func (m VertexBufferMemory) SetUint8s(offsetInBytes int, data []uint8) {
	if len(data) > 0 {
		dest := m.Memory + uintptr(offsetInBytes)
		for i := range data {
			*((*uint8)(unsafe.Pointer(dest))) = data[i]
			dest++
		}
	}
}

// SetBytes copies the given slice to the memory, starting at the given offset
// in bytes.
func (m VertexBufferMemory) SetBytes(offsetInBytes int, data []byte) {
	if len(data) > 0 {
		dest := m.Memory + uintptr(offsetInBytes)
		for i := range data {
			*((*byte)(unsafe.Pointer(dest))) = data[i]
			dest++
		}
	}
}

// GetFloat32s copies data from memory to the given slice, starting at the given
// offset in bytes.
func (m VertexBufferMemory) GetFloat32s(offsetInBytes int, data []float32) {
	if len(data) > 0 {
		src := m.Memory + uintptr(offsetInBytes)
		for i := range data {
			data[i] = *((*float32)(unsafe.Pointer(src)))
			src += 4
		}
	}
}

// GetFloat64s copies data from memory to the given slice, starting at the given
// offset in bytes.
func (m VertexBufferMemory) GetFloat64s(offsetInBytes int, data []float64) {
	if len(data) > 0 {
		src := m.Memory + uintptr(offsetInBytes)
		for i := range data {
			data[i] = *((*float64)(unsafe.Pointer(src)))
			src += 8
		}
	}
}

// GetInt32s copies data from memory to the given slice, starting at the given
// offset in bytes.
func (m VertexBufferMemory) GetInt32s(offsetInBytes int, data []int32) {
	if len(data) > 0 {
		src := m.Memory + uintptr(offsetInBytes)
		for i := range data {
			data[i] = *((*int32)(unsafe.Pointer(src)))
			src += 4
		}
	}
}

// GetUint32s copies data from memory to the given slice, starting at the given
// offset in bytes.
func (m VertexBufferMemory) GetUint32s(offsetInBytes int, data []uint32) {
	if len(data) > 0 {
		src := m.Memory + uintptr(offsetInBytes)
		for i := range data {
			data[i] = *((*uint32)(unsafe.Pointer(src)))
			src += 4
		}
	}
}

// GetInt16s copies data from memory to the given slice, starting at the given
// offset in bytes.
func (m VertexBufferMemory) GetInt16s(offsetInBytes int, data []int16) {
	if len(data) > 0 {
		src := m.Memory + uintptr(offsetInBytes)
		for i := range data {
			data[i] = *((*int16)(unsafe.Pointer(src)))
			src += 2
		}
	}
}

// GetUint16s copies data from memory to the given slice, starting at the given
// offset in bytes.
func (m VertexBufferMemory) GetUint16s(offsetInBytes int, data []uint16) {
	if len(data) > 0 {
		src := m.Memory + uintptr(offsetInBytes)
		for i := range data {
			data[i] = *((*uint16)(unsafe.Pointer(src)))
			src += 2
		}
	}
}

// GetInt8s copies data from memory to the given slice, starting at the given
// offset in bytes.
func (m VertexBufferMemory) GetInt8s(offsetInBytes int, data []int8) {
	if len(data) > 0 {
		src := m.Memory + uintptr(offsetInBytes)
		for i := range data {
			data[i] = *((*int8)(unsafe.Pointer(src)))
			src++
		}
	}
}

// GetUint8s copies data from memory to the given slice, starting at the given
// offset in bytes.
func (m VertexBufferMemory) GetUint8s(offsetInBytes int, data []uint8) {
	if len(data) > 0 {
		src := m.Memory + uintptr(offsetInBytes)
		for i := range data {
			data[i] = *((*uint8)(unsafe.Pointer(src)))
			src++
		}
	}
}

// GetBytes copies data from memory to the given slice, starting at the given
// offset in bytes.
func (m VertexBufferMemory) GetBytes(offsetInBytes int, data []byte) {
	if len(data) > 0 {
		src := m.Memory + uintptr(offsetInBytes)
		for i := range data {
			data[i] = *((*byte)(unsafe.Pointer(src)))
			src++
		}
	}
}
