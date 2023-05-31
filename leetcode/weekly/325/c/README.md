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

## 二分答案·题单

#### 二分答案（按照难度分排序）
- [875. 爱吃香蕉的珂珂](https://leetcode.cn/problems/koko-eating-bananas/)
- [1283. 使结果不超过阈值的最小除数](https://leetcode.cn/problems/find-the-smallest-divisor-given-a-threshold/)
- [2187. 完成旅途的最少时间](https://leetcode.cn/problems/minimum-time-to-complete-trips/)
- [2226. 每个小孩最多能分到多少糖果](https://leetcode.cn/problems/maximum-candies-allocated-to-k-children/)
- [1870. 准时到达的列车最小时速](https://leetcode.cn/problems/minimum-speed-to-arrive-on-time/)
- [1011. 在 D 天内送达包裹的能力](https://leetcode.cn/problems/capacity-to-ship-packages-within-d-days/)
- [2064. 分配给商店的最多商品的最小值](https://leetcode.cn/problems/minimized-maximum-of-products-distributed-to-any-store/)
- [1760. 袋子里最少数目的球](https://leetcode.cn/problems/minimum-limit-of-balls-in-a-bag/)
- [1482. 制作 m 束花所需的最少天数](https://leetcode.cn/problems/minimum-number-of-days-to-make-m-bouquets/)
- [1642. 可以到达的最远建筑](https://leetcode.cn/problems/furthest-building-you-can-reach/)
- [1898. 可移除字符的最大数目](https://leetcode.cn/problems/maximum-number-of-removable-characters/)
- [778. 水位上升的泳池中游泳](https://leetcode.cn/problems/swim-in-rising-water/)
- [2258. 逃离火灾](https://leetcode.cn/problems/escape-the-spreading-fire/)

#### 第 k 小/大（部分题目还可以用堆解决）
- [373. 查找和最小的 K 对数字](https://leetcode.cn/problems/find-k-pairs-with-smallest-sums/)
- [378. 有序矩阵中第 K 小的元素](https://leetcode.cn/problems/kth-smallest-element-in-a-sorted-matrix/)
- [719. 找出第 K 小的数对距离](https://leetcode.cn/problems/find-k-th-smallest-pair-distance/)
- [786. 第 K 个最小的素数分数](https://leetcode.cn/problems/k-th-smallest-prime-fraction/)
- [1439. 有序矩阵中的第 k 个最小数组和](https://leetcode.cn/problems/find-the-kth-smallest-sum-of-a-matrix-with-sorted-rows/)
- [2040. 两个有序数组的第 K 小乘积](https://leetcode.cn/problems/kth-smallest-product-of-two-sorted-arrays/)
- [2386. 找出数组的第 K 大和](https://leetcode.cn/problems/find-the-k-sum-of-an-array/)

#### 最小化最大值
- [2439. 最小化数组中的最大值](https://leetcode.cn/problems/minimize-maximum-of-array/)
- [2513. 最小化两个数组中的最大值](https://leetcode.cn/problems/minimize-the-maximum-of-two-arrays/)
- [2560. 打家劫舍 IV](https://leetcode.cn/problems/house-robber-iv/)
- [2616. 最小化数对的最大差值](https://leetcode.cn/problems/minimize-the-maximum-difference-of-pairs/)

#### 最大化最小值
- [1552. 两球之间的磁力](https://leetcode.cn/problems/magnetic-force-between-two-balls/)
- [2517. 礼盒的最大甜蜜度](https://leetcode.cn/problems/maximum-tastiness-of-candy-basket/)
- [2528. 最大化城市的最小供电站数目](https://leetcode.cn/problems/maximize-the-minimum-powered-city/)

[往期每日一题题解（按 tag 分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

---

欢迎关注[ biIibiIi@灵茶山艾府](https://space.bilibili.com/206214)，高质量算法教学，持续输出中~
