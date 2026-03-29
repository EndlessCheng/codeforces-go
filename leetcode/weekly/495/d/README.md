**前置题目**：[399. 除法求值](https://leetcode.cn/problems/evaluate-division/)，[我的题解](https://leetcode.cn/problems/evaluate-division/solutions/3805496/dai-quan-bing-cha-ji-pythonjavacgojsrust-a4a1/)。

在边权只有 $0$ 和 $1$ 的情况下，环的边权和是偶数，等价于有偶数个 $1$，等价于边权的**异或和**等于 $0$。

设 $x$ 和 $y$ 是某个环上的两个节点，我们可以把这个环拆分成从 $x$ 到 $y$ 的两条不同的简单路径。

由于 $s\oplus s = 0$，所以**当且仅当两条路径的异或和相同，环的边权和是偶数**。

如何验证是否满足要求？

- 如果添加的边连接了图的两个连通块，那么没有形成新的环，直接连边。
- 如果添加的边连接了同一个连通块中的两个点 $x$ 和 $y$，那么当且仅当添加的边的边权等于从 $x$ 到 $y$ 的路径的异或和（根据上面的结论，路径异或和是唯一的），我们才能添加这条边。

如何快速求出从 $x$ 到 $y$ 的路径的异或和？这可以用**带权并查集**维护。请先完成 [399. 除法求值](https://leetcode.cn/problems/evaluate-division/)。

对于本题，可以维护节点到其代表元 $\textit{root}$ 的路径异或和。从 $x$ 到 $y$ 的路径，可以拆分成先从 $x$ 到 $\textit{root}$，再从 $\textit{root}$ 到 $y$（重复走的边权会被异或抵消掉），这两条路径的异或和再计算异或，即为从 $x$ 到 $y$ 的路径异或和。

> **注**：即使边权范围改成 $[0,10^9]$，这个做法也是适用的。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
# 完整模板见 https://leetcode.cn/circle/discuss/mOr1u6/
class UnionFind:
    def __init__(self, n: int):
        self.fa = list(range(n))  # fa[x] 是 x 的代表元
        self.dis = [0] * n  # dis[x] = 从 x 到 fa[x] 的路径异或和

    def find(self, x: int) -> int:
        fa = self.fa
        if fa[x] != x:
            root = self.find(fa[x])
            self.dis[x] ^= self.dis[fa[x]]
            fa[x] = root
        return fa[x]

    def merge(self, from_: int, to: int, value: int) -> bool:
        x, y = self.find(from_), self.find(to)
        dis = self.dis
        if x == y:
            return dis[from_] ^ dis[to] == value
        dis[x] = value ^ dis[to] ^ dis[from_]
        self.fa[x] = y
        return True


class Solution:
    def numberOfEdgesAdded(self, n: int, edges: List[List[int]]) -> int:
        uf = UnionFind(n)
        ans = 0
        for x, y, w in edges:
            if uf.merge(x, y, w):
                ans += 1
        return ans
```

```java [sol-Java]
// 完整模板见 https://leetcode.cn/circle/discuss/mOr1u6/
class UnionFind {
    private final int[] fa; // fa[x] 是 x 的代表元
    private final int[] dis; // dis[x] = 从 x 到 fa[x] 的路径异或和

    public UnionFind(int n) {
        fa = new int[n];
        dis = new int[n];
        for (int i = 0; i < n; i++) {
            fa[i] = i;
        }
    }

    public int find(int x) {
        if (fa[x] != x) {
            int root = find(fa[x]);
            dis[x] ^= dis[fa[x]];
            fa[x] = root;
        }
        return fa[x];
    }

    public boolean merge(int from, int to, int value) {
        int x = find(from), y = find(to);
        if (x == y) {
            return (dis[from] ^ dis[to]) == value;
        }
        dis[x] = value ^ dis[to] ^ dis[from];
        fa[x] = y;
        return true;
    }
}

class Solution {
    public int numberOfEdgesAdded(int n, int[][] edges) {
        UnionFind uf = new UnionFind(n);
        int ans = 0;
        for (int[] e : edges) {
            if (uf.merge(e[0], e[1], e[2])) {
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
// 完整模板见 https://leetcode.cn/circle/discuss/mOr1u6/
// 根据题目用 UnionFind<int> uf(n) 或者 UnionFind<long long> uf(n) 初始化
template<typename T>
class UnionFind {
public:
    vector<int> fa; // fa[x] 是 x 的代表元
    vector<T> dis; // dis[x] = 从 x 到 fa[x] 的路径异或和

    UnionFind(int n) : fa(n), dis(n) {
        ranges::iota(fa, 0); // iota(fa.begin(), fa.end(), 0);
    }

    int find(int x) {
        if (fa[x] != x) {
            int root = find(fa[x]);
            dis[x] ^= dis[fa[x]];
            fa[x] = root;
        }
        return fa[x];
    }

    bool merge(int from, int to, T value) {
        int x = find(from), y = find(to);
        if (x == y) {
            return (dis[from] ^ dis[to]) == value;
        }
        dis[x] = value ^ dis[to] ^ dis[from];
        fa[x] = y;
        return true;
    }
};

class Solution {
public:
    int numberOfEdgesAdded(int n, vector<vector<int>>& edges) {
        int ans = 0;
        UnionFind<int> uf(n);
        for (auto& e : edges) {
            if (uf.merge(e[0], e[1], e[2])) {
                ans++;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
// 完整模板见 https://leetcode.cn/circle/discuss/mOr1u6/
type unionFind struct {
	fa  []int // fa[x] 是 x 的代表元
	dis []int // dis[x] = 从 x 到 fa[x] 的路径异或和
}

func newUnionFind(n int) unionFind {
	fa := make([]int, n)
	dis := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return unionFind{fa, dis}
}

func (u unionFind) find(x int) int {
	if u.fa[x] != x {
		root := u.find(u.fa[x])
		u.dis[x] ^= u.dis[u.fa[x]]
		u.fa[x] = root
	}
	return u.fa[x]
}

func (u unionFind) merge(from, to, value int) bool {
	x, y := u.find(from), u.find(to)
	if x == y {
		return u.dis[from]^u.dis[to] == value
	}
	u.dis[x] = value ^ u.dis[to] ^ u.dis[from]
	u.fa[x] = y
	return true
}

func numberOfEdgesAdded(n int, edges [][]int) (ans int) {
	uf := newUnionFind(n)
	for _, e := range edges {
		if uf.merge(e[0], e[1], e[2]) {
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面数据结构题单的「**§7.6 带权并查集（边权并查集）**」。

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
