package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"sort"
)

func main() {
	toBytes := func(g [][]string) [][]byte {
		n, m := len(g), len(g[0])
		bytes := make([][]byte, n)
		for i := range bytes {
			bytes[i] = make([]byte, m)
			for j := range bytes[i] {
				bytes[i][j] = g[i][j][0]
			}
		}
		return bytes
	}

	_ = MustBuildTreeNode

	_ = []interface{}{toBytes, ListNode{}, TreeNode{}}
}

// LC 41
func firstMissingPositive(a []int) int {
	n := len(a)
	for i, v := range a {
		for 0 < v && v <= n && v != a[v-1] {
			a[i], a[v-1] = a[v-1], a[i]
			v = a[i]
		}
	}
	for i, v := range a {
		if i+1 != v {
			return i + 1
		}
	}
	return n + 1
}

// LC99
func recoverTree(root *TreeNode) {
	nodes := []*TreeNode{}
	var f func(o *TreeNode)
	f = func(o *TreeNode) {
		if o == nil {
			return
		}
		f(o.Left)
		nodes = append(nodes, o)
		f(o.Right)
	}
	f(root)
	so := make([]*TreeNode, len(nodes))
	copy(so, nodes)
	sort.Slice(so, func(i, j int) bool { return so[i].Val < so[j].Val })
	do := []*TreeNode{}
	for i, o := range nodes {
		if o.Val != so[i].Val {
			do = append(do, o)
		}
	}
	do[0].Val, do[1].Val = do[1].Val, do[0].Val
}

// LC 124
func maxPathSum(root *TreeNode) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ans := int(-1e18)
	var f func(*TreeNode) int
	f = func(o *TreeNode) int {
		if o == nil {
			return -1e18
		}
		l := max(f(o.Left), 0)
		r := max(f(o.Right), 0)
		ans = max(ans, o.Val+l+r)
		return o.Val + max(l, r)
	}
	f(root)
	return ans
}

// LC 332
func findItinerary(tickets [][]string) []string {
	g := map[string][]string{}
	for _, p := range tickets {
		g[p[0]] = append(g[p[0]], p[1])
	}
	for _, vs := range g {
		sort.Strings(vs)
	}

	path := make([]string, 0, len(tickets)+1)
	var f func(string)
	f = func(v string) {
		for len(g[v]) > 0 {
			w := g[v][0]
			g[v] = g[v][1:]
			f(w)
		}
		path = append(path, v)
	}
	f("JFK")

	for i, j := 0, len(path)-1; i < j; i++ {
		path[i], path[j] = path[j], path[i]
		j--
	}
	return path
}
