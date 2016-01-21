#undef INTERFACE
/*
 * Copyright (C) 2002-2003 Jason Edmeades
 *                         Raphael Junqueira
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with this library; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301, USA
 */

#ifndef _D3D9_H_
#define _D3D9_H_

#ifndef DIRECT3D_VERSION
#define DIRECT3D_VERSION  0x0900
#endif

#include <stdlib.h>

#define COM_NO_WINDOWS_H
#include <objbase.h>
#include <windows.h>

// d3d9types.h
/*
 * Copyright (C) 2002-2003 Jason Edmeades 
 * Copyright (C) 2002-2003 Raphael Junqueira
 * Copyright (C) 2005 Oliver Stieber
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with this library; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301, USA
 */

#ifndef __WINE_D3D9TYPES_H
#define __WINE_D3D9TYPES_H

/*****************************************************************************
 * Direct 3D v9 #defines
 */
#define D3DCLEAR_TARGET   0x00000001
#define D3DCLEAR_ZBUFFER  0x00000002
#define D3DCLEAR_STENCIL  0x00000004

#define D3DCLIPPLANE0 (1 << 0)
#define D3DCLIPPLANE1 (1 << 1)
#define D3DCLIPPLANE2 (1 << 2)
#define D3DCLIPPLANE3 (1 << 3)
#define D3DCLIPPLANE4 (1 << 4)
#define D3DCLIPPLANE5 (1 << 5)

#define D3DCOLOR_ARGB(a,r,g,b)       ((D3DCOLOR)((((a)&0xff)<<24)|(((r)&0xff)<<16)|(((g)&0xff)<<8)|((b)&0xff)))
#define D3DCOLOR_COLORVALUE(r,g,b,a) D3DCOLOR_RGBA((DWORD)((r)*255.f),(DWORD)((g)*255.f),(DWORD)((b)*255.f),(DWORD)((a)*255.f))
#define D3DCOLOR_RGBA(r,g,b,a)       D3DCOLOR_ARGB(a,r,g,b)
#define D3DCOLOR_XRGB(r,g,b)         D3DCOLOR_ARGB(0xff,r,g,b)
#define D3DCOLOR_XYUV(y,u,v)         D3DCOLOR_ARGB(0xFF,y,u,v)
#define D3DCOLOR_AYUV(a,y,u,v)       D3DCOLOR_ARGB(a,y,u,v)

#define D3DCS_LEFT                   0x001
#define D3DCS_RIGHT                  0x002
#define D3DCS_TOP                    0x004
#define D3DCS_BOTTOM                 0x008
#define D3DCS_FRONT                  0x010
#define D3DCS_BACK                   0x020
#define D3DCS_PLANE0                 0x040
#define D3DCS_PLANE1                 0x080
#define D3DCS_PLANE2                 0x100
#define D3DCS_PLANE3                 0x200
#define D3DCS_PLANE4                 0x400
#define D3DCS_PLANE5                 0x800
#define D3DCS_ALL                    0xFFF

#define D3DFVF_TEXTUREFORMAT1 3
#define D3DFVF_TEXTUREFORMAT2 0
#define D3DFVF_TEXTUREFORMAT3 1
#define D3DFVF_TEXTUREFORMAT4 2
#define D3DFVF_TEXCOORDSIZE1(CoordIndex) (D3DFVF_TEXTUREFORMAT1 << (CoordIndex*2 + 16))
#define D3DFVF_TEXCOORDSIZE2(CoordIndex) (D3DFVF_TEXTUREFORMAT2)
#define D3DFVF_TEXCOORDSIZE3(CoordIndex) (D3DFVF_TEXTUREFORMAT3 << (CoordIndex*2 + 16))
#define D3DFVF_TEXCOORDSIZE4(CoordIndex) (D3DFVF_TEXTUREFORMAT4 << (CoordIndex*2 + 16))

#define D3DLOCK_READONLY           0x0010
#define D3DLOCK_NOSYSLOCK          0x0800
#define D3DLOCK_NOOVERWRITE        0x1000
#define D3DLOCK_DISCARD            0x2000
#define D3DLOCK_DONOTWAIT          0x4000
#define D3DLOCK_NO_DIRTY_UPDATE    0x8000

#define D3DMAXUSERCLIPPLANES       32
#define D3DCLIPPLANE0              (1 << 0)
#define D3DCLIPPLANE1              (1 << 1)
#define D3DCLIPPLANE2              (1 << 2)
#define D3DCLIPPLANE3              (1 << 3)
#define D3DCLIPPLANE4              (1 << 4)
#define D3DCLIPPLANE5              (1 << 5)


#define D3DRENDERSTATE_WRAPBIAS    128U

/* MSDN has this in d3d9caps.h, but it should be here */
#define D3DTSS_TCI_PASSTHRU                       0x00000
#define D3DTSS_TCI_CAMERASPACENORMAL              0x10000
#define D3DTSS_TCI_CAMERASPACEPOSITION            0x20000
#define D3DTSS_TCI_CAMERASPACEREFLECTIONVECTOR    0x30000
#define D3DTSS_TCI_SPHEREMAP                      0x40000


#define D3DTS_WORLD  D3DTS_WORLDMATRIX(0)
#define D3DTS_WORLD1 D3DTS_WORLDMATRIX(1)
#define D3DTS_WORLD2 D3DTS_WORLDMATRIX(2)
#define D3DTS_WORLD3 D3DTS_WORLDMATRIX(3)
#define D3DTS_WORLDMATRIX(index) (D3DTRANSFORMSTATETYPE)(index + 256)

#define D3DUSAGE_RENDERTARGET       0x00000001
#define D3DUSAGE_DEPTHSTENCIL       0x00000002
#define D3DUSAGE_WRITEONLY          0x00000008
#define D3DUSAGE_SOFTWAREPROCESSING 0x00000010
#define D3DUSAGE_DONOTCLIP          0x00000020
#define D3DUSAGE_POINTS             0x00000040
#define D3DUSAGE_RTPATCHES          0x00000080
#define D3DUSAGE_NPATCHES           0x00000100
#define D3DUSAGE_DYNAMIC            0x00000200
#define D3DUSAGE_AUTOGENMIPMAP      0x00000400
#define D3DUSAGE_DMAP               0x00004000

#define D3DUSAGE_QUERY_FILTER                   0x00020000
#define D3DUSAGE_QUERY_LEGACYBUMPMAP            0x00008000
#define D3DUSAGE_QUERY_POSTPIXELSHADER_BLENDING 0x00080000
#define D3DUSAGE_QUERY_SRGBREAD                 0x00010000
#define D3DUSAGE_QUERY_SRGBWRITE                0x00040000
#define D3DUSAGE_QUERY_VERTEXTEXTURE            0x00100000
#define D3DUSAGE_QUERY_WRAPANDMIP               0x00200000

#define D3DWRAP_U        1
#define D3DWRAP_V        2
#define D3DWRAP_W        4
#define D3DWRAPCOORD_0   1
#define D3DWRAPCOORD_1   2
#define D3DWRAPCOORD_2   4
#define D3DWRAPCOORD_3   8

#define MAX_DEVICE_IDENTIFIER_STRING        512

#define D3DFVF_RESERVED0           0x0001
#define D3DFVF_POSITION_MASK       0x400E
#define D3DFVF_XYZ                 0x0002
#define D3DFVF_XYZRHW              0x0004
#define D3DFVF_XYZB1               0x0006
#define D3DFVF_XYZB2               0x0008
#define D3DFVF_XYZB3               0x000a
#define D3DFVF_XYZB4               0x000c
#define D3DFVF_XYZB5               0x000e
#define D3DFVF_XYZW                0x4002
#define D3DFVF_NORMAL              0x0010
#define D3DFVF_PSIZE               0x0020
#define D3DFVF_DIFFUSE             0x0040
#define D3DFVF_SPECULAR            0x0080
#define D3DFVF_TEXCOUNT_MASK       0x0f00
#define D3DFVF_TEXCOUNT_SHIFT           8
#define D3DFVF_TEX0                0x0000
#define D3DFVF_TEX1                0x0100
#define D3DFVF_TEX2                0x0200
#define D3DFVF_TEX3                0x0300
#define D3DFVF_TEX4                0x0400
#define D3DFVF_TEX5                0x0500
#define D3DFVF_TEX6                0x0600
#define D3DFVF_TEX7                0x0700
#define D3DFVF_TEX8                0x0800
#define D3DFVF_LASTBETA_UBYTE4     0x1000
#define D3DFVF_LASTBETA_D3DCOLOR   0x8000
#define D3DFVF_RESERVED2           0x6000

#define D3DTA_SELECTMASK        0x0000000f
#define D3DTA_DIFFUSE           0x00000000
#define D3DTA_CURRENT           0x00000001
#define D3DTA_TEXTURE           0x00000002
#define D3DTA_TFACTOR           0x00000003
#define D3DTA_SPECULAR          0x00000004
#define D3DTA_TEMP              0x00000005
#define D3DTA_CONSTANT          0x00000006
#define D3DTA_COMPLEMENT        0x00000010
#define D3DTA_ALPHAREPLICATE    0x00000020

#define D3DCOLORWRITEENABLE_RED   (1<<0)
#define D3DCOLORWRITEENABLE_GREEN (1<<1)
#define D3DCOLORWRITEENABLE_BLUE  (1<<2)
#define D3DCOLORWRITEENABLE_ALPHA (1<<3)

#define D3DPV_DONOTCOPYDATA         (1 << 0)

#define D3DSTREAMSOURCE_INDEXEDDATA  (1 << 30)
#define D3DSTREAMSOURCE_INSTANCEDATA (2 << 30)

#define D3D_MAX_SIMULTANEOUS_RENDERTARGETS 4

#define MAXD3DDECLLENGTH         64 /* +end marker */
#define MAXD3DDECLMETHOD         D3DDECLMETHOD_LOOKUPPRESAMPLED
#define MAXD3DDECLTYPE           D3DDECLTYPE_UNUSED
#define MAXD3DDECLUSAGE          D3DDECLUSAGE_SAMPLE
#define MAXD3DDECLUSAGEINDEX     15

#define D3DDMAPSAMPLER 256
#define D3DVERTEXTEXTURESAMPLER0 (D3DDMAPSAMPLER+1)
#define D3DVERTEXTEXTURESAMPLER1 (D3DDMAPSAMPLER+2)
#define D3DVERTEXTEXTURESAMPLER2 (D3DDMAPSAMPLER+3)
#define D3DVERTEXTEXTURESAMPLER3 (D3DDMAPSAMPLER+4)

#ifndef MAKEFOURCC
#define MAKEFOURCC(ch0, ch1, ch2, ch3)  \
    ((DWORD)(BYTE)(ch0) | ((DWORD)(BYTE)(ch1) << 8) |  \
    ((DWORD)(BYTE)(ch2) << 16) | ((DWORD)(BYTE)(ch3) << 24 ))
#endif

/* Constants used by D3DPRESENT_PARAMETERS. when creating a device or swapchain */

#define D3DPRESENTFLAG_LOCKABLE_BACKBUFFER  0x00000001 /* Create a lockable backbuffer */
#define D3DPRESENTFLAG_DISCARD_DEPTHSTENCIL 0x00000002 /* Discard Z buffer */
#define D3DPRESENTFLAG_DEVICECLIP           0x00000004 /* Clip the window blited into the client area 2k + xp only */
#define D3DPRESENTFLAG_VIDEO                0x00000010 /* backbuffer 'may' contain video data */
#define D3DPRESENTFLAG_NOAUTOROTATE         0x00000020 /* d3d9ex, ignore display rotation */
#define D3DPRESENTFLAG_UNPRUNEDMODE         0x00000040 /* d3d9ex, specify invalid display modes */

#define D3DPRESENT_BACK_BUFFERS_MAX         3
#define D3DPRESENT_RATE_DEFAULT             0x00000000

/**************************** 
 * Vertex Shaders Declaration
 */

typedef enum _D3DDECLUSAGE {
  D3DDECLUSAGE_POSITION     = 0,
  D3DDECLUSAGE_BLENDWEIGHT  = 1,
  D3DDECLUSAGE_BLENDINDICES = 2,
  D3DDECLUSAGE_NORMAL       = 3,      
  D3DDECLUSAGE_PSIZE        = 4,       
  D3DDECLUSAGE_TEXCOORD     = 5,    
  D3DDECLUSAGE_TANGENT      = 6,     
  D3DDECLUSAGE_BINORMAL     = 7,    
  D3DDECLUSAGE_TESSFACTOR   = 8,  
  D3DDECLUSAGE_POSITIONT    = 9,   
  D3DDECLUSAGE_COLOR        = 10,       
  D3DDECLUSAGE_FOG          = 11,        
  D3DDECLUSAGE_DEPTH        = 12,      
  D3DDECLUSAGE_SAMPLE       = 13     
} D3DDECLUSAGE;

#define D3DMAXDECLUSAGE         D3DDECLUSAGE_SAMPLE
#define D3DMAXDECLUSAGEINDEX    15
#define D3DMAXDECLLENGTH        18
#define D3DMAXDECLUSAGE_DX8     D3DDECLUSAGE_TEXCOORD

typedef enum _D3DDECLMETHOD {
  D3DDECLMETHOD_DEFAULT          = 0,
  D3DDECLMETHOD_PARTIALU         = 1,
  D3DDECLMETHOD_PARTIALV         = 2,
  D3DDECLMETHOD_CROSSUV          = 3,
  D3DDECLMETHOD_UV               = 4,
  D3DDECLMETHOD_LOOKUP           = 5,
  D3DDECLMETHOD_LOOKUPPRESAMPLED = 6
} D3DDECLMETHOD;


#define D3DMAXDECLMETHOD        D3DDECLMETHOD_LOOKUPPRESAMPLED

typedef enum _D3DDECLTYPE {
  D3DDECLTYPE_FLOAT1    =  0,
  D3DDECLTYPE_FLOAT2    =  1,
  D3DDECLTYPE_FLOAT3    =  2,
  D3DDECLTYPE_FLOAT4    =  3,
  D3DDECLTYPE_D3DCOLOR  =  4,
  D3DDECLTYPE_UBYTE4    =  5,
  D3DDECLTYPE_SHORT2    =  6,
  D3DDECLTYPE_SHORT4    =  7,
  /* VS 2.0 */
  D3DDECLTYPE_UBYTE4N   =  8,
  D3DDECLTYPE_SHORT2N   =  9,
  D3DDECLTYPE_SHORT4N   = 10,
  D3DDECLTYPE_USHORT2N  = 11,
  D3DDECLTYPE_USHORT4N  = 12,
  D3DDECLTYPE_UDEC3     = 13,
  D3DDECLTYPE_DEC3N     = 14,
  D3DDECLTYPE_FLOAT16_2 = 15,
  D3DDECLTYPE_FLOAT16_4 = 16,
  D3DDECLTYPE_UNUSED    = 17,
} D3DDECLTYPE;

#define D3DMAXDECLTYPE          D3DDECLTYPE_UNUSED

typedef struct _D3DVERTEXELEMENT9 {
  WORD    Stream;
  WORD    Offset;
  BYTE    Type;
  BYTE    Method;
  BYTE    Usage;
  BYTE    UsageIndex;
} D3DVERTEXELEMENT9, *LPD3DVERTEXELEMENT9;


typedef enum _D3DQUERYTYPE {
    D3DQUERYTYPE_VCACHE = 4,
    D3DQUERYTYPE_RESOURCEMANAGER = 5,
    D3DQUERYTYPE_VERTEXSTATS = 6,
    D3DQUERYTYPE_EVENT = 8,
    D3DQUERYTYPE_OCCLUSION = 9,
    D3DQUERYTYPE_TIMESTAMP = 10,
    D3DQUERYTYPE_TIMESTAMPDISJOINT = 11,
    D3DQUERYTYPE_TIMESTAMPFREQ = 12,
    D3DQUERYTYPE_PIPELINETIMINGS = 13,
    D3DQUERYTYPE_INTERFACETIMINGS = 14,
    D3DQUERYTYPE_VERTEXTIMINGS = 15,
    D3DQUERYTYPE_PIXELTIMINGS = 16,
    D3DQUERYTYPE_BANDWIDTHTIMINGS = 17,
    D3DQUERYTYPE_CACHEUTILIZATION = 18
} D3DQUERYTYPE;

#define D3DISSUE_BEGIN   (1 << 1)
#define D3DISSUE_END     (1 << 0)
#define D3DGETDATA_FLUSH (1 << 0)


#define D3DDECL_END() {0xFF,0,D3DDECLTYPE_UNUSED,0,0,0}
#define D3DDP_MAXTEXCOORD   8


#define D3DVSD_MAKETOKENTYPE(TokenType) \
  ((TokenType << D3DVSD_TOKENTYPESHIFT) & D3DVSD_TOKENTYPEMASK)

#define D3DVSD_CONST(ConstantAddress, Count) \
  (D3DVSD_MAKETOKENTYPE(D3DVSD_TOKEN_CONSTMEM) | ((Count) << D3DVSD_CONSTCOUNTSHIFT) | (ConstantAddress))

#define D3DVSD_END() 0xFFFFFFFF

#define D3DVSD_NOP() 0x00000000

#define D3DVSD_REG(VertexRegister, Type) \
  (D3DVSD_MAKETOKENTYPE(D3DVSD_TOKEN_STREAMDATA) | ((Type) << D3DVSD_DATATYPESHIFT) | (VertexRegister))

#define D3DVSD_SKIP(Count) \
  (D3DVSD_MAKETOKENTYPE(D3DVSD_TOKEN_STREAMDATA) | 0x10000000 | ((Count) << D3DVSD_SKIPCOUNTSHIFT))

#define D3DVSD_STREAM(StreamNumber) \
  (D3DVSD_MAKETOKENTYPE(D3DVSD_TOKEN_STREAM) | (StreamNumber))

#define D3DVSD_STREAM_TESS() \
  (D3DVSD_MAKETOKENTYPE(D3DVSD_TOKEN_STREAM) | (D3DVSD_STREAMTESSMASK))

#define D3DVSD_TESSNORMAL(RegisterIn, RegisterOut) \
  (D3DVSD_MAKETOKENTYPE(D3DVSD_TOKEN_TESSELLATOR) | ((RegisterIn) << D3DVSD_VERTEXREGINSHIFT) | ((0x02) << D3DVSD_DATATYPESHIFT) | (RegisterOut))

#define D3DVSD_TESSUV(Register) \
  (D3DVSD_MAKETOKENTYPE(D3DVSD_TOKEN_TESSELLATOR) | 0x10000000 | ((0x01) << D3DVSD_DATATYPESHIFT) | (Register))


/********************************
 * Pixel/Vertex Shaders Functions
 */

/** Maximum number of supported texture coordinates sets operation */
#define D3DDP_MAXTEXCOORD   8

/** opcode token mask */
#define D3DSI_OPCODE_MASK       0x0000FFFF
#define D3DSI_INSTLENGTH_MASK   0x0F000000
#define D3DSI_INSTLENGTH_SHIFT  24

/** opcodes types for PS and VS */
typedef enum _D3DSHADER_INSTRUCTION_OPCODE_TYPE {
  D3DSIO_NOP          =  0,
  D3DSIO_MOV          =  1,
  D3DSIO_ADD          =  2,
  D3DSIO_SUB          =  3,
  D3DSIO_MAD          =  4,
  D3DSIO_MUL          =  5,
  D3DSIO_RCP          =  6,
  D3DSIO_RSQ          =  7,
  D3DSIO_DP3          =  8,
  D3DSIO_DP4          =  9,
  D3DSIO_MIN          = 10,
  D3DSIO_MAX          = 11,
  D3DSIO_SLT          = 12,
  D3DSIO_SGE          = 13,
  D3DSIO_EXP          = 14,
  D3DSIO_LOG          = 15,
  D3DSIO_LIT          = 16,
  D3DSIO_DST          = 17,
  D3DSIO_LRP          = 18,
  D3DSIO_FRC          = 19,
  D3DSIO_M4x4         = 20,
  D3DSIO_M4x3         = 21,
  D3DSIO_M3x4         = 22,
  D3DSIO_M3x3         = 23,
  D3DSIO_M3x2         = 24,
  D3DSIO_CALL         = 25,
  D3DSIO_CALLNZ       = 26,
  D3DSIO_LOOP         = 27,
  D3DSIO_RET          = 28,
  D3DSIO_ENDLOOP      = 29,
  D3DSIO_LABEL        = 30,
  D3DSIO_DCL          = 31,
  D3DSIO_POW          = 32,
  D3DSIO_CRS          = 33,
  D3DSIO_SGN          = 34,
  D3DSIO_ABS          = 35,
  D3DSIO_NRM          = 36,
  D3DSIO_SINCOS       = 37,
  D3DSIO_REP          = 38,
  D3DSIO_ENDREP       = 39,
  D3DSIO_IF           = 40,
  D3DSIO_IFC          = 41,
  D3DSIO_ELSE         = 42,
  D3DSIO_ENDIF        = 43,
  D3DSIO_BREAK        = 44,
  D3DSIO_BREAKC       = 45,
  D3DSIO_MOVA         = 46,
  D3DSIO_DEFB         = 47,
  D3DSIO_DEFI         = 48,

  D3DSIO_TEXCOORD     = 64,
  D3DSIO_TEXKILL      = 65,
  D3DSIO_TEX          = 66,
  D3DSIO_TEXBEM       = 67,
  D3DSIO_TEXBEML      = 68,
  D3DSIO_TEXREG2AR    = 69,
  D3DSIO_TEXREG2GB    = 70,
  D3DSIO_TEXM3x2PAD   = 71,
  D3DSIO_TEXM3x2TEX   = 72,
  D3DSIO_TEXM3x3PAD   = 73,
  D3DSIO_TEXM3x3TEX   = 74,
  D3DSIO_TEXM3x3DIFF  = 75,
  D3DSIO_TEXM3x3SPEC  = 76,
  D3DSIO_TEXM3x3VSPEC = 77,
  D3DSIO_EXPP         = 78,
  D3DSIO_LOGP         = 79,
  D3DSIO_CND          = 80,
  D3DSIO_DEF          = 81,
  D3DSIO_TEXREG2RGB   = 82,
  D3DSIO_TEXDP3TEX    = 83,
  D3DSIO_TEXM3x2DEPTH = 84,
  D3DSIO_TEXDP3       = 85,
  D3DSIO_TEXM3x3      = 86,
  D3DSIO_TEXDEPTH     = 87,
  D3DSIO_CMP          = 88,
  D3DSIO_BEM          = 89,
  D3DSIO_DP2ADD       = 90,
  D3DSIO_DSX          = 91,
  D3DSIO_DSY          = 92,
  D3DSIO_TEXLDD       = 93,
  D3DSIO_SETP         = 94,
  D3DSIO_TEXLDL       = 95,
  D3DSIO_BREAKP       = 96,
  
  D3DSIO_PHASE        = 0xFFFD,
  D3DSIO_COMMENT      = 0xFFFE,
  D3DSIO_END          = 0XFFFF,

  D3DSIO_FORCE_DWORD  = 0X7FFFFFFF /** for 32-bit alignment */
} D3DSHADER_INSTRUCTION_OPCODE_TYPE;

#define D3DSINCOSCONST1   -1.5500992e-006f, -2.1701389e-005f,  0.0026041667f, 0.00026041668f
#define D3DSINCOSCONST2   -0.020833334f,    -0.12500000f,      1.0f,          0.50000000f

#define D3DSHADER_INSTRUCTION_PREDICATED    (1 << 28)

#define D3DSI_TEXLD_PROJECT 0x00010000
#define D3DSI_TEXLD_BIAS    0x00020000

/** for parallelism */
#define D3DSI_COISSUE 0x40000000

#define D3DSP_DCL_USAGE_SHIFT 0
#define D3DSP_DCL_USAGE_MASK  0x0000000f

#define D3DSP_DCL_USAGEINDEX_SHIFT 16
#define D3DSP_DCL_USAGEINDEX_MASK  0x000f0000

#define D3DSP_TEXTURETYPE_SHIFT 27
#define D3DSP_TEXTURETYPE_MASK  0x78000000

typedef enum _D3DSAMPLER_TEXTURE_TYPE {
  D3DSTT_UNKNOWN      = 0 << D3DSP_TEXTURETYPE_SHIFT,
  D3DSTT_1D           = 1 << D3DSP_TEXTURETYPE_SHIFT,
  D3DSTT_2D           = 2 << D3DSP_TEXTURETYPE_SHIFT,
  D3DSTT_CUBE         = 3 << D3DSP_TEXTURETYPE_SHIFT,
  D3DSTT_VOLUME       = 4 << D3DSP_TEXTURETYPE_SHIFT,

  D3DSTT_FORCE_DWORD  = 0x7FFFFFFF
} D3DSAMPLER_TEXTURE_TYPE;

#define D3DSP_REGNUM_MASK       0x000007FF

/** destination parameter modifiers (.xyzw) */
#define D3DSP_WRITEMASK_0       0x00010000 /* .x r */
#define D3DSP_WRITEMASK_1       0x00020000 /* .y g */
#define D3DSP_WRITEMASK_2       0x00040000 /* .z b */
#define D3DSP_WRITEMASK_3       0x00080000 /* .w a */
#define D3DSP_WRITEMASK_ALL     0x000F0000 /* all */

#define D3DSP_DSTMOD_SHIFT      20
#define D3DSP_DSTMOD_MASK       (0xF << D3DSP_DSTMOD_SHIFT)

typedef enum _D3DSHADER_PARAM_DSTMOD_TYPE {
  D3DSPDM_NONE             = 0 << D3DSP_DSTMOD_SHIFT,
  D3DSPDM_SATURATE         = 1 << D3DSP_DSTMOD_SHIFT,
  D3DSPDM_PARTIALPRECISION = 2 << D3DSP_DSTMOD_SHIFT,
  D3DSPDM_MSAMPCENTROID    = 4 << D3DSP_DSTMOD_SHIFT,

  D3DSPDM_FORCE_DWORD  = 0x7FFFFFFF
} D3DSHADER_PARAM_DSTMOD_TYPE;

/** destination param */
#define D3DSP_DSTSHIFT_SHIFT     24
#define D3DSP_DSTSHIFT_MASK      (0xF << D3DSP_DSTSHIFT_SHIFT)

/** destination/source reg type */
#define D3DSP_REGTYPE_SHIFT      28
#define D3DSP_REGTYPE_SHIFT2     8
#define D3DSP_REGTYPE_MASK       (0x7 << D3DSP_REGTYPE_SHIFT)
#define D3DSP_REGTYPE_MASK2      0x00001800

typedef enum _D3DSHADER_PARAM_REGISTER_TYPE {
  D3DSPR_TEMP         =  0, 
  D3DSPR_INPUT        =  1,
  D3DSPR_CONST        =  2,
  D3DSPR_ADDR         =  3,
  D3DSPR_TEXTURE      =  3,
  D3DSPR_RASTOUT      =  4,
  D3DSPR_ATTROUT      =  5,
  D3DSPR_TEXCRDOUT    =  6,
  D3DSPR_OUTPUT       =  6,
  D3DSPR_CONSTINT     =  7,
  D3DSPR_COLOROUT     =  8,
  D3DSPR_DEPTHOUT     =  9,
  D3DSPR_SAMPLER      = 10,
  D3DSPR_CONST2       = 11,
  D3DSPR_CONST3       = 12,
  D3DSPR_CONST4       = 13,
  D3DSPR_CONSTBOOL    = 14,
  D3DSPR_LOOP         = 15,
  D3DSPR_TEMPFLOAT16  = 16,
  D3DSPR_MISCTYPE     = 17,
  D3DSPR_LABEL        = 18,
  D3DSPR_PREDICATE    = 19,

  D3DSPR_FORCE_DWORD  = 0x7FFFFFFF
} D3DSHADER_PARAM_REGISTER_TYPE;

typedef enum _D3DSHADER_MISCTYPE_OFFSETS {
    D3DSMO_POSITION  = 0,
    D3DSMO_FACE      = 1
} D3DSHADER_MISCTYPE_OFFSETS;

typedef enum _D3DVS_RASTOUT_OFFSETS {
  D3DSRO_POSITION     = 0,
  D3DSRO_FOG          = 1,
  D3DSRO_POINT_SIZE   = 2,

  D3DSRO_FORCE_DWORD  = 0x7FFFFFFF
} D3DVS_RASTOUT_OFFSETS;

#define D3DVS_ADDRESSMODE_SHIFT  13
#define D3DVS_ADDRESSMODE_MASK   (0x1 << D3DVS_ADDRESSMODE_SHIFT)

typedef enum _D3DVS_ADDRESSMODE_TYPE {
  D3DVS_ADDRMODE_ABSOLUTE     = 0 << D3DVS_ADDRESSMODE_SHIFT,
  D3DVS_ADDRMODE_RELATIVE     = 1 << D3DVS_ADDRESSMODE_SHIFT,

  D3DVS_ADDRMODE_FORCE_DWORD  = 0x7FFFFFFF
} D3DVS_ADDRESSMODE_TYPE;

#define D3DSHADER_ADDRESSMODE_SHIFT 13
#define D3DSHADER_ADDRESSMODE_MASK  (1 << D3DSHADER_ADDRESSMODE_SHIFT)

typedef enum _D3DSHADER_ADDRESSMODE_TYPE {
  D3DSHADER_ADDRMODE_ABSOLUTE    = 0 << D3DSHADER_ADDRESSMODE_SHIFT,
  D3DSHADER_ADDRMODE_RELATIVE    = 1 << D3DSHADER_ADDRESSMODE_SHIFT,

  D3DSHADER_ADDRMODE_FORCE_DWORD = 0x7FFFFFFF
} D3DSHADER_ADDRESSMODE_TYPE;


#define D3DVS_SWIZZLE_SHIFT      16
#define D3DVS_SWIZZLE_MASK       (0xFF << D3DVS_SWIZZLE_SHIFT)

#define D3DSP_SWIZZLE_SHIFT      16
#define D3DSP_SWIZZLE_MASK       (0xFF << D3DSP_SWIZZLE_SHIFT)

#define D3DVS_X_X       (0 << D3DVS_SWIZZLE_SHIFT)
#define D3DVS_X_Y       (1 << D3DVS_SWIZZLE_SHIFT)
#define D3DVS_X_Z       (2 << D3DVS_SWIZZLE_SHIFT)
#define D3DVS_X_W       (3 << D3DVS_SWIZZLE_SHIFT)

#define D3DVS_Y_X       (0 << (D3DVS_SWIZZLE_SHIFT + 2))
#define D3DVS_Y_Y       (1 << (D3DVS_SWIZZLE_SHIFT + 2))
#define D3DVS_Y_Z       (2 << (D3DVS_SWIZZLE_SHIFT + 2))
#define D3DVS_Y_W       (3 << (D3DVS_SWIZZLE_SHIFT + 2))

#define D3DVS_Z_X       (0 << (D3DVS_SWIZZLE_SHIFT + 4))
#define D3DVS_Z_Y       (1 << (D3DVS_SWIZZLE_SHIFT + 4))
#define D3DVS_Z_Z       (2 << (D3DVS_SWIZZLE_SHIFT + 4))
#define D3DVS_Z_W       (3 << (D3DVS_SWIZZLE_SHIFT + 4))

#define D3DVS_W_X       (0 << (D3DVS_SWIZZLE_SHIFT + 6))
#define D3DVS_W_Y       (1 << (D3DVS_SWIZZLE_SHIFT + 6))
#define D3DVS_W_Z       (2 << (D3DVS_SWIZZLE_SHIFT + 6))
#define D3DVS_W_W       (3 << (D3DVS_SWIZZLE_SHIFT + 6))

#define D3DVS_NOSWIZZLE (D3DVS_X_X | D3DVS_Y_Y | D3DVS_Z_Z | D3DVS_W_W)

#define D3DSP_NOSWIZZLE \
    ((0 << (D3DSP_SWIZZLE_SHIFT + 0)) | (1 << (D3DSP_SWIZZLE_SHIFT + 2)) | (2 << (D3DSP_SWIZZLE_SHIFT + 4)) | (3 << (D3DSP_SWIZZLE_SHIFT + 6)))

#define D3DSP_SRCMOD_SHIFT      24
#define D3DSP_SRCMOD_MASK       (0xF << D3DSP_SRCMOD_SHIFT)

typedef enum _D3DSHADER_PARAM_SRCMOD_TYPE {
  D3DSPSM_NONE         =  0 << D3DSP_SRCMOD_SHIFT,
  D3DSPSM_NEG          =  1 << D3DSP_SRCMOD_SHIFT,
  D3DSPSM_BIAS         =  2 << D3DSP_SRCMOD_SHIFT,
  D3DSPSM_BIASNEG      =  3 << D3DSP_SRCMOD_SHIFT,
  D3DSPSM_SIGN         =  4 << D3DSP_SRCMOD_SHIFT,
  D3DSPSM_SIGNNEG      =  5 << D3DSP_SRCMOD_SHIFT,
  D3DSPSM_COMP         =  6 << D3DSP_SRCMOD_SHIFT,
  D3DSPSM_X2           =  7 << D3DSP_SRCMOD_SHIFT,
  D3DSPSM_X2NEG        =  8 << D3DSP_SRCMOD_SHIFT,
  D3DSPSM_DZ           =  9 << D3DSP_SRCMOD_SHIFT,
  D3DSPSM_DW           = 10 << D3DSP_SRCMOD_SHIFT,
  D3DSPSM_ABS          = 11 << D3DSP_SRCMOD_SHIFT,
  D3DSPSM_ABSNEG       = 12 << D3DSP_SRCMOD_SHIFT,
  D3DSPSM_NOT          = 13 << D3DSP_SRCMOD_SHIFT,

  D3DSPSM_FORCE_DWORD  = 0x7FFFFFFF
} D3DSHADER_PARAM_SRCMOD_TYPE;

#define D3DPS_VERSION(major, minor) (0xFFFF0000 | ((major) << 8) | (minor))
#define D3DVS_VERSION(major, minor) (0xFFFE0000 | ((major) << 8) | (minor))
#define D3DSHADER_VERSION_MAJOR(version) (((version) >> 8) & 0xFF)
#define D3DSHADER_VERSION_MINOR(version) (((version) >> 0) & 0xFF)

#define D3DSI_COMMENTSIZE_SHIFT 16
#define D3DSI_COMMENTSIZE_MASK (0x7FFF << D3DSI_COMMENTSIZE_SHIFT)

#define D3DSHADER_COMMENT(commentSize) \
  ((((commentSize) << D3DSI_COMMENTSIZE_SHIFT) & D3DSI_COMMENTSIZE_MASK) | D3DSIO_COMMENT)

#define D3DPS_END() 0x0000FFFF
#define D3DVS_END() 0x0000FFFF


/*****************************************************************************
 * Direct 3D v8 enumerated types
 */
typedef enum _D3DBACKBUFFER_TYPE {
    D3DBACKBUFFER_TYPE_MONO         = 0,
    D3DBACKBUFFER_TYPE_LEFT         = 1,
    D3DBACKBUFFER_TYPE_RIGHT        = 2,

    D3DBACKBUFFER_TYPE_FORCE_DWORD  = 0x7fffffff
} D3DBACKBUFFER_TYPE;

#define D3DPRESENT_BACK_BUFFER_MAX 3

typedef enum _D3DBASISTYPE {
   D3DBASIS_BEZIER        = 0,
   D3DBASIS_BSPLINE       = 1,
   D3DBASIS_CATMULL_ROM   = 2, /* In D3D8 this used to be D3DBASIS_INTERPOLATE */

   D3DBASIS_FORCE_DWORD   = 0x7fffffff
} D3DBASISTYPE;

typedef enum _D3DBLEND {
    D3DBLEND_ZERO               =  1,
    D3DBLEND_ONE                =  2,
    D3DBLEND_SRCCOLOR           =  3,
    D3DBLEND_INVSRCCOLOR        =  4,
    D3DBLEND_SRCALPHA           =  5,
    D3DBLEND_INVSRCALPHA        =  6,
    D3DBLEND_DESTALPHA          =  7,
    D3DBLEND_INVDESTALPHA       =  8,
    D3DBLEND_DESTCOLOR          =  9,
    D3DBLEND_INVDESTCOLOR       = 10,
    D3DBLEND_SRCALPHASAT        = 11,
    D3DBLEND_BOTHSRCALPHA       = 12,
    D3DBLEND_BOTHINVSRCALPHA    = 13,
    D3DBLEND_BLENDFACTOR        = 14,
    D3DBLEND_INVBLENDFACTOR     = 15,
    D3DBLEND_FORCE_DWORD        = 0x7fffffff
} D3DBLEND;

typedef enum _D3DBLENDOP {
    D3DBLENDOP_ADD              = 1,
    D3DBLENDOP_SUBTRACT         = 2,
    D3DBLENDOP_REVSUBTRACT      = 3,
    D3DBLENDOP_MIN              = 4,
    D3DBLENDOP_MAX              = 5,

    D3DBLENDOP_FORCE_DWORD      = 0x7fffffff
} D3DBLENDOP;

typedef enum _D3DCMPFUNC {
    D3DCMP_NEVER                = 1,
    D3DCMP_LESS                 = 2,
    D3DCMP_EQUAL                = 3,
    D3DCMP_LESSEQUAL            = 4,
    D3DCMP_GREATER              = 5,
    D3DCMP_NOTEQUAL             = 6,
    D3DCMP_GREATEREQUAL         = 7,
    D3DCMP_ALWAYS               = 8,

    D3DCMP_FORCE_DWORD          = 0x7fffffff
} D3DCMPFUNC;

typedef enum _D3DCUBEMAP_FACES {
    D3DCUBEMAP_FACE_POSITIVE_X     = 0,
    D3DCUBEMAP_FACE_NEGATIVE_X     = 1,
    D3DCUBEMAP_FACE_POSITIVE_Y     = 2,
    D3DCUBEMAP_FACE_NEGATIVE_Y     = 3,
    D3DCUBEMAP_FACE_POSITIVE_Z     = 4,
    D3DCUBEMAP_FACE_NEGATIVE_Z     = 5,

    D3DCUBEMAP_FACE_FORCE_DWORD    = 0xffffffff
} D3DCUBEMAP_FACES;

typedef enum _D3DCULL {
    D3DCULL_NONE                = 1,
    D3DCULL_CW                  = 2,
    D3DCULL_CCW                 = 3,

    D3DCULL_FORCE_DWORD         = 0x7fffffff
} D3DCULL;

typedef enum _D3DDEBUGMONITORTOKENS {
    D3DDMT_ENABLE          = 0,
    D3DDMT_DISABLE         = 1,

    D3DDMT_FORCE_DWORD     = 0x7fffffff
} D3DDEBUGMONITORTOKENS;

typedef enum _D3DDEGREETYPE {
    D3DDEGREE_LINEAR      = 1,
    D3DDEGREE_QUADRATIC   = 2,
    D3DDEGREE_CUBIC       = 3,
    D3DDEGREE_QUINTIC     = 5,
    
    D3DDEGREE_FORCE_DWORD   = 0x7fffffff
} D3DDEGREETYPE;

typedef enum _D3DDEVTYPE {
    D3DDEVTYPE_HAL         = 1,
    D3DDEVTYPE_REF         = 2,
    D3DDEVTYPE_SW          = 3,
    D3DDEVTYPE_NULLREF     = 4,

    D3DDEVTYPE_FORCE_DWORD = 0xffffffff
} D3DDEVTYPE;

typedef enum _D3DFILLMODE {
    D3DFILL_POINT               = 1,
    D3DFILL_WIREFRAME           = 2,
    D3DFILL_SOLID               = 3,

    D3DFILL_FORCE_DWORD         = 0x7fffffff
} D3DFILLMODE;

typedef enum _D3DFOGMODE {
    D3DFOG_NONE                 = 0,
    D3DFOG_EXP                  = 1,
    D3DFOG_EXP2                 = 2,
    D3DFOG_LINEAR               = 3,

    D3DFOG_FORCE_DWORD          = 0x7fffffff
} D3DFOGMODE;

typedef enum _D3DFORMAT {
    D3DFMT_UNKNOWN              =   0,

    D3DFMT_R8G8B8               =  20,
    D3DFMT_A8R8G8B8             =  21,
    D3DFMT_X8R8G8B8             =  22,
    D3DFMT_R5G6B5               =  23,
    D3DFMT_X1R5G5B5             =  24,
    D3DFMT_A1R5G5B5             =  25,
    D3DFMT_A4R4G4B4             =  26,
    D3DFMT_R3G3B2               =  27,
    D3DFMT_A8                   =  28,
    D3DFMT_A8R3G3B2             =  29,
    D3DFMT_X4R4G4B4             =  30,
    D3DFMT_A2B10G10R10          =  31,
    D3DFMT_A8B8G8R8             =  32,
    D3DFMT_X8B8G8R8             =  33,
    D3DFMT_G16R16               =  34,
    D3DFMT_A2R10G10B10          =  35,
    D3DFMT_A16B16G16R16         =  36,
  

    D3DFMT_A8P8                 =  40,
    D3DFMT_P8                   =  41,

    D3DFMT_L8                   =  50,
    D3DFMT_A8L8                 =  51,
    D3DFMT_A4L4                 =  52,

    D3DFMT_V8U8                 =  60,
    D3DFMT_L6V5U5               =  61,
    D3DFMT_X8L8V8U8             =  62,
    D3DFMT_Q8W8V8U8             =  63,
    D3DFMT_V16U16               =  64,
    D3DFMT_A2W10V10U10          =  67,

    D3DFMT_UYVY                 =  MAKEFOURCC('U', 'Y', 'V', 'Y'),
    D3DFMT_YUY2                 =  MAKEFOURCC('Y', 'U', 'Y', '2'),
    D3DFMT_DXT1                 =  MAKEFOURCC('D', 'X', 'T', '1'),
    D3DFMT_DXT2                 =  MAKEFOURCC('D', 'X', 'T', '2'),
    D3DFMT_DXT3                 =  MAKEFOURCC('D', 'X', 'T', '3'),
    D3DFMT_DXT4                 =  MAKEFOURCC('D', 'X', 'T', '4'),
    D3DFMT_DXT5                 =  MAKEFOURCC('D', 'X', 'T', '5'),
    D3DFMT_MULTI2_ARGB8         =  MAKEFOURCC('M', 'E', 'T', '1'),
    D3DFMT_G8R8_G8B8            =  MAKEFOURCC('G', 'R', 'G', 'B'),
    D3DFMT_R8G8_B8G8            =  MAKEFOURCC('R', 'G', 'B', 'G'),

    D3DFMT_D16_LOCKABLE         =  70,
    D3DFMT_D32                  =  71,
    D3DFMT_D15S1                =  73,
    D3DFMT_D24S8                =  75,
    D3DFMT_D24X8                =  77,
    D3DFMT_D24X4S4              =  79,
    D3DFMT_D16                  =  80,
    D3DFMT_L16                  =  81,
    D3DFMT_D32F_LOCKABLE        =  82,
    D3DFMT_D24FS8               =  83,

    D3DFMT_VERTEXDATA           = 100,
    D3DFMT_INDEX16              = 101,
    D3DFMT_INDEX32              = 102,
    D3DFMT_Q16W16V16U16         = 110,
    /* Floating point formats */
    D3DFMT_R16F                 = 111,
    D3DFMT_G16R16F              = 112,
    D3DFMT_A16B16G16R16F        = 113,
    
    /* IEEE formats */
    D3DFMT_R32F                 = 114,
    D3DFMT_G32R32F              = 115,
    D3DFMT_A32B32G32R32F        = 116,
    
    D3DFMT_CxV8U8               = 117,


    D3DFMT_FORCE_DWORD          = 0xFFFFFFFF
} D3DFORMAT;

typedef enum _D3DLIGHTTYPE {
    D3DLIGHT_POINT          = 1,
    D3DLIGHT_SPOT           = 2,
    D3DLIGHT_DIRECTIONAL    = 3,

    D3DLIGHT_FORCE_DWORD    = 0x7fffffff
} D3DLIGHTTYPE;

typedef enum _D3DMATERIALCOLORSOURCE {
    D3DMCS_MATERIAL         = 0,
    D3DMCS_COLOR1           = 1,
    D3DMCS_COLOR2           = 2,

    D3DMCS_FORCE_DWORD      = 0x7fffffff
} D3DMATERIALCOLORSOURCE;

typedef enum _D3DMULTISAMPLE_TYPE {
    D3DMULTISAMPLE_NONE            =  0,
    D3DMULTISAMPLE_NONMASKABLE     =  1,
    D3DMULTISAMPLE_2_SAMPLES       =  2,
    D3DMULTISAMPLE_3_SAMPLES       =  3,
    D3DMULTISAMPLE_4_SAMPLES       =  4,
    D3DMULTISAMPLE_5_SAMPLES       =  5,
    D3DMULTISAMPLE_6_SAMPLES       =  6,
    D3DMULTISAMPLE_7_SAMPLES       =  7,
    D3DMULTISAMPLE_8_SAMPLES       =  8,
    D3DMULTISAMPLE_9_SAMPLES       =  9,
    D3DMULTISAMPLE_10_SAMPLES      = 10,
    D3DMULTISAMPLE_11_SAMPLES      = 11,
    D3DMULTISAMPLE_12_SAMPLES      = 12,
    D3DMULTISAMPLE_13_SAMPLES      = 13,
    D3DMULTISAMPLE_14_SAMPLES      = 14,
    D3DMULTISAMPLE_15_SAMPLES      = 15,
    D3DMULTISAMPLE_16_SAMPLES      = 16,

    D3DMULTISAMPLE_FORCE_DWORD     = 0x7fffffff
} D3DMULTISAMPLE_TYPE;

#if 0
typedef enum _D3DORDERTYPE {
   D3DORDER_LINEAR      = 1,
   D3DORDER_QUADRATIC   = 2,
   D3DORDER_CUBIC       = 3,
   D3DORDER_QUINTIC     = 5,

   D3DORDER_FORCE_DWORD = 0x7fffffff
} D3DORDERTYPE;
#endif
typedef enum _D3DPATCHEDGESTYLE {
   D3DPATCHEDGE_DISCRETE    = 0,
   D3DPATCHEDGE_CONTINUOUS  = 1,

   D3DPATCHEDGE_FORCE_DWORD = 0x7fffffff,
} D3DPATCHEDGESTYLE;

typedef enum _D3DPOOL {
    D3DPOOL_DEFAULT                 = 0,
    D3DPOOL_MANAGED                 = 1,
    D3DPOOL_SYSTEMMEM               = 2,
    D3DPOOL_SCRATCH                 = 3,

    D3DPOOL_FORCE_DWORD             = 0x7fffffff
} D3DPOOL;

typedef enum _D3DPRIMITIVETYPE {
    D3DPT_POINTLIST             = 1,
    D3DPT_LINELIST              = 2,
    D3DPT_LINESTRIP             = 3,
    D3DPT_TRIANGLELIST          = 4,
    D3DPT_TRIANGLESTRIP         = 5,
    D3DPT_TRIANGLEFAN           = 6,

    D3DPT_FORCE_DWORD           = 0x7fffffff
} D3DPRIMITIVETYPE;

typedef enum _D3DRENDERSTATETYPE {
    D3DRS_ZENABLE                   =   7,
    D3DRS_FILLMODE                  =   8,
    D3DRS_SHADEMODE                 =   9,
    D3DRS_ZWRITEENABLE              =  14,
    D3DRS_ALPHATESTENABLE           =  15,
    D3DRS_LASTPIXEL                 =  16,
    D3DRS_SRCBLEND                  =  19,
    D3DRS_DESTBLEND                 =  20,
    D3DRS_CULLMODE                  =  22,
    D3DRS_ZFUNC                     =  23,
    D3DRS_ALPHAREF                  =  24,
    D3DRS_ALPHAFUNC                 =  25,
    D3DRS_DITHERENABLE              =  26,
    D3DRS_ALPHABLENDENABLE          =  27,
    D3DRS_FOGENABLE                 =  28,
    D3DRS_SPECULARENABLE            =  29,
    D3DRS_FOGCOLOR                  =  34,
    D3DRS_FOGTABLEMODE              =  35,
    D3DRS_FOGSTART                  =  36,
    D3DRS_FOGEND                    =  37,
    D3DRS_FOGDENSITY                =  38,
    D3DRS_RANGEFOGENABLE            =  48,
    D3DRS_STENCILENABLE             =  52,
    D3DRS_STENCILFAIL               =  53,
    D3DRS_STENCILZFAIL              =  54,
    D3DRS_STENCILPASS               =  55,
    D3DRS_STENCILFUNC               =  56,
    D3DRS_STENCILREF                =  57,
    D3DRS_STENCILMASK               =  58,
    D3DRS_STENCILWRITEMASK          =  59,
    D3DRS_TEXTUREFACTOR             =  60,
    D3DRS_WRAP0                     = 128,
    D3DRS_WRAP1                     = 129,
    D3DRS_WRAP2                     = 130,
    D3DRS_WRAP3                     = 131,
    D3DRS_WRAP4                     = 132,
    D3DRS_WRAP5                     = 133,
    D3DRS_WRAP6                     = 134,
    D3DRS_WRAP7                     = 135,
    D3DRS_CLIPPING                  = 136,
    D3DRS_LIGHTING                  = 137,
    D3DRS_AMBIENT                   = 139,
    D3DRS_FOGVERTEXMODE             = 140,
    D3DRS_COLORVERTEX               = 141,
    D3DRS_LOCALVIEWER               = 142,
    D3DRS_NORMALIZENORMALS          = 143,
    D3DRS_DIFFUSEMATERIALSOURCE     = 145,
    D3DRS_SPECULARMATERIALSOURCE    = 146,
    D3DRS_AMBIENTMATERIALSOURCE     = 147,
    D3DRS_EMISSIVEMATERIALSOURCE    = 148,
    D3DRS_VERTEXBLEND               = 151,
    D3DRS_CLIPPLANEENABLE           = 152,
    D3DRS_POINTSIZE                 = 154,
    D3DRS_POINTSIZE_MIN             = 155,
    D3DRS_POINTSPRITEENABLE         = 156,
    D3DRS_POINTSCALEENABLE          = 157,
    D3DRS_POINTSCALE_A              = 158,
    D3DRS_POINTSCALE_B              = 159,
    D3DRS_POINTSCALE_C              = 160,
    D3DRS_MULTISAMPLEANTIALIAS      = 161,
    D3DRS_MULTISAMPLEMASK           = 162,
    D3DRS_PATCHEDGESTYLE            = 163,
    D3DRS_DEBUGMONITORTOKEN         = 165,
    D3DRS_POINTSIZE_MAX             = 166,
    D3DRS_INDEXEDVERTEXBLENDENABLE  = 167,
    D3DRS_COLORWRITEENABLE          = 168,
    D3DRS_TWEENFACTOR               = 170,
    D3DRS_BLENDOP                   = 171,
    D3DRS_POSITIONDEGREE            = 172,
    D3DRS_NORMALDEGREE              = 173,
    D3DRS_SCISSORTESTENABLE         = 174,
    D3DRS_SLOPESCALEDEPTHBIAS       = 175,
    D3DRS_ANTIALIASEDLINEENABLE     = 176,
    D3DRS_MINTESSELLATIONLEVEL      = 178,
    D3DRS_MAXTESSELLATIONLEVEL      = 179,
    D3DRS_ADAPTIVETESS_X            = 180,
    D3DRS_ADAPTIVETESS_Y            = 181,
    D3DRS_ADAPTIVETESS_Z            = 182,
    D3DRS_ADAPTIVETESS_W            = 183,
    D3DRS_ENABLEADAPTIVETESSELLATION= 184,
    D3DRS_TWOSIDEDSTENCILMODE       = 185,
    D3DRS_CCW_STENCILFAIL           = 186,
    D3DRS_CCW_STENCILZFAIL          = 187,
    D3DRS_CCW_STENCILPASS           = 188,
    D3DRS_CCW_STENCILFUNC           = 189,
    D3DRS_COLORWRITEENABLE1         = 190,
    D3DRS_COLORWRITEENABLE2         = 191,
    D3DRS_COLORWRITEENABLE3         = 192,
    D3DRS_BLENDFACTOR               = 193,
    D3DRS_SRGBWRITEENABLE           = 194,
    D3DRS_DEPTHBIAS                 = 195,
    D3DRS_WRAP8                     = 198,
    D3DRS_WRAP9                     = 199,
    D3DRS_WRAP10                    = 200,
    D3DRS_WRAP11                    = 201,
    D3DRS_WRAP12                    = 202,
    D3DRS_WRAP13                    = 203,
    D3DRS_WRAP14                    = 204,
    D3DRS_WRAP15                    = 205,
    D3DRS_SEPARATEALPHABLENDENABLE  = 206,
    D3DRS_SRCBLENDALPHA             = 207,
    D3DRS_DESTBLENDALPHA            = 208,
    D3DRS_BLENDOPALPHA              = 209,

    D3DRS_FORCE_DWORD               = 0x7fffffff
} D3DRENDERSTATETYPE;

typedef enum _D3DRESOURCETYPE {
    D3DRTYPE_SURFACE                =  1,
    D3DRTYPE_VOLUME                 =  2,
    D3DRTYPE_TEXTURE                =  3,
    D3DRTYPE_VOLUMETEXTURE          =  4,
    D3DRTYPE_CUBETEXTURE            =  5,
    D3DRTYPE_VERTEXBUFFER           =  6,
    D3DRTYPE_INDEXBUFFER            =  7,

    D3DRTYPE_FORCE_DWORD            = 0x7fffffff
} D3DRESOURCETYPE;

#define D3DRTYPECOUNT (D3DRTYPE_INDEXBUFFER+1)

typedef enum _D3DSHADEMODE {
    D3DSHADE_FLAT               = 1,
    D3DSHADE_GOURAUD            = 2,
    D3DSHADE_PHONG              = 3,

    D3DSHADE_FORCE_DWORD        = 0x7fffffff
} D3DSHADEMODE;

typedef enum _D3DSTATEBLOCKTYPE {
    D3DSBT_ALL           = 1,
    D3DSBT_PIXELSTATE    = 2,
    D3DSBT_VERTEXSTATE   = 3,

    D3DSBT_FORCE_DWORD   = 0xffffffff
} D3DSTATEBLOCKTYPE;

typedef enum _D3DSTENCILOP {
    D3DSTENCILOP_KEEP           = 1,
    D3DSTENCILOP_ZERO           = 2,
    D3DSTENCILOP_REPLACE        = 3,
    D3DSTENCILOP_INCRSAT        = 4,
    D3DSTENCILOP_DECRSAT        = 5,
    D3DSTENCILOP_INVERT         = 6,
    D3DSTENCILOP_INCR           = 7,
    D3DSTENCILOP_DECR           = 8,

    D3DSTENCILOP_FORCE_DWORD    = 0x7fffffff
} D3DSTENCILOP;

typedef enum _D3DSWAPEFFECT {
    D3DSWAPEFFECT_DISCARD         = 1,
    D3DSWAPEFFECT_FLIP            = 2,
    D3DSWAPEFFECT_COPY            = 3,
    D3DSWAPEFFECT_OVERLAY         = 4,
    D3DSWAPEFFECT_FLIPEX          = 5,
    D3DSWAPEFFECT_FORCE_DWORD     = 0xFFFFFFFF
} D3DSWAPEFFECT;

typedef enum _D3DTEXTUREADDRESS {
    D3DTADDRESS_WRAP            = 1,
    D3DTADDRESS_MIRROR          = 2,
    D3DTADDRESS_CLAMP           = 3,
    D3DTADDRESS_BORDER          = 4,
    D3DTADDRESS_MIRRORONCE      = 5,

    D3DTADDRESS_FORCE_DWORD     = 0x7fffffff
} D3DTEXTUREADDRESS;

typedef enum _D3DTEXTUREFILTERTYPE {
    D3DTEXF_NONE            = 0,
    D3DTEXF_POINT           = 1,
    D3DTEXF_LINEAR          = 2,
    D3DTEXF_ANISOTROPIC     = 3,
    D3DTEXF_FLATCUBIC       = 4,
    D3DTEXF_GAUSSIANCUBIC   = 5,
    D3DTEXF_PYRAMIDALQUAD   = 6,
    D3DTEXF_GAUSSIANQUAD    = 7,
    D3DTEXF_FORCE_DWORD     = 0x7fffffff
} D3DTEXTUREFILTERTYPE;

typedef enum _D3DTEXTUREOP {
    D3DTOP_DISABLE                   =  1,
    D3DTOP_SELECTARG1                =  2,
    D3DTOP_SELECTARG2                =  3,
    D3DTOP_MODULATE                  =  4,
    D3DTOP_MODULATE2X                =  5,
    D3DTOP_MODULATE4X                =  6,
    D3DTOP_ADD                       =  7,
    D3DTOP_ADDSIGNED                 =  8,
    D3DTOP_ADDSIGNED2X               =  9,
    D3DTOP_SUBTRACT                  = 10,
    D3DTOP_ADDSMOOTH                 = 11,
    D3DTOP_BLENDDIFFUSEALPHA         = 12,
    D3DTOP_BLENDTEXTUREALPHA         = 13,
    D3DTOP_BLENDFACTORALPHA          = 14,
    D3DTOP_BLENDTEXTUREALPHAPM       = 15,
    D3DTOP_BLENDCURRENTALPHA         = 16,
    D3DTOP_PREMODULATE               = 17,
    D3DTOP_MODULATEALPHA_ADDCOLOR    = 18,
    D3DTOP_MODULATECOLOR_ADDALPHA    = 19,
    D3DTOP_MODULATEINVALPHA_ADDCOLOR = 20,
    D3DTOP_MODULATEINVCOLOR_ADDALPHA = 21,
    D3DTOP_BUMPENVMAP                = 22,
    D3DTOP_BUMPENVMAPLUMINANCE       = 23,
    D3DTOP_DOTPRODUCT3               = 24,
    D3DTOP_MULTIPLYADD               = 25,
    D3DTOP_LERP                      = 26,

    D3DTOP_FORCE_DWORD               = 0x7fffffff,
} D3DTEXTUREOP;

typedef enum _D3DTEXTURESTAGESTATETYPE {
    D3DTSS_COLOROP               =  1,
    D3DTSS_COLORARG1             =  2,
    D3DTSS_COLORARG2             =  3,
    D3DTSS_ALPHAOP               =  4,
    D3DTSS_ALPHAARG1             =  5,
    D3DTSS_ALPHAARG2             =  6,
    D3DTSS_BUMPENVMAT00          =  7,
    D3DTSS_BUMPENVMAT01          =  8,
    D3DTSS_BUMPENVMAT10          =  9,
    D3DTSS_BUMPENVMAT11          = 10,
    D3DTSS_TEXCOORDINDEX         = 11,
    D3DTSS_BUMPENVLSCALE         = 22,
    D3DTSS_BUMPENVLOFFSET        = 23,
    D3DTSS_TEXTURETRANSFORMFLAGS = 24,
    D3DTSS_COLORARG0             = 26,
    D3DTSS_ALPHAARG0             = 27,
    D3DTSS_RESULTARG             = 28,
    D3DTSS_CONSTANT              = 32,

    D3DTSS_FORCE_DWORD           = 0x7fffffff
} D3DTEXTURESTAGESTATETYPE;

typedef enum _D3DTEXTURETRANSFORMFLAGS {
    D3DTTFF_DISABLE         =   0,
    D3DTTFF_COUNT1          =   1,
    D3DTTFF_COUNT2          =   2,
    D3DTTFF_COUNT3          =   3,
    D3DTTFF_COUNT4          =   4,
    D3DTTFF_PROJECTED       = 256,

    D3DTTFF_FORCE_DWORD     = 0x7fffffff
} D3DTEXTURETRANSFORMFLAGS;

typedef enum _D3DTRANSFORMSTATETYPE {
    D3DTS_VIEW            =  2,
    D3DTS_PROJECTION      =  3,
    D3DTS_TEXTURE0        = 16,
    D3DTS_TEXTURE1        = 17,
    D3DTS_TEXTURE2        = 18,
    D3DTS_TEXTURE3        = 19,
    D3DTS_TEXTURE4        = 20,
    D3DTS_TEXTURE5        = 21,
    D3DTS_TEXTURE6        = 22,
    D3DTS_TEXTURE7        = 23,

    D3DTS_FORCE_DWORD     = 0x7fffffff
} D3DTRANSFORMSTATETYPE;

typedef enum _D3DVERTEXBLENDFLAGS {
    D3DVBF_DISABLE  =   0,
    D3DVBF_1WEIGHTS =   1,
    D3DVBF_2WEIGHTS =   2,
    D3DVBF_3WEIGHTS =   3,
    D3DVBF_TWEENING = 255,
    D3DVBF_0WEIGHTS = 256
} D3DVERTEXBLENDFLAGS;

typedef enum _D3DZBUFFERTYPE {
    D3DZB_FALSE                 = 0,
    D3DZB_TRUE                  = 1,
    D3DZB_USEW                  = 2,

    D3DZB_FORCE_DWORD           = 0x7fffffff
} D3DZBUFFERTYPE;

typedef enum _D3DSAMPLERSTATETYPE {
    D3DSAMP_ADDRESSU       = 1,
    D3DSAMP_ADDRESSV       = 2,
    D3DSAMP_ADDRESSW       = 3,
    D3DSAMP_BORDERCOLOR    = 4,
    D3DSAMP_MAGFILTER      = 5,
    D3DSAMP_MINFILTER      = 6,
    D3DSAMP_MIPFILTER      = 7,
    D3DSAMP_MIPMAPLODBIAS  = 8,
    D3DSAMP_MAXMIPLEVEL    = 9,
    D3DSAMP_MAXANISOTROPY  = 10,
    D3DSAMP_SRGBTEXTURE    = 11,
    D3DSAMP_ELEMENTINDEX   = 12,
    D3DSAMP_DMAPOFFSET     = 13,
                                
    D3DSAMP_FORCE_DWORD   = 0x7fffffff,
} D3DSAMPLERSTATETYPE;


/*****************************************************************************
 * Direct 3D v9 typedefs
 */
#ifndef D3DCOLOR_DEFINED
typedef DWORD D3DCOLOR;
#define D3DCOLOR_DEFINED
#endif

/*****************************************************************************
 * Direct 3D v9 structures
 */
typedef struct _D3DADAPTER_IDENTIFIER9 {
    char            Driver[MAX_DEVICE_IDENTIFIER_STRING];
    char            Description[MAX_DEVICE_IDENTIFIER_STRING];
    char            DeviceName[32];
    LARGE_INTEGER   DriverVersion; 

    DWORD           VendorId;
    DWORD           DeviceId;
    DWORD           SubSysId;
    DWORD           Revision;

    GUID            DeviceIdentifier;

    DWORD           WHQLLevel;
} D3DADAPTER_IDENTIFIER9;

typedef struct _D3DBOX {
    UINT                Left;
    UINT                Top;
    UINT                Right;
    UINT                Bottom;
    UINT                Front;
    UINT                Back;
} D3DBOX;

typedef struct _D3DCLIPSTATUS9 {
   DWORD ClipUnion;
   DWORD ClipIntersection;
} D3DCLIPSTATUS9;

#ifndef D3DCOLORVALUE_DEFINED
typedef struct _D3DCOLORVALUE {
    float r;
    float g;
    float b;
    float a;
} D3DCOLORVALUE;
#define D3DCOLORVALUE_DEFINED
#endif

typedef struct _D3DDEVICE_CREATION_PARAMETERS {
    UINT          AdapterOrdinal;
    D3DDEVTYPE    DeviceType;
    HWND          hFocusWindow;
    DWORD         BehaviorFlags;
} D3DDEVICE_CREATION_PARAMETERS;

typedef struct _D3DDEVINFO_D3D9BANDWIDTHTIMINGS {
    float         MaxBandwidthUtilized;
    float         FrontEndUploadMemoryUtilizedPercent;
    float         VertexRateUtilizedPercent;
    float         TriangleSetupRateUtilizedPercent;
    float         FillRateUtilizedPercent;
} D3DDEVINFO_D3D9BANDWIDTHTIMINGS;

typedef struct _D3DDEVINFO_D3D9CACHEUTILIZATION {
    float         TextureCacheHitRate;
    float         PostTransformVertexCacheHitRate;
} D3DDEVINFO_D3D9CACHEUTILIZATION;

typedef struct _D3DDEVINFO_D3D9INTERFACETIMINGS {
    float         WaitingForGPUToUseApplicationResourceTimePercent;
    float         WaitingForGPUToAcceptMoreCommandsTimePercent;
    float         WaitingForGPUToStayWithinLatencyTimePercent;
    float         WaitingForGPUExclusiveResourceTimePercent;
    float         WaitingForGPUOtherTimePercent;
} D3DDEVINFO_D3D9INTERFACETIMINGS;

typedef struct _D3DDEVINFO_D3D9PIPELINETIMINGS {
    float         VertexProcessingTimePercent;
    float         PixelProcessingTimePercent;
    float         OtherGPUProcessingTimePercent;
    float         GPUIdleTimePercent;
} D3DDEVINFO_D3D9PIPELINETIMINGS;

typedef struct _D3DDEVINFO_D3D9STAGETIMINGS {
    float         MemoryProcessingPercent;
    float         ComputationProcessingPercent;
} D3DDEVINFO_D3D9STAGETIMINGS;


/* Vertex cache optimization hints. */
typedef struct D3DDEVINFO_VCACHE {
    /* Must be a 4 char code FOURCC (e.g. CACH) */
    DWORD         Pattern; 
    /* 0 to get the longest  strips, 1 vertex cache */
    DWORD         OptMethod; 
     /* Cache size to use (only valid if OptMethod==1) */
    DWORD         CacheSize;
    /* internal for deciding when to restart strips, non user modifiable (only valid if OptMethod==1) */
    DWORD         MagicNumber; 
} D3DDEVINFO_VCACHE;

typedef struct D3DRESOURCESTATS {
    WINBOOL             bThrashing;
    DWORD               ApproxBytesDownloaded;
    DWORD               NumEvicts;
    DWORD               NumVidCreates;
    DWORD               LastPri;
    DWORD               NumUsed;
    DWORD               NumUsedInVidMem;
    DWORD               WorkingSet;
    DWORD               WorkingSetBytes;
    DWORD               TotalManaged;
    DWORD               TotalBytes;
} D3DRESOURCESTATS;

typedef struct _D3DDEVINFO_D3DRESOURCEMANAGER {
    D3DRESOURCESTATS stats[D3DRTYPECOUNT];
} D3DDEVINFO_D3DRESOURCEMANAGER;

typedef struct _D3DDEVINFO_D3DVERTEXSTATS {
    DWORD NumRenderedTriangles;
    DWORD NumExtraClippingTriangles;
} D3DDEVINFO_D3DVERTEXSTATS;

typedef struct _D3DDISPLAYMODE {
    UINT            Width;
    UINT            Height;
    UINT            RefreshRate;
    D3DFORMAT       Format;
} D3DDISPLAYMODE;

typedef struct _D3DGAMMARAMP {
    WORD                red  [256];
    WORD                green[256];
    WORD                blue [256];
} D3DGAMMARAMP;

typedef struct _D3DINDEXBUFFER_DESC {
    D3DFORMAT           Format;
    D3DRESOURCETYPE     Type;
    DWORD               Usage;
    D3DPOOL             Pool;
    UINT                Size;
} D3DINDEXBUFFER_DESC;

#ifndef D3DVECTOR_DEFINED
typedef struct _D3DVECTOR {
    float x;
    float y;
    float z;
} D3DVECTOR;
#define D3DVECTOR_DEFINED
#endif

typedef struct _D3DLIGHT9 {
    D3DLIGHTTYPE    Type;
    D3DCOLORVALUE   Diffuse;
    D3DCOLORVALUE   Specular;
    D3DCOLORVALUE   Ambient;
    D3DVECTOR       Position;
    D3DVECTOR       Direction;
    float           Range;
    float           Falloff;
    float           Attenuation0;
    float           Attenuation1;
    float           Attenuation2;
    float           Theta;
    float           Phi;
} D3DLIGHT9;

typedef struct _D3DLINEPATTERN {
    WORD    wRepeatFactor;
    WORD    wLinePattern;
} D3DLINEPATTERN;

typedef struct _D3DLOCKED_BOX {
    INT                 RowPitch;
    INT                 SlicePitch;
    void*               pBits;
} D3DLOCKED_BOX;

typedef struct _D3DLOCKED_RECT {
    INT                 Pitch;
    void*               pBits;
} D3DLOCKED_RECT;

typedef struct _D3DMATERIAL9 {
    D3DCOLORVALUE   Diffuse;
    D3DCOLORVALUE   Ambient;
    D3DCOLORVALUE   Specular;
    D3DCOLORVALUE   Emissive;
    float           Power;
} D3DMATERIAL9;

#ifndef D3DMATRIX_DEFINED
typedef struct _D3DMATRIX {
    union {
        struct {
            float        _11, _12, _13, _14;
            float        _21, _22, _23, _24;
            float        _31, _32, _33, _34;
            float        _41, _42, _43, _44;
        } DUMMYSTRUCTNAME;
        float m[4][4];
    } DUMMYUNIONNAME;
} D3DMATRIX;
#define D3DMATRIX_DEFINED
#endif

typedef struct _D3DPRESENT_PARAMETERS_ {
    UINT                    BackBufferWidth;
    UINT                    BackBufferHeight;
    D3DFORMAT               BackBufferFormat;
    UINT                    BackBufferCount;

    D3DMULTISAMPLE_TYPE     MultiSampleType;
    DWORD                   MultiSampleQuality;

    D3DSWAPEFFECT           SwapEffect;
    HWND                    hDeviceWindow;
    WINBOOL                 Windowed;
    WINBOOL                 EnableAutoDepthStencil;
    D3DFORMAT               AutoDepthStencilFormat;
    DWORD                   Flags;

    UINT                    FullScreen_RefreshRateInHz;
    UINT                    PresentationInterval;

} D3DPRESENT_PARAMETERS;

typedef struct _D3DRANGE {
    UINT                Offset;
    UINT                Size;
} D3DRANGE;

typedef struct _D3DRASTER_STATUS {
    WINBOOL         InVBlank;
    UINT            ScanLine;
} D3DRASTER_STATUS;

#ifndef D3DRECT_DEFINED
typedef struct _D3DRECT {
    LONG x1;
    LONG y1;
    LONG x2;
    LONG y2;
} D3DRECT;
#define D3DRECT_DEFINED
#endif

typedef struct _D3DRECTPATCH_INFO {
    UINT                StartVertexOffsetWidth;
    UINT                StartVertexOffsetHeight;
    UINT                Width;
    UINT                Height;
    UINT                Stride;
    D3DBASISTYPE        Basis;
    D3DDEGREETYPE       Degree;
} D3DRECTPATCH_INFO;

typedef struct _D3DSURFACE_DESC {
    D3DFORMAT           Format;
    D3DRESOURCETYPE     Type;
    DWORD               Usage;
    D3DPOOL             Pool;
    D3DMULTISAMPLE_TYPE MultiSampleType;
    DWORD               MultiSampleQuality;
    UINT                Width;
    UINT                Height;
} D3DSURFACE_DESC;

typedef struct _D3DTRIPATCH_INFO {
    UINT                StartVertexOffset;
    UINT                NumVertices;
    D3DBASISTYPE        Basis;
    D3DDEGREETYPE       Degree;
} D3DTRIPATCH_INFO;

typedef struct _D3DVERTEXBUFFER_DESC {
    D3DFORMAT           Format;
    D3DRESOURCETYPE     Type;
    DWORD               Usage;
    D3DPOOL             Pool;
    UINT                Size;
    DWORD               FVF;
} D3DVERTEXBUFFER_DESC;

typedef struct _D3DVIEWPORT9 {
    DWORD       X;
    DWORD       Y;
    DWORD       Width;
    DWORD       Height;
    float       MinZ;
    float       MaxZ;
} D3DVIEWPORT9;

typedef struct _D3DVOLUME_DESC {
    D3DFORMAT           Format;
    D3DRESOURCETYPE     Type;
    DWORD               Usage;
    D3DPOOL             Pool;

    UINT                Width;
    UINT                Height;
    UINT                Depth;
} D3DVOLUME_DESC;

/* Parts added with d3d9ex */
#if !defined(D3D_DISABLE_9EX)
typedef enum D3DSCANLINEORDERING
{
    D3DSCANLINEORDERING_UNKNOWN,
    D3DSCANLINEORDERING_PROGRESSIVE,
    D3DSCANLINEORDERING_INTERLACED,
} D3DSCANLINEORDERING;


typedef struct D3DDISPLAYMODEFILTER
{
    UINT                Size;
    D3DFORMAT           Format;
    D3DSCANLINEORDERING ScanLineOrdering;
} D3DDISPLAYMODEFILTER;

typedef struct D3DDISPLAYMODEEX
{
    UINT                Size;
    UINT                Width;
    UINT                Height;
    UINT                RefreshRate;
    D3DFORMAT           Format;
    D3DSCANLINEORDERING ScanLineOrdering;
} D3DDISPLAYMODEEX;

typedef enum D3DDISPLAYROTATION
{
    D3DDISPLAYROTATION_IDENTITY = 1,
    D3DDISPLAYROTATION_90,
    D3DDISPLAYROTATION_180,
    D3DDISPLAYROTATION_270
} D3DDISPLAYROTATION;

typedef enum _D3DCOMPOSERECTSOP{
    D3DCOMPOSERECTS_COPY        = 1,
    D3DCOMPOSERECTS_OR,
    D3DCOMPOSERECTS_AND,
    D3DCOMPOSERECTS_NEG,
    D3DCOMPOSERECTS_FORCE_DWORD = 0x7fffffff
} D3DCOMPOSERECTSOP;

typedef struct _D3DPRESENTSTATS
{
    UINT          PresentCount;
    UINT          PresentRefreshCount;
    UINT          SyncRefreshCount;
    LARGE_INTEGER SyncQPCTime;
    LARGE_INTEGER SyncGPUTime;
} D3DPRESENTSTATS;

#endif /* D3D_DISABLE_9EX */

typedef enum _D3DSHADER_COMPARISON
{
    D3DSPC_RESERVED0 = 0,
    D3DSPC_GT,
    D3DSPC_EQ,
    D3DSPC_GE,
    D3DSPC_LT,
    D3DSPC_NE,
    D3DSPC_LE,
    D3DSPC_RESERVED1,
} D3DSHADER_COMPARISON;

#endif /* __WINE_D3D9TYPES_H */

// end of d3d9types.h

// d3d9caps.h
/*
 * Copyright (C) 2002-2003 Jason Edmeades
 *                         Raphael Junqueira
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with this library; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301, USA
 */

#ifndef __WINE_D3D9CAPS_H
#define __WINE_D3D9CAPS_H

/*
 * Definitions
 */
#define D3DCAPS_READ_SCANLINE 0x20000

#define D3DCURSORCAPS_COLOR   1
#define D3DCURSORCAPS_LOWRES  2


#define D3DDEVCAPS2_STREAMOFFSET                        0x00000001
#define D3DDEVCAPS2_DMAPNPATCH                          0x00000002
#define D3DDEVCAPS2_ADAPTIVETESSRTPATCH                 0x00000004
#define D3DDEVCAPS2_ADAPTIVETESSNPATCH                  0x00000008
#define D3DDEVCAPS2_CAN_STRETCHRECT_FROM_TEXTURES       0x00000010
#define D3DDEVCAPS2_PRESAMPLEDDMAPNPATCH                0x00000020
#define D3DDEVCAPS2_VERTEXELEMENTSCANSHARESTREAMOFFSET  0x00000040

#define D3DDEVCAPS_EXECUTESYSTEMMEMORY     0x0000010
#define D3DDEVCAPS_EXECUTEVIDEOMEMORY      0x0000020
#define D3DDEVCAPS_TLVERTEXSYSTEMMEMORY    0x0000040
#define D3DDEVCAPS_TLVERTEXVIDEOMEMORY     0x0000080
#define D3DDEVCAPS_TEXTURESYSTEMMEMORY     0x0000100
#define D3DDEVCAPS_TEXTUREVIDEOMEMORY      0x0000200
#define D3DDEVCAPS_DRAWPRIMTLVERTEX        0x0000400
#define D3DDEVCAPS_CANRENDERAFTERFLIP      0x0000800
#define D3DDEVCAPS_TEXTURENONLOCALVIDMEM   0x0001000
#define D3DDEVCAPS_DRAWPRIMITIVES2         0x0002000
#define D3DDEVCAPS_SEPARATETEXTUREMEMORIES 0x0004000
#define D3DDEVCAPS_DRAWPRIMITIVES2EX       0x0008000
#define D3DDEVCAPS_HWTRANSFORMANDLIGHT     0x0010000
#define D3DDEVCAPS_CANBLTSYSTONONLOCAL     0x0020000
#define D3DDEVCAPS_HWRASTERIZATION         0x0080000
#define D3DDEVCAPS_PUREDEVICE              0x0100000
#define D3DDEVCAPS_QUINTICRTPATCHES        0x0200000
#define D3DDEVCAPS_RTPATCHES               0x0400000
#define D3DDEVCAPS_RTPATCHHANDLEZERO       0x0800000
#define D3DDEVCAPS_NPATCHES                0x1000000

#define D3DFVFCAPS_TEXCOORDCOUNTMASK  0x00FFFF
#define D3DFVFCAPS_DONOTSTRIPELEMENTS 0x080000
#define D3DFVFCAPS_PSIZE              0x100000

#define D3DLINECAPS_TEXTURE           0x01
#define D3DLINECAPS_ZTEST             0x02
#define D3DLINECAPS_BLEND             0x04
#define D3DLINECAPS_ALPHACMP          0x08
#define D3DLINECAPS_FOG               0x10
#define D3DLINECAPS_ANTIALIAS         0x20

#define D3DPBLENDCAPS_ZERO            0x00000001
#define D3DPBLENDCAPS_ONE             0x00000002
#define D3DPBLENDCAPS_SRCCOLOR        0x00000004
#define D3DPBLENDCAPS_INVSRCCOLOR     0x00000008
#define D3DPBLENDCAPS_SRCALPHA        0x00000010
#define D3DPBLENDCAPS_INVSRCALPHA     0x00000020
#define D3DPBLENDCAPS_DESTALPHA       0x00000040
#define D3DPBLENDCAPS_INVDESTALPHA    0x00000080
#define D3DPBLENDCAPS_DESTCOLOR       0x00000100
#define D3DPBLENDCAPS_INVDESTCOLOR    0x00000200
#define D3DPBLENDCAPS_SRCALPHASAT     0x00000400
#define D3DPBLENDCAPS_BOTHSRCALPHA    0x00000800
#define D3DPBLENDCAPS_BOTHINVSRCALPHA 0x00001000
#define D3DPBLENDCAPS_BLENDFACTOR     0x00002000

#define D3DPCMPCAPS_NEVER        0x01
#define D3DPCMPCAPS_LESS         0x02
#define D3DPCMPCAPS_EQUAL        0x04
#define D3DPCMPCAPS_LESSEQUAL    0x08
#define D3DPCMPCAPS_GREATER      0x10
#define D3DPCMPCAPS_NOTEQUAL     0x20
#define D3DPCMPCAPS_GREATEREQUAL 0x40
#define D3DPCMPCAPS_ALWAYS       0x80

#define D3DPMISCCAPS_MASKZ                      0x00000002
#define D3DPMISCCAPS_LINEPATTERNREP             0x00000004
#define D3DPMISCCAPS_CULLNONE                   0x00000010
#define D3DPMISCCAPS_CULLCW                     0x00000020
#define D3DPMISCCAPS_CULLCCW                    0x00000040
#define D3DPMISCCAPS_COLORWRITEENABLE           0x00000080
#define D3DPMISCCAPS_CLIPPLANESCALEDPOINTS      0x00000100
#define D3DPMISCCAPS_CLIPTLVERTS                0x00000200
#define D3DPMISCCAPS_TSSARGTEMP                 0x00000400
#define D3DPMISCCAPS_BLENDOP                    0x00000800
#define D3DPMISCCAPS_NULLREFERENCE              0x00001000
#define D3DPMISCCAPS_INDEPENDENTWRITEMASKS      0x00004000
#define D3DPMISCCAPS_PERSTAGECONSTANT           0x00008000
#define D3DPMISCCAPS_FOGANDSPECULARALPHA        0x00010000
#define D3DPMISCCAPS_SEPARATEALPHABLEND         0x00020000
#define D3DPMISCCAPS_MRTINDEPENDENTBITDEPTHS    0x00040000
#define D3DPMISCCAPS_MRTPOSTPIXELSHADERBLENDING 0x00080000
#define D3DPMISCCAPS_FOGVERTEXCLAMPED           0x00100000


#define D3DPRASTERCAPS_DITHER                     0x00000001
#define D3DPRASTERCAPS_PAT                        0x00000008
#define D3DPRASTERCAPS_ZTEST                      0x00000010
#define D3DPRASTERCAPS_FOGVERTEX                  0x00000080
#define D3DPRASTERCAPS_FOGTABLE                   0x00000100
#define D3DPRASTERCAPS_ANTIALIASEDGES             0x00001000
#define D3DPRASTERCAPS_MIPMAPLODBIAS              0x00002000
#define D3DPRASTERCAPS_ZBIAS                      0x00004000
#define D3DPRASTERCAPS_ZBUFFERLESSHSR             0x00008000
#define D3DPRASTERCAPS_FOGRANGE                   0x00010000
#define D3DPRASTERCAPS_ANISOTROPY                 0x00020000
#define D3DPRASTERCAPS_WBUFFER                    0x00040000
#define D3DPRASTERCAPS_WFOG                       0x00100000
#define D3DPRASTERCAPS_ZFOG                       0x00200000
#define D3DPRASTERCAPS_COLORPERSPECTIVE           0x00400000
#define D3DPRASTERCAPS_SCISSORTEST                0x01000000
#define D3DPRASTERCAPS_SLOPESCALEDEPTHBIAS        0x02000000
#define D3DPRASTERCAPS_DEPTHBIAS                  0x04000000
#define D3DPRASTERCAPS_MULTISAMPLE_TOGGLE         0x08000000

#define D3DPRESENT_INTERVAL_DEFAULT               0x00000000
#define D3DPRESENT_INTERVAL_ONE                   0x00000001
#define D3DPRESENT_INTERVAL_TWO                   0x00000002
#define D3DPRESENT_INTERVAL_THREE                 0x00000004
#define D3DPRESENT_INTERVAL_FOUR                  0x00000008
#define D3DPRESENT_INTERVAL_IMMEDIATE             0x80000000

#define D3DPSHADECAPS_COLORGOURAUDRGB             0x00008
#define D3DPSHADECAPS_SPECULARGOURAUDRGB          0x00200
#define D3DPSHADECAPS_ALPHAGOURAUDBLEND           0x04000
#define D3DPSHADECAPS_FOGGOURAUD                  0x80000

#define D3DPTADDRESSCAPS_WRAP                     0x01
#define D3DPTADDRESSCAPS_MIRROR                   0x02
#define D3DPTADDRESSCAPS_CLAMP                    0x04
#define D3DPTADDRESSCAPS_BORDER                   0x08
#define D3DPTADDRESSCAPS_INDEPENDENTUV            0x10
#define D3DPTADDRESSCAPS_MIRRORONCE               0x20

#define D3DPTEXTURECAPS_PERSPECTIVE              0x00000001
#define D3DPTEXTURECAPS_POW2                     0x00000002
#define D3DPTEXTURECAPS_ALPHA                    0x00000004
#define D3DPTEXTURECAPS_SQUAREONLY               0x00000020
#define D3DPTEXTURECAPS_TEXREPEATNOTSCALEDBYSIZE 0x00000040
#define D3DPTEXTURECAPS_ALPHAPALETTE             0x00000080
#define D3DPTEXTURECAPS_NONPOW2CONDITIONAL       0x00000100
#define D3DPTEXTURECAPS_PROJECTED                0x00000400
#define D3DPTEXTURECAPS_CUBEMAP                  0x00000800
#define D3DPTEXTURECAPS_VOLUMEMAP                0x00002000
#define D3DPTEXTURECAPS_MIPMAP                   0x00004000
#define D3DPTEXTURECAPS_MIPVOLUMEMAP             0x00008000
#define D3DPTEXTURECAPS_MIPCUBEMAP               0x00010000
#define D3DPTEXTURECAPS_CUBEMAP_POW2             0x00020000
#define D3DPTEXTURECAPS_VOLUMEMAP_POW2           0x00040000
#define D3DPTEXTURECAPS_NOPROJECTEDBUMPENV       0x00200000

#define D3DPTFILTERCAPS_MINFPOINT                0x00000100
#define D3DPTFILTERCAPS_MINFLINEAR               0x00000200
#define D3DPTFILTERCAPS_MINFANISOTROPIC          0x00000400
#define D3DPTFILTERCAPS_MINFPYRAMIDALQUAD        0x00000800
#define D3DPTFILTERCAPS_MINFGAUSSIANQUAD         0x00001000
#define D3DPTFILTERCAPS_MIPFPOINT                0x00010000
#define D3DPTFILTERCAPS_MIPFLINEAR               0x00020000
#define D3DPTFILTERCAPS_MAGFPOINT                0x01000000
#define D3DPTFILTERCAPS_MAGFLINEAR               0x02000000
#define D3DPTFILTERCAPS_MAGFANISOTROPIC          0x04000000
#define D3DPTFILTERCAPS_MAGFPYRAMIDALQUAD        0x08000000
#define D3DPTFILTERCAPS_MAGFGAUSSIANQUAD         0x10000000

#define D3DSTENCILCAPS_KEEP                      0x01
#define D3DSTENCILCAPS_ZERO                      0x02
#define D3DSTENCILCAPS_REPLACE                   0x04
#define D3DSTENCILCAPS_INCRSAT                   0x08
#define D3DSTENCILCAPS_DECRSAT                   0x10
#define D3DSTENCILCAPS_INVERT                    0x20
#define D3DSTENCILCAPS_INCR                      0x40
#define D3DSTENCILCAPS_DECR                      0x80
#define D3DSTENCILCAPS_TWOSIDED                  0x100

#define D3DTEXOPCAPS_DISABLE                     0x0000001
#define D3DTEXOPCAPS_SELECTARG1                  0x0000002
#define D3DTEXOPCAPS_SELECTARG2                  0x0000004
#define D3DTEXOPCAPS_MODULATE                    0x0000008
#define D3DTEXOPCAPS_MODULATE2X                  0x0000010
#define D3DTEXOPCAPS_MODULATE4X                  0x0000020
#define D3DTEXOPCAPS_ADD                         0x0000040
#define D3DTEXOPCAPS_ADDSIGNED                   0x0000080
#define D3DTEXOPCAPS_ADDSIGNED2X                 0x0000100
#define D3DTEXOPCAPS_SUBTRACT                    0x0000200
#define D3DTEXOPCAPS_ADDSMOOTH                   0x0000400
#define D3DTEXOPCAPS_BLENDDIFFUSEALPHA           0x0000800
#define D3DTEXOPCAPS_BLENDTEXTUREALPHA           0x0001000
#define D3DTEXOPCAPS_BLENDFACTORALPHA            0x0002000
#define D3DTEXOPCAPS_BLENDTEXTUREALPHAPM         0x0004000
#define D3DTEXOPCAPS_BLENDCURRENTALPHA           0x0008000
#define D3DTEXOPCAPS_PREMODULATE                 0x0010000
#define D3DTEXOPCAPS_MODULATEALPHA_ADDCOLOR      0x0020000
#define D3DTEXOPCAPS_MODULATECOLOR_ADDALPHA      0x0040000
#define D3DTEXOPCAPS_MODULATEINVALPHA_ADDCOLOR   0x0080000
#define D3DTEXOPCAPS_MODULATEINVCOLOR_ADDALPHA   0x0100000
#define D3DTEXOPCAPS_BUMPENVMAP                  0x0200000
#define D3DTEXOPCAPS_BUMPENVMAPLUMINANCE         0x0400000
#define D3DTEXOPCAPS_DOTPRODUCT3                 0x0800000
#define D3DTEXOPCAPS_MULTIPLYADD                 0x1000000
#define D3DTEXOPCAPS_LERP                        0x2000000

#define D3DVTXPCAPS_TEXGEN                         0x00000001
#define D3DVTXPCAPS_MATERIALSOURCE7                0x00000002
#define D3DVTXPCAPS_DIRECTIONALLIGHTS              0x00000008
#define D3DVTXPCAPS_POSITIONALLIGHTS               0x00000010
#define D3DVTXPCAPS_LOCALVIEWER                    0x00000020
#define D3DVTXPCAPS_TWEENING                       0x00000040
#define D3DVTXPCAPS_TEXGEN_SPHEREMAP               0x00000100
#define D3DVTXPCAPS_NO_TEXGEN_NONLOCALVIEWER       0x00000200

#define D3DDTCAPS_UBYTE4                           0x00000001
#define D3DDTCAPS_UBYTE4N                          0x00000002
#define D3DDTCAPS_SHORT2N                          0x00000004
#define D3DDTCAPS_SHORT4N                          0x00000008
#define D3DDTCAPS_USHORT2N                         0x00000010
#define D3DDTCAPS_USHORT4N                         0x00000020
#define D3DDTCAPS_UDEC3                            0x00000040
#define D3DDTCAPS_DEC3N                            0x00000080
#define D3DDTCAPS_FLOAT16_2                        0x00000100
#define D3DDTCAPS_FLOAT16_4                        0x00000200

#define D3DCAPS3_ALPHA_FULLSCREEN_FLIP_OR_DISCARD  0x00000020
#define D3DCAPS3_LINEAR_TO_SRGB_PRESENTATION       0x00000080
#define D3DCAPS3_COPY_TO_VIDMEM                    0x00000100
#define D3DCAPS3_COPY_TO_SYSTEMMEM                 0x00000200
#define D3DCAPS3_RESERVED                          0x8000001F

#define D3DCAPS2_NO2DDURING3DSCENE                 0x00000002
#define D3DCAPS2_FULLSCREENGAMMA                   0x00020000
#define D3DCAPS2_CANRENDERWINDOWED                 0x00080000
#define D3DCAPS2_CANCALIBRATEGAMMA                 0x00100000
#define D3DCAPS2_RESERVED                          0x02000000
#define D3DCAPS2_CANMANAGERESOURCE                 0x10000000
#define D3DCAPS2_DYNAMICTEXTURES                   0x20000000
#define D3DCAPS2_CANAUTOGENMIPMAP                  0x40000000


#define D3DVS20_MAX_DYNAMICFLOWCONTROLDEPTH  24
#define D3DVS20_MIN_DYNAMICFLOWCONTROLDEPTH  0
#define D3DVS20_MAX_NUMTEMPS                 32
#define D3DVS20_MIN_NUMTEMPS                 12
#define D3DVS20_MAX_STATICFLOWCONTROLDEPTH   4
#define D3DVS20_MIN_STATICFLOWCONTROLDEPTH   1

#define D3DVS20CAPS_PREDICATION              (1 << 0)

#define D3DPS20CAPS_ARBITRARYSWIZZLE         (1 << 0)
#define D3DPS20CAPS_GRADIENTINSTRUCTIONS     (1 << 1)
#define D3DPS20CAPS_PREDICATION              (1 << 2)
#define D3DPS20CAPS_NODEPENDENTREADLIMIT     (1 << 3)
#define D3DPS20CAPS_NOTEXINSTRUCTIONLIMIT    (1 << 4)

#define D3DPS20_MAX_DYNAMICFLOWCONTROLDEPTH  24
#define D3DPS20_MIN_DYNAMICFLOWCONTROLDEPTH  0
#define D3DPS20_MAX_NUMTEMPS                 32
#define D3DPS20_MIN_NUMTEMPS                 12
#define D3DPS20_MAX_STATICFLOWCONTROLDEPTH   4
#define D3DPS20_MIN_STATICFLOWCONTROLDEPTH   0
#define D3DPS20_MAX_NUMINSTRUCTIONSLOTS      512
#define D3DPS20_MIN_NUMINSTRUCTIONSLOTS      96

#define D3DMIN30SHADERINSTRUCTIONS          512
#define D3DMAX30SHADERINSTRUCTIONS          32768


typedef struct _D3DVSHADERCAPS2_0 {
  DWORD  Caps;
  INT    DynamicFlowControlDepth;
  INT    NumTemps;
  INT    StaticFlowControlDepth;
} D3DVSHADERCAPS2_0;

typedef struct _D3DPSHADERCAPS2_0 {
  DWORD  Caps;
  INT    DynamicFlowControlDepth;
  INT    NumTemps;
  INT    StaticFlowControlDepth;
  INT    NumInstructionSlots;
} D3DPSHADERCAPS2_0;

/*
 * The d3dcaps9 structure
 */
typedef struct _D3DCAPS9 {
  D3DDEVTYPE          DeviceType;
  UINT                AdapterOrdinal;
  
  DWORD               Caps;
  DWORD               Caps2;
  DWORD               Caps3;
  DWORD               PresentationIntervals;
  
  DWORD               CursorCaps;
  
  DWORD               DevCaps;
  
  DWORD               PrimitiveMiscCaps;
  DWORD               RasterCaps;
  DWORD               ZCmpCaps;
  DWORD               SrcBlendCaps;
  DWORD               DestBlendCaps;
  DWORD               AlphaCmpCaps;
  DWORD               ShadeCaps;
  DWORD               TextureCaps;
  DWORD               TextureFilterCaps;
  DWORD               CubeTextureFilterCaps;
  DWORD               VolumeTextureFilterCaps;
  DWORD               TextureAddressCaps;
  DWORD               VolumeTextureAddressCaps;
  
  DWORD               LineCaps;
  
  DWORD               MaxTextureWidth, MaxTextureHeight;
  DWORD               MaxVolumeExtent;
  
  DWORD               MaxTextureRepeat;
  DWORD               MaxTextureAspectRatio;
  DWORD               MaxAnisotropy;
  float               MaxVertexW;
  
  float               GuardBandLeft;
  float               GuardBandTop;
  float               GuardBandRight;
  float               GuardBandBottom;
  
  float               ExtentsAdjust;
  DWORD               StencilCaps;
  
  DWORD               FVFCaps;
  DWORD               TextureOpCaps;
  DWORD               MaxTextureBlendStages;
  DWORD               MaxSimultaneousTextures;
  
  DWORD               VertexProcessingCaps;
  DWORD               MaxActiveLights;
  DWORD               MaxUserClipPlanes;
  DWORD               MaxVertexBlendMatrices;
  DWORD               MaxVertexBlendMatrixIndex;
  
  float               MaxPointSize;
  
  DWORD               MaxPrimitiveCount;
  DWORD               MaxVertexIndex;
  DWORD               MaxStreams;
  DWORD               MaxStreamStride;
  
  DWORD               VertexShaderVersion;
  DWORD               MaxVertexShaderConst;
  
  DWORD               PixelShaderVersion;
  float               PixelShader1xMaxValue;

  /* DX 9 */
  DWORD               DevCaps2;

  float               MaxNpatchTessellationLevel;
  DWORD               Reserved5;

  UINT                MasterAdapterOrdinal;   
  UINT                AdapterOrdinalInGroup;  
  UINT                NumberOfAdaptersInGroup;
  DWORD               DeclTypes;              
  DWORD               NumSimultaneousRTs;     
  DWORD               StretchRectFilterCaps;  
  D3DVSHADERCAPS2_0   VS20Caps;
  D3DPSHADERCAPS2_0   PS20Caps;
  DWORD               VertexTextureFilterCaps;
  DWORD               MaxVShaderInstructionsExecuted;
  DWORD               MaxPShaderInstructionsExecuted;
  DWORD               MaxVertexShader30InstructionSlots; 
  DWORD               MaxPixelShader30InstructionSlots;

} D3DCAPS9;

#endif

// end of d3d9caps.h

/*****************************************************************************
 * Behavior Flags for IDirect3D8::CreateDevice
 */
#define D3DCREATE_FPU_PRESERVE                  0x00000002
#define D3DCREATE_MULTITHREADED                 0x00000004
#define D3DCREATE_PUREDEVICE                    0x00000010
#define D3DCREATE_SOFTWARE_VERTEXPROCESSING     0x00000020
#define D3DCREATE_HARDWARE_VERTEXPROCESSING     0x00000040
#define D3DCREATE_MIXED_VERTEXPROCESSING        0x00000080
#define D3DCREATE_DISABLE_DRIVER_MANAGEMENT     0x00000100
#define D3DCREATE_ADAPTERGROUP_DEVICE           0x00000200
#define D3DCREATE_DISABLE_DRIVER_MANAGEMENT_EX  0x00000400
#define D3DCREATE_NOWINDOWCHANGES               0x00000800
#define D3DCREATE_DISABLE_PSGP_THREADING        0x00002000
#define D3DCREATE_ENABLE_PRESENTSTATS           0x00004000
#define D3DCREATE_DISABLE_PRINTSCREEN           0x00008000
#define D3DCREATE_SCREENSAVER                   0x10000000

/*****************************************************************************
 * Flags for SetPrivateData
 */
#define D3DSPD_IUNKNOWN                         0x00000001


/*****************************************************************************
 * #defines and error codes
 */
#define D3D_SDK_VERSION                         32
#define D3DADAPTER_DEFAULT                      0
#define D3DENUM_NO_WHQL_LEVEL                   0x00000002
#define D3DPRESENT_BACK_BUFFERS_MAX             3
#define D3DSGR_NO_CALIBRATION                   0x00000000
#define D3DSGR_CALIBRATE                        0x00000001

#define _FACD3D  0x876
#define MAKE_D3DHRESULT( code )                 MAKE_HRESULT( 1, _FACD3D, code )
#define MAKE_D3DSTATUS( code )                  MAKE_HRESULT( 0, _FACD3D, code )

/*****************************************************************************
 * Direct3D Errors
 */
#define D3D_OK                                  S_OK
#define D3DERR_WRONGTEXTUREFORMAT               MAKE_D3DHRESULT(2072)
#define D3DERR_UNSUPPORTEDCOLOROPERATION        MAKE_D3DHRESULT(2073)
#define D3DERR_UNSUPPORTEDCOLORARG              MAKE_D3DHRESULT(2074)
#define D3DERR_UNSUPPORTEDALPHAOPERATION        MAKE_D3DHRESULT(2075)
#define D3DERR_UNSUPPORTEDALPHAARG              MAKE_D3DHRESULT(2076)
#define D3DERR_TOOMANYOPERATIONS                MAKE_D3DHRESULT(2077)
#define D3DERR_CONFLICTINGTEXTUREFILTER         MAKE_D3DHRESULT(2078)
#define D3DERR_UNSUPPORTEDFACTORVALUE           MAKE_D3DHRESULT(2079)
#define D3DERR_CONFLICTINGRENDERSTATE           MAKE_D3DHRESULT(2081)
#define D3DERR_UNSUPPORTEDTEXTUREFILTER         MAKE_D3DHRESULT(2082)
#define D3DERR_CONFLICTINGTEXTUREPALETTE        MAKE_D3DHRESULT(2086)
#define D3DERR_DRIVERINTERNALERROR              MAKE_D3DHRESULT(2087)
#define D3DERR_NOTFOUND                         MAKE_D3DHRESULT(2150)
#define D3DERR_MOREDATA                         MAKE_D3DHRESULT(2151)
#define D3DERR_DEVICELOST                       MAKE_D3DHRESULT(2152)
#define D3DERR_DEVICENOTRESET                   MAKE_D3DHRESULT(2153)
#define D3DERR_NOTAVAILABLE                     MAKE_D3DHRESULT(2154)
#define D3DERR_OUTOFVIDEOMEMORY                 MAKE_D3DHRESULT(380)
#define D3DERR_INVALIDDEVICE                    MAKE_D3DHRESULT(2155)
#define D3DERR_INVALIDCALL                      MAKE_D3DHRESULT(2156)
#define D3DERR_DRIVERINVALIDCALL                MAKE_D3DHRESULT(2157)
#define D3DERR_WASSTILLDRAWING                  MAKE_D3DHRESULT(540)
#define D3DOK_NOAUTOGEN                         MAKE_D3DSTATUS(2159)

#define D3DERR_DEVICEREMOVED                    MAKE_D3DHRESULT(2160)
#define D3DERR_DEVICEHUNG                       MAKE_D3DHRESULT(2164)
#define S_NOT_RESIDENT                          MAKE_D3DSTATUS(2165)
#define S_RESIDENT_IN_SHARED_MEMORY             MAKE_D3DSTATUS(2166)
#define S_PRESENT_MODE_CHANGED                  MAKE_D3DSTATUS(2167)
#define S_PRESENT_OCCLUDED                      MAKE_D3DSTATUS(2168)
#define D3DERR_UNSUPPORTEDOVERLAY               MAKE_D3DHRESULT(2171)
#define D3DERR_UNSUPPORTEDOVERLAYFORMAT         MAKE_D3DHRESULT(2172)
#define D3DERR_CANNOTPROTECTCONTENT             MAKE_D3DHRESULT(2173)
#define D3DERR_UNSUPPORTEDCRYPTO                MAKE_D3DHRESULT(2174)
#define D3DERR_PRESENT_STATISTICS_DISJOINT      MAKE_D3DHRESULT(2180)


/*****************************************************************************
 * Predeclare the interfaces
 */
DEFINE_GUID(IID_IDirect3D9,                   0x81BDCBCA, 0x64D4, 0x426D, 0xAE, 0x8D, 0xAD, 0x1, 0x47, 0xF4, 0x27, 0x5C);
typedef struct IDirect3D9 *LPDIRECT3D9, *PDIRECT3D9;

DEFINE_GUID(IID_IDirect3D9Ex,                 0x02177241, 0x69FC, 0x400C, 0x8F, 0xF1, 0x93, 0xA4, 0x4D, 0xF6, 0x86, 0x1D);
typedef struct IDirect3D9Ex *LPDIRECT3D9EX, *PDIRECT3D9EX;

DEFINE_GUID(IID_IDirect3DDevice9,             0xd0223b96, 0xbf7a, 0x43fd, 0x92, 0xbd, 0xa4, 0x3b, 0xd, 0x82, 0xb9, 0xeb);
typedef struct IDirect3DDevice9 *LPDIRECT3DDEVICE9;

DEFINE_GUID(IID_IDirect3DDevice9Ex,           0xb18b10ce, 0x2649, 0x405a, 0x87, 0xf, 0x95, 0xf7, 0x77, 0xd4, 0x31, 0x3a);
typedef struct IDirect3DDevice9Ex *LPDIRECT3DDEVICE9EX, *PDIRECT3DDEVICE9EX;

DEFINE_GUID(IID_IDirect3DResource9,           0x5eec05d, 0x8f7d, 0x4362, 0xb9, 0x99, 0xd1, 0xba, 0xf3, 0x57, 0xc7, 0x4);
typedef struct IDirect3DResource9 *LPDIRECT3DRESOURCE9, *PDIRECT3DRESOURCE9;

DEFINE_GUID(IID_IDirect3DVertexBuffer9,       0xb64bb1b5, 0xfd70, 0x4df6, 0xbf, 0x91, 0x19, 0xd0, 0xa1, 0x24, 0x55, 0xe3);
typedef struct IDirect3DVertexBuffer9 *LPDIRECT3DVERTEXBUFFER9, *PDIRECT3DVERTEXBUFFER9;

DEFINE_GUID(IID_IDirect3DVolume9,             0x24f416e6, 0x1f67, 0x4aa7, 0xb8, 0x8e, 0xd3, 0x3f, 0x6f, 0x31, 0x28, 0xa1);
typedef struct IDirect3DVolume9 *LPDIRECT3DVOLUME9, *PDIRECT3DVOLUME9;

DEFINE_GUID(IID_IDirect3DSwapChain9,          0x794950f2, 0xadfc, 0x458a, 0x90, 0x5e, 0x10, 0xa1, 0xb, 0xb, 0x50, 0x3b);
typedef struct IDirect3DSwapChain9 *LPDIRECT3DSWAPCHAIN9, *PDIRECT3DSWAPCHAIN9;

DEFINE_GUID(IID_IDirect3DSwapChain9Ex,        0x91886caf, 0x1c3d, 0x4d2e, 0xa0, 0xab, 0x3e, 0x4c, 0x7d, 0x8d, 0x33, 0x3);
typedef struct IDirect3DSwapChain9Ex *LPDIRECT3DSWAPCHAIN9EX, *PDIRECT3DSWAPCHAIN9EX;

DEFINE_GUID(IID_IDirect3DSurface9,            0xcfbaf3a, 0x9ff6, 0x429a, 0x99, 0xb3, 0xa2, 0x79, 0x6a, 0xf8, 0xb8, 0x9b);
typedef struct IDirect3DSurface9 *LPDIRECT3DSURFACE9, *PDIRECT3DSURFACE9;

DEFINE_GUID(IID_IDirect3DIndexBuffer9,        0x7c9dd65e, 0xd3f7, 0x4529, 0xac, 0xee, 0x78, 0x58, 0x30, 0xac, 0xde, 0x35);
typedef struct IDirect3DIndexBuffer9 *LPDIRECT3DINDEXBUFFER9, *PDIRECT3DINDEXBUFFER9;

DEFINE_GUID(IID_IDirect3DBaseTexture9,        0x580ca87e, 0x1d3c, 0x4d54, 0x99, 0x1d, 0xb7, 0xd3, 0xe3, 0xc2, 0x98, 0xce);
typedef struct IDirect3DBaseTexture9 *LPDIRECT3DBASETEXTURE9, *PDIRECT3DBASETEXTURE9;

DEFINE_GUID(IID_IDirect3DTexture9,            0x85c31227, 0x3de5, 0x4f00, 0x9b, 0x3a, 0xf1, 0x1a, 0xc3, 0x8c, 0x18, 0xb5);
typedef struct IDirect3DTexture9 *LPDIRECT3DTEXTURE9, *PDIRECT3DTEXTURE9;

DEFINE_GUID(IID_IDirect3DCubeTexture9,        0xfff32f81, 0xd953, 0x473a, 0x92, 0x23, 0x93, 0xd6, 0x52, 0xab, 0xa9, 0x3f);
typedef struct IDirect3DCubeTexture9 *LPDIRECT3DCUBETEXTURE9, *PDIRECT3DCUBETEXTURE9;

DEFINE_GUID(IID_IDirect3DVolumeTexture9,      0x2518526c, 0xe789, 0x4111, 0xa7, 0xb9, 0x47, 0xef, 0x32, 0x8d, 0x13, 0xe6);
typedef struct IDirect3DVolumeTexture9 *LPDIRECT3DVOLUMETEXTURE9, *PDIRECT3DVOLUMETEXTURE9;

DEFINE_GUID(IID_IDirect3DVertexDeclaration9,  0xdd13c59c, 0x36fa, 0x4098, 0xa8, 0xfb, 0xc7, 0xed, 0x39, 0xdc, 0x85, 0x46);
typedef struct IDirect3DVertexDeclaration9 *LPDIRECT3DVERTEXDECLARATION9;

DEFINE_GUID(IID_IDirect3DVertexShader9,       0xefc5557e, 0x6265, 0x4613, 0x8a, 0x94, 0x43, 0x85, 0x78, 0x89, 0xeb, 0x36);
typedef struct IDirect3DVertexShader9 *LPDIRECT3DVERTEXSHADER9;

DEFINE_GUID(IID_IDirect3DPixelShader9,        0x6d3bdbdc, 0x5b02, 0x4415, 0xb8, 0x52, 0xce, 0x5e, 0x8b, 0xcc, 0xb2, 0x89);
typedef struct IDirect3DPixelShader9 *LPDIRECT3DPIXELSHADER9;

DEFINE_GUID(IID_IDirect3DStateBlock9,         0xb07c4fe5, 0x310d, 0x4ba8, 0xa2, 0x3c, 0x4f, 0xf, 0x20, 0x6f, 0x21, 0x8b);
typedef struct IDirect3DStateBlock9 *LPDIRECT3DSTATEBLOCK9;

DEFINE_GUID(IID_IDirect3DQuery9,              0xd9771460, 0xa695, 0x4f26, 0xbb, 0xd3, 0x27, 0xb8, 0x40, 0xb5, 0x41, 0xcc);
typedef struct IDirect3DQuery9 *LPDIRECT3DQUERY9, *PDIRECT3DQUERY9;

/*****************************************************************************
 * IDirect3D9 interface
 */
#undef INTERFACE
#define INTERFACE IDirect3D9
DECLARE_INTERFACE_(IDirect3D9,IUnknown)
{
    /*** IUnknown methods ***/
    STDMETHOD_(HRESULT,QueryInterface)(THIS_ REFIID riid, void** ppvObject) PURE;
    STDMETHOD_(ULONG,AddRef)(THIS) PURE;
    STDMETHOD_(ULONG,Release)(THIS) PURE;
    /*** IDirect3D9 methods ***/
    STDMETHOD(RegisterSoftwareDevice)(THIS_ void* pInitializeFunction) PURE;
    STDMETHOD_(UINT, GetAdapterCount)(THIS) PURE;
    STDMETHOD(GetAdapterIdentifier)(THIS_ UINT Adapter, DWORD Flags, D3DADAPTER_IDENTIFIER9* pIdentifier) PURE;
    STDMETHOD_(UINT, GetAdapterModeCount)(THIS_ UINT Adapter, D3DFORMAT Format) PURE;
    STDMETHOD(EnumAdapterModes)(THIS_ UINT Adapter, D3DFORMAT Format, UINT Mode, D3DDISPLAYMODE* pMode) PURE;
    STDMETHOD(GetAdapterDisplayMode)(THIS_ UINT Adapter, D3DDISPLAYMODE* pMode) PURE;
    STDMETHOD(CheckDeviceType)(THIS_ UINT iAdapter, D3DDEVTYPE DevType, D3DFORMAT DisplayFormat, D3DFORMAT BackBufferFormat, WINBOOL bWindowed) PURE;
    STDMETHOD(CheckDeviceFormat)(THIS_ UINT Adapter, D3DDEVTYPE DeviceType, D3DFORMAT AdapterFormat, DWORD Usage, D3DRESOURCETYPE RType, D3DFORMAT CheckFormat) PURE;
    STDMETHOD(CheckDeviceMultiSampleType)(THIS_ UINT Adapter, D3DDEVTYPE DeviceType, D3DFORMAT SurfaceFormat, WINBOOL Windowed, D3DMULTISAMPLE_TYPE MultiSampleType, DWORD* pQualityLevels) PURE;
    STDMETHOD(CheckDepthStencilMatch)(THIS_ UINT Adapter, D3DDEVTYPE DeviceType, D3DFORMAT AdapterFormat, D3DFORMAT RenderTargetFormat, D3DFORMAT DepthStencilFormat) PURE;
    STDMETHOD(CheckDeviceFormatConversion)(THIS_ UINT Adapter, D3DDEVTYPE DeviceType, D3DFORMAT SourceFormat, D3DFORMAT TargetFormat) PURE;
    STDMETHOD(GetDeviceCaps)(THIS_ UINT Adapter, D3DDEVTYPE DeviceType, D3DCAPS9* pCaps) PURE;
    STDMETHOD_(HMONITOR, GetAdapterMonitor)(THIS_ UINT Adapter) PURE;
    STDMETHOD(CreateDevice)(THIS_ UINT Adapter, D3DDEVTYPE DeviceType, HWND hFocusWindow, DWORD BehaviorFlags, D3DPRESENT_PARAMETERS* pPresentationParameters, struct IDirect3DDevice9** ppReturnedDeviceInterface) PURE;
};
#undef INTERFACE

#if !defined(__cplusplus) || defined(CINTERFACE)
/*** IUnknown methods ***/
#define IDirect3D9_QueryInterface(p,a,b) (p)->lpVtbl->QueryInterface(p,a,b)
#define IDirect3D9_AddRef(p)             (p)->lpVtbl->AddRef(p)
#define IDirect3D9_Release(p)            (p)->lpVtbl->Release(p)
/*** IDirect3D9 methods ***/
#define IDirect3D9_RegisterSoftwareDevice(p,a)                (p)->lpVtbl->RegisterSoftwareDevice(p,a)
#define IDirect3D9_GetAdapterCount(p)                         (p)->lpVtbl->GetAdapterCount(p)
#define IDirect3D9_GetAdapterIdentifier(p,a,b,c)              (p)->lpVtbl->GetAdapterIdentifier(p,a,b,c)
#define IDirect3D9_GetAdapterModeCount(p,a,b)                 (p)->lpVtbl->GetAdapterModeCount(p,a,b)
#define IDirect3D9_EnumAdapterModes(p,a,b,c,d)                (p)->lpVtbl->EnumAdapterModes(p,a,b,c,d)
#define IDirect3D9_GetAdapterDisplayMode(p,a,b)               (p)->lpVtbl->GetAdapterDisplayMode(p,a,b)
#define IDirect3D9_CheckDeviceType(p,a,b,c,d,e)               (p)->lpVtbl->CheckDeviceType(p,a,b,c,d,e)
#define IDirect3D9_CheckDeviceFormat(p,a,b,c,d,e,f)           (p)->lpVtbl->CheckDeviceFormat(p,a,b,c,d,e,f)
#define IDirect3D9_CheckDeviceMultiSampleType(p,a,b,c,d,e,f)  (p)->lpVtbl->CheckDeviceMultiSampleType(p,a,b,c,d,e,f)
#define IDirect3D9_CheckDepthStencilMatch(p,a,b,c,d,e)        (p)->lpVtbl->CheckDepthStencilMatch(p,a,b,c,d,e)
#define IDirect3D9_CheckDeviceFormatConversion(p,a,b,c,d)     (p)->lpVtbl->CheckDeviceFormatConversion(p,a,b,c,d)
#define IDirect3D9_GetDeviceCaps(p,a,b,c)                     (p)->lpVtbl->GetDeviceCaps(p,a,b,c)
#define IDirect3D9_GetAdapterMonitor(p,a)                     (p)->lpVtbl->GetAdapterMonitor(p,a)
#define IDirect3D9_CreateDevice(p,a,b,c,d,e,f)                (p)->lpVtbl->CreateDevice(p,a,b,c,d,e,f)
#else
/*** IUnknown methods ***/
#define IDirect3D9_QueryInterface(p,a,b) (p)->QueryInterface(a,b)
#define IDirect3D9_AddRef(p)             (p)->AddRef()
#define IDirect3D9_Release(p)            (p)->Release()
/*** IDirect3D9 methods ***/
#define IDirect3D9_RegisterSoftwareDevice(p,a)                (p)->RegisterSoftwareDevice(a)
#define IDirect3D9_GetAdapterCount(p)                         (p)->GetAdapterCount()
#define IDirect3D9_GetAdapterIdentifier(p,a,b,c)              (p)->GetAdapterIdentifier(a,b,c)
#define IDirect3D9_GetAdapterModeCount(p,a,b)                 (p)->GetAdapterModeCount(a,b)
#define IDirect3D9_EnumAdapterModes(p,a,b,c,d)                (p)->EnumAdapterModes(a,b,c,d)
#define IDirect3D9_GetAdapterDisplayMode(p,a,b)               (p)->GetAdapterDisplayMode(a,b)
#define IDirect3D9_CheckDeviceType(p,a,b,c,d,e)               (p)->CheckDeviceType(a,b,c,d,e)
#define IDirect3D9_CheckDeviceFormat(p,a,b,c,d,e,f)           (p)->CheckDeviceFormat(a,b,c,d,e,f)
#define IDirect3D9_CheckDeviceMultiSampleType(p,a,b,c,d,e,f)  (p)->CheckDeviceMultiSampleType(a,b,c,d,e,f)
#define IDirect3D9_CheckDepthStencilMatch(p,a,b,c,d,e)        (p)->CheckDepthStencilMatch(a,b,c,d,e)
#define IDirect3D9_CheckDeviceFormatConversion(p,a,b,c,d)     (p)->CheckDeviceFormatConversion(a,b,c,d)
#define IDirect3D9_GetDeviceCaps(p,a,b,c)                     (p)->GetDeviceCaps(a,b,c)
#define IDirect3D9_GetAdapterMonitor(p,a)                     (p)->GetAdapterMonitor(a)
#define IDirect3D9_CreateDevice(p,a,b,c,d,e,f)                (p)->CreateDevice(a,b,c,d,e,f)
#endif


#if !defined(D3D_DISABLE_9EX)
/*****************************************************************************
 * IDirect3D9Ex interface
 */
#define INTERFACE IDirect3D9Ex
DECLARE_INTERFACE_(IDirect3D9Ex,IDirect3D9)
{
    /*** IUnknown methods ***/
    STDMETHOD_(HRESULT,QueryInterface)(THIS_ REFIID riid, void** ppvObject) PURE;
    STDMETHOD_(ULONG,AddRef)(THIS) PURE;
    STDMETHOD_(ULONG,Release)(THIS) PURE;
    /*** IDirect3D9 methods ***/

    /* Note: Microsoft's d3d9.h does not declare IDirect3D9Ex::RegisterSoftwareDevice . This would mean that
     * the offsets of the other methods in the Vtable change too. This is wrong. In Microsoft's
     * d3d9.dll, the offsets for the other functions are still compatible with IDirect3D9.
     * This is probably because even in MS header IDirect3D9Ex inherits from IDirect3D9, which makes the
     * C++ interface compatible, and nobody uses the C interface in Windows world.
     */
    STDMETHOD(RegisterSoftwareDevice)(THIS_ void* pInitializeFunction) PURE;

    STDMETHOD_(UINT, GetAdapterCount)(THIS) PURE;
    STDMETHOD(GetAdapterIdentifier)(THIS_ UINT Adapter, DWORD Flags, D3DADAPTER_IDENTIFIER9* pIdentifier) PURE;
    STDMETHOD_(UINT, GetAdapterModeCount)(THIS_ UINT Adapter, D3DFORMAT Format) PURE;
    STDMETHOD(EnumAdapterModes)(THIS_ UINT Adapter, D3DFORMAT Format, UINT Mode, D3DDISPLAYMODE* pMode) PURE;
    STDMETHOD(GetAdapterDisplayMode)(THIS_ UINT Adapter, D3DDISPLAYMODE* pMode) PURE;
    STDMETHOD(CheckDeviceType)(THIS_ UINT iAdapter, D3DDEVTYPE DevType, D3DFORMAT DisplayFormat, D3DFORMAT BackBufferFormat, WINBOOL bWindowed) PURE;
    STDMETHOD(CheckDeviceFormat)(THIS_ UINT Adapter, D3DDEVTYPE DeviceType, D3DFORMAT AdapterFormat, DWORD Usage, D3DRESOURCETYPE RType, D3DFORMAT CheckFormat) PURE;
    STDMETHOD(CheckDeviceMultiSampleType)(THIS_ UINT Adapter, D3DDEVTYPE DeviceType, D3DFORMAT SurfaceFormat, WINBOOL Windowed, D3DMULTISAMPLE_TYPE MultiSampleType, DWORD* pQualityLevels) PURE;
    STDMETHOD(CheckDepthStencilMatch)(THIS_ UINT Adapter, D3DDEVTYPE DeviceType, D3DFORMAT AdapterFormat, D3DFORMAT RenderTargetFormat, D3DFORMAT DepthStencilFormat) PURE;
    STDMETHOD(CheckDeviceFormatConversion)(THIS_ UINT Adapter, D3DDEVTYPE DeviceType, D3DFORMAT SourceFormat, D3DFORMAT TargetFormat) PURE;
    STDMETHOD(GetDeviceCaps)(THIS_ UINT Adapter, D3DDEVTYPE DeviceType, D3DCAPS9* pCaps) PURE;
    STDMETHOD_(HMONITOR, GetAdapterMonitor)(THIS_ UINT Adapter) PURE;
    STDMETHOD(CreateDevice)(THIS_ UINT Adapter, D3DDEVTYPE DeviceType, HWND hFocusWindow, DWORD BehaviorFlags, D3DPRESENT_PARAMETERS* pPresentationParameters, struct IDirect3DDevice9** ppReturnedDeviceInterface) PURE;
    /*** IDirect3D9Ex methods ***/
    STDMETHOD_(UINT, GetAdapterModeCountEx)(THIS_ UINT Adapter, CONST D3DDISPLAYMODEFILTER *pFilter) PURE;
    STDMETHOD(EnumAdapterModesEx)(THIS_ UINT Adapter, CONST D3DDISPLAYMODEFILTER *pFilter, UINT Mode, D3DDISPLAYMODEEX* pMode) PURE;
    STDMETHOD(GetAdapterDisplayModeEx)(THIS_ UINT Adapter, D3DDISPLAYMODEEX *pMode, D3DDISPLAYROTATION *pRotation);
    STDMETHOD(CreateDeviceEx)(THIS_ UINT Adapter, D3DDEVTYPE DeviceType, HWND hFocusWindow, DWORD BehaviorFlags, D3DPRESENT_PARAMETERS* pPresentationParameters, D3DDISPLAYMODEEX* pFullscreenDisplayMode, struct IDirect3DDevice9Ex **ppReturnedDeviceInterface) PURE;
    STDMETHOD(GetAdapterLUID)(THIS_ UINT Adatper, LUID *pLUID) PURE;
};
#undef INTERFACE

#if !defined(__cplusplus) || defined(CINTERFACE)
/*** IUnknown methods ***/
#define IDirect3D9Ex_QueryInterface(p,a,b) (p)->lpVtbl->QueryInterface(p,a,b)
#define IDirect3D9Ex_AddRef(p)             (p)->lpVtbl->AddRef(p)
#define IDirect3D9Ex_Release(p)            (p)->lpVtbl->Release(p)
/*** IDirect3D9 methods ***/
#define IDirect3D9Ex_RegisterSoftwareDevice(p,a)                (p)->lpVtbl->RegisterSoftwareDevice(p,a)
#define IDirect3D9Ex_GetAdapterCount(p)                         (p)->lpVtbl->GetAdapterCount(p)
#define IDirect3D9Ex_GetAdapterIdentifier(p,a,b,c)              (p)->lpVtbl->GetAdapterIdentifier(p,a,b,c)
#define IDirect3D9Ex_GetAdapterModeCount(p,a,b)                 (p)->lpVtbl->GetAdapterModeCount(p,a,b)
#define IDirect3D9Ex_EnumAdapterModes(p,a,b,c,d)                (p)->lpVtbl->EnumAdapterModes(p,a,b,c,d)
#define IDirect3D9Ex_GetAdapterDisplayMode(p,a,b)               (p)->lpVtbl->GetAdapterDisplayMode(p,a,b)
#define IDirect3D9Ex_CheckDeviceType(p,a,b,c,d,e)               (p)->lpVtbl->CheckDeviceType(p,a,b,c,d,e)
#define IDirect3D9Ex_CheckDeviceFormat(p,a,b,c,d,e,f)           (p)->lpVtbl->CheckDeviceFormat(p,a,b,c,d,e,f)
#define IDirect3D9Ex_CheckDeviceMultiSampleType(p,a,b,c,d,e,f)  (p)->lpVtbl->CheckDeviceMultiSampleType(p,a,b,c,d,e,f)
#define IDirect3D9Ex_CheckDepthStencilMatch(p,a,b,c,d,e)        (p)->lpVtbl->CheckDepthStencilMatch(p,a,b,c,d,e)
#define IDirect3D9Ex_CheckDeviceFormatConversion(p,a,b,c,d)     (p)->lpVtbl->CheckDeviceFormatConversion(p,a,b,c,d)
#define IDirect3D9Ex_GetDeviceCaps(p,a,b,c)                     (p)->lpVtbl->GetDeviceCaps(p,a,b,c)
#define IDirect3D9Ex_GetAdapterMonitor(p,a)                     (p)->lpVtbl->GetAdapterMonitor(p,a)
#define IDirect3D9Ex_CreateDevice(p,a,b,c,d,e,f)                (p)->lpVtbl->CreateDevice(p,a,b,c,d,e,f)
/*** IDirect3D9Ex methods ***/
#define IDirect3D9Ex_GetAdapterModeCountEx(p,a,b)               (p)->lpVtbl->GetAdapterModeCountEx(p,a,b)
#define IDirect3D9Ex_EnumAdapterModesEx(p,a,b,c,d)              (p)->lpVtbl->EnumAdapterModesEx(p,a,b,c,d)
#define IDirect3D9Ex_GetAdapterDisplayModeEx(p,a,b,c)           (p)->lpVtbl->GetAdapterDisplayModeEx(p,a,b,c)
#define IDirect3D9Ex_CreateDeviceEx(p,a,b,c,d,e,f,g)            (p)->lpVtbl->CreateDeviceEx(p,a,b,c,d,e,f,g)
#define IDirect3D9Ex_GetAdapterLUID(p,a,b)                      (p)->lpVtbl->GetAdapterLUID(p,a,b)
#else
/*** IUnknown methods ***/
#define IDirect3D9Ex_QueryInterface(p,a,b) (p)->QueryInterface(a,b)
#define IDirect3D9Ex_AddRef(p)             (p)->AddRef()
#define IDirect3D9Ex_Release(p)            (p)->Release()
/*** IDirect3D9 methods ***/
#define IDirect3D9Ex_RegisterSoftwareDevice(p,a)                (p)->RegisterSoftwareDevice(a)
#define IDirect3D9Ex_GetAdapterCount(p)                         (p)->GetAdapterCount()
#define IDirect3D9Ex_GetAdapterIdentifier(p,a,b,c)              (p)->GetAdapterIdentifier(a,b,c)
#define IDirect3D9Ex_GetAdapterModeCount(p,a,b)                 (p)->GetAdapterModeCount(a,b)
#define IDirect3D9Ex_EnumAdapterModes(p,a,b,c,d)                (p)->EnumAdapterModes(a,b,c,d)
#define IDirect3D9Ex_GetAdapterDisplayMode(p,a,b)               (p)->GetAdapterDisplayMode(a,b)
#define IDirect3D9Ex_CheckDeviceType(p,a,b,c,d,e)               (p)->CheckDeviceType(a,b,c,d,e)
#define IDirect3D9Ex_CheckDeviceFormat(p,a,b,c,d,e,f)           (p)->CheckDeviceFormat(a,b,c,d,e,f)
#define IDirect3D9Ex_CheckDeviceMultiSampleType(p,a,b,c,d,e,f)  (p)->CheckDeviceMultiSampleType(a,b,c,d,e,f)
#define IDirect3D9Ex_CheckDepthStencilMatch(p,a,b,c,d,e)        (p)->CheckDepthStencilMatch(a,b,c,d,e)
#define IDirect3D9Ex_CheckDeviceFormatConversion(p,a,b,c,d)     (p)->CheckDeviceFormatConversion(a,b,c,d)
#define IDirect3D9Ex_GetDeviceCaps(p,a,b,c)                     (p)->GetDeviceCaps(a,b,c)
#define IDirect3D9Ex_GetAdapterMonitor(p,a)                     (p)->GetAdapterMonitor(a)
#define IDirect3D9Ex_CreateDevice(p,a,b,c,d,e,f)                (p)->CreateDevice(a,b,c,d,e,f)
#endif

#endif /* D3D_DISABLE_9EX */

/*****************************************************************************
 * IDirect3DVolume9 interface
 */
#define INTERFACE IDirect3DVolume9
DECLARE_INTERFACE_(IDirect3DVolume9,IUnknown)
{
    /*** IUnknown methods ***/
    STDMETHOD_(HRESULT,QueryInterface)(THIS_ REFIID riid, void** ppvObject) PURE;
    STDMETHOD_(ULONG,AddRef)(THIS) PURE;
    STDMETHOD_(ULONG,Release)(THIS) PURE;
    /*** IDirect3DVolume9 methods ***/
    STDMETHOD(GetDevice)(THIS_ struct IDirect3DDevice9** ppDevice) PURE;
    STDMETHOD(SetPrivateData)(THIS_ REFGUID refguid, CONST void* pData, DWORD SizeOfData, DWORD Flags) PURE;
    STDMETHOD(GetPrivateData)(THIS_ REFGUID refguid, void* pData, DWORD* pSizeOfData) PURE;
    STDMETHOD(FreePrivateData)(THIS_ REFGUID refguid) PURE;
    STDMETHOD(GetContainer)(THIS_ REFIID riid, void** ppContainer) PURE;
    STDMETHOD(GetDesc)(THIS_ D3DVOLUME_DESC* pDesc) PURE;
    STDMETHOD(LockBox)(THIS_ D3DLOCKED_BOX* pLockedVolume, CONST D3DBOX* pBox, DWORD Flags) PURE;
    STDMETHOD(UnlockBox)(THIS) PURE;
};
#undef INTERFACE

#if !defined(__cplusplus) || defined(CINTERFACE)
/*** IUnknown methods ***/
#define IDirect3DVolume9_QueryInterface(p,a,b)        (p)->lpVtbl->QueryInterface(p,a,b)
#define IDirect3DVolume9_AddRef(p)                    (p)->lpVtbl->AddRef(p)
#define IDirect3DVolume9_Release(p)                   (p)->lpVtbl->Release(p)
/*** IDirect3DVolume9 methods ***/
#define IDirect3DVolume9_GetDevice(p,a)               (p)->lpVtbl->GetDevice(p,a)
#define IDirect3DVolume9_SetPrivateData(p,a,b,c,d)    (p)->lpVtbl->SetPrivateData(p,a,b,c,d)
#define IDirect3DVolume9_GetPrivateData(p,a,b,c)      (p)->lpVtbl->GetPrivateData(p,a,b,c)
#define IDirect3DVolume9_FreePrivateData(p,a)         (p)->lpVtbl->FreePrivateData(p,a)
#define IDirect3DVolume9_GetContainer(p,a,b)          (p)->lpVtbl->GetContainer(p,a,b)
#define IDirect3DVolume9_GetDesc(p,a)                 (p)->lpVtbl->GetDesc(p,a)
#define IDirect3DVolume9_LockBox(p,a,b,c)             (p)->lpVtbl->LockBox(p,a,b,c)
#define IDirect3DVolume9_UnlockBox(p)                 (p)->lpVtbl->UnlockBox(p)
#else
/*** IUnknown methods ***/
#define IDirect3DVolume9_QueryInterface(p,a,b)        (p)->QueryInterface(a,b)
#define IDirect3DVolume9_AddRef(p)                    (p)->AddRef()
#define IDirect3DVolume9_Release(p)                   (p)->Release()
/*** IDirect3DVolume9 methods ***/
#define IDirect3DVolume9_GetDevice(p,a)               (p)->GetDevice(a)
#define IDirect3DVolume9_SetPrivateData(p,a,b,c,d)    (p)->SetPrivateData(a,b,c,d)
#define IDirect3DVolume9_GetPrivateData(p,a,b,c)      (p)->GetPrivateData(a,b,c)
#define IDirect3DVolume9_FreePrivateData(p,a)         (p)->FreePrivateData(a)
#define IDirect3DVolume9_GetContainer(p,a,b)          (p)->GetContainer(a,b)
#define IDirect3DVolume9_GetDesc(p,a)                 (p)->GetDesc(a)
#define IDirect3DVolume9_LockBox(p,a,b,c)             (p)->LockBox(a,b,c)
#define IDirect3DVolume9_UnlockBox(p)                 (p)->UnlockBox()
#endif

/*****************************************************************************
 * IDirect3DSwapChain9 interface
 */
#define INTERFACE IDirect3DSwapChain9
DECLARE_INTERFACE_(IDirect3DSwapChain9,IUnknown)
{
    /*** IUnknown methods ***/
    STDMETHOD_(HRESULT,QueryInterface)(THIS_ REFIID riid, void **ppvObject) PURE;
    STDMETHOD_(ULONG,AddRef)(THIS) PURE;
    STDMETHOD_(ULONG,Release)(THIS) PURE;
    /*** IDirect3DSwapChain9 methods ***/
    STDMETHOD(Present)(THIS_ CONST RECT *pSourceRect, CONST RECT *pDestRect, HWND hDestWindowOverride, CONST RGNDATA *pDirtyRegion, DWORD dwFlags) PURE;
    STDMETHOD(GetFrontBufferData)(THIS_ struct IDirect3DSurface9 *pDestSurface) PURE;
    STDMETHOD(GetBackBuffer)(THIS_ UINT iBackBuffer, D3DBACKBUFFER_TYPE Type, struct IDirect3DSurface9 **ppBackBuffer) PURE;
    STDMETHOD(GetRasterStatus)(THIS_ D3DRASTER_STATUS *pRasterStatus) PURE;
    STDMETHOD(GetDisplayMode)(THIS_ D3DDISPLAYMODE *pMode) PURE;
    STDMETHOD(GetDevice)(THIS_ struct IDirect3DDevice9 **ppDevice) PURE;
    STDMETHOD(GetPresentParameters)(THIS_ D3DPRESENT_PARAMETERS *pPresentationParameters) PURE;
};
#undef INTERFACE

#if !defined(__cplusplus) || defined(CINTERFACE)
/*** IUnknown methods ***/
#define IDirect3DSwapChain9_QueryInterface(p,a,b)        (p)->lpVtbl->QueryInterface(p,a,b)
#define IDirect3DSwapChain9_AddRef(p)                    (p)->lpVtbl->AddRef(p)
#define IDirect3DSwapChain9_Release(p)                   (p)->lpVtbl->Release(p)
/*** IDirect3DSwapChain9 methods ***/
#define IDirect3DSwapChain9_Present(p,a,b,c,d,e)         (p)->lpVtbl->Present(p,a,b,c,d,e)
#define IDirect3DSwapChain9_GetFrontBufferData(p,a)      (p)->lpVtbl->GetFrontBufferData(p,a)
#define IDirect3DSwapChain9_GetBackBuffer(p,a,b,c)       (p)->lpVtbl->GetBackBuffer(p,a,b,c)
#define IDirect3DSwapChain9_GetRasterStatus(p,a)         (p)->lpVtbl->GetRasterStatus(p,a)
#define IDirect3DSwapChain9_GetDisplayMode(p,a)          (p)->lpVtbl->GetDisplayMode(p,a)
#define IDirect3DSwapChain9_GetDevice(p,a)               (p)->lpVtbl->GetDevice(p,a)
#define IDirect3DSwapChain9_GetPresentParameters(p,a)    (p)->lpVtbl->GetPresentParameters(p,a)
#else
/*** IUnknown methods ***/
#define IDirect3DSwapChain9_QueryInterface(p,a,b)        (p)->QueryInterface(a,b)
#define IDirect3DSwapChain9_AddRef(p)                    (p)->AddRef()
#define IDirect3DSwapChain9_Release(p)                   (p)->Release()
/*** IDirect3DSwapChain9 methods ***/
#define IDirect3DSwapChain9_Present(p,a,b,c,d,e)         (p)->Present(a,b,c,d,e)
#define IDirect3DSwapChain9_GetFrontBufferData(p,a)      (p)->GetFrontBufferData(a)
#define IDirect3DSwapChain9_GetBackBuffer(p,a,b,c)       (p)->GetBackBuffer(a,b,c)
#define IDirect3DSwapChain9_GetRasterStatus(p,a)         (p)->GetRasterStatus(a)
#define IDirect3DSwapChain9_GetDisplayMode(p,a)          (p)->GetDisplayMode(a)
#define IDirect3DSwapChain9_GetDevice(p,a)               (p)->GetDevice(a)
#define IDirect3DSwapChain9_GetPresentParameters(p,a)    (p)->GetPresentParameters(a)
#endif

#if !defined(D3D_DISABLE_9EX)

/*****************************************************************************
 * IDirect3DSwapChain9Ex interface
 */
#define INTERFACE IDirect3DSwapChain9Ex
DECLARE_INTERFACE_(IDirect3DSwapChain9Ex,IDirect3DSwapChain9)
{
    /*** IUnknown methods ***/
    STDMETHOD_(HRESULT,QueryInterface)(THIS_ REFIID riid, void **ppvObject) PURE;
    STDMETHOD_(ULONG,AddRef)(THIS) PURE;
    STDMETHOD_(ULONG,Release)(THIS) PURE;
    /*** IDirect3DSwapChain9 methods ***/
    STDMETHOD(Present)(THIS_ CONST RECT *pSourceRect, CONST RECT *pDestRect, HWND hDestWindowOverride, CONST RGNDATA *pDirtyRegion, DWORD dwFlags) PURE;
    STDMETHOD(GetFrontBufferData)(THIS_ struct IDirect3DSurface9 *pDestSurface) PURE;
    STDMETHOD(GetBackBuffer)(THIS_ UINT iBackBuffer, D3DBACKBUFFER_TYPE Type, struct IDirect3DSurface9 **ppBackBuffer) PURE;
    STDMETHOD(GetRasterStatus)(THIS_ D3DRASTER_STATUS *pRasterStatus) PURE;
    STDMETHOD(GetDisplayMode)(THIS_ D3DDISPLAYMODE *pMode) PURE;
    STDMETHOD(GetDevice)(THIS_ struct IDirect3DDevice9 **ppDevice) PURE;
    STDMETHOD(GetPresentParameters)(THIS_ D3DPRESENT_PARAMETERS *pPresentationParameters) PURE;
    /*** IDirect3DSwapChain9Ex methods ***/
    STDMETHOD(GetLastPresentCount)(THIS_ UINT *pLastPresentCount) PURE;
    STDMETHOD(GetPresentStats)(THIS_ D3DPRESENTSTATS *pPresentationStatistics) PURE;
    STDMETHOD(GetDisplayModeEx)(THIS_ D3DDISPLAYMODEEX *pMode, D3DDISPLAYROTATION *pRotation) PURE;
};
#undef INTERFACE

#if !defined(__cplusplus) || defined(CINTERFACE)
/*** IUnknown methods ***/
#define IDirect3DSwapChain9Ex_QueryInterface(p,a,b)      (p)->lpVtbl->QueryInterface(p,a,b)
#define IDirect3DSwapChain9Ex_AddRef(p)                  (p)->lpVtbl->AddRef(p)
#define IDirect3DSwapChain9Ex_Release(p)                 (p)->lpVtbl->Release(p)
/*** IDirect3DSwapChain9 methods ***/
#define IDirect3DSwapChain9Ex_Present(p,a,b,c,d,e)       (p)->lpVtbl->Present(p,a,b,c,d,e)
#define IDirect3DSwapChain9Ex_GetFrontBufferData(p,a)    (p)->lpVtbl->GetFrontBufferData(p,a)
#define IDirect3DSwapChain9EX_GetBackBuffer(p,a,b,c)     (p)->lpVtbl->GetBackBuffer(p,a,b,c)
#define IDirect3DSwapChain9EX_GetRasterStatus(p,a)       (p)->lpVtbl->GetRasterStatus(p,a)
#define IDirect3DSwapChain9Ex_GetDisplayMode(p,a)        (p)->lpVtbl->GetDisplayMode(p,a)
#define IDirect3DSwapChain9Ex_GetDevice(p,a)             (p)->lpVtbl->GetDevice(p,a)
#define IDirect3DSwapChain9Ex_GetPresentParameters(p,a)  (p)->lpVtbl->GetPresentParameters(p,a)
/*** IDirect3DSwapChain9Ex methods ***/
#define IDirect3DSwapChain9Ex_GetLastPresentCount(p,a)   (p)->lpVtbl->GetLastPresentCount(p,a)
#define IDirect3DSwapChain9Ex_GetPresentStats(p,a)       (p)->lpVtbl->GetPresentStats(p,a)
#define IDirect3DSwapChain9Ex_GetDisplayModeEx(p,a,b)    (p)->lpVtbl->GetDisplayModeEx(p,a,b)
#else
/*** IUnknown methods ***/
#define IDirect3DSwapChain9Ex_QueryInterface(p,a,b)      (p)->QueryInterface(a,b)
#define IDirect3DSwapChain9Ex_AddRef(p)                  (p)->AddRef()
#define IDirect3DSwapChain9Ex_Release(p)                 (p)->Release()
/*** IDirect3DSwapChain9 methods ***/
#define IDirect3DSwapChain9Ex_Present(p,a,b,c,d,e)       (p)->Present(a,b,c,d,e)
#define IDirect3DSwapChain9Ex_GetFrontBufferData(p,a)    (p)->GetFrontBufferData(a)
#define IDirect3DSwapChain9Ex_GetBackBuffer(p,a,b,c)     (p)->GetBackBuffer(a,b,c)
#define IDirect3DSwapChain9Ex_GetRasterStatus(p,a)       (p)->GetRasterStatus(a)
#define IDirect3DSwapChain9Ex_GetDisplayMode(p,a)        (p)->GetDisplayMode(a)
#define IDirect3DSwapChain9Ex_GetDevice(p,a)             (p)->GetDevice(a)
#define IDirect3DSwapChain9Ex_GetPresentParameters(p,a)  (p)->GetPresentParameters(a)
/*** IDirect3DSwapChain9Ex methods ***/
#define IDirect3DSwapChain9Ex_GetLastPresentCount(p,a)   (p)->GetLastPresentCount(a)
#define IDirect3DSwapChain9Ex_GetPresentStats(p,a)       (p)->GetPresentStats(a)
#define IDirect3DSwapChain9Ex_GetDisplayModeEx(p,a,b)    (p)->GetDisplayModeEx(a,b)
#endif

#endif /* D3D_DISABLE_9EX */

/*****************************************************************************
 * IDirect3DResource9 interface
 */
#define INTERFACE IDirect3DResource9
DECLARE_INTERFACE_(IDirect3DResource9,IUnknown)
{
    /*** IUnknown methods ***/
    STDMETHOD_(HRESULT,QueryInterface)(THIS_ REFIID riid, void** ppvObject) PURE;
    STDMETHOD_(ULONG,AddRef)(THIS) PURE;
    STDMETHOD_(ULONG,Release)(THIS) PURE;
    /*** IDirect3DResource9 methods ***/
    STDMETHOD(GetDevice)(THIS_ struct IDirect3DDevice9** ppDevice) PURE;
    STDMETHOD(SetPrivateData)(THIS_ REFGUID refguid, CONST void* pData, DWORD SizeOfData, DWORD Flags) PURE;
    STDMETHOD(GetPrivateData)(THIS_ REFGUID refguid, void* pData, DWORD* pSizeOfData) PURE;
    STDMETHOD(FreePrivateData)(THIS_ REFGUID refguid) PURE;
    STDMETHOD_(DWORD, SetPriority)(THIS_ DWORD PriorityNew) PURE;
    STDMETHOD_(DWORD, GetPriority)(THIS) PURE;
    STDMETHOD_(void, PreLoad)(THIS) PURE;
    STDMETHOD_(D3DRESOURCETYPE, GetType)(THIS) PURE;
};
#undef INTERFACE

#if !defined(__cplusplus) || defined(CINTERFACE)
/*** IUnknown methods ***/
#define IDirect3DResource9_QueryInterface(p,a,b)        (p)->lpVtbl->QueryInterface(p,a,b)
#define IDirect3DResource9_AddRef(p)                    (p)->lpVtbl->AddRef(p)
#define IDirect3DResource9_Release(p)                   (p)->lpVtbl->Release(p)
/*** IDirect3DResource9 methods ***/
#define IDirect3DResource9_GetDevice(p,a)               (p)->lpVtbl->GetDevice(p,a)
#define IDirect3DResource9_SetPrivateData(p,a,b,c,d)    (p)->lpVtbl->SetPrivateData(p,a,b,c,d)
#define IDirect3DResource9_GetPrivateData(p,a,b,c)      (p)->lpVtbl->GetPrivateData(p,a,b,c)
#define IDirect3DResource9_FreePrivateData(p,a)         (p)->lpVtbl->FreePrivateData(p,a)
#define IDirect3DResource9_SetPriority(p,a)             (p)->lpVtbl->SetPriority(p,a)
#define IDirect3DResource9_GetPriority(p)               (p)->lpVtbl->GetPriority(p)
#define IDirect3DResource9_PreLoad(p)                   (p)->lpVtbl->PreLoad(p)
#define IDirect3DResource9_GetType(p)                   (p)->lpVtbl->GetType(p)
#else
/*** IUnknown methods ***/
#define IDirect3DResource9_QueryInterface(p,a,b)        (p)->QueryInterface(a,b)
#define IDirect3DResource9_AddRef(p)                    (p)->AddRef()
#define IDirect3DResource9_Release(p)                   (p)->Release()
/*** IDirect3DResource9 methods ***/
#define IDirect3DResource9_GetDevice(p,a)               (p)->GetDevice(a)
#define IDirect3DResource9_SetPrivateData(p,a,b,c,d)    (p)->SetPrivateData(a,b,c,d)
#define IDirect3DResource9_GetPrivateData(p,a,b,c)      (p)->GetPrivateData(a,b,c)
#define IDirect3DResource9_FreePrivateData(p,a)         (p)->FreePrivateData(a)
#define IDirect3DResource9_SetPriority(p,a)             (p)->SetPriority(a)
#define IDirect3DResource9_GetPriority(p)               (p)->GetPriority()
#define IDirect3DResource9_PreLoad(p)                   (p)->PreLoad()
#define IDirect3DResource9_GetType(p)                   (p)->GetType()
#endif

/*****************************************************************************
 * IDirect3DSurface9 interface
 */
#define INTERFACE IDirect3DSurface9
DECLARE_INTERFACE_(IDirect3DSurface9,IDirect3DResource9)
{
    /*** IUnknown methods ***/
    STDMETHOD_(HRESULT,QueryInterface)(THIS_ REFIID riid, void** ppvObject) PURE;
    STDMETHOD_(ULONG,AddRef)(THIS) PURE;
    STDMETHOD_(ULONG,Release)(THIS) PURE;
    /*** IDirect3DResource9 methods ***/
    STDMETHOD(GetDevice)(THIS_ struct IDirect3DDevice9** ppDevice) PURE;
    STDMETHOD(SetPrivateData)(THIS_ REFGUID refguid, CONST void* pData, DWORD SizeOfData, DWORD Flags) PURE;
    STDMETHOD(GetPrivateData)(THIS_ REFGUID refguid, void* pData, DWORD* pSizeOfData) PURE;
    STDMETHOD(FreePrivateData)(THIS_ REFGUID refguid) PURE;
    STDMETHOD_(DWORD, SetPriority)(THIS_ DWORD PriorityNew) PURE;
    STDMETHOD_(DWORD, GetPriority)(THIS) PURE;
    STDMETHOD_(void, PreLoad)(THIS) PURE;
    STDMETHOD_(D3DRESOURCETYPE, GetType)(THIS) PURE;
    /*** IDirect3DSurface9 methods ***/
    STDMETHOD(GetContainer)(THIS_ REFIID riid, void** ppContainer) PURE;
    STDMETHOD(GetDesc)(THIS_ D3DSURFACE_DESC* pDesc) PURE;
    STDMETHOD(LockRect)(THIS_ D3DLOCKED_RECT* pLockedRect, CONST RECT* pRect, DWORD Flags) PURE;
    STDMETHOD(UnlockRect)(THIS) PURE;
    STDMETHOD(GetDC)(THIS_ HDC* phdc) PURE;
    STDMETHOD(ReleaseDC)(THIS_ HDC hdc) PURE;
};
#undef INTERFACE

#if !defined(__cplusplus) || defined(CINTERFACE)
/*** IUnknown methods ***/
#define IDirect3DSurface9_QueryInterface(p,a,b)        (p)->lpVtbl->QueryInterface(p,a,b)
#define IDirect3DSurface9_AddRef(p)                    (p)->lpVtbl->AddRef(p)
#define IDirect3DSurface9_Release(p)                   (p)->lpVtbl->Release(p)
/*** IDirect3DSurface9 methods: IDirect3DResource9 ***/
#define IDirect3DSurface9_GetDevice(p,a)               (p)->lpVtbl->GetDevice(p,a)
#define IDirect3DSurface9_SetPrivateData(p,a,b,c,d)    (p)->lpVtbl->SetPrivateData(p,a,b,c,d)
#define IDirect3DSurface9_GetPrivateData(p,a,b,c)      (p)->lpVtbl->GetPrivateData(p,a,b,c)
#define IDirect3DSurface9_FreePrivateData(p,a)         (p)->lpVtbl->FreePrivateData(p,a)
#define IDirect3DSurface9_SetPriority(p,a)             (p)->lpVtbl->SetPriority(p,a)
#define IDirect3DSurface9_GetPriority(p)               (p)->lpVtbl->GetPriority(p)
#define IDirect3DSurface9_PreLoad(p)                   (p)->lpVtbl->PreLoad(p)
#define IDirect3DSurface9_GetType(p)                   (p)->lpVtbl->GetType(p)
/*** IDirect3DSurface9 methods ***/
#define IDirect3DSurface9_GetContainer(p,a,b)          (p)->lpVtbl->GetContainer(p,a,b)
#define IDirect3DSurface9_GetDesc(p,a)                 (p)->lpVtbl->GetDesc(p,a)
#define IDirect3DSurface9_LockRect(p,a,b,c)            (p)->lpVtbl->LockRect(p,a,b,c)
#define IDirect3DSurface9_UnlockRect(p)                (p)->lpVtbl->UnlockRect(p)
#define IDirect3DSurface9_GetDC(p,a)                   (p)->lpVtbl->GetDC(p,a)
#define IDirect3DSurface9_ReleaseDC(p,a)               (p)->lpVtbl->ReleaseDC(p,a)
#else
/*** IUnknown methods ***/
#define IDirect3DSurface9_QueryInterface(p,a,b)        (p)->QueryInterface(a,b)
#define IDirect3DSurface9_AddRef(p)                    (p)->AddRef()
#define IDirect3DSurface9_Release(p)                   (p)->Release()
/*** IDirect3DSurface9 methods: IDirect3DResource9 ***/
#define IDirect3DSurface9_GetDevice(p,a)               (p)->GetDevice(a)
#define IDirect3DSurface9_SetPrivateData(p,a,b,c,d)    (p)->SetPrivateData(a,b,c,d)
#define IDirect3DSurface9_GetPrivateData(p,a,b,c)      (p)->GetPrivateData(a,b,c)
#define IDirect3DSurface9_FreePrivateData(p,a)         (p)->FreePrivateData(a)
#define IDirect3DSurface9_SetPriority(p,a)             (p)->SetPriority(a)
#define IDirect3DSurface9_GetPriority(p)               (p)->GetPriority()
#define IDirect3DSurface9_PreLoad(p)                   (p)->PreLoad()
#define IDirect3DSurface9_GetType(p)                   (p)->GetType()
/*** IDirect3DSurface9 methods ***/
#define IDirect3DSurface9_GetContainer(p,a,b)          (p)->GetContainer(a,b)
#define IDirect3DSurface9_GetDesc(p,a)                 (p)->GetDesc(a)
#define IDirect3DSurface9_LockRect(p,a,b,c)            (p)->LockRect(a,b,c)
#define IDirect3DSurface9_UnlockRect(p)                (p)->UnlockRect()
#define IDirect3DSurface9_GetDC(p,a)                   (p)->GetDC(a)
#define IDirect3DSurface9_ReleaseDC(p,a)               (p)->ReleaseDC(a)
#endif

/*****************************************************************************
 * IDirect3DVertexBuffer9 interface
 */
#define INTERFACE IDirect3DVertexBuffer9
DECLARE_INTERFACE_(IDirect3DVertexBuffer9,IDirect3DResource9)
{
    /*** IUnknown methods ***/
    STDMETHOD_(HRESULT,QueryInterface)(THIS_ REFIID riid, void** ppvObject) PURE;
    STDMETHOD_(ULONG,AddRef)(THIS) PURE;
    STDMETHOD_(ULONG,Release)(THIS) PURE;
    /*** IDirect3DResource9 methods ***/
    STDMETHOD(GetDevice)(THIS_ struct IDirect3DDevice9** ppDevice) PURE;
    STDMETHOD(SetPrivateData)(THIS_ REFGUID refguid, CONST void* pData, DWORD SizeOfData, DWORD Flags) PURE;
    STDMETHOD(GetPrivateData)(THIS_ REFGUID refguid, void* pData, DWORD* pSizeOfData) PURE;
    STDMETHOD(FreePrivateData)(THIS_ REFGUID refguid) PURE;
    STDMETHOD_(DWORD, SetPriority)(THIS_ DWORD PriorityNew) PURE;
    STDMETHOD_(DWORD, GetPriority)(THIS) PURE;
    STDMETHOD_(void, PreLoad)(THIS) PURE;
    STDMETHOD_(D3DRESOURCETYPE, GetType)(THIS) PURE;
    /*** IDirect3DVertexBuffer9 methods ***/
    STDMETHOD(Lock)(THIS_ UINT OffsetToLock, UINT SizeToLock, void** ppbData, DWORD Flags) PURE;
    STDMETHOD(Unlock)(THIS) PURE;
    STDMETHOD(GetDesc)(THIS_ D3DVERTEXBUFFER_DESC* pDesc) PURE;
};
#undef INTERFACE

#if !defined(__cplusplus) || defined(CINTERFACE)
/*** IUnknown methods ***/
#define IDirect3DVertexBuffer9_QueryInterface(p,a,b)        (p)->lpVtbl->QueryInterface(p,a,b)
#define IDirect3DVertexBuffer9_AddRef(p)                    (p)->lpVtbl->AddRef(p)
#define IDirect3DVertexBuffer9_Release(p)                   (p)->lpVtbl->Release(p)
/*** IDirect3DVertexBuffer9 methods: IDirect3DResource9 ***/
#define IDirect3DVertexBuffer9_GetDevice(p,a)               (p)->lpVtbl->GetDevice(p,a)
#define IDirect3DVertexBuffer9_SetPrivateData(p,a,b,c,d)    (p)->lpVtbl->SetPrivateData(p,a,b,c,d)
#define IDirect3DVertexBuffer9_GetPrivateData(p,a,b,c)      (p)->lpVtbl->GetPrivateData(p,a,b,c)
#define IDirect3DVertexBuffer9_FreePrivateData(p,a)         (p)->lpVtbl->FreePrivateData(p,a)
#define IDirect3DVertexBuffer9_SetPriority(p,a)             (p)->lpVtbl->SetPriority(p,a)
#define IDirect3DVertexBuffer9_GetPriority(p)               (p)->lpVtbl->GetPriority(p)
#define IDirect3DVertexBuffer9_PreLoad(p)                   (p)->lpVtbl->PreLoad(p)
#define IDirect3DVertexBuffer9_GetType(p)                   (p)->lpVtbl->GetType(p)
/*** IDirect3DVertexBuffer9 methods ***/
#define IDirect3DVertexBuffer9_Lock(p,a,b,c,d)              (p)->lpVtbl->Lock(p,a,b,c,d)
#define IDirect3DVertexBuffer9_Unlock(p)                    (p)->lpVtbl->Unlock(p)
#define IDirect3DVertexBuffer9_GetDesc(p,a)                 (p)->lpVtbl->GetDesc(p,a)
#else
/*** IUnknown methods ***/
#define IDirect3DVertexBuffer9_QueryInterface(p,a,b)        (p)->QueryInterface(a,b)
#define IDirect3DVertexBuffer9_AddRef(p)                    (p)->AddRef()
#define IDirect3DVertexBuffer9_Release(p)                   (p)->Release()
/*** IDirect3DVertexBuffer9 methods: IDirect3DResource9 ***/
#define IDirect3DVertexBuffer9_GetDevice(p,a)               (p)->GetDevice(a)
#define IDirect3DVertexBuffer9_SetPrivateData(p,a,b,c,d)    (p)->SetPrivateData(a,b,c,d)
#define IDirect3DVertexBuffer9_GetPrivateData(p,a,b,c)      (p)->GetPrivateData(a,b,c)
#define IDirect3DVertexBuffer9_FreePrivateData(p,a)         (p)->FreePrivateData(a)
#define IDirect3DVertexBuffer9_SetPriority(p,a)             (p)->SetPriority(a)
#define IDirect3DVertexBuffer9_GetPriority(p)               (p)->GetPriority()
#define IDirect3DVertexBuffer9_PreLoad(p)                   (p)->PreLoad()
#define IDirect3DVertexBuffer9_GetType(p)                   (p)->GetType()
/*** IDirect3DVertexBuffer9 methods ***/
#define IDirect3DVertexBuffer9_Lock(p,a,b,c,d)              (p)->Lock(a,b,c,d)
#define IDirect3DVertexBuffer9_Unlock(p)                    (p)->Unlock()
#define IDirect3DVertexBuffer9_GetDesc(p,a)                 (p)->GetDesc(a)
#endif

/*****************************************************************************
 * IDirect3DIndexBuffer9 interface
 */
#define INTERFACE IDirect3DIndexBuffer9
DECLARE_INTERFACE_(IDirect3DIndexBuffer9,IDirect3DResource9)
{
    /*** IUnknown methods ***/
    STDMETHOD_(HRESULT,QueryInterface)(THIS_ REFIID riid, void** ppvObject) PURE;
    STDMETHOD_(ULONG,AddRef)(THIS) PURE;
    STDMETHOD_(ULONG,Release)(THIS) PURE;
    /*** IDirect3DResource9 methods ***/
    STDMETHOD(GetDevice)(THIS_ struct IDirect3DDevice9** ppDevice) PURE;
    STDMETHOD(SetPrivateData)(THIS_ REFGUID refguid, CONST void* pData, DWORD SizeOfData, DWORD Flags) PURE;
    STDMETHOD(GetPrivateData)(THIS_ REFGUID refguid, void* pData, DWORD* pSizeOfData) PURE;
    STDMETHOD(FreePrivateData)(THIS_ REFGUID refguid) PURE;
    STDMETHOD_(DWORD, SetPriority)(THIS_ DWORD PriorityNew) PURE;
    STDMETHOD_(DWORD, GetPriority)(THIS) PURE;
    STDMETHOD_(void, PreLoad)(THIS) PURE;
    STDMETHOD_(D3DRESOURCETYPE, GetType)(THIS) PURE;
    /*** IDirect3DIndexBuffer9 methods ***/
    STDMETHOD(Lock)(THIS_ UINT OffsetToLock, UINT SizeToLock, void** ppbData, DWORD Flags) PURE;
    STDMETHOD(Unlock)(THIS) PURE;
    STDMETHOD(GetDesc)(THIS_ D3DINDEXBUFFER_DESC* pDesc) PURE;
};
#undef INTERFACE

#if !defined(__cplusplus) || defined(CINTERFACE)
/*** IUnknown methods ***/
#define IDirect3DIndexBuffer9_QueryInterface(p,a,b)        (p)->lpVtbl->QueryInterface(p,a,b)
#define IDirect3DIndexBuffer9_AddRef(p)                    (p)->lpVtbl->AddRef(p)
#define IDirect3DIndexBuffer9_Release(p)                   (p)->lpVtbl->Release(p)
/*** IDirect3DIndexBuffer9 methods: IDirect3DResource9 ***/
#define IDirect3DIndexBuffer9_GetDevice(p,a)               (p)->lpVtbl->GetDevice(p,a)
#define IDirect3DIndexBuffer9_SetPrivateData(p,a,b,c,d)    (p)->lpVtbl->SetPrivateData(p,a,b,c,d)
#define IDirect3DIndexBuffer9_GetPrivateData(p,a,b,c)      (p)->lpVtbl->GetPrivateData(p,a,b,c)
#define IDirect3DIndexBuffer9_FreePrivateData(p,a)         (p)->lpVtbl->FreePrivateData(p,a)
#define IDirect3DIndexBuffer9_SetPriority(p,a)             (p)->lpVtbl->SetPriority(p,a)
#define IDirect3DIndexBuffer9_GetPriority(p)               (p)->lpVtbl->GetPriority(p)
#define IDirect3DIndexBuffer9_PreLoad(p)                   (p)->lpVtbl->PreLoad(p)
#define IDirect3DIndexBuffer9_GetType(p)                   (p)->lpVtbl->GetType(p)
/*** IDirect3DIndexBuffer9 methods ***/
#define IDirect3DIndexBuffer9_Lock(p,a,b,c,d)              (p)->lpVtbl->Lock(p,a,b,c,d)
#define IDirect3DIndexBuffer9_Unlock(p)                    (p)->lpVtbl->Unlock(p)
#define IDirect3DIndexBuffer9_GetDesc(p,a)                 (p)->lpVtbl->GetDesc(p,a)
#else
/*** IUnknown methods ***/
#define IDirect3DIndexBuffer9_QueryInterface(p,a,b)        (p)->QueryInterface(a,b)
#define IDirect3DIndexBuffer9_AddRef(p)                    (p)->AddRef()
#define IDirect3DIndexBuffer9_Release(p)                   (p)->Release()
/*** IDirect3DIndexBuffer9 methods: IDirect3DResource9 ***/
#define IDirect3DIndexBuffer9_GetDevice(p,a)               (p)->GetDevice(a)
#define IDirect3DIndexBuffer9_SetPrivateData(p,a,b,c,d)    (p)->SetPrivateData(a,b,c,d)
#define IDirect3DIndexBuffer9_GetPrivateData(p,a,b,c)      (p)->GetPrivateData(a,b,c)
#define IDirect3DIndexBuffer9_FreePrivateData(p,a)         (p)->FreePrivateData(a)
#define IDirect3DIndexBuffer9_SetPriority(p,a)             (p)->SetPriority(a)
#define IDirect3DIndexBuffer9_GetPriority(p)               (p)->GetPriority()
#define IDirect3DIndexBuffer9_PreLoad(p)                   (p)->PreLoad()
#define IDirect3DIndexBuffer9_GetType(p)                   (p)->GetType()
/*** IDirect3DIndexBuffer9 methods ***/
#define IDirect3DIndexBuffer9_Lock(p,a,b,c,d)              (p)->Lock(a,b,c,d)
#define IDirect3DIndexBuffer9_Unlock(p)                    (p)->Unlock()
#define IDirect3DIndexBuffer9_GetDesc(p,a)                 (p)->GetDesc(a)
#endif

/*****************************************************************************
 * IDirect3DBaseTexture9 interface
 */
#define INTERFACE IDirect3DBaseTexture9
DECLARE_INTERFACE_(IDirect3DBaseTexture9,IDirect3DResource9)
{
    /*** IUnknown methods ***/
    STDMETHOD_(HRESULT,QueryInterface)(THIS_ REFIID riid, void** ppvObject) PURE;
    STDMETHOD_(ULONG,AddRef)(THIS) PURE;
    STDMETHOD_(ULONG,Release)(THIS) PURE;
    /*** IDirect3DResource9 methods ***/
    STDMETHOD(GetDevice)(THIS_ struct IDirect3DDevice9** ppDevice) PURE;
    STDMETHOD(SetPrivateData)(THIS_ REFGUID refguid, CONST void* pData, DWORD SizeOfData, DWORD Flags) PURE;
    STDMETHOD(GetPrivateData)(THIS_ REFGUID refguid, void* pData, DWORD* pSizeOfData) PURE;
    STDMETHOD(FreePrivateData)(THIS_ REFGUID refguid) PURE;
    STDMETHOD_(DWORD, SetPriority)(THIS_ DWORD PriorityNew) PURE;
    STDMETHOD_(DWORD, GetPriority)(THIS) PURE;
    STDMETHOD_(void, PreLoad)(THIS) PURE;
    STDMETHOD_(D3DRESOURCETYPE, GetType)(THIS) PURE;
    /*** IDirect3DBaseTexture9 methods ***/
    STDMETHOD_(DWORD, SetLOD)(THIS_ DWORD LODNew) PURE;
    STDMETHOD_(DWORD, GetLOD)(THIS) PURE;
    STDMETHOD_(DWORD, GetLevelCount)(THIS) PURE;
    STDMETHOD(SetAutoGenFilterType)(THIS_ D3DTEXTUREFILTERTYPE FilterType) PURE;
    STDMETHOD_(D3DTEXTUREFILTERTYPE, GetAutoGenFilterType)(THIS) PURE;
    STDMETHOD_(void, GenerateMipSubLevels)(THIS) PURE;
};
#undef INTERFACE

#if !defined(__cplusplus) || defined(CINTERFACE)
/*** IUnknown methods ***/
#define IDirect3DBaseTexture9_QueryInterface(p,a,b)  (p)->lpVtbl->QueryInterface(p,a,b)
#define IDirect3DBaseTexture9_AddRef(p)              (p)->lpVtbl->AddRef(p)
#define IDirect3DBaseTexture9_Release(p)             (p)->lpVtbl->Release(p)
/*** IDirect3DBaseTexture9 methods: IDirect3DResource9 ***/
#define IDirect3DBaseTexture9_GetDevice(p,a)             (p)->lpVtbl->GetDevice(p,a)
#define IDirect3DBaseTexture9_SetPrivateData(p,a,b,c,d)  (p)->lpVtbl->SetPrivateData(p,a,b,c,d)
#define IDirect3DBaseTexture9_GetPrivateData(p,a,b,c)    (p)->lpVtbl->GetPrivateData(p,a,b,c)
#define IDirect3DBaseTexture9_FreePrivateData(p,a)       (p)->lpVtbl->FreePrivateData(p,a)
#define IDirect3DBaseTexture9_SetPriority(p,a)           (p)->lpVtbl->SetPriority(p,a)
#define IDirect3DBaseTexture9_GetPriority(p)             (p)->lpVtbl->GetPriority(p)
#define IDirect3DBaseTexture9_PreLoad(p)                 (p)->lpVtbl->PreLoad(p)
#define IDirect3DBaseTexture9_GetType(p)                 (p)->lpVtbl->GetType(p)
/*** IDirect3DBaseTexture9 methods ***/
#define IDirect3DBaseTexture9_SetLOD(p,a)                (p)->lpVtbl->SetLOD(p,a)
#define IDirect3DBaseTexture9_GetLOD(p)                  (p)->lpVtbl->GetLOD(p)
#define IDirect3DBaseTexture9_GetLevelCount(p)           (p)->lpVtbl->GetLevelCount(p)
#define IDirect3DBaseTexture9_SetAutoGenFilterType(p,a)  (p)->lpVtbl->SetAutoGenFilterType(p,a)
#define IDirect3DBaseTexture9_GetAutoGenFilterType(p)    (p)->lpVtbl->GetAutoGenFilterType(p)
#define IDirect3DBaseTexture9_GenerateMipSubLevels(p)    (p)->lpVtbl->GenerateMipSubLevels(p)
#else
/*** IUnknown methods ***/
#define IDirect3DBaseTexture9_QueryInterface(p,a,b)  (p)->QueryInterface(a,b)
#define IDirect3DBaseTexture9_AddRef(p)              (p)->AddRef()
#define IDirect3DBaseTexture9_Release(p)             (p)->Release()
/*** IDirect3DBaseTexture9 methods: IDirect3DResource9 ***/
#define IDirect3DBaseTexture9_GetDevice(p,a)             (p)->GetDevice(a)
#define IDirect3DBaseTexture9_SetPrivateData(p,a,b,c,d)  (p)->SetPrivateData(a,b,c,d)
#define IDirect3DBaseTexture9_GetPrivateData(p,a,b,c)    (p)->GetPrivateData(a,b,c)
#define IDirect3DBaseTexture9_FreePrivateData(p,a)       (p)->FreePrivateData(a)
#define IDirect3DBaseTexture9_SetPriority(p,a)           (p)->SetPriority(a)
#define IDirect3DBaseTexture9_GetPriority(p)             (p)->GetPriority()
#define IDirect3DBaseTexture9_PreLoad(p)                 (p)->PreLoad()
#define IDirect3DBaseTexture9_GetType(p)                 (p)->GetType()
/*** IDirect3DBaseTexture9 methods ***/
#define IDirect3DBaseTexture9_SetLOD(p,a)                (p)->SetLOD(a)
#define IDirect3DBaseTexture9_GetLOD(p)                  (p)->GetLOD()
#define IDirect3DBaseTexture9_GetLevelCount(p)           (p)->GetLevelCount()
#define IDirect3DBaseTexture9_SetAutoGenFilterType(p,a)  (p)->SetAutoGenFilterType(a)
#define IDirect3DBaseTexture9_GetAutoGenFilterType(p)    (p)->GetAutoGenFilterType()
#define IDirect3DBaseTexture9_GenerateMipSubLevels(p)    (p)->GenerateMipSubLevels()
#endif

/*****************************************************************************
 * IDirect3DCubeTexture9 interface
 */
#define INTERFACE IDirect3DCubeTexture9
DECLARE_INTERFACE_(IDirect3DCubeTexture9,IDirect3DBaseTexture9)
{
    /*** IUnknown methods ***/
    STDMETHOD_(HRESULT,QueryInterface)(THIS_ REFIID riid, void** ppvObject) PURE;
    STDMETHOD_(ULONG,AddRef)(THIS) PURE;
    STDMETHOD_(ULONG,Release)(THIS) PURE;
    /*** IDirect3DResource9 methods ***/
    STDMETHOD(GetDevice)(THIS_ struct IDirect3DDevice9** ppDevice) PURE;
    STDMETHOD(SetPrivateData)(THIS_ REFGUID refguid, CONST void* pData, DWORD SizeOfData, DWORD Flags) PURE;
    STDMETHOD(GetPrivateData)(THIS_ REFGUID refguid, void* pData, DWORD* pSizeOfData) PURE;
    STDMETHOD(FreePrivateData)(THIS_ REFGUID refguid) PURE;
    STDMETHOD_(DWORD, SetPriority)(THIS_ DWORD PriorityNew) PURE;
    STDMETHOD_(DWORD, GetPriority)(THIS) PURE;
    STDMETHOD_(void, PreLoad)(THIS) PURE;
    STDMETHOD_(D3DRESOURCETYPE, GetType)(THIS) PURE;
    /*** IDirect3DBaseTexture9 methods ***/
    STDMETHOD_(DWORD, SetLOD)(THIS_ DWORD LODNew) PURE;
    STDMETHOD_(DWORD, GetLOD)(THIS) PURE;
    STDMETHOD_(DWORD, GetLevelCount)(THIS) PURE;
    STDMETHOD(SetAutoGenFilterType)(THIS_ D3DTEXTUREFILTERTYPE FilterType) PURE;
    STDMETHOD_(D3DTEXTUREFILTERTYPE, GetAutoGenFilterType)(THIS) PURE;
    STDMETHOD_(void, GenerateMipSubLevels)(THIS) PURE;
    /*** IDirect3DCubeTexture9 methods ***/
    STDMETHOD(GetLevelDesc)(THIS_ UINT Level,D3DSURFACE_DESC* pDesc) PURE;
    STDMETHOD(GetCubeMapSurface)(THIS_ D3DCUBEMAP_FACES FaceType, UINT Level, IDirect3DSurface9** ppCubeMapSurface) PURE;
    STDMETHOD(LockRect)(THIS_ D3DCUBEMAP_FACES FaceType, UINT Level, D3DLOCKED_RECT* pLockedRect, CONST RECT* pRect, DWORD Flags) PURE;
    STDMETHOD(UnlockRect)(THIS_ D3DCUBEMAP_FACES FaceType, UINT Level) PURE;
    STDMETHOD(AddDirtyRect)(THIS_ D3DCUBEMAP_FACES FaceType, CONST RECT* pDirtyRect) PURE;
};
#undef INTERFACE

#if !defined(__cplusplus) || defined(CINTERFACE)
/*** IUnknown methods ***/
#define IDirect3DCubeTexture9_QueryInterface(p,a,b)       (p)->lpVtbl->QueryInterface(p,a,b)
#define IDirect3DCubeTexture9_AddRef(p)                   (p)->lpVtbl->AddRef(p)
#define IDirect3DCubeTexture9_Release(p)                  (p)->lpVtbl->Release(p)
/*** IDirect3DCubeTexture9 methods: IDirect3DResource9 ***/
#define IDirect3DCubeTexture9_GetDevice(p,a)              (p)->lpVtbl->GetDevice(p,a)
#define IDirect3DCubeTexture9_SetPrivateData(p,a,b,c,d)   (p)->lpVtbl->SetPrivateData(p,a,b,c,d)
#define IDirect3DCubeTexture9_GetPrivateData(p,a,b,c)     (p)->lpVtbl->GetPrivateData(p,a,b,c)
#define IDirect3DCubeTexture9_FreePrivateData(p,a)        (p)->lpVtbl->FreePrivateData(p,a)
#define IDirect3DCubeTexture9_SetPriority(p,a)            (p)->lpVtbl->SetPriority(p,a)
#define IDirect3DCubeTexture9_GetPriority(p)              (p)->lpVtbl->GetPriority(p)
#define IDirect3DCubeTexture9_PreLoad(p)                  (p)->lpVtbl->PreLoad(p)
#define IDirect3DCubeTexture9_GetType(p)                  (p)->lpVtbl->GetType(p)
/*** IDirect3DCubeTexture9 methods: IDirect3DBaseTexture9 ***/
#define IDirect3DCubeTexture9_SetLOD(p,a)                 (p)->lpVtbl->SetLOD(p,a)
#define IDirect3DCubeTexture9_GetLOD(p)                   (p)->lpVtbl->GetLOD(p)
#define IDirect3DCubeTexture9_GetLevelCount(p)            (p)->lpVtbl->GetLevelCount(p)
#define IDirect3DCubeTexture9_SetAutoGenFilterType(p,a)   (p)->lpVtbl->SetAutoGenFilterType(p,a)
#define IDirect3DCubeTexture9_GetAutoGenFilterType(p)     (p)->lpVtbl->GetAutoGenFilterType(p)
#define IDirect3DCubeTexture9_GenerateMipSubLevels(p)     (p)->lpVtbl->GenerateMipSubLevels(p)
/*** IDirect3DCubeTexture9 methods ***/
#define IDirect3DCubeTexture9_GetLevelDesc(p,a,b)         (p)->lpVtbl->GetLevelDesc(p,a,b)
#define IDirect3DCubeTexture9_GetCubeMapSurface(p,a,b,c)  (p)->lpVtbl->GetCubeMapSurface(p,a,b,c)
#define IDirect3DCubeTexture9_LockRect(p,a,b,c,d,e)       (p)->lpVtbl->LockRect(p,a,b,c,d,e)
#define IDirect3DCubeTexture9_UnlockRect(p,a,b)           (p)->lpVtbl->UnlockRect(p,a,b)
#define IDirect3DCubeTexture9_AddDirtyRect(p,a,b)         (p)->lpVtbl->AddDirtyRect(p,a,b)
#else
/*** IUnknown methods ***/
#define IDirect3DCubeTexture9_QueryInterface(p,a,b)       (p)->QueryInterface(a,b)
#define IDirect3DCubeTexture9_AddRef(p)                   (p)->AddRef()
#define IDirect3DCubeTexture9_Release(p)                  (p)->Release()
/*** IDirect3DCubeTexture9 methods: IDirect3DResource9 ***/
#define IDirect3DCubeTexture9_GetDevice(p,a)              (p)->GetDevice(a)
#define IDirect3DCubeTexture9_SetPrivateData(p,a,b,c,d)   (p)->SetPrivateData(a,b,c,d)
#define IDirect3DCubeTexture9_GetPrivateData(p,a,b,c)     (p)->GetPrivateData(a,b,c)
#define IDirect3DCubeTexture9_FreePrivateData(p,a)        (p)->FreePrivateData(a)
#define IDirect3DCubeTexture9_SetPriority(p,a)            (p)->SetPriority(a)
#define IDirect3DCubeTexture9_GetPriority(p)              (p)->GetPriority()
#define IDirect3DCubeTexture9_PreLoad(p)                  (p)->PreLoad()
#define IDirect3DCubeTexture9_GetType(p)                  (p)->GetType()
/*** IDirect3DCubeTexture9 methods: IDirect3DBaseTexture9 ***/
#define IDirect3DCubeTexture9_SetLOD(p,a)                 (p)->SetLOD(a)
#define IDirect3DCubeTexture9_GetLOD(p)                   (p)->GetLOD()
#define IDirect3DCubeTexture9_GetLevelCount(p)            (p)->GetLevelCount()
#define IDirect3DCubeTexture9_SetAutoGenFilterType(p,a)   (p)->SetAutoGenFilterType(a)
#define IDirect3DCubeTexture9_GetAutoGenFilterType(p)     (p)->GetAutoGenFilterType()
#define IDirect3DCubeTexture9_GenerateMipSubLevels(p)     (p)->GenerateMipSubLevels()
/*** IDirect3DCubeTexture9 methods ***/
#define IDirect3DCubeTexture9_GetLevelDesc(p,a,b)         (p)->GetLevelDesc(a,b)
#define IDirect3DCubeTexture9_GetCubeMapSurface(p,a,b,c)  (p)->GetCubeMapSurface(a,b,c)
#define IDirect3DCubeTexture9_LockRect(p,a,b,c,d,e)       (p)->LockRect(a,b,c,d,e)
#define IDirect3DCubeTexture9_UnlockRect(p,a,b)           (p)->UnlockRect(a,b)
#define IDirect3DCubeTexture9_AddDirtyRect(p,a,b)         (p)->AddDirtyRect(a,b)
#endif

/*****************************************************************************
 * IDirect3DTexture9 interface
 */
#define INTERFACE IDirect3DTexture9
DECLARE_INTERFACE_(IDirect3DTexture9,IDirect3DBaseTexture9)
{
    /*** IUnknown methods ***/
    STDMETHOD_(HRESULT,QueryInterface)(THIS_ REFIID riid, void** ppvObject) PURE;
    STDMETHOD_(ULONG,AddRef)(THIS) PURE;
    STDMETHOD_(ULONG,Release)(THIS) PURE;
    /*** IDirect3DResource9 methods ***/
    STDMETHOD(GetDevice)(THIS_ struct IDirect3DDevice9** ppDevice) PURE;
    STDMETHOD(SetPrivateData)(THIS_ REFGUID refguid, CONST void* pData, DWORD SizeOfData, DWORD Flags) PURE;
    STDMETHOD(GetPrivateData)(THIS_ REFGUID refguid, void* pData, DWORD* pSizeOfData) PURE;
    STDMETHOD(FreePrivateData)(THIS_ REFGUID refguid) PURE;
    STDMETHOD_(DWORD, SetPriority)(THIS_ DWORD PriorityNew) PURE;
    STDMETHOD_(DWORD, GetPriority)(THIS) PURE;
    STDMETHOD_(void, PreLoad)(THIS) PURE;
    STDMETHOD_(D3DRESOURCETYPE, GetType)(THIS) PURE;
    /*** IDirect3DBaseTexture9 methods ***/
    STDMETHOD_(DWORD, SetLOD)(THIS_ DWORD LODNew) PURE;
    STDMETHOD_(DWORD, GetLOD)(THIS) PURE;
    STDMETHOD_(DWORD, GetLevelCount)(THIS) PURE;
    STDMETHOD(SetAutoGenFilterType)(THIS_ D3DTEXTUREFILTERTYPE FilterType) PURE;
    STDMETHOD_(D3DTEXTUREFILTERTYPE, GetAutoGenFilterType)(THIS) PURE;
    STDMETHOD_(void, GenerateMipSubLevels)(THIS) PURE;
    /*** IDirect3DTexture9 methods ***/
    STDMETHOD(GetLevelDesc)(THIS_ UINT Level, D3DSURFACE_DESC* pDesc) PURE;
    STDMETHOD(GetSurfaceLevel)(THIS_ UINT Level, IDirect3DSurface9** ppSurfaceLevel) PURE;
    STDMETHOD(LockRect)(THIS_ UINT Level, D3DLOCKED_RECT* pLockedRect, CONST RECT* pRect, DWORD Flags) PURE;
    STDMETHOD(UnlockRect)(THIS_ UINT Level) PURE;
    STDMETHOD(AddDirtyRect)(THIS_ CONST RECT* pDirtyRect) PURE;
};
#undef INTERFACE

#if !defined(__cplusplus) || defined(CINTERFACE)
/*** IUnknown methods ***/
#define IDirect3DTexture9_QueryInterface(p,a,b)      (p)->lpVtbl->QueryInterface(p,a,b)
#define IDirect3DTexture9_AddRef(p)                  (p)->lpVtbl->AddRef(p)
#define IDirect3DTexture9_Release(p)                 (p)->lpVtbl->Release(p)
/*** IDirect3DTexture9 methods: IDirect3DResource9 ***/
#define IDirect3DTexture9_GetDevice(p,a)             (p)->lpVtbl->GetDevice(p,a)
#define IDirect3DTexture9_SetPrivateData(p,a,b,c,d)  (p)->lpVtbl->SetPrivateData(p,a,b,c,d)
#define IDirect3DTexture9_GetPrivateData(p,a,b,c)    (p)->lpVtbl->GetPrivateData(p,a,b,c)
#define IDirect3DTexture9_FreePrivateData(p,a)       (p)->lpVtbl->FreePrivateData(p,a)
#define IDirect3DTexture9_SetPriority(p,a)           (p)->lpVtbl->SetPriority(p,a)
#define IDirect3DTexture9_GetPriority(p)             (p)->lpVtbl->GetPriority(p)
#define IDirect3DTexture9_PreLoad(p)                 (p)->lpVtbl->PreLoad(p)
#define IDirect3DTexture9_GetType(p)                 (p)->lpVtbl->GetType(p)
/*** IDirect3DTexture9 methods: IDirect3DBaseTexture9 ***/
#define IDirect3DTexture9_SetLOD(p,a)                (p)->lpVtbl->SetLOD(p,a)
#define IDirect3DTexture9_GetLOD(p)                  (p)->lpVtbl->GetLOD(p)
#define IDirect3DTexture9_GetLevelCount(p)           (p)->lpVtbl->GetLevelCount(p)
#define IDirect3DTexture9_SetAutoGenFilterType(p,a)  (p)->lpVtbl->SetAutoGenFilterType(p,a)
#define IDirect3DTexture9_GetAutoGenFilterType(p)    (p)->lpVtbl->GetAutoGenFilterType(p)
#define IDirect3DTexture9_GenerateMipSubLevels(p)    (p)->lpVtbl->GenerateMipSubLevels(p)
/*** IDirect3DTexture9 methods ***/
#define IDirect3DTexture9_GetLevelDesc(p,a,b)        (p)->lpVtbl->GetLevelDesc(p,a,b)
#define IDirect3DTexture9_GetSurfaceLevel(p,a,b)     (p)->lpVtbl->GetSurfaceLevel(p,a,b)
#define IDirect3DTexture9_LockRect(p,a,b,c,d)        (p)->lpVtbl->LockRect(p,a,b,c,d)
#define IDirect3DTexture9_UnlockRect(p,a)            (p)->lpVtbl->UnlockRect(p,a)
#define IDirect3DTexture9_AddDirtyRect(p,a)          (p)->lpVtbl->AddDirtyRect(p,a)
#else
/*** IUnknown methods ***/
#define IDirect3DTexture9_QueryInterface(p,a,b)      (p)->QueryInterface(a,b)
#define IDirect3DTexture9_AddRef(p)                  (p)->AddRef()
#define IDirect3DTexture9_Release(p)                 (p)->Release()
/*** IDirect3DTexture9 methods: IDirect3DResource9 ***/
#define IDirect3DTexture9_GetDevice(p,a)             (p)->GetDevice(a)
#define IDirect3DTexture9_SetPrivateData(p,a,b,c,d)  (p)->SetPrivateData(a,b,c,d)
#define IDirect3DTexture9_GetPrivateData(p,a,b,c)    (p)->GetPrivateData(a,b,c)
#define IDirect3DTexture9_FreePrivateData(p,a)       (p)->FreePrivateData(a)
#define IDirect3DTexture9_SetPriority(p,a)           (p)->SetPriority(a)
#define IDirect3DTexture9_GetPriority(p)             (p)->GetPriority()
#define IDirect3DTexture9_PreLoad(p)                 (p)->PreLoad()
#define IDirect3DTexture9_GetType(p)                 (p)->GetType()
/*** IDirect3DTexture9 methods: IDirect3DBaseTexture9 ***/
#define IDirect3DTexture9_SetLOD(p,a)                (p)->SetLOD(a)
#define IDirect3DTexture9_GetLOD(p)                  (p)->GetLOD()
#define IDirect3DTexture9_GetLevelCount(p)           (p)->GetLevelCount()
#define IDirect3DTexture9_SetAutoGenFilterType(p,a)  (p)->SetAutoGenFilterType(a)
#define IDirect3DTexture9_GetAutoGenFilterType(p)    (p)->GetAutoGenFilterType()
#define IDirect3DTexture9_GenerateMipSubLevels(p)    (p)->GenerateMipSubLevels()
/*** IDirect3DTexture9 methods ***/
#define IDirect3DTexture9_GetLevelDesc(p,a,b)        (p)->GetLevelDesc(a,b)
#define IDirect3DTexture9_GetSurfaceLevel(p,a,b)     (p)->GetSurfaceLevel(a,b)
#define IDirect3DTexture9_LockRect(p,a,b,c,d)        (p)->LockRect(a,b,c,d)
#define IDirect3DTexture9_UnlockRect(p,a)            (p)->UnlockRect(a)
#define IDirect3DTexture9_AddDirtyRect(p,a)          (p)->AddDirtyRect(a)
#endif

/*****************************************************************************
 * IDirect3DVolumeTexture9 interface
 */
#define INTERFACE IDirect3DVolumeTexture9
DECLARE_INTERFACE_(IDirect3DVolumeTexture9,IDirect3DBaseTexture9)
{
    /*** IUnknown methods ***/
    STDMETHOD_(HRESULT,QueryInterface)(THIS_ REFIID riid, void** ppvObject) PURE;
    STDMETHOD_(ULONG,AddRef)(THIS) PURE;
    STDMETHOD_(ULONG,Release)(THIS) PURE;
    /*** IDirect3DResource9 methods ***/
    STDMETHOD(GetDevice)(THIS_ struct IDirect3DDevice9** ppDevice) PURE;
    STDMETHOD(SetPrivateData)(THIS_ REFGUID refguid, CONST void* pData, DWORD SizeOfData, DWORD Flags) PURE;
    STDMETHOD(GetPrivateData)(THIS_ REFGUID refguid, void* pData, DWORD* pSizeOfData) PURE;
    STDMETHOD(FreePrivateData)(THIS_ REFGUID refguid) PURE;
    STDMETHOD_(DWORD, SetPriority)(THIS_ DWORD PriorityNew) PURE;
    STDMETHOD_(DWORD, GetPriority)(THIS) PURE;
    STDMETHOD_(void, PreLoad)(THIS) PURE;
    STDMETHOD_(D3DRESOURCETYPE, GetType)(THIS) PURE;
    /*** IDirect3DBaseTexture9 methods ***/
    STDMETHOD_(DWORD, SetLOD)(THIS_ DWORD LODNew) PURE;
    STDMETHOD_(DWORD, GetLOD)(THIS) PURE;
    STDMETHOD_(DWORD, GetLevelCount)(THIS) PURE;
    STDMETHOD(SetAutoGenFilterType)(THIS_ D3DTEXTUREFILTERTYPE FilterType) PURE;
    STDMETHOD_(D3DTEXTUREFILTERTYPE, GetAutoGenFilterType)(THIS) PURE;
    STDMETHOD_(void, GenerateMipSubLevels)(THIS) PURE;
    /*** IDirect3DVolumeTexture9 methods ***/
    STDMETHOD(GetLevelDesc)(THIS_ UINT Level, D3DVOLUME_DESC *pDesc) PURE;
    STDMETHOD(GetVolumeLevel)(THIS_ UINT Level, IDirect3DVolume9** ppVolumeLevel) PURE;
    STDMETHOD(LockBox)(THIS_ UINT Level, D3DLOCKED_BOX* pLockedVolume, CONST D3DBOX* pBox, DWORD Flags) PURE;
    STDMETHOD(UnlockBox)(THIS_ UINT Level) PURE;
    STDMETHOD(AddDirtyBox)(THIS_ CONST D3DBOX* pDirtyBox) PURE;
};
#undef INTERFACE

#if !defined(__cplusplus) || defined(CINTERFACE)
/*** IUnknown methods ***/
#define IDirect3DVolumeTexture9_QueryInterface(p,a,b) (p)->lpVtbl->QueryInterface(p,a,b)
#define IDirect3DVolumeTexture9_AddRef(p) (p)->lpVtbl->AddRef(p)
#define IDirect3DVolumeTexture9_Release(p) (p)->lpVtbl->Release(p)
/*** IDirect3DVolumeTexture9 methods: IDirect3DResource9 ***/
#define IDirect3DVolumeTexture9_GetDevice(p,a) (p)->lpVtbl->GetDevice(p,a)
#define IDirect3DVolumeTexture9_SetPrivateData(p,a,b,c,d) (p)->lpVtbl->SetPrivateData(p,a,b,c,d)
#define IDirect3DVolumeTexture9_GetPrivateData(p,a,b,c) (p)->lpVtbl->GetPrivateData(p,a,b,c)
#define IDirect3DVolumeTexture9_FreePrivateData(p,a) (p)->lpVtbl->FreePrivateData(p,a)
#define IDirect3DVolumeTexture9_SetPriority(p,a) (p)->lpVtbl->SetPriority(p,a)
#define IDirect3DVolumeTexture9_GetPriority(p) (p)->lpVtbl->GetPriority(p)
#define IDirect3DVolumeTexture9_PreLoad(p) (p)->lpVtbl->PreLoad(p)
#define IDirect3DVolumeTexture9_GetType(p) (p)->lpVtbl->GetType(p)
/*** IDirect3DVolumeTexture9 methods: IDirect3DBaseTexture9 ***/
#define IDirect3DVolumeTexture9_SetLOD(p,a) (p)->lpVtbl->SetLOD(p,a)
#define IDirect3DVolumeTexture9_GetLOD(p) (p)->lpVtbl->GetLOD(p)
#define IDirect3DVolumeTexture9_GetLevelCount(p) (p)->lpVtbl->GetLevelCount(p)
#define IDirect3DVolumeTexture9_SetAutoGenFilterType(p,a) (p)->lpVtbl->SetAutoGenFilterType(p,a)
#define IDirect3DVolumeTexture9_GetAutoGenFilterType(p) (p)->lpVtbl->GetAutoGenFilterType(p)
#define IDirect3DVolumeTexture9_GenerateMipSubLevels(p) (p)->lpVtbl->GenerateMipSubLevels(p)
/*** IDirect3DVolumeTexture9 methods ***/
#define IDirect3DVolumeTexture9_GetLevelDesc(p,a,b) (p)->lpVtbl->GetLevelDesc(p,a,b)
#define IDirect3DVolumeTexture9_GetVolumeLevel(p,a,b) (p)->lpVtbl->GetVolumeLevel(p,a,b)
#define IDirect3DVolumeTexture9_LockBox(p,a,b,c,d) (p)->lpVtbl->LockBox(p,a,b,c,d)
#define IDirect3DVolumeTexture9_UnlockBox(p,a) (p)->lpVtbl->UnlockBox(p,a)
#define IDirect3DVolumeTexture9_AddDirtyBox(p,a) (p)->lpVtbl->AddDirtyBox(p,a)
#else
/*** IUnknown methods ***/
#define IDirect3DVolumeTexture9_QueryInterface(p,a,b) (p)->QueryInterface(a,b)
#define IDirect3DVolumeTexture9_AddRef(p) (p)->AddRef()
#define IDirect3DVolumeTexture9_Release(p) (p)->Release()
/*** IDirect3DVolumeTexture9 methods: IDirect3DResource9 ***/
#define IDirect3DVolumeTexture9_GetDevice(p,a) (p)->GetDevice(a)
#define IDirect3DVolumeTexture9_SetPrivateData(p,a,b,c,d) (p)->SetPrivateData(a,b,c,d)
#define IDirect3DVolumeTexture9_GetPrivateData(p,a,b,c) (p)->GetPrivateData(a,b,c)
#define IDirect3DVolumeTexture9_FreePrivateData(p,a) (p)->FreePrivateData(a)
#define IDirect3DVolumeTexture9_SetPriority(p,a) (p)->SetPriority(a)
#define IDirect3DVolumeTexture9_GetPriority(p) (p)->GetPriority()
#define IDirect3DVolumeTexture9_PreLoad(p) (p)->PreLoad()
#define IDirect3DVolumeTexture9_GetType(p) (p)->GetType()
/*** IDirect3DVolumeTexture9 methods: IDirect3DBaseTexture9 ***/
#define IDirect3DVolumeTexture9_SetLOD(p,a) (p)->SetLOD(a)
#define IDirect3DVolumeTexture9_GetLOD(p) (p)->GetLOD()
#define IDirect3DVolumeTexture9_GetLevelCount(p) (p)->GetLevelCount()
#define IDirect3DVolumeTexture9_SetAutoGenFilterType(p,a) (p)->SetAutoGenFilterType(a)
#define IDirect3DVolumeTexture9_GetAutoGenFilterType(p) (p)->GetAutoGenFilterType()
#define IDirect3DVolumeTexture9_GenerateMipSubLevels(p) (p)->GenerateMipSubLevels()
/*** IDirect3DVolumeTexture9 methods ***/
#define IDirect3DVolumeTexture9_GetLevelDesc(p,a,b) (p)->GetLevelDesc(a,b)
#define IDirect3DVolumeTexture9_GetVolumeLevel(p,a,b) (p)->GetVolumeLevel(a,b)
#define IDirect3DVolumeTexture9_LockBox(p,a,b,c,d) (p)->LockBox(a,b,c,d)
#define IDirect3DVolumeTexture9_UnlockBox(p,a) (p)->UnlockBox(a)
#define IDirect3DVolumeTexture9_AddDirtyBox(p,a) (p)->AddDirtyBox(a)
#endif

/*****************************************************************************
 * IDirect3DVertexDeclaration9 interface
 */
#define INTERFACE IDirect3DVertexDeclaration9
DECLARE_INTERFACE_(IDirect3DVertexDeclaration9,IUnknown)
{
    /*** IUnknown methods ***/
    STDMETHOD_(HRESULT,QueryInterface)(THIS_ REFIID riid, void** ppvObject) PURE;
    STDMETHOD_(ULONG,AddRef)(THIS) PURE;
    STDMETHOD_(ULONG,Release)(THIS) PURE;
    /*** IDirect3DVertexDeclaration9 methods ***/
    STDMETHOD(GetDevice)(THIS_ struct IDirect3DDevice9** ppDevice) PURE;
    STDMETHOD(GetDeclaration)(THIS_ D3DVERTEXELEMENT9*, UINT* pNumElements) PURE;
};
#undef INTERFACE

#if !defined(__cplusplus) || defined(CINTERFACE)
/*** IUnknown methods ***/
#define IDirect3DVertexDeclaration9_QueryInterface(p,a,b)  (p)->lpVtbl->QueryInterface(p,a,b)
#define IDirect3DVertexDeclaration9_AddRef(p)              (p)->lpVtbl->AddRef(p)
#define IDirect3DVertexDeclaration9_Release(p)             (p)->lpVtbl->Release(p)
/*** IDirect3DVertexShader9 methods ***/
#define IDirect3DVertexDeclaration9_GetDevice(p,a)         (p)->lpVtbl->GetDevice(p,a)
#define IDirect3DVertexDeclaration9_GetDeclaration(p,a,b)  (p)->lpVtbl->GetDeclaration(p,a,b)
#else
/*** IUnknown methods ***/
#define IDirect3DVertexDeclaration9_QueryInterface(p,a,b)  (p)->QueryInterface(a,b)
#define IDirect3DVertexDeclaration9_AddRef(p)              (p)->AddRef()
#define IDirect3DVertexDeclaration9_Release(p)             (p)->Release()
/*** IDirect3DVertexShader9 methods ***/
#define IDirect3DVertexDeclaration9_GetDevice(p,a)         (p)->GetDevice(a)
#define IDirect3DVertexDeclaration9_GetDeclaration(p,a,b)  (p)->GetDeclaration(a,b)
#endif

/*****************************************************************************
 * IDirect3DVertexShader9 interface
 */
#define INTERFACE IDirect3DVertexShader9
DECLARE_INTERFACE_(IDirect3DVertexShader9,IUnknown)
{
    /*** IUnknown methods ***/
    STDMETHOD_(HRESULT,QueryInterface)(THIS_ REFIID riid, void** ppvObject) PURE;
    STDMETHOD_(ULONG,AddRef)(THIS) PURE;
    STDMETHOD_(ULONG,Release)(THIS) PURE;
    /*** IDirect3DVertexShader9 methods ***/
    STDMETHOD(GetDevice)(THIS_ struct IDirect3DDevice9** ppDevice) PURE;
    STDMETHOD(GetFunction)(THIS_ void*, UINT* pSizeOfData) PURE;
};
#undef INTERFACE

#if !defined(__cplusplus) || defined(CINTERFACE)
/*** IUnknown methods ***/
#define IDirect3DVertexShader9_QueryInterface(p,a,b)  (p)->lpVtbl->QueryInterface(p,a,b)
#define IDirect3DVertexShader9_AddRef(p)              (p)->lpVtbl->AddRef(p)
#define IDirect3DVertexShader9_Release(p)             (p)->lpVtbl->Release(p)
/*** IDirect3DVertexShader9 methods ***/
#define IDirect3DVertexShader9_GetDevice(p,a)         (p)->lpVtbl->GetDevice(p,a)
#define IDirect3DVertexShader9_GetFunction(p,a,b)     (p)->lpVtbl->GetFunction(p,a,b)
#else
/*** IUnknown methods ***/
#define IDirect3DVertexShader9_QueryInterface(p,a,b)  (p)->QueryInterface(a,b)
#define IDirect3DVertexShader9_AddRef(p)              (p)->AddRef()
#define IDirect3DVertexShader9_Release(p)             (p)->Release()
/*** IDirect3DVertexShader9 methods ***/
#define IDirect3DVertexShader9_GetDevice(p,a)         (p)->GetDevice(a)
#define IDirect3DVertexShader9_GetFunction(p,a,b)     (p)->GetFunction(a,b)
#endif

/*****************************************************************************
 * IDirect3DPixelShader9 interface
 */
#define INTERFACE IDirect3DPixelShader9
DECLARE_INTERFACE_(IDirect3DPixelShader9,IUnknown)
{
    /*** IUnknown methods ***/
    STDMETHOD_(HRESULT,QueryInterface)(THIS_ REFIID riid, void** ppvObject) PURE;
    STDMETHOD_(ULONG,AddRef)(THIS) PURE;
    STDMETHOD_(ULONG,Release)(THIS) PURE;
    /*** IDirect3DPixelShader9 methods ***/
    STDMETHOD(GetDevice)(THIS_ struct IDirect3DDevice9** ppDevice) PURE;
    STDMETHOD(GetFunction)(THIS_ void*, UINT* pSizeOfData) PURE;
};
#undef INTERFACE

#if !defined(__cplusplus) || defined(CINTERFACE)
/*** IUnknown methods ***/
#define IDirect3DPixelShader9_QueryInterface(p,a,b)  (p)->lpVtbl->QueryInterface(p,a,b)
#define IDirect3DPixelShader9_AddRef(p)              (p)->lpVtbl->AddRef(p)
#define IDirect3DPixelShader9_Release(p)             (p)->lpVtbl->Release(p)
/*** IDirect3DPixelShader9 methods ***/
#define IDirect3DPixelShader9_GetDevice(p,a)         (p)->lpVtbl->GetDevice(p,a)
#define IDirect3DPixelShader9_GetFunction(p,a,b)     (p)->lpVtbl->GetFunction(p,a,b)
#else
/*** IUnknown methods ***/
#define IDirect3DPixelShader9_QueryInterface(p,a,b)  (p)->QueryInterface(a,b)
#define IDirect3DPixelShader9_AddRef(p)              (p)->AddRef()
#define IDirect3DPixelShader9_Release(p)             (p)->Release()
/*** IDirect3DPixelShader9 methods ***/
#define IDirect3DPixelShader9_GetDevice(p,a)         (p)->GetDevice(a)
#define IDirect3DPixelShader9_GetFunction(p,a,b)     (p)->GetFunction(a,b)
#endif

/*****************************************************************************
 * IDirect3DStateBlock9 interface
 */
#define INTERFACE IDirect3DStateBlock9
DECLARE_INTERFACE_(IDirect3DStateBlock9,IUnknown)
{
    /*** IUnknown methods ***/
    STDMETHOD_(HRESULT,QueryInterface)(THIS_ REFIID riid, void** ppvObject) PURE;
    STDMETHOD_(ULONG,AddRef)(THIS) PURE;
    STDMETHOD_(ULONG,Release)(THIS) PURE;
    /*** IDirect3DStateBlock9 methods ***/
    STDMETHOD(GetDevice)(THIS_ struct IDirect3DDevice9** ppDevice) PURE;
    STDMETHOD(Capture)(THIS) PURE;
    STDMETHOD(Apply)(THIS) PURE;
};
#undef INTERFACE

#if !defined(__cplusplus) || defined(CINTERFACE)
/*** IUnknown methods ***/
#define IDirect3DStateBlock9_QueryInterface(p,a,b)  (p)->lpVtbl->QueryInterface(p,a,b)
#define IDirect3DStateBlock9_AddRef(p)              (p)->lpVtbl->AddRef(p)
#define IDirect3DStateBlock9_Release(p)             (p)->lpVtbl->Release(p)
/*** IDirect3DStateBlock9 methods ***/
#define IDirect3DStateBlock9_GetDevice(p,a)         (p)->lpVtbl->GetDevice(p,a)
#define IDirect3DStateBlock9_Capture(p)             (p)->lpVtbl->Capture(p)
#define IDirect3DStateBlock9_Apply(p)               (p)->lpVtbl->Apply(p)
#else
/*** IUnknown methods ***/
#define IDirect3DStateBlock9_QueryInterface(p,a,b)  (p)->QueryInterface(a,b)
#define IDirect3DStateBlock9_AddRef(p)              (p)->AddRef()
#define IDirect3DStateBlock9_Release(p)             (p)->Release()
/*** IDirect3DStateBlock9 methods ***/
#define IDirect3DStateBlock9_GetDevice(p,a)         (p)->GetDevice(a)
#define IDirect3DStateBlock9_Capture(p)             (p)->Capture()
#define IDirect3DStateBlock9_Apply(p)               (p)->Apply()
#endif

/*****************************************************************************
 * IDirect3DQuery9 interface
 */
#define INTERFACE IDirect3DQuery9
DECLARE_INTERFACE_(IDirect3DQuery9,IUnknown)
{
    /*** IUnknown methods ***/
    STDMETHOD_(HRESULT,QueryInterface)(THIS_ REFIID riid, void** ppvObject) PURE;
    STDMETHOD_(ULONG,AddRef)(THIS) PURE;
    STDMETHOD_(ULONG,Release)(THIS) PURE;
    /*** IDirect3DQuery9 methods ***/
    STDMETHOD(GetDevice)(THIS_ struct IDirect3DDevice9** ppDevice) PURE;
    STDMETHOD_(D3DQUERYTYPE, GetType)(THIS) PURE;
    STDMETHOD_(DWORD, GetDataSize)(THIS) PURE;
    STDMETHOD(Issue)(THIS_ DWORD dwIssueFlags) PURE;
    STDMETHOD(GetData)(THIS_ void* pData, DWORD dwSize, DWORD dwGetDataFlags) PURE;
};
#undef INTERFACE

#if !defined(__cplusplus) || defined(CINTERFACE)
/*** IUnknown methods ***/
#define IDirect3DQuery9_QueryInterface(p,a,b) (p)->lpVtbl->QueryInterface(p,a,b)
#define IDirect3DQuery9_AddRef(p) (p)->lpVtbl->AddRef(p)
#define IDirect3DQuery9_Release(p) (p)->lpVtbl->Release(p)
/*** IDirect3DQuery9 ***/
#define IDirect3DQuery9_GetDevice(p,a) (p)->lpVtbl->GetDevice(p,a)
#define IDirect3DQuery9_GetType(p) (p)->lpVtbl->GetType(p)
#define IDirect3DQuery9_GetDataSize(p) (p)->lpVtbl->GetDataSize(p)
#define IDirect3DQuery9_Issue(p,a) (p)->lpVtbl->Issue(p,a)
#define IDirect3DQuery9_GetData(p,a,b,c) (p)->lpVtbl->GetData(p,a,b,c)
#else
/*** IUnknown methods ***/
#define IDirect3DQuery9_QueryInterface(p,a,b) (p)->QueryInterface(a,b)
#define IDirect3DQuery9_AddRef(p) (p)->AddRef()
#define IDirect3DQuery9_Release(p) (p)->Release()
/*** IDirect3DQuery9 ***/
#define IDirect3DQuery9_GetDevice(p,a) (p)->GetDevice(a)
#define IDirect3DQuery9_GetType(p) (p)->GetType()
#define IDirect3DQuery9_GetDataSize(p) (p)->GetDataSize()
#define IDirect3DQuery9_Issue(p,a) (p)->Issue(a)
#define IDirect3DQuery9_GetData(p,a,b,c) (p)->GetData(a,b,c)
#endif

/*****************************************************************************
 * IDirect3DDevice9 interface
 */
#define INTERFACE IDirect3DDevice9
DECLARE_INTERFACE_(IDirect3DDevice9,IUnknown)
{
    /*** IUnknown methods ***/
    STDMETHOD_(HRESULT,QueryInterface)(THIS_ REFIID riid, void** ppvObject) PURE;
    STDMETHOD_(ULONG,AddRef)(THIS) PURE;
    STDMETHOD_(ULONG,Release)(THIS) PURE;
    /*** IDirect3DDevice9 methods ***/
    STDMETHOD(TestCooperativeLevel)(THIS) PURE;
    STDMETHOD_(UINT, GetAvailableTextureMem)(THIS) PURE;
    STDMETHOD(EvictManagedResources)(THIS) PURE;
    STDMETHOD(GetDirect3D)(THIS_ IDirect3D9** ppD3D9) PURE;
    STDMETHOD(GetDeviceCaps)(THIS_ D3DCAPS9* pCaps) PURE;
    STDMETHOD(GetDisplayMode)(THIS_ UINT iSwapChain, D3DDISPLAYMODE* pMode) PURE;
    STDMETHOD(GetCreationParameters)(THIS_ D3DDEVICE_CREATION_PARAMETERS *pParameters) PURE;
    STDMETHOD(SetCursorProperties)(THIS_ UINT XHotSpot, UINT YHotSpot, IDirect3DSurface9* pCursorBitmap) PURE;
    STDMETHOD_(void, SetCursorPosition)(THIS_ int X,int Y, DWORD Flags) PURE;
    STDMETHOD_(WINBOOL, ShowCursor)(THIS_ WINBOOL bShow) PURE;
    STDMETHOD(CreateAdditionalSwapChain)(THIS_ D3DPRESENT_PARAMETERS* pPresentationParameters, IDirect3DSwapChain9** pSwapChain) PURE;
    STDMETHOD(GetSwapChain)(THIS_ UINT iSwapChain, IDirect3DSwapChain9** pSwapChain) PURE;
    STDMETHOD_(UINT, GetNumberOfSwapChains)(THIS) PURE;
    STDMETHOD(Reset)(THIS_ D3DPRESENT_PARAMETERS* pPresentationParameters) PURE;
    STDMETHOD(Present)(THIS_ CONST RECT* pSourceRect, CONST RECT* pDestRect, HWND hDestWindowOverride, CONST RGNDATA* pDirtyRegion) PURE;
    STDMETHOD(GetBackBuffer)(THIS_ UINT iSwapChain, UINT iBackBuffer, D3DBACKBUFFER_TYPE Type, IDirect3DSurface9** ppBackBuffer) PURE;
    STDMETHOD(GetRasterStatus)(THIS_ UINT iSwapChain, D3DRASTER_STATUS* pRasterStatus) PURE;
    STDMETHOD(SetDialogBoxMode)(THIS_ WINBOOL bEnableDialogs) PURE;
    STDMETHOD_(void, SetGammaRamp)(THIS_ UINT iSwapChain, DWORD Flags, CONST D3DGAMMARAMP* pRamp) PURE;
    STDMETHOD_(void, GetGammaRamp)(THIS_ UINT iSwapChain, D3DGAMMARAMP* pRamp) PURE;
    STDMETHOD(CreateTexture)(THIS_ UINT Width, UINT Height, UINT Levels, DWORD Usage, D3DFORMAT Format, D3DPOOL Pool, IDirect3DTexture9** ppTexture, HANDLE* pSharedHandle) PURE;
    STDMETHOD(CreateVolumeTexture)(THIS_ UINT Width, UINT Height, UINT Depth, UINT Levels, DWORD Usage, D3DFORMAT Format, D3DPOOL Pool, IDirect3DVolumeTexture9** ppVolumeTexture, HANDLE* pSharedHandle) PURE;
    STDMETHOD(CreateCubeTexture)(THIS_ UINT EdgeLength, UINT Levels, DWORD Usage, D3DFORMAT Format, D3DPOOL Pool, IDirect3DCubeTexture9** ppCubeTexture, HANDLE* pSharedHandle) PURE;
    STDMETHOD(CreateVertexBuffer)(THIS_ UINT Length, DWORD Usage, DWORD FVF, D3DPOOL Pool, IDirect3DVertexBuffer9** ppVertexBuffer, HANDLE* pSharedHandle) PURE;
    STDMETHOD(CreateIndexBuffer)(THIS_ UINT Length, DWORD Usage, D3DFORMAT Format, D3DPOOL Pool, IDirect3DIndexBuffer9** ppIndexBuffer, HANDLE* pSharedHandle) PURE;
    STDMETHOD(CreateRenderTarget)(THIS_ UINT Width, UINT Height, D3DFORMAT Format, D3DMULTISAMPLE_TYPE MultiSample, DWORD MultisampleQuality, WINBOOL Lockable, IDirect3DSurface9** ppSurface, HANDLE* pSharedHandle) PURE;
    STDMETHOD(CreateDepthStencilSurface)(THIS_ UINT Width, UINT Height, D3DFORMAT Format, D3DMULTISAMPLE_TYPE MultiSample, DWORD MultisampleQuality, WINBOOL Discard, IDirect3DSurface9** ppSurface, HANDLE* pSharedHandle) PURE;
    STDMETHOD(UpdateSurface)(THIS_ IDirect3DSurface9* pSourceSurface, CONST RECT* pSourceRect, IDirect3DSurface9* pDestinationSurface, CONST POINT* pDestPoint) PURE;
    STDMETHOD(UpdateTexture)(THIS_ IDirect3DBaseTexture9* pSourceTexture, IDirect3DBaseTexture9* pDestinationTexture) PURE;
    STDMETHOD(GetRenderTargetData)(THIS_ IDirect3DSurface9* pRenderTarget, IDirect3DSurface9* pDestSurface) PURE;
    STDMETHOD(GetFrontBufferData)(THIS_ UINT iSwapChain, IDirect3DSurface9* pDestSurface) PURE;
    STDMETHOD(StretchRect)(THIS_ IDirect3DSurface9* pSourceSurface, CONST RECT* pSourceRect, IDirect3DSurface9* pDestSurface, CONST RECT* pDestRect, D3DTEXTUREFILTERTYPE Filter) PURE;
    STDMETHOD(ColorFill)(THIS_ IDirect3DSurface9* pSurface, CONST RECT* pRect, D3DCOLOR color) PURE;
    STDMETHOD(CreateOffscreenPlainSurface)(THIS_ UINT Width, UINT Height, D3DFORMAT Format, D3DPOOL Pool, IDirect3DSurface9** ppSurface, HANDLE* pSharedHandle) PURE;
    STDMETHOD(SetRenderTarget)(THIS_ DWORD RenderTargetIndex, IDirect3DSurface9* pRenderTarget) PURE;
    STDMETHOD(GetRenderTarget)(THIS_ DWORD RenderTargetIndex, IDirect3DSurface9** ppRenderTarget) PURE;
    STDMETHOD(SetDepthStencilSurface)(THIS_ IDirect3DSurface9* pNewZStencil) PURE;
    STDMETHOD(GetDepthStencilSurface)(THIS_ IDirect3DSurface9** ppZStencilSurface) PURE;
    STDMETHOD(BeginScene)(THIS) PURE;
    STDMETHOD(EndScene)(THIS) PURE;
    STDMETHOD(Clear)(THIS_ DWORD Count, CONST D3DRECT* pRects, DWORD Flags, D3DCOLOR Color, float Z, DWORD Stencil) PURE;
    STDMETHOD(SetTransform)(THIS_ D3DTRANSFORMSTATETYPE State, CONST D3DMATRIX* pMatrix) PURE;
    STDMETHOD(GetTransform)(THIS_ D3DTRANSFORMSTATETYPE State, D3DMATRIX* pMatrix) PURE;
    STDMETHOD(MultiplyTransform)(THIS_ D3DTRANSFORMSTATETYPE, CONST D3DMATRIX*) PURE;
    STDMETHOD(SetViewport)(THIS_ CONST D3DVIEWPORT9* pViewport) PURE;
    STDMETHOD(GetViewport)(THIS_ D3DVIEWPORT9* pViewport) PURE;
    STDMETHOD(SetMaterial)(THIS_ CONST D3DMATERIAL9* pMaterial) PURE;
    STDMETHOD(GetMaterial)(THIS_ D3DMATERIAL9* pMaterial) PURE;
    STDMETHOD(SetLight)(THIS_ DWORD Index, CONST D3DLIGHT9*) PURE;
    STDMETHOD(GetLight)(THIS_ DWORD Index, D3DLIGHT9*) PURE;
    STDMETHOD(LightEnable)(THIS_ DWORD Index, WINBOOL Enable) PURE;
    STDMETHOD(GetLightEnable)(THIS_ DWORD Index, WINBOOL* pEnable) PURE;
    STDMETHOD(SetClipPlane)(THIS_ DWORD Index, CONST float* pPlane) PURE;
    STDMETHOD(GetClipPlane)(THIS_ DWORD Index, float* pPlane) PURE;
    STDMETHOD(SetRenderState)(THIS_ D3DRENDERSTATETYPE State, DWORD Value) PURE;
    STDMETHOD(GetRenderState)(THIS_ D3DRENDERSTATETYPE State, DWORD* pValue) PURE;
    STDMETHOD(CreateStateBlock)(THIS_ D3DSTATEBLOCKTYPE Type, IDirect3DStateBlock9** ppSB) PURE;
    STDMETHOD(BeginStateBlock)(THIS) PURE;
    STDMETHOD(EndStateBlock)(THIS_ IDirect3DStateBlock9** ppSB) PURE;
    STDMETHOD(SetClipStatus)(THIS_ CONST D3DCLIPSTATUS9* pClipStatus) PURE;
    STDMETHOD(GetClipStatus)(THIS_ D3DCLIPSTATUS9* pClipStatus) PURE;
    STDMETHOD(GetTexture)(THIS_ DWORD Stage, IDirect3DBaseTexture9** ppTexture) PURE;
    STDMETHOD(SetTexture)(THIS_ DWORD Stage, IDirect3DBaseTexture9* pTexture) PURE;
    STDMETHOD(GetTextureStageState)(THIS_ DWORD Stage, D3DTEXTURESTAGESTATETYPE Type, DWORD* pValue) PURE;
    STDMETHOD(SetTextureStageState)(THIS_ DWORD Stage, D3DTEXTURESTAGESTATETYPE Type, DWORD Value) PURE;
    STDMETHOD(GetSamplerState)(THIS_ DWORD Sampler, D3DSAMPLERSTATETYPE Type, DWORD* pValue) PURE;
    STDMETHOD(SetSamplerState)(THIS_ DWORD Sampler, D3DSAMPLERSTATETYPE Type, DWORD Value) PURE;
    STDMETHOD(ValidateDevice)(THIS_ DWORD* pNumPasses) PURE;
    STDMETHOD(SetPaletteEntries)(THIS_ UINT PaletteNumber, CONST PALETTEENTRY* pEntries) PURE;
    STDMETHOD(GetPaletteEntries)(THIS_ UINT PaletteNumber,PALETTEENTRY* pEntries) PURE;
    STDMETHOD(SetCurrentTexturePalette)(THIS_ UINT PaletteNumber) PURE;
    STDMETHOD(GetCurrentTexturePalette)(THIS_ UINT *PaletteNumber) PURE;
    STDMETHOD(SetScissorRect)(THIS_ CONST RECT* pRect) PURE;
    STDMETHOD(GetScissorRect)(THIS_ RECT* pRect) PURE;
    STDMETHOD(SetSoftwareVertexProcessing)(THIS_ WINBOOL bSoftware) PURE;
    STDMETHOD_(WINBOOL, GetSoftwareVertexProcessing)(THIS) PURE;
    STDMETHOD(SetNPatchMode)(THIS_ float nSegments) PURE;
    STDMETHOD_(float, GetNPatchMode)(THIS) PURE;
    STDMETHOD(DrawPrimitive)(THIS_ D3DPRIMITIVETYPE PrimitiveType, UINT StartVertex, UINT PrimitiveCount) PURE;
    STDMETHOD(DrawIndexedPrimitive)(THIS_ D3DPRIMITIVETYPE, INT BaseVertexIndex, UINT MinVertexIndex, UINT NumVertices, UINT startIndex, UINT primCount) PURE;
    STDMETHOD(DrawPrimitiveUP)(THIS_ D3DPRIMITIVETYPE PrimitiveType, UINT PrimitiveCount, CONST void* pVertexStreamZeroData, UINT VertexStreamZeroStride) PURE;
    STDMETHOD(DrawIndexedPrimitiveUP)(THIS_ D3DPRIMITIVETYPE PrimitiveType, UINT MinVertexIndex, UINT NumVertices, UINT PrimitiveCount, CONST void* pIndexData, D3DFORMAT IndexDataFormat, CONST void* pVertexStreamZeroData, UINT VertexStreamZeroStride) PURE;
    STDMETHOD(ProcessVertices)(THIS_ UINT SrcStartIndex, UINT DestIndex, UINT VertexCount, IDirect3DVertexBuffer9* pDestBuffer, IDirect3DVertexDeclaration9* pVertexDecl, DWORD Flags) PURE;
    STDMETHOD(CreateVertexDeclaration)(THIS_ CONST D3DVERTEXELEMENT9* pVertexElements, IDirect3DVertexDeclaration9** ppDecl) PURE;
    STDMETHOD(SetVertexDeclaration)(THIS_ IDirect3DVertexDeclaration9* pDecl) PURE;
    STDMETHOD(GetVertexDeclaration)(THIS_ IDirect3DVertexDeclaration9** ppDecl) PURE;
    STDMETHOD(SetFVF)(THIS_ DWORD FVF) PURE;
    STDMETHOD(GetFVF)(THIS_ DWORD* pFVF) PURE;
    STDMETHOD(CreateVertexShader)(THIS_ CONST DWORD* pFunction, IDirect3DVertexShader9** ppShader) PURE;
    STDMETHOD(SetVertexShader)(THIS_ IDirect3DVertexShader9* pShader) PURE;
    STDMETHOD(GetVertexShader)(THIS_ IDirect3DVertexShader9** ppShader) PURE;
    STDMETHOD(SetVertexShaderConstantF)(THIS_ UINT StartRegister, CONST float* pConstantData, UINT Vector4fCount) PURE;
    STDMETHOD(GetVertexShaderConstantF)(THIS_ UINT StartRegister, float* pConstantData, UINT Vector4fCount) PURE;
    STDMETHOD(SetVertexShaderConstantI)(THIS_ UINT StartRegister, CONST int* pConstantData, UINT Vector4iCount) PURE;
    STDMETHOD(GetVertexShaderConstantI)(THIS_ UINT StartRegister, int* pConstantData, UINT Vector4iCount) PURE;
    STDMETHOD(SetVertexShaderConstantB)(THIS_ UINT StartRegister, CONST WINBOOL* pConstantData, UINT  BoolCount) PURE;
    STDMETHOD(GetVertexShaderConstantB)(THIS_ UINT StartRegister, WINBOOL* pConstantData, UINT BoolCount) PURE;
    STDMETHOD(SetStreamSource)(THIS_ UINT StreamNumber, IDirect3DVertexBuffer9* pStreamData, UINT OffsetInBytes, UINT Stride) PURE;
    STDMETHOD(GetStreamSource)(THIS_ UINT StreamNumber, IDirect3DVertexBuffer9** ppStreamData, UINT* OffsetInBytes, UINT* pStride) PURE;
    STDMETHOD(SetStreamSourceFreq)(THIS_ UINT StreamNumber, UINT Divider) PURE;
    STDMETHOD(GetStreamSourceFreq)(THIS_ UINT StreamNumber, UINT* Divider) PURE;
    STDMETHOD(SetIndices)(THIS_ IDirect3DIndexBuffer9* pIndexData) PURE;
    STDMETHOD(GetIndices)(THIS_ IDirect3DIndexBuffer9** ppIndexData) PURE;
    STDMETHOD(CreatePixelShader)(THIS_ CONST DWORD* pFunction, IDirect3DPixelShader9** ppShader) PURE;
    STDMETHOD(SetPixelShader)(THIS_ IDirect3DPixelShader9* pShader) PURE;
    STDMETHOD(GetPixelShader)(THIS_ IDirect3DPixelShader9** ppShader) PURE;
    STDMETHOD(SetPixelShaderConstantF)(THIS_ UINT StartRegister, CONST float* pConstantData, UINT Vector4fCount) PURE;
    STDMETHOD(GetPixelShaderConstantF)(THIS_ UINT StartRegister, float* pConstantData, UINT Vector4fCount) PURE;
    STDMETHOD(SetPixelShaderConstantI)(THIS_ UINT StartRegister, CONST int* pConstantData, UINT Vector4iCount) PURE;
    STDMETHOD(GetPixelShaderConstantI)(THIS_ UINT StartRegister, int* pConstantData, UINT Vector4iCount) PURE;
    STDMETHOD(SetPixelShaderConstantB)(THIS_ UINT StartRegister, CONST WINBOOL* pConstantData, UINT  BoolCount) PURE;
    STDMETHOD(GetPixelShaderConstantB)(THIS_ UINT StartRegister, WINBOOL* pConstantData, UINT BoolCount) PURE;
    STDMETHOD(DrawRectPatch)(THIS_ UINT Handle, CONST float* pNumSegs, CONST D3DRECTPATCH_INFO* pRectPatchInfo) PURE;
    STDMETHOD(DrawTriPatch)(THIS_ UINT Handle, CONST float* pNumSegs, CONST D3DTRIPATCH_INFO* pTriPatchInfo) PURE;
    STDMETHOD(DeletePatch)(THIS_ UINT Handle) PURE;
    STDMETHOD(CreateQuery)(THIS_ D3DQUERYTYPE Type, IDirect3DQuery9** ppQuery) PURE;
};
#undef INTERFACE

#if !defined(__cplusplus) || defined(CINTERFACE)
/*** IUnknown methods ***/
#define IDirect3DDevice9_QueryInterface(p,a,b) (p)->lpVtbl->QueryInterface(p,a,b)
#define IDirect3DDevice9_AddRef(p)             (p)->lpVtbl->AddRef(p)
#define IDirect3DDevice9_Release(p)            (p)->lpVtbl->Release(p)
/*** IDirect3DDevice9 methods ***/
#define IDirect3DDevice9_TestCooperativeLevel(p)                       (p)->lpVtbl->TestCooperativeLevel(p)
#define IDirect3DDevice9_GetAvailableTextureMem(p)                     (p)->lpVtbl->GetAvailableTextureMem(p)
#define IDirect3DDevice9_EvictManagedResources(p)                      (p)->lpVtbl->EvictManagedResources(p)
#define IDirect3DDevice9_GetDirect3D(p,a)                              (p)->lpVtbl->GetDirect3D(p,a)
#define IDirect3DDevice9_GetDeviceCaps(p,a)                            (p)->lpVtbl->GetDeviceCaps(p,a)
#define IDirect3DDevice9_GetDisplayMode(p,a,b)                         (p)->lpVtbl->GetDisplayMode(p,a,b)
#define IDirect3DDevice9_GetCreationParameters(p,a)                    (p)->lpVtbl->GetCreationParameters(p,a)
#define IDirect3DDevice9_SetCursorProperties(p,a,b,c)                  (p)->lpVtbl->SetCursorProperties(p,a,b,c)
#define IDirect3DDevice9_SetCursorPosition(p,a,b,c)                    (p)->lpVtbl->SetCursorPosition(p,a,b,c)
#define IDirect3DDevice9_ShowCursor(p,a)                               (p)->lpVtbl->ShowCursor(p,a)
#define IDirect3DDevice9_CreateAdditionalSwapChain(p,a,b)              (p)->lpVtbl->CreateAdditionalSwapChain(p,a,b)
#define IDirect3DDevice9_GetSwapChain(p,a,b)                           (p)->lpVtbl->GetSwapChain(p,a,b)
#define IDirect3DDevice9_GetNumberOfSwapChains(p)                      (p)->lpVtbl->GetNumberOfSwapChains(p)
#define IDirect3DDevice9_Reset(p,a)                                    (p)->lpVtbl->Reset(p,a)
#define IDirect3DDevice9_Present(p,a,b,c,d)                            (p)->lpVtbl->Present(p,a,b,c,d)
#define IDirect3DDevice9_GetBackBuffer(p,a,b,c,d)                      (p)->lpVtbl->GetBackBuffer(p,a,b,c,d)
#define IDirect3DDevice9_GetRasterStatus(p,a,b)                        (p)->lpVtbl->GetRasterStatus(p,a,b)
#define IDirect3DDevice9_SetDialogBoxMode(p,a)                         (p)->lpVtbl->SetDialogBoxMode(p,a)
#define IDirect3DDevice9_SetGammaRamp(p,a,b,c)                         (p)->lpVtbl->SetGammaRamp(p,a,b,c)
#define IDirect3DDevice9_GetGammaRamp(p,a,b)                           (p)->lpVtbl->GetGammaRamp(p,a,b)
#define IDirect3DDevice9_CreateTexture(p,a,b,c,d,e,f,g,h)              (p)->lpVtbl->CreateTexture(p,a,b,c,d,e,f,g,h)
#define IDirect3DDevice9_CreateVolumeTexture(p,a,b,c,d,e,f,g,h,i)      (p)->lpVtbl->CreateVolumeTexture(p,a,b,c,d,e,f,g,h,i)
#define IDirect3DDevice9_CreateCubeTexture(p,a,b,c,d,e,f,g)            (p)->lpVtbl->CreateCubeTexture(p,a,b,c,d,e,f,g)
#define IDirect3DDevice9_CreateVertexBuffer(p,a,b,c,d,e,f)             (p)->lpVtbl->CreateVertexBuffer(p,a,b,c,d,e,f)
#define IDirect3DDevice9_CreateIndexBuffer(p,a,b,c,d,e,f)              (p)->lpVtbl->CreateIndexBuffer(p,a,b,c,d,e,f)
#define IDirect3DDevice9_CreateRenderTarget(p,a,b,c,d,e,f,g,h)         (p)->lpVtbl->CreateRenderTarget(p,a,b,c,d,e,f,g,h)
#define IDirect3DDevice9_CreateDepthStencilSurface(p,a,b,c,d,e,f,g,h)  (p)->lpVtbl->CreateDepthStencilSurface(p,a,b,c,d,e,f,g,h)
#define IDirect3DDevice9_UpdateSurface(p,a,b,c,d)                      (p)->lpVtbl->UpdateSurface(p,a,b,c,d)
#define IDirect3DDevice9_UpdateTexture(p,a,b)                          (p)->lpVtbl->UpdateTexture(p,a,b)
#define IDirect3DDevice9_GetRenderTargetData(p,a,b)                    (p)->lpVtbl->GetRenderTargetData(p,a,b)
#define IDirect3DDevice9_GetFrontBufferData(p,a,b)                     (p)->lpVtbl->GetFrontBufferData(p,a,b)
#define IDirect3DDevice9_StretchRect(p,a,b,c,d,e)                      (p)->lpVtbl->StretchRect(p,a,b,c,d,e)
#define IDirect3DDevice9_ColorFill(p,a,b,c)                            (p)->lpVtbl->ColorFill(p,a,b,c)
#define IDirect3DDevice9_CreateOffscreenPlainSurface(p,a,b,c,d,e,f)    (p)->lpVtbl->CreateOffscreenPlainSurface(p,a,b,c,d,e,f)
#define IDirect3DDevice9_SetRenderTarget(p,a,b)                        (p)->lpVtbl->SetRenderTarget(p,a,b)
#define IDirect3DDevice9_GetRenderTarget(p,a,b)                        (p)->lpVtbl->GetRenderTarget(p,a,b)
#define IDirect3DDevice9_SetDepthStencilSurface(p,a)                   (p)->lpVtbl->SetDepthStencilSurface(p,a)
#define IDirect3DDevice9_GetDepthStencilSurface(p,a)                   (p)->lpVtbl->GetDepthStencilSurface(p,a)
#define IDirect3DDevice9_BeginScene(p)                                 (p)->lpVtbl->BeginScene(p)
#define IDirect3DDevice9_EndScene(p)                                   (p)->lpVtbl->EndScene(p)
#define IDirect3DDevice9_Clear(p,a,b,c,d,e,f)                          (p)->lpVtbl->Clear(p,a,b,c,d,e,f)
#define IDirect3DDevice9_SetTransform(p,a,b)                           (p)->lpVtbl->SetTransform(p,a,b)
#define IDirect3DDevice9_GetTransform(p,a,b)                           (p)->lpVtbl->GetTransform(p,a,b)
#define IDirect3DDevice9_MultiplyTransform(p,a,b)                      (p)->lpVtbl->MultiplyTransform(p,a,b)
#define IDirect3DDevice9_SetViewport(p,a)                              (p)->lpVtbl->SetViewport(p,a)
#define IDirect3DDevice9_GetViewport(p,a)                              (p)->lpVtbl->GetViewport(p,a)
#define IDirect3DDevice9_SetMaterial(p,a)                              (p)->lpVtbl->SetMaterial(p,a)
#define IDirect3DDevice9_GetMaterial(p,a)                              (p)->lpVtbl->GetMaterial(p,a)
#define IDirect3DDevice9_SetLight(p,a,b)                               (p)->lpVtbl->SetLight(p,a,b)
#define IDirect3DDevice9_GetLight(p,a,b)                               (p)->lpVtbl->GetLight(p,a,b)
#define IDirect3DDevice9_LightEnable(p,a,b)                            (p)->lpVtbl->LightEnable(p,a,b)
#define IDirect3DDevice9_GetLightEnable(p,a,b)                         (p)->lpVtbl->GetLightEnable(p,a,b)
#define IDirect3DDevice9_SetClipPlane(p,a,b)                           (p)->lpVtbl->SetClipPlane(p,a,b)
#define IDirect3DDevice9_GetClipPlane(p,a,b)                           (p)->lpVtbl->GetClipPlane(p,a,b)
#define IDirect3DDevice9_SetRenderState(p,a,b)                         (p)->lpVtbl->SetRenderState(p,a,b)
#define IDirect3DDevice9_GetRenderState(p,a,b)                         (p)->lpVtbl->GetRenderState(p,a,b)
#define IDirect3DDevice9_CreateStateBlock(p,a,b)                       (p)->lpVtbl->CreateStateBlock(p,a,b)
#define IDirect3DDevice9_BeginStateBlock(p)                            (p)->lpVtbl->BeginStateBlock(p)
#define IDirect3DDevice9_EndStateBlock(p,a)                            (p)->lpVtbl->EndStateBlock(p,a)
#define IDirect3DDevice9_SetClipStatus(p,a)                            (p)->lpVtbl->SetClipStatus(p,a)
#define IDirect3DDevice9_GetClipStatus(p,a)                            (p)->lpVtbl->GetClipStatus(p,a)
#define IDirect3DDevice9_GetTexture(p,a,b)                             (p)->lpVtbl->GetTexture(p,a,b)
#define IDirect3DDevice9_SetTexture(p,a,b)                             (p)->lpVtbl->SetTexture(p,a,b)
#define IDirect3DDevice9_GetTextureStageState(p,a,b,c)                 (p)->lpVtbl->GetTextureStageState(p,a,b,c)
#define IDirect3DDevice9_SetTextureStageState(p,a,b,c)                 (p)->lpVtbl->SetTextureStageState(p,a,b,c)
#define IDirect3DDevice9_GetSamplerState(p,a,b,c)                      (p)->lpVtbl->GetSamplerState(p,a,b,c)
#define IDirect3DDevice9_SetSamplerState(p,a,b,c)                      (p)->lpVtbl->SetSamplerState(p,a,b,c)
#define IDirect3DDevice9_ValidateDevice(p,a)                           (p)->lpVtbl->ValidateDevice(p,a)
#define IDirect3DDevice9_SetPaletteEntries(p,a,b)                      (p)->lpVtbl->SetPaletteEntries(p,a,b)
#define IDirect3DDevice9_GetPaletteEntries(p,a,b)                      (p)->lpVtbl->GetPaletteEntries(p,a,b)
#define IDirect3DDevice9_SetCurrentTexturePalette(p,a)                 (p)->lpVtbl->SetCurrentTexturePalette(p,a)
#define IDirect3DDevice9_GetCurrentTexturePalette(p,a)                 (p)->lpVtbl->GetCurrentTexturePalette(p,a)
#define IDirect3DDevice9_SetScissorRect(p,a)                           (p)->lpVtbl->SetScissorRect(p,a)
#define IDirect3DDevice9_GetScissorRect(p,a)                           (p)->lpVtbl->GetScissorRect(p,a)
#define IDirect3DDevice9_SetSoftwareVertexProcessing(p,a)              (p)->lpVtbl->SetSoftwareVertexProcessing(p,a)
#define IDirect3DDevice9_GetSoftwareVertexProcessing(p)                (p)->lpVtbl->GetSoftwareVertexProcessing(p)
#define IDirect3DDevice9_SetNPatchMode(p,a)                            (p)->lpVtbl->SetNPatchMode(p,a)
#define IDirect3DDevice9_GetNPatchMode(p)                              (p)->lpVtbl->GetNPatchMode(p)
#define IDirect3DDevice9_DrawPrimitive(p,a,b,c)                        (p)->lpVtbl->DrawPrimitive(p,a,b,c)
#define IDirect3DDevice9_DrawIndexedPrimitive(p,a,b,c,d,e,f)           (p)->lpVtbl->DrawIndexedPrimitive(p,a,b,c,d,e,f)
#define IDirect3DDevice9_DrawPrimitiveUP(p,a,b,c,d)                    (p)->lpVtbl->DrawPrimitiveUP(p,a,b,c,d)
#define IDirect3DDevice9_DrawIndexedPrimitiveUP(p,a,b,c,d,e,f,g,h)     (p)->lpVtbl->DrawIndexedPrimitiveUP(p,a,b,c,d,e,f,g,h)
#define IDirect3DDevice9_ProcessVertices(p,a,b,c,d,e,f)                (p)->lpVtbl->ProcessVertices(p,a,b,c,d,e,f)
#define IDirect3DDevice9_CreateVertexDeclaration(p,a,b)                (p)->lpVtbl->CreateVertexDeclaration(p,a,b)
#define IDirect3DDevice9_SetVertexDeclaration(p,a)                     (p)->lpVtbl->SetVertexDeclaration(p,a)
#define IDirect3DDevice9_GetVertexDeclaration(p,a)                     (p)->lpVtbl->GetVertexDeclaration(p,a)
#define IDirect3DDevice9_SetFVF(p,a)                                   (p)->lpVtbl->SetFVF(p,a)
#define IDirect3DDevice9_GetFVF(p,a)                                   (p)->lpVtbl->GetFVF(p,a)
#define IDirect3DDevice9_CreateVertexShader(p,a,b)                     (p)->lpVtbl->CreateVertexShader(p,a,b)
#define IDirect3DDevice9_SetVertexShader(p,a)                          (p)->lpVtbl->SetVertexShader(p,a)
#define IDirect3DDevice9_GetVertexShader(p,a)                          (p)->lpVtbl->GetVertexShader(p,a)
#define IDirect3DDevice9_SetVertexShaderConstantF(p,a,b,c)             (p)->lpVtbl->SetVertexShaderConstantF(p,a,b,c)
#define IDirect3DDevice9_GetVertexShaderConstantF(p,a,b,c)             (p)->lpVtbl->GetVertexShaderConstantF(p,a,b,c)
#define IDirect3DDevice9_SetVertexShaderConstantI(p,a,b,c)             (p)->lpVtbl->SetVertexShaderConstantI(p,a,b,c)
#define IDirect3DDevice9_GetVertexShaderConstantI(p,a,b,c)             (p)->lpVtbl->GetVertexShaderConstantI(p,a,b,c)
#define IDirect3DDevice9_SetVertexShaderConstantB(p,a,b,c)             (p)->lpVtbl->SetVertexShaderConstantB(p,a,b,c)
#define IDirect3DDevice9_GetVertexShaderConstantB(p,a,b,c)             (p)->lpVtbl->GetVertexShaderConstantB(p,a,b,c)
#define IDirect3DDevice9_SetStreamSource(p,a,b,c,d)                    (p)->lpVtbl->SetStreamSource(p,a,b,c,d)
#define IDirect3DDevice9_GetStreamSource(p,a,b,c,d)                    (p)->lpVtbl->GetStreamSource(p,a,b,c,d)
#define IDirect3DDevice9_SetStreamSourceFreq(p,a,b)                    (p)->lpVtbl->SetStreamSourceFreq(p,a,b)
#define IDirect3DDevice9_GetStreamSourceFreq(p,a,b)                    (p)->lpVtbl->GetStreamSourceFreq(p,a,b)
#define IDirect3DDevice9_SetIndices(p,a)                               (p)->lpVtbl->SetIndices(p,a)
#define IDirect3DDevice9_GetIndices(p,a)                               (p)->lpVtbl->GetIndices(p,a)
#define IDirect3DDevice9_CreatePixelShader(p,a,b)                      (p)->lpVtbl->CreatePixelShader(p,a,b)
#define IDirect3DDevice9_SetPixelShader(p,a)                           (p)->lpVtbl->SetPixelShader(p,a)
#define IDirect3DDevice9_GetPixelShader(p,a)                           (p)->lpVtbl->GetPixelShader(p,a)
#define IDirect3DDevice9_SetPixelShaderConstantF(p,a,b,c)              (p)->lpVtbl->SetPixelShaderConstantF(p,a,b,c)
#define IDirect3DDevice9_GetPixelShaderConstantF(p,a,b,c)              (p)->lpVtbl->GetPixelShaderConstantF(p,a,b,c)
#define IDirect3DDevice9_SetPixelShaderConstantI(p,a,b,c)              (p)->lpVtbl->SetPixelShaderConstantI(p,a,b,c)
#define IDirect3DDevice9_GetPixelShaderConstantI(p,a,b,c)              (p)->lpVtbl->GetPixelShaderConstantI(p,a,b,c)
#define IDirect3DDevice9_SetPixelShaderConstantB(p,a,b,c)              (p)->lpVtbl->SetPixelShaderConstantB(p,a,b,c)
#define IDirect3DDevice9_GetPixelShaderConstantB(p,a,b,c)              (p)->lpVtbl->GetPixelShaderConstantB(p,a,b,c)
#define IDirect3DDevice9_DrawRectPatch(p,a,b,c)                        (p)->lpVtbl->DrawRectPatch(p,a,b,c)
#define IDirect3DDevice9_DrawTriPatch(p,a,b,c)                         (p)->lpVtbl->DrawTriPatch(p,a,b,c)
#define IDirect3DDevice9_DeletePatch(p,a)                              (p)->lpVtbl->DeletePatch(p,a)
#define IDirect3DDevice9_CreateQuery(p,a,b)                            (p)->lpVtbl->CreateQuery(p,a,b)
#else
/*** IUnknown methods ***/
#define IDirect3DDevice9_QueryInterface(p,a,b) (p)->QueryInterface(a,b)
#define IDirect3DDevice9_AddRef(p)             (p)->AddRef()
#define IDirect3DDevice9_Release(p)            (p)->Release()
/*** IDirect3DDevice9 methods ***/
#define IDirect3DDevice9_TestCooperativeLevel(p)                       (p)->TestCooperativeLevel()
#define IDirect3DDevice9_GetAvailableTextureMem(p)                     (p)->GetAvailableTextureMem()
#define IDirect3DDevice9_EvictManagedResources(p)                      (p)->EvictManagedResources()
#define IDirect3DDevice9_GetDirect3D(p,a)                              (p)->GetDirect3D(a)
#define IDirect3DDevice9_GetDeviceCaps(p,a)                            (p)->GetDeviceCaps(a)
#define IDirect3DDevice9_GetDisplayMode(p,a,b)                         (p)->GetDisplayMode(a,b)
#define IDirect3DDevice9_GetCreationParameters(p,a)                    (p)->GetCreationParameters(a)
#define IDirect3DDevice9_SetCursorProperties(p,a,b,c)                  (p)->SetCursorProperties(a,b,c)
#define IDirect3DDevice9_SetCursorPosition(p,a,b,c)                    (p)->SetCursorPosition(a,b,c)
#define IDirect3DDevice9_ShowCursor(p,a)                               (p)->ShowCursor(a)
#define IDirect3DDevice9_CreateAdditionalSwapChain(p,a,b)              (p)->CreateAdditionalSwapChain(a,b)
#define IDirect3DDevice9_GetSwapChain(p,a,b)                           (p)->GetSwapChain(a,b)
#define IDirect3DDevice9_GetNumberOfSwapChains(p)                      (p)->GetNumberOfSwapChains()
#define IDirect3DDevice9_Reset(p,a)                                    (p)->Reset(a)
#define IDirect3DDevice9_Present(p,a,b,c,d)                            (p)->Present(a,b,c,d)
#define IDirect3DDevice9_GetBackBuffer(p,a,b,c,d)                      (p)->GetBackBuffer(a,b,c,d)
#define IDirect3DDevice9_GetRasterStatus(p,a,b)                        (p)->GetRasterStatus(a,b)
#define IDirect3DDevice9_SetDialogBoxMode(p,a)                         (p)->SetDialogBoxMode(a)
#define IDirect3DDevice9_SetGammaRamp(p,a,b,c)                         (p)->SetGammaRamp(a,b,c)
#define IDirect3DDevice9_GetGammaRamp(p,a,b)                           (p)->GetGammaRamp(a,b)
#define IDirect3DDevice9_CreateTexture(p,a,b,c,d,e,f,g,h)              (p)->CreateTexture(a,b,c,d,e,f,g,h)
#define IDirect3DDevice9_CreateVolumeTexture(p,a,b,c,d,e,f,g,h,i)      (p)->CreateVolumeTexture(a,b,c,d,e,f,g,h,i)
#define IDirect3DDevice9_CreateCubeTexture(p,a,b,c,d,e,f,g)            (p)->CreateCubeTexture(a,b,c,d,e,f,g)
#define IDirect3DDevice9_CreateVertexBuffer(p,a,b,c,d,e,f)             (p)->CreateVertexBuffer(a,b,c,d,e,f)
#define IDirect3DDevice9_CreateIndexBuffer(p,a,b,c,d,e,f)              (p)->CreateIndexBuffer(a,b,c,d,e,f)
#define IDirect3DDevice9_CreateRenderTarget(p,a,b,c,d,e,f,g,h)         (p)->CreateRenderTarget(a,b,c,d,e,f,g,h)
#define IDirect3DDevice9_CreateDepthStencilSurface(p,a,b,c,d,e,f,g,h)  (p)->CreateDepthStencilSurface(a,b,c,d,e,f,g,h)
#define IDirect3DDevice9_UpdateSurface(p,a,b,c,d)                      (p)->UpdateSurface(a,b,c,d)
#define IDirect3DDevice9_UpdateTexture(p,a,b)                          (p)->UpdateTexture(a,b)
#define IDirect3DDevice9_GetRenderTargetData(p,a,b)                    (p)->GetRenderTargetData(a,b)
#define IDirect3DDevice9_GetFrontBufferData(p,a,b)                     (p)->GetFrontBufferData(a,b)
#define IDirect3DDevice9_StretchRect(p,a,b,c,d,e)                      (p)->StretchRect(a,b,c,d,e)
#define IDirect3DDevice9_ColorFill(p,a,b,c)                            (p)->ColorFill(a,b,c)
#define IDirect3DDevice9_CreateOffscreenPlainSurface(p,a,b,c,d,e,f)    (p)->CreateOffscreenPlainSurface(a,b,c,d,e,f)
#define IDirect3DDevice9_SetRenderTarget(p,a,b)                        (p)->SetRenderTarget(a,b)
#define IDirect3DDevice9_GetRenderTarget(p,a,b)                        (p)->GetRenderTarget(a,b)
#define IDirect3DDevice9_SetDepthStencilSurface(p,a)                   (p)->SetDepthStencilSurface(a)
#define IDirect3DDevice9_GetDepthStencilSurface(p,a)                   (p)->GetDepthStencilSurface(a)
#define IDirect3DDevice9_BeginScene(p)                                 (p)->BeginScene()
#define IDirect3DDevice9_EndScene(p)                                   (p)->EndScene()
#define IDirect3DDevice9_Clear(p,a,b,c,d,e,f)                          (p)->Clear(a,b,c,d,e,f)
#define IDirect3DDevice9_SetTransform(p,a,b)                           (p)->SetTransform(a,b)
#define IDirect3DDevice9_GetTransform(p,a,b)                           (p)->GetTransform(a,b)
#define IDirect3DDevice9_MultiplyTransform(p,a,b)                      (p)->MultiplyTransform(a,b)
#define IDirect3DDevice9_SetViewport(p,a)                              (p)->SetViewport(a)
#define IDirect3DDevice9_GetViewport(p,a)                              (p)->GetViewport(a)
#define IDirect3DDevice9_SetMaterial(p,a)                              (p)->SetMaterial(a)
#define IDirect3DDevice9_GetMaterial(p,a)                              (p)->GetMaterial(a)
#define IDirect3DDevice9_SetLight(p,a,b)                               (p)->SetLight(a,b)
#define IDirect3DDevice9_GetLight(p,a,b)                               (p)->GetLight(a,b)
#define IDirect3DDevice9_LightEnable(p,a,b)                            (p)->LightEnable(a,b)
#define IDirect3DDevice9_GetLightEnable(p,a,b)                         (p)->GetLightEnable(a,b)
#define IDirect3DDevice9_SetClipPlane(p,a,b)                           (p)->SetClipPlane(a,b)
#define IDirect3DDevice9_GetClipPlane(p,a,b)                           (p)->GetClipPlane(a,b)
#define IDirect3DDevice9_SetRenderState(p,a,b)                         (p)->SetRenderState(a,b)
#define IDirect3DDevice9_GetRenderState(p,a,b)                         (p)->GetRenderState(a,b)
#define IDirect3DDevice9_CreateStateBlock(p,a,b)                       (p)->CreateStateBlock(a,b)
#define IDirect3DDevice9_BeginStateBlock(p)                            (p)->BeginStateBlock()
#define IDirect3DDevice9_EndStateBlock(p,a)                            (p)->EndStateBlock(a)
#define IDirect3DDevice9_SetClipStatus(p,a)                            (p)->SetClipStatus(a)
#define IDirect3DDevice9_GetClipStatus(p,a)                            (p)->GetClipStatus(a)
#define IDirect3DDevice9_GetTexture(p,a,b)                             (p)->GetTexture(a,b)
#define IDirect3DDevice9_SetTexture(p,a,b)                             (p)->SetTexture(a,b)
#define IDirect3DDevice9_GetTextureStageState(p,a,b,c)                 (p)->GetTextureStageState(a,b,c)
#define IDirect3DDevice9_SetTextureStageState(p,a,b,c)                 (p)->SetTextureStageState(a,b,c)
#define IDirect3DDevice9_GetSamplerState(p,a,b,c)                      (p)->GetSamplerState(a,b,c)
#define IDirect3DDevice9_SetSamplerState(p,a,b,c)                      (p)->SetSamplerState(a,b,c)
#define IDirect3DDevice9_ValidateDevice(p,a)                           (p)->ValidateDevice(a)
#define IDirect3DDevice9_SetPaletteEntries(p,a,b)                      (p)->SetPaletteEntries(a,b)
#define IDirect3DDevice9_GetPaletteEntries(p,a,b)                      (p)->GetPaletteEntries(a,b)
#define IDirect3DDevice9_SetCurrentTexturePalette(p,a)                 (p)->SetCurrentTexturePalette(a)
#define IDirect3DDevice9_GetCurrentTexturePalette(p,a)                 (p)->GetCurrentTexturePalette(a)
#define IDirect3DDevice9_SetScissorRect(p,a)                           (p)->SetScissorRect(a)
#define IDirect3DDevice9_GetScissorRect(p,a)                           (p)->GetScissorRect(a)
#define IDirect3DDevice9_SetSoftwareVertexProcessing(p,a)              (p)->SetSoftwareVertexProcessing(a)
#define IDirect3DDevice9_GetSoftwareVertexProcessing(p)                (p)->GetSoftwareVertexProcessing()
#define IDirect3DDevice9_SetNPatchMode(p,a)                            (p)->SetNPatchMode(a)
#define IDirect3DDevice9_GetNPatchMode(p)                              (p)->GetNPatchMode()
#define IDirect3DDevice9_DrawPrimitive(p,a,b,c)                        (p)->DrawPrimitive(a,b,c)
#define IDirect3DDevice9_DrawIndexedPrimitive(p,a,b,c,d,e,f)           (p)->DrawIndexedPrimitive(a,b,c,d,e,f)
#define IDirect3DDevice9_DrawPrimitiveUP(p,a,b,c,d)                    (p)->DrawPrimitiveUP(a,b,c,d)
#define IDirect3DDevice9_DrawIndexedPrimitiveUP(p,a,b,c,d,e,f,g,h)     (p)->DrawIndexedPrimitiveUP(a,b,c,d,e,f,g,h)
#define IDirect3DDevice9_ProcessVertices(p,a,b,c,d,e,f)                (p)->ProcessVertices(a,b,c,d,e,f)
#define IDirect3DDevice9_CreateVertexDeclaration(p,a,b)                (p)->CreateVertexDeclaration(a,b)
#define IDirect3DDevice9_SetVertexDeclaration(p,a)                     (p)->SetVertexDeclaration(a)
#define IDirect3DDevice9_GetVertexDeclaration(p,a)                     (p)->GetVertexDeclaration(a)
#define IDirect3DDevice9_SetFVF(p,a)                                   (p)->SetFVF(a)
#define IDirect3DDevice9_GetFVF(p,a)                                   (p)->GetFVF(a)
#define IDirect3DDevice9_CreateVertexShader(p,a,b)                     (p)->CreateVertexShader(a,b)
#define IDirect3DDevice9_SetVertexShader(p,a)                          (p)->SetVertexShader(a)
#define IDirect3DDevice9_GetVertexShader(p,a)                          (p)->GetVertexShader(a)
#define IDirect3DDevice9_SetVertexShaderConstantF(p,a,b,c)             (p)->SetVertexShaderConstantF(a,b,c)
#define IDirect3DDevice9_GetVertexShaderConstantF(p,a,b,c)             (p)->GetVertexShaderConstantF(a,b,c)
#define IDirect3DDevice9_SetVertexShaderConstantI(p,a,b,c)             (p)->SetVertexShaderConstantI(a,b,c)
#define IDirect3DDevice9_GetVertexShaderConstantI(p,a,b,c)             (p)->GetVertexShaderConstantI(a,b,c)
#define IDirect3DDevice9_SetVertexShaderConstantB(p,a,b,c)             (p)->SetVertexShaderConstantB(a,b,c)
#define IDirect3DDevice9_GetVertexShaderConstantB(p,a,b,c)             (p)->GetVertexShaderConstantB(a,b,c)
#define IDirect3DDevice9_SetStreamSource(p,a,b,c,d)                    (p)->SetStreamSource(a,b,c,d)
#define IDirect3DDevice9_GetStreamSource(p,a,b,c,d)                    (p)->GetStreamSource(a,b,c,d)
#define IDirect3DDevice9_SetStreamSourceFreq(p,a,b)                    (p)->SetStreamSourceFreq(a,b)
#define IDirect3DDevice9_GetStreamSourceFreq(p,a,b)                    (p)->GetStreamSourceFreq(a,b)
#define IDirect3DDevice9_SetIndices(p,a)                               (p)->SetIndices(a)
#define IDirect3DDevice9_GetIndices(p,a)                               (p)->GetIndices(a)
#define IDirect3DDevice9_CreatePixelShader(p,a,b)                      (p)->CreatePixelShader(a,b)
#define IDirect3DDevice9_SetPixelShader(p,a)                           (p)->SetPixelShader(a)
#define IDirect3DDevice9_GetPixelShader(p,a)                           (p)->GetPixelShader(a)
#define IDirect3DDevice9_SetPixelShaderConstantF(p,a,b,c)              (p)->SetPixelShaderConstantF(a,b,c)
#define IDirect3DDevice9_GetPixelShaderConstantF(p,a,b,c)              (p)->GetPixelShaderConstantF(a,b,c)
#define IDirect3DDevice9_SetPixelShaderConstantI(p,a,b,c)              (p)->SetPixelShaderConstantI(a,b,c)
#define IDirect3DDevice9_GetPixelShaderConstantI(p,a,b,c)              (p)->GetPixelShaderConstantI(a,b,c)
#define IDirect3DDevice9_SetPixelShaderConstantB(p,a,b,c)              (p)->SetPixelShaderConstantB(a,b,c)
#define IDirect3DDevice9_GetPixelShaderConstantB(p,a,b,c)              (p)->GetPixelShaderConstantB(a,b,c)
#define IDirect3DDevice9_DrawRectPatch(p,a,b,c)                        (p)->DrawRectPatch(a,b,c)
#define IDirect3DDevice9_DrawTriPatch(p,a,b,c)                         (p)->DrawTriPatch(a,b,c)
#define IDirect3DDevice9_DeletePatch(p,a)                              (p)->DeletePatch(a)
#define IDirect3DDevice9_CreateQuery(p,a,b)                            (p)->CreateQuery(a,b)
#endif


/*****************************************************************************
 * IDirect3DDevice9Ex interface
 */
#define INTERFACE IDirect3DDevice9Ex
DECLARE_INTERFACE_(IDirect3DDevice9Ex,IDirect3DDevice9)
{
    /*** IUnknown methods ***/
    STDMETHOD_(HRESULT,QueryInterface)(THIS_ REFIID riid, void** ppvObject) PURE;
    STDMETHOD_(ULONG,AddRef)(THIS) PURE;
    STDMETHOD_(ULONG,Release)(THIS) PURE;
    /*** IDirect3DDevice9 methods ***/
    STDMETHOD(TestCooperativeLevel)(THIS) PURE;
    STDMETHOD_(UINT, GetAvailableTextureMem)(THIS) PURE;
    STDMETHOD(EvictManagedResources)(THIS) PURE;
    STDMETHOD(GetDirect3D)(THIS_ IDirect3D9** ppD3D9) PURE;
    STDMETHOD(GetDeviceCaps)(THIS_ D3DCAPS9* pCaps) PURE;
    STDMETHOD(GetDisplayMode)(THIS_ UINT iSwapChain, D3DDISPLAYMODE* pMode) PURE;
    STDMETHOD(GetCreationParameters)(THIS_ D3DDEVICE_CREATION_PARAMETERS *pParameters) PURE;
    STDMETHOD(SetCursorProperties)(THIS_ UINT XHotSpot, UINT YHotSpot, IDirect3DSurface9* pCursorBitmap) PURE;
    STDMETHOD_(void, SetCursorPosition)(THIS_ int X,int Y, DWORD Flags) PURE;
    STDMETHOD_(WINBOOL, ShowCursor)(THIS_ WINBOOL bShow) PURE;
    STDMETHOD(CreateAdditionalSwapChain)(THIS_ D3DPRESENT_PARAMETERS* pPresentationParameters, IDirect3DSwapChain9** pSwapChain) PURE;
    STDMETHOD(GetSwapChain)(THIS_ UINT iSwapChain, IDirect3DSwapChain9** pSwapChain) PURE;
    STDMETHOD_(UINT, GetNumberOfSwapChains)(THIS) PURE;
    STDMETHOD(Reset)(THIS_ D3DPRESENT_PARAMETERS* pPresentationParameters) PURE;
    STDMETHOD(Present)(THIS_ CONST RECT* pSourceRect, CONST RECT* pDestRect, HWND hDestWindowOverride, CONST RGNDATA* pDirtyRegion) PURE;
    STDMETHOD(GetBackBuffer)(THIS_ UINT iSwapChain, UINT iBackBuffer, D3DBACKBUFFER_TYPE Type, IDirect3DSurface9** ppBackBuffer) PURE;
    STDMETHOD(GetRasterStatus)(THIS_ UINT iSwapChain, D3DRASTER_STATUS* pRasterStatus) PURE;
    STDMETHOD(SetDialogBoxMode)(THIS_ WINBOOL bEnableDialogs) PURE;
    STDMETHOD_(void, SetGammaRamp)(THIS_ UINT iSwapChain, DWORD Flags, CONST D3DGAMMARAMP* pRamp) PURE;
    STDMETHOD_(void, GetGammaRamp)(THIS_ UINT iSwapChain, D3DGAMMARAMP* pRamp) PURE;
    STDMETHOD(CreateTexture)(THIS_ UINT Width, UINT Height, UINT Levels, DWORD Usage, D3DFORMAT Format, D3DPOOL Pool, IDirect3DTexture9** ppTexture, HANDLE* pSharedHandle) PURE;
    STDMETHOD(CreateVolumeTexture)(THIS_ UINT Width, UINT Height, UINT Depth, UINT Levels, DWORD Usage, D3DFORMAT Format, D3DPOOL Pool, IDirect3DVolumeTexture9** ppVolumeTexture, HANDLE* pSharedHandle) PURE;
    STDMETHOD(CreateCubeTexture)(THIS_ UINT EdgeLength, UINT Levels, DWORD Usage, D3DFORMAT Format, D3DPOOL Pool, IDirect3DCubeTexture9** ppCubeTexture, HANDLE* pSharedHandle) PURE;
    STDMETHOD(CreateVertexBuffer)(THIS_ UINT Length, DWORD Usage, DWORD FVF, D3DPOOL Pool, IDirect3DVertexBuffer9** ppVertexBuffer, HANDLE* pSharedHandle) PURE;
    STDMETHOD(CreateIndexBuffer)(THIS_ UINT Length, DWORD Usage, D3DFORMAT Format, D3DPOOL Pool, IDirect3DIndexBuffer9** ppIndexBuffer, HANDLE* pSharedHandle) PURE;
    STDMETHOD(CreateRenderTarget)(THIS_ UINT Width, UINT Height, D3DFORMAT Format, D3DMULTISAMPLE_TYPE MultiSample, DWORD MultisampleQuality, WINBOOL Lockable, IDirect3DSurface9** ppSurface, HANDLE* pSharedHandle) PURE;
    STDMETHOD(CreateDepthStencilSurface)(THIS_ UINT Width, UINT Height, D3DFORMAT Format, D3DMULTISAMPLE_TYPE MultiSample, DWORD MultisampleQuality, WINBOOL Discard, IDirect3DSurface9** ppSurface, HANDLE* pSharedHandle) PURE;
    STDMETHOD(UpdateSurface)(THIS_ IDirect3DSurface9* pSourceSurface, CONST RECT* pSourceRect, IDirect3DSurface9* pDestinationSurface, CONST POINT* pDestPoint) PURE;
    STDMETHOD(UpdateTexture)(THIS_ IDirect3DBaseTexture9* pSourceTexture, IDirect3DBaseTexture9* pDestinationTexture) PURE;
    STDMETHOD(GetRenderTargetData)(THIS_ IDirect3DSurface9* pRenderTarget, IDirect3DSurface9* pDestSurface) PURE;
    STDMETHOD(GetFrontBufferData)(THIS_ UINT iSwapChain, IDirect3DSurface9* pDestSurface) PURE;
    STDMETHOD(StretchRect)(THIS_ IDirect3DSurface9* pSourceSurface, CONST RECT* pSourceRect, IDirect3DSurface9* pDestSurface, CONST RECT* pDestRect, D3DTEXTUREFILTERTYPE Filter) PURE;
    STDMETHOD(ColorFill)(THIS_ IDirect3DSurface9* pSurface, CONST RECT* pRect, D3DCOLOR color) PURE;
    STDMETHOD(CreateOffscreenPlainSurface)(THIS_ UINT Width, UINT Height, D3DFORMAT Format, D3DPOOL Pool, IDirect3DSurface9** ppSurface, HANDLE* pSharedHandle) PURE;
    STDMETHOD(SetRenderTarget)(THIS_ DWORD RenderTargetIndex, IDirect3DSurface9* pRenderTarget) PURE;
    STDMETHOD(GetRenderTarget)(THIS_ DWORD RenderTargetIndex, IDirect3DSurface9** ppRenderTarget) PURE;
    STDMETHOD(SetDepthStencilSurface)(THIS_ IDirect3DSurface9* pNewZStencil) PURE;
    STDMETHOD(GetDepthStencilSurface)(THIS_ IDirect3DSurface9** ppZStencilSurface) PURE;
    STDMETHOD(BeginScene)(THIS) PURE;
    STDMETHOD(EndScene)(THIS) PURE;
    STDMETHOD(Clear)(THIS_ DWORD Count, CONST D3DRECT* pRects, DWORD Flags, D3DCOLOR Color, float Z, DWORD Stencil) PURE;
    STDMETHOD(SetTransform)(THIS_ D3DTRANSFORMSTATETYPE State, CONST D3DMATRIX* pMatrix) PURE;
    STDMETHOD(GetTransform)(THIS_ D3DTRANSFORMSTATETYPE State, D3DMATRIX* pMatrix) PURE;
    STDMETHOD(MultiplyTransform)(THIS_ D3DTRANSFORMSTATETYPE, CONST D3DMATRIX*) PURE;
    STDMETHOD(SetViewport)(THIS_ CONST D3DVIEWPORT9* pViewport) PURE;
    STDMETHOD(GetViewport)(THIS_ D3DVIEWPORT9* pViewport) PURE;
    STDMETHOD(SetMaterial)(THIS_ CONST D3DMATERIAL9* pMaterial) PURE;
    STDMETHOD(GetMaterial)(THIS_ D3DMATERIAL9* pMaterial) PURE;
    STDMETHOD(SetLight)(THIS_ DWORD Index, CONST D3DLIGHT9*) PURE;
    STDMETHOD(GetLight)(THIS_ DWORD Index, D3DLIGHT9*) PURE;
    STDMETHOD(LightEnable)(THIS_ DWORD Index, WINBOOL Enable) PURE;
    STDMETHOD(GetLightEnable)(THIS_ DWORD Index, WINBOOL* pEnable) PURE;
    STDMETHOD(SetClipPlane)(THIS_ DWORD Index, CONST float* pPlane) PURE;
    STDMETHOD(GetClipPlane)(THIS_ DWORD Index, float* pPlane) PURE;
    STDMETHOD(SetRenderState)(THIS_ D3DRENDERSTATETYPE State, DWORD Value) PURE;
    STDMETHOD(GetRenderState)(THIS_ D3DRENDERSTATETYPE State, DWORD* pValue) PURE;
    STDMETHOD(CreateStateBlock)(THIS_ D3DSTATEBLOCKTYPE Type, IDirect3DStateBlock9** ppSB) PURE;
    STDMETHOD(BeginStateBlock)(THIS) PURE;
    STDMETHOD(EndStateBlock)(THIS_ IDirect3DStateBlock9** ppSB) PURE;
    STDMETHOD(SetClipStatus)(THIS_ CONST D3DCLIPSTATUS9* pClipStatus) PURE;
    STDMETHOD(GetClipStatus)(THIS_ D3DCLIPSTATUS9* pClipStatus) PURE;
    STDMETHOD(GetTexture)(THIS_ DWORD Stage, IDirect3DBaseTexture9** ppTexture) PURE;
    STDMETHOD(SetTexture)(THIS_ DWORD Stage, IDirect3DBaseTexture9* pTexture) PURE;
    STDMETHOD(GetTextureStageState)(THIS_ DWORD Stage, D3DTEXTURESTAGESTATETYPE Type, DWORD* pValue) PURE;
    STDMETHOD(SetTextureStageState)(THIS_ DWORD Stage, D3DTEXTURESTAGESTATETYPE Type, DWORD Value) PURE;
    STDMETHOD(GetSamplerState)(THIS_ DWORD Sampler, D3DSAMPLERSTATETYPE Type, DWORD* pValue) PURE;
    STDMETHOD(SetSamplerState)(THIS_ DWORD Sampler, D3DSAMPLERSTATETYPE Type, DWORD Value) PURE;
    STDMETHOD(ValidateDevice)(THIS_ DWORD* pNumPasses) PURE;
    STDMETHOD(SetPaletteEntries)(THIS_ UINT PaletteNumber, CONST PALETTEENTRY* pEntries) PURE;
    STDMETHOD(GetPaletteEntries)(THIS_ UINT PaletteNumber,PALETTEENTRY* pEntries) PURE;
    STDMETHOD(SetCurrentTexturePalette)(THIS_ UINT PaletteNumber) PURE;
    STDMETHOD(GetCurrentTexturePalette)(THIS_ UINT *PaletteNumber) PURE;
    STDMETHOD(SetScissorRect)(THIS_ CONST RECT* pRect) PURE;
    STDMETHOD(GetScissorRect)(THIS_ RECT* pRect) PURE;
    STDMETHOD(SetSoftwareVertexProcessing)(THIS_ WINBOOL bSoftware) PURE;
    STDMETHOD_(WINBOOL, GetSoftwareVertexProcessing)(THIS) PURE;
    STDMETHOD(SetNPatchMode)(THIS_ float nSegments) PURE;
    STDMETHOD_(float, GetNPatchMode)(THIS) PURE;
    STDMETHOD(DrawPrimitive)(THIS_ D3DPRIMITIVETYPE PrimitiveType, UINT StartVertex, UINT PrimitiveCount) PURE;
    STDMETHOD(DrawIndexedPrimitive)(THIS_ D3DPRIMITIVETYPE, INT BaseVertexIndex, UINT MinVertexIndex, UINT NumVertices, UINT startIndex, UINT primCount) PURE;
    STDMETHOD(DrawPrimitiveUP)(THIS_ D3DPRIMITIVETYPE PrimitiveType, UINT PrimitiveCount, CONST void* pVertexStreamZeroData, UINT VertexStreamZeroStride) PURE;
    STDMETHOD(DrawIndexedPrimitiveUP)(THIS_ D3DPRIMITIVETYPE PrimitiveType, UINT MinVertexIndex, UINT NumVertices, UINT PrimitiveCount, CONST void* pIndexData, D3DFORMAT IndexDataFormat, CONST void* pVertexStreamZeroData, UINT VertexStreamZeroStride) PURE;
    STDMETHOD(ProcessVertices)(THIS_ UINT SrcStartIndex, UINT DestIndex, UINT VertexCount, IDirect3DVertexBuffer9* pDestBuffer, IDirect3DVertexDeclaration9* pVertexDecl, DWORD Flags) PURE;
    STDMETHOD(CreateVertexDeclaration)(THIS_ CONST D3DVERTEXELEMENT9* pVertexElements, IDirect3DVertexDeclaration9** ppDecl) PURE;
    STDMETHOD(SetVertexDeclaration)(THIS_ IDirect3DVertexDeclaration9* pDecl) PURE;
    STDMETHOD(GetVertexDeclaration)(THIS_ IDirect3DVertexDeclaration9** ppDecl) PURE;
    STDMETHOD(SetFVF)(THIS_ DWORD FVF) PURE;
    STDMETHOD(GetFVF)(THIS_ DWORD* pFVF) PURE;
    STDMETHOD(CreateVertexShader)(THIS_ CONST DWORD* pFunction, IDirect3DVertexShader9** ppShader) PURE;
    STDMETHOD(SetVertexShader)(THIS_ IDirect3DVertexShader9* pShader) PURE;
    STDMETHOD(GetVertexShader)(THIS_ IDirect3DVertexShader9** ppShader) PURE;
    STDMETHOD(SetVertexShaderConstantF)(THIS_ UINT StartRegister, CONST float* pConstantData, UINT Vector4fCount) PURE;
    STDMETHOD(GetVertexShaderConstantF)(THIS_ UINT StartRegister, float* pConstantData, UINT Vector4fCount) PURE;
    STDMETHOD(SetVertexShaderConstantI)(THIS_ UINT StartRegister, CONST int* pConstantData, UINT Vector4iCount) PURE;
    STDMETHOD(GetVertexShaderConstantI)(THIS_ UINT StartRegister, int* pConstantData, UINT Vector4iCount) PURE;
    STDMETHOD(SetVertexShaderConstantB)(THIS_ UINT StartRegister, CONST WINBOOL* pConstantData, UINT  BoolCount) PURE;
    STDMETHOD(GetVertexShaderConstantB)(THIS_ UINT StartRegister, WINBOOL* pConstantData, UINT BoolCount) PURE;
    STDMETHOD(SetStreamSource)(THIS_ UINT StreamNumber, IDirect3DVertexBuffer9* pStreamData, UINT OffsetInBytes, UINT Stride) PURE;
    STDMETHOD(GetStreamSource)(THIS_ UINT StreamNumber, IDirect3DVertexBuffer9** ppStreamData, UINT* OffsetInBytes, UINT* pStride) PURE;
    STDMETHOD(SetStreamSourceFreq)(THIS_ UINT StreamNumber, UINT Divider) PURE;
    STDMETHOD(GetStreamSourceFreq)(THIS_ UINT StreamNumber, UINT* Divider) PURE;
    STDMETHOD(SetIndices)(THIS_ IDirect3DIndexBuffer9* pIndexData) PURE;
    STDMETHOD(GetIndices)(THIS_ IDirect3DIndexBuffer9** ppIndexData) PURE;
    STDMETHOD(CreatePixelShader)(THIS_ CONST DWORD* pFunction, IDirect3DPixelShader9** ppShader) PURE;
    STDMETHOD(SetPixelShader)(THIS_ IDirect3DPixelShader9* pShader) PURE;
    STDMETHOD(GetPixelShader)(THIS_ IDirect3DPixelShader9** ppShader) PURE;
    STDMETHOD(SetPixelShaderConstantF)(THIS_ UINT StartRegister, CONST float* pConstantData, UINT Vector4fCount) PURE;
    STDMETHOD(GetPixelShaderConstantF)(THIS_ UINT StartRegister, float* pConstantData, UINT Vector4fCount) PURE;
    STDMETHOD(SetPixelShaderConstantI)(THIS_ UINT StartRegister, CONST int* pConstantData, UINT Vector4iCount) PURE;
    STDMETHOD(GetPixelShaderConstantI)(THIS_ UINT StartRegister, int* pConstantData, UINT Vector4iCount) PURE;
    STDMETHOD(SetPixelShaderConstantB)(THIS_ UINT StartRegister, CONST WINBOOL* pConstantData, UINT  BoolCount) PURE;
    STDMETHOD(GetPixelShaderConstantB)(THIS_ UINT StartRegister, WINBOOL* pConstantData, UINT BoolCount) PURE;
    STDMETHOD(DrawRectPatch)(THIS_ UINT Handle, CONST float* pNumSegs, CONST D3DRECTPATCH_INFO* pRectPatchInfo) PURE;
    STDMETHOD(DrawTriPatch)(THIS_ UINT Handle, CONST float* pNumSegs, CONST D3DTRIPATCH_INFO* pTriPatchInfo) PURE;
    STDMETHOD(DeletePatch)(THIS_ UINT Handle) PURE;
    STDMETHOD(CreateQuery)(THIS_ D3DQUERYTYPE Type, IDirect3DQuery9** ppQuery) PURE;
    /* IDirect3DDevice9Ex methods */
#if !defined(D3D_DISABLE_9EX)
    STDMETHOD(SetConvolutionMonoKernel)(THIS_ UINT width, UINT height, float *rows, float *columns) PURE;
    STDMETHOD(ComposeRects)(THIS_ IDirect3DSurface9 *src_surface, IDirect3DSurface9 *dst_surface,
            IDirect3DVertexBuffer9 *src_descs, UINT rect_count, IDirect3DVertexBuffer9 *dst_descs,
            D3DCOMPOSERECTSOP operation, INT offset_x, INT offset_y) PURE;
    STDMETHOD(PresentEx)(THIS_ CONST RECT *pSourceRect, CONST RECT *pDestRect, HWND hDestWindowOverride, CONST RGNDATA *pDirtyRegion, DWORD dwFlags) PURE;
    STDMETHOD(GetGPUThreadPriority)(THIS_ INT *pPriority) PURE;
    STDMETHOD(SetGPUThreadPriority)(THIS_ INT Priority) PURE;
    STDMETHOD(WaitForVBlank)(THIS_ UINT iSwapChain) PURE;
    STDMETHOD(CheckResourceResidency)(THIS_ IDirect3DResource9 **resources, UINT32 resource_count) PURE;
    STDMETHOD(SetMaximumFrameLatency)(THIS_ UINT MaxLatency) PURE;
    STDMETHOD(GetMaximumFrameLatency)(THIS_ UINT *pMaxLatenxy) PURE;
    STDMETHOD(CheckDeviceState)(THIS_ HWND dst_window) PURE;
    STDMETHOD(CreateRenderTargetEx)(THIS_ UINT Width, UINT Height, D3DFORMAT Format, D3DMULTISAMPLE_TYPE MultiSample, DWORD MultiSampleQuality, WINBOOL Lockable, IDirect3DSurface9 ** ppSurface, HANDLE *pSharedHandle, DWORD Usage) PURE;
    STDMETHOD(CreateOffscreenPlainSurfaceEx)(THIS_ UINT Width, UINT Height, D3DFORMAT Format, D3DPOOL Pool, IDirect3DSurface9 **ppSurface, HANDLE *pSharedHandle, DWORD Usage) PURE;
    STDMETHOD(CreateDepthStencilSurfaceEx)(THIS_ UINT width, UINT height, D3DFORMAT format,
            D3DMULTISAMPLE_TYPE multisample_type, DWORD multisample_quality, WINBOOL discard,
            IDirect3DSurface9 **surface, HANDLE *shared_handle, DWORD usage) PURE;
    STDMETHOD(ResetEx)(THIS_ D3DPRESENT_PARAMETERS *pPresentationParameters, D3DDISPLAYMODEEX *pFullscreenDisplayMode) PURE;
    STDMETHOD(GetDisplayModeEx)(THIS_ UINT iSwapChain, D3DDISPLAYMODEEX *pMode, D3DDISPLAYROTATION *pRotation) PURE;
#endif /* D3D_DISABLE_9EX */
};
#undef INTERFACE

#if !defined(__cplusplus) || defined(CINTERFACE)
/*** IUnknown methods ***/
#define IDirect3DDevice9Ex_QueryInterface(p,a,b) (p)->lpVtbl->QueryInterface(p,a,b)
#define IDirect3DDevice9Ex_AddRef(p)             (p)->lpVtbl->AddRef(p)
#define IDirect3DDevice9Ex_Release(p)            (p)->lpVtbl->Release(p)
/*** IDirect3DDevice9 methods ***/
#define IDirect3DDevice9Ex_TestCooperativeLevel(p)                       (p)->lpVtbl->TestCooperativeLevel(p)
#define IDirect3DDevice9Ex_GetAvailableTextureMem(p)                     (p)->lpVtbl->GetAvailableTextureMem(p)
#define IDirect3DDevice9Ex_EvictManagedResources(p)                      (p)->lpVtbl->EvictManagedResources(p)
#define IDirect3DDevice9Ex_GetDirect3D(p,a)                              (p)->lpVtbl->GetDirect3D(p,a)
#define IDirect3DDevice9Ex_GetDeviceCaps(p,a)                            (p)->lpVtbl->GetDeviceCaps(p,a)
#define IDirect3DDevice9Ex_GetDisplayMode(p,a,b)                         (p)->lpVtbl->GetDisplayMode(p,a,b)
#define IDirect3DDevice9Ex_GetCreationParameters(p,a)                    (p)->lpVtbl->GetCreationParameters(p,a)
#define IDirect3DDevice9Ex_SetCursorProperties(p,a,b,c)                  (p)->lpVtbl->SetCursorProperties(p,a,b,c)
#define IDirect3DDevice9Ex_SetCursorPosition(p,a,b,c)                    (p)->lpVtbl->SetCursorPosition(p,a,b,c)
#define IDirect3DDevice9Ex_ShowCursor(p,a)                               (p)->lpVtbl->ShowCursor(p,a)
#define IDirect3DDevice9Ex_CreateAdditionalSwapChain(p,a,b)              (p)->lpVtbl->CreateAdditionalSwapChain(p,a,b)
#define IDirect3DDevice9Ex_GetSwapChain(p,a,b)                           (p)->lpVtbl->GetSwapChain(p,a,b)
#define IDirect3DDevice9Ex_GetNumberOfSwapChains(p)                      (p)->lpVtbl->GetNumberOfSwapChains(p)
#define IDirect3DDevice9Ex_Reset(p,a)                                    (p)->lpVtbl->Reset(p,a)
#define IDirect3DDevice9Ex_Present(p,a,b,c,d)                            (p)->lpVtbl->Present(p,a,b,c,d)
#define IDirect3DDevice9Ex_GetBackBuffer(p,a,b,c,d)                      (p)->lpVtbl->GetBackBuffer(p,a,b,c,d)
#define IDirect3DDevice9Ex_GetRasterStatus(p,a,b)                        (p)->lpVtbl->GetRasterStatus(p,a,b)
#define IDirect3DDevice9Ex_SetDialogBoxMode(p,a)                         (p)->lpVtbl->SetDialogBoxMode(p,a)
#define IDirect3DDevice9Ex_SetGammaRamp(p,a,b,c)                         (p)->lpVtbl->SetGammaRamp(p,a,b,c)
#define IDirect3DDevice9Ex_GetGammaRamp(p,a,b)                           (p)->lpVtbl->GetGammaRamp(p,a,b)
#define IDirect3DDevice9Ex_CreateTexture(p,a,b,c,d,e,f,g,h)              (p)->lpVtbl->CreateTexture(p,a,b,c,d,e,f,g,h)
#define IDirect3DDevice9Ex_CreateVolumeTexture(p,a,b,c,d,e,f,g,h,i)      (p)->lpVtbl->CreateVolumeTexture(p,a,b,c,d,e,f,g,h,i)
#define IDirect3DDevice9Ex_CreateCubeTexture(p,a,b,c,d,e,f,g)            (p)->lpVtbl->CreateCubeTexture(p,a,b,c,d,e,f,g)
#define IDirect3DDevice9Ex_CreateVertexBuffer(p,a,b,c,d,e,f)             (p)->lpVtbl->CreateVertexBuffer(p,a,b,c,d,e,f)
#define IDirect3DDevice9Ex_CreateIndexBuffer(p,a,b,c,d,e,f)              (p)->lpVtbl->CreateIndexBuffer(p,a,b,c,d,e,f)
#define IDirect3DDevice9Ex_CreateRenderTarget(p,a,b,c,d,e,f,g,h)         (p)->lpVtbl->CreateRenderTarget(p,a,b,c,d,e,f,g,h)
#define IDirect3DDevice9Ex_CreateDepthStencilSurface(p,a,b,c,d,e,f,g,h)  (p)->lpVtbl->CreateDepthStencilSurface(p,a,b,c,d,e,f,g,h)
#define IDirect3DDevice9Ex_UpdateSurface(p,a,b,c,d)                      (p)->lpVtbl->UpdateSurface(p,a,b,c,d)
#define IDirect3DDevice9Ex_UpdateTexture(p,a,b)                          (p)->lpVtbl->UpdateTexture(p,a,b)
#define IDirect3DDevice9Ex_GetRenderTargetData(p,a,b)                    (p)->lpVtbl->GetRenderTargetData(p,a,b)
#define IDirect3DDevice9Ex_GetFrontBufferData(p,a,b)                     (p)->lpVtbl->GetFrontBufferData(p,a,b)
#define IDirect3DDevice9Ex_StretchRect(p,a,b,c,d,e)                      (p)->lpVtbl->StretchRect(p,a,b,c,d,e)
#define IDirect3DDevice9Ex_ColorFill(p,a,b,c)                            (p)->lpVtbl->ColorFill(p,a,b,c)
#define IDirect3DDevice9Ex_CreateOffscreenPlainSurface(p,a,b,c,d,e,f)    (p)->lpVtbl->CreateOffscreenPlainSurface(p,a,b,c,d,e,f)
#define IDirect3DDevice9Ex_SetRenderTarget(p,a,b)                        (p)->lpVtbl->SetRenderTarget(p,a,b)
#define IDirect3DDevice9Ex_GetRenderTarget(p,a,b)                        (p)->lpVtbl->GetRenderTarget(p,a,b)
#define IDirect3DDevice9Ex_SetDepthStencilSurface(p,a)                   (p)->lpVtbl->SetDepthStencilSurface(p,a)
#define IDirect3DDevice9Ex_GetDepthStencilSurface(p,a)                   (p)->lpVtbl->GetDepthStencilSurface(p,a)
#define IDirect3DDevice9Ex_BeginScene(p)                                 (p)->lpVtbl->BeginScene(p)
#define IDirect3DDevice9Ex_EndScene(p)                                   (p)->lpVtbl->EndScene(p)
#define IDirect3DDevice9Ex_Clear(p,a,b,c,d,e,f)                          (p)->lpVtbl->Clear(p,a,b,c,d,e,f)
#define IDirect3DDevice9Ex_SetTransform(p,a,b)                           (p)->lpVtbl->SetTransform(p,a,b)
#define IDirect3DDevice9Ex_GetTransform(p,a,b)                           (p)->lpVtbl->GetTransform(p,a,b)
#define IDirect3DDevice9Ex_MultiplyTransform(p,a,b)                      (p)->lpVtbl->MultiplyTransform(p,a,b)
#define IDirect3DDevice9Ex_SetViewport(p,a)                              (p)->lpVtbl->SetViewport(p,a)
#define IDirect3DDevice9Ex_GetViewport(p,a)                              (p)->lpVtbl->GetViewport(p,a)
#define IDirect3DDevice9Ex_SetMaterial(p,a)                              (p)->lpVtbl->SetMaterial(p,a)
#define IDirect3DDevice9Ex_GetMaterial(p,a)                              (p)->lpVtbl->GetMaterial(p,a)
#define IDirect3DDevice9Ex_SetLight(p,a,b)                               (p)->lpVtbl->SetLight(p,a,b)
#define IDirect3DDevice9Ex_GetLight(p,a,b)                               (p)->lpVtbl->GetLight(p,a,b)
#define IDirect3DDevice9Ex_LightEnable(p,a,b)                            (p)->lpVtbl->LightEnable(p,a,b)
#define IDirect3DDevice9Ex_GetLightEnable(p,a,b)                         (p)->lpVtbl->GetLightEnable(p,a,b)
#define IDirect3DDevice9Ex_SetClipPlane(p,a,b)                           (p)->lpVtbl->SetClipPlane(p,a,b)
#define IDirect3DDevice9Ex_GetClipPlane(p,a,b)                           (p)->lpVtbl->GetClipPlane(p,a,b)
#define IDirect3DDevice9Ex_SetRenderState(p,a,b)                         (p)->lpVtbl->SetRenderState(p,a,b)
#define IDirect3DDevice9Ex_GetRenderState(p,a,b)                         (p)->lpVtbl->GetRenderState(p,a,b)
#define IDirect3DDevice9Ex_CreateStateBlock(p,a,b)                       (p)->lpVtbl->CreateStateBlock(p,a,b)
#define IDirect3DDevice9Ex_BeginStateBlock(p)                            (p)->lpVtbl->BeginStateBlock(p)
#define IDirect3DDevice9Ex_EndStateBlock(p,a)                            (p)->lpVtbl->EndStateBlock(p,a)
#define IDirect3DDevice9Ex_SetClipStatus(p,a)                            (p)->lpVtbl->SetClipStatus(p,a)
#define IDirect3DDevice9Ex_GetClipStatus(p,a)                            (p)->lpVtbl->GetClipStatus(p,a)
#define IDirect3DDevice9Ex_GetTexture(p,a,b)                             (p)->lpVtbl->GetTexture(p,a,b)
#define IDirect3DDevice9Ex_SetTexture(p,a,b)                             (p)->lpVtbl->SetTexture(p,a,b)
#define IDirect3DDevice9Ex_GetTextureStageState(p,a,b,c)                 (p)->lpVtbl->GetTextureStageState(p,a,b,c)
#define IDirect3DDevice9Ex_SetTextureStageState(p,a,b,c)                 (p)->lpVtbl->SetTextureStageState(p,a,b,c)
#define IDirect3DDevice9Ex_GetSamplerState(p,a,b,c)                      (p)->lpVtbl->GetSamplerState(p,a,b,c)
#define IDirect3DDevice9Ex_SetSamplerState(p,a,b,c)                      (p)->lpVtbl->SetSamplerState(p,a,b,c)
#define IDirect3DDevice9Ex_ValidateDevice(p,a)                           (p)->lpVtbl->ValidateDevice(p,a)
#define IDirect3DDevice9Ex_SetPaletteEntries(p,a,b)                      (p)->lpVtbl->SetPaletteEntries(p,a,b)
#define IDirect3DDevice9Ex_GetPaletteEntries(p,a,b)                      (p)->lpVtbl->GetPaletteEntries(p,a,b)
#define IDirect3DDevice9Ex_SetCurrentTexturePalette(p,a)                 (p)->lpVtbl->SetCurrentTexturePalette(p,a)
#define IDirect3DDevice9Ex_GetCurrentTexturePalette(p,a)                 (p)->lpVtbl->GetCurrentTexturePalette(p,a)
#define IDirect3DDevice9Ex_SetScissorRect(p,a)                           (p)->lpVtbl->SetScissorRect(p,a)
#define IDirect3DDevice9Ex_GetScissorRect(p,a)                           (p)->lpVtbl->GetScissorRect(p,a)
#define IDirect3DDevice9Ex_SetSoftwareVertexProcessing(p,a)              (p)->lpVtbl->SetSoftwareVertexProcessing(p,a)
#define IDirect3DDevice9Ex_GetSoftwareVertexProcessing(p)                (p)->lpVtbl->GetSoftwareVertexProcessing(p)
#define IDirect3DDevice9Ex_SetNPatchMode(p,a)                            (p)->lpVtbl->SetNPatchMode(p,a)
#define IDirect3DDevice9Ex_GetNPatchMode(p)                              (p)->lpVtbl->GetNPatchMode(p)
#define IDirect3DDevice9Ex_DrawPrimitive(p,a,b,c)                        (p)->lpVtbl->DrawPrimitive(p,a,b,c)
#define IDirect3DDevice9Ex_DrawIndexedPrimitive(p,a,b,c,d,e,f)           (p)->lpVtbl->DrawIndexedPrimitive(p,a,b,c,d,e,f)
#define IDirect3DDevice9Ex_DrawPrimitiveUP(p,a,b,c,d)                    (p)->lpVtbl->DrawPrimitiveUP(p,a,b,c,d)
#define IDirect3DDevice9Ex_DrawIndexedPrimitiveUP(p,a,b,c,d,e,f,g,h)     (p)->lpVtbl->DrawIndexedPrimitiveUP(p,a,b,c,d,e,f,g,h)
#define IDirect3DDevice9Ex_ProcessVertices(p,a,b,c,d,e,f)                (p)->lpVtbl->ProcessVertices(p,a,b,c,d,e,f)
#define IDirect3DDevice9Ex_CreateVertexDeclaration(p,a,b)                (p)->lpVtbl->CreateVertexDeclaration(p,a,b)
#define IDirect3DDevice9Ex_SetVertexDeclaration(p,a)                     (p)->lpVtbl->SetVertexDeclaration(p,a)
#define IDirect3DDevice9Ex_GetVertexDeclaration(p,a)                     (p)->lpVtbl->GetVertexDeclaration(p,a)
#define IDirect3DDevice9Ex_SetFVF(p,a)                                   (p)->lpVtbl->SetFVF(p,a)
#define IDirect3DDevice9Ex_GetFVF(p,a)                                   (p)->lpVtbl->GetFVF(p,a)
#define IDirect3DDevice9Ex_CreateVertexShader(p,a,b)                     (p)->lpVtbl->CreateVertexShader(p,a,b)
#define IDirect3DDevice9Ex_SetVertexShader(p,a)                          (p)->lpVtbl->SetVertexShader(p,a)
#define IDirect3DDevice9Ex_GetVertexShader(p,a)                          (p)->lpVtbl->GetVertexShader(p,a)
#define IDirect3DDevice9Ex_SetVertexShaderConstantF(p,a,b,c)             (p)->lpVtbl->SetVertexShaderConstantF(p,a,b,c)
#define IDirect3DDevice9Ex_GetVertexShaderConstantF(p,a,b,c)             (p)->lpVtbl->GetVertexShaderConstantF(p,a,b,c)
#define IDirect3DDevice9Ex_SetVertexShaderConstantI(p,a,b,c)             (p)->lpVtbl->SetVertexShaderConstantI(p,a,b,c)
#define IDirect3DDevice9Ex_GetVertexShaderConstantI(p,a,b,c)             (p)->lpVtbl->GetVertexShaderConstantI(p,a,b,c)
#define IDirect3DDevice9Ex_SetVertexShaderConstantB(p,a,b,c)             (p)->lpVtbl->SetVertexShaderConstantB(p,a,b,c)
#define IDirect3DDevice9Ex_GetVertexShaderConstantB(p,a,b,c)             (p)->lpVtbl->GetVertexShaderConstantB(p,a,b,c)
#define IDirect3DDevice9Ex_SetStreamSource(p,a,b,c,d)                    (p)->lpVtbl->SetStreamSource(p,a,b,c,d)
#define IDirect3DDevice9Ex_GetStreamSource(p,a,b,c,d)                    (p)->lpVtbl->GetStreamSource(p,a,b,c,d)
#define IDirect3DDevice9Ex_SetStreamSourceFreq(p,a,b)                    (p)->lpVtbl->SetStreamSourceFreq(p,a,b)
#define IDirect3DDevice9Ex_GetStreamSourceFreq(p,a,b)                    (p)->lpVtbl->GetStreamSourceFreq(p,a,b)
#define IDirect3DDevice9Ex_SetIndices(p,a)                               (p)->lpVtbl->SetIndices(p,a)
#define IDirect3DDevice9Ex_GetIndices(p,a)                               (p)->lpVtbl->GetIndices(p,a)
#define IDirect3DDevice9Ex_CreatePixelShader(p,a,b)                      (p)->lpVtbl->CreatePixelShader(p,a,b)
#define IDirect3DDevice9Ex_SetPixelShader(p,a)                           (p)->lpVtbl->SetPixelShader(p,a)
#define IDirect3DDevice9Ex_GetPixelShader(p,a)                           (p)->lpVtbl->GetPixelShader(p,a)
#define IDirect3DDevice9Ex_SetPixelShaderConstantF(p,a,b,c)              (p)->lpVtbl->SetPixelShaderConstantF(p,a,b,c)
#define IDirect3DDevice9Ex_GetPixelShaderConstantF(p,a,b,c)              (p)->lpVtbl->GetPixelShaderConstantF(p,a,b,c)
#define IDirect3DDevice9Ex_SetPixelShaderConstantI(p,a,b,c)              (p)->lpVtbl->SetPixelShaderConstantI(p,a,b,c)
#define IDirect3DDevice9Ex_GetPixelShaderConstantI(p,a,b,c)              (p)->lpVtbl->GetPixelShaderConstantI(p,a,b,c)
#define IDirect3DDevice9Ex_SetPixelShaderConstantB(p,a,b,c)              (p)->lpVtbl->SetPixelShaderConstantB(p,a,b,c)
#define IDirect3DDevice9Ex_GetPixelShaderConstantB(p,a,b,c)              (p)->lpVtbl->GetPixelShaderConstantB(p,a,b,c)
#define IDirect3DDevice9Ex_DrawRectPatch(p,a,b,c)                        (p)->lpVtbl->DrawRectPatch(p,a,b,c)
#define IDirect3DDevice9Ex_DrawTriPatch(p,a,b,c)                         (p)->lpVtbl->DrawTriPatch(p,a,b,c)
#define IDirect3DDevice9Ex_DeletePatch(p,a)                              (p)->lpVtbl->DeletePatch(p,a)
#define IDirect3DDevice9Ex_CreateQuery(p,a,b)                            (p)->lpVtbl->CreateQuery(p,a,b)
/* IDirect3DDevice9Ex */
#define IDirect3DDevice9Ex_SetConvolutionMonoKernel(p,a,b,c,d)           (p)->lpVtbl->SetConvolutionMonoKernel(p,a,b,c,d)
#define IDirect3DDevice9Ex_ComposeRects(p,a,b,c,d,e,f,g,h)               (p)->lpVtbl->ComposeRects(p,a,b,c,d,e,f,g,h)
#define IDirect3DDevice9Ex_PresentEx(p,a,b,c,d,e)                        (p)->lpVtbl->PresentEx(p,a,b,c,d,e)
#define IDirect3DDevice9Ex_GetGPUThreadPriority(p,a)                     (p)->lpVtbl->GetGPUThreadPriority(p,a)
#define IDirect3DDevice9Ex_SetGPUThreadPriority(p,a)                     (p)->lpVtbl->SetGPUThreadPriority(p,a)
#define IDirect3DDevice9Ex_WaitForVBlank(p,a)                            (p)->lpVtbl->WaitForVBlank(p,a)
#define IDirect3DDevice9Ex_CheckResourceResidency(p,a,b)                 (p)->lpVtbl->CheckResourceResidency(p,a,b)
#define IDirect3DDevice9Ex_SetMaximumFrameLatency(p,a)                   (p)->lpVtbl->SetMaximumFrameLatency(p,a)
#define IDirect3DDevice9Ex_GetMaximumFrameLatency(p,a)                   (p)->lpVtbl->GetMaximumFrameLatency(p,a)
#define IDirect3DDevice9Ex_CheckDeviceState(p,a)                         (p)->lpVtbl->CheckDeviceState(p,a)
#define IDirect3DDevice9Ex_CreateRenderTargetEx(p,a,b,c,d,e,f,g,h,i)     (p)->lpVtbl->CreateRenderTargetEx(p,a,b,c,d,e,f,g,h,i)
#define IDirect3DDevice9Ex_CreateOffscreenPlainSurfaceEx(p,a,b,c,d,e,f,g)(p)->lpVtbl->CreateOffscreenPlainSurfaceEx(p,a,b,c,d,e,f,g)
#define IDirect3DDevice9Ex_CreateDepthStencilSurfaceEx(p,a,b,c,d,e,f,g,h,i)(p)->lpVtbl->CreateDepthStencilSurfaceEx(p,a,b,c,d,e,f,g,h,i)
#define IDirect3DDevice9Ex_ResetEx(p,a,b)                                 (p)->lpVtbl->ResetEx(p,a,b)
#define IDirect3DDevice9Ex_GetDisplayModeEx(p,a,b,c)                     (p)->lpVtbl->GetDisplayModeEx(p,a,b,c)
#else
/*** IUnknown methods ***/
#define IDirect3DDevice9Ex_QueryInterface(p,a,b) (p)->QueryInterface(a,b)
#define IDirect3DDevice9Ex_AddRef(p)             (p)->AddRef()
#define IDirect3DDevice9Ex_Release(p)            (p)->Release()
/*** IDirect3DDevice9 methods ***/
#define IDirect3DDevice9Ex_TestCooperativeLevel(p)                       (p)->TestCooperativeLevel()
#define IDirect3DDevice9Ex_GetAvailableTextureMem(p)                     (p)->GetAvailableTextureMem()
#define IDirect3DDevice9Ex_EvictManagedResources(p)                      (p)->EvictManagedResources()
#define IDirect3DDevice9Ex_GetDirect3D(p,a)                              (p)->GetDirect3D(a)
#define IDirect3DDevice9Ex_GetDeviceCaps(p,a)                            (p)->GetDeviceCaps(a)
#define IDirect3DDevice9Ex_GetDisplayMode(p,a,b)                         (p)->GetDisplayMode(a,b)
#define IDirect3DDevice9Ex_GetCreationParameters(p,a)                    (p)->GetCreationParameters(a)
#define IDirect3DDevice9Ex_SetCursorProperties(p,a,b,c)                  (p)->SetCursorProperties(a,b,c)
#define IDirect3DDevice9Ex_SetCursorPosition(p,a,b,c)                    (p)->SetCursorPosition(a,b,c)
#define IDirect3DDevice9Ex_ShowCursor(p,a)                               (p)->ShowCursor(a)
#define IDirect3DDevice9Ex_CreateAdditionalSwapChain(p,a,b)              (p)->CreateAdditionalSwapChain(a,b)
#define IDirect3DDevice9Ex_GetSwapChain(p,a,b)                           (p)->GetSwapChain(a,b)
#define IDirect3DDevice9Ex_GetNumberOfSwapChains(p)                      (p)->GetNumberOfSwapChains()
#define IDirect3DDevice9Ex_Reset(p,a)                                    (p)->Reset(a)
#define IDirect3DDevice9Ex_Present(p,a,b,c,d)                            (p)->Present(a,b,c,d)
#define IDirect3DDevice9Ex_GetBackBuffer(p,a,b,c,d)                      (p)->GetBackBuffer(a,b,c,d)
#define IDirect3DDevice9Ex_GetRasterStatus(p,a,b)                        (p)->GetRasterStatus(a,b)
#define IDirect3DDevice9Ex_SetDialogBoxMode(p,a)                         (p)->SetDialogBoxMode(a)
#define IDirect3DDevice9Ex_SetGammaRamp(p,a,b,c)                         (p)->SetGammaRamp(a,b,c)
#define IDirect3DDevice9Ex_GetGammaRamp(p,a,b)                           (p)->GetGammaRamp(a,b)
#define IDirect3DDevice9Ex_CreateTexture(p,a,b,c,d,e,f,g,h)              (p)->CreateTexture(a,b,c,d,e,f,g,h)
#define IDirect3DDevice9Ex_CreateVolumeTexture(p,a,b,c,d,e,f,g,h,i)      (p)->CreateVolumeTexture(a,b,c,d,e,f,g,h,i)
#define IDirect3DDevice9Ex_CreateCubeTexture(p,a,b,c,d,e,f,g)            (p)->CreateCubeTexture(a,b,c,d,e,f,g)
#define IDirect3DDevice9Ex_CreateVertexBuffer(p,a,b,c,d,e,f)             (p)->CreateVertexBuffer(a,b,c,d,e,f)
#define IDirect3DDevice9Ex_CreateIndexBuffer(p,a,b,c,d,e,f)              (p)->CreateIndexBuffer(a,b,c,d,e,f)
#define IDirect3DDevice9Ex_CreateRenderTarget(p,a,b,c,d,e,f,g,h)         (p)->CreateRenderTarget(a,b,c,d,e,f,g,h)
#define IDirect3DDevice9Ex_CreateDepthStencilSurface(p,a,b,c,d,e,f,g,h)  (p)->CreateDepthStencilSurface(a,b,c,d,e,f,g,h)
#define IDirect3DDevice9Ex_UpdateSurface(p,a,b,c,d)                      (p)->UpdateSurface(a,b,c,d)
#define IDirect3DDevice9Ex_UpdateTexture(p,a,b)                          (p)->UpdateTexture(a,b)
#define IDirect3DDevice9Ex_GetRenderTargetData(p,a,b)                    (p)->GetRenderTargetData(a,b)
#define IDirect3DDevice9Ex_GetFrontBufferData(p,a,b)                     (p)->GetFrontBufferData(a,b)
#define IDirect3DDevice9Ex_StretchRect(p,a,b,c,d,e)                      (p)->StretchRect(a,b,c,d,e)
#define IDirect3DDevice9Ex_ColorFill(p,a,b,c)                            (p)->ColorFill(a,b,c)
#define IDirect3DDevice9Ex_CreateOffscreenPlainSurface(p,a,b,c,d,e,f)    (p)->CreateOffscreenPlainSurface(a,b,c,d,e,f)
#define IDirect3DDevice9Ex_SetRenderTarget(p,a,b)                        (p)->SetRenderTarget(a,b)
#define IDirect3DDevice9Ex_GetRenderTarget(p,a,b)                        (p)->GetRenderTarget(a,b)
#define IDirect3DDevice9Ex_SetDepthStencilSurface(p,a)                   (p)->SetDepthStencilSurface(a)
#define IDirect3DDevice9Ex_GetDepthStencilSurface(p,a)                   (p)->GetDepthStencilSurface(a)
#define IDirect3DDevice9Ex_BeginScene(p)                                 (p)->BeginScene()
#define IDirect3DDevice9Ex_EndScene(p)                                   (p)->EndScene()
#define IDirect3DDevice9Ex_Clear(p,a,b,c,d,e,f)                          (p)->Clear(a,b,c,d,e,f)
#define IDirect3DDevice9Ex_SetTransform(p,a,b)                           (p)->SetTransform(a,b)
#define IDirect3DDevice9Ex_GetTransform(p,a,b)                           (p)->GetTransform(a,b)
#define IDirect3DDevice9Ex_MultiplyTransform(p,a,b)                      (p)->MultiplyTransform(a,b)
#define IDirect3DDevice9Ex_SetViewport(p,a)                              (p)->SetViewport(a)
#define IDirect3DDevice9Ex_GetViewport(p,a)                              (p)->GetViewport(a)
#define IDirect3DDevice9Ex_SetMaterial(p,a)                              (p)->SetMaterial(a)
#define IDirect3DDevice9Ex_GetMaterial(p,a)                              (p)->GetMaterial(a)
#define IDirect3DDevice9Ex_SetLight(p,a,b)                               (p)->SetLight(a,b)
#define IDirect3DDevice9Ex_GetLight(p,a,b)                               (p)->GetLight(a,b)
#define IDirect3DDevice9Ex_LightEnable(p,a,b)                            (p)->LightEnable(a,b)
#define IDirect3DDevice9Ex_GetLightEnable(p,a,b)                         (p)->GetLightEnable(a,b)
#define IDirect3DDevice9Ex_SetClipPlane(p,a,b)                           (p)->SetClipPlane(a,b)
#define IDirect3DDevice9Ex_GetClipPlane(p,a,b)                           (p)->GetClipPlane(a,b)
#define IDirect3DDevice9Ex_SetRenderState(p,a,b)                         (p)->SetRenderState(a,b)
#define IDirect3DDevice9Ex_GetRenderState(p,a,b)                         (p)->GetRenderState(a,b)
#define IDirect3DDevice9Ex_CreateStateBlock(p,a,b)                       (p)->CreateStateBlock(a,b)
#define IDirect3DDevice9Ex_BeginStateBlock(p)                            (p)->BeginStateBlock()
#define IDirect3DDevice9Ex_EndStateBlock(p,a)                            (p)->EndStateBlock(a)
#define IDirect3DDevice9Ex_SetClipStatus(p,a)                            (p)->SetClipStatus(a)
#define IDirect3DDevice9Ex_GetClipStatus(p,a)                            (p)->GetClipStatus(a)
#define IDirect3DDevice9Ex_GetTexture(p,a,b)                             (p)->GetTexture(a,b)
#define IDirect3DDevice9Ex_SetTexture(p,a,b)                             (p)->SetTexture(a,b)
#define IDirect3DDevice9Ex_GetTextureStageState(p,a,b,c)                 (p)->GetTextureStageState(a,b,c)
#define IDirect3DDevice9Ex_SetTextureStageState(p,a,b,c)                 (p)->SetTextureStageState(a,b,c)
#define IDirect3DDevice9Ex_GetSamplerState(p,a,b,c)                      (p)->GetSamplerState(a,b,c)
#define IDirect3DDevice9Ex_SetSamplerState(p,a,b,c)                      (p)->SetSamplerState(a,b,c)
#define IDirect3DDevice9Ex_ValidateDevice(p,a)                           (p)->ValidateDevice(a)
#define IDirect3DDevice9Ex_SetPaletteEntries(p,a,b)                      (p)->SetPaletteEntries(a,b)
#define IDirect3DDevice9Ex_GetPaletteEntries(p,a,b)                      (p)->GetPaletteEntries(a,b)
#define IDirect3DDevice9Ex_SetCurrentTexturePalette(p,a)                 (p)->SetCurrentTexturePalette(a)
#define IDirect3DDevice9Ex_GetCurrentTexturePalette(p,a)                 (p)->GetCurrentTexturePalette(a)
#define IDirect3DDevice9Ex_SetScissorRect(p,a)                           (p)->SetScissorRect(a)
#define IDirect3DDevice9Ex_GetScissorRect(p,a)                           (p)->GetScissorRect(a)
#define IDirect3DDevice9Ex_SetSoftwareVertexProcessing(p,a)              (p)->SetSoftwareVertexProcessing(a)
#define IDirect3DDevice9Ex_GetSoftwareVertexProcessing(p)                (p)->GetSoftwareVertexProcessing()
#define IDirect3DDevice9Ex_SetNPatchMode(p,a)                            (p)->SetNPatchMode(a)
#define IDirect3DDevice9Ex_GetNPatchMode(p)                              (p)->GetNPatchMode()
#define IDirect3DDevice9Ex_DrawPrimitive(p,a,b,c)                        (p)->DrawPrimitive(a,b,c)
#define IDirect3DDevice9Ex_DrawIndexedPrimitive(p,a,b,c,d,e,f)           (p)->DrawIndexedPrimitive(a,b,c,d,e,f)
#define IDirect3DDevice9Ex_DrawPrimitiveUP(p,a,b,c,d)                    (p)->DrawPrimitiveUP(a,b,c,d)
#define IDirect3DDevice9Ex_DrawIndexedPrimitiveUP(p,a,b,c,d,e,f,g,h)     (p)->DrawIndexedPrimitiveUP(a,b,c,d,e,f,g,h)
#define IDirect3DDevice9Ex_ProcessVertices(p,a,b,c,d,e,f)                (p)->ProcessVertices(a,b,c,d,e,f)
#define IDirect3DDevice9Ex_CreateVertexDeclaration(p,a,b)                (p)->CreateVertexDeclaration(a,b)
#define IDirect3DDevice9Ex_SetVertexDeclaration(p,a)                     (p)->SetVertexDeclaration(a)
#define IDirect3DDevice9Ex_GetVertexDeclaration(p,a)                     (p)->GetVertexDeclaration(a)
#define IDirect3DDevice9Ex_SetFVF(p,a)                                   (p)->SetFVF(a)
#define IDirect3DDevice9Ex_GetFVF(p,a)                                   (p)->GetFVF(a)
#define IDirect3DDevice9Ex_CreateVertexShader(p,a,b)                     (p)->CreateVertexShader(a,b)
#define IDirect3DDevice9Ex_SetVertexShader(p,a)                          (p)->SetVertexShader(a)
#define IDirect3DDevice9Ex_GetVertexShader(p,a)                          (p)->GetVertexShader(a)
#define IDirect3DDevice9Ex_SetVertexShaderConstantF(p,a,b,c)             (p)->SetVertexShaderConstantF(a,b,c)
#define IDirect3DDevice9Ex_GetVertexShaderConstantF(p,a,b,c)             (p)->GetVertexShaderConstantF(a,b,c)
#define IDirect3DDevice9Ex_SetVertexShaderConstantI(p,a,b,c)             (p)->SetVertexShaderConstantI(a,b,c)
#define IDirect3DDevice9Ex_GetVertexShaderConstantI(p,a,b,c)             (p)->GetVertexShaderConstantI(a,b,c)
#define IDirect3DDevice9Ex_SetVertexShaderConstantB(p,a,b,c)             (p)->SetVertexShaderConstantB(a,b,c)
#define IDirect3DDevice9Ex_GetVertexShaderConstantB(p,a,b,c)             (p)->GetVertexShaderConstantB(a,b,c)
#define IDirect3DDevice9Ex_SetStreamSource(p,a,b,c,d)                    (p)->SetStreamSource(a,b,c,d)
#define IDirect3DDevice9Ex_GetStreamSource(p,a,b,c,d)                    (p)->GetStreamSource(a,b,c,d)
#define IDirect3DDevice9Ex_SetStreamSourceFreq(p,a,b)                    (p)->SetStreamSourceFreq(a,b)
#define IDirect3DDevice9Ex_GetStreamSourceFreq(p,a,b)                    (p)->GetStreamSourceFreq(a,b)
#define IDirect3DDevice9Ex_SetIndices(p,a)                               (p)->SetIndices(a)
#define IDirect3DDevice9Ex_GetIndices(p,a)                               (p)->GetIndices(a)
#define IDirect3DDevice9Ex_CreatePixelShader(p,a,b)                      (p)->CreatePixelShader(a,b)
#define IDirect3DDevice9Ex_SetPixelShader(p,a)                           (p)->SetPixelShader(a)
#define IDirect3DDevice9Ex_GetPixelShader(p,a)                           (p)->GetPixelShader(a)
#define IDirect3DDevice9Ex_SetPixelShaderConstantF(p,a,b,c)              (p)->SetPixelShaderConstantF(a,b,c)
#define IDirect3DDevice9Ex_GetPixelShaderConstantF(p,a,b,c)              (p)->GetPixelShaderConstantF(a,b,c)
#define IDirect3DDevice9Ex_SetPixelShaderConstantI(p,a,b,c)              (p)->SetPixelShaderConstantI(a,b,c)
#define IDirect3DDevice9Ex_GetPixelShaderConstantI(p,a,b,c)              (p)->GetPixelShaderConstantI(a,b,c)
#define IDirect3DDevice9Ex_SetPixelShaderConstantB(p,a,b,c)              (p)->SetPixelShaderConstantB(a,b,c)
#define IDirect3DDevice9Ex_GetPixelShaderConstantB(p,a,b,c)              (p)->GetPixelShaderConstantB(a,b,c)
#define IDirect3DDevice9Ex_DrawRectPatch(p,a,b,c)                        (p)->DrawRectPatch(a,b,c)
#define IDirect3DDevice9Ex_DrawTriPatch(p,a,b,c)                         (p)->DrawTriPatch(a,b,c)
#define IDirect3DDevice9Ex_DeletePatch(p,a)                              (p)->DeletePatch(a)
#define IDirect3DDevice9Ex_CreateQuery(p,a,b)                            (p)->CreateQuery(a,b)
/* IDirect3DDevice9Ex */
#define IDirect3DDevice9Ex_SetConvolutionMonoKernel(p,a,b,c,d)           (p)->SetConvolutionMonoKernel(a,b,c,d)
#define IDirect3DDevice9Ex_ComposeRects(p,a,b,c,d,e,f,g,h)               (p)->ComposeRects(a,b,c,d,e,f,g,h)
#define IDirect3DDevice9Ex_PresentEx(p,a,b,c,d,e)                        (p)->PresentEx(a,b,c,d,e)
#define IDirect3DDevice9Ex_GetGPUThreadPriority(p,a)                     (p)->GetGPUThreadPriority(a)
#define IDirect3DDevice9Ex_SetGPUThreadPriority(p,a)                     (p)->SetGPUThreadPriority(a)
#define IDirect3DDevice9Ex_WaitForVBlank(p,a)                            (p)->WaitForVBlank(a)
#define IDirect3DDevice9Ex_CheckResourceResidency(p,a,b)                 (p)->CheckResourceResidency(a,b)
#define IDirect3DDevice9Ex_SetMaximumFrameLatency(p,a)                   (p)->SetMaximumFrameLatency(a)
#define IDirect3DDevice9Ex_GetMaximumFrameLatency(p,a)                   (p)->GetMaximumFrameLatency(a)
#define IDirect3DDevice9Ex_CheckDeviceState(p,a)                         (p)->CheckDeviceState(a)
#define IDirect3DDevice9Ex_CreateRenderTargetEx(p,a,b,c,d,e,f,g,h,i)     (p)->CreateRenderTargetEx(a,b,c,d,e,f,g,h,i)
#define IDirect3DDevice9Ex_CreateOffscreenPlainSurfaceEx(p,a,b,c,d,e,f,g)(p)->CreateOffscreenPlainSurfaceEx(a,b,c,d,e,f,g)
#define IDirect3DDevice9Ex_CreateDepthStencilSurfaceEx(p,a,b,c,d,e,f,g,h,i)(p)->CreateDepthStencilSurfaceEx(a,b,c,d,e,f,g,h,i)
#define IDirect3DDevice9Ex_ResetEx(p,a,b)                                (p)->ResetEx(a,b)
#define IDirect3DDevice9Ex_GetDisplayModeEx(p,a,b,c)                     (p)->GetDisplayModeEx(a,b,c)
#endif

#ifdef __cplusplus
extern "C" {
#endif  /* defined(__cplusplus) */

int         WINAPI D3DPERF_BeginEvent(D3DCOLOR,LPCWSTR);
int         WINAPI D3DPERF_EndEvent(void);
DWORD       WINAPI D3DPERF_GetStatus(void);
WINBOOL     WINAPI D3DPERF_QueryRepeatFrame(void);
void        WINAPI D3DPERF_SetMarker(D3DCOLOR,LPCWSTR);
void        WINAPI D3DPERF_SetOptions(DWORD);
void        WINAPI D3DPERF_SetRegion(D3DCOLOR,LPCWSTR);

/* Define the main entrypoint as well */
IDirect3D9* WINAPI Direct3DCreate9(UINT SDKVersion);

#ifdef __cplusplus
} /* extern "C" */
#endif /* defined(__cplusplus) */


#endif /* _D3D9_H_ */
