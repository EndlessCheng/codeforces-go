package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// 94. 二叉树的中序遍历
func inorderTraversal(root *TreeNode) (ans []int) {
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)              // 左
		ans = append(ans, node.Val) // 根（这行代码移到前面就是前序，移到后面就是后序）
		dfs(node.Right)             // 右
	}
	dfs(root)
	return
}

// 108. 将有序数组转换为二叉搜索树
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	m := len(nums) / 2
	return &TreeNode{
		Val:   nums[m],
		Left:  sortedArrayToBST(nums[:m]),
		Right: sortedArrayToBST(nums[m+1:]),
	}
}

func balanceBST(root *TreeNode) *TreeNode {
	nums := inorderTraversal(root)
	return sortedArrayToBST(nums)
}
