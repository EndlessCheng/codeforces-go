package main

// https://space.bilibili.com/206214
type fenwick []int

func (t fenwick) update(i, val int) {
	for ; i < len(t); i += i & -i {
		t[i] += val
	}
}

func (t fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += t[i]
	}
	return
}

func (t fenwick) query(l, r int) int {
	if l > r {
		return 0
	}
	return t.pre(r) - t.pre(l-1)
}

func minOperations(nums, locked []int, k int) (ans int) {
	n := len(nums)
	t := make(fenwick, n+1)
	for i := 0; i < n; {
		if nums[i] == i+1 { // nums[i] 已在正确位置上
			i++
			continue
		}
		start := i
		mn, mx := n+1, 0
		// 这一段的最小值必须是 start+1，且值域大小必须等于段长度
		for ; i < n && (mn != start+1 || mx-mn+1 != i-start); i++ {
			if locked[i] > 0 {
				if k == 0 { // 无法解锁
					return -1
				}
				k--
				ans++
			}
			x := nums[i]
			ans += t.query(x, mx) // x 与其左侧元素的逆序对个数
			t.update(x, 1)
			mn = min(mn, x)
			mx = max(mx, x)
		}
	}
	return
}
