package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func calc14(x, k int) int {
	ds := []int{}
	for d := 1; d*d <= x; d++ {
		if x%d == 0 {
			ds = append(ds, d)
			if d*d < x {
				ds = append(ds, x/d)
			}
		}
	}
	slices.Sort(ds)
	n := len(ds)
	f := make([]int, n)
	for i := 1; i < n; i++ {
		f[i] = 1e9
		for j := i - 1; j >= 0 && ds[i]/ds[j] <= k; j-- {
			if ds[i]%ds[j] == 0 {
				f[i] = min(f[i], f[j]+1)
			}
		}
	}
	return f[n-1]
}

func cf2114F(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	var T, x, y, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x, &y, &k)
		g := gcd(x, y)
		ans := calc14(x/g, k) + calc14(y/g, k)
		if ans >= 1e9 {
			ans = -1
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2114F(bufio.NewReader(os.Stdin), os.Stdout) }
