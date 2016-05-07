package d3d9

// PSVersion creates a pixel shader version token.
func PSVersion(major, minor uint) uint32 {
	return uint32(0xFFFF0000 | ((major) << 8) | (minor))
}

// VSVersion creates a vertex shader version token.
func VSVersion(major, minor uint) uint32 {
	return uint32(0xFFFE0000 | ((major) << 8) | (minor))
}

// TSWorldMatrix maps indices in the range 0 through 255 to the corresponding
// transform states.
func TSWorldMatrix(index int) TRANSFORMSTATETYPE {
	return TRANSFORMSTATETYPE(index + 256)
}

// FVFTexCoordSize1 constructs bit patterns that are used to identify texture
// coordinate formats within a FVF description. The results of these macros can
// be combined within a FVF description by using the OR operator.
func FVFTexCoordSize1(CoordIndex uint) uint32 {
	return uint32(int(FVF_TEXTUREFORMAT1) << (CoordIndex*2 + 16))
}

// FVFTexCoordSize2 constructs bit patterns that are used to identify texture
// coordinate formats within a FVF description. The results of these macros can
// be combined within a FVF description by using the OR operator.
func FVFTexCoordSize2(CoordIndex uint) uint32 {
	return FVF_TEXTUREFORMAT2
}

// FVFTexCoordSize3 constructs bit patterns that are used to identify texture
// coordinate formats within a FVF description. The results of these macros can
// be combined within a FVF description by using the OR operator.
func FVFTexCoordSize3(CoordIndex uint) uint32 {
	return FVF_TEXTUREFORMAT3 << (CoordIndex*2 + 16)
}

// FVFTexCoordSize4 constructs bit patterns that are used to identify texture
// coordinate formats within a FVF description. The results of these macros can
// be combined within a FVF description by using the OR operator.
func FVFTexCoordSize4(CoordIndex uint) uint32 {
	return FVF_TEXTUREFORMAT4 << (CoordIndex*2 + 16)
}
