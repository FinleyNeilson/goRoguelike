package utils

// Flatten converts (x, y) into a 1D index for a row-major grid with given width.
func Flatten(x, y, width int) int {
	return y*width + x
}

// Unflatten converts a 1D index into (x, y) coordinates for a row-major grid.
func Unflatten(index, width int) (x, y int) {
	x = index % width
	y = index / width
	return
}
