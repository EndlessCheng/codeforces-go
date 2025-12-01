首先，统计每一行的点的个数，如果这一行有 $c$ 个点，那么从 $c$ 个点中选 $2$ 个点，有 $\dfrac{c(c-1)}{2}$ 种选法，可以组成一条水平边，即梯形的顶边或底边。

枚举每一行，设这一行有 $k=\dfrac{c(c-1)}{2}$ 条水平边，那么另外一条边就是之前遍历过的行的边数 $s$。根据乘法原理，之前遍历过的行与这一行，一共可以组成

$$
s\cdot k
$$

个水平梯形，加入答案。

⚠**注意**：另外一条边不能是其余所有行，这会导致重复计算。

在最坏情况下，有两行，每行 $\dfrac{n}{2}$ 个点，组成约 $\dfrac{n^2}{8}$ 条线段，答案约为 $\dfrac{n^4}{64} = 1.5625\times 10^{18}$，这不超过 $64$ 位整数最大值，所以无需在循环中取模。

[本题视频讲解](https://www.bilibili.com/video/BV1tbg8z3EaP/?t=3m40s) 详细介绍了本题的计算过程和注意事项，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countTrapezoids(self, points: List[List[int]]) -> int:
        MOD = 1_000_000_007
        cnt = Counter(p[1] for p in points)  # 统计每一行（水平线）有多少个点
        ans = s = 0
        for c in cnt.values():
            k = c * (c - 1) // 2
            ans += s * k
            s += k
        return ans % MOD
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    public int countTrapezoids(int[][] points) {
        Map<Integer, Integer> cnt = new HashMap<>(points.length, 1); // 预分配空间
        for (int[] p : points) {
            cnt.merge(p[1], 1, Integer::sum); // 统计每一行（水平线）有多少个点
        }

        long ans = 0, s = 0;
        for (int c : cnt.values()) {
            long k = (long) c * (c - 1) / 2;
            ans += s * k;
            s += k;
        }
        return (int) (ans % MOD);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countTrapezoids(vector<vector<int>>& points) {
        const int MOD = 1'000'000'007;
        unordered_map<int, int> cnt;
        for (auto& p : points) {
            cnt[p[1]]++; // 统计每一行（水平线）有多少个点
        }

        long long ans = 0, s = 0;
        for (auto& [_, c] : cnt) {
            long long k = 1LL * c * (c - 1) / 2;
            ans += s * k;
            s += k;
        }
        return ans % MOD;
    }
};
```

```go [sol-Go]
func countTrapezoids(points [][]int) (ans int) {
	const mod = 1_000_000_007
	cnt := make(map[int]int, len(points)) // 预分配空间
	for _, p := range points {
		cnt[p[1]]++ // 统计每一行（水平线）有多少个点
	}

	s := 0
	for _, c := range cnt {
		k := c * (c - 1) / 2
		ans += s * k
		s += k
	}
	return ans % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{points}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面数据结构题单的「**§0.1 枚举右，维护左**」。

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
