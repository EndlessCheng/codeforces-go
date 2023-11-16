从特殊到一般。

### 一

如果 $a$ 中只有 $0$ 和 $1$，我们只需要知道区间内是否有两个 $0$。

如果有两个 $0$，那么最小 OR 是 $0$，否则是 $1$。

相当于只需要知道最小的两个数是多少，就能确定 OR 最小是多少。

### 二

如果 $a$ 中只有 $0,1,2,3$，在最坏情况下，最少需要知道几个数呢？

考虑下面这三个二进制数，其中 $?$ 表示 $0$ 或者 $1$。

$$
\begin{aligned}
0?\\
1?\\
1?
\end{aligned}
$$

如果只选择 $0?$ 和 $1?$，这在 $00$ 和 $10$ 的情况下是没问题的，但当它们分别是

$$
\begin{aligned}
01\\
10\\
10
\end{aligned}
$$

只选择 $01$ 和 $10$ 会算出 OR 为 $11$，但是选择 $10$ 和 $10$ 会算出 OR 为 $10$。

**猜想**：在 $a$ 中只有 $0,1,2,3$ 的情况下，至少要知道最小的 $3$ 个数，才能保证**一定**可以得到 OR 的最小值。

**证明**：分类讨论。

1. 如果有两个 $0?$，那么 OR 的最高位肯定是 $0$，问题变成一个比特，也就是 $a$ 中只有 $0$ 和 $1$ 的情况。我们已经知道，这只需要知道最小的 $2$ 个数。
2. 如果没有 $0?$ 只有 $1?$，那么 OR 的最高位肯定是 $1$，所以同上，变成一个比特的问题，也只需要知道最小的 $2$ 个数。
3. 如果恰好有一个 $0?$，其余的是 $1?$，继续分类讨论：
    - 如果 $0?$ 和 $1?$ 的 OR 最小，那么 $1?$ 这边只需要选最小的数。
    - 如果 $1?$ 和 $1?$ 的 OR 最小，问题变成上面讨论的第 2 点，需要知道 $1?$ 中最小的 $2$ 个数。
    - 所以知道最小的 $3$ 个数就行，OR 的最小值一定是这 $3$ 个数中的 $2$ 个数的 OR。

### 三

如果 $a[i]$ 的范围是 $[0,7]$，至少要知道最小的几个数，才能保证**一定**可以得到 OR 的最小值？

至少要知道最小的 $4$ 个数，证明方式同上。

按照如下方式构造，可以使 OR 的最小值一定来自第三小和第四小的数的 OR。

$$
\begin{aligned}
011\\
101\\
110\\
110
\end{aligned}
$$

如果 $a[i]$ 的范围是 $[0,15]$，构造方法如下：

$$
\begin{aligned}
0111\\
1011\\
1101\\
1110\\
1110
\end{aligned}
$$

### 四

总的来说，通过**数学归纳法**可以证明，OR 的最小值一定是最小的 $31$ 个数中选 $2$ 个数的 OR。

所以用**线段树**维护区间内最小的 $31$ 个数，问题就变成 $C(31,2)$ 的暴力枚举了。

```go
package main
import ("bufio";."fmt";"os")
func min(a, b int) int { if b < a { return b }; return a }

type seg [][]int

// 合并两个有序数组，保留前 k 个数
func merge(a, b []int) []int {
	const k = 31
	i, n := 0, len(a)
	j, m := 0, len(b)
	res := make([]int, 0, min(n+m, k))
	for len(res) < k {
		if i == n {
			res = append(res, b[j:min(j+k-len(res), m)]...)
			break
		}
		if j == m {
			res = append(res, a[i:min(i+k-len(res), n)]...)
			break
		}
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	return res
}

func (t seg) build(a []int, o, l, r int) {
	if l == r {
		t[o] = a[l-1 : l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t[o] = merge(t[o<<1], t[o<<1|1])
}

func (t seg) query(o, l, r, L, R int) []int {
	if L <= l && r <= R {
		return t[o]
	}
	m := (l + r) >> 1
	if R <= m {
		return t.query(o<<1, l, m, L, R)
	}
	if m < L {
		return t.query(o<<1|1, m+1, r, L, R)
	}
	return merge(t.query(o<<1, l, m, L, R), t.query(o<<1|1, m+1, r, L, R))
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var T, n, q, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		t := make(seg, n*4)
		t.build(a, 1, 1, n)
		for Fscan(in, &q); q > 0; q-- {
			Fscan(in, &l, &r)
			b := t.query(1, 1, n, l, r)
			ans := 1 << 30
			for i, v := range b {
				for _, w := range b[:i] {
					ans = min(ans, v|w)
				}
			}
			Fprintln(out, ans)
		}
	}
}
```

**时间复杂度**：$\mathcal{O}(nk + q(k\log n+k^2))$，其中 $k=31$。

**空间复杂度**：$\mathcal{O}(nk)$。
