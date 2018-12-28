/*
Package d3d9 provides a wrapper for Microsoft's Direct3D9 API in pure Go.
It can only be used on Windows.

When running a Direct3D application you need to have d3d9.dll installed on the
system, which fortunately has been deployed with all Windows versions since XP.
This means if you have Go installed, you also have the DLL installed. This also
means that your application can be deployed without the DirectX DLLs and if you
have no other dependencies you can just give the executable file to the users.

To get going with this wrapper there are some basic samples that demonstrate how
to set up and use the library at:
https://github.com/gonutz/d3d9/tree/master/samples
*/
package d3d9

import (
	"errors"
	"syscall"
	"unsafe"
)

var (
	dll             = syscall.NewLazyDLL("d3d9.dll")
	direct3DCreate9 = dll.NewProc("Direct3DCreate9")
)

// Create takes the SDK version as argument and returns a new Direct3D object.
// Use SDK_VERSION as the parameter.
func Create(version uint) (*Direct3D, error) {
	obj, _, _ := direct3DCreate9.Call(uintptr(version))
	if obj == 0 {
		return nil, errors.New("Direct3DCreate9 returned nil")
	}
	return (*Direct3D)(unsafe.Pointer(obj)), nil
}

// Direct3D and its methods are used to create Direct3D objects and set up the
// environment. Its interface includes methods for enumerating and retrieving
// capabilities of the device.
type Direct3D struct {
	vtbl *direct3DVtbl
}

type direct3DVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	RegisterSoftwareDevice      uintptr
	GetAdapterCount             uintptr
	GetAdapterIdentifier        uintptr
	GetAdapterModeCount         uintptr
	EnumAdapterModes            uintptr
	GetAdapterDisplayMode       uintptr
	CheckDeviceType             uintptr
	CheckDeviceFormat           uintptr
	CheckDeviceMultiSampleType  uintptr
	CheckDepthStencilMatch      uintptr
	CheckDeviceFormatConversion uintptr
	GetDeviceCaps               uintptr
	GetAdapterMonitor           uintptr
	CreateDevice                uintptr
}

// AddRef increments the reference count for an interface on an object. This
// method should be called for every new copy of a pointer to an interface on an
// object.
func (obj *Direct3D) AddRef() uint32 {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.AddRef,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint32(ret)
}

// Release has to be called when finished using the object to free its
// associated resources.
func (obj *Direct3D) Release() uint32 {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint32(ret)
}

// RegisterSoftwareDevice registers a pluggable software device. Software
// devices provide software rasterization enabling applications to access a
// variety of software rasterizers.
func (obj *Direct3D) RegisterSoftwareDevice(initFunc uintptr) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.RegisterSoftwareDevice,
		2,
		uintptr(unsafe.Pointer(obj)),
		initFunc,
		0,
	)
	return toErr(ret)
}

// GetAdapterCount returns the number of adapters on the system.
func (obj *Direct3D) GetAdapterCount() uint {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetAdapterCount,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint(ret)
}

// GetAdapterIdentifier describes the physical display adapters present in the
// system when the IDirect3D9 interface was instantiated.
// Adapter must be a value from 0 to GetAdapterCount()-1.
func (obj *Direct3D) GetAdapterIdentifier(
	adapter uint,
	flags uint32,
) (id ADAPTER_IDENTIFIER, err Error) {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.GetAdapterIdentifier,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(adapter),
		uintptr(flags),
		uintptr(unsafe.Pointer(&id)),
		0,
		0,
	)
	err = toErr(ret)
	return
}

// GetAdapterModeCount returns the number of display modes available on this
// adapter.
func (obj *Direct3D) GetAdapterModeCount(adapter uint, format FORMAT) uint {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetAdapterModeCount,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(adapter),
		uintptr(format),
	)
	return uint(ret)
}

// EnumAdapterModes queries the device to determine whether the specified
// adapter supports the requested format and display mode. This method could be
// used in a loop to enumerate all the available adapter modes.
//
// Allowed formats are as follows:
//   FMT_A1R5G5B5
//   FMT_A2R10G10B10
//   FMT_A8R8G8B8
//   FMT_R5G6B5
//   FMT_X1R5G5B5
//   FMT_X8R8G8B8
//
// EnumAdapterModes treats pixel formats 565 and 555 as equivalent, and returns
// the correct version. The difference comes into play only when the application
// locks the back buffer and there is an explicit flag that the application must
// set in order to accomplish this.
func (obj *Direct3D) EnumAdapterModes(
	adapter uint,
	format FORMAT,
	mode uint,
) (displayMode DISPLAYMODE, err Error) {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.EnumAdapterModes,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(adapter),
		uintptr(format),
		uintptr(mode),
		uintptr(unsafe.Pointer(&displayMode)),
		0,
	)
	err = toErr(ret)
	return
}

// GetAdapterDisplayMode retrieves the current display mode of the adapter.
func (obj *Direct3D) GetAdapterDisplayMode(
	adapter uint,
) (mode DISPLAYMODE, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetAdapterDisplayMode,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(adapter),
		uintptr(unsafe.Pointer(&mode)),
	)
	err = toErr(ret)
	return
}

// CheckDeviceType verifies whether a hardware accelerated device type can be
// used on this adapter.
func (obj *Direct3D) CheckDeviceType(
	adapter uint,
	deviceType DEVTYPE,
	displayFormat FORMAT,
	backBufferFormat FORMAT,
	windowed bool,
) Error {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.CheckDeviceType,
		6,
		uintptr(unsafe.Pointer(obj)),
		uintptr(adapter),
		uintptr(deviceType),
		uintptr(displayFormat),
		uintptr(backBufferFormat),
		uintptrBool(windowed),
	)
	return toErr(ret)
}

func uintptrBool(b bool) uintptr {
	if b {
		return 1
	}
	return 0
}

// CheckDeviceFormat determines whether a surface format is available as a
// specified resource type and can be used as a texture, depth-stencil buffer,
// or render target, or any combination of the three, on a device representing
// this adapter.
func (obj *Direct3D) CheckDeviceFormat(
	adapter uint,
	deviceType DEVTYPE,
	adapterFormat FORMAT,
	usage uint32,
	resType RESOURCETYPE,
	checkFormat FORMAT,
) Error {
	ret, _, _ := syscall.Syscall9(
		obj.vtbl.CheckDeviceFormat,
		7,
		uintptr(unsafe.Pointer(obj)),
		uintptr(adapter),
		uintptr(deviceType),
		uintptr(adapterFormat),
		uintptr(usage),
		uintptr(resType),
		uintptr(checkFormat),
		0,
		0,
	)
	return toErr(ret)
}

// CheckDeviceMultiSampleType determines if a multisampling technique is
// available on this device.
func (obj *Direct3D) CheckDeviceMultiSampleType(
	adapter uint,
	deviceType DEVTYPE,
	surfaceFormat FORMAT,
	windowed bool,
	multiSampleType MULTISAMPLE_TYPE,
) (qualityLevels uint32, err Error) {
	ret, _, _ := syscall.Syscall9(
		obj.vtbl.CheckDeviceMultiSampleType,
		7,
		uintptr(unsafe.Pointer(obj)),
		uintptr(adapter),
		uintptr(deviceType),
		uintptr(surfaceFormat),
		uintptrBool(windowed),
		uintptr(multiSampleType),
		uintptr(unsafe.Pointer(&qualityLevels)),
		0,
		0,
	)
	err = toErr(ret)
	return
}

// CheckDepthStencilMatch determines whether a depth-stencil format is
// compatible with a render-target format in a particular display mode.
func (obj *Direct3D) CheckDepthStencilMatch(
	adapter uint,
	deviceType DEVTYPE,
	adapterFormat FORMAT,
	renderTargetFormat FORMAT,
	depthStencilFormat FORMAT,
) Error {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.CheckDepthStencilMatch,
		6,
		uintptr(unsafe.Pointer(obj)),
		uintptr(adapter),
		uintptr(deviceType),
		uintptr(adapterFormat),
		uintptr(renderTargetFormat),
		uintptr(depthStencilFormat),
	)
	return toErr(ret)
}

// CheckDeviceFormatConversion tests the device to see if it supports conversion
// from one display format to another.
func (obj *Direct3D) CheckDeviceFormatConversion(
	adapter uint,
	deviceType DEVTYPE,
	sourceFormat FORMAT,
	targetFormat FORMAT,
) Error {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.CheckDeviceFormatConversion,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(adapter),
		uintptr(deviceType),
		uintptr(sourceFormat),
		uintptr(targetFormat),
		0,
	)
	return toErr(ret)
}

// GetDeviceCaps retrieves device-specific information about a device.
func (obj *Direct3D) GetDeviceCaps(
	adapter uint,
	deviceType DEVTYPE,
) (caps CAPS, err Error) {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.GetDeviceCaps,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(adapter),
		uintptr(deviceType),
		uintptr(unsafe.Pointer(&caps)),
		0,
		0,
	)
	err = toErr(ret)
	return
}

// GetAdapterMonitor returns the handle of the monitor associated with the
// Direct3D object.
func (obj *Direct3D) GetAdapterMonitor(adapter uint) HMONITOR {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetAdapterMonitor,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(adapter),
		0,
	)
	return HMONITOR(ret)
}

// CreateDevice creates a device to represent the display adapter.
func (obj *Direct3D) CreateDevice(
	adapter uint,
	deviceType DEVTYPE,
	focusWindow HWND,
	behaviorFlags uint32,
	params PRESENT_PARAMETERS,
) (*Device, PRESENT_PARAMETERS, Error) {
	var device *Device
	ret, _, _ := syscall.Syscall9(
		obj.vtbl.CreateDevice,
		7,
		uintptr(unsafe.Pointer(obj)),
		uintptr(adapter),
		uintptr(deviceType),
		uintptr(focusWindow),
		uintptr(behaviorFlags),
		uintptr(unsafe.Pointer(&params)),
		uintptr(unsafe.Pointer(&device)),
		0,
		0,
	)
	return device, params, toErr(ret)
}
