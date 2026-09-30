package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2038I(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	Fscan(in, &n, &m)
	s := make([]string, n+1)
	v1 := make([]int, n)
	for i := 1; i <= n; i++ {
		Fscan(in, &s[i])
		s[i] = "$" + s[i] + s[i]
		v1[i-1] = i
	}

	ans := make([]int, m+1)
	v2 := make([]int, 0, n)
	for t := m + m; t >= 1; t-- {
		v2 = v2[:0]
		for _, i := range v1 {
			if s[i][t] == '1' {
				v2 = append(v2, i)
			}
		}
		for _, i := range v1 {
			if s[i][t] == '0' {
				v2 = append(v2, i)
			}
		}
		v1, v2 = v2, v1
		if t <= m {
			ans[t] = v1[0]
		}
	}

	for i := 1; i <= m; i++ {
		Fprint(out, ans[i], " ")
	}
}

//func main() { cf2038I(bufio.NewReader(os.Stdin), os.Stdout) }
