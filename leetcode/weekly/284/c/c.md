设 $\textit{nums}$ 的长度为 $n$，根据题意：

- 如果 $n=1$，那么我们只能在栈不为空时删除栈顶，栈为空时将 $\textit{nums}[0]$ 入栈。因此 $k$ 为奇数时，$k$ 次操作后栈为空，返回 $-1$；$k$ 为偶数时则返回 $\textit{nums}[0]$。
- 如果 $k=0$，无法执行任何操作，直接返回 $\textit{nums}[0]$。

其余情况，按数组元素的下标 $i$ 分类讨论：

- 如果 $i=0$，我们可以不断地删除-添加 $\textit{nums}[0]$，如果 $k$ 为偶数，那么最后栈顶为 $\textit{nums}[0]$；如果 $k$ 为奇数（这里要求 $k>1$），我们可以在倒数第二步删除 $\textit{nums}[1]$，最后一步将 $\textit{nums}[0]$ 入栈，从而保证 $\textit{nums}[0]$ 可以为栈顶。 
- 如果 $0<i<k-1$，我们仍然可以仿造上述流程操作。
- 如果 $i=k-1$，最后一步操作只能删除 $\textit{nums}[i]$，所以无法将 $\textit{nums}[i]$ 置于栈顶。
- 如果 $i=k$，那么可以删除前 $k$ 个元素，将 $\textit{nums}[i]$ 置于栈顶。
- 如果 $i>k$，$\textit{nums}[i]$ 前面的元素无法删除，所以无法将 $\textit{nums}[i]$ 置于栈顶。

综上所述，我们可以让 $i<k-1$ 或 $i=k$ 的数组元素作为 $k$ 次操作后的栈顶。这些元素的最大值即为答案。

```Python [sol1-Python3]
class Solution:
    def maximumTop(self, nums: List[int], k: int) -> int:
        return max(num for i, num in enumerate(nums) if i < k - 1 or i == k) if len(nums) > 1 or k % 2 == 0 else -1
```

```go [sol1-Go]
func maximumTop(a []int, k int) (ans int) {
	n := len(a)
	if n == 1 || k == 0 {
		if k%2 == 1 { return -1 }
		return a[0]
	}
	// 删除 a[k-1] 以及 a[k+1:]，下面直接取 a 的最大值
	if k < n {
		a = append(a[:k-1], a[k]) 
	} else if k == n {
		a = a[:n-1]
	}
	for _, v := range a {
		if v > ans { ans = v }
	}
	return
}
```
