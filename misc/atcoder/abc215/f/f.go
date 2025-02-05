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
	var n int
	Fscan(in, &n)
	a := make([]struct{ x, y int }, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y)
	}
	sort.Slice(a, func(i, j int) bool { return a[i].x < a[j].x })
	ans := sort.Search(a[n-1].x-a[0].x, func(low int) bool {
		low++
		mn, mx := int(1e9), 0
		l := 0
		for _, p := range a {
			for a[l].x <= p.x-low {
				mn = min(mn, a[l].y)
				mx = max(mx, a[l].y)
				l++
			}
			if p.y-mn >= low || mx-p.y >= low {
				return false
			}
		}
		return true
	})
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
