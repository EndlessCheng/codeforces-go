package main

// O(m^3) 二分图最大带权匹配

// github.com/EndlessCheng/codeforces-go
const inf int = 1e9

func KuhnMunkres(wt [][]int) (sum int) {
	match := make([]int, len(wt)) // 右部点匹配了哪一个左部点
	la := make([]int, len(wt))
	for i, row := range wt {
		la[i] = -inf
		for _, v := range row {
			if v > la[i] {
				la[i] = v
			}
		}
	}
	lb := make([]int, len(wt))
	slack := make([]int, len(wt))
	for i := 1; i < len(wt); i++ {
		vb := make([]bool, len(wt))
		for j := 1; j < len(wt); j++ {
			slack[j] = inf
		}
		last := make([]int, len(wt)) // 右部点在交错树中的上一个右部点，用于倒推得到交错路
		y := 0
		match[0] = i // 一开始假设有一条 i-0 的匹配
		for {
			vb[y] = true
			x, nextY := match[y], 0
			delta := inf
			for j := 1; j < len(wt); j++ {
				if !vb[j] {
					if d := la[x] + lb[j] - wt[x][j]; d < slack[j] {
						slack[j] = d
						last[j] = y
					}
					if slack[j] < delta {
						delta = slack[j]
						nextY = j
					}
				}
			}
			// 当 delta=0 时，相当于沿着相等子图向下搜索一层
			// 当 delta>0 时，相当于直接回到最小边（新加入相等子图的边）处开始搜索
			if delta > 0 {
				for j := 0; j < len(wt); j++ {
					if vb[j] {
						la[match[j]] -= delta
						lb[j] += delta
					} else {
						slack[j] -= delta
					}
				}
			}
			y = nextY
			if match[y] == 0 {
				break
			}
		}
		// 倒推更新增广路
		for ; y > 0; y = last[y] {
			match[y] = match[last[y]]
		}
	}
	for w := 1; w < len(wt); w++ {
		sum += wt[match[w]][w]
	}
	return
}

func maxCompatibilitySum(students [][]int, mentors [][]int) int {
	m := len(students)
	wt := make([][]int, m+1)
	wt[0] = make([]int, m+1)
	for i := range wt[0] {
		wt[0][i] = -inf
	}
	for i, st := range students {
		wt[i+1] = make([]int, m+1)
		wt[i+1][0] = -inf
		for j, mt := range mentors {
			for k, v := range st {
				if v == mt[k] {
					wt[i+1][j+1]++
				}
			}
		}
	}
	return KuhnMunkres(wt)
}
