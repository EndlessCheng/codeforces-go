package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var T, n, m, lim int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &lim)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		f := func(A []int) bool {
			a := append([]int(nil), A...)
			sort.Ints(a)
			s := 0
			for i, j := 0, len(a)-1; i < m && i < j; i++ {
				s += (a[i] - a[j]) * (a[i] - a[j])
				j--
			}
			return s <= lim
		}
		ans := 0
		for r := 0; r < n; { // 注意这里是 <
			l := r
			for sz := 1; sz > 0; {
				if r+sz <= n && f(a[l:r+sz]) {
					r += sz
					sz *= 2
				} else {
					sz /= 2
				}
			}
			ans++
		}
		Fprintln(out, ans)
	}
}

func main() { run(os.Stdin, os.Stdout) }
