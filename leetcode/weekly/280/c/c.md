我们可以将 $\textit{beans}$ 从小到大排序后，枚举最终非空袋子中魔法豆的数目 $v$，将小于 $v$ 的魔法豆全部清空，大于 $v$ 的魔法豆减少至 $v$。这样所有非空袋子中的魔法豆就均相等了。

考虑计算最多能保留多少个魔法豆。

设数组 $\textit{beans}$ 的长度为 $n$，对于第 $i$ 个袋子，我们可以至多保留

$$
(n-i) \cdot \textit{beans}[i]
$$

个魔法豆。

用 $\sum\textit{beans}[i]$ 减所有保留魔法豆的最大值，即为答案。

```go
func minimumRemoval(beans []int) int64 {
	sort.Ints(beans)
	sum, mx := 0, 0
	for i, v := range beans {
		sum += v
		mx = max(mx, (len(beans)-i)*v)
	}
	return int64(sum - mx)
}

func max(a, b int) int { if b > a { return b }; return a }
```

附 Python 一行写法：

```Python 
class Solution:
    def minimumRemoval(self, beans: List[int]) -> int:
        return sum(beans) - max((len(beans) - i) * v for i, v in enumerate(sorted(beans)))
```
