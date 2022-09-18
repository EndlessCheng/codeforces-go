下午 2 点在 B 站直播讲周赛和双周赛的题目，[欢迎关注](https://space.bilibili.com/206214/dynamic)~

---

#### 方法一：BFS

BFS 这棵树，对于奇数层，直接交换层里面的所有元素值（交换的是元素值，不是节点）。

```py [sol1-Python3]
class Solution:
    def reverseOddLevels(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        q, level = [root], 0
        while q[0].left:
            q = list(chain.from_iterable((node.left, node.right) for node in q))
            if level == 0:
                for i in range(len(q) // 2):
                    x, y = q[i], q[len(q) - 1 - i]
                    x.val, y.val = y.val, x.val
            level ^= 1
        return root
```

```go [sol1-Go]
func reverseOddLevels(root *TreeNode) *TreeNode {
	q := []*TreeNode{root}
	for level := 0; q[0].Left != nil; level ^= 1 {
		next := make([]*TreeNode, 0, len(q)*2)
		for _, node := range q {
			next = append(next, node.Left, node.Right)
		}
		q = next
		if level == 0 {
			for i, n := 0, len(q); i < n/2; i++ {
				x, y := q[i], q[n-1-i]
				x.Val, y.Val = y.Val, x.Val
			}
		}
	}
	return root
}
```

#### 方法二：DFS

依然是交换值的思路，通过同时递归左右子树实现。

```py [sol2-Python3]
class Solution:
    def reverseOddLevels(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        def dfs(node1: Optional[TreeNode], node2: Optional[TreeNode], is_odd_level: bool) -> None:
            if node1 is None: return
            if is_odd_level: node1.val, node2.val = node2.val, node1.val
            dfs(node1.left, node2.right, not is_odd_level)
            dfs(node1.right, node2.left, not is_odd_level)
        dfs(root.left, root.right, True)
        return root
```

```go [sol2-Go]
func dfs(node1 *TreeNode, node2 *TreeNode, isOddLevel bool) {
	if node1 == nil {
		return
	}
	if isOddLevel {
		node1.Val, node2.Val = node2.Val, node1.Val
	}
	dfs(node1.Left, node2.Right, !isOddLevel)
	dfs(node1.Right, node2.Left, !isOddLevel)
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	dfs(root.Left, root.Right, true)
	return root
}
```

