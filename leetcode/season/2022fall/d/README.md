个人赛五道题目的 [视频讲解](https://www.bilibili.com/video/BV1zN4y1K762) 已出炉，欢迎点赞三连，在评论区分享你对这场比赛的看法~

---

定义状态 (当前节点，祖先节点开关 2 的切换次数的奇偶性，父节点是否切换了开关 3)，每个状态表示从当前状态出发，最少需要操作多少次开关，可以关闭子树所有节点的灯。

跑一个树形 DP。如果当前受到祖先节点的开关影响后，变成开灯状态，那么可以：

- 操作开关 1；
- 操作开关 2；
- 操作开关 3；
- 操作开关 123；
- 这四种操作取最小值。

如果变成关灯状态，那么可以：

- 不操作任何一个开关；
- 操作开关 12；
- 操作开关 13；
- 操作开关 23；
- 这四种操作取最小值。

```py [sol1-Python3]
class Solution:
    def closeLampInTree(self, root: TreeNode) -> int:
        @cache
        def f(node: TreeNode, switch2: bool, switch3: bool) -> int:
            if node is None:
                return 0
            if (node.val == 1) == (switch2 == switch3):
                res1 = f(node.left, switch2, False) + f(node.right, switch2, False) + 1
                res2 = f(node.left, not switch2, False) + f(node.right, not switch2, False) + 1
                res3 = f(node.left, switch2, True) + f(node.right, switch2, True) + 1
                res123 = f(node.left, not switch2, True) + f(node.right, not switch2, True) + 3
                return min(res1, res2, res3, res123)
            else:
                res0 = f(node.left, switch2, False) + f(node.right, switch2, False)
                res12 = f(node.left, not switch2, False) + f(node.right, not switch2, False) + 2
                res13 = f(node.left, switch2, True) + f(node.right, switch2, True) + 2
                res23 = f(node.left, not switch2, True) + f(node.right, not switch2, True) + 2
                return min(res0, res12, res13, res23)
        return f(root, False, False)
```

```go [sol1-Go]
func closeLampInTree(root *TreeNode) (ans int) {
	type pair struct {
		node             *TreeNode
		switch2, switch3 bool
	}
	dp := map[pair]int{}
	var f func(*TreeNode, bool, bool) int
	f = func(node *TreeNode, switch2, switch3 bool) int {
		if node == nil {
			return 0
		}
		p := pair{node, switch2, switch3}
		if res, ok := dp[p]; ok {
			return res
		}
		if node.Val == 1 == (switch2 == switch3) {
			res1 := f(node.Left, switch2, false) + f(node.Right, switch2, false) + 1
			res2 := f(node.Left, !switch2, false) + f(node.Right, !switch2, false) + 1
			res3 := f(node.Left, switch2, true) + f(node.Right, switch2, true) + 1
			r123 := f(node.Left, !switch2, true) + f(node.Right, !switch2, true) + 3
			dp[p] = min(res1, res2, res3, r123)
		} else {
			res0 := f(node.Left, switch2, false) + f(node.Right, switch2, false)
			res12 := f(node.Left, !switch2, false) + f(node.Right, !switch2, false) + 2
			res13 := f(node.Left, switch2, true) + f(node.Right, switch2, true) + 2
			res23 := f(node.Left, !switch2, true) + f(node.Right, !switch2, true) + 2
			dp[p] = min(res0, res12, res13, res23)
		}
		return dp[p]
	}
	return f(root, false, false)
}

func min(a, b, c, d int) int {
	if b < a {
		a = b
	}
	if c < a {
		a = c
	}
	if d < a {
		a = d
	}
	return a
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为二叉树的节点个数。
- 空间复杂度：$O(n)$。
