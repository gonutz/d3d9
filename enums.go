package d3d9

//#include "d3d9wrapper.h"
import "C"

type BACKBUFFER_TYPE uint32

const (
	BACKBUFFER_TYPE_MONO  BACKBUFFER_TYPE = 0
	BACKBUFFER_TYPE_LEFT                  = 1
	BACKBUFFER_TYPE_RIGHT                 = 2
)

type BASISTYPE uint32

const (
	BASIS_BEZIER      BASISTYPE = 0
	BASIS_BSPLINE               = 1
	BASIS_CATMULL_ROM           = 2
)

type BLEND uint32

const (
	BLEND_ZERO            BLEND = 1
	BLEND_ONE                   = 2
	BLEND_SRCCOLOR              = 3
	BLEND_INVSRCCOLOR           = 4
	BLEND_SRCALPHA              = 5
	BLEND_INVSRCALPHA           = 6
	BLEND_DESTALPHA             = 7
	BLEND_INVDESTALPHA          = 8
	BLEND_DESTCOLOR             = 9
	BLEND_INVDESTCOLOR          = 10
	BLEND_SRCALPHASAT           = 11
	BLEND_BOTHSRCALPHA          = 12
	BLEND_BOTHINVSRCALPHA       = 13
	BLEND_BLENDFACTOR           = 14
	BLEND_INVBLENDFACTOR        = 15
)

type BLENDOP uint32

const (
	BLENDOP_ADD         BLENDOP = 1
	BLENDOP_SUBTRACT            = 2
	BLENDOP_REVSUBTRACT         = 3
	BLENDOP_MIN                 = 4
	BLENDOP_MAX                 = 5
)

type CMPFUNC uint32

const (
	CMP_NEVER        CMPFUNC = 1
	CMP_LESS                 = 2
	CMP_EQUAL                = 3
	CMP_LESSEQUAL            = 4
	CMP_GREATER              = 5
	CMP_NOTEQUAL             = 6
	CMP_GREATEREQUAL         = 7
	CMP_ALWAYS               = 8
)

type CUBEMAP_FACES uint32

const (
	CUBEMAP_FACE_POSITIVE_X CUBEMAP_FACES = 0
	CUBEMAP_FACE_NEGATIVE_X               = 1
	CUBEMAP_FACE_POSITIVE_Y               = 2
	CUBEMAP_FACE_NEGATIVE_Y               = 3
	CUBEMAP_FACE_POSITIVE_Z               = 4
	CUBEMAP_FACE_NEGATIVE_Z               = 5
)

type CULL uint32

const (
	CULL_NONE CULL = 1
	CULL_CW        = 2
	CULL_CCW       = 3
)

type DEBUGMONITORTOKENS uint32

const (
	DMT_ENABLE  DEBUGMONITORTOKENS = 0
	DMT_DISABLE                    = 1
)

type DECLMETHOD byte

const (
	DECLMETHOD_DEFAULT          DECLMETHOD = 0
	DECLMETHOD_PARTIALU                    = 1
	DECLMETHOD_PARTIALV                    = 2
	DECLMETHOD_CROSSUV                     = 3
	DECLMETHOD_UV                          = 4
	DECLMETHOD_LOOKUP                      = 5
	DECLMETHOD_LOOKUPPRESAMPLED            = 6
)

type DECLTYPE byte

const (
	DECLTYPE_FLOAT1    DECLTYPE = 0
	DECLTYPE_FLOAT2             = 1
	DECLTYPE_FLOAT3             = 2
	DECLTYPE_FLOAT4             = 3
	DECLTYPE_D3DCOLOR           = 4
	DECLTYPE_UBYTE4             = 5
	DECLTYPE_SHORT2             = 6
	DECLTYPE_SHORT4             = 7
	DECLTYPE_UBYTE4N            = 8
	DECLTYPE_SHORT2N            = 9
	DECLTYPE_SHORT4N            = 10
	DECLTYPE_USHORT2N           = 11
	DECLTYPE_USHORT4N           = 12
	DECLTYPE_UDEC3              = 13
	DECLTYPE_DEC3N              = 14
	DECLTYPE_FLOAT16_2          = 15
	DECLTYPE_FLOAT16_4          = 16
	DECLTYPE_UNUSED             = 17
)

type DECLUSAGE byte

const (
	DECLUSAGE_POSITION     DECLUSAGE = 0
	DECLUSAGE_BLENDWEIGHT            = 1
	DECLUSAGE_BLENDINDICES           = 2
	DECLUSAGE_NORMAL                 = 3
	DECLUSAGE_PSIZE                  = 4
	DECLUSAGE_TEXCOORD               = 5
	DECLUSAGE_TANGENT                = 6
	DECLUSAGE_BINORMAL               = 7
	DECLUSAGE_TESSFACTOR             = 8
	DECLUSAGE_POSITIONT              = 9
	DECLUSAGE_COLOR                  = 10
	DECLUSAGE_FOG                    = 11
	DECLUSAGE_DEPTH                  = 12
	DECLUSAGE_SAMPLE                 = 13
)

type DEGREETYPE uint32

const (
	DEGREE_LINEAR    DEGREETYPE = 1
	DEGREE_QUADRATIC            = 2
	DEGREE_CUBIC                = 3
	DEGREE_QUINTIC              = 5
)

type DEVTYPE uint32

const (
	DEVTYPE_HAL     DEVTYPE = 1
	DEVTYPE_REF             = 2
	DEVTYPE_SW              = 3
	DEVTYPE_NULLREF         = 4
)

type FILLMODE uint32

const (
	FILL_POINT     FILLMODE = 1
	FILL_WIREFRAME          = 2
	FILL_SOLID              = 3
)

type FOGMODE uint32

const (
	FOG_NONE   FOGMODE = 0
	FOG_EXP            = 1
	FOG_EXP2           = 2
	FOG_LINEAR         = 3
)

type FORMAT uint32

const (
	FMT_UNKNOWN       FORMAT = 0
	FMT_R8G8B8               = 20
	FMT_A8R8G8B8             = 21
	FMT_X8R8G8B8             = 22
	FMT_R5G6B5               = 23
	FMT_X1R5G5B5             = 24
	FMT_A1R5G5B5             = 25
	FMT_A4R4G4B4             = 26
	FMT_R3G3B2               = 27
	FMT_A8                   = 28
	FMT_A8R3G3B2             = 29
	FMT_X4R4G4B4             = 30
	FMT_A2B10G10R10          = 31
	FMT_A8B8G8R8             = 32
	FMT_X8B8G8R8             = 33
	FMT_G16R16               = 34
	FMT_A2R10G10B10          = 35
	FMT_A16B16G16R16         = 36
	FMT_A8P8                 = 40
	FMT_P8                   = 41
	FMT_L8                   = 50
	FMT_A8L8                 = 51
	FMT_A4L4                 = 52
	FMT_V8U8                 = 60
	FMT_L6V5U5               = 61
	FMT_X8L8V8U8             = 62
	FMT_Q8W8V8U8             = 63
	FMT_V16U16               = 64
	FMT_A2W10V10U10          = 67
	FMT_UYVY                 = 1498831189
	FMT_R8G8_B8G8            = 1195525970
	FMT_YUY2                 = 844715353
	FMT_G8R8_G8B8            = 1111970375
	FMT_DXT1                 = 827611204
	FMT_DXT2                 = 844388420
	FMT_DXT3                 = 861165636
	FMT_DXT4                 = 877942852
	FMT_DXT5                 = 894720068
	FMT_D16_LOCKABLE         = 70
	FMT_D32                  = 71
	FMT_D15S1                = 73
	FMT_D24S8                = 75
	FMT_D24X8                = 77
	FMT_D24X4S4              = 79
	FMT_D16                  = 80
	FMT_L16                  = 81
	FMT_D32F_LOCKABLE        = 82
	FMT_D24FS8               = 83
	FMT_VERTEXDATA           = 100
	FMT_INDEX16              = 101
	FMT_INDEX32              = 102
	FMT_Q16W16V16U16         = 110
	FMT_MULTI2_ARGB8         = 827606349
	FMT_R16F                 = 111
	FMT_G16R16F              = 112
	FMT_A16B16G16R16F        = 113
	FMT_R32F                 = 114
	FMT_G32R32F              = 115
	FMT_A32B32G32R32F        = 116
	FMT_CxV8U8               = 117
)

type LIGHTTYPE uint32

const (
	LIGHT_POINT       LIGHTTYPE = 1
	LIGHT_SPOT                  = 2
	LIGHT_DIRECTIONAL           = 3
)

type MATERIALCOLORSOURCE uint32

const (
	MCS_MATERIAL MATERIALCOLORSOURCE = 0
	MCS_COLOR1                       = 1
	MCS_COLOR2                       = 2
)

type MULTISAMPLE_TYPE uint32

const (
	MULTISAMPLE_NONE        MULTISAMPLE_TYPE = 0
	MULTISAMPLE_NONMASKABLE                  = 1
	MULTISAMPLE_2_SAMPLES                    = 2
	MULTISAMPLE_3_SAMPLES                    = 3
	MULTISAMPLE_4_SAMPLES                    = 4
	MULTISAMPLE_5_SAMPLES                    = 5
	MULTISAMPLE_6_SAMPLES                    = 6
	MULTISAMPLE_7_SAMPLES                    = 7
	MULTISAMPLE_8_SAMPLES                    = 8
	MULTISAMPLE_9_SAMPLES                    = 9
	MULTISAMPLE_10_SAMPLES                   = 10
	MULTISAMPLE_11_SAMPLES                   = 11
	MULTISAMPLE_12_SAMPLES                   = 12
	MULTISAMPLE_13_SAMPLES                   = 13
	MULTISAMPLE_14_SAMPLES                   = 14
	MULTISAMPLE_15_SAMPLES                   = 15
	MULTISAMPLE_16_SAMPLES                   = 16
)

type PATCHEDGESTYLE uint32

const (
	PATCHEDGE_DISCRETE   PATCHEDGESTYLE = 0
	PATCHEDGE_CONTINUOUS                = 1
)

type POOL uint32

const (
	POOL_DEFAULT   POOL = 0
	POOL_MANAGED        = 1
	POOL_SYSTEMMEM      = 2
	POOL_SCRATCH        = 3
)

type PRIMITIVETYPE uint32

const (
	PT_POINTLIST     PRIMITIVETYPE = 0
	PT_LINELIST                    = 1
	PT_LINESTRIP                   = 2
	PT_TRIANGLELIST                = 3
	PT_TRIANGLESTRIP               = 4
	PT_TRIANGLEFAN                 = 5
)

type QUERYTYPE byte

const (
	QUERYTYPE_VCACHE            QUERYTYPE = 4
	QUERYTYPE_RESOURCEMANAGER             = 5
	QUERYTYPE_VERTEXSTATS                 = 6
	QUERYTYPE_EVENT                       = 8
	QUERYTYPE_OCCLUSION                   = 9
	QUERYTYPE_TIMESTAMP                   = 10
	QUERYTYPE_TIMESTAMPDISJOINT           = 11
	QUERYTYPE_TIMESTAMPFREQ               = 12
	QUERYTYPE_PIPELINETIMINGS             = 13
	QUERYTYPE_INTERFACETIMINGS            = 14
	QUERYTYPE_VERTEXTIMINGS               = 15
	QUERYTYPE_PIXELTIMINGS                = 16
	QUERYTYPE_BANDWIDTHTIMINGS            = 17
	QUERYTYPE_CACHEUTILIZATION            = 18
)

type RENDERSTATETYPE uint32

const (
	RS_ZENABLE                    RENDERSTATETYPE = 7
	RS_FILLMODE                                   = 8
	RS_SHADEMODE                                  = 9
	RS_ZWRITEENABLE                               = 14
	RS_ALPHATESTENABLE                            = 15
	RS_LASTPIXEL                                  = 16
	RS_SRCBLEND                                   = 19
	RS_DESTBLEND                                  = 20
	RS_CULLMODE                                   = 22
	RS_ZFUNC                                      = 23
	RS_ALPHAREF                                   = 24
	RS_ALPHAFUNC                                  = 25
	RS_DITHERENABLE                               = 26
	RS_ALPHABLENDENABLE                           = 27
	RS_FOGENABLE                                  = 28
	RS_SPECULARENABLE                             = 29
	RS_FOGCOLOR                                   = 34
	RS_FOGTABLEMODE                               = 35
	RS_FOGSTART                                   = 36
	RS_FOGEND                                     = 37
	RS_FOGDENSITY                                 = 38
	RS_RANGEFOGENABLE                             = 48
	RS_STENCILENABLE                              = 52
	RS_STENCILFAIL                                = 53
	RS_STENCILZFAIL                               = 54
	RS_STENCILPASS                                = 55
	RS_STENCILFUNC                                = 56
	RS_STENCILREF                                 = 57
	RS_STENCILMASK                                = 58
	RS_STENCILWRITEMASK                           = 59
	RS_TEXTUREFACTOR                              = 60
	RS_WRAP0                                      = 128
	RS_WRAP1                                      = 129
	RS_WRAP2                                      = 130
	RS_WRAP3                                      = 131
	RS_WRAP4                                      = 132
	RS_WRAP5                                      = 133
	RS_WRAP6                                      = 134
	RS_WRAP7                                      = 135
	RS_CLIPPING                                   = 136
	RS_LIGHTING                                   = 137
	RS_AMBIENT                                    = 139
	RS_FOGVERTEXMODE                              = 140
	RS_COLORVERTEX                                = 141
	RS_LOCALVIEWER                                = 142
	RS_NORMALIZENORMALS                           = 143
	RS_DIFFUSEMATERIALSOURCE                      = 145
	RS_SPECULARMATERIALSOURCE                     = 146
	RS_AMBIENTMATERIALSOURCE                      = 147
	RS_EMISSIVEMATERIALSOURCE                     = 148
	RS_VERTEXBLEND                                = 151
	RS_CLIPPLANEENABLE                            = 152
	RS_POINTSIZE                                  = 154
	RS_POINTSIZE_MIN                              = 155
	RS_POINTSPRITEENABLE                          = 156
	RS_POINTSCALEENABLE                           = 157
	RS_POINTSCALE_A                               = 158
	RS_POINTSCALE_B                               = 159
	RS_POINTSCALE_C                               = 160
	RS_MULTISAMPLEANTIALIAS                       = 161
	RS_MULTISAMPLEMASK                            = 162
	RS_PATCHEDGESTYLE                             = 163
	RS_DEBUGMONITORTOKEN                          = 165
	RS_POINTSIZE_MAX                              = 166
	RS_INDEXEDVERTEXBLENDENABLE                   = 167
	RS_COLORWRITEENABLE                           = 168
	RS_TWEENFACTOR                                = 170
	RS_BLENDOP                                    = 171
	RS_POSITIONDEGREE                             = 172
	RS_NORMALDEGREE                               = 173
	RS_SCISSORTESTENABLE                          = 174
	RS_SLOPESCALEDEPTHBIAS                        = 175
	RS_ANTIALIASEDLINEENABLE                      = 176
	RS_MINTESSELLATIONLEVEL                       = 178
	RS_MAXTESSELLATIONLEVEL                       = 179
	RS_ADAPTIVETESS_X                             = 180
	RS_ADAPTIVETESS_Y                             = 181
	RS_ADAPTIVETESS_Z                             = 182
	RS_ADAPTIVETESS_W                             = 183
	RS_ENABLEADAPTIVETESSELLATION                 = 184
	RS_TWOSIDEDSTENCILMODE                        = 185
	RS_CCW_STENCILFAIL                            = 186
	RS_CCW_STENCILZFAIL                           = 187
	RS_CCW_STENCILPASS                            = 188
	RS_CCW_STENCILFUNC                            = 189
	RS_COLORWRITEENABLE1                          = 190
	RS_COLORWRITEENABLE2                          = 191
	RS_COLORWRITEENABLE3                          = 192
	RS_BLENDFACTOR                                = 193
	RS_SRGBWRITEENABLE                            = 194
	RS_DEPTHBIAS                                  = 195
	RS_WRAP8                                      = 198
	RS_WRAP9                                      = 199
	RS_WRAP10                                     = 200
	RS_WRAP11                                     = 201
	RS_WRAP12                                     = 202
	RS_WRAP13                                     = 203
	RS_WRAP14                                     = 204
	RS_WRAP15                                     = 205
	RS_SEPARATEALPHABLENDENABLE                   = 206
	RS_SRCBLENDALPHA                              = 207
	RS_DESTBLENDALPHA                             = 208
	RS_BLENDOPALPHA                               = 209
)

type RESOURCETYPE uint32

const (
	RTYPE_SURFACE       RESOURCETYPE = 1
	RTYPE_VOLUME                     = 2
	RTYPE_TEXTURE                    = 3
	RTYPE_VOLUMETEXTURE              = 4
	RTYPE_CUBETEXTURE                = 5
	RTYPE_VERTEXBUFFER               = 6
	RTYPE_INDEXBUFFER                = 7
)

type SAMPLER_TEXTURE_TYPE uint32

const (
	STT_UNKNOWN SAMPLER_TEXTURE_TYPE = 0
	STT_2D                           = 268435456
	STT_CUBE                         = 402653184
	STT_VOLUME                       = 536870912
)

type SAMPLERSTATETYPE uint32

const (
	SAMP_ADDRESSU      SAMPLERSTATETYPE = 1
	SAMP_ADDRESSV                       = 2
	SAMP_ADDRESSW                       = 3
	SAMP_BORDERCOLOR                    = 4
	SAMP_MAGFILTER                      = 5
	SAMP_MINFILTER                      = 6
	SAMP_MIPFILTER                      = 7
	SAMP_MIPMAPLODBIAS                  = 8
	SAMP_MAXMIPLEVEL                    = 9
	SAMP_MAXANISOTROPY                  = 10
	SAMP_SRGBTEXTURE                    = 11
	SAMP_ELEMENTINDEX                   = 12
	SAMP_DMAPOFFSET                     = 13
)

type SHADEMODE uint32

const (
	SHADE_FLAT    SHADEMODE = 1
	SHADE_GOURAUD           = 2
	SHADE_PHONG             = 3
)

type STATEBLOCKTYPE uint32

const (
	SBT_ALL         STATEBLOCKTYPE = 1
	SBT_PIXELSTATE                 = 2
	SBT_VERTEXSTATE                = 3
)

type STENCILOP uint32

const (
	STENCILOP_KEEP    STENCILOP = 1
	STENCILOP_ZERO              = 2
	STENCILOP_REPLACE           = 3
	STENCILOP_INCRSAT           = 4
	STENCILOP_DECRSAT           = 5
	STENCILOP_INVERT            = 6
	STENCILOP_INCR              = 7
	STENCILOP_DECR              = 8
)

type SWAPEFFECT uint32

const (
	SWAPEFFECT_DISCARD SWAPEFFECT = 1
	SWAPEFFECT_FLIP               = 2
	SWAPEFFECT_COPY               = 3
)

type TEXTUREADDRESS uint32

const (
	TADDRESS_WRAP       TEXTUREADDRESS = 1
	TADDRESS_MIRROR                    = 2
	TADDRESS_CLAMP                     = 3
	TADDRESS_BORDER                    = 4
	TADDRESS_MIRRORONCE                = 5
)

type TEXTUREFILTERTYPE uint32

const (
	TEXF_NONE          TEXTUREFILTERTYPE = 0
	TEXF_POINT                           = 1
	TEXF_LINEAR                          = 2
	TEXF_ANISOTROPIC                     = 3
	TEXF_PYRAMIDALQUAD                   = 6
	TEXF_GAUSSIANQUAD                    = 7
)

type TEXTUREOP uint32

const (
	TOP_DISABLE                   TEXTUREOP = 1
	TOP_SELECTARG1                          = 2
	TOP_SELECTARG2                          = 3
	TOP_MODULATE                            = 4
	TOP_MODULATE2X                          = 5
	TOP_MODULATE4X                          = 6
	TOP_ADD                                 = 7
	TOP_ADDSIGNED                           = 8
	TOP_ADDSIGNED2X                         = 9
	TOP_SUBTRACT                            = 10
	TOP_ADDSMOOTH                           = 11
	TOP_BLENDDIFFUSEALPHA                   = 12
	TOP_BLENDTEXTUREALPHA                   = 13
	TOP_BLENDFACTORALPHA                    = 14
	TOP_BLENDTEXTUREALPHAPM                 = 15
	TOP_BLENDCURRENTALPHA                   = 16
	TOP_PREMODULATE                         = 17
	TOP_MODULATEALPHA_ADDCOLOR              = 18
	TOP_MODULATECOLOR_ADDALPHA              = 19
	TOP_MODULATEINVALPHA_ADDCOLOR           = 20
	TOP_MODULATEINVCOLOR_ADDALPHA           = 21
	TOP_BUMPENVMAP                          = 22
	TOP_BUMPENVMAPLUMINANCE                 = 23
	TOP_DOTPRODUCT3                         = 24
	TOP_MULTIPLYADD                         = 25
	TOP_LERP                                = 26
)

type TEXTURESTAGESTATETYPE uint32

const (
	TSS_COLOROP               TEXTURESTAGESTATETYPE = 1
	TSS_COLORARG1                                   = 2
	TSS_COLORARG2                                   = 3
	TSS_ALPHAOP                                     = 4
	TSS_ALPHAARG1                                   = 5
	TSS_ALPHAARG2                                   = 6
	TSS_BUMPENVMAT00                                = 7
	TSS_BUMPENVMAT01                                = 8
	TSS_BUMPENVMAT10                                = 9
	TSS_BUMPENVMAT11                                = 10
	TSS_TEXCOORDINDEX                               = 11
	TSS_BUMPENVLSCALE                               = 22
	TSS_BUMPENVLOFFSET                              = 23
	TSS_TEXTURETRANSFORMFLAGS                       = 24
	TSS_COLORARG0                                   = 26
	TSS_ALPHAARG0                                   = 27
	TSS_RESULTARG                                   = 28
	TSS_CONSTANT                                    = 32
)

type TEXTURETRANSFORMFLAGS uint32

const (
	TTFF_DISABLE   TEXTURETRANSFORMFLAGS = 0
	TTFF_COUNT1                          = 1
	TTFF_COUNT2                          = 2
	TTFF_COUNT3                          = 3
	TTFF_COUNT4                          = 4
	TTFF_PROJECTED                       = 256
)

type TRANSFORMSTATETYPE uint32

const (
	TS_VIEW       TRANSFORMSTATETYPE = 2
	TS_PROJECTION                    = 3
	TS_TEXTURE0                      = 16
	TS_TEXTURE1                      = 17
	TS_TEXTURE2                      = 18
	TS_TEXTURE3                      = 19
	TS_TEXTURE4                      = 20
	TS_TEXTURE5                      = 21
	TS_TEXTURE6                      = 22
	TS_TEXTURE7                      = 23
)

type VERTEXBLENDFLAGS uint16

const (
	VBF_DISABLE  VERTEXBLENDFLAGS = 0
	VBF_1WEIGHTS                  = 1
	VBF_2WEIGHTS                  = 2
	VBF_3WEIGHTS                  = 3
	VBF_TWEENING                  = 255
	VBF_0WEIGHTS                  = 256
)

type ZBUFFERTYPE uint32

const (
	ZB_FALSE ZBUFFERTYPE = 0
	ZB_TRUE              = 1
	ZB_USEW              = 2
)
