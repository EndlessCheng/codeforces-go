[本题视频讲解](https://www.bilibili.com/video/BV1994y1A7oo/)

## 前置知识：中位数贪心

为方便描述，将 $\textit{nums}$ 简记为 $a$。

**定理**：将 $a$ 的所有元素变为 $a$ 的**中位数**是最优的。

**证明**：设 $a$ 的长度为 $n$，设要将所有 $a[i]$ 变为 $x$。假设 $a$ 已经从小到大排序。首先，如果 $x$ 取在区间 $[a[0],a[n-1]]$ 之外，那么 $x$ 向区间方向移动可以使距离和变小；同时，如果 $x$ 取在区间 $[a[0],a[n-1]]$ 之内，无论如何移动 $x$，它到 $a[0]$ 和 $a[n-1]$ 的距离和都是一个定值 $a[n-1]-a[0]$，那么去掉 $a[0]$ 和 $a[n-1]$ 这两个最左最右的数，问题规模缩小。不断缩小问题规模，如果最后剩下 $1$ 个数，那么 $x$ 就取它；如果最后剩下 $2$ 个数，那么 $x$ 取这两个数之间的任意值都可以（包括这两个数）。因此 $x$ 可以取 $a[n/2]$。

## 前置知识：前缀和

对于数组 $\textit{nums}$，定义它的前缀和 $\textit{s}[0]=0$，$\textit{s}[i+1] = \sum\limits_{j=0}^{i}\textit{nums}[j]$。

根据这个定义，有 $s[i+1]=s[i]+\textit{nums}[i]$。

例如 $\textit{nums}=[1,2,1,2]$，对应的前缀和数组为 $s=[0,1,3,4,6]$。

通过前缀和，我们可以把**子数组的元素和转换成两个前缀和的差**，即

$$
\sum_{j=\textit{left}}^{\textit{right}}\textit{nums}[j] = \sum\limits_{j=0}^{\textit{right}}\textit{nums}[j] - \sum\limits_{j=0}^{\textit{left}-1}\textit{nums}[j] = \textit{s}[\textit{right}+1] - \textit{s}[\textit{left}]
$$

例如 $\textit{nums}$ 的子数组 $[2,1,2]$ 的和就可以用 $s[4]-s[1]=6-1=5$ 算出来。

**注**：为方便计算，常用左闭右开区间 $[\textit{left},\textit{right})$ 来表示从 $\textit{nums}[\textit{left}]$ 到 $\textit{nums}[\textit{right}-1]$ 的子数组，此时子数组的和为 $\textit{s}[\textit{right}] - \textit{s}[\textit{left}]$，子数组的长度为 $\textit{right}-\textit{left}$。

**注 2**：$s[0]=0$ 表示一个空数组的元素和。为什么要额外定义它？想一想，如果要计算的子数组恰好是一个前缀（从 $\textit{nums}[0]$ 开始），你要用 $s[\textit{right}]$ 减去谁呢？通过定义 $s[0]=0$，任意子数组（包括前缀）都可以表示为两个前缀和的差。

## 前置知识：滑动窗口

请看视频[【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)

## 方法一：滑动窗口+前缀和

把数组排序后，要变成一样的数必然在一个连续子数组中，那么用**滑动窗口**来做，枚举子数组的右端点 $\textit{right}$，然后维护子数组的左端点 $\textit{left}$。

根据中位数贪心，最优做法是把子数组内的元素都变成子数组的中位数，操作次数如果超过 $k$，就必须移动左端点。

求出数组的前缀和，就可以 $\mathcal{O}(1)$ 算出操作次数了，具体请看我之前写的 [这篇题解的图](https://leetcode.cn/problems/minimum-operations-to-make-all-array-elements-equal/solution/yi-tu-miao-dong-pai-xu-qian-zhui-he-er-f-nf55/)。

```py [sol-Python3]
class Solution:
    def maxFrequencyScore(self, nums: List[int], k: int) -> int:
        nums.sort()

        n = len(nums)
        s = list(accumulate(nums, initial=0))  # nums 的前缀和

        # 把 nums[l] 到 nums[r] 都变成 nums[i]
        def distance_sum(l: int, i: int, r: int) -> int:
            left = nums[i] * (i - l) - (s[i] - s[l])
            right = s[r + 1] - s[i + 1] - nums[i] * (r - i)
            return left + right

        ans = left = 0
        for i in range(n):
            while distance_sum(left, (left + i) // 2, i) > k:
                left += 1
            ans = max(ans, i - left + 1)
        return ans
```

```java [sol-Java]
class Solution {
    public int maxFrequencyScore(int[] nums, long k) {
        Arrays.sort(nums);

        int n = nums.length;
        long[] s = new long[n + 1];
        for (int i = 0; i < n; i++) {
            s[i + 1] = s[i] + nums[i];
        }

        int ans = 0, left = 0;
        for (int i = 0; i < n; i++) {
            while (distanceSum(s, nums, left, (left + i) / 2, i) > k) {
                left++;
            }
            ans = Math.max(ans, i - left + 1);
        }
        return ans;
    }

    // 把 nums[l] 到 nums[r] 都变成 nums[i]
    long distanceSum(long[] s, int[] nums, int l, int i, int r) {
        long left = (long) nums[i] * (i - l) - (s[i] - s[l]);
        long right = s[r + 1] - s[i + 1] - (long) nums[i] * (r - i);
        return left + right;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxFrequencyScore(vector<int> &nums, long long k) {
        sort(nums.begin(), nums.end());

        int n = nums.size();
        vector<long long> s(n + 1, 0);
        for (int i = 0; i < n; i++) {
            s[i + 1] = s[i] + nums[i];
        }

        // 把 nums[l] 到 nums[r] 都变成 nums[i]
        auto distance_sum = [&](int l, int i, int r) -> long long {
            long long left = (long long) nums[i] * (i - l) - (s[i] - s[l]);
            long long right = s[r + 1] - s[i + 1] - (long long) nums[i] * (r - i);
            return left + right;
        };

        int ans = 0, left = 0;
        for (int i = 0; i < n; i++) {
            while (distance_sum(left, (left + i) / 2, i) > k) {
                left++;
            }
            ans = max(ans, i - left + 1);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxFrequencyScore(nums []int, K int64) (ans int) {
	k := int(K)
	slices.Sort(nums)

	n := len(nums)
	sum := make([]int, n+1)
	for i, v := range nums {
		sum[i+1] = sum[i] + v
	}

	// 把 nums[l] 到 nums[r] 都变成 nums[i]
	distanceSum := func(l, i, r int) int {
		left := nums[i]*(i-l) - (sum[i] - sum[l])
		right := sum[r+1] - sum[i+1] - nums[i]*(r-i)
		return left + right
	}

	left := 0
	for i := range nums {
		for distanceSum(left, (left+i)/2, i) > k {
			left++
		}
		ans = max(ans, i-left+1)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：滑动窗口+贡献法

进一步优化，可以做到 $\mathcal{O}(1)$ 额外空间。

请看 [视频讲解](https://www.bilibili.com/video/BV1994y1A7oo/)。

总结视频中讲的要点，关键之处在于，随着窗口右端点向右移动，每个数字会经历三个过程：

1. 移入窗口，对操作次数的贡献是 $\textit{nums}[i]$。
2. 变成正中间的数字，对操作次数的贡献是 $0$。
3. 变成左半边的数字，对操作次数的贡献是 $-\textit{nums}[i]$。

当一个数字 $x=\textit{nums}[\textit{right}]$ 移入窗口时：

- 如果 $x$ 移入前的窗口长度是偶数，那么原来窗口第二个中位数 $\textit{mid}$ 会在 $x$ 移入后，变成正中间的数字。原来的贡献是 $\textit{mid}$，现在的贡献是 $0$，那么贡献的变化量就是 $-\textit{mid}$。再算上 $x$ 的贡献，操作次数的变化量是 $x-\textit{mid}$。
- 如果 $x$ 移入前的窗口长度是奇数，那么原来窗口正中间的中位数 $\textit{mid}$ 会在 $x$ 移入后，变成左半边的数字（第一个中位数）。原来的贡献是 $0$，现在的贡献是 $-\textit{mid}$，那么贡献的变化量就是 $-\textit{mid}$。再算上 $x$ 的贡献，操作次数的变化量是 $x-\textit{mid}$。

所以一个数字 $x=\textit{nums}[\textit{right}]$ 移入窗口时，只需要把操作次数增加

$$
x - \textit{mid}_1
$$

其中 $\textit{mid}_1$ 是 $x$ 移入窗口后的第一个中位数（等价于 $x$ 移入前的第二个中位数），即 $\textit{nums}[(\textit{left} + \textit{right}) / 2]$。

对于元素移出窗口的情况，也是类似的。结论是只需要把操作次数增加

$$
\textit{nums}[\textit{left}] - \textit{mid}_2
$$

其中 $\textit{mid}_2$ 是移出 $\textit{nums}[\textit{left}]$ 前的第二个中位数，即 $\textit{nums}[(\textit{left} + \textit{right} + 1) / 2]$。

```py [sol-Python3]
class Solution:
    def maxFrequencyScore(self, nums: List[int], k: int) -> int:
        nums.sort()
        ans = left = s = 0  # s 是窗口元素与窗口中位数的差之和
        for right, x in enumerate(nums):
            s += x - nums[(left + right) // 2]
            while s > k:
                s += nums[left] - nums[(left + right + 1) // 2]
                left += 1
            ans = max(ans, right - left + 1)
        return ans
```

```java [sol-Java]
class Solution {
    public int maxFrequencyScore(int[] nums, long k) {
        Arrays.sort(nums);
        int ans = 0, left = 0;
        long s = 0; // 窗口元素与窗口中位数的差之和
        for (int right = 0; right < nums.length; right++) {
            s += nums[right] - nums[(left + right) / 2];
            while (s > k) {
                s += nums[left] - nums[(left + right + 1) / 2];
                left++;
            }
            ans = Math.max(ans, right - left + 1);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxFrequencyScore(vector<int> &nums, long long k) {
        sort(nums.begin(), nums.end());
        int ans = 0, left = 0;
        long long s = 0; // 窗口元素与窗口中位数的差之和
        for (int right = 0; right < nums.size(); right++) {
            s += nums[right] - nums[(left + right) / 2];
            while (s > k) {
                s += nums[left] - nums[(left + right + 1) / 2];
                left++;
            }
            ans = max(ans, right - left + 1);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxFrequencyScore(nums []int, k int64) int {
	slices.Sort(nums)
	ans, left := 0, 0
	s := int64(0) // 窗口元素与窗口中位数的差之和
	for right, x := range nums {
		s += int64(x - nums[(left+right)/2])
		for s > k {
			s += int64(nums[left] - nums[(left+right+1)/2])
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 中位数贪心（右边数字为难度分）

- [462. 最小操作次数使数组元素相等 II](https://leetcode.cn/problems/minimum-moves-to-equal-array-elements-ii/)
- [2033. 获取单值网格的最小操作数](https://leetcode.cn/problems/minimum-operations-to-make-a-uni-value-grid/) 1672
- [2448. 使数组相等的最小开销](https://leetcode.cn/problems/minimum-cost-to-make-array-equal/) 2005
- [2607. 使子数组元素和相等](https://leetcode.cn/problems/make-k-subarray-sums-equal/) 2071
- [1703. 得到连续 K 个 1 的最少相邻交换次数](https://leetcode.cn/problems/minimum-adjacent-swaps-for-k-consecutive-ones/) 2467
