注意题目的这句话：保证单位 $0$ 可以通过唯一的转换路径（不需要反向转换）转换为**任何**其他单位。

这意味着输入的是 $n$ 个点和 $n-1$ 条边的**连通图**，说明输入的是一棵**树**。

建树。然后从 $0$ 开始 DFS 这棵树，同时把从 $0$ 到 $i$ 的边权乘起来，即为 $\textit{ans}[i]$。

为了计算乘积，递归的过程中额外传入参数 $\textit{mul}$。

注意取模。关于模运算的知识点，见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1J2jAziENo/?t=2m21s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def baseUnitConversions(self, conversions: List[List[int]]) -> List[int]:
        MOD = 1_000_000_007
        n = len(conversions) + 1
        g = [[] for _ in range(n)]
        for x, y, weight in conversions:
            g[x].append((y, weight))

        ans = [0] * n
        def dfs(x: int, mul: int) -> None:
            ans[x] = mul
            for y, weight in g[x]:
                dfs(y, mul * weight % MOD)
        dfs(0, 1)
        return ans
```

```java [sol-Java]
class Solution {
    public int[] baseUnitConversions(int[][] conversions) {
        int n = conversions.length + 1;
        List<int[]>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : conversions) {
            g[e[0]].add(new int[]{e[1], e[2]});
        }

        int[] ans = new int[n];
        dfs(0, 1, g, ans);
        return ans;
    }

    private void dfs(int x, long mul, List<int[]>[] g, int[] ans) {
        ans[x] = (int) mul;
        for (int[] e : g[x]) {
            dfs(e[0], mul * e[1] % 1_000_000_007, g, ans);
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> baseUnitConversions(vector<vector<int>>& conversions) {
        const int MOD = 1'000'000'007;
        int n = conversions.size() + 1;
        vector<vector<pair<int, int>>> g(n);
        for (auto& e : conversions) {
            g[e[0]].emplace_back(e[1], e[2]);
        }

        vector<int> ans(n);
        auto dfs = [&](this auto&& dfs, int x, long long mul) -> void {
            ans[x] = mul;
            for (auto& [y, weight] : g[x]) {
                dfs(y, mul * weight % MOD);
            }
        };
        dfs(0, 1);
        return ans;
    }
};
```

```go [sol-Go]
func baseUnitConversions(conversions [][]int) []int {
	const mod = 1_000_000_007
	n := len(conversions) + 1
	type edge struct{ to, weight int }
	g := make([][]edge, n)
	for _, e := range conversions {
		x, y, weight := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, weight})
	}

	ans := make([]int, n)
	var dfs func(int, int)
	dfs = func(x, mul int) {
		ans[x] = mul
		for _, e := range g[x] {
			dfs(e.to, mul*e.weight%mod)
		}
	}
	dfs(0, 1)
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{conversions}$ 的长度。树中每个节点恰好访问一次。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，见下面一般树题单的「**§3.2 自顶向下 DFS**」。

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
11. 【本题相关】[链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
