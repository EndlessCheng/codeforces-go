对于物品 $i$，统计满足 $\textit{factor}_j$ 是 $\textit{factor}_i$ 的倍数的物品 $j$ 的个数（**不**包括物品 $i$），记作 $\textit{cnt}_i$。那么：

- 前 $\textit{cnt}_i$ 次购买物品 $i$，每次可以获得 $2$ 个物品。
- 继续购买物品 $i$，每次只能获得 $1$ 个物品。

想象有多家商场做活动，**都是买一送一**，去哪家商场买东西？优先去最便宜的。

设最便宜的物品的价格为 $\textit{minPrice}$。分类讨论：

- 对于价格 $< 2\cdot \textit{minPrice}$ 的物品，购买 $\textit{cnt}_i$ 次，每次可以获得 $2$ 个物品。为了最大化所获得的物品总数，贪心地，**按照价格从低到高购买**。
- 对于价格 $\ge 2\cdot \textit{minPrice}$ 的物品，**无需购买**，因为这不如买两个最便宜的物品。

代码实现时，我们需要快速计算 $\textit{cnt}_i$。枚举 $x = 1,2,\ldots, n$ 以及 $x$ 的倍数 $y=x,2x,\ldots$ 计算 $y$ 的个数之和。由调和级数可知，这只需要 $\mathcal{O}(n\log n)$ 的时间。其中 $n$ 是 $\textit{items}$ 的长度。注意题目保证 $\textit{factor}_i\le n$。

[本题视频讲解](https://www.bilibili.com/video/BV1KwVn6zEZB/?t=19m10s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maximumSaleItems(self, items: list[list[int]], budget: int) -> int:
        n = len(items)
        cnt_factor = [0] * (n + 1)
        min_price = inf
        for factor, price in items:
            cnt_factor[factor] += 1
            min_price = min(min_price, price)

        cnt_multi = [0] * (n + 1)
        a = []

        for factor, price in items:
            if price >= min_price * 2:
                continue

            if cnt_multi[factor] == 0:  # 之前没有计算过
                for j in range(factor, n + 1, factor):
                    cnt_multi[factor] += cnt_factor[j]

            if cnt_multi[factor] > 1:
                a.append((price, cnt_multi[factor] - 1))  # factor 的倍数不包括该物品

        a.sort(key=lambda p: p[0])

        ans = 0
        for price, cnt in a:
            if budget < price:  # 没钱了
                break
            c = min(cnt, budget // price)  # 该物品最多买 c 个
            budget -= price * c
            ans += c * 2
        return ans + budget // min_price  # 剩余的钱买最便宜的物品
```

```java [sol-Java]
class Solution {
    public int maximumSaleItems(int[][] items, int budget) {
        int n = items.length;
        int[] cntFactor = new int[n + 1];
        int minPrice = Integer.MAX_VALUE;
        for (int[] p : items) {
            cntFactor[p[0]]++;
            minPrice = Math.min(minPrice, p[1]);
        }

        int[] cntMulti = new int[n + 1];
        List<int[]> a = new ArrayList<>();

        for (int[] p : items) {
            int factor = p[0], price = p[1];
            if (price >= minPrice * 2) {
                continue;
            }

            if (cntMulti[factor] == 0) { // 之前没有计算过
                for (int j = factor; j <= n; j += factor) {
                    cntMulti[factor] += cntFactor[j];
                }
            }

            if (cntMulti[factor] > 1) {
                a.add(new int[]{price, cntMulti[factor] - 1}); // factor 的倍数不包括该物品
            }
        }

        a.sort((p, q) -> p[0] - q[0]);

        int ans = 0;
        for (int[] p : a) {
            int price = p[0], cnt = p[1];
            if (budget < price) { // 没钱了
                break;
            }
            int c = Math.min(cnt, budget / price); // 该物品最多买 c 个
            budget -= price * c;
            ans += c * 2;
        }
        return ans + budget / minPrice; // 剩余的钱买最便宜的物品
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumSaleItems(vector<vector<int>>& items, int budget) {
        int n = items.size();
        vector<int> cnt_factor(n + 1);
        int min_price = INT_MAX;
        for (auto& p : items) {
            cnt_factor[p[0]]++;
            min_price = min(min_price, p[1]);
        }
        vector<int> cnt_multi(n + 1);
        vector<pair<int, int>> a;

        for (auto& p : items) {
            int factor = p[0], price = p[1];
            if (price >= min_price * 2) {
                continue;
            }

            int& cnt = cnt_multi[factor]; // 注意这里是引用
            if (cnt == 0) { // 之前没有计算过
                for (int j = factor; j <= n; j += factor) {
                    cnt += cnt_factor[j];
                }
            }

            if (cnt > 1) {
                a.emplace_back(price, cnt - 1); // factor 的倍数不包括该物品
            }
        }

        ranges::sort(a);

        int ans = 0;
        for (auto& [price, cnt] : a) {
            if (budget < price) { // 没钱了
                break;
            }
            int c = min(cnt, budget / price); // 该物品最多买 c 个
            budget -= price * c;
            ans += c * 2;
        }
        return ans + budget / min_price; // 剩余的钱买最便宜的物品
    }
};
```

```go [sol-Go]
func maximumSaleItems(items [][]int, budget int) (ans int) {
	n := len(items)
	cntFactor := make([]int, n+1)
	minPrice := math.MaxInt
	for _, p := range items {
		cntFactor[p[0]]++
		minPrice = min(minPrice, p[1])
	}
	cntMulti := make([]int, n+1)
	type pair struct{ price, cnt int }
	a := []pair{}

	for _, p := range items {
		factor, price := p[0], p[1]
		if price >= minPrice*2 {
			continue
		}

		if cntMulti[factor] == 0 { // 之前没有计算过
			for j := factor; j <= n; j += factor {
				cntMulti[factor] += cntFactor[j]
			}
		}

		if cntMulti[factor] > 1 {
			a = append(a, pair{price, cntMulti[factor] - 1}) // factor 的倍数不包括该物品
		}
	}

	slices.SortFunc(a, func(a, b pair) int { return a.price - b.price })

	for _, p := range a {
		if budget < p.price { // 没钱了
			break
		}
		c := min(p.cnt, budget/p.price) // 该物品最多买 c 个
		budget -= p.price * c
		ans += c * 2
	}

	// 剩余的钱买最便宜的物品
	return ans + budget/minPrice
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{items}$ 的长度。由调和级数可知，计算 `cntMulti` 的总时间复杂度为 $\mathcal{O}(n\log n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面贪心题单的「**§1.1 从最小/最大开始贪心**」。

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
