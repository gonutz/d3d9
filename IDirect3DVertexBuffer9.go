package d3d9

/*
#include <d3d9.h>
#include <string.h>

HRESULT IDirect3DVertexBuffer9GetDesc(IDirect3DVertexBuffer9* obj, D3DVERTEXBUFFER_DESC* pDesc) {
	return obj->lpVtbl->GetDesc(obj, pDesc);
}

HRESULT IDirect3DVertexBuffer9Lock(IDirect3DVertexBuffer9* obj, UINT OffsetToLock, UINT SizeToLock, VOID** ppbData, DWORD Flags) {
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

type VertexBuffer struct {
	Resource
	handle *C.IDirect3DVertexBuffer9
}

func (obj VertexBuffer) Release() {
	C.IDirect3DVertexBuffer9Release(obj.handle)
}

func (obj VertexBuffer) GetDesc() (pDesc VERTEXBUFFER_DESC, err error) {
	var c_pDesc C.D3DVERTEXBUFFER_DESC
	err = toErr(C.IDirect3DVertexBuffer9GetDesc(obj.handle, &c_pDesc))
	pDesc.fromC(&c_pDesc)
	return
}

func (obj VertexBuffer) Lock(OffsetToLock uint, SizeToLock uint, Flags uint32) (ppbData unsafe.Pointer, err error) {
	err = toErr(C.IDirect3DVertexBuffer9Lock(obj.handle, C.UINT(OffsetToLock), C.UINT(SizeToLock), &ppbData, C.DWORD(Flags)))
	return
}

func (obj VertexBuffer) Unlock() (err error) {
	err = toErr(C.IDirect3DVertexBuffer9Unlock(obj.handle))
	return
}

func (obj VertexBuffer) LockedSetFloats(OffsetToLock uint, Flags uint32, data []float32) (err error) {
	var buffer unsafe.Pointer
	byteCount := uint(len(data) * 4)
	buffer, err = obj.Lock(OffsetToLock, byteCount, Flags)
	if err != nil {
		return
	}
	C.memcpy(buffer, unsafe.Pointer(&data[0]), C.size_t(byteCount))
	return obj.Unlock()
}

func (obj VertexBuffer) LockedSetUints(OffsetToLock uint, Flags uint32, data []uint32) (err error) {
	var buffer unsafe.Pointer
	byteCount := uint(len(data) * 4)
	buffer, err = obj.Lock(OffsetToLock, byteCount, Flags)
	if err != nil {
		return
	}
	C.memcpy(buffer, unsafe.Pointer(&data[0]), C.size_t(byteCount))
	return obj.Unlock()
}
