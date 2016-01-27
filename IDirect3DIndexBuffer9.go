package d3d9

/*
#include "d3d9wrapper.h"

HRESULT IDirect3DIndexBuffer9GetDesc(IDirect3DIndexBuffer9* obj, D3DINDEXBUFFER_DESC* pDesc) {
	return obj->lpVtbl->GetDesc(obj, pDesc);
}

HRESULT IDirect3DIndexBuffer9Lock(IDirect3DIndexBuffer9* obj, UINT OffsetToLock, UINT SizeToLock, VOID** ppbData, DWORD Flags) {
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

func (obj IndexBuffer) GetDesc() (pDesc INDEXBUFFER_DESC, err error) {
	var c_pDesc C.D3DINDEXBUFFER_DESC
	err = toErr(C.IDirect3DIndexBuffer9GetDesc(obj.handle, &c_pDesc))
	pDesc.fromC(&c_pDesc)
	return
}

func (obj IndexBuffer) Lock(OffsetToLock uint, SizeToLock uint, Flags uint32) (ppbData unsafe.Pointer, err error) {
	err = toErr(C.IDirect3DIndexBuffer9Lock(obj.handle, C.UINT(OffsetToLock), C.UINT(SizeToLock), &ppbData, C.DWORD(Flags)))
	return
}

func (obj IndexBuffer) Unlock() (err error) {
	err = toErr(C.IDirect3DIndexBuffer9Unlock(obj.handle))
	return
}

func (obj IndexBuffer) LockedSetFloats(OffsetToLock uint, Flags uint32, data []float32) (err error) {
	if len(data) == 0 {
		return nil
	}
	var buffer unsafe.Pointer
	byteCount := uint(len(data) * 4)
	buffer, err = obj.Lock(OffsetToLock, byteCount, Flags)
	if err != nil {
		return
	}
	C.memcpy(buffer, unsafe.Pointer(&data[0]), C.size_t(byteCount))
	return obj.Unlock()
}

func (obj IndexBuffer) LockedSetUints(OffsetToLock uint, Flags uint32, data []uint32) (err error) {
	if len(data) == 0 {
		return nil
	}
	var buffer unsafe.Pointer
	byteCount := uint(len(data) * 4)
	buffer, err = obj.Lock(OffsetToLock, byteCount, Flags)
	if err != nil {
		return
	}
	C.memcpy(buffer, unsafe.Pointer(&data[0]), C.size_t(byteCount))
	return obj.Unlock()
}
