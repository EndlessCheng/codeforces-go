package main

import (
	. "fmt"
	"sort"
)

func collections() {
	ifElseI := func(cond bool, r1, r2 int) int {
		if cond {
			return r1
		}
		return r2
	}
	ifElseS := func(cond bool, r1, r2 string) string {
		if cond {
			return r1
		}
		return r2
	}

	const mod int = 1e9 + 7
	dirOffset4 := [...][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	_ = []interface{}{Print, ifElseI, ifElseS, dirOffset4, min, max}
}

func shiftGrid(grid [][]int, k int) [][]int {
	n, m := len(grid), len(grid[0])
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, m)
	}
	for i, gi := range grid {
		for j, gij := range gi {
			newPos := i*m + j + k
			ni, nj := newPos/m%n, newPos%m
			g[ni][nj] = gij
		}
	}
	return g
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type FindElements struct {
	has map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	f := FindElements{map[int]bool{}}
	f.dfs(root, 0)
	return f
}
func (f *FindElements) dfs(o *TreeNode, v int) {
	if o != nil {
		f.has[v] = true
		f.dfs(o.Left, v*2+1)
		f.dfs(o.Right, v*2+2)
	}
}
func (f *FindElements) Find(target int) bool {
	return f.has[target]
}

func maxSumDivThree(nums []int) int {
	arr1 := []int{}
	arr2 := []int{}
	ans := 0
	for _, v := range nums {
		if v%3 == 0 {
			ans += v
		} else if v%3 == 1 {
			arr1 = append(arr1, v)
		} else {
			arr2 = append(arr2, v)
		}
	}
	sort.Ints(arr1)
	sort.Ints(arr2)
	n1, n2 := len(arr1), len(arr2)

	if len(arr1) >= 6 {
		start := n1%3 + 3
		for _, v := range arr1[start:] {
			ans += v
		}
		arr1 = arr1[:start]
	}
	if len(arr2) >= 6 {
		start := n2%3 + 3
		for _, v := range arr2[start:] {
			ans += v
		}
		arr2 = arr2[:start]
	}
	n1, n2 = len(arr1), len(arr2)

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	base := ans
	for i := 0; i <= n1; i++ {
		for j := 0; j <= n2; j++ {
			// choose i arr1 and j arr2
			if (i+2*j)%3 != 0 {
				continue
			}
			tmpSum := base
			for ii := n1 - 1; ii >= n1-i; ii-- {
				tmpSum += arr1[ii]
			}
			for ii := n2 - 1; ii >= n2-j; ii-- {
				tmpSum += arr2[ii]
			}
			ans = max(ans, tmpSum)
		}
	}
	return ans
}

func minPushBox(grid [][]byte) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	n, m := len(grid), len(grid[0])
	isValid := func(i, j int) bool {
		return i >= 0 && i < n && j >= 0 && j < m && grid[i][j] != '#'
	}

	findPos := func(c byte) (int, int) {
		for i, gi := range grid {
			for j, gij := range gi {
				if gij == c {
					return i, j
				}
			}
		}
		panic(c)
	}
	initI, initJ := findPos('S')
	finalI, finalJ := findPos('T')

	dirOffset4 := [...][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	var ti, tj int
	var dfsGrid func(i, j int) bool
	var vis [20][20]bool

	dfsGrid = func(i, j int) bool {
		if !isValid(i, j) {
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
		if !isValid(i, j) {
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
		if !isValid(i, j) {
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
			if !isValid(i+dir[0], j+dir[1]) {
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

	boxI, boxJ := findPos('B')
	return dfs(boxI, boxJ)
}

func main() {
	//Println(shiftGrid([][]int{
	//	{1, 2}, {3, 4}, {5, 6},
	//}, 100))
	//
	//Println(maxSumDivThree([]int{3, 6, 5, 1, 8}))
	//Println(maxSumDivThree([]int{4}))
	//Println(maxSumDivThree([]int{1, 2, 3, 4, 4}))
	//Println(maxSumDivThree([]int{5, 2, 2, 2}))
	//Println(maxSumDivThree([]int{13, 21, 7, 27, 40, 18, 37, 7, 31, 5}))
	//Println(maxSumDivThree([]int{366, 809, 6, 792, 822, 181, 210, 588, 344, 618, 341, 410, 121, 864, 191, 749, 637, 169, 123, 472, 358, 908, 235, 914, 322, 946, 738, 754, 908, 272, 267, 326, 587, 267, 803, 281, 586, 707, 94, 627, 724, 469, 568, 57, 103, 984, 787, 552, 14, 545, 866, 494, 263, 157, 479, 823, 835, 100, 495, 773, 729, 921, 348, 871, 91, 386, 183, 979, 716, 806, 639, 290, 612, 322, 289, 910, 484, 300, 195, 546, 499, 213, 8, 623, 490, 473, 603, 721, 793, 418, 551, 331, 598, 670, 960, 483, 154, 317, 834, 352}))

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
