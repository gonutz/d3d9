package d3d9

//#include "d3d9wrapper.h"
import "C"

const (
	// ADAPTER_DEFAULT is used to specify the primary display adapter.
	ADAPTER_DEFAULT = C.D3DADAPTER_DEFAULT
	// CAPS_READ_SCANLINE means the display hardware is capable of returning
	// the current scan line.
	CAPS_READ_SCANLINE = C.D3DCAPS_READ_SCANLINE
	// CAPS2_CANAUTOGENMIPMAP indicates that the driver is capable of
	// automatically generating mipmaps.
	CAPS2_CANAUTOGENMIPMAP = C.D3DCAPS2_CANAUTOGENMIPMAP
	// CAPS2_CANCALIBRATEGAMMA indicates that the system has a calibrator
	// installed that can automatically adjust the gamma ramp so that the
	// result is identical on all systems that have a calibrator. To invoke the
	// calibrator when setting new gamma levels, use the SGR_CALIBRATE flag
	// when calling SetGammaRamp. Calibrating gamma ramps incurs some
	// processing overhead and should not be used frequently.
	CAPS2_CANCALIBRATEGAMMA = C.D3DCAPS2_CANCALIBRATEGAMMA
	// CAPS2_CANMANAGERESOURCE indicates that the driver is capable of managing
	// resources. On such drivers, POOL_MANAGED resources will be managed by
	// the driver. To have Direct3D override the driver so that Direct3D
	// manages resources, use the CREATE_DISABLE_DRIVER_MANAGEMENT flag when
	// calling CreateDevice.
	CAPS2_CANMANAGERESOURCE = C.D3DCAPS2_CANMANAGERESOURCE
	// CAPS2_DYNAMICTEXTURES indicates that the driver supports dynamic
	// textures.
	CAPS2_DYNAMICTEXTURES = C.D3DCAPS2_DYNAMICTEXTURES
	// CAPS2_FULLSCREENGAMMA indicates that the driver supports dynamic gamma
	// ramp adjustment in full-screen mode.
	CAPS2_FULLSCREENGAMMA = C.D3DCAPS2_FULLSCREENGAMMA
	// CAPS3_ALPHA_FULLSCREEN_FLIP_OR_DISCARD indicates that the device can
	// respect the RS_ALPHABLENDENABLE render state in full-screen mode while
	// using the FLIP or DISCARD swap effect. This only applies when the
	// RS_SRCBLEND or RS_DESTBLEND states are set to one of the following:
	// BLEND_DESTALPHA, BLEND_INVDESTALPHA, BLEND_DESTCOLOR, BLEND_INVDESTCOLOR
	CAPS3_ALPHA_FULLSCREEN_FLIP_OR_DISCARD = C.D3DCAPS3_ALPHA_FULLSCREEN_FLIP_OR_DISCARD
	// CAPS3_COPY_TO_VIDMEM indicates that the device can accelerate a memory
	// copy from system memory to local video memory. This cap guarantees that
	// UpdateSurface and UpdateTexture calls will be hardware accelerated. If
	// this cap is absent, these calls will succeed but will be slower.
	CAPS3_COPY_TO_VIDMEM = C.D3DCAPS3_COPY_TO_VIDMEM
	// CAPS3_COPY_TO_SYSTEMMEM indicates that the device can accelerate a
	// memory copy from local video memory to system memory. This cap
	// guarantees that GetRenderTargetData calls will be hardware accelerated.
	// If this cap is absent, this call will succeed but will be slower.
	CAPS3_COPY_TO_SYSTEMMEM = C.D3DCAPS3_COPY_TO_SYSTEMMEM
	// CAPS3_LINEAR_TO_SRGB_PRESENTATION indicates that the device can perform
	// gamma correction from a windowed back buffer (containing linear content)
	// to an sRGB desktop.
	CAPS3_LINEAR_TO_SRGB_PRESENTATION = C.D3DCAPS3_LINEAR_TO_SRGB_PRESENTATION
	// CLEAR_STENCIL clears the stencil buffer.
	CLEAR_STENCIL = C.D3DCLEAR_STENCIL
	// CLEAR_TARGET clears a render target, or all targets in a multiple render
	// target.
	CLEAR_TARGET = C.D3DCLEAR_TARGET
	// CLEAR_ZBUFFER clears the depth buffer.
	CLEAR_ZBUFFER = C.D3DCLEAR_ZBUFFER
	// CREATE_ADAPTERGROUP_DEVICE asks the device to drive all the heads that
	// this master adapter owns. The flag is illegal on nonmaster adapters. If
	// this flag is set, the presentation parameters passed to CreateDevice
	// should point to an array of PRESENT_PARAMETERS. The number of elements
	// in PRESENT_PARAMETERS should equal the number of adapters defined by the
	// NumberOfAdaptersInGroup member of the CAPS9 structure. The DirectX
	// runtime will assign each element to each head in the numerical order
	// specified by the AdapterOrdinalInGroup member of CAPS9.
	CREATE_ADAPTERGROUP_DEVICE = C.D3DCREATE_ADAPTERGROUP_DEVICE
	// CREATE_DISABLE_DRIVER_MANAGEMENT instructs Direct3D to manage resources
	// instead of the driver. Direct3D calls will not fail for resource errors
	// such as insufficient video memory.
	CREATE_DISABLE_DRIVER_MANAGEMENT = C.D3DCREATE_DISABLE_DRIVER_MANAGEMENT
	// CREATE_DISABLE_DRIVER_MANAGEMENT_EX instructs Direct3D to manage
	// resources instead of the driver. Unlike
	// CREATE_DISABLE_DRIVER_MANAGEMENT, CREATE_DISABLE_DRIVER_MANAGEMENT_EX
	// will return errors for conditions such as insufficient video memory.
	CREATE_DISABLE_DRIVER_MANAGEMENT_EX = C.D3DCREATE_DISABLE_DRIVER_MANAGEMENT_EX
	// CREATE_DISABLE_PRINTSCREEN causes the runtime not register hotkeys for
	// Printscreen, Ctrl-Printscreen and Alt-Printscreen to capture the desktop
	// or window content.
	// This flag is available in Direct3D 9Ex only.
	CREATE_DISABLE_PRINTSCREEN = C.D3DCREATE_DISABLE_PRINTSCREEN
	// CREATE_DISABLE_PSGP_THREADING restricts computation to the main
	// application thread. If the flag is not set, the runtime may perform
	// software vertex processing and other computations in worker thread to
	// improve performance on multi-processor systems. Differences between
	// Windows XP and Windows Vista: This flag is available on Windows Vista,
	// Windows Server 2008, and Windows 7.
	CREATE_DISABLE_PSGP_THREADING = C.D3DCREATE_DISABLE_PSGP_THREADING
	// CREATE_ENABLE_PRESENTSTATS enables the gathering of present statistics
	// on the device. Calls to GetPresentStatistics will return valid data.
	// This flag is available in Direct3D 9Ex only.
	CREATE_ENABLE_PRESENTSTATS = C.D3DCREATE_ENABLE_PRESENTSTATS
	// CREATE_FPU_PRESERVE sets the precision for Direct3D floating-point
	// calculations to the precision used by the calling thread. If you do not
	// specify this flag, Direct3D defaults to single-precision
	// round-to-nearest mode for two reasons:
	// 1) Double-precision mode will reduce Direct3D performance.
	// 2) Portions of Direct3D assume floating-point unit exceptions are
	// masked; unmasking these exceptions may result in undefined behavior.
	CREATE_FPU_PRESERVE = C.D3DCREATE_FPU_PRESERVE
	// CREATE_HARDWARE_VERTEXPROCESSING specifies hardware vertex processing.
	CREATE_HARDWARE_VERTEXPROCESSING = C.D3DCREATE_HARDWARE_VERTEXPROCESSING
	// CREATE_MIXED_VERTEXPROCESSING specifies mixed (both software and
	// hardware) vertex processing.
	CREATE_MIXED_VERTEXPROCESSING = C.D3DCREATE_MIXED_VERTEXPROCESSING
	// CREATE_MULTITHREADED indicates that the application requests Direct3D to
	// be multithread safe. This makes a Direct3D thread take ownership of its
	// global critical section more frequently, which can degrade performance.
	// If an application processes window messages in one thread while making
	// Direct3D API calls in another, the application must use this flag when
	// creating the device. This window must also be destroyed before unloading
	// d3d9.dll.
	CREATE_MULTITHREADED = C.D3DCREATE_MULTITHREADED
	// CREATE_NOWINDOWCHANGES indicates that Direct3D must not alter the focus
	// window in any way.
	// Note: If this flag is set, the application must fully support all focus
	// management events, such as ALT+TAB and mouse click events.
	CREATE_NOWINDOWCHANGES = C.D3DCREATE_NOWINDOWCHANGES
	// CREATE_PUREDEVICE specifies that Direct3D does not support Get* calls
	// for anything that can be stored in state blocks. It also tells Direct3D
	// not to provide any emulation services for vertex processing. This means
	// that if the device does not support vertex processing, then the
	// application can use only post-transformed vertices.
	CREATE_PUREDEVICE = C.D3DCREATE_PUREDEVICE
	// CREATE_SCREENSAVER allows screensavers during a fullscreen application.
	// Without this flag, Direct3D will disable screensavers for as long as the
	// calling application is fullscreen. If the calling application is already
	// a screensaver, this flag has no effect.
	// This flag is available in Direct3D 9Ex only.
	CREATE_SCREENSAVER = C.D3DCREATE_SCREENSAVER
	// CREATE_SOFTWARE_VERTEXPROCESSING specifies software vertex processing.
	CREATE_SOFTWARE_VERTEXPROCESSING = C.D3DCREATE_SOFTWARE_VERTEXPROCESSING
	// CS_ALL is a combination of all clip flags.
	CS_ALL = C.D3DCS_ALL
	// CS_LEFT means all vertices are clipped by the left plane of the viewing
	// frustum.
	CS_LEFT = C.D3DCS_LEFT
	// CS_RIGHT means all vertices are clipped by the right plane of the
	// viewing frustum.
	CS_RIGHT = C.D3DCS_RIGHT
	// CS_TOP means all vertices are clipped by the top plane of the viewing
	// frustum.
	CS_TOP = C.D3DCS_TOP
	// CS_BOTTOM means all vertices are clipped by the bottom plane of the
	// viewing frustum.
	CS_BOTTOM = C.D3DCS_BOTTOM
	// CS_FRONT means all vertices are clipped by the front plane of the
	// viewing frustum.
	CS_FRONT = C.D3DCS_FRONT
	// CS_BACK means all vertices are clipped by the back plane of the viewing
	// frustum.
	CS_BACK = C.D3DCS_BACK
	// CS_PLANE0 uses application-defined clipping planes.
	CS_PLANE0 = C.D3DCS_PLANE0
	// CS_PLANE1 uses application-defined clipping planes.
	CS_PLANE1 = C.D3DCS_PLANE1
	// CS_PLANE2 uses application-defined clipping planes.
	CS_PLANE2 = C.D3DCS_PLANE2
	// CS_PLANE3 uses application-defined clipping planes.
	CS_PLANE3 = C.D3DCS_PLANE3
	// CS_PLANE4 uses application-defined clipping planes.
	CS_PLANE4 = C.D3DCS_PLANE4
	// CS_PLANE5 uses application-defined clipping planes.
	CS_PLANE5 = C.D3DCS_PLANE5
	// CURSORCAPS_COLOR indicates that the driver supports hardware color
	// cursor in at least high resolution modes (height >= 400).
	CURSORCAPS_COLOR = C.D3DCURSORCAPS_COLOR
	// CURSORCAPS_LOWRES indicates that the driver supports hardware color
	// cursor in low resolution modes (height < 400).
	CURSORCAPS_LOWRES = C.D3DCURSORCAPS_LOWRES
	// DEVCAPS_CANBLTSYSTONONLOCAL means the device supports blits from
	// system-memory textures to nonlocal video-memory textures.
	DEVCAPS_CANBLTSYSTONONLOCAL = C.D3DDEVCAPS_CANBLTSYSTONONLOCAL
	// DEVCAPS_CANRENDERAFTERFLIP means the device can queue rendering commands
	// after a page flip. Applications do not change their behavior if this
	// flag is set; this capability means that the device is relatively fast.
	DEVCAPS_CANRENDERAFTERFLIP = C.D3DDEVCAPS_CANRENDERAFTERFLIP
	// DEVCAPS_DRAWPRIMITIVES2 means the device can support at least a DirectX
	// 5-compliant driver.
	DEVCAPS_DRAWPRIMITIVES2 = C.D3DDEVCAPS_DRAWPRIMITIVES2
	// DEVCAPS_DRAWPRIMITIVES2EX means the device can support at least a
	// DirectX 7-compliant driver.
	DEVCAPS_DRAWPRIMITIVES2EX = C.D3DDEVCAPS_DRAWPRIMITIVES2EX
	// DEVCAPS_DRAWPRIMTLVERTEX means the device exports an
	// IDirect3DDevice9::DrawPrimitive-aware hal.
	DEVCAPS_DRAWPRIMTLVERTEX = C.D3DDEVCAPS_DRAWPRIMTLVERTEX
	// DEVCAPS_EXECUTESYSTEMMEMORY means the device can use execute buffers
	// from system memory.
	DEVCAPS_EXECUTESYSTEMMEMORY = C.D3DDEVCAPS_EXECUTESYSTEMMEMORY
	// DEVCAPS_EXECUTEVIDEOMEMORY means the device can use execute buffers from
	// video memory.
	DEVCAPS_EXECUTEVIDEOMEMORY = C.D3DDEVCAPS_EXECUTEVIDEOMEMORY
	// DEVCAPS_HWRASTERIZATION means the device has hardware acceleration for
	// scene rasterization.
	DEVCAPS_HWRASTERIZATION = C.D3DDEVCAPS_HWRASTERIZATION
	// DEVCAPS_HWTRANSFORMANDLIGHT means the device can support transformation
	// and lighting in hardware.
	DEVCAPS_HWTRANSFORMANDLIGHT = C.D3DDEVCAPS_HWTRANSFORMANDLIGHT
	// DEVCAPS_NPATCHES means the device supports N patches.
	DEVCAPS_NPATCHES = C.D3DDEVCAPS_NPATCHES
	// DEVCAPS_PUREDEVICE means the device can support rasterization,
	// transform, lighting, and shading in hardware.
	DEVCAPS_PUREDEVICE = C.D3DDEVCAPS_PUREDEVICE
	// DEVCAPS_QUINTICRTPATCHES means the device supports quintic Bézier
	// curves and B-splines.
	DEVCAPS_QUINTICRTPATCHES = C.D3DDEVCAPS_QUINTICRTPATCHES
	// DEVCAPS_RTPATCHES means the device supports rectangular and triangular
	// patches.
	DEVCAPS_RTPATCHES = C.D3DDEVCAPS_RTPATCHES
	// DEVCAPS_RTPATCHHANDLEZERO means the hardware architecture does not
	// require caching of any information, and uncached patches (handle zero)
	// will be drawn as efficiently as cached ones. Note that setting
	// DEVCAPS_RTPATCHHANDLEZERO does not mean that a patch with handle zero
	// can be drawn. A handle-zero patch can always be drawn whether this cap
	// is set or not.
	DEVCAPS_RTPATCHHANDLEZERO = C.D3DDEVCAPS_RTPATCHHANDLEZERO
	// DEVCAPS_SEPARATETEXTUREMEMORIES means the device is texturing from
	// separate memory pools.
	DEVCAPS_SEPARATETEXTUREMEMORIES = C.D3DDEVCAPS_SEPARATETEXTUREMEMORIES
	// DEVCAPS_TEXTURENONLOCALVIDMEM means the device can retrieve textures
	// from non-local video memory.
	DEVCAPS_TEXTURENONLOCALVIDMEM = C.D3DDEVCAPS_TEXTURENONLOCALVIDMEM
	// DEVCAPS_TEXTURESYSTEMMEMORY means the device can retrieve textures from
	// system memory.
	DEVCAPS_TEXTURESYSTEMMEMORY = C.D3DDEVCAPS_TEXTURESYSTEMMEMORY
	// DEVCAPS_TEXTUREVIDEOMEMORY means the device can retrieve textures from
	// device memory.
	DEVCAPS_TEXTUREVIDEOMEMORY = C.D3DDEVCAPS_TEXTUREVIDEOMEMORY
	// DEVCAPS_TLVERTEXSYSTEMMEMORY means the device can use buffers from
	// system memory for transformed and lit vertices.
	DEVCAPS_TLVERTEXSYSTEMMEMORY = C.D3DDEVCAPS_TLVERTEXSYSTEMMEMORY
	// DEVCAPS_TLVERTEXVIDEOMEMORY means the device can use buffers from video
	// memory for transformed and lit vertices.
	DEVCAPS_TLVERTEXVIDEOMEMORY = C.D3DDEVCAPS_TLVERTEXVIDEOMEMORY
	// DEVCAPS2_ADAPTIVETESSRTPATCH indicates that the device supports adaptive
	// tessellation of RT-patches
	DEVCAPS2_ADAPTIVETESSRTPATCH = C.D3DDEVCAPS2_ADAPTIVETESSRTPATCH
	// DEVCAPS2_ADAPTIVETESSNPATCH indicates that the device supports adaptive
	// tessellation of N-patches.
	DEVCAPS2_ADAPTIVETESSNPATCH = C.D3DDEVCAPS2_ADAPTIVETESSNPATCH
	// DEVCAPS2_CAN_STRETCHRECT_FROM_TEXTURES indicates that the device
	// supports StretchRect using a texture as the source.
	DEVCAPS2_CAN_STRETCHRECT_FROM_TEXTURES = C.D3DDEVCAPS2_CAN_STRETCHRECT_FROM_TEXTURES
	// DEVCAPS2_DMAPNPATCH indicates that the device supports displacement maps
	// for N-patches.
	DEVCAPS2_DMAPNPATCH = C.D3DDEVCAPS2_DMAPNPATCH
	// DEVCAPS2_PRESAMPLEDDMAPNPATCH indicates that the device supports
	// presampled displacement maps for N-patches.
	DEVCAPS2_PRESAMPLEDDMAPNPATCH = C.D3DDEVCAPS2_PRESAMPLEDDMAPNPATCH
	// DEVCAPS2_STREAMOFFSET indicates that the device supports stream offsets.
	DEVCAPS2_STREAMOFFSET = C.D3DDEVCAPS2_STREAMOFFSET
	// DEVCAPS2_VERTEXELEMENTSCANSHARESTREAMOFFSET means that multiple vertex
	// elements can share the same offset in a stream if
	// DEVCAPS2_VERTEXELEMENTSCANSHARESTREAMOFFSET is set by the device and the
	// vertex declaration does not have an element with DECLUSAGE_POSITIONT0.
	DEVCAPS2_VERTEXELEMENTSCANSHARESTREAMOFFSET = C.D3DDEVCAPS2_VERTEXELEMENTSCANSHARESTREAMOFFSET
	// DTCAPS_UBYTE4 is a 4D unsigned byte.
	DTCAPS_UBYTE4 = C.D3DDTCAPS_UBYTE4
	// DTCAPS_UBYTE4N is a normalized, 4D unsigned byte. Each of the four bytes
	// is normalized by dividing to 255.0.
	DTCAPS_UBYTE4N = C.D3DDTCAPS_UBYTE4N
	// DTCAPS_SHORT2N is a normalized, 2D signed short, expanded to (first
	// byte/32767.0, second byte/32767.0, 0, 1).
	DTCAPS_SHORT2N = C.D3DDTCAPS_SHORT2N
	// DTCAPS_SHORT4N is a normalized, 4D signed short, expanded to (first
	// byte/32767.0, second byte/32767.0, third byte/32767.0, fourth
	// byte/32767.0).
	DTCAPS_SHORT4N = C.D3DDTCAPS_SHORT4N
	// DTCAPS_USHORT2N is a normalized, 2D unsigned short, expanded to (first
	// byte/65535.0, second byte/65535.0, 0, 1).
	DTCAPS_USHORT2N = C.D3DDTCAPS_USHORT2N
	// DTCAPS_USHORT4N is a normalized 4D unsigned short, expanded to (first
	// byte/65535.0, second byte/65535.0, third byte/65535.0, fourth
	// byte/65535.0).
	DTCAPS_USHORT4N = C.D3DDTCAPS_USHORT4N
	// DTCAPS_UDEC3 is a 3D unsigned 10 10 10 format expanded to (value, value,
	// value, 1).
	DTCAPS_UDEC3 = C.D3DDTCAPS_UDEC3
	// DTCAPS_DEC3N is a 3D signed 10 10 10 format normalized and expanded to
	// (v[0]/511.0, v[1]/511.0, v[2]/511.0, 1).
	DTCAPS_DEC3N = C.D3DDTCAPS_DEC3N
	// DTCAPS_FLOAT16_2 is a 2D 16-bit floating point numbers.
	DTCAPS_FLOAT16_2 = C.D3DDTCAPS_FLOAT16_2
	// DTCAPS_FLOAT16_4 is a 4D 16-bit floating point numbers.
	DTCAPS_FLOAT16_4 = C.D3DDTCAPS_FLOAT16_4
	// OK_NOAUTOGEN is a success code. However, the autogeneration of mipmaps
	// is not supported for this format. This means that resource creation will
	// succeed but the mipmap levels will not be automatically generated.
	OK_NOAUTOGEN = C.D3DOK_NOAUTOGEN
	// ERR_CONFLICTINGRENDERSTATE indicates that the currently set render
	// states cannot be used together.
	ERR_CONFLICTINGRENDERSTATE = C.D3DERR_CONFLICTINGRENDERSTATE
	// ERR_CONFLICTINGTEXTUREFILTER indicates that the current texture filters
	// cannot be used together.
	ERR_CONFLICTINGTEXTUREFILTER = C.D3DERR_CONFLICTINGTEXTUREFILTER
	// ERR_CONFLICTINGTEXTUREPALETTE indicates that the current textures cannot
	// be used simultaneously.
	ERR_CONFLICTINGTEXTUREPALETTE = C.D3DERR_CONFLICTINGTEXTUREPALETTE
	// ERR_DEVICEHUNG indicates that the device that returned this code caused
	// the hardware adapter to be reset by the OS. Most applications should
	// destroy the device and quit. Applications that must continue should
	// destroy all video memory objects (surfaces, textures, state blocks etc)
	// and call Reset() to put the device in a default state. If the
	// application then continues rendering in the same way, the device will
	// return to this state.
	// Applies to Direct3D 9Ex only.
	ERR_DEVICEHUNG = C.D3DERR_DEVICEHUNG
	// ERR_DEVICELOST indicates that the device has been lost but cannot be
	// reset at this time. Therefore, rendering is not possible. A Direct3D
	// device object other than the one that returned this code caused the
	// hardware adapter to be reset by the OS. Delete all video memory objects
	// (surfaces, textures, state blocks) and call Reset() to return the device
	// to a default state. If the application continues rendering without a
	// reset, the rendering calls will succeed.
	ERR_DEVICELOST = C.D3DERR_DEVICELOST
	// ERR_DEVICENOTRESET indicates that the device has been lost but can be
	// reset at this time.
	ERR_DEVICENOTRESET = C.D3DERR_DEVICENOTRESET
	// ERR_DEVICEREMOVED indicates that he hardware adapter has been removed.
	// Application must destroy the device, do enumeration of adapters and
	// create another Direct3D device. If application continues rendering
	// without calling Reset, the rendering calls will succeed.
	// Applies to Direct3D 9Ex only.
	ERR_DEVICEREMOVED = C.D3DERR_DEVICEREMOVED
	// ERR_DRIVERINTERNALERROR Internal driver error. Applications should
	// destroy and recreate the device when receiving this error.
	ERR_DRIVERINTERNALERROR = C.D3DERR_DRIVERINTERNALERROR
	// ERR_DRIVERINVALIDCALL is not used.
	ERR_DRIVERINVALIDCALL = C.D3DERR_DRIVERINVALIDCALL
	// ERR_INVALIDCALL indicates that the method call is invalid. For example,
	// a method's parameter may not be an invalid pointer.
	ERR_INVALIDCALL = C.D3DERR_INVALIDCALL
	// ERR_INVALIDDEVICE indicates that the requested device type is not valid.
	ERR_INVALIDDEVICE = C.D3DERR_INVALIDDEVICE
	// ERR_MOREDATA indicates that there is more data available than the
	// specified buffer size can hold.
	ERR_MOREDATA = C.D3DERR_MOREDATA
	// ERR_NOTAVAILABLE indicates that this device does not support the queried
	// technique.
	ERR_NOTAVAILABLE = C.D3DERR_NOTAVAILABLE
	// ERR_NOTFOUND indicates that the requested item was not found.
	ERR_NOTFOUND = C.D3DERR_NOTFOUND
	// _OK indicates that no error occurred.
	OK = C.D3D_OK
	// ERR_OUTOFVIDEOMEMORY indicates that Direct3D does not have enough
	// display memory to perform the operation. The device is using more
	// resources in a single scene than can fit simultaneously into video
	// memory. Present, PresentEx, or CheckDeviceState can return this error.
	// Recovery is similar to ERR_DEVICEHUNG, though the application may want
	// to reduce its per-frame memory usage as well to avoid having the error
	// recur.
	ERR_OUTOFVIDEOMEMORY = C.D3DERR_OUTOFVIDEOMEMORY
	// ERR_TOOMANYOPERATIONS indicates that the application is requesting more
	// texture-filtering operations than the device supports.
	ERR_TOOMANYOPERATIONS = C.D3DERR_TOOMANYOPERATIONS
	// ERR_UNSUPPORTEDALPHAARG indicates that the device does not support a
	// specified texture-blending argument for the alpha channel.
	ERR_UNSUPPORTEDALPHAARG = C.D3DERR_UNSUPPORTEDALPHAARG
	// ERR_UNSUPPORTEDALPHAOPERATION indicates that the device does not support
	// a specified texture-blending operation for the alpha channel.
	ERR_UNSUPPORTEDALPHAOPERATION = C.D3DERR_UNSUPPORTEDALPHAOPERATION
	// ERR_UNSUPPORTEDCOLORARG indicates that the device does not support a
	// specified texture-blending argument for color values.
	ERR_UNSUPPORTEDCOLORARG = C.D3DERR_UNSUPPORTEDCOLORARG
	// ERR_UNSUPPORTEDCOLOROPERATION indicates that the device does not support
	// a specified texture-blending operation for color values.
	ERR_UNSUPPORTEDCOLOROPERATION = C.D3DERR_UNSUPPORTEDCOLOROPERATION
	// ERR_UNSUPPORTEDFACTORVALUE indicates that the device does not support
	// the specified texture factor value. Not used; provided only to support
	// older drivers.
	ERR_UNSUPPORTEDFACTORVALUE = C.D3DERR_UNSUPPORTEDFACTORVALUE
	// ERR_UNSUPPORTEDTEXTUREFILTER indicates that the device does not support
	// the specified texture filter.
	ERR_UNSUPPORTEDTEXTUREFILTER = C.D3DERR_UNSUPPORTEDTEXTUREFILTER
	// ERR_WASSTILLDRAWING indicates that the previous blit operation that is
	// transferring information to or from this surface is incomplete.
	ERR_WASSTILLDRAWING = C.D3DERR_WASSTILLDRAWING
	// ERR_WRONGTEXTUREFORMAT indicates that the pixel format of the texture
	// surface is not valid.
	ERR_WRONGTEXTUREFORMAT = C.D3DERR_WRONGTEXTUREFORMAT
	// E_FAIL indicates that an undetermined error occurred inside the Direct3D
	// subsystem.
	E_FAIL = C.E_FAIL
	// E_INVALIDARG indicates that an invalid parameter was passed to the
	// returning function.
	E_INVALIDARG = C.E_INVALIDARG
	// E_NOINTERFACE indicates that no object interface is available.
	E_NOINTERFACE = C.E_NOINTERFACE
	// E_NOTIMPL indicates that a method is not implemented.
	E_NOTIMPL = C.E_NOTIMPL
	// E_OUTOFMEMORY indicates that Direct3D could not allocate sufficient
	// memory to complete the call.
	E_OUTOFMEMORY = C.E_OUTOFMEMORY
	// S_NOT_RESIDENT indicates that at least one allocation that comprises the
	// resources is on disk. Direct3D 9Ex only.
	S_NOT_RESIDENT = C.S_NOT_RESIDENT
	// S_RESIDENT_IN_SHARED_MEMORY indicates that no allocations that comprise
	// the resources are on disk. However, at least one allocation is not in
	// GPU-accessible memory. Direct3D 9Ex only.
	S_RESIDENT_IN_SHARED_MEMORY = C.S_RESIDENT_IN_SHARED_MEMORY
	// ERR_UNSUPPORTEDOVERLAY indicates that the device does not support
	// overlay for the specified size or display mode.
	// Direct3D 9Ex under Windows 7 only.
	ERR_UNSUPPORTEDOVERLAY = C.D3DERR_UNSUPPORTEDOVERLAY
	// ERR_UNSUPPORTEDOVERLAYFORMAT indicates that the device does not support
	// overlay for the specified surface format.
	// Direct3D 9Ex under Windows 7 only.
	ERR_UNSUPPORTEDOVERLAYFORMAT = C.D3DERR_UNSUPPORTEDOVERLAYFORMAT
	// ERR_CANNOTPROTECTCONTENT indicates that the specified content cannot be
	// protected.
	// Direct3D 9Ex under Windows 7 only.
	ERR_CANNOTPROTECTCONTENT = C.D3DERR_CANNOTPROTECTCONTENT
	// ERR_UNSUPPORTEDCRYPTO indicates that the specified cryptographic
	// algorithm is not supported.
	// Direct3D 9Ex under Windows 7 only.
	ERR_UNSUPPORTEDCRYPTO = C.D3DERR_UNSUPPORTEDCRYPTO
	// ERR_PRESENT_STATISTICS_DISJOINT indicates that the present statistics
	// have no orderly sequence.
	// Direct3D 9Ex under Windows 7 only.
	ERR_PRESENT_STATISTICS_DISJOINT = C.D3DERR_PRESENT_STATISTICS_DISJOINT
	// FVFCAPS_DONOTSTRIPELEMENTS means it is preferable that vertex elements
	// not be stripped. That is, if the vertex format contains elements that
	// are not used with the current render states, there is no need to
	// regenerate the vertices. If this capability flag is not present,
	// stripping extraneous elements from the vertex format provides better
	// performance.
	FVFCAPS_DONOTSTRIPELEMENTS = C.D3DFVFCAPS_DONOTSTRIPELEMENTS
	// FVFCAPS_PSIZE means the point size is determined by either the render
	// state or the vertex data. If an FVF is used, point size can come from
	// point size data in the vertex declaration. Otherwise, point size is
	// determined by the render state RS_POINTSIZE. If the application provides
	// point size in both (the render state and the vertex declaration), the
	// vertex data overrides the render-state data.
	FVFCAPS_PSIZE = C.D3DFVFCAPS_PSIZE
	// FVFCAPS_TEXCOORDCOUNTMASK masks the low WORD of FVFCaps. These bits,
	// cast to the WORD data type, describe the total number of texture
	// coordinate sets that the device can simultaneously use for multiple
	// texture blending. (You can use up to eight texture coordinate sets for
	// any vertex, but the device can blend using only the specified number of
	// texture coordinate sets.)
	FVFCAPS_TEXCOORDCOUNTMASK = C.D3DFVFCAPS_TEXCOORDCOUNTMASK
	// FVF_DIFFUSE means the vertex format includes a diffuse color component.
	// It is a COLOR in ARGB order.
	FVF_DIFFUSE = C.D3DFVF_DIFFUSE
	// FVF_NORMAL means the vertex format includes a vertex normal vector. This
	// flag cannot be used with the FVF_XYZRHW flag. The normal consists of
	// three float32s.
	FVF_NORMAL = C.D3DFVF_NORMAL
	// FVF_PSIZE means vertex format specified in point size. This size is
	// expressed in camera space units for vertices that are not transformed
	// and lit, and in device-space units for transformed and lit vertices.
	FVF_PSIZE = C.D3DFVF_PSIZE
	// FVF_SPECULAR means the vertex format includes a specular color component.
	FVF_SPECULAR = C.D3DFVF_SPECULAR
	// FVF_XYZ means the vertex format includes the position of an
	// untransformed vertex. This flag cannot be used with the FVF_XYZRHW flag.
	FVF_XYZ = C.D3DFVF_XYZ
	// FVF_XYZRHW means the vertex format includes the position of a
	// transformed vertex. This flag cannot be used with the FVF_XYZ or
	// FVF_NORMAL flags.
	FVF_XYZRHW = C.D3DFVF_XYZRHW
	// FVF_XYZB1 means the vertex format contains position data, and 1
	// weighting (beta) value to use for multimatrix vertex blending
	// operations. Currently, Direct3D can blend with up to three weighting
	// values and four blending matrices.
	FVF_XYZB1 = C.D3DFVF_XYZB1
	// FVF_XYZB2 means the vertex format contains position data, and 2
	// weighting (beta) values to use for multimatrix vertex blending
	// operations. Currently, Direct3D can blend with up to three weighting
	// values and four blending matrices.
	FVF_XYZB2 = C.D3DFVF_XYZB2
	// FVF_XYZB3 means the vertex format contains position data, and 3
	// weighting (beta) values to use for multimatrix vertex blending
	// operations. Currently, Direct3D can blend with up to three weighting
	// values and four blending matrices.
	FVF_XYZB3 = C.D3DFVF_XYZB3
	// FVF_XYZB4 means the vertex format contains position data, and 4
	// weighting (beta) values to use for multimatrix vertex blending
	// operations. Currently, Direct3D can blend with up to three weighting
	// values and four blending matrices.
	FVF_XYZB4 = C.D3DFVF_XYZB4
	// FVF_XYZB5 means the vertex format contains position data, and 5
	// weighting (beta) values to use for multimatrix vertex blending
	// operations. Currently, Direct3D can blend with up to three weighting
	// values and four blending matrices.
	FVF_XYZB5 = C.D3DFVF_XYZB5
	// FVF_XYZW means the vertex format contains transformed and clipped (x, y,
	// z, w) data. ProcessVertices does not invoke the clipper, instead
	// outputting data in clip coordinates. This constant is designed for, and
	// can only be used with, the programmable vertex pipeline.
	FVF_XYZW = C.D3DFVF_XYZW
	// FVF_TEX0 is the number of texture coordinate sets for this vertex. The
	// actual values for these flags are not sequential.
	FVF_TEX0 = C.D3DFVF_TEX0
	// FVF_TEX1 is the number of texture coordinate sets for this vertex. The
	// actual values for these flags are not sequential.
	FVF_TEX1 = C.D3DFVF_TEX1
	// FVF_TEX2 is the number of texture coordinate sets for this vertex. The
	// actual values for these flags are not sequential.
	FVF_TEX2 = C.D3DFVF_TEX2
	// FVF_TEX3 is the number of texture coordinate sets for this vertex. The
	// actual values for these flags are not sequential.
	FVF_TEX3 = C.D3DFVF_TEX3
	// FVF_TEX4 is the number of texture coordinate sets for this vertex. The
	// actual values for these flags are not sequential.
	FVF_TEX4 = C.D3DFVF_TEX4
	// FVF_TEX5 is the number of texture coordinate sets for this vertex. The
	// actual values for these flags are not sequential.
	FVF_TEX5 = C.D3DFVF_TEX5
	// FVF_TEX6 is the number of texture coordinate sets for this vertex. The
	// actual values for these flags are not sequential.
	FVF_TEX6 = C.D3DFVF_TEX6
	// FVF_TEX7 is the number of texture coordinate sets for this vertex. The
	// actual values for these flags are not sequential.
	FVF_TEX7 = C.D3DFVF_TEX7
	// FVF_TEX8 is the number of texture coordinate sets for this vertex. The
	// actual values for these flags are not sequential.
	FVF_TEX8 = C.D3DFVF_TEX8
	// FVF_TEXTUREFORMAT1 means 1 floating point value.
	FVF_TEXTUREFORMAT1 = C.D3DFVF_TEXTUREFORMAT1
	// FVF_TEXTUREFORMAT2 means 2 floating point value.
	FVF_TEXTUREFORMAT2 = C.D3DFVF_TEXTUREFORMAT2
	// FVF_TEXTUREFORMAT3 means 3 floating point value.
	FVF_TEXTUREFORMAT3 = C.D3DFVF_TEXTUREFORMAT3
	// FVF_TEXTUREFORMAT4 means 4 floating point value.
	FVF_TEXTUREFORMAT4 = C.D3DFVF_TEXTUREFORMAT4
	// FVF_POSITION_MASK means the format has a mask for position bits.
	FVF_POSITION_MASK = C.D3DFVF_POSITION_MASK
	// FVF_TEXCOUNT_MASK means the format has a mask value for texture flag
	// bits.
	FVF_TEXCOUNT_MASK = C.D3DFVF_TEXCOUNT_MASK
	// FVF_LASTBETA_D3DCOLOR means the last beta field in the vertex position
	// data will be of type COLOR. The data in the beta fields are used with
	// matrix palette skinning to specify matrix indices.
	FVF_LASTBETA_D3DCOLOR = C.D3DFVF_LASTBETA_D3DCOLOR
	// FVF_LASTBETA_UBYTE4 means the last beta field in the vertex position
	// data will be of type UBYTE4. The data in the beta fields are used with
	// matrix palette skinning to specify matrix indices.
	FVF_LASTBETA_UBYTE4 = C.D3DFVF_LASTBETA_UBYTE4
	// FVF_TEXCOUNT_SHIFT is the number of bits by which to shift an integer
	// value that identifies the number of texture coordinates for a vertex.
	FVF_TEXCOUNT_SHIFT = C.D3DFVF_TEXCOUNT_SHIFT
	// LINECAPS_ALPHACMP means it supports alpha-test comparisons.
	LINECAPS_ALPHACMP = C.D3DLINECAPS_ALPHACMP
	// LINECAPS_ANTIALIAS means antialiased lines are supported.
	LINECAPS_ANTIALIAS = C.D3DLINECAPS_ANTIALIAS
	// LINECAPS_BLEND means it supports source-blending.
	LINECAPS_BLEND = C.D3DLINECAPS_BLEND
	// LINECAPS_FOG means it supports fog.
	LINECAPS_FOG = C.D3DLINECAPS_FOG
	// LINECAPS_TEXTURE means it supports texture-mapping.
	LINECAPS_TEXTURE = C.D3DLINECAPS_TEXTURE
	// LINECAPS_ZTEST means it supports z-buffer comparisons.
	LINECAPS_ZTEST = C.D3DLINECAPS_ZTEST
	// LOCK_DISCARD instructs the application to discard all memory within the
	// locked region. For vertex and index buffers, the entire buffer will be
	// discarded. This option is only valid when the resource is created with
	// dynamic usage.
	LOCK_DISCARD = C.D3DLOCK_DISCARD
	// LOCK_DONOTWAIT allows an application to gain back CPU cycles if the
	// driver cannot lock the surface immediately. If this flag is set and the
	// driver cannot lock the surface immediately, the lock call will return
	// ERR_WASSTILLDRAWING. This flag can only be used when locking a surface
	// created using CreateOffscreenPlainSurface, CreateRenderTarget, or
	// CreateDepthStencilSurface. This flag can also be used with a back buffer.
	LOCK_DONOTWAIT = C.D3DLOCK_DONOTWAIT
	// LOCK_NO_DIRTY_UPDATE by default, a lock on a resource adds a dirty
	// region to that resource. This option prevents any changes to the dirty
	// state of the resource. Applications should use this option when they
	// have additional information about the set of regions changed during the
	// lock operation.
	LOCK_NO_DIRTY_UPDATE = C.D3DLOCK_NO_DIRTY_UPDATE
	// LOCK_NOOVERWRITE indicates that memory that was referred to in a drawing
	// call since the last lock without this flag will not be modified during
	// the lock. This can enable optimizations when the application is
	// appending data to a resource. Specifying this flag enables the driver to
	// return immediately if the resource is in use, otherwise, the driver must
	// finish using the resource before returning from locking.
	LOCK_NOOVERWRITE = C.D3DLOCK_NOOVERWRITE
	// LOCK_NOSYSLOCK the default behavior of a video memory lock is to reserve
	// a system-wide critical section, guaranteeing that no display mode
	// changes will occur for the duration of the lock. This option causes the
	// system-wide critical section not to be held for the duration of the lock.
	// The lock operation is time consuming, but can enable the system to
	// perform other duties, such as moving the mouse cursor. This option is
	// useful for long-duration locks, such as the lock of the back buffer for
	// software rendering that would otherwise adversely affect system
	// responsiveness.
	LOCK_NOSYSLOCK = C.D3DLOCK_NOSYSLOCK
	// LOCK_READONLY means the application will not write to the buffer. This
	// enables resources stored in non-native formats to save the recompression
	// step when unlocking.
	LOCK_READONLY = C.D3DLOCK_READONLY
	// PBLENDCAPS_BLENDFACTOR means the driver supports both BLEND_BLENDFACTOR
	// and BLEND_INVBLENDFACTOR.
	PBLENDCAPS_BLENDFACTOR = C.D3DPBLENDCAPS_BLENDFACTOR
	// PBLENDCAPS_BOTHINVSRCALPHA means the source blend factor is (1 - As, 1 -
	// As, 1 - As, 1 - As) and destination blend factor is (As, As, As, As);
	// the destination blend selection is overridden.
	PBLENDCAPS_BOTHINVSRCALPHA = C.D3DPBLENDCAPS_BOTHINVSRCALPHA
	// PBLENDCAPS_BOTHSRCALPHA means the driver supports the BLEND_BOTHSRCALPHA
	// blend mode. (This blend mode is obsolete.)
	PBLENDCAPS_BOTHSRCALPHA = C.D3DPBLENDCAPS_BOTHSRCALPHA
	// PBLENDCAPS_DESTALPHA means the blend factor is (Ad, Ad, Ad, Ad).
	PBLENDCAPS_DESTALPHA = C.D3DPBLENDCAPS_DESTALPHA
	// PBLENDCAPS_DESTCOLOR means the blend factor is (Rd, Gd, Bd, Ad).
	PBLENDCAPS_DESTCOLOR = C.D3DPBLENDCAPS_DESTCOLOR
	// PBLENDCAPS_INVDESTALPHA means the blend factor is (1 - Ad, 1 - Ad, 1 -
	// Ad, 1 - Ad).
	PBLENDCAPS_INVDESTALPHA = C.D3DPBLENDCAPS_INVDESTALPHA
	// PBLENDCAPS_INVDESTCOLOR means the blend factor is (1 - Rd, 1 - Gd, 1 -
	// Bd, 1 - Ad).
	PBLENDCAPS_INVDESTCOLOR = C.D3DPBLENDCAPS_INVDESTCOLOR
	// PBLENDCAPS_INVSRCALPHA means the blend factor is (1 - As, 1 - As, 1 -
	// As, 1 - As).
	PBLENDCAPS_INVSRCALPHA = C.D3DPBLENDCAPS_INVSRCALPHA
	// PBLENDCAPS_INVSRCCOLOR means the blend factor is (1 - Rs, 1 - Gs, 1 -
	// Bs, 1 - As).
	PBLENDCAPS_INVSRCCOLOR = C.D3DPBLENDCAPS_INVSRCCOLOR
	// PBLENDCAPS_ONE means the blend factor is (1, 1, 1, 1).
	PBLENDCAPS_ONE = C.D3DPBLENDCAPS_ONE
	// PBLENDCAPS_SRCALPHA means the blend factor is (As, As, As, As).
	PBLENDCAPS_SRCALPHA = C.D3DPBLENDCAPS_SRCALPHA
	// PBLENDCAPS_SRCALPHASAT means the blend factor is (f, f, f, 1); f =
	// min(As, 1 - Ad).
	PBLENDCAPS_SRCALPHASAT = C.D3DPBLENDCAPS_SRCALPHASAT
	// PBLENDCAPS_SRCCOLOR means the blend factor is (Rs, Gs, Bs, As).
	PBLENDCAPS_SRCCOLOR = C.D3DPBLENDCAPS_SRCCOLOR
	// PBLENDCAPS_ZERO means the blend factor is (0, 0, 0, 0).
	PBLENDCAPS_ZERO = C.D3DPBLENDCAPS_ZERO
	// PCMPCAPS_ALWAYS always passes the z-test.
	PCMPCAPS_ALWAYS = C.D3DPCMPCAPS_ALWAYS
	// PCMPCAPS_EQUAL passes the z-test if the new z equals the current z.
	PCMPCAPS_EQUAL = C.D3DPCMPCAPS_EQUAL
	// PCMPCAPS_GREATER passes the z-test if the new z is greater than the
	// current z.
	PCMPCAPS_GREATER = C.D3DPCMPCAPS_GREATER
	// PCMPCAPS_GREATEREQUAL passes the z-test if the new z is greater than or
	// equal to the current z.
	PCMPCAPS_GREATEREQUAL = C.D3DPCMPCAPS_GREATEREQUAL
	// PCMPCAPS_LESS passes the z-test if the new z is less than the current z.
	PCMPCAPS_LESS = C.D3DPCMPCAPS_LESS
	// PCMPCAPS_LESSEQUAL passes the z-test if the new z is less than or equal
	// to the current z.
	PCMPCAPS_LESSEQUAL = C.D3DPCMPCAPS_LESSEQUAL
	// PCMPCAPS_NEVER always fails the z-test.
	PCMPCAPS_NEVER = C.D3DPCMPCAPS_NEVER
	// PCMPCAPS_NOTEQUAL passes the z-test if the new z does not equal the
	// current z.
	PCMPCAPS_NOTEQUAL = C.D3DPCMPCAPS_NOTEQUAL
	// PMISCCAPS_MASKZ means the device can enable and disable modification of
	// the depth buffer on pixel operations.
	PMISCCAPS_MASKZ = C.D3DPMISCCAPS_MASKZ
	// PMISCCAPS_CULLNONE means the driver does not perform triangle culling.
	// This corresponds to the CULL_NONE member of the CULL enumerated type.
	PMISCCAPS_CULLNONE = C.D3DPMISCCAPS_CULLNONE
	// PMISCCAPS_CULLCW means the driver supports clockwise triangle culling
	// through the RS_CULLMODE state. (This applies only to triangle
	// primitives.) This flag corresponds to the CULL_CW member of the CULL
	// enumerated type.
	PMISCCAPS_CULLCW = C.D3DPMISCCAPS_CULLCW
	// PMISCCAPS_CULLCCW means the driver supports counterclockwise culling
	// through the RS_CULLMODE state. (This applies only to triangle
	// primitives.) This flag corresponds to the CULL_CCW member of the CULL
	// enumerated type.
	PMISCCAPS_CULLCCW = C.D3DPMISCCAPS_CULLCCW
	// PMISCCAPS_COLORWRITEENABLE means the device supports per-channel writes
	// for the render-target color buffer through the RS_COLORWRITEENABLE state.
	PMISCCAPS_COLORWRITEENABLE = C.D3DPMISCCAPS_COLORWRITEENABLE
	// PMISCCAPS_CLIPPLANESCALEDPOINTS means the device correctly clips scaled
	// points of size greater than 1.0 to user-defined clipping planes.
	PMISCCAPS_CLIPPLANESCALEDPOINTS = C.D3DPMISCCAPS_CLIPPLANESCALEDPOINTS
	// PMISCCAPS_CLIPTLVERTS means the device clips post-transformed vertex
	// primitives.
	// Specify USAGE_DONOTCLIP when the pipeline should not do any clipping.
	// For this case, additional software clipping may need to be performed at
	// draw time, requiring the vertex buffer to be in system memory.
	PMISCCAPS_CLIPTLVERTS = C.D3DPMISCCAPS_CLIPTLVERTS
	// PMISCCAPS_TSSARGTEMP means the device supports TA for temporary register.
	PMISCCAPS_TSSARGTEMP = C.D3DPMISCCAPS_TSSARGTEMP
	// PMISCCAPS_BLENDOP means the device supports alpha-blending operations
	// other than BLENDOP_ADD.
	PMISCCAPS_BLENDOP = C.D3DPMISCCAPS_BLENDOP
	// PMISCCAPS_NULLREFERENCE is a reference device that does not render.
	PMISCCAPS_NULLREFERENCE = C.D3DPMISCCAPS_NULLREFERENCE
	// PMISCCAPS_INDEPENDENTWRITEMASKS means the device supports independent
	// write masks for multiple element textures or multiple render targets.
	PMISCCAPS_INDEPENDENTWRITEMASKS = C.D3DPMISCCAPS_INDEPENDENTWRITEMASKS
	// PMISCCAPS_PERSTAGECONSTANT means the device supports per-stage
	// constants. See TSS_CONSTANT in TEXTURESTAGESTATETYPE.
	PMISCCAPS_PERSTAGECONSTANT = C.D3DPMISCCAPS_PERSTAGECONSTANT
	// PMISCCAPS_FOGANDSPECULARALPHA means the device supports separate fog and
	// specular alpha. Many devices use the specular alpha channel to store the
	// fog factor.
	PMISCCAPS_FOGANDSPECULARALPHA = C.D3DPMISCCAPS_FOGANDSPECULARALPHA
	// PMISCCAPS_SEPARATEALPHABLEND means the device supports separate blend
	// settings for the alpha channel.
	PMISCCAPS_SEPARATEALPHABLEND = C.D3DPMISCCAPS_SEPARATEALPHABLEND
	// PMISCCAPS_MRTINDEPENDENTBITDEPTHS means the device supports different
	// bit depths for multiple render targets.
	PMISCCAPS_MRTINDEPENDENTBITDEPTHS = C.D3DPMISCCAPS_MRTINDEPENDENTBITDEPTHS
	// PMISCCAPS_MRTPOSTPIXELSHADERBLENDING means the device supports
	// post-pixel shader operations for multiple render targets.
	PMISCCAPS_MRTPOSTPIXELSHADERBLENDING = C.D3DPMISCCAPS_MRTPOSTPIXELSHADERBLENDING
	// PMISCCAPS_FOGVERTEXCLAMPED means the device clamps fog blend factor per
	// vertex.
	PMISCCAPS_FOGVERTEXCLAMPED = C.D3DPMISCCAPS_FOGVERTEXCLAMPED
	// PRASTERCAPS_ANISOTROPY means the device supports anisotropic filtering.
	PRASTERCAPS_ANISOTROPY = C.D3DPRASTERCAPS_ANISOTROPY
	// PRASTERCAPS_COLORPERSPECTIVE means the device iterates colors
	// perspective correctly.
	PRASTERCAPS_COLORPERSPECTIVE = C.D3DPRASTERCAPS_COLORPERSPECTIVE
	// PRASTERCAPS_DITHER means the device can dither to improve color
	// resolution.
	PRASTERCAPS_DITHER = C.D3DPRASTERCAPS_DITHER
	// PRASTERCAPS_DEPTHBIAS means the device supports legacy depth bias. For
	// true depth bias, see PRASTERCAPS_SLOPESCALEDEPTHBIAS.
	PRASTERCAPS_DEPTHBIAS = C.D3DPRASTERCAPS_DEPTHBIAS
	// PRASTERCAPS_FOGRANGE means the device supports range-based fog. In
	// range-based fog, the distance of an object from the viewer is used to
	// compute fog effects, not the depth of the object (that is, the
	// z-coordinate) in the scene.
	PRASTERCAPS_FOGRANGE = C.D3DPRASTERCAPS_FOGRANGE
	// PRASTERCAPS_FOGTABLE means the device calculates the fog value by
	// referring to a lookup table containing fog values that are indexed to
	// the depth of a given pixel.
	PRASTERCAPS_FOGTABLE = C.D3DPRASTERCAPS_FOGTABLE
	// PRASTERCAPS_FOGVERTEX means the device calculates the fog value during
	// the lighting operation and interpolates the fog value during
	// rasterization.
	PRASTERCAPS_FOGVERTEX = C.D3DPRASTERCAPS_FOGVERTEX
	// PRASTERCAPS_MIPMAPLODBIAS means the device supports level-of-detail bias
	// adjustments. These bias adjustments enable an application to make a
	// mipmap appear crisper or less sharp than it normally would. For more
	// information about level-of-detail bias in mipmaps, see
	// SAMP_MIPMAPLODBIAS.
	PRASTERCAPS_MIPMAPLODBIAS = C.D3DPRASTERCAPS_MIPMAPLODBIAS
	// PRASTERCAPS_MULTISAMPLE_TOGGLE means the device supports toggling
	// multisampling on and off between Device.BeginScene and Device.EndScene
	// (using RS_MULTISAMPLEANTIALIAS).
	PRASTERCAPS_MULTISAMPLE_TOGGLE = C.D3DPRASTERCAPS_MULTISAMPLE_TOGGLE
	// PRASTERCAPS_SCISSORTEST means the device supports scissor test.
	PRASTERCAPS_SCISSORTEST = C.D3DPRASTERCAPS_SCISSORTEST
	// PRASTERCAPS_SLOPESCALEDEPTHBIAS means the device performs true
	// slope-scale based depth bias. This is in contrast to the legacy style
	// depth bias.
	PRASTERCAPS_SLOPESCALEDEPTHBIAS = C.D3DPRASTERCAPS_SLOPESCALEDEPTHBIAS
	// PRASTERCAPS_WBUFFER means the device supports depth buffering using w.
	PRASTERCAPS_WBUFFER = C.D3DPRASTERCAPS_WBUFFER
	// PRASTERCAPS_WFOG means the device supports w-based fog. W-based fog is
	// used when a perspective projection matrix is specified, but affine
	// projections still use z-based fog. The system considers a projection
	// matrix that contains a nonzero value in the [3][4] element to be a
	// perspective projection matrix.
	PRASTERCAPS_WFOG = C.D3DPRASTERCAPS_WFOG
	// PRASTERCAPS_ZBUFFERLESSHSR means the device can perform hidden-surface
	// removal (HSR) without requiring the application to sort polygons and
	// without requiring the allocation of a depth-buffer. This leaves more
	// video memory for textures. The method used to perform HSR is
	// hardware-dependent and is transparent to the application.
	// Z-bufferless HSR is performed if no depth-buffer surface is associated
	// with the rendering-target surface and the depth-buffer comparison test
	// is enabled (that is, when the state value associated with the RS_ZENABLE
	// enumeration constant is set to TRUE).
	PRASTERCAPS_ZBUFFERLESSHSR = C.D3DPRASTERCAPS_ZBUFFERLESSHSR
	// PRASTERCAPS_ZFOG means the device supports z-based fog.
	PRASTERCAPS_ZFOG = C.D3DPRASTERCAPS_ZFOG
	// PRASTERCAPS_ZTEST means the device can perform z-test operations. This
	// effectively renders a primitive and indicates whether any z pixels have
	// been rendered.
	PRASTERCAPS_ZTEST = C.D3DPRASTERCAPS_ZTEST
	// PRESENT_INTERVAL_DEFAULT is nearly equivalent to PRESENT_INTERVAL_ONE.
	PRESENT_INTERVAL_DEFAULT = C.D3DPRESENT_INTERVAL_DEFAULT
	// PRESENT_INTERVAL_ONE means the driver will wait for the vertical retrace
	// period (the runtime will "beam follow" to prevent tearing). Present
	// operations will not be affected more frequently than the screen refresh;
	// the runtime will complete at most one Present operation per adapter
	// refresh period. This is equivalent to using SWAPEFFECT_COPYVSYNC in
	// DirectX 8.1. This option is always available for both windowed and
	// full-screen swap chains.
	PRESENT_INTERVAL_ONE = C.D3DPRESENT_INTERVAL_ONE
	// PRESENT_INTERVAL_TWO means the driver will wait for the vertical retrace
	// period. Present operations will not be affected more frequently than
	// every second screen refresh. Check the PresentationIntervals cap (see
	// CAPS) to see if PRESENT_INTERVAL_TWO is supported by the driver.
	PRESENT_INTERVAL_TWO = C.D3DPRESENT_INTERVAL_TWO
	// PRESENT_INTERVAL_THREE means the driver will wait for the vertical
	// retrace period. Present operations will not be affected more frequently
	// than every third screen refresh. Check the PresentationIntervals cap
	// (see CAPS) to see if PRESENT_INTERVAL_THREE is supported by the driver.
	PRESENT_INTERVAL_THREE = C.D3DPRESENT_INTERVAL_THREE
	// PRESENT_INTERVAL_FOUR means the driver will wait for the vertical
	// retrace period. Present operations will not be affected more frequently
	// than every fourth screen refresh. Check the PresentationIntervals member
	// (see CAPS) to see if PRESENT_INTERVAL_FOUR is supported by the driver.
	PRESENT_INTERVAL_FOUR = C.D3DPRESENT_INTERVAL_FOUR
	// PRESENT_INTERVAL_IMMEDIATE means the runtime updates the window client
	// area immediately and might do so more than once during the adapter
	// refresh period. This is equivalent to using SWAPEFFECT_COPY in DirectX
	// 8. Present operations might be affected immediately. This option is
	// always available for both windowed and full-screen swap chains.
	PRESENT_INTERVAL_IMMEDIATE = C.D3DPRESENT_INTERVAL_IMMEDIATE
	// PRESENT_BACK_BUFFERS_MAX is the maximum number of back buffers supported
	// in Direct3D 9.
	PRESENT_BACK_BUFFERS_MAX = C.D3DPRESENT_BACK_BUFFERS_MAX
	// PRESENTFLAG_DEVICECLIP clips a windowed Present blit into the window
	// client area, within the monitor screen area of the video adapter that
	// created the Direct3D device. PRESENTFLAG_DEVICECLIP is not valid with
	// SWAPEFFECT_FLIPEX.
	PRESENTFLAG_DEVICECLIP = C.D3DPRESENTFLAG_DEVICECLIP
	// PRESENTFLAG_DISCARD_DEPTHSTENCIL set this flag when the device or swap
	// chain is created to enable z-buffer discarding. If this flag is set, the
	// contents of the depth stencil buffer will be invalid after calling
	// either Present, or SetDepthStencilSurface with a different depth surface.
	// Discarding z-buffer data can increase performance and is driver
	// dependent. The debug runtime will enforce discarding by clearing the
	// z-buffer to some constant value after calling either Present, or
	// SetDepthStencilSurface with a different depth surface.
	// Discarding z-buffer data is illegal for all lockable formats,
	// FMT_D16_LOCKABLE and FMT_D32F_LOCKABLE. Any use of CreateDevice
	// specifying a lockable format and z-buffer discarding will fail.
	PRESENTFLAG_DISCARD_DEPTHSTENCIL = C.D3DPRESENTFLAG_DISCARD_DEPTHSTENCIL
	// PRESENTFLAG_LOCKABLE_BACKBUFFER set this flag if the application
	// requires the ability to lock the back buffer directly. Note that back
	// buffers are not lockable unless the application specifies
	// PRESENTFLAG_LOCKABLE_BACKBUFFER when calling CreateDevice or Reset.
	// Lockable back buffers incur a performance cost on some graphics hardware
	// configurations.
	// Performing a lock operation (or using UpdateSurface to write) on the
	// lockable back buffer decreases performance on many cards. In this case,
	// consider using textured triangles to move data to the back buffer.
	PRESENTFLAG_LOCKABLE_BACKBUFFER = C.D3DPRESENTFLAG_LOCKABLE_BACKBUFFER
	// PRESENTFLAG_VIDEO is a hint to the driver that the back buffers will
	// contain video data.
	PRESENTFLAG_VIDEO = C.D3DPRESENTFLAG_VIDEO
	// PSHADECAPS_ALPHAGOURAUDBLEND means the device can support an alpha
	// component for Gouraud-blended transparency (the SHADE_GOURAUD state for
	// the SHADEMODE enumerated type). In this mode, the alpha color component
	// of a primitive is provided at vertices and interpolated across a face
	// along with the other color components.
	PSHADECAPS_ALPHAGOURAUDBLEND = C.D3DPSHADECAPS_ALPHAGOURAUDBLEND
	// PSHADECAPS_COLORGOURAUDRGB means the device can support colored Gouraud
	// shading. In this mode, the per-vertex color components (red, green, and
	// blue) are interpolated across a triangle face.
	PSHADECAPS_COLORGOURAUDRGB = C.D3DPSHADECAPS_COLORGOURAUDRGB
	// PSHADECAPS_FOGGOURAUD means the device can support fog in the Gouraud
	// shading mode.
	PSHADECAPS_FOGGOURAUD = C.D3DPSHADECAPS_FOGGOURAUD
	// PSHADECAPS_SPECULARGOURAUDRGB means the device supports Gouraud shading
	// of specular highlights.
	PSHADECAPS_SPECULARGOURAUDRGB = C.D3DPSHADECAPS_SPECULARGOURAUDRGB
	// PS20_MAX_DYNAMICFLOWCONTROLDEPTH is the maximum level of nesting of
	// dynamic flow control instructions (break, breakc, ifc).
	PS20_MAX_DYNAMICFLOWCONTROLDEPTH = C.D3DPS20_MAX_DYNAMICFLOWCONTROLDEPTH
	// PS20_MIN_DYNAMICFLOWCONTROLDEPTH is the minimum level of nesting of
	// dynamic flow control instructions (break, breakc, ifc).
	PS20_MIN_DYNAMICFLOWCONTROLDEPTH = C.D3DPS20_MIN_DYNAMICFLOWCONTROLDEPTH
	// PS20_MAX_NUMTEMPS is the maximum number of supported temporary registers.
	PS20_MAX_NUMTEMPS = C.D3DPS20_MAX_NUMTEMPS
	// PS20_MIN_NUMTEMPS is the minimum number of supported temporary registers.
	PS20_MIN_NUMTEMPS = C.D3DPS20_MIN_NUMTEMPS
	// PS20_MAX_STATICFLOWCONTROLDEPTH is the maximum depth of nesting of the
	// loop - vs/rep - vs and call - vs/callnz bool - vs instructions.
	PS20_MAX_STATICFLOWCONTROLDEPTH = C.D3DPS20_MAX_STATICFLOWCONTROLDEPTH
	// PS20_MIN_STATICFLOWCONTROLDEPTH is the minimum depth of nesting of the
	// loop - vs/rep - vs and call - vs/callnz bool - vs instructions.
	PS20_MIN_STATICFLOWCONTROLDEPTH = C.D3DPS20_MIN_STATICFLOWCONTROLDEPTH
	// PS20_MAX_NUMINSTRUCTIONSLOTS is the maximum number of supported
	// instructions.
	PS20_MAX_NUMINSTRUCTIONSLOTS = C.D3DPS20_MAX_NUMINSTRUCTIONSLOTS
	// PS20_MIN_NUMINSTRUCTIONSLOTS is the minimum number of supported
	// instructions.
	PS20_MIN_NUMINSTRUCTIONSLOTS = C.D3DPS20_MIN_NUMINSTRUCTIONSLOTS
	// PTADDRESSCAPS_BORDER means the device supports setting coordinates
	// outside the range [0.0, 1.0] to the border color, as specified by the
	// SAMP_BORDERCOLOR texture-stage state.
	PTADDRESSCAPS_BORDER = C.D3DPTADDRESSCAPS_BORDER
	// PTADDRESSCAPS_CLAMP means the device can clamp textures to addresses.
	PTADDRESSCAPS_CLAMP = C.D3DPTADDRESSCAPS_CLAMP
	// PTADDRESSCAPS_INDEPENDENTUV means the device can separate the
	// texture-addressing modes of the u and v coordinates of the texture. This
	// ability corresponds to the SAMP_ADDRESSU and SAMP_ADDRESSV render-state
	// values.
	PTADDRESSCAPS_INDEPENDENTUV = C.D3DPTADDRESSCAPS_INDEPENDENTUV
	// PTADDRESSCAPS_MIRROR means the device can mirror textures to addresses.
	PTADDRESSCAPS_MIRROR = C.D3DPTADDRESSCAPS_MIRROR
	// PTADDRESSCAPS_MIRRORONCE means the device can take the absolute value of
	// the texture coordinate (thus, mirroring around 0) and then clamp to the
	// maximum value.
	PTADDRESSCAPS_MIRRORONCE = C.D3DPTADDRESSCAPS_MIRRORONCE
	// PTADDRESSCAPS_WRAP means the device can wrap textures to addresses.
	PTADDRESSCAPS_WRAP = C.D3DPTADDRESSCAPS_WRAP
	// PTEXTURECAPS_ALPHA means alpha in texture pixels is supported.
	PTEXTURECAPS_ALPHA = C.D3DPTEXTURECAPS_ALPHA
	// PTEXTURECAPS_ALPHAPALETTE means the device can draw alpha from texture
	// palettes.
	PTEXTURECAPS_ALPHAPALETTE = C.D3DPTEXTURECAPS_ALPHAPALETTE
	// PTEXTURECAPS_CUBEMAP means the device upports cube textures.
	PTEXTURECAPS_CUBEMAP = C.D3DPTEXTURECAPS_CUBEMAP
	// PTEXTURECAPS_CUBEMAP_POW2 means the device requires that cube texture
	// maps have dimensions specified as powers of two.
	PTEXTURECAPS_CUBEMAP_POW2 = C.D3DPTEXTURECAPS_CUBEMAP_POW2
	// PTEXTURECAPS_MIPCUBEMAP means the device supports mipmapped cube
	// textures.
	PTEXTURECAPS_MIPCUBEMAP = C.D3DPTEXTURECAPS_MIPCUBEMAP
	// PTEXTURECAPS_MIPMAP means the device supports mipmapped textures.
	PTEXTURECAPS_MIPMAP = C.D3DPTEXTURECAPS_MIPMAP
	// PTEXTURECAPS_MIPVOLUMEMAP means the device supports mipmapped volume
	// textures.
	PTEXTURECAPS_MIPVOLUMEMAP = C.D3DPTEXTURECAPS_MIPVOLUMEMAP
	// PTEXTURECAPS_NONPOW2CONDITIONAL means PTEXTURECAPS_POW2 is also set,
	// conditionally supports the use of 2D textures with dimensions that are
	// not powers of two. A device that exposes this capability can use such a
	// texture if all of the following requirements are met.
	// - The texture addressing mode for the texture stage is set to
	// TADDRESS_CLAMP.
	// - Texture wrapping for the texture stage is disabled (RS_WRAP n set to
	// 0).
	// - Mipmapping is not in use (use magnification filter only).
	// - Texture formats must not be FMT_DXT1 through FMT_DXT5.
	// If this flag is not set, and PTEXTURECAPS_POW2 is also not set, then
	// unconditional support is provided for 2D textures with dimensions that
	// are not powers of two.
	// A texture that is not a power of two cannot be set at a stage that will
	// be read based on a shader computation (such as the bem - ps and texm3x3
	// - ps instructions in pixel shaders versions 1_0 to 1_3). For example,
	// these textures can be used to store bumps that will be fed into texture
	// reads, but not the environment maps that are used in texbem - ps,
	// texbeml - ps, and texm3x3spec - ps. This means that a texture with
	// dimensions that are not powers of two cannot be addressed or sampled
	// using texture coordinates computed within the shader. This type of
	// operation is known as a dependent read and cannot be performed on these
	// types of textures.
	PTEXTURECAPS_NONPOW2CONDITIONAL = C.D3DPTEXTURECAPS_NONPOW2CONDITIONAL
	// PTEXTURECAPS_NOPROJECTEDBUMPENV means the device does not support a
	// projected bump-environment loopkup operation in programmable and fixed
	// function shaders.
	PTEXTURECAPS_NOPROJECTEDBUMPENV = C.D3DPTEXTURECAPS_NOPROJECTEDBUMPENV
	// PTEXTURECAPS_PERSPECTIVE means perspective correction texturing is
	// supported.
	PTEXTURECAPS_PERSPECTIVE = C.D3DPTEXTURECAPS_PERSPECTIVE
	// PTEXTURECAPS_POW2 means that if PTEXTURECAPS_NONPOW2CONDITIONAL is not
	// set, all textures must have widths and heights specified as powers of
	// two. This requirement does not apply to either cube textures or volume
	// textures.
	// If PTEXTURECAPS_NONPOW2CONDITIONAL is also set, conditionally supports
	// the use of 2D textures with dimensions that are not powers of two. See
	// PTEXTURECAPS_NONPOW2CONDITIONAL description.
	// If this flag is not set, and PTEXTURECAPS_NONPOW2CONDITIONAL is also not
	// set, then unconditional support is provided for 2D textures with
	// dimensions that are not powers of two.
	PTEXTURECAPS_POW2 = C.D3DPTEXTURECAPS_POW2
	// PTEXTURECAPS_PROJECTED means the device supports the TTFF_PROJECTED
	// texture transformation flag. When applied, the device divides
	// transformed texture coordinates by the last texture coordinate. If this
	// capability is present, then the projective divide occurs per pixel. If
	// this capability is not present, but the projective divide needs to occur
	// anyway, then it is performed on a per-vertex basis by the Direct3D
	// runtime.
	PTEXTURECAPS_PROJECTED = C.D3DPTEXTURECAPS_PROJECTED
	// PTEXTURECAPS_SQUAREONLY means all textures must be square.
	PTEXTURECAPS_SQUAREONLY = C.D3DPTEXTURECAPS_SQUAREONLY
	// PTEXTURECAPS_TEXREPEATNOTSCALEDBYSIZE means texture indices are not
	// scaled by the texture size prior to interpolation.
	PTEXTURECAPS_TEXREPEATNOTSCALEDBYSIZE = C.D3DPTEXTURECAPS_TEXREPEATNOTSCALEDBYSIZE
	// PTEXTURECAPS_VOLUMEMAP means the device supports volume textures.
	PTEXTURECAPS_VOLUMEMAP = C.D3DPTEXTURECAPS_VOLUMEMAP
	// PTEXTURECAPS_VOLUMEMAP_POW2 means the device requires that volume
	// texture maps have dimensions specified as powers of two.
	PTEXTURECAPS_VOLUMEMAP_POW2 = C.D3DPTEXTURECAPS_VOLUMEMAP_POW2
	// PTFILTERCAPS_MAGFPOINT means the device supports per-stage point-sample
	// filtering for magnifying textures. The point-sample magnification filter
	// is represented by the TEXF_POINT member of the TEXTUREFILTERTYPE
	// enumerated type.
	PTFILTERCAPS_MAGFPOINT = C.D3DPTFILTERCAPS_MAGFPOINT
	// PTFILTERCAPS_MAGFLINEAR means the device supports per-stage bilinear
	// interpolation filtering for magnifying mipmaps. The bilinear
	// interpolation mipmapping filter is represented by the TEXF_LINEAR member
	// of the TEXTUREFILTERTYPE enumerated type.
	PTFILTERCAPS_MAGFLINEAR = C.D3DPTFILTERCAPS_MAGFLINEAR
	// PTFILTERCAPS_MAGFANISOTROPIC means the device supports per-stage
	// anisotropic filtering for magnifying textures. The anisotropic
	// magnification filter is represented by the TEXF_ANISOTROPIC member of
	// the TEXTUREFILTERTYPE enumerated type.
	PTFILTERCAPS_MAGFANISOTROPIC = C.D3DPTFILTERCAPS_MAGFANISOTROPIC
	// PTFILTERCAPS_MAGFPYRAMIDALQUAD means the device supports per-stage
	// pyramidal sample filtering for magnifying textures. The pyramidal
	// magnifying filter is represented by the TEXF_PYRAMIDALQUAD member of the
	// TEXTUREFILTERTYPE enumerated type.
	PTFILTERCAPS_MAGFPYRAMIDALQUAD = C.D3DPTFILTERCAPS_MAGFPYRAMIDALQUAD
	// PTFILTERCAPS_MAGFGAUSSIANQUAD means the device supports per-stage
	// Gaussian quad filtering for magnifying textures.
	PTFILTERCAPS_MAGFGAUSSIANQUAD = C.D3DPTFILTERCAPS_MAGFGAUSSIANQUAD
	// PTFILTERCAPS_MINFPOINT means the device supports per-stage point-sample
	// filtering for minifying textures. The point-sample minification filter
	// is represented by the TEXF_POINT member of the TEXTUREFILTERTYPE
	// enumerated type.
	PTFILTERCAPS_MINFPOINT = C.D3DPTFILTERCAPS_MINFPOINT
	// PTFILTERCAPS_MINFLINEAR means the device supports per-stage linear
	// filtering for minifying textures. The linear minification filter is
	// represented by the TEXF_LINEAR member of the TEXTUREFILTERTYPE
	// enumerated type.
	PTFILTERCAPS_MINFLINEAR = C.D3DPTFILTERCAPS_MINFLINEAR
	// PTFILTERCAPS_MINFANISOTROPIC means the device supports per-stage
	// anisotropic filtering for minifying textures. The anisotropic
	// minification filter is represented by the TEXF_ANISOTROPIC member of the
	// TEXTUREFILTERTYPE enumerated type.
	PTFILTERCAPS_MINFANISOTROPIC = C.D3DPTFILTERCAPS_MINFANISOTROPIC
	// PTFILTERCAPS_MINFPYRAMIDALQUAD means the device supports per-stage
	// pyramidal sample filtering for minifying textures.
	PTFILTERCAPS_MINFPYRAMIDALQUAD = C.D3DPTFILTERCAPS_MINFPYRAMIDALQUAD
	// PTFILTERCAPS_MINFGAUSSIANQUAD means the device supports per-stage
	// Gaussian quad filtering for minifying textures.
	PTFILTERCAPS_MINFGAUSSIANQUAD = C.D3DPTFILTERCAPS_MINFGAUSSIANQUAD
	// PTFILTERCAPS_MIPFPOINT means the device supports per-stage point-sample
	// filtering for mipmaps. The point-sample mipmapping filter is represented
	// by the TEXF_POINT member of the TEXTUREFILTERTYPE enumerated type.
	PTFILTERCAPS_MIPFPOINT = C.D3DPTFILTERCAPS_MIPFPOINT
	// PTFILTERCAPS_MIPFLINEAR means the device supports per-stage bilinear
	// interpolation filtering for mipmaps. The bilinear interpolation
	// mipmapping filter is represented by the TEXF_LINEAR member of the
	// TEXTUREFILTERTYPE enumerated type.
	PTFILTERCAPS_MIPFLINEAR = C.D3DPTFILTERCAPS_MIPFLINEAR
	// SPD_IUNKNOWN is used in Resource.SetPrivateData. The data at is a
	// pointer to an IUnknown interface. The size parameter must be set to the
	// size of a pointer to IUnknown. Direct3D automatically calls IUnknown
	// through the data pointer when the private data is destroyed. Private
	// data will be destroyed by a subsequent call to Resource.SetPrivateData
	// with the same GUID, a subsequent call to Resource.FreePrivateData, or
	// when the Direct3D object is released.
	SPD_IUNKNOWN = C.D3DSPD_IUNKNOWN
	// STENCILCAPS_KEEP does not update the entry in the stencil buffer. This
	// is the default value.
	STENCILCAPS_KEEP = C.D3DSTENCILCAPS_KEEP
	// STENCILCAPS_ZERO sets the stencil-buffer entry to 0.
	STENCILCAPS_ZERO = C.D3DSTENCILCAPS_ZERO
	// STENCILCAPS_REPLACE replaces the stencil-buffer entry with reference
	// value.
	STENCILCAPS_REPLACE = C.D3DSTENCILCAPS_REPLACE
	// STENCILCAPS_INCRSAT increments the stencil-buffer entry, clamping to the
	// maximum value.
	STENCILCAPS_INCRSAT = C.D3DSTENCILCAPS_INCRSAT
	// STENCILCAPS_DECRSAT decrements the stencil-buffer entry, clamping to
	// zero.
	STENCILCAPS_DECRSAT = C.D3DSTENCILCAPS_DECRSAT
	// STENCILCAPS_INVERT inverts the bits in the stencil-buffer entry.
	STENCILCAPS_INVERT = C.D3DSTENCILCAPS_INVERT
	// STENCILCAPS_INCR increments the stencil-buffer entry, wrapping to zero
	// if the new value exceeds the maximum value.
	STENCILCAPS_INCR = C.D3DSTENCILCAPS_INCR
	// STENCILCAPS_DECR decrements the stencil-buffer entry, wrapping to the
	// maximum value if the new value is less than zero.
	STENCILCAPS_DECR = C.D3DSTENCILCAPS_DECR
	// STENCILCAPS_TWOSIDED means the device supports two-sided stencil.
	STENCILCAPS_TWOSIDED = C.D3DSTENCILCAPS_TWOSIDED
	// TA_CONSTANT selects a constant from a texture stage. The default value
	// is 0xFFFFFFFF.
	TA_CONSTANT = C.D3DTA_CONSTANT
	// TA_CURRENT means the texture argument is the result of the previous
	// blending stage. In the first texture stage (stage 0), this argument is
	// equivalent to TA_DIFFUSE. If the previous blending stage uses a bump-map
	// texture (the TOP_BUMPENVMAP operation), the system chooses the texture
	// from the stage before the bump-map texture. If s represents the current
	// texture stage and s - 1 contains a bump-map texture, this argument
	// becomes the result output by texture stage s - 2. Permissions are
	// read/write.
	TA_CURRENT = C.D3DTA_CURRENT
	// TA_DIFFUSE means the texture argument is the diffuse color interpolated
	// from vertex components during Gouraud shading. If the vertex does not
	// contain a diffuse color, the default color is 0xFFFFFFFF. Permissions
	// are read-only.
	TA_DIFFUSE = C.D3DTA_DIFFUSE
	// TA_SELECTMASK is a mask value for all arguments; not used when setting
	// texture arguments.
	TA_SELECTMASK = C.D3DTA_SELECTMASK
	// TA_SPECULAR means the texture argument is the specular color
	// interpolated from vertex components during Gouraud shading. If the
	// vertex does not contain a specular color, the default color is
	// 0xFFFFFFFF. Permissions are read-only.
	TA_SPECULAR = C.D3DTA_SPECULAR
	// TA_TEMP means the texture argument is a temporary register color for
	// read or write. TA_TEMP is supported if the PMISCCAPS_TSSARGTEMP device
	// capability is present. The default value for the register is (0.0, 0.0,
	// 0.0, 0.0). Permissions are read/write.
	TA_TEMP = C.D3DTA_TEMP
	// TA_TEXTURE means the texture argument is the texture color for this
	// texture stage. Permissions are read-only.
	TA_TEXTURE = C.D3DTA_TEXTURE
	// TA_TFACTOR means the texture argument is the texture factor set in a
	// previous call to the SetRenderState with the RS_TEXTUREFACTOR
	// render-state value. Permissions are read-only.
	TA_TFACTOR = C.D3DTA_TFACTOR
	// TA_ALPHAREPLICATE replicates the alpha information to all color channels
	// before the operation completes. This is a read modifier.
	TA_ALPHAREPLICATE = C.D3DTA_ALPHAREPLICATE
	// TA_COMPLEMENT takes the complement of the argument x, (1.0 - x). This is
	// a read modifier.
	TA_COMPLEMENT = C.D3DTA_COMPLEMENT
	// TEXOPCAPS_ADD means the TOP_ADD texture-blending operation is supported.
	TEXOPCAPS_ADD = C.D3DTEXOPCAPS_ADD
	// TEXOPCAPS_ADDSIGNED means the TOP_ADDSIGNED texture-blending operation
	// is supported.
	TEXOPCAPS_ADDSIGNED = C.D3DTEXOPCAPS_ADDSIGNED
	// TEXOPCAPS_ADDSIGNED2X means the TOP_ADDSIGNED2X texture-blending
	// operation is supported.
	TEXOPCAPS_ADDSIGNED2X = C.D3DTEXOPCAPS_ADDSIGNED2X
	// TEXOPCAPS_ADDSMOOTH means the TOP_ADDSMOOTH texture-blending operation
	// is supported.
	TEXOPCAPS_ADDSMOOTH = C.D3DTEXOPCAPS_ADDSMOOTH
	// TEXOPCAPS_BLENDCURRENTALPHA means the TOP_BLENDCURRENTALPHA
	// texture-blending operation is supported.
	TEXOPCAPS_BLENDCURRENTALPHA = C.D3DTEXOPCAPS_BLENDCURRENTALPHA
	// TEXOPCAPS_BLENDDIFFUSEALPHA means the TOP_BLENDDIFFUSEALPHA
	// texture-blending operation is supported.
	TEXOPCAPS_BLENDDIFFUSEALPHA = C.D3DTEXOPCAPS_BLENDDIFFUSEALPHA
	// TEXOPCAPS_BLENDFACTORALPHA means the TOP_BLENDFACTORALPHA
	// texture-blending operation is supported.
	TEXOPCAPS_BLENDFACTORALPHA = C.D3DTEXOPCAPS_BLENDFACTORALPHA
	// TEXOPCAPS_BLENDTEXTUREALPHA means the TOP_BLENDTEXTUREALPHA
	// texture-blending operation is supported.
	TEXOPCAPS_BLENDTEXTUREALPHA = C.D3DTEXOPCAPS_BLENDTEXTUREALPHA
	// TEXOPCAPS_BLENDTEXTUREALPHAPM means the TOP_BLENDTEXTUREALPHAPM
	// texture-blending operation is supported.
	TEXOPCAPS_BLENDTEXTUREALPHAPM = C.D3DTEXOPCAPS_BLENDTEXTUREALPHAPM
	// TEXOPCAPS_BUMPENVMAP means the TOP_BUMPENVMAP texture-blending operation
	// is supported.
	TEXOPCAPS_BUMPENVMAP = C.D3DTEXOPCAPS_BUMPENVMAP
	// TEXOPCAPS_BUMPENVMAPLUMINANCE means the TOP_BUMPENVMAPLUMINANCE
	// texture-blending operation is supported.
	TEXOPCAPS_BUMPENVMAPLUMINANCE = C.D3DTEXOPCAPS_BUMPENVMAPLUMINANCE
	// TEXOPCAPS_DISABLE means the TOP_DISABLE texture-blending operation is
	// supported.
	TEXOPCAPS_DISABLE = C.D3DTEXOPCAPS_DISABLE
	// TEXOPCAPS_DOTPRODUCT3 means the TOP_DOTPRODUCT3 texture-blending
	// operation is supported.
	TEXOPCAPS_DOTPRODUCT3 = C.D3DTEXOPCAPS_DOTPRODUCT3
	// TEXOPCAPS_LERP means the TOP_LERP texture-blending operation is
	// supported.
	TEXOPCAPS_LERP = C.D3DTEXOPCAPS_LERP
	// TEXOPCAPS_MODULATE means the TOP_MODULATE texture-blending operation is
	// supported.
	TEXOPCAPS_MODULATE = C.D3DTEXOPCAPS_MODULATE
	// TEXOPCAPS_MODULATE2X means the TOP_MODULATE2X texture-blending operation
	// is supported.
	TEXOPCAPS_MODULATE2X = C.D3DTEXOPCAPS_MODULATE2X
	// TEXOPCAPS_MODULATE4X means the TOP_MODULATE4X texture-blending operation
	// is supported.
	TEXOPCAPS_MODULATE4X = C.D3DTEXOPCAPS_MODULATE4X
	// TEXOPCAPS_MODULATEALPHA_ADDCOLOR means the TOP_MODULATEALPHA_ADDCOLOR
	// texture-blending operation is supported.
	TEXOPCAPS_MODULATEALPHA_ADDCOLOR = C.D3DTEXOPCAPS_MODULATEALPHA_ADDCOLOR
	// TEXOPCAPS_MODULATECOLOR_ADDALPHA means the TOP_MODULATECOLOR_ADDALPHA
	// texture-blending operation is supported.
	TEXOPCAPS_MODULATECOLOR_ADDALPHA = C.D3DTEXOPCAPS_MODULATECOLOR_ADDALPHA
	// TEXOPCAPS_MODULATEINVALPHA_ADDCOLOR means the
	// TOP_MODULATEINVALPHA_ADDCOLOR texture-blending operation is supported.
	TEXOPCAPS_MODULATEINVALPHA_ADDCOLOR = C.D3DTEXOPCAPS_MODULATEINVALPHA_ADDCOLOR
	// TEXOPCAPS_MODULATEINVCOLOR_ADDALPHA means the
	// TOP_MODULATEINVCOLOR_ADDALPHA texture-blending operation is supported.
	TEXOPCAPS_MODULATEINVCOLOR_ADDALPHA = C.D3DTEXOPCAPS_MODULATEINVCOLOR_ADDALPHA
	// TEXOPCAPS_MULTIPLYADD means the TOP_MULTIPLYADD texture-blending
	// operation is supported.
	TEXOPCAPS_MULTIPLYADD = C.D3DTEXOPCAPS_MULTIPLYADD
	// TEXOPCAPS_PREMODULATE means the TOP_PREMODULATE texture-blending
	// operation is supported.
	TEXOPCAPS_PREMODULATE = C.D3DTEXOPCAPS_PREMODULATE
	// TEXOPCAPS_SELECTARG1 means the TOP_SELECTARG1 texture-blending operation
	// is supported.
	TEXOPCAPS_SELECTARG1 = C.D3DTEXOPCAPS_SELECTARG1
	// TEXOPCAPS_SELECTARG2 means the TOP_SELECTARG2 texture-blending operation
	// is supported.
	TEXOPCAPS_SELECTARG2 = C.D3DTEXOPCAPS_SELECTARG2
	// TEXOPCAPS_SUBTRACT means the TOP_SUBTRACT texture-blending operation is
	// supported.
	TEXOPCAPS_SUBTRACT = C.D3DTEXOPCAPS_SUBTRACT
	// TSS_TCI_PASSTHRU uses the specified texture coordinates contained within
	// the vertex format. This value resolves to zero.
	TSS_TCI_PASSTHRU = C.D3DTSS_TCI_PASSTHRU
	// TSS_TCI_CAMERASPACENORMAL uses the vertex normal, transformed to camera
	// space, as the input texture coordinates for this stage's texture
	// transformation.
	TSS_TCI_CAMERASPACENORMAL = C.D3DTSS_TCI_CAMERASPACENORMAL
	// TSS_TCI_CAMERASPACEPOSITION uses the vertex position, transformed to
	// camera space, as the input texture coordinates for this stage's texture
	// transformation.
	TSS_TCI_CAMERASPACEPOSITION = C.D3DTSS_TCI_CAMERASPACEPOSITION
	// TSS_TCI_CAMERASPACEREFLECTIONVECTOR uses the reflection vector,
	// transformed to camera space, as the input texture coordinate for this
	// stage's texture transformation. The reflection vector is computed from
	// the input vertex position and normal vector.
	TSS_TCI_CAMERASPACEREFLECTIONVECTOR = C.D3DTSS_TCI_CAMERASPACEREFLECTIONVECTOR
	// TSS_TCI_SPHEREMAP uses the specified texture coordinates for sphere
	// mapping.
	TSS_TCI_SPHEREMAP = C.D3DTSS_TCI_SPHEREMAP
	// USAGE_AUTOGENMIPMAP means the resource will automatically generate
	// mipmaps. See Automatic Generation of Mipmaps (Direct3D 9). Automatic
	// generation of mipmaps is not supported for volume textures and depth
	// stencil surfaces/textures. This usage is not valid for a resource in
	// system memory (POOL_SYSTEMMEM).
	USAGE_AUTOGENMIPMAP = C.D3DUSAGE_AUTOGENMIPMAP
	// USAGE_DEPTHSTENCIL means the resource will be a depth stencil buffer.
	// USAGE_DEPTHSTENCIL can only be used with POOL_DEFAULT.
	USAGE_DEPTHSTENCIL = C.D3DUSAGE_DEPTHSTENCIL
	// USAGE_DMAP means the resource will be a displacement map.
	USAGE_DMAP = C.D3DUSAGE_DMAP
	// USAGE_DONOTCLIP indicates that the vertex buffer content will never
	// require clipping. When rendering with buffers that have this flag set,
	// the RS_CLIPPING render state must be set to false.
	USAGE_DONOTCLIP = C.D3DUSAGE_DONOTCLIP
	// USAGE_DYNAMIC indicates that the vertex buffer requires dynamic memory
	// use. This is useful for drivers because it enables them to decide where
	// to place the buffer. In general, static vertex buffers are placed in
	// video memory and dynamic vertex buffers are placed in AGP memory. Note
	// that there is no separate static use. If you do not specify
	// USAGE_DYNAMIC, the vertex buffer is made static. USAGE_DYNAMIC is
	// strictly enforced through the LOCK_DISCARD and LOCK_NOOVERWRITE locking
	// flags. As a result, LOCK_DISCARD and LOCK_NOOVERWRITE are valid only on
	// vertex buffers created with USAGE_DYNAMIC. They are not valid flags on
	// static vertex buffers.
	// USAGE_DYNAMIC and POOL_MANAGED are incompatible and should not be used
	// together.
	// Textures can specify USAGE_DYNAMIC. However, managed textures cannot use
	// USAGE_DYNAMIC.
	USAGE_DYNAMIC = C.D3DUSAGE_DYNAMIC
	// USAGE_NPATCHES indicates that the vertex buffer is to be used for
	// drawing N-patches.
	USAGE_NPATCHES = C.D3DUSAGE_NPATCHES
	// USAGE_POINTS indicates that the vertex or index buffer will be used for
	// drawing point sprites. The buffer will be loaded in system memory if
	// software vertex processing is needed to emulate point sprites.
	USAGE_POINTS = C.D3DUSAGE_POINTS
	// USAGE_RENDERTARGET means the resource will be a render target.
	// USAGE_RENDERTARGET can only be used with POOL_DEFAULT.
	USAGE_RENDERTARGET = C.D3DUSAGE_RENDERTARGET
	// USAGE_RTPATCHES indicates that the vertex buffer is to be used for
	// drawing high-order primitives.
	USAGE_RTPATCHES = C.D3DUSAGE_RTPATCHES
	// USAGE_SOFTWAREPROCESSING if this flag is used, vertex processing is done
	// in software. If this flag is not used, vertex processing is done in
	// hardware.
	// The USAGE_SOFTWAREPROCESSING flag can be set when mixed-mode or software
	// vertex processing (CREATE_MIXED_VERTEXPROCESSING /
	// CREATE_SOFTWARE_VERTEXPROCESSING) is enabled for that device.
	// USAGE_SOFTWAREPROCESSING must be set for buffers to be used with
	// software vertex processing in mixed mode, but it should not be set for
	// the best possible performance when using hardware index processing in
	// mixed mode (CREATE_HARDWARE_VERTEXPROCESSING). However, setting
	// USAGE_SOFTWAREPROCESSING is the only option when a single buffer is used
	// with both hardware and software vertex processing.
	// USAGE_SOFTWAREPROCESSING is allowed for mixed and software devices.
	// USAGE_SOFTWAREPROCESSING is used with CheckDeviceFormat to find out if a
	// particular texture format can be used as a vertex texture during
	// software vertex processing. If it can, the texture must be created in
	// POOL_SCRATCH.
	USAGE_SOFTWAREPROCESSING = C.D3DUSAGE_SOFTWAREPROCESSING
	// USAGE_WRITEONLY informs the system that the application writes only to
	// the vertex buffer. Using this flag enables the driver to choose the best
	// memory location for efficient write operations and rendering. Attempts
	// to read from a vertex buffer that is created with this capability will
	// fail. Buffers created with POOL_DEFAULT that do not specify
	// USAGE_WRITEONLY may suffer a severe performance penalty. USAGE_WRITEONLY
	// only affects the performance of POOL_DEFAULT buffers.
	USAGE_WRITEONLY = C.D3DUSAGE_WRITEONLY
	// USAGE_QUERY_FILTER queries the resource format to see if it supports
	// texture filter types other than TEXF_POINT (which is always supported).
	USAGE_QUERY_FILTER = C.D3DUSAGE_QUERY_FILTER
	// USAGE_QUERY_LEGACYBUMPMAP queries the resource about a legacy bump map.
	USAGE_QUERY_LEGACYBUMPMAP = C.D3DUSAGE_QUERY_LEGACYBUMPMAP
	// USAGE_QUERY_POSTPIXELSHADER_BLENDING queries the resource to verify
	// support for post pixel shader blending support. If CheckDeviceFormat
	// fails with USAGE_QUERY_POSTPIXELSHADER_BLENDING, post pixel blending
	// operations are not supported. These include alpha test, pixel fog,
	// render-target blending, color write enable, and dithering.
	USAGE_QUERY_POSTPIXELSHADER_BLENDING = C.D3DUSAGE_QUERY_POSTPIXELSHADER_BLENDING
	// USAGE_QUERY_SRGBREAD queries the resource to verify if a texture
	// supports gamma correction during a read operation.
	USAGE_QUERY_SRGBREAD = C.D3DUSAGE_QUERY_SRGBREAD
	// USAGE_QUERY_SRGBWRITE queries the resource to verify if a texture
	// supports gamma correction during a write operation.
	USAGE_QUERY_SRGBWRITE = C.D3DUSAGE_QUERY_SRGBWRITE
	// USAGE_QUERY_VERTEXTEXTURE queries the resource to verify support for
	// vertex shader texture sampling.
	USAGE_QUERY_VERTEXTEXTURE = C.D3DUSAGE_QUERY_VERTEXTEXTURE
	// USAGE_QUERY_WRAPANDMIP queries the resource to verify support for
	// texture wrapping and mip-mapping.
	USAGE_QUERY_WRAPANDMIP = C.D3DUSAGE_QUERY_WRAPANDMIP
	// VERTEXTEXTURESAMPLER0 identifies the texture sampler used by vertex
	// shaders.
	VERTEXTEXTURESAMPLER0 = C.D3DVERTEXTEXTURESAMPLER0
	// VERTEXTEXTURESAMPLER1 identifies the texture sampler used by vertex
	// shaders.
	VERTEXTEXTURESAMPLER1 = C.D3DVERTEXTEXTURESAMPLER1
	// VERTEXTEXTURESAMPLER2 identifies the texture sampler used by vertex
	// shaders.
	VERTEXTEXTURESAMPLER2 = C.D3DVERTEXTEXTURESAMPLER2
	// VERTEXTEXTURESAMPLER3 identifies the texture sampler used by vertex
	// shaders.
	VERTEXTEXTURESAMPLER3 = C.D3DVERTEXTEXTURESAMPLER3
	// VS20CAPS_PREDICATION indicates that instruction predication is
	// supported. See setp_comp - vs.
	VS20CAPS_PREDICATION = C.D3DVS20CAPS_PREDICATION
	// VS20_MAX_DYNAMICFLOWCONTROLDEPTH is the maximum level of nesting of
	// dynamic flow control instructions (break - vs, break_comp - vs, breakp -
	// vs, if_comp - vs, if_comp - vs, if pred - vs).
	VS20_MAX_DYNAMICFLOWCONTROLDEPTH = C.D3DVS20_MAX_DYNAMICFLOWCONTROLDEPTH
	// VS20_MIN_DYNAMICFLOWCONTROLDEPTH is the minimum level of nesting of
	// dynamic flow control instructions (break - vs, break_comp - vs, breakp -
	// vs, if_comp - vs, if_comp - vs, if pred - vs).
	VS20_MIN_DYNAMICFLOWCONTROLDEPTH = C.D3DVS20_MIN_DYNAMICFLOWCONTROLDEPTH
	// VS20_MAX_NUMTEMPS is the maximum number of temporary registers supported.
	VS20_MAX_NUMTEMPS = C.D3DVS20_MAX_NUMTEMPS
	// VS20_MIN_NUMTEMPS is the minimum number of temporary registers supported.
	VS20_MIN_NUMTEMPS = C.D3DVS20_MIN_NUMTEMPS
	// VS20_MAX_STATICFLOWCONTROLDEPTH is the maximum depth of nesting of the
	// loop - vs/rep - vs and call - vs/callnz bool - vs instructions.
	VS20_MAX_STATICFLOWCONTROLDEPTH = C.D3DVS20_MAX_STATICFLOWCONTROLDEPTH
	// VS20_MIN_STATICFLOWCONTROLDEPTH is the minimum depth of nesting of the
	// loop - vs/rep - vs and call - vs/callnz bool - vs instructions.
	VS20_MIN_STATICFLOWCONTROLDEPTH = C.D3DVS20_MIN_STATICFLOWCONTROLDEPTH
	// VTXPCAPS_DIRECTIONALLIGHTS means the device can do directional lights.
	VTXPCAPS_DIRECTIONALLIGHTS = C.D3DVTXPCAPS_DIRECTIONALLIGHTS
	// VTXPCAPS_LOCALVIEWER means the device can do local viewer.
	VTXPCAPS_LOCALVIEWER = C.D3DVTXPCAPS_LOCALVIEWER
	// VTXPCAPS_MATERIALSOURCE7 means the device supports the color material
	// states: RS_AMBIENTMATERIALSOURCE, RS_DIFFUSEMATERIALSOURCE,
	// RS_EMISSIVEMATERIALSOURCE, and RS_SPECULARMATERIALSOURCE.
	VTXPCAPS_MATERIALSOURCE7 = C.D3DVTXPCAPS_MATERIALSOURCE7
	// VTXPCAPS_NO_TEXGEN_NONLOCALVIEWER means the device does not support
	// texture generation in non-local viewer mode.
	VTXPCAPS_NO_TEXGEN_NONLOCALVIEWER = C.D3DVTXPCAPS_NO_TEXGEN_NONLOCALVIEWER
	// VTXPCAPS_POSITIONALLIGHTS means the device can do positional lights
	// (includes point and spot).
	VTXPCAPS_POSITIONALLIGHTS = C.D3DVTXPCAPS_POSITIONALLIGHTS
	// VTXPCAPS_TEXGEN means the device can do texgen.
	VTXPCAPS_TEXGEN = C.D3DVTXPCAPS_TEXGEN
	// VTXPCAPS_TEXGEN_SPHEREMAP means the device supports TSS_TCI_SPHEREMAP.
	VTXPCAPS_TEXGEN_SPHEREMAP = C.D3DVTXPCAPS_TEXGEN_SPHEREMAP
	// VTXPCAPS_TWEENING means the device can do vertex tweening.
	VTXPCAPS_TWEENING = C.D3DVTXPCAPS_TWEENING
	// S_OK indicates that no error occurred.
	S_OK = C.S_OK
	// S_PRESENT_OCCLUDED indicates that the presentation area is occluded.
	// Occlusion means that the presentation window is minimized or another
	// device entered the fullscreen mode on the same monitor as the
	// presentation window and the presentation window is completely on that
	// monitor. Occlusion will not occur if the client area is covered by
	// another Window.
	// Occluded applications can continue rendering and all calls will succeed,
	// but the occluded presentation window will not be updated. Preferably the
	// application should stop rendering to the presentation window using the
	// device and keep calling CheckDeviceState until S_OK or
	// S_PRESENT_MODE_CHANGED returns.
	// Direct3D 9Ex only.
	S_PRESENT_OCCLUDED = C.S_PRESENT_OCCLUDED
	// S_PRESENT_MODE_CHANGED indicates that the desktop display mode has been
	// changed. The application can continue rendering, but there might be
	// color conversion/stretching. Pick a back buffer format similar to the
	// current display mode, and call Reset to recreate the swap chains. The
	// device will leave this state after a Reset is called. Direct3D 9Ex only.
	S_PRESENT_MODE_CHANGED = C.S_PRESENT_MODE_CHANGED
	// _SDK_VERSION is must be used as the argument to Create.
	SDK_VERSION = C.D3D_SDK_VERSION
	// _MAX_SIMULTANEOUS_RENDERTARGETS is the maximum number of rendertargets.
	MAX_SIMULTANEOUS_RENDERTARGETS = C.D3D_MAX_SIMULTANEOUS_RENDERTARGETS
	// DMAPSAMPLER is 256, which is the maximum number of texture samplers.
	DMAPSAMPLER = C.D3DDMAPSAMPLER
	// DP_MAXTEXCOORD is the maximum number of texture coordinates.
	DP_MAXTEXCOORD = C.D3DDP_MAXTEXCOORD
	// MAXD3DDECLLENGTH is the maximum number of elements in a vertex
	// declaration (does not include "end" marker vertex element).
	MAXD3DDECLLENGTH = C.MAXD3DDECLLENGTH
	// MAXD3DDECLUSAGEINDEX is the maximum index (0-15) that can be used in a
	// vertex declaration.
	MAXD3DDECLUSAGEINDEX = C.MAXD3DDECLUSAGEINDEX
	// CLIPPLANE0 can be used to enable user-defined clipping plane 0. It is
	// defined as a convenience when setting values for the RS_CLIPPLANEENABLE
	// render state.
	CLIPPLANE0 = C.D3DCLIPPLANE0
	// CLIPPLANE1 can be used to enable user-defined clipping plane 1. It is
	// defined as a convenience when setting values for the RS_CLIPPLANEENABLE
	// render state.
	CLIPPLANE1 = C.D3DCLIPPLANE1
	// CLIPPLANE2 can be used to enable user-defined clipping plane 2. It is
	// defined as a convenience when setting values for the RS_CLIPPLANEENABLE
	// render state.
	CLIPPLANE2 = C.D3DCLIPPLANE2
	// CLIPPLANE3 can be used to enable user-defined clipping plane 3. It is
	// defined as a convenience when setting values for the RS_CLIPPLANEENABLE
	// render state.
	CLIPPLANE3 = C.D3DCLIPPLANE3
	// CLIPPLANE4 can be used to enable user-defined clipping plane 4. It is
	// defined as a convenience when setting values for the RS_CLIPPLANEENABLE
	// render state.
	CLIPPLANE4 = C.D3DCLIPPLANE4
	// CLIPPLANE5 can be used to enable user-defined clipping plane 5. It is
	// defined as a convenience when setting values for the RS_CLIPPLANEENABLE
	// render state.
	CLIPPLANE5 = C.D3DCLIPPLANE5
	// ISSUE_BEGIN is a value used by Issue to issue a query begin.
	ISSUE_BEGIN = C.D3DISSUE_BEGIN
	// ISSUE_END is a value used by Issue to issue a query end.
	ISSUE_END = C.D3DISSUE_END
	// GETDATA_FLUSH is the value passed to GetData to flush query data.
	GETDATA_FLUSH = C.D3DGETDATA_FLUSH
)
