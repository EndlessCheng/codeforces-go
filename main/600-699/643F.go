package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf643F(in io.Reader, out io.Writer) {
	var n, p, q uint32
	Fscan(in, &n, &p, &q)
	f := make([]uint32, p+1)
	f[0] = 1
	for x := 29; x >= 0; x-- {
		for i := p; i > 0; i-- {
			for j := range i {
				f[i] += f[j] * f[i-j]
			}
		}
		if n>>x&1 != 0 {
			for i := p; i > 0; i-- {
				f[i] += f[i-1]
			}
		}
	}

	ans := uint32(0)
	for i := uint32(1); i <= q; i++ {
		var x, y uint32 = 0, 1
		for j := uint32(0); j <= p && j < n; j++ {
			x += y * f[j]
			y *= i
		}
		ans ^= i * x
	}
	Fprint(out, ans)
}

//func main() { cf643F(bufio.NewReader(os.Stdin), os.Stdout) }
