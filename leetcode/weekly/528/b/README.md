例如 $x$ 有三个儿子，子树最大分数和分别为 $[-5,-2,3]$。把最小的 $-5$ 变成 $0$，可以让子树 $x$ 的分数和变得尽量大。

一般地，对于节点 $x$，我们需要知道 $x$ 的儿子子树 $y$ 的最大分数和。这是一个**和原问题相似的、规模更小的子问题**，可以**递归**解决。

我们可以先累加 $x$ 的儿子子树 $y$ 的最大分数和，然后从中减去最小的。如果最小的分数和 $\ge 0$ 且 $\textit{op}[x] = 0$，那么无需减去最小的分数和。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def maxScore(self, parent: list[int], op: list[int], score: list[int]) -> int:
        n = len(parent)
        g = [[] for _ in range(n)]
        for i in range(1, n):
            g[parent[i]].append(i)

        # 返回子树 x 的最大分数和
        def dfs(x: int) -> int:
            sums = score[x]
            if not g[x]:  # x 是叶子
                return sums
            mn = inf
            for y in g[x]:
                s = dfs(y)  # 子树 y 的最大分数和
                sums += s  # 累加分数和
                mn = min(mn, s)  # 维护最小的分数和
            if mn < 0 or op[x]:
                sums -= mn  # 去掉最小的分数和
            return sums

        return dfs(0)
```

```java [sol-Java]
class Solution {
    public int maxScore(int[] parent, int[] op, int[] score) {
        int n = parent.length;
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, _ -> new ArrayList<>());
        for (int i = 1; i < n; i++) {
            g[parent[i]].add(i);
        }

        return dfs(0, g, op, score);
    }

    // 返回子树 x 的最大分数和
    private int dfs(int x, List<Integer>[] g, int[] op, int[] score) {
        int sum = score[x];
        if (g[x].isEmpty()) { // x 是叶子
            return sum;
        }
        int mn = Integer.MAX_VALUE;
        for (int y : g[x]) {
            int s = dfs(y, g, op, score); // 子树 y 的最大分数和
            sum += s; // 累加分数和
            mn = Math.min(mn, s); // 维护最小的分数和
        }
        if (mn < 0 || op[x] > 0) {
            sum -= mn; // 去掉最小的分数和
        }
        return sum;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxScore(vector<int>& parent, vector<int>& op, vector<int>& score) {
        int n = parent.size();
        vector<vector<int>> g(n);
        for (int i = 1; i < n; i++) {
            g[parent[i]].push_back(i);
        }

        // 返回子树 x 的最大分数和
        auto dfs = [&](this auto&& dfs, int x) -> int {
            int sum = score[x];
            if (g[x].empty()) { // x 是叶子
                return sum;
            }
            int mn = INT_MAX;
            for (int y : g[x]) {
                int s = dfs(y); // 子树 y 的最大分数和
                sum += s; // 累加分数和
                mn = min(mn, s); // 维护最小的分数和
            }
            if (mn < 0 || op[x]) {
                sum -= mn; // 去掉最小的分数和
            }
            return sum;
        };

        return dfs(0);
    }
};
```

```go [sol-Go]
func maxScore(parent []int, op []int, score []int) int {
	n := len(parent)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parent[i]
		g[p] = append(g[p], i)
	}

	// 返回子树 x 的最大分数和
	var dfs func(int) int
	dfs = func(x int) int {
		sum := score[x]
		if g[x] == nil { // x 是叶子
			return sum
		}
		mn := math.MaxInt
		for _, y := range g[x] {
			s := dfs(y)     // 子树 y 的最大分数和
			sum += s        // 累加分数和
			mn = min(mn, s) // 维护最小的分数和
		}
		if mn < 0 || op[x] > 0 {
			sum -= mn // 去掉最小的分数和
		}
		return sum
	}

	return dfs(0)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{parent}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面动态规划题单的「**§12.5 其他树形 DP**」。

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
