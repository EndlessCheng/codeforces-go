package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1904C(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		if k > 2 {
			Fprintln(out, 0)
			continue
		}
		slices.Sort(a)
		ans := a[0]
		for i := 1; i < n; i++ {
			ans = min(ans, a[i]-a[i-1])
		}
		if k == 2 {
			for i, v := range a {
				j := i - 1
				for _, w := range a[:i] {
					d := v - w
					for j >= 0 && d < a[j] {
						j--
					}
					ans = min(ans, a[j+1]-d)
					if j >= 0 {
						ans = min(ans, d-a[j])
					}
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1904C(bufio.NewReader(os.Stdin), os.Stdout) }
