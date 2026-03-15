## 引入

如果一对点，横坐标相同或者纵坐标相同，那么在这两个点之间连一条边，可以得到一个无向图。我们需要在这个图上，找到最大的两个连通块。

如果 $\mathcal{O}(n^2)$ 枚举所有点对，就太慢了。

对于 $x$ 相同的点，我们可以把这些点都连到同一个代表（或者说中介）上。对于 $y$ 相同的点，同理，把这些点都连到同一个代表（或者说中介）上。

如果两个点可以通过中介互相到达，那么这两个点属于同一个连通块。

这样只需要连 $2n$ 条边，大大地提高了效率。

进一步地，直接把中介 $x$ 和中介 $y$ 相连，这样只需要连 $n$ 条边。

## 思路

激活点 $(x,y)$，也同时激活了 $x$ 这一行，$y$ 这一列。

问题可以抽象成如下图论问题：

- 在行和列之间连边，**相连的行列可以被一个坐标点直接或间接地激活到**。具体地，在节点 $x$ 和节点 $y + \textit{offset}$ 之间连边，其中 $\textit{offset} = 3\times 10^9$。增加 $\textit{offset}$ 是为了区分行列。
- 找到图中最大的两个连通块。这两个连通块可以通过额外添加一个坐标点相连。
- 连通块的大小取决于连通块中的坐标点的个数。

这可以用**并查集**解决。本题数据范围很大，可以用哈希表记录代表元。

[本题视频讲解](https://www.bilibili.com/video/BV1DvwTzbE1n/?t=27m34s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxActivated(self, points: List[List[int]]) -> int:
        # 哈希表并查集
        fa = {}

        def find(x: int) -> int:
            if x not in fa:
                fa[x] = x
            if fa[x] != x:
                fa[x] = find(fa[x])
                return fa[x]
            return x

        OFFSET = 3 * 10 ** 9
        for x, y in points:
            fa[find(x)] = find(y + OFFSET)

        # 统计连通块的大小
        size = Counter(find(p[0]) for p in points)

        mx1 = mx2 = 0
        for sz in size.values():
            if sz > mx1:
                mx2 = mx1
                mx1 = sz
            elif sz > mx2:
                mx2 = sz

        return mx1 + mx2 + 1
```

```java [sol-Java]
class Solution {
    public int maxActivated(int[][] points) {
        // 哈希表并查集
        Map<Long, Long> fa = new HashMap<>();

        final long OFFSET = (long) 3e9;
        for (int[] p : points) {
            long fx = find(p[0], fa);
            long fy = find(p[1] + OFFSET, fa);
            fa.put(fx, fy);
        }

        // 统计连通块的大小
        Map<Long, Integer> size = new HashMap<>();
        for (int[] p : points) {
            size.merge(find(p[0], fa), 1, Integer::sum);
        }

        int mx1 = 0, mx2 = 0;
        for (int sz : size.values()) {
            if (sz > mx1) {
                mx2 = mx1;
                mx1 = sz;
            } else if (sz > mx2) {
                mx2 = sz;
            }
        }

        return mx1 + mx2 + 1;
    }

    private long find(long x, Map<Long, Long> fa) {
        if (!fa.containsKey(x)) {
            fa.put(x, x);
        }
        long fx = fa.get(x);
        if (fx != x) {
            long root = find(fx, fa);
            fa.put(x, root);
            return root;
        }
        return x;
    }
}
```

```cpp [sol-C++]
class Solution {
    // 哈希表并查集
    unordered_map<long long, long long> fa;

    long long find(long long x) {
        auto it = fa.find(x);
        if (it == fa.end()) {
            fa[x] = x;
            return x;
        }
        auto& fx = it->second;
        if (fx != x) {
            fx = find(fx);
        }
        return fx;
    }

public:
    int maxActivated(vector<vector<int>>& points) {
        constexpr long long OFFSET = 3e9;
        for (auto& p : points) {
            fa[find(p[0])] = find(p[1] + OFFSET);
        }

        unordered_map<long long, int> size;
        for (auto& p : points) {
            size[find(p[0])]++; // 统计连通块的大小
        }

        int mx1 = 0, mx2 = 0;
        for (auto& [_, sz] : size) {
            if (sz > mx1) {
                mx2 = mx1;
                mx1 = sz;
            } else if (sz > mx2) {
                mx2 = sz;
            }
        }

        return mx1 + mx2 + 1;
    }
};
```

```go [sol-Go]
func maxActivated(points [][]int) int {
	// 哈希表并查集
	fa := map[int]int{}
	var find func(int) int
	find = func(x int) int {
		fx, ok := fa[x]
		if !ok {
			fa[x] = x
			fx = x
		}
		if fx != x {
			fa[x] = find(fx)
			return fa[x]
		}
		return x
	}

	const offset int = 3e9
	for _, p := range points {
		fa[find(p[0])] = find(p[1] + offset)
	}

	size := map[int]int{}
	for _, p := range points {
		size[find(p[0])]++ // 统计连通块的大小
	}

	mx1, mx2 := 0, 0
	for _, sz := range size {
		if sz > mx1 {
			mx2 = mx1
			mx1 = sz
		} else if sz > mx2 {
			mx2 = sz
		}
	}
	return mx1 + mx2 + 1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{points}$ 的长度。并查集我只写了路径压缩，单次操作的均摊复杂度是 $\mathcal{O}(\log n)$。
- 空间复杂度：$\mathcal{O}(n)$。

注：如果用建图 + DFS 实现，可以做到 $\mathcal{O}(n)$ 时间。

## 相似题目

[947. 移除最多的同行或同列石头](https://leetcode.cn/problems/most-stones-removed-with-same-row-or-column/)

## 专题训练

见下面数据结构题单的「**七、并查集**」。

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
