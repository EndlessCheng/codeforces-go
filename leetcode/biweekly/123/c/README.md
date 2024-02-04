[视频讲解](https://www.bilibili.com/video/BV14C411r7nN/) 第三题。

## 前置知识：前缀和

对于数组 $a$，定义它的前缀和 $\textit{s}[0]=0$，$\textit{s}[i+1] = \sum\limits_{j=0}^{i}a[j]$。

根据这个定义，有 $s[i+1]=s[i]+a[i]$。

例如 $a=[1,2,1,2]$，对应的前缀和数组为 $s=[0,1,3,4,6]$。

通过前缀和，我们可以把**连续子数组的元素和转换成两个前缀和的差**，$a[\textit{left}]$ 到 $a[\textit{right}]$ 的元素和等于

$$
\sum_{j=\textit{left}}^{\textit{right}}a[j] = \sum\limits_{j=0}^{\textit{right}}a[j] - \sum\limits_{j=0}^{\textit{left}-1}a[j] = \textit{s}[\textit{right}+1] - \textit{s}[\textit{left}]
$$

例如 $a$ 的子数组 $[2,1,2]$ 的和就可以用 $s[4]-s[1]=6-1=5$ 算出来。

**注**：$s[0]=0$ 表示一个空数组的元素和。为什么要额外定义它？想一想，如果要计算的子数组恰好是一个前缀（从 $a[0]$ 到 $a[\textit{right}]$），你要用 $s[\textit{right}+1]$ 减去谁呢？通过定义 $s[0]=0$，任意子数组（包括前缀）都可以表示为两个前缀和的差。

## 思路

为方便描述，把 $\textit{nums}$ 简称为 $a$。

子数组 $a[i..j]$ 的元素和为 

$$
s[j+1]-s[i]
$$

枚举 $j$，我们需要找到最小的 $s[i]$，满足 $|a[i]-a[j]|=k$，即 $a[i] = a[j]-k$ 或 $a[i]=a[j]+k$。

定义哈希表 $\textit{minS}$，键为 $a[i]$，值为相同 $a[i]$ 下的 $s[i]$ 的最小值。

子数组最后一个数为 $a[j]$ 时，子数组的最大元素和为

$$
\begin{aligned}
& s[j+1]-\textit{minS}[a[i]]\\
=\ &s[j+1]-\min(\textit{minS}[a[j]-k],\textit{minS}[a[j]+k])
\end{aligned}
$$

```py [sol-Python3]
class Solution:
    def maximumSubarraySum(self, nums: List[int], k: int) -> int:
        ans = -inf
        min_s = defaultdict(lambda: inf)
        s = 0
        for x in nums:
            ans = max(ans, s + x - min(min_s[x - k], min_s[x + k]))
            min_s[x] = min(min_s[x], s)
            s += x
        return ans if ans > -inf else 0
```

```java [sol-Java]
class Solution {
    public long maximumSubarraySum(int[] nums, int k) {
        long ans = Long.MIN_VALUE;
        long sum = 0;
        Map<Integer, Long> minS = new HashMap<>();
        for (int x : nums) {
            long s1 = minS.getOrDefault(x - k, Long.MAX_VALUE / 2);
            long s2 = minS.getOrDefault(x + k, Long.MAX_VALUE / 2);
            ans = Math.max(ans, sum + x - Math.min(s1, s2));
            minS.merge(x, sum, Math::min);
            sum += x;
        }
        return ans > Long.MIN_VALUE / 4 ? ans : 0;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumSubarraySum(vector<int> &nums, int k) {
        long long ans = LLONG_MIN, sum = 0;
        unordered_map<int, long long> min_s;
        for (int x: nums) {
            auto it = min_s.find(x + k);
            if (it != min_s.end()) {
                ans = max(ans, sum + x - it->second);
            }

            it = min_s.find(x - k);
            if (it != min_s.end()) {
                ans = max(ans, sum + x - it->second);
            }

            it = min_s.find(x);
            if (it == min_s.end() || sum < it->second) {
                min_s[x] = sum;
            }

            sum += x;
        }
        return ans == LLONG_MIN ? 0 : ans;
    }
};
```

```go [sol-Go]
func maximumSubarraySum(nums []int, k int) int64 {
	ans := math.MinInt
	minS := map[int]int{}
	sum := 0
	for _, x := range nums {
		s, ok := minS[x+k]
		if ok {
			ans = max(ans, sum+x-s)
		}

		s, ok = minS[x-k]
		if ok {
			ans = max(ans, sum+x-s)
		}

		s, ok = minS[x]
		if !ok || sum < s {
			minS[x] = sum
		}

		sum += x
	}
	if ans == math.MinInt {
		return 0
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $a$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 练习：前缀和与哈希表

- [930. 和相同的二元子数组](https://leetcode.cn/problems/binary-subarrays-with-sum/) 1592
- [560. 和为 K 的子数组](https://leetcode.cn/problems/subarray-sum-equals-k/)
- [1524. 和为奇数的子数组数目](https://leetcode.cn/problems/number-of-sub-arrays-with-odd-sum/) 1611
- [974. 和可被 K 整除的子数组](https://leetcode.cn/problems/subarray-sums-divisible-by-k/) 1676
- [523. 连续的子数组和](https://leetcode.cn/problems/continuous-subarray-sum/)
- [525. 连续数组](https://leetcode.cn/problems/contiguous-array/)
- [1124. 表现良好的最长时间段](https://leetcode.cn/problems/longest-well-performing-interval/) 1908
- [2488. 统计中位数为 K 的子数组](https://leetcode.cn/problems/count-subarrays-with-median-k/) 1999
- [1590. 使数组和能被 P 整除](https://leetcode.cn/problems/make-sum-divisible-by-p/) 2039
- [2949. 统计美丽子字符串 II](https://leetcode.cn/problems/count-beautiful-substrings-ii/) 2445
- [面试题 17.05. 字母与数字](https://leetcode.cn/problems/find-longest-subarray-lcci/)
- [1983. 范围和相等的最宽索引对](https://leetcode.cn/problems/widest-pair-of-indices-with-equal-range-sum/)（会员题）
- [2489. 固定比率的子字符串数](https://leetcode.cn/problems/number-of-substrings-with-fixed-ratio/)（会员题）

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
