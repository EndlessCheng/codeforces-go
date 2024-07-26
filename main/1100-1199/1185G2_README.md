## 初步分析

原问题可以拆分成如下两个问题：

1. 不考虑总时长（以及歌曲的排列），计算从这 $n$ 首歌中选出 $i$ 个类型 $1$，$j$ 个类型 $2$，$k$ 个类型 $3$，且相邻歌曲类型不同的方案数。
2. 不考虑相邻歌曲类型不同（以及歌曲的排列），计算从这 $n$ 首歌中选出 $i$ 个类型 $1$，$j$ 个类型 $2$，$k$ 个类型 $3$，且总时长为 $t$ 的方案数。

两个方案数相乘，再乘上排列数 $i!j!k!$，就可以得到从这 $n$ 首歌中选出 $i$ 个类型 $1$，$j$ 个类型 $2$，$k$ 个类型 $3$，总时长为 $t$ 且相邻歌曲类型不同的方案数。

> 另一种理解方式：相当于有若干个格子和三种颜色，问题 1 负责涂色（保证相邻格子颜色不同），问题 2（乘上 $i!j!k!$）负责给每个格子放上具体的歌曲（颜色要匹配）。

枚举所有的 $i,j,k$，累加方案数，即为答案。

## 问题 1

不考虑总时长（以及歌曲的排列），计算从这 $n$ 首歌中选出 $i$ 个类型 $1$，$j$ 个类型 $2$，$k$ 个类型 $3$，且相邻歌曲类型不同的方案数。

用**状态机 DP** 解决。

定义 $c_{i,j,k,x}$ 表示从这 $n$ 首歌中选出 $i$ 个类型 $1$，$j$ 个类型 $2$，$k$ 个类型 $3$，最后一首歌的类型为 $x$ 且相邻歌曲类型不同的方案数。

枚举上一首歌的类型，那么有

$$
\begin{aligned}
c_{i+1,j,k,1} &= c_{i,j,k,2} + c_{i,j,k,3} \\
c_{i,j+1,k,2} &= c_{i,j,k,1} + c_{i,j,k,3} \\
c_{i,j,k+1,3} &= c_{i,j,k,1} + c_{i,j,k,2} \\
\end{aligned}
$$

这里用刷表法转移。

初始值 $c_{1,0,0,1} = c_{0,1,0,2} = c_{0,0,1,3} = 1$。

## 问题 2

不考虑相邻歌曲类型不同（以及歌曲的排列），计算从这 $n$ 首歌中选出 $i$ 个类型 $1$，$j$ 个类型 $2$，$k$ 个类型 $3$，且总时长为 $t$ 的方案数。

用多维 **0-1 背包**解决。

### 优化前

定义 $f_{p,i,j,k,t}$ 表示从前 $p$ 首歌中选出 $i$ 个类型 $1$，$j$ 个类型 $2$，$k$ 个类型 $3$，且总时长为 $t$ 的方案数。这需要 $\mathcal{O}(n^4T)$ 的时间和空间，太慢了。

### 优化

考虑继续拆分问题：

- 计算从这 $n$ 首歌中选出 $i$ 个类型 $1$，且总时长为 $t$ 的方案数。记作 $f_{i,t}$。
- 计算从这 $n$ 首歌中选出 $j$ 个类型 $2$，$k$ 个类型 $3$，且总时长为 $T-t$ 的方案数。记作 $g_{j,k,T-t}$。

两个方案数相乘，就是问题 2 的方案数。

这两个更小的问题，可以分别用 $\mathcal{O}(n^2T)$ 和 $\mathcal{O}(n^3T)$ 的多维 0-1 背包解决。

## 汇总

枚举 $i,j,k$。

从这 $n$ 首歌中选出 $i$ 个类型 $1$，$j$ 个类型 $2$，$k$ 个类型 $3$，总时长为 $t$ 且相邻歌曲类型不同的方案数为

$$
w_{i,j,k} = i!j!k! \cdot (c_{i,j,k,1} + c_{i,j,k,2} + c_{i,j,k,3}) \cdot \sum_{t=0}^T f_{i,t}g_{j,k,T-t}
$$

答案为

$$
\sum_{i=0}^{\textit{cnt}_1}\sum_{j=0}^{\textit{cnt}_2}\sum_{k=0}^{\textit{cnt}_3} w_{i,j,k}
$$

其中 $\textit{cnt}_x$ 为输入的 $n$ 首歌中的类型为 $x$ 的歌曲个数。

代码实现时，类型改成 $0,1,2$。

先计算问题 2，然后计算问题 1。计算问题 2 和计算答案可以在同一个枚举 $i,j,k$ 的三重循环中解决。

```go
package main
import ."fmt"

const mod = 1_000_000_007

func add(a *int, b int) {
	*a = (*a + b) % mod
}

func main() {
	var n, tot, w, tp int
	Scan(&n, &tot)

	f := make([][]int, n+2)
	for i := range f {
		f[i] = make([]int, tot+1)
	}
	f[0][0] = 1
	g := make([][][]int, n+2)
	for i := range g {
		g[i] = make([][]int, n+2)
		for j := range g[i] {
			g[i][j] = make([]int, tot+1)
		}
	}
	g[0][0][0] = 1

	cnt := [3]int{}
	for i := 0; i < n; i++ {
		Scan(&w, &tp)
		tp--
		if tp == 0 {
			for j := cnt[0]; j >= 0; j-- {
				for t := tot; t >= w; t-- {
					add(&f[j+1][t], f[j][t-w])
				}
			}
		} else {
			is := [3]int{}
			is[tp] = 1
			for j := cnt[1]; j >= 0; j-- {
				for k := cnt[2]; k >= 0; k-- {
					for t := tot; t >= w; t-- {
						add(&g[j+is[1]][k+is[2]][t], g[j][k][t-w])
					}
				}
			}
		}
		cnt[tp]++
	}

	fac := make([]int, n+1)
	fac[0] = 1
	for i := 1; i <= n; i++ {
		fac[i] = fac[i-1] * i % mod
	}

	ans := 0
	c := make([][][][3]int, cnt[0]+2)
	for i := range c {
		c[i] = make([][][3]int, cnt[1]+2)
		for j := range c[i] {
			c[i][j] = make([][3]int, cnt[2]+2)
		}
	}
	c[1][0][0][0] = 1
	c[0][1][0][1] = 1
	c[0][0][1][2] = 1
	for i, mat := range c[:cnt[0]+1] {
		for j, row := range mat[:cnt[1]+1] {
			for k, comb := range row[:cnt[2]+1] {
				sum := 0
				for t, fit := range f[i] {
					sum = (sum + fit*g[j][k][tot-t]) % mod
				}
				add(&ans, fac[i]*fac[j]%mod*fac[k]%mod*(comb[0]+comb[1]+comb[2])%mod*sum)

				add(&c[i+1][j][k][0], comb[1]+comb[2])
				add(&c[i][j+1][k][1], comb[0]+comb[2])
				add(&c[i][j][k+1][2], comb[0]+comb[1])
			}
		}
	}
	Print(ans)
}
```

**时间复杂度**：$\mathcal{O}(n^3T)$。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
