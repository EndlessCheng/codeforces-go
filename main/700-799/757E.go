package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF757E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mod, mx int = 1e9 + 7, 1e6 + 1
	lpf := [mx]int{1: 1}
	for i := 2; i < mx; i++ {
		if lpf[i] == 0 {
			for j := i; j < mx; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}

	f := [mx][20]int{}
	f[0][0] = 1
	for j := 1; j < 20; j++ {
		f[0][j] = 2
	}
	for i := 1; i < mx; i++ {
		f[i][0] = 1
		for j := 1; j < 20; j++ {
			f[i][j] = (f[i][j-1] + f[i-1][j]) % mod
		}
	}

	var q, r, n int
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &r, &n)
		ans := 1
		for n > 1 {
			p := lpf[n]
			e := 1
			for n /= p; lpf[n] == p; n /= p {
				e++
			}
			ans = ans * f[r][e] % mod
		}
		Fprintln(out, ans)
	}
}

//func main() { CF757E(os.Stdin, os.Stdout) }
