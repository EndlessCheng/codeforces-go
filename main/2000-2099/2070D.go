package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2070D(in io.Reader, out io.Writer) {
	const mod = 998244353
	var T, n, p int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n)
		d := make([]int, n)
		row := make([][]int, n)
		for i := 1; i < n; i++ {
			Fscan(in, &p)
			p--
			g[p] = append(g[p], i)
			d[i] = d[p] + 1
			row[d[i]] = append(row[d[i]], i)
		}

		f := make([]int, n)
		pre := 0
		for i := n - 1; i > 0; i-- {
			cur := 0
			for _, v := range row[i] {
				sum := pre
				for _, w := range g[v] {
					sum -= f[w]
				}
				f[v] = (sum + 1) % mod
				cur += f[v]
			}
			pre = cur
		}
		Fprintln(out, ((pre+1)%mod+mod)%mod)
	}
}

//func main() { cf2070D(bufio.NewReader(os.Stdin), os.Stdout) }
