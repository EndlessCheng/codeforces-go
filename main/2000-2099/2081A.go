package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2081A(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 1_000_000_007
	const inv2 = (mod + 1) / 2
	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		cnt1 := make([]int, n)
		c1 := 0
		for i, b := range s {
			if b == '0' {
				c1 = 0
			} else {
				c1++
				cnt1[i] = c1
			}
		}

		f := make([][2]int, n+1)
		for i := 1; i < n; i++ {
			for j := range 2 {
				res := f[i][0] + 1
				if j > 0 || s[i] == '1' {
					c := cnt1[i-j] + j
					res = (res + f[i-c+1][1] + c) * inv2 % mod
				}
				f[i+1][j] = res
			}
		}
		Fprintln(out, f[n][0])
	}
}

//func main() { cf2081A(bufio.NewReader(os.Stdin), os.Stdout) }
