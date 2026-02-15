## 前置题目/知识点

1. [2791. 树中可以形成回文的路径数](https://leetcode.cn/problems/count-paths-that-can-form-a-palindrome-in-a-tree/)，[我的题解](https://leetcode.cn/problems/count-paths-that-can-form-a-palindrome-in-a-tree/solutions/2355288/yong-wei-yun-suan-chu-li-by-endlesscheng-n9ws/)。
2. 最近公共祖先（LCA），[模板讲解](https://leetcode.cn/problems/kth-ancestor-of-a-tree-node/solution/mo-ban-jiang-jie-shu-shang-bei-zeng-suan-v3rw/)。
3. [3515. 带权树中的最短路径](https://leetcode.cn/problems/shortest-path-in-a-weighted-tree/)，[我的题解](https://leetcode.cn/problems/shortest-path-in-a-weighted-tree/solutions/3649372/dfs-shi-jian-chuo-chai-fen-shu-zhuang-sh-h8q3/)。
4. [231. 2 的幂](https://leetcode.cn/problems/power-of-two/)，[我的题解](https://leetcode.cn/problems/power-of-two/solutions/2973442/yan-ge-zheng-ming-yi-xing-xie-fa-pythonj-h04o/)。

## 处理 query

根据 2791 题，问题等价于：

- 从 $x$ 到 $y$ 的路径上，至多有一个字母的出现次数是奇数。

用二进制数表示路径中的字母出现次数的奇偶性。如果二进制数是 $0$ 或者 $2^k\ (k\ge 0)$，那么至多有一个字母的出现次数是奇数。

定义 $\textit{XOR}[i]$ 表示从 $0$ 到 $i$ 的路径的字母出现次数的奇偶性（对应的二进制数）。

根据异或的性质，从 $x$ 到 $y$ 的路径的字母出现次数的奇偶性，等于如下三者的异或和：

- $0$ 到 $x$ 的路径的字母出现次数的奇偶性 $\textit{XOR}[x]$。
- $0$ 到 $y$ 的路径的字母出现次数的奇偶性 $\textit{XOR}[y]$。
- 上面两条路径异或后，把 $x$ 和 $y$ 的 $\textit{lca}$ 抵消掉了，所以要添加回来，即异或 $s[\textit{lca}]$ 对应的二进制数 `1 << (s[lca] - 'a')`。

$\textit{XOR}[i]$ 可以通过一次自顶向下的 DFS 求出。

## 处理 update

想一想，当我们修改 $s[x]$ 后，哪些 $\textit{XOR}[i]$ 会变？变成什么了？

修改 $s[x]$ 后，在子树 $x$ 中的节点 $i$ 的 $\textit{XOR}[i]$ 会变。

去掉原来的 $s[x]$，改成新的字母 $c$，用位运算解决，把这些 $\textit{XOR}[i]$ 都异或 `val = (1 << (s[x] - 'a')) ^ (1 << (c - 'a'))`。

根据 3515 题，求出节点 $x$ 的进出时间戳 $\textit{tin}[x]$ 和 $\textit{tout}[x]$，那么对于子树 $x$ 中的所有节点，其进出时间戳都在闭区间 $[\textit{tin}[x],\textit{tout}[x]]$ 中。

即把子树中的所有 $\textit{XOR}[i]$ 都异或同一个数，变成区间异或同一个数，即「区间更新」操作。

所以本题是「区间更新，单点查询」，最适合的数据结构是**差分树状数组**，转换成「单点更新，前缀查询」，做法同 3515 题。

代码实现时，差分树状数组保存的是「区间更新」操作的异或结果，没有保存初始值 $\textit{XOR}[i]$。获取 $\textit{XOR}[i]$ 修改后的值，可以用其初始值 $\textit{XOR}[i]$ 异或差分树状数组的前缀查询结果。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
# 模板来自我的题单 https://leetcode.cn/circle/discuss/mOr1u6/
class FenwickTree:
    def __init__(self, n: int):
        self.tree = [0] * (n + 1)  # 使用下标 1 到 n

    # a[i] ^= val
    # 1 <= i <= n
    # 时间复杂度 O(log n)
    def update(self, i: int, val: int) -> None:
        t = self.tree
        while i < len(t):
            t[i] ^= val
            i += i & -i

    # 计算前缀异或和 a[1] ^ ... ^ a[i]
    # 1 <= i <= n
    # 时间复杂度 O(log n)
    def pre(self, i: int) -> int:
        t = self.tree
        res = 0
        while i > 0:
            res ^= t[i]
            i &= i - 1
        return res


# 模板来自我的题单 https://leetcode.cn/circle/discuss/K0n2gO/
class LcaBinaryLifting:
    def __init__(self, edges: List[List[int]], s: List[int]):
        n = len(edges) + 1
        m = n.bit_length()
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        depth = [0] * n
        pa = [[-1] * n for _ in range(m)]
        tin = [0] * n  # DFS 时间戳
        tout = [0] * n
        clock = 0
        path_xor_from_root = [0] * n  # 从根开始的路径的字母出现次数的奇偶性
        path_xor_from_root[0] = 1 << s[0]

        def dfs(x: int, fa: int) -> None:
            pa[0][x] = fa
            nonlocal clock
            clock += 1
            tin[x] = clock
            for y in g[x]:
                if y != fa:
                    depth[y] = depth[x] + 1
                    path_xor_from_root[y] = path_xor_from_root[x] ^ (1 << s[y])
                    dfs(y, x)
            tout[x] = clock

        dfs(0, -1)

        for i in range(m - 1):
            for x in range(n):
                if (p := pa[i][x]) != -1:
                    pa[i + 1][x] = pa[i][p]

        self.depth = depth
        self.pa = pa
        self.tin = tin
        self.tout = tout
        self.path_xor_from_root = path_xor_from_root

    # 返回 node 的第 k 个祖先节点
    # 如果不存在，返回 -1
    def get_kth_ancestor(self, node: int, k: int) -> int:
        pa = self.pa
        for i in range(k.bit_length()):
            if k >> i & 1:
                node = pa[i][node]
                if node < 0:
                    return -1
        return node

    # 返回 x 和 y 的最近公共祖先
    def get_lca(self, x: int, y: int) -> int:
        if self.depth[x] > self.depth[y]:
            x, y = y, x
        # 使 y 和 x 在同一深度
        y = self.get_kth_ancestor(y, self.depth[y] - self.depth[x])
        if y == x:
            return x
        pa = self.pa
        for i in range(len(pa) - 1, -1, -1):
            px, py = pa[i][x], pa[i][y]
            if px != py:
                x, y = px, py  # 同时往上跳 2**i 步
        return pa[0][x]


class Solution:
    def palindromePath(self, n: int, edges: list[list[int]], s: str, queries: list[str]) -> list[bool]:
        ord_a = ord('a')
        t = [ord(ch) - ord_a for ch in s]  # 映射成 [0, 25] 中的整数

        g = LcaBinaryLifting(edges, t)
        tin = g.tin
        tout = g.tout
        path_xor_from_root = g.path_xor_from_root

        f = FenwickTree(n)  # 注意树状数组是异或运算
        ans = []

        for q in queries:
            op, x, y = q.split()
            x = int(x)
            if op[0] == 'u':
                c = ord(y) - ord_a
                val = (1 << t[x]) ^ (1 << c)  # 擦除旧的，换上新的
                t[x] = c
                # 子树 x 全部异或 val，转换成对区间 [tin[x], tout[x]] 的差分更新
                f.update(tin[x], val)
                f.update(tout[x] + 1, val)
            else:
                y = int(y)
                lca = g.get_lca(x, y)
                res = path_xor_from_root[x] ^ path_xor_from_root[y] ^ f.pre(tin[x]) ^ f.pre(tin[y]) ^ (1 << t[lca])
                ans.append(res & (res - 1) == 0)  # 至多一个字母的出现次数是奇数

        return ans
```

```java [sol-Java]
// 模板来自我的题单 https://leetcode.cn/circle/discuss/mOr1u6/
class FenwickTree {
    private final int[] tree;

    public FenwickTree(int n) {
        tree = new int[n + 1]; // 使用下标 1 到 n
    }

    // a[i] ^= val
    // 1 <= i <= n
    // 时间复杂度 O(log n)
    public void update(int i, int val) {
        for (; i < tree.length; i += i & -i) {
            tree[i] ^= val;
        }
    }

    // 求前缀异或和 a[1] ^ ... ^ a[i]
    // 1 <= i <= n
    // 时间复杂度 O(log n)
    public int pre(int i) {
        int res = 0;
        for (; i > 0; i &= i - 1) {
            res ^= tree[i];
        }
        return res;
    }
}

// 模板来自我的题单 https://leetcode.cn/circle/discuss/K0n2gO/
class LcaBinaryLifting {
    private final int[] depth;
    private final int[][] pa;
    private int clock = 0;

    public final int[] tin;
    public final int[] tout;
    public final int[] pathXorFromRoot;

    LcaBinaryLifting(int[][] edges, char[] s) {
        int n = edges.length + 1;
        int m = 32 - Integer.numberOfLeadingZeros(n); // n 的二进制长度
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, e -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            g[x].add(y);
            g[y].add(x);
        }

        depth = new int[n];
        pa = new int[m][n];
        tin = new int[n]; // DFS 时间戳
        tout = new int[n];
        pathXorFromRoot = new int[n]; // 从根开始的路径的字母出现次数的奇偶性
        pathXorFromRoot[0] = 1 << (s[0] - 'a');

        dfs(0, -1, g, s);

        for (int i = 0; i < m - 1; i++) {
            for (int x = 0; x < n; x++) {
                int p = pa[i][x];
                pa[i + 1][x] = p < 0 ? -1 : pa[i][p];
            }
        }
    }

    private void dfs(int x, int fa, List<Integer>[] g, char[] s) {
        pa[0][x] = fa;
        tin[x] = ++clock;
        for (int y : g[x]) {
            if (y != fa) {
                depth[y] = depth[x] + 1;
                pathXorFromRoot[y] = pathXorFromRoot[x] ^ (1 << (s[y] - 'a'));
                dfs(y, x, g, s);
            }
        }
        tout[x] = clock;
    }

    // 返回 node 的第 k 个祖先节点
    // 如果不存在，返回 -1
    private int getKthAncestor(int node, int k) {
        for (; k > 0 && node >= 0; k &= k - 1) {
            node = pa[Integer.numberOfTrailingZeros(k)][node];
        }
        return node;
    }

    // 返回 x 和 y 的最近公共祖先（节点编号从 0 开始）
    public int getLCA(int x, int y) {
        if (depth[x] > depth[y]) {
            int tmp = y;
            y = x;
            x = tmp;
        }
        // 使 y 和 x 在同一深度
        y = getKthAncestor(y, depth[y] - depth[x]);
        if (y == x) {
            return x;
        }
        for (int i = pa.length - 1; i >= 0; i--) {
            int px = pa[i][x], py = pa[i][y];
            if (px != py) {
                x = px;
                y = py; // 同时往上跳 2^i 步
            }
        }
        return pa[0][x];
    }
}

class Solution {
    public List<Boolean> palindromePath(int n, int[][] edges, String s, String[] queries) {
        char[] t = s.toCharArray();
        LcaBinaryLifting g = new LcaBinaryLifting(edges, t);
        FenwickTree f = new FenwickTree(n); // 注意树状数组是异或运算
        List<Boolean> ans = new ArrayList<>();

        for (String q : queries) {
            String[] parts = q.split(" ");
            int x = Integer.parseInt(parts[1]);
            if (parts[0].charAt(0) == 'u') {
                char c = parts[2].charAt(0);
                int val = (1 << (t[x] - 'a')) ^ (1 << (c - 'a')); // 擦除旧的，换上新的
                t[x] = c;
                // 子树 x 全部异或 val，转换成对区间 [tin[x], tout[x]] 的差分更新
                f.update(g.tin[x], val);
                f.update(g.tout[x] + 1, val);
            } else {
                int y = Integer.parseInt(parts[2]);
                int lca = g.getLCA(x, y);
                int res = g.pathXorFromRoot[x] ^ g.pathXorFromRoot[y] ^ f.pre(g.tin[x]) ^ f.pre(g.tin[y]) ^ (1 << (t[lca] - 'a'));
                ans.add((res & (res - 1)) == 0); // 至多一个字母的出现次数是奇数
            }
        }

        return ans;
    }
}
```

```cpp [sol-C++]
// 模板来自我的题单 https://leetcode.cn/circle/discuss/mOr1u6/
// 根据题目用 FenwickTree<int> t(n) 或者 FenwickTree<long long> t(n) 初始化
template<typename T>
class FenwickTree {
    vector<T> tree;

public:
    // 使用下标 1 到 n
    FenwickTree(int n) : tree(n + 1) {}

    // a[i] ^= val
    // 1 <= i <= n
    // 时间复杂度 O(log n)
    void update(int i, T val) {
        for (; i < tree.size(); i += i & -i) {
            tree[i] ^= val;
        }
    }

    // 求前缀异或和 a[1] ^ ... ^ a[i]
    // 1 <= i <= n
    // 时间复杂度 O(log n)
    T pre(int i) const {
        T res = 0;
        for (; i > 0; i &= i - 1) {
            res ^= tree[i];
        }
        return res;
    }
};

// 模板来自我的题单 https://leetcode.cn/circle/discuss/K0n2gO/
class LcaBinaryLifting {
    vector<int> depth;
    vector<vector<int>> pa;

public:
    vector<int> tin; // DFS 时间戳
    vector<int> tout;
    vector<int> path_xor_from_root; // 从根开始的路径的字母出现次数的奇偶性

    LcaBinaryLifting(vector<vector<int>>& edges, string& s) {
        int n = edges.size() + 1;
        int m = bit_width((uint32_t) n);
        vector<vector<int>> g(n);
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        depth.resize(n);
        pa.resize(m, vector<int>(n, -1));
        tin.resize(n);
        tout.resize(n);
        path_xor_from_root.resize(n);
        path_xor_from_root[0] = 1 << (s[0] - 'a');
        int clock = 0;

        auto dfs = [&](this auto&& dfs, int x, int fa) -> void {
            pa[0][x] = fa;
            tin[x] = ++clock;
            for (int y : g[x]) {
                if (y != fa) {
                    depth[y] = depth[x] + 1;
                    path_xor_from_root[y] = path_xor_from_root[x] ^ (1 << (s[y] - 'a'));
                    dfs(y, x);
                }
            }
            tout[x] = clock;
        };
        dfs(0, -1);

        for (int i = 0; i < m - 1; i++) {
            for (int x = 0; x < n; x++) {
                if (int p = pa[i][x]; p != -1) {
                    pa[i + 1][x] = pa[i][p];
                }
            }
        }
    }

    // 返回 node 的第 k 个祖先节点
    // 如果不存在，返回 -1
    int get_kth_ancestor(int node, int k) {
        for (; k > 0 && node >= 0; k &= k - 1) {
            node = pa[countr_zero((uint32_t) k)][node];
        }
        return node;
    }

    // 返回 x 和 y 的最近公共祖先（节点编号从 0 开始）
    int get_lca(int x, int y) {
        if (depth[x] > depth[y]) {
            swap(x, y);
        }
        y = get_kth_ancestor(y, depth[y] - depth[x]); // 使 y 和 x 在同一深度
        if (y == x) {
            return x;
        }
        for (int i = pa.size() - 1; i >= 0; i--) {
            int px = pa[i][x], py = pa[i][y];
            if (px != py) {
                x = px;
                y = py; // 同时往上跳 2^i 步
            }
        }
        return pa[0][x];
    }
};

class Solution {
public:
    vector<bool> palindromePath(int n, vector<vector<int>>& edges, string s, vector<string>& queries) {
        LcaBinaryLifting g(edges, s);
        FenwickTree<int> f(n); // 注意树状数组是异或运算
        vector<bool> ans;

        string op, x_str, y_str;
        for (auto& q : queries) {
            stringstream ss(q);
            ss >> op >> x_str >> y_str;
            int x = stoi(x_str);
            if (op[0] == 'u') {
                char c = y_str[0];
                int val = (1 << (s[x] - 'a')) ^ (1 << (c - 'a')); // 擦除旧的，换上新的
                s[x] = c;
                // 子树 x 全部异或 val，转换成对区间 [tin[x], tout[x]] 的差分更新
                f.update(g.tin[x], val);
                f.update(g.tout[x] + 1, val);
            } else {
                int y = stoi(y_str);
                int lca = g.get_lca(x, y);
                int res = g.path_xor_from_root[x] ^ g.path_xor_from_root[y] ^ f.pre(g.tin[x]) ^ f.pre(g.tin[y]) ^ (1 << (s[lca] - 'a'));
                ans.push_back((res & (res - 1)) == 0); // 至多一个字母的出现次数是奇数
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
// 模板来自我的题单 https://leetcode.cn/circle/discuss/mOr1u6/
type fenwick []int

func newFenwickTree(n int) fenwick {
	return make(fenwick, n+1) // 使用下标 1 到 n
}

// a[i] ^= val
// 1 <= i <= n
// 时间复杂度 O(log n)
func (f fenwick) update(i int, val int) {
	for ; i < len(f); i += i & -i {
		f[i] ^= val
	}
}

// 计算前缀异或和 a[1] ^ ... ^ a[i]
// 1 <= i <= n
// 时间复杂度 O(log n)
func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res ^= f[i]
	}
	return
}

func palindromePath(n int, edges [][]int, s string, queries []string) (ans []bool) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	mx := bits.Len(uint(n))
	pa := make([][16]int, n)
	dep := make([]int, n)
	timeIn := make([]int, n) // DFS 时间戳
	timeOut := make([]int, n)
	clock := 0
	pathXorFromRoot := make([]int, n) // 从根开始的路径中的字母奇偶性的集合
	pathXorFromRoot[0] = 1 << (s[0] - 'a')

	var dfs func(int, int)
	dfs = func(x, p int) {
		pa[x][0] = p
		clock++
		timeIn[x] = clock
		for _, y := range g[x] {
			if y != p {
				dep[y] = dep[x] + 1
				pathXorFromRoot[y] = pathXorFromRoot[x] ^ 1<<(s[y]-'a')
				dfs(y, x)
			}
		}
		timeOut[x] = clock
	}
	dfs(0, -1)

	for i := range mx - 1 {
		for x := range pa {
			p := pa[x][i]
			if p != -1 {
				pa[x][i+1] = pa[p][i]
			} else {
				pa[x][i+1] = -1
			}
		}
	}

	uptoDep := func(x, d int) int {
		for k := uint32(dep[x] - d); k > 0; k &= k - 1 {
			x = pa[x][bits.TrailingZeros32(k)]
		}
		return x
	}

	// 返回 x 和 y 的最近公共祖先
	getLCA := func(x, y int) int {
		if dep[x] > dep[y] {
			x, y = y, x
		}
		y = uptoDep(y, dep[x]) // 使 y 和 x 在同一深度
		if y == x {
			return x
		}
		for i := mx - 1; i >= 0; i-- {
			px, py := pa[x][i], pa[y][i]
			if px != py {
				x, y = px, py // 同时往上跳 2^i 步
			}
		}
		return pa[x][0]
	}

	// 上面全是模板，下面开始本题逻辑

	t := []byte(s)
	f := newFenwickTree(n) // 注意树状数组是异或运算
	for _, q := range queries {
		if q[0] == 'u' {
			x, _ := strconv.Atoi(q[7 : len(q)-2])
			c := q[len(q)-1]
			val := 1<<(t[x]-'a') ^ 1<<(c-'a') // 擦除旧的，换上新的
			t[x] = c
			// 子树 x 全部异或 val，转换成对区间 [timeIn[x], timeOut[x]] 的差分更新
			f.update(timeIn[x], val)
			f.update(timeOut[x]+1, val)
		} else {
			q = q[6:]
			i := strings.IndexByte(q, ' ')
			x, _ := strconv.Atoi(q[:i])
			y, _ := strconv.Atoi(q[i+1:])
			lca := getLCA(x, y)
			res := pathXorFromRoot[x] ^ pathXorFromRoot[y] ^ f.pre(timeIn[x]) ^ f.pre(timeIn[y]) ^ 1<<(t[lca]-'a')
			ans = append(ans, res&(res-1) == 0) // 至多一个字母的出现次数是奇数
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+q)\log n)$，其中 $q$ 是 $\textit{queries}$ 的长度。注意，把字符串转成整数的时间复杂度也是 $\mathcal{O}(\log n)$，因为字符串中的数字长度不超过 $n-1$ 的十进制长度。
- 空间复杂度：$\mathcal{O}(n\log n)$。返回值不计入。

## 专题训练

1. 树题单的「**§3.7 DFS 时间戳**」和「**§3.8 最近公共祖先（LCA）**」。
2. 数据结构题单的「**§1.4 状态压缩前缀和**」和「**§8.1 树状数组**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
