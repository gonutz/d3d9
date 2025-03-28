This sample displays a rotating rectangle. It shows you how to use vertex
buffers and shaders with a transformation matrix. The shaders are compiled on 
from source using the [dxc](https://github.com/gonutz/dxc) library, which
provides a wrapper around the DirectX compiler DLL.

Note that the rectangle's aspect ratio is that of the window, so if you resize
the window, the rectangle changes size. That is because we use the virtual
coordinate system of Direct3D which has x and y axes go from -1 to 1. Usually
you would scale these to the window resolution. This is omitted for simplicity
of the example.

For window creation this sample uses
[the Windows API](https://github.com/gonutz/w32). Direct3D needs a handle to
the window it runs in so you need a method for setting this up. Other libraries
that you can use include [SDL2](https://github.com/veandco/go-sdl2),
[Allen Dang's gform library](https://github.com/AllenDang/gform) and the
[walk library](https://github.com/lxn/walk).

If you simply build this sample with `go build` the resulting program will keep
a console window open while running. Use the `build.bat` to build instead, it
passes the flag `-H=windowsgui` to the linker which gets rid of the console
window, like so:

	go build -ldflags="-H=windowsgui"
