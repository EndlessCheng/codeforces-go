package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1043D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m int
	Fscan(in, &n, &m)
	a := make([][]int, m)
	pos := make([][]int, m)
	for i := range a {
		a[i] = make([]int, n)
		pos[i] = make([]int, n+1)
		for j := range a[i] {
			Fscan(in, &a[i][j])
			pos[i][a[i][j]] = j
		}
	}

	ans, c := int64(0), 0
	ps := make([]int, m)
	ps[0] = -2
	for _, v := range a[0] {
		for i, p := range pos {
			ps[i]++
			if ps[i] != p[v] {
				for i, p := range pos {
					ps[i] = p[v]
				}
				c = 0
				break
			}
		}
		c++
		ans += int64(c)
	}
	Fprint(out, ans)
}

//func main() { CF1043D(os.Stdin, os.Stdout) }
