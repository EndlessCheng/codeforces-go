package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1735E(in io.Reader, _w io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		slices.Sort(a)
		b := make([]int, n)
		for i := range b {
			Fscan(in, &b[i])
		}

		f := func(k int) bool {
			cnt := map[int]int{}
			for _, v := range b {
				cnt[v]++
			}
			tp := make([]int8, n)
			for i := n - 1; i >= 0; i-- {
				x := a[i] + k
				y := abs(a[i] - k)
				if c, _ := cnt[x]; c > 0 {
					cnt[x]--
					tp[i] = 1
					continue
				}
				if c, _ := cnt[y]; c > 0 {
					cnt[y]--
					tp[i] = 2
					continue
				}
				return false
			}

			Fprintln(out, "YES")
			p := int(1e9)
			if k > 1e9 {
				p = 2e9 - k
			}
			for i, t := range tp {
				if t == 1 {
					Fprint(out, p-a[i], " ")
				} else {
					Fprint(out, p+a[i], " ")
				}
			}
			Fprintln(out)
			Fprintln(out, p, p+k)
			return true
		}

		for _, v := range b {
			if f(a[0]+v) || f(abs(a[0]-v)) {
				continue o
			}
		}
		Fprintln(out, "NO")
	}
}

//func main() { cf1735E(bufio.NewReader(os.Stdin), os.Stdout) }
