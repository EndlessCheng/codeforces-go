package main

// O(n) 一次遍历 + O(1) 空间

// github.com/EndlessCheng/codeforces-go
func minimumBuckets(street string) (ans int) {
	bucketPos := -2 // 上一个水桶的放置位置
	for i, ch := range street {
		if ch == 'H' { // 遍历每个房屋
			if bucketPos == i-1 { // 左侧已有水桶，不做任何处理
			} else if i+1 < len(street) && street[i+1] == '.' { // 贪心：先考虑在右侧放水桶
				bucketPos = i + 1
				ans++
			} else if i > 0 && street[i-1] == '.' { // 再考虑在左侧放水桶
				ans++
			} else { // 无解：左右均无空位
				return -1
			}
		}
	}
	return
}
