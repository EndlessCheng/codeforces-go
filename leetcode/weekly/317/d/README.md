既然是求树的高度，我们可以先跑一遍 DFS，求出每棵子树的高度 $\textit{height}$（这里定义成最长路径的**节点数**）。

然后再 DFS 一遍这棵树，同时维护当前节点深度 $\textit{depth}$（从 $0$ 开始），以及删除当前子树后剩余部分的树的高度 $\textit{restH}$（这里定义成最长路径的**边数**）。

具体做法如下：

- 往左走，递归前算一下从根节点到当前节点右子树最深节点的长度，即 $\textit{depth} + \textit{height}[\textit{node}.\textit{right}]$，与 $\textit{restH}$ 取最大值，然后往下递归；
- 往右走，递归前算一下从根节点到当前节点左子树最深节点的长度，即 $\textit{depth} + \textit{height}[\textit{node}.\textit{left}]$，与 $\textit{restH}$ 取最大值，然后往下递归。

每个节点的答案即为递归到该节点时的 $\textit{restH}$ 值。

代码实现时可以直接把答案覆盖到 $\textit{queries}$ 数组中。

```py [sol1-Python3]
class Solution:
    def treeQueries(self, root: Optional[TreeNode], queries: List[int]) -> List[int]:
        height = defaultdict(int)  # 每棵子树的高度
        def get_height(node: Optional[TreeNode]) -> int:
            if node is None: return 0
            height[node] = 1 + max(get_height(node.left), get_height(node.right))
            return height[node]
        get_height(root)

        res = [0] * (len(height) + 1)  # 每个节点的答案
        def dfs(node: Optional[TreeNode], depth: int, rest_h: int) -> None:
            if node is None: return
            depth += 1
            res[node.val] = rest_h
            dfs(node.left, depth, max(rest_h, depth + height[node.right]))
            dfs(node.right, depth, max(rest_h, depth + height[node.left]))
        dfs(root, -1, 0)

        for i, q in enumerate(queries):
            queries[i] = res[q]
        return queries
```

```java [sol1-Java]
class Solution {
    private Map<TreeNode, Integer> height = new HashMap<>(); // 每棵子树的高度
    private int[] res; // 每个节点的答案

    public int[] treeQueries(TreeNode root, int[] queries) {
        getHeight(root);
        height.put(null, 0); // 简化 dfs 的代码，这样不用写 getOrDefault
        res = new int[height.size()];
        dfs(root, -1, 0);
        for (var i = 0; i < queries.length; i++)
            queries[i] = res[queries[i]];
        return queries;
    }

    private int getHeight(TreeNode node) {
        if (node == null) return 0;
        var h = 1 + Math.max(getHeight(node.left), getHeight(node.right));
        height.put(node, h);
        return h;
    }

    private void dfs(TreeNode node, int depth, int restH) {
        if (node == null) return;
        ++depth;
        res[node.val] = restH;
        dfs(node.left, depth, Math.max(restH, depth + height.get(node.right)));
        dfs(node.right, depth, Math.max(restH, depth + height.get(node.left)));
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    vector<int> treeQueries(TreeNode *root, vector<int> &queries) {
        unordered_map<TreeNode*, int> height; // 每棵子树的高度
        function<int(TreeNode*)> get_height = [&](TreeNode *node) -> int {
            return node ? height[node] = 1 + max(get_height(node->left), get_height(node->right)) : 0;
        };
        get_height(root);

        int res[height.size() + 1]; // 每个节点的答案
        function<void(TreeNode*, int, int)> dfs = [&](TreeNode *node, int depth, int rest_h) {
            if (node == nullptr) return;
            ++depth;
            res[node->val] = rest_h;
            dfs(node->left, depth, max(rest_h, depth + height[node->right]));
            dfs(node->right, depth, max(rest_h, depth + height[node->left]));
        };
        dfs(root, -1, 0);

        for (auto &q : queries) q = res[q];
        return queries;
    }
};
```

```go [sol1-Go]
func treeQueries(root *TreeNode, queries []int) []int {
	height := map[*TreeNode]int{} // 每棵子树的高度
	var getHeight func(*TreeNode) int
	getHeight = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		height[node] = 1 + max(getHeight(node.Left), getHeight(node.Right))
		return height[node]
	}
	getHeight(root)

	res := make([]int, len(height)+1) // 每个节点的答案
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, depth, restH int) {
		if node == nil {
			return
		}
		depth++
		res[node.Val] = restH
		dfs(node.Left, depth, max(restH, depth+height[node.Right]))
		dfs(node.Right, depth, max(restH, depth+height[node.Left]))
	}
	dfs(root, -1, 0)

	for i, q := range queries {
		queries[i] = res[q]
	}
	return queries
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为二叉树的节点个数。
- 空间复杂度：$O(n)$。
