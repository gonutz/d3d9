package d3d9

/*
#include "d3d9wrapper.h"

HRESULT IDirect3DIndexBuffer9GetDesc(
		IDirect3DIndexBuffer9* obj,
		D3DINDEXBUFFER_DESC* pDesc) {
	return obj->lpVtbl->GetDesc(obj, pDesc);
}

HRESULT IDirect3DIndexBuffer9Lock(
		IDirect3DIndexBuffer9* obj,
		UINT OffsetToLock,
		UINT SizeToLock,
		VOID** ppbData,
		DWORD Flags) {
	return obj->lpVtbl->Lock(obj, OffsetToLock, SizeToLock, ppbData, Flags);
}

HRESULT IDirect3DIndexBuffer9Unlock(IDirect3DIndexBuffer9* obj) {
	return obj->lpVtbl->Unlock(obj);
}

void IDirect3DIndexBuffer9Release(IDirect3DIndexBuffer9* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"
import "unsafe"

// IndexBuffer and its methods are used to manipulate an index buffer resource.
type IndexBuffer struct {
	Resource
	handle *C.IDirect3DIndexBuffer9
}

// Release has to be called when finished using the object to free its
// associated resources.
func (obj IndexBuffer) Release() {
	C.IDirect3DIndexBuffer9Release(obj.handle)
}

// GetDesc retrieves a description of the index buffer resource.
func (obj IndexBuffer) GetDesc() (pDesc INDEXBUFFER_DESC, err Error) {
	var cpDesc C.D3DINDEXBUFFER_DESC
	err = toErr(C.IDirect3DIndexBuffer9GetDesc(obj.handle, &cpDesc))
	pDesc.fromC(&cpDesc)
	return
}

// Lock locks a range of index data and obtains a pointer to the index buffer
// memory.
func (obj IndexBuffer) Lock(
	OffsetToLock uint,
	SizeToLock uint,
	Flags uint32,
) (
	data IndexBufferMemory,
	err Error,
) {
	var ppbData unsafe.Pointer
	err = toErr(C.IDirect3DIndexBuffer9Lock(
		obj.handle,
		C.UINT(OffsetToLock),
		C.UINT(SizeToLock),
		&ppbData,
		C.DWORD(Flags),
	))
	data = IndexBufferMemory{ppbData}
	return
}

// IndexBufferMemory encapsulates a raw memory pointer and provides methods to
// get and set typical slice types.
type IndexBufferMemory struct {
	Memory unsafe.Pointer
}

// SetUint32s copies the given slice to the memory, starting at the given
// offset in bytes.
func (m IndexBufferMemory) SetUint32s(offsetInBytes int, data []uint32) {
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
func (m IndexBufferMemory) GetUint32s(offsetInBytes int, data []uint32) {
	if len(data) > 0 {
		C.memcpy(
			unsafe.Pointer(&data[0]),
			unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInBytes)),
			C.size_t(len(data)*4),
		)
	}
}

// SetUint16s copies the given slice to the memory, starting at the given
// offset in bytes.
func (m IndexBufferMemory) SetUint16s(offsetInBytes int, data []uint16) {
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
func (m IndexBufferMemory) GetUint16s(offsetInBytes int, data []uint16) {
	if len(data) > 0 {
		C.memcpy(
			unsafe.Pointer(&data[0]),
			unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInBytes)),
			C.size_t(len(data)*2),
		)
	}
}

// Unlock unlocks index data.
func (obj IndexBuffer) Unlock() (err Error) {
	err = toErr(C.IDirect3DIndexBuffer9Unlock(obj.handle))
	return
}
