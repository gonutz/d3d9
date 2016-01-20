package d3d9

//#include "d3d9wrapper.h"
import "C"

type COLOR uint32

// TODO remove this eventually
func (c COLOR) toC() C.D3DCOLOR {
	return (C.D3DCOLOR)(c)
}

func ColorARGB(a, r, g, b uint) COLOR {
	return COLOR((a&0xFF)<<24 | (r&0xFF)<<16 | (g&0xFF)<<8 | b&0xFF)
}

func ColorRGBA(r, g, b, a uint) COLOR {
	return ColorARGB(a, r, g, b)
}

func ColorXRGB(r, g, b uint) COLOR {
	return ColorARGB(0xFF, r, g, b)
}

func ColorRGB(r, g, b uint) COLOR {
	return ColorARGB(0xFF, r, g, b)
}

func ColorXYUV(y, u, v uint) COLOR {
	return ColorARGB(0xFF, y, u, v)
}

func ColorYUV(y, u, v uint) COLOR {
	return ColorARGB(0xFF, y, u, v)
}

func ColorAYUV(a, y, u, v uint) COLOR {
	return ColorARGB(a, y, u, v)
}

func ColorValue(r, g, b, a float32) COLOR {
	return ColorARGB(uint(a*255), uint(r*255), uint(g*255), uint(b*255))
}
