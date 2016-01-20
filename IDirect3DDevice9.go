package d3d9

/*
#include "d3d9wrapper.h"

HRESULT IDirect3DDevice9BeginScene(IDirect3DDevice9* obj) {
	return obj->lpVtbl->BeginScene(obj);
}

HRESULT IDirect3DDevice9BeginStateBlock(IDirect3DDevice9* obj) {
	return obj->lpVtbl->BeginStateBlock(obj);
}

HRESULT IDirect3DDevice9Clear(IDirect3DDevice9* obj, DWORD Count, D3DRECT* pRects, DWORD Flags, D3DCOLOR Color, float Z, DWORD Stencil) {
	return obj->lpVtbl->Clear(obj, Count, pRects, Flags, Color, Z, Stencil);
}

HRESULT IDirect3DDevice9ColorFill(IDirect3DDevice9* obj, IDirect3DSurface9* pSurface, RECT* pRect, D3DCOLOR color) {
	return obj->lpVtbl->ColorFill(obj, pSurface, pRect, color);
}

HRESULT IDirect3DDevice9CreateAdditionalSwapChain(IDirect3DDevice9* obj, D3DPRESENT_PARAMETERS* pPresentationParameters, IDirect3DSwapChain9** ppSwapChain) {
	return obj->lpVtbl->CreateAdditionalSwapChain(obj, pPresentationParameters, ppSwapChain);
}

HRESULT IDirect3DDevice9CreateCubeTexture(IDirect3DDevice9* obj, UINT EdgeLength, UINT Levels, DWORD Usage, D3DFORMAT Format, D3DPOOL Pool, IDirect3DCubeTexture9** ppCubeTexture, HANDLE* pSharedHandle) {
	return obj->lpVtbl->CreateCubeTexture(obj, EdgeLength, Levels, Usage, Format, Pool, ppCubeTexture, pSharedHandle);
}

HRESULT IDirect3DDevice9CreateDepthStencilSurface(IDirect3DDevice9* obj, UINT Width, UINT Height, D3DFORMAT Format, D3DMULTISAMPLE_TYPE MultiSample, DWORD MultisampleQuality, BOOL Discard, IDirect3DSurface9** ppSurface, HANDLE* pSharedHandle) {
	return obj->lpVtbl->CreateDepthStencilSurface(obj, Width, Height, Format, MultiSample, MultisampleQuality, Discard, ppSurface, pSharedHandle);
}

HRESULT IDirect3DDevice9CreateIndexBuffer(IDirect3DDevice9* obj, UINT Length, DWORD Usage, D3DFORMAT Format, D3DPOOL Pool, IDirect3DIndexBuffer9** ppIndexBuffer, HANDLE* pSharedHandle) {
	return obj->lpVtbl->CreateIndexBuffer(obj, Length, Usage, Format, Pool, ppIndexBuffer, pSharedHandle);
}

HRESULT IDirect3DDevice9CreateOffscreenPlainSurface(IDirect3DDevice9* obj, UINT Width, UINT Height, D3DFORMAT Format, D3DPOOL Pool, IDirect3DSurface9** ppSurface, HANDLE* pSharedHandle) {
	return obj->lpVtbl->CreateOffscreenPlainSurface(obj, Width, Height, Format, Pool, ppSurface, pSharedHandle);
}

HRESULT IDirect3DDevice9CreatePixelShader(IDirect3DDevice9* obj, DWORD* pFunction, IDirect3DPixelShader9** ppShader) {
	return obj->lpVtbl->CreatePixelShader(obj, pFunction, ppShader);
}

HRESULT IDirect3DDevice9CreateQuery(IDirect3DDevice9* obj, D3DQUERYTYPE Type, IDirect3DQuery9** ppQuery) {
	return obj->lpVtbl->CreateQuery(obj, Type, ppQuery);
}

HRESULT IDirect3DDevice9CreateRenderTarget(IDirect3DDevice9* obj, UINT Width, UINT Height, D3DFORMAT Format, D3DMULTISAMPLE_TYPE MultiSample, DWORD MultisampleQuality, BOOL Lockable, IDirect3DSurface9** ppSurface, HANDLE* pSharedHandle) {
	return obj->lpVtbl->CreateRenderTarget(obj, Width, Height, Format, MultiSample, MultisampleQuality, Lockable, ppSurface, pSharedHandle);
}

HRESULT IDirect3DDevice9CreateStateBlock(IDirect3DDevice9* obj, D3DSTATEBLOCKTYPE Type, IDirect3DStateBlock9** ppSB) {
	return obj->lpVtbl->CreateStateBlock(obj, Type, ppSB);
}

HRESULT IDirect3DDevice9CreateTexture(IDirect3DDevice9* obj, UINT Width, UINT Height, UINT Levels, DWORD Usage, D3DFORMAT Format, D3DPOOL Pool, IDirect3DTexture9** ppTexture, HANDLE* pSharedHandle) {
	return obj->lpVtbl->CreateTexture(obj, Width, Height, Levels, Usage, Format, Pool, ppTexture, pSharedHandle);
}

HRESULT IDirect3DDevice9CreateVertexBuffer(IDirect3DDevice9* obj, UINT Length, DWORD Usage, DWORD FVF, D3DPOOL Pool, IDirect3DVertexBuffer9** ppVertexBuffer, HANDLE* pSharedHandle) {
	return obj->lpVtbl->CreateVertexBuffer(obj, Length, Usage, FVF, Pool, ppVertexBuffer, pSharedHandle);
}

HRESULT IDirect3DDevice9CreateVertexDeclaration(IDirect3DDevice9* obj, D3DVERTEXELEMENT9* pVertexElements, IDirect3DVertexDeclaration9** ppDecl) {
	return obj->lpVtbl->CreateVertexDeclaration(obj, pVertexElements, ppDecl);
}

HRESULT IDirect3DDevice9CreateVertexShader(IDirect3DDevice9* obj, DWORD* pFunction, IDirect3DVertexShader9** ppShader) {
	return obj->lpVtbl->CreateVertexShader(obj, pFunction, ppShader);
}

HRESULT IDirect3DDevice9CreateVolumeTexture(IDirect3DDevice9* obj, UINT Width, UINT Height, UINT Depth, UINT Levels, DWORD Usage, D3DFORMAT Format, D3DPOOL Pool, IDirect3DVolumeTexture9** ppVolumeTexture, HANDLE* pSharedHandle) {
	return obj->lpVtbl->CreateVolumeTexture(obj, Width, Height, Depth, Levels, Usage, Format, Pool, ppVolumeTexture, pSharedHandle);
}

HRESULT IDirect3DDevice9DeletePatch(IDirect3DDevice9* obj, UINT Handle) {
	return obj->lpVtbl->DeletePatch(obj, Handle);
}

HRESULT IDirect3DDevice9DrawIndexedPrimitive(IDirect3DDevice9* obj, D3DPRIMITIVETYPE Type, INT BaseVertexIndex, UINT MinIndex, UINT NumVertices, UINT StartIndex, UINT PrimitiveCount) {
	return obj->lpVtbl->DrawIndexedPrimitive(obj, Type, BaseVertexIndex, MinIndex, NumVertices, StartIndex, PrimitiveCount);
}

HRESULT IDirect3DDevice9DrawIndexedPrimitiveUP(IDirect3DDevice9* obj, D3DPRIMITIVETYPE PrimitiveType, UINT MinVertexIndex, UINT NumVertices, UINT PrimitiveCount, void* pIndexData, D3DFORMAT IndexDataFormat, void* pVertexStreamZeroData, UINT VertexStreamZeroStride) {
	return obj->lpVtbl->DrawIndexedPrimitiveUP(obj, PrimitiveType, MinVertexIndex, NumVertices, PrimitiveCount, pIndexData, IndexDataFormat, pVertexStreamZeroData, VertexStreamZeroStride);
}

HRESULT IDirect3DDevice9DrawPrimitive(IDirect3DDevice9* obj, D3DPRIMITIVETYPE PrimitiveType, UINT StartVertex, UINT PrimitiveCount) {
	return obj->lpVtbl->DrawPrimitive(obj, PrimitiveType, StartVertex, PrimitiveCount);
}

HRESULT IDirect3DDevice9DrawPrimitiveUP(IDirect3DDevice9* obj, D3DPRIMITIVETYPE PrimitiveType, UINT PrimitiveCount, void* pVertexStreamZeroData, UINT VertexStreamZeroStride) {
	return obj->lpVtbl->DrawPrimitiveUP(obj, PrimitiveType, PrimitiveCount, pVertexStreamZeroData, VertexStreamZeroStride);
}

HRESULT IDirect3DDevice9DrawRectPatch(IDirect3DDevice9* obj, UINT Handle, float* pNumSegs, D3DRECTPATCH_INFO* pRectPatchInfo) {
	return obj->lpVtbl->DrawRectPatch(obj, Handle, pNumSegs, pRectPatchInfo);
}

HRESULT IDirect3DDevice9DrawTriPatch(IDirect3DDevice9* obj, UINT Handle, float* pNumSegs, D3DTRIPATCH_INFO* pTriPatchInfo) {
	return obj->lpVtbl->DrawTriPatch(obj, Handle, pNumSegs, pTriPatchInfo);
}

HRESULT IDirect3DDevice9EndScene(IDirect3DDevice9* obj) {
	return obj->lpVtbl->EndScene(obj);
}

HRESULT IDirect3DDevice9EndStateBlock(IDirect3DDevice9* obj, IDirect3DStateBlock9** ppSB) {
	return obj->lpVtbl->EndStateBlock(obj, ppSB);
}

HRESULT IDirect3DDevice9EvictManagedResources(IDirect3DDevice9* obj) {
	return obj->lpVtbl->EvictManagedResources(obj);
}

UINT IDirect3DDevice9GetAvailableTextureMem(IDirect3DDevice9* obj) {
	return obj->lpVtbl->GetAvailableTextureMem(obj);
}

HRESULT IDirect3DDevice9GetBackBuffer(IDirect3DDevice9* obj, UINT iSwapChain, UINT BackBuffer, D3DBACKBUFFER_TYPE Type, IDirect3DSurface9** ppBackBuffer) {
	return obj->lpVtbl->GetBackBuffer(obj, iSwapChain, BackBuffer, Type, ppBackBuffer);
}

HRESULT IDirect3DDevice9GetClipPlane(IDirect3DDevice9* obj, DWORD Index, float* pPlane) {
	return obj->lpVtbl->GetClipPlane(obj, Index, pPlane);
}

HRESULT IDirect3DDevice9GetClipStatus(IDirect3DDevice9* obj, D3DCLIPSTATUS9* pClipStatus) {
	return obj->lpVtbl->GetClipStatus(obj, pClipStatus);
}

HRESULT IDirect3DDevice9GetCreationParameters(IDirect3DDevice9* obj, D3DDEVICE_CREATION_PARAMETERS* pParameters) {
	return obj->lpVtbl->GetCreationParameters(obj, pParameters);
}

HRESULT IDirect3DDevice9GetCurrentTexturePalette(IDirect3DDevice9* obj, UINT* pPaletteNumber) {
	return obj->lpVtbl->GetCurrentTexturePalette(obj, pPaletteNumber);
}

HRESULT IDirect3DDevice9GetDepthStencilSurface(IDirect3DDevice9* obj, IDirect3DSurface9** ppZStencilSurface) {
	return obj->lpVtbl->GetDepthStencilSurface(obj, ppZStencilSurface);
}

HRESULT IDirect3DDevice9GetDeviceCaps(IDirect3DDevice9* obj, D3DCAPS9* pCaps) {
	return obj->lpVtbl->GetDeviceCaps(obj, pCaps);
}

HRESULT IDirect3DDevice9GetDirect3D(IDirect3DDevice9* obj, IDirect3D9** ppD3D9) {
	return obj->lpVtbl->GetDirect3D(obj, ppD3D9);
}

HRESULT IDirect3DDevice9GetDisplayMode(IDirect3DDevice9* obj, UINT iSwapChain, D3DDISPLAYMODE* pMode) {
	return obj->lpVtbl->GetDisplayMode(obj, iSwapChain, pMode);
}

HRESULT IDirect3DDevice9GetFrontBufferData(IDirect3DDevice9* obj, UINT iSwapChain, IDirect3DSurface9* pDestSurface) {
	return obj->lpVtbl->GetFrontBufferData(obj, iSwapChain, pDestSurface);
}

HRESULT IDirect3DDevice9GetFVF(IDirect3DDevice9* obj, DWORD* pFVF) {
	return obj->lpVtbl->GetFVF(obj, pFVF);
}

void IDirect3DDevice9GetGammaRamp(IDirect3DDevice9* obj, UINT iSwapChain, D3DGAMMARAMP* pRamp) {
	obj->lpVtbl->GetGammaRamp(obj, iSwapChain, pRamp);
}

HRESULT IDirect3DDevice9GetIndices(IDirect3DDevice9* obj, IDirect3DIndexBuffer9** ppIndexData) {
	return obj->lpVtbl->GetIndices(obj, ppIndexData);
}

HRESULT IDirect3DDevice9GetLight(IDirect3DDevice9* obj, DWORD Index, D3DLIGHT9* pLight) {
	return obj->lpVtbl->GetLight(obj, Index, pLight);
}

HRESULT IDirect3DDevice9GetLightEnable(IDirect3DDevice9* obj, DWORD Index, BOOL* pEnable) {
	return obj->lpVtbl->GetLightEnable(obj, Index, pEnable);
}

HRESULT IDirect3DDevice9GetMaterial(IDirect3DDevice9* obj, D3DMATERIAL9* pMaterial) {
	return obj->lpVtbl->GetMaterial(obj, pMaterial);
}

FLOAT IDirect3DDevice9GetNPatchMode(IDirect3DDevice9* obj) {
	return obj->lpVtbl->GetNPatchMode(obj);
}

UINT IDirect3DDevice9GetNumberOfSwapChains(IDirect3DDevice9* obj) {
	return obj->lpVtbl->GetNumberOfSwapChains(obj);
}

HRESULT IDirect3DDevice9GetPaletteEntries(IDirect3DDevice9* obj, UINT PaletteNumber, PALETTEENTRY* pEntries) {
	return obj->lpVtbl->GetPaletteEntries(obj, PaletteNumber, pEntries);
}

HRESULT IDirect3DDevice9GetPixelShader(IDirect3DDevice9* obj, IDirect3DPixelShader9** ppShader) {
	return obj->lpVtbl->GetPixelShader(obj, ppShader);
}

HRESULT IDirect3DDevice9GetPixelShaderConstantB(IDirect3DDevice9* obj, UINT StartRegister, BOOL* pConstantData, UINT BoolCount) {
	return obj->lpVtbl->GetPixelShaderConstantB(obj, StartRegister, pConstantData, BoolCount);
}

HRESULT IDirect3DDevice9GetPixelShaderConstantF(IDirect3DDevice9* obj, UINT StartRegister, float* pConstantData, UINT Vector4fCount) {
	return obj->lpVtbl->GetPixelShaderConstantF(obj, StartRegister, pConstantData, Vector4fCount);
}

HRESULT IDirect3DDevice9GetPixelShaderConstantI(IDirect3DDevice9* obj, UINT StartRegister, int* pConstantData, UINT Vector4iCount) {
	return obj->lpVtbl->GetPixelShaderConstantI(obj, StartRegister, pConstantData, Vector4iCount);
}

HRESULT IDirect3DDevice9GetRasterStatus(IDirect3DDevice9* obj, UINT iSwapChain, D3DRASTER_STATUS* pRasterStatus) {
	return obj->lpVtbl->GetRasterStatus(obj, iSwapChain, pRasterStatus);
}

HRESULT IDirect3DDevice9GetRenderState(IDirect3DDevice9* obj, D3DRENDERSTATETYPE State, DWORD* pValue) {
	return obj->lpVtbl->GetRenderState(obj, State, pValue);
}

HRESULT IDirect3DDevice9GetRenderTarget(IDirect3DDevice9* obj, DWORD RenderTargetIndex, IDirect3DSurface9** ppRenderTarget) {
	return obj->lpVtbl->GetRenderTarget(obj, RenderTargetIndex, ppRenderTarget);
}

HRESULT IDirect3DDevice9GetRenderTargetData(IDirect3DDevice9* obj, IDirect3DSurface9* pRenderTarget, IDirect3DSurface9* pDestSurface) {
	return obj->lpVtbl->GetRenderTargetData(obj, pRenderTarget, pDestSurface);
}

HRESULT IDirect3DDevice9GetSamplerState(IDirect3DDevice9* obj, DWORD Sampler, D3DSAMPLERSTATETYPE Type, DWORD* pValue) {
	return obj->lpVtbl->GetSamplerState(obj, Sampler, Type, pValue);
}

HRESULT IDirect3DDevice9GetScissorRect(IDirect3DDevice9* obj, RECT* pRect) {
	return obj->lpVtbl->GetScissorRect(obj, pRect);
}

BOOL IDirect3DDevice9GetSoftwareVertexProcessing(IDirect3DDevice9* obj) {
	return obj->lpVtbl->GetSoftwareVertexProcessing(obj);
}

HRESULT IDirect3DDevice9GetStreamSource(IDirect3DDevice9* obj, UINT StreamNumber, IDirect3DVertexBuffer9** ppStreamData, UINT* pOffsetInBytes, UINT* pStride) {
	return obj->lpVtbl->GetStreamSource(obj, StreamNumber, ppStreamData, pOffsetInBytes, pStride);
}

HRESULT IDirect3DDevice9GetStreamSourceFreq(IDirect3DDevice9* obj, UINT StreamNumber, UINT* pDivider) {
	return obj->lpVtbl->GetStreamSourceFreq(obj, StreamNumber, pDivider);
}

HRESULT IDirect3DDevice9GetSwapChain(IDirect3DDevice9* obj, UINT iSwapChain, IDirect3DSwapChain9** ppSwapChain) {
	return obj->lpVtbl->GetSwapChain(obj, iSwapChain, ppSwapChain);
}

HRESULT IDirect3DDevice9GetTexture(IDirect3DDevice9* obj, DWORD Stage, IDirect3DBaseTexture9** ppTexture) {
	return obj->lpVtbl->GetTexture(obj, Stage, ppTexture);
}

HRESULT IDirect3DDevice9GetTextureStageState(IDirect3DDevice9* obj, DWORD Stage, D3DTEXTURESTAGESTATETYPE Type, DWORD* pValue) {
	return obj->lpVtbl->GetTextureStageState(obj, Stage, Type, pValue);
}

HRESULT IDirect3DDevice9GetTransform(IDirect3DDevice9* obj, D3DTRANSFORMSTATETYPE State, D3DMATRIX* pMatrix) {
	return obj->lpVtbl->GetTransform(obj, State, pMatrix);
}

HRESULT IDirect3DDevice9GetVertexDeclaration(IDirect3DDevice9* obj, IDirect3DVertexDeclaration9** ppDecl) {
	return obj->lpVtbl->GetVertexDeclaration(obj, ppDecl);
}

HRESULT IDirect3DDevice9GetVertexShader(IDirect3DDevice9* obj, IDirect3DVertexShader9** ppShader) {
	return obj->lpVtbl->GetVertexShader(obj, ppShader);
}

HRESULT IDirect3DDevice9GetVertexShaderConstantB(IDirect3DDevice9* obj, UINT StartRegister, BOOL* pConstantData, UINT BoolCount) {
	return obj->lpVtbl->GetVertexShaderConstantB(obj, StartRegister, pConstantData, BoolCount);
}

HRESULT IDirect3DDevice9GetVertexShaderConstantF(IDirect3DDevice9* obj, UINT StartRegister, float* pConstantData, UINT Vector4fCount) {
	return obj->lpVtbl->GetVertexShaderConstantF(obj, StartRegister, pConstantData, Vector4fCount);
}

HRESULT IDirect3DDevice9GetVertexShaderConstantI(IDirect3DDevice9* obj, UINT StartRegister, int* pConstantData, UINT Vector4iCount) {
	return obj->lpVtbl->GetVertexShaderConstantI(obj, StartRegister, pConstantData, Vector4iCount);
}

HRESULT IDirect3DDevice9GetViewport(IDirect3DDevice9* obj, D3DVIEWPORT9* pViewport) {
	return obj->lpVtbl->GetViewport(obj, pViewport);
}

HRESULT IDirect3DDevice9LightEnable(IDirect3DDevice9* obj, DWORD LightIndex, BOOL bEnable) {
	return obj->lpVtbl->LightEnable(obj, LightIndex, bEnable);
}

HRESULT IDirect3DDevice9MultiplyTransform(IDirect3DDevice9* obj, D3DTRANSFORMSTATETYPE State, D3DMATRIX* pMatrix) {
	return obj->lpVtbl->MultiplyTransform(obj, State, pMatrix);
}

HRESULT IDirect3DDevice9Present(IDirect3DDevice9* obj, RECT* pSourceRect, RECT* pDestRect, HWND hDestWindowOverride, RGNDATA* pDirtyRegion) {
	return obj->lpVtbl->Present(obj, pSourceRect, pDestRect, hDestWindowOverride, pDirtyRegion);
}

HRESULT IDirect3DDevice9ProcessVertices(IDirect3DDevice9* obj, UINT SrcStartIndex, UINT DestIndex, UINT VertexCount, IDirect3DVertexBuffer9* pDestBuffer, IDirect3DVertexDeclaration9* pVertexDecl, DWORD Flags) {
	return obj->lpVtbl->ProcessVertices(obj, SrcStartIndex, DestIndex, VertexCount, pDestBuffer, pVertexDecl, Flags);
}

HRESULT IDirect3DDevice9Reset(IDirect3DDevice9* obj, D3DPRESENT_PARAMETERS* pPresentationParameters) {
	return obj->lpVtbl->Reset(obj, pPresentationParameters);
}

HRESULT IDirect3DDevice9SetClipPlane(IDirect3DDevice9* obj, DWORD Index, float* pPlane) {
	return obj->lpVtbl->SetClipPlane(obj, Index, pPlane);
}

HRESULT IDirect3DDevice9SetClipStatus(IDirect3DDevice9* obj, D3DCLIPSTATUS9* pClipStatus) {
	return obj->lpVtbl->SetClipStatus(obj, pClipStatus);
}

HRESULT IDirect3DDevice9SetCurrentTexturePalette(IDirect3DDevice9* obj, UINT PaletteNumber) {
	return obj->lpVtbl->SetCurrentTexturePalette(obj, PaletteNumber);
}

void IDirect3DDevice9SetCursorPosition(IDirect3DDevice9* obj, INT X, INT Y, DWORD Flags) {
	obj->lpVtbl->SetCursorPosition(obj, X, Y, Flags);
}

HRESULT IDirect3DDevice9SetCursorProperties(IDirect3DDevice9* obj, UINT XHotSpot, UINT YHotSpot, IDirect3DSurface9* pCursorBitmap) {
	return obj->lpVtbl->SetCursorProperties(obj, XHotSpot, YHotSpot, pCursorBitmap);
}

HRESULT IDirect3DDevice9SetDepthStencilSurface(IDirect3DDevice9* obj, IDirect3DSurface9* pNewZStencil) {
	return obj->lpVtbl->SetDepthStencilSurface(obj, pNewZStencil);
}

HRESULT IDirect3DDevice9SetDialogBoxMode(IDirect3DDevice9* obj, BOOL bEnableDialogs) {
	return obj->lpVtbl->SetDialogBoxMode(obj, bEnableDialogs);
}

HRESULT IDirect3DDevice9SetFVF(IDirect3DDevice9* obj, DWORD FVF) {
	return obj->lpVtbl->SetFVF(obj, FVF);
}

void IDirect3DDevice9SetGammaRamp(IDirect3DDevice9* obj, UINT iSwapChain, DWORD Flags, D3DGAMMARAMP* pRamp) {
	obj->lpVtbl->SetGammaRamp(obj, iSwapChain, Flags, pRamp);
}

HRESULT IDirect3DDevice9SetIndices(IDirect3DDevice9* obj, IDirect3DIndexBuffer9* pIndexData) {
	return obj->lpVtbl->SetIndices(obj, pIndexData);
}

HRESULT IDirect3DDevice9SetLight(IDirect3DDevice9* obj, DWORD Index, D3DLIGHT9* pLight) {
	return obj->lpVtbl->SetLight(obj, Index, pLight);
}

HRESULT IDirect3DDevice9SetMaterial(IDirect3DDevice9* obj, D3DMATERIAL9* pMaterial) {
	return obj->lpVtbl->SetMaterial(obj, pMaterial);
}

HRESULT IDirect3DDevice9SetNPatchMode(IDirect3DDevice9* obj, float nSegments) {
	return obj->lpVtbl->SetNPatchMode(obj, nSegments);
}

HRESULT IDirect3DDevice9SetPaletteEntries(IDirect3DDevice9* obj, UINT PaletteNumber, PALETTEENTRY* pEntries) {
	return obj->lpVtbl->SetPaletteEntries(obj, PaletteNumber, pEntries);
}

HRESULT IDirect3DDevice9SetPixelShader(IDirect3DDevice9* obj, IDirect3DPixelShader9* pShader) {
	return obj->lpVtbl->SetPixelShader(obj, pShader);
}

HRESULT IDirect3DDevice9SetPixelShaderConstantB(IDirect3DDevice9* obj, UINT StartRegister, BOOL* pConstantData, UINT BoolCount) {
	return obj->lpVtbl->SetPixelShaderConstantB(obj, StartRegister, pConstantData, BoolCount);
}

HRESULT IDirect3DDevice9SetPixelShaderConstantF(IDirect3DDevice9* obj, UINT StartRegister, float* pConstantData, UINT Vector4fCount) {
	return obj->lpVtbl->SetPixelShaderConstantF(obj, StartRegister, pConstantData, Vector4fCount);
}

HRESULT IDirect3DDevice9SetPixelShaderConstantI(IDirect3DDevice9* obj, UINT StartRegister, int* pConstantData, UINT Vector4iCount) {
	return obj->lpVtbl->SetPixelShaderConstantI(obj, StartRegister, pConstantData, Vector4iCount);
}

HRESULT IDirect3DDevice9SetRenderState(IDirect3DDevice9* obj, D3DRENDERSTATETYPE State, DWORD Value) {
	return obj->lpVtbl->SetRenderState(obj, State, Value);
}

HRESULT IDirect3DDevice9SetRenderTarget(IDirect3DDevice9* obj, DWORD RenderTargetIndex, IDirect3DSurface9* pRenderTarget) {
	return obj->lpVtbl->SetRenderTarget(obj, RenderTargetIndex, pRenderTarget);
}

HRESULT IDirect3DDevice9SetSamplerState(IDirect3DDevice9* obj, DWORD Sampler, D3DSAMPLERSTATETYPE Type, DWORD Value) {
	return obj->lpVtbl->SetSamplerState(obj, Sampler, Type, Value);
}

HRESULT IDirect3DDevice9SetScissorRect(IDirect3DDevice9* obj, RECT* pRect) {
	return obj->lpVtbl->SetScissorRect(obj, pRect);
}

HRESULT IDirect3DDevice9SetSoftwareVertexProcessing(IDirect3DDevice9* obj, BOOL bSoftware) {
	return obj->lpVtbl->SetSoftwareVertexProcessing(obj, bSoftware);
}

HRESULT IDirect3DDevice9SetStreamSource(IDirect3DDevice9* obj, UINT StreamNumber, IDirect3DVertexBuffer9* pStreamData, UINT OffsetInBytes, UINT Stride) {
	return obj->lpVtbl->SetStreamSource(obj, StreamNumber, pStreamData, OffsetInBytes, Stride);
}

HRESULT IDirect3DDevice9SetStreamSourceFreq(IDirect3DDevice9* obj, UINT StreamNumber, UINT FrequencyParameter) {
	return obj->lpVtbl->SetStreamSourceFreq(obj, StreamNumber, FrequencyParameter);
}

HRESULT IDirect3DDevice9SetTexture(IDirect3DDevice9* obj, DWORD Sampler, IDirect3DBaseTexture9* pTexture) {
	return obj->lpVtbl->SetTexture(obj, Sampler, pTexture);
}

HRESULT IDirect3DDevice9SetTextureStageState(IDirect3DDevice9* obj, DWORD Stage, D3DTEXTURESTAGESTATETYPE Type, DWORD Value) {
	return obj->lpVtbl->SetTextureStageState(obj, Stage, Type, Value);
}

HRESULT IDirect3DDevice9SetTransform(IDirect3DDevice9* obj, D3DTRANSFORMSTATETYPE State, D3DMATRIX* pMatrix) {
	return obj->lpVtbl->SetTransform(obj, State, pMatrix);
}

HRESULT IDirect3DDevice9SetVertexDeclaration(IDirect3DDevice9* obj, IDirect3DVertexDeclaration9* pDecl) {
	return obj->lpVtbl->SetVertexDeclaration(obj, pDecl);
}

HRESULT IDirect3DDevice9SetVertexShader(IDirect3DDevice9* obj, IDirect3DVertexShader9* pShader) {
	return obj->lpVtbl->SetVertexShader(obj, pShader);
}

HRESULT IDirect3DDevice9SetVertexShaderConstantB(IDirect3DDevice9* obj, UINT StartRegister, BOOL* pConstantData, UINT BoolCount) {
	return obj->lpVtbl->SetVertexShaderConstantB(obj, StartRegister, pConstantData, BoolCount);
}

HRESULT IDirect3DDevice9SetVertexShaderConstantF(IDirect3DDevice9* obj, UINT StartRegister, float* pConstantData, UINT Vector4fCount) {
	return obj->lpVtbl->SetVertexShaderConstantF(obj, StartRegister, pConstantData, Vector4fCount);
}

HRESULT IDirect3DDevice9SetVertexShaderConstantI(IDirect3DDevice9* obj, UINT StartRegister, int* pConstantData, UINT Vector4iCount) {
	return obj->lpVtbl->SetVertexShaderConstantI(obj, StartRegister, pConstantData, Vector4iCount);
}

HRESULT IDirect3DDevice9SetViewport(IDirect3DDevice9* obj, D3DVIEWPORT9* pViewport) {
	return obj->lpVtbl->SetViewport(obj, pViewport);
}

BOOL IDirect3DDevice9ShowCursor(IDirect3DDevice9* obj, BOOL bShow) {
	return obj->lpVtbl->ShowCursor(obj, bShow);
}

HRESULT IDirect3DDevice9StretchRect(IDirect3DDevice9* obj, IDirect3DSurface9* pSourceSurface, RECT* pSourceRect, IDirect3DSurface9* pDestSurface, RECT* pDestRect, D3DTEXTUREFILTERTYPE Filter) {
	return obj->lpVtbl->StretchRect(obj, pSourceSurface, pSourceRect, pDestSurface, pDestRect, Filter);
}

HRESULT IDirect3DDevice9TestCooperativeLevel(IDirect3DDevice9* obj) {
	return obj->lpVtbl->TestCooperativeLevel(obj);
}

HRESULT IDirect3DDevice9UpdateSurface(IDirect3DDevice9* obj, IDirect3DSurface9* pSourceSurface, RECT* pSourceRect, IDirect3DSurface9* pDestinationSurface, POINT* pDestinationPoint) {
	return obj->lpVtbl->UpdateSurface(obj, pSourceSurface, pSourceRect, pDestinationSurface, pDestinationPoint);
}

HRESULT IDirect3DDevice9UpdateTexture(IDirect3DDevice9* obj, IDirect3DBaseTexture9* pSourceTexture, IDirect3DBaseTexture9* pDestinationTexture) {
	return obj->lpVtbl->UpdateTexture(obj, pSourceTexture, pDestinationTexture);
}

HRESULT IDirect3DDevice9ValidateDevice(IDirect3DDevice9* obj, DWORD* pNumPasses) {
	return obj->lpVtbl->ValidateDevice(obj, pNumPasses);
}

void IDirect3DDevice9Release(IDirect3DDevice9* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"
import "unsafe"

type Device struct {
	handle *C.IDirect3DDevice9
}

func (obj Device) Release() {
	C.IDirect3DDevice9Release(obj.handle)
}

func (obj Device) BeginScene() (err error) {
	err = toErr(C.IDirect3DDevice9BeginScene(obj.handle))
	return
}

func (obj Device) BeginStateBlock() (err error) {
	err = toErr(C.IDirect3DDevice9BeginStateBlock(obj.handle))
	return
}

func (obj Device) Clear(pRects []D3DRECT, Flags uint32, Color COLOR, Z float32, Stencil uint32) (err error) {
	if len(pRects) == 0 {
		err = toErr(C.IDirect3DDevice9Clear(obj.handle, 0, nil, (C.DWORD)(Flags), Color.toC(), (C.float)(Z), (C.DWORD)(Stencil)))
	} else {
		c_pRects := make([]C.D3DRECT, len(pRects))
		for i := range c_pRects {
			c_pRects[i] = pRects[i].toC()
		}
		err = toErr(C.IDirect3DDevice9Clear(obj.handle, (C.DWORD)(len(pRects)), &c_pRects[0], (C.DWORD)(Flags), Color.toC(), (C.float)(Z), (C.DWORD)(Stencil)))
	}
	return
}

func (obj Device) ColorFill(pSurface Surface, pRect RECT, color COLOR) (err error) {
	c_pRect := pRect.toC()
	err = toErr(C.IDirect3DDevice9ColorFill(obj.handle, pSurface.handle, &c_pRect, (C.D3DCOLOR)(color)))
	return
}

func (obj Device) CreateAdditionalSwapChain(inpPresentationParameters PRESENT_PARAMETERS) (ppSwapChain SwapChain, outpPresentationParameters PRESENT_PARAMETERS, err error) {
	c_pPresentationParameters := inpPresentationParameters.toC()
	var c_ppSwapChain *C.IDirect3DSwapChain9
	err = toErr(C.IDirect3DDevice9CreateAdditionalSwapChain(obj.handle, &c_pPresentationParameters, &c_ppSwapChain))
	outpPresentationParameters.fromC(&c_pPresentationParameters)
	ppSwapChain = SwapChain{c_ppSwapChain}
	return
}

func (obj Device) CreateCubeTexture(EdgeLength uint, Levels uint, Usage uint32, Format FORMAT, Pool POOL, pSharedHandle unsafe.Pointer) (ppCubeTexture CubeTexture, err error) {
	var c_ppCubeTexture *C.IDirect3DCubeTexture9
	err = toErr(C.IDirect3DDevice9CreateCubeTexture(obj.handle, (C.UINT)(EdgeLength), (C.UINT)(Levels), (C.DWORD)(Usage), (C.D3DFORMAT)(Format), (C.D3DPOOL)(Pool), &c_ppCubeTexture, (*C.HANDLE)(pSharedHandle)))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(c_ppCubeTexture))}
	base := BaseTexture{resource, (*C.IDirect3DBaseTexture9)(unsafe.Pointer(c_ppCubeTexture))}
	ppCubeTexture = CubeTexture{base, c_ppCubeTexture}
	return
}

func (obj Device) CreateDepthStencilSurface(Width uint, Height uint, Format FORMAT, MultiSample MULTISAMPLE_TYPE, MultisampleQuality uint32, Discard bool, pSharedHandle unsafe.Pointer) (ppSurface Surface, err error) {
	var c_ppSurface *C.IDirect3DSurface9
	err = toErr(C.IDirect3DDevice9CreateDepthStencilSurface(obj.handle, (C.UINT)(Width), (C.UINT)(Height), (C.D3DFORMAT)(Format), (C.D3DMULTISAMPLE_TYPE)(MultiSample), (C.DWORD)(MultisampleQuality), toBOOL(Discard), &c_ppSurface, (*C.HANDLE)(pSharedHandle)))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(c_ppSurface))}
	ppSurface = Surface{resource, c_ppSurface}
	return
}

func (obj Device) CreateIndexBuffer(Length uint, Usage uint32, Format FORMAT, Pool POOL, pSharedHandle unsafe.Pointer) (ppIndexBuffer IndexBuffer, err error) {
	var c_ppIndexBuffer *C.IDirect3DIndexBuffer9
	err = toErr(C.IDirect3DDevice9CreateIndexBuffer(obj.handle, (C.UINT)(Length), (C.DWORD)(Usage), (C.D3DFORMAT)(Format), (C.D3DPOOL)(Pool), &c_ppIndexBuffer, (*C.HANDLE)(pSharedHandle)))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(c_ppIndexBuffer))}
	ppIndexBuffer = IndexBuffer{resource, c_ppIndexBuffer}
	return
}

func (obj Device) CreateOffscreenPlainSurface(Width uint, Height uint, Format FORMAT, Pool POOL, pSharedHandle unsafe.Pointer) (ppSurface Surface, err error) {
	var c_ppSurface *C.IDirect3DSurface9
	err = toErr(C.IDirect3DDevice9CreateOffscreenPlainSurface(obj.handle, (C.UINT)(Width), (C.UINT)(Height), (C.D3DFORMAT)(Format), (C.D3DPOOL)(Pool), &c_ppSurface, (*C.HANDLE)(pSharedHandle)))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(c_ppSurface))}
	ppSurface = Surface{resource, c_ppSurface}
	return
}

func (obj Device) CreatePixelShader(pFunction unsafe.Pointer) (ppShader PixelShader, err error) {
	var c_ppShader *C.IDirect3DPixelShader9
	err = toErr(C.IDirect3DDevice9CreatePixelShader(obj.handle, (*C.DWORD)(pFunction), &c_ppShader))
	ppShader = PixelShader{c_ppShader}
	return
}

func (obj Device) CreatePixelShaderFromBytes(pFunction []byte) (ppShader PixelShader, err error) {
	var c_ppShader *C.IDirect3DPixelShader9
	err = toErr(C.IDirect3DDevice9CreatePixelShader(obj.handle, (*C.DWORD)(unsafe.Pointer(&pFunction[0])), &c_ppShader))
	ppShader = PixelShader{c_ppShader}
	return
}

func (obj Device) CreateQuery(Type QUERYTYPE) (ppQuery Query, err error) {
	var c_ppQuery *C.IDirect3DQuery9
	err = toErr(C.IDirect3DDevice9CreateQuery(obj.handle, (C.D3DQUERYTYPE)(Type), &c_ppQuery))
	ppQuery = Query{c_ppQuery}
	return
}

func (obj Device) CreateRenderTarget(Width uint, Height uint, Format FORMAT, MultiSample MULTISAMPLE_TYPE, MultisampleQuality uint32, Lockable bool, pSharedHandle unsafe.Pointer) (ppSurface Surface, err error) {
	var c_ppSurface *C.IDirect3DSurface9
	err = toErr(C.IDirect3DDevice9CreateRenderTarget(obj.handle, (C.UINT)(Width), (C.UINT)(Height), (C.D3DFORMAT)(Format), (C.D3DMULTISAMPLE_TYPE)(MultiSample), (C.DWORD)(MultisampleQuality), toBOOL(Lockable), &c_ppSurface, (*C.HANDLE)(pSharedHandle)))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(c_ppSurface))}
	ppSurface = Surface{resource, c_ppSurface}
	return
}

func (obj Device) CreateStateBlock(Type STATEBLOCKTYPE) (ppSB StateBlock, err error) {
	var c_ppSB *C.IDirect3DStateBlock9
	err = toErr(C.IDirect3DDevice9CreateStateBlock(obj.handle, (C.D3DSTATEBLOCKTYPE)(Type), &c_ppSB))
	ppSB = StateBlock{c_ppSB}
	return
}

func (obj Device) CreateTexture(Width uint, Height uint, Levels uint, Usage uint32, Format FORMAT, Pool POOL, pSharedHandle unsafe.Pointer) (ppTexture Texture, err error) {
	var c_ppTexture *C.IDirect3DTexture9
	err = toErr(C.IDirect3DDevice9CreateTexture(obj.handle, (C.UINT)(Width), (C.UINT)(Height), (C.UINT)(Levels), (C.DWORD)(Usage), (C.D3DFORMAT)(Format), (C.D3DPOOL)(Pool), &c_ppTexture, (*C.HANDLE)(pSharedHandle)))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(c_ppTexture))}
	base := BaseTexture{resource, (*C.IDirect3DBaseTexture9)(unsafe.Pointer(c_ppTexture))}
	ppTexture = Texture{base, c_ppTexture}
	return
}

func (obj Device) CreateVertexBuffer(Length uint, Usage uint32, FVF uint32, Pool POOL, pSharedHandle unsafe.Pointer) (ppVertexBuffer VertexBuffer, err error) {
	var c_ppVertexBuffer *C.IDirect3DVertexBuffer9
	err = toErr(C.IDirect3DDevice9CreateVertexBuffer(obj.handle, (C.UINT)(Length), (C.DWORD)(Usage), (C.DWORD)(FVF), (C.D3DPOOL)(Pool), &c_ppVertexBuffer, (*C.HANDLE)(pSharedHandle)))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(c_ppVertexBuffer))}
	ppVertexBuffer = VertexBuffer{resource, c_ppVertexBuffer}
	return
}

func (obj Device) CreateVertexDeclaration(pVertexElements []VERTEXELEMENT) (ppDecl VertexDeclaration, err error) {
	c_pVertexElements := make([]C.D3DVERTEXELEMENT9, len(pVertexElements))
	for i := range c_pVertexElements {
		c_pVertexElements[i] = pVertexElements[i].toC()
	}
	var c_ppDecl *C.IDirect3DVertexDeclaration9
	err = toErr(C.IDirect3DDevice9CreateVertexDeclaration(obj.handle, &c_pVertexElements[0], &c_ppDecl))
	ppDecl = VertexDeclaration{c_ppDecl}
	return
}

func DeclEnd() VERTEXELEMENT {
	return VERTEXELEMENT{0xFF, 0, DECLTYPE_UNUSED, 0, 0, 0}
}

func (obj Device) CreateVertexShader(pFunction unsafe.Pointer) (ppShader VertexShader, err error) {
	var c_ppShader *C.IDirect3DVertexShader9
	err = toErr(C.IDirect3DDevice9CreateVertexShader(obj.handle, (*C.DWORD)(pFunction), &c_ppShader))
	ppShader = VertexShader{c_ppShader}
	return
}

func (obj Device) CreateVertexShaderFromBytes(pFunction []byte) (ppShader VertexShader, err error) {
	var c_ppShader *C.IDirect3DVertexShader9
	err = toErr(C.IDirect3DDevice9CreateVertexShader(obj.handle, (*C.DWORD)(unsafe.Pointer(&pFunction[0])), &c_ppShader))
	ppShader = VertexShader{c_ppShader}
	return
}

func (obj Device) CreateVolumeTexture(Width uint, Height uint, Depth uint, Levels uint, Usage uint32, Format FORMAT, Pool POOL, pSharedHandle unsafe.Pointer) (ppVolumeTexture VolumeTexture, err error) {
	var c_ppVolumeTexture *C.IDirect3DVolumeTexture9
	err = toErr(C.IDirect3DDevice9CreateVolumeTexture(obj.handle, (C.UINT)(Width), (C.UINT)(Height), (C.UINT)(Depth), (C.UINT)(Levels), (C.DWORD)(Usage), (C.D3DFORMAT)(Format), (C.D3DPOOL)(Pool), &c_ppVolumeTexture, (*C.HANDLE)(pSharedHandle)))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(c_ppVolumeTexture))}
	base := BaseTexture{resource, (*C.IDirect3DBaseTexture9)(unsafe.Pointer(c_ppVolumeTexture))}
	ppVolumeTexture = VolumeTexture{base, c_ppVolumeTexture}
	return
}

func (obj Device) DeletePatch(Handle uint) (err error) {
	err = toErr(C.IDirect3DDevice9DeletePatch(obj.handle, (C.UINT)(Handle)))
	return
}

func (obj Device) DrawIndexedPrimitive(Type PRIMITIVETYPE, BaseVertexIndex int, MinIndex uint, NumVertices uint, StartIndex uint, PrimitiveCount uint) (err error) {
	err = toErr(C.IDirect3DDevice9DrawIndexedPrimitive(obj.handle, (C.D3DPRIMITIVETYPE)(Type), (C.INT)(BaseVertexIndex), (C.UINT)(MinIndex), (C.UINT)(NumVertices), (C.UINT)(StartIndex), (C.UINT)(PrimitiveCount)))
	return
}

func (obj Device) DrawIndexedPrimitiveUP(PrimitiveType PRIMITIVETYPE, MinVertexIndex uint, NumVertices uint, PrimitiveCount uint, pIndexData unsafe.Pointer, IndexDataFormat FORMAT, pVertexStreamZeroData unsafe.Pointer, VertexStreamZeroStride uint) (err error) {
	err = toErr(C.IDirect3DDevice9DrawIndexedPrimitiveUP(obj.handle, (C.D3DPRIMITIVETYPE)(PrimitiveType), (C.UINT)(MinVertexIndex), (C.UINT)(NumVertices), (C.UINT)(PrimitiveCount), pIndexData, (C.D3DFORMAT)(IndexDataFormat), pVertexStreamZeroData, (C.UINT)(VertexStreamZeroStride)))
	return
}

func (obj Device) DrawPrimitive(PrimitiveType PRIMITIVETYPE, StartVertex uint, PrimitiveCount uint) (err error) {
	err = toErr(C.IDirect3DDevice9DrawPrimitive(obj.handle, (C.D3DPRIMITIVETYPE)(PrimitiveType), (C.UINT)(StartVertex), (C.UINT)(PrimitiveCount)))
	return
}

func (obj Device) DrawPrimitiveUP(PrimitiveType PRIMITIVETYPE, PrimitiveCount uint, pVertexStreamZeroData unsafe.Pointer, VertexStreamZeroStride uint) (err error) {
	err = toErr(C.IDirect3DDevice9DrawPrimitiveUP(obj.handle, (C.D3DPRIMITIVETYPE)(PrimitiveType), (C.UINT)(PrimitiveCount), pVertexStreamZeroData, (C.UINT)(VertexStreamZeroStride)))
	return
}

func (obj Device) DrawRectPatch(Handle uint, pNumSegs [4]float32, pRectPatchInfo RECTPATCH_INFO) (err error) {
	c_pRectPatchInfo := pRectPatchInfo.toC()
	err = toErr(C.IDirect3DDevice9DrawRectPatch(obj.handle, C.UINT(Handle), (*C.float)(&pNumSegs[0]), &c_pRectPatchInfo))
	return
}

func (obj Device) DrawTriPatch(Handle uint, pNumSegs [3]float32, pTriPatchInfo TRIPATCH_INFO) (err error) {
	c_pTriPatchInfo := pTriPatchInfo.toC()
	err = toErr(C.IDirect3DDevice9DrawTriPatch(obj.handle, C.UINT(Handle), (*C.float)(&pNumSegs[0]), &c_pTriPatchInfo))
	return
}

func (obj Device) EndScene() (err error) {
	err = toErr(C.IDirect3DDevice9EndScene(obj.handle))
	return
}

func (obj Device) EndStateBlock() (ppSB StateBlock, err error) {
	var c_ppSB *C.IDirect3DStateBlock9
	err = toErr(C.IDirect3DDevice9EndStateBlock(obj.handle, &c_ppSB))
	ppSB = StateBlock{c_ppSB}
	return
}

func (obj Device) EvictManagedResources() (err error) {
	err = toErr(C.IDirect3DDevice9EvictManagedResources(obj.handle))
	return
}

func (obj Device) GetAvailableTextureMem() uint {
	return (uint)(C.IDirect3DDevice9GetAvailableTextureMem(obj.handle))
}

func (obj Device) GetBackBuffer(iSwapChain uint, BackBuffer uint, Type BACKBUFFER_TYPE) (ppBackBuffer Surface, err error) {
	var c_ppBackBuffer *C.IDirect3DSurface9
	err = toErr(C.IDirect3DDevice9GetBackBuffer(obj.handle, (C.UINT)(iSwapChain), (C.UINT)(BackBuffer), (C.D3DBACKBUFFER_TYPE)(Type), &c_ppBackBuffer))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(c_ppBackBuffer))}
	ppBackBuffer = Surface{resource, c_ppBackBuffer}
	return
}

func (obj Device) GetClipPlane(Index uint32) (pPlane [4]float32, err error) {
	err = toErr(C.IDirect3DDevice9GetClipPlane(obj.handle, (C.DWORD)(Index), (*C.float)(&pPlane[0])))
	return
}

func (obj Device) GetClipStatus() (pClipStatus CLIPSTATUS, err error) {
	var c_pClipStatus C.D3DCLIPSTATUS9
	err = toErr(C.IDirect3DDevice9GetClipStatus(obj.handle, &c_pClipStatus))
	pClipStatus.fromC(&c_pClipStatus)
	return
}

func (obj Device) GetCreationParameters() (pParameters DEVICE_CREATION_PARAMETERS, err error) {
	var c_pParameters C.D3DDEVICE_CREATION_PARAMETERS
	err = toErr(C.IDirect3DDevice9GetCreationParameters(obj.handle, &c_pParameters))
	pParameters.fromC(&c_pParameters)
	return
}

func (obj Device) GetCurrentTexturePalette() (pPaletteNumber uint, err error) {
	var c_pPaletteNumber C.UINT
	err = toErr(C.IDirect3DDevice9GetCurrentTexturePalette(obj.handle, &c_pPaletteNumber))
	pPaletteNumber = (uint)(c_pPaletteNumber)
	return
}

func (obj Device) GetDepthStencilSurface() (ppZStencilSurface Surface, err error) {
	var c_ppZStencilSurface *C.IDirect3DSurface9
	err = toErr(C.IDirect3DDevice9GetDepthStencilSurface(obj.handle, &c_ppZStencilSurface))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(c_ppZStencilSurface))}
	ppZStencilSurface = Surface{resource, c_ppZStencilSurface}
	return
}

func (obj Device) GetDeviceCaps() (pCaps CAPS, err error) {
	var c_pCaps C.D3DCAPS9
	err = toErr(C.IDirect3DDevice9GetDeviceCaps(obj.handle, &c_pCaps))
	pCaps.fromC(&c_pCaps)
	return
}

func (obj Device) GetDirect3D() (ppD3D9 Direct3D, err error) {
	var c_ppD3D9 *C.IDirect3D9
	err = toErr(C.IDirect3DDevice9GetDirect3D(obj.handle, &c_ppD3D9))
	ppD3D9 = Direct3D{c_ppD3D9}
	return
}

func (obj Device) GetDisplayMode(iSwapChain uint) (pMode DISPLAYMODE, err error) {
	var c_pMode C.D3DDISPLAYMODE
	err = toErr(C.IDirect3DDevice9GetDisplayMode(obj.handle, (C.UINT)(iSwapChain), &c_pMode))
	pMode.fromC(&c_pMode)
	return
}

func (obj Device) GetFrontBufferData(iSwapChain uint, pDestSurface Surface) (err error) {
	err = toErr(C.IDirect3DDevice9GetFrontBufferData(obj.handle, (C.UINT)(iSwapChain), pDestSurface.handle))
	return
}

func (obj Device) GetFVF() (pFVF uint32, err error) {
	var c_pFVF C.DWORD
	err = toErr(C.IDirect3DDevice9GetFVF(obj.handle, &c_pFVF))
	pFVF = (uint32)(c_pFVF)
	return
}

func (obj Device) GetGammaRamp(iSwapChain uint) (pRamp GAMMARAMP) {
	var c_pRamp C.D3DGAMMARAMP
	C.IDirect3DDevice9GetGammaRamp(obj.handle, (C.UINT)(iSwapChain), &c_pRamp)
	pRamp.fromC(&c_pRamp)
	return
}

func (obj Device) GetIndices() (ppIndexData IndexBuffer, err error) {
	var c_ppIndexData *C.IDirect3DIndexBuffer9
	err = toErr(C.IDirect3DDevice9GetIndices(obj.handle, &c_ppIndexData))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(c_ppIndexData))}
	ppIndexData = IndexBuffer{resource, c_ppIndexData}
	return
}

func (obj Device) GetLight(Index uint32) (pLight LIGHT, err error) {
	var c_pLight C.D3DLIGHT9
	err = toErr(C.IDirect3DDevice9GetLight(obj.handle, (C.DWORD)(Index), &c_pLight))
	pLight.fromC(&c_pLight)
	return
}

func (obj Device) GetLightEnable(Index uint32) (pEnable bool, err error) {
	var c_pEnable C.BOOL
	err = toErr(C.IDirect3DDevice9GetLightEnable(obj.handle, (C.DWORD)(Index), &c_pEnable))
	pEnable = c_pEnable != 0
	return
}

func (obj Device) GetMaterial() (pMaterial MATERIAL, err error) {
	var c_pMaterial C.D3DMATERIAL9
	err = toErr(C.IDirect3DDevice9GetMaterial(obj.handle, &c_pMaterial))
	pMaterial.fromC(&c_pMaterial)
	return
}

func (obj Device) GetNPatchMode() float32 {
	return (float32)(C.IDirect3DDevice9GetNPatchMode(obj.handle))
}

func (obj Device) GetNumberOfSwapChains() uint {
	return (uint)(C.IDirect3DDevice9GetNumberOfSwapChains(obj.handle))
}

func (obj Device) GetPaletteEntries(PaletteNumber uint) (pEntries PALETTEENTRY, err error) {
	var c_pEntries C.PALETTEENTRY
	err = toErr(C.IDirect3DDevice9GetPaletteEntries(obj.handle, (C.UINT)(PaletteNumber), &c_pEntries))
	pEntries.fromC(&c_pEntries)
	return
}

func (obj Device) GetPixelShader() (ppShader PixelShader, err error) {
	var c_ppShader *C.IDirect3DPixelShader9
	err = toErr(C.IDirect3DDevice9GetPixelShader(obj.handle, &c_ppShader))
	ppShader = PixelShader{c_ppShader}
	return
}

// pConstantData will be filled with the constant data
func (obj Device) GetPixelShaderConstantB(StartRegister uint, pConstantData []bool) (err error) {
	c_pConstantData := make([]C.BOOL, len(pConstantData))
	err = toErr(C.IDirect3DDevice9GetPixelShaderConstantB(obj.handle, (C.UINT)(StartRegister), &c_pConstantData[0], (C.UINT)(len(pConstantData))))
	for i := range c_pConstantData {
		pConstantData[i] = c_pConstantData[i] != 0
	}
	return
}

// pConstantData will be filled with the constant data
func (obj Device) GetPixelShaderConstantF(StartRegister uint, pConstantData []float32) (err error) {
	err = toErr(C.IDirect3DDevice9GetPixelShaderConstantF(obj.handle, (C.UINT)(StartRegister), (*C.float)(&pConstantData[0]), (C.UINT)(len(pConstantData)/4)))
	return
}

// pConstantData will be filled with the constant data
func (obj Device) GetPixelShaderConstantI(StartRegister uint, pConstantData []int32) (err error) {
	err = toErr(C.IDirect3DDevice9GetPixelShaderConstantI(obj.handle, (C.UINT)(StartRegister), (*C.int)(&pConstantData[0]), (C.UINT)(len(pConstantData)/4)))
	return
}

func (obj Device) GetRasterStatus(iSwapChain uint) (pRasterStatus RASTER_STATUS, err error) {
	var c_pRasterStatus C.D3DRASTER_STATUS
	err = toErr(C.IDirect3DDevice9GetRasterStatus(obj.handle, (C.UINT)(iSwapChain), &c_pRasterStatus))
	pRasterStatus.fromC(&c_pRasterStatus)
	return
}

func (obj Device) GetRenderState(State RENDERSTATETYPE) (pValue uint32, err error) {
	var c_pValue C.DWORD
	err = toErr(C.IDirect3DDevice9GetRenderState(obj.handle, (C.D3DRENDERSTATETYPE)(State), &c_pValue))
	pValue = (uint32)(c_pValue)
	return
}

func (obj Device) GetRenderTarget(RenderTargetIndex uint32) (ppRenderTarget Surface, err error) {
	var c_ppRenderTarget *C.IDirect3DSurface9
	err = toErr(C.IDirect3DDevice9GetRenderTarget(obj.handle, (C.DWORD)(RenderTargetIndex), &c_ppRenderTarget))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(c_ppRenderTarget))}
	ppRenderTarget = Surface{resource, c_ppRenderTarget}
	return
}

func (obj Device) GetRenderTargetData(pRenderTarget Surface, pDestSurface Surface) (err error) {
	err = toErr(C.IDirect3DDevice9GetRenderTargetData(obj.handle, pRenderTarget.handle, pDestSurface.handle))
	return
}

func (obj Device) GetSamplerState(Sampler uint32, Type SAMPLERSTATETYPE) (pValue uint32, err error) {
	var c_pValue C.DWORD
	err = toErr(C.IDirect3DDevice9GetSamplerState(obj.handle, (C.DWORD)(Sampler), (C.D3DSAMPLERSTATETYPE)(Type), &c_pValue))
	pValue = (uint32)(c_pValue)
	return
}

func (obj Device) GetScissorRect() (pRect RECT, err error) {
	var c_pRect C.RECT
	err = toErr(C.IDirect3DDevice9GetScissorRect(obj.handle, &c_pRect))
	pRect.fromC(&c_pRect)
	return
}

func (obj Device) GetSoftwareVertexProcessing() bool {
	return C.IDirect3DDevice9GetSoftwareVertexProcessing(obj.handle) != 0
}

func (obj Device) GetStreamSource(StreamNumber uint) (ppStreamData VertexBuffer, pOffsetInBytes uint, pStride uint, err error) {
	var c_ppStreamData *C.IDirect3DVertexBuffer9
	var c_pOffsetInBytes C.UINT
	var c_pStride C.UINT
	err = toErr(C.IDirect3DDevice9GetStreamSource(obj.handle, (C.UINT)(StreamNumber), &c_ppStreamData, &c_pOffsetInBytes, &c_pStride))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(c_ppStreamData))}
	ppStreamData = VertexBuffer{resource, c_ppStreamData}
	pOffsetInBytes = (uint)(c_pOffsetInBytes)
	pStride = (uint)(c_pStride)
	return
}

func (obj Device) GetStreamSourceFreq(StreamNumber uint) (pDivider uint, err error) {
	var c_pDivider C.UINT
	err = toErr(C.IDirect3DDevice9GetStreamSourceFreq(obj.handle, (C.UINT)(StreamNumber), &c_pDivider))
	pDivider = (uint)(c_pDivider)
	return
}

func (obj Device) GetSwapChain(iSwapChain uint) (ppSwapChain SwapChain, err error) {
	var c_ppSwapChain *C.IDirect3DSwapChain9
	err = toErr(C.IDirect3DDevice9GetSwapChain(obj.handle, (C.UINT)(iSwapChain), &c_ppSwapChain))
	ppSwapChain = SwapChain{c_ppSwapChain}
	return
}

func (obj Device) GetTexture(Stage uint32) (ppTexture BaseTexture, err error) {
	var c_ppTexture *C.IDirect3DBaseTexture9
	err = toErr(C.IDirect3DDevice9GetTexture(obj.handle, (C.DWORD)(Stage), &c_ppTexture))
	resource := Resource{(*C.IDirect3DResource9)(unsafe.Pointer(c_ppTexture))}
	ppTexture = BaseTexture{resource, c_ppTexture}
	return
}

func (obj Device) GetTextureStageState(Stage uint32, Type TEXTURESTAGESTATETYPE) (pValue uint32, err error) {
	var c_pValue C.DWORD
	err = toErr(C.IDirect3DDevice9GetTextureStageState(obj.handle, (C.DWORD)(Stage), (C.D3DTEXTURESTAGESTATETYPE)(Type), &c_pValue))
	pValue = (uint32)(c_pValue)
	return
}

func (obj Device) GetTransform(State TRANSFORMSTATETYPE) (pMatrix MATRIX, err error) {
	var c_pMatrix C.D3DMATRIX
	err = toErr(C.IDirect3DDevice9GetTransform(obj.handle, (C.D3DTRANSFORMSTATETYPE)(State), &c_pMatrix))
	pMatrix.fromC(&c_pMatrix)
	return
}

func (obj Device) GetVertexDeclaration() (ppDecl VertexDeclaration, err error) {
	var c_ppDecl *C.IDirect3DVertexDeclaration9
	err = toErr(C.IDirect3DDevice9GetVertexDeclaration(obj.handle, &c_ppDecl))
	ppDecl = VertexDeclaration{c_ppDecl}
	return
}

func (obj Device) GetVertexShader() (ppShader VertexShader, err error) {
	var c_ppShader *C.IDirect3DVertexShader9
	err = toErr(C.IDirect3DDevice9GetVertexShader(obj.handle, &c_ppShader))
	ppShader = VertexShader{c_ppShader}
	return
}

// pConstantData will be filled with the constant data
func (obj Device) GetVertexShaderConstantB(StartRegister uint, pConstantData []bool) (err error) {
	c_pConstantData := make([]C.BOOL, len(pConstantData))
	err = toErr(C.IDirect3DDevice9GetVertexShaderConstantB(obj.handle, (C.UINT)(StartRegister), &c_pConstantData[0], (C.UINT)(len(pConstantData))))
	for i := range c_pConstantData {
		pConstantData[i] = c_pConstantData[i] != 0
	}
	return
}

// pConstantData will be filled with the constant data
func (obj Device) GetVertexShaderConstantF(StartRegister uint, pConstantData []float32) (err error) {
	err = toErr(C.IDirect3DDevice9GetVertexShaderConstantF(obj.handle, (C.UINT)(StartRegister), (*C.float)(&pConstantData[0]), (C.UINT)(len(pConstantData)/4)))
	return
}

// pConstantData will be filled with the constant data
func (obj Device) GetVertexShaderConstantI(StartRegister uint, pConstantData []int32) (err error) {
	err = toErr(C.IDirect3DDevice9GetVertexShaderConstantI(obj.handle, (C.UINT)(StartRegister), (*C.int)(&pConstantData[0]), (C.UINT)(len(pConstantData)/4)))
	return
}

func (obj Device) GetViewport() (pViewport VIEWPORT, err error) {
	var c_pViewport C.D3DVIEWPORT9
	err = toErr(C.IDirect3DDevice9GetViewport(obj.handle, &c_pViewport))
	pViewport.fromC(&c_pViewport)
	return
}

func (obj Device) LightEnable(LightIndex uint32, bEnable bool) (err error) {
	err = toErr(C.IDirect3DDevice9LightEnable(obj.handle, (C.DWORD)(LightIndex), toBOOL(bEnable)))
	return
}

func (obj Device) MultiplyTransform(State TRANSFORMSTATETYPE, pMatrix MATRIX) (err error) {
	c_pMatrix := pMatrix.toC()
	err = toErr(C.IDirect3DDevice9MultiplyTransform(obj.handle, (C.D3DTRANSFORMSTATETYPE)(State), &c_pMatrix))
	return
}

func (obj Device) Present(pSourceRect *RECT, pDestRect *RECT, hDestWindowOverride unsafe.Pointer, pDirtyRegion *RGNDATA) (err error) {
	if pSourceRect == nil && pDestRect == nil && pDirtyRegion == nil {
		err = toErr(C.IDirect3DDevice9Present(obj.handle, nil, nil, (C.HWND)(hDestWindowOverride), nil))
		return
	}
	// TODO what if some are nil and some are not?
	var c_sourceRect, c_destRect C.RECT
	var c_dirtyRegion C.RGNDATA

	var c_pSourceRect *C.RECT = nil
	var c_pDestRect *C.RECT = nil
	var c_pDirtyRegion *C.RGNDATA = nil

	if pSourceRect != nil {
		c_sourceRect = pSourceRect.toC()
		c_pSourceRect = &c_sourceRect
	}
	if pDestRect != nil {
		c_destRect = pDestRect.toC()
		c_pDestRect = &c_destRect
	}
	if pDirtyRegion != nil {
		c_dirtyRegion = pDirtyRegion.toC()
		c_pDirtyRegion = &c_dirtyRegion
	}

	err = toErr(C.IDirect3DDevice9Present(obj.handle, c_pSourceRect, c_pDestRect, (C.HWND)(hDestWindowOverride), c_pDirtyRegion))
	return
}

func (obj Device) ProcessVertices(SrcStartIndex uint, DestIndex uint, VertexCount uint, pDestBuffer VertexBuffer, pVertexDecl VertexDeclaration, Flags uint32) (err error) {
	err = toErr(C.IDirect3DDevice9ProcessVertices(obj.handle, (C.UINT)(SrcStartIndex), (C.UINT)(DestIndex), (C.UINT)(VertexCount), pDestBuffer.handle, pVertexDecl.handle, (C.DWORD)(Flags)))
	return
}

func (obj Device) Reset(inpPresentationParameters PRESENT_PARAMETERS) (outpPresentationParameters PRESENT_PARAMETERS, err error) {
	c_pPresentationParameters := inpPresentationParameters.toC()
	err = toErr(C.IDirect3DDevice9Reset(obj.handle, &c_pPresentationParameters))
	outpPresentationParameters.fromC(&c_pPresentationParameters)
	return
}

func (obj Device) SetClipPlane(Index uint32, pPlane [4]float32) (err error) {
	err = toErr(C.IDirect3DDevice9SetClipPlane(obj.handle, C.DWORD(Index), (*C.float)(&pPlane[0])))
	return
}

func (obj Device) SetClipStatus(pClipStatus CLIPSTATUS) (err error) {
	c_pClipStatus := pClipStatus.toC()
	err = toErr(C.IDirect3DDevice9SetClipStatus(obj.handle, &c_pClipStatus))
	return
}

func (obj Device) SetCurrentTexturePalette(PaletteNumber uint) (err error) {
	err = toErr(C.IDirect3DDevice9SetCurrentTexturePalette(obj.handle, (C.UINT)(PaletteNumber)))
	return
}

func (obj Device) SetCursorPosition(X int, Y int, Flags uint32) {
	C.IDirect3DDevice9SetCursorPosition(obj.handle, (C.INT)(X), (C.INT)(Y), (C.DWORD)(Flags))
}

func (obj Device) SetCursorProperties(XHotSpot uint, YHotSpot uint, pCursorBitmap Surface) (err error) {
	err = toErr(C.IDirect3DDevice9SetCursorProperties(obj.handle, (C.UINT)(XHotSpot), (C.UINT)(YHotSpot), pCursorBitmap.handle))
	return
}

func (obj Device) SetDepthStencilSurface(pNewZStencil Surface) (err error) {
	err = toErr(C.IDirect3DDevice9SetDepthStencilSurface(obj.handle, pNewZStencil.handle))
	return
}

func (obj Device) SetDialogBoxMode(bEnableDialogs bool) (err error) {
	err = toErr(C.IDirect3DDevice9SetDialogBoxMode(obj.handle, toBOOL(bEnableDialogs)))
	return
}

func (obj Device) SetFVF(FVF uint32) (err error) {
	err = toErr(C.IDirect3DDevice9SetFVF(obj.handle, (C.DWORD)(FVF)))
	return
}

func (obj Device) SetGammaRamp(iSwapChain uint, Flags uint32, pRamp GAMMARAMP) {
	c_pRamp := pRamp.toC()
	C.IDirect3DDevice9SetGammaRamp(obj.handle, (C.UINT)(iSwapChain), (C.DWORD)(Flags), &c_pRamp)
}

func (obj Device) SetIndices(pIndexData IndexBuffer) (err error) {
	err = toErr(C.IDirect3DDevice9SetIndices(obj.handle, pIndexData.handle))
	return
}

func (obj Device) SetLight(Index uint32, pLight LIGHT) (err error) {
	c_pLight := pLight.toC()
	err = toErr(C.IDirect3DDevice9SetLight(obj.handle, (C.DWORD)(Index), &c_pLight))
	return
}

func (obj Device) SetMaterial(pMaterial MATERIAL) (err error) {
	c_pMaterial := pMaterial.toC()
	err = toErr(C.IDirect3DDevice9SetMaterial(obj.handle, &c_pMaterial))
	return
}

func (obj Device) SetNPatchMode(nSegments float32) (err error) {
	err = toErr(C.IDirect3DDevice9SetNPatchMode(obj.handle, (C.float)(nSegments)))
	return
}

func (obj Device) SetPaletteEntries(PaletteNumber uint, pEntries PALETTEENTRY) (err error) {
	c_pEntries := pEntries.toC()
	err = toErr(C.IDirect3DDevice9SetPaletteEntries(obj.handle, (C.UINT)(PaletteNumber), &c_pEntries))
	return
}

func (obj Device) SetPixelShader(pShader PixelShader) (err error) {
	err = toErr(C.IDirect3DDevice9SetPixelShader(obj.handle, pShader.handle))
	return
}

func (obj Device) SetPixelShaderConstantB(StartRegister uint, pConstantData []bool) (err error) {
	data := make([]C.BOOL, len(pConstantData))
	for i := range data {
		data[i] = toBOOL(pConstantData[i])
	}
	err = toErr(C.IDirect3DDevice9SetPixelShaderConstantB(obj.handle, C.UINT(StartRegister), (*C.BOOL)(&data[0]), (C.UINT)(len(pConstantData))))
	return
}

func (obj Device) SetPixelShaderConstantF(StartRegister uint, pConstantData []float32) (err error) {
	err = toErr(C.IDirect3DDevice9SetPixelShaderConstantF(obj.handle, C.UINT(StartRegister), (*C.float)(&pConstantData[0]), (C.UINT)((len(pConstantData)+3)/4)))
	return
}

func (obj Device) SetPixelShaderConstantI(StartRegister uint, pConstantData []int32) (err error) {
	err = toErr(C.IDirect3DDevice9SetPixelShaderConstantI(obj.handle, C.UINT(StartRegister), (*C.int)(&pConstantData[0]), (C.UINT)((len(pConstantData)+3)/4)))
	return
}

func (obj Device) SetRenderState(State RENDERSTATETYPE, Value uint32) (err error) {
	err = toErr(C.IDirect3DDevice9SetRenderState(obj.handle, (C.D3DRENDERSTATETYPE)(State), (C.DWORD)(Value)))
	return
}

func (obj Device) SetRenderTarget(RenderTargetIndex uint32, pRenderTarget Surface) (err error) {
	err = toErr(C.IDirect3DDevice9SetRenderTarget(obj.handle, (C.DWORD)(RenderTargetIndex), pRenderTarget.handle))
	return
}

func (obj Device) SetSamplerState(Sampler uint32, Type SAMPLERSTATETYPE, Value uint32) (err error) {
	err = toErr(C.IDirect3DDevice9SetSamplerState(obj.handle, (C.DWORD)(Sampler), (C.D3DSAMPLERSTATETYPE)(Type), (C.DWORD)(Value)))
	return
}

func (obj Device) SetScissorRect(pRect RECT) (err error) {
	c_pRect := pRect.toC()
	err = toErr(C.IDirect3DDevice9SetScissorRect(obj.handle, &c_pRect))
	return
}

func (obj Device) SetSoftwareVertexProcessing(bSoftware bool) (err error) {
	err = toErr(C.IDirect3DDevice9SetSoftwareVertexProcessing(obj.handle, toBOOL(bSoftware)))
	return
}

func (obj Device) SetStreamSource(StreamNumber uint, pStreamData VertexBuffer, OffsetInBytes uint, Stride uint) (err error) {
	err = toErr(C.IDirect3DDevice9SetStreamSource(obj.handle, (C.UINT)(StreamNumber), pStreamData.handle, (C.UINT)(OffsetInBytes), (C.UINT)(Stride)))
	return
}

func (obj Device) SetStreamSourceFreq(StreamNumber uint, FrequencyParameter uint) (err error) {
	err = toErr(C.IDirect3DDevice9SetStreamSourceFreq(obj.handle, (C.UINT)(StreamNumber), (C.UINT)(FrequencyParameter)))
	return
}

func (obj Device) SetTexture(Sampler uint32, pTexture BaseTexture) (err error) {
	err = toErr(C.IDirect3DDevice9SetTexture(obj.handle, (C.DWORD)(Sampler), pTexture.handle))
	return
}

func (obj Device) SetTextureStageState(Stage uint32, Type TEXTURESTAGESTATETYPE, Value uint32) (err error) {
	err = toErr(C.IDirect3DDevice9SetTextureStageState(obj.handle, (C.DWORD)(Stage), (C.D3DTEXTURESTAGESTATETYPE)(Type), (C.DWORD)(Value)))
	return
}

func (obj Device) SetTransform(State TRANSFORMSTATETYPE, pMatrix MATRIX) (err error) {
	c_pMatrix := pMatrix.toC()
	err = toErr(C.IDirect3DDevice9SetTransform(obj.handle, (C.D3DTRANSFORMSTATETYPE)(State), &c_pMatrix))
	return
}

func (obj Device) SetVertexDeclaration(pDecl VertexDeclaration) (err error) {
	err = toErr(C.IDirect3DDevice9SetVertexDeclaration(obj.handle, pDecl.handle))
	return
}

func (obj Device) SetVertexShader(pShader VertexShader) (err error) {
	err = toErr(C.IDirect3DDevice9SetVertexShader(obj.handle, pShader.handle))
	return
}

func (obj Device) SetVertexShaderConstantB(StartRegister uint, pConstantData []bool) (err error) {
	data := make([]C.BOOL, len(pConstantData))
	for i := range data {
		data[i] = toBOOL(pConstantData[i])
	}
	err = toErr(C.IDirect3DDevice9SetVertexShaderConstantB(obj.handle, C.UINT(StartRegister), (*C.BOOL)(&data[0]), (C.UINT)(len(pConstantData))))
	return
}

func (obj Device) SetVertexShaderConstantF(StartRegister uint, pConstantData []float32) (err error) {
	err = toErr(C.IDirect3DDevice9SetVertexShaderConstantF(obj.handle, C.UINT(StartRegister), (*C.float)(&pConstantData[0]), (C.UINT)((len(pConstantData)+3)/4)))
	return
}

func (obj Device) SetVertexShaderConstantI(StartRegister uint, pConstantData []int32) (err error) {
	err = toErr(C.IDirect3DDevice9SetVertexShaderConstantI(obj.handle, C.UINT(StartRegister), (*C.int)(&pConstantData[0]), (C.UINT)((len(pConstantData)+3)/4)))
	return
}

func (obj Device) SetViewport(pViewport VIEWPORT) (err error) {
	c_pViewport := pViewport.toC()
	err = toErr(C.IDirect3DDevice9SetViewport(obj.handle, &c_pViewport))
	return
}

func (obj Device) ShowCursor(bShow bool) bool {
	return C.IDirect3DDevice9ShowCursor(obj.handle, toBOOL(bShow)) != 0
}

func (obj Device) StretchRect(pSourceSurface Surface, pSourceRect RECT, pDestSurface Surface, pDestRect RECT, Filter TEXTUREFILTERTYPE) (err error) {
	c_pSourceRect := pSourceRect.toC()
	c_pDestRect := pDestRect.toC()
	err = toErr(C.IDirect3DDevice9StretchRect(obj.handle, pSourceSurface.handle, &c_pSourceRect, pDestSurface.handle, &c_pDestRect, (C.D3DTEXTUREFILTERTYPE)(Filter)))
	return
}

func (obj Device) TestCooperativeLevel() (err error) {
	err = toErr(C.IDirect3DDevice9TestCooperativeLevel(obj.handle))
	return
}

func (obj Device) UpdateSurface(pSourceSurface Surface, pSourceRect RECT, pDestinationSurface Surface, pDestinationPoint POINT) (err error) {
	c_pSourceRect := pSourceRect.toC()
	c_pDestinationPoint := pDestinationPoint.toC()
	err = toErr(C.IDirect3DDevice9UpdateSurface(obj.handle, pSourceSurface.handle, &c_pSourceRect, pDestinationSurface.handle, &c_pDestinationPoint))
	return
}

func (obj Device) UpdateTexture(pSourceTexture BaseTexture, pDestinationTexture BaseTexture) (err error) {
	err = toErr(C.IDirect3DDevice9UpdateTexture(obj.handle, pSourceTexture.handle, pDestinationTexture.handle))
	return
}

func (obj Device) ValidateDevice() (pNumPasses uint32, err error) {
	var c_pNumPasses C.DWORD
	err = toErr(C.IDirect3DDevice9ValidateDevice(obj.handle, &c_pNumPasses))
	pNumPasses = (uint32)(c_pNumPasses)
	return
}
