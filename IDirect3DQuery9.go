package d3d9

/*
#include <d3d9.h>

HRESULT IDirect3DQuery9GetData(IDirect3DQuery9* obj, void* pData, DWORD dwSize, DWORD dwGetDataFlags) {
	return obj->lpVtbl->GetData(obj, pData, dwSize, dwGetDataFlags);
}

DWORD IDirect3DQuery9GetDataSize(IDirect3DQuery9* obj) {
	return obj->lpVtbl->GetDataSize(obj);
}

HRESULT IDirect3DQuery9GetDevice(IDirect3DQuery9* obj, IDirect3DDevice9** pDevice) {
	return obj->lpVtbl->GetDevice(obj, pDevice);
}

D3DQUERYTYPE IDirect3DQuery9GetType(IDirect3DQuery9* obj) {
	return obj->lpVtbl->GetType(obj);
}

HRESULT IDirect3DQuery9Issue(IDirect3DQuery9* obj, DWORD dwIssueFlags) {
	return obj->lpVtbl->Issue(obj, dwIssueFlags);
}

void IDirect3DQuery9Release(IDirect3DQuery9* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"
import "unsafe"

type Query struct {
	handle *C.IDirect3DQuery9
}

func (obj Query) Release() {
	C.IDirect3DQuery9Release(obj.handle)
}

func (obj Query) GetData(pData unsafe.Pointer, dwSize uint32, dwGetDataFlags uint32) (err error) {
	err = toErr(C.IDirect3DQuery9GetData(obj.handle, pData, (C.DWORD)(dwSize), (C.DWORD)(dwGetDataFlags)))
	return
}

func (obj Query) GetDataSize() uint32 {
	return uint32(C.IDirect3DQuery9GetDataSize(obj.handle))
}

func (obj Query) GetDevice() (pDevice Device, err error) {
	var c_pDevice *C.IDirect3DDevice9
	err = toErr(C.IDirect3DQuery9GetDevice(obj.handle, &c_pDevice))
	pDevice = Device{c_pDevice}
	return
}

func (obj Query) GetType() QUERYTYPE {
	return (QUERYTYPE)(C.IDirect3DQuery9GetType(obj.handle))
}

func (obj Query) Issue(dwIssueFlags uint32) (err error) {
	err = toErr(C.IDirect3DQuery9Issue(obj.handle, (C.DWORD)(dwIssueFlags)))
	return
}
