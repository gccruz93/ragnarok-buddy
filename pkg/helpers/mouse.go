package helpers

func IsMouseHover(mx, my int, x, y, x2, y2 float64) bool {
	return (float64(mx) >= x && float64(mx) <= x2) && (float64(my) >= y && float64(my) <= y2)
}
