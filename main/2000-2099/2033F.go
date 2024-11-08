package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2033F(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		if m == 1 {
			Fprintln(out, n%mod)
			continue
		}
		f0, f1 := 1, 1
		for i := 3; ; i++ {
			nf := (f0 + f1) % m
			if nf == 0 {
				Fprintln(out, n%mod*i%mod)
				break
			}
			f0 = f1
			f1 = nf
		}
	}
}

//func main() { cf2033F(bufio.NewReader(os.Stdin), os.Stdout) }
