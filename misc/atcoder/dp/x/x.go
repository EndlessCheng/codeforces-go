package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	a := make([]struct{ w, s, v int }, n)
	for i := range a {
		Fscan(in, &a[i].w, &a[i].s, &a[i].v)
	}
	sort.Slice(a, func(i, j int) bool { return a[i].w+a[i].s < a[j].w+a[j].s })

	const mx int = 1e4
	f := [mx + 1]int{}
	for _, t := range a {
		ans = max(ans, f[t.s]+t.v)
		for j := mx; j >= t.w; j-- {
			f[j] = max(f[j], f[min(j-t.w, t.s)]+t.v)
		}
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
