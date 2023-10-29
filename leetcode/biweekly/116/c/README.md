请看视频讲解[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)

小优化：例如 $\textit{nums}$ 前两个数的和是 $5$，那么枚举 $j$ 的时候，最大只需要枚举 $5$，而不是 $\textit{target}$。

```py [sol-Python3]
class Solution:
    def lengthOfLongestSubsequence(self, nums: List[int], target: int) -> int:
        f = [0] + [-inf] * target
        s = 0
        for x in nums:
            s = min(s + x, target)
            for j in range(s, x - 1, -1):
                # f[j] = max(f[j], f[j - x] + 1)
                # 手写 max 效率更高
                if f[j] < f[j - x] + 1:
                    f[j] = f[j - x] + 1
        return f[-1] if f[-1] > 0 else -1
```

```java [sol-Java]
class Solution {
    public int lengthOfLongestSubsequence(List<Integer> nums, int target) {
        int[] f = new int[target + 1];
        Arrays.fill(f, Integer.MIN_VALUE);
        f[0] = 0;
        int s = 0;
        for (int x : nums) {
            s = Math.min(s + x, target);
            for (int j = s; j >= x; j--) {
                f[j] = Math.max(f[j], f[j - x] + 1);
            }
        }
        return f[target] > 0 ? f[target] : -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int lengthOfLongestSubsequence(vector<int> &nums, int target) {
        vector<int> f(target + 1, INT_MIN);
        f[0] = 0;
        int s = 0;
        for (int x : nums) {
            s = min(s + x, target);
            for (int j = s; j >= x; j--) {
                f[j] = max(f[j], f[j - x] + 1);
            }
        }
        return f[target] > 0 ? f[target] : -1;
    }
};
```

```go [sol-Go]
func lengthOfLongestSubsequence(nums []int, target int) int {
	f := make([]int, target+1)
	for i := 1; i <= target; i++ {
		f[i] = math.MinInt
	}
	s := 0
	for _, x := range nums {
		s = min(s+x, target)
		for j := s; j >= x; j-- {
			f[j] = max(f[j], f[j-x]+1)
		}
	}
	if f[target] > 0 {
		return f[target]
	}
	return -1
}

func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\cdot \textit{target})$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(\textit{target})$。
