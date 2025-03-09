package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	const mx int = 1e6 + 1
	var cnt, ans [mx]int
	for i := range a {
		Fscan(in, &a[i])
		cnt[a[i]]++
	}

	for i := 1; i < mx; i++ {
		s := 0
		for j := i; j < mx; j += i {
			s += cnt[j]
		}
		if s >= k {
			for j := i; j < mx; j += i {
				ans[j] = i
			}
		}
	}
	for _, v := range a {
		Fprintln(out, ans[v])
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
