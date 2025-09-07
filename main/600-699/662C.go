package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf662C(in io.Reader, out io.Writer) {
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
	for j := range m {
		v := 0
		for i, s := range a {
			v |= int(s[j]-'0') << i
		}
		f[0][v]++
	}
	for i := range n {
		for j := n; j > 0; j-- {
			for k := range f[j] {
				f[j][k] += f[j-1][k^1<<i]
			}
		}
	}

	ans := int(1e9)
	for k := range 1 << n {
		s := 0
		for i, row := range f {
			s += min(i, n-i) * row[k]
		}
		ans = min(ans, s)
	}
	Fprint(out, ans)
}

//func main() { cf662C(bufio.NewReader(os.Stdin), os.Stdout) }
