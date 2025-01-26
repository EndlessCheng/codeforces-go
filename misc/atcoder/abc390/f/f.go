package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	pos := make([][]int, n+1)
	for i := range a {
		Fscan(in, &a[i])
		pos[a[i]] = append(pos[a[i]], i)
	}
	for v, ps := range pos {
		for idx, i := range ps {
			r := n
			if idx+1 < len(ps) {
				r = ps[idx+1]
			}
			ps1 := pos[v-1]
			j := sort.SearchInts(ps1, i)
			if j < len(ps1) {
				r = min(r, ps1[j])
			}
			l := -1
			if j > 0 {
				l = ps1[j-1]
			}
			ans += (i - l) * (r - i)
		}
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
