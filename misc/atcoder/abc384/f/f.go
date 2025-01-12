package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
	"slices"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		ans += a[i] / (a[i] & -a[i])
	}
	mx := bits.Len(uint(slices.Max(a)))
	for m := 1; m <= 1<<mx; m <<= 1 {
		type pair struct{ c, s int }
		cnt := map[int]pair{}
		s := 0
		for _, v := range a {
			p := cnt[(m-v%m)%m]
			s += p.c*v + p.s
			p = cnt[v%m]
			cnt[v%m] = pair{p.c + 1, p.s + v}
		}
		if m == 1 {
			ans += s
		} else {
			ans -= s / m
		}
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
