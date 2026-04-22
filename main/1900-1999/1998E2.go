package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1998E2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]int, n+1)
		p := make([]int, n+1)
		l := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			p[i] = p[i-1] + a[i]
			l[i] = i
		}

		var dfs func(int, int) int
		dfs = func(x, i int) int {
			x = l[x]
			if x == 1 || p[i]-p[x-1] < a[x-1] {
				return x
			}
			l[x] = dfs(x-1, i)
			return l[x]
		}

		cnt := map[int]int{}
		for i := 1; i <= n; i++ {
			cnt[i]++
			nc := map[int]int{}
			for key, w := range cnt {
				if key == i || a[i] <= p[i-1]-p[key-1] {
					nx := dfs(key, i)
					nc[nx] += w
				}
			}
			cnt = nc
			if v, ok := cnt[1]; ok {
				Fprint(out, v, " ")
			}
		}
		Fprintln(out)
	}
}

//func main() { cf1998E2(bufio.NewReader(os.Stdin), os.Stdout) }
