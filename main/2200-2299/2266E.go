package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2266E(in io.Reader, out io.Writer) {
	const mx = 200_001
	pf := [mx][]int{}
	for i := 2; i < mx; i++ {
		if pf[i] == nil {
			for j := i; j < mx; j += i {
				pf[j] = append(pf[j], i)
			}
		}
	}

	var T, n, k, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		f := make([]int, n+1)
		for i := k + 1; i <= n; i++ {
			f[i] = 1e9
			for _, p := range pf[i] {
				f[i] = min(f[i], f[i/p]*p+1)
			}
		}

		ans := 0
		for range n {
			Fscan(in, &v)
			ans += f[v]
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2266E(bufio.NewReader(os.Stdin), os.Stdout) }
