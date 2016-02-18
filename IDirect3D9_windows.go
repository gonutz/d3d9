/*
Package d3d9 provides a wrapper for Microsoft's Direct3D9 API. The package can
thus only be used on Windows.
To build the package you need a working C compiler. It was tested with MinGW,
both the 32 bit and 64 bit versions work. You do not need to install any
additional libraries, this wrapper works with the standard MinGW installation.
When running a Direct3D application you need to have d3d9.dll installed on the
system, which fortunately has been deployed with all Windows versions since XP.
This means if you have Go installed, you also have the DLL installed.
*/
package d3d9

/*
#cgo CFLAGS: -DD3D_DISABLE_9EX

#include "d3d9wrapper.h"

HRESULT IDirect3D9CheckDepthStencilMatch(
		IDirect3D9* obj,
		UINT Adapter,
		D3DDEVTYPE DeviceType,
		D3DFORMAT AdapterFormat,
		D3DFORMAT RenderTargetFormat,
		D3DFORMAT DepthStencilFormat) {
	return obj->lpVtbl->CheckDepthStencilMatch(obj, Adapter, DeviceType,
		AdapterFormat, RenderTargetFormat, DepthStencilFormat);
}

HRESULT IDirect3D9CheckDeviceFormat(
		IDirect3D9* obj,
		UINT Adapter,
		D3DDEVTYPE DeviceType,
		D3DFORMAT AdapterFormat,
		DWORD Usage,
		D3DRESOURCETYPE RType,
		D3DFORMAT CheckFormat) {
	return obj->lpVtbl->CheckDeviceFormat(obj, Adapter, DeviceType,
		AdapterFormat, Usage, RType, CheckFormat);
}

HRESULT IDirect3D9CheckDeviceFormatConversion(
		IDirect3D9* obj,
		UINT Adapter,
		D3DDEVTYPE DeviceType,
		D3DFORMAT SourceFormat,
		D3DFORMAT TargetFormat) {
	return obj->lpVtbl->CheckDeviceFormatConversion(obj, Adapter, DeviceType,
		SourceFormat, TargetFormat);
}

HRESULT IDirect3D9CheckDeviceMultiSampleType(
		IDirect3D9* obj,
		UINT Adapter,
		D3DDEVTYPE DeviceType,
		D3DFORMAT SurfaceFormat,
		BOOL Windowed,
		D3DMULTISAMPLE_TYPE MultiSampleType,
		DWORD* pQualityLevels) {
	return obj->lpVtbl->CheckDeviceMultiSampleType(obj, Adapter, DeviceType,
		SurfaceFormat, Windowed, MultiSampleType, pQualityLevels);
}

HRESULT IDirect3D9CheckDeviceType(
		IDirect3D9* obj,
		UINT Adapter,
		D3DDEVTYPE DeviceType,
		D3DFORMAT DisplayFormat,
		D3DFORMAT BackBufferFormat,
		BOOL Windowed) {
	return obj->lpVtbl->CheckDeviceType(obj, Adapter, DeviceType, DisplayFormat,
		BackBufferFormat, Windowed);
}

HRESULT IDirect3D9CreateDevice(
		IDirect3D9* obj,
		UINT Adapter,
		D3DDEVTYPE DeviceType,
		HWND hFocusWindow,
		DWORD BehaviorFlags,
		D3DPRESENT_PARAMETERS* pPresentationParameters,
		IDirect3DDevice9** ppReturnedDeviceInterface) {
	return obj->lpVtbl->CreateDevice(obj, Adapter, DeviceType, hFocusWindow,
		BehaviorFlags, pPresentationParameters, ppReturnedDeviceInterface);
}

HRESULT IDirect3D9EnumAdapterModes(
		IDirect3D9* obj,
		UINT Adapter,
		D3DFORMAT Format,
		UINT Mode,
		D3DDISPLAYMODE* pMode) {
	return obj->lpVtbl->EnumAdapterModes(obj, Adapter, Format, Mode, pMode);
}

UINT IDirect3D9GetAdapterCount(IDirect3D9* obj) {
	return obj->lpVtbl->GetAdapterCount(obj);
}

HRESULT IDirect3D9GetAdapterDisplayMode(
		IDirect3D9* obj,
		UINT Adapter,
		D3DDISPLAYMODE* pMode) {
	return obj->lpVtbl->GetAdapterDisplayMode(obj, Adapter, pMode);
}

HRESULT IDirect3D9GetAdapterIdentifier(
		IDirect3D9* obj,
		UINT Adapter,
		DWORD Flags,
		D3DADAPTER_IDENTIFIER9* pIdentifier) {
	return obj->lpVtbl->GetAdapterIdentifier(obj, Adapter, Flags, pIdentifier);
}

UINT IDirect3D9GetAdapterModeCount(
		IDirect3D9* obj,
		UINT Adapter,
		D3DFORMAT Format) {
	return obj->lpVtbl->GetAdapterModeCount(obj, Adapter, Format);
}

HMONITOR IDirect3D9GetAdapterMonitor(IDirect3D9* obj, UINT Adapter) {
	return obj->lpVtbl->GetAdapterMonitor(obj, Adapter);
}

HRESULT IDirect3D9GetDeviceCaps(
		IDirect3D9* obj,
		UINT Adapter,
		D3DDEVTYPE DeviceType,
		D3DCAPS9* pCaps) {
	return obj->lpVtbl->GetDeviceCaps(obj, Adapter, DeviceType, pCaps);
}

HRESULT IDirect3D9RegisterSoftwareDevice(
		IDirect3D9* obj,
		void* pInitializeFunction) {
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

// Init loads the d3d9.dll and returns an error if it can not find it. Call
// init once before using this library and call Close once when you are
// finished.
func Init() (err error) {
	dll, err = syscall.LoadLibrary("d3d9.dll")
	return
}

// Close has to be called once after using this library, call Init once before
// calling any other function. Close will free the d3d9.dll library resource.
func Close() {
	syscall.FreeLibrary(dll)
}

// Direct3D and its methods are used to create Direct3D objects and set up the
// environment. Its interface includes methods for enumerating and retrieving
// capabilities of the device.
type Direct3D struct {
	handle *C.IDirect3D9
}

// Release has to be called when finished using the object to free its
// associated resources.
func (obj Direct3D) Release() {
	C.IDirect3D9Release(obj.handle)
}

// Create takes the SDK version as argument and returns a new Direct3D object.
// Use SDK_VERSION as the parameter.
func Create(version uint) (obj Direct3D, err error) {
	Direct3DCreate9, err := syscall.GetProcAddress(dll, "Direct3DCreate9")
	if err != nil {
		return Direct3D{}, err
	}
	result, _, callErr := syscall.Syscall(
		Direct3DCreate9,
		1,
		uintptr(version),
		0,
		0,
	)
	if callErr != 0 {
		return Direct3D{}, callErr
	}
	if result == 0 {
		return Direct3D{}, errors.New("Direct3DCreate9 returned nil")
	}
	return Direct3D{(*C.IDirect3D9)(unsafe.Pointer(result))}, nil
}

// CheckDepthStencilMatch determines whether a depth-stencil format is
// compatible with a render-target format in a particular display mode.
func (obj Direct3D) CheckDepthStencilMatch(
	Adapter uint,
	DeviceType DEVTYPE,
	AdapterFormat FORMAT,
	RenderTargetFormat FORMAT,
	DepthStencilFormat FORMAT,
) (
	err error,
) {
	err = toErr(C.IDirect3D9CheckDepthStencilMatch(
		obj.handle,
		(C.UINT)(Adapter),
		(C.D3DDEVTYPE)(DeviceType),
		(C.D3DFORMAT)(AdapterFormat),
		(C.D3DFORMAT)(RenderTargetFormat),
		(C.D3DFORMAT)(DepthStencilFormat),
	))
	return
}

// CheckDeviceFormat determines whether a surface format is available as a
// specified resource type and can be used as a texture, depth-stencil buffer,
// or render target, or any combination of the three, on a device representing
// this adapter.
func (obj Direct3D) CheckDeviceFormat(
	Adapter uint,
	DeviceType DEVTYPE,
	AdapterFormat FORMAT,
	Usage uint32,
	RType RESOURCETYPE,
	CheckFormat FORMAT,
) (
	err error,
) {
	err = toErr(C.IDirect3D9CheckDeviceFormat(
		obj.handle,
		(C.UINT)(Adapter),
		(C.D3DDEVTYPE)(DeviceType),
		(C.D3DFORMAT)(AdapterFormat),
		(C.DWORD)(Usage),
		(C.D3DRESOURCETYPE)(RType),
		(C.D3DFORMAT)(CheckFormat),
	))
	return
}

// CheckDeviceFormatConversion tests the device to see if it supports conversion
// from one display format to another.
func (obj Direct3D) CheckDeviceFormatConversion(
	Adapter uint,
	DeviceType DEVTYPE,
	SourceFormat FORMAT,
	TargetFormat FORMAT,
) (
	err error,
) {
	err = toErr(C.IDirect3D9CheckDeviceFormatConversion(
		obj.handle,
		(C.UINT)(Adapter),
		(C.D3DDEVTYPE)(DeviceType),
		(C.D3DFORMAT)(SourceFormat),
		(C.D3DFORMAT)(TargetFormat),
	))
	return
}

// CheckDeviceMultiSampleType determines if a multisampling technique is
// available on this device.
func (obj Direct3D) CheckDeviceMultiSampleType(
	Adapter uint,
	DeviceType DEVTYPE,
	SurfaceFormat FORMAT,
	Windowed bool,
	MultiSampleType MULTISAMPLE_TYPE,
) (
	pQualityLevels uint32,
	err error,
) {
	var cpQualityLevels C.DWORD
	err = toErr(C.IDirect3D9CheckDeviceMultiSampleType(obj.handle,
		(C.UINT)(Adapter),
		(C.D3DDEVTYPE)(DeviceType),
		(C.D3DFORMAT)(SurfaceFormat),
		toBOOL(Windowed),
		(C.D3DMULTISAMPLE_TYPE)(MultiSampleType),
		&cpQualityLevels,
	))
	pQualityLevels = (uint32)(cpQualityLevels)
	return
}

// CheckDeviceType verifies whether a hardware accelerated device type can be
// used on this adapter.
func (obj Direct3D) CheckDeviceType(
	Adapter uint,
	DeviceType DEVTYPE,
	DisplayFormat FORMAT,
	BackBufferFormat FORMAT,
	Windowed bool,
) (
	err error,
) {
	err = toErr(C.IDirect3D9CheckDeviceType(
		obj.handle,
		(C.UINT)(Adapter),
		(C.D3DDEVTYPE)(DeviceType),
		(C.D3DFORMAT)(DisplayFormat),
		(C.D3DFORMAT)(BackBufferFormat),
		toBOOL(Windowed),
	))
	return
}

// CreateDevice reates a device to represent the display adapter.
func (obj Direct3D) CreateDevice(
	Adapter uint,
	DeviceType DEVTYPE,
	hFocusWindow unsafe.Pointer,
	BehaviorFlags uint32,
	inpPresentationParameters PRESENT_PARAMETERS,
) (
	ppReturnedDeviceInterface Device,
	outpPresentationParameters PRESENT_PARAMETERS,
	err error,
) {
	cpPresentationParameters := inpPresentationParameters.toC()
	var cppReturnedDeviceInterface *C.IDirect3DDevice9
	err = toErr(C.IDirect3D9CreateDevice(
		obj.handle,
		(C.UINT)(Adapter),
		(C.D3DDEVTYPE)(DeviceType),
		(C.HWND)(hFocusWindow),
		(C.DWORD)(BehaviorFlags),
		&cpPresentationParameters,
		&cppReturnedDeviceInterface,
	))
	outpPresentationParameters.fromC(&cpPresentationParameters)
	ppReturnedDeviceInterface = Device{cppReturnedDeviceInterface}
	return
}

// EnumAdapterModes queries the device to determine whether the specified
// adapter supports the requested format and display mode. This method could be
// used in a loop to enumerate all the available adapter modes.
func (obj Direct3D) EnumAdapterModes(
	Adapter uint,
	Format FORMAT,
	Mode uint,
) (
	pMode DISPLAYMODE,
	err error,
) {
	var cpMode C.D3DDISPLAYMODE
	err = toErr(C.IDirect3D9EnumAdapterModes(
		obj.handle,
		(C.UINT)(Adapter),
		(C.D3DFORMAT)(Format),
		(C.UINT)(Mode),
		&cpMode,
	))
	pMode.fromC(&cpMode)
	return
}

// GetAdapterCount returns the number of adapters on the system.
func (obj Direct3D) GetAdapterCount() uint {
	return (uint)(C.IDirect3D9GetAdapterCount(obj.handle))
}

// GetAdapterDisplayMode retrieves the current display mode of the adapter.
func (obj Direct3D) GetAdapterDisplayMode(
	Adapter uint,
) (
	pMode DISPLAYMODE,
	err error,
) {
	var cpMode C.D3DDISPLAYMODE
	err = toErr(C.IDirect3D9GetAdapterDisplayMode(
		obj.handle,
		(C.UINT)(Adapter),
		&cpMode,
	))
	pMode.fromC(&cpMode)
	return
}

// GetAdapterIdentifier describes the physical display adapters present in the
// system when the IDirect3D9 interface was instantiated.
// Adapter must be a value from 0 to GetAdapterCount()-1.
func (obj Direct3D) GetAdapterIdentifier(
	Adapter uint,
	Flags uint32,
) (
	pIdentifier ADAPTER_IDENTIFIER,
	err error,
) {
	var cpIdentifier C.D3DADAPTER_IDENTIFIER9
	err = toErr(C.IDirect3D9GetAdapterIdentifier(
		obj.handle,
		(C.UINT)(Adapter),
		(C.DWORD)(Flags),
		&cpIdentifier,
	))
	pIdentifier.fromC(&cpIdentifier)
	return
}

// GetAdapterModeCount returns the number of display modes available on this
// adapter.
func (obj Direct3D) GetAdapterModeCount(Adapter uint, Format FORMAT) uint {
	return (uint)(C.IDirect3D9GetAdapterModeCount(
		obj.handle,
		(C.UINT)(Adapter),
		(C.D3DFORMAT)(Format),
	))
}

// GetAdapterMonitor returns the handle of the monitor associated with the
// Direct3D object.
func (obj Direct3D) GetAdapterMonitor(Adapter uint) unsafe.Pointer {
	return (unsafe.Pointer)(C.IDirect3D9GetAdapterMonitor(
		obj.handle,
		(C.UINT)(Adapter),
	))
}

// GetDeviceCaps retrieves device-specific information about a device.
func (obj Direct3D) GetDeviceCaps(
	Adapter uint,
	DeviceType DEVTYPE,
) (
	pCaps CAPS,
	err error,
) {
	var cpCaps C.D3DCAPS9
	err = toErr(C.IDirect3D9GetDeviceCaps(
		obj.handle,
		(C.UINT)(Adapter),
		(C.D3DDEVTYPE)(DeviceType),
		&cpCaps,
	))
	pCaps.fromC(&cpCaps)
	return
}

// RegisterSoftwareDevice registers a pluggable software device. Software
// devices provide software rasterization enabling applications to access a
// variety of software rasterizers.
func (obj Direct3D) RegisterSoftwareDevice(
	pInitializeFunction unsafe.Pointer,
) (
	err error,
) {
	err = toErr(C.IDirect3D9RegisterSoftwareDevice(
		obj.handle,
		pInitializeFunction,
	))
	return
}
