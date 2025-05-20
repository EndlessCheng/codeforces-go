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
		// s[i] = 0，那么上一位必须进位且本位上取整，f[i] = f[i-1] * (1/2)
		// s[i] = 1，那么上一位进位，或者上一位没有进位且本位上取整，f[i] = f[i-1] + (1-f[i-1]) * (1/2) = (f[i-1]+1) * (1/2)
		f := 0
		for i := n - 1; i > 0; i-- {
			f = (f + int(s[i]&1)) * inv2 % mod
		}
		Fprintln(out, (n-1+f)%mod)
	}
}

//func main() { cf2081A(bufio.NewReader(os.Stdin), os.Stdout) }
