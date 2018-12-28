package d3d9

// BACKBUFFER_TYPE defines constants that describe the type of a back buffer.
type BACKBUFFER_TYPE uint32

const (
	// BACKBUFFER_TYPE_MONO specifies a nonstereo swap chain.
	BACKBUFFER_TYPE_MONO = 0
	// BACKBUFFER_TYPE_LEFT specifies the left side of a stereo pair in a swap
	// chain.
	BACKBUFFER_TYPE_LEFT = 1
	// BACKBUFFER_TYPE_RIGHT specifies the right side of a stereo pair in a
	// swap chain.
	BACKBUFFER_TYPE_RIGHT = 2
)

// BASISTYPE defines the basis type of a high-order patch surface.
type BASISTYPE uint32

const (
	// BASIS_BEZIER means input vertices are treated as a series of Bézier
	// patches. The number of vertices specified must be divisible by 4.
	// Portions of the mesh beyond this criterion will not be rendered. Full
	// continuity is assumed between sub-patches in the interior of the surface
	// rendered by each call. Only the vertices at the corners of each
	// sub-patch are guaranteed to lie on the resulting surface.
	BASIS_BEZIER = 0
	// BASIS_BSPLINE means input vertices are treated as control points of a
	// B-spline surface. The number of apertures rendered is two fewer than the
	// number of apertures in that direction. In general, the generated surface
	// does not contain the control vertices specified.
	BASIS_BSPLINE = 1
	// BASIS_CATMULL_ROM means an interpolating basis defines the surface so
	// that the surface goes through all the input vertices specified. In
	// DirectX 8, this was BASIS_INTERPOLATE.
	BASIS_CATMULL_ROM = 2
)

// BLEND defines the supported blend mode.
type BLEND uint32

const (
	// BLEND_ZERO Blend factor is (0, 0, 0, 0).
	BLEND_ZERO = 1
	// BLEND_ONE Blend factor is (1, 1, 1, 1).
	BLEND_ONE = 2
	// BLEND_SRCCOLOR Blend factor is (Rs, Gs, Bs, As).
	BLEND_SRCCOLOR = 3
	// BLEND_INVSRCCOLOR Blend factor is (1 - Rs, 1 - Gs, 1 - Bs, 1 - As).
	BLEND_INVSRCCOLOR = 4
	// BLEND_SRCALPHA Blend factor is (As, As, As, As).
	BLEND_SRCALPHA = 5
	// BLEND_INVSRCALPHA Blend factor is ( 1 - As, 1 - As, 1 - As, 1 - As).
	BLEND_INVSRCALPHA = 6
	// BLEND_DESTALPHA Blend factor is (Ad, Ad, Ad, Ad).
	BLEND_DESTALPHA = 7
	// BLEND_INVDESTALPHA Blend factor is (1 - Ad, 1 - Ad, 1 - Ad, 1 - Ad).
	BLEND_INVDESTALPHA = 8
	// BLEND_DESTCOLOR Blend factor is (Rd, Gd, Bd, Ad).
	BLEND_DESTCOLOR = 9
	// BLEND_INVDESTCOLOR Blend factor is (1 - Rd, 1 - Gd, 1 - Bd, 1 - Ad).
	BLEND_INVDESTCOLOR = 10
	// BLEND_SRCALPHASAT Blend factor is (f, f, f, 1); where f = min(As, 1 -
	// Ad).
	BLEND_SRCALPHASAT = 11
	// BLEND_BOTHSRCALPHA is obsolete. Starting with DirectX 6, you can achieve
	// the same effect by setting the source and destination blend factors to
	// BLEND_SRCALPHA and BLEND_INVSRCALPHA in separate calls.
	BLEND_BOTHSRCALPHA = 12
	// BLEND_BOTHINVSRCALPHA is obsolete. Source blend factor is (1 - As, 1 -
	// As, 1 - As, 1 - As), and destination blend factor is (As, As, As, As);
	// the destination blend selection is overridden. This blend mode is
	// supported only for the RS_SRCBLEND render state.
	BLEND_BOTHINVSRCALPHA = 13
	// BLEND_BLENDFACTOR is a constant color blending factor used by the
	// frame-buffer blender. This blend mode is supported only if
	// PBLENDCAPS_BLENDFACTOR is set in the SrcBlendCaps or DestBlendCaps
	// members of CAPS.
	BLEND_BLENDFACTOR = 14
	// BLEND_INVBLENDFACTOR is an inverted constant color-blending factor used
	// by the frame-buffer blender. This blend mode is supported only if the
	// PBLENDCAPS_BLENDFACTOR bit is set in the SrcBlendCaps or DestBlendCaps
	// members of CAPS.
	BLEND_INVBLENDFACTOR = 15
)

// BLENDOP defines the supported blend operations.
type BLENDOP uint32

const (
	// BLENDOP_ADD means the result is the destination added to the source.
	// Result = Source + Destination
	BLENDOP_ADD = 1
	// BLENDOP_SUBTRACT means the result is the destination subtracted from to
	// the source. Result = Source - Destination
	BLENDOP_SUBTRACT = 2
	// BLENDOP_REVSUBTRACT means the result is the source subtracted from the
	// destination. Result = Destination - Source
	BLENDOP_REVSUBTRACT = 3
	// BLENDOP_MIN means the result is the minimum of the source and
	// destination. Result = MIN(Source, Destination)
	BLENDOP_MIN = 4
	// BLENDOP_MAX means the result is the maximum of the source and
	// destination. Result = MAX(Source, Destination)
	BLENDOP_MAX = 5
)

// CMPFUNC defines the supported compare functions.
type CMPFUNC uint32

const (
	// CMP_NEVER always fails the test.
	CMP_NEVER = 1
	// CMP_LESS accepts the new pixel if its value is less than the value of
	// the current pixel.
	CMP_LESS = 2
	// CMP_EQUAL accepts the new pixel if its value equals the value of the
	// current pixel.
	CMP_EQUAL = 3
	// CMP_LESSEQUAL accepts the new pixel if its value is less than or equal
	// to the value of the current pixel.
	CMP_LESSEQUAL = 4
	// CMP_GREATER accepts the new pixel if its value is greater than the value
	// of the current pixel.
	CMP_GREATER = 5
	// CMP_NOTEQUAL accepts the new pixel if its value does not equal the value
	// of the current pixel.
	CMP_NOTEQUAL = 6
	// CMP_GREATEREQUAL accepts the new pixel if its value is greater than or
	// equal to the value of the current pixel.
	CMP_GREATEREQUAL = 7
	// CMP_ALWAYS always passes the test.
	CMP_ALWAYS = 8
)

// CUBEMAP_FACES defines the faces of a cubemap.
type CUBEMAP_FACES uint32

const (
	// CUBEMAP_FACE_POSITIVE_X is the positive x-face of the cubemap.
	CUBEMAP_FACE_POSITIVE_X = 0
	// CUBEMAP_FACE_NEGATIVE_X is the negative x-face of the cubemap.
	CUBEMAP_FACE_NEGATIVE_X = 1
	// CUBEMAP_FACE_POSITIVE_Y is the positive y-face of the cubemap.
	CUBEMAP_FACE_POSITIVE_Y = 2
	// CUBEMAP_FACE_NEGATIVE_Y is the negative y-face of the cubemap.
	CUBEMAP_FACE_NEGATIVE_Y = 3
	// CUBEMAP_FACE_POSITIVE_Z is the positive z-face of the cubemap.
	CUBEMAP_FACE_POSITIVE_Z = 4
	// CUBEMAP_FACE_NEGATIVE_Z is the negative z-face of the cubemap.
	CUBEMAP_FACE_NEGATIVE_Z = 5
)

// CULL defines the supported culling modes.
type CULL uint32

const (
	// CULL_NONE does not cull back faces.
	CULL_NONE = 1
	// CULL_CW culls back faces with clockwise vertices.
	CULL_CW = 2
	// CULL_CCW culls back faces with counterclockwise vertices.
	CULL_CCW = 3
)

// DEBUGMONITORTOKENS defines the debug monitor tokens.
type DEBUGMONITORTOKENS uint32

const (
	// DMT_ENABLE enables the debug monitor.
	DMT_ENABLE = 0
	// DMT_DISABLE disables the debug monitor.
	DMT_DISABLE = 1
)

// DECLMETHOD defines the vertex declaration method which is a predefined
// operation performed by the tessellator (or any procedural geometry routine
// on the vertex data during tessellation).
type DECLMETHOD byte

const (
	// DECLMETHOD_DEFAULT is the default value. The tessellator copies the
	// vertex data (spline data for patches) as is, with no additional
	// calculations. When the tessellator is used, this element is
	// interpolated. Otherwise vertex data is copied into the input register.
	// The input and output type can be any value, but are always the same type.
	DECLMETHOD_DEFAULT = 0
	// DECLMETHOD_PARTIALU computes the tangent at a point on the rectangle or
	// triangle patch in the U direction. The input type can be one of the
	// following: DECLTYPE_D3DCOLOR, DECLTYPE_FLOAT3, DECLTYPE_FLOAT4,
	// DECLTYPE_SHORT4, DECLTYPE_UBYTE4.
	// The output type is always DECLTYPE_FLOAT3.
	DECLMETHOD_PARTIALU = 1
	// DECLMETHOD_PARTIALV Computes the tangent at a point on the rectangle or
	// triangle patch in the V direction. The input type can be one of the
	// following: DECLTYPE_D3DCOLOR, DECLTYPE_FLOAT3, DECLTYPE_FLOAT4,
	// DECLTYPE_SHORT4, DECLTYPE_UBYTE4.
	// The output type is always DECLTYPE_FLOAT3.
	DECLMETHOD_PARTIALV = 2
	// DECLMETHOD_CROSSUV computes the normal at a point on the rectangle or
	// triangle patch by taking the cross product of two tangents. The input
	// type can be one of the following: DECLTYPE_D3DCOLOR, DECLTYPE_FLOAT3,
	// DECLTYPE_FLOAT4, DECLTYPE_SHORT4, DECLTYPE_UBYTE4.
	// The output type is always DECLTYPE_FLOAT3.
	DECLMETHOD_CROSSUV = 3
	// DECLMETHOD_UV copies out the U, V values at a point on the rectangle or
	// triangle patch. This results in a 2D float. The input type must be set
	// to DECLTYPE_UNUSED. The output data type is always DECLTYPE_FLOAT2. The
	// input stream and offset are also unused (but must be set to 0).
	DECLMETHOD_UV = 4
	// DECLMETHOD_LOOKUP looks up a displacement map. The input type can be one
	// of the following: DECLTYPE_FLOAT2, DECLTYPE_FLOAT3, DECLTYPE_FLOAT4.
	// Only the X and Y components are used for the texture map lookup. The
	// output type is always DECLTYPE_FLOAT1. The device must support
	// displacement mapping. This constant is supported only by the
	// programmable pipeline on N-patch data, if N-patches are enabled.
	DECLMETHOD_LOOKUP = 5
	// DECLMETHOD_LOOKUPPRESAMPLED looks up a presampled displacement map. The
	// input type must be set to DECLTYPE_UNUSED; the stream index and the
	// stream offset must be set to 0. The output type for this operation is
	// always DECLTYPE_FLOAT1. The device must support displacement mapping.
	// This constant is supported only by the programmable pipeline on N-patch
	// data, if N-patches are enabled. This method can be used only with
	// DECLUSAGE_SAMPLE.
	DECLMETHOD_LOOKUPPRESAMPLED = 6
)

// DECLTYPE defines a vertex declaration data type.
type DECLTYPE byte

const (
	// DECLTYPE_FLOAT1 is a one-component float expanded to (float, 0, 0, 1).
	DECLTYPE_FLOAT1 = 0
	// DECLTYPE_FLOAT2 is a two-component float expanded to (float, float, 0,
	// 1).
	DECLTYPE_FLOAT2 = 1
	// DECLTYPE_FLOAT3 is a three-component float expanded to (float, float,
	// float, 1).
	DECLTYPE_FLOAT3 = 2
	// DECLTYPE_FLOAT4 is a four-component float expanded to (float, float,
	// float, float).
	DECLTYPE_FLOAT4 = 3
	// DECLTYPE_D3DCOLOR is a four-component, packed, unsigned bytes mapped to
	// 0 to 1 range. Input is a COLOR and is expanded to RGBA order.
	DECLTYPE_D3DCOLOR = 4
	// DECLTYPE_UBYTE4 is a four-component, unsigned byte.
	DECLTYPE_UBYTE4 = 5
	// DECLTYPE_SHORT2 is a two-component, signed short expanded to (value,
	// value, 0, 1).
	DECLTYPE_SHORT2 = 6
	// DECLTYPE_SHORT4 is a four-component, signed short expanded to (value,
	// value, value, value).
	DECLTYPE_SHORT4 = 7
	// DECLTYPE_UBYTE4N is a four-component byte with each byte normalized by
	// dividing with 255.0f.
	DECLTYPE_UBYTE4N = 8
	// DECLTYPE_SHORT2N is a normalized, two-component, signed short, expanded
	// to (first short/32767.0, second short/32767.0, 0, 1).
	DECLTYPE_SHORT2N = 9
	// DECLTYPE_SHORT4N is a normalized, four-component, signed short, expanded
	// to (first short/32767.0, second short/32767.0, third short/32767.0,
	// fourth short/32767.0).
	DECLTYPE_SHORT4N = 10
	// DECLTYPE_USHORT2N is a normalized, two-component, unsigned short,
	// expanded to (first short/65535.0, short short/65535.0, 0, 1).
	DECLTYPE_USHORT2N = 11
	// DECLTYPE_USHORT4N is a normalized, four-component, unsigned short,
	// expanded to (first short/65535.0, second short/65535.0, third
	// short/65535.0, fourth short/65535.0).
	DECLTYPE_USHORT4N = 12
	// DECLTYPE_UDEC3 is a three-component, unsigned, 10 10 10 format expanded
	// to (value, value, value, 1).
	DECLTYPE_UDEC3 = 13
	// DECLTYPE_DEC3N is a three-component, signed, 10 10 10 format normalized
	// and expanded to (v[0]/511.0, v[1]/511.0, v[2]/511.0, 1).
	DECLTYPE_DEC3N = 14
	// DECLTYPE_FLOAT16_2 is a two-component, 16-bit, floating point expanded
	// to (value, value, 0, 1).
	DECLTYPE_FLOAT16_2 = 15
	// DECLTYPE_FLOAT16_4 is a four-component, 16-bit, floating point expanded
	// to (value, value, value, value).
	DECLTYPE_FLOAT16_4 = 16
	// DECLTYPE_UNUSED means the type field in the declaration is unused. This
	// is designed for use with DECLMETHOD_UV and DECLMETHOD_LOOKUPPRESAMPLED.
	DECLTYPE_UNUSED = 17
)

// DECLUSAGE identifies the intended use of vertex data.
type DECLUSAGE byte

const (
	// DECLUSAGE_POSITION means position data ranging from (-1,-1) to (1,1).
	// Use DECLUSAGE_POSITION with a usage index of 0 to specify untransformed
	// position for fixed function vertex processing and the n-patch
	// tessellator. Use DECLUSAGE_POSITION with a usage index of 1 to specify
	// untransformed position in the fixed function vertex shader for vertex
	// tweening.
	DECLUSAGE_POSITION = 0
	// DECLUSAGE_BLENDWEIGHT means blending weight data. Use
	// DECLUSAGE_BLENDWEIGHT with a usage index of 0 to specify the blend
	// weights used in indexed and nonindexed vertex blending.
	DECLUSAGE_BLENDWEIGHT = 1
	// DECLUSAGE_BLENDINDICES means blending indices data. Use
	// DECLUSAGE_BLENDINDICES with a usage index of 0 to specify matrix indices
	// for indexed paletted skinning.
	DECLUSAGE_BLENDINDICES = 2
	// DECLUSAGE_NORMAL means vertex normal data. Use DECLUSAGE_NORMAL with a
	// usage index of 0 to specify vertex normals for fixed function vertex
	// processing and the n-patch tessellator. Use DECLUSAGE_NORMAL with a
	// usage index of 1 to specify vertex normals for fixed function vertex
	// processing for vertex tweening.
	DECLUSAGE_NORMAL = 3
	// DECLUSAGE_PSIZE means point size data. Use DECLUSAGE_PSIZE with a usage
	// index of 0 to specify the point-size attribute used by the setup engine
	// of the rasterizer to expand a point into a quad for the point-sprite
	// functionality.
	DECLUSAGE_PSIZE = 4
	// DECLUSAGE_TEXCOORD means texture coordinate data. Use
	// DECLUSAGE_TEXCOORD, n to specify texture coordinates in fixed function
	// vertex processing and in pixel shaders prior to ps_3_0. These can be
	// used to pass user defined data.
	DECLUSAGE_TEXCOORD = 5
	// DECLUSAGE_TANGENT means vertex tangent data.
	DECLUSAGE_TANGENT = 6
	// DECLUSAGE_BINORMAL means vertex binormal data.
	DECLUSAGE_BINORMAL = 7
	// DECLUSAGE_TESSFACTOR means a single positive floating point value. Use
	// DECLUSAGE_TESSFACTOR with a usage index of 0 to specify a tessellation
	// factor used in the tessellation unit to control the rate of tessellation.
	DECLUSAGE_TESSFACTOR = 8
	// DECLUSAGE_POSITIONT means the vertex data contains transformed position
	// data ranging from (0,0) to (viewport width, viewport height). Use
	// DECLUSAGE_POSITIONT with a usage index of 0 to specify transformed
	// position. When a declaration containing this is set, the pipeline does
	// not perform vertex processing.
	DECLUSAGE_POSITIONT = 9
	// DECLUSAGE_COLOR means the vertex data contains diffuse or specular
	// color. Use DECLUSAGE_COLOR with a usage index of 0 to specify the
	// diffuse color in the fixed function vertex shader and pixel shaders
	// prior to ps_3_0. Use D3DDECLUSAGE_COLOR with a usage index of 1 to
	// specify the specular color in the fixed function vertex shader and pixel
	// shaders prior to ps_3_0.
	DECLUSAGE_COLOR = 10
	// DECLUSAGE_FOG means the vertex data contains fog data. Use DECLUSAGE_FOG
	// with a usage index of 0 to specify a fog blend value used after pixel
	// shading finishes. This applies to pixel shaders prior to version ps_3_0.
	DECLUSAGE_FOG = 11
	// DECLUSAGE_DEPTH means the vertex data contains depth data.
	DECLUSAGE_DEPTH = 12
	// DECLUSAGE_SAMPLE means the vertex data contains sampler data. Use
	// DECLUSAGE_SAMPLE with a usage index of 0 to specify the displacement
	// value to look up. It can be used only with
	// DECLUSAGE_LOOKUPPRESAMPLED or DECLUSAGE_LOOKUP.
	DECLUSAGE_SAMPLE = 13
)

// DEGREETYPE defines the degree of the variables in the equation that
// describes a curve.
type DEGREETYPE uint32

const (
	// DEGREE_LINEAR means the curve is described by variables of first order.
	DEGREE_LINEAR = 1
	// DEGREE_QUADRATIC means the curve is described by variables of second
	// order.
	DEGREE_QUADRATIC = 2
	// DEGREE_CUBIC means the curve is described by variables of third order.
	DEGREE_CUBIC = 3
	// DEGREE_QUINTIC means the curve is described by variables of fourth order.
	DEGREE_QUINTIC = 5
)

// DEVTYPE defines device types.
type DEVTYPE uint32

const (
	// DEVTYPE_HAL indicates hardware rasterization. Shading is done with
	// software, hardware, or mixed transform and lighting.
	DEVTYPE_HAL = 1
	// DEVTYPE_REF means Direct3D features are implemented in software;
	// however, the reference rasterizer does make use of special CPU
	// instructions whenever it can.
	// The reference device is installed by the Windows SDK 8.0 or later and is
	// intended as an aid in debugging for development only.
	DEVTYPE_REF = 2
	// DEVTYPE_SW is a pluggable software device that has been registered with
	// Direct3D.RegisterSoftwareDevice.
	DEVTYPE_SW = 3
	// DEVTYPE_NULLREF initializes Direct3D on a computer that has neither
	// hardware nor reference rasterization available, and enable resources for
	// 3D content creation.
	DEVTYPE_NULLREF = 4
)

// FILLMODE defines constants describing the fill mode.
type FILLMODE uint32

const (
	// FILL_POINT fills points.
	FILL_POINT = 1
	// FILL_WIREFRAME fills wireframes.
	FILL_WIREFRAME = 2
	// FILL_SOLID fills solids.
	FILL_SOLID = 3
)

// FOGMODE defines constants that describe the fog mode.
type FOGMODE uint32

const (
	// FOG_NONE means no fog effect.
	FOG_NONE = 0
	// FOG_EXP means the fog effect intensifies exponentially, according to the
	// following formula:
	// f = 1/(e^(d*density))
	FOG_EXP = 1
	// FOG_EXP2 means the fog effect intensifies exponentially with the square
	// of the distance, according to the following formula:
	// f = 1/(e^[(d*density)^2])
	FOG_EXP2 = 2
	// FOG_LINEAR means the fog effect intensifies linearly between the start
	// and end points, according to the following formula:
	// f = (end - d)/(end - start)
	// This is the only fog mode currently supported.
	FOG_LINEAR = 3
)

// FORMAT defines the various types of surface formats.
type FORMAT uint32

func (f FORMAT) String() string {
	switch f {
	default:
		return "UNKNOWN"
	case FMT_R8G8B8:
		return "R8G8B8"
	case FMT_A8R8G8B8:
		return "A8R8G8B8"
	case FMT_X8R8G8B8:
		return "X8R8G8B8"
	case FMT_R5G6B5:
		return "R5G6B5"
	case FMT_X1R5G5B5:
		return "X1R5G5B5"
	case FMT_A1R5G5B5:
		return "A1R5G5B5"
	case FMT_A4R4G4B4:
		return "A4R4G4B4"
	case FMT_R3G3B2:
		return "R3G3B2"
	case FMT_A8:
		return "A8"
	case FMT_A8R3G3B2:
		return "A8R3G3B2"
	case FMT_X4R4G4B4:
		return "X4R4G4B4"
	case FMT_A2B10G10R10:
		return "A2B10G10R10"
	case FMT_A8B8G8R8:
		return "A8B8G8R8"
	case FMT_X8B8G8R8:
		return "X8B8G8R8"
	case FMT_G16R16:
		return "G16R16"
	case FMT_A2R10G10B10:
		return "A2R10G10B10"
	case FMT_A16B16G16R16:
		return "A16B16G16R16"
	case FMT_A8P8:
		return "A8P8"
	case FMT_P8:
		return "P8"
	case FMT_L8:
		return "L8"
	case FMT_A8L8:
		return "A8L8"
	case FMT_A4L4:
		return "A4L4"
	case FMT_V8U8:
		return "V8U8"
	case FMT_L6V5U5:
		return "L6V5U5"
	case FMT_X8L8V8U8:
		return "X8L8V8U8"
	case FMT_Q8W8V8U8:
		return "Q8W8V8U8"
	case FMT_V16U16:
		return "V16U16"
	case FMT_A2W10V10U10:
		return "A2W10V10U10"
	case FMT_UYVY:
		return "UYVY"
	case FMT_R8G8_B8G8:
		return "R8G8_B8G8"
	case FMT_YUY2:
		return "YUY2"
	case FMT_G8R8_G8B8:
		return "G8R8_G8B8"
	case FMT_DXT1:
		return "DXT1"
	case FMT_DXT2:
		return "DXT2"
	case FMT_DXT3:
		return "DXT3"
	case FMT_DXT4:
		return "DXT4"
	case FMT_DXT5:
		return "DXT5"
	case FMT_D16_LOCKABLE:
		return "D16_LOCKABLE"
	case FMT_D32:
		return "D32"
	case FMT_D15S1:
		return "D15S1"
	case FMT_D24S8:
		return "D24S8"
	case FMT_D24X8:
		return "D24X8"
	case FMT_D24X4S4:
		return "D24X4S4"
	case FMT_D16:
		return "D16"
	case FMT_L16:
		return "L16"
	case FMT_D32F_LOCKABLE:
		return "D32F_LOCKABLE"
	case FMT_D24FS8:
		return "D24FS8"
	case FMT_VERTEXDATA:
		return "VERTEXDATA"
	case FMT_INDEX16:
		return "INDEX16"
	case FMT_INDEX32:
		return "INDEX32"
	case FMT_Q16W16V16U16:
		return "Q16W16V16U16"
	case FMT_MULTI2_ARGB8:
		return "MULTI2_ARGB8"
	case FMT_R16F:
		return "R16F"
	case FMT_G16R16F:
		return "G16R16F"
	case FMT_A16B16G16R16F:
		return "A16B16G16R16F"
	case FMT_R32F:
		return "R32F"
	case FMT_G32R32F:
		return "G32R32F"
	case FMT_A32B32G32R32F:
		return "A32B32G32R32F"
	case FMT_CxV8U8:
		return "CxV8U8"
	}
}

const (
	// FMT_UNKNOWN means the surface format is unknown.
	FMT_UNKNOWN = 0
	// FMT_R8G8B8 means 24-bit RGB pixel format with 8 bits per channel.
	FMT_R8G8B8 = 20
	// FMT_A8R8G8B8 means 32-bit ARGB pixel format with alpha, using 8 bits per
	// channel.
	FMT_A8R8G8B8 = 21
	// FMT_X8R8G8B8 means 32-bit RGB pixel format, where 8 bits are reserved
	// for each color.
	FMT_X8R8G8B8 = 22
	// FMT_R5G6B5 means 16-bit RGB pixel format with 5 bits for red, 6 bits for
	// green, and 5 bits for blue.
	FMT_R5G6B5 = 23
	// FMT_X1R5G5B5 means 16-bit pixel format where 5 bits are reserved for
	// each color.
	FMT_X1R5G5B5 = 24
	// FMT_A1R5G5B5 means 16-bit pixel format where 5 bits are reserved for
	// each color and 1 bit is reserved for alpha.
	FMT_A1R5G5B5 = 25
	// FMT_A4R4G4B4 means 16-bit ARGB pixel format with 4 bits for each channel.
	FMT_A4R4G4B4 = 26
	// FMT_R3G3B2 means 8-bit RGB texture format using 3 bits for red, 3 bits
	// for green, and 2 bits for blue.
	FMT_R3G3B2 = 27
	// FMT_A8 means 8-bit alpha only.
	FMT_A8 = 28
	// FMT_A8R3G3B2 means 16-bit ARGB texture format using 8 bits for alpha, 3
	// bits each for red and green, and 2 bits for blue.
	FMT_A8R3G3B2 = 29
	// FMT_X4R4G4B4 means 16-bit RGB pixel format using 4 bits for each color.
	FMT_X4R4G4B4 = 30
	// FMT_A2B10G10R10 means 32-bit pixel format using 10 bits for each color
	// and 2 bits for alpha.
	FMT_A2B10G10R10 = 31
	// FMT_A8B8G8R8 means 32-bit ARGB pixel format with alpha, using 8 bits per
	// channel.
	FMT_A8B8G8R8 = 32
	// FMT_X8B8G8R8 means 32-bit RGB pixel format, where 8 bits are reserved
	// for each color.
	FMT_X8B8G8R8 = 33
	// FMT_G16R16 means 32-bit pixel format using 16 bits each for green and
	// red.
	FMT_G16R16 = 34
	// FMT_A2R10G10B10 means 32-bit pixel format using 10 bits each for red,
	// green, and blue, and 2 bits for alpha.
	FMT_A2R10G10B10 = 35
	// FMT_A16B16G16R16 means 64-bit pixel format using 16 bits for each
	// component.
	FMT_A16B16G16R16 = 36
	// FMT_A8P8 means 8-bit color indexed with 8 bits of alpha.
	FMT_A8P8 = 40
	// FMT_P8 means 8-bit color indexed.
	FMT_P8 = 41
	// FMT_L8 means 8-bit luminance only.
	FMT_L8 = 50
	// FMT_A8L8 means 16-bit using 8 bits each for alpha and luminance.
	FMT_A8L8 = 51
	// FMT_A4L4 means 8-bit using 4 bits each for alpha and luminance.
	FMT_A4L4 = 52
	// FMT_V8U8 means 16-bit bump-map format using 8 bits each for u and v data.
	FMT_V8U8 = 60
	// FMT_L6V5U5 means 16-bit bump-map format with luminance using 6 bits for
	// luminance, and 5 bits each for v and u.
	FMT_L6V5U5 = 61
	// FMT_X8L8V8U8 means 32-bit bump-map format with luminance using 8 bits
	// for each channel.
	FMT_X8L8V8U8 = 62
	// FMT_Q8W8V8U8 means 32-bit bump-map format using 8 bits for each channel.
	FMT_Q8W8V8U8 = 63
	// FMT_V16U16 means 32-bit bump-map format using 16 bits for each channel.
	FMT_V16U16 = 64
	// FMT_A2W10V10U10 means 32-bit bump-map format using 2 bits for alpha and
	// 10 bits each for w, v, and u.
	FMT_A2W10V10U10 = 67
	// FMT_UYVY means UYVY format (PC98 compliance).
	FMT_UYVY = 1498831189
	// FMT_R8G8_B8G8 is a 16-bit packed RGB format analogous to UYVY (U0Y0,
	// V0Y1, U2Y2, and so on). It requires a pixel pair in order to properly
	// represent the color value. The first pixel in the pair contains 8 bits
	// of green (in the low 8 bits) and 8 bits of red (in the high 8 bits). The
	// second pixel contains 8 bits of green (in the low 8 bits) and 8 bits of
	// blue (in the high 8 bits). Together, the two pixels share the red and
	// blue components, while each has a unique green component (R0G0, B0G1,
	// R2G2, and so on). The texture sampler does not normalize the colors when
	// looking up into a pixel shader; they remain in the range of 0.0f to
	// 255.0f. This is true for all programmable pixel shader models. For the
	// fixed function pixel shader, the hardware should normalize to the 0.f to
	// 1.f range and essentially treat it as the YUY2 texture. Hardware that
	// exposes this format must have PixelShader1xMaxValue member of CAPS set
	// to a value capable of handling that range.
	FMT_R8G8_B8G8 = 1195525970
	// FMT_YUY2 means YUY2 format (PC98 compliance).
	FMT_YUY2 = 844715353
	// FMT_G8R8_G8B8 is a 16-bit packed RGB format analogous to YUY2 (Y0U0,
	// Y1V0, Y2U2, and so on). It requires a pixel pair in order to properly
	// represent the color value. The first pixel in the pair contains 8 bits
	// of green (in the high 8 bits) and 8 bits of red (in the low 8 bits). The
	// second pixel contains 8 bits of green (in the high 8 bits) and 8 bits of
	// blue (in the low 8 bits). Together, the two pixels share the red and
	// blue components, while each has a unique green component (G0R0, G1B0,
	// G2R2, and so on). The texture sampler does not normalize the colors when
	// looking up into a pixel shader; they remain in the range of 0.0f to
	// 255.0f. This is true for all programmable pixel shader models. For the
	// fixed function pixel shader, the hardware should normalize to the 0.f to
	// 1.f range and essentially treat it as the YUY2 texture. Hardware that
	// exposes this format must have the PixelShader1xMaxValue member of CAPS
	// set to a value capable of handling that range.
	FMT_G8R8_G8B8 = 1111970375
	// FMT_DXT1 means DXT1 compression texture format.
	FMT_DXT1 = 827611204
	// FMT_DXT2 means DXT2 compression texture format.
	FMT_DXT2 = 844388420
	// FMT_DXT3 means DXT3 compression texture format.
	FMT_DXT3 = 861165636
	// FMT_DXT4 means DXT4 compression texture format.
	FMT_DXT4 = 877942852
	// FMT_DXT5 means DXT5 compression texture format.
	FMT_DXT5 = 894720068
	// FMT_D16_LOCKABLE means 16-bit z-buffer bit depth.
	FMT_D16_LOCKABLE = 70
	// FMT_D32 means 32-bit z-buffer bit depth.
	FMT_D32 = 71
	// FMT_D15S1 means 16-bit z-buffer bit depth where 15 bits are reserved for
	// the depth channel and 1 bit is reserved for the stencil channel.
	FMT_D15S1 = 73
	// FMT_D24S8 means 32-bit z-buffer bit depth using 24 bits for the depth
	// channel and 8 bits for the stencil channel.
	FMT_D24S8 = 75
	// FMT_D24X8 means 32-bit z-buffer bit depth using 24 bits for the depth
	// channel.
	FMT_D24X8 = 77
	// FMT_D24X4S4 means 32-bit z-buffer bit depth using 24 bits for the depth
	// channel and 4 bits for the stencil channel.
	FMT_D24X4S4 = 79
	// FMT_D16 means 16-bit z-buffer bit depth.
	FMT_D16 = 80
	// FMT_L16 means 16-bit luminance only.
	FMT_L16 = 81
	// FMT_D32F_LOCKABLE is a lockable format where the depth value is
	// represented as a standard IEEE floating-point number.
	FMT_D32F_LOCKABLE = 82
	// FMT_D24FS8 is a non-lockable format that contains 24 bits of depth (in a
	// 24-bit floating point format - 20e4) and 8 bits of stencil.
	FMT_D24FS8 = 83
	// FMT_VERTEXDATA describes a vertex buffer surface.
	FMT_VERTEXDATA = 100
	// FMT_INDEX16 means 16-bit index buffer bit depth.
	FMT_INDEX16 = 101
	// FMT_INDEX32 means 32-bit index buffer bit depth.
	FMT_INDEX32 = 102
	// FMT_Q16W16V16U16 means 64-bit bump-map format using 16 bits for each
	// component.
	FMT_Q16W16V16U16 = 110
	// FMT_MULTI2_ARGB8 means a multiElement texture (not compressed).
	FMT_MULTI2_ARGB8 = 827606349
	// FMT_R16F means 16-bit float format using 16 bits for the red channel.
	FMT_R16F = 111
	// FMT_G16R16F means 32-bit float format using 16 bits for the red channel
	// and 16 bits for the green channel.
	FMT_G16R16F = 112
	// FMT_A16B16G16R16F means 64-bit float format using 16 bits for the each
	// channel (alpha, blue, green, red).
	FMT_A16B16G16R16F = 113
	// FMT_R32F means 32-bit float format using 32 bits for the red channel.
	FMT_R32F = 114
	// FMT_G32R32F means 64-bit float format using 32 bits for the red channel
	// and 32 bits for the green channel.
	FMT_G32R32F = 115
	// FMT_A32B32G32R32F means 128-bit float format using 32 bits for the each
	// channel (alpha, blue, green, red).
	FMT_A32B32G32R32F = 116
	// FMT_CxV8U8 means 16-bit normal compression format. The texture sampler
	// computes the C channel from: C = sqrt(1 - U² - V²).
	FMT_CxV8U8 = 117
)

// LIGHTTYPE defines the light type.
type LIGHTTYPE uint32

const (
	// LIGHT_POINT means the light is a point source. The light has a position
	// in space and radiates light in all directions.
	LIGHT_POINT = 1
	// LIGHT_SPOT means the light is a spotlight source. This light is like a
	// point light, except that the illumination is limited to a cone. This
	// light type has a direction and several other parameters that determine
	// the shape of the cone it produces.
	LIGHT_SPOT = 2
	// LIGHT_DIRECTIONAL means the light is a directional light source. This is
	// equivalent to using a point light source at an infinite distance.
	LIGHT_DIRECTIONAL = 3
)

// MATERIALCOLORSOURCE defines the location at which a color or color component
// must be accessed for lighting calculations.
type MATERIALCOLORSOURCE uint32

const (
	// MCS_MATERIAL uses the color from the current material.
	MCS_MATERIAL = 0
	// MCS_COLOR1 uses the diffuse vertex color.
	MCS_COLOR1 = 1
	// MCS_COLOR2 uses the specular vertex color.
	MCS_COLOR2 = 2
)

// MULTISAMPLE_TYPE defines the levels of full-scene multisampling that the
// device can apply.
type MULTISAMPLE_TYPE uint32

const (
	// MULTISAMPLE_NONE means no level of full-scene multisampling is available.
	MULTISAMPLE_NONE = 0
	// MULTISAMPLE_NONMASKABLE enables the multisample quality value.
	MULTISAMPLE_NONMASKABLE = 1
	// MULTISAMPLE_2_SAMPLES means 2 levels of full-scene multisampling
	// available.
	MULTISAMPLE_2_SAMPLES = 2
	// MULTISAMPLE_3_SAMPLES means 3 levels of full-scene multisampling
	// available.
	MULTISAMPLE_3_SAMPLES = 3
	// MULTISAMPLE_4_SAMPLES means 4 levels of full-scene multisampling
	// available.
	MULTISAMPLE_4_SAMPLES = 4
	// MULTISAMPLE_5_SAMPLES means 5 levels of full-scene multisampling
	// available.
	MULTISAMPLE_5_SAMPLES = 5
	// MULTISAMPLE_6_SAMPLES means 6 levels of full-scene multisampling
	// available.
	MULTISAMPLE_6_SAMPLES = 6
	// MULTISAMPLE_7_SAMPLES means 7 levels of full-scene multisampling
	// available.
	MULTISAMPLE_7_SAMPLES = 7
	// MULTISAMPLE_8_SAMPLES means 8 levels of full-scene multisampling
	// available.
	MULTISAMPLE_8_SAMPLES = 8
	// MULTISAMPLE_9_SAMPLES means 9 levels of full-scene multisampling
	// available.
	MULTISAMPLE_9_SAMPLES = 9
	// MULTISAMPLE_10_SAMPLES means 10 levels of full-scene multisampling
	// available.
	MULTISAMPLE_10_SAMPLES = 10
	// MULTISAMPLE_11_SAMPLES means 11 levels of full-scene multisampling
	// available.
	MULTISAMPLE_11_SAMPLES = 11
	// MULTISAMPLE_12_SAMPLES means 12 levels of full-scene multisampling
	// available.
	MULTISAMPLE_12_SAMPLES = 12
	// MULTISAMPLE_13_SAMPLES means 13 levels of full-scene multisampling
	// available.
	MULTISAMPLE_13_SAMPLES = 13
	// MULTISAMPLE_14_SAMPLES means 14 levels of full-scene multisampling
	// available.
	MULTISAMPLE_14_SAMPLES = 14
	// MULTISAMPLE_15_SAMPLES means 15 levels of full-scene multisampling
	// available.
	MULTISAMPLE_15_SAMPLES = 15
	// MULTISAMPLE_16_SAMPLES means 16 levels of full-scene multisampling
	// available.
	MULTISAMPLE_16_SAMPLES = 16
)

// PATCHEDGESTYLE defines whether the current tessellation mode is discrete or
// continuous.
type PATCHEDGESTYLE uint32

const (
	// PATCHEDGE_DISCRETE means discrete edge style. In discrete mode, you can
	// specify float tessellation but it will be truncated to integers.
	PATCHEDGE_DISCRETE = 0
	// PATCHEDGE_CONTINUOUS means continuous edge style. In continuous mode,
	// tessellation is specified as float values that can be smoothly varied to
	// reduce "popping" artifacts.
	PATCHEDGE_CONTINUOUS = 1
)

// POOL defines the memory class that holds the buffers for a resource.
type POOL uint32

const (
	// POOL_DEFAULT means resources are placed in the memory pool most
	// appropriate for the set of usages requested for the given resource. This
	// is usually video memory, including both local video memory and AGP
	// memory. The POOL_DEFAULT pool is separate from POOL_MANAGED and
	// POOL_SYSTEMMEM, and it specifies that the resource is placed in the
	// preferred memory for device access. Note that POOL_DEFAULT never
	// indicates that either POOL_MANAGED or POOL_SYSTEMMEM should be chosen as
	// the memory pool type for this resource. Textures placed in the
	// POOL_DEFAULT pool cannot be locked unless they are dynamic textures or
	// they are private, FOURCC, driver formats. To access unlockable textures,
	// you must use functions such as Device.UpdateSurface,
	// Device.UpdateTexture, Device.GetFrontBufferData, and
	// Device.GetRenderTargetData. POOL_MANAGED is probably a better choice
	// than POOL_DEFAULT for most applications. Note that some textures created
	// in driver-proprietary pixel formats, unknown to the Direct3D runtime,
	// can be locked. Also note that - unlike textures - swap chain back
	// buffers, render targets, vertex buffers, and index buffers can be
	// locked. When a device is lost, resources created using POOL_DEFAULT must
	// be released before calling Device.Reset.
	// When creating resources with POOL_DEFAULT, if video card memory is
	// already committed, managed resources will be evicted to free enough
	// memory to satisfy the request.
	POOL_DEFAULT = 0
	// POOL_MANAGED means resources are copied automatically to
	// device-accessible memory as needed. Managed resources are backed by
	// system memory and do not need to be recreated when a device is lost.
	// Managed resources can be locked. Only the system-memory copy is directly
	// modified. Direct3D copies your changes to driver-accessible memory as
	// needed.
	// Differences between Direct3D 9 and Direct3D 9Ex: POOL_MANAGED is valid
	// with Device; however, it is not valid with DeviceEx.
	POOL_MANAGED = 1
	// POOL_SYSTEMMEM means resources are placed in memory that is not
	// typically accessible by the Direct3D device. This memory allocation
	// consumes system RAM but does not reduce pageable RAM. These resources do
	// not need to be recreated when a device is lost. Resources in this pool
	// can be locked and can be used as the source for a Device.UpdateSurface
	// or Device.UpdateTexture operation to a memory resource created with
	// POOL_DEFAULT.
	POOL_SYSTEMMEM = 2
	// POOL_SCRATCH means resources are placed in system RAM and do not need to
	// be recreated when a device is lost. These resources are not bound by
	// device size or format restrictions. Because of this, these resources
	// cannot be accessed by the Direct3D device nor set as textures or render
	// targets. However, these resources can always be created, locked, and
	// copied.
	POOL_SCRATCH = 3
)

// PRIMITIVETYPE defines the primitives supported by Direct3D.
type PRIMITIVETYPE uint32

const (
	// PT_POINTLIST renders the vertices as a collection of isolated points.
	// This value is unsupported for indexed primitives.
	PT_POINTLIST = 1
	// PT_LINELIST renders the vertices as a list of isolated straight line
	// segments.
	PT_LINELIST = 2
	// PT_LINESTRIP renders the vertices as a single polyline.
	PT_LINESTRIP = 3
	// PT_TRIANGLELIST renders the specified vertices as a sequence of isolated
	// triangles. Each group of three vertices defines a separate triangle.
	// Back-face culling is affected by the current winding-order render state.
	PT_TRIANGLELIST = 4
	// PT_TRIANGLESTRIP renders the vertices as a triangle strip. The
	// backface-culling flag is automatically flipped on even-numbered
	// triangles.
	PT_TRIANGLESTRIP = 5
	// PT_TRIANGLEFAN renders the vertices as a triangle fan.
	PT_TRIANGLEFAN = 6
)

// QUERYTYPE identifies the query type.
type QUERYTYPE byte

const (
	// QUERYTYPE_VCACHE queries for driver hints about data layout for vertex
	// caching.
	QUERYTYPE_VCACHE = 4
	// QUERYTYPE_RESOURCEMANAGER queryies the resource manager. For this query,
	// the device behavior flags must include CREATE_DISABLE_DRIVER_MANAGEMENT.
	QUERYTYPE_RESOURCEMANAGER = 5
	// QUERYTYPE_VERTEXSTATS queries vertex statistics.
	QUERYTYPE_VERTEXSTATS = 6
	// QUERYTYPE_EVENT queries for any and all asynchronous events that have
	// been issued from API calls.
	QUERYTYPE_EVENT = 8
	// QUERYTYPE_OCCLUSION returns the number of pixels that pass z-testing.
	// These pixels are for primitives drawn between the issue of ISSUE_BEGIN
	// and ISSUE_END. This enables an application to check the occlusion result
	// against 0. Zero is fully occluded, which means the pixels are not
	// visible from the current camera position.
	QUERYTYPE_OCCLUSION = 9
	// QUERYTYPE_TIMESTAMP returns a 64-bit timestamp.
	QUERYTYPE_TIMESTAMP = 10
	// QUERYTYPE_TIMESTAMPDISJOINT notifies an application if the counter
	// frequency has changed from the QUERYTYPE_TIMESTAMP.
	QUERYTYPE_TIMESTAMPDISJOINT = 11
	// QUERYTYPE_TIMESTAMPFREQ results in true if the values from
	// QUERYTYPE_TIMESTAMP queries cannot be guaranteed to be continuous
	// throughout the duration of the QUERYTYPE_TIMESTAMPDISJOINT query.
	// Otherwise, the query result is false.
	QUERYTYPE_TIMESTAMPFREQ = 12
	// QUERYTYPE_PIPELINETIMINGS returns percent of time processing pipeline
	// data.
	QUERYTYPE_PIPELINETIMINGS = 13
	// QUERYTYPE_INTERFACETIMINGS returns percent of time processing data in
	// the driver.
	QUERYTYPE_INTERFACETIMINGS = 14
	// QUERYTYPE_VERTEXTIMINGS returns percent of time processing vertex shader
	// data.
	QUERYTYPE_VERTEXTIMINGS = 15
	// QUERYTYPE_PIXELTIMINGS returns percent of time processing pixel shader
	// data.
	QUERYTYPE_PIXELTIMINGS = 16
	// QUERYTYPE_BANDWIDTHTIMINGS is for throughput measurement comparisons for
	// help in understanding the performance of an application.
	QUERYTYPE_BANDWIDTHTIMINGS = 17
	// QUERYTYPE_CACHEUTILIZATION measures the cache hit-rate performance for
	// textures and indexed vertices.
	QUERYTYPE_CACHEUTILIZATION = 18
)

// RENDERSTATETYPE defines set-up states for all kinds of vertex and pixel
// processing. Some render states set up vertex processing, and some set up
// pixel processing. Render states can be saved and restored using stateblocks.
type RENDERSTATETYPE uint32

const (
	// RS_ZENABLE is the depth-buffering state as one member of the ZBUFFERTYPE
	// enumerated type. Set this state to ZB_TRUE to enable z-buffering,
	// ZB_USEW to enable w-buffering, or ZB_FALSE to disable depth buffering.
	// The default value for this render state is ZB_TRUE if a depth stencil
	// was created along with the swap chain by setting the
	// EnableAutoDepthStencil member of the PRESENT_PARAMETERS structure to
	// true, and ZB_FALSE otherwise.
	RS_ZENABLE = 7
	// RS_FILLMODE expects one or more members of the D3DFILLMODE enumerated
	// type. The default value is D3DFILL_SOLID.
	RS_FILLMODE = 8
	// RS_SHADEMODE expects one or more members of the D3DSHADEMODE enumerated
	// type. The default value is D3DSHADE_GOURAUD.
	RS_SHADEMODE = 9
	// RS_ZWRITEENABLE can be set to true to enable the application to write to
	// the depth buffer. The default value is true. This member enables an
	// application to prevent the system from updating the depth buffer with
	// new depth values. If false, depth comparisons are still made according
	// to the render state RS_ZFUNC, assuming that depth buffering is taking
	// place, but depth values are not written to the buffer.
	RS_ZWRITEENABLE = 14
	// RS_ALPHATESTENABLE can be set to true to enable per pixel alpha testing.
	// If the test passes, the pixel is processed by the frame buffer.
	// Otherwise, all frame-buffer processing is skipped for the pixel.
	// The test is done by comparing the incoming alpha value with the
	// reference alpha value, using the comparison function provided by the
	// RS_ALPHAFUNC render state. The reference alpha value is determined by
	// the value set for RS_ALPHAREF.
	// The default value of this parameter is false.
	RS_ALPHATESTENABLE = 15
	// RS_LASTPIXEL is true by default, which enables drawing of the last pixel
	// in a line. To prevent drawing of the last pixel, set this value to false.
	RS_LASTPIXEL = 16
	// RS_SRCBLEND expectes one member of the BLEND enumerated type. The
	// default value is BLEND_ONE.
	RS_SRCBLEND = 19
	// RS_DESTBLEND expectes one member of the BLEND enumerated type. The
	// default value is BLEND_ZERO.
	RS_DESTBLEND = 20
	// RS_CULLMODE specifies how back-facing triangles are culled, if at all.
	// This can be set to one member of the CULL enumerated type. The default
	// value is CULL_CCW.
	RS_CULLMODE = 22
	// RS_ZFUNC expectes one member of the CMPFUNC enumerated type. The default
	// value is CMP_LESSEQUAL. This member enables an application to accept or
	// reject a pixel, based on its distance from the camera.
	// The depth value of the pixel is compared with the depth-buffer value. If
	// the depth value of the pixel passes the comparison function, the pixel
	// is written.
	// The depth value is written to the depth buffer only if the render state
	// is true.
	// Software rasterizers and many hardware accelerators work faster if the
	// depth test fails, because there is no need to filter and modulate the
	// texture if the pixel is not going to be rendered.
	RS_ZFUNC = 23
	// RS_ALPHAREF expectes a value that specifies a reference alpha value
	// against which pixels are tested when alpha testing is enabled. This is
	// an 8-bit value placed in the low 8 bits of the render-state value.
	// Values can range from 0x00000000 through 0x000000FF. The default value
	// is 0.
	RS_ALPHAREF = 24
	// RS_ALPHAFUNC expectes one member of the CMPFUNC enumerated type. The
	// default value is CMP_ALWAYS. This member enables an application to
	// accept or reject a pixel, based on its alpha value.
	RS_ALPHAFUNC = 25
	// RS_DITHERENABLE can be set to true to enable dithering. The default
	// value is false.
	RS_DITHERENABLE = 26
	// RS_ALPHABLENDENABLE can be set to true to enable alpha-blended
	// transparency. The default value is false.
	// The type of alpha blending is determined by the RS_SRCBLEND and
	// RS_DESTBLEND render states.
	RS_ALPHABLENDENABLE = 27
	// RS_FOGENABLE can be set to true to enable fog blending. The default
	// value is false.
	RS_FOGENABLE = 28
	// RS_SPECULARENABLE can be set to true to enable specular highlights. The
	// default value is false.
	// Specular highlights are calculated as though every vertex in the object
	// being lit is at the object's origin. This gives the expected results as
	// long as the object is modeled around the origin and the distance from
	// the light to the object is relatively large. In other cases, the results
	// as undefined.
	// When this member is set to true, the specular color is added to the base
	// color after the texture cascade but before alpha blending.
	RS_SPECULARENABLE = 29
	// RS_FOGCOLOR expectes a value whose type is COLOR. The default value is 0.
	RS_FOGCOLOR = 34
	// RS_FOGTABLEMODE is the fog formula to be used for pixel fog. Set to one
	// of the members of the FOGMODE enumerated type. The default value is
	// FOG_NONE.
	RS_FOGTABLEMODE = 35
	// RS_FOGSTART is the depth at which pixel or vertex fog effects begin for
	// linear fog mode. The default value is 0.0f. Depth is specified in world
	// space for vertex fog and either device space [0.0, 1.0] or world space
	// for pixel fog. For pixel fog, these values are in device space when the
	// system uses z for fog calculations and world-world space when the system
	// is using eye-relative fog (w-fog).
	// Use Device.SetRenderStateFloat to set this value.
	RS_FOGSTART = 36
	// RS_FOGEND is the depth at which pixel or vertex fog effects end for
	// linear fog mode. The default value is 1.0f. Depth is specified in world
	// space for vertex fog and either device space [0.0, 1.0] or world space
	// for pixel fog. For pixel fog, these values are in device space when the
	// system uses z for fog calculations and in world space when the system is
	// using eye-relative fog (w-fog).
	// Use Device.SetRenderStateFloat to set this value.
	RS_FOGEND = 37
	// RS_FOGDENSITY is the fog density for pixel or vertex fog used in the
	// exponential fog modes FOG_EXP and FOG_EXP2). Valid density values range
	// from 0.0 through 1.0. The default value is 1.0.
	// Use Device.SetRenderStateFloat to set this value.
	RS_FOGDENSITY = 38
	// RS_RANGEFOGENABLE can be set to true to enable range-based vertex fog.
	// The default value is false, in which case the system uses depth-based
	// fog. In range-based fog, the distance of an object from the viewer is
	// used to compute fog effects, not the depth of the object (that is, the
	// z-coordinate) in the scene. In range-based fog, all fog methods work as
	// usual, except that they use range instead of depth in the computations.
	// Range is the correct factor to use for fog computations, but depth is
	// commonly used instead because range is time-consuming to compute and
	// depth is generally already available. Using depth to calculate fog has
	// the undesirable effect of having the fogginess of peripheral objects
	// change as the viewer's eye moves - in this case, the depth changes and
	// the range remains constant.
	// Because no hardware currently supports per-pixel range-based fog, range
	// correction is offered only for vertex fog.
	// Use Device.SetRenderStateBool to set this value.
	RS_RANGEFOGENABLE = 48
	// RS_STENCILENABLE can be set to true to enable stenciling, or false to
	// disable stenciling. The default value is false.
	// Use Device.SetRenderStateBool to set this value.
	RS_STENCILENABLE = 52
	// RS_STENCILFAIL is the stencil operation to perform if the stencil test
	// fails. Values are from the STENCILOP enumerated type. The default value
	// is STENCILOP_KEEP.
	RS_STENCILFAIL = 53
	// RS_STENCILZFAIL is the stencil operation to perform if the stencil test
	// passes and the depth test (z-test) fails. Values are from the STENCILOP
	// enumerated type. The default value is STENCILOP_KEEP.
	RS_STENCILZFAIL = 54
	// RS_STENCILPASS is the stencil operation to perform if both the stencil
	// and the depth (z) tests pass. Values are from the STENCILOP enumerated
	// type. The default value is STENCILOP_KEEP.
	RS_STENCILPASS = 55
	// RS_STENCILFUNC is the comparison function for the stencil test. Values
	// are from the CMPFUNC enumerated type. The default value is CMP_ALWAYS.
	// The comparison function is used to compare the reference value to a
	// stencil buffer entry. This comparison applies only to the bits in the
	// reference value and stencil buffer entry that are set in the stencil
	// mask (set by the RS_STENCILMASK render state). If true, the stencil test
	// passes.
	RS_STENCILFUNC = 56
	// RS_STENCILREF is an int reference value for the stencil test. The
	// default value is 0.
	RS_STENCILREF = 57
	// RS_STENCILMASK is the mask applied to the reference value and each
	// stencil buffer entry to determine the significant bits for the stencil
	// test. The default mask is 0xFFFFFFFF.
	RS_STENCILMASK = 58
	// RS_STENCILWRITEMASK is the write mask applied to values written into the
	// stencil buffer. The default mask is 0xFFFFFFFF.
	RS_STENCILWRITEMASK = 59
	// RS_TEXTUREFACTOR is the color used for multiple-texture blending with
	// the TA_TFACTOR texture-blending argument or the TOP_BLENDFACTORALPHA
	// texture-blending operation. The associated value is a COLOR variable.
	// The default value is opaque white (0xFFFFFFFF).
	RS_TEXTUREFACTOR = 60
	// RS_WRAP0 is the texture-wrapping behavior for multiple sets of texture
	// coordinates. Valid values for this render state can be any combination
	// of the WRAPCOORD_0 (or WRAP_U), WRAPCOORD_1 (or WRAP_V), WRAPCOORD_2 (or
	// WRAP_W), and WRAPCOORD_3 flags. These cause the system to wrap in the
	// direction of the first, second, third, and fourth dimensions, sometimes
	// referred to as the s, t, r, and q directions, for a given texture. The
	// default value for this render state is 0 (wrapping disabled in all
	// directions).
	RS_WRAP0 = 128
	// RS_WRAP1 see RS_WRAP0.
	RS_WRAP1 = 129
	// RS_WRAP2 see RS_WRAP0.
	RS_WRAP2 = 130
	// RS_WRAP3 see RS_WRAP0.
	RS_WRAP3 = 131
	// RS_WRAP4 see RS_WRAP0.
	RS_WRAP4 = 132
	// RS_WRAP5 see RS_WRAP0.
	RS_WRAP5 = 133
	// RS_WRAP6 see RS_WRAP0.
	RS_WRAP6 = 134
	// RS_WRAP7 see RS_WRAP0.
	RS_WRAP7 = 135
	// RS_CLIPPING can be set to true to enable primitive clipping by Direct3D,
	// or false to disable it. The default value is true.
	// Use Device.SetRenderStateBool to set this value.
	RS_CLIPPING = 136
	// RS_LIGHTING can be set to true to enable Direct3D lighting, or false to
	// disable it. The default value is true. Only vertices that include a
	// vertex normal are properly lit; vertices that do not contain a normal
	// employ a dot product of 0 in all lighting calculations.
	// Use Device.SetRenderStateBool to set this value.
	RS_LIGHTING = 137
	// RS_AMBIENT is the ambient light color. This value is of type COLOR. The
	// default value is 0.
	RS_AMBIENT = 139
	// RS_FOGVERTEXMODE is the fog formula to be used for vertex fog. Set to
	// one member of the FOGMODE enumerated type. The default value is FOG_NONE.
	RS_FOGVERTEXMODE = 140
	// RS_COLORVERTEX can be set to true to enable per-vertex color or false to
	// disable it. The default value is true. Enabling per-vertex color allows
	// the system to include the color defined for individual vertices in its
	// lighting calculations.
	// For more information, see the following render states:
	// RS_DIFFUSEMATERIALSOURCE, RS_SPECULARMATERIALSOURCE,
	// RS_AMBIENTMATERIALSOURCE, RS_EMISSIVEMATERIALSOURCE.
	// Use Device.SetRenderStateBool to set this value.
	RS_COLORVERTEX = 141
	// RS_LOCALVIEWER can be set to true to enable camera-relative specular
	// highlights, or false to use orthogonal specular highlights. The default
	// value is true. Applications that use orthogonal projection should
	// specify false.
	// Use Device.SetRenderStateBool to set this value.
	RS_LOCALVIEWER = 142
	// RS_NORMALIZENORMALS can be set to true to enable automatic normalization
	// of vertex normals, or false to disable it. The default value is false.
	// Enabling this feature causes the system to normalize the vertex normals
	// for vertices after transforming them to camera space, which can be
	// computationally time-consuming.
	// Use Device.SetRenderStateBool to set this value.
	RS_NORMALIZENORMALS = 143
	// RS_DIFFUSEMATERIALSOURCE is the diffuse color source for lighting
	// calculations. Valid values are members of the MATERIALCOLORSOURCE
	// enumerated type. The default value is MCS_COLOR1. The value for this
	// render state is used only if the RS_COLORVERTEX render state is set to
	// true.
	RS_DIFFUSEMATERIALSOURCE = 145
	// RS_SPECULARMATERIALSOURCE is the specular color source for lighting
	// calculations. Valid values are members of the MATERIALCOLORSOURCE
	// enumerated type. The default value is MCS_COLOR2.
	RS_SPECULARMATERIALSOURCE = 146
	// RS_AMBIENTMATERIALSOURCE is the ambient color source for lighting
	// calculations. Valid values are members of the MATERIALCOLORSOURCE
	// enumerated type. The default value is MCS_MATERIAL.
	RS_AMBIENTMATERIALSOURCE = 147
	// RS_EMISSIVEMATERIALSOURCE is the emissive color source for lighting
	// calculations. Valid values are members of the MATERIALCOLORSOURCE
	// enumerated type. The default value is MCS_MATERIAL.
	RS_EMISSIVEMATERIALSOURCE = 148
	// RS_VERTEXBLEND is the number of matrices to use to perform geometry
	// blending, if any. Valid values are members of the VERTEXBLENDFLAGS
	// enumerated type. The default value is VBF_DISABLE.
	RS_VERTEXBLEND = 151
	// RS_CLIPPLANEENABLE enables or disables user-defined clipping planes.
	// Valid values are any uint32 in which the status of each bit (set or not
	// set) toggles the activation state of a corresponding user-defined
	// clipping plane. The least significant bit (bit 0) controls the first
	// clipping plane at index 0, and subsequent bits control the activation of
	// clipping planes at higher indexes. If a bit is set, the system applies
	// the appropriate clipping plane during scene rendering. The default value
	// is 0.
	// The CLIPPLANEn macros are defined to provide a convenient way to enable
	// clipping planes.
	RS_CLIPPLANEENABLE = 152
	// RS_POINTSIZE is a float value that specifies the size to use for point
	// size computation in cases where point size is not specified for each
	// vertex. This value is not used when the vertex contains point size. This
	// value is in screen space units if RS_POINTSCALEENABLE is false;
	// otherwise this value is in world space units. The default value is the
	// value a driver returns. If a driver returns 0 or 1, the default value is
	// 64, which allows software point size emulation.
	// Use Device.SetRenderStateFloat to set this value.
	RS_POINTSIZE = 154
	// RS_POINTSIZE_MIN is a float value that specifies the minimum size of
	// point primitives. Point primitives are clamped to this size during
	// rendering. Setting this to values smaller than 1.0 results in points
	// dropping out when the point does not cover a pixel center and
	// antialiasing is disabled or being rendered with reduced intensity when
	// antialiasing is enabled. The default value is 1.0f. The range for this
	// value is greater than or equal to 0.0f.
	// Use Device.SetRenderStateFloat to set this value.
	RS_POINTSIZE_MIN = 155
	// RS_POINTSPRITEENABLE is a bool value. When true, texture coordinates of
	// point primitives are set so that full textures are mapped on each point.
	// When false, the vertex texture coordinates are used for the entire
	// point. The default value is false. You can achieve DirectX 7 style
	// single-pixel points by setting RS_POINTSCALEENABLE to false and
	// RS_POINTSIZE to 1.0, which are the default values.
	// Use Device.SetRenderStateBool to set this value.
	RS_POINTSPRITEENABLE = 156
	// RS_POINTSCALEENABLE is a bool value that controls computation of size
	// for point primitives. When true, the point size is interpreted as a
	// camera space value and is scaled by the distance function and the
	// frustum to viewport y-axis scaling to compute the final screen-space
	// point size. When false, the point size is interpreted as screen space
	// and used directly. The default value is false.
	// Use Device.SetRenderStateBool to set this value.
	RS_POINTSCALEENABLE = 157
	// RS_POINTSCALE_A is a float value that controls for distance-based size
	// attenuation for point primitives. Active only when RS_POINTSCALEENABLE
	// is true. The default value is 1.0f. The range for this value is greater
	// than or equal to 0.0f.
	// Use Device.SetRenderStateFloat to set this value.
	RS_POINTSCALE_A = 158
	// RS_POINTSCALE_B is a float value that controls for distance-based size
	// attenuation for point primitives. Active only when RS_POINTSCALEENABLE
	// is true. The default value is 0.0f. The range for this value is greater
	// than or equal to 0.0f.
	// Use Device.SetRenderStateFloat to set this value.
	RS_POINTSCALE_B = 159
	// RS_POINTSCALE_C is a float value that controls for distance-based size
	// attenuation for point primitives. Active only when RS_POINTSCALEENABLE
	// is true. The default value is 0.0f. The range for this value is greater
	// than or equal to 0.0f.
	// Use Device.SetRenderStateFloat to set this value.
	RS_POINTSCALE_C = 160
	// RS_MULTISAMPLEANTIALIAS is a bool value that determines how individual
	// samples are computed when using a multisample render-target buffer. When
	// set to true, the multiple samples are computed so that full-scene
	// antialiasing is performed by sampling at different sample positions for
	// each multiple sample. When set to false, the multiple samples are all
	// written with the same sample value, sampled at the pixel center, which
	// allows non-antialiased rendering to a multisample buffer. This render
	// state has no effect when rendering to a single sample buffer. The
	// default value is true.
	// Use Device.SetRenderStateBool to set this value.
	RS_MULTISAMPLEANTIALIAS = 161
	// RS_MULTISAMPLEMASK is a uint32 mask. Each bit in this mask, starting at
	// the least significant bit (LSB), controls modification of one of the
	// samples in a multisample render target. Thus, for an 8-sample render
	// target, the low byte contains the eight write enables for each of the
	// eight samples. This render state has no effect when rendering to a
	// single sample buffer. The default value is 0xFFFFFFFF.
	// This render state enables use of a multisample buffer as an accumulation
	// buffer, doing multipass rendering of geometry where each pass updates a
	// subset of samples.
	// If there are n multisamples and k enabled samples, the resulting
	// intensity of the rendered image should be k/n. Each component RGB of
	// every pixel is factored by k/n.
	RS_MULTISAMPLEMASK = 162
	// RS_PATCHEDGESTYLE sets whether patch edges will use float style
	// tessellation. Possible values are defined by the PATCHEDGESTYLE
	// enumerated type. The default value is PATCHEDGE_DISCRETE.
	RS_PATCHEDGESTYLE = 163
	// RS_DEBUGMONITORTOKEN is set only for debugging the monitor. Possible
	// values are defined by the DEBUGMONITORTOKENS enumerated type. Note that
	// if RS_DEBUGMONITORTOKEN is set, the call is treated as passing a token
	// to the debug monitor. For example, if - after passing DMT_ENABLE or
	// DMT_DISABLE to RS_DEBUGMONITORTOKEN - other token values are passed in,
	// the state (enabled or disabled) of the debug monitor will still persist.
	// This state is only useful for debug builds. The debug monitor defaults
	// to DMT_ENABLE.
	RS_DEBUGMONITORTOKEN = 165
	// RS_POINTSIZE_MAX is a float value that specifies the maximum size to
	// which point sprites will be clamped. The value must be less than or
	// equal to the MaxPointSize member of CAPS and greater than or equal to
	// RS_POINTSIZE_MIN. The default value is 64.0.
	// Use Device.SetRenderStateFloat to set this value.
	RS_POINTSIZE_MAX = 166
	// RS_INDEXEDVERTEXBLENDENABLE is a bool value that enables or disables
	// indexed vertex blending. The default value is false. When set to true,
	// indexed vertex blending is enabled. When set to false, indexed vertex
	// blending is disabled. If this render state is enabled, the user must
	// pass matrix indices as a packed DWORDwith every vertex. When the render
	// state is disabled and vertex blending is enabled through the
	// RS_VERTEXBLEND state, it is equivalent to having matrix indices 0, 1, 2,
	// 3 in every vertex.
	// Use Device.SetRenderStateBool to set this value.
	RS_INDEXEDVERTEXBLENDENABLE = 167
	// RS_COLORWRITEENABLE is a uint32 value that enables a per-channel write
	// for the render-target color buffer. A set bit results in the color
	// channel being updated during 3D rendering. A clear bit results in the
	// color channel being unaffected. This functionality is available if the
	// PMISCCAPS_COLORWRITEENABLE capabilities bit is set in the
	// PrimitiveMiscCaps member of the CAPS structure for the device. This
	// render state does not affect the clear operation. The default value is
	// 0x0000000F.
	// Valid values for this render state can be any combination of the
	// COLORWRITEENABLE_ALPHA, COLORWRITEENABLE_BLUE, COLORWRITEENABLE_GREEN,
	// or COLORWRITEENABLE_RED flags.
	RS_COLORWRITEENABLE = 168
	// RS_TWEENFACTOR is a float value that controls the tween factor. The
	// default value is 0.0f.
	// Use Device.SetRenderStateFloat to set this value.
	RS_TWEENFACTOR = 170
	// RS_BLENDOP is used to select the arithmetic operation applied when the
	// alpha blending render state, RS_ALPHABLENDENABLE, is set to true. Valid
	// values are defined by the BLENDOP enumerated type. The default value is
	// BLENDOP_ADD.
	// If the PMISCCAPS_BLENDOP device capability is not supported, then
	// BLENDOP_ADD is performed.
	RS_BLENDOP = 171
	// RS_POSITIONDEGREE is the N-patch position interpolation degree. The
	// values can be DEGREE_CUBIC (default) or DEGREE_LINEAR. For more
	// information, see DEGREETYPE.
	RS_POSITIONDEGREE = 172
	// RS_NORMALDEGREE is the N-patch normal interpolation degree. The values
	// can be DEGREE_LINEAR (default) or DEGREE_QUADRATIC. For more
	// information, see DEGREETYPE.
	RS_NORMALDEGREE = 173
	// RS_SCISSORTESTENABLE can be set to true to enable scissor testing and
	// false to disable it. The default value is false.
	RS_SCISSORTESTENABLE = 174
	// RS_SLOPESCALEDEPTHBIAS is used to determine how much bias can be applied
	// to co-planar primitives to reduce z-fighting. The default value is 0.
	// bias = (max * RS_SLOPESCALEDEPTHBIAS) + RS_DEPTHBIAS.
	// where max is the maximum depth slope of the triangle being rendered.
	RS_SLOPESCALEDEPTHBIAS = 175
	// RS_ANTIALIASEDLINEENABLE can be set to true to enable line antialiasing,
	// false to disable line antialiasing. The default value is false.
	// When rendering to a multisample render target, RS_ANTIALIASEDLINEENABLE
	// is ignored and all lines are rendered aliased.
	RS_ANTIALIASEDLINEENABLE = 176
	// RS_MINTESSELLATIONLEVEL is the minimum tessellation level. The default
	// value is 1.0f.
	RS_MINTESSELLATIONLEVEL = 178
	// RS_MAXTESSELLATIONLEVEL is the maximum tessellation level. The default
	// value is 1.0f.
	RS_MAXTESSELLATIONLEVEL = 179
	// RS_ADAPTIVETESS_X is the amount to adaptively tessellate, in the x
	// direction. Default value is 0.0f.
	RS_ADAPTIVETESS_X = 180
	// RS_ADAPTIVETESS_Y is the amount to adaptively tessellate, in the y
	// direction. Default value is 0.0f.
	RS_ADAPTIVETESS_Y = 181
	// RS_ADAPTIVETESS_Z is the amount to adaptively tessellate, in the z
	// direction. Default value is 1.0f.
	RS_ADAPTIVETESS_Z = 182
	// RS_ADAPTIVETESS_W is the amount to adaptively tessellate, in the w
	// direction. Default value is 0.0f.
	RS_ADAPTIVETESS_W = 183
	// RS_ENABLEADAPTIVETESSELLATION can be set to true to enable adaptive
	// tessellation, false to disable it. The default value is false.
	RS_ENABLEADAPTIVETESSELLATION = 184
	// RS_TWOSIDEDSTENCILMODE can be set to true to enable two-sided
	// stenciling, false disables it. The default value is false. The
	// application should set RS_CULLMODE to CULL_NONE to enable two-sided
	// stencil mode. If the triangle winding order is clockwise, the
	// RS_STENCIL* operations will be used. If the winding order is
	// counterclockwise, the RS_CCW_STENCIL* operations will be used.
	// To see if two-sided stencil is supported, check the StencilCaps member
	// of CAPS for STENCILCAPS_TWOSIDED. See also STENCILCAPS.
	RS_TWOSIDEDSTENCILMODE = 185
	// RS_CCW_STENCILFAIL is the stencil operation to perform if CCW stencil
	// test fails. Values are from the STENCILOP enumerated type. The default
	// value is STENCILOP_KEEP.
	RS_CCW_STENCILFAIL = 186
	// RS_CCW_STENCILZFAIL is the stencil operation to perform if CCW stencil
	// test passes and z-test fails. Values are from the STENCILOP enumerated
	// type. The default value is STENCILOP_KEEP.
	RS_CCW_STENCILZFAIL = 187
	// RS_CCW_STENCILPASS is the stencil operation to perform if both CCW
	// stencil and z-tests pass. Values are from the STENCILOP enumerated type.
	// The default value is STENCILOP_KEEP.
	RS_CCW_STENCILPASS = 188
	// RS_CCW_STENCILFUNC is the comparison function. CCW stencil test passes
	// if ((ref & mask) stencil function (stencil & mask)) is true. Values are
	// from the CMPFUNC enumerated type. The default value is CMP_ALWAYS.
	RS_CCW_STENCILFUNC = 189
	// RS_COLORWRITEENABLE1 specifies additional ColorWriteEnable values for
	// the devices. See RS_COLORWRITEENABLE. This functionality is available if
	// the PMISCCAPS_INDEPENDENTWRITEMASKS capabilities bit is set in the
	// PrimitiveMiscCaps member of the CAPS structure for the device. The
	// default value is 0x0000000F.
	RS_COLORWRITEENABLE1 = 190
	// RS_COLORWRITEENABLE2 specifies additional ColorWriteEnable values for
	// the devices. See RS_COLORWRITEENABLE. This functionality is available if
	// the PMISCCAPS_INDEPENDENTWRITEMASKS capabilities bit is set in the
	// PrimitiveMiscCaps member of the CAPS structure for the device. The
	// default value is 0x0000000F.
	RS_COLORWRITEENABLE2 = 191
	// RS_COLORWRITEENABLE3 Additional ColorWriteEnable values for the devices.
	// See RS_COLORWRITEENABLE. This functionality is available if the
	// PMISCCAPS_INDEPENDENTWRITEMASKS capabilities bit is set in the
	// PrimitiveMiscCaps member of the CAPS structure for the device. The
	// default value is 0x0000000F.
	RS_COLORWRITEENABLE3 = 192
	// RS_BLENDFACTOR is a COLOR used for a constant blend-factor during alpha
	// blending. This functionality is available if the PBLENDCAPS_BLENDFACTOR
	// capabilities bit is set in the SrcBlendCaps member of CAPS or the
	// DestBlendCaps member of CAPS. See RENDERSTATETYPE. The default value is
	// 0xFFFFFFFF.
	RS_BLENDFACTOR = 193
	// RS_SRGBWRITEENABLE enables render-target writes to be gamma corrected to
	// sRGB. The format must expose USAGE_SRGBWRITE. The default value is 0.
	RS_SRGBWRITEENABLE = 194
	// RS_DEPTHBIAS is a floating-point value that is used for comparison of
	// depth values. The default value is 0.
	// Use Device.SetRenderStateFloat to set this value.
	RS_DEPTHBIAS = 195
	// RS_WRAP8 see RS_WRAP0.
	RS_WRAP8 = 198
	// RS_WRAP9 see RS_WRAP0.
	RS_WRAP9 = 199
	// RS_WRAP10 see RS_WRAP0.
	RS_WRAP10 = 200
	// RS_WRAP11 see RS_WRAP0.
	RS_WRAP11 = 201
	// RS_WRAP12 see RS_WRAP0.
	RS_WRAP12 = 202
	// RS_WRAP13 see RS_WRAP0.
	RS_WRAP13 = 203
	// RS_WRAP14 see RS_WRAP0.
	RS_WRAP14 = 204
	// RS_WRAP15 see RS_WRAP0.
	RS_WRAP15 = 205
	// RS_SEPARATEALPHABLENDENABLE can be set to true to enable the separate
	// blend mode for the alpha channel. The default value is false.
	// When set to false, the render-target blending factors and operations
	// applied to alpha are forced to be the same as those defined for color.
	// This mode is effectively hardwired to false on implementations that
	// don't set the cap PMISCCAPS_SEPARATEALPHABLEND. See PMISCCAPS.
	// The type of separate alpha blending is determined by the
	// RS_SRCBLENDALPHA and RS_DESTBLENDALPHA render states.
	RS_SEPARATEALPHABLENDENABLE = 206
	// RS_SRCBLENDALPHA is one member of the BLEND enumerated type. This value
	// is ignored unless RS_SEPARATEALPHABLENDENABLE is true. The default value
	// is BLEND_ONE.
	RS_SRCBLENDALPHA = 207
	// RS_DESTBLENDALPHA is one member of the BLEND enumerated type. This value
	// is ignored unless RS_SEPARATEALPHABLENDENABLE is true. The default value
	// is BLEND_ZERO.
	RS_DESTBLENDALPHA = 208
	// RS_BLENDOPALPHA is a value used to select the arithmetic operation
	// applied to separate alpha blending when the render state,
	// RS_SEPARATEALPHABLENDENABLE, is set to true.
	// Valid values are defined by the BLENDOP enumerated type. The default
	// value is BLENDOP_ADD.
	// If the PMISCCAPS_BLENDOP device capability is not supported, then
	// BLENDOP_ADD is performed. See PMISCCAPS.
	RS_BLENDOPALPHA = 209
)

// RESOURCETYPE defines resource types.
type RESOURCETYPE uint32

const (
	// RTYPE_SURFACE is a surface resource.
	RTYPE_SURFACE = 1
	// RTYPE_VOLUME is a volume resource.
	RTYPE_VOLUME = 2
	// RTYPE_TEXTURE is a texture resource.
	RTYPE_TEXTURE = 3
	// RTYPE_VOLUMETEXTURE is a volume texture resource.
	RTYPE_VOLUMETEXTURE = 4
	// RTYPE_CUBETEXTURE is a cube texture resource.
	RTYPE_CUBETEXTURE = 5
	// RTYPE_VERTEXBUFFER is a vertex buffer resource.
	RTYPE_VERTEXBUFFER = 6
	// RTYPE_INDEXBUFFER is an index buffer resource.
	RTYPE_INDEXBUFFER = 7
)

// SAMPLER_TEXTURE_TYPE defines the sampler texture types for vertex shaders.
type SAMPLER_TEXTURE_TYPE uint32

const (
	// STT_UNKNOWN is an uninitialized value. The value of this element is
	// SP_TEXTURETYPE_SHIFT.
	STT_UNKNOWN = 0
	// STT_2D declars a 2D texture. The value of this element is
	// SP_TEXTURETYPE_SHIFT * 4.
	STT_2D = 268435456
	// STT_CUBE declars a cube texture. The value of this element is
	// SP_TEXTURETYPE_SHIFT * 8.
	STT_CUBE = 402653184
	// STT_VOLUME declars a volume texture. The value of this element is
	// SP_TEXTURETYPE_SHIFT * 16.
	STT_VOLUME = 536870912
)

// SAMPLERSTATETYPE defines texture sampling operations such as texture
// addressing and texture filtering. Some sampler states set-up vertex
// processing, and some set-up pixel processing. Sampler states can be saved
// and restored using stateblocks.
type SAMPLERSTATETYPE uint32

const (
	// SAMP_ADDRESSU specifies the texture-address mode for the u coordinate.
	// The default is TADDRESS_WRAP. For more information, see TEXTUREADDRESS.
	SAMP_ADDRESSU = 1
	// SAMP_ADDRESSV specifies the texture-address mode for the v coordinate.
	// The default is TADDRESS_WRAP. For more information, see TEXTUREADDRESS.
	SAMP_ADDRESSV = 2
	// SAMP_ADDRESSW specifies the texture-address mode for the w coordinate.
	// The default is TADDRESS_WRAP. For more information, see TEXTUREADDRESS.
	SAMP_ADDRESSW = 3
	// SAMP_BORDERCOLOR specifies the border color or type COLOR. The default
	// color is 0x00000000.
	SAMP_BORDERCOLOR = 4
	// SAMP_MAGFILTER specifies the magnification filter of type
	// TEXTUREFILTERTYPE. The default value is TEXF_POINT.
	SAMP_MAGFILTER = 5
	// SAMP_MINFILTER specifies the minification filter of type
	// TEXTUREFILTERTYPE. The default value is TEXF_POINT.
	SAMP_MINFILTER = 6
	// SAMP_MIPFILTER specifies the mipmap filter to use during minification.
	// See TEXTUREFILTERTYPE. The default value is TEXF_NONE.
	SAMP_MIPFILTER = 7
	// SAMP_MIPMAPLODBIAS specifies the mipmap level-of-detail bias. The
	// default value is zero.
	SAMP_MIPMAPLODBIAS = 8
	// SAMP_MAXMIPLEVEL specifies the level-of-detail index of the largest map
	// to use. Values range from 0 to (n - 1) where 0 is the largest. The
	// default value is zero.
	SAMP_MAXMIPLEVEL = 9
	// SAMP_MAXANISOTROPY specifies the maximum anisotropy. Values range from 1
	// to the value that is specified in the MaxAnisotropy member of the CAPS
	// structure. The default value is 1.
	SAMP_MAXANISOTROPY = 10
	// SAMP_SRGBTEXTURE specifies the gamma correction value. The default value
	// is 0, which means gamma is 1.0 and no correction is required. Otherwise,
	// this value means that the sampler should assume gamma of 2.2 on the
	// content and convert it to linear (gamma 1.0) before presenting it to the
	// pixel shader.
	SAMP_SRGBTEXTURE = 11
	// SAMP_ELEMENTINDEX indicates which element index to use, when a
	// multielement texture is assigned to the sampler. The default value is 0.
	SAMP_ELEMENTINDEX = 12
	// SAMP_DMAPOFFSET specifies the vertex offset in the presampled
	// displacement map. This is a constant used by the tessellator, its
	// default value is 0.
	SAMP_DMAPOFFSET = 13
)

// SHADEMODE defines constants that describe the supported shading modes.
type SHADEMODE uint32

const (
	// SHADE_FLAT means flat shading mode. The color and specular component of
	// the first vertex in the triangle are used to determine the color and
	// specular component of the face. These colors remain constant across the
	// triangle; that is, they are not interpolated. The specular alpha is
	// interpolated.
	SHADE_FLAT = 1
	// SHADE_GOURAUD means gouraud shading mode. The color and specular
	// components of the face are determined by a linear interpolation between
	// all three of the triangle's vertices.
	SHADE_GOURAUD = 2
	// SHADE_PHONG is not supported.
	SHADE_PHONG = 3
)

// STATEBLOCKTYPE predefines sets of pipeline state used by state blocks.
type STATEBLOCKTYPE uint32

const (
	// SBT_ALL captures the current device state.
	SBT_ALL = 1
	// SBT_PIXELSTATE capture the current pixel state.
	SBT_PIXELSTATE = 2
	// SBT_VERTEXSTATE capture the current vertex state.
	SBT_VERTEXSTATE = 3
)

// STENCILOP defines stencil-buffer operations.
type STENCILOP uint32

const (
	// STENCILOP_KEEP does not update the entry in the stencil buffer. This is
	// the default value.
	STENCILOP_KEEP = 1
	// STENCILOP_ZERO sets the stencil-buffer entry to 0.
	STENCILOP_ZERO = 2
	// STENCILOP_REPLACE replaces the stencil-buffer entry with a reference
	// value.
	STENCILOP_REPLACE = 3
	// STENCILOP_INCRSAT increments the stencil-buffer entry, clamping to the
	// maximum value.
	STENCILOP_INCRSAT = 4
	// STENCILOP_DECRSAT decrements the stencil-buffer entry, clamping to zero.
	STENCILOP_DECRSAT = 5
	// STENCILOP_INVERT inverts the bits in the stencil-buffer entry.
	STENCILOP_INVERT = 6
	// STENCILOP_INCR increments the stencil-buffer entry, wrapping to zero if
	// the new value exceeds the maximum value.
	STENCILOP_INCR = 7
	// STENCILOP_DECR decrements the stencil-buffer entry, wrapping to the
	// maximum value if the new value is less than zero.
	STENCILOP_DECR = 8
)

// SWAPEFFECT defines swap effects.
type SWAPEFFECT uint32

const (
	// SWAPEFFECT_DISCARD When a swap chain is created with a swap effect of
	// SWAPEFFECT_FLIP or SWAPEFFECT_COPY, the runtime will guarantee that an
	// Device.Present operation will not affect the content of any of the back
	// buffers. Unfortunately, meeting this guarantee can involve substantial
	// video memory or processing overheads, especially when implementing flip
	// semantics for a windowed swap chain or copy semantics for a full-screen
	// swap chain. An application may use the SWAPEFFECT_DISCARD swap effect to
	// avoid these overheads and to enable the display driver to select the
	// most efficient presentation technique for the swap chain. This is also
	// the only swap effect that may be used when specifying a value other than
	// MULTISAMPLE_NONE for the MultiSampleType member of PRESENT_PARAMETERS.
	// Like a swap chain that uses SWAPEFFECT_FLIP, a swap chain that uses
	// SWAPEFFECT_DISCARD might include more than one back buffer, any of which
	// may be accessed using Device.GetBackBuffer or SwapChain.GetBackBuffer.
	// The swap chain is best envisaged as a queue in which 0 always indexes
	// the back buffer that will be displayed by the next Present operation and
	// from which buffers are discarded when they have been displayed.
	// An application that uses this swap effect cannot make any assumptions
	// about the contents of a discarded back buffer and should therefore
	// update an entire back buffer before invoking a Present operation that
	// would display it. Although this is not enforced, the debug version of
	// the runtime will overwrite the contents of discarded back buffers with
	// random data to enable developers to verify that their applications are
	// updating the entire back buffer surfaces correctly.
	SWAPEFFECT_DISCARD = 1
	// SWAPEFFECT_FLIP The swap chain might include multiple back buffers and
	// is best envisaged as a circular queue that includes the front buffer.
	// Within this queue, the back buffers are always numbered sequentially
	// from 0 to (n - 1), where n is the number of back buffers, so that 0
	// denotes the least recently presented buffer. When Present is invoked,
	// the queue is "rotated" so that the front buffer becomes back buffer (n -
	// 1), while the back buffer 0 becomes the new front buffer.
	SWAPEFFECT_FLIP = 2
	// SWAPEFFECT_COPY This swap effect may be specified only for a swap chain
	// comprising a single back buffer. Whether the swap chain is windowed or
	// full-screen, the runtime will guarantee the semantics implied by a
	// copy-based Present operation, namely that the operation leaves the
	// content of the back buffer unchanged, instead of replacing it with the
	// content of the front buffer as a flip-based Present operation would.
	// For a full-screen swap chain, the runtime uses a combination of flip
	// operations and copy operations, supported if necessary by hidden back
	// buffers, to accomplish the Present operation. Accordingly, the
	// presentation is synchronized with the display adapter's vertical retrace
	// and its rate is constrained by the chosen presentation interval. A swap
	// chain specified with the PRESENT_INTERVAL_IMMEDIATE flag is the only
	// exception. (Refer to the description of the PresentationIntervals member
	// of the PRESENT_PARAMETERS structure.) In this case, a Present operation
	// copies the back buffer content directly to the front buffer without
	// waiting for the vertical retrace.
	SWAPEFFECT_COPY = 3
)

// TEXTUREADDRESS defines constants that describe the supported
// texture-addressing modes.
type TEXTUREADDRESS uint32

const (
	// TADDRESS_WRAP tiles the texture at every integer junction. For example,
	// for u values between 0 and 3, the texture is repeated three times; no
	// mirroring is performed.
	TADDRESS_WRAP = 1
	// TADDRESS_MIRROR is similar to TADDRESS_WRAP, except that the texture is
	// flipped at every integer junction. For u values between 0 and 1, for
	// example, the texture is addressed normally; between 1 and 2, the texture
	// is flipped (mirrored); between 2 and 3, the texture is normal again; and
	// so on.
	TADDRESS_MIRROR = 2
	// TADDRESS_CLAMP means texture coordinates outside the range [0.0, 1.0]
	// are set to the texture color at 0.0 or 1.0, respectively.
	TADDRESS_CLAMP = 3
	// TADDRESS_BORDER means texture coordinates outside the range [0.0, 1.0]
	// are set to the border color.
	TADDRESS_BORDER = 4
	// TADDRESS_MIRRORONCE is similar to TADDRESS_MIRROR and TADDRESS_CLAMP.
	// Takes the absolute value of the texture coordinate (thus, mirroring
	// around 0), and then clamps to the maximum value. The most common usage
	// is for volume textures, where support for the full TADDRESS_MIRRORONCE
	// texture-addressing mode is not necessary, but the data is symmetric
	// around the one axis.
	TADDRESS_MIRRORONCE = 5
)

// TEXTUREFILTERTYPE defines texture filtering modes for a texture stage.
type TEXTUREFILTERTYPE uint32

const (
	// TEXF_NONE when used with SAMP_MIPFILTER, disables mipmapping.
	TEXF_NONE = 0
	// TEXF_POINT when used with SAMP_ MAGFILTER or SAMP_MINFILTER, specifies
	// that point filtering is to be used as the texture magnification or
	// minification filter respectively. When used with SAMP_MIPFILTER, enables
	// mipmapping and specifies that the rasterizer chooses the color from the
	// texel of the nearest mip level.
	TEXF_POINT = 1
	// TEXF_LINEAR when used with SAMP_ MAGFILTER or SAMP_MINFILTER, specifies
	// that linear filtering is to be used as the texture magnification or
	// minification filter respectively. When used with SAMP_MIPFILTER, enables
	// mipmapping and trilinear filtering; it specifies that the rasterizer
	// interpolates between the two nearest mip levels.
	TEXF_LINEAR = 2
	// TEXF_ANISOTROPIC when used with SAMP_ MAGFILTER or SAMP_MINFILTER,
	// specifies that anisotropic texture filtering used as a texture
	// magnification or minification filter respectively. Compensates for
	// distortion caused by the difference in angle between the texture polygon
	// and the plane of the screen. Use with SAMP_MIPFILTER is undefined.
	TEXF_ANISOTROPIC = 3
	// TEXF_PYRAMIDALQUAD is a 4-sample tent filter used as a texture
	// magnification or minification filter. Use with SAMP_MIPFILTER is
	// undefined.
	TEXF_PYRAMIDALQUAD = 6
	// TEXF_GAUSSIANQUAD is a 4-sample Gaussian filter used as a texture
	// magnification or minification filter. Use with SAMP_MIPFILTER is
	// undefined.
	TEXF_GAUSSIANQUAD = 7
)

// TEXTUREOP defines per-stage texture-blending operations.
type TEXTUREOP uint32

const (
	// TOP_DISABLE disables output from this texture stage and all stages with
	// a higher index. To disable texture mapping, set this as the color
	// operation for the first texture stage (stage 0). Alpha operations cannot
	// be disabled when color operations are enabled. Setting the alpha
	// operation to TOP_DISABLE when color blending is enabled causes undefined
	// behavior.
	TOP_DISABLE = 1
	// TOP_SELECTARG1 uses this texture stage's first color or alpha argument,
	// unmodified, as the output. This operation affects the color argument
	// when used with the TSS_COLOROP texture-stage state, and the alpha
	// argument when used with TSS_ALPHAOP.
	// S(RGBA) = Arg1
	TOP_SELECTARG1 = 2
	// TOP_SELECTARG2 Use this texture stage's second color or alpha argument,
	// unmodified, as the output. This operation affects the color argument
	// when used with the D3DTSS_COLOROP texture stage state, and the alpha
	// argument when used with D3DTSS_ALPHAOP.
	// S(RGBA) = Arg2
	TOP_SELECTARG2 = 3
	// TOP_MODULATE multiplies the components of the arguments.
	// S(RGBA) = Arg1*Arg2
	TOP_MODULATE = 4
	// TOP_MODULATE2X multiplies the components of the arguments, and shift the
	// products to the left 1 bit (effectively multiplying them by 2) for
	// brightening.
	// S(RGBA) = (Arg1*Arg2)<<1
	TOP_MODULATE2X = 5
	// TOP_MODULATE4X multiplies the components of the arguments, and shift the
	// products to the left 2 bits (effectively multiplying them by 4) for
	// brightening.
	// S(RGBA) = (Arg1*Arg2)<<2
	TOP_MODULATE4X = 6
	// TOP_ADD adds the components of the arguments.
	// S(RGBA) = Arg1+Arg2
	TOP_ADD = 7
	// TOP_ADDSIGNED adds the components of the arguments with a - 0.5 bias,
	// making the effective range of values from - 0.5 through 0.5.
	// S(RGBA) = Arg1+Arg2-0.5
	TOP_ADDSIGNED = 8
	// TOP_ADDSIGNED2X adds the components of the arguments with a - 0.5 bias,
	// and shift the products to the left 1 bit.
	// S(RGBA) = (Arg1+Arg2-0.5)<<1
	TOP_ADDSIGNED2X = 9
	// TOP_SUBTRACT subtracts the components of the second argument from those
	// of the first argument.
	// S(RGBA) = Arg1-Arg2
	TOP_SUBTRACT = 10
	// TOP_ADDSMOOTH adds the first and second arguments; then subtract their
	// product from the sum.
	// S(RGBA) = Arg1+Arg2 - Arg1*Arg2 = Arg1+Arg2*(1-Arg1)
	TOP_ADDSMOOTH = 11
	// TOP_BLENDDIFFUSEALPHA linearly blends this texture stage, using the
	// interpolated alpha from each vertex.
	// S(RGBA) = Arg1*Alpha + Arg2*(1-Alpha)
	TOP_BLENDDIFFUSEALPHA = 12
	// TOP_BLENDTEXTUREALPHA linearly blends this texture stage, using the
	// alpha from this stage's texture.
	// S(RGBA) = Arg1*Alpha + Arg2*(1-Alpha)
	TOP_BLENDTEXTUREALPHA = 13
	// TOP_BLENDFACTORALPHA linearly blends this texture stage, using a scalar
	// alpha set with the D3DRS_TEXTUREFACTOR render state.
	// S(RGBA) = Arg1*Alpha + Arg2*(1-Alpha)
	TOP_BLENDFACTORALPHA = 14
	// TOP_BLENDTEXTUREALPHAPM linearly blends a texture stage that uses a
	// premultiplied alpha.
	// S(RGBA) = Arg1 + Arg2*(1-Alpha)
	TOP_BLENDTEXTUREALPHAPM = 15
	// TOP_BLENDCURRENTALPHA linearly blends this texture stage, using the
	// alpha taken from the previous texture stage.
	// S(RGBA) = Arg1*Alpha + Arg2*(1-Alpha)
	TOP_BLENDCURRENTALPHA = 16
	// TOP_PREMODULATE is set in stage n. The output of stage n is arg1.
	// Additionally, if there is a texture in stage n + 1, any TA_CURRENT in
	// stage n + 1 is premultiplied by texture in stage n + 1.
	TOP_PREMODULATE = 17
	// TOP_MODULATEALPHA_ADDCOLOR modulates the color of the second argument,
	// using the alpha of the first argument; then add the result to argument
	// one. This operation is supported only for color operations
	// (D3DTSS_COLOROP).
	// S(RGBA) = Arg1(RGB)+Arg1(A)*Arg2(RGB)
	TOP_MODULATEALPHA_ADDCOLOR = 18
	// TOP_MODULATECOLOR_ADDALPHA modulates the arguments; then add the alpha
	// of the first argument. This operation is supported only for color
	// operations (D3DTSS_COLOROP).
	// S(RGBA) = Arg1(RGB)*Arg2(RGB)+Arg1(A)
	TOP_MODULATECOLOR_ADDALPHA = 19
	// TOP_MODULATEINVALPHA_ADDCOLOR is similar to TOP_MODULATEALPHA_ADDCOLOR,
	// but use the inverse of the alpha of the first argument. This operation
	// is supported only for color operations (TSS_COLOROP).
	// S(RGBA) = (1-Arg1(A))*Arg2(RGB)+Arg1(RGB)
	TOP_MODULATEINVALPHA_ADDCOLOR = 20
	// TOP_MODULATEINVCOLOR_ADDALPHA is similar to TOP_MODULATECOLOR_ADDALPHA,
	// but use the inverse of the color of the first argument. This operation
	// is supported only for color operations (TSS_COLOROP).
	// S(RGBA) = (1-Arg1(RGB))*Arg2(RGB)+Arg1(A)
	TOP_MODULATEINVCOLOR_ADDALPHA = 21
	// TOP_BUMPENVMAP performs per-pixel bump mapping, using the environment
	// map in the next texture stage, without luminance. This operation is
	// supported only for color operations (TSS_COLOROP).
	TOP_BUMPENVMAP = 22
	// TOP_BUMPENVMAPLUMINANCE performs per-pixel bump mapping, using the
	// environment map in the next texture stage, with luminance. This
	// operation is supported only for color operations (TSS_COLOROP).
	TOP_BUMPENVMAPLUMINANCE = 23
	// TOP_DOTPRODUCT3 modulates the components of each argument as signed
	// components, add their products; then replicate the sum to all color
	// channels, including alpha. This operation is supported for color and
	// alpha operations.
	// S(RGBA) = Arg1(R)*Arg2(R) + Arg1(G)*Arg2(G) + Arg1(B)*Arg2(B)
	// In DirectX 6 and DirectX 7, multitexture operations the above inputs are
	// all shifted down by half (y = x - 0.5) before use to simulate signed
	// data, and the scalar result is automatically clamped to positive values
	// and replicated to all three output channels. Also, note that as a color
	// operation this does not updated the alpha it just updates the RGB
	// components.
	// However, in DirectX 8.1 shaders you can specify that the output be
	// routed to the .rgb or the .a components or both (the default). You can
	// also specify a separate scalar operation on the alpha channel.
	TOP_DOTPRODUCT3 = 24
	// TOP_MULTIPLYADD performs a multiply-accumulate operation. It takes the
	// last two arguments, multiplies them together, and adds them to the
	// remaining input/source argument, and places that into the result.
	// S(RGBA) = Arg1 + Arg2*Arg3
	TOP_MULTIPLYADD = 25
	// TOP_LERP linearly interpolates between the second and third source
	// arguments by a proportion specified in the first source argument.
	// S(RGBA) = Arg1*Arg2 + (1-Arg1)*Arg3
	TOP_LERP = 26
)

// TEXTURESTAGESTATETYPE defines multi-blender texture operations. Some sampler
// states set up vertex processing, and some set up pixel processing. Texture
// stage states can be saved and restored using stateblocks.
type TEXTURESTAGESTATETYPE uint32

const (
	// TSS_COLOROP texture-stage state is a texture color blending operation
	// identified by one member of the TEXTUREOP enumerated type. The default
	// value for the first texture stage (stage 0) is TOP_MODULATE; for all
	// other stages the default is TOP_DISABLE.
	TSS_COLOROP = 1
	// TSS_COLORARG1 texture-stage state is the first color argument for the
	// stage, identified by one of the TA. The default argument is TA_TEXTURE.
	// Specify TA_TEMP to select a temporary register color for read or write.
	// TA_TEMP is supported if the PMISCCAPS_TSSARGTEMP device capability is
	// present. The default value for the register is (0.0, 0.0, 0.0, 0.0).
	TSS_COLORARG1 = 2
	// TSS_COLORARG2 texture-stage state is the second color argument for the
	// stage, identified by TA. The default argument is TA_CURRENT. Specify
	// TA_TEMP to select a temporary register color for read or write. TA_TEMP
	// is supported if the PMISCCAPS_TSSARGTEMP device capability is present.
	// The default value for the register is (0.0, 0.0, 0.0, 0.0)
	TSS_COLORARG2 = 3
	// TSS_ALPHAOP texture-stage state is a texture alpha blending operation
	// identified by one member of the TEXTUREOP enumerated type. The default
	// value for the first texture stage (stage 0) is TOP_SELECTARG1, and for
	// all other stages the default is TOP_DISABLE.
	TSS_ALPHAOP = 4
	// TSS_ALPHAARG1 texture-stage state is the first alpha argument for the
	// stage, identified by by TA. The default argument is TA_TEXTURE. If no
	// texture is set for this stage, the default argument is TA_DIFFUSE.
	// Specify TA_TEMP to select a temporary register color for read or write.
	// TA_TEMP is supported if the PMISCCAPS_TSSARGTEMP device capability is
	// present. The default value for the register is (0.0, 0.0, 0.0, 0.0).
	TSS_ALPHAARG1 = 5
	// TSS_ALPHAARG2 texture-stage state is the second alpha argument for the
	// stage, identified by by TA. The default argument is TA_CURRENT. Specify
	// TA_TEMP to select a temporary register color for read or write. TA_TEMP
	// is supported if the PMISCCAPS_TSSARGTEMP device capability is present.
	// The default value for the register is (0.0, 0.0, 0.0, 0.0).
	TSS_ALPHAARG2 = 6
	// TSS_BUMPENVMAT00 texture-stage state is a floating-point value for the
	// [0][0] coefficient in a bump-mapping matrix. The default value is 0.0.
	TSS_BUMPENVMAT00 = 7
	// TSS_BUMPENVMAT01 texture-stage state is a floating-point value for the
	// [0][1] coefficient in a bump-mapping matrix. The default value is 0.0.
	TSS_BUMPENVMAT01 = 8
	// TSS_BUMPENVMAT10 texture-stage state is a floating-point value for the
	// [1][0] coefficient in a bump-mapping matrix. The default value is 0.0.
	TSS_BUMPENVMAT10 = 9
	// TSS_BUMPENVMAT11 texture-stage state is a floating-point value for the
	// [1][1] coefficient in a bump-mapping matrix. The default value is 0.0.
	TSS_BUMPENVMAT11 = 10
	// TSS_TEXCOORDINDEX is an index of the texture coordinate set to use with
	// this texture stage. You can specify up to eight sets of texture
	// coordinates per vertex. If a vertex does not include a set of texture
	// coordinates at the specified index, the system defaults to the u and v
	// coordinates (0,0).
	// When rendering using vertex shaders, each stage's texture coordinate
	// index must be set to its default value. The default index for each stage
	// is equal to the stage index. Set this state to the zero-based index of
	// the coordinate set for each vertex that this texture stage uses.
	// Additionally, applications can include, as logical OR with the index
	// being set, one of the constants to request that Direct3D automatically
	// generate the input texture coordinates for a texture transformation.
	// With the exception of TSS_TCI_PASSTHRU, which resolves to zero, if any
	// of the following values is included with the index being set, the system
	// uses the index strictly to determine texture wrapping mode. These flags
	// are most useful when performing environment mapping.
	TSS_TEXCOORDINDEX = 11
	// TSS_BUMPENVLSCALE is a floating-point scale value for bump-map
	// luminance. The default value is 0.0.
	TSS_BUMPENVLSCALE = 22
	// TSS_BUMPENVLOFFSET is a floating-point offset value for bump-map
	// luminance. The default value is 0.0.
	TSS_BUMPENVLOFFSET = 23
	// TSS_TEXTURETRANSFORMFLAGS is a member of the TEXTURETRANSFORMFLAGS
	// enumerated type that controls the transformation of texture coordinates
	// for this texture stage. The default value is TTFF_DISABLE.
	TSS_TEXTURETRANSFORMFLAGS = 24
	// TSS_COLORARG0 specifies settings for the third color operand for triadic
	// operations (multiply, add, and linearly interpolate), identified by TA.
	// This setting is supported if the TEXOPCAPS_MULTIPLYADD or TEXOPCAPS_LERP
	// device capabilities are present. The default argument is TA_CURRENT.
	// Specify TA_TEMP to select a temporary register color for read or write.
	// TA_TEMP is supported if the PMISCCAPS_TSSARGTEMP device capability is
	// present. The default value for the register is (0.0, 0.0, 0.0, 0.0).
	TSS_COLORARG0 = 26
	// TSS_ALPHAARG0 specifies settings for the alpha channel selector operand
	// for triadic operations (multiply, add, and linearly interpolate),
	// identified by TA. This setting is supported if the TEXOPCAPS_MULTIPLYADD
	// or TEXOPCAPS_LERP device capabilities are present. The default argument
	// is TA_CURRENT. Specify TA_TEMP to select a temporary register color for
	// read or write. TA_TEMP is supported if the PMISCCAPS_TSSARGTEMP device
	// capability is present. The default argument is (0.0, 0.0, 0.0, 0.0).
	TSS_ALPHAARG0 = 27
	// TSS_RESULTARG is a setting to select destination register for the result
	// of this stage, identified by TA. This value can be set to TA_CURRENT
	// (the default value) or to TA_TEMP, which is a single temporary register
	// that can be read into subsequent stages as an input argument. The final
	// color passed to the fog blender and frame buffer is taken from
	// TA_CURRENT, so the last active texture stage state must be set to write
	// to current. This setting is supported if the PMISCCAPS_TSSARGTEMP device
	// capability is present.
	TSS_RESULTARG = 28
	// TSS_CONSTANT is a per-stage constant color. To see if a device supports
	// a per-stage constant color, see the PMISCCAPS_PERSTAGECONSTANT constant
	// in PMISCCAPS. TSS_CONSTANT is used by TA_CONSTANT.
	TSS_CONSTANT = 32
)

// TEXTURETRANSFORMFLAGS defines texture coordinate transformation values.
type TEXTURETRANSFORMFLAGS uint32

const (
	// TTFF_DISABLE means texture coordinates are passed directly to the
	// rasterizer.
	TTFF_DISABLE = 0
	// TTFF_COUNT1 means the rasterizer should expect 1D texture coordinates.
	// This value is used by fixed function vertex processing; it should be set
	// to 0 when using a programmable vertex shader.
	TTFF_COUNT1 = 1
	// TTFF_COUNT2 means the rasterizer should expect 2D texture coordinates.
	// This value is used by fixed function vertex processing; it should be set
	// to 0 when using a programmable vertex shader.
	TTFF_COUNT2 = 2
	// TTFF_COUNT3 means the rasterizer should expect 3D texture coordinates.
	// This value is used by fixed function vertex processing; it should be set
	// to 0 when using a programmable vertex shader.
	TTFF_COUNT3 = 3
	// TTFF_COUNT4 means the rasterizer should expect 4D texture coordinates.
	// This value is used by fixed function vertex processing; it should be set
	// to 0 when using a programmable vertex shader.
	TTFF_COUNT4 = 4
	// TTFF_PROJECTED is a flag honored by the fixed function pixel pipeline,
	// as well as the programmable pixel pipeline in versions ps_1_1 to ps_1_3.
	// When texture projection is enabled for a texture stage, all four
	// floating point values must be written to the corresponding texture
	// register. Each texture coordinate is divided by the last element before
	// being passed to the rasterizer. For example, if this flag is specified
	// with the TTFF_COUNT3 flag, the first and second texture coordinates are
	// divided by the third coordinate before being passed to the rasterizer.
	TTFF_PROJECTED = 256
)

// TRANSFORMSTATETYPE defines constants that describe transformation state
// values.
type TRANSFORMSTATETYPE uint32

const (
	// TS_VIEW identifies the transformation matrix being set as the view
	// transformation matrix. The default value is nil (the identity matrix).
	TS_VIEW = 2
	// TS_PROJECTION identifies the transformation matrix being set as the
	// projection transformation matrix. The default value is nil (the identity
	// matrix).
	TS_PROJECTION = 3
	// TS_TEXTURE0 identifies the transformation matrix being set for texture
	// stage 0.
	TS_TEXTURE0 = 16
	// TS_TEXTURE1 identifies the transformation matrix being set for texture
	// stage 1.
	TS_TEXTURE1 = 17
	// TS_TEXTURE2 identifies the transformation matrix being set for texture
	// stage 2.
	TS_TEXTURE2 = 18
	// TS_TEXTURE3 identifies the transformation matrix being set for texture
	// stage 3.
	TS_TEXTURE3 = 19
	// TS_TEXTURE4 identifies the transformation matrix being set for texture
	// stage 4.
	TS_TEXTURE4 = 20
	// TS_TEXTURE5 identifies the transformation matrix being set for texture
	// stage 5.
	TS_TEXTURE5 = 21
	// TS_TEXTURE6 identifies the transformation matrix being set for texture
	// stage 6.
	TS_TEXTURE6 = 22
	// TS_TEXTURE7 identifies the transformation matrix being set for texture
	// stage 7.
	TS_TEXTURE7 = 23
)

// VERTEXBLENDFLAGS defines flags used to control the number or matrices that
// the system applies when performing multimatrix vertex blending.
type VERTEXBLENDFLAGS uint16

const (
	// VBF_DISABLE disables vertex blending; apply only the world matrix set by
	// the TS_WORLDMATRIX macro, where the index value for the transformation
	// state is 0.
	VBF_DISABLE = 0
	// VBF_1WEIGHTS enables vertex blending between the two matrices set by the
	// TS_WORLDMATRIX macro, where the index value for the transformation
	// states are 0 and 1.
	VBF_1WEIGHTS = 1
	// VBF_2WEIGHTS enables vertex blending between the three matrices set by
	// the TS_WORLDMATRIX macro, where the index value for the transformation
	// states are 0, 1, and 2.
	VBF_2WEIGHTS = 2
	// VBF_3WEIGHTS enables vertex blending between the four matrices set by
	// the TS_WORLDMATRIX macro, where the index value for the transformation
	// states are 0, 1, 2, and 3.
	VBF_3WEIGHTS = 3
	// VBF_TWEENING means vertex blending is done by using the value assigned
	// to RS_TWEENFACTOR.
	VBF_TWEENING = 255
	// VBF_0WEIGHTS uses a single matrix with a weight of 1.0.
	VBF_0WEIGHTS = 256
)

// ZBUFFERTYPE defines constants that describe depth-buffer formats.
type ZBUFFERTYPE uint32

const (
	// ZB_FALSE disables depth buffering.
	ZB_FALSE = 0
	// ZB_TRUE enables z-buffering.
	ZB_TRUE = 1
	// ZB_USEW enables w-buffering.
	ZB_USEW = 2
)
