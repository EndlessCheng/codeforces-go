package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF52B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	a := make([]string, n)
	r := make([]int, n)
	c := make([]int, m)
	for i := range a {
		Fscan(in, &a[i])
		r[i] = strings.Count(a[i], "*")
		for j, v := range a[i] {
			if v == '*' {
				c[j]++
			}
		}
	}
	ans := int64(0)
	for i, row := range a {
		for j, v := range row {
			if v == '*' {
				ans += int64((r[i] - 1) * (c[j] - 1))
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF52B(os.Stdin, os.Stdout) }
