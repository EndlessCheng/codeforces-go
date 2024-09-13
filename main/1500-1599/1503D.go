package main

import (
	. "fmt"
	"io"
	"slices"
)

func cf1503D(in io.Reader, out io.Writer) {
	var n, x, y, ans int
	Fscan(in, &n)
	to := make([]int, n*2+1)
	pos := make([]byte, n*2+1)
	for i := 0; i < n; i++ {
		Fscan(in, &x, &y)
		pos[y] = 1
		to[x] = y
		to[y] = x
	}

	a := make([]int, n)
	al, ar := 0, n-1
	b := make([]int, n)
	bl, br := 0, n-1
	vis := make([]bool, n*2+2)
	mn, mx := 1, n*2
	for mn < mx {
		cnt := [2]int{}
		t := mn
		for t > 0 {
			last := 0
			for ; mn <= t; mn++ {
				if vis[mn] {
					continue
				}
				cnt[pos[mn]]++

				vis[mn] = true
				a[al] = mn
				al++

				last = to[mn]
				vis[last] = true
				b[bl] = last
				bl++
			}
			if last == 0 {
				break
			}

			t = last
			last = 0
			for ; mx >= t; mx-- {
				if vis[mx] {
					continue
				}
				cnt[pos[mx]]++

				vis[mx] = true
				a[ar] = mx
				ar--

				last = to[mx]
				vis[last] = true
				b[br] = last
				br--
			}
			t = last
		}
		ans += min(cnt[0], cnt[1])
	}

	slices.Reverse(b)
	if !slices.IsSorted(a) || !slices.IsSorted(b) {
		Fprint(out, -1)
	} else {
		Fprint(out, ans)
	}
}

//func main() { cf1503D(bufio.NewReader(os.Stdin), os.Stdout) }
