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

type IndexBuffer struct {
	Resource
	handle *C.IDirect3DIndexBuffer9
}

func (obj IndexBuffer) Release() {
	C.IDirect3DIndexBuffer9Release(obj.handle)
}

// GetDesc retrieves a description of the index buffer resource.
func (obj IndexBuffer) GetDesc() (pDesc INDEXBUFFER_DESC, err error) {
	var c_pDesc C.D3DINDEXBUFFER_DESC
	err = toErr(C.IDirect3DIndexBuffer9GetDesc(obj.handle, &c_pDesc))
	pDesc.fromC(&c_pDesc)
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
	err error,
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

type IndexBufferMemory struct {
	Memory unsafe.Pointer
}

func (m IndexBufferMemory) SetUint32s(offsetInInts int, data []uint32) {
	C.memcpy(
		unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInInts*4)),
		unsafe.Pointer(&data[0]),
		C.size_t(len(data)*4),
	)
}

func (m IndexBufferMemory) GetUint32s(offsetInInts int, data []uint32) {
	C.memcpy(
		unsafe.Pointer(&data[0]),
		unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInInts*4)),
		C.size_t(len(data)*4),
	)
}

func (m IndexBufferMemory) SetUint16s(offsetInInts int, data []uint16) {
	C.memcpy(
		unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInInts*2)),
		unsafe.Pointer(&data[0]),
		C.size_t(len(data)*2),
	)
}

func (m IndexBufferMemory) GetUint16s(offsetInInts int, data []uint16) {
	C.memcpy(
		unsafe.Pointer(&data[0]),
		unsafe.Pointer(uintptr(int(uintptr(m.Memory))+offsetInInts*2)),
		C.size_t(len(data)*2),
	)
}

// Unlock unlocks index data.
func (obj IndexBuffer) Unlock() (err error) {
	err = toErr(C.IDirect3DIndexBuffer9Unlock(obj.handle))
	return
}
