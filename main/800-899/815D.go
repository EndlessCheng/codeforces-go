package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf815D(in io.Reader, out io.Writer) {
	var n, a, b, c int
	Fscan(in, &n, &a, &b, &c)
	type pair struct{ x, y int }
	v := make([][]pair, a+1)
	m := make([]int, b+2)
	for range n {
		var i, j, k int
		Fscan(in, &i, &j, &k)
		v[i] = append(v[i], pair{j, k})
		m[j] = max(m[j], k)
	}
	for i := b; i >= 0; i-- {
		m[i] = max(m[i], m[i+1])
	}
	s := make([]int, b+1)
	for i := 1; i <= b; i++ {
		s[i] = s[i-1] + m[i]
	}

	ans := 0
	i, j, k := b, 0, 0
	for x := a; x > 0; x-- {
		for _, p := range v[x] {
			j = max(j, p.x)
			k = max(k, p.y)
		}
		for i = max(i, j); i > j && m[i] < k; i-- {
		}
		ans += (b-i)*(c-k) + (i-j)*c - s[i] + s[j]
	}
	Fprint(out, ans)
}

//func main() { cf815D(bufio.NewReader(os.Stdin), os.Stdout) }
