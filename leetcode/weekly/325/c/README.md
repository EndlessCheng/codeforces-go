## 视频讲解

前置知识：[【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

本题视频讲解：[【周赛 325】](https://www.bilibili.com/video/BV1FV4y1F7v7/)第三题。

## 思路

「任意两种糖果价格绝对差的**最小值**」等价于「排序后，任意两种**相邻**糖果价格绝对差的最小值」。

如果题目中有「最大化最小值」或者「最小化最大值」，一般都是二分答案，请记住这个套路。

为什么？对于本题来说，甜蜜度越大，能选择的糖果越少，有**单调性**，所以可以二分。

定义 $f(d)$ 表示甜蜜度至少为 $d$ 时，至多能选多少类糖果。

二分答案 $d$：

- 如果 $f(d)\ge k$，说明答案至少为 $d$。
- 如果 $f(d)< k$，说明答案至多为 $d-1$。
- 二分结束后，设答案为 $\textit{d}_0$，那么 $f(d_0)\ge k$ 且 $f(d_0+1)< k$。

如何计算 $f(d)$？对 $\textit{price}$ 从小到大排序，贪心地计算 $f(d)$：从 $\textit{price}[0]$ 开始选；假设上一个选的数是 $\textit{pre}$，那么当 $\textit{price}[i] \ge \textit{pre}+d$ 时，才可以选 $\textit{price}[i]$。

二分下界可以取 $1$，上界可以取 $\left\lfloor\dfrac{\max(\textit{price})-\min(\textit{price})}{k-1}\right\rfloor$，这是因为**最小值不会超过平均值**。（平均值指选了 $\textit{price}$ 最小最大以及中间的一些糖果，相邻糖果差值的平均值。）

> 请注意，二分的区间的定义是：尚未确定 $f(d)$ 与 $k$ 的大小关系的 $d$ 的值组成的集合（范围）。在区间左侧外面的 $d$ 都是 $f(d)\ge k$ 的，在区间右侧外面的 $d$ 都是 $f(d)< k$ 的。在理解二分时，请牢记区间的定义及其性质。

### 答疑

**问**：为什么二分出来的答案，一定来自数组中价格的差？有没有可能，二分出来的答案不是任何价格的差？

**答**：反证法。如果答案 $d$ 不是任何价格的差，也就是说，礼盒中任意两种糖果的价格的绝对差都大于 $d$，也就是大于等于 $d+1$。那么对于 $d+1$ 来说，它也可以满足 `f(d + 1) == true`，这与循环不变量相矛盾。

```py [sol-Python3]
class Solution:
    def maximumTastiness(self, price: List[int], k: int) -> int:
        price.sort()

        def f(d: int) -> int:
            cnt, pre = 1, price[0]
            for p in price:
                if p >= pre + d:
                    cnt += 1
                    pre = p
            return cnt

        # 二分模板·其三（开区间写法）https://www.bilibili.com/video/BV1AP41137w7/
        left, right = 0, (price[-1] - price[0]) // (k - 1) + 1
        while left + 1 < right:  # 开区间不为空
            # 循环不变量：
            # f(left) >= k
            # f(right) < k
            mid = (left + right) // 2
            if f(mid) >= k: left = mid  # 下一轮二分 (mid, right)
            else: right = mid  # 下一轮二分 (left, mid)
        return left
```

```java [sol-Java]
class Solution {
    public int maximumTastiness(int[] price, int k) {
        Arrays.sort(price);

        // 二分模板·其三（开区间写法）https://www.bilibili.com/video/BV1AP41137w7/
        int left = 0, right = (price[price.length - 1] - price[0]) / (k - 1) + 1;
        while (left + 1 < right) { // 开区间不为空
            // 循环不变量：
            // f(left) >= k
            // f(right) < k
            int mid = left + (right - left) / 2;
            if (f(price, mid) >= k) left = mid; // 下一轮二分 (mid, right)
            else right = mid; // 下一轮二分 (left, mid)
        }
        return left;
    }

    private int f(int[] price, int d) {
        int cnt = 1, pre = price[0];
        for (int p : price) {
            if (p >= pre + d) {
                cnt++;
                pre = p;
            }
        }
        return cnt;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumTastiness(vector<int> &price, int k) {
        sort(price.begin(), price.end());

        auto f = [&](int d) -> int {
            int cnt = 1, pre = price[0];
            for (int p: price) {
                if (p >= pre + d) {
                    cnt++;
                    pre = p;
                }
            }
            return cnt;
        };

        // 二分模板·其三（开区间写法）https://www.bilibili.com/video/BV1AP41137w7/
        int left = 0, right = (price.back() - price[0]) / (k - 1) + 1;
        while (left + 1 < right) { // 开区间不为空
            // 循环不变量：
            // f(left) >= k
            // f(right) < k
            int mid = left + (right - left) / 2;
            (f(mid) >= k ? left : right) = mid;
        }
        return left;
    }
};
```

```go [sol-Go]
func maximumTastiness(price []int, k int) int {
	sort.Ints(price)
	return sort.Search((price[len(price)-1]-price[0])/(k-1), func(d int) bool {
		d++ // 二分最小的 f(d+1) < k，从而知道最大的 f(d) >= k
		cnt, pre := 1, price[0]
		for _, p := range price[1:] {
			if p >= pre+d {
				cnt++
				pre = p
			}
		}
		return cnt < k
	})
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + n\log U)$，其中 $n$ 为 $\textit{price}$ 的长度，$U=\left\lfloor\dfrac{\max(\textit{price})-\min(\textit{price})}{k-1}\right\rfloor$。
- 空间复杂度：$\mathcal{O}(1)$，忽略排序的空间，仅用到若干额外变量。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
