package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1900D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mx int = 1e5 + 1
	phi := [mx]int{}
	ds := [mx][]int{}
	for i := 1; i < mx; i++ {
		phi[i] = i
		for j := i; j < mx; j += i {
			ds[j] = append(ds[j], i)
		}
	}
	for i := 2; i < mx; i++ {
		if phi[i] == i {
			for j := i; j < mx; j += i {
				phi[j] = phi[j] / i * (i - 1)
			}
		}
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		sort.Ints(a)

		ans := 0
		cnt := [mx]int{}
		for i, x := range a {
			for _, d := range ds[x] {
				ans += (n - 1 - i) * phi[d] * cnt[d]
				cnt[d]++
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1900D(os.Stdin, os.Stdout) }
