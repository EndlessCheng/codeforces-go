## 题意（版本 A）

说，有个小偷公司，给小偷定的 KPI 是偷至少 $k$ 间房子，要求偷的房子不能相邻。

张三作为其中的一个小偷，他不想偷太多，否则一不小心就「数额巨大」，这可太刑了。所以张三计划，在他偷过的房子中，偷走的最大金额要尽量地小。

这个最小值是多少呢？

## 题意（版本 B）

给定数组 $\textit{nums}$，从中选择一个长度至少为 $k$ 的子序列 $A$，要求 $A$ 中没有任何元素在 $\textit{nums}$ 中是相邻的。

最小化 $\max(A)$。

## 方法一：二分+DP

有关二分的三种写法，请看[【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。本文采用**开区间**写法。

看到「最大化最小值」或者「最小化最大值」就要想到**二分答案**，这是一个固定的套路。

对于本题，「偷走的最大金额」越小，能偷的房子就越少，反之越多。例如 $\textit{nums}=[1,4,2,3]$，在最大金额为 $2$ 时，$\textit{nums}$ 中只有 $1$ 和 $2$ 是可以偷的；在最大金额为 $4$ 时，$\textit{nums}$ 中 $1,2,3,4$ 都可以偷。

一般地，二分的值越小，越不能/能满足要求；二分的值越大，越能/不能满足要求。有单调性的保证，就可以二分答案了。

把二分中点 $\textit{mid}$ 记作 $\textit{mx}$，仿照 [198. 打家劫舍](https://leetcode.cn/problems/house-robber/)，定义 $f[i]$ 表示从 $\textit{nums}[0]$ 到 $\textit{nums}[i]$ 中偷金额不超过 $\textit{mx}$ 的房屋，最多能偷多少间房屋。如果 $f[n-1]\ge k$ 则表示答案至多为 $\textit{mx}$，否则表示答案必须超过 $\textit{mx}$。

用「选或不选」来分类讨论：

- 不选 $\textit{nums}[i]$：$f[i] = f[i-1]$；
- 选 $\textit{nums}[i]$，前提是 $\textit{nums}[i] \le \textit{mx}$：$f[i] = f[i-2]+1$。

这两取最大值，即

$$
f[i] = \max(f[i-1], f[i-2] + 1)
$$

代码实现时，可以用两个变量滚动计算。具体请看[【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)。

#### 答疑

**问**：有没有可能，二分出来的答案，不在 $\textit{nums}$ 中？

**答**：不可能。二分出来的答案，一定在 $\textit{nums}$ 中。证明如下：

设答案为 $\textit{ans}$，也就是当最大金额为 $\textit{ans}$ 时，可以偷至少 $k$ 间房子。如果 $\textit{ans}$ 不在 $\textit{nums}$ 中，那么当最大金额为 $\textit{ans}-1$ 时，也可以偷至少 $k$ 间房子。这与二分算法相矛盾：根据视频中讲的红蓝染色法，循环结束时，$\textit{ans}$ 和 $\textit{ans}-1$ 的颜色必然是不同的，即 $\textit{ans}$ 可以满足题目要求，而 $\textit{ans}-1$ 不满足题目要求。所以，二分出来的答案，一定在 $\textit{nums}$ 中。

```py [sol1-Python3]
class Solution:
    def minCapability(self, nums: List[int], k: int) -> int:
        # solve(mx) 返回最大金额为 mx 时，最多可以偷多少间房子
        def solve(mx: int) -> int:
            f0 = f1 = 0
            for x in nums:
                if x > mx:
                    f0 = f1
                else:
                    f0, f1 = f1, max(f1, f0 + 1)
            return f1
        return bisect_left(range(max(nums)), k, key=solve)
```

```java [sol1-Java]
class Solution {
    public int minCapability(int[] nums, int k) {
        int left = 0, right = 0;
        for (int x : nums) {
            right = Math.max(right, x);
        }
        while (left + 1 < right) { // 开区间写法
            int mid = (left + right) >>> 1;
            if (check(nums, k, mid)) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }

    private boolean check(int[] nums, int k, int mx) {
        int f0 = 0, f1 = 0;
        for (int x : nums) {
            if (x > mx) {
                f0 = f1;
            } else {
                int tmp = f1;
                f1 = Math.max(f1, f0 + 1);
                f0 = tmp;
            }
        }
        return f1 >= k;
    }
}
```

```cpp [sol1-C++]
class Solution {
    bool check(vector<int> &nums, int k, int mx) {
        int f0 = 0, f1 = 0;
        for (int x: nums) {
            if (x > mx) {
                f0 = f1;
            } else {
                int tmp = f1;
                f1 = max(f1, f0 + 1);
                f0 = tmp;
            }
        }
        return f1 >= k;
    }

public:
    int minCapability(vector<int> &nums, int k) {
        int left = 0, right = *max_element(nums.begin(), nums.end());
        while (left + 1 < right) { // 开区间写法
            int mid = left + (right - left) / 2;
            (check(nums, k, mid) ? right : left) = mid;
        }
        return right;
    }
};
```

```go [sol1-Go]
func minCapability(nums []int, k int) int {
	return sort.Search(1e9, func(mx int) bool {
		f0, f1 := 0, 0
		for _, x := range nums {
			if x <= mx {
				f0, f1 = f1, max(f1, f0+1)
			} else {
				f0 = f1
			}
		}
		return f1 >= k
	})
}

func max(a, b int) int { if b > a { return b }; return a }
```

```js [sol1-JavaScript]
var minCapability = function (nums, k) {
    function check(mx) {
        let f0 = 0, f1 = 0
        for (const x of nums) {
            if (x > mx) {
                f0 = f1;
            } else {
                [f0, f1] = [f1, Math.max(f1, f0 + 1)]
            }
        }
        return f1 >= k;
    }

    let left = 0, right = Math.max(...nums);
    while (left + 1 < right) { // 开区间写法
        const mid = (left + right) >> 1;
        if (check(mid)) {
            right = mid;
        } else {
            left = mid;
        }
    }
    return right;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

## 方法二：二分+贪心

也可以用贪心做。

考虑到只需要计算个数，在从左到右遍历的情况下只要当前房子可以偷，就立刻偷。

例如 $\textit{nums}=[1,2,3,4],\ mx=3$，如果不偷 $\textit{nums}[0]=1$ 去偷 $\textit{nums}[1]=2$，那么只能偷一间房子。而如果偷 $\textit{nums}[0]=1$ 和 $\textit{nums}[2]=3$，就可以偷两间房子。

严格证明如下：

根据递推式 $f[i] = \max(f[i-1], f[i-2] + 1)$ 可知，

$$
f[i] \ge f[i-1]
$$

所以 $f$ 数组是递增的。

此外，$f[i]-f[i-1]$ 不会超过 $1$，毕竟我们统计的是「个数」，从 $f[i-1]$ 到 $f[i]$ 最多增加 $1$。

因此

$$
f[i-2]+1\ge f[i-1]
$$

必然成立。也就是说，如果 $\textit{nums}[i]\le mx$，则

$$
f[i] = f[i-2] + 1
$$

上式表明，在从左到右遍历 $\textit{nums}$ 时，能偷就偷。如果 $\textit{nums}[i]\le mx$，我们可以偷 $\textit{nums}[i]$，并跳过 $\textit{nums}[i+1]$。

```py [sol-Python3]
class Solution:
    def minCapability(self, nums: List[int], k: int) -> int:
        def solve(mx: int) -> int:
            cnt = i = 0
            while i < len(nums):
                if nums[i] > mx:  # 不偷
                    i += 1
                else:  # 立刻偷
                    cnt += 1
                    i += 2  # 跳过下一间房子
            return cnt
        return bisect_left(range(max(nums)), k, key=solve)
```

```java [sol-Java]
class Solution {
    public int minCapability(int[] nums, int k) {
        int left = 0, right = 0;
        for (int x : nums) {
            right = Math.max(right, x);
        }
        while (left + 1 < right) { // 开区间写法
            int mid = (left + right) >>> 1;
            if (check(nums, k, mid)) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }

    private boolean check(int[] nums, int k, int mx) {
        int cnt = 0, n = nums.length;
        for (int i = 0; i < n; i++) {
            if (nums[i] <= mx) {
                cnt++; // 偷 nums[i]
                i++; // 跳过下一间房子
            }
        }
        return cnt >= k;
    }
}
```

```cpp [sol-C++]
class Solution {
    bool check(vector<int> &nums, int k, int mx) {
        int cnt = 0, n = nums.size();
        for (int i = 0; i < n; i++) {
            if (nums[i] <= mx) {
                cnt++; // 偷 nums[i]
                i++; // 跳过下一间房子
            }
        }
        return cnt >= k;
    }

public:
    int minCapability(vector<int> &nums, int k) {
        int left = 0, right = *max_element(nums.begin(), nums.end());
        while (left + 1 < right) { // 开区间写法
            int mid = left + (right - left) / 2;
            (check(nums, k, mid) ? right : left) = mid;
        }
        return right;
    }
};
```

```go [sol-Go]
func minCapability(nums []int, k int) int {
	return sort.Search(1e9, func(mx int) bool {
		cnt, n := 0, len(nums)
		for i := 0; i < n; i++ {
			if nums[i] <= mx {
				cnt++ // 偷 nums[i]
				i++   // 跳过下一间房子
			}
		}
		return cnt >= k
	})
}
```

```js [sol-JavaScript]
var minCapability = function (nums, k) {
    function check(mx) {
        let cnt = 0;
        for (let i = 0; i < nums.length; i++) {
            if (nums[i] <= mx) {
                cnt++; // 偷 nums[i]
                i++; // 跳过下一间房子
            }
        }
        return cnt >= k;
    }

    let left = 0, right = Math.max(...nums);
    while (left + 1 < right) { // 开区间写法
        const mid = (left + right) >> 1;
        if (check(mid)) {
            right = mid;
        } else {
            left = mid;
        }
    }
    return right;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

## 练习

#### 最小化最大值

- [2439. 最小化数组中的最大值](https://leetcode.cn/problems/minimize-maximum-of-array/)
- [2513. 最小化两个数组中的最大值](https://leetcode.cn/problems/minimize-the-maximum-of-two-arrays/)
- [2616. 最小化数对的最大差值](https://leetcode.cn/problems/minimize-the-maximum-difference-of-pairs/)

#### 最大化最小值

- [1552. 两球之间的磁力](https://leetcode.cn/problems/magnetic-force-between-two-balls/)
- [2517. 礼盒的最大甜蜜度](https://leetcode.cn/problems/maximum-tastiness-of-candy-basket/)
- [2528. 最大化城市的最小供电站数目](https://leetcode.cn/problems/maximize-the-minimum-powered-city/)

#### 二分答案
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

#### 第 K 小/大（部分题目也可以用堆解决）

- [373. 查找和最小的 K 对数字](https://leetcode.cn/problems/find-k-pairs-with-smallest-sums/)
- [378. 有序矩阵中第 K 小的元素](https://leetcode.cn/problems/kth-smallest-element-in-a-sorted-matrix/)
- [719. 找出第 K 小的数对距离](https://leetcode.cn/problems/find-k-th-smallest-pair-distance/)
- [786. 第 K 个最小的素数分数](https://leetcode.cn/problems/k-th-smallest-prime-fraction/)
- [1439. 有序矩阵中的第 k 个最小数组和](https://leetcode.cn/problems/find-the-kth-smallest-sum-of-a-matrix-with-sorted-rows/)
- [2040. 两个有序数组的第 K 小乘积](https://leetcode.cn/problems/kth-smallest-product-of-two-sorted-arrays/)
- [2386. 找出数组的第 K 大和](https://leetcode.cn/problems/find-the-k-sum-of-an-array/)

[往期题解精选（按 tag 分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
