package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF922E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}

	var n, tot, ans int
	var base, incCap, incMana, cost int64
	Fscan(in, &n, &base, &incCap, &incMana)
	c := make([]int, n)
	for i := range c {
		Fscan(in, &c[i])
		tot += c[i]
	}

	f := make([]int64, tot+1)
	f[0] = base
	for _, c := range c {
		Fscan(in, &cost)
		for j := ans; j >= 0; j-- {
			for k := 1; k <= c && int64(k)*cost <= f[j]; k++ {
				f[j+k] = max(f[j+k], f[j]-int64(k)*cost) // 枚举召唤次数
				if j+k > ans {
					ans = j + k
				}
			}
		}
		// 移至下一棵树，回复魔力，但不能超过魔力上限
		for j := 0; j <= ans; j++ {
			f[j] = min(f[j]+incMana, base+int64(j)*incCap)
		}
	}
	Fprint(out, ans)
}

//func main() { CF922E(os.Stdin, os.Stdout) }
