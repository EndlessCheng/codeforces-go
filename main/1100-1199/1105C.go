package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1105C(in io.Reader, out io.Writer) {
	const mod int64 = 1e9 + 7
	var n, l, r int64
	Fscan(in, &n, &l, &r)
	l--
	c := [3]int64{r/3 - l/3, (r+2)/3 - (l+2)/3, (r+1)/3 - (l+1)/3}
	f := [3]int64{1}
	for ; n > 0; n-- {
		f = [3]int64{
			(f[0]*c[0] + f[1]*c[2] + f[2]*c[1]) % mod,
			(f[0]*c[1] + f[1]*c[0] + f[2]*c[2]) % mod,
			(f[0]*c[2] + f[1]*c[1] + f[2]*c[0]) % mod,
		}
	}
	Fprint(out, f[0])
}

//func main() { CF1105C(os.Stdin, os.Stdout) }
