### 提示 1

每个比特位分别计算方案数，最后答案为每个比特位的方案数的乘积。

现在来思考 $a$ 中只有 $0$ 和 $1$ 的方案数。

### 提示 2

对于每个约束：

- 如果 $x$（在这个比特位上）是 $0$，意味着区间 $[l,r]$ 至少要有一个 $0$。
- 如果 $x$（在这个比特位上）是 $1$，意味着区间 $[l,r]$ 必须都是 $1$。

### 提示 3

考虑哪些位置填入 $0$。

定义 $f[i]$ 表示考虑下标 $[1,i]$，且下标 $i$ 填 $0$ 的方案数。

如果下标 $i$ 必须填 $1$，则 $f[i] = 0$。

否则，枚举上一个 $0$ 的位置 $j$，换言之，从 $j+1$ 到 $i-1$ 全部填 $1$。

由于有「区间 $[l,r]$ 至少要有一个 $0$」的约束，$j$ 不能太小，所以

$$
f[i] = f[\textit{minJ}] + f[\textit{minJ}+1] + \cdots + f[i-1]
$$

特别地，我们规定 $f[0]=1$，表示从 $1$ 到 $i-1$ 全部是 $1$，有一种方案（但不一定可以从 $f[0]$ 转移到 $f[i]$）。

现在需要思考的问题是，对于每个 $i$，其转移来源 $j$ 的最小值 $\textit{minJ}$ 是多少？

### 提示 4

讨论 $f[i]$ 的转移来源的最小值：

- 如果 $i-1$ 恰好是某个约束区间 $[l,r]$ 的右端点，且这个区间至少要有一个 $0$，那么转移来源不能小于 $l$，否则从 $l$ 到 $r$ 就都是 $1$ 了。当然，转移来源也不能小于 $f[i-1]$ 对应的 $\textit{minJ}$，因为左边的其它约束区间也要满足。
- 如果 $i-1$ 不是某个约束区间 $[l,r]$ 的右端点，那么 $f[i]$ 对应的 $\textit{minJ}$ 和 $f[i-1]$ 对应的 $\textit{minJ}$ 是一样的。

我们可以预处理数组 $\textit{maxL}[r]$ 表示至少要有一个 $0$ 的约束区间右端点为 $r$ 时，左端点 $l$ 的最大值。

### 提示 5

如何快速计算 $f[\textit{minJ}] + f[\textit{minJ}+1] + \cdots + f[i-1]$？

注意到 $\textit{minJ}$ 不会减小，所以可以用滑动窗口优化，用一个变量 $\textit{sumF}$ 维护这个和式。

如何快速知道下标 $i$ 必须填 $1$？

用差分数组标记必须全为 $1$ 的区间。

```go
package main
import("bufio";."fmt";"os")
func max(a, b int) int { if b > a { return b }; return a }

func main() {
	in := bufio.NewReader(os.Stdin)
	const mod = 998244353
	var n, k, m int
	Fscan(in, &n, &k, &m)
	cons := make([]struct{ l, r, x int }, m)
	for i := range cons {
		Fscan(in, &cons[i].l, &cons[i].r, &cons[i].x)
	}

	ans := 1
	f := make([]int, n+2)
	f[0] = 1
	for b := 0; b < k; b++ {
		maxL := make([]int, n+1)
		d := make([]int, n+2)
		for _, p := range cons {
			if p.x>>b&1 == 0 {
				maxL[p.r] = max(maxL[p.r], p.l)
			} else {
				d[p.l]++
				d[p.r+1]--
			}
		}

		sumF := 1
		sumD := 0
		left := 0
		for i := 1; i <= n+1; i++ {
			for left < maxL[i-1] {
				sumF -= f[left]
				left++
			}
			sumD += d[i]
			if sumD > 0 {
				f[i] = 0
				continue
			}
			sumF %= mod
			f[i] = sumF
			sumF *= 2
		}
		// f[n+1] 相当于枚举最后一个 0 的下标，计算这些 f[i] 之和
		ans = ans * f[n+1] % mod
	}
	Print((ans%mod + mod) % mod) // 保证答案非负
}
```

时间复杂度：$\mathcal{O}(k\cdot(n+m))$。
