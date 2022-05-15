位运算题目经典技巧：逐个考虑每一个比特位。

假设元素值只有 $0$ 和 $1$，那么解法就很简单了：由于不能选 $0$（这会导致按位与为 $0$），我们选择所有的 $1$，答案即为 $1$ 的个数。

将上述结论推广，考虑每一个比特位，统计这一位上的 $1$ 的个数，取所有个数的最大值作为答案。

由于 $2^{23} < 10^7<2^{24}$，枚举比特位可以从 $0$ 枚举到 $23$。

```Python [sol1-Python3]
class Solution:
    def largestCombination(self, candidates: List[int]) -> int:
        return max(sum((num >> i) & 1 for num in candidates) for i in range(24))
```

```go [sol1-Go]
func largestCombination(candidates []int) (ans int) {
	for i := 0; i < 24; i++ {
		s := 0
		for _, v := range candidates {
			s += v >> i & 1
		}
		ans = max(ans, s)
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

