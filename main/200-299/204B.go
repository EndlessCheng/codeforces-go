package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf204B(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	type pair struct{ x, y int }
	a := make([]pair, n)
	cnt := map[int]int{}
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y)
		cnt[a[i].x]++
		cnt[a[i].y]++
	}

	ans := n + 1
	m := (n + 1) / 2
	for v, c := range cnt {
		if c < m {
			continue
		}
		front, back := 0, 0
		for _, p := range a {
			if p.x == v {
				front++
			} else if p.y == v {
				back++
			}
		}
		need := m - front
		if need <= 0 {
			Fprint(out, 0)
			return
		}
		if need <= back {
			ans = min(ans, need)
		}
	}
	if ans > n {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { cf204B(bufio.NewReader(os.Stdin), os.Stdout) }
