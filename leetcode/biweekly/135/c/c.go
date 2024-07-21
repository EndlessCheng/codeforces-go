package main

// https://space.bilibili.com/206214
func minChanges(nums []int, k int) int {
	n := len(nums)
	d := make([]int, k+2)
	for i := 0; i < n/2; i++ {
		p, q := nums[i], nums[n-1-i]
		if p > q { // 保证 p <= q
			p, q = q, p
		}
		x := q - p
		mx := max(q, k-p)
		// [0, x-1] 全部 +1：把 q-p 改成小于 x 的，只需要改 p 或 q 中的一个数
		d[0]++
		d[x]--
		// [x+1, mx] 全部 +1：把 q-p 改成大于 x 小于等于 mx 的，也只需要改 p 或 q 中的一个数
		d[x+1]++
		d[mx+1]--
		// [mx+1, k] 全部 +2：把 q-p 改成大于 mx 的，p 和 q 都需要改
		d[mx+1] += 2
	}

	ans := n
	minModify := 0
	for _, v := range d {
		minModify += v
		ans = min(ans, minModify)
	}
	return ans
}

func minChanges2(nums []int, k int) int {
	cnt := make([]int, k+1)
	cnt2 := make([]int, k+1)
	n := len(nums)
	for i := 0; i < n/2; i++ {
		p, q := nums[i], nums[n-1-i]
		if p > q { // 保证 p <= q
			p, q = q, p
		}
		cnt[q-p]++
		cnt2[max(q, k-p)]++
	}

	ans := n
	sum2 := 0 // 统计有多少对 (p,q) 都要改
	for x, c := range cnt {
		// 其他 n/2-c 对 (p,q) 至少要改一个数，在此基础上，有额外的 sum2 对 (p,q) 还要再改一个数
		ans = min(ans, n/2-c+sum2)
		// 对于后面的更大的 x，当前的这 cnt2[x] 对 (p,q) 都要改
		sum2 += cnt2[x]
	}
	return ans
}
