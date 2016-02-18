# d3d9
This library is a Go wrapper for Microsoft's Direct3D9 API.

# Build
Besides a working Go installation you need to have a C compiler installed. The library was tested with [MinGW (32 bit)](https://sourceforge.net/projects/mingw/files/) on Windows XP and with MinGW (64 bit) on Windows 8.1. No additional Go or C libraries are needed to build, once you have set up your C compiler just use

    go get -u github.com/gonutz/d3d9

to build and install.

To run a Direct3D9 application you need to have d3d9.dll installed on your system. Since Windows XP this is pre-installed with all Windows versions, so if you have Go installed you will have d3d9.dll as well. This also means that you will not need to ship any additional DLL when building with this library.

# Usage
All Direct3D9 interfaces are translated to Go types and their methods are translated to functions on the types so the usage is very close to the C++ API.

There are some differences in the names in Go, since the package is named `d3d9`, all names in that package drop the `D3D` and `9` parts because they would be redundant. The changes are:

- Interfaces drop the `IDirect3D` prefix and the `9` suffix, e.g. `IDirect3DDevice9` becomes `d3d9.Device`. The only exceptions is `IDirect3D9` which in Go becomes `Direct3D`.
- Constants and enumerations drop the `D3D` prefix, otherwise they are the same and keep the upper case convention so users of Direct3D can easily find what they are looking for. For example `D3DADAPTER_DEFAULT` becomes `d3d9.ADAPTER_DEFAULT`, `D3DFMT_R8G8B8` becomes `d3d9.FMT_R8G8B8` etc.
- Structs, like constants, only drop the `D3D` prefix, they too keep the upper case naming convention, so `D3DRANGE` becomes `RANGE`. There is one exception to this, `D3DRECT` is still `D3DRECT` because the API also uses Windows' `RECT` type and these are two distinct types.
- Error constants also drop the `D3D` prefix so `D3DERR_OUTOFVIDEOMEMORY` becomes `ERR_OUTOFVIDEOMEMORY`. However, the interface functions do not return these constants, they return Go `error`s instead of `HRESULT`s.

Note that Direct3D9 needs a window handle for setting it up. This means that you need some way to create a native window and get the handle to pass it to d3d9. In the samples you can see how to do it using the [SDL2 Go wrapper](https://github.com/veandco/go-sdl2) and [Allen Dang's w32](https://github.com/AllenDang/w32). You can also use another Windows wrapper library, like the [walk library](https://github.com/lxn/walk). You can also write a little CGo code to set up a window yourself. This library does not provide window creation or event handling functionality, only the Direct3D9 wrapper.

When you write a program using this library, make sure to add this code in your main package:

    func init() {
	    runtime.LockOSThread()
	}

# Documentation
See the [GoDoc](https://godoc.org/github.com/gonutz/d3d9) for the Go API. The functions are only documented very generally, to get more detailed information about the Direct3D9 API see the [MSDN documentation](https://msdn.microsoft.com/en-us/library/windows/desktop/bb172964%28v=vs.85%29.aspx).

# Status
The code started out as a rough version generated from the MSDN online documentation and was then manually refined to work well with Go and its conventions. It wraps the whole Direct3D9 API right now and adds some additional convenience functions for easy Go usage. However, LOCKED_RECT only has one function for setting the whole data rectangle at once, there is no simple way to set only sub-rectangles right now. Similarly the CubeTexture's LOCKED_BOX has no wrapper functions for getting and setting its data, yet. For now the raw unsafe.Pointer memory pointers can be used along with the pitch values and in the future such functions can be added when needed.

# Help improve this library

Only real world use and feedback can improve the usability of this library, so please use it, fork it, send pull requests and create issues to help improve this library.
