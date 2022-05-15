两个特殊楼层间的楼层都是连续的，取最大值作为答案。

代码实现时可以把 $\textit{bottom}-1$ 和 $\textit{top}+1$ 视作两个特殊楼层，从而简化代码逻辑。

```go
func maxConsecutive(bottom, top int, a []int) (ans int) {
	a = append(a, bottom-1, top+1)
	sort.Ints(a)
	for i := 1; i < len(a); i++ {
		ans = max(ans, a[i]-a[i-1]-1)
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```
