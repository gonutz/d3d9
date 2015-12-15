This example initializes Direct3D9 and simple changes the background color in an endless loop.
It shows the basic setup and teardown of the library.

For window creating this sample uses [SDL2](https://github.com/veandco/go-sdl2). Direct3D needs a handle to the window it runs in so you need a method for setting this up. Other libraries that you can use include [Allen Dang's w32](https://github.com/AllenDang/w32) or the [walk library](https://github.com/lxn/walk).

If you simply build this sample with `go build` the resulting program will keep a console window open while running. Use the `build.bat` to build instead, it passes the flag `-H=windowsgui` to the linker which gets rid of the console window.