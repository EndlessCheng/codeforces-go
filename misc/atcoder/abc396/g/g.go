package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	a := make([]string, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, 1<<n)
	}
	for j := 0; j < m; j++ {
		v := 0
		for i, s := range a {
			v |= int(s[j]-'0') << i
		}
		f[0][v]++
	}
	for i := 0; i < n; i++ {
		for j := n; j > 0; j-- {
			for k := range f[j] {
				f[j][k] += f[j-1][k^1<<i]
			}
		}
	}

	ans := int(1e9)
	for k := 0; k < 1<<n; k++ {
		s := 0
		for i, row := range f {
			s += min(i, n-i) * row[k]
		}
		ans = min(ans, s)
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
