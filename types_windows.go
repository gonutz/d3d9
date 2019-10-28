package d3d9

import "unsafe"

const MAX_DEVICE_IDENTIFIER_STRING = 512

// ADAPTER_IDENTIFIER contains information identifying the adapter.
type ADAPTER_IDENTIFIER struct {
	Driver           [MAX_DEVICE_IDENTIFIER_STRING]byte
	Description      [MAX_DEVICE_IDENTIFIER_STRING]byte
	DeviceName       [32]byte
	DriverVersion    int64
	VendorId         uint32
	DeviceId         uint32
	SubSysId         uint32
	Revision         uint32
	DeviceIdentifier GUID
	WHQLLevel        uint32
}

func zeroTerminatedString(b []byte) string {
	for i := range b {
		if b[i] == 0 {
			return string(b[:i])
		}
	}
	return string(b)
}

// GetDriver returns the Driver member as a Go string with trailing 0s removed.
func (id *ADAPTER_IDENTIFIER) GetDriver() string {
	if id == nil {
		return ""
	}
	return zeroTerminatedString(id.Driver[:])
}

// GetDescription returns the Description member as a Go string with trailing 0s
// removed.
func (id *ADAPTER_IDENTIFIER) GetDescription() string {
	if id == nil {
		return ""
	}
	return zeroTerminatedString(id.Description[:])
}

// GetDeviceName returns the DeviceName member as a Go string with trailing 0s
// removed.
func (id *ADAPTER_IDENTIFIER) GetDeviceName() string {
	if id == nil {
		return ""
	}
	return zeroTerminatedString(id.DeviceName[:])
}

// GetVersion splits the DriverVersion into its four semantic parts: product,
// version, sub-version and build.
func (id *ADAPTER_IDENTIFIER) GetVersion() (product, ver, subVer, build int) {
	if id != nil {
		v := uint64(id.DriverVersion)
		product = int((v & 0xFFFF000000000000) >> 48)
		ver = int((v & 0x0000FFFF00000000) >> 32)
		subVer = int((v & 0x00000000FFFF0000) >> 16)
		build = int(v & 0x000000000000FFFF)
	}
	return
}

// GUID is a globally unique identifier.
type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]uint8
}

type DISPLAYMODE struct {
	Width       uint32
	Height      uint32
	RefreshRate uint32
	Format      FORMAT
}

// CAPS represents the capabilities of the hardware exposed through the Direct3D
// object.
type CAPS struct {
	DeviceType                        DEVTYPE
	AdapterOrdinal                    uint32
	Caps                              uint32
	Caps2                             uint32
	Caps3                             uint32
	PresentationIntervals             uint32
	CursorCaps                        uint32
	DevCaps                           uint32
	PrimitiveMiscCaps                 uint32
	RasterCaps                        uint32
	ZCmpCaps                          uint32
	SrcBlendCaps                      uint32
	DestBlendCaps                     uint32
	AlphaCmpCaps                      uint32
	ShadeCaps                         uint32
	TextureCaps                       uint32
	TextureFilterCaps                 uint32
	CubeTextureFilterCaps             uint32
	VolumeTextureFilterCaps           uint32
	TextureAddressCaps                uint32
	VolumeTextureAddressCaps          uint32
	LineCaps                          uint32
	MaxTextureWidth                   uint32
	MaxTextureHeight                  uint32
	MaxVolumeExtent                   uint32
	MaxTextureRepeat                  uint32
	MaxTextureAspectRatio             uint32
	MaxAnisotropy                     uint32
	MaxVertexW                        float32
	GuardBandLeft                     float32
	GuardBandTop                      float32
	GuardBandRight                    float32
	GuardBandBottom                   float32
	ExtentsAdjust                     float32
	StencilCaps                       uint32
	FVFCaps                           uint32
	TextureOpCaps                     uint32
	MaxTextureBlendStages             uint32
	MaxSimultaneousTextures           uint32
	VertexProcessingCaps              uint32
	MaxActiveLights                   uint32
	MaxUserClipPlanes                 uint32
	MaxVertexBlendMatrices            uint32
	MaxVertexBlendMatrixIndex         uint32
	MaxPointSize                      float32
	MaxPrimitiveCount                 uint32
	MaxVertexIndex                    uint32
	MaxStreams                        uint32
	MaxStreamStride                   uint32
	VertexShaderVersion               uint32
	MaxVertexShaderConst              uint32
	PixelShaderVersion                uint32
	PixelShader1xMaxValue             float32
	DevCaps2                          uint32
	MasterAdapterOrdinal              uint32
	AdapterOrdinalInGroup             uint32
	NumberOfAdaptersInGroup           uint32
	DeclTypes                         uint32
	NumSimultaneousRTs                uint32
	StretchRectFilterCaps             uint32
	VS20Caps                          VSHADERCAPS2_0
	PS20Caps                          PSHADERCAPS2_0
	VertexTextureFilterCaps           uint32
	MaxVShaderInstructionsExecuted    uint32
	MaxPShaderInstructionsExecuted    uint32
	MaxVertexShader30InstructionSlots uint32
	MaxPixelShader30InstructionSlots  uint32
}

// VSHADERCAPS2_0 contains vertex shader capabilities.
type VSHADERCAPS2_0 struct {
	Caps                    uint32
	DynamicFlowControlDepth int32
	NumTemps                int32
	StaticFlowControlDepth  int32
}

// PSHADERCAPS2_0 describes pixel shader driver caps.
type PSHADERCAPS2_0 struct {
	Caps                    uint32
	DynamicFlowControlDepth int32
	NumTemps                int32
	StaticFlowControlDepth  int32
	NumInstructionSlots     int32
}

// PRESENT_PARAMETERS describes the presentation parameters.
type PRESENT_PARAMETERS struct {
	BackBufferWidth            uint32
	BackBufferHeight           uint32
	BackBufferFormat           FORMAT
	BackBufferCount            uint32
	MultiSampleType            MULTISAMPLE_TYPE
	MultiSampleQuality         uint32
	SwapEffect                 SWAPEFFECT
	HDeviceWindow              HWND
	Windowed                   int32
	EnableAutoDepthStencil     int32
	AutoDepthStencilFormat     FORMAT
	Flags                      uint32
	FullScreen_RefreshRateInHz uint32
	PresentationInterval       uint32
}

type (
	HANDLE   uintptr
	HWND     HANDLE
	HMONITOR HANDLE
	HDC      HANDLE
)

// RECT describes a rectangle.
type RECT struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

// RGNDATA contains region data.
type RGNDATA struct {
	Rdh    RGNDATAHEADER
	Buffer [1]byte
}

// RGNDATAHEADER describes region data.
type RGNDATAHEADER struct {
	DwSize   uint32
	IType    uint32
	NCount   uint32
	NRgnSize uint32
	RcBound  RECT
}

// DEVICE_CREATION_PARAMETERS describes the creation parameters for a device.
type DEVICE_CREATION_PARAMETERS struct {
	AdapterOrdinal uint32
	DeviceType     DEVTYPE
	HFocusWindow   HWND
	BehaviorFlags  uint32
}

// RASTER_STATUS describes the raster status.
type RASTER_STATUS struct {
	InVBlank int32
	ScanLine uint32
}

// GAMMARAMP contains red, green, and blue ramp data.
type GAMMARAMP struct {
	Red   [256]uint16
	Green [256]uint16
	Blue  [256]uint16
}

// POINT describes a 2D point.
type POINT struct {
	X int32
	Y int32
}

// MATRIX describes a matrix.
type MATRIX [16]float32

// VIEWPORT defines the window dimensions of a render-target surface onto which
// a 3D volume projects.
type VIEWPORT struct {
	X      uint32
	Y      uint32
	Width  uint32
	Height uint32
	MinZ   float32
	MaxZ   float32
}

// MATERIAL specifies material properties.
type MATERIAL struct {
	Diffuse  COLORVALUE
	Ambient  COLORVALUE
	Specular COLORVALUE
	Emissive COLORVALUE
	Power    float32
}

// COLORVALUE describes color values.
type COLORVALUE struct {
	R float32
	G float32
	B float32
	A float32
}

// LIGHT defines a set of lighting properties.
type LIGHT struct {
	Type         LIGHTTYPE
	Diffuse      COLORVALUE
	Specular     COLORVALUE
	Ambient      COLORVALUE
	Position     VECTOR
	Direction    VECTOR
	Range        float32
	Falloff      float32
	Attenuation0 float32
	Attenuation1 float32
	Attenuation2 float32
	Theta        float32
	Phi          float32
}

// VECTOR defines a vector.
type VECTOR struct {
	X float32
	Y float32
	Z float32
}

// CLIPSTATUS describes the current clip status.
type CLIPSTATUS struct {
	ClipUnion        uint32
	ClipIntersection uint32
}

// PALETTEENTRY specifies the color and usage of an entry in a logical palette.
type PALETTEENTRY struct {
	PeRed   byte
	PeGreen byte
	PeBlue  byte
	PeFlags byte
}

// VERTEXELEMENT defines the vertex data layout. Each vertex can contain one or
// more data types, and each data type is described by a vertex element.
type VERTEXELEMENT struct {
	Stream     uint16
	Offset     uint16
	Type       DECLTYPE
	Method     DECLMETHOD
	Usage      DECLUSAGE
	UsageIndex byte
}

// RECTPATCH_INFO describes a rectangular high-order patch.
type RECTPATCH_INFO struct {
	StartVertexOffsetWidth  uint32
	StartVertexOffsetHeight uint32
	Width                   uint32
	Height                  uint32
	Stride                  uint32
	Basis                   BASISTYPE
	Degree                  DEGREETYPE
}

// TRIPATCH_INFO describes a triangular high-order patch.
type TRIPATCH_INFO struct {
	StartVertexOffset uint32
	NumVertices       uint32
	Basis             BASISTYPE
	Degree            DEGREETYPE
}

// BOX defines a volume.
type BOX struct {
	Left   uint32
	Top    uint32
	Right  uint32
	Bottom uint32
	Front  uint32
	Back   uint32
}

// VOLUME_DESC describes a volume.
type VOLUME_DESC struct {
	Format FORMAT
	Type   RESOURCETYPE
	Usage  uint32
	Pool   POOL
	Width  uint32
	Height uint32
	Depth  uint32
}

// LOCKED_BOX describes a locked box (volume).
type LOCKED_BOX struct {
	RowPitch   int32
	SlicePitch int32
	PBits      uintptr
}

// VERTEXBUFFER_DESC describes a vertex buffer.
type VERTEXBUFFER_DESC struct {
	Format FORMAT
	Type   RESOURCETYPE
	Usage  uint32
	Pool   POOL
	Size   uint32
	FVF    uint32
}

// INDEXBUFFER_DESC describes an index buffer.
type INDEXBUFFER_DESC struct {
	Format FORMAT
	Type   RESOURCETYPE
	Usage  uint32
	Pool   POOL
	Size   uint32
}

// SURFACE_DESC describes a surface.
type SURFACE_DESC struct {
	Format             FORMAT
	Type               RESOURCETYPE
	Usage              uint32
	Pool               POOL
	MultiSampleType    MULTISAMPLE_TYPE
	MultiSampleQuality uint32
	Width              uint32
	Height             uint32
}

// LOCKED_RECT describes a locked rectangular region.
type LOCKED_RECT struct {
	Pitch int32
	PBits uintptr
}

// SetAllBytes will fill the whole rect with the given data, taking into account
// the rect's pitch. The given byte slice is expected to have the given stride
// in bytes, i.e. one line in the given data is <srcStride> bytes in length.
func (r LOCKED_RECT) SetAllBytes(data []byte, srcStride int) {
	if len(data) == 0 {
		return
	}

	dest := r.PBits
	destStride := int(r.Pitch)
	src := uintptr(unsafe.Pointer(&data[0]))
	height := (len(data) + srcStride - 1) / srcStride

	stride := srcStride
	if destStride < srcStride {
		stride = destStride
	}
	destSkip := uintptr(destStride - stride)
	srcSkip := uintptr(srcStride - stride)
	d := dest
	s := src
	if srcStride%8 == 0 && destStride%8 == 0 {
		// in this case we can speed up copying by using 8 byte wide uint64s
		// instead of copying byte for byte
		for y := 0; y < height; y++ {
			for x := 0; x < stride; x += 8 {
				*((*uint64)(unsafe.Pointer(d))) = *((*uint64)(unsafe.Pointer(s)))
				d += 8
				s += 8
			}
			d += destSkip
			s += srcSkip
		}
	} else if srcStride%4 == 0 && destStride%4 == 0 {
		// in this case we can speed up copying by using 4 byte wide uint32s
		// instead of copying byte for byte
		for y := 0; y < height; y++ {
			for x := 0; x < stride; x += 4 {
				*((*uint32)(unsafe.Pointer(d))) = *((*uint32)(unsafe.Pointer(s)))
				d += 4
				s += 4
			}
			d += destSkip
			s += srcSkip
		}
	} else {
		// in the unlikely case that stride is neither a multiple of 8 nor 4
		// bytes, just copy byte for byte
		for y := 0; y < height; y++ {
			for x := 0; x < stride; x++ {
				*((*byte)(unsafe.Pointer(d))) = *((*byte)(unsafe.Pointer(s)))
				d++
				s++
			}
			d += destSkip
			s += srcSkip
		}
	}
}

// DEVINFO_D3D9BANDWIDTHTIMINGS contains throughput metrics for help in
// understanding the performance of an application.
type DEVINFO_D3D9BANDWIDTHTIMINGS struct {
	MaxBandwidthUtilized                float32
	FrontEndUploadMemoryUtilizedPercent float32
	VertexRateUtilizedPercent           float32
	TriangleSetupRateUtilizedPercent    float32
	FillRateUtilizedPercent             float32
}

// DEVINFO_D3D9CACHEUTILIZATION measures the cache hit rate performance for
// textures and indexed vertices.
type DEVINFO_D3D9CACHEUTILIZATION struct {
	TextureCacheHitRate             float32
	PostTransformVertexCacheHitRate float32
}

// DEVINFO_D3D9INTERFACETIMINGS contains the percent of time processing data in
// the driver. These statistics may help identify cases when the driver is
// waiting for other resources.
type DEVINFO_D3D9INTERFACETIMINGS struct {
	WaitingForGPUToUseApplicationResourceTimePercent float32
	WaitingForGPUToAcceptMoreCommandsTimePercent     float32
	WaitingForGPUToStayWithinLatencyTimePercent      float32
	WaitingForGPUExclusiveResourceTimePercent        float32
	WaitingForGPUOtherTimePercent                    float32
}

// DEVINFO_D3D9PIPELINETIMINGS contains the percent of time processing data in
// the pipeline.
type DEVINFO_D3D9PIPELINETIMINGS struct {
	VertexProcessingTimePercent   float32
	PixelProcessingTimePercent    float32
	OtherGPUProcessingTimePercent float32
	GPUIdleTimePercent            float32
}

// DEVINFO_D3D9STAGETIMINGS contains the percent of time processing shader data.
type DEVINFO_D3D9STAGETIMINGS struct {
	MemoryProcessingPercent      float32
	ComputationProcessingPercent float32
}

// DEVINFO_D3DVERTEXSTATS reports the number of triangles that have been
// processed and clipped by the runtime's software vertex processing.
type DEVINFO_D3DVERTEXSTATS struct {
	NumRenderedTriangles      uint32
	NumExtraClippingTriangles uint32
}

// DEVINFO_VCACHE contains vertex cache optimization hints.
type DEVINFO_VCACHE struct {
	Pattern     uint32
	OptMethod   uint32
	CacheSize   uint32
	MagicNumber uint32
}

// RESOURCESTATS contains resource statistics gathered by the
// DEVINFO_ResourceManager when using the asynchronous query mechanism.
type RESOURCESTATS struct {
	BThrashing            uint32
	ApproxBytesDownloaded uint32
	NumEvicts             uint32
	NumVidCreates         uint32
	LastPri               uint32
	NumUsed               uint32
	NumUsedInVidMem       uint32
	WorkingSet            uint32
	WorkingSetBytes       uint32
	TotalManaged          uint32
	TotalBytes            uint32
}

// LUID is a 64-bit value guaranteed to be unique only on the system on which it
// was generated. The uniqueness of a locally unique identifier (LUID) is
// guaranteed only until the system is restarted.
type LUID struct {
	LowPart  uint32
	HighPart int32
}

// RANGE defines a range.
type RANGE struct {
	Offset uint32
	Size   uint32
}
