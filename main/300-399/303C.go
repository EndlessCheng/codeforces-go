package main

import (
	. "fmt"
	"io"
)

// todo 时间复杂度的严格分析？

// https://github.com/EndlessCheng
func cf303C(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var n, k int
	Fscan(in, &n, &k)
	const mx int = 1e6 + 1
	cnt := [mx]int{}
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		for _, v := range a[:i] {
			cnt[abs(a[i]-v)]++
		}
	}

	vis := [mx]int{}
o:
	for m := n - k; m < mx; m++ {
		sum := 0
		for i := m; i < mx; i += m {
			sum += cnt[i]
		}
		if sum > k*(k+1)/2 {
			continue
		}
		del := 0
		for _, v := range a {
			if vis[v%m] == m {
				del++
				if del > k {
					continue o
				}
			} else {
				vis[v%m] = m
			}
		}
		Fprint(out, m)
		return
	}
}

//func main() { cf303C(bufio.NewReader(os.Stdin), os.Stdout) }
