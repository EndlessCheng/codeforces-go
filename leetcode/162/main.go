package main

import (
	. "fmt"
)

var _ = Print

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func ifElseI(cond bool, a, b int) int {
	if cond {
		return a
	}
	return b
}
func ifElseS(cond bool, a, b string) string {
	if cond {
		return a
	}
	return b
}

const mod int = 1e9 + 7

func closedIsland(grid [][]int) (aa int) {
	v := [200][200]bool{}
	var dfs func(int, int) bool
	n := len(grid)
	m := len(grid[0])
	var d4 = [...][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	dfs = func(i, j int) bool {
		if i < 0 || i >= n || j < 0 || j >= m {
			return false
		}
		if grid[i][j] == 1 {
			return true
		}
		if v[i][j] {
			return true
		}
		v[i][j] = true
		// FIXME: 反思：太久没写简单的 DFS 了导致没遍历完连通分量就提前 return 了
		res := true
		for _, dir := range d4 {
			if !dfs(i+dir[0], j+dir[1]) {
				res = false
			}
		}
		return res
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 && !v[i][j] {
				if dfs(i, j) {
					aa++
				}
			}
		}
	}
	return
}

func reconstructMatrix(upper int, lower int, colsum []int) (ans [][]int) {
	ans = [][]int{}
	sum := 0
	for _, s := range colsum {
		sum += s
	}
	if sum != upper+lower {
		return
	}
	n := len(colsum)
	ans = [][]int{make([]int, n), make([]int, n)}
	// fill 2
	for i, sum := range colsum {
		if sum == 2 {
			ans[0][i] = 1
			ans[1][i] = 1
			upper--
			lower--
			// FIXME 反思：比赛时太进张，没有仔细想想其余 return [] 的情况
			if upper < 0 || lower < 0 {
				return [][]int{}
			}
		}
	}
	// fill 1
	for i, sum := range colsum {
		if sum == 1 {
			if upper > 0 {
				ans[0][i] = 1
				upper--
			} else if lower > 0 {
				ans[1][i] = 1
				lower--
			} else {
				return [][]int{}
			}
		}
	}
	return
}

func oddCells(n int, m int, indices [][]int) (cnt int) {
	x := [50][50]int{}
	for _, ind := range indices {
		row, col := ind[0], ind[1]
		for j := 0; j < m; j++ {
			x[row][j]++
		}
		for i := 0; i < n; i++ {
			x[i][col]++
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if x[i][j]&1 == 1 {
				cnt++
			}
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxScoreWords(words []string, letters []byte, score []int) int {
	n := len(words)
	wScores := make([]int, n)
	//cnts := make([][26]int, n)
	for i, w := range words {
		for _, c := range w {
			wScores[i] += score[c-'a']
			//cnts[i][c-'a']++
		}
	}
	cnt := make([]int, 26)
	for _, l := range letters {
		cnt[l-'a']++
	}
	ans := 0
outer:
	for i := 0; i < (1 << uint(n)); i++ {
		cnt2 := make([]int, 26)
		sum := 0
		copy(cnt2, cnt)
		for j := 0; j < n; j++ {
			if i>>uint(j)&1 == 1 {
				for _, c := range words[j] {
					cnt2[c-'a']--
					if cnt2[c-'a'] < 0 {
						continue outer
					}
				}
				sum += wScores[j]
			}
		}
		ans = max(ans, sum)
	}
	return ans
}

func main() {
	Println(closedIsland([][]int{
		{0, 0, 1, 1, 0, 1, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 0, 1, 1, 1, 0},
		{1, 0, 1, 1, 1, 0, 0, 1, 1, 0},
		{0, 1, 1, 0, 0, 0, 0, 1, 0, 1},
		{0, 0, 0, 0, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 1, 1, 1},
		{1, 0, 1, 0, 1, 1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 1, 0, 1, 0, 1},
		{1, 1, 1, 0, 1, 1, 0, 1, 1, 0},
	}))
	//Println(closedIsland([][]int{
	//	{0, 0, 1, 1, 0, 1, 0, 0, 1, 0},
	//	{1, 1, *, 1, 1, 0, 1, 1, 1, 0},
	//	{1, *, 1, 1, 1, 0, 0, 1, 1, 0},
	//	{0, 1, 1, 0, 0, 0, 0, 1, 0, 1},
	//	{0, 0, 0, 0, 0, 0, 1, 1, 1, 0},
	//	{0, 1, 0, 1, 0, 1, 0, 1, 1, 1},
	//	{1, 0, 1, 0, 1, 1, 0, 0, 0, 1},
	//	{1, 1, 1, 1, 1, 1, 0, 0, 0, 0},
	//	{1, 1, 1, 0, 0, 1, 0, 1, 0, 1},
	//	{1, 1, 1, 0, 1, 1, 0, 1, 1, 0},
	//}))

	//Println(reconstructMatrix(9, 2, []int{0, 1, 2, 0, 0, 0, 0, 0, 2, 1, 2, 1, 2}))
}
