package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1383E(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var s string
	Fscan(in, &s)
	n := len(s)
	a := make([]int, n+1)
	f := make([]int, n+1)
	pos := make([]int, n+2)
	for i, b := range s {
		a[i+1] = (1 - int(b-'0')) * (a[i] + 1)
	}
	for i := n; i > 0; i-- {
		ex := 0
		if a[i] <= a[n] {
			ex = 1
		}
		f[i] = (f[pos[a[i]+1]] + f[pos[0]] + ex) % mod
		pos[a[i]] = i
	}
	if pos[0] > 0 {
		Fprint(out, f[pos[0]]*pos[0]%mod)
	} else {
		Fprint(out, n)
	}
}

//func main() { cf1383E(bufio.NewReader(os.Stdin), os.Stdout) }
