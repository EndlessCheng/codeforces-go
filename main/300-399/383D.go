package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF383D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod int = 1e9 + 7
	var n, bias, s, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		bias += a[i]
	}
	f := make([]int, bias*2+1)
	f[bias] = 1
	for _, v := range a {
		ans = (ans + (f[bias+v]+f[bias-v])%mod) % mod
		nf := make([]int, len(f))
		for i := bias - s; i <= bias+s; i++ {
			nf[i-v] = (nf[i-v] + f[i]) % mod
			nf[i+v] = (nf[i+v] + f[i]) % mod
		}
		nf[bias] = (nf[bias] + 1) % mod
		f = nf
		s += v
	}
	Fprint(out, ans)
}

//func main() { CF383D(os.Stdin, os.Stdout) }
