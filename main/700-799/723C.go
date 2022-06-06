package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF723C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, need int
	Fscan(in, &n, &m)
	avg := n / m
	a := make([]int, n)
	cnt := make([]int, m+1)
	ex := []int{}
	for i := range a {
		Fscan(in, &a[i])
		if v := a[i]; v > m {
			ex = append(ex, i)
		} else if cnt[v]++; cnt[v] > avg {
			ex = append(ex, i)
		}
	}
	for _, c := range cnt[1:] {
		if c < avg {
			need += avg - c
		}
	}
	Fprintln(out, avg, need)

	for v, i := 1, 0; v <= m; v++ {
		for c := avg - cnt[v]; c > 0; c-- {
			a[ex[i]] = v
			i++
		}
	}
	for _, v := range a {
		Fprint(out, v, " ")
	}
}

//func main() { CF723C(os.Stdin, os.Stdout) }
