先判断无解的情况。根据题意，对于节点 $x$ 来说：

1. $x$ 至多有三个邻居。
2. $x$ 的邻居的颜色，不能与 $x$ 的颜色相同。
3. $x$ 至多有一个父节点，所以在 $x$ 的邻居中，「上一个颜色」的个数至多一个。R 的上一个颜色是 W，W 的上一个颜色是 B，B 的上一个颜色是 R。

如果满足上述要求，那么还要满足什么要求，节点 $x$ 才能作为树的根？

1. $x$ 至多有两个邻居。
2. $x$ 没有父节点，所以在 $x$ 的邻居中，「上一个颜色」的个数必须是 $0$。

如果有两个点都可以作为树的根，会发生什么情况？

举个例子，看示例 1 的图，如果有两个白点都可以作为根，那么当其中一个白点是根时：

- 如果另一个白点只有一个邻居（度为 $1$），那么另一个白点变成叶子，其父节点是红点，不符合要求（红点的儿子必须是蓝点）。
- 如果另一个白点有两个邻居，那么另一个白点有一个父节点和一个儿子，其父节点是红点，不符合要求（红点的儿子必须是蓝点）。

一般地，如果有多个点可以作为根，那么当其中一个点是根的时候，其余点会有儿子变成其父亲，导致颜色不满足要求。所以**只有一个点能作为根**。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

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
            # 一个点至多三个邻居
            if len(to) > 3:
                return -1

            cnt = [0] * 3
            for y in to:
                cnt[mp[colors[y]]] += 1

            color_x = mp[colors[x]]
            # 不能有邻居和 x 颜色相同，或者 x 的父节点的颜色个数超过 1
            if cnt[color_x] > 0 or cnt[(color_x - 1) % 3] > 1:
                return -1

            if len(to) < 3 and cnt[(color_x - 1) % 3] == 0:
                # 不能有多个点能是根，那样当一个点是根的时候，另一个点不满足要求
                if ans >= 0:
                    return -1
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
        for (int x = 0; x < g.length; x++) {
            List<Integer> to = g[x];
            // 一个点至多三个邻居
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

            if (to.size() < 3 && cnt[(colorX + 2) % 3] == 0) {
                // 不能有多个点能是根，那样当一个点是根的时候，另一个点不满足要求
                if (ans >= 0) {
                    return -1;
                }
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
        for (int x = 0; x < g.size(); x++) {
            auto& to = g[x];
            // 一个点至多三个邻居
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

            if (to.size() < 3 && cnt[(color_x + 2) % 3] == 0) {
                // 不能有多个点能是根，那样当一个点是根的时候，另一个点不满足要求
                if (ans >= 0) {
                    return -1;
                }
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
		// 一个点至多三个邻居
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

		if len(to) < 3 && cnt[(colorX+2)%3] == 0 {
			// 不能有多个点能是根，那样当一个点是根的时候，另一个点不满足要求
			if ans >= 0 {
				return -1
			}
			ans = x
		}
	}

	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{colors}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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
