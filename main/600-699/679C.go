package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf679C(in io.Reader, out io.Writer) {
	var n, k, ans, ts int
	Fscan(in, &n, &k)
	a := make([]string, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	dir4 := []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	id := make([][]int, n)
	for i := range id {
		id[i] = make([]int, n)
	}
	size := []int{0}
	var dfs func(int, int)
	dfs = func(i, j int) {
		id[i][j] = len(size) - 1
		size[len(size)-1]++
		for _, d := range dir4 {
			x, y := i+d.x, j+d.y
			if 0 <= x && x < n && 0 <= y && y < n && a[x][y] != 'X' && id[x][y] == 0 {
				dfs(x, y)
			}
		}
	}
	for i, row := range a {
		for j, v := range row {
			if v == 'X' || id[i][j] > 0 {
				continue
			}
			size = append(size, 0)
			dfs(i, j)
		}
	}

	vis := make([]int, len(size))
	for top := range n - k + 1 {
		for r := range n {
			for _, row := range id[top : top+k] {
				size[row[r]]--
			}

			l := r + 1 - k
			if l < 0 {
				continue
			}

			s := 0
			ts++
			for i := max(top-1, 0); i <= min(top+k, n-1); i++ {
				j, step, c := l, 1, k
				if i != top-1 && i != top+k {
					j, step, c = l-1, k+1, 2
				}
				for range c {
					if 0 <= j && j < n {
						p := id[i][j]
						if p > 0 && vis[p] != ts {
							vis[p] = ts
							s += size[p]
						}
					}
					j += step
				}
			}
			ans = max(ans, s)

			for _, row := range id[top : top+k] {
				size[row[l]]++
			}
		}

		for l := n - k + 1; l < n; l++ {
			for _, row := range id[top : top+k] {
				size[row[l]]++
			}
		}
	}
	Fprint(out, ans+k*k)
}

//func main() { cf679C(bufio.NewReader(os.Stdin), os.Stdout) }
