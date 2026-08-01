package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2225D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353
	var T, n, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x)
		ans := ((n+1)/4 - x/4) % mod * ((x/4 + 1) % mod)
		ans += ((n+3)/4 - (x+2)/4) % mod * ((x + 2) / 4 % mod)
		Fprintln(out, ans%mod)
	}
}

//func main() { cf2225D(bufio.NewReader(os.Stdin), os.Stdout) }
