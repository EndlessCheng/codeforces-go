package main

// https://space.bilibili.com/206214
func shadowPairs(nums []int) (ans int64) {
	type pair struct{ x, cnt int }
	st := []pair{{}} // 栈底哨兵
	size := 0 // 栈的大小（cnt 之和）
	for _, x := range nums {
		for st[len(st)-1].x > x {
			size -= st[len(st)-1].cnt
			st = st[:len(st)-1]
		}

		ans += int64(size)
		if st[len(st)-1].x == x {
			// 恰好等于 x 的 nums[i] 不能构成影子对，要减掉
			ans -= int64(st[len(st)-1].cnt)
			st[len(st)-1].cnt++
		} else {
			st = append(st, pair{x, 1})
		}
		size++
	}
	return
}
