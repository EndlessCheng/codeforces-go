对于物品 $i$，统计满足 $\textit{factor}_j$ 是 $\textit{factor}_i$ 的倍数的物品 $j$ 的个数（包括物品 $i$），记作 $\textit{cnt}_i$。那么 $\textit{cnt}_i$ 就是**首次**购买物品 $i$ 时，所获得的物品个数。

如果每个物品只能购买一次，那么本题是标准的 **0-1 背包问题**：

- 给你一个容量为 $\textit{budget}$ 的背包，以及 $n$ 个物品，其中物品 $i$ 的体积为 $\textit{price}_i$，价值为 $\textit{cnt}_i$。在**至多装满**背包的情况下，所选物品的价值之和最大是多少？

关于 0-1 背包问题的做法，请看 [0-1 背包【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)。

设装入背包的物品的花费至多为 $i$ 时，获得了 $f_i$ 个物品。

本题可以重复购买物品。由于重复购买物品时，无法免费获取物品，所以贪心地，**重复购买最便宜的物品**，设其价格为 $\textit{minPrice}$。

枚举装入背包的物品的花费至多为 $i$，那么用于重复购买物品的预算为 $\textit{budget} - i$，可以额外购买 $\left\lfloor\dfrac{\textit{budget} - i}{\textit{minPrice}}\right\rfloor$ 个物品。所以可以获得的物品最大总数为

$$
\max_{i=0}^{\textit{budget}} f_i + \left\lfloor\dfrac{\textit{budget} - i}{\textit{minPrice}}\right\rfloor
$$

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

## 优化前

```py [sol-Python3]
class Solution:
    def maximumSaleItems(self, items: list[list[int]], budget: int) -> int:
        f = [0] * (budget + 1)
        min_price = inf

        for factor, price in items:
            min_price = min(min_price, price)

            cnt = 0  # 统计 factor 的倍数（包括 factor）
            for factor_j, _ in items:
                if factor_j % factor == 0:
                    cnt += 1

            # 视作一个体积为 price，价值为 cnt 的物品
            for j in range(budget, price - 1, -1):
                f[j] = max(f[j], f[j - price] + cnt)

        return max(fi + (budget - i) // min_price for i, fi in enumerate(f))
```

```java [sol-Java]
class Solution {
    public int maximumSaleItems(int[][] items, int budget) {
        int[] f = new int[budget + 1];
        int minPrice = Integer.MAX_VALUE;

        for (int[] p : items) {
            int factor = p[0], price = p[1];
            minPrice = Math.min(minPrice, price);

            int cnt = 0; // 统计 factor 的倍数（包括 factor）
            for (int[] q : items) {
                if (q[0] % factor == 0) {
                    cnt++;
                }
            }

            // 视作一个体积为 price，价值为 cnt 的物品
            for (int j = budget; j >= price; j--) {
                f[j] = Math.max(f[j], f[j - price] + cnt);
            }
        }

        int ans = 0;
        for (int i = 0; i <= budget; i++) {
            ans = Math.max(ans, f[i] + (budget - i) / minPrice);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumSaleItems(vector<vector<int>>& items, int budget) {
        vector<int> f(budget + 1);
        int min_price = INT_MAX;

        for (auto& p : items) {
            int factor = p[0], price = p[1];
            min_price = min(min_price, price);

            int cnt = 0; // 统计 factor 的倍数（包括 factor）
            for (auto& q : items) {
                if (q[0] % factor == 0) {
                    cnt++;
                }
            }

            // 视作一个体积为 price，价值为 cnt 的物品
            for (int j = budget; j >= price; j--) {
                f[j] = max(f[j], f[j - price] + cnt);
            }
        }

        int ans = 0;
        for (int i = 0; i <= budget; i++) {
            ans = max(ans, f[i] + (budget - i) / min_price);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumSaleItems(items [][]int, budget int) (ans int) {
	f := make([]int, budget+1)
	minPrice := math.MaxInt

	for _, p := range items {
		factor, price := p[0], p[1]
		minPrice = min(minPrice, price)

		cnt := 0 // 统计 factor 的倍数（包括 factor）
		for _, q := range items {
			if q[0]%factor == 0 {
				cnt++
			}
		}

		// 视作一个体积为 price，价值为 cnt 的物品
		for j := budget; j >= price; j-- {
			f[j] = max(f[j], f[j-price]+cnt)
		}
	}

	for i, cnt := range f {
		ans = max(ans, cnt+(budget-i)/minPrice)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n(n+\textit{budget}))$，其中 $n$ 是 $\textit{items}$ 的长度。
- 空间复杂度：$\mathcal{O}(\textit{budget})$。

## 优化

设 $U=\max(\textit{factor}_i)$。

统计 $\textit{factor}_i$ 的倍数个数时，上面代码写了一个 $\mathcal{O}(n^2)$ 的暴力。改成枚举 $x = 1,2,\ldots,U$ 以及 $x$ 的倍数 $y=x,2x,\ldots$ 计算 $y$ 的个数之和。由调和级数可知，这只需要 $\mathcal{O}(U\log U)$ 的时间。

```py [sol-Python3]
class Solution:
    def maximumSaleItems(self, items: list[list[int]], budget: int) -> int:
        max_factor = max(p[0] for p in items)
        cnt_factor = [0] * (max_factor + 1)
        for factor, _ in items:
            cnt_factor[factor] += 1

        cnt_multi = [0] * (max_factor + 1)
        f = [0] * (budget + 1)
        min_price = inf
        sum_price = 0

        for factor, price in items:
            min_price = min(min_price, price)

            if cnt_multi[factor] == 0:  # 之前没有计算过
                for j in range(factor, max_factor + 1, factor):
                    cnt_multi[factor] += cnt_factor[j]
            cnt = cnt_multi[factor]

            # 视作一个体积为 price，价值为 cnt 的物品
            # 优化：已遍历的物品的体积和至多为 sum_price，大于这个值的体积和无法凑出来
            sum_price = min(sum_price + price, budget)
            for j in range(sum_price, price - 1, -1):
                v = f[j - price] + cnt
                if v > f[j]: f[j] = v  # 手写 max 更快

        return max(fi + (budget - i) // min_price for i, fi in enumerate(f))
```

```java [sol-Java]
class Solution {
    public int maximumSaleItems(int[][] items, int budget) {
        int maxFactor = 0;
        int minPrice = Integer.MAX_VALUE;
        for (int[] p : items) {
            maxFactor = Math.max(maxFactor, p[0]);
            minPrice = Math.min(minPrice, p[1]);
        }

        int[] cntFactor = new int[maxFactor + 1];
        for (int[] p : items) {
            cntFactor[p[0]]++;
        }
        int[] cntMulti = new int[maxFactor + 1];
        int[] f = new int[budget + 1];
        int sumPrice = 0;

        for (int[] p : items) {
            int factor = p[0], price = p[1];

            if (cntMulti[factor] == 0) { // 之前没有计算过
                for (int j = factor; j <= maxFactor; j += factor) {
                    cntMulti[factor] += cntFactor[j];
                }
            }
            int cnt = cntMulti[factor];

            // 视作一个体积为 price，价值为 cnt 的物品
            // 优化：已遍历的物品的体积和至多为 sumPrice，大于这个值的体积和无法凑出来
            sumPrice = Math.min(sumPrice + price, budget);
            for (int j = sumPrice; j >= price; j--) {
                f[j] = Math.max(f[j], f[j - price] + cnt);
            }
        }

        int ans = 0;
        for (int i = 0; i <= budget; i++) {
            ans = Math.max(ans, f[i] + (budget - i) / minPrice);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumSaleItems(vector<vector<int>>& items, int budget) {
        int max_factor = 0;
        int min_price = INT_MAX;
        for (auto& p : items) {
            max_factor = max(max_factor, p[0]);
            min_price = min(min_price, p[1]);
        }

        vector<int> cnt_factor(max_factor + 1);
        for (auto& p : items) {
            cnt_factor[p[0]]++;
        }
        vector<int> cnt_multi(max_factor + 1);
        vector<int> f(budget + 1);
        int sum_price = 0;

        for (auto& p : items) {
            int factor = p[0], price = p[1];

            int& cnt = cnt_multi[factor]; // 注意这里是引用
            if (cnt == 0) { // 之前没有计算过
                for (int j = factor; j <= max_factor; j += factor) {
                    cnt += cnt_factor[j];
                }
            }

            // 视作一个体积为 price，价值为 cnt 的物品
            // 优化：已遍历的物品的体积和至多为 sum_price，大于这个值的体积和无法凑出来
            sum_price = min(sum_price + price, budget);
            for (int j = sum_price; j >= price; j--) {
                f[j] = max(f[j], f[j - price] + cnt);
            }
        }

        int ans = 0;
        for (int i = 0; i <= budget; i++) {
            ans = max(ans, f[i] + (budget - i) / min_price);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumSaleItems(items [][]int, budget int) (ans int) {
	maxFactor := 0
	minPrice := math.MaxInt
	for _, p := range items {
		maxFactor = max(maxFactor, p[0])
		minPrice = min(minPrice, p[1])
	}

	cntFactor := make([]int, maxFactor+1)
	for _, p := range items {
		cntFactor[p[0]]++
	}
	cntMulti := make([]int, maxFactor+1)
	f := make([]int, budget+1)
	sumPrice := 0

	for _, p := range items {
		factor, price := p[0], p[1]

		if cntMulti[factor] == 0 { // 之前没有计算过
			for j := factor; j <= maxFactor; j += factor {
				cntMulti[factor] += cntFactor[j]
			}
		}
		cnt := cntMulti[factor]

		// 视作一个体积为 price，价值为 cnt 的物品
		// 优化：已遍历的物品的体积和至多为 sumPrice，大于这个值的体积和无法凑出来
		sumPrice = min(sumPrice+price, budget)
		for j := sumPrice; j >= price; j-- {
			f[j] = max(f[j], f[j-price]+cnt)
		}
	}

	for i, cnt := range f {
		ans = max(ans, cnt+(budget-i)/minPrice)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\cdot\textit{budget} + U\log U)$，其中 $n$ 是 $\textit{items}$ 的长度，$U=\max(\textit{factor}_i)$。由调和级数可知，计算 `cntMulti` 的总时间复杂度为 $\mathcal{O}(U\log U)$。
- 空间复杂度：$\mathcal{O}(\textit{budget} + U)$。

## 专题训练

见下面动态规划题单的「**§3.1 0-1 背包**」。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
