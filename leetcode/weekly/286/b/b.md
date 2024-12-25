遍历数组，用栈来模拟这个过程（实际不需要栈，后面会说明）：

- 如果栈大小为偶数，可以随意加入元素；
- 如果栈大小为奇数，那么加入的元素不能和栈顶相同。

遍历结束后，若栈大小为奇数，则移除栈顶。

最后栈大小就是保留的元素，用数组大小减去栈大小就是删除的元素个数。

实际上不需要栈，用一个变量表示栈的奇偶性即可。

```go [sol1-Go]
func minDeletion(a []int) (ans int) {
	odd := false // 栈大小的奇偶性
	for i, n := 0, len(a); i < n; {
		start := i
		// 注意这里的 i 和外层循环的 i 是同一个变量，因此时间复杂度为 O(n)
		for i < n && a[i] == a[start] { i++ }
		l := i - start // 连续相同元素个数
		if !odd { // 只能放一个元素
			ans += l - 1
			odd = true
		} else if l == 1 {
			odd = false
		} else { // 放两个元素，栈
			ans += l - 2
		}
	}
	if odd { // 栈大小必须为偶数
		ans++
	}
	return
}
```

