package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// github.com/EndlessCheng/codeforces-go
func maxSumBST(root *TreeNode) (ans int) {
	var f func(o *TreeNode) (sum, min, max int, ok bool)
	f = func(o *TreeNode) (sum, min, max int, ok bool) {
		if o == nil {
			return 0, 1e9, -1e9, true
		}
		sum, min, max, ok = f(o.Left)
		sum1, min1, max1, ok1 := f(o.Right)
		if ok && max < o.Val && ok1 && min1 > o.Val {
			sum += o.Val + sum1
			if sum > ans {
				ans = sum
			}
			if o.Left == nil {
				min = o.Val
			}
			max = max1
			if o.Right == nil {
				max = o.Val
			}
			return
		}
		ok = false
		return
	}
	f(root)
	return
}

// 这种写法是错误的，见 [9,4,10,null,null,6,11] => 27
func maxSumBST2(root *TreeNode) (ans int) {
	const inf int = 2e9
	var dfs func(*TreeNode, int, int) int
	dfs = func(o *TreeNode, min, max int) int {
		if o == nil {
			return 0
		}
		if o.Val <= min || o.Val >= max {
			dfs(o, -inf, inf)
			return -inf
		}
		vl := dfs(o.Left, min, o.Val)
		vr := dfs(o.Right, o.Val, max)
		if vl == -inf || vr == -inf {
			return -inf
		}
		v := vl + vr + o.Val
		if v > ans {
			ans = v
		}
		return v
	}
	dfs(root, -inf, inf)
	return
}
