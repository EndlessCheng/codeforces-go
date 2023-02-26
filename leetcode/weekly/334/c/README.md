# 方法一：二分答案

### 前置知识：二分

见[【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

### 提示 1

如果 $2\cdot\textit{nums}\le \textit{nums}[j]$，则称 $\textit{nums}[i]$ 和 $\textit{nums}[j]$ **匹配**。

如果可以匹配 $k$ 对，那么也可以匹配小于 $k$ 对，去掉一些数对即可做到。

如果无法匹配 $k$ 对，那么也无法匹配大于 $k$ 对（反证法）。

因此答案有单调性，可以二分。

### 提示 2

检测能不能匹配 $k$ 对。

要让哪些数匹配呢？

**结论：从小到大排序后，如果存在 $k$ 对匹配，那么一定可以让最小的 $k$ 个数和最大的 $k$ 个数匹配。**

证明：假设不是最小的 $k$ 个数和最大的 $k$ 个数匹配，那么我们总是可以把 $\textit{nums}[i]$ 替换成比它小的且不在匹配中的数，这仍然是匹配的；同理，把 $\textit{nums}[j]$ 替换成比它大的且不在匹配中的数，这仍然是匹配的。所以如果存在 $k$ 对匹配，那么一定可以让最小的 $k$ 个数和最大的 $k$ 个数匹配。

反过来说，如果最小的 $k$ 个数无法和最大的 $k$ 个数匹配，则任意 $k$ 对都无法匹配。

从小到大排序后，$\textit{nums}[0]$ 肯定要匹配 $\textit{nums}[n-k]$。如果不这样做，$\textit{nums}[0]$ 匹配了在 $\textit{nums}[n-k]$ 右侧的数，相当于占了一个位置，那么一定有个更大的 $\textit{nums}[i]$ 要匹配 $\textit{nums}[n-k]$，这不一定能匹配上。

所以 $\textit{nums}[i]$ 一定要匹配 $\textit{nums}[n-k+i]$。

如果对于所有的 $0\le i < k$ 都有 $2\cdot\textit{nums}[i]\le\textit{nums}[n-k+i]$，那么可以匹配 $k$ 对。

```py [sol1-Python3]
class Solution:
    def maxNumOfMarkedIndices(self, nums: List[int]) -> int:
        nums.sort()
        left, right = 0, len(nums) // 2 + 1  # 开区间
        while left + 1 < right:
            k = (left + right) // 2
            if all(nums[i] * 2 <= nums[i - k] for i in range(k)):
                left = k
            else:
                right = k
        return left * 2
```

```java [sol1-Java]
class Solution {
    public int maxNumOfMarkedIndices(int[] nums) {
        Arrays.sort(nums);
        int left = 0, right = nums.length / 2 + 1; // 开区间
        while (left + 1 < right) {
            int mid = (left + right) >>> 1;
            if (check(nums, mid)) left = mid;
            else right = mid;
        }
        return left * 2;
    }

    private boolean check(int[] nums, int k) {
        for (int i = 0; i < k; ++i)
            if (nums[i] * 2 > nums[nums.length - k + i])
                return false;
        return true;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int maxNumOfMarkedIndices(vector<int> &nums) {
        sort(nums.begin(), nums.end());

        auto check = [&](int k) -> bool {
            for (int i = 0; i < k; ++i)
                if (nums[i] * 2 > nums[nums.size() - k + i])
                    return false;
            return true;
        };

        int left = 0, right = nums.size() / 2 + 1; // 开区间
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            (check(mid) ? left : right) = mid;
        }
        return left * 2;
    }
};
```

```go [sol1-Go]
func maxNumOfMarkedIndices(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return 2 * sort.Search(n/2, func(k int) bool {
		k++
		for i, x := range nums[:k] {
			if x*2 > nums[n-k+i] {
				return true
			}
		}
		return false
	})
}
```

### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$。忽略排序的栈开销，仅用到若干额外变量。

# 方法二：双指针

从方法一的匹配方式可知，我们需要用左半部分的数，去匹配右半部分的数。

从右半部分中，找到第一个满足 $2\cdot\textit{nums}[0]\le \textit{nums}[j]$ 的 $j$，那么 $\textit{nums}[1]$ 只能匹配右半部分中的下标大于 $j$ 的数，依此类推。

这可以用双指针实现。

附：[视频讲解](https://www.bilibili.com/video/BV1wj411G7sH/)

```py [sol2-Python3]
class Solution:
    def maxNumOfMarkedIndices(self, nums: List[int]) -> int:
        nums.sort()
        i = 0
        for x in nums[(len(nums) + 1) // 2:]:
            if nums[i] * 2 <= x:
                i += 1
        return i * 2
```

```java [sol2-Java]
class Solution {
    public int maxNumOfMarkedIndices(int[] nums) {
        Arrays.sort(nums);
        int i = 0, n = nums.length;
        for (int j = (n + 1) / 2; j < n; ++j)
            if (nums[i] * 2 <= nums[j])
                ++i;
        return i * 2;
    }
}
```

```cpp [sol2-C++]
class Solution {
public:
    int maxNumOfMarkedIndices(vector<int> &nums) {
        sort(nums.begin(), nums.end());
        int i = 0, n = nums.size();
        for (int j = (n + 1) / 2; j < n; ++j)
            if (nums[i] * 2 <= nums[j])
                ++i;
        return i * 2;
    }
};
```

```go [sol2-Go]
func maxNumOfMarkedIndices(nums []int) int {
	sort.Ints(nums)
	i := 0
	for _, x := range nums[(len(nums)+1)/2:] {
		if nums[i]*2 <= x {
			i++
		}
	}
	return i * 2
}
```

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$。忽略排序的栈开销，仅用到若干额外变量。
