package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2259G(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		s := make([]int, n+1)
		for i := range a {
			Fscan(in, &a[i])
			s[i+1] = s[i] + a[i]
			a[i] -= i * k
		}

		ans := make([]any, n)
		j := n - 1
		for i := n - 1; i > 0; i-- {
			b := a[i-1] - k
			for a[j] <= b {
				j--
			}
			ans[i] = s[j+1] - s[i+1] - (i+j+1)*(j-i)/2*k - (j-i)*b
		}
		ans[0] = 0
		Fprintln(out, ans...)
	}
}

//func main() { cf2259G(bufio.NewReader(os.Stdin), os.Stdout) }
