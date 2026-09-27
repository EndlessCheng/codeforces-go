package main

// https://space.bilibili.com/206214
func canTransform(source, target []int) bool {
	diff := 0
	for i, x := range source {
		diff += x - target[i]
	}
	return diff == 0
}
