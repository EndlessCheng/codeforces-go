package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// github.com/EndlessCheng/codeforces-go
const mx int = 5e4 + 1

func canMerge(trees []*TreeNode) *TreeNode {
	isSub := [mx]bool{}      // 对于根节点，需要知道其是否为另一棵二叉搜索树的子节点
	roots := [mx]*TreeNode{} // 对于子节点，我们需要知道其数值所对应的二叉搜索树的根节点是哪个
	for _, rt := range trees {
		if rt.Left != nil {
			if isSub[rt.Left.Val] { // 由于二叉搜索树上不能有两个值相同的节点，所以 trees 中也不能有两个值相同的叶节点
				return nil
			}
			isSub[rt.Left.Val] = true
		}
		if rt.Right != nil {
			if isSub[rt.Right.Val] {
				return nil
			}
			isSub[rt.Right.Val] = true
		}
		roots[rt.Val] = rt
	}

	var root *TreeNode
	for _, rt := range trees {
		if !isSub[rt.Val] { // 根节点不应是另一棵二叉搜索树的子节点，否则二叉搜索树上会出现两个值相同的节点
			if root != nil { // 根节点应只有一个，否则会构成森林
				return nil
			}
			root = rt
		}
	}
	if root == nil { // 未找到根节点
		return nil
	}

	cnt := 0
	// 一边构建，一边判断是否合法
	var build func(*TreeNode, int, int) *TreeNode
	build = func(node *TreeNode, l, r int) *TreeNode {
		cnt++
		if node.Left != nil {
			if node.Left.Val <= l {
				return nil
			}
			if lo := roots[node.Left.Val]; lo != nil {
				node.Left = build(lo, l, node.Val)
				if node.Left == nil {
					return nil
				}
			}
		}
		if node.Right != nil {
			if node.Right.Val >= r {
				return nil
			}
			if ro := roots[node.Right.Val]; ro != nil {
				node.Right = build(ro, node.Val, r)
				if node.Right == nil {
					return nil
				}
			}
		}
		return node
	}
	root = build(root, 0, mx)
	if cnt == len(trees) { // 所有 trees[i] 均参与构建二叉搜索树
		return root
	}
	return nil
}
