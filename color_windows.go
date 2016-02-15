package d3d9

//#include "d3d9wrapper.h"
import "C"

type COLOR uint32

func ColorARGB(a, r, g, b uint8) COLOR {
	return COLOR(a)<<24 | COLOR(r)<<16 | COLOR(g)<<8 | COLOR(b)
}

func ColorRGBA(r, g, b, a uint8) COLOR {
	return ColorARGB(a, r, g, b)
}

func ColorXRGB(r, g, b uint8) COLOR {
	return ColorARGB(0xFF, r, g, b)
}

func ColorRGB(r, g, b uint8) COLOR {
	return ColorARGB(0xFF, r, g, b)
}

func ColorXYUV(y, u, v uint8) COLOR {
	return ColorARGB(0xFF, y, u, v)
}

func ColorYUV(y, u, v uint8) COLOR {
	return ColorARGB(0xFF, y, u, v)
}

func ColorAYUV(a, y, u, v uint8) COLOR {
	return ColorARGB(a, y, u, v)
}

func ColorValue(r, g, b, a float32) COLOR {
	return ColorARGB(uint8(a*255), uint8(r*255), uint8(g*255), uint8(b*255))
}
