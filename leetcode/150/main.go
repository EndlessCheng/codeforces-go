package main

import (
	. "fmt"
	"index/suffixarray"
	"reflect"
	"unsafe"
)

// 41 / 3865  Virtual	19 0:36:51	   0:02:45 0:12:03 0:33:12 0:36:51

var _ = Print

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

func countCharacters(words []string, chars string) int {
	cnt := [26]int{}
	for _, c := range chars {
		cnt[c-'a']++
	}
	ans := 0
	for _, w := range words {
		cnt2 := [26]int{}
		for _, c := range w {
			cnt2[c-'a']++
		}
		ok := true
		for i, c2 := range cnt2 {
			if c2 > cnt[i] {
				ok = false
				break
			}
		}
		if ok {
			ans += len(w)
		}
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	maxVal := int(-2e9)
	maxPos := 1
	q := [2][]*TreeNode{{}, {root}}
	for i := 1; len(q[i&1]) > 0; i++ {
		s := 0
		q[(i+1)&1] = []*TreeNode{}
		for _, o := range q[i&1] {
			s += o.Val
			if o.Left != nil {
				q[(i+1)&1] = append(q[(i+1)&1], o.Left)
			}
			if o.Right != nil {
				q[(i+1)&1] = append(q[(i+1)&1], o.Right)
			}
		}
		if s > maxVal {
			maxVal = s
			maxPos = i
		}
	}
	return maxPos
}

func maxDistance(grid [][]int) int {
	n := len(grid)
	type p struct{ i, j int }
	lands := []p{}
	vis := [100][100]bool{}
	left := n * n
	for i, g := range grid {
		for j, c := range g {
			if c == 1 {
				lands = append(lands, p{i, j})
				vis[i][j] = true
				left--
			}
		}
	}
	if left == 0 || left == n*n {
		return -1
	}

	dirOffset4 := [...][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	searchDirOffset4 := func(n, centerI, centerJ, dis int) {
		for i, dir := range dirOffset4 {
			dx := dirOffset4[(i+1)%4][0] - dir[0]
			dy := dirOffset4[(i+1)%4][1] - dir[1]
			x := centerI + dir[0]*dis
			y := centerJ + dir[1]*dis
			for _i := 0; _i < dis; _i++ {
				if x >= 0 && x < n && y >= 0 && y < n {
					if !vis[x][y] {
						vis[x][y] = true
						left--
					}
				}
				x += dx
				y += dy
			}
		}
	}
	ans := 1
	for ; left > 0; ans++ {
		for _, land := range lands {
			searchDirOffset4(n, land.i, land.j, ans)
		}
	}
	return ans - 1
}

func lastSubstring(s string) string {
	sa := *(*[]int)(unsafe.Pointer(reflect.ValueOf(suffixarray.New([]byte(s))).Elem().FieldByName("sa").UnsafeAddr()))
	return s[sa[len(sa)-1]:]
}

func main() {
	Println(countCharacters([]string{"cat", "bt", "hat", "tree"}, "atach"))
	Println(countCharacters([]string{"hello", "world", "leetcode"}, "welldonehoneyr"))

	Println(maxDistance([][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}))
	Println(maxDistance([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}}))

	Println(lastSubstring("abab"))
	Println(lastSubstring("leetcode"))
}
