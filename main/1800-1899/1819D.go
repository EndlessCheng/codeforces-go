package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1819D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m, k int
	vis := [2e5 + 1]int{}
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([][]int, n+1)
		used := []int{}
		for i := 1; i <= n; i++ {
			Fscan(in, &k)
			a[i] = make([]int, k)
			for j := range a[i] {
				Fscan(in, &a[i][j])
				used = append(used, a[i][j])
			}
		}

		rem := make([]int, n+2)
		zero := make([]int, n+2)

		ok := false
		cur := 0
		for i := n; i > 0; i-- {
			for _, x := range a[i] {
				cur = min(cur+1, m)
				if vis[x] != 0 {
					ok = true
				}
				vis[x] = i
			}
			if len(a[i]) == 0 {
				cur = m
			}
			if ok {
				rem[i] = 0
			} else {
				rem[i] = cur
			}
		}

		for _, x := range used {
			vis[x] = 0
		}

		empty := 0
		ans := rem[1]
		limit := 0

		for i := 1; i <= n; i++ {
			zero[i+1] = zero[i]
			if len(a[i]) == 0 {
				ans = max(ans, rem[i+1])
				empty = i
				zero[i+1] = i
				continue
			}

			mx := empty
			for _, x := range a[i] {
				mx = max(mx, vis[x])
			}
			if mx > 0 && limit <= zero[mx] {
				zero[i+1] = i
				ans = max(ans, rem[i+1])
			}
			for _, x := range a[i] {
				limit = max(limit, vis[x])
				vis[x] = i
			}
		}
		Fprintln(out, ans)
		for _, x := range used {
			vis[x] = 0
		}
	}
}

//func main() { cf1819D(bufio.NewReader(os.Stdin), os.Stdout) }
