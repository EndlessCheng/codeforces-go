[视频讲解](https://www.bilibili.com/video/BV1uG4y157Rc)。

枚举所有 $[0,n]$ 内的关门时间，我们需要知道 $j$ 前面的 `N` 的个数以及 $j$ 及其后面的 `Y` 的个数。

我们可以先统计出 $\textit{customers}$ 中 `Y` 的个数，即 $j=0$ 的代价。然后枚举 $[1,n]$ 内的 $j$，如果 $\textit{customers}[j-1]$ 是 `N`，则代价加一，否则代价减一。

遍历中，代价的最小值对应的 $j$ 即为答案。

```py [sol1-Python3]
class Solution:
    def bestClosingTime(self, customers: str) -> int:
        ans = 0
        min_cost = cost = customers.count('Y')
        for i, c in enumerate(customers, 1):
            if c == 'N': 
                cost += 1
            else:
                cost -= 1
                if cost < min_cost:
                    min_cost = cost
                    ans = i
        return ans
```

```go [sol1-Go]
func bestClosingTime(customers string) (ans int) {
	cost := strings.Count(customers, "Y")
	minCost := cost
	for i, c := range customers {
		if c == 'N' {
			cost++
		} else {
			cost--
			if cost < minCost {
				minCost = cost
				ans = i + 1
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{customers}$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干额外变量。
