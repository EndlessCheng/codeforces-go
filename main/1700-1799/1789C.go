package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1789C(in io.Reader, out io.Writer) {
	var T, n, m, p, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]int, n)
		sumLen := make([]int, n+m+1)
		for i := range a {
			Fscan(in, &a[i])
			sumLen[a[i]] = m + 1
		}
		for i := 1; i <= m; i++ {
			Fscan(in, &p, &v)
			p--
			sumLen[a[p]] -= m + 1 - i
			a[p] = v
			sumLen[v] += m + 1 - i
		}
		ans := n * m * (m + 1)
		for _, k := range sumLen {
			ans -= k * (k - 1) / 2
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1789C(bufio.NewReader(os.Stdin), os.Stdout) }
