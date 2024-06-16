package main

import (
	"bufio"
	. "fmt"
	"io"
)

func cf1978F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx = 1000001
	pf := [mx][]int{}
	for i := 2; i < mx; i++ {
		if pf[i] == nil {
			for j := i; j < mx; j += i {
				pf[j] = append(pf[j], i)
			}
		}
	}

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		ans := n*2 - 1
		for i := range a {
			Fscan(in, &a[i])
			if a[i] == 1 {
				ans += n - 2
			}
		}
		if a[0] == 1 {
			ans++
		}
		fa := make([]int, n*2)
		var find func(int) int
		find = func(x int) int {
			if fa[x] != x {
				fa[x] = find(fa[x])
			}
			return fa[x]
		}
		pre := map[int]int{}
		for i := 1; i < n*2; i++ {
			fa[i] = i
			for _, p := range pf[a[i%n]] {
				if j := pre[p]; j > 0 && i-j <= k {
					x, y := find(i), find(j)
					if x != y {
						fa[x] = y
						ans--
					}
				}
				pre[p] = i
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1978F(bufio.NewReader(os.Stdin), os.Stdout) }
