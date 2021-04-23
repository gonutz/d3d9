This sample switches to fullscreen mode and displays a pink background. There
is no interactivity, to quit the program use Alt+F4 or ESCAPE.

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
