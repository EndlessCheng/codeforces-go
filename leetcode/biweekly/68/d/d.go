package main

import (
	"fmt"
	"math"
	"math/bits"
)

/* 拆分成四个问题：计算对数 + 对 1e5 取模 + 计算尾零个数 + 判断剩余数字的长度

我们拆分成四个问题来算：

- 计算 $\textit{pre}$；
- 计算 $\textit{suf}$；
- 计算 $C$；
- 判断剩余数字的长度是否超过 $10$。

#### 1. 计算 $\textit{pre}$

我们可以通过取以 $10$ 为底的对数的方式，将乘法转换成加法，即如下法则：

$$
a\cdot b = 10^{\log_{10}a}\cdot 10^{\log_{10}b} = 10^{\log_{10}a+\log_{10}b}
$$

记最后得到的指数为 $e$，则有

$$
\textit{pre} = 10^{e-\lfloor e \rfloor} \cdot 10000 = 10^{e-\lfloor e \rfloor + 4}
$$

#### 2. 计算 $C$

先来看怎么计算尾零。这相当于求乘积中能分解出来的 $10$ 的个数。

我们可以将所有整数分解质因子，那么分解出来的 $2$ 的幂次之和，以及 $5$ 的幂次之和，这两者的较小值就是最后乘积中能分解出来的 $10$ 的个数。

#### 3. 计算 $\textit{suf}$

我们可以将每个数字的所有因子 $2$ 和 $5$ 去掉，然后将剩下的数字相乘，由于我们只取末 $5$ 位，所以在乘法的过程中可以对 $10^5$ 取模。

由于可能会多去掉一些 $2$ 或 $5$，在遍历 $[\textit{left},\textit{right}]$ 结束后还需要再重新乘上多去掉的 $2$ 或 $5$。

#### 4. 判断剩余数字的长度是否超过 $10$

在上一条的过程中额外计算一个乘积，判断其是否大于或等于 $10^{10}$。

*/

// github.com/EndlessCheng/codeforces-go
func abbreviateProduct(left, right int) string {
	e, cnt2, cnt5, suf, mul := 0.0, 0, 0, 1, 1
	update := func(x int) {
		suf = suf * x % 1e5
		if mul != -1 {
			mul *= x
			if mul >= 1e10 {
				mul = -1
			}
		}
	}

	for i := left; i <= right; i++ {
		e += math.Log10(float64(i))
		x := i
		tz := bits.TrailingZeros(uint(x)) // 因子 2 的个数
		cnt2 += tz
		x >>= tz
		for ; x%5 == 0; x /= 5 {
			cnt5++ // 因子 5 的个数
		}
		update(x)
	}
	cnt10 := min(cnt2, cnt5)
	for i := cnt10; i < cnt2; i++ {
		update(2) // 补上多拆出来的 2
	}
	for i := cnt10; i < cnt5; i++ {
		update(5) // 补上多拆出来的 5
	}

	if mul != -1 { // 不需要缩写
		return fmt.Sprintf("%de%d", mul, cnt10)
	}
	pre := int(math.Pow(10, e-math.Floor(e)+4))
	return fmt.Sprintf("%d...%05de%d", pre, suf, cnt10)
}

func min(a, b int) int { if a > b { return b }; return a }
