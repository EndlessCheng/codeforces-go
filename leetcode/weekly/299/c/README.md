本题 [视频讲解](https://www.bilibili.com/video/BV1pW4y1r7xs) 已出炉，欢迎点赞三连~

---

设 $s_1 = \sum\limits_{i}\textit{nums}_1[i]$

交换 $[\textit{left},\textit{right}]$ 范围内的子数组后，对于 $\textit{nums}'_1$ 有

$$
\sum\limits_{i}\textit{nums}'_1[i] = s_1 - (\textit{nums}_1[\textit{left}] + \cdots + \textit{nums}_1[\textit{right}]) + (\textit{nums}_2[\textit{left}] + \cdots + \textit{nums}_2[\textit{right}])
$$

合并相同下标，等号右侧变形为

$$
s_1 + (\textit{nums}_2[\textit{left}]-\textit{nums}_1[\textit{left}]) + \cdots + (\textit{nums}_2[\textit{right}]-\textit{nums}_1[\textit{right}])
$$

设 $\textit{diff}[i] = \textit{nums}_2[i]-\textit{nums}_1[i]$，那么问题可以转换成求 $\textit{diff}$ 数组的 [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/)（允许子数组为空）。

对于 $\textit{nums}_2$ 也同理，求这两者的最大值，即为答案。

#### 复杂度分析

- 时间复杂度：$O(n)$。
- 空间复杂度：$O(1)$。仅需要几个额外的变量。

```py [sol1-Python3]
class Solution:
    def solve(self, nums1: List[int], nums2: List[int]) -> int:
        maxSum = s = 0
        for d in (y - x for x, y in zip(nums1, nums2)):
            s = max(s + d, 0)
            maxSum = max(maxSum, s)
        return sum(nums1) + maxSum

    def maximumsSplicedArray(self, nums1: List[int], nums2: List[int]) -> int:
        return max(self.solve(nums1, nums2), self.solve(nums2, nums1))
```

```java [sol1-Java]
class Solution {
    public int maximumsSplicedArray(int[] nums1, int[] nums2) {
        return Math.max(solve(nums1, nums2), solve(nums2, nums1));
    }

    int solve(int[] nums1, int[] nums2) {
        int s1 = 0, maxSum = 0;
        for (int i = 0, s = 0; i < nums1.length; ++i) {
            s1 += nums1[i];
            s = Math.max(s + nums2[i] - nums1[i], 0);
            maxSum = Math.max(maxSum, s);
        }
        return s1 + maxSum;
    }
}
```

```cpp [sol1-C++]
class Solution {
    int solve(vector<int> &nums1, vector<int> &nums2) {
        int s1 = 0, maxSum = 0;
        for (int i = 0, s = 0; i < nums1.size(); ++i) {
            s1 += nums1[i];
            s = max(s + nums2[i] - nums1[i], 0);
            maxSum = max(maxSum, s);
        }
        return s1 + maxSum;
    }

public:
    int maximumsSplicedArray(vector<int> &nums1, vector<int> &nums2) {
        return max(solve(nums1, nums2), solve(nums2, nums1));
    }
};
```

```go [sol1-Go]
func solve(nums1, nums2 []int) int {
	s1, maxSum, s := 0, 0, 0
	for i, x := range nums1 {
		s1 += x
		s = max(s+nums2[i]-x, 0)
		maxSum = max(maxSum, s)
	}
	return s1 + maxSum
}

func maximumsSplicedArray(nums1, nums2 []int) int {
	return max(solve(nums1, nums2), solve(nums2, nums1))
}

func max(a, b int) int { if b > a { return b }; return a }
```
