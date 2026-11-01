package main

// https://space.bilibili.com/206214
// 模板来源 https://leetcode.cn/discuss/post/3583665/
type fenwick []int

// a[i] 增加 val
// 1 <= i <= n
// 时间复杂度 O(log n)
func (f fenwick) update(i int, val int) {
	for ; i < len(f); i += i & -i {
		f[i] += val
	}
}

// 求前缀和 a[1] + ... + a[i]
// 1 <= i <= n
// 时间复杂度 O(log n)
func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

// 求区间和 a[l] + ... + a[r]
// 1 <= l <= r <= n
// 时间复杂度 O(log n)
func (f fenwick) query(l, r int) int {
	return f.pre(r) - f.pre(l-1)
}

func minOperations(nums, locked []int, k int) (ans int) {
	n := len(nums)
	t := make(fenwick, n+1)
	for i := 0; i < n; {
		if nums[i] == i+1 { // nums[i] 已在正确位置上
			i++
			continue
		}
		mx := -1
		for ; i < n && mx != i; i++ {
			if locked[i] > 0 {
				if k == 0 { // 无法解锁
					return -1
				}
				k--
				ans++
			}
			x := nums[i]
			v := n + 1 - x
			ans += t.pre(v) // x 与其左侧元素的逆序对个数
			t.update(v, 1)
			mx = max(mx, x)
		}
	}
	return
}
