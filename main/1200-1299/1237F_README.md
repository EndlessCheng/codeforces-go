下文用 $n$ 表示行数，$m$ 表示列数。

## 从特殊到一般

想一想，如果只能竖着放骨牌，有多少种方案？

设初始有 $\textit{emptyC}$ 列是空的（没有骨牌），枚举竖着放**恰好** $i$ 个骨牌，那么问题可以拆分为：

1. 有 $\textit{emptyC}$ 列可以竖着放骨牌，从中选 $i$ 列的**排列数**，即 $P_{\textit{emptyC}}^{i}$。
2. 有 $n$ 行，其中某些行不能放骨牌（因为初始已经放了骨牌），从中选择 $i$ 个**连续两行**（用来竖着放骨牌）的方案数。

由于第一个问题只和列有关，第二个问题只和行有关，所以两个问题互相独立，根据乘法原理，可以分别计算再相乘。

### 第一个问题

为什么要计算的是排列数，而不是组合数？

假如可以：

- 在第 $1$ 列的第 $1,2$ 行竖着放一个骨牌。
- 在第 $2$ 列的第 $3,4$ 行竖着放一个骨牌。
- 在第 $3$ 列的第 $5,6$ 行竖着放一个骨牌。

根据题目约束，这 $3$ 个骨牌也可以这样放：

- 在第 $1$ 列的第 $3,4$ 行竖着放一个骨牌。
- 在第 $2$ 列的第 $5,6$ 行竖着放一个骨牌。
- 在第 $3$ 列的第 $1,2$ 行竖着放一个骨牌。

也就是说，只要我能竖着放 $3$ 个骨牌，那么这 $3$ 个骨牌之间就有 $3!$ 种不同方案。

### 第二个问题

这可以用 DP 解决。

定义 $f_{i,j}$ 表示从前 $i$ 行选恰好 $j$ 个连续两行的方案数。

考虑选或不选，有

$$
f_{i,j} = f_{i-1,j} + f_{i-2,j-1}
$$

其中 $f_{i-2,j-1}$ 只有当第 $i-1$ 行和第 $i$ 行都是空的才能转移（选）。

初始值 $f_{i,0} = 1$。

有 $n$ 行，其中某些行不能放骨牌，从中选择 $i$ 个连续两行（用来竖着放骨牌）的方案数即 $f_{n,i}$。

## 回到原问题

假设我们已经把 $i$ 个骨牌竖着放在棋盘上，那么剩下要解决的问题，就是在此基础上，计算横着放 $j$ 个骨牌的方案数。思考方式同上。

但是，如果横着放 $j$ 个骨牌，就会导致有额外的 $2j$ 列是不能竖着放骨牌的，所以对竖着放骨牌的排列数会有影响。具体分析如下。

设初始有 $\textit{emptyR}$ 行是空的（没有骨牌），有 $\textit{emptyC}$ 列是空的。

枚举竖着放**恰好** $i$ 个骨牌，横着放**恰好** $j$ 个骨牌，那么问题可以拆分为：

1. 有 $\textit{emptyC}-2j$ 列可以竖着放骨牌，从中选 $i$ 列的排列数，即 $P_{\textit{emptyC}-2j}^{i}$。
2. 有 $\textit{emptyR}-2i$ 行可以横着放骨牌，从中选 $j$ 行的排列数，即 $P_{\textit{emptyR}-2i}^{j}$。
3. 有 $n$ 行，其中某些行不能放骨牌，从中选择 $i$ 个连续两行（用来竖着放骨牌）的方案数。
4. 有 $m$ 列，其中某些列不能放骨牌，从中选择 $j$ 个连续两列（用来横着放骨牌）的方案数。

和 $f_{n,i}$ 一样，用同样的方法计算有 $m$ 列，其中若干列不能放骨牌，从中选择 $j$ 个连续两列（用来横着放骨牌）的方案数，记作 $g_{m,j}$。

上述四个问题的方案数相乘，即

$$
f_{n,i}\cdot g_{m,j}\cdot P_{\textit{emptyC}-2j}^{i}\cdot P_{\textit{emptyR}-2i}^{j}
$$

枚举 $i$ 和 $j$，答案为

$$
\sum_{i=0}^{\lfloor n/2 \rfloor} \sum_{j=0}^{\lfloor m/2 \rfloor}f_{n,i}\cdot g_{m,j}\cdot P_{\textit{emptyC}-2j}^{i}\cdot P_{\textit{emptyR}-2i}^{j}
$$

## AC 代码（Golang）

```go
package main
import . "fmt"

func main() {
	const mod = 998244353
	const mx = 3600
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	F := [mx + 1]int{1}
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * i % mod
	}
	invF := [...]int{mx: pow(F[mx], mod-2)}
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
	perm := func(n, k int) int {
		return F[n] * invF[n-k] % mod
	}

	var n, m, k, r1, c1, r2, c2, ans int
	Scan(&n, &m, &k)
	banR := make([]bool, n)
	banC := make([]bool, m)
	for ; k > 0; k-- {
		Scan(&r1, &c1, &r2, &c2)
		banR[r1-1] = true
		banR[r2-1] = true
		banC[c1-1] = true
		banC[c2-1] = true
	}

	calc := func(ban []bool) ([]int, int) {
		n := len(ban)
		f := make([][]int, n+1)
		for i := range f {
			f[i] = make([]int, n/2+1)
			f[i][0] = 1
		}
		for i := 1; i < n; i++ {
			for j := 1; j <= (i+1)/2; j++ {
				f[i+1][j] = f[i][j]
				if !ban[i] && !ban[i-1] {
					f[i+1][j] = (f[i+1][j] + f[i-1][j-1]) % mod
				}
			}
		}
		empty := 0
		for _, b := range ban {
			if !b {
				empty++
			}
		}
		return f[n], empty
	}

	f, emptyR := calc(banR)
	g, emptyC := calc(banC)
	for i, v := range f { // i 个竖放
		for j, w := range g { // j 个横放
			if j > emptyR-i*2 || i > emptyC-j*2 {
				break
			}
			ans = (ans + v*w%mod*perm(emptyC-j*2, i)%mod*perm(emptyR-i*2, j)) % mod
		}
	}
	Print(ans)
}
```

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
