package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF584E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var n, tot int
	Fscan(in, &n)
	a := make([]int, n)
	pa := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		a[i]--
		pa[a[i]] = i
	}
	b := make([]int, n)
	pb := make([]int, n)
	for i := range b {
		Fscan(in, &b[i])
		b[i]--
		pb[b[i]] = i
	}

	for i, pa := range pa {
		tot += abs(pa - pb[i])
	}
	Fprintln(out, tot/2)
	ans := [][2]int{}
	for _, v := range b {
		for i := pa[v]; a[i] != b[i]; {
			j := pb[a[i]]
			for ; pb[a[j]] < i; j++ { // 左边还有要换的，先换它
			}
			ans = append(ans, [2]int{i + 1, j + 1})
			a[i], a[j] = a[j], a[i]
			pa[a[i]], pa[a[j]] = pa[a[j]], pa[a[i]]
			i = j
		}
	}
	Fprintln(out, len(ans))
	for _, p := range ans {
		Fprintln(out, p[0], p[1])
	}
}

//func main() { CF584E(os.Stdin, os.Stdout) }
