package main

// https://space.bilibili.com/206214
func countNonDecreasingSubarrays(nums []int, k int) (ans int64) {
	n := len(nums)
	cnt, rTree := 0, 0 // rTree 表示窗口最右边那棵树在 st 中的下标
	type pair struct{ val, size int } // 根节点的值, 树的大小
	st := []pair{}
	r := n - 1
	for l := n - 1; l >= 0; l-- {
		// x 进入窗口
		x := nums[l]
		size := 1 // 统计以 x 为根的树的大小
		for len(st) > 0 && x >= st[len(st)-1].val {
			// 以 val 为根的树，现在合并到 x 的下面（x 和 val 连一条边）
			p := st[len(st)-1]
			st = st[:len(st)-1]
			size += p.size
			cnt += (x - p.val) * p.size // 树 val 中的数都变成 x
		}
		st = append(st, pair{x, size})

		// 如果从 st 中弹出树包含 rTree，那么 rTree 现在指向栈顶这棵树
		rTree = min(rTree, len(st)-1)

		// 当 cnt 大于 k 时，缩小窗口
		for cnt > k {
			// 操作次数的减少量，等于 nums[r] 所在树的根节点值减去 nums[r]
			cnt -= st[rTree].val - nums[r]
			r--
			// nums[r] 离开窗口后，树的大小减一
			st[rTree].size--
			if st[rTree].size == 0 { // 这棵树是空的
				rTree++ // rTree 指向左边下一棵树
			}
		}

		ans += int64(r - l + 1)
	}
	return
}

func countNonDecreasingSubarrays2(nums []int, k int) (ans int64) {
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
