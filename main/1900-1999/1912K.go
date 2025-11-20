package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1912K(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n, v uint
	Fscan(in, &n)
	cnt := [2]int{}
	cnt2 := [4]int{}
	f := [4]int{}
	for range n {
		Fscan(in, &v)
		v %= 2
		f[2|v] = (f[2|v] + f[(v^1)<<1|1] + cnt2[(v^1)<<1|1]) % mod
		f[v] = (f[v] + f[v<<1] + cnt2[v<<1]) % mod
		cnt2[2|v] = (cnt2[2|v] + cnt[1]) % mod
		cnt2[v] = (cnt2[v] + cnt[0]) % mod
		cnt[v]++
	}
	Fprint(out, (f[0]+f[1]+f[2]+f[3])%mod)
}

//func main() { cf1912K(bufio.NewReader(os.Stdin), os.Stdout) }
