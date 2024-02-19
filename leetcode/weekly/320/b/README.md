首先，题目没说二叉搜索树是**平衡**的，最坏情况下这棵树是一条链，此时单次询问的复杂度是 $\mathcal{O}(n)$ 的，其中 $n$ 为二叉搜索树的节点个数。

为了加快回答询问的速度，可以通过一次 [94. 二叉树的中序遍历](https://leetcode.cn/problems/binary-tree-inorder-traversal/) 得到有一个**严格递增**数组 $a$，再在 $a$ 上做二分查找，就可以做到单次询问 $\mathcal{O}(\log n)$ 的时间了。

关于二分查找的原理，请看[【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

设 $j$ 是大于等于 $q=\textit{queries}_i$ 的第一个数的下标，如果不存在则 $j=n$。

对于 $\textit{max}_i$：

- 如果 $j<n$，那么 $\textit{max}_i = a[j]$。
- 否则 $\textit{max}_i = -1$。

对于 $\textit{min}_i$：

- 如果 $j<n$ 且 $a[j] = q$，那么 $\textit{min}_i = a[j]$。
- 否则如果 $j>0$，那么 $\textit{min}_i = a[j-1]$。
- 否则 $\textit{min}_i = -1$。

```py [sol-Python3]
class Solution:
    def closestNodes(self, root: Optional[TreeNode], queries: List[int]) -> List[List[int]]:
        a = []
        def dfs(node: Optional[TreeNode]) -> None:
            if node is None:
                return
            dfs(node.left)
            a.append(node.val)
            dfs(node.right)
        dfs(root)

        n = len(a)
        ans = []
        for q in queries:
            j = bisect_left(a, q)
            mx = a[j] if j < n else -1
            if j == n or a[j] != q:  # a[j]>q, a[j-1]<q
                j -= 1
            mn = a[j] if j >= 0 else -1
            ans.append([mn, mx])
        return ans
```

```java [sol-Java]
class Solution {
    public List<List<Integer>> closestNodes(TreeNode root, List<Integer> queries) {
        List<Integer> arr = new ArrayList<>();
        dfs(root, arr);

        int n = arr.size();
        int[] a = new int[n];
        for (int i = 0; i < n; i++) {
            a[i] = arr.get(i); // 转成数组，效率更高
        }

        List<List<Integer>> ans = new ArrayList<>(queries.size()); // 预分配空间
        for (int q : queries) {
            int j = lowerBound(a, q);
            int mx = j == n ? -1 : a[j];
            if (j == n || a[j] != q) { // a[j]>q, a[j-1]<q
                j--;
            }
            int mn = j < 0 ? -1 : a[j];
            ans.add(List.of(mn, mx));
        }
        return ans;
    }

    private void dfs(TreeNode node, List<Integer> a) {
        if (node == null) {
            return;
        }
        dfs(node.left, a);
        a.add(node.val);
        dfs(node.right, a);
    }

    // 见 https://www.bilibili.com/video/BV1AP41137w7/
    private int lowerBound(int[] a, int target) {
        int left = -1, right = a.length; // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            int mid = (left + right) >>> 1; // 比 /2 快
            if (a[mid] >= target) {
                right = mid; // 范围缩小到 (left, mid)
            } else {
                left = mid; // 范围缩小到 (mid, right)
            }
        }
        return right;
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<int> a;

    void dfs(TreeNode *node) {
        if (node == nullptr) {
            return;
        }
        dfs(node->left);
        a.push_back(node->val);
        dfs(node->right);
    };

public:
    vector<vector<int>> closestNodes(TreeNode *root, vector<int> &queries) {
        dfs(root);
        int n = a.size();
        vector<vector<int>> ans;
        for (int q : queries) {
            int j = ranges::lower_bound(a, q) - a.begin();
            int mx = j < n ? a[j] : -1;
            if (j == n || a[j] != q) { // a[j]>q, a[j-1]<q
                j--;
            }
            int mn = j >= 0 ? a[j] : -1;
            ans.push_back({mn, mx});
        }
        return ans;
    }
};
```

```go [sol-Go]
func closestNodes(root *TreeNode, queries []int) [][]int {
	a := []int{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		a = append(a, node.Val)
		dfs(node.Right)
	}
	dfs(root)

	ans := make([][]int, len(queries))
	for i, q := range queries {
		mn, mx := -1, -1
		j, ok := slices.BinarySearch(a, q)
		if j < len(a) {
			mx = a[j]
		}
		if !ok { // a[j]>q, a[j-1]<q
			j--
		}
		if j >= 0 {
			mn = a[j]
		}
		ans[i] = []int{mn, mx}
	}
	return ans
}
```

```js [sol-JavaScript]
var closestNodes = function(root, queries) {
    const a = [];
    function dfs(node) {
        if (node === null) {
            return;
        }
        dfs(node.left);
        a.push(node.val);
        dfs(node.right);
    }
    dfs(root);

    const n = a.length;
    const ans = [];
    for (const q of queries) {
        let j = lowerBound(a, q);
        const mx = j < n ? a[j] : -1;
        if (j === n || a[j] !== q) { // a[j]>q, a[j-1]<q
            j--;
        }
        const mn = j >= 0 ? a[j] : -1;
        ans.push([mn, mx]);
    }
    return ans;
};

// 见 https://www.bilibili.com/video/BV1AP41137w7/
var lowerBound = function(a, target) {
    let left = -1, right = a.length; // 开区间 (left, right)
    while (left + 1 < right) { // 区间不为空
        const mid = Math.floor((left + right) / 2);
        if (a[mid] >= target) {
            right = mid; // 范围缩小到 (left, mid)
        } else {
            left = mid; // 范围缩小到 (mid, right)
        }
    }
    return right;
}
```

```rust [sol-Rust]
use std::rc::Rc;
use std::cell::RefCell;

impl Solution {
    pub fn closest_nodes(root: Option<Rc<RefCell<TreeNode>>>, queries: Vec<i32>) -> Vec<Vec<i32>> {
        fn dfs(node: Option<&Rc<RefCell<TreeNode>>>, a: &mut Vec<i32>) {
            if let Some(x) = node {
                let x = x.borrow();
                dfs(x.left.as_ref(), a);
                a.push(x.val);
                dfs(x.right.as_ref(), a);
            }
        }
        let mut a = Vec::new();
        dfs(root.as_ref(), &mut a);

        let n = a.len();
        let mut ans = Vec::new();
        for q in queries {
            let mut j = a.partition_point(|&x| x < q);
            let mx = if j < n { a[j] } else { -1 };
            let mn = if j < n && a[j] == q { q } else if j > 0 { a[j - 1] } else { -1 };
            ans.push(vec![mn, mx]);
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + q\log n)$，其中 $n$ 为二叉搜索树的节点个数，$q$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。返回值的空间不计入。

[【题单】二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
