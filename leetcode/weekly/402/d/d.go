package main

// https://space.bilibili.com/206214
type fenwick []int

func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] += val
	}
}

func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return res
}

func (f fenwick) query(l, r int) int {
	if r < l {
		return 0
	}
	return f.pre(r) - f.pre(l-1)
}

func countOfPeaks(nums []int, queries [][]int) (ans []int) {
	n := len(nums)
	f := make(fenwick, n-1)
	update := func(i, val int) {
		if nums[i-1] < nums[i] && nums[i] > nums[i+1] {
			f.update(i, val)
		}
	}
	for i := 1; i < n-1; i++ {
		update(i, 1)
	}
	for _, q := range queries {
		if q[0] == 1 {
			ans = append(ans, f.query(q[1]+1, q[2]-1))
			continue
		}
		i := q[1]
		for j := max(i-1, 1); j <= min(i+1, n-2); j++ {
			update(j, -1)
		}
		nums[i] = q[2]
		for j := max(i-1, 1); j <= min(i+1, n-2); j++ {
			update(j, 1)
		}
	}
	return
}
