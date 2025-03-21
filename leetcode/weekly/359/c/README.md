请看 [视频讲解](https://www.bilibili.com/video/BV1Rx4y1f75Y/) 第三题。

## 前置知识：动态规划入门

详见 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)

## 思路

定义 $f[i+1]$ 表示销售编号不超过 $i$ 的房屋的最大盈利。

考虑编号为 $i$ 的房屋卖或不卖：

- 不卖，有 $f[i+1]=f[i]$。
- 卖，那么遍历所有 $\textit{end}_j=i$ 的购买请求，有 $f[i+1] = \max (f[\textit{start}_j]+\textit{gold}_j)$。为了方便遍历，可以先把所有 $\textit{end}$ 相同的数据归类。
- 两种情况取最大值。

初始值 $f[0]=0$。

最终答案为 $f[n]$。

```py [sol-Python3]
class Solution:
    def maximizeTheProfit(self, n: int, offers: List[List[int]]) -> int:
        groups = [[] for _ in range(n)]
        for start, end, gold in offers:
            groups[end].append((start, gold))

        f = [0] * (n + 1)
        for end, g in enumerate(groups):
            f[end + 1] = f[end]
            for start, gold in g:
                f[end + 1] = max(f[end + 1], f[start] + gold)
        return f[n]
```

```java [sol-Java]
class Solution {
    public int maximizeTheProfit(int n, List<List<Integer>> offers) {
        List<int[]>[] groups = new ArrayList[n];
        Arrays.setAll(groups, i -> new ArrayList<>());
        for (List<Integer> offer : offers) {
            groups[offer.get(1)].add(new int[]{offer.get(0), offer.get(2)});
        }

        int[] f = new int[n + 1];
        for (int end = 0; end < n; end++) {
            f[end + 1] = f[end];
            for (int[] p : groups[end]) {
                f[end + 1] = Math.max(f[end + 1], f[p[0]] + p[1]);
            }
        }
        return f[n];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximizeTheProfit(int n, vector<vector<int>>& offers) {
        vector<vector<pair<int, int>>> groups(n);
        for (auto& offer : offers) {
            groups[offer[1]].emplace_back(offer[0], offer[2]);
        }

        vector<int> f(n + 1);
        for (int end = 0; end < n; end++) {
            f[end + 1] = f[end];
            for (auto& [start, gold]: groups[end]) {
                f[end + 1] = max(f[end + 1], f[start] + gold);
            }
        }
        return f[n];
    }
};
```

```go [sol-Go]
func maximizeTheProfit(n int, offers [][]int) int {
	type pair struct{ start, gold int }
	groups := make([][]pair, n)
	for _, offer := range offers {
		start, end, gold := offer[0], offer[1], offer[2]
		groups[end] = append(groups[end], pair{start, gold})
	}

	f := make([]int, n+1)
	for end, g := range groups {
		f[end+1] = f[end]
		for _, p := range g {
			f[end+1] = max(f[end+1], f[p.start]+p.gold)
		}
	}
	return f[n]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $m$ 为 $\textit{offers}$ 的长度。
- 空间复杂度：$\mathcal{O}(n+m)$。

## 相似题目

- [2008. 出租车的最大盈利](https://leetcode.cn/problems/maximum-earnings-from-taxi/)（和本题几乎一样）
- [1235. 规划兼职工作](https://leetcode.cn/problems/maximum-profit-in-job-scheduling/)（数据范围更大的情况，[我的题解](https://leetcode.cn/problems/maximum-profit-in-job-scheduling/solution/dong-tai-gui-hua-er-fen-cha-zhao-you-hua-zkcg/)）
- [1751. 最多可以参加的会议数目 II](https://leetcode.cn/problems/maximum-number-of-events-that-can-be-attended-ii/)（区间个数限制）
- [2054. 两个最好的不重叠活动](https://leetcode.cn/problems/two-best-non-overlapping-events/)

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
