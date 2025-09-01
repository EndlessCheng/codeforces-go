package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1093F(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n, k, l, v int
	Fscan(in, &n, &k, &l)
	f := make([][]int, n+1)
	f[0] = make([]int, k+1)
	sumF := make([]int, n+1)
	sumF[0] = 1
	cnt := make([]int, k+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		f[i] = make([]int, k+1)
		for x := 1; x <= k; x++ {
			if v > 0 && x != v {
				cnt[x] = 0
				continue
			}
			cnt[x]++
			if cnt[x] >= l {
				f[i][x] = (sumF[i-1] - sumF[i-l] + f[i-l][x]) % mod
			} else {
				f[i][x] = sumF[i-1]
			}
			sumF[i] += f[i][x]
		}
		sumF[i] %= mod
	}
	Fprint(out, (sumF[n]+mod)%mod)
}

//func main() { cf1093F(bufio.NewReader(os.Stdin), os.Stdout) }
