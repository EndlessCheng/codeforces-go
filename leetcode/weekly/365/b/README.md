## 方法一：枚举 j

枚举 $j$，为了让 $(\textit{nums}[i] - \textit{nums}[j]) * \textit{nums}[k]$ 尽量大，我们需要知道 $j$ 左侧元素的最大值，和 $j$ 右侧元素的最大值。

也就是 $\textit{nums}$ 的前缀最大值 $\textit{preMax}$ 和后缀最大值 $\textit{sufMax}$，这都可以用递推预处理出来：

- $\textit{preMax}[i] = \max(\textit{preMax}[i-1], \textit{nums}[i])$
- $\textit{sufMax}[i] = \max(\textit{sufMax}[i+1], \textit{nums}[i])$

代码实现时，可以只预处理 $\textit{sufMax}$ 数组，$\textit{preMax}$ 可以在计算答案的同时算出来。

```py [sol-Python3]
class Solution:
    def maximumTripletValue(self, nums: List[int]) -> int:
        n = len(nums)
        suf_max = [0] * (n + 1)
        for i in range(n - 1, 1, -1):
            suf_max[i] = max(suf_max[i + 1], nums[i])
        ans = pre_max = 0
        for j, x in enumerate(nums):
            ans = max(ans, (pre_max - x) * suf_max[j + 1])
            pre_max = max(pre_max, x)
        return ans
```

```java [sol-Java]
class Solution {
    public long maximumTripletValue(int[] nums) {
        int n = nums.length;
        int[] sufMax = new int[n + 1];
        for (int i = n - 1; i > 1; i--) {
            sufMax[i] = Math.max(sufMax[i + 1], nums[i]);
        }
        long ans = 0;
        int preMax = nums[0];
        for (int j = 1; j < n - 1; j++) {
            ans = Math.max(ans, (long) (preMax - nums[j]) * sufMax[j + 1]);
            preMax = Math.max(preMax, nums[j]);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumTripletValue(vector<int> &nums) {
        int n = nums.size();
        vector<int> suf_max(n + 1, 0);
        for (int i = n - 1; i > 1; i--) {
            suf_max[i] = max(suf_max[i + 1], nums[i]);
        }
        long long ans = 0;
        int pre_max = nums[0];
        for (int j = 1; j < n - 1; j++) {
            ans = max(ans, (long long) (pre_max - nums[j]) * suf_max[j + 1]);
            pre_max = max(pre_max, nums[j]);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumTripletValue(nums []int) int64 {
	ans := 0
	n := len(nums)
	sufMax := make([]int, n+1)
	for i := n - 1; i > 1; i-- {
		sufMax[i] = max(sufMax[i+1], nums[i])
	}
	preMax := 0
	for j, x := range nums {
		ans = max(ans, (preMax-x)*sufMax[j+1])
		preMax = max(preMax, x)
	}
	return int64(ans)
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：枚举 k

枚举 $k$，我们需要知道 $k$ 左边 $\textit{nums}[i] - \textit{nums}[j]$ 的最大值。

类似 [121. 买卖股票的最佳时机](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/)，我们可以在遍历的过程中，维护 $\textit{nums}[i]$ 的最大值 $\textit{preMax}$，同时维护 $\textit{preMax}$ 减当前元素的最大值 $\textit{maxDiff}$，这就是 $k$ 左边 $\textit{nums}[i] - \textit{nums}[j]$ 的最大值。

```py [sol-Python3]
class Solution:
    def maximumTripletValue(self, nums: List[int]) -> int:
        ans = max_diff = pre_max = 0
        for x in nums:
            ans = max(ans, max_diff * x)
            max_diff = max(max_diff, pre_max - x)
            pre_max = max(pre_max, x)
        return ans
```

```java [sol-Java]
class Solution {
    public long maximumTripletValue(int[] nums) {
        long ans = 0;
        int maxDiff = 0, preMax = 0;
        for (int x : nums) {
            ans = Math.max(ans, (long) maxDiff * x);
            maxDiff = Math.max(maxDiff, preMax - x);
            preMax = Math.max(preMax, x);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumTripletValue(vector<int> &nums) {
        long long ans = 0;
        int max_diff = 0, pre_max = 0;
        for (int x : nums) {
            ans = max(ans, (long long) max_diff * x);
            max_diff = max(max_diff, pre_max - x);
            pre_max = max(pre_max, x);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumTripletValue(nums []int) int64 {
	ans, maxDiff, preMax := 0, 0, 0
	for _, x := range nums {
		ans = max(ans, maxDiff*x)
		maxDiff = max(maxDiff, preMax-x)
		preMax = max(preMax, x)
	}
	return int64(ans)
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

