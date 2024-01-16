package rectangle_overlap

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	x10, y10, x11, y11 := rec1[0], rec1[1], rec1[2], rec1[3]
	x20, y20, x21, y21 := rec2[0], rec2[1], rec2[2], rec2[3]

	return !(x10 >= x21 || x11 <= x20 || y10 >= y21 || y11 <= y20)
}
