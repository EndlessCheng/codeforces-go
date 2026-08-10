package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2246C(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		g, d1 := 1, 1
		for i := range a {
			Fscan(in, &a[i])
			if i > 0 && a[i] != a[i-1] {
				g++
				if a[0] < 0 && a[i]-a[i-1] == 1 {
					d1++
				}
			}
		}
		Fprintln(out, pow(2, n-g)*d1%mod)
	}
}

//func main() { cf2246C(bufio.NewReader(os.Stdin), os.Stdout) }
