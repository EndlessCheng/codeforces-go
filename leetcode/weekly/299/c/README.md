下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

设 $s_1 = \sum\textit{nums}_1$

交换子数组后，对于 $\textit{nums}'_1$ 有

$$
\sum\textit{nums}'_1 = s_1 - (\textit{nums}_1[\textit{left}] + \cdots + \textit{nums}_1[\textit{right}]) + (\textit{nums}_2[\textit{left}] + \cdots + \textit{nums}_2[\textit{right}])
$$

上式可以变形为

$$
s_1 + (\textit{nums}_2[\textit{left}]-\textit{nums}_1[\textit{left}]) + \cdots + (\textit{nums}_2[\textit{right}]-\textit{nums}_1[\textit{right}])
$$

设 $\textit{diff}[i] = \textit{nums}_2[i]-\textit{nums}_1[i]$，那么问题可以转换成求 $\textit{diff}$ 数组的 [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/)。

对于 $\textit{nums}_2$ 也同理，求这两者的最大值，即为答案。

```go [sol1-Go]
func maxSubarraySum(a []int) int {
	maxS, s := 0, 0
	for _, v := range a {
		s = max(s+v, 0)
		maxS = max(maxS, s)
	}
	return maxS
}

func maximumsSplicedArray(nums1, nums2 []int) (ans int) {
	n := len(nums1)
	f := func(a, b []int) {
		diff := make([]int, n)
		for i, v := range b {
			diff[i] = v - a[i]
		}
		sum := 0
		for _, v := range a {
			sum += v
		}
		sum += maxSubarraySum(diff)
		ans = max(ans, sum)
	}
	f(nums1, nums2)
	f(nums2, nums1)
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

