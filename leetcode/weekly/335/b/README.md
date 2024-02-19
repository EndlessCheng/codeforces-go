首先，利用 BFS，可以得到二叉树每一层的节点值之和，具体请看[【基础算法精讲 13】](https://www.bilibili.com/video/BV1hG4y1277i/)，本文用双数组实现。

BFS 的同时，把每一层的节点值之和保存到一个列表 $a$ 中，把 $a$ 排序后就可以得到第 $k$ 大。此外，也可以用快速选择得到第 $k$ 大。

如果 $k$ 大于 $a$ 的长度，返回 $-1$。

```py [sol-Python3]
class Solution:
    def kthLargestLevelSum(self, root: Optional[TreeNode], k: int) -> int:
        a = []
        q = [root]
        while q:
            s = 0
            tmp = q
            q = []
            for node in tmp:
                s += node.val
                if node.left:  q.append(node.left)
                if node.right: q.append(node.right)
            a.append(s)
        if k > len(a):
            return -1
        a.sort()
        return a[-k]
```

```java [sol-Java]
class Solution {
    public long kthLargestLevelSum(TreeNode root, int k) {
        List<Long> a = new ArrayList<>();
        List<TreeNode> q = List.of(root);
        while (!q.isEmpty()) {
            long sum = 0;
            List<TreeNode> tmp = q;
            q = new ArrayList<>();
            for (TreeNode node : tmp) {
                sum += node.val;
                if (node.left != null)  q.add(node.left);
                if (node.right != null) q.add(node.right);
            }
            a.add(sum);
        }
        int n = a.size();
        if (k > n) {
            return -1;
        }
        Collections.sort(a);
        return a.get(n - k);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long kthLargestLevelSum(TreeNode *root, int k) {
        vector<long long> a;
        vector<TreeNode*> q = {root};
        while (!q.empty()) {
            long long sum = 0;
            vector<TreeNode*> nxt;
            for (auto node : q) {
                sum += node->val;
                if (node->left)  nxt.push_back(node->left);
                if (node->right) nxt.push_back(node->right);
            }
            a.push_back(sum);
            q = move(nxt);
        }
        int n = a.size();
        if (k > n) {
            return -1;
        }
        ranges::sort(a);
        return a[n - k];
    }
};
```

```cpp [sol-C++ 快速选择]
class Solution {
public:
    long long kthLargestLevelSum(TreeNode *root, int k) {
        vector<long long> a;
        vector<TreeNode*> q = {root};
        while (!q.empty()) {
            long long sum = 0;
            vector<TreeNode*> nxt;
            for (auto node : q) {
                sum += node->val;
                if (node->left)  nxt.push_back(node->left);
                if (node->right) nxt.push_back(node->right);
            }
            a.push_back(sum);
            q = move(nxt);
        }
        int n = a.size();
        if (k > n) {
            return -1;
        }
        ranges::nth_element(a, a.begin() + (n - k));
        return a[n - k];
    }
};
```

```go [sol-Go]
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	a := []int64{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		sum := int64(0)
		tmp := q
		q = nil
		for _, node := range tmp {
			sum += int64(node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		a = append(a, sum)
	}
	n := len(a)
	if k > n {
		return -1
	}
	slices.Sort(a)
	return a[n-k]
}
```

```js [sol-JavaScript]
var kthLargestLevelSum = function(root, k) {
    const a = [];
    let q = [root];
    while (q.length) {
        let sum = 0;
        const tmp = q;
        q = [];
        for (const node of tmp) {
            sum += node.val;
            if (node.left)  q.push(node.left);
            if (node.right) q.push(node.right);
        }
        a.push(sum);
    }
    if (k > a.length) {
        return -1;
    }
    a.sort((a, b) => b - a);
    return a[k - 1];
};
```

```rust [sol-Rust]
use std::rc::Rc;
use std::cell::RefCell;

impl Solution {
    pub fn kth_largest_level_sum(root: Option<Rc<RefCell<TreeNode>>>, k: i32) -> i64 {
        let mut a = Vec::new();
        let mut cur = vec![root.unwrap()];
        while !cur.is_empty() {
            let mut sum = 0i64;
            let mut nxt = Vec::new();
            for node in cur {
                let mut x = node.borrow_mut();
                sum += x.val as i64;
                if let Some(left) = x.left.take() {
                    nxt.push(left);
                }
                if let Some(right) = x.right.take() {
                    nxt.push(right);
                }
            }
            cur = nxt;
            a.push(sum);
        }
        let k = k as usize;
        if k > a.len() {
            return -1;
        }
        a.sort_unstable();
        a[a.len() - k]
    }
}
```

### 复杂度分析

- 时间复杂度：排序写法 $\mathcal{O}(n\log n)$，快速选择写法 $\mathcal{O}(n)$，其中 $n$ 为二叉树的节点个数。
- 空间复杂度：$\mathcal{O}(n)$。

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
