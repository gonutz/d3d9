package d3d9

/*
#include "d3d9wrapper.h"

HRESULT IDirect3DDevice9BeginScene(IDirect3DDevice9* obj) {
	return obj->lpVtbl->BeginScene(obj);
}

HRESULT IDirect3DDevice9BeginStateBlock(IDirect3DDevice9* obj) {
	return obj->lpVtbl->BeginStateBlock(obj);
}

HRESULT IDirect3DDevice9Clear(
		IDirect3DDevice9* obj,
		DWORD Count,
		D3DRECT* pRects,
		DWORD Flags,
		D3DCOLOR Color,
		float Z,
		DWORD Stencil) {
	return obj->lpVtbl->Clear(obj, Count, pRects, Flags, Color, Z, Stencil);
}

HRESULT IDirect3DDevice9ColorFill(
		IDirect3DDevice9* obj,
		IDirect3DSurface9* pSurface,
		RECT* pRect,
		D3DCOLOR color) {
	return obj->lpVtbl->ColorFill(obj, pSurface, pRect, color);
}

HRESULT IDirect3DDevice9CreateAdditionalSwapChain(
		IDirect3DDevice9* obj,
		D3DPRESENT_PARAMETERS* pPresentationParameters,
		IDirect3DSwapChain9** ppSwapChain) {
	return obj->lpVtbl->CreateAdditionalSwapChain(obj, pPresentationParameters,
		ppSwapChain);
}

HRESULT IDirect3DDevice9CreateCubeTexture(
		IDirect3DDevice9* obj,
		UINT EdgeLength,
		UINT Levels,
		DWORD Usage,
		D3DFORMAT Format,
		D3DPOOL Pool,
		IDirect3DCubeTexture9** ppCubeTexture,
		HANDLE* pSharedHandle) {
	return obj->lpVtbl->CreateCubeTexture(obj, EdgeLength, Levels, Usage,
		Format, Pool, ppCubeTexture, pSharedHandle);
}

HRESULT IDirect3DDevice9CreateDepthStencilSurface(
		IDirect3DDevice9* obj,
		UINT Width,
		UINT Height,
		D3DFORMAT Format,
		D3DMULTISAMPLE_TYPE MultiSample,
		DWORD MultisampleQuality,
		BOOL Discard,
		IDirect3DSurface9** ppSurface,
		HANDLE* pSharedHandle) {
	return obj->lpVtbl->CreateDepthStencilSurface(obj, Width, Height, Format,
		MultiSample, MultisampleQuality, Discard, ppSurface, pSharedHandle);
}

HRESULT IDirect3DDevice9CreateIndexBuffer(
		IDirect3DDevice9* obj,
		UINT Length,
		DWORD Usage,
		D3DFORMAT Format,
		D3DPOOL Pool,
		IDirect3DIndexBuffer9** ppIndexBuffer,
		HANDLE* pSharedHandle) {
	return obj->lpVtbl->CreateIndexBuffer(obj, Length, Usage, Format, Pool,
		ppIndexBuffer, pSharedHandle);
}

HRESULT IDirect3DDevice9CreateOffscreenPlainSurface(
		IDirect3DDevice9* obj,
		UINT Width,
		UINT Height,
		D3DFORMAT Format,
		D3DPOOL Pool,
		IDirect3DSurface9** ppSurface,
		HANDLE* pSharedHandle) {
	return obj->lpVtbl->CreateOffscreenPlainSurface(obj, Width, Height, Format,
		Pool, ppSurface, pSharedHandle);
}

HRESULT IDirect3DDevice9CreatePixelShader(
		IDirect3DDevice9* obj,
		DWORD* pFunction,
		IDirect3DPixelShader9** ppShader) {
	return obj->lpVtbl->CreatePixelShader(obj, pFunction, ppShader);
}

HRESULT IDirect3DDevice9CreateQuery(
		IDirect3DDevice9* obj,
		D3DQUERYTYPE Type,
		IDirect3DQuery9** ppQuery) {
	return obj->lpVtbl->CreateQuery(obj, Type, ppQuery);
}

HRESULT IDirect3DDevice9CreateRenderTarget(
		IDirect3DDevice9* obj,
		UINT Width,
		UINT Height,
		D3DFORMAT Format,
		D3DMULTISAMPLE_TYPE MultiSample,
		DWORD MultisampleQuality,
		BOOL Lockable,
		IDirect3DSurface9** ppSurface,
		HANDLE* pSharedHandle) {
	return obj->lpVtbl->CreateRenderTarget(obj, Width, Height, Format,
		MultiSample, MultisampleQuality, Lockable, ppSurface, pSharedHandle);
}

HRESULT IDirect3DDevice9CreateStateBlock(
		IDirect3DDevice9* obj,
		D3DSTATEBLOCKTYPE Type,
		IDirect3DStateBlock9** ppSB) {
	return obj->lpVtbl->CreateStateBlock(obj, Type, ppSB);
}

HRESULT IDirect3DDevice9CreateTexture(
		IDirect3DDevice9* obj,
		UINT Width,
		UINT Height,
		UINT Levels,
		DWORD Usage,
		D3DFORMAT Format,
		D3DPOOL Pool,
		IDirect3DTexture9** ppTexture,
		HANDLE* pSharedHandle) {
	return obj->lpVtbl->CreateTexture(obj, Width, Height, Levels, Usage, Format,
		Pool, ppTexture, pSharedHandle);
}

HRESULT IDirect3DDevice9CreateVertexBuffer(
		IDirect3DDevice9* obj,
		UINT Length,
		DWORD Usage,
		DWORD FVF,
		D3DPOOL Pool,
		IDirect3DVertexBuffer9** ppVertexBuffer,
		HANDLE* pSharedHandle) {
	return obj->lpVtbl->CreateVertexBuffer(obj, Length, Usage, FVF, Pool,
		ppVertexBuffer, pSharedHandle);
}

HRESULT IDirect3DDevice9CreateVertexDeclaration(
		IDirect3DDevice9* obj,
		D3DVERTEXELEMENT9* pVertexElements,
		IDirect3DVertexDeclaration9** ppDecl) {
	return obj->lpVtbl->CreateVertexDeclaration(obj, pVertexElements, ppDecl);
}

HRESULT IDirect3DDevice9CreateVertexShader(
		IDirect3DDevice9* obj,
		DWORD* pFunction,
		IDirect3DVertexShader9** ppShader) {
	return obj->lpVtbl->CreateVertexShader(obj, pFunction, ppShader);
}

HRESULT IDirect3DDevice9CreateVolumeTexture(
		IDirect3DDevice9* obj,
		UINT Width,
		UINT Height,
		UINT Depth,
		UINT Levels,
		DWORD Usage,
		D3DFORMAT Format,
		D3DPOOL Pool,
		IDirect3DVolumeTexture9** ppVolumeTexture,
		HANDLE* pSharedHandle) {
	return obj->lpVtbl->CreateVolumeTexture(obj, Width, Height, Depth, Levels,
		Usage, Format, Pool, ppVolumeTexture, pSharedHandle);
}

HRESULT IDirect3DDevice9DeletePatch(IDirect3DDevice9* obj, UINT Handle) {
	return obj->lpVtbl->DeletePatch(obj, Handle);
}

HRESULT IDirect3DDevice9DrawIndexedPrimitive(
		IDirect3DDevice9* obj,
		D3DPRIMITIVETYPE Type,
		INT BaseVertexIndex,
		UINT MinIndex,
		UINT NumVertices,
		UINT StartIndex,
		UINT PrimitiveCount) {
	return obj->lpVtbl->DrawIndexedPrimitive(obj, Type, BaseVertexIndex,
		MinIndex, NumVertices, StartIndex, PrimitiveCount);
}

HRESULT IDirect3DDevice9DrawIndexedPrimitiveUP(
		IDirect3DDevice9* obj,
		D3DPRIMITIVETYPE PrimitiveType,
		UINT MinVertexIndex,
		UINT NumVertices,
		UINT PrimitiveCount,
		void* pIndexData,
		D3DFORMAT IndexDataFormat,
		void* pVertexStreamZeroData,
		UINT VertexStreamZeroStride) {
	return obj->lpVtbl->DrawIndexedPrimitiveUP(obj, PrimitiveType,
		MinVertexIndex, NumVertices, PrimitiveCount, pIndexData,
		IndexDataFormat, pVertexStreamZeroData, VertexStreamZeroStride);
}

HRESULT IDirect3DDevice9DrawPrimitive(
		IDirect3DDevice9* obj,
		D3DPRIMITIVETYPE PrimitiveType,
		UINT StartVertex,
		UINT PrimitiveCount) {
	return obj->lpVtbl->DrawPrimitive(obj, PrimitiveType, StartVertex,
		PrimitiveCount);
}

HRESULT IDirect3DDevice9DrawPrimitiveUP(
		IDirect3DDevice9* obj,
		D3DPRIMITIVETYPE PrimitiveType,
		UINT PrimitiveCount,
		void* pVertexStreamZeroData,
		UINT VertexStreamZeroStride) {
	return obj->lpVtbl->DrawPrimitiveUP(obj, PrimitiveType, PrimitiveCount,
		pVertexStreamZeroData, VertexStreamZeroStride);
}

HRESULT IDirect3DDevice9DrawRectPatch(
		IDirect3DDevice9* obj,
		UINT Handle,
		float* pNumSegs,
		D3DRECTPATCH_INFO* pRectPatchInfo) {
	return obj->lpVtbl->DrawRectPatch(obj, Handle, pNumSegs, pRectPatchInfo);
}

HRESULT IDirect3DDevice9DrawTriPatch(
		IDirect3DDevice9* obj,
		UINT Handle,
		float* pNumSegs,
		D3DTRIPATCH_INFO* pTriPatchInfo) {
	return obj->lpVtbl->DrawTriPatch(obj, Handle, pNumSegs, pTriPatchInfo);
}

HRESULT IDirect3DDevice9EndScene(IDirect3DDevice9* obj) {
	return obj->lpVtbl->EndScene(obj);
}

HRESULT IDirect3DDevice9EndStateBlock(
		IDirect3DDevice9* obj,
		IDirect3DStateBlock9** ppSB) {
	return obj->lpVtbl->EndStateBlock(obj, ppSB);
}

HRESULT IDirect3DDevice9EvictManagedResources(IDirect3DDevice9* obj) {
	return obj->lpVtbl->EvictManagedResources(obj);
}

UINT IDirect3DDevice9GetAvailableTextureMem(IDirect3DDevice9* obj) {
	return obj->lpVtbl->GetAvailableTextureMem(obj);
}

HRESULT IDirect3DDevice9GetBackBuffer(
		IDirect3DDevice9* obj,
		UINT iSwapChain,
		UINT BackBuffer,
		D3DBACKBUFFER_TYPE Type,
		IDirect3DSurface9** ppBackBuffer) {
	return obj->lpVtbl->GetBackBuffer(obj, iSwapChain, BackBuffer, Type,
		ppBackBuffer);
}

HRESULT IDirect3DDevice9GetClipPlane(
		IDirect3DDevice9* obj,
		DWORD Index,
		float* pPlane) {
	return obj->lpVtbl->GetClipPlane(obj, Index, pPlane);
}

HRESULT IDirect3DDevice9GetClipStatus(
		IDirect3DDevice9* obj,
		D3DCLIPSTATUS9* pClipStatus) {
	return obj->lpVtbl->GetClipStatus(obj, pClipStatus);
}

HRESULT IDirect3DDevice9GetCreationParameters(
		IDirect3DDevice9* obj,
		D3DDEVICE_CREATION_PARAMETERS* pParameters) {
	return obj->lpVtbl->GetCreationParameters(obj, pParameters);
}

HRESULT IDirect3DDevice9GetCurrentTexturePalette(
		IDirect3DDevice9* obj,
		UINT* pPaletteNumber) {
	return obj->lpVtbl->GetCurrentTexturePalette(obj, pPaletteNumber);
}

HRESULT IDirect3DDevice9GetDepthStencilSurface(
		IDirect3DDevice9* obj,
		IDirect3DSurface9** ppZStencilSurface) {
	return obj->lpVtbl->GetDepthStencilSurface(obj, ppZStencilSurface);
}

HRESULT IDirect3DDevice9GetDeviceCaps(IDirect3DDevice9* obj, D3DCAPS9* pCaps) {
	return obj->lpVtbl->GetDeviceCaps(obj, pCaps);
}

HRESULT IDirect3DDevice9GetDirect3D(
		IDirect3DDevice9* obj,
		IDirect3D9** ppD3D9) {
	return obj->lpVtbl->GetDirect3D(obj, ppD3D9);
}

HRESULT IDirect3DDevice9GetDisplayMode(
		IDirect3DDevice9* obj,
		UINT iSwapChain,
		D3DDISPLAYMODE* pMode) {
	return obj->lpVtbl->GetDisplayMode(obj, iSwapChain, pMode);
}

HRESULT IDirect3DDevice9GetFrontBufferData(
		IDirect3DDevice9* obj,
		UINT iSwapChain,
		IDirect3DSurface9* pDestSurface) {
	return obj->lpVtbl->GetFrontBufferData(obj, iSwapChain, pDestSurface);
}

HRESULT IDirect3DDevice9GetFVF(IDirect3DDevice9* obj, DWORD* pFVF) {
	return obj->lpVtbl->GetFVF(obj, pFVF);
}

void IDirect3DDevice9GetGammaRamp(
		IDirect3DDevice9* obj,
		UINT iSwapChain,
		D3DGAMMARAMP* pRamp) {
	obj->lpVtbl->GetGammaRamp(obj, iSwapChain, pRamp);
}

HRESULT IDirect3DDevice9GetIndices(
		IDirect3DDevice9* obj,
		IDirect3DIndexBuffer9** ppIndexData) {
	return obj->lpVtbl->GetIndices(obj, ppIndexData);
}

HRESULT IDirect3DDevice9GetLight(
		IDirect3DDevice9* obj,
		DWORD Index,
		D3DLIGHT9* pLight) {
	return obj->lpVtbl->GetLight(obj, Index, pLight);
}

HRESULT IDirect3DDevice9GetLightEnable(
		IDirect3DDevice9* obj,
		DWORD Index,
		BOOL* pEnable) {
	return obj->lpVtbl->GetLightEnable(obj, Index, pEnable);
}

HRESULT IDirect3DDevice9GetMaterial(
		IDirect3DDevice9* obj,
		D3DMATERIAL9* pMaterial) {
	return obj->lpVtbl->GetMaterial(obj, pMaterial);
}

FLOAT IDirect3DDevice9GetNPatchMode(IDirect3DDevice9* obj) {
	return obj->lpVtbl->GetNPatchMode(obj);
}

UINT IDirect3DDevice9GetNumberOfSwapChains(IDirect3DDevice9* obj) {
	return obj->lpVtbl->GetNumberOfSwapChains(obj);
}

HRESULT IDirect3DDevice9GetPaletteEntries(
		IDirect3DDevice9* obj,
		UINT PaletteNumber,
		PALETTEENTRY* pEntries) {
	return obj->lpVtbl->GetPaletteEntries(obj, PaletteNumber, pEntries);
}

HRESULT IDirect3DDevice9GetPixelShader(
		IDirect3DDevice9* obj,
		IDirect3DPixelShader9** ppShader) {
	return obj->lpVtbl->GetPixelShader(obj, ppShader);
}

HRESULT IDirect3DDevice9GetPixelShaderConstantB(
		IDirect3DDevice9* obj,
		UINT StartRegister,
		BOOL* pConstantData,
		UINT BoolCount) {
	return obj->lpVtbl->GetPixelShaderConstantB(obj, StartRegister,
		pConstantData, BoolCount);
}

HRESULT IDirect3DDevice9GetPixelShaderConstantF(
		IDirect3DDevice9* obj,
		UINT StartRegister,
		float* pConstantData,
		UINT Vector4fCount) {
	return obj->lpVtbl->GetPixelShaderConstantF(obj, StartRegister,
		pConstantData, Vector4fCount);
}

HRESULT IDirect3DDevice9GetPixelShaderConstantI(
		IDirect3DDevice9* obj,
		UINT StartRegister,
		int* pConstantData,
		UINT Vector4iCount) {
	return obj->lpVtbl->GetPixelShaderConstantI(obj, StartRegister,
		pConstantData, Vector4iCount);
}

HRESULT IDirect3DDevice9GetRasterStatus(
		IDirect3DDevice9* obj,
		UINT iSwapChain,
		D3DRASTER_STATUS* pRasterStatus) {
	return obj->lpVtbl->GetRasterStatus(obj, iSwapChain, pRasterStatus);
}

HRESULT IDirect3DDevice9GetRenderState(
		IDirect3DDevice9* obj,
		D3DRENDERSTATETYPE State,
		DWORD* pValue) {
	return obj->lpVtbl->GetRenderState(obj, State, pValue);
}

HRESULT IDirect3DDevice9GetRenderTarget(
		IDirect3DDevice9* obj,
		DWORD RenderTargetIndex,
		IDirect3DSurface9** ppRenderTarget) {
	return obj->lpVtbl->GetRenderTarget(obj, RenderTargetIndex, ppRenderTarget);
}

HRESULT IDirect3DDevice9GetRenderTargetData(
		IDirect3DDevice9* obj,
		IDirect3DSurface9* pRenderTarget,
		IDirect3DSurface9* pDestSurface) {
	return obj->lpVtbl->GetRenderTargetData(obj, pRenderTarget, pDestSurface);
}

HRESULT IDirect3DDevice9GetSamplerState(
		IDirect3DDevice9* obj,
		DWORD Sampler,
		D3DSAMPLERSTATETYPE Type,
		DWORD* pValue) {
	return obj->lpVtbl->GetSamplerState(obj, Sampler, Type, pValue);
}

HRESULT IDirect3DDevice9GetScissorRect(IDirect3DDevice9* obj, RECT* pRect) {
	return obj->lpVtbl->GetScissorRect(obj, pRect);
}

BOOL IDirect3DDevice9GetSoftwareVertexProcessing(IDirect3DDevice9* obj) {
	return obj->lpVtbl->GetSoftwareVertexProcessing(obj);
}

HRESULT IDirect3DDevice9GetStreamSource(
		IDirect3DDevice9* obj,
		UINT StreamNumber,
		IDirect3DVertexBuffer9** ppStreamData,
		UINT* pOffsetInBytes,
		UINT* pStride) {
	return obj->lpVtbl->GetStreamSource(obj, StreamNumber, ppStreamData,
		pOffsetInBytes, pStride);
}

HRESULT IDirect3DDevice9GetStreamSourceFreq(
		IDirect3DDevice9* obj,
		UINT StreamNumber,
		UINT* pDivider) {
	return obj->lpVtbl->GetStreamSourceFreq(obj, StreamNumber, pDivider);
}

HRESULT IDirect3DDevice9GetSwapChain(
		IDirect3DDevice9* obj,
		UINT iSwapChain,
		IDirect3DSwapChain9** ppSwapChain) {
	return obj->lpVtbl->GetSwapChain(obj, iSwapChain, ppSwapChain);
}

HRESULT IDirect3DDevice9GetTexture(
		IDirect3DDevice9* obj,
		DWORD Stage,
		IDirect3DBaseTexture9** ppTexture) {
	return obj->lpVtbl->GetTexture(obj, Stage, ppTexture);
}

HRESULT IDirect3DDevice9GetTextureStageState(
		IDirect3DDevice9* obj,
		DWORD Stage,
		D3DTEXTURESTAGESTATETYPE Type,
		DWORD* pValue) {
	return obj->lpVtbl->GetTextureStageState(obj, Stage, Type, pValue);
}

HRESULT IDirect3DDevice9GetTransform(
		IDirect3DDevice9* obj,
		D3DTRANSFORMSTATETYPE State,
		D3DMATRIX* pMatrix) {
	return obj->lpVtbl->GetTransform(obj, State, pMatrix);
}

HRESULT IDirect3DDevice9GetVertexDeclaration(
		IDirect3DDevice9* obj,
		IDirect3DVertexDeclaration9** ppDecl) {
	return obj->lpVtbl->GetVertexDeclaration(obj, ppDecl);
}

HRESULT IDirect3DDevice9GetVertexShader(
		IDirect3DDevice9* obj,
		IDirect3DVertexShader9** ppShader) {
	return obj->lpVtbl->GetVertexShader(obj, ppShader);
}

HRESULT IDirect3DDevice9GetVertexShaderConstantB(
		IDirect3DDevice9* obj,
		UINT StartRegister,
		BOOL* pConstantData,
		UINT BoolCount) {
	return obj->lpVtbl->GetVertexShaderConstantB(obj, StartRegister,
		pConstantData, BoolCount);
}

HRESULT IDirect3DDevice9GetVertexShaderConstantF(
		IDirect3DDevice9* obj,
		UINT StartRegister,
		float* pConstantData,
		UINT Vector4fCount) {
	return obj->lpVtbl->GetVertexShaderConstantF(obj, StartRegister,
		pConstantData, Vector4fCount);
}

HRESULT IDirect3DDevice9GetVertexShaderConstantI(
		IDirect3DDevice9* obj,
		UINT StartRegister,
		int* pConstantData,
		UINT Vector4iCount) {
	return obj->lpVtbl->GetVertexShaderConstantI(obj, StartRegister,
		pConstantData, Vector4iCount);
}

HRESULT IDirect3DDevice9GetViewport(
		IDirect3DDevice9* obj,
		D3DVIEWPORT9* pViewport) {
	return obj->lpVtbl->GetViewport(obj, pViewport);
}

HRESULT IDirect3DDevice9LightEnable(
		IDirect3DDevice9* obj,
		DWORD LightIndex,
		BOOL bEnable) {
	return obj->lpVtbl->LightEnable(obj, LightIndex, bEnable);
}

HRESULT IDirect3DDevice9MultiplyTransform(
		IDirect3DDevice9* obj,
		D3DTRANSFORMSTATETYPE State,
		D3DMATRIX* pMatrix) {
	return obj->lpVtbl->MultiplyTransform(obj, State, pMatrix);
}

HRESULT IDirect3DDevice9Present(
		IDirect3DDevice9* obj,
		RECT* pSourceRect,
		RECT* pDestRect,
		HWND hDestWindowOverride,
		RGNDATA* pDirtyRegion) {
	return obj->lpVtbl->Present(obj, pSourceRect, pDestRect,
		hDestWindowOverride, pDirtyRegion);
}

HRESULT IDirect3DDevice9ProcessVertices(
		IDirect3DDevice9* obj,
		UINT SrcStartIndex,
		UINT DestIndex,
		UINT VertexCount,
		IDirect3DVertexBuffer9* pDestBuffer,
		IDirect3DVertexDeclaration9* pVertexDecl,
		DWORD Flags) {
	return obj->lpVtbl->ProcessVertices(obj, SrcStartIndex, DestIndex,
		VertexCount, pDestBuffer, pVertexDecl, Flags);
}

HRESULT IDirect3DDevice9Reset(
		IDirect3DDevice9* obj,
		D3DPRESENT_PARAMETERS* pPresentationParameters) {
	return obj->lpVtbl->Reset(obj, pPresentationParameters);
}

HRESULT IDirect3DDevice9SetClipPlane(
		IDirect3DDevice9* obj,
		DWORD Index,
		float* pPlane) {
	return obj->lpVtbl->SetClipPlane(obj, Index, pPlane);
}

HRESULT IDirect3DDevice9SetClipStatus(
		IDirect3DDevice9* obj,
		D3DCLIPSTATUS9* pClipStatus) {
	return obj->lpVtbl->SetClipStatus(obj, pClipStatus);
}

HRESULT IDirect3DDevice9SetCurrentTexturePalette(
		IDirect3DDevice9* obj,
		UINT PaletteNumber) {
	return obj->lpVtbl->SetCurrentTexturePalette(obj, PaletteNumber);
}

void IDirect3DDevice9SetCursorPosition(
		IDirect3DDevice9* obj,
		INT X,
		INT Y,
		DWORD Flags) {
	obj->lpVtbl->SetCursorPosition(obj, X, Y, Flags);
}

HRESULT IDirect3DDevice9SetCursorProperties(
		IDirect3DDevice9* obj,
		UINT XHotSpot,
		UINT YHotSpot,
		IDirect3DSurface9* pCursorBitmap) {
	return obj->lpVtbl->SetCursorProperties(obj, XHotSpot, YHotSpot,
		pCursorBitmap);
}

HRESULT IDirect3DDevice9SetDepthStencilSurface(
		IDirect3DDevice9* obj,
		IDirect3DSurface9* pNewZStencil) {
	return obj->lpVtbl->SetDepthStencilSurface(obj, pNewZStencil);
}

HRESULT IDirect3DDevice9SetDialogBoxMode(
		IDirect3DDevice9* obj,
		BOOL bEnableDialogs) {
	return obj->lpVtbl->SetDialogBoxMode(obj, bEnableDialogs);
}

HRESULT IDirect3DDevice9SetFVF(IDirect3DDevice9* obj, DWORD FVF) {
	return obj->lpVtbl->SetFVF(obj, FVF);
}

void IDirect3DDevice9SetGammaRamp(
		IDirect3DDevice9* obj,
		UINT iSwapChain,
		DWORD Flags,
		D3DGAMMARAMP* pRamp) {
	obj->lpVtbl->SetGammaRamp(obj, iSwapChain, Flags, pRamp);
}

HRESULT IDirect3DDevice9SetIndices(
		IDirect3DDevice9* obj,
		IDirect3DIndexBuffer9* pIndexData) {
	return obj->lpVtbl->SetIndices(obj, pIndexData);
}

HRESULT IDirect3DDevice9SetLight(
		IDirect3DDevice9* obj,
		DWORD Index,
		D3DLIGHT9* pLight) {
	return obj->lpVtbl->SetLight(obj, Index, pLight);
}

HRESULT IDirect3DDevice9SetMaterial(
		IDirect3DDevice9* obj,
		D3DMATERIAL9* pMaterial) {
	return obj->lpVtbl->SetMaterial(obj, pMaterial);
}

HRESULT IDirect3DDevice9SetNPatchMode(IDirect3DDevice9* obj, float nSegments) {
	return obj->lpVtbl->SetNPatchMode(obj, nSegments);
}

HRESULT IDirect3DDevice9SetPaletteEntries(
		IDirect3DDevice9* obj,
		UINT PaletteNumber,
		PALETTEENTRY* pEntries) {
	return obj->lpVtbl->SetPaletteEntries(obj, PaletteNumber, pEntries);
}

HRESULT IDirect3DDevice9SetPixelShader(
		IDirect3DDevice9* obj,
		IDirect3DPixelShader9* pShader) {
	return obj->lpVtbl->SetPixelShader(obj, pShader);
}

HRESULT IDirect3DDevice9SetPixelShaderConstantB(
		IDirect3DDevice9* obj,
		UINT StartRegister,
		BOOL* pConstantData,
		UINT BoolCount) {
	return obj->lpVtbl->SetPixelShaderConstantB(obj, StartRegister,
		pConstantData, BoolCount);
}

HRESULT IDirect3DDevice9SetPixelShaderConstantF(
		IDirect3DDevice9* obj,
		UINT StartRegister,
		float* pConstantData,
		UINT Vector4fCount) {
	return obj->lpVtbl->SetPixelShaderConstantF(obj, StartRegister,
		pConstantData, Vector4fCount);
}

HRESULT IDirect3DDevice9SetPixelShaderConstantI(
		IDirect3DDevice9* obj,
		UINT StartRegister,
		int* pConstantData,
		UINT Vector4iCount) {
	return obj->lpVtbl->SetPixelShaderConstantI(obj, StartRegister,
		pConstantData, Vector4iCount);
}

HRESULT IDirect3DDevice9SetRenderState(
		IDirect3DDevice9* obj,
		D3DRENDERSTATETYPE State,
		DWORD Value) {
	return obj->lpVtbl->SetRenderState(obj, State, Value);
}

HRESULT IDirect3DDevice9SetRenderTarget(
		IDirect3DDevice9* obj,
		DWORD RenderTargetIndex,
		IDirect3DSurface9* pRenderTarget) {
	return obj->lpVtbl->SetRenderTarget(obj, RenderTargetIndex, pRenderTarget);
}

HRESULT IDirect3DDevice9SetSamplerState(
		IDirect3DDevice9* obj,
		DWORD Sampler,
		D3DSAMPLERSTATETYPE Type,
		DWORD Value) {
	return obj->lpVtbl->SetSamplerState(obj, Sampler, Type, Value);
}

HRESULT IDirect3DDevice9SetScissorRect(IDirect3DDevice9* obj, RECT* pRect) {
	return obj->lpVtbl->SetScissorRect(obj, pRect);
}

HRESULT IDirect3DDevice9SetSoftwareVertexProcessing(
		IDirect3DDevice9* obj,
		BOOL bSoftware) {
	return obj->lpVtbl->SetSoftwareVertexProcessing(obj, bSoftware);
}

HRESULT IDirect3DDevice9SetStreamSource(
		IDirect3DDevice9* obj,
		UINT StreamNumber,
		IDirect3DVertexBuffer9* pStreamData,
		UINT OffsetInBytes,
		UINT Stride) {
	return obj->lpVtbl->SetStreamSource(obj, StreamNumber, pStreamData,
		OffsetInBytes, Stride);
}

HRESULT IDirect3DDevice9SetStreamSourceFreq(
		IDirect3DDevice9* obj,
		UINT StreamNumber,
		UINT FrequencyParameter) {
	return obj->lpVtbl->SetStreamSourceFreq(obj, StreamNumber,
		FrequencyParameter);
}

HRESULT IDirect3DDevice9SetTexture(
		IDirect3DDevice9* obj,
		DWORD Sampler,
		IDirect3DBaseTexture9* pTexture) {
	return obj->lpVtbl->SetTexture(obj, Sampler, pTexture);
}

HRESULT IDirect3DDevice9SetTextureStageState(
		IDirect3DDevice9* obj,
		DWORD Stage,
		D3DTEXTURESTAGESTATETYPE Type,
		DWORD Value) {
	return obj->lpVtbl->SetTextureStageState(obj, Stage, Type, Value);
}

HRESULT IDirect3DDevice9SetTransform(
		IDirect3DDevice9* obj,
		D3DTRANSFORMSTATETYPE State,
		D3DMATRIX* pMatrix) {
	return obj->lpVtbl->SetTransform(obj, State, pMatrix);
}

HRESULT IDirect3DDevice9SetVertexDeclaration(
		IDirect3DDevice9* obj,
		IDirect3DVertexDeclaration9* pDecl) {
	return obj->lpVtbl->SetVertexDeclaration(obj, pDecl);
}

HRESULT IDirect3DDevice9SetVertexShader(
		IDirect3DDevice9* obj,
		IDirect3DVertexShader9* pShader) {
	return obj->lpVtbl->SetVertexShader(obj, pShader);
}

HRESULT IDirect3DDevice9SetVertexShaderConstantB(
		IDirect3DDevice9* obj,
		UINT StartRegister,
		BOOL* pConstantData,
		UINT BoolCount) {
	return obj->lpVtbl->SetVertexShaderConstantB(obj, StartRegister,
		pConstantData, BoolCount);
}

HRESULT IDirect3DDevice9SetVertexShaderConstantF(
		IDirect3DDevice9* obj,
		UINT StartRegister,
		float* pConstantData,
		UINT Vector4fCount) {
	return obj->lpVtbl->SetVertexShaderConstantF(obj, StartRegister,
		pConstantData, Vector4fCount);
}

HRESULT IDirect3DDevice9SetVertexShaderConstantI(
		IDirect3DDevice9* obj,
		UINT StartRegister,
		int* pConstantData,
		UINT Vector4iCount) {
	return obj->lpVtbl->SetVertexShaderConstantI(obj, StartRegister,
		pConstantData, Vector4iCount);
}

HRESULT IDirect3DDevice9SetViewport(
		IDirect3DDevice9* obj,
		D3DVIEWPORT9* pViewport) {
	return obj->lpVtbl->SetViewport(obj, pViewport);
}

BOOL IDirect3DDevice9ShowCursor(IDirect3DDevice9* obj, BOOL bShow) {
	return obj->lpVtbl->ShowCursor(obj, bShow);
}

HRESULT IDirect3DDevice9StretchRect(
		IDirect3DDevice9* obj,
		IDirect3DSurface9* pSourceSurface,
		RECT* pSourceRect,
		IDirect3DSurface9* pDestSurface,
		RECT* pDestRect,
		D3DTEXTUREFILTERTYPE Filter) {
	return obj->lpVtbl->StretchRect(obj, pSourceSurface, pSourceRect,
		pDestSurface, pDestRect, Filter);
}

HRESULT IDirect3DDevice9TestCooperativeLevel(IDirect3DDevice9* obj) {
	return obj->lpVtbl->TestCooperativeLevel(obj);
}

HRESULT IDirect3DDevice9UpdateSurface(
		IDirect3DDevice9* obj,
		IDirect3DSurface9* pSourceSurface,
		RECT* pSourceRect,
		IDirect3DSurface9* pDestinationSurface,
		POINT* pDestinationPoint) {
	return obj->lpVtbl->UpdateSurface(obj, pSourceSurface, pSourceRect,
		pDestinationSurface, pDestinationPoint);
}

HRESULT IDirect3DDevice9UpdateTexture(
		IDirect3DDevice9* obj,
		IDirect3DBaseTexture9* pSourceTexture,
		IDirect3DBaseTexture9* pDestinationTexture) {
	return obj->lpVtbl->UpdateTexture(obj, pSourceTexture, pDestinationTexture);
}

HRESULT IDirect3DDevice9ValidateDevice(
		IDirect3DDevice9* obj,
		DWORD* pNumPasses) {
	return obj->lpVtbl->ValidateDevice(obj, pNumPasses);
}

void IDirect3DDevice9Release(IDirect3DDevice9* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"
import "unsafe"

// Device and its methods are used to perform DrawPrimitive-based rendering,
// create resources, work with system-level variables, adjust gamma ramp levels,
// work with palettes and create shaders.
type Device struct {
	handle *C.IDirect3DDevice9
}

// Release has to be called when finished using the object to free its
// associated resources.
func (obj Device) Release() {
	C.IDirect3DDevice9Release(obj.handle)
}

// BeginScene begins a scene.
// Applications must call BeginScene before performing any rendering and must
// call EndScene when rendering is complete and before calling BeginScene again.
func (obj Device) BeginScene() (err Error) {
	err = toErr(C.IDirect3DDevice9BeginScene(obj.handle))
	return
}

// BeginStateBlock signals Direct3D to begin recording a device-state block.
func (obj Device) BeginStateBlock() (err Error) {
	err = toErr(C.IDirect3DDevice9BeginStateBlock(obj.handle))
	return
}

// Clear clears one or more surfaces such as a render target, multiple render
// targets, a stencil buffer, and a depth buffer.
func (obj Device) Clear(
	pRects []D3DRECT,
	Flags uint32,
	Color COLOR,
	Z float32,
	Stencil uint32,
) (
	err Error,
) {
	if len(pRects) == 0 {
		err = toErr(C.IDirect3DDevice9Clear(
			obj.handle,
			0,
			nil,
			(C.DWORD)(Flags),
			(C.D3DCOLOR)(Color),
			(C.float)(Z),
			(C.DWORD)(Stencil),
		))
	} else {
		cpRects := make([]C.D3DRECT, len(pRects))
		for i := range cpRects {
			cpRects[i] = pRects[i].toC()
		}
		err = toErr(C.IDirect3DDevice9Clear(
			obj.handle,
			(C.DWORD)(len(pRects)),
			&cpRects[0],
			(C.DWORD)(Flags),
			(C.D3DCOLOR)(Color),
			(C.float)(Z),
			(C.DWORD)(Stencil),
		))
	}
	return
}

// ColorFill allows an application to fill a rectangular area of a
// POOL_DEFAULT surface with a specified color.
func (obj Device) ColorFill(
	pSurface Surface,
	pRect *RECT,
	color COLOR,
) (
	err Error,
) {
	if pRect == nil {
		err = toErr(C.IDirect3DDevice9ColorFill(
			obj.handle,
			pSurface.handle,
			nil,
			(C.D3DCOLOR)(color),
		))
	} else {
		cpRect := pRect.toC()
		err = toErr(C.IDirect3DDevice9ColorFill(
			obj.handle,
			pSurface.handle,
			&cpRect,
			(C.D3DCOLOR)(color),
		))
	}
	return
}

// CreateAdditionalSwapChain creates an additional swap chain for rendering
// multiple views.
func (obj Device) CreateAdditionalSwapChain(
	inpPresentationParameters PRESENT_PARAMETERS,
) (
	ppSwapChain SwapChain,
	outpPresentationParameters PRESENT_PARAMETERS,
	err Error,
) {
	cpPresentationParameters := inpPresentationParameters.toC()
	var cppSwapChain *C.IDirect3DSwapChain9
	err = toErr(C.IDirect3DDevice9CreateAdditionalSwapChain(
		obj.handle,
		&cpPresentationParameters,
		&cppSwapChain,
	))
	outpPresentationParameters.fromC(&cpPresentationParameters)
	ppSwapChain = SwapChain{cppSwapChain}
	return
}

// CreateCubeTexture creates a cube texture resource.
func (obj Device) CreateCubeTexture(
	EdgeLength uint,
	Levels uint,
	Usage uint32,
	Format FORMAT,
	Pool POOL,
	pSharedHandle unsafe.Pointer,
) (
	ppCubeTexture CubeTexture,
	err Error,
) {
	var cppCubeTexture *C.IDirect3DCubeTexture9
	err = toErr(C.IDirect3DDevice9CreateCubeTexture(
		obj.handle,
		(C.UINT)(EdgeLength),
		(C.UINT)(Levels),
		(C.DWORD)(Usage),
		(C.D3DFORMAT)(Format),
		(C.D3DPOOL)(Pool),
		&cppCubeTexture,
		(*C.HANDLE)(pSharedHandle),
	))
	resource := Resource{
		(*C.IDirect3DResource9)(unsafe.Pointer(cppCubeTexture)),
	}
	base := BaseTexture{
		resource, (*C.IDirect3DBaseTexture9)(unsafe.Pointer(cppCubeTexture)),
	}
	ppCubeTexture = CubeTexture{base, cppCubeTexture}
	return
}

// CreateDepthStencilSurface creates a depth-stencil resource.
func (obj Device) CreateDepthStencilSurface(
	Width uint,
	Height uint,
	Format FORMAT,
	MultiSample MULTISAMPLE_TYPE,
	MultisampleQuality uint32,
	Discard bool,
	pSharedHandle unsafe.Pointer,
) (
	ppSurface Surface,
	err Error,
) {
	var cppSurface *C.IDirect3DSurface9
	err = toErr(C.IDirect3DDevice9CreateDepthStencilSurface(
		obj.handle,
		(C.UINT)(Width),
		(C.UINT)(Height),
		(C.D3DFORMAT)(Format),
		(C.D3DMULTISAMPLE_TYPE)(MultiSample),
		(C.DWORD)(MultisampleQuality),
		toBOOL(Discard),
		&cppSurface,
		(*C.HANDLE)(pSharedHandle),
	))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(cppSurface))}
	ppSurface = Surface{resource, cppSurface}
	return
}

// CreateIndexBuffer creates an index buffer.
func (obj Device) CreateIndexBuffer(
	Length uint,
	Usage uint32,
	Format FORMAT,
	Pool POOL,
	pSharedHandle unsafe.Pointer,
) (
	ppIndexBuffer IndexBuffer,
	err Error,
) {
	var cppIndexBuffer *C.IDirect3DIndexBuffer9
	err = toErr(C.IDirect3DDevice9CreateIndexBuffer(
		obj.handle,
		(C.UINT)(Length),
		(C.DWORD)(Usage),
		(C.D3DFORMAT)(Format),
		(C.D3DPOOL)(Pool),
		&cppIndexBuffer,
		(*C.HANDLE)(pSharedHandle),
	))
	resource := Resource{
		(*C.IDirect3DResource9)(unsafe.Pointer(cppIndexBuffer)),
	}
	ppIndexBuffer = IndexBuffer{resource, cppIndexBuffer}
	return
}

// CreateOffscreenPlainSurface creates an off-screen surface.
func (obj Device) CreateOffscreenPlainSurface(
	Width uint,
	Height uint,
	Format FORMAT,
	Pool POOL,
	pSharedHandle unsafe.Pointer,
) (
	ppSurface Surface,
	err Error,
) {
	var cppSurface *C.IDirect3DSurface9
	err = toErr(C.IDirect3DDevice9CreateOffscreenPlainSurface(
		obj.handle,
		(C.UINT)(Width),
		(C.UINT)(Height),
		(C.D3DFORMAT)(Format),
		(C.D3DPOOL)(Pool),
		&cppSurface,
		(*C.HANDLE)(pSharedHandle),
	))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(cppSurface))}
	ppSurface = Surface{resource, cppSurface}
	return
}

// CreatePixelShader creates a pixel shader.
func (obj Device) CreatePixelShader(
	pFunction unsafe.Pointer,
) (
	ppShader PixelShader,
	err Error,
) {
	var cppShader *C.IDirect3DPixelShader9
	err = toErr(C.IDirect3DDevice9CreatePixelShader(
		obj.handle,
		(*C.DWORD)(pFunction),
		&cppShader,
	))
	ppShader = PixelShader{cppShader}
	return
}

// CreatePixelShaderFromBytes creates a pixel shader from a byte slice.
func (obj Device) CreatePixelShaderFromBytes(
	pFunction []byte,
) (
	ppShader PixelShader,
	err Error,
) {
	return obj.CreatePixelShader(unsafe.Pointer(&pFunction[0]))
}

// CreateQuery creates a status query.
func (obj Device) CreateQuery(Type QUERYTYPE) (ppQuery Query, err Error) {
	var cppQuery *C.IDirect3DQuery9
	err = toErr(C.IDirect3DDevice9CreateQuery(
		obj.handle,
		(C.D3DQUERYTYPE)(Type),
		&cppQuery,
	))
	ppQuery = Query{cppQuery}
	return
}

// CreateRenderTarget creates a render-target surface.
func (obj Device) CreateRenderTarget(
	Width uint,
	Height uint,
	Format FORMAT,
	MultiSample MULTISAMPLE_TYPE,
	MultisampleQuality uint32,
	Lockable bool,
	pSharedHandle unsafe.Pointer,
) (
	ppSurface Surface,
	err Error,
) {
	var cppSurface *C.IDirect3DSurface9
	err = toErr(C.IDirect3DDevice9CreateRenderTarget(
		obj.handle,
		(C.UINT)(Width),
		(C.UINT)(Height),
		(C.D3DFORMAT)(Format),
		(C.D3DMULTISAMPLE_TYPE)(MultiSample),
		(C.DWORD)(MultisampleQuality),
		toBOOL(Lockable),
		&cppSurface,
		(*C.HANDLE)(pSharedHandle),
	))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(cppSurface))}
	ppSurface = Surface{resource, cppSurface}
	return
}

// CreateStateBlock creates a new state block that contains the values for all
// device states, vertex-related states, or pixel-related states.
func (obj Device) CreateStateBlock(
	Type STATEBLOCKTYPE,
) (
	ppSB StateBlock,
	err Error,
) {
	var cppSB *C.IDirect3DStateBlock9
	err = toErr(C.IDirect3DDevice9CreateStateBlock(
		obj.handle,
		(C.D3DSTATEBLOCKTYPE)(Type),
		&cppSB,
	))
	ppSB = StateBlock{cppSB}
	return
}

// CreateTexture creates a texture resource.
func (obj Device) CreateTexture(
	Width uint,
	Height uint,
	Levels uint,
	Usage uint32,
	Format FORMAT,
	Pool POOL,
	pSharedHandle unsafe.Pointer,
) (
	ppTexture Texture,
	err Error,
) {
	var cppTexture *C.IDirect3DTexture9
	err = toErr(C.IDirect3DDevice9CreateTexture(
		obj.handle,
		(C.UINT)(Width),
		(C.UINT)(Height),
		(C.UINT)(Levels),
		(C.DWORD)(Usage),
		(C.D3DFORMAT)(Format),
		(C.D3DPOOL)(Pool),
		&cppTexture,
		(*C.HANDLE)(pSharedHandle),
	))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(cppTexture))}
	base := BaseTexture{
		resource,
		(*C.IDirect3DBaseTexture9)(unsafe.Pointer(cppTexture)),
	}
	ppTexture = Texture{base, cppTexture}
	return
}

// CreateVertexBuffer creates a vertex buffer.
func (obj Device) CreateVertexBuffer(
	Length uint,
	Usage uint32,
	FVF uint32,
	Pool POOL,
	pSharedHandle unsafe.Pointer,
) (
	ppVertexBuffer VertexBuffer,
	err Error,
) {
	var cppVertexBuffer *C.IDirect3DVertexBuffer9
	err = toErr(C.IDirect3DDevice9CreateVertexBuffer(
		obj.handle,
		(C.UINT)(Length),
		(C.DWORD)(Usage),
		(C.DWORD)(FVF),
		(C.D3DPOOL)(Pool),
		&cppVertexBuffer,
		(*C.HANDLE)(pSharedHandle),
	))
	resource := Resource{
		(*C.IDirect3DResource9)(unsafe.Pointer(cppVertexBuffer)),
	}
	ppVertexBuffer = VertexBuffer{resource, cppVertexBuffer}
	return
}

// CreateVertexDeclaration creates a vertex shader declaration from the device
// and the vertex elements.
// The last element should be DeclEnd().
func (obj Device) CreateVertexDeclaration(
	pVertexElements []VERTEXELEMENT,
) (
	ppDecl VertexDeclaration,
	err Error,
) {
	cpVertexElements := make([]C.D3DVERTEXELEMENT9, len(pVertexElements))
	for i := range cpVertexElements {
		cpVertexElements[i] = pVertexElements[i].toC()
	}
	var cppDecl *C.IDirect3DVertexDeclaration9
	err = toErr(C.IDirect3DDevice9CreateVertexDeclaration(
		obj.handle,
		&cpVertexElements[0],
		&cppDecl,
	))
	ppDecl = VertexDeclaration{cppDecl}
	return
}

// DeclEnd marks the end of a slice of VERTEXELEMENTs.
func DeclEnd() VERTEXELEMENT {
	return VERTEXELEMENT{0xFF, 0, DECLTYPE_UNUSED, 0, 0, 0}
}

// CreateVertexShader creates a vertex shader.
func (obj Device) CreateVertexShader(
	pFunction unsafe.Pointer,
) (
	ppShader VertexShader,
	err Error,
) {
	var cppShader *C.IDirect3DVertexShader9
	err = toErr(C.IDirect3DDevice9CreateVertexShader(
		obj.handle,
		(*C.DWORD)(pFunction),
		&cppShader,
	))
	ppShader = VertexShader{cppShader}
	return
}

// CreateVertexShaderFromBytes creates a vertex shader from a byte slice.
func (obj Device) CreateVertexShaderFromBytes(
	pFunction []byte,
) (
	ppShader VertexShader,
	err Error,
) {
	return obj.CreateVertexShader(unsafe.Pointer(&pFunction[0]))
}

// CreateVolumeTexture creates a volume texture resource.
func (obj Device) CreateVolumeTexture(
	Width uint,
	Height uint,
	Depth uint,
	Levels uint,
	Usage uint32,
	Format FORMAT,
	Pool POOL,
	pSharedHandle unsafe.Pointer,
) (
	ppVolumeTexture VolumeTexture,
	err Error,
) {
	var cppVolumeTexture *C.IDirect3DVolumeTexture9
	err = toErr(C.IDirect3DDevice9CreateVolumeTexture(
		obj.handle,
		(C.UINT)(Width),
		(C.UINT)(Height),
		(C.UINT)(Depth),
		(C.UINT)(Levels),
		(C.DWORD)(Usage),
		(C.D3DFORMAT)(Format),
		(C.D3DPOOL)(Pool),
		&cppVolumeTexture,
		(*C.HANDLE)(pSharedHandle),
	))
	resource := Resource{
		(*C.IDirect3DResource9)(unsafe.Pointer(cppVolumeTexture)),
	}
	base := BaseTexture{
		resource,
		(*C.IDirect3DBaseTexture9)(unsafe.Pointer(cppVolumeTexture)),
	}
	ppVolumeTexture = VolumeTexture{base, cppVolumeTexture}
	return
}

// DeletePatch frees a cached high-order patch.
func (obj Device) DeletePatch(Handle uint) (err Error) {
	err = toErr(C.IDirect3DDevice9DeletePatch(obj.handle, (C.UINT)(Handle)))
	return
}

// DrawIndexedPrimitive renders the specified geometric
// primitive into an array of vertices, based on indexing.
func (obj Device) DrawIndexedPrimitive(
	Type PRIMITIVETYPE,
	BaseVertexIndex int,
	MinIndex uint,
	NumVertices uint,
	StartIndex uint,
	PrimitiveCount uint,
) (err Error) {
	err = toErr(C.IDirect3DDevice9DrawIndexedPrimitive(
		obj.handle,
		(C.D3DPRIMITIVETYPE)(Type),
		(C.INT)(BaseVertexIndex),
		(C.UINT)(MinIndex),
		(C.UINT)(NumVertices),
		(C.UINT)(StartIndex),
		(C.UINT)(PrimitiveCount),
	))
	return
}

// DrawIndexedPrimitiveUP renders the specified geometric primitive with data
// specified by a user memory pointer.
func (obj Device) DrawIndexedPrimitiveUP(
	PrimitiveType PRIMITIVETYPE,
	MinVertexIndex uint,
	NumVertices uint,
	PrimitiveCount uint,
	pIndexData unsafe.Pointer,
	IndexDataFormat FORMAT,
	pVertexStreamZeroData unsafe.Pointer,
	VertexStreamZeroStride uint,
) (
	err Error,
) {
	err = toErr(C.IDirect3DDevice9DrawIndexedPrimitiveUP(
		obj.handle,
		(C.D3DPRIMITIVETYPE)(PrimitiveType),
		(C.UINT)(MinVertexIndex),
		(C.UINT)(NumVertices),
		(C.UINT)(PrimitiveCount),
		pIndexData,
		(C.D3DFORMAT)(IndexDataFormat),
		pVertexStreamZeroData,
		(C.UINT)(VertexStreamZeroStride),
	))
	return
}

// DrawIndexedPrimitiveUPuint32 renders the specified geometric primitive with
// data specified by a user memory pointer.
func (obj Device) DrawIndexedPrimitiveUPuint32(
	PrimitiveType PRIMITIVETYPE,
	MinVertexIndex uint,
	NumVertices uint,
	PrimitiveCount uint,
	indexData []uint32,
	pVertexStreamZeroData unsafe.Pointer,
	VertexStreamZeroStride uint,
) (
	err Error,
) {
	return obj.DrawIndexedPrimitiveUP(
		PrimitiveType,
		MinVertexIndex,
		NumVertices,
		PrimitiveCount,
		unsafe.Pointer(&indexData[0]),
		FMT_INDEX32,
		pVertexStreamZeroData,
		VertexStreamZeroStride,
	)
}

// DrawIndexedPrimitiveUPuint16 renders the specified geometric primitive with
// data specified by a user memory pointer.
func (obj Device) DrawIndexedPrimitiveUPuint16(
	PrimitiveType PRIMITIVETYPE,
	MinVertexIndex uint,
	NumVertices uint,
	PrimitiveCount uint,
	indexData []uint16,
	pVertexStreamZeroData unsafe.Pointer,
	VertexStreamZeroStride uint,
) (
	err Error,
) {
	return obj.DrawIndexedPrimitiveUP(
		PrimitiveType,
		MinVertexIndex,
		NumVertices,
		PrimitiveCount,
		unsafe.Pointer(&indexData[0]),
		FMT_INDEX16,
		pVertexStreamZeroData,
		VertexStreamZeroStride,
	)
}

// DrawPrimitive renders a sequence of nonindexed, geometric primitives of the
// specified type from the current set of data input streams.
func (obj Device) DrawPrimitive(
	PrimitiveType PRIMITIVETYPE,
	StartVertex uint,
	PrimitiveCount uint,
) (
	err Error) {

	err = toErr(C.IDirect3DDevice9DrawPrimitive(
		obj.handle,
		(C.D3DPRIMITIVETYPE)(PrimitiveType),
		(C.UINT)(StartVertex),
		(C.UINT)(PrimitiveCount),
	))
	return
}

// DrawPrimitiveUP renders data specified by a user memory pointer as a sequence
// of geometric primitives of the specified type.
func (obj Device) DrawPrimitiveUP(
	PrimitiveType PRIMITIVETYPE,
	PrimitiveCount uint,
	pVertexStreamZeroData unsafe.Pointer,
	VertexStreamZeroStride uint,
) (
	err Error,
) {
	err = toErr(C.IDirect3DDevice9DrawPrimitiveUP(
		obj.handle,
		(C.D3DPRIMITIVETYPE)(PrimitiveType),
		(C.UINT)(PrimitiveCount),
		pVertexStreamZeroData,
		(C.UINT)(VertexStreamZeroStride),
	))
	return
}

// DrawRectPatch draws a rectangular patch using the currently set streams.
func (obj Device) DrawRectPatch(
	Handle uint,
	pNumSegs [4]float32,
	pRectPatchInfo *RECTPATCH_INFO,
) (
	err Error,
) {
	if pRectPatchInfo == nil {
		err = toErr(C.IDirect3DDevice9DrawRectPatch(
			obj.handle,
			C.UINT(Handle),
			(*C.float)(&pNumSegs[0]),
			nil,
		))
	} else {
		cpRectPatchInfo := pRectPatchInfo.toC()
		err = toErr(C.IDirect3DDevice9DrawRectPatch(
			obj.handle,
			C.UINT(Handle),
			(*C.float)(&pNumSegs[0]),
			&cpRectPatchInfo,
		))
	}
	return
}

// DrawTriPatch draws a triangular patch using the currently set streams.
func (obj Device) DrawTriPatch(
	Handle uint,
	pNumSegs [3]float32,
	pTriPatchInfo *TRIPATCH_INFO,
) (
	err Error,
) {
	if pTriPatchInfo == nil {
		err = toErr(C.IDirect3DDevice9DrawTriPatch(
			obj.handle,
			C.UINT(Handle),
			(*C.float)(&pNumSegs[0]),
			nil,
		))
	} else {
		cpTriPatchInfo := pTriPatchInfo.toC()
		err = toErr(C.IDirect3DDevice9DrawTriPatch(
			obj.handle,
			C.UINT(Handle),
			(*C.float)(&pNumSegs[0]),
			&cpTriPatchInfo,
		))
	}
	return
}

// EndScene ends a scene that was begun by calling BeginScene.
func (obj Device) EndScene() (err Error) {
	err = toErr(C.IDirect3DDevice9EndScene(obj.handle))
	return
}

// EndStateBlock signals Direct3D to stop recording a device-state block and
// retrieve a pointer to the state block interface.
func (obj Device) EndStateBlock() (ppSB StateBlock, err Error) {
	var cppSB *C.IDirect3DStateBlock9
	err = toErr(C.IDirect3DDevice9EndStateBlock(obj.handle, &cppSB))
	ppSB = StateBlock{cppSB}
	return
}

// EvictManagedResources evicts all managed resources, including both Direct3D
// and driver-managed resources.
func (obj Device) EvictManagedResources() (err Error) {
	err = toErr(C.IDirect3DDevice9EvictManagedResources(obj.handle))
	return
}

// GetAvailableTextureMem returns an estimate of the amount of available texture
// memory.
func (obj Device) GetAvailableTextureMem() uint {
	return (uint)(C.IDirect3DDevice9GetAvailableTextureMem(obj.handle))
}

// GetBackBuffer retrieves a back buffer from the device's swap chain.
// Call Release on the returned surface when finished using it.
func (obj Device) GetBackBuffer(
	iSwapChain uint,
	BackBuffer uint,
	Type BACKBUFFER_TYPE,
) (
	ppBackBuffer Surface,
	err Error,
) {
	var cppBackBuffer *C.IDirect3DSurface9
	err = toErr(C.IDirect3DDevice9GetBackBuffer(
		obj.handle,
		(C.UINT)(iSwapChain),
		(C.UINT)(BackBuffer),
		(C.D3DBACKBUFFER_TYPE)(Type),
		&cppBackBuffer,
	))
	resource := Resource{
		(*C.IDirect3DResource9)(unsafe.Pointer(cppBackBuffer)),
	}
	ppBackBuffer = Surface{resource, cppBackBuffer}
	return
}

// GetClipPlane retrieves the coefficients of a user-defined clipping plane for
// the device.
func (obj Device) GetClipPlane(Index uint32) (pPlane [4]float32, err Error) {
	err = toErr(C.IDirect3DDevice9GetClipPlane(
		obj.handle,
		(C.DWORD)(Index),
		(*C.float)(&pPlane[0]),
	))
	return
}

// GetClipStatus retrieves the clip status.
func (obj Device) GetClipStatus() (pClipStatus CLIPSTATUS, err Error) {
	var cpClipStatus C.D3DCLIPSTATUS9
	err = toErr(C.IDirect3DDevice9GetClipStatus(obj.handle, &cpClipStatus))
	pClipStatus.fromC(&cpClipStatus)
	return
}

// GetCreationParameters retrieves the creation parameters of the device.
func (obj Device) GetCreationParameters() (
	pParameters DEVICE_CREATION_PARAMETERS,
	err Error,
) {
	var cpParameters C.D3DDEVICE_CREATION_PARAMETERS
	err = toErr(C.IDirect3DDevice9GetCreationParameters(
		obj.handle,
		&cpParameters,
	))
	pParameters.fromC(&cpParameters)
	return
}

// GetCurrentTexturePalette retrieves the current texture palette.
func (obj Device) GetCurrentTexturePalette() (pPaletteNumber uint, err Error) {
	var cpPaletteNumber C.UINT
	err = toErr(C.IDirect3DDevice9GetCurrentTexturePalette(
		obj.handle,
		&cpPaletteNumber,
	))
	pPaletteNumber = (uint)(cpPaletteNumber)
	return
}

// GetDepthStencilSurface returns the depth-stencil surface owned by the
// Direct3DDevice object.
// Call Release on the returned surface when finished using it.
func (obj Device) GetDepthStencilSurface() (
	ppZStencilSurface Surface,
	err Error,
) {
	var cppZStencilSurface *C.IDirect3DSurface9
	err = toErr(C.IDirect3DDevice9GetDepthStencilSurface(
		obj.handle,
		&cppZStencilSurface,
	))
	resource := Resource{
		(*C.IDirect3DResource9)(unsafe.Pointer(cppZStencilSurface)),
	}
	ppZStencilSurface = Surface{resource, cppZStencilSurface}
	return
}

// GetDeviceCaps retrieves the capabilities of the rendering device.
func (obj Device) GetDeviceCaps() (pCaps CAPS, err Error) {
	var cpCaps C.D3DCAPS9
	err = toErr(C.IDirect3DDevice9GetDeviceCaps(obj.handle, &cpCaps))
	pCaps.fromC(&cpCaps)
	return
}

// GetDirect3D returns an interface to the instance of the Direct3D object that
// created the device.
// Call Release on the returned object when finished using it.
func (obj Device) GetDirect3D() (ppD3D9 Direct3D, err Error) {
	var cppD3D9 *C.IDirect3D9
	err = toErr(C.IDirect3DDevice9GetDirect3D(obj.handle, &cppD3D9))
	ppD3D9 = Direct3D{cppD3D9}
	return
}

// GetDisplayMode retrieves the display mode's spatial resolution, color
// resolution, and refresh frequency.
func (obj Device) GetDisplayMode(
	iSwapChain uint,
) (
	pMode DISPLAYMODE,
	err Error,
) {
	var cpMode C.D3DDISPLAYMODE
	err = toErr(C.IDirect3DDevice9GetDisplayMode(
		obj.handle,
		(C.UINT)(iSwapChain),
		&cpMode,
	))
	pMode.fromC(&cpMode)
	return
}

// GetFrontBufferData generates a copy of the device's front buffer and places
// that copy in a system memory buffer provided by the application.
func (obj Device) GetFrontBufferData(
	iSwapChain uint,
	pDestSurface Surface,
) (
	err Error,
) {
	err = toErr(C.IDirect3DDevice9GetFrontBufferData(
		obj.handle,
		(C.UINT)(iSwapChain),
		pDestSurface.handle,
	))
	return
}

// GetFVF gets the fixed vertex function declaration.
func (obj Device) GetFVF() (pFVF uint32, err Error) {
	var cpFVF C.DWORD
	err = toErr(C.IDirect3DDevice9GetFVF(obj.handle, &cpFVF))
	pFVF = (uint32)(cpFVF)
	return
}

// GetGammaRamp retrieves the gamma correction ramp for the swap chain.
func (obj Device) GetGammaRamp(iSwapChain uint) (pRamp GAMMARAMP) {
	var cpRamp C.D3DGAMMARAMP
	C.IDirect3DDevice9GetGammaRamp(obj.handle, (C.UINT)(iSwapChain), &cpRamp)
	pRamp.fromC(&cpRamp)
	return
}

// GetIndices retrieves index data.
func (obj Device) GetIndices() (ppIndexData IndexBuffer, err Error) {
	var cppIndexData *C.IDirect3DIndexBuffer9
	err = toErr(C.IDirect3DDevice9GetIndices(obj.handle, &cppIndexData))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(cppIndexData))}
	ppIndexData = IndexBuffer{resource, cppIndexData}
	return
}

// GetLight retrieves a set of lighting properties that this device uses.
func (obj Device) GetLight(Index uint32) (pLight LIGHT, err Error) {
	var cpLight C.D3DLIGHT9
	err = toErr(C.IDirect3DDevice9GetLight(
		obj.handle,
		(C.DWORD)(Index),
		&cpLight,
	))
	pLight.fromC(&cpLight)
	return
}

// GetLightEnable retrieves the activity status - enabled or disabled - for a
// set of lighting parameters within a device.
func (obj Device) GetLightEnable(Index uint32) (pEnable bool, err Error) {
	var cpEnable C.BOOL
	err = toErr(C.IDirect3DDevice9GetLightEnable(
		obj.handle,
		(C.DWORD)(Index),
		&cpEnable,
	))
	pEnable = cpEnable != 0
	return
}

// GetMaterial retrieves the current material properties for the device.
func (obj Device) GetMaterial() (pMaterial MATERIAL, err Error) {
	var cpMaterial C.D3DMATERIAL9
	err = toErr(C.IDirect3DDevice9GetMaterial(obj.handle, &cpMaterial))
	pMaterial.fromC(&cpMaterial)
	return
}

// GetNPatchMode returns the N-patch mode segments.
func (obj Device) GetNPatchMode() float32 {
	return (float32)(C.IDirect3DDevice9GetNPatchMode(obj.handle))
}

// GetNumberOfSwapChains returns the number of implicit swap chains.
func (obj Device) GetNumberOfSwapChains() uint {
	return (uint)(C.IDirect3DDevice9GetNumberOfSwapChains(obj.handle))
}

// GetPaletteEntries retrieves palette entries.
func (obj Device) GetPaletteEntries(
	PaletteNumber uint,
) (
	pEntries [256]PALETTEENTRY,
	err Error,
) {
	var cpEntries [256]C.PALETTEENTRY
	err = toErr(C.IDirect3DDevice9GetPaletteEntries(
		obj.handle,
		(C.UINT)(PaletteNumber),
		&cpEntries[0],
	))
	for i := range pEntries {
		pEntries[i].fromC(&cpEntries[i])
	}
	return
}

// GetPixelShader retrieves the currently set pixel shader.
func (obj Device) GetPixelShader() (ppShader PixelShader, err Error) {
	var cppShader *C.IDirect3DPixelShader9
	err = toErr(C.IDirect3DDevice9GetPixelShader(obj.handle, &cppShader))
	ppShader = PixelShader{cppShader}
	return
}

// GetPixelShaderConstantB returns a Boolean shader constant. The given bool
// slice is filled with the constant data.
func (obj Device) GetPixelShaderConstantB(
	StartRegister uint,
	pConstantData []bool,
) (
	err Error,
) {
	cpConstantData := make([]C.BOOL, len(pConstantData))
	err = toErr(C.IDirect3DDevice9GetPixelShaderConstantB(
		obj.handle,
		(C.UINT)(StartRegister),
		&cpConstantData[0],
		(C.UINT)(len(pConstantData)),
	))
	for i := range cpConstantData {
		pConstantData[i] = cpConstantData[i] != 0
	}
	return
}

// GetPixelShaderConstantF returns a floating-point shader constant. The given
// float slice is filled with the constant data. The length of the slice must be
// a multiple of four.
func (obj Device) GetPixelShaderConstantF(
	StartRegister uint,
	pConstantData []float32,
) (
	err Error,
) {
	err = toErr(C.IDirect3DDevice9GetPixelShaderConstantF(
		obj.handle,
		(C.UINT)(StartRegister),
		(*C.float)(&pConstantData[0]),
		(C.UINT)(len(pConstantData)/4),
	))
	return
}

// GetPixelShaderConstantI returns an integer shader constant. The given int
// slice is filled with the constant data. The length of the slice must be a
// multiple of four.
func (obj Device) GetPixelShaderConstantI(
	StartRegister uint,
	pConstantData []int32,
) (
	err Error,
) {
	err = toErr(C.IDirect3DDevice9GetPixelShaderConstantI(
		obj.handle,
		(C.UINT)(StartRegister),
		(*C.int)(&pConstantData[0]),
		(C.UINT)(len(pConstantData)/4),
	))
	return
}

// GetRasterStatus returns information describing the raster of the monitor on
// which the swap chain is presented.
func (obj Device) GetRasterStatus(
	iSwapChain uint,
) (
	pRasterStatus RASTER_STATUS,
	err Error,
) {
	var cpRasterStatus C.D3DRASTER_STATUS
	err = toErr(C.IDirect3DDevice9GetRasterStatus(
		obj.handle,
		(C.UINT)(iSwapChain),
		&cpRasterStatus,
	))
	pRasterStatus.fromC(&cpRasterStatus)
	return
}

// GetRenderState retrieves a render-state value for a device.
func (obj Device) GetRenderState(
	State RENDERSTATETYPE,
) (
	pValue uint32,
	err Error,
) {
	var cpValue C.DWORD
	err = toErr(C.IDirect3DDevice9GetRenderState(
		obj.handle,
		(C.D3DRENDERSTATETYPE)(State),
		&cpValue,
	))
	pValue = (uint32)(cpValue)
	return
}

// GetRenderTarget retrieves a render-target surface.
func (obj Device) GetRenderTarget(
	RenderTargetIndex uint32,
) (
	ppRenderTarget Surface,
	err Error,
) {
	var cppRenderTarget *C.IDirect3DSurface9
	err = toErr(C.IDirect3DDevice9GetRenderTarget(
		obj.handle,
		(C.DWORD)(RenderTargetIndex),
		&cppRenderTarget,
	))
	resource := Resource{
		(*C.IDirect3DResource9)(unsafe.Pointer(cppRenderTarget)),
	}
	ppRenderTarget = Surface{resource, cppRenderTarget}
	return
}

// GetRenderTargetData copies the render-target data from device memory to
// system memory.
func (obj Device) GetRenderTargetData(
	pRenderTarget Surface,
	pDestSurface Surface,
) (
	err Error,
) {
	err = toErr(C.IDirect3DDevice9GetRenderTargetData(
		obj.handle,
		pRenderTarget.handle,
		pDestSurface.handle,
	))
	return
}

// GetSamplerState return the sampler state value.
func (obj Device) GetSamplerState(
	Sampler uint32,
	Type SAMPLERSTATETYPE,
) (
	pValue uint32,
	err Error,
) {
	var cpValue C.DWORD
	err = toErr(C.IDirect3DDevice9GetSamplerState(
		obj.handle,
		(C.DWORD)(Sampler),
		(C.D3DSAMPLERSTATETYPE)(Type),
		&cpValue,
	))
	pValue = (uint32)(cpValue)
	return
}

// GetScissorRect returns the scissor rectangle.
func (obj Device) GetScissorRect() (pRect RECT, err Error) {
	var cpRect C.RECT
	err = toErr(C.IDirect3DDevice9GetScissorRect(obj.handle, &cpRect))
	pRect.fromC(&cpRect)
	return
}

// GetSoftwareVertexProcessing returns true if software vertex processing is
// set. Otherwise, it returns false.
func (obj Device) GetSoftwareVertexProcessing() bool {
	return C.IDirect3DDevice9GetSoftwareVertexProcessing(obj.handle) != 0
}

// GetStreamSource retrieves a vertex buffer bound to the specified data stream.
func (obj Device) GetStreamSource(
	StreamNumber uint,
) (
	ppStreamData VertexBuffer,
	pOffsetInBytes uint,
	pStride uint,
	err Error,
) {
	var cppStreamData *C.IDirect3DVertexBuffer9
	var cpOffsetInBytes C.UINT
	var cpStride C.UINT
	err = toErr(C.IDirect3DDevice9GetStreamSource(
		obj.handle,
		(C.UINT)(StreamNumber),
		&cppStreamData,
		&cpOffsetInBytes,
		&cpStride,
	))
	resource := Resource{
		(*C.IDirect3DResource9)(unsafe.Pointer(cppStreamData)),
	}
	ppStreamData = VertexBuffer{resource, cppStreamData}
	pOffsetInBytes = (uint)(cpOffsetInBytes)
	pStride = (uint)(cpStride)
	return
}

// GetStreamSourceFreq returns the stream source frequency divider value.
func (obj Device) GetStreamSourceFreq(
	StreamNumber uint,
) (
	pDivider uint,
	err Error,
) {
	var cpDivider C.UINT
	err = toErr(C.IDirect3DDevice9GetStreamSourceFreq(
		obj.handle,
		(C.UINT)(StreamNumber),
		&cpDivider,
	))
	pDivider = (uint)(cpDivider)
	return
}

// GetSwapChain returns a pointer to a swap chain.
func (obj Device) GetSwapChain(
	iSwapChain uint,
) (
	ppSwapChain SwapChain,
	err Error,
) {
	var cppSwapChain *C.IDirect3DSwapChain9
	err = toErr(C.IDirect3DDevice9GetSwapChain(
		obj.handle,
		(C.UINT)(iSwapChain),
		&cppSwapChain,
	))
	ppSwapChain = SwapChain{cppSwapChain}
	return
}

// GetTexture retrieves a texture assigned to a stage for a device.
func (obj Device) GetTexture(Stage uint32) (ppTexture BaseTexture, err Error) {
	var cppTexture *C.IDirect3DBaseTexture9
	err = toErr(C.IDirect3DDevice9GetTexture(
		obj.handle,
		(C.DWORD)(Stage),
		&cppTexture,
	))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(cppTexture))}
	ppTexture = BaseTexture{resource, cppTexture}
	return
}

// GetTextureStageState retrieves a state value for an assigned texture.
func (obj Device) GetTextureStageState(
	Stage uint32,
	Type TEXTURESTAGESTATETYPE,
) (
	pValue uint32,
	err Error,
) {
	var cpValue C.DWORD
	err = toErr(C.IDirect3DDevice9GetTextureStageState(
		obj.handle,
		(C.DWORD)(Stage),
		(C.D3DTEXTURESTAGESTATETYPE)(Type),
		&cpValue,
	))
	pValue = (uint32)(cpValue)
	return
}

// GetTransform retrieves a matrix describing a transformation state.
func (obj Device) GetTransform(
	State TRANSFORMSTATETYPE,
) (
	pMatrix MATRIX,
	err Error,
) {
	var cpMatrix C.D3DMATRIX
	err = toErr(C.IDirect3DDevice9GetTransform(
		obj.handle,
		(C.D3DTRANSFORMSTATETYPE)(State),
		&cpMatrix,
	))
	pMatrix.fromC(&cpMatrix)
	return
}

// GetVertexDeclaration returns a vertex shader declaration.
func (obj Device) GetVertexDeclaration() (ppDecl VertexDeclaration, err Error) {
	var cppDecl *C.IDirect3DVertexDeclaration9
	err = toErr(C.IDirect3DDevice9GetVertexDeclaration(obj.handle, &cppDecl))
	ppDecl = VertexDeclaration{cppDecl}
	return
}

// GetVertexShader retrieves the currently set vertex shader.
func (obj Device) GetVertexShader() (ppShader VertexShader, err Error) {
	var cppShader *C.IDirect3DVertexShader9
	err = toErr(C.IDirect3DDevice9GetVertexShader(obj.handle, &cppShader))
	ppShader = VertexShader{cppShader}
	return
}

// GetVertexShaderConstantB returns a Boolean shader constant. The given bool
// slice is filled with the constant data.
func (obj Device) GetVertexShaderConstantB(
	StartRegister uint,
	pConstantData []bool,
) (
	err Error,
) {
	cpConstantData := make([]C.BOOL, len(pConstantData))
	err = toErr(C.IDirect3DDevice9GetVertexShaderConstantB(
		obj.handle,
		(C.UINT)(StartRegister),
		&cpConstantData[0],
		(C.UINT)(len(pConstantData)),
	))
	for i := range cpConstantData {
		pConstantData[i] = cpConstantData[i] != 0
	}
	return
}

// GetVertexShaderConstantF returns a floating-point shader constant. The given
// float slice is filled with the constant data. The length of the slice must be
// a multiple of four.
func (obj Device) GetVertexShaderConstantF(
	StartRegister uint,
	pConstantData []float32,
) (
	err Error,
) {
	err = toErr(C.IDirect3DDevice9GetVertexShaderConstantF(
		obj.handle,
		(C.UINT)(StartRegister),
		(*C.float)(&pConstantData[0]),
		(C.UINT)(len(pConstantData)/4),
	))
	return
}

// GetVertexShaderConstantI returns an integer shader constant. The given int
// slice is filled with the constant data. The length of the slice must be a
// multiple of four.
func (obj Device) GetVertexShaderConstantI(
	StartRegister uint,
	pConstantData []int32,
) (
	err Error,
) {
	err = toErr(C.IDirect3DDevice9GetVertexShaderConstantI(
		obj.handle,
		(C.UINT)(StartRegister),
		(*C.int)(&pConstantData[0]),
		(C.UINT)(len(pConstantData)/4),
	))
	return
}

// GetViewport retrieves the viewport parameters currently set for the device.
func (obj Device) GetViewport() (pViewport VIEWPORT, err Error) {
	var cpViewport C.D3DVIEWPORT9
	err = toErr(C.IDirect3DDevice9GetViewport(obj.handle, &cpViewport))
	pViewport.fromC(&cpViewport)
	return
}

// LightEnable enables or disables a set of lighting parameters within a device.
func (obj Device) LightEnable(LightIndex uint32, bEnable bool) (err Error) {
	err = toErr(C.IDirect3DDevice9LightEnable(
		obj.handle,
		(C.DWORD)(LightIndex),
		toBOOL(bEnable),
	))
	return
}

// MultiplyTransform multiplies a device's world, view, or projection matrices
// by a specified matrix.
func (obj Device) MultiplyTransform(
	State TRANSFORMSTATETYPE,
	pMatrix MATRIX,
) (
	err Error,
) {
	cpMatrix := pMatrix.toC()
	err = toErr(C.IDirect3DDevice9MultiplyTransform(
		obj.handle,
		(C.D3DTRANSFORMSTATETYPE)(State),
		&cpMatrix,
	))
	return
}

// Present presents the contents of the next buffer in the sequence of back
// buffers owned by the device.
func (obj Device) Present(
	pSourceRect *RECT,
	pDestRect *RECT,
	hDestWindowOverride unsafe.Pointer,
	pDirtyRegion *RGNDATA,
) (
	err Error,
) {
	if pSourceRect == nil && pDestRect == nil && pDirtyRegion == nil {
		err = toErr(C.IDirect3DDevice9Present(
			obj.handle,
			nil,
			nil,
			(C.HWND)(hDestWindowOverride),
			nil,
		))
		return
	}

	var csourceRect C.RECT
	var cpSourceRect *C.RECT
	if pSourceRect != nil {
		csourceRect = pSourceRect.toC()
		cpSourceRect = &csourceRect
	}

	var cdestRect C.RECT
	var cpDestRect *C.RECT
	if pDestRect != nil {
		cdestRect = pDestRect.toC()
		cpDestRect = &cdestRect
	}

	var cdirtyRegion C.RGNDATA
	var cpDirtyRegion *C.RGNDATA
	if pDirtyRegion != nil {
		cdirtyRegion = pDirtyRegion.toC()
		cpDirtyRegion = &cdirtyRegion
	}

	err = toErr(C.IDirect3DDevice9Present(
		obj.handle,
		cpSourceRect,
		cpDestRect,
		(C.HWND)(hDestWindowOverride),
		cpDirtyRegion,
	))
	return
}

// ProcessVertices applies the vertex processing defined by the vertex shader to
// the set of input data streams, generating a single stream of interleaved
// vertex data to the destination vertex buffer.
func (obj Device) ProcessVertices(
	SrcStartIndex uint,
	DestIndex uint,
	VertexCount uint,
	pDestBuffer VertexBuffer,
	pVertexDecl VertexDeclaration,
	Flags uint32,
) (
	err Error,
) {
	err = toErr(C.IDirect3DDevice9ProcessVertices(
		obj.handle,
		(C.UINT)(SrcStartIndex),
		(C.UINT)(DestIndex),
		(C.UINT)(VertexCount),
		pDestBuffer.handle,
		pVertexDecl.handle,
		(C.DWORD)(Flags),
	))
	return
}

// Reset resets the type, size, and format of the swap chain.
func (obj Device) Reset(
	inpPresentationParameters PRESENT_PARAMETERS,
) (
	outpPresentationParameters PRESENT_PARAMETERS,
	err Error,
) {
	cpPresentationParameters := inpPresentationParameters.toC()
	err = toErr(C.IDirect3DDevice9Reset(obj.handle, &cpPresentationParameters))
	outpPresentationParameters.fromC(&cpPresentationParameters)
	return
}

// SetClipPlane sets the coefficients of a user-defined clipping plane for the
// device.
func (obj Device) SetClipPlane(Index uint32, pPlane [4]float32) (err Error) {
	err = toErr(C.IDirect3DDevice9SetClipPlane(
		obj.handle,
		C.DWORD(Index),
		(*C.float)(&pPlane[0]),
	))
	return
}

// SetClipStatus sets the clip status.
func (obj Device) SetClipStatus(pClipStatus CLIPSTATUS) (err Error) {
	cpClipStatus := pClipStatus.toC()
	err = toErr(C.IDirect3DDevice9SetClipStatus(obj.handle, &cpClipStatus))
	return
}

// SetCurrentTexturePalette sets the current texture palette.
func (obj Device) SetCurrentTexturePalette(PaletteNumber uint) (err Error) {
	err = toErr(C.IDirect3DDevice9SetCurrentTexturePalette(
		obj.handle,
		(C.UINT)(PaletteNumber),
	))
	return
}

// SetCursorPosition sets the cursor position and update options.
func (obj Device) SetCursorPosition(X int, Y int, Flags uint32) {
	C.IDirect3DDevice9SetCursorPosition(
		obj.handle,
		(C.INT)(X),
		(C.INT)(Y),
		(C.DWORD)(Flags),
	)
}

// SetCursorProperties sets properties for the cursor.
func (obj Device) SetCursorProperties(
	XHotSpot uint,
	YHotSpot uint,
	pCursorBitmap Surface,
) (
	err Error,
) {
	err = toErr(C.IDirect3DDevice9SetCursorProperties(
		obj.handle,
		(C.UINT)(XHotSpot),
		(C.UINT)(YHotSpot),
		pCursorBitmap.handle,
	))
	return
}

// SetDepthStencilSurface sets the depth stencil surface.
func (obj Device) SetDepthStencilSurface(pNewZStencil Surface) (err Error) {
	err = toErr(C.IDirect3DDevice9SetDepthStencilSurface(
		obj.handle,
		pNewZStencil.handle,
	))
	return
}

// SetDialogBoxMode allows the use of GDI dialog boxes in full-screen mode
// applications.
func (obj Device) SetDialogBoxMode(bEnableDialogs bool) (err Error) {
	err = toErr(C.IDirect3DDevice9SetDialogBoxMode(
		obj.handle,
		toBOOL(bEnableDialogs),
	))
	return
}

// SetFVF sets the current vertex stream declaration.
func (obj Device) SetFVF(FVF uint32) (err Error) {
	err = toErr(C.IDirect3DDevice9SetFVF(obj.handle, (C.DWORD)(FVF)))
	return
}

// SetGammaRamp sets the gamma correction ramp for the implicit swap chain. This
// method will affect the entire screen (not just the active window if you are
// running in windowed mode).
func (obj Device) SetGammaRamp(iSwapChain uint, Flags uint32, pRamp GAMMARAMP) {
	cpRamp := pRamp.toC()
	C.IDirect3DDevice9SetGammaRamp(obj.handle,
		(C.UINT)(iSwapChain),
		(C.DWORD)(Flags),
		&cpRamp,
	)
}

// SetIndices sets index data.
func (obj Device) SetIndices(pIndexData IndexBuffer) (err Error) {
	err = toErr(C.IDirect3DDevice9SetIndices(obj.handle, pIndexData.handle))
	return
}

// SetLight assigns a set of lighting properties for this device.
func (obj Device) SetLight(Index uint32, pLight LIGHT) (err Error) {
	cpLight := pLight.toC()
	err = toErr(C.IDirect3DDevice9SetLight(
		obj.handle,
		(C.DWORD)(Index),
		&cpLight,
	))
	return
}

// SetMaterial sets the material properties for the device.
func (obj Device) SetMaterial(pMaterial MATERIAL) (err Error) {
	cpMaterial := pMaterial.toC()
	err = toErr(C.IDirect3DDevice9SetMaterial(obj.handle, &cpMaterial))
	return
}

// SetNPatchMode enables or disables N-patches.
// nSegments specifies the number of subdivision segments. If the number of
// segments is less than 1.0, N-patches are disabled. The default value is 0.0.
func (obj Device) SetNPatchMode(nSegments float32) (err Error) {
	err = toErr(C.IDirect3DDevice9SetNPatchMode(
		obj.handle,
		(C.float)(nSegments),
	))
	return
}

// SetPaletteEntries sets palette entries.
func (obj Device) SetPaletteEntries(
	PaletteNumber uint,
	pEntries [256]PALETTEENTRY,
) (
	err Error,
) {
	var cpEntries [256]C.PALETTEENTRY
	for i := range pEntries {
		cpEntries[i] = pEntries[i].toC()
	}
	err = toErr(C.IDirect3DDevice9SetPaletteEntries(
		obj.handle,
		(C.UINT)(PaletteNumber),
		&cpEntries[0],
	))
	return
}

// SetPixelShader sets the current pixel shader to a previously created pixel
// shader.
func (obj Device) SetPixelShader(pShader PixelShader) (err Error) {
	err = toErr(C.IDirect3DDevice9SetPixelShader(obj.handle, pShader.handle))
	return
}

// SetPixelShaderConstantB sets a Boolean shader constant.
func (obj Device) SetPixelShaderConstantB(
	StartRegister uint,
	pConstantData []bool,
) (
	err Error,
) {
	data := make([]C.BOOL, len(pConstantData))
	for i := range data {
		data[i] = toBOOL(pConstantData[i])
	}
	err = toErr(C.IDirect3DDevice9SetPixelShaderConstantB(
		obj.handle,
		C.UINT(StartRegister),
		(*C.BOOL)(&data[0]),
		(C.UINT)(len(pConstantData)),
	))
	return
}

// SetPixelShaderConstantF sets a floating-point shader constant. The length of
// the slice must be a multiple of four.
func (obj Device) SetPixelShaderConstantF(
	StartRegister uint,
	pConstantData []float32,
) (
	err Error,
) {
	err = toErr(C.IDirect3DDevice9SetPixelShaderConstantF(
		obj.handle,
		C.UINT(StartRegister),
		(*C.float)(&pConstantData[0]),
		(C.UINT)((len(pConstantData)+3)/4),
	))
	return
}

// SetPixelShaderConstantI sets an integer shader constant. The length of the
// slice must be a multiple of four.
func (obj Device) SetPixelShaderConstantI(
	StartRegister uint,
	pConstantData []int32,
) (
	err Error,
) {
	err = toErr(C.IDirect3DDevice9SetPixelShaderConstantI(
		obj.handle,
		C.UINT(StartRegister),
		(*C.int)(&pConstantData[0]),
		(C.UINT)((len(pConstantData)+3)/4),
	))
	return
}

// SetRenderState sets a single device render-state parameter.
func (obj Device) SetRenderState(
	State RENDERSTATETYPE,
	Value uint32,
) (
	err Error,
) {
	err = toErr(C.IDirect3DDevice9SetRenderState(
		obj.handle,
		(C.D3DRENDERSTATETYPE)(State),
		(C.DWORD)(Value),
	))
	return
}

func (obj Device) SetRenderStateFloat(
	State RENDERSTATETYPE,
	Value float32,
) (
	err Error,
) {
	err = toErr(C.IDirect3DDevice9SetRenderState(
		obj.handle,
		(C.D3DRENDERSTATETYPE)(State),
		*((*C.DWORD)(unsafe.Pointer(&Value))),
	))
	return
}

func (obj Device) SetRenderStateBool(
	State RENDERSTATETYPE,
	Value bool,
) (
	err Error,
) {
	if Value {
		err = toErr(C.IDirect3DDevice9SetRenderState(
			obj.handle,
			(C.D3DRENDERSTATETYPE)(State),
			1,
		))
	} else {
		err = toErr(C.IDirect3DDevice9SetRenderState(
			obj.handle,
			(C.D3DRENDERSTATETYPE)(State),
			0,
		))
	}
	return
}

// SetRenderTarget sets a new color buffer for the device.
func (obj Device) SetRenderTarget(
	RenderTargetIndex uint32,
	pRenderTarget Surface,
) (
	err Error,
) {
	err = toErr(C.IDirect3DDevice9SetRenderTarget(
		obj.handle,
		(C.DWORD)(RenderTargetIndex),
		pRenderTarget.handle,
	))
	return
}

// SetSamplerState sets the sampler state value.
func (obj Device) SetSamplerState(
	Sampler uint32,
	Type SAMPLERSTATETYPE,
	Value uint32,
) (
	err Error,
) {
	err = toErr(C.IDirect3DDevice9SetSamplerState(
		obj.handle,
		(C.DWORD)(Sampler),
		(C.D3DSAMPLERSTATETYPE)(Type),
		(C.DWORD)(Value),
	))
	return
}

// SetScissorRect sets the scissor rectangle.
func (obj Device) SetScissorRect(pRect RECT) (err Error) {
	cpRect := pRect.toC()
	err = toErr(C.IDirect3DDevice9SetScissorRect(obj.handle, &cpRect))
	return
}

// SetSoftwareVertexProcessing can be used to switch between software and
// hardware vertex processing.
func (obj Device) SetSoftwareVertexProcessing(bSoftware bool) (err Error) {
	err = toErr(C.IDirect3DDevice9SetSoftwareVertexProcessing(
		obj.handle,
		toBOOL(bSoftware),
	))
	return
}

// SetStreamSource binds a vertex buffer to a device data stream.
func (obj Device) SetStreamSource(
	StreamNumber uint,
	pStreamData VertexBuffer,
	OffsetInBytes uint,
	Stride uint,
) (
	err Error,
) {
	err = toErr(C.IDirect3DDevice9SetStreamSource(
		obj.handle,
		(C.UINT)(StreamNumber),
		pStreamData.handle,
		(C.UINT)(OffsetInBytes),
		(C.UINT)(Stride),
	))
	return
}

// SetStreamSourceFreq sets the stream source frequency divider value. This may
// be used to draw several instances of geometry.
func (obj Device) SetStreamSourceFreq(
	StreamNumber uint,
	FrequencyParameter uint,
) (
	err Error,
) {
	err = toErr(C.IDirect3DDevice9SetStreamSourceFreq(
		obj.handle,
		(C.UINT)(StreamNumber),
		(C.UINT)(FrequencyParameter),
	))
	return
}

// SetTexture assigns a texture to a stage for a device.
func (obj Device) SetTexture(Sampler uint32, pTexture BaseTexture) (err Error) {
	err = toErr(C.IDirect3DDevice9SetTexture(
		obj.handle,
		(C.DWORD)(Sampler),
		pTexture.handle,
	))
	return
}

// SetTextureStageState sets the state value for the currently assigned texture.
func (obj Device) SetTextureStageState(
	Stage uint32,
	Type TEXTURESTAGESTATETYPE,
	Value uint32,
) (
	err Error,
) {
	err = toErr(C.IDirect3DDevice9SetTextureStageState(
		obj.handle,
		(C.DWORD)(Stage),
		(C.D3DTEXTURESTAGESTATETYPE)(Type),
		(C.DWORD)(Value),
	))
	return
}

// SetTransform sets a single device transformation-related state.
func (obj Device) SetTransform(
	State TRANSFORMSTATETYPE,
	pMatrix MATRIX,
) (
	err Error,
) {
	cpMatrix := pMatrix.toC()
	err = toErr(C.IDirect3DDevice9SetTransform(
		obj.handle,
		(C.D3DTRANSFORMSTATETYPE)(State),
		&cpMatrix,
	))
	return
}

// SetVertexDeclaration sets a vertex declaration.
func (obj Device) SetVertexDeclaration(pDecl VertexDeclaration) (err Error) {
	err = toErr(C.IDirect3DDevice9SetVertexDeclaration(
		obj.handle,
		pDecl.handle,
	))
	return
}

// SetVertexShader sets the vertex shader.
func (obj Device) SetVertexShader(pShader VertexShader) (err Error) {
	err = toErr(C.IDirect3DDevice9SetVertexShader(obj.handle, pShader.handle))
	return
}

// SetVertexShaderConstantB sets a Boolean vertex shader constant.
func (obj Device) SetVertexShaderConstantB(
	StartRegister uint,
	pConstantData []bool,
) (
	err Error,
) {
	data := make([]C.BOOL, len(pConstantData))
	for i := range data {
		data[i] = toBOOL(pConstantData[i])
	}
	err = toErr(C.IDirect3DDevice9SetVertexShaderConstantB(
		obj.handle,
		C.UINT(StartRegister),
		(*C.BOOL)(&data[0]),
		(C.UINT)(len(pConstantData)),
	))
	return
}

// SetVertexShaderConstantF sets a floating-point vertex shader constant. The
// length of the slice must be a multiple of four.
func (obj Device) SetVertexShaderConstantF(
	StartRegister uint,
	pConstantData []float32,
) (
	err Error,
) {
	err = toErr(C.IDirect3DDevice9SetVertexShaderConstantF(
		obj.handle,
		C.UINT(StartRegister),
		(*C.float)(&pConstantData[0]),
		(C.UINT)((len(pConstantData)+3)/4),
	))
	return
}

// SetVertexShaderConstantI sets an integer vertex shader constant. The length
// of the slice must be a multiple of four.
func (obj Device) SetVertexShaderConstantI(
	StartRegister uint,
	pConstantData []int32,
) (
	err Error,
) {
	err = toErr(C.IDirect3DDevice9SetVertexShaderConstantI(
		obj.handle,
		C.UINT(StartRegister),
		(*C.int)(&pConstantData[0]),
		(C.UINT)((len(pConstantData)+3)/4),
	))
	return
}

// SetViewport sets the viewport parameters for the device.
func (obj Device) SetViewport(pViewport VIEWPORT) (err Error) {
	cpViewport := pViewport.toC()
	err = toErr(C.IDirect3DDevice9SetViewport(obj.handle, &cpViewport))
	return
}

// ShowCursor displays or hides the cursor.
func (obj Device) ShowCursor(bShow bool) bool {
	return C.IDirect3DDevice9ShowCursor(obj.handle, toBOOL(bShow)) != 0
}

// StretchRect copies the contents of the source rectangle to the destination
// rectangle. The source rectangle can be stretched and filtered by the copy.
// This function is often used to change the aspect ratio of a video stream.
func (obj Device) StretchRect(
	pSourceSurface Surface,
	pSourceRect *RECT,
	pDestSurface Surface,
	pDestRect *RECT,
	Filter TEXTUREFILTERTYPE,
) (
	err Error,
) {
	if pSourceRect == nil && pDestRect == nil {
		err = toErr(C.IDirect3DDevice9StretchRect(
			obj.handle,
			pSourceSurface.handle,
			nil,
			pDestSurface.handle,
			nil,
			(C.D3DTEXTUREFILTERTYPE)(Filter),
		))
		return
	}

	var csourceRect C.RECT
	var cpSourceRect *C.RECT
	if pSourceRect != nil {
		csourceRect = pSourceRect.toC()
		cpSourceRect = &csourceRect
	}

	var cdestRect C.RECT
	var cpDestRect *C.RECT
	if pDestRect != nil {
		cdestRect = pDestRect.toC()
		cpDestRect = &cdestRect
	}

	err = toErr(C.IDirect3DDevice9StretchRect(
		obj.handle,
		pSourceSurface.handle,
		cpSourceRect,
		pDestSurface.handle,
		cpDestRect,
		(C.D3DTEXTUREFILTERTYPE)(Filter),
	))
	return
}

// TestCooperativeLevel reports the current cooperative-level status of the
// Direct3D device for a windowed or full-screen application.
func (obj Device) TestCooperativeLevel() (err Error) {
	err = toErr(C.IDirect3DDevice9TestCooperativeLevel(obj.handle))
	return
}

// UpdateSurface copies rectangular subsets of pixels from one surface to another.
func (obj Device) UpdateSurface(
	pSourceSurface Surface,
	pSourceRect *RECT,
	pDestinationSurface Surface,
	pDestinationPoint *POINT,
) (
	err Error,
) {
	if pSourceRect == nil && pDestinationPoint == nil {
		err = toErr(C.IDirect3DDevice9UpdateSurface(
			obj.handle,
			pSourceSurface.handle,
			nil,
			pDestinationSurface.handle,
			nil,
		))
		return
	}

	var csourceRect C.RECT
	var cpSourceRect *C.RECT
	if pSourceRect != nil {
		csourceRect = pSourceRect.toC()
		cpSourceRect = &csourceRect
	}

	var cdestinationPoint C.POINT
	var cpDestinationPoint *C.POINT
	if pDestinationPoint != nil {
		cdestinationPoint = pDestinationPoint.toC()
		cpDestinationPoint = &cdestinationPoint
	}

	err = toErr(C.IDirect3DDevice9UpdateSurface(
		obj.handle,
		pSourceSurface.handle,
		cpSourceRect,
		pDestinationSurface.handle,
		cpDestinationPoint,
	))
	return
}

// UpdateTexture updates the dirty portions of a texture.
func (obj Device) UpdateTexture(
	pSourceTexture BaseTexture,
	pDestinationTexture BaseTexture,
) (
	err Error,
) {
	err = toErr(C.IDirect3DDevice9UpdateTexture(
		obj.handle,
		pSourceTexture.handle,
		pDestinationTexture.handle,
	))
	return
}

// ValidateDevice reports the device's ability to render the current
// texture-blending operations and arguments in a single pass.
func (obj Device) ValidateDevice() (pNumPasses uint32, err Error) {
	var cpNumPasses C.DWORD
	err = toErr(C.IDirect3DDevice9ValidateDevice(obj.handle, &cpNumPasses))
	pNumPasses = (uint32)(cpNumPasses)
	return
}
