### 本题视频讲解

见[【双周赛 102】](https://www.bilibili.com/video/BV1Es4y1N7v1/) 第二题。

### 思路

一边遍历，一边计算前缀最大值 $\textit{mx}$，以及前缀的得分之和 $\textit{s}$。

```py [sol1-Python3]
class Solution:
    def findPrefixScore(self, nums: List[int]) -> List[int]:
        ans = []
        mx = s = 0
        for x in nums:
            mx = max(mx, x)  # 前缀最大值
            s += x + mx  # 累加前缀的得分
            ans.append(s)
        return ans
```

```go [sol1-Go]
func findPrefixScore(nums []int) []int64 {
	ans := make([]int64, len(nums))
	mx, s := 0, 0
	for i, x := range nums {
		mx = max(mx, x) // 前缀最大值
		s += x + mx // 累加前缀的得分
		ans[i] = int64(s)
	}
	return ans
}

func max(a, b int) int { if a < b { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$。返回值不计入，仅用到若干额外变量。
