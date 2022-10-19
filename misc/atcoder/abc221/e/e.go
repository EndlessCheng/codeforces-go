package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 998244353
	const mod2 = (mod + 1) / 2

	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := append([]int(nil), a...)
	sort.Ints(b)

	tree := make([]int, n+1)
	p2, invP2 := 1, 1
	for _, v := range a {
		v = sort.SearchInts(b, v) + 1
		for j := v; j > 0; j &= j - 1 {
			ans = (ans + tree[j]*p2) % mod
		}
		p2 = p2 * 2 % mod
		invP2 = invP2 * mod2 % mod
		for ; v <= n; v += v & -v {
			tree[v] = (tree[v] + invP2) % mod
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
