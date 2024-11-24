package main

// https://space.bilibili.com/206214
func minArraySum(nums []int, k, op1, op2 int) int {
	f := make([][]int, op1+1)
	for i := range f {
		f[i] = make([]int, op2+1)
	}
	for _, x := range nums {
		var y int
		if (x+1)/2 >= k {
			y = (x+1)/2 - k
		} else {
			y = (x - k + 1) / 2
		}
		for p := op1; p >= 0; p-- {
			for q := op2; q >= 0; q-- {
				res := f[p][q] + x
				if p > 0 {
					res = min(res, f[p-1][q]+(x+1)/2)
				}
				if q > 0 && x >= k {
					res = min(res, f[p][q-1]+x-k)
					if p > 0 {
						res = min(res, f[p-1][q-1]+y)
					}
				}
				f[p][q] = res
			}
		}
	}
	return f[op1][op2]
}

func minArraySum2(nums []int, k, op1, op2 int) int {
	n := len(nums)
	f := make([][][]int, n+1)
	for i := range f {
		f[i] = make([][]int, op1+1)
		for j := range f[i] {
			f[i][j] = make([]int, op2+1)
		}
	}
	for i, x := range nums {
		for p := 0; p <= op1; p++ {
			for q := 0; q <= op2; q++ {
				res := f[i][p][q] + x
				if p > 0 {
					res = min(res, f[i][p-1][q]+(x+1)/2)
				}
				if q > 0 && x >= k {
					res = min(res, f[i][p][q-1]+x-k)
					if p > 0 {
						var y int
						if (x+1)/2 >= k {
							y = (x+1)/2 - k
						} else {
							y = (x - k + 1) / 2
						}
						res = min(res, f[i][p-1][q-1]+y)
					}
				}
				f[i+1][p][q] = res
			}
		}
	}
	return f[n][op1][op2]
}

func minArraySum3(nums []int, k, op1, op2 int) int {
	n := len(nums)
	memo := make([][][]int, n)
	for i := range memo {
		memo[i] = make([][]int, op1+1)
		for j := range memo[i] {
			memo[i][j] = make([]int, op2+1)
			for k := range memo[i][j] {
				memo[i][j][k] = -1 // -1 表示没有计算过
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, op1, op2 int) int {
		if i < 0 {
			return 0
		}
		p := &memo[i][op1][op2]
		if *p != -1 { // 之前计算过
			return *p
		}
		x := nums[i]
		res := dfs(i-1, op1, op2) + x
		if op1 > 0 {
			res = min(res, dfs(i-1, op1-1, op2)+(x+1)/2)
		}
		if op2 > 0 && x >= k {
			res = min(res, dfs(i-1, op1, op2-1)+x-k)
			if op1 > 0 {
				var y int
				if (x+1)/2 >= k {
					y = (x+1)/2 - k // 先除再减更优
				} else {
					y = (x - k + 1) / 2 // 只能先减再除
				}
				res = min(res, dfs(i-1, op1-1, op2-1)+y)
			}
		}
		*p = res // 记忆化
		return res
	}
	return dfs(n-1, op1, op2)
}
