package main

import "math"

// https://space.bilibili.com/206214
func subtreeInversionSum1(edges [][]int, nums []int, k int) int64 {
	n := len(nums)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	memo := make([][][2]int, n)
	for i := range memo {
		memo[i] = make([][2]int, k)
		for j := range memo[i] {
			for p := range memo[i][j] {
				memo[i][j][p] = math.MinInt
			}
		}
	}
	var dfs func(int, int, int, int) int
	dfs = func(x, fa, cd, parity int) int {
		p := &memo[x][cd][parity]
		if *p != math.MinInt {
			return *p
		}

		// 不反转
		res := nums[x] * (1 - parity*2)
		for _, y := range g[x] {
			if y != fa {
				res += dfs(y, x, max(cd-1, 0), parity)
			}
		}

		// 反转
		if cd == 0 {
			s := nums[x] * (parity*2 - 1)
			for _, y := range g[x] {
				if y != fa {
					s += dfs(y, x, k-1, parity^1) // 重置 CD
				}
			}
			res = max(res, s)
		}

		*p = res
		return res
	}
	return int64(dfs(0, -1, 0, 0))
}

func subtreeInversionSum2(edges [][]int, nums []int, k int) int64 {
	n := len(nums)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int, int) [][2]int
	dfs = func(x, fa int) [][2]int {
		v := nums[x]
		res := make([][2]int, k)
		for cd := range res {
			res[cd] = [2]int{v, -v}
		}

		s0, s1 := -v, v
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			fy := dfs(y, x)
			// 不反转
			for cd := range res {
				res[cd][0] += fy[max(cd-1, 0)][0]
				res[cd][1] += fy[max(cd-1, 0)][1]
			}
			// 反转
			s0 += fy[k-1][1]
			s1 += fy[k-1][0]
		}
		// 反转
		res[0][0] = max(res[0][0], s0)
		res[0][1] = max(res[0][1], s1)

		return res
	}

	return int64(dfs(0, -1)[0][0])
}

func subtreeInversionSum(edges [][]int, nums []int, k int) int64 {
	n := len(nums)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	f := [][2]int{}
	var dfs func(int, int) (int, int, int)
	dfs = func(x, fa int) (int, int, int) {
		f = append(f, [2]int{}) // 用于刷表

		s := nums[x]             // 子树和
		notInv0, notInv1 := 0, 0 // 不反转 x 时的额外增量（0 表示上面反转了偶数次，1 表示上面反转了奇数次）
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			sy, y0, y1 := dfs(y, x)
			s += sy
			// 不反转 x，反转次数的奇偶性不变
			notInv0 += y0
			notInv1 += y1
		}

		subRes := f[len(f)-1] // 被刷表后的结果
		f = f[:len(f)-1]

		// 反转 x
		// x 上面反转了偶数次，反转 x 会带来 -2 倍子树和的增量，且对于 x 的 k 级后代来说，上面反转了奇数次（所以是 subRes1）
		inv0 := subRes[1] - s*2
		// x 上面反转了奇数次，反转 x 会带来 2 倍子树和的增量，且对于 x 的 k 级后代来说，上面反转了偶数次（所以是 subRes0）
		inv1 := subRes[0] + s*2

		res0 := max(notInv0, inv0)
		res1 := max(notInv1, inv1)

		// 刷表法：更新 x 的 k 级祖先的状态
		if len(f) >= k {
			f[len(f)-k][0] += res0
			f[len(f)-k][1] += res1
		}

		return s, res0, res1
	}

	s, res0, _ := dfs(0, -1)
	return int64(s + res0) // 对于根节点来说，上面一定反转了偶数次（0 次）
}
