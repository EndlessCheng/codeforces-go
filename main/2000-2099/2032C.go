package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2032C(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		slices.Sort(a)

		ans := n - 2
		j := 0
		for i := 2; i < n; i++ {
			for a[j]+a[j+1] <= a[i] {
				j++
			}
			ans = min(ans, n-1-i+j)
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2032C(bufio.NewReader(os.Stdin), os.Stdout) }
