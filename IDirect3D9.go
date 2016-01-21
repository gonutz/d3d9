package d3d9

/*
#cgo CFLAGS: -DD3D_DISABLE_9EX

#include "d3d9wrapper.h"

HRESULT IDirect3D9CheckDepthStencilMatch(IDirect3D9* obj, UINT Adapter, D3DDEVTYPE DeviceType, D3DFORMAT AdapterFormat, D3DFORMAT RenderTargetFormat, D3DFORMAT DepthStencilFormat) {
	return obj->lpVtbl->CheckDepthStencilMatch(obj, Adapter, DeviceType, AdapterFormat, RenderTargetFormat, DepthStencilFormat);
}

HRESULT IDirect3D9CheckDeviceFormat(IDirect3D9* obj, UINT Adapter, D3DDEVTYPE DeviceType, D3DFORMAT AdapterFormat, DWORD Usage, D3DRESOURCETYPE RType, D3DFORMAT CheckFormat) {
	return obj->lpVtbl->CheckDeviceFormat(obj, Adapter, DeviceType, AdapterFormat, Usage, RType, CheckFormat);
}

HRESULT IDirect3D9CheckDeviceFormatConversion(IDirect3D9* obj, UINT Adapter, D3DDEVTYPE DeviceType, D3DFORMAT SourceFormat, D3DFORMAT TargetFormat) {
	return obj->lpVtbl->CheckDeviceFormatConversion(obj, Adapter, DeviceType, SourceFormat, TargetFormat);
}

HRESULT IDirect3D9CheckDeviceMultiSampleType(IDirect3D9* obj, UINT Adapter, D3DDEVTYPE DeviceType, D3DFORMAT SurfaceFormat, BOOL Windowed, D3DMULTISAMPLE_TYPE MultiSampleType, DWORD* pQualityLevels) {
	return obj->lpVtbl->CheckDeviceMultiSampleType(obj, Adapter, DeviceType, SurfaceFormat, Windowed, MultiSampleType, pQualityLevels);
}

HRESULT IDirect3D9CheckDeviceType(IDirect3D9* obj, UINT Adapter, D3DDEVTYPE DeviceType, D3DFORMAT DisplayFormat, D3DFORMAT BackBufferFormat, BOOL Windowed) {
	return obj->lpVtbl->CheckDeviceType(obj, Adapter, DeviceType, DisplayFormat, BackBufferFormat, Windowed);
}

HRESULT IDirect3D9CreateDevice(IDirect3D9* obj, UINT Adapter, D3DDEVTYPE DeviceType, HWND hFocusWindow, DWORD BehaviorFlags, D3DPRESENT_PARAMETERS* pPresentationParameters, IDirect3DDevice9** ppReturnedDeviceInterface) {
	return obj->lpVtbl->CreateDevice(obj, Adapter, DeviceType, hFocusWindow, BehaviorFlags, pPresentationParameters, ppReturnedDeviceInterface);
}

HRESULT IDirect3D9EnumAdapterModes(IDirect3D9* obj, UINT Adapter, D3DFORMAT Format, UINT Mode, D3DDISPLAYMODE* pMode) {
	return obj->lpVtbl->EnumAdapterModes(obj, Adapter, Format, Mode, pMode);
}

UINT IDirect3D9GetAdapterCount(IDirect3D9* obj) {
	return obj->lpVtbl->GetAdapterCount(obj);
}

HRESULT IDirect3D9GetAdapterDisplayMode(IDirect3D9* obj, UINT Adapter, D3DDISPLAYMODE* pMode) {
	return obj->lpVtbl->GetAdapterDisplayMode(obj, Adapter, pMode);
}

HRESULT IDirect3D9GetAdapterIdentifier(IDirect3D9* obj, UINT Adapter, DWORD Flags, D3DADAPTER_IDENTIFIER9* pIdentifier) {
	return obj->lpVtbl->GetAdapterIdentifier(obj, Adapter, Flags, pIdentifier);
}

UINT IDirect3D9GetAdapterModeCount(IDirect3D9* obj, UINT Adapter, D3DFORMAT Format) {
	return obj->lpVtbl->GetAdapterModeCount(obj, Adapter, Format);
}

HMONITOR IDirect3D9GetAdapterMonitor(IDirect3D9* obj, UINT Adapter) {
	return obj->lpVtbl->GetAdapterMonitor(obj, Adapter);
}

HRESULT IDirect3D9GetDeviceCaps(IDirect3D9* obj, UINT Adapter, D3DDEVTYPE DeviceType, D3DCAPS9* pCaps) {
	return obj->lpVtbl->GetDeviceCaps(obj, Adapter, DeviceType, pCaps);
}

HRESULT IDirect3D9RegisterSoftwareDevice(IDirect3D9* obj, void* pInitializeFunction) {
	return obj->lpVtbl->RegisterSoftwareDevice(obj, pInitializeFunction);
}

void IDirect3D9Release(IDirect3D9* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"
import (
	"errors"
	"syscall"
	"unsafe"
)

var dll syscall.Handle

func Init() (err error) {
	dll, err = syscall.LoadLibrary("d3d9.dll")
	return
}

func Close() {
	syscall.FreeLibrary(dll)
}

type Direct3D struct {
	handle *C.IDirect3D9
}

func (obj Direct3D) Release() {
	C.IDirect3D9Release(obj.handle)
}

func Create(version uint) (obj Direct3D, err error) {
	Direct3DCreate9, err := syscall.GetProcAddress(dll, "Direct3DCreate9")
	if err != nil {
		return Direct3D{}, err
	}
	result, _, callErr := syscall.Syscall(Direct3DCreate9, 1, uintptr(version), 0, 0)
	if callErr != 0 {
		return Direct3D{}, callErr
	}
	if result == 0 {
		return Direct3D{}, errors.New("Direct3DCreate9 returned nil")
	}
	return Direct3D{(*C.IDirect3D9)(unsafe.Pointer(result))}, nil
}

func (obj Direct3D) CheckDepthStencilMatch(Adapter uint, DeviceType DEVTYPE, AdapterFormat FORMAT, RenderTargetFormat FORMAT, DepthStencilFormat FORMAT) (err error) {
	err = toErr(C.IDirect3D9CheckDepthStencilMatch(obj.handle, (C.UINT)(Adapter), (C.D3DDEVTYPE)(DeviceType), (C.D3DFORMAT)(AdapterFormat), (C.D3DFORMAT)(RenderTargetFormat), (C.D3DFORMAT)(DepthStencilFormat)))
	return
}

func (obj Direct3D) CheckDeviceFormat(Adapter uint, DeviceType DEVTYPE, AdapterFormat FORMAT, Usage uint32, RType RESOURCETYPE, CheckFormat FORMAT) (err error) {
	err = toErr(C.IDirect3D9CheckDeviceFormat(obj.handle, (C.UINT)(Adapter), (C.D3DDEVTYPE)(DeviceType), (C.D3DFORMAT)(AdapterFormat), (C.DWORD)(Usage), (C.D3DRESOURCETYPE)(RType), (C.D3DFORMAT)(CheckFormat)))
	return
}

func (obj Direct3D) CheckDeviceFormatConversion(Adapter uint, DeviceType DEVTYPE, SourceFormat FORMAT, TargetFormat FORMAT) (err error) {
	err = toErr(C.IDirect3D9CheckDeviceFormatConversion(obj.handle, (C.UINT)(Adapter), (C.D3DDEVTYPE)(DeviceType), (C.D3DFORMAT)(SourceFormat), (C.D3DFORMAT)(TargetFormat)))
	return
}

func (obj Direct3D) CheckDeviceMultiSampleType(Adapter uint, DeviceType DEVTYPE, SurfaceFormat FORMAT, Windowed bool, MultiSampleType MULTISAMPLE_TYPE) (pQualityLevels uint32, err error) {
	var c_pQualityLevels C.DWORD
	err = toErr(C.IDirect3D9CheckDeviceMultiSampleType(obj.handle,
		(C.UINT)(Adapter),
		(C.D3DDEVTYPE)(DeviceType),
		(C.D3DFORMAT)(SurfaceFormat),
		toBOOL(Windowed),
		(C.D3DMULTISAMPLE_TYPE)(MultiSampleType),
		&c_pQualityLevels),
	)
	pQualityLevels = (uint32)(c_pQualityLevels)
	return
}

func (obj Direct3D) CheckDeviceType(Adapter uint, DeviceType DEVTYPE, DisplayFormat FORMAT, BackBufferFormat FORMAT, Windowed bool) (err error) {
	err = toErr(C.IDirect3D9CheckDeviceType(obj.handle, (C.UINT)(Adapter), (C.D3DDEVTYPE)(DeviceType), (C.D3DFORMAT)(DisplayFormat), (C.D3DFORMAT)(BackBufferFormat), toBOOL(Windowed)))
	return
}

func (obj Direct3D) CreateDevice(Adapter uint, DeviceType DEVTYPE, hFocusWindow unsafe.Pointer, BehaviorFlags uint32, inpPresentationParameters PRESENT_PARAMETERS) (ppReturnedDeviceInterface Device, outpPresentationParameters PRESENT_PARAMETERS, err error) {
	c_pPresentationParameters := inpPresentationParameters.toC()
	var c_ppReturnedDeviceInterface *C.IDirect3DDevice9
	err = toErr(C.IDirect3D9CreateDevice(obj.handle, (C.UINT)(Adapter), (C.D3DDEVTYPE)(DeviceType), (C.HWND)(hFocusWindow), (C.DWORD)(BehaviorFlags), &c_pPresentationParameters, &c_ppReturnedDeviceInterface))
	outpPresentationParameters.fromC(&c_pPresentationParameters)
	ppReturnedDeviceInterface = Device{c_ppReturnedDeviceInterface}
	return
}

func (obj Direct3D) EnumAdapterModes(Adapter uint, Format FORMAT, Mode uint) (pMode DISPLAYMODE, err error) {
	var c_pMode C.D3DDISPLAYMODE
	err = toErr(C.IDirect3D9EnumAdapterModes(obj.handle, (C.UINT)(Adapter), (C.D3DFORMAT)(Format), (C.UINT)(Mode), &c_pMode))
	pMode.fromC(&c_pMode)
	return
}

func (obj Direct3D) GetAdapterCount() uint {
	return (uint)(C.IDirect3D9GetAdapterCount(obj.handle))
}

func (obj Direct3D) GetAdapterDisplayMode(Adapter uint) (pMode DISPLAYMODE, err error) {
	var c_pMode C.D3DDISPLAYMODE
	err = toErr(C.IDirect3D9GetAdapterDisplayMode(obj.handle, (C.UINT)(Adapter), &c_pMode))
	pMode.fromC(&c_pMode)
	return
}

func (obj Direct3D) GetAdapterIdentifier(Adapter uint, Flags uint32) (pIdentifier ADAPTER_IDENTIFIER, err error) {
	var c_pIdentifier C.D3DADAPTER_IDENTIFIER9
	err = toErr(C.IDirect3D9GetAdapterIdentifier(obj.handle, (C.UINT)(Adapter), (C.DWORD)(Flags), &c_pIdentifier))
	pIdentifier.fromC(&c_pIdentifier)
	return
}

func (obj Direct3D) GetAdapterModeCount(Adapter uint, Format FORMAT) uint {
	return (uint)(C.IDirect3D9GetAdapterModeCount(obj.handle, (C.UINT)(Adapter), (C.D3DFORMAT)(Format)))
}

func (obj Direct3D) GetAdapterMonitor(Adapter uint) unsafe.Pointer {
	return (unsafe.Pointer)(C.IDirect3D9GetAdapterMonitor(obj.handle, (C.UINT)(Adapter)))
}

func (obj Direct3D) GetDeviceCaps(Adapter uint, DeviceType DEVTYPE) (pCaps CAPS, err error) {
	var c_pCaps C.D3DCAPS9
	err = toErr(C.IDirect3D9GetDeviceCaps(obj.handle, (C.UINT)(Adapter), (C.D3DDEVTYPE)(DeviceType), &c_pCaps))
	pCaps.fromC(&c_pCaps)
	return
}

func (obj Direct3D) RegisterSoftwareDevice(pInitializeFunction unsafe.Pointer) (err error) {
	err = toErr(C.IDirect3D9RegisterSoftwareDevice(obj.handle, pInitializeFunction))
	return
}
