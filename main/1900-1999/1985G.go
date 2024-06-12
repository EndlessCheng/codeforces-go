package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1985G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
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
	var T, l, r, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &l, &r, &k)
		if k > 9 {
			Fprintln(out, 0)
		} else {
			Fprintln(out, (pow(9/k+1, r)-pow(9/k+1, l)+mod)%mod)
		}
	}
}

//func main() { cf1985G(bufio.NewReader(os.Stdin), os.Stdout) }
