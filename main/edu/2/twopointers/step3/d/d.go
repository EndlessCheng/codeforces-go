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
	var n int
	a := [4][]int{}
	for i := range a {
		Fscan(in, &n)
		a[i] = make([]int, n)
		for j := range a[i] {
			Fscan(in, &a[i][j])
		}
		sort.Ints(a[i])
	}

	mi, ans := int(1e9), [4]int{}
o:
	for i, b := range a {
		for _, v := range b {
			mx, res := v, [4]int{}
			for j, c := range a {
				if j != i {
					p := sort.SearchInts(c, v)
					if p == len(c) {
						continue o
					}
					res[j] = c[p]
					if res[j] > mx {
						mx = res[j]
					}
				}
			}
			if mx-v < mi {
				res[i] = v
				mi, ans = mx-v, res
			}
		}
	}
	Fprint(out, ans[0], ans[1], ans[2], ans[3])
}

func main() { run(os.Stdin, os.Stdout) }
