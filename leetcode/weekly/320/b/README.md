首先，题目没说二叉搜索树是**平衡**的，最坏情况下二叉搜索树是一条链。

因此需要通过一次 [94. 二叉树的中序遍历](https://leetcode.cn/problems/binary-tree-inorder-traversal/) 得到有一个有序数组，再在数组上做二分查找。

在有序数组中求小于等于和大于等于，和 [34. 在排序数组中查找元素的第一个和最后一个位置](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/) 是一样的。

我在 [二分查找又死循环了？一个视频讲透二分本质！【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/) 这期视频中详细讲解了二分查找，欢迎观看。

```py [sol1-Python3]
class Solution:
    def closestNodes(self, root: Optional[TreeNode], queries: List[int]) -> List[List[int]]:
        a = []
        def dfs(o: Optional[TreeNode]) -> None:
            if o is None: return
            dfs(o.left)
            a.append(o.val)
            dfs(o.right)
        dfs(root)

        ans = []
        for q in queries:
            j = bisect_right(a, q)
            min = a[j - 1] if j else -1
            j = bisect_left(a, q)
            max = a[j] if j < len(a) else -1
            ans.append([min, max])
        return ans
```

```go [sol1-Go]
func closestNodes(root *TreeNode, queries []int) [][]int {
	a := []int{}
	var dfs func(*TreeNode)
	dfs = func(o *TreeNode) {
		if o == nil {
			return
		}
		dfs(o.Left)
		a = append(a, o.Val)
		dfs(o.Right)
	}
	dfs(root)

	ans := make([][]int, len(queries))
	for i, q := range queries {
		min, max := -1, -1
		// 这是怎么转换的，可以看我上面贴的视频链接
		j := sort.SearchInts(a, q+1) - 1
		if j >= 0 {
			min = a[j]
		}
		j = sort.SearchInts(a, q)
		if j < len(a) {
			max = a[j]
		}
		ans[i] = []int{min, max}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$O(q\log n)$，其中 $q$ 为 $\textit{queries}$ 的长度，$n$ 为二叉搜索树的节点个数。
- 空间复杂度：$O(n)$。
