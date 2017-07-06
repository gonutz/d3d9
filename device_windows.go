package d3d9

import (
	"syscall"
	"unsafe"
)

// Device and its methods are used to perform DrawPrimitive-based rendering,
// create resources, work with system-level variables, adjust gamma ramp levels,
// work with palettes and create shaders.
type Device struct {
	vtbl *deviceVtbl
}

type deviceVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	TestCooperativeLevel        uintptr
	GetAvailableTextureMem      uintptr
	EvictManagedResources       uintptr
	GetDirect3D                 uintptr
	GetDeviceCaps               uintptr
	GetDisplayMode              uintptr
	GetCreationParameters       uintptr
	SetCursorProperties         uintptr
	SetCursorPosition           uintptr
	ShowCursor                  uintptr
	CreateAdditionalSwapChain   uintptr
	GetSwapChain                uintptr
	GetNumberOfSwapChains       uintptr
	Reset                       uintptr
	Present                     uintptr
	GetBackBuffer               uintptr
	GetRasterStatus             uintptr
	SetDialogBoxMode            uintptr
	SetGammaRamp                uintptr
	GetGammaRamp                uintptr
	CreateTexture               uintptr
	CreateVolumeTexture         uintptr
	CreateCubeTexture           uintptr
	CreateVertexBuffer          uintptr
	CreateIndexBuffer           uintptr
	CreateRenderTarget          uintptr
	CreateDepthStencilSurface   uintptr
	UpdateSurface               uintptr
	UpdateTexture               uintptr
	GetRenderTargetData         uintptr
	GetFrontBufferData          uintptr
	StretchRect                 uintptr
	ColorFill                   uintptr
	CreateOffscreenPlainSurface uintptr
	SetRenderTarget             uintptr
	GetRenderTarget             uintptr
	SetDepthStencilSurface      uintptr
	GetDepthStencilSurface      uintptr
	BeginScene                  uintptr
	EndScene                    uintptr
	Clear                       uintptr
	SetTransform                uintptr
	GetTransform                uintptr
	MultiplyTransform           uintptr
	SetViewport                 uintptr
	GetViewport                 uintptr
	SetMaterial                 uintptr
	GetMaterial                 uintptr
	SetLight                    uintptr
	GetLight                    uintptr
	LightEnable                 uintptr
	GetLightEnable              uintptr
	SetClipPlane                uintptr
	GetClipPlane                uintptr
	SetRenderState              uintptr
	GetRenderState              uintptr
	CreateStateBlock            uintptr
	BeginStateBlock             uintptr
	EndStateBlock               uintptr
	SetClipStatus               uintptr
	GetClipStatus               uintptr
	GetTexture                  uintptr
	SetTexture                  uintptr
	GetTextureStageState        uintptr
	SetTextureStageState        uintptr
	GetSamplerState             uintptr
	SetSamplerState             uintptr
	ValidateDevice              uintptr
	SetPaletteEntries           uintptr
	GetPaletteEntries           uintptr
	SetCurrentTexturePalette    uintptr
	GetCurrentTexturePalette    uintptr
	SetScissorRect              uintptr
	GetScissorRect              uintptr
	SetSoftwareVertexProcessing uintptr
	GetSoftwareVertexProcessing uintptr
	SetNPatchMode               uintptr
	GetNPatchMode               uintptr
	DrawPrimitive               uintptr
	DrawIndexedPrimitive        uintptr
	DrawPrimitiveUP             uintptr
	DrawIndexedPrimitiveUP      uintptr
	ProcessVertices             uintptr
	CreateVertexDeclaration     uintptr
	SetVertexDeclaration        uintptr
	GetVertexDeclaration        uintptr
	SetFVF                      uintptr
	GetFVF                      uintptr
	CreateVertexShader          uintptr
	SetVertexShader             uintptr
	GetVertexShader             uintptr
	SetVertexShaderConstantF    uintptr
	GetVertexShaderConstantF    uintptr
	SetVertexShaderConstantI    uintptr
	GetVertexShaderConstantI    uintptr
	SetVertexShaderConstantB    uintptr
	GetVertexShaderConstantB    uintptr
	SetStreamSource             uintptr
	GetStreamSource             uintptr
	SetStreamSourceFreq         uintptr
	GetStreamSourceFreq         uintptr
	SetIndices                  uintptr
	GetIndices                  uintptr
	CreatePixelShader           uintptr
	SetPixelShader              uintptr
	GetPixelShader              uintptr
	SetPixelShaderConstantF     uintptr
	GetPixelShaderConstantF     uintptr
	SetPixelShaderConstantI     uintptr
	GetPixelShaderConstantI     uintptr
	SetPixelShaderConstantB     uintptr
	GetPixelShaderConstantB     uintptr
	DrawRectPatch               uintptr
	DrawTriPatch                uintptr
	DeletePatch                 uintptr
	CreateQuery                 uintptr
}

// AddRef increments the reference count for an interface on an object. This
// method should be called for every new copy of a pointer to an interface on an
// object.
func (obj *Device) AddRef() uint32 {
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
func (obj *Device) Release() uint32 {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint32(ret)
}

// TestCooperativeLevel reports the current cooperative-level status of the
// Direct3D device for a windowed or full-screen application.
func (obj *Device) TestCooperativeLevel() Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.TestCooperativeLevel,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return toErr(ret)
}

// GetAvailableTextureMem returns an estimate of the amount of available texture
// memory.
func (obj *Device) GetAvailableTextureMem() uint {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetAvailableTextureMem,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint(ret)
}

// EvictManagedResources evicts all managed resources, including both Direct3D
// and driver-managed resources.
func (obj *Device) EvictManagedResources() Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.EvictManagedResources,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return toErr(ret)
}

// GetDirect3D returns an interface to the instance of the Direct3D object that
// created the device.
// Call Release on the returned object when finished using it.
func (obj *Device) GetDirect3D() (d3d *Direct3D, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetDirect3D,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&d3d)),
		0,
	)
	err = toErr(ret)
	return
}

// GetDeviceCaps retrieves the capabilities of the rendering device.
func (obj *Device) GetDeviceCaps() (CAPS, Error) {
	var caps CAPS
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetDeviceCaps,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&caps)),
		0,
	)
	return caps, toErr(ret)
}

// GetDisplayMode retrieves the display mode's spatial resolution, color
// resolution, and refresh frequency.
func (obj *Device) GetDisplayMode(swapChain uint) (DISPLAYMODE, Error) {
	var mode DISPLAYMODE
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetDisplayMode,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(swapChain),
		uintptr(unsafe.Pointer(&mode)),
	)
	return mode, toErr(ret)
}

// GetCreationParameters retrieves the creation parameters of the device.
func (obj *Device) GetCreationParameters() (DEVICE_CREATION_PARAMETERS, Error) {
	var params DEVICE_CREATION_PARAMETERS
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetCreationParameters,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&params)),
		0,
	)
	return params, toErr(ret)
}

// SetCursorProperties sets properties for the cursor.
func (obj *Device) SetCursorProperties(
	xHotSpot uint,
	yHotSpot uint,
	cursorBitmap *Surface,
) Error {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.SetCursorProperties,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(xHotSpot),
		uintptr(yHotSpot),
		uintptr(unsafe.Pointer(cursorBitmap)),
		0,
		0,
	)
	return toErr(ret)
}

// SetCursorPosition sets the cursor position and update options.
func (obj *Device) SetCursorPosition(x int, y int, flags uint32) {
	syscall.Syscall6(
		obj.vtbl.SetCursorPosition,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(x),
		uintptr(y),
		uintptr(flags),
		0,
		0,
	)
}

// ShowCursor displays or hides the cursor.
func (obj *Device) ShowCursor(show bool) bool {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.ShowCursor,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptrBool(show),
		0,
	)
	return ret != 0
}

// CreateAdditionalSwapChain creates an additional swap chain for rendering
// multiple views.
func (obj *Device) CreateAdditionalSwapChain(params PRESENT_PARAMETERS) (
	*SwapChain, PRESENT_PARAMETERS, Error,
) {
	var chain *SwapChain
	ret, _, _ := syscall.Syscall(
		obj.vtbl.CreateAdditionalSwapChain,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&params)),
		uintptr(unsafe.Pointer(&chain)),
	)
	return chain, params, toErr(ret)
}

// GetSwapChain returns a pointer to a swap chain.
func (obj *Device) GetSwapChain(swapChain uint) (*SwapChain, Error) {
	var chain *SwapChain
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetSwapChain,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(swapChain),
		uintptr(unsafe.Pointer(&chain)),
	)
	return chain, toErr(ret)
}

// GetNumberOfSwapChains returns the number of implicit swap chains.
func (obj *Device) GetNumberOfSwapChains() uint {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetNumberOfSwapChains,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint(ret)
}

// Reset resets the type, size, and format of the swap chain.
func (obj *Device) Reset(params PRESENT_PARAMETERS) (PRESENT_PARAMETERS, Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Reset,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&params)),
		0,
	)
	return params, toErr(ret)
}

// Present presents the contents of the next buffer in the sequence of back
// buffers owned by the device.
func (obj *Device) Present(
	sourceRect *RECT,
	destRect *RECT,
	destWindowOverride HWND,
	dirtyRegion *RGNDATA,
) Error {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.Present,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(sourceRect)),
		uintptr(unsafe.Pointer(destRect)),
		uintptr(destWindowOverride),
		uintptr(unsafe.Pointer(dirtyRegion)),
		0,
	)
	return toErr(ret)
}

// GetBackBuffer retrieves a back buffer from the device's swap chain.
// Call Release on the returned surface when finished using it.
func (obj *Device) GetBackBuffer(
	swapChain uint,
	backBuffer uint,
	typ BACKBUFFER_TYPE,
) (*Surface, Error) {
	var surface *Surface
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.GetBackBuffer,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(swapChain),
		uintptr(backBuffer),
		uintptr(typ),
		uintptr(unsafe.Pointer(&surface)),
		0,
	)
	return surface, toErr(ret)
}

// GetRasterStatus returns information describing the raster of the monitor on
// which the swap chain is presented.
func (obj *Device) GetRasterStatus(swapChain uint) (RASTER_STATUS, Error) {
	var status RASTER_STATUS
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetRasterStatus,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(swapChain),
		uintptr(unsafe.Pointer(&status)),
	)
	return status, toErr(ret)
}

// SetDialogBoxMode allows the use of GDI dialog boxes in full-screen mode
// applications.
func (obj *Device) SetDialogBoxMode(enableDialogs bool) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.SetDialogBoxMode,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptrBool(enableDialogs),
		0,
	)
	return toErr(ret)
}

// SetGammaRamp sets the gamma correction ramp for the implicit swap chain. This
// method will affect the entire screen (not just the active window if you are
// running in windowed mode).
func (obj *Device) SetGammaRamp(swapChain uint, flags uint32, ramp GAMMARAMP) {
	syscall.Syscall6(
		obj.vtbl.SetGammaRamp,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(swapChain),
		uintptr(flags),
		uintptr(unsafe.Pointer(&ramp)),
		0,
		0,
	)
}

// GetGammaRamp retrieves the gamma correction ramp for the swap chain.
func (obj *Device) GetGammaRamp(swapChain uint) GAMMARAMP {
	var ramp GAMMARAMP
	syscall.Syscall(
		obj.vtbl.GetGammaRamp,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(swapChain),
		uintptr(unsafe.Pointer(&ramp)),
	)
	return ramp
}

// CreateTexture creates a texture resource.
func (obj *Device) CreateTexture(
	width uint,
	height uint,
	levels uint,
	usage uint32,
	format FORMAT,
	pool POOL,
	sharedHandle uintptr,
) (*Texture, Error) {
	var tex *Texture
	ret, _, _ := syscall.Syscall9(
		obj.vtbl.CreateTexture,
		9,
		uintptr(unsafe.Pointer(obj)),
		uintptr(width),
		uintptr(height),
		uintptr(levels),
		uintptr(usage),
		uintptr(format),
		uintptr(pool),
		uintptr(unsafe.Pointer(&tex)),
		sharedHandle,
	)
	return tex, toErr(ret)
}

// CreateVolumeTexture creates a volume texture resource.
func (obj *Device) CreateVolumeTexture(
	width uint,
	height uint,
	depth uint,
	levels uint,
	usage uint32,
	format FORMAT,
	pool POOL,
	sharedHandle uintptr,
) (*VolumeTexture, Error) {
	var tex *VolumeTexture
	ret, _, _ := syscall.Syscall12(
		obj.vtbl.CreateVolumeTexture,
		10,
		uintptr(unsafe.Pointer(obj)),
		uintptr(width),
		uintptr(height),
		uintptr(depth),
		uintptr(levels),
		uintptr(usage),
		uintptr(format),
		uintptr(pool),
		uintptr(unsafe.Pointer(&tex)),
		sharedHandle,
		0,
		0,
	)
	return tex, toErr(ret)
}

// CreateCubeTexture creates a cube texture resource.
func (obj *Device) CreateCubeTexture(
	edgeLength uint,
	levels uint,
	usage uint32,
	format FORMAT,
	pool POOL,
	sharedHandle uintptr,
) (*CubeTexture, Error) {
	var tex *CubeTexture
	ret, _, _ := syscall.Syscall9(
		obj.vtbl.CreateCubeTexture,
		8,
		uintptr(unsafe.Pointer(obj)),
		uintptr(edgeLength),
		uintptr(levels),
		uintptr(usage),
		uintptr(format),
		uintptr(pool),
		uintptr(unsafe.Pointer(&tex)),
		sharedHandle,
		0,
	)
	return tex, toErr(ret)
}

// CreateVertexBuffer creates a vertex buffer.
func (obj *Device) CreateVertexBuffer(
	length uint,
	usage uint32,
	fvf uint32,
	pool POOL,
	sharedHandle uintptr,
) (*VertexBuffer, Error) {
	var buf *VertexBuffer
	ret, _, _ := syscall.Syscall9(
		obj.vtbl.CreateVertexBuffer,
		7,
		uintptr(unsafe.Pointer(obj)),
		uintptr(length),
		uintptr(usage),
		uintptr(fvf),
		uintptr(pool),
		uintptr(unsafe.Pointer(&buf)),
		sharedHandle,
		0,
		0,
	)
	return buf, toErr(ret)
}

// CreateIndexBuffer creates an index buffer.
func (obj *Device) CreateIndexBuffer(
	length uint,
	usage uint32,
	format FORMAT,
	pool POOL,
	sharedHandle uintptr,
) (*IndexBuffer, Error) {
	var buf *IndexBuffer
	ret, _, _ := syscall.Syscall9(
		obj.vtbl.CreateIndexBuffer,
		7,
		uintptr(unsafe.Pointer(obj)),
		uintptr(length),
		uintptr(usage),
		uintptr(format),
		uintptr(pool),
		uintptr(unsafe.Pointer(&buf)),
		sharedHandle,
		0,
		0,
	)
	return buf, toErr(ret)
}

// CreateRenderTarget creates a render-target surface.
func (obj *Device) CreateRenderTarget(
	width uint,
	height uint,
	format FORMAT,
	multiSample MULTISAMPLE_TYPE,
	multisampleQuality uint32,
	lockable bool,
	sharedHandle uintptr,
) (*Surface, Error) {
	var surface *Surface
	ret, _, _ := syscall.Syscall9(
		obj.vtbl.CreateRenderTarget,
		9,
		uintptr(unsafe.Pointer(obj)),
		uintptr(width),
		uintptr(height),
		uintptr(format),
		uintptr(multiSample),
		uintptr(multisampleQuality),
		uintptrBool(lockable),
		uintptr(unsafe.Pointer(&surface)),
		sharedHandle,
	)
	return surface, toErr(ret)
}

// CreateDepthStencilSurface creates a depth-stencil resource.
func (obj *Device) CreateDepthStencilSurface(
	width uint,
	height uint,
	format FORMAT,
	multiSample MULTISAMPLE_TYPE,
	multisampleQuality uint32,
	discard bool,
	sharedHandle uintptr,
) (*Surface, Error) {
	var surface *Surface
	ret, _, _ := syscall.Syscall9(
		obj.vtbl.CreateDepthStencilSurface,
		9,
		uintptr(unsafe.Pointer(obj)),
		uintptr(width),
		uintptr(height),
		uintptr(format),
		uintptr(multiSample),
		uintptr(multisampleQuality),
		uintptrBool(discard),
		uintptr(unsafe.Pointer(&surface)),
		sharedHandle,
	)
	return surface, toErr(ret)
}

// UpdateSurface copies rectangular subsets of pixels from one surface to another.
func (obj *Device) UpdateSurface(
	sourceSurface *Surface,
	sourceRect *RECT,
	destSurface *Surface,
	destPoint *POINT,
) Error {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.UpdateSurface,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(sourceSurface)),
		uintptr(unsafe.Pointer(sourceRect)),
		uintptr(unsafe.Pointer(destSurface)),
		uintptr(unsafe.Pointer(destPoint)),
		0,
	)
	return toErr(ret)
}

// UpdateTexture updates the dirty portions of a texture.
func (obj *Device) UpdateTexture(sourceTexture, destTexture *BaseTexture) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.UpdateTexture,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(sourceTexture)),
		uintptr(unsafe.Pointer(destTexture)),
	)
	return toErr(ret)
}

// GetRenderTargetData copies the render-target data from device memory to
// system memory.
func (obj *Device) GetRenderTargetData(renderTarget, destSurface *Surface) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetRenderTargetData,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(renderTarget)),
		uintptr(unsafe.Pointer(destSurface)),
	)
	return toErr(ret)
}

// GetFrontBufferData generates a copy of the device's front buffer and places
// that copy in a system memory buffer provided by the application.
func (obj *Device) GetFrontBufferData(swapChain uint, destSurface *Surface) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetFrontBufferData,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(swapChain),
		uintptr(unsafe.Pointer(destSurface)),
	)
	return toErr(ret)
}

// StretchRect copies the contents of the source rectangle to the destination
// rectangle. The source rectangle can be stretched and filtered by the copy.
// This function is often used to change the aspect ratio of a video stream.
func (obj *Device) StretchRect(
	sourceSurface *Surface,
	sourceRect *RECT,
	destSurface *Surface,
	destRect *RECT,
	filter TEXTUREFILTERTYPE,
) Error {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.StretchRect,
		6,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(sourceSurface)),
		uintptr(unsafe.Pointer(sourceRect)),
		uintptr(unsafe.Pointer(destSurface)),
		uintptr(unsafe.Pointer(destRect)),
		uintptr(filter),
	)
	return toErr(ret)
}

// ColorFill allows an application to fill a rectangular area of a
// POOL_DEFAULT surface with a specified color.
func (obj *Device) ColorFill(
	surface *Surface,
	rect *RECT,
	color COLOR,
) Error {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.ColorFill,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(surface)),
		uintptr(unsafe.Pointer(rect)),
		uintptr(color),
		0,
		0,
	)
	return toErr(ret)
}

// CreateOffscreenPlainSurface creates an off-screen surface.
func (obj *Device) CreateOffscreenPlainSurface(
	width uint,
	height uint,
	format FORMAT,
	pool POOL,
	sharedHandle uintptr,
) (*Surface, Error) {
	var surface *Surface
	ret, _, _ := syscall.Syscall9(
		obj.vtbl.CreateOffscreenPlainSurface,
		7,
		uintptr(unsafe.Pointer(obj)),
		uintptr(width),
		uintptr(height),
		uintptr(format),
		uintptr(pool),
		uintptr(unsafe.Pointer(&surface)),
		sharedHandle,
		0,
		0,
	)
	return surface, toErr(ret)
}

// SetRenderTarget sets a new color buffer for the device.
func (obj *Device) SetRenderTarget(
	renderTargetIndex uint32,
	renderTarget *Surface,
) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.SetRenderTarget,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(renderTargetIndex),
		uintptr(unsafe.Pointer(renderTarget)),
	)
	return toErr(ret)
}

// GetRenderTarget retrieves a render-target surface.
func (obj *Device) GetRenderTarget(renderTargetIndex uint32) (*Surface, Error) {
	var surface *Surface
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetRenderTarget,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(renderTargetIndex),
		uintptr(unsafe.Pointer(surface)),
	)
	return surface, toErr(ret)
}

// SetDepthStencilSurface sets the depth stencil surface.
func (obj *Device) SetDepthStencilSurface(newZStencil *Surface) (err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.SetDepthStencilSurface,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(newZStencil)),
		0,
	)
	return toErr(ret)
}

// GetDepthStencilSurface returns the depth-stencil surface owned by the
// Direct3DDevice object.
// Call Release on the returned surface when finished using it.
func (obj *Device) GetDepthStencilSurface() (*Surface, Error) {
	var surface *Surface
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetDepthStencilSurface,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&surface)),
		0,
	)
	return surface, toErr(ret)
}

// BeginScene begins a scene.
// Applications must call BeginScene before performing any rendering and must
// call EndScene when rendering is complete and before calling BeginScene again.
func (obj *Device) BeginScene() Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.BeginScene,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return toErr(ret)
}

// EndScene ends a scene that was begun by calling BeginScene.
func (obj *Device) EndScene() (err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.EndScene,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return toErr(ret)
}

// Clear clears one or more surfaces such as a render target, multiple render
// targets, a stencil buffer, and a depth buffer.
func (obj *Device) Clear(
	rects []RECT,
	flags uint32,
	color COLOR,
	z float32,
	stencil uint32,
) Error {
	var rectPtr *RECT
	if len(rects) > 0 {
		rectPtr = &rects[0]
	}
	ret, _, _ := syscall.Syscall9(
		obj.vtbl.Clear,
		7,
		uintptr(unsafe.Pointer(obj)),
		uintptr(len(rects)),
		uintptr(unsafe.Pointer(rectPtr)),
		uintptr(flags),
		uintptr(color),
		uintptr(z),
		uintptr(stencil),
		0,
		0,
	)
	return toErr(ret)
}

// SetTransform sets a single device transformation-related state.
func (obj *Device) SetTransform(state TRANSFORMSTATETYPE, matrix MATRIX) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.SetTransform,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(state),
		uintptr(unsafe.Pointer(&matrix[0])),
	)
	return toErr(ret)
}

// GetTransform retrieves a matrix describing a transformation state.
func (obj *Device) GetTransform(state TRANSFORMSTATETYPE) (m MATRIX, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetTransform,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(state),
		uintptr(unsafe.Pointer(&m[0])),
	)
	err = toErr(ret)
	return
}

// MultiplyTransform multiplies a device's world, view, or projection matrices
// by a specified matrix.
func (obj *Device) MultiplyTransform(state TRANSFORMSTATETYPE, m MATRIX) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.MultiplyTransform,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(state),
		uintptr(unsafe.Pointer(&m[0])),
	)
	return toErr(ret)
}

// SetViewport sets the viewport parameters for the device.
func (obj *Device) SetViewport(v VIEWPORT) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.SetViewport,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&v)),
		0,
	)
	return toErr(ret)
}

// GetViewport retrieves the viewport parameters currently set for the device.
func (obj *Device) GetViewport() (v VIEWPORT, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetViewport,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&v)),
		0,
	)
	err = toErr(ret)
	return
}

// SetMaterial sets the material properties for the device.
func (obj *Device) SetMaterial(m MATERIAL) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.SetMaterial,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&m)),
		0,
	)
	return toErr(ret)
}

// GetMaterial retrieves the current material properties for the device.
func (obj *Device) GetMaterial() (m MATERIAL, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetMaterial,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&m)),
		0,
	)
	err = toErr(ret)
	return
}

// SetLight assigns a set of lighting properties for this device.
func (obj *Device) SetLight(index uint32, light LIGHT) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.SetLight,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(index),
		uintptr(unsafe.Pointer(&light)),
	)
	return toErr(ret)
}

// GetLight retrieves a set of lighting properties that this device uses.
func (obj *Device) GetLight(index uint32) (light LIGHT, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetLight,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(index),
		uintptr(unsafe.Pointer(&light)),
	)
	err = toErr(ret)
	return
}

// LightEnable enables or disables a set of lighting parameters within a device.
func (obj *Device) LightEnable(lightIndex uint32, enable bool) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.LightEnable,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(lightIndex),
		uintptrBool(enable),
	)
	return toErr(ret)
}

// GetLightEnable retrieves the activity status - enabled or disabled - for a
// set of lighting parameters within a device.
func (obj *Device) GetLightEnable(index uint32) (bool, Error) {
	var e int32
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetLightEnable,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(index),
		uintptr(unsafe.Pointer(&e)),
	)
	return e != 0, toErr(ret)
}

// SetClipPlane sets the coefficients of a user-defined clipping plane for the
// device.
func (obj *Device) SetClipPlane(index uint32, plane [4]float32) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.SetClipPlane,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(index),
		uintptr(unsafe.Pointer(&plane[0])),
	)
	return toErr(ret)
}

// GetClipPlane retrieves the coefficients of a user-defined clipping plane for
// the device.
func (obj *Device) GetClipPlane(index uint32) (plane [4]float32, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetClipPlane,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(index),
		uintptr(unsafe.Pointer(&plane[0])),
	)
	err = toErr(ret)
	return
}

// SetRenderState sets a single device render-state parameter.
func (obj *Device) SetRenderState(state RENDERSTATETYPE, value uint32) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.SetRenderState,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(state),
		uintptr(value),
	)
	return toErr(ret)
}

func (obj *Device) SetRenderStateFloat(state RENDERSTATETYPE, value float32) Error {
	return obj.SetRenderState(state, *((*uint32)(unsafe.Pointer(&value))))
}

func (obj *Device) SetRenderStateBool(state RENDERSTATETYPE, value bool) Error {
	if value {
		return obj.SetRenderState(state, 1)
	} else {
		return obj.SetRenderState(state, 0)
	}
}

// GetRenderState retrieves a render-state value for a device.
func (obj *Device) GetRenderState(state RENDERSTATETYPE) (value uint32, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetRenderState,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(state),
		uintptr(unsafe.Pointer(&value)),
	)
	err = toErr(ret)
	return
}

// CreateStateBlock creates a new state block that contains the values for all
// device states, vertex-related states, or pixel-related states.
func (obj *Device) CreateStateBlock(typ STATEBLOCKTYPE) (*StateBlock, Error) {
	var block *StateBlock
	ret, _, _ := syscall.Syscall(
		obj.vtbl.CreateStateBlock,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(typ),
		uintptr(unsafe.Pointer(&block)),
	)
	return block, toErr(ret)
}

// BeginStateBlock signals Direct3D to begin recording a device-state block.
func (obj *Device) BeginStateBlock() Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.BeginStateBlock,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return toErr(ret)
}

// EndStateBlock signals Direct3D to stop recording a device-state block and
// retrieve a pointer to the state block interface.
func (obj *Device) EndStateBlock() (*StateBlock, Error) {
	var block *StateBlock
	ret, _, _ := syscall.Syscall(
		obj.vtbl.EndStateBlock,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&block)),
		0,
	)
	return block, toErr(ret)
}

// SetClipStatus sets the clip status.
func (obj *Device) SetClipStatus(clipStatus CLIPSTATUS) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.SetClipStatus,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&clipStatus)),
		0,
	)
	return toErr(ret)
}

// GetClipStatus retrieves the clip status.
func (obj *Device) GetClipStatus() (clipStatus CLIPSTATUS, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetClipStatus,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&clipStatus)),
		0,
	)
	err = toErr(ret)
	return
}

// GetTexture retrieves a texture assigned to a stage for a device.
func (obj *Device) GetTexture(stage uint32) (*BaseTexture, Error) {
	var tex *BaseTexture
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetTexture,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(stage),
		uintptr(unsafe.Pointer(&tex)),
	)
	return tex, toErr(ret)
}

type BaseTextureImpl interface {
	baseTexturePointer() uintptr
}

// SetTexture assigns a texture to a stage for a device.
func (obj *Device) SetTexture(sampler uint32, texture BaseTextureImpl) Error {
	var base uintptr
	if texture != nil {
		base = texture.baseTexturePointer()
	}
	ret, _, _ := syscall.Syscall(
		obj.vtbl.SetTexture,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(sampler),
		base,
	)
	return toErr(ret)
}

// GetTextureStageState retrieves a state value for an assigned texture.
func (obj *Device) GetTextureStageState(
	stage uint32,
	typ TEXTURESTAGESTATETYPE,
) (value uint32, err Error) {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.GetTextureStageState,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(stage),
		uintptr(typ),
		uintptr(unsafe.Pointer(&value)),
		0,
		0,
	)
	err = toErr(ret)
	return
}

// SetTextureStageState sets the state value for the currently assigned texture.
func (obj *Device) SetTextureStageState(
	stage uint32,
	typ TEXTURESTAGESTATETYPE,
	value uint32,
) Error {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.SetTextureStageState,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(stage),
		uintptr(typ),
		uintptr(value),
		0,
		0,
	)
	return toErr(ret)
}

// GetSamplerState return the sampler state value.
func (obj *Device) GetSamplerState(
	sampler uint32,
	typ SAMPLERSTATETYPE,
) (value uint32, err Error) {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.GetSamplerState,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(sampler),
		uintptr(typ),
		uintptr(unsafe.Pointer(&value)),
		0,
		0,
	)
	err = toErr(ret)
	return
}

// SetSamplerState sets the sampler state value.
func (obj *Device) SetSamplerState(
	sampler uint32,
	typ SAMPLERSTATETYPE,
	value uint32,
) Error {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.SetSamplerState,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(sampler),
		uintptr(typ),
		uintptr(value),
		0,
		0,
	)
	return toErr(ret)
}

// ValidateDevice reports the device's ability to render the current
// texture-blending operations and arguments in a single pass.
func (obj *Device) ValidateDevice() (numPasses uint32, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.ValidateDevice,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&numPasses)),
		0,
	)
	err = toErr(ret)
	return
}

// SetPaletteEntries sets palette entries.
func (obj *Device) SetPaletteEntries(
	paletteNumber uint,
	entries [256]PALETTEENTRY,
) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.SetPaletteEntries,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(paletteNumber),
		uintptr(unsafe.Pointer(&entries[0])),
	)
	return toErr(ret)
}

// GetPaletteEntries retrieves palette entries.
func (obj *Device) GetPaletteEntries(
	paletteNumber uint,
) (entries [256]PALETTEENTRY, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetPaletteEntries,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(paletteNumber),
		uintptr(unsafe.Pointer(&entries[0])),
	)
	err = toErr(ret)
	return
}

// SetCurrentTexturePalette sets the current texture palette.
func (obj *Device) SetCurrentTexturePalette(paletteNumber uint) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.SetCurrentTexturePalette,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(paletteNumber),
		0,
	)
	return toErr(ret)
}

// GetCurrentTexturePalette retrieves the current texture palette.
func (obj *Device) GetCurrentTexturePalette() (paletteNumber uint, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetCurrentTexturePalette,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&paletteNumber)),
		0,
	)
	err = toErr(ret)
	return
}

// SetScissorRect sets the scissor rectangle.
func (obj *Device) SetScissorRect(r RECT) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.SetScissorRect,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&r)),
		0,
	)
	return toErr(ret)
}

// GetScissorRect returns the scissor rectangle.
func (obj *Device) GetScissorRect() (r RECT, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetScissorRect,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&r)),
		0,
	)
	err = toErr(ret)
	return
}

// SetSoftwareVertexProcessing can be used to switch between software and
// hardware vertex processing.
func (obj *Device) SetSoftwareVertexProcessing(software bool) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.SetSoftwareVertexProcessing,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptrBool(software),
		0,
	)
	return toErr(ret)
}

// GetSoftwareVertexProcessing returns true if software vertex processing is
// set. Otherwise, it returns false.
func (obj *Device) GetSoftwareVertexProcessing() bool {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetSoftwareVertexProcessing,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return ret != 0
}

// SetNPatchMode enables or disables N-patches.
// segmentCount specifies the number of subdivision segments. If the number of
// segments is less than 1.0, N-patches are disabled. The default value is 0.0.
func (obj *Device) SetNPatchMode(segmentCount float32) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.SetNPatchMode,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(segmentCount),
		0,
	)
	return toErr(ret)
}

// GetNPatchMode returns the N-patch mode segments.
func (obj *Device) GetNPatchMode() float32 {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetNPatchMode,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return float32(ret)
}

// DrawPrimitive renders a sequence of nonindexed, geometric primitives of the
// specified type from the current set of data input streams.
func (obj *Device) DrawPrimitive(
	typ PRIMITIVETYPE,
	startVertex uint,
	primitiveCount uint,
) Error {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.DrawPrimitive,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(typ),
		uintptr(startVertex),
		uintptr(primitiveCount),
		0,
		0,
	)
	return toErr(ret)
}

// DrawIndexedPrimitive renders the specified geometric
// primitive into an array of vertices, based on indexing.
func (obj *Device) DrawIndexedPrimitive(
	typ PRIMITIVETYPE,
	baseVertexIndex int,
	minIndex uint,
	numVertices uint,
	startIndex uint,
	primitiveCount uint,
) Error {
	ret, _, _ := syscall.Syscall9(
		obj.vtbl.DrawIndexedPrimitive,
		7,
		uintptr(unsafe.Pointer(obj)),
		uintptr(typ),
		uintptr(baseVertexIndex),
		uintptr(minIndex),
		uintptr(numVertices),
		uintptr(startIndex),
		uintptr(primitiveCount),
		0,
		0,
	)
	return toErr(ret)
}

// DrawPrimitiveUP renders data specified by a user memory pointer as a sequence
// of geometric primitives of the specified type.
func (obj *Device) DrawPrimitiveUP(
	typ PRIMITIVETYPE,
	primitiveCount uint,
	vertexStreamZeroData uintptr,
	vertexStreamZeroStride uint,
) Error {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.DrawPrimitiveUP,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(typ),
		uintptr(primitiveCount),
		vertexStreamZeroData,
		uintptr(vertexStreamZeroStride),
		0,
	)
	return toErr(ret)
}

// DrawIndexedPrimitiveUP renders the specified geometric primitive with data
// specified by a user memory pointer.
func (obj *Device) DrawIndexedPrimitiveUP(
	typ PRIMITIVETYPE,
	minVertexIndex uint,
	numVertices uint,
	primitiveCount uint,
	indexData uintptr,
	indexDataFormat FORMAT,
	vertexStreamZeroData uintptr,
	vertexStreamZeroStride uint,
) Error {
	ret, _, _ := syscall.Syscall9(
		obj.vtbl.DrawIndexedPrimitiveUP,
		9,
		uintptr(unsafe.Pointer(obj)),
		uintptr(typ),
		uintptr(minVertexIndex),
		uintptr(numVertices),
		uintptr(primitiveCount),
		indexData,
		uintptr(indexDataFormat),
		vertexStreamZeroData,
		uintptr(vertexStreamZeroStride),
	)
	return toErr(ret)
}

// DrawIndexedPrimitiveUPuint32 renders the specified geometric primitive with
// data specified by a user memory pointer.
func (obj *Device) DrawIndexedPrimitiveUPuint32(
	primitiveType PRIMITIVETYPE,
	minVertexIndex uint,
	numVertices uint,
	primitiveCount uint,
	indexData []uint32,
	vertexStreamZeroData uintptr,
	vertexStreamZeroStride uint,
) Error {
	return obj.DrawIndexedPrimitiveUP(
		primitiveType,
		minVertexIndex,
		numVertices,
		primitiveCount,
		uintptr(unsafe.Pointer(&indexData[0])),
		FMT_INDEX32,
		vertexStreamZeroData,
		vertexStreamZeroStride,
	)
}

// DrawIndexedPrimitiveUPuint16 renders the specified geometric primitive with
// data specified by a user memory pointer.
func (obj *Device) DrawIndexedPrimitiveUPuint16(
	primitiveType PRIMITIVETYPE,
	minVertexIndex uint,
	numVertices uint,
	primitiveCount uint,
	indexData []uint16,
	vertexStreamZeroData uintptr,
	vertexStreamZeroStride uint,
) (
	err Error,
) {
	return obj.DrawIndexedPrimitiveUP(
		primitiveType,
		minVertexIndex,
		numVertices,
		primitiveCount,
		uintptr(unsafe.Pointer(&indexData[0])),
		FMT_INDEX16,
		vertexStreamZeroData,
		vertexStreamZeroStride,
	)
}

// ProcessVertices applies the vertex processing defined by the vertex shader to
// the set of input data streams, generating a single stream of interleaved
// vertex data to the destination vertex buffer.
func (obj *Device) ProcessVertices(
	srcStartIndex uint,
	destIndex uint,
	vertexCount uint,
	destBuffer *VertexBuffer,
	vertexDecl *VertexDeclaration,
	flags uint32,
) Error {
	ret, _, _ := syscall.Syscall9(
		obj.vtbl.ProcessVertices,
		7,
		uintptr(unsafe.Pointer(obj)),
		uintptr(srcStartIndex),
		uintptr(destIndex),
		uintptr(vertexCount),
		uintptr(unsafe.Pointer(destBuffer)),
		uintptr(unsafe.Pointer(vertexDecl)),
		uintptr(flags),
		0,
		0,
	)
	return toErr(ret)
}

// CreateVertexDeclaration creates a vertex shader declaration from the device
// and the vertex elements.
// The last element should be DeclEnd().
func (obj *Device) CreateVertexDeclaration(
	vertexElements []VERTEXELEMENT,
) (*VertexDeclaration, Error) {
	var elems uintptr
	if len(vertexElements) > 0 {
		elems = uintptr(unsafe.Pointer(&vertexElements[0]))
	}
	var decl *VertexDeclaration
	ret, _, _ := syscall.Syscall(
		obj.vtbl.CreateVertexDeclaration,
		3,
		uintptr(unsafe.Pointer(obj)),
		elems,
		uintptr(unsafe.Pointer(&decl)),
	)
	return decl, toErr(ret)
}

// DeclEnd marks the end of a slice of VERTEXELEMENTs.
func DeclEnd() VERTEXELEMENT {
	return VERTEXELEMENT{0xFF, 0, DECLTYPE_UNUSED, 0, 0, 0}
}

// SetVertexDeclaration sets a vertex declaration.
func (obj *Device) SetVertexDeclaration(decl *VertexDeclaration) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.SetVertexDeclaration,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(decl)),
		0,
	)
	return toErr(ret)
}

// GetVertexDeclaration returns a vertex shader declaration.
func (obj *Device) GetVertexDeclaration() (decl VertexDeclaration, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetVertexDeclaration,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&decl)),
		0,
	)
	err = toErr(ret)
	return
}

// SetFVF sets the current vertex stream declaration.
func (obj *Device) SetFVF(fvf uint32) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.SetFVF,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(fvf),
		0,
	)
	return toErr(ret)
}

// GetFVF gets the fixed vertex function declaration.
func (obj *Device) GetFVF() (fvf uint32, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetFVF,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&fvf)),
		0,
	)
	err = toErr(ret)
	return
}

// CreateVertexShader creates a vertex shader.
func (obj *Device) CreateVertexShader(function uintptr) (*VertexShader, Error) {
	var shader *VertexShader
	ret, _, _ := syscall.Syscall(
		obj.vtbl.CreateVertexShader,
		3,
		uintptr(unsafe.Pointer(obj)),
		function,
		uintptr(unsafe.Pointer(&shader)),
	)
	return shader, toErr(ret)
}

// CreateVertexShaderFromBytes creates a vertex shader from a byte slice.
func (obj *Device) CreateVertexShaderFromBytes(code []byte) (*VertexShader, Error) {
	return obj.CreateVertexShader(uintptr(unsafe.Pointer(&code[0])))
}

// SetVertexShader sets the vertex shader.
func (obj *Device) SetVertexShader(shader *VertexShader) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.SetVertexShader,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(shader)),
		0,
	)
	return toErr(ret)
}

// GetVertexShader retrieves the currently set vertex shader.
func (obj *Device) GetVertexShader() (*VertexShader, Error) {
	var shader *VertexShader
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetVertexShader,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&shader)),
		0,
	)
	return shader, toErr(ret)
}

// SetVertexShaderConstantF sets a floating-point vertex shader constant. The
// length of the slice must be a multiple of four.
func (obj *Device) SetVertexShaderConstantF(
	startRegister uint,
	constantData []float32,
) Error {
	if len(constantData) == 0 {
		return nil
	}
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.SetVertexShaderConstantF,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(startRegister),
		uintptr(unsafe.Pointer(&constantData[0])),
		uintptr((len(constantData)+3)/4),
		0,
		0,
	)
	return toErr(ret)
}

// GetVertexShaderConstantF returns a floating-point shader constant. The given
// float slice is filled with the constant data. The length of the slice must be
// a multiple of four.
func (obj *Device) GetVertexShaderConstantF(
	startRegister uint,
	constantData []float32,
) Error {
	if len(constantData) == 0 {
		return nil
	}
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.GetVertexShaderConstantF,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(startRegister),
		uintptr(unsafe.Pointer(&constantData[0])),
		uintptr(len(constantData)/4),
		0,
		0,
	)
	return toErr(ret)
}

// SetVertexShaderConstantI sets an integer vertex shader constant. The length
// of the slice must be a multiple of four.
func (obj *Device) SetVertexShaderConstantI(
	startRegister uint,
	constantData []int32,
) Error {
	if len(constantData) == 0 {
		return nil
	}
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.SetVertexShaderConstantI,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(startRegister),
		uintptr(unsafe.Pointer(&constantData[0])),
		uintptr((len(constantData)+3)/4),
		0,
		0,
	)
	return toErr(ret)
}

// GetVertexShaderConstantI returns an integer shader constant. The given int
// slice is filled with the constant data. The length of the slice must be a
// multiple of four.
func (obj *Device) GetVertexShaderConstantI(
	startRegister uint,
	constantData []int32,
) Error {
	if len(constantData) == 0 {
		return nil
	}
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.GetVertexShaderConstantI,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(startRegister),
		uintptr(unsafe.Pointer(&constantData[0])),
		uintptr(len(constantData)/4),
		0,
		0,
	)
	return toErr(ret)
}

// SetVertexShaderConstantB sets a Boolean vertex shader constant.
func (obj *Device) SetVertexShaderConstantB(
	startRegister uint,
	constantData []int32,
) Error {
	if len(constantData) == 0 {
		return nil
	}
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.SetVertexShaderConstantB,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(startRegister),
		uintptr(unsafe.Pointer(&constantData[0])),
		uintptr(len(constantData)),
		0,
		0,
	)
	return toErr(ret)
}

// GetVertexShaderConstantB returns a Boolean shader constant. The given bool
// slice is filled with the constant data.
func (obj *Device) GetVertexShaderConstantB(
	startRegister uint,
	constantData []int32,
) Error {
	if len(constantData) == 0 {
		return nil
	}
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.GetVertexShaderConstantB,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(startRegister),
		uintptr(unsafe.Pointer(&constantData[0])),
		uintptr(len(constantData)),
		0,
		0,
	)
	return toErr(ret)
}

// SetStreamSource binds a vertex buffer to a device data stream.
func (obj *Device) SetStreamSource(
	streamNumber uint,
	streamData *VertexBuffer,
	offsetInBytes uint,
	stride uint,
) Error {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.SetStreamSource,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(streamNumber),
		uintptr(unsafe.Pointer(streamData)),
		uintptr(offsetInBytes),
		uintptr(stride),
		0,
	)
	return toErr(ret)
}

// GetStreamSource retrieves a vertex buffer bound to the specified data stream.
func (obj *Device) GetStreamSource(
	streamNumber uint,
) (
	streamData *VertexBuffer,
	offsetInBytes uint,
	stride uint,
	err Error,
) {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.GetStreamSource,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(streamNumber),
		uintptr(unsafe.Pointer(&streamData)),
		uintptr(unsafe.Pointer(&offsetInBytes)),
		uintptr(unsafe.Pointer(&stride)),
		0,
	)
	err = toErr(ret)
	return
}

// SetStreamSourceFreq sets the stream source frequency divider value. This may
// be used to draw several instances of geometry.
func (obj *Device) SetStreamSourceFreq(
	streamNumber uint,
	frequencyParameter uint,
) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.SetStreamSourceFreq,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(streamNumber),
		uintptr(frequencyParameter),
	)
	return toErr(ret)
}

// GetStreamSourceFreq returns the stream source frequency divider value.
func (obj *Device) GetStreamSourceFreq(
	streamNumber uint,
) (divider uint32, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetStreamSourceFreq,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(streamNumber),
		uintptr(unsafe.Pointer(&divider)),
	)
	err = toErr(ret)
	return
}

// SetIndices sets index data.
func (obj *Device) SetIndices(indexData *IndexBuffer) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.SetIndices,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(indexData)),
		0,
	)
	return toErr(ret)
}

// GetIndices retrieves index data.
func (obj *Device) GetIndices() (indexData *IndexBuffer, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetIndices,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&indexData)),
		0,
	)
	err = toErr(ret)
	return
}

// CreatePixelShader creates a pixel shader.
func (obj *Device) CreatePixelShader(
	function uintptr,
) (shader *PixelShader, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.CreatePixelShader,
		3,
		uintptr(unsafe.Pointer(obj)),
		function,
		uintptr(unsafe.Pointer(&shader)),
	)
	err = toErr(ret)
	return
}

// CreatePixelShaderFromBytes creates a pixel shader from a byte slice.
func (obj *Device) CreatePixelShaderFromBytes(code []byte) (*PixelShader, Error) {
	return obj.CreatePixelShader(uintptr(unsafe.Pointer(&code[0])))
}

// SetPixelShader sets the current pixel shader to a previously created pixel
// shader.
func (obj *Device) SetPixelShader(shader *PixelShader) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.SetPixelShader,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(shader)),
		0,
	)
	return toErr(ret)
}

// GetPixelShader retrieves the currently set pixel shader.
func (obj *Device) GetPixelShader() (shader *PixelShader, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetPixelShader,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&shader)),
		0,
	)
	err = toErr(ret)
	return
}

// SetPixelShaderConstantF sets a floating-point shader constant. The length of
// the slice must be a multiple of four.
func (obj *Device) SetPixelShaderConstantF(
	startRegister uint,
	constantData []float32,
) Error {
	if len(constantData) == 0 {
		return nil
	}
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.SetPixelShaderConstantF,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(startRegister),
		uintptr(unsafe.Pointer(&constantData[0])),
		uintptr((len(constantData)+3)/4),
		0,
		0,
	)
	return toErr(ret)
}

// GetPixelShaderConstantF returns a floating-point shader constant. The given
// float slice is filled with the constant data. The length of the slice must be
// a multiple of four.
func (obj *Device) GetPixelShaderConstantF(
	startRegister uint,
	constantData []float32,
) Error {
	if len(constantData) == 0 {
		return nil
	}
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.GetPixelShaderConstantF,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(startRegister),
		uintptr(unsafe.Pointer(&constantData[0])),
		uintptr(len(constantData)/4),
		0,
		0,
	)
	return toErr(ret)
}

// SetPixelShaderConstantI sets an integer shader constant. The length of the
// slice must be a multiple of four.
func (obj *Device) SetPixelShaderConstantI(
	startRegister uint,
	constantData []int32,
) Error {
	if len(constantData) == 0 {
		return nil
	}
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.SetPixelShaderConstantI,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(startRegister),
		uintptr(unsafe.Pointer(&constantData[0])),
		uintptr((len(constantData)+3)/4),
		0,
		0,
	)
	return toErr(ret)
}

// GetPixelShaderConstantI returns an integer shader constant. The given int
// slice is filled with the constant data. The length of the slice must be a
// multiple of four.
func (obj *Device) GetPixelShaderConstantI(
	startRegister uint,
	constantData []int32,
) Error {
	if len(constantData) == 0 {
		return nil
	}
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.GetPixelShaderConstantI,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(startRegister),
		uintptr(unsafe.Pointer(&constantData[0])),
		uintptr(len(constantData)/4),
		0,
		0,
	)
	return toErr(ret)
}

// SetPixelShaderConstantB sets a Boolean shader constant.
func (obj *Device) SetPixelShaderConstantB(
	startRegister uint,
	constantData []int32,
) Error {
	if len(constantData) == 0 {
		return nil
	}
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.SetPixelShaderConstantB,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(startRegister),
		uintptr(unsafe.Pointer(&constantData[0])),
		uintptr(len(constantData)),
		0,
		0,
	)
	return toErr(ret)
}

// GetPixelShaderConstantB returns a Boolean shader constant. The given bool
// slice is filled with the constant data.
func (obj *Device) GetPixelShaderConstantB(
	startRegister uint,
	constantData []int32,
) Error {
	if len(constantData) == 0 {
		return nil
	}
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.GetPixelShaderConstantB,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(startRegister),
		uintptr(unsafe.Pointer(&constantData[0])),
		uintptr(len(constantData)),
		0,
		0,
	)
	return toErr(ret)
}

// DrawRectPatch draws a rectangular patch using the currently set streams.
func (obj *Device) DrawRectPatch(
	handle uint,
	numSegs [4]float32,
	rectPatchInfo *RECTPATCH_INFO,
) Error {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.DrawRectPatch,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(handle),
		uintptr(unsafe.Pointer(&numSegs[0])),
		uintptr(unsafe.Pointer(rectPatchInfo)),
		0,
		0,
	)
	return toErr(ret)
}

// DrawTriPatch draws a triangular patch using the currently set streams.
func (obj *Device) DrawTriPatch(
	handle uint,
	numSegs [3]float32,
	triPatchInfo *TRIPATCH_INFO,
) Error {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.DrawTriPatch,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(handle),
		uintptr(unsafe.Pointer(&numSegs[0])),
		uintptr(unsafe.Pointer(triPatchInfo)),
		0,
		0,
	)
	return toErr(ret)
}

// DeletePatch frees a cached high-order patch.
func (obj *Device) DeletePatch(handle uint) Error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.DeletePatch,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(handle),
		0,
	)
	return toErr(ret)
}

// CreateQuery creates a status query.
func (obj *Device) CreateQuery(typ QUERYTYPE) (query *Query, err Error) {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.CreateQuery,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(typ),
		uintptr(unsafe.Pointer(&query)),
	)
	err = toErr(ret)
	return
}
