package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1158D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	x := make([]int, n+1)
	y := make([]int, n+1)
	vis := make([]bool, n+1)
	s := make([]byte, n)
	for i := 1; i <= n; i++ {
		Fscan(in, &x[i], &y[i])
	}
	var str string
	Fscan(in, &str)
	copy(s[2:], str)

	le := func(a, b, c int) bool {
		return int64(x[b]-x[a])*int64(y[c]-y[a]) > int64(x[c]-x[a])*int64(y[b]-y[a])
	}

	j := 0
	for i := 1; i <= n; i++ {
		if j == 0 || x[i] < x[j] {
			j = i
		}
	}

	vis[j] = true
	Fprint(out, j, " ")
	for i := 2; i < n; i++ {
		l := 0
		if s[i] == 'L' {
			for k := 1; k <= n; k++ {
				if !vis[k] && (l == 0 || !le(j, l, k)) {
					l = k
				}
			}
		} else {
			for k := 1; k <= n; k++ {
				if !vis[k] && (l == 0 || le(j, l, k)) {
					l = k
				}
			}
		}
		j = l
		vis[j] = true
		Fprint(out, j, " ")
	}
	for i := 1; i <= n; i++ {
		if !vis[i] {
			Fprint(out, i)
		}
	}
}

//func main() { cf1158D(bufio.NewReader(os.Stdin), os.Stdout) }
