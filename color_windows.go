package d3d9

// COLOR represents a color with red, green, blue and alpha channels, 8 bits
// each, in a 32 bit unsinged integer.
type COLOR uint32

// ColorARGB converts the given color channel values to a color. Each channel
// ranges from 0 (no intensity) to 255 (full intensity).
func ColorARGB(a, r, g, b uint8) COLOR {
	return COLOR(a)<<24 | COLOR(r)<<16 | COLOR(g)<<8 | COLOR(b)
}

// ColorRGBA converts the given color channel values to a color. Each channel
// ranges from 0 (no intensity) to 255 (full intensity).
func ColorRGBA(r, g, b, a uint8) COLOR {
	return ColorARGB(a, r, g, b)
}

// ColorXRGB converts the given color channel values to a color. Each channel
// ranges from 0 (no intensity) to 255 (full intensity).
func ColorXRGB(r, g, b uint8) COLOR {
	return ColorARGB(0xFF, r, g, b)
}

// ColorRGB converts the given color channel values to a color. Each channel
// ranges from 0 (no intensity) to 255 (full intensity).
func ColorRGB(r, g, b uint8) COLOR {
	return ColorARGB(0xFF, r, g, b)
}

// ColorXYUV converts the given color channel values to a color. Each channel
// ranges from 0 (no intensity) to 255 (full intensity).
func ColorXYUV(y, u, v uint8) COLOR {
	return ColorARGB(0xFF, y, u, v)
}

// ColorYUV converts the given color channel values to a color. Each channel
// ranges from 0 (no intensity) to 255 (full intensity).
func ColorYUV(y, u, v uint8) COLOR {
	return ColorARGB(0xFF, y, u, v)
}

// ColorAYUV converts the given color channel values to a color. Each channel
// ranges from 0 (no intensity) to 255 (full intensity).
func ColorAYUV(a, y, u, v uint8) COLOR {
	return ColorARGB(a, y, u, v)
}

// ColorValue converts the given color channel values to a color. Each channel
// ranges from 0.0 (no intensity) to 1.0 (full intensity).
func ColorValue(r, g, b, a float32) COLOR {
	return ColorARGB(uint8(a*255), uint8(r*255), uint8(g*255), uint8(b*255))
}
