package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1148C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n+1)
	p := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		p[a[i]] = i
	}

	type pr struct{ x, y int }
	ans := []pr{}
	for i := 1; i <= n; i++ {
		j := p[i]
		if j == i {
			continue
		}
		sz := len(ans)
		if j-i >= n/2 {
			ans = append(ans, pr{i, j})
			p[a[i]] = j
		} else if j <= n/2 {
			ans = append(ans, pr{j, n}, pr{i, n})
			p[a[i]] = n
			p[a[n]] = j
		} else if i <= n/2 {
			ans = append(ans, pr{1, j}, pr{1, n}, pr{i, n}, pr{1, j})
			p[a[i]] = n
			p[a[n]] = j
		} else {
			ans = append(ans, pr{1, j}, pr{1, i}, pr{1, j})
			p[a[i]] = j
		}
		for _, p := range ans[sz:] {
			a[p.x], a[p.y] = a[p.y], a[p.x]
		}
	}
	Fprintln(out, len(ans))
	for _, p := range ans {
		Fprintln(out, p.x, p.y)
	}
}

//func main() { CF1148C(os.Stdin, os.Stdout) }
