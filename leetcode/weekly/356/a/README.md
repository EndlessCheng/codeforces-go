下午两点[【b站@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，欢迎关注！

```py [sol-Python3]
class Solution:
    def numberOfEmployeesWhoMetTarget(self, hours: List[int], target: int) -> int:
        return sum(h >= target for h in hours)
```

```go [sol-Go]
func numberOfEmployeesWhoMetTarget(hours []int, target int) (ans int) {
	for _, h := range hours {
		if h >= target {
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{hours}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
