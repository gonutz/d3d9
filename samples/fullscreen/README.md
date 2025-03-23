This sample switches to fullscreen mode and displays a pink background. There
is no interactivity, to quit the program use Alt+F4 or ESCAPE.

Note that this uses "real fullscreen" as opposed to "windowed fullscreen". That
means that this program really switches your monitor's display mode. This
typically takes time and does not play nicely with switching windows, e.g.
Alt+Tab is often a problem in games going full screen to a different
resolution.
You usually want to go fullscreen to the same resolution that the monitor is in
right now, or use windowed fullscreen where you really just set up a window
without borders that covers the whole screen and run your game in windowed
mode.

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
