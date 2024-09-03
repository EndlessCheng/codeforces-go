package main

import (
	. "fmt"
	"io"
)

func cf1969C(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		f := make([][11]int, n+1)
		for i := range a {
			Fscan(in, &a[i])
			for j := 0; j <= k; j++ {
				res := f[i][j] + a[i]
				mn := a[i]
				for l := i - 1; l >= max(i-j, 0); l-- {
					mn = min(mn, a[l])
					t := i - l
					res = min(res, f[l][j-t]+mn*(t+1))
				}
				f[i+1][j] = res
			}
		}
		Fprintln(out, f[n][k])
	}
}

//func main() { cf1969C(bufio.NewReader(os.Stdin), os.Stdout) }
