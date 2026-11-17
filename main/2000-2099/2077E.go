package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2077E(in io.Reader, out io.Writer) {
	const mod = 998244353
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		st := [2][]int{}
		ans := 0
		for i := range n {
			Fscan(in, &v)
			a[i] = v
			ans = (ans + v*(i+1)%mod*(n-i)) % mod
			p := i & 1
			for v != 0 && len(st[p]) > 0 {
				j := st[p][len(st[p])-1]
				w := min(a[j], v)
				ans = (ans - w*(j+1)%mod*(n-i)) % mod
				a[j] -= w
				v -= w
				if a[j] == 0 {
					st[p] = st[p][:len(st[p])-1]
				}
			}
			st[p^1] = append(st[p^1], i)
		}
		Fprintln(out, (ans+mod)%mod)
	}
}

//func main() { cf2077E(bufio.NewReader(os.Stdin), os.Stdout) }
