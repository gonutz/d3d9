This example initializes Direct3D9 and draws a static pink triangle in the
center of a black background. It shows how to create and use vertex buffers and
shaders. The shaders are compiled on from source using the
[dxc](https://github.com/gonutz/dxc) library, which provides a wrapper around
the DirectX compiler DLL.

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
