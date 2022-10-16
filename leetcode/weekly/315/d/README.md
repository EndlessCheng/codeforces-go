#### 提示 1

首先考虑一个简单的情况，$\textit{nums}$ 的所有元素都在 $[\textit{minK},\textit{maxK}]$ 范围内。

在这种情况下，相当于要统计同时包含 $\textit{minK}$ 和 $\textit{maxK}$ 的子数组的个数。

我们可以枚举子数组的右端点。遍历 $\textit{nums}$，记录 $\textit{minK}$ 上一次出现的位置 $\textit{minI}$ 和 $\textit{maxK}$ 上一次出现的位置 $\textit{maxI}$，当遍历到 $\textit{nums}[i]$ 时，如果 $\textit{minK}$ 和 $\textit{maxK}$ 之前出现过，则左端点 $\le\min(\textit{minI},\textit{maxI})$ 的子数组都是合法的，合法子数组的个数为 $\min(\textit{minI},\textit{maxI})+1$。

#### 提示 2

回到原问题，由于子数组不能包含在 $[\textit{minK},\textit{maxK}]$ 范围之外的元素，因此我们还需要记录上一个在 $[\textit{minK},\textit{maxK}]$ 范围之外的 $\textit{nums}[i]$ 的下标，记作 $i_0$。此时合法子数组的个数为 $\min(\textit{minI},\textit{maxI})-i_0$。

代码实现时：

- 为方便计算，可以初始化 $\textit{minI},\ \textit{maxI},\ i_0$ 均为 $-1$。
- 如果 $\min(\textit{minI},\textit{maxI})-i_0 < 0$，则表示在 $i_0$ 右侧 $\textit{minK}$ 和 $\textit{maxK}$ 没有同时出现，此时合法子数组的个数为 $0$。

```py [sol1-Python3]
class Solution:
    def countSubarrays(self, nums: List[int], min_k: int, max_k: int) -> int:
        ans = 0
        min_i = max_i = i0 = -1
        for i, x in enumerate(nums):
            if x == min_k: min_i = i
            if x == max_k: max_i = i
            if not min_k <= x <= max_k: i0 = i  # 子数组不能包含 nums[i0]
            ans += max(min(min_i, max_i) - i0, 0)
            # 注：上面这行代码，改为手动算 min max 会更快
            # j = min_i if min_i < max_i else max_i
            # if j > i0: ans += j - i0
        return ans
```

```java [sol1-Java]
class Solution {
    public long countSubarrays(int[] nums, int minK, int maxK) {
        var ans = 0L;
        int n = nums.length, minI = -1, maxI = -1, i0 = -1;
        for (var i = 0; i < n; ++i) {
            var x = nums[i];
            if (x == minK) minI = i;
            if (x == maxK) maxI = i;
            if (x < minK || x > maxK) i0 = i; // 子数组不能包含 nums[i0]
            ans += Math.max(Math.min(minI, maxI) - i0, 0);
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    long long countSubarrays(vector<int> &nums, int min_k, int max_k) {
        long long ans = 0L;
        int n = nums.size(), min_i = -1, max_i = -1, i0 = -1;
        for (int i = 0; i < n; ++i) {
            int x = nums[i];
            if (x == min_k) min_i = i;
            if (x == max_k) max_i = i;
            if (x < min_k || x > max_k) i0 = i; // 子数组不能包含 nums[i0]
            ans += max(min(min_i, max_i) - i0, 0);
        }
        return ans;
    }
};
```

```go [sol1-Go]
func countSubarrays(nums []int, minK, maxK int) (ans int64) {
	minI, maxI, i0 := -1, -1, -1
	for i, x := range nums {
		if x == minK {
			minI = i
		}
		if x == maxK {
			maxI = i
		}
		if x < minK || x > maxK {
			i0 = i // 子数组不能包含 nums[i0]
		}
		ans += int64(max(min(minI, maxI)-i0, 0))
	}
	return
}

func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干变量。
