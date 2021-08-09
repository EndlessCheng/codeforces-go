package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF920G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx int = 1e6 + 1
	ds := [mx][]int{}
	for i := 1; i < mx; i++ {
		for j := i; j < mx; j += i {
			ds[j] = append(ds[j], i)
		}
	}
	mu := [mx]int8{1: 1}
	for i := 1; i < mx; i++ {
		for j := i * 2; j < mx; j += i {
			mu[j] -= mu[i]
		}
	}
	f := func(x, p int) (res int) {
		for _, d := range ds[p] {
			res += x / d * int(mu[d])
		}
		return
	}

	var T, x, p, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x, &p, &k)
		k += f(x, p)
		Fprintln(out, sort.Search(7e6, func(x int) bool { return f(x, p) >= k }))
	}
}

//func main() { CF920G(os.Stdin, os.Stdout) }
