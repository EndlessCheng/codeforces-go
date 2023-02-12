package main

// https://space.bilibili.com/206214
func findTheArrayConcVal(a []int) (ans int64) {
	for len(a) > 1 {
		x := a[0]
		for y := a[len(a)-1]; y > 0; y /= 10 {
			x *= 10
		}
		ans += int64(x + a[len(a)-1])
		a = a[1 : len(a)-1]
	}
	if len(a) > 0 {
		ans += int64(a[0])
	}
	return
}
