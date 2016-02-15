package d3d9

/*
#include "d3d9wrapper.h"

HRESULT IDirect3DQuery9GetData(
		IDirect3DQuery9* obj,
		void* pData,
		DWORD dwSize,
		DWORD dwGetDataFlags) {
	return obj->lpVtbl->GetData(obj, pData, dwSize, dwGetDataFlags);
}

DWORD IDirect3DQuery9GetDataSize(IDirect3DQuery9* obj) {
	return obj->lpVtbl->GetDataSize(obj);
}

HRESULT IDirect3DQuery9GetDevice(
		IDirect3DQuery9* obj,
		IDirect3DDevice9** pDevice) {
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

// GetData polls a queried resource to get the query state or a query result.
func (obj Query) GetData(pData []byte, dwGetDataFlags uint32) (err error) {
	if len(pData) == 0 {
		err = toErr(C.IDirect3DQuery9GetData(
			obj.handle,
			nil,
			0,
			(C.DWORD)(dwGetDataFlags),
		))
	} else {
		err = toErr(C.IDirect3DQuery9GetData(
			obj.handle,
			unsafe.Pointer(&pData[0]),
			(C.DWORD)(len(pData)),
			(C.DWORD)(dwGetDataFlags),
		))
	}
	return
}

// GetDataSize returns the number of bytes in the query data.
func (obj Query) GetDataSize() uint32 {
	return uint32(C.IDirect3DQuery9GetDataSize(obj.handle))
}

// GetDevice returns the device that is being queried.
func (obj Query) GetDevice() (pDevice Device, err error) {
	var c_pDevice *C.IDirect3DDevice9
	err = toErr(C.IDirect3DQuery9GetDevice(obj.handle, &c_pDevice))
	pDevice = Device{c_pDevice}
	return
}

// GetType returns the query type.
func (obj Query) GetType() QUERYTYPE {
	return (QUERYTYPE)(C.IDirect3DQuery9GetType(obj.handle))
}

// Issue issues a query.
func (obj Query) Issue(dwIssueFlags uint32) (err error) {
	err = toErr(C.IDirect3DQuery9Issue(obj.handle, (C.DWORD)(dwIssueFlags)))
	return
}
