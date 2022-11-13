这是一个经典问题，做法是置换环。

例如在数组 $[2,0,1,4,3]$ 中，$[2,0,1]$ 和 $[4,3]$ 分别是两个置换环，环与环之间是数字是不需要发生交换的，只会在环内发生交换。

怎么找到环呢？从第一个数开始，把这个数字当成下标去访问数组，不断循环直到回到这个数本身。

我们只需要计算每个环内需要多少次交换。对于每个环，交换次数为环的大小减一。

代码实现时需要离散化。

```py [sol1-Python3]
class Solution:
    def minimumOperations(self, root: Optional[TreeNode]) -> int:
        ans, q = 0, [root]
        while q:
            a = []
            tmp = q
            q = []
            for node in tmp:
                a.append(node.val)
                if node.left: q.append(node.left)
                if node.right: q.append(node.right)

            b = sorted(a)
            a = [bisect_left(b, v) for v in a]  # 离散化

            ans += len(a)
            vis = [False] * len(a)
            for v in a:
                if vis[v]: continue
                while not vis[v]:
                    vis[v] = True
                    v = a[v]
                ans -= 1
        return ans
```

```go [sol1-Go]
func minimumOperations(root *TreeNode) (ans int) {
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := len(q)
		a := make([]int, n)
		tmp := q
		q = nil
		for i, node := range tmp {
			a[i] = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		b := append([]int{}, a...)
		sort.Ints(b)
		for i, v := range a {
			a[i] = sort.SearchInts(b, v) // 离散化
		}

		ans += n
		vis := make([]bool, n)
		for _, v := range a {
			if !vis[v] {
				for ; !vis[v]; v = a[v] {
					vis[v] = true
				}
				ans--
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为二叉树的节点个数。对于完全二叉树而言，最后一层的节点个数可以达到 $O(n)$。
- 空间复杂度：$O(n)$。
