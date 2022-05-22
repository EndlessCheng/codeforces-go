按照 $\textit{day}_i$ 排序后，如果相邻两条线段的斜率不同，那么必然需要一条新的线段。

代码实现时，可以假定第一天之前的斜率为 $\dfrac{1}{0}$，从而简化判断逻辑。

```go
func minimumLines(a [][]int) (ans int) {
	sort.Slice(a, func(i, j int) bool { return a[i][0] < a[j][0] }) // 按照 day 排序
	for i, preDY, preDX := 1, 1, 0; i < len(a); i++ {
		dy, dx := a[i][1]-a[i-1][1], a[i][0]-a[i-1][0]
		if dy*preDX != preDY*dx { // 与上一条线段的斜率不同
			ans++
			preDY, preDX = dy, dx
		}
	}
	return
}
```
