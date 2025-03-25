package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1492C(in io.Reader, out io.Writer) {
	var n, m, ans int
	var s, t []byte
	Fscan(in, &n, &m, &s, &t)

	l := make([]int, m)
	k := 0
	for i, b := range t {
		for s[k] != b {
			k++
		}
		l[i] = k
		k++
	}

	r := make([]int, m)
	k = n - 1
	for i := m - 1; i >= 0; i-- {
		b := t[i]
		for s[k] != b {
			k--
		}
		r[i] = k
		k--
	}

	for i := 1; i < m; i++ {
		ans = max(ans, l[i]-l[i-1], r[i]-r[i-1], r[i]-l[i-1])
	}
	Fprint(out, ans)
}

//func main() { cf1492C(bufio.NewReader(os.Stdin), os.Stdout) }
