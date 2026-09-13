package main

// https://space.bilibili.com/206214
// 判断闭区间 [l1, r1] 和 [l2, r2] 是否严格相交
func isIntervalOverlap(l1, r1, l2, r2 int) bool {
	return r1 > l2 && r2 > l1
}

func isRectangleOverlap(rec1, rec2 []int) bool {
	return isIntervalOverlap(rec1[0], rec1[2], rec2[0], rec2[2]) &&
		isIntervalOverlap(rec1[1], rec1[3], rec2[1], rec2[3])
}
