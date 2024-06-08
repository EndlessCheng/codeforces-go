package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1879D(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n, v, xor, ans int
	Fscan(in, &n)
	cnt := [30][2]int{}
	for i := range cnt {
		cnt[i][0] = 1
	}
	sum := [30][2]int{}
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		xor ^= v
		for j := 0; j < 30; j++ {
			b := xor >> j & 1
			ans = (ans + (i*cnt[j][b^1]-sum[j][b^1])%mod<<j) % mod
			cnt[j][b]++
			sum[j][b] += i
		}
	}
	Fprint(out, ans)
}

//func main() { cf1879D(bufio.NewReader(os.Stdin), os.Stdout) }
