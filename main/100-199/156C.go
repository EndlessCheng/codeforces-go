package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf156C(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	const mx = 2500
	f := [101][mx + 1]int{}
	g := [mx + 1]int{1}
	for i := 1; i < len(f); i++ {
		const c = 25
		for j := 1; j <= mx; j++ {
			g[j] = (g[j] + g[j-1]) % mod
		}
		for j := mx; j > c; j-- {
			g[j] -= g[j-c-1]
		}
		f[i] = g
	}

	var T int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		tot := 0
		for _, v := range s {
			tot += int(v - 'a')
		}
		Fprintln(out, ((f[len(s)][tot]-1)%mod+mod)%mod)
	}
}

//func main() { cf156C(bufio.NewReader(os.Stdin), os.Stdout) }
