package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf360C(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n, k int
	var s string
	Fscan(in, &n, &k, &s)
	// 要求位置 i 有「决定性差异」，即在第 i 位 t[i]≠s[i]
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, k+1)
	}
	f[0][0] = 1
	sum := make([]int, k+1)
	sum[0] = 1
	for i, b := range s {
		i++
		for j := range f[i] {
			f[i][j] = sum[j] * int(b-'a')
			for sz := 1; sz <= min(i, j/(n-i+1)); sz++ {
				f[i][j] += f[i-sz][j-sz*(n-i+1)] * int('z'-b) // 注意这里的 n-i+1 把右边的情况都计算进来了
			}
			f[i][j] %= mod
			sum[j] += f[i][j]
		}
	}
	Fprint(out, sum[k]%mod)
}

//func main() { cf360C(os.Stdin, os.Stdout) }
