package main

import (
	. "fmt"
	"sort"
)

var _ = Print

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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	has map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	f := FindElements{
		has: map[int]bool{},
	}
	root.Val = 0
	f.dfs(root)
	return f
}

func (f *FindElements) dfs(root *TreeNode) {
	f.has[root.Val] = true
	if root.Left != nil {
		root.Left.Val = root.Val*2 + 1
		f.dfs(root.Left)
	}
	if root.Right != nil {
		root.Right.Val = root.Val*2 + 2
		f.dfs(root.Right)
	}
}

func (f *FindElements) Find(target int) bool {
	return f.has[target]
}

func shiftGrid(grid [][]int, k int) [][]int {
	n := len(grid)
	m := len(grid[0])
	g := make([][]int, len(grid))
	for i := range g {
		g[i] = make([]int, len(grid[0]))
	}
	for i := range grid {
		for j := range grid[i] {
			newPos := i*m + j + k
			ni, nj := newPos/m%n, newPos%m
			g[ni][nj] = grid[i][j]
		}
	}
	return g
}
func maxSumDivThree(nums []int) int {
	n1 := []int{}
	n2 := []int{}
	ans := 0
	for _, v := range nums {
		if v%3 == 0 {
			ans += v
		} else if v%3 == 1 {
			n1 = append(n1, v)
		} else {
			n2 = append(n2, v)
		}
	}

	sort.Ints(n1)
	sort.Ints(n2)
	if len(n1) >= 6 {
		start := len(n1)%3 + 3
		for _, v := range n1[start:] {
			ans += v
		}
		n1 = n1[:start]
	}
	if len(n2) >= 6 {
		start := len(n2)%3 + 3
		for _, v := range n2[start:] {
			ans += v
		}
		n2 = n2[:start]
	}

	// choose
	old := ans
	min1 := len(n1)
	min2 := len(n2)
	for i := 0; i <= min1; i++ {
		for j := 0; j <= min2; j++ {
			// choose i n1 ans j n2
			if (i+2*j)%3 != 0 {
				continue
			}
			tmp := old
			cnt := 0
			if i > 0 {
				for ii := len(n1) - 1; ii >= 0; ii-- {
					tmp += n1[ii]
					cnt++
					if cnt == i {
						break
					}
				}
			}
			cnt = 0
			if j > 0 {
				for ii := len(n2) - 1; ii >= 0; ii-- {
					tmp += n2[ii]
					cnt++
					if cnt == j {
						break
					}
				}
			}
			//Println(tmp,i,j)
			ans = max(ans, tmp)
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minPushBox(grid [][]byte) int {
	n, m := len(grid), len(grid[0])

	var initI, initJ, finalI, finalJ int

	for i, gi := range grid {
		for j, gij := range gi {
			if gij == 'T' {
				finalI, finalJ = i, j
				grid[i][j] = '.'
			} else if gij == 'S' {
				initI, initJ = i, j
				grid[i][j] = '.'
			}
		}
	}

	dirOffset4 := [...][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	var ti, tj int
	var dfsGrid func(i, j int) bool
	var vis [20][20]bool

	dfsGrid = func(i, j int) bool {
		if i < 0 || i == n || j < 0 || j == m {
			return false
		}
		if grid[i][j] != '.' {
			return false
		}
		if i == ti && j == tj {
			return true
		}

		if vis[i][j] {
			return false
		}
		vis[i][j] = true

		res := false
		for _, dir := range dirOffset4 {
			if dfsGrid(i+dir[0], j+dir[1]) {
				res = true
			}
		}
		return res
	}

	canReach := func(si, sj, i, j int) bool {
		if i < 0 || i == n || j < 0 || j == m {
			return false
		}
		if grid[i][j] != '.' {
			return false
		}
		ti, tj = i, j

		//	var si, sj int
		//outer:
		//	for i, gi := range grid {
		//		for j, gij := range gi {
		//			if gij == 'S' {
		//				si, sj = i, j
		//				break outer
		//			}
		//		}
		//	}
		vis = [20][20]bool{}
		return dfsGrid(si, sj)
	}

	canPushTo := func(i, j int) bool {
		if i < 0 || i == n || j < 0 || j == m {
			return false
		}
		return grid[i][j] == '.'
	}

	const inf int = 1e8

	// min steps when box in (i,j)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i == n || j < 0 || j == m {
			return -1
		}
		if i == finalI && j == finalJ {
			return 0
		}
		if dp[i][j] != inf {
			return dp[i][j]
		}
		dp[i][j] = 2 * inf
		oldI, oldJ := initI, initJ
		ans := inf
		for _, dir := range dirOffset4 {
			if !canReach(oldI, oldJ, i-dir[0], j-dir[1]) {
				continue
			}
			if !canPushTo(i+dir[0], j+dir[1]) {
				continue
			}
			initI, initJ = i+dir[0], j+dir[1]
			//Println(initI, initJ)
			newAns := dfs(initI, initJ)
			if newAns == -1 {
				continue
			}
			ans = min(ans, newAns+1)
		}
		initI, initJ = oldI, oldJ
		if ans >= inf {
			ans = -1
		}
		dp[i][j] = ans
		return ans
	}

	var boxI, boxJ int
outer:
	for i, gi := range grid {
		for j, gij := range gi {
			if gij == 'B' {
				boxI, boxJ = i, j
				grid[i][j] = '.'
				break outer
			}
		}
	}

	return dfs(boxI, boxJ)
}

func main() {
	//Println(shiftGrid([][]int{
	//	{1, 2}, {3, 4}, {5, 6},
	//}, 100))

	//Println(maxSumDivThree([]int{3, 6, 5, 1, 8}))
	//Println(maxSumDivThree([]int{4}))
	//Println(maxSumDivThree([]int{1, 2, 3, 4, 4}))
	//Println(maxSumDivThree([]int{5, 2, 2, 2}))
	//Println(maxSumDivThree([]int{13,21,7,27,40,18,37,7,31,5}))
	//Println(maxSumDivThree([]int{	366,809,6,792,822,181,210,588,344,618,341,410,121,864,191,749,637,169,123,472,358,908,235,914,322,946,738,754,908,272,267,326,587,267,803,281,586,707,94,627,724,469,568,57,103,984,787,552,14,545,866,494,263,157,479,823,835,100,495,773,729,921,348,871,91,386,183,979,716,806,639,290,612,322,289,910,484,300,195,546,499,213,8,623,490,473,603,721,793,418,551,331,598,670,960,483,154,317,834,352}))

	toBytes := func(g [][]string) [][]byte {
		res := make([][]byte, len(g))
		for i := range res {
			res[i] = make([]byte, len(g[0]))
			for j := range res[i] {
				res[i][j] = g[i][j][0]
			}
		}
		return res
	}
	Println(minPushBox(toBytes([][]string{
		{"#", "#", "#", "#", "#", "#"},
		{"#", "T", "#", "#", "#", "#"},
		{"#", ".", ".", "B", ".", "#"},
		{"#", ".", "#", "#", ".", "#"},
		{"#", ".", ".", ".", "S", "#"},
		{"#", "#", "#", "#", "#", "#"},
	})))
	Println(minPushBox(toBytes([][]string{
		{"#", "#", "#", "#", "#", "#"},
		{"#", "T", "#", "#", "#", "#"},
		{"#", ".", ".", "B", ".", "#"},
		{"#", "#", "#", "#", ".", "#"},
		{"#", ".", ".", ".", "S", "#"},
		{"#", "#", "#", "#", "#", "#"},
	})))
	Println(minPushBox(toBytes([][]string{
		{"#", "#", "#", "#", "#", "#"},
		{"#", "T", ".", ".", "#", "#"},
		{"#", ".", "#", "B", ".", "#"},
		{"#", ".", ".", ".", ".", "#"},
		{"#", ".", ".", ".", "S", "#"},
		{"#", "#", "#", "#", "#", "#"},
	})))
	Println(minPushBox(toBytes([][]string{
		{"#", ".", ".", "#", "#", "#", "#", "#"},
		{"#", ".", ".", "T", "#", ".", ".", "#"},
		{"#", ".", ".", ".", "#", "B", ".", "#"},
		{"#", ".", ".", ".", ".", ".", ".", "#"},
		{"#", ".", ".", ".", "#", ".", "S", "#"},
		{"#", ".", ".", "#", "#", "#", "#", "#"},
	})))
}
