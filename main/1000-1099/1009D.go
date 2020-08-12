package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1009D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var n, m int
	Fscan(in, &n, &m)
	if m < n-1 {
		Fprint(out, "Impossible")
		return
	}
	ans := make([][2]int, 0, m)
o:
	for i := 1; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if gcd(i, j) == 1 {
				ans = append(ans, [2]int{i, j})
				if len(ans) == m {
					break o
				}
			}
		}
	}
	if len(ans) < m {
		Fprint(out, "Impossible")
		return
	}
	Fprintln(out, "Possible")
	for _, e := range ans {
		Fprintln(out, e[0], e[1])
	}
}

//func main() { CF1009D(os.Stdin, os.Stdout) }
