### 本题视频讲解

见[【周赛 340】](https://www.bilibili.com/video/BV1iN411w7my/)。

### 前置知识：二分

见 [二分查找【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

### 提示 1

看到「最大化最小值」或者「最小化最大值」就要想到**二分答案**，这是一个固定的套路。

为什么？一般来说，二分的值越大，越能/不能满足要求；二分的值越小，越不能/能满足要求，有单调性，可以二分。

> 更多题目见文末的题单。

### 提示 2

二分数对中的最大差值 $mx$。

由于下标和答案无关，可以先排序。为了让匹配的数对尽量多，应尽量选相邻的元素，这样更能满足要求。例如 $[1,2,3,4]$，如果 $1,3$ 匹配，$2,4$ 匹配，最大差值是 $2$；而如果 $1,2$ 相邻匹配，$3,4$ 相邻匹配，最大差值只有 $1$。

我们来算一算最多能匹配多少个数对：

- 如果可以选 $\textit{nums}[0]$ 和 $\textit{nums}[1]$，那么答案等于「$n-2$ 个数的最多数对个数」$+1$。
- 如果不选 $\textit{nums}[0]$，那么答案等于「$n-1$ 个数的最多数对个数」。
- 这两种情况取最大值。

这看上去很像 [198. 打家劫舍](https://leetcode.cn/problems/house-robber/)，可以用动态规划实现。

也可以用贪心做：

- 注意到，「$n-1$ 个数的最多数对个数」不会超过「$n-3$ 个数的最多数对个数」$+1$。这里 $+1$ 表示选 $\textit{nums}[1]$ 和 $\textit{nums}[2]$。
- 由于「$n-2$ 个数的最多数对个数」$\ge$「$n-3$ 个数的最多数对个数」，所以如果可以选 $\textit{nums}[0]$ 和 $\textit{nums}[1]$，那么直接选就行。
- 依此类推，不断缩小问题规模。所以遍历一遍数组就能求出最多数对个数，具体见代码。

```py [sol1-Python3]
class Solution:
    def minimizeMax(self, nums: List[int], p: int) -> int:
        nums.sort()
        def f(mx: int) -> int:
            cnt = i = 0
            while i < len(nums) - 1:
                if nums[i + 1] - nums[i] <= mx:  # 都选
                    cnt += 1
                    i += 2
                else:  # 不选 nums[i]
                    i += 1
            return cnt
        return bisect_left(range(nums[-1] - nums[0]), p, key=f)
```

```java [sol1-Java]
class Solution {
    public int minimizeMax(int[] nums, int p) {
        Arrays.sort(nums);
        int n = nums.length, left = -1, right = nums[n - 1] - nums[0]; // 开区间
        while (left + 1 < right) { // 开区间
            int mid = (left + right) >>> 1, cnt = 0;
            for (int i = 0; i < n - 1; ++i)
                if (nums[i + 1] - nums[i] <= mid) { // 都选
                    ++cnt;
                    ++i;
                }
            if (cnt >= p) right = mid;
            else left = mid;
        }
        return right;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int minimizeMax(vector<int> &nums, int p) {
        sort(nums.begin(), nums.end());
        int left = -1, right = nums.back() - nums[0]; // 开区间
        while (left + 1 < right) { // 开区间
            int mid = left + (right - left) / 2, cnt = 0;
            for (int i = 0; i < nums.size() - 1; ++i)
                if (nums[i + 1] - nums[i] <= mid) { // 都选
                    ++cnt;
                    ++i;
                }
            (cnt >= p ? right : left) = mid;
        }
        return right;
    }
};
```

```go [sol1-Go]
func minimizeMax(nums []int, p int) int {
	sort.Ints(nums)
	n := len(nums)
	return sort.Search(nums[n-1]-nums[0], func(mx int) bool {
		cnt := 0
		for i := 0; i < n-1; i++ {
			if nums[i+1]-nums[i] <= mx { // 都选
				cnt++
				i++
			}
		}
		return cnt >= p
	})
}
```

### 复杂度分析

- 时间复杂度：$O(n\log n + n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})-\min(\textit{nums})$。
- 空间复杂度：$O(1)$。忽略排序时的栈空间，仅用到若干额外变量。

### 二分答案·题单

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

#### 最小化最大值
- [2439. 最小化数组中的最大值](https://leetcode.cn/problems/minimize-maximum-of-array/)
- [2513. 最小化两个数组中的最大值](https://leetcode.cn/problems/minimize-the-maximum-of-two-arrays/)
- [2560. 打家劫舍 IV](https://leetcode.cn/problems/house-robber-iv/)
- [2616. 最小化数对的最大差值](https://leetcode.cn/problems/minimize-the-maximum-difference-of-pairs/)

#### 最大化最小值
- [1552. 两球之间的磁力](https://leetcode.cn/problems/magnetic-force-between-two-balls/)
- [2517. 礼盒的最大甜蜜度](https://leetcode.cn/problems/maximum-tastiness-of-candy-basket/)
- [2528. 最大化城市的最小供电站数目](https://leetcode.cn/problems/maximize-the-minimum-powered-city/)
