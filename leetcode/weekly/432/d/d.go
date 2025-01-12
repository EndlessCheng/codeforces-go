package main

// https://space.bilibili.com/206214
func countNonDecreasingSubarrays(nums []int, k int) (ans int64) {
	n := len(nums)
	g := make([][]int, n)
	posR := make([]int, n)
	st := []int{}
	for i, x := range nums {
		for len(st) > 0 && x >= nums[st[len(st)-1]] {
			posR[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
		// 循环结束后，栈顶就是左侧 > x 的最近元素了
		if len(st) > 0 {
			left := st[len(st)-1]
			g[left] = append(g[left], i)
		}
		st = append(st, i)
	}
	for _, i := range st {
		posR[i] = n
	}

	cnt := 0
	l := 0
	q := []int{} // 单调队列维护最大值
	for r, x := range nums {
		// x 进入窗口
		for len(q) > 0 && nums[q[len(q)-1]] <= x {
			q = q[:len(q)-1] // 维护 q 的单调性
		}
		q = append(q, r)

		// 由于队首到队尾单调递减，所以窗口最大值就是队首
		cnt += max(nums[q[0]]-x, 0)

		for cnt > k {
			out := nums[l] // 离开窗口的元素
			for _, i := range g[l] {
				if i > r {
					break
				}
				cnt -= (out - nums[i]) * (min(posR[i], r+1) - i)
			}
			l++

			// 队首已经离开窗口了
			if q[0] < l {
				q = q[1:]
			}
		}

		ans += int64(r - l + 1)
	}
	return
}
