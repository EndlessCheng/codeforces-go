package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf908E(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var m, n int
	var s string
	Fscan(in, &m, &n)
	bell := make([]int, m+1)
	b := make([]int, m+1)
	b[0] = 1
	bell[0] = 1
	for i := 1; i <= m; i++ {
		pre := b[0]
		b[0] = b[i-1]
		bell[i] = b[0]
		for j := 1; j <= i; j++ {
			b[j], pre = (b[j-1]+pre)%mod, b[j]
		}
	}

	a := make([]int, m)
	for j := range n {
		Fscan(in, &s)
		for i, b := range s {
			a[i] |= int(b&1) << j
		}
	}

	cnt := map[int]int{}
	for _, v := range a {
		cnt[v]++
	}

	ans := 1
	for _, c := range cnt {
		ans = ans * bell[c] % mod
	}
	Fprint(out, ans)
}

//func main() { cf908E(bufio.NewReader(os.Stdin), os.Stdout) }
