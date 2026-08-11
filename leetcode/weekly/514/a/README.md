根据 [排序不等式](https://baike.baidu.com/item/%E6%8E%92%E5%BA%8F%E4%B8%8D%E7%AD%89%E5%BC%8F/7775728)，最优做法是把最大的折扣用在最贵的商品上。

把 $\textit{prices}$ 和 $\textit{discounts}$ 降序排序，遍历，计算并累加折扣后的价格。

代码实现时，可以把价格乘以 $100$，这样折扣后的价格是整数。最后返回时再除以 $100$。这样只有一次浮点运算。

[本题视频讲解](https://www.bilibili.com/video/BV1ryuy6WEDs/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minPrice(self, prices: list[int], discounts: list[int]) -> float:
        prices.sort(reverse=True)
        discounts.sort(reverse=True)

        # 更快的写法见另一份代码【Python3 写法二】
        s = 0
        for i, p in enumerate(prices):
            d = discounts[i] if i < len(discounts) else 0
            s += p * (100 - d)
        return s / 100
```

```py [sol-Python3 写法二]
class Solution:
    def minPrice(self, prices: list[int], discounts: list[int]) -> float:
        prices.sort(reverse=True)
        discounts.sort(reverse=True)

        discount = sum(p * d for p, d in zip(prices, discounts))
        return (sum(prices) * 100 - discount) / 100
```

```java [sol-Java]
class Solution {
    public double minPrice(int[] prices, int[] discounts) {
        Arrays.sort(prices);
        Arrays.sort(discounts);

        long ans = 0;
        for (int i = 0; i < prices.length; i++) {
            int d = i < discounts.length ? discounts[discounts.length - 1 - i] : 0;
            ans += prices[prices.length - 1 - i] * (100 - d);
        }
        return ans / 100.;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    double minPrice(vector<int>& prices, vector<int>& discounts) {
        ranges::sort(prices, greater());
        ranges::sort(discounts, greater());

        long long ans = 0;
        for (int i = 0; i < prices.size(); i++) {
            int d = i < discounts.size() ? discounts[i] : 0;
            ans += prices[i] * (100 - d);
        }
        return ans / 100.;
    }
};
```

```go [sol-Go]
func minPrice(prices, discounts []int) float64 {
	slices.SortFunc(prices, func(a, b int) int { return b - a })
	slices.SortFunc(discounts, func(a, b int) int { return b - a })

	ans := 0
	for i, p := range prices {
		d := 0
		if i < len(discounts) {
			d = discounts[i]
		}
		ans += p * (100 - d)
	}
	return float64(ans) / 100
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + m\log m)$，其中 $n$ 是 $\textit{prices}$ 的长度，$m$ 是 $\textit{discounts}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 进一步优化

由于瓶颈在排序上，考虑如何优化排序。

1. $\textit{prices}$ 只需要对前 $\min(n,m)$ 大排序。可以先跑一遍快速选择算法，再排序。
2. $\textit{discounts}$ 的元素范围只有 $[1,100]$，可以改用**计数排序**。

## 专题训练

见下面贪心题单的「**§4.3 排序不等式**」。

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
