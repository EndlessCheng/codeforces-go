package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p4933(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n int
	Fscan(in, &n)
	ans := n * (n + 1) / 2
	f := make([]map[int]int, n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		x := a[i]
		f[i] = map[int]int{}
		for j, y := range a[:i] {
			d := x - y
			cnt := f[j][d]
			ans += cnt
			f[i][d] = (f[i][d] + cnt + 1) % mod
		}
	}
	Fprint(out, ans%mod)
}

//func main() { p4933(bufio.NewReader(os.Stdin), os.Stdout) }
