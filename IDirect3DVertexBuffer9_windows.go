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
// offset in items (not in bytes).
func (m VertexBufferMemory) SetFloat32s(offsetInFloats int, data []float32) {
	C.memcpy(
		unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInFloats*4)),
		unsafe.Pointer(&data[0]),
		C.size_t(len(data)*4),
	)
}

// GetFloat32s copies data from memory to the given slice, starting at the given
// offset in items (not in bytes).
func (m VertexBufferMemory) GetFloat32s(offsetInFloats int, data []float32) {
	C.memcpy(
		unsafe.Pointer(&data[0]),
		unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInFloats*4)),
		C.size_t(len(data)*4),
	)
}

// SetFloat64s copies the given slice to the memory, starting at the given
// offset in items (not in bytes).
func (m VertexBufferMemory) SetFloat64s(offsetInFloats int, data []float64) {
	C.memcpy(
		unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInFloats*8)),
		unsafe.Pointer(&data[0]),
		C.size_t(len(data)*8),
	)
}

// GetFloat64s copies data from memory to the given slice, starting at the given
// offset in items (not in bytes).
func (m VertexBufferMemory) GetFloat64s(offsetInFloats int, data []float64) {
	C.memcpy(
		unsafe.Pointer(&data[0]),
		unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInFloats*8)),
		C.size_t(len(data)*8),
	)
}

// SetInt32s copies the given slice to the memory, starting at the given
// offset in items (not in bytes).
func (m VertexBufferMemory) SetInt32s(offsetInInts int, data []int32) {
	C.memcpy(
		unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInInts*4)),
		unsafe.Pointer(&data[0]),
		C.size_t(len(data)*4),
	)
}

// GetInt32s copies data from memory to the given slice, starting at the given
// offset in items (not in bytes).
func (m VertexBufferMemory) GetInt32s(offsetInInts int, data []int32) {
	C.memcpy(
		unsafe.Pointer(&data[0]),
		unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInInts*4)),
		C.size_t(len(data)*4),
	)
}

// SetUint32s copies the given slice to the memory, starting at the given
// offset in items (not in bytes).
func (m VertexBufferMemory) SetUint32s(offsetInInts int, data []uint32) {
	C.memcpy(
		unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInInts*4)),
		unsafe.Pointer(&data[0]),
		C.size_t(len(data)*4),
	)
}

// GetUint32s copies data from memory to the given slice, starting at the given
// offset in items (not in bytes).
func (m VertexBufferMemory) GetUint32s(offsetInInts int, data []uint32) {
	C.memcpy(
		unsafe.Pointer(&data[0]),
		unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInInts*4)),
		C.size_t(len(data)*4),
	)
}

// SetInt16s copies the given slice to the memory, starting at the given
// offset in items (not in bytes).
func (m VertexBufferMemory) SetInt16s(offsetInInts int, data []int16) {
	C.memcpy(
		unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInInts*2)),
		unsafe.Pointer(&data[0]),
		C.size_t(len(data)*2),
	)
}

// GetInt16s copies data from memory to the given slice, starting at the given
// offset in items (not in bytes).
func (m VertexBufferMemory) GetInt16s(offsetInInts int, data []int16) {
	C.memcpy(
		unsafe.Pointer(&data[0]),
		unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInInts*2)),
		C.size_t(len(data)*2),
	)
}

// SetUint16s copies the given slice to the memory, starting at the given
// offset in items (not in bytes).
func (m VertexBufferMemory) SetUint16s(offsetInInts int, data []uint16) {
	C.memcpy(
		unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInInts*2)),
		unsafe.Pointer(&data[0]),
		C.size_t(len(data)*2),
	)
}

// GetUint16s copies data from memory to the given slice, starting at the given
// offset in items (not in bytes).
func (m VertexBufferMemory) GetUint16s(offsetInInts int, data []uint16) {
	C.memcpy(
		unsafe.Pointer(&data[0]),
		unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInInts*2)),
		C.size_t(len(data)*2),
	)
}

// SetInt8s copies the given slice to the memory, starting at the given
// offset in items (not in bytes).
func (m VertexBufferMemory) SetInt8s(offsetInInts int, data []int8) {
	C.memcpy(
		unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInInts)),
		unsafe.Pointer(&data[0]),
		C.size_t(len(data)),
	)
}

// GetInt8s copies data from memory to the given slice, starting at the given
// offset in items (not in bytes).
func (m VertexBufferMemory) GetInt8s(offsetInInts int, data []int8) {
	C.memcpy(
		unsafe.Pointer(&data[0]),
		unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInInts)),
		C.size_t(len(data)),
	)
}

// SetUint8s copies the given slice to the memory, starting at the given
// offset in items (not in bytes). You can also use it to set a byte slice, just
// cast the []uint8 slice to []byte.
func (m VertexBufferMemory) SetUint8s(offsetInInts int, data []uint8) {
	C.memcpy(
		unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInInts)),
		unsafe.Pointer(&data[0]),
		C.size_t(len(data)),
	)
}

// GetUint8s copies data from memory to the given slice, starting at the given
// offset in items (not in bytes). You can also use it to get a byte slice, just
// cast the result form []uint8 to []byte.
func (m VertexBufferMemory) GetUint8s(offsetInInts int, data []uint8) {
	C.memcpy(
		unsafe.Pointer(&data[0]),
		unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInInts)),
		C.size_t(len(data)),
	)
}

// Unlock unlocks vertex data.
func (obj VertexBuffer) Unlock() (err Error) {
	err = toErr(C.IDirect3DVertexBuffer9Unlock(obj.handle))
	return
}
