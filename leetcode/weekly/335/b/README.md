下午两点【biIibiIi@灵茶山艾府】直播讲题，记得关注哦~

---

### 前置知识：二叉树 BFS

见[【基础算法精讲 13】](https://www.bilibili.com/video/BV1hG4y1277i/)。

### 思路

BFS 二叉树，记录每一层的节点值之和，排序后取第 $k$ 大（也可以用快速选择）。

```py [sol1-Python3]
class Solution:
    def kthLargestLevelSum(self, root: Optional[TreeNode], k: int) -> int:
        q = [root]
        sum = []
        while q:
            tmp, s = q, 0
            q = []
            for node in tmp:
                s += node.val
                if node.left: q.append(node.left)
                if node.right: q.append(node.right)
            sum.append(s)
        sum.sort()  # 也可以用快速选择
        return -1 if len(sum) < k else sum[-k]
```

```go [sol1-Go]
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	q := []*TreeNode{root}
	sum := []int{}
	for len(q) > 0 {
		tmp, s := q, 0
		q = nil
		for _, node := range tmp {
			s += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		sum = append(sum, s)
	}
	n := len(sum)
	if n < k {
		return -1
	}
	sort.Ints(sum) // 也可以用快速选择
	return int64(sum[n-k])
}
```

### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为二叉树的节点个数。
- 空间复杂度：$O(n)$。
