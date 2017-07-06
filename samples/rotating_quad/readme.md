This sample displays a rotating rectangle. It shows you how to use vertex buffers and shaders with a transformation matrix. The shaders were compiled with `FXC.exe`, the DirectX effects compiler, which is included in the DirectX SDK and the Windows Platform SDK.

The script `build.bat` only builds the executable. The necessary shader objects are included in this repository.

The shaders are defined in `vs.code` (vertex shader) and `ps.code` (pixel shader). Use the script `build_all.bat` to recompile the shaders and then build the executable. You must have `FXC.exe` installed and the variable `FXC` in the `build_all.bat` script must point to it. The script will then compile `vs.code` and `ps.code` into `vs.object` and `ps.object` object files. These object files are then converted to byte slices that are embedded in the Go code. This is done with the `bin2go` tool. This makes it unnecessary to load files at run-time, the shader object code is shipped with the executable.

For window creation this sample uses [Allen Dang's gform library](https://github.com/AllenDang/gform). Direct3D needs a handle to the window it runs in so you need a method for setting this up. Other libraries that you can use include [SDL2](https://github.com/veandco/go-sdl2) and the [walk library](https://github.com/lxn/walk).

If you simply build this sample with `go build` the resulting program will keep a console window open while running. Use the `build.bat` to build instead, it passes the flag `-H=windowsgui` to the linker which gets rid of the console window.