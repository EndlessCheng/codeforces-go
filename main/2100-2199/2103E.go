package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2103E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
		}
		if slices.IsSorted(a[1:]) {
			Fprintln(out, 0)
			continue
		}

		mp := map[int]int{}
		x, y := -1, -1
		for i := 1; i <= n; i++ {
			if j, ok := mp[k-a[i]]; ok {
				x, y = j, i
				break
			}
			mp[a[i]] = i
		}

		if x == -1 && y == -1 {
			Fprintln(out, -1)
			continue
		}

		type tuple struct{ i, j, x int }
		ans := []tuple{}
		if x != 1 {
			ans = append(ans, tuple{y, x, a[y] - a[1]})
			a[y] = a[1]
			a[x] = k - a[y]
			ans = append(ans, tuple{x, 1, a[x] - a[1]})
			a[x], a[1] = a[1], a[x]
			x = 1
		}

		if y != n {
			ans = append(ans, tuple{x, y, a[x] - a[n]})
			a[x] = a[n]
			a[y] = k - a[x]
			ans = append(ans, tuple{y, n, a[y] - a[n]})
			a[y], a[n] = a[n], a[y]
			y = n
		}

		b := make([]int, n+1)
		for i := 2; i < n; i++ {
			b[i] = i
		}
		slices.SortStableFunc(b[2:n], func(x, y int) int { return a[x] - a[y] })
		pos := make([]int, n+1)
		for i := 2; i < n; i++ {
			pos[b[i]] = i
		}

		for i := 2; i < n; i++ {
			if i == b[i] {
				continue
			}
			j := b[i]
			x, y = a[i], a[j]
			ans = append(ans, tuple{1, n, y + a[1] - k}, tuple{1, j, x - y}, tuple{1, i, y - x})
			a[i], a[j] = a[j], a[i]
			b[pos[i]], b[pos[j]] = b[pos[j]], b[pos[i]]
			pos[i], pos[j] = pos[j], pos[i]
			a[1] = k - y
			a[n] = y
		}

		ans = append(ans, tuple{1, n, a[1]})
		Fprintln(out, len(ans))
		for _, d := range ans {
			Fprintln(out, d.i, d.j, d.x)
		}
	}
}

//func main() { cf2103E(bufio.NewReader(os.Stdin), os.Stdout) }
