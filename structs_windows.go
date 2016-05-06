package d3d9

/*
#include "d3d9wrapper.h"

void setLargeInteger(void* i, long long to) {
  *(LARGE_INTEGER*)i = (LARGE_INTEGER)to;
}

long long getLargeInteger(void* i) {
  return *((long long*)i);
}

void toCMatrix(D3DMATRIX* m, float* values) {
  m->_11 = *values;
  m->_12 = *(values+1);
  m->_13 = *(values+2);
  m->_14 = *(values+3);
  m->_21 = *(values+4);
  m->_22 = *(values+5);
  m->_23 = *(values+6);
  m->_24 = *(values+7);
  m->_31 = *(values+8);
  m->_32 = *(values+9);
  m->_33 = *(values+10);
  m->_34 = *(values+11);
  m->_41 = *(values+12);
  m->_42 = *(values+13);
  m->_43 = *(values+14);
  m->_44 = *(values+15);
}

void toGoMatrix(D3DMATRIX* m, float* values) {
  *values      = m->_11;
  *(values+1)  = m->_12;
  *(values+2)  = m->_13;
  *(values+3)  = m->_14;
  *(values+4)  = m->_21;
  *(values+5)  = m->_22;
  *(values+6)  = m->_23;
  *(values+7)  = m->_24;
  *(values+8)  = m->_31;
  *(values+9)  = m->_32;
  *(values+10) = m->_33;
  *(values+11) = m->_34;
  *(values+12) = m->_41;
  *(values+13) = m->_42;
  *(values+14) = m->_43;
  *(values+15) = m->_44;
}

void setWholeLockedRect(
		void* dest,
		int destStride,
		void* source,
		int sourceStride,
		int height) {
	int stride = sourceStride;
	if (destStride < sourceStride)
		stride = destStride;
	int destSkip = destStride - stride;
	int sourceSkip = sourceStride - stride;
	int x, y;
	char* d = (char*) dest;
	char* s = (char*) source;
	for (y = 0; y < height; ++y) {
		for (x = 0; x < stride; ++x)
			*d++ = *s++;
		d += destSkip;
		s += sourceSkip;
	}
}
*/
import "C"
import "unsafe"

func toBOOL(b bool) C.BOOL {
	if b {
		return 1
	}
	return 0
}

// ADAPTER_IDENTIFIER contains information identifying the adapter.
type ADAPTER_IDENTIFIER struct {
	Driver           [C.MAX_DEVICE_IDENTIFIER_STRING]byte
	Description      [C.MAX_DEVICE_IDENTIFIER_STRING]byte
	DeviceName       [32]byte
	DriverVersion    int64
	VendorId         uint32
	DeviceId         uint32
	SubSysId         uint32
	Revision         uint32
	DeviceIdentifier GUID
	WHQLLevel        uint32
}

func (s *ADAPTER_IDENTIFIER) toC() C.D3DADAPTER_IDENTIFIER9 {
	var c C.D3DADAPTER_IDENTIFIER9
	for i := range s.Driver {
		c.Driver[i] = (C.char)(s.Driver[i])
	}
	for i := range s.Description {
		c.Description[i] = (C.char)(s.Description[i])
	}
	for i := range s.DeviceName {
		c.DeviceName[i] = (C.char)(s.DeviceName[i])
	}
	C.setLargeInteger((unsafe.Pointer)(&c.DriverVersion), (C.longlong)(s.DriverVersion))
	c.VendorId = (C.DWORD)(s.VendorId)
	c.DeviceId = (C.DWORD)(s.DeviceId)
	c.SubSysId = (C.DWORD)(s.SubSysId)
	c.Revision = (C.DWORD)(s.Revision)
	c.DeviceIdentifier = s.DeviceIdentifier.toC()
	c.WHQLLevel = (C.DWORD)(s.WHQLLevel)
	return c
}

func (s *ADAPTER_IDENTIFIER) fromC(c *C.D3DADAPTER_IDENTIFIER9) {
	for i := range s.Driver {
		s.Driver[i] = (byte)(c.Driver[i])
	}
	for i := range s.Description {
		s.Description[i] = (byte)(c.Description[i])
	}
	for i := range s.DeviceName {
		s.DeviceName[i] = (byte)(c.DeviceName[i])
	}
	s.DriverVersion = (int64)(C.getLargeInteger((unsafe.Pointer)(&c.DriverVersion)))
	s.VendorId = (uint32)(c.VendorId)
	s.DeviceId = (uint32)(c.DeviceId)
	s.SubSysId = (uint32)(c.SubSysId)
	s.Revision = (uint32)(c.Revision)
	s.DeviceIdentifier.fromC(&c.DeviceIdentifier)
	s.WHQLLevel = (uint32)(c.WHQLLevel)
}

// BOX defines a volume.
type BOX struct {
	Left   uint
	Top    uint
	Right  uint
	Bottom uint
	Front  uint
	Back   uint
}

func (s *BOX) toC() C.D3DBOX {
	var c C.D3DBOX
	c.Left = (C.UINT)(s.Left)
	c.Top = (C.UINT)(s.Top)
	c.Right = (C.UINT)(s.Right)
	c.Bottom = (C.UINT)(s.Bottom)
	c.Front = (C.UINT)(s.Front)
	c.Back = (C.UINT)(s.Back)
	return c
}

func (s *BOX) fromC(c *C.D3DBOX) {
	s.Left = (uint)(c.Left)
	s.Top = (uint)(c.Top)
	s.Right = (uint)(c.Right)
	s.Bottom = (uint)(c.Bottom)
	s.Front = (uint)(c.Front)
	s.Back = (uint)(c.Back)
}

// CAPS represents the capabilities of the hardware exposed through the Direct3D
// object.
type CAPS struct {
	DeviceType                        DEVTYPE
	AdapterOrdinal                    uint
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
	MasterAdapterOrdinal              uint
	AdapterOrdinalInGroup             uint
	NumberOfAdaptersInGroup           uint
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

func (s *CAPS) toC() C.D3DCAPS9 {
	var c C.D3DCAPS9
	c.DeviceType = (C.D3DDEVTYPE)(s.DeviceType)
	c.AdapterOrdinal = (C.UINT)(s.AdapterOrdinal)
	c.Caps = (C.DWORD)(s.Caps)
	c.Caps2 = (C.DWORD)(s.Caps2)
	c.Caps3 = (C.DWORD)(s.Caps3)
	c.PresentationIntervals = (C.DWORD)(s.PresentationIntervals)
	c.CursorCaps = (C.DWORD)(s.CursorCaps)
	c.DevCaps = (C.DWORD)(s.DevCaps)
	c.PrimitiveMiscCaps = (C.DWORD)(s.PrimitiveMiscCaps)
	c.RasterCaps = (C.DWORD)(s.RasterCaps)
	c.ZCmpCaps = (C.DWORD)(s.ZCmpCaps)
	c.SrcBlendCaps = (C.DWORD)(s.SrcBlendCaps)
	c.DestBlendCaps = (C.DWORD)(s.DestBlendCaps)
	c.AlphaCmpCaps = (C.DWORD)(s.AlphaCmpCaps)
	c.ShadeCaps = (C.DWORD)(s.ShadeCaps)
	c.TextureCaps = (C.DWORD)(s.TextureCaps)
	c.TextureFilterCaps = (C.DWORD)(s.TextureFilterCaps)
	c.CubeTextureFilterCaps = (C.DWORD)(s.CubeTextureFilterCaps)
	c.VolumeTextureFilterCaps = (C.DWORD)(s.VolumeTextureFilterCaps)
	c.TextureAddressCaps = (C.DWORD)(s.TextureAddressCaps)
	c.VolumeTextureAddressCaps = (C.DWORD)(s.VolumeTextureAddressCaps)
	c.LineCaps = (C.DWORD)(s.LineCaps)
	c.MaxTextureWidth = (C.DWORD)(s.MaxTextureWidth)
	c.MaxTextureHeight = (C.DWORD)(s.MaxTextureHeight)
	c.MaxVolumeExtent = (C.DWORD)(s.MaxVolumeExtent)
	c.MaxTextureRepeat = (C.DWORD)(s.MaxTextureRepeat)
	c.MaxTextureAspectRatio = (C.DWORD)(s.MaxTextureAspectRatio)
	c.MaxAnisotropy = (C.DWORD)(s.MaxAnisotropy)
	c.MaxVertexW = (C.float)(s.MaxVertexW)
	c.GuardBandLeft = (C.float)(s.GuardBandLeft)
	c.GuardBandTop = (C.float)(s.GuardBandTop)
	c.GuardBandRight = (C.float)(s.GuardBandRight)
	c.GuardBandBottom = (C.float)(s.GuardBandBottom)
	c.ExtentsAdjust = (C.float)(s.ExtentsAdjust)
	c.StencilCaps = (C.DWORD)(s.StencilCaps)
	c.FVFCaps = (C.DWORD)(s.FVFCaps)
	c.TextureOpCaps = (C.DWORD)(s.TextureOpCaps)
	c.MaxTextureBlendStages = (C.DWORD)(s.MaxTextureBlendStages)
	c.MaxSimultaneousTextures = (C.DWORD)(s.MaxSimultaneousTextures)
	c.VertexProcessingCaps = (C.DWORD)(s.VertexProcessingCaps)
	c.MaxActiveLights = (C.DWORD)(s.MaxActiveLights)
	c.MaxUserClipPlanes = (C.DWORD)(s.MaxUserClipPlanes)
	c.MaxVertexBlendMatrices = (C.DWORD)(s.MaxVertexBlendMatrices)
	c.MaxVertexBlendMatrixIndex = (C.DWORD)(s.MaxVertexBlendMatrixIndex)
	c.MaxPointSize = (C.float)(s.MaxPointSize)
	c.MaxPrimitiveCount = (C.DWORD)(s.MaxPrimitiveCount)
	c.MaxVertexIndex = (C.DWORD)(s.MaxVertexIndex)
	c.MaxStreams = (C.DWORD)(s.MaxStreams)
	c.MaxStreamStride = (C.DWORD)(s.MaxStreamStride)
	c.VertexShaderVersion = (C.DWORD)(s.VertexShaderVersion)
	c.MaxVertexShaderConst = (C.DWORD)(s.MaxVertexShaderConst)
	c.PixelShaderVersion = (C.DWORD)(s.PixelShaderVersion)
	c.PixelShader1xMaxValue = (C.float)(s.PixelShader1xMaxValue)
	c.DevCaps2 = (C.DWORD)(s.DevCaps2)
	c.MasterAdapterOrdinal = (C.UINT)(s.MasterAdapterOrdinal)
	c.AdapterOrdinalInGroup = (C.UINT)(s.AdapterOrdinalInGroup)
	c.NumberOfAdaptersInGroup = (C.UINT)(s.NumberOfAdaptersInGroup)
	c.DeclTypes = (C.DWORD)(s.DeclTypes)
	c.NumSimultaneousRTs = (C.DWORD)(s.NumSimultaneousRTs)
	c.StretchRectFilterCaps = (C.DWORD)(s.StretchRectFilterCaps)
	c.VS20Caps = s.VS20Caps.toC()
	c.PS20Caps = s.PS20Caps.toC()
	c.VertexTextureFilterCaps = (C.DWORD)(s.VertexTextureFilterCaps)
	c.MaxVShaderInstructionsExecuted = (C.DWORD)(s.MaxVShaderInstructionsExecuted)
	c.MaxPShaderInstructionsExecuted = (C.DWORD)(s.MaxPShaderInstructionsExecuted)
	c.MaxVertexShader30InstructionSlots = (C.DWORD)(s.MaxVertexShader30InstructionSlots)
	c.MaxPixelShader30InstructionSlots = (C.DWORD)(s.MaxPixelShader30InstructionSlots)
	return c
}

func (s *CAPS) fromC(c *C.D3DCAPS9) {
	s.DeviceType = (DEVTYPE)(c.DeviceType)
	s.AdapterOrdinal = (uint)(c.AdapterOrdinal)
	s.Caps = (uint32)(c.Caps)
	s.Caps2 = (uint32)(c.Caps2)
	s.Caps3 = (uint32)(c.Caps3)
	s.PresentationIntervals = (uint32)(c.PresentationIntervals)
	s.CursorCaps = (uint32)(c.CursorCaps)
	s.DevCaps = (uint32)(c.DevCaps)
	s.PrimitiveMiscCaps = (uint32)(c.PrimitiveMiscCaps)
	s.RasterCaps = (uint32)(c.RasterCaps)
	s.ZCmpCaps = (uint32)(c.ZCmpCaps)
	s.SrcBlendCaps = (uint32)(c.SrcBlendCaps)
	s.DestBlendCaps = (uint32)(c.DestBlendCaps)
	s.AlphaCmpCaps = (uint32)(c.AlphaCmpCaps)
	s.ShadeCaps = (uint32)(c.ShadeCaps)
	s.TextureCaps = (uint32)(c.TextureCaps)
	s.TextureFilterCaps = (uint32)(c.TextureFilterCaps)
	s.CubeTextureFilterCaps = (uint32)(c.CubeTextureFilterCaps)
	s.VolumeTextureFilterCaps = (uint32)(c.VolumeTextureFilterCaps)
	s.TextureAddressCaps = (uint32)(c.TextureAddressCaps)
	s.VolumeTextureAddressCaps = (uint32)(c.VolumeTextureAddressCaps)
	s.LineCaps = (uint32)(c.LineCaps)
	s.MaxTextureWidth = (uint32)(c.MaxTextureWidth)
	s.MaxTextureHeight = (uint32)(c.MaxTextureHeight)
	s.MaxVolumeExtent = (uint32)(c.MaxVolumeExtent)
	s.MaxTextureRepeat = (uint32)(c.MaxTextureRepeat)
	s.MaxTextureAspectRatio = (uint32)(c.MaxTextureAspectRatio)
	s.MaxAnisotropy = (uint32)(c.MaxAnisotropy)
	s.MaxVertexW = (float32)(c.MaxVertexW)
	s.GuardBandLeft = (float32)(c.GuardBandLeft)
	s.GuardBandTop = (float32)(c.GuardBandTop)
	s.GuardBandRight = (float32)(c.GuardBandRight)
	s.GuardBandBottom = (float32)(c.GuardBandBottom)
	s.ExtentsAdjust = (float32)(c.ExtentsAdjust)
	s.StencilCaps = (uint32)(c.StencilCaps)
	s.FVFCaps = (uint32)(c.FVFCaps)
	s.TextureOpCaps = (uint32)(c.TextureOpCaps)
	s.MaxTextureBlendStages = (uint32)(c.MaxTextureBlendStages)
	s.MaxSimultaneousTextures = (uint32)(c.MaxSimultaneousTextures)
	s.VertexProcessingCaps = (uint32)(c.VertexProcessingCaps)
	s.MaxActiveLights = (uint32)(c.MaxActiveLights)
	s.MaxUserClipPlanes = (uint32)(c.MaxUserClipPlanes)
	s.MaxVertexBlendMatrices = (uint32)(c.MaxVertexBlendMatrices)
	s.MaxVertexBlendMatrixIndex = (uint32)(c.MaxVertexBlendMatrixIndex)
	s.MaxPointSize = (float32)(c.MaxPointSize)
	s.MaxPrimitiveCount = (uint32)(c.MaxPrimitiveCount)
	s.MaxVertexIndex = (uint32)(c.MaxVertexIndex)
	s.MaxStreams = (uint32)(c.MaxStreams)
	s.MaxStreamStride = (uint32)(c.MaxStreamStride)
	s.VertexShaderVersion = (uint32)(c.VertexShaderVersion)
	s.MaxVertexShaderConst = (uint32)(c.MaxVertexShaderConst)
	s.PixelShaderVersion = (uint32)(c.PixelShaderVersion)
	s.PixelShader1xMaxValue = (float32)(c.PixelShader1xMaxValue)
	s.DevCaps2 = (uint32)(c.DevCaps2)
	s.MasterAdapterOrdinal = (uint)(c.MasterAdapterOrdinal)
	s.AdapterOrdinalInGroup = (uint)(c.AdapterOrdinalInGroup)
	s.NumberOfAdaptersInGroup = (uint)(c.NumberOfAdaptersInGroup)
	s.DeclTypes = (uint32)(c.DeclTypes)
	s.NumSimultaneousRTs = (uint32)(c.NumSimultaneousRTs)
	s.StretchRectFilterCaps = (uint32)(c.StretchRectFilterCaps)
	s.VS20Caps.fromC(&c.VS20Caps)
	s.PS20Caps.fromC(&c.PS20Caps)
	s.VertexTextureFilterCaps = (uint32)(c.VertexTextureFilterCaps)
	s.MaxVShaderInstructionsExecuted = (uint32)(c.MaxVShaderInstructionsExecuted)
	s.MaxPShaderInstructionsExecuted = (uint32)(c.MaxPShaderInstructionsExecuted)
	s.MaxVertexShader30InstructionSlots = (uint32)(c.MaxVertexShader30InstructionSlots)
	s.MaxPixelShader30InstructionSlots = (uint32)(c.MaxPixelShader30InstructionSlots)
}

// CLIPSTATUS describes the current clip status.
type CLIPSTATUS struct {
	ClipUnion        uint32
	ClipIntersection uint32
}

func (s *CLIPSTATUS) toC() C.D3DCLIPSTATUS9 {
	var c C.D3DCLIPSTATUS9
	c.ClipUnion = (C.DWORD)(s.ClipUnion)
	c.ClipIntersection = (C.DWORD)(s.ClipIntersection)
	return c
}

func (s *CLIPSTATUS) fromC(c *C.D3DCLIPSTATUS9) {
	s.ClipUnion = (uint32)(c.ClipUnion)
	s.ClipIntersection = (uint32)(c.ClipIntersection)
}

// COLORVALUE describes color values.
type COLORVALUE struct {
	R float32
	G float32
	B float32
	A float32
}

func (s *COLORVALUE) toC() C.D3DCOLORVALUE {
	var c C.D3DCOLORVALUE
	c.r = (C.float)(s.R)
	c.g = (C.float)(s.G)
	c.b = (C.float)(s.B)
	c.a = (C.float)(s.A)
	return c
}

func (s *COLORVALUE) fromC(c *C.D3DCOLORVALUE) {
	s.R = (float32)(c.r)
	s.G = (float32)(c.g)
	s.B = (float32)(c.b)
	s.A = (float32)(c.a)
}

// DEVICE_CREATION_PARAMETERS describes the creation parameters for a device.
type DEVICE_CREATION_PARAMETERS struct {
	AdapterOrdinal uint
	DeviceType     DEVTYPE
	HFocusWindow   unsafe.Pointer
	BehaviorFlags  uint32
}

func (s *DEVICE_CREATION_PARAMETERS) toC() C.D3DDEVICE_CREATION_PARAMETERS {
	var c C.D3DDEVICE_CREATION_PARAMETERS
	c.AdapterOrdinal = (C.UINT)(s.AdapterOrdinal)
	c.DeviceType = (C.D3DDEVTYPE)(s.DeviceType)
	c.hFocusWindow = (C.HWND)(s.HFocusWindow)
	c.BehaviorFlags = (C.DWORD)(s.BehaviorFlags)
	return c
}

func (s *DEVICE_CREATION_PARAMETERS) fromC(c *C.D3DDEVICE_CREATION_PARAMETERS) {
	s.AdapterOrdinal = (uint)(c.AdapterOrdinal)
	s.DeviceType = (DEVTYPE)(c.DeviceType)
	s.HFocusWindow = (unsafe.Pointer)(c.hFocusWindow)
	s.BehaviorFlags = (uint32)(c.BehaviorFlags)
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

func (s *DEVINFO_D3D9BANDWIDTHTIMINGS) toC() C.D3DDEVINFO_D3D9BANDWIDTHTIMINGS {
	var c C.D3DDEVINFO_D3D9BANDWIDTHTIMINGS
	c.MaxBandwidthUtilized = (C.float)(s.MaxBandwidthUtilized)
	c.FrontEndUploadMemoryUtilizedPercent = (C.float)(s.FrontEndUploadMemoryUtilizedPercent)
	c.VertexRateUtilizedPercent = (C.float)(s.VertexRateUtilizedPercent)
	c.TriangleSetupRateUtilizedPercent = (C.float)(s.TriangleSetupRateUtilizedPercent)
	c.FillRateUtilizedPercent = (C.float)(s.FillRateUtilizedPercent)
	return c
}

func (s *DEVINFO_D3D9BANDWIDTHTIMINGS) fromC(c *C.D3DDEVINFO_D3D9BANDWIDTHTIMINGS) {
	s.MaxBandwidthUtilized = (float32)(c.MaxBandwidthUtilized)
	s.FrontEndUploadMemoryUtilizedPercent = (float32)(c.FrontEndUploadMemoryUtilizedPercent)
	s.VertexRateUtilizedPercent = (float32)(c.VertexRateUtilizedPercent)
	s.TriangleSetupRateUtilizedPercent = (float32)(c.TriangleSetupRateUtilizedPercent)
	s.FillRateUtilizedPercent = (float32)(c.FillRateUtilizedPercent)
}

// DEVINFO_D3D9CACHEUTILIZATION measures the cache hit rate performance for
// textures and indexed vertices.
type DEVINFO_D3D9CACHEUTILIZATION struct {
	TextureCacheHitRate             float32
	PostTransformVertexCacheHitRate float32
}

func (s *DEVINFO_D3D9CACHEUTILIZATION) toC() C.D3DDEVINFO_D3D9CACHEUTILIZATION {
	var c C.D3DDEVINFO_D3D9CACHEUTILIZATION
	c.TextureCacheHitRate = (C.float)(s.TextureCacheHitRate)
	c.PostTransformVertexCacheHitRate = (C.float)(s.PostTransformVertexCacheHitRate)
	return c
}

func (s *DEVINFO_D3D9CACHEUTILIZATION) fromC(c *C.D3DDEVINFO_D3D9CACHEUTILIZATION) {
	s.TextureCacheHitRate = (float32)(c.TextureCacheHitRate)
	s.PostTransformVertexCacheHitRate = (float32)(c.PostTransformVertexCacheHitRate)
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

func (s *DEVINFO_D3D9INTERFACETIMINGS) toC() C.D3DDEVINFO_D3D9INTERFACETIMINGS {
	var c C.D3DDEVINFO_D3D9INTERFACETIMINGS
	c.WaitingForGPUToUseApplicationResourceTimePercent = (C.float)(s.WaitingForGPUToUseApplicationResourceTimePercent)
	c.WaitingForGPUToAcceptMoreCommandsTimePercent = (C.float)(s.WaitingForGPUToAcceptMoreCommandsTimePercent)
	c.WaitingForGPUToStayWithinLatencyTimePercent = (C.float)(s.WaitingForGPUToStayWithinLatencyTimePercent)
	c.WaitingForGPUExclusiveResourceTimePercent = (C.float)(s.WaitingForGPUExclusiveResourceTimePercent)
	c.WaitingForGPUOtherTimePercent = (C.float)(s.WaitingForGPUOtherTimePercent)
	return c
}

func (s *DEVINFO_D3D9INTERFACETIMINGS) fromC(c *C.D3DDEVINFO_D3D9INTERFACETIMINGS) {
	s.WaitingForGPUToUseApplicationResourceTimePercent = (float32)(c.WaitingForGPUToUseApplicationResourceTimePercent)
	s.WaitingForGPUToAcceptMoreCommandsTimePercent = (float32)(c.WaitingForGPUToAcceptMoreCommandsTimePercent)
	s.WaitingForGPUToStayWithinLatencyTimePercent = (float32)(c.WaitingForGPUToStayWithinLatencyTimePercent)
	s.WaitingForGPUExclusiveResourceTimePercent = (float32)(c.WaitingForGPUExclusiveResourceTimePercent)
	s.WaitingForGPUOtherTimePercent = (float32)(c.WaitingForGPUOtherTimePercent)
}

// DEVINFO_D3D9PIPELINETIMINGS contains the percent of time processing data in
// the pipeline.
type DEVINFO_D3D9PIPELINETIMINGS struct {
	VertexProcessingTimePercent   float32
	PixelProcessingTimePercent    float32
	OtherGPUProcessingTimePercent float32
	GPUIdleTimePercent            float32
}

func (s *DEVINFO_D3D9PIPELINETIMINGS) toC() C.D3DDEVINFO_D3D9PIPELINETIMINGS {
	var c C.D3DDEVINFO_D3D9PIPELINETIMINGS
	c.VertexProcessingTimePercent = (C.float)(s.VertexProcessingTimePercent)
	c.PixelProcessingTimePercent = (C.float)(s.PixelProcessingTimePercent)
	c.OtherGPUProcessingTimePercent = (C.float)(s.OtherGPUProcessingTimePercent)
	c.GPUIdleTimePercent = (C.float)(s.GPUIdleTimePercent)
	return c
}

func (s *DEVINFO_D3D9PIPELINETIMINGS) fromC(c *C.D3DDEVINFO_D3D9PIPELINETIMINGS) {
	s.VertexProcessingTimePercent = (float32)(c.VertexProcessingTimePercent)
	s.PixelProcessingTimePercent = (float32)(c.PixelProcessingTimePercent)
	s.OtherGPUProcessingTimePercent = (float32)(c.OtherGPUProcessingTimePercent)
	s.GPUIdleTimePercent = (float32)(c.GPUIdleTimePercent)
}

// DEVINFO_D3D9STAGETIMINGS contains the percent of time processing shader data.
type DEVINFO_D3D9STAGETIMINGS struct {
	MemoryProcessingPercent      float32
	ComputationProcessingPercent float32
}

func (s *DEVINFO_D3D9STAGETIMINGS) toC() C.D3DDEVINFO_D3D9STAGETIMINGS {
	var c C.D3DDEVINFO_D3D9STAGETIMINGS
	c.MemoryProcessingPercent = (C.float)(s.MemoryProcessingPercent)
	c.ComputationProcessingPercent = (C.float)(s.ComputationProcessingPercent)
	return c
}

func (s *DEVINFO_D3D9STAGETIMINGS) fromC(c *C.D3DDEVINFO_D3D9STAGETIMINGS) {
	s.MemoryProcessingPercent = (float32)(c.MemoryProcessingPercent)
	s.ComputationProcessingPercent = (float32)(c.ComputationProcessingPercent)
}

// DEVINFO_D3DVERTEXSTATS reports the number of triangles that have been
// processed and clipped by the runtime's software vertex processing.
type DEVINFO_D3DVERTEXSTATS struct {
	NumRenderedTriangles      uint32
	NumExtraClippingTriangles uint32
}

func (s *DEVINFO_D3DVERTEXSTATS) toC() C.D3DDEVINFO_D3DVERTEXSTATS {
	var c C.D3DDEVINFO_D3DVERTEXSTATS
	c.NumRenderedTriangles = (C.DWORD)(s.NumRenderedTriangles)
	c.NumExtraClippingTriangles = (C.DWORD)(s.NumExtraClippingTriangles)
	return c
}

func (s *DEVINFO_D3DVERTEXSTATS) fromC(c *C.D3DDEVINFO_D3DVERTEXSTATS) {
	s.NumRenderedTriangles = (uint32)(c.NumRenderedTriangles)
	s.NumExtraClippingTriangles = (uint32)(c.NumExtraClippingTriangles)
}

// DEVINFO_VCACHE contains vertex cache optimization hints.
type DEVINFO_VCACHE struct {
	Pattern     uint32
	OptMethod   uint32
	CacheSize   uint32
	MagicNumber uint32
}

func (s *DEVINFO_VCACHE) toC() C.D3DDEVINFO_VCACHE {
	var c C.D3DDEVINFO_VCACHE
	c.Pattern = (C.DWORD)(s.Pattern)
	c.OptMethod = (C.DWORD)(s.OptMethod)
	c.CacheSize = (C.DWORD)(s.CacheSize)
	c.MagicNumber = (C.DWORD)(s.MagicNumber)
	return c
}

func (s *DEVINFO_VCACHE) fromC(c *C.D3DDEVINFO_VCACHE) {
	s.Pattern = (uint32)(c.Pattern)
	s.OptMethod = (uint32)(c.OptMethod)
	s.CacheSize = (uint32)(c.CacheSize)
	s.MagicNumber = (uint32)(c.MagicNumber)
}

// DISPLAYMODE describes the display mode.
type DISPLAYMODE struct {
	Width       uint
	Height      uint
	RefreshRate uint
	Format      FORMAT
}

func (s *DISPLAYMODE) toC() C.D3DDISPLAYMODE {
	var c C.D3DDISPLAYMODE
	c.Width = (C.UINT)(s.Width)
	c.Height = (C.UINT)(s.Height)
	c.RefreshRate = (C.UINT)(s.RefreshRate)
	c.Format = (C.D3DFORMAT)(s.Format)
	return c
}

func (s *DISPLAYMODE) fromC(c *C.D3DDISPLAYMODE) {
	s.Width = (uint)(c.Width)
	s.Height = (uint)(c.Height)
	s.RefreshRate = (uint)(c.RefreshRate)
	s.Format = (FORMAT)(c.Format)
}

// GAMMARAMP contains red, green, and blue ramp data.
type GAMMARAMP struct {
	Red   [256]uint16
	Green [256]uint16
	Blue  [256]uint16
}

func (s *GAMMARAMP) toC() C.D3DGAMMARAMP {
	var c C.D3DGAMMARAMP
	for i := range s.Red {
		c.red[i] = (C.WORD)(s.Red[i])
	}
	for i := range s.Green {
		c.green[i] = (C.WORD)(s.Green[i])
	}
	for i := range s.Blue {
		c.blue[i] = (C.WORD)(s.Blue[i])
	}
	return c
}

func (s *GAMMARAMP) fromC(c *C.D3DGAMMARAMP) {
	for i := range s.Red {
		s.Red[i] = (uint16)(c.red[i])
	}
	for i := range s.Green {
		s.Green[i] = (uint16)(c.green[i])
	}
	for i := range s.Blue {
		s.Blue[i] = (uint16)(c.blue[i])
	}
}

// INDEXBUFFER_DESC describes an index buffer.
type INDEXBUFFER_DESC struct {
	Format FORMAT
	Type   RESOURCETYPE
	Usage  uint32
	Pool   POOL
	Size   uint
}

func (s *INDEXBUFFER_DESC) toC() C.D3DINDEXBUFFER_DESC {
	var c C.D3DINDEXBUFFER_DESC
	c.Format = (C.D3DFORMAT)(s.Format)
	c.Type = (C.D3DRESOURCETYPE)(s.Type)
	c.Usage = (C.DWORD)(s.Usage)
	c.Pool = (C.D3DPOOL)(s.Pool)
	c.Size = (C.UINT)(s.Size)
	return c
}

func (s *INDEXBUFFER_DESC) fromC(c *C.D3DINDEXBUFFER_DESC) {
	s.Format = (FORMAT)(c.Format)
	s.Type = (RESOURCETYPE)(c.Type)
	s.Usage = (uint32)(c.Usage)
	s.Pool = (POOL)(c.Pool)
	s.Size = (uint)(c.Size)
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

func (s *LIGHT) toC() C.D3DLIGHT9 {
	var c C.D3DLIGHT9
	c.Type = (C.D3DLIGHTTYPE)(s.Type)
	c.Diffuse = s.Diffuse.toC()
	c.Specular = s.Specular.toC()
	c.Ambient = s.Ambient.toC()
	c.Position = s.Position.toC()
	c.Direction = s.Direction.toC()
	c.Range = (C.float)(s.Range)
	c.Falloff = (C.float)(s.Falloff)
	c.Attenuation0 = (C.float)(s.Attenuation0)
	c.Attenuation1 = (C.float)(s.Attenuation1)
	c.Attenuation2 = (C.float)(s.Attenuation2)
	c.Theta = (C.float)(s.Theta)
	c.Phi = (C.float)(s.Phi)
	return c
}

func (s *LIGHT) fromC(c *C.D3DLIGHT9) {
	s.Type = (LIGHTTYPE)(c.Type)
	s.Diffuse.fromC(&c.Diffuse)
	s.Specular.fromC(&c.Specular)
	s.Ambient.fromC(&c.Ambient)
	s.Position.fromC(&c.Position)
	s.Direction.fromC(&c.Direction)
	s.Range = (float32)(c.Range)
	s.Falloff = (float32)(c.Falloff)
	s.Attenuation0 = (float32)(c.Attenuation0)
	s.Attenuation1 = (float32)(c.Attenuation1)
	s.Attenuation2 = (float32)(c.Attenuation2)
	s.Theta = (float32)(c.Theta)
	s.Phi = (float32)(c.Phi)
}

// LOCKED_BOX describes a locked box (volume).
type LOCKED_BOX struct {
	RowPitch   int
	SlicePitch int
	PBits      unsafe.Pointer
}

func (s *LOCKED_BOX) toC() C.D3DLOCKED_BOX {
	var c C.D3DLOCKED_BOX
	c.RowPitch = (C.INT)(s.RowPitch)
	c.SlicePitch = (C.INT)(s.SlicePitch)
	c.pBits = s.PBits
	return c
}

func (s *LOCKED_BOX) fromC(c *C.D3DLOCKED_BOX) {
	s.RowPitch = (int)(c.RowPitch)
	s.SlicePitch = (int)(c.SlicePitch)
	s.PBits = (unsafe.Pointer)(c.pBits)
}

// LOCKED_RECT describes a locked rectangular region.
type LOCKED_RECT struct {
	Pitch int
	PBits unsafe.Pointer
}

func (s *LOCKED_RECT) toC() C.D3DLOCKED_RECT {
	var c C.D3DLOCKED_RECT
	c.Pitch = (C.INT)(s.Pitch)
	c.pBits = s.PBits
	return c
}

func (s *LOCKED_RECT) fromC(c *C.D3DLOCKED_RECT) {
	s.Pitch = (int)(c.Pitch)
	s.PBits = (unsafe.Pointer)(c.pBits)
}

// SetAllBytes will fill the whole rect with the given data, taking into account
// the rect's pitch. The given byte slice is expected to have the given stride
// in bytes, i.e. one line in the given data is <stride> bytes in length. If the
// byte slice is tightly packed, stride will be 0.
func (r LOCKED_RECT) SetAllBytes(data []byte, stride int) {
	C.setWholeLockedRect(
		r.PBits,
		C.int(r.Pitch),
		unsafe.Pointer(&data[0]),
		C.int(stride),
		C.int(len(data)/stride),
	)
}

// MATERIAL specifies material properties.
type MATERIAL struct {
	Diffuse  COLORVALUE
	Ambient  COLORVALUE
	Specular COLORVALUE
	Emissive COLORVALUE
	Power    float32
}

func (s *MATERIAL) toC() C.D3DMATERIAL9 {
	var c C.D3DMATERIAL9
	c.Diffuse = s.Diffuse.toC()
	c.Ambient = s.Ambient.toC()
	c.Specular = s.Specular.toC()
	c.Emissive = s.Emissive.toC()
	c.Power = (C.float)(s.Power)
	return c
}

func (s *MATERIAL) fromC(c *C.D3DMATERIAL9) {
	s.Diffuse.fromC(&c.Diffuse)
	s.Ambient.fromC(&c.Ambient)
	s.Specular.fromC(&c.Specular)
	s.Emissive.fromC(&c.Emissive)
	s.Power = (float32)(c.Power)
}

// MATRIX escribes a matrix.
type MATRIX [16]float32

func (s *MATRIX) toC() C.D3DMATRIX {
	var c C.D3DMATRIX
	C.toCMatrix(&c, (*C.float)(unsafe.Pointer(&s[0])))
	return c
}

func (s *MATRIX) fromC(c *C.D3DMATRIX) {
	C.toGoMatrix(c, (*C.float)(unsafe.Pointer(&(*s)[0])))
}

// PRESENT_PARAMETERS describes the presentation parameters.
type PRESENT_PARAMETERS struct {
	BackBufferWidth            uint
	BackBufferHeight           uint
	BackBufferFormat           FORMAT
	BackBufferCount            uint
	MultiSampleType            MULTISAMPLE_TYPE
	MultiSampleQuality         uint32
	SwapEffect                 SWAPEFFECT
	HDeviceWindow              unsafe.Pointer
	Windowed                   bool
	EnableAutoDepthStencil     bool
	AutoDepthStencilFormat     FORMAT
	Flags                      uint32
	FullScreen_RefreshRateInHz uint
	PresentationInterval       uint
}

func (s *PRESENT_PARAMETERS) toC() C.D3DPRESENT_PARAMETERS {
	var c C.D3DPRESENT_PARAMETERS
	c.BackBufferWidth = (C.UINT)(s.BackBufferWidth)
	c.BackBufferHeight = (C.UINT)(s.BackBufferHeight)
	c.BackBufferFormat = (C.D3DFORMAT)(s.BackBufferFormat)
	c.BackBufferCount = (C.UINT)(s.BackBufferCount)
	c.MultiSampleType = (C.D3DMULTISAMPLE_TYPE)(s.MultiSampleType)
	c.MultiSampleQuality = (C.DWORD)(s.MultiSampleQuality)
	c.SwapEffect = (C.D3DSWAPEFFECT)(s.SwapEffect)
	c.hDeviceWindow = (C.HWND)(s.HDeviceWindow)
	c.Windowed = 0
	if s.Windowed {
		c.Windowed = 1
	}
	c.EnableAutoDepthStencil = 0
	if s.EnableAutoDepthStencil {
		c.EnableAutoDepthStencil = 1
	}
	c.AutoDepthStencilFormat = (C.D3DFORMAT)(s.AutoDepthStencilFormat)
	c.Flags = (C.DWORD)(s.Flags)
	c.FullScreen_RefreshRateInHz = (C.UINT)(s.FullScreen_RefreshRateInHz)
	c.PresentationInterval = (C.UINT)(s.PresentationInterval)
	return c
}

func (s *PRESENT_PARAMETERS) fromC(c *C.D3DPRESENT_PARAMETERS) {
	s.BackBufferWidth = (uint)(c.BackBufferWidth)
	s.BackBufferHeight = (uint)(c.BackBufferHeight)
	s.BackBufferFormat = (FORMAT)(c.BackBufferFormat)
	s.BackBufferCount = (uint)(c.BackBufferCount)
	s.MultiSampleType = (MULTISAMPLE_TYPE)(c.MultiSampleType)
	s.MultiSampleQuality = (uint32)(c.MultiSampleQuality)
	s.SwapEffect = (SWAPEFFECT)(c.SwapEffect)
	s.HDeviceWindow = (unsafe.Pointer)(c.hDeviceWindow)
	s.Windowed = c.Windowed != 0
	s.EnableAutoDepthStencil = c.EnableAutoDepthStencil != 0
	s.AutoDepthStencilFormat = (FORMAT)(c.AutoDepthStencilFormat)
	s.Flags = (uint32)(c.Flags)
	s.FullScreen_RefreshRateInHz = (uint)(c.FullScreen_RefreshRateInHz)
	s.PresentationInterval = (uint)(c.PresentationInterval)
}

// PSHADERCAPS2_0 describes pixel shader driver caps.
type PSHADERCAPS2_0 struct {
	Caps                    uint32
	DynamicFlowControlDepth int
	NumTemps                int
	StaticFlowControlDepth  int
	NumInstructionSlots     int
}

func (s *PSHADERCAPS2_0) toC() C.D3DPSHADERCAPS2_0 {
	var c C.D3DPSHADERCAPS2_0
	c.Caps = (C.DWORD)(s.Caps)
	c.DynamicFlowControlDepth = (C.INT)(s.DynamicFlowControlDepth)
	c.NumTemps = (C.INT)(s.NumTemps)
	c.StaticFlowControlDepth = (C.INT)(s.StaticFlowControlDepth)
	c.NumInstructionSlots = (C.INT)(s.NumInstructionSlots)
	return c
}

func (s *PSHADERCAPS2_0) fromC(c *C.D3DPSHADERCAPS2_0) {
	s.Caps = (uint32)(c.Caps)
	s.DynamicFlowControlDepth = (int)(c.DynamicFlowControlDepth)
	s.NumTemps = (int)(c.NumTemps)
	s.StaticFlowControlDepth = (int)(c.StaticFlowControlDepth)
	s.NumInstructionSlots = (int)(c.NumInstructionSlots)
}

// RANGE defines a range.
type RANGE struct {
	Offset uint
	Size   uint
}

func (s *RANGE) toC() C.D3DRANGE {
	var c C.D3DRANGE
	c.Offset = (C.UINT)(s.Offset)
	c.Size = (C.UINT)(s.Size)
	return c
}

func (s *RANGE) fromC(c *C.D3DRANGE) {
	s.Offset = (uint)(c.Offset)
	s.Size = (uint)(c.Size)
}

// RASTER_STATUS describes the raster status.
type RASTER_STATUS struct {
	InVBlank bool
	ScanLine uint
}

func (s *RASTER_STATUS) toC() C.D3DRASTER_STATUS {
	var c C.D3DRASTER_STATUS
	c.InVBlank = 0
	if s.InVBlank {
		c.InVBlank = 1
	}
	c.ScanLine = (C.UINT)(s.ScanLine)
	return c
}

func (s *RASTER_STATUS) fromC(c *C.D3DRASTER_STATUS) {
	s.InVBlank = c.InVBlank != 0
	s.ScanLine = (uint)(c.ScanLine)
}

// D3DRECT defines a rectangle.
type D3DRECT struct {
	X1 int32
	Y1 int32
	X2 int32
	Y2 int32
}

func (s *D3DRECT) toC() C.D3DRECT {
	var c C.D3DRECT
	c.x1 = (C.LONG)(s.X1)
	c.y1 = (C.LONG)(s.Y1)
	c.x2 = (C.LONG)(s.X2)
	c.y2 = (C.LONG)(s.Y2)
	return c
}

func (s *D3DRECT) fromC(c *C.D3DRECT) {
	s.X1 = (int32)(c.x1)
	s.Y1 = (int32)(c.y1)
	s.X2 = (int32)(c.x2)
	s.Y2 = (int32)(c.y2)
}

// RECTPATCH_INFO describes a rectangular high-order patch.
type RECTPATCH_INFO struct {
	StartVertexOffsetWidth  uint
	StartVertexOffsetHeight uint
	Width                   uint
	Height                  uint
	Stride                  uint
	Basis                   BASISTYPE
	Degree                  DEGREETYPE
}

func (s *RECTPATCH_INFO) toC() C.D3DRECTPATCH_INFO {
	var c C.D3DRECTPATCH_INFO
	c.StartVertexOffsetWidth = (C.UINT)(s.StartVertexOffsetWidth)
	c.StartVertexOffsetHeight = (C.UINT)(s.StartVertexOffsetHeight)
	c.Width = (C.UINT)(s.Width)
	c.Height = (C.UINT)(s.Height)
	c.Stride = (C.UINT)(s.Stride)
	c.Basis = (C.D3DBASISTYPE)(s.Basis)
	c.Degree = (C.D3DDEGREETYPE)(s.Degree)
	return c
}

func (s *RECTPATCH_INFO) fromC(c *C.D3DRECTPATCH_INFO) {
	s.StartVertexOffsetWidth = (uint)(c.StartVertexOffsetWidth)
	s.StartVertexOffsetHeight = (uint)(c.StartVertexOffsetHeight)
	s.Width = (uint)(c.Width)
	s.Height = (uint)(c.Height)
	s.Stride = (uint)(c.Stride)
	s.Basis = (BASISTYPE)(c.Basis)
	s.Degree = (DEGREETYPE)(c.Degree)
}

// RESOURCESTATS contains resource statistics gathered by the
// DEVINFO_ResourceManager when using the asynchronous query mechanism.
type RESOURCESTATS struct {
	BThrashing            bool
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

func (s *RESOURCESTATS) toC() C.D3DRESOURCESTATS {
	var c C.D3DRESOURCESTATS
	c.bThrashing = 0
	if s.BThrashing {
		c.bThrashing = 1
	}
	c.ApproxBytesDownloaded = (C.DWORD)(s.ApproxBytesDownloaded)
	c.NumEvicts = (C.DWORD)(s.NumEvicts)
	c.NumVidCreates = (C.DWORD)(s.NumVidCreates)
	c.LastPri = (C.DWORD)(s.LastPri)
	c.NumUsed = (C.DWORD)(s.NumUsed)
	c.NumUsedInVidMem = (C.DWORD)(s.NumUsedInVidMem)
	c.WorkingSet = (C.DWORD)(s.WorkingSet)
	c.WorkingSetBytes = (C.DWORD)(s.WorkingSetBytes)
	c.TotalManaged = (C.DWORD)(s.TotalManaged)
	c.TotalBytes = (C.DWORD)(s.TotalBytes)
	return c
}

func (s *RESOURCESTATS) fromC(c *C.D3DRESOURCESTATS) {
	s.BThrashing = c.bThrashing != 0
	s.ApproxBytesDownloaded = (uint32)(c.ApproxBytesDownloaded)
	s.NumEvicts = (uint32)(c.NumEvicts)
	s.NumVidCreates = (uint32)(c.NumVidCreates)
	s.LastPri = (uint32)(c.LastPri)
	s.NumUsed = (uint32)(c.NumUsed)
	s.NumUsedInVidMem = (uint32)(c.NumUsedInVidMem)
	s.WorkingSet = (uint32)(c.WorkingSet)
	s.WorkingSetBytes = (uint32)(c.WorkingSetBytes)
	s.TotalManaged = (uint32)(c.TotalManaged)
	s.TotalBytes = (uint32)(c.TotalBytes)
}

// SURFACE_DESC describes a surface.
type SURFACE_DESC struct {
	Format             FORMAT
	Type               RESOURCETYPE
	Usage              uint32
	Pool               POOL
	MultiSampleType    MULTISAMPLE_TYPE
	MultiSampleQuality uint32
	Width              uint
	Height             uint
}

func (s *SURFACE_DESC) toC() C.D3DSURFACE_DESC {
	var c C.D3DSURFACE_DESC
	c.Format = (C.D3DFORMAT)(s.Format)
	c.Type = (C.D3DRESOURCETYPE)(s.Type)
	c.Usage = (C.DWORD)(s.Usage)
	c.Pool = (C.D3DPOOL)(s.Pool)
	c.MultiSampleType = (C.D3DMULTISAMPLE_TYPE)(s.MultiSampleType)
	c.MultiSampleQuality = (C.DWORD)(s.MultiSampleQuality)
	c.Width = (C.UINT)(s.Width)
	c.Height = (C.UINT)(s.Height)
	return c
}

func (s *SURFACE_DESC) fromC(c *C.D3DSURFACE_DESC) {
	s.Format = (FORMAT)(c.Format)
	s.Type = (RESOURCETYPE)(c.Type)
	s.Usage = (uint32)(c.Usage)
	s.Pool = (POOL)(c.Pool)
	s.MultiSampleType = (MULTISAMPLE_TYPE)(c.MultiSampleType)
	s.MultiSampleQuality = (uint32)(c.MultiSampleQuality)
	s.Width = (uint)(c.Width)
	s.Height = (uint)(c.Height)
}

// TRIPATCH_INFO describes a triangular high-order patch.
type TRIPATCH_INFO struct {
	StartVertexOffset uint
	NumVertices       uint
	Basis             BASISTYPE
	Degree            DEGREETYPE
}

func (s *TRIPATCH_INFO) toC() C.D3DTRIPATCH_INFO {
	var c C.D3DTRIPATCH_INFO
	c.StartVertexOffset = (C.UINT)(s.StartVertexOffset)
	c.NumVertices = (C.UINT)(s.NumVertices)
	c.Basis = (C.D3DBASISTYPE)(s.Basis)
	c.Degree = (C.D3DDEGREETYPE)(s.Degree)
	return c
}

func (s *TRIPATCH_INFO) fromC(c *C.D3DTRIPATCH_INFO) {
	s.StartVertexOffset = (uint)(c.StartVertexOffset)
	s.NumVertices = (uint)(c.NumVertices)
	s.Basis = (BASISTYPE)(c.Basis)
	s.Degree = (DEGREETYPE)(c.Degree)
}

// VECTOR defines a vector.
type VECTOR struct {
	X float32
	Y float32
	Z float32
}

func (s *VECTOR) toC() C.D3DVECTOR {
	var c C.D3DVECTOR
	c.x = (C.float)(s.X)
	c.y = (C.float)(s.Y)
	c.z = (C.float)(s.Z)
	return c
}

func (s *VECTOR) fromC(c *C.D3DVECTOR) {
	s.X = (float32)(c.x)
	s.Y = (float32)(c.y)
	s.Z = (float32)(c.z)
}

// VERTEXBUFFER_DESC describes a vertex buffer.
type VERTEXBUFFER_DESC struct {
	Format FORMAT
	Type   RESOURCETYPE
	Usage  uint32
	Pool   POOL
	Size   uint
	FVF    uint32
}

func (s *VERTEXBUFFER_DESC) toC() C.D3DVERTEXBUFFER_DESC {
	var c C.D3DVERTEXBUFFER_DESC
	c.Format = (C.D3DFORMAT)(s.Format)
	c.Type = (C.D3DRESOURCETYPE)(s.Type)
	c.Usage = (C.DWORD)(s.Usage)
	c.Pool = (C.D3DPOOL)(s.Pool)
	c.Size = (C.UINT)(s.Size)
	c.FVF = (C.DWORD)(s.FVF)
	return c
}

func (s *VERTEXBUFFER_DESC) fromC(c *C.D3DVERTEXBUFFER_DESC) {
	s.Format = (FORMAT)(c.Format)
	s.Type = (RESOURCETYPE)(c.Type)
	s.Usage = (uint32)(c.Usage)
	s.Pool = (POOL)(c.Pool)
	s.Size = (uint)(c.Size)
	s.FVF = (uint32)(c.FVF)
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

func (s *VERTEXELEMENT) toC() C.D3DVERTEXELEMENT9 {
	var c C.D3DVERTEXELEMENT9
	c.Stream = (C.WORD)(s.Stream)
	c.Offset = (C.WORD)(s.Offset)
	c.Type = (C.BYTE)(s.Type)
	c.Method = (C.BYTE)(s.Method)
	c.Usage = (C.BYTE)(s.Usage)
	c.UsageIndex = (C.BYTE)(s.UsageIndex)
	return c
}

func (s *VERTEXELEMENT) fromC(c *C.D3DVERTEXELEMENT9) {
	s.Stream = (uint16)(c.Stream)
	s.Offset = (uint16)(c.Offset)
	s.Type = (DECLTYPE)(c.Type)
	s.Method = (DECLMETHOD)(c.Method)
	s.Usage = (DECLUSAGE)(c.Usage)
	s.UsageIndex = (byte)(c.UsageIndex)
}

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

func (s *VIEWPORT) toC() C.D3DVIEWPORT9 {
	var c C.D3DVIEWPORT9
	c.X = (C.DWORD)(s.X)
	c.Y = (C.DWORD)(s.Y)
	c.Width = (C.DWORD)(s.Width)
	c.Height = (C.DWORD)(s.Height)
	c.MinZ = (C.float)(s.MinZ)
	c.MaxZ = (C.float)(s.MaxZ)
	return c
}

func (s *VIEWPORT) fromC(c *C.D3DVIEWPORT9) {
	s.X = (uint32)(c.X)
	s.Y = (uint32)(c.Y)
	s.Width = (uint32)(c.Width)
	s.Height = (uint32)(c.Height)
	s.MinZ = (float32)(c.MinZ)
	s.MaxZ = (float32)(c.MaxZ)
}

// VOLUME_DESC describes a volume.
type VOLUME_DESC struct {
	Format FORMAT
	Type   RESOURCETYPE
	Usage  uint32
	Pool   POOL
	Width  uint
	Height uint
	Depth  uint
}

func (s *VOLUME_DESC) toC() C.D3DVOLUME_DESC {
	var c C.D3DVOLUME_DESC
	c.Format = (C.D3DFORMAT)(s.Format)
	c.Type = (C.D3DRESOURCETYPE)(s.Type)
	c.Usage = (C.DWORD)(s.Usage)
	c.Pool = (C.D3DPOOL)(s.Pool)
	c.Width = (C.UINT)(s.Width)
	c.Height = (C.UINT)(s.Height)
	c.Depth = (C.UINT)(s.Depth)
	return c
}

func (s *VOLUME_DESC) fromC(c *C.D3DVOLUME_DESC) {
	s.Format = (FORMAT)(c.Format)
	s.Type = (RESOURCETYPE)(c.Type)
	s.Usage = (uint32)(c.Usage)
	s.Pool = (POOL)(c.Pool)
	s.Width = (uint)(c.Width)
	s.Height = (uint)(c.Height)
	s.Depth = (uint)(c.Depth)
}

// VSHADERCAPS2_0 contains vertex shader capabilities.
type VSHADERCAPS2_0 struct {
	Caps                    uint32
	DynamicFlowControlDepth int
	NumTemps                int
	StaticFlowControlDepth  int
}

func (s *VSHADERCAPS2_0) toC() C.D3DVSHADERCAPS2_0 {
	var c C.D3DVSHADERCAPS2_0
	c.Caps = (C.DWORD)(s.Caps)
	c.DynamicFlowControlDepth = (C.INT)(s.DynamicFlowControlDepth)
	c.NumTemps = (C.INT)(s.NumTemps)
	c.StaticFlowControlDepth = (C.INT)(s.StaticFlowControlDepth)
	return c
}

func (s *VSHADERCAPS2_0) fromC(c *C.D3DVSHADERCAPS2_0) {
	s.Caps = (uint32)(c.Caps)
	s.DynamicFlowControlDepth = (int)(c.DynamicFlowControlDepth)
	s.NumTemps = (int)(c.NumTemps)
	s.StaticFlowControlDepth = (int)(c.StaticFlowControlDepth)
}

// PALETTEENTRY specifies the color and usage of an entry in a logical palette.
type PALETTEENTRY struct {
	PeRed   byte
	PeGreen byte
	PeBlue  byte
	PeFlags byte
}

func (s *PALETTEENTRY) toC() C.PALETTEENTRY {
	var c C.PALETTEENTRY
	c.peRed = (C.BYTE)(s.PeRed)
	c.peGreen = (C.BYTE)(s.PeGreen)
	c.peBlue = (C.BYTE)(s.PeBlue)
	c.peFlags = (C.BYTE)(s.PeFlags)
	return c
}

func (s *PALETTEENTRY) fromC(c *C.PALETTEENTRY) {
	s.PeRed = (byte)(c.peRed)
	s.PeGreen = (byte)(c.peGreen)
	s.PeBlue = (byte)(c.peBlue)
	s.PeFlags = (byte)(c.peFlags)
}

// GUID is a globally unique identifier.
type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]uint8
}

func (s *GUID) toC() C.GUID {
	var c C.GUID
	c.Data1 = (C.ulong)(s.Data1)
	c.Data2 = (C.ushort)(s.Data2)
	c.Data3 = (C.ushort)(s.Data3)
	for i := range s.Data4 {
		c.Data4[i] = (C.uchar)(s.Data4[i])
	}
	return c
}

func (s *GUID) fromC(c *C.GUID) {
	s.Data1 = (uint32)(c.Data1)
	s.Data2 = (uint16)(c.Data2)
	s.Data3 = (uint16)(c.Data3)
	for i := range s.Data4 {
		s.Data4[i] = (uint8)(c.Data4[i])
	}
}

// LUID is a locally unique identifier.
type LUID struct {
	LowPart  uint32
	HighPart int32
}

func (s *LUID) toC() C.LUID {
	var c C.LUID
	c.LowPart = (C.DWORD)(s.LowPart)
	c.HighPart = (C.LONG)(s.HighPart)
	return c
}

func (s *LUID) fromC(c *C.LUID) {
	s.LowPart = (uint32)(c.LowPart)
	s.HighPart = (int32)(c.HighPart)
}

// RECT describes a rectangle.
type RECT struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

func (s *RECT) toC() C.RECT {
	var c C.RECT
	c.left = (C.LONG)(s.Left)
	c.top = (C.LONG)(s.Top)
	c.right = (C.LONG)(s.Right)
	c.bottom = (C.LONG)(s.Bottom)
	return c
}

func (s *RECT) fromC(c *C.RECT) {
	s.Left = (int32)(c.left)
	s.Top = (int32)(c.top)
	s.Right = (int32)(c.right)
	s.Bottom = (int32)(c.bottom)
}

// RGNDATAHEADER describes region data.
type RGNDATAHEADER struct {
	DwSize   uint32
	IType    uint32
	NCount   uint32
	NRgnSize uint32
	RcBound  RECT
}

func (s *RGNDATAHEADER) toC() C.RGNDATAHEADER {
	var c C.RGNDATAHEADER
	c.dwSize = (C.DWORD)(s.DwSize)
	c.iType = (C.DWORD)(s.IType)
	c.nCount = (C.DWORD)(s.NCount)
	c.nRgnSize = (C.DWORD)(s.NRgnSize)
	c.rcBound = s.RcBound.toC()
	return c
}

func (s *RGNDATAHEADER) fromC(c *C.RGNDATAHEADER) {
	s.DwSize = (uint32)(c.dwSize)
	s.IType = (uint32)(c.iType)
	s.NCount = (uint32)(c.nCount)
	s.NRgnSize = (uint32)(c.nRgnSize)
	s.RcBound.fromC(&c.rcBound)
}

// RGNDATA contains region data.
type RGNDATA struct {
	Rdh    RGNDATAHEADER
	Buffer [1]byte
}

func (s *RGNDATA) toC() C.RGNDATA {
	var c C.RGNDATA
	c.rdh = s.Rdh.toC()
	for i := range s.Buffer {
		c.Buffer[i] = (C.char)(s.Buffer[i])
	}
	return c
}

func (s *RGNDATA) fromC(c *C.RGNDATA) {
	s.Rdh.fromC(&c.rdh)
	for i := range s.Buffer {
		s.Buffer[i] = (byte)(c.Buffer[i])
	}
}

// POINT describes a 2D point.
type POINT struct {
	X int32
	Y int32
}

func (s *POINT) toC() C.POINT {
	var c C.POINT
	c.x = (C.LONG)(s.X)
	c.y = (C.LONG)(s.Y)
	return c
}

func (s *POINT) fromC(c *C.POINT) {
	s.X = (int32)(c.x)
	s.Y = (int32)(c.y)
}
