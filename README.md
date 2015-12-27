# d3d9
This library allows you to use Direct3D9 from your Go programs. It is a thin wrapper around the C API. It was mostly generated from the MSDN documentation and then manually improved where necessary.

# Usage
All Direct3D9 interfaces are translated to Go types. The D3D9 part of the names was dropped, e.g. `IDirect3DDevice9` is now `d3d9.Device`.
Interface methods are simply translated to methods on the Go types and have the same name.

For structures the `D3D` prefix and `9` suffix was dropped, e.g. `D3DCAPS9` is `d3d9.CAPS`.

Constants and enumerations were changed in the same way.

# Building
This library provides a Cgo wrapper around the C++ API so make sure you have a C compiler and the DirectX header files included. Specifically you need `d3d9.h` in your include path.

For initialization the `d3d9.dll` is loaded so it must be installed on the target PC (every version of Windows from XP up should include this, so if you can install Go on your PC, you probably already have it). No static `.lib` files are necessary for building.

Note that Direct3D9 needs a window handle for setting it up. This means that you need some way to create a native window and get the handle to pass it to d3d9. In the samples you can see how to do it using the [SDL2 Go wrapper](https://github.com/veandco/go-sdl2). You can also use another Windows wrapper library, like [Allen Dang's w32](https://github.com/AllenDang/w32) or the [walk library](https://github.com/lxn/walk). You could even write a little Cgo code to set up a window yourself. This library does not provide window creation or event handling functionality, only a Direct3D9 wrapper.

# Status
Right now the code is mostly generated from the MSDN documentation and will need more tweaks and refinement to work more fluently with Go.

For example, to transfer data to and from the GPU, like shader object code or vertex buffer data, the Direct3D9 API uses void* pointers which translate to unsafe.Pointers in Go. This is not a good interface to work with, though. That is why there are already some helper functions like `d3d9.Device.CreatePixelShaderFromBytes` which takes a byte slice and creates and sets a pixel shader in one call. For common operations like this, convenience methods come in very handy and reduce the friction when working in Go.

# Help improve this library

Only real world use and feedback can improve the usability of this library, so please use it, fork it, send pull requests and issue issues to help bring Direct3D to Go!
