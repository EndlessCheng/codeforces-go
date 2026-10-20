## 方法一：统计邻居颜色

先判断无解的情况。根据题意，对于节点 $x$ 来说：

1. $x$ 至多有三个邻居（父节点和左右儿子）。
2. $x$ 的邻居的颜色，不能与 $x$ 的颜色相同。
3. 由于 $x$ 至多有一个父节点，所以在 $x$ 的邻居中，「上一个颜色」的个数至多一个。R 的上一个颜色是 W，W 的上一个颜色是 B，B 的上一个颜色是 R。

如果满足上述要求，那么还要满足什么要求，节点 $x$ 才能作为树的根？

1. $x$ 至多有两个邻居。
2. 由于 $x$ 没有父节点，所以在 $x$ 的邻居中，「上一个颜色」的个数必须是零。

如果有多个点都可以作为树的根，会发生什么情况？

1. 当其中一个点是根的时候，其余点会有儿子变成其父亲，导致颜色不符合要求。所以**只能有一个合法根节点**。
2. 实际上，如果有多个根，那么考察两个根节点子树的交集，必然有一个节点有两个父节点，但这已经在无解的情况三中判断了。所以无需讨论有多个根的情况。

[本题视频讲解](https://www.bilibili.com/video/BV1N2nu6xEPW/)，欢迎点赞关注~

### 写法一

```py [sol-Python3]
mp = {'W': 0, 'R': 1, 'B': 2}

class Solution:
    def findColoredRoot(self, edges: list[list[int]], colors: str) -> int:
        n = len(colors)
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        ans = -1
        for x, to in enumerate(g):
            # 一个点至多三个邻居（父节点和左右儿子）
            if len(to) > 3:
                return -1

            cnt = [0] * 3
            for y in to:
                cnt[mp[colors[y]]] += 1

            color_x = mp[colors[x]]
            # 不能有邻居和 x 颜色相同，或者 x 的父节点的颜色个数超过 1
            if cnt[color_x] > 0 or cnt[color_x - 1] > 1:
                return -1

            # 判断 x 能否是根
            if len(to) < 3 and cnt[color_x - 1] == 0:
                ans = x

        return ans
```

```java [sol-Java]
class Solution {
    private static final int[] mp = new int[128];

    static {
        mp['W'] = 0;
        mp['R'] = 1;
        mp['B'] = 2;
    }

    public int findColoredRoot(int[][] edges, String colors) {
        char[] s = colors.toCharArray();
        int n = s.length;
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, _ -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            g[x].add(y);
            g[y].add(x);
        }

        int ans = -1;
        for (int x = 0; x < n; x++) {
            List<Integer> to = g[x];
            // 一个点至多三个邻居（父节点和左右儿子）
            if (to.size() > 3) {
                return -1;
            }

            int[] cnt = new int[3];
            for (int y : to) {
                cnt[mp[s[y]]]++;
            }

            int colorX = mp[s[x]];
            // 不能有邻居和 x 颜色相同，或者 x 的父节点的颜色个数超过 1
            if (cnt[colorX] > 0 || cnt[(colorX + 2) % 3] > 1) {
                return -1;
            }

            // 判断 x 能否是根
            if (to.size() < 3 && cnt[(colorX + 2) % 3] == 0) {
                ans = x;
            }
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
    int map_color(char c) {
        if (c == 'W') {
            return 0;
        }
        if (c == 'R') {
            return 1;
        }
        return 2;
    }

public:
    int findColoredRoot(vector<vector<int>>& edges, string colors) {
        int n = colors.size();
        vector<vector<int>> g(n);
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        int ans = -1;
        for (int x = 0; x < n; x++) {
            auto& to = g[x];
            // 一个点至多三个邻居（父节点和左右儿子）
            if (to.size() > 3) {
                return -1;
            }

            int cnt[3]{};
            for (int y : to) {
                cnt[map_color(colors[y])]++;
            }

            int color_x = map_color(colors[x]);
            // 不能有邻居和 x 颜色相同，或者 x 的父节点的颜色个数超过 1
            if (cnt[color_x] > 0 || cnt[(color_x + 2) % 3] > 1) {
                return -1;
            }

            // 判断 x 能否是根
            if (to.size() < 3 && cnt[(color_x + 2) % 3] == 0) {
                ans = x;
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
var mp = [...]int{'W': 0, 'R': 1, 'B': 2}

func findColoredRoot(edges [][]int, colors string) int {
	n := len(colors)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	ans := -1
	for x, to := range g {
		// 一个点至多三个邻居（父节点和左右儿子）
		if len(to) > 3 {
			return -1
		}

		cnt := [3]int{}
		for _, y := range to {
			cnt[mp[colors[y]]]++
		}

		colorX := mp[colors[x]]
		// 不能有邻居和 x 颜色相同，或者 x 的父节点的颜色个数超过 1
		if cnt[colorX] > 0 || cnt[(colorX+2)%3] > 1 {
			return -1
		}

		// 判断 x 能否是根
		if len(to) < 3 && cnt[(colorX+2)%3] == 0 {
			ans = x
		}
	}

	return ans
}
```

### 写法二

不创建 $g$，改为直接统计每个节点的邻居颜色个数。

由于内存更加紧凑，部分语言可以明显加快速度。

```py [sol-Python3]
mp = {'W': 0, 'R': 1, 'B': 2}

class Solution:
    def findColoredRoot(self, edges: list[list[int]], colors: str) -> int:
        n = len(colors)
        cnts = [[0] * 3 for _ in range(n)]
        for x, y in edges:
            cnts[x][mp[colors[y]]] += 1
            cnts[y][mp[colors[x]]] += 1

        ans = -1
        for x, cnt in enumerate(cnts):
            sz = sum(cnt)
            color_x = mp[colors[x]]
            if sz > 3 or cnt[color_x] > 0 or cnt[color_x - 1] > 1:
                return -1

            if sz < 3 and cnt[color_x - 1] == 0:
                ans = x

        return ans
```

```java [sol-Java]
class Solution {
    private static final int[] mp = new int[128];

    static {
        mp['W'] = 0;
        mp['R'] = 1;
        mp['B'] = 2;
    }

    public int findColoredRoot(int[][] edges, String colors) {
        char[] s = colors.toCharArray();
        int n = s.length;
        int[][] cnts = new int[n][3];
        for (int[] e : edges) {
            int x = e[0], y = e[1];
            cnts[x][mp[s[y]]]++;
            cnts[y][mp[s[x]]]++;
        }

        int ans = -1;
        for (int x = 0; x < n; x++) {
            int[] cnt = cnts[x];
            int sz = cnt[0] + cnt[1] + cnt[2];
            int colorX = mp[s[x]];
            if (sz > 3 || cnt[colorX] > 0 || cnt[(colorX + 2) % 3] > 1) {
                return -1;
            }

            if (sz < 3 && cnt[(colorX + 2) % 3] == 0) {
                ans = x;
            }
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
    int map_color(char c) {
        if (c == 'W') {
            return 0;
        }
        if (c == 'R') {
            return 1;
        }
        return 2;
    }

public:
    int findColoredRoot(vector<vector<int>>& edges, string colors) {
        int n = colors.size();
        vector<array<int, 3>> cnts(n);
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            cnts[x][map_color(colors[y])]++;
            cnts[y][map_color(colors[x])]++;
        }

        int ans = -1;
        for (int x = 0; x < n; x++) {
            auto& cnt = cnts[x];
            int sz = cnt[0] + cnt[1] + cnt[2];
            int color_x = map_color(colors[x]);
            if (sz > 3 || cnt[color_x] > 0 || cnt[(color_x + 2) % 3] > 1) {
                return -1;
            }

            if (sz < 3 && cnt[(color_x + 2) % 3] == 0) {
                ans = x;
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
var mp = [...]int{'W': 0, 'R': 1, 'B': 2}

func findColoredRoot(edges [][]int, colors string) int {
	n := len(colors)
	cnts := make([][3]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		cnts[x][mp[colors[y]]]++
		cnts[y][mp[colors[x]]]++
	}

	ans := -1
	for x, cnt := range cnts {
		sz := cnt[0] + cnt[1] + cnt[2]
		colorX := mp[colors[x]]
		if sz > 3 || cnt[colorX] > 0 || cnt[(colorX+2)%3] > 1 {
			return -1
		}

		if sz < 3 && cnt[(colorX+2)%3] == 0 {
			ans = x
		}
	}

	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{colors}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：统计入度出度

对于一条边的两个节点：

- 如果两个节点颜色相同，那么直接返回 $-1$。
- 如果两个节点颜色不同，根据「分层颜色模式」，其中一个点一定是另一个点的父亲：红点是蓝点的父亲，蓝点是白点的父亲，白点是红点的父亲。据此，我们可以直接确定每条边的方向，从而统计每个节点的**入度**（父节点个数）和**出度**（儿子个数）。

对于二叉树，每个节点的入度至多为 $1$，出度至多为 $2$。如果每个节点都满足该要求，那么：

- 由于树只有 $n-1$ 条边，所以入度之和至多为 $n-1$，所以至少有一个点没有入度，即至少有一个合法根节点。
- 如果有至少两个合法根节点，那么考察两个根节点子树的交集，必然有一个节点有两个父节点，不满足入度至多为 $1$，所以至多有一个合法根节点。

综上所述，只要每个节点的入度至多为 $1$，出度至多为 $2$，那么这棵树恰好有一个合法根节点。

### 优化前

```py [sol-Python3]
child_color = {'R': 'B', 'B': 'W', 'W': 'R'}

class Solution:
    def findColoredRoot(self, edges: list[list[int]], colors: str) -> int:
        n = len(colors)
        in_deg = [0] * n
        out_deg = [0] * n

        for x, y in edges:
            cx, cy = colors[x], colors[y]
            if cx == cy:
                return -1

            # 确定谁是谁的父节点
            if child_color[cx] != cy:
                x, y = y, x

            # x 是 y 的父节点
            out_deg[x] += 1
            in_deg[y] += 1
            if out_deg[x] > 2 or in_deg[y] > 1:
                return -1

        for i, deg in enumerate(in_deg):
            if deg == 0:
                return i
```

```java [sol-Java]
class Solution {
    private static final char[] childColor = new char[128];

    static {
        childColor['R'] = 'B';
        childColor['B'] = 'W';
        childColor['W'] = 'R';
    }

    public int findColoredRoot(int[][] edges, String colors) {
        char[] s = colors.toCharArray();
        int n = s.length;
        int[] inDeg = new int[n];
        int[] outDeg = new int[n];

        for (int[] e : edges) {
            int x = e[0], y = e[1];
            char cx = s[x], cy = s[y];
            if (cx == cy) {
                return -1;
            }

            // 确定谁是谁的父节点
            if (childColor[cx] != cy) {
                int tmp = x;
                x = y;
                y = tmp;
            }

            // x 是 y 的父节点
            outDeg[x]++;
            inDeg[y]++;
            if (outDeg[x] > 2 || inDeg[y] > 1) {
                return -1;
            }
        }

        for (int i = 0; ; i++) {
            if (inDeg[i] == 0) {
                return i;
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
    char get_child_color(char ch) {
        if (ch == 'R') {
            return 'B';
        }
        if (ch == 'B') {
            return 'W';
        }
        return 'R';
    }

public:
    int findColoredRoot(vector<vector<int>>& edges, string colors) {
        int n = colors.size();
        vector<int> in_deg(n), out_deg(n);

        for (auto& e : edges) {
            int x = e[0], y = e[1];
            char cx = colors[x], cy = colors[y];
            if (cx == cy) {
                return -1;
            }

            // 确定谁是谁的父节点
            if (get_child_color(cx) != cy) {
                swap(x, y);
            }

            // x 是 y 的父节点
            out_deg[x]++;
            in_deg[y]++;
            if (out_deg[x] > 2 || in_deg[y] > 1) {
                return -1;
            }
        }

        for (int i = 0; ; i++) {
            if (in_deg[i] == 0) {
                return i;
            }
        }
    }
};
```

```go [sol-Go]
var childColor = [...]byte{'R': 'B', 'B': 'W', 'W': 'R'}

func findColoredRoot(edges [][]int, colors string) int {
	n := len(colors)
	inDeg := make([]int, n)
	outDeg := make([]int, n)

	for _, e := range edges {
		x, y := e[0], e[1]
		cx, cy := colors[x], colors[y]
		if cx == cy {
			return -1
		}

		// 确定谁是谁的父节点
		if childColor[cx] != cy {
			x, y = y, x
		}

		// x 是 y 的父节点
		outDeg[x]++
		inDeg[y]++
		if outDeg[x] > 2 || inDeg[y] > 1 {
			return -1
		}
	}

	for i := 0; ; i++ {
		if inDeg[i] == 0 {
			return i
		}
	}
}
```

### 优化：一次遍历

把树上每个节点编号异或一遍，再把所有儿子的编号异或一遍。由于同一个值异或两次等于 $0$，所以最终异或和恰好等于根节点的值。

这个技巧来自 [136. 只出现一次的数字](https://leetcode.cn/problems/single-number/)，请看 [我的题解](https://leetcode.cn/problems/single-number/solutions/2481594/li-yong-yi-huo-de-xing-zhi-fu-ti-dan-pyt-oizc/)。

树上每个节点编号的异或和，即 $0$ 到 $n-1$ 的异或和，可以 $\mathcal{O}(1)$ 计算，原理见 [1486 题解](https://leetcode.cn/problems/xor-operation-in-an-array/solutions/2793723/o1-gong-shi-tui-dao-pythonjavaccgojsrust-le23/) 的「$0$ 到 $n$ 的异或和」。

```py [sol-Python3]
child_color = {'R': 'B', 'B': 'W', 'W': 'R'}

# 返回 0 到 n 的异或和
def xor_n(n: int) -> int:
    return (n, 1, n + 1, 0)[n % 4]

class Solution:
    def findColoredRoot(self, edges: list[list[int]], colors: str) -> int:
        n = len(colors)
        in_deg = [0] * n
        out_deg = [0] * n
        root = xor_n(n - 1)

        for x, y in edges:
            cx, cy = colors[x], colors[y]
            if cx == cy:
                return -1

            # 确定谁是谁的父节点
            if child_color[cx] != cy:
                x, y = y, x

            # x 是 y 的父节点
            out_deg[x] += 1
            in_deg[y] += 1
            if out_deg[x] > 2 or in_deg[y] > 1:
                return -1

            root ^= y

        return root
```

```java [sol-Java]
class Solution {
    private static final char[] childColor = new char[128];

    static {
        childColor['R'] = 'B';
        childColor['B'] = 'W';
        childColor['W'] = 'R';
    }

    public int findColoredRoot(int[][] edges, String colors) {
        char[] s = colors.toCharArray();
        int n = s.length;
        byte[] inDeg = new byte[n]; // 优化空间，减少 cache miss
        byte[] outDeg = new byte[n];
        int root = xorN(n - 1);

        for (int[] e : edges) {
            int x = e[0], y = e[1];
            char cx = s[x], cy = s[y];
            if (cx == cy) {
                return -1;
            }

            // 确定谁是谁的父节点
            if (childColor[cx] != cy) {
                int tmp = x;
                x = y;
                y = tmp;
            }

            // x 是 y 的父节点
            outDeg[x]++;
            inDeg[y]++;
            if (outDeg[x] > 2 || inDeg[y] > 1) {
                return -1;
            }

            root ^= y;
        }

        return root;
    }

    // 返回 0 到 n 的异或和
    private int xorN(int n) {
        return switch (n % 4) {
            case 0 -> n;
            case 1 -> 1;
            case 2 -> n + 1;
            default -> 0;
        };
    }
}
```

```cpp [sol-C++]
class Solution {
    // 返回 0 到 n 的异或和
    int xor_n(int n) {
        switch (n % 4) {
            case 0: return n;
            case 1: return 1;
            case 2: return n + 1;
            default: return 0;
        }
    }

    char get_child_color(char ch) {
        if (ch == 'R') {
            return 'B';
        }
        if (ch == 'B') {
            return 'W';
        }
        return 'R';
    }

public:
    int findColoredRoot(vector<vector<int>>& edges, string colors) {
        int n = colors.size();
        vector<int8_t> in_deg(n), out_deg(n); // 优化空间，减少 cache miss
        int root = xor_n(n - 1);

        for (auto& e : edges) {
            int x = e[0], y = e[1];
            char cx = colors[x], cy = colors[y];
            if (cx == cy) {
                return -1;
            }

            // 确定谁是谁的父节点
            if (get_child_color(cx) != cy) {
                swap(x, y);
            }

            // x 是 y 的父节点
            out_deg[x]++;
            in_deg[y]++;
            if (out_deg[x] > 2 || in_deg[y] > 1) {
                return -1;
            }

            root ^= y;
        }

        return root;
    }
};
```

```go [sol-Go]
// 返回 0 到 n 的异或和
func xorN(n int) int {
	switch n % 4 {
	case 0: return n
	case 1: return 1
	case 2: return n + 1
	default: return 0
	}
}

var childColor = [...]byte{'R': 'B', 'B': 'W', 'W': 'R'}

func findColoredRoot(edges [][]int, colors string) int {
	n := len(colors)
	inDeg := make([]int8, n) // 优化空间，减少 cache miss
	outDeg := make([]int8, n)
	root := xorN(n - 1)

	for _, e := range edges {
		x, y := e[0], e[1]
		cx, cy := colors[x], colors[y]
		if cx == cy {
			return -1
		}

		// 确定谁是谁的父节点
		if childColor[cx] != cy {
			x, y = y, x
		}

		// x 是 y 的父节点
		outDeg[x]++
		inDeg[y]++
		if outDeg[x] > 2 || inDeg[y] > 1 {
			return -1
		}

		root ^= y
	}

	return root
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{colors}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

**注**：用到了这个异或技巧的题目还有 [2196. 根据描述创建二叉树](https://leetcode.cn/problems/create-binary-tree-from-descriptions/)。

## 专题训练

见下面思维题单的「**§5.8 分类讨论**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
