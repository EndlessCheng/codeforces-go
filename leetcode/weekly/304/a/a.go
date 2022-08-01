package main

/*

下午 2 点在 B 站直播讲周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

*/

// https://space.bilibili.com/206214
func minimumOperations(nums []int) int {
	set := map[int]struct{}{}
	for _, x := range nums {
		if x > 0 {
			set[x] = struct{}{}
		}
	}
	return len(set)
}
