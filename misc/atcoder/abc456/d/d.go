package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	const mod = 998244353
	var s string
	Fscan(in, &s)
	f := [3]int{}
	for _, c := range s {
		f[c-'a'] = (f[0] + f[1] + f[2] + 1) % mod
	}
	Fprint(out, (f[0]+f[1]+f[2])%mod)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
