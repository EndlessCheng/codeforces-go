package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"math"
)

func maxLevelSum1(root *TreeNode) (ans int) {
	maxSum := math.MinInt
	q := []*TreeNode{root}

	for level := 1; q != nil; level++ {
		tmp := q
		q = nil
		s := 0

		for _, node := range tmp {
			s += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		if s > maxSum {
			maxSum = s
			ans = level
		}
	}

	return ans
}

func maxLevelSum(root *TreeNode) (ans int) {
	rowSum := []int{}

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if len(rowSum) == level {
			rowSum = append(rowSum, node.Val)
		} else {
			rowSum[level] += node.Val
		}

		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}

	dfs(root, 0)

	for i, s := range rowSum {
		if s > rowSum[ans] {
			ans = i
		}
	}
	return ans + 1
}
