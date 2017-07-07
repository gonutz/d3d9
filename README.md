# d3d9
This library is a pure Go wrapper for Microsoft's Direct3D9 API.

# Games
Here you can see the library in action, it was used for two 48 hour game jams.

This is [my entry](https://github.com/gonutz/gophette) to the [GopherGala 2016](http://gophergala.com/):

![Gophette Screenshot](https://raw.githubusercontent.com/gonutz/gophette/master/screenshots/race.png)

And here is [my entry](https://github.com/gonutz/ld36) for the [Ludum Dare Game Jam 2016](http://ludumdare.com/compo/ludum-dare-36/?action=preview&uid=110557):

![Reinventing the Wheel Screenshot](http://ludumdare.com/compo/wp-content/compo2//570486/110557-shot0-1472432554.png-eq-900-500.jpg)

# Installation
Get and build the library with:

    go get -u github.com/gonutz/d3d9

To run a Direct3D9 application you need to have `d3d9.dll` installed on your system. Luckily, all Windows versions starting with Windows XP have this pre-installed. This means that you will not need to ship any additional DLL with your application when using this library.

# Obsolete CGo Version 1
This library does not use CGo anymore. With this change the API has incompatibly changed in some places. If you have a project using the old API and do not want to update everything right now, you can use a copy of version 1.0.0 of this library at [github.com/gonutz/cgo/d3d9](https://github.com/gonutz/cgo/tree/master/d3d9). All you have to do is go get that library and update your imports from `github.com/gonutz/d3d9` to `github.com/gonutz/cgo/d3d9`.

# Usage
All Direct3D9 interfaces are translated to Go types and their methods are translated to functions on the types so the usage is very close to the C++ API.

There are some differences in the names in Go, since the package is named `d3d9`, all names in that package drop the `D3D` and `9` parts because they would be redundant. The changes are:

- Interfaces drop the `IDirect3D` prefix and the `9` suffix, e.g. `IDirect3DDevice9` becomes `d3d9.Device`. The only exception is `IDirect3D9` which in Go becomes `Direct3D`.
- Constants and enumerations drop the `D3D` prefix, otherwise they are the same and keep the upper case convention so users of Direct3D can easily find what they are looking for. For example `D3DADAPTER_DEFAULT` becomes `d3d9.ADAPTER_DEFAULT`, `D3DFMT_R8G8B8` becomes `d3d9.FMT_R8G8B8` etc.
- Structs, like constants, only drop the `D3D` prefix, they too keep the upper case naming convention, so `D3DRANGE` becomes `d3d9.RANGE`.
- Error constants also drop the `D3D` prefix so `D3DERR_OUTOFVIDEOMEMORY` becomes `ERR_OUTOFVIDEOMEMORY`. However, the interface functions do not return these constants, they return Go `error`s instead of `HRESULT`s.
- Instead of returning `HRESULT`, functions return `d3d9.Error` which implements the standard Go error interface and has an additional function `Code() int32` which returns the error code. This code can be checked against the defined error constants. If a function succeeds it returns `nil` (and not `D3D_OK`) as the d3d9.Error.

Note that Direct3D9 needs a window handle for setting it up. This means that you need some way to create a native window and get the handle to pass it to d3d9. In the samples you can see how to do it using the [SDL2 Go wrapper](https://github.com/veandco/go-sdl2) and [Allen Dang's w32](https://github.com/AllenDang/w32). You could also use other Windows wrapper libraries, like the [walk library](https://github.com/lxn/walk), or just write a little CGo code to set up a window yourself. This library does not provide window creation or event handling functionality, only the Direct3D9 wrapper.

All calls to Direct3D9 must happen from the same thread that creates the Direct3D9 object so make sure to add this code in your main package:

    func init() {
	    runtime.LockOSThread()
	}

There are some additional convenience functions. `IndexBuffer` and `VertexBuffer` have `Lock` functions which return `void*` pointers in the C API and would thus return `uintptr`s in Go. You can use these pointers to read and write various types from/to that memory. However, using `uintptr` or `unsafe.Pointer` is not idiomatic Go so the `Lock` functions return a wrapper around the `uintptr` instead, providing functions of the form `GetFloat32s` and `SetFloat32s` which take a slice of `[]float32` and handle copying the data for you. See the documentation of these functions for further information.

Similarly, when locking a two-dimensional resource, like a texture, you will get a `d3d9.LOCKED_RECT`. It too has a wrapper function for setting its 2D data. There is however no function to read the data, yet, and no functions for reading and writing sub-rectangles which could be useful. These functions can be added when needed by a real application to see the use case first and implement the functions according to that instead of guessing what might work. If you need such a function, create a pull request or an issue and it can be incorporated in this library.

# Documentation
See the [GoDoc](https://godoc.org/github.com/gonutz/d3d9) for the Go API. The functions are only documented very generally, to get more detailed information about the Direct3D9 API see the [MSDN documentation](https://msdn.microsoft.com/en-us/library/windows/desktop/bb172964%28v=vs.85%29.aspx).

# Examples
The samples folder contains some example programs that show the basics of setting up and using the library.

A real world usage of this library is [my entry](https://github.com/gonutz/gophette) to the [GopherGala 2016](http://gophergala.com/). It was first built using SDL2 and after the competition I ported it to DirectX, see [`main_windows.go`](https://github.com/gonutz/gophette/blob/master/main_windows.go) for the implementation.

Another real world example is [my entry](https://github.com/gonutz/ld36) for the [Ludum Dare Game Jam 2016](http://ludumdare.com/compo/ludum-dare-36/?action=preview&uid=110557) which is Windows only.

The [prototype library](https://github.com/gonutz/prototype) also uses `d3d9` on Windows.

# Status
The code started out as a rough version generated from the MSDN online documentation and was then manually refined to work well with Go and its conventions. It wraps the whole Direct3D9 API right now and adds some additional convenience functions for easy Go usage. However, `LOCKED_RECT` only has one function for setting the whole data rectangle at once, there is no simple way to set only sub-rectangles right now. Similarly the CubeTexture's `LOCKED_BOX` has no wrapper functions for getting and setting its data, yet. For now the raw `uintptr` memory pointers can be used along with the pitch values and in the future such functions can be added when needed.

# Help improve this library

Only real world use and feedback can improve the usability of this library, so please use it, fork it, send pull requests and create issues to help improve this library.
