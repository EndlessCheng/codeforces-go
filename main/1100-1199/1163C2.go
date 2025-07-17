package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1163C2(in io.Reader, out io.Writer) {
	var n, x, y, ans int
	Fscan(in, &n)
	a := make([]struct{ x, y int }, n)
	cnt := map[float64]int{}
	type pair struct{ k, b float64 }
	vis := map[pair]bool{}
	for i := range a {
		Fscan(in, &x, &y)
		a[i].x, a[i].y = x, y
		for _, p := range a[:i] {
			dy := y - p.y
			dx := x - p.x
			k := 1e9
			kb := pair{k, float64(x)}
			if dx != 0 {
				k = float64(dy) / float64(dx)
				kb = pair{k, float64(y*dx-dy*x) / float64(dx)}
			}
			if vis[kb] {
				continue
			}
			vis[kb] = true
			cnt[k]++
			ans += len(vis) - cnt[k]
		}
	}
	Fprint(out, ans)
}

//func main() { cf1163C2(bufio.NewReader(os.Stdin), os.Stdout) }
