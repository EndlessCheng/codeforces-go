package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1999E(in io.Reader, out io.Writer) {
	const mx = 200001
	f := [mx]int{}
	for i := 1; i < mx; i++ {
		f[i] = f[i/3] + 1
	}
	for i := 2; i < mx; i++ {
		f[i] += f[i-1]
	}
	var T, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &l, &r)
		Fprintln(out, f[r]+f[l]-f[l-1]*2)
	}
}

//func main() { cf1999E(bufio.NewReader(os.Stdin), os.Stdout) }
