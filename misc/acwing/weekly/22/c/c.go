package main

import (
	. "fmt"
	"io"
	"math/big"
	"os"
)

/* 脑筋急转弯
由于 $a_i\le b_i$，结合 $a$ 递增 $b$ 递减的要求，只要满足 $a_m\le b_m$ 就满足 $a_i\le b_i$。

于是我们完全可以把 $b$ 倒过来拼在 $a$ 的后面，这样题目就简化成了求长为 $2m$，元素范围在 $[1,n]$ 的非严格单调递增数组个数。

这是一个经典组合数学问题，等价于把 $2m$ 个无区别的球放入 $n$ 个有区别的盒子中，且允许空盒的方案数，这里第 $i$ 个盒子放的球就表示值为 $i$ 的元素。

我们采用隔板法来解决：把 $n$ 个盒子当做 $n-1$ 个隔板，球加上隔板总共有 $2m+n-1$ 个位置，从中选择 $n-1$ 个位置放隔板，这样就把 $2m$ 个球划分成了 $n$ 份，放入对应的盒子中。

因此方案数为 $C(2m+n-1,n-1)$。
*/

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var n, m int64
	Fscan(in, &n, &m)
	Fprint(out, new(big.Int).Rem(new(big.Int).Binomial(2*m+n-1, n-1), big.NewInt(1e9+7)))
}

func main() { run(os.Stdin, os.Stdout) }
