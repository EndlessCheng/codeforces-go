package main

import (
	"math"
	"math/bits"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func maxDistance(side int, points [][]int, k int) int {
	n := len(points)
	a := make([]int, n)
	for i, p := range points {
		x, y := p[0], p[1]
		if x == 0 {
			a[i] = y
		} else if y == side {
			a[i] = side + x
		} else if x == side {
			a[i] = side*3 - y
		} else {
			a[i] = side*4 - x
		}
	}
	slices.Sort(a)

	f := make([]int, n+1)
	end := make([]int, n)
	ans := sort.Search(side, func(low int) bool {
		low++
		j := n
		for i := n - 1; i >= 0; i-- {
			for a[j-1] >= a[i]+low {
				j--
			}
			f[i] = f[j] + 1
			if f[i] == 1 {
				end[i] = i // i 自己就是最后一个点
			} else {
				end[i] = end[j]
			}
			if f[i] == k && a[end[i]]-a[i] <= side*4-low {
				return false
			}
		}
		return true
	})
	return ans
}

func maxDistance4(side int, points [][]int, k int) int {
	n := len(points)
	a := make([]int, n, n+1)
	for i, p := range points {
		x, y := p[0], p[1]
		if x == 0 {
			a[i] = y
		} else if y == side {
			a[i] = side + x
		} else if x == side {
			a[i] = side*3 - y
		} else {
			a[i] = side*4 - x
		}
	}
	slices.Sort(a)
	a = append(a, math.MaxInt) // 哨兵

	g := make([][]int, n+1)
	ans := sort.Search(side, func(low int) bool {
		low++
		clear(g)
		j := n
		for i := n - 1; i >= 0; i-- {
			for a[j-1] >= a[i]+low {
				j--
			}
			g[j] = append(g[j], i) // 建树
		}

		st := []int{}
		var dfs func(int) bool
		dfs = func(x int) bool {
			st = append(st, a[x])
			m := len(st)
			// 注意栈中多了一个 a[n]，所以是 m > k 不是 >=
			if m > k && st[m-k]-a[x] <= side*4-low {
				return true
			}
			for _, y := range g[x] {
				if dfs(y) {
					return true
				}
			}
			st = st[:m-1] // 恢复现场
			return false
		}
		return !dfs(n)
	})
	return ans
}

func maxDistance3(side int, points [][]int, k int) int {
	n := len(points)
	a := make([]int, n)
	for i, p := range points {
		x, y := p[0], p[1]
		if x == 0 {
			a[i] = y
		} else if y == side {
			a[i] = side + x
		} else if x == side {
			a[i] = side*3 - y
		} else {
			a[i] = side*4 - x
		}
	}
	slices.Sort(a)

	k-- // 往后跳 k-1 步，这里先减一，方便计算
	highBit := bits.Len(uint(k)) - 1
	nxt := make([][5]int, n+1) // 5 可以改为 highBit+1（用 array 而不是 slice，提高访问效率）
	for j := range nxt[n] {
		nxt[n][j] = n // 哨兵
	}

	ans := sort.Search(side, func(low int) bool {
		low++
		// 预处理倍增数组 nxt
		j := n
		for i := n - 1; i >= 0; i-- {
			for a[j-1] >= a[i]+low {
				j--
			}
			nxt[i][0] = j
			for k := 1; k <= highBit; k++ {
				nxt[i][k] = nxt[nxt[i][k-1]][k-1]
			}
		}

		// 枚举起点
		for i, start := range a {
			// 往后跳 k-1 步（注意上面把 k 减一了）
			cur := i
			for j := highBit; j >= 0; j-- {
				if k>>j&1 > 0 {
					cur = nxt[cur][j]
				}
			}
			if cur == n { // 出界
				break
			}
			if a[cur]-start <= side*4-low {
				return false
			}
		}
		return true
	})
	return ans
}

func maxDistance22(side int, points [][]int, k int) int {
	a := make([]int, len(points))
	for i, p := range points {
		x, y := p[0], p[1]
		if x == 0 {
			a[i] = y
		} else if y == side {
			a[i] = side + x
		} else if x == side {
			a[i] = side*3 - y
		} else {
			a[i] = side*4 - x
		}
	}
	slices.Sort(a)

	ans := sort.Search(side, func(low int) bool {
		low++
		idx := make([]int, k)
		cur := a[0]
		for j, i := 1, 0; j < k; j++ {
			i += sort.Search(len(a)-i, func(j int) bool { return a[i+j] >= cur+low })
			if i == len(a) {
				return true
			}
			idx[j] = i
			cur = a[i]
		}
		if cur-a[0] <= side*4-low {
			return false
		}

		// 第一个指针移动到第二个指针的位置，就不用继续枚举了
		end0 := idx[1]
		for idx[0]++; idx[0] < end0; idx[0]++ {
			for j := 1; j < k; j++ {
				for a[idx[j]] < a[idx[j-1]]+low {
					idx[j]++
					if idx[j] == len(a) {
						return true
					}
				}
			}
			if a[idx[k-1]]-a[idx[0]] <= side*4-low {
				return false
			}
		}
		return true
	})
	return ans
}

func maxDistance21(side int, points [][]int, k int) int {
	a := make([]int, len(points))
	for i, p := range points {
		x, y := p[0], p[1]
		if x == 0 {
			a[i] = y
		} else if y == side {
			a[i] = side + x
		} else if x == side {
			a[i] = side*3 - y
		} else {
			a[i] = side*4 - x
		}
	}
	slices.Sort(a)

	ans := sort.Search(side, func(low int) bool {
		low++
		idx := make([]int, k)
		for {
			for j := 1; j < k; j++ {
				for a[idx[j]] < a[idx[j-1]]+low {
					idx[j]++
					if idx[j] == len(a) {
						return true
					}
				}
			}
			if a[idx[k-1]]-a[idx[0]] <= side*4-low {
				return false
			}
			idx[0]++
		}
	})
	return ans
}
