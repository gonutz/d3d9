package d3d9

/*
#include "d3d9wrapper.h"

HRESULT IDirect3DVertexBuffer9GetDesc(
		IDirect3DVertexBuffer9* obj,
		D3DVERTEXBUFFER_DESC* pDesc) {
	return obj->lpVtbl->GetDesc(obj, pDesc);
}

HRESULT IDirect3DVertexBuffer9Lock(
		IDirect3DVertexBuffer9* obj,
		UINT OffsetToLock,
		UINT SizeToLock,
		VOID** ppbData,
		DWORD Flags) {
	return obj->lpVtbl->Lock(obj, OffsetToLock, SizeToLock, ppbData, Flags);
}

HRESULT IDirect3DVertexBuffer9Unlock(IDirect3DVertexBuffer9* obj) {
	return obj->lpVtbl->Unlock(obj);
}

void IDirect3DVertexBuffer9Release(IDirect3DVertexBuffer9* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"
import "unsafe"

// VertexBuffer and its methods are used to manipulate vertex buffer resources.
type VertexBuffer struct {
	Resource
	handle *C.IDirect3DVertexBuffer9
}

// Release has to be called when finished using the object to free its
// associated resources.
func (obj VertexBuffer) Release() {
	C.IDirect3DVertexBuffer9Release(obj.handle)
}

// GetDesc retrieves a description of the vertex buffer resource.
func (obj VertexBuffer) GetDesc() (pDesc VERTEXBUFFER_DESC, err Error) {
	var cpDesc C.D3DVERTEXBUFFER_DESC
	err = toErr(C.IDirect3DVertexBuffer9GetDesc(obj.handle, &cpDesc))
	pDesc.fromC(&cpDesc)
	return
}

// Lock locks a range of vertex data and obtains a pointer to the vertex buffer
// memory.
func (obj VertexBuffer) Lock(
	OffsetToLock uint,
	SizeToLock uint,
	Flags uint32,
) (
	data VertexBufferMemory,
	err Error,
) {
	var ppbData unsafe.Pointer
	err = toErr(C.IDirect3DVertexBuffer9Lock(
		obj.handle,
		C.UINT(OffsetToLock),
		C.UINT(SizeToLock),
		&ppbData,
		C.DWORD(Flags),
	))
	data = VertexBufferMemory{ppbData}
	return
}

// VertexBufferMemory encapsulates a raw memory pointer and provides methods to
// get and set typical slice types.
type VertexBufferMemory struct {
	Memory unsafe.Pointer
}

// SetFloat32s copies the given slice to the memory, starting at the given
// offset in bytes.
func (m VertexBufferMemory) SetFloat32s(offsetInBytes int, data []float32) {
	if len(data) > 0 {
		C.memcpy(
			unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInBytes)),
			unsafe.Pointer(&data[0]),
			C.size_t(len(data)*4),
		)
	}
}

// GetFloat32s copies data from memory to the given slice, starting at the given
// offset in bytes.
func (m VertexBufferMemory) GetFloat32s(offsetInBytes int, data []float32) {
	if len(data) > 0 {
		C.memcpy(
			unsafe.Pointer(&data[0]),
			unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInBytes)),
			C.size_t(len(data)*4),
		)
	}
}

// SetFloat64s copies the given slice to the memory, starting at the given
// offset in bytes.
func (m VertexBufferMemory) SetFloat64s(offsetInBytes int, data []float64) {
	if len(data) > 0 {
		C.memcpy(
			unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInBytes)),
			unsafe.Pointer(&data[0]),
			C.size_t(len(data)*8),
		)
	}
}

// GetFloat64s copies data from memory to the given slice, starting at the given
// offset in bytes.
func (m VertexBufferMemory) GetFloat64s(offsetInBytes int, data []float64) {
	if len(data) > 0 {
		C.memcpy(
			unsafe.Pointer(&data[0]),
			unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInBytes)),
			C.size_t(len(data)*8),
		)
	}
}

// SetInt32s copies the given slice to the memory, starting at the given
// offset in bytes.
func (m VertexBufferMemory) SetInt32s(offsetInBytes int, data []int32) {
	if len(data) > 0 {
		C.memcpy(
			unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInBytes)),
			unsafe.Pointer(&data[0]),
			C.size_t(len(data)*4),
		)
	}
}

// GetInt32s copies data from memory to the given slice, starting at the given
// offset in bytes.
func (m VertexBufferMemory) GetInt32s(offsetInBytes int, data []int32) {
	if len(data) > 0 {
		C.memcpy(
			unsafe.Pointer(&data[0]),
			unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInBytes)),
			C.size_t(len(data)*4),
		)
	}
}

// SetUint32s copies the given slice to the memory, starting at the given
// offset in bytes.
func (m VertexBufferMemory) SetUint32s(offsetInBytes int, data []uint32) {
	if len(data) > 0 {
		C.memcpy(
			unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInBytes)),
			unsafe.Pointer(&data[0]),
			C.size_t(len(data)*4),
		)
	}
}

// GetUint32s copies data from memory to the given slice, starting at the given
// offset in bytes.
func (m VertexBufferMemory) GetUint32s(offsetInBytes int, data []uint32) {
	if len(data) > 0 {
		C.memcpy(
			unsafe.Pointer(&data[0]),
			unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInBytes)),
			C.size_t(len(data)*4),
		)
	}
}

// SetInt16s copies the given slice to the memory, starting at the given
// offset in bytes.
func (m VertexBufferMemory) SetInt16s(offsetInBytes int, data []int16) {
	if len(data) > 0 {
		C.memcpy(
			unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInBytes)),
			unsafe.Pointer(&data[0]),
			C.size_t(len(data)*2),
		)
	}
}

// GetInt16s copies data from memory to the given slice, starting at the given
// offset in bytes.
func (m VertexBufferMemory) GetInt16s(offsetInBytes int, data []int16) {
	if len(data) > 0 {
		C.memcpy(
			unsafe.Pointer(&data[0]),
			unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInBytes)),
			C.size_t(len(data)*2),
		)
	}
}

// SetUint16s copies the given slice to the memory, starting at the given
// offset in bytes.
func (m VertexBufferMemory) SetUint16s(offsetInBytes int, data []uint16) {
	if len(data) > 0 {
		C.memcpy(
			unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInBytes)),
			unsafe.Pointer(&data[0]),
			C.size_t(len(data)*2),
		)
	}
}

// GetUint16s copies data from memory to the given slice, starting at the given
// offset in bytes.
func (m VertexBufferMemory) GetUint16s(offsetInBytes int, data []uint16) {
	if len(data) > 0 {
		C.memcpy(
			unsafe.Pointer(&data[0]),
			unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInBytes)),
			C.size_t(len(data)*2),
		)
	}
}

// SetInt8s copies the given slice to the memory, starting at the given
// offset in bytes.
func (m VertexBufferMemory) SetInt8s(offsetInBytes int, data []int8) {
	if len(data) > 0 {
		C.memcpy(
			unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInBytes)),
			unsafe.Pointer(&data[0]),
			C.size_t(len(data)),
		)
	}
}

// GetInt8s copies data from memory to the given slice, starting at the given
// offset in bytes.
func (m VertexBufferMemory) GetInt8s(offsetInBytes int, data []int8) {
	if len(data) > 0 {
		C.memcpy(
			unsafe.Pointer(&data[0]),
			unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInBytes)),
			C.size_t(len(data)),
		)
	}
}

// SetUint8s copies the given slice to the memory, starting at the given offset
// in bytes.
func (m VertexBufferMemory) SetUint8s(offsetInBytes int, data []uint8) {
	if len(data) > 0 {
		C.memcpy(
			unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInBytes)),
			unsafe.Pointer(&data[0]),
			C.size_t(len(data)),
		)
	}
}

// GetUint8s copies data from memory to the given slice, starting at the given
// offset in bytes.
func (m VertexBufferMemory) GetUint8s(offsetInBytes int, data []uint8) {
	if len(data) > 0 {
		C.memcpy(
			unsafe.Pointer(&data[0]),
			unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInBytes)),
			C.size_t(len(data)),
		)
	}
}

// SetBytes copies the given slice to the memory, starting at the given offset
// in bytes.
func (m VertexBufferMemory) SetBytes(offsetInBytes int, data []byte) {
	if len(data) > 0 {
		C.memcpy(
			unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInBytes)),
			unsafe.Pointer(&data[0]),
			C.size_t(len(data)),
		)
	}
}

// GetBytes copies data from memory to the given slice, starting at the given
// offset in bytes.
func (m VertexBufferMemory) GetBytes(offsetInBytes int, data []byte) {
	if len(data) > 0 {
		C.memcpy(
			unsafe.Pointer(&data[0]),
			unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInBytes)),
			C.size_t(len(data)),
		)
	}
}

// Unlock unlocks vertex data.
func (obj VertexBuffer) Unlock() (err Error) {
	err = toErr(C.IDirect3DVertexBuffer9Unlock(obj.handle))
	return
}
