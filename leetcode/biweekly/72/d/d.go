package main

// github.com/EndlessCheng/codeforces-go
type fenwick []int

func newFenwickTree(n int) fenwick {
	return make(fenwick, n+1) // 使用下标 1 到 n
}

// a[i] 增加 val
// 1 <= i <= n
func (f fenwick) update(i int, val int) {
	for ; i < len(f); i += i & -i {
		f[i] += val
	}
}

// 求前缀和 a[1] + ... + a[i]
// 1 <= i <= n
func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

func goodTriplets(nums1, nums2 []int) (ans int64) {
	n := len(nums1)
	p := make([]int, n)
	for i, x := range nums1 {
		p[x] = i
	}

	t := newFenwickTree(n)
	for i, y := range nums2[:n-1] {
		y = p[y]
		less := t.pre(y)
		ans += int64(less) * int64(n-1-y-(i-less))
		t.update(y+1, 1)
	}
	return
}
