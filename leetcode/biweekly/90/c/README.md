下午两点在B站讲这场双周赛的题目，[欢迎关注](https://space.bilibili.com/206214)~

---

根据题意，模 $\textit{space}$ 相同的数是可以被这些数中最小的那个数摧毁的。例如 $\textit{space}=10$，那么 $3,13,23,33,\cdots$ 都可以被 $3$ 摧毁。

因此按模 $\textit{space}$ 的结果分组，最大组的最小值就是答案。

```py [sol1-Python3]
class Solution:
    def destroyTargets(self, nums: List[int], space: int) -> int:
        g = defaultdict(list)
        for x in nums:
            g[x % space].append(x)
        mx = ans = 0
        for a in g.values():
            m, low = len(a), min(a)
            if m > mx or m == mx and low < ans:
                mx, ans = m, low
        return ans
```

```go [sol1-Go]
func destroyTargets(nums []int, space int) (ans int) {
	g := map[int][]int{}
	for _, x := range nums {
		g[x%space] = append(g[x%space], x)
	}
	mx := 0
	for _, a := range g {
		m := len(a)
		low := a[0]
		for _, x := range a {
			if x < low {
				low = x
			}
		}
		if m > mx || m == mx && low < ans {
			ans = low
			mx = m
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。
