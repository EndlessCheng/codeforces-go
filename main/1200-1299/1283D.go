package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1283D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, n)
	vis := make(map[int]bool, n+m*2)
	for i := range a {
		Fscan(in, &a[i])
		vis[a[i]] = true
	}
	q := []int{}
	add := func(v int) {
		for i := v - 1; i < v+2; i += 2 {
			if !vis[i] {
				vis[i] = true
				q = append(q, i)
			}
		}
	}
	for _, v := range a {
		add(v)
	}

	ans := int64(0)
	y := make([]int, 0, m)
	for d := int64(1); ; d++ {
		tmp := q
		q = nil
		for _, v := range tmp {
			ans += d
			y = append(y, v)
			if len(y) == m {
				Fprintln(out, ans)
				for _, v := range y {
					Fprint(out, v, " ")
				}
				return
			}
			add(v)
		}
	}
}

//func main() { CF1283D(os.Stdin, os.Stdout) }
