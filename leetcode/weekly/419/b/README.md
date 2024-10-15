根据完美二叉树的定义，一棵高为 $h$ 的完美二叉树，从上往下，每一层恰好有 $1,2,4,8,\cdots,2^{h-1}$ 个节点，所以子树大小为

$$
1+2+4+8+\cdots+2^{h-1} = 2^h-1
$$

由于知道高度就知道了子树大小，DFS 只需返回子树高度。

DFS 的同时，用一个数组 $\textit{hs}$ 维护合法子树的高度。

分类讨论：

- 如果当前节点是空节点，返回 $0$。
- 否则，判断左右子树的高度是否都不为 $-1$ 且相同，如果不是，返回 $-1$；如果是，把子树高度加一加入 $\textit{hs}$，然后返回子树高度加一。

DFS 结束后：

- 设 $m$ 为 $\textit{hs}$ 的大小。
- 如果 $k>m$，返回 $-1$。
- 否则设 $\textit{hs}$ 从大到小排序后的下标为 $k-1$ 的元素为 $h$（也可以是从小到大排序后的下标为 $m-k$ 的元素），返回 $2^h - 1$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1zU2zYiEa4/)，欢迎点赞关注~

### 写法一：快速排序/快速选择

```py [sol-Python3]
class Solution:
    def kthLargestPerfectSubtree(self, root: Optional[TreeNode], k: int) -> int:
        hs = []
        def dfs(node: Optional[TreeNode]) -> int:
            if node is None:
                return 0
            left_h = dfs(node.left)
            right_h = dfs(node.right)
            if left_h < 0 or left_h != right_h:
                return -1  # 不合法
            hs.append(left_h + 1)
            return left_h + 1
        dfs(root)

        if k > len(hs):
            return -1
        hs.sort()
        return (1 << hs[-k]) - 1
```

```java [sol-Java]
class Solution {
    public int kthLargestPerfectSubtree(TreeNode root, int k) {
        List<Integer> hs = new ArrayList<>();
        dfs(root, hs);

        if (k > hs.size()) {
            return -1;
        }
        Collections.sort(hs);
        return (1 << hs.get(hs.size() - k)) - 1;
    }

    private int dfs(TreeNode node, List<Integer> hs) {
        if (node == null) {
            return 0;
        }
        int leftH = dfs(node.left, hs);
        int rightH = dfs(node.right, hs);
        if (leftH < 0 || leftH != rightH) {
            return -1; // 不合法
        }
        hs.add(leftH + 1);
        return leftH + 1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int kthLargestPerfectSubtree(TreeNode* root, int k) {
        vector<int> hs;
        auto dfs = [&](auto&& dfs, TreeNode* node) -> int {
            if (node == nullptr) {
                return 0;
            }
            int left_h = dfs(dfs, node->left);
            int right_h = dfs(dfs, node->right);
            if (left_h < 0 || left_h != right_h) {
                return -1; // 不合法
            }
            hs.push_back(left_h + 1);
            return left_h + 1;
        };
        dfs(dfs, root);

        if (k > hs.size()) {
            return -1;
        }
        // ranges::sort(hs);
        ranges::nth_element(hs, hs.end() - k);
        return (1 << hs[hs.size() - k]) - 1;
    }
};
```

```go [sol-Go]
func kthLargestPerfectSubtree(root *TreeNode, k int) int {
	hs := []int{}
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftH := dfs(node.Left)
		rightH := dfs(node.Right)
		if leftH < 0 || leftH != rightH {
			return -1 // 不合法
		}
		hs = append(hs, leftH+1)
		return leftH + 1
	}
	dfs(root)

	if k > len(hs) {
		return -1
	}
	slices.Sort(hs)
	return 1<<hs[len(hs)-k] - 1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m\log m)$ 或者 $\mathcal{O}(n)$，其中 $n$ 是二叉树的节点个数，$m$ 是 $\textit{hs}$ 的长度。如果使用快速选择，则时间复杂度为 $\mathcal{O}(n)$，见 C++ 代码中的 `ranges::nth_element`。
- 空间复杂度：$\mathcal{O}(n)$。递归需要 $\mathcal{O}(n)$ 的栈空间。

### 写法二：计数排序

统计每个高度的出现次数。

叶子的高度改成 $0$，充分利用数组空间。

```py [sol-Python3]
class Solution:
    def kthLargestPerfectSubtree(self, root: Optional[TreeNode], k: int) -> int:
        cnt = [0] * 10
        def dfs(node: Optional[TreeNode]) -> int:
            if node is None:
                return -1
            left_h = dfs(node.left)
            right_h = dfs(node.right)
            if left_h == -2 or left_h != right_h:
                return -2  # 不合法
            cnt[left_h + 1] += 1
            return left_h + 1
        dfs(root)

        for i in range(len(cnt) - 1, -1, -1):
            c = cnt[i]
            if c >= k:
                return (1 << (i + 1)) - 1
            k -= c
        return -1
```

```java [sol-Java]
class Solution {
    public int kthLargestPerfectSubtree(TreeNode root, int k) {
        int[] cnt = new int[10];
        dfs(root, cnt);

        for (int i = cnt.length - 1; i >= 0; i--) {
            int c = cnt[i];
            if (c >= k) {
                return (1 << (i + 1)) - 1;
            }
            k -= c;
        }
        return -1;
    }

    private int dfs(TreeNode node, int[] cnt) {
        if (node == null) {
            return -1;
        }
        int leftH = dfs(node.left, cnt);
        int rightH = dfs(node.right, cnt);
        if (leftH == -2 || leftH != rightH) {
            return -2; // 不合法
        }
        cnt[leftH + 1]++;
        return leftH + 1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int kthLargestPerfectSubtree(TreeNode* root, int k) {
        int cnt[10]{};
        auto dfs = [&](auto&& dfs, TreeNode* node) -> int {
            if (node == nullptr) {
                return -1;
            }
            int left_h = dfs(dfs, node->left);
            int right_h = dfs(dfs, node->right);
            if (left_h == -2 || left_h != right_h) {
                return -2; // 不合法
            }
            cnt[left_h + 1]++;
            return left_h + 1;
        };
        dfs(dfs, root);

        for (int i = 9; i >= 0; i--) {
            int c = cnt[i];
            if (c >= k) {
                return (1 << (i + 1)) - 1;
            }
            k -= c;
        }
        return -1;
    }
};
```

```go [sol-Go]
func kthLargestPerfectSubtree(root *TreeNode, k int) int {
	cnt := [10]int{}
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		leftH := dfs(node.Left)
		rightH := dfs(node.Right)
		if leftH == -2 || leftH != rightH {
			return -2 // 不合法
		}
		cnt[leftH+1]++
		return leftH + 1
	}
	dfs(root)

	for i := len(cnt) - 1; i >= 0; i-- {
		c := cnt[i]
		if c >= k {
			return 1<<(i+1) - 1
		}
		k -= c
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是二叉树的节点个数。
- 空间复杂度：$\mathcal{O}(n)$。递归需要 $\mathcal{O}(n)$ 的栈空间。

更多相似题目，见下面链表二叉树题单中的「**§2.3 自底向上 DFS**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
