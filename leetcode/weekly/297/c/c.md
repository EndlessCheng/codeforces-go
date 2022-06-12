定义 $f[i][j]$ 表示前 $i$ 个孩子分配的饼干集合为 $j$ 时，前 $i$ 个孩子的不公平程度的最小值。

考虑给第 $i$ 个孩子分配的饼干集合为 $s$，设集合 $s$ 的元素和为 $\textit{sum}[s]$，那么此时前 $i$ 个孩子的不公平程度为

$$
\max(f[i-1][j \setminus s], \textit{sum}[s])
$$

其中 $j \setminus s$ 表示从集合 $j$ 中去掉集合 $s$ 的元素后剩余的元素组成的集合。

枚举所有 $j$ 的子集 $s$，取上式的最小值即为 $f[i][j]$。

代码实现时，通过倒序枚举 $j$，第一个维度可以省略。此外，$\textit{sum}$ 可以通过预处理得到。

```go
func distributeCookies(a []int, k int) int {
	n := 1 << len(a)
	sum := make([]int, n)
	for i, v := range a {
		for mask, bit := 0, 1<<i; mask < bit; mask++ {
			sum[bit|mask] = sum[mask] + v
		}
	}

	f := append([]int{}, sum...)
	for i := 1; i < k; i++ {
		for j := n - 1; j > 0; j-- {
			for s := j; s > 0; s = (s - 1) & j {
				f[j] = min(f[j], max(f[j^s], sum[s]))
			}
		}
	}
	return f[n-1]
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
```
