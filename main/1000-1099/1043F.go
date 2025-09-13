package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1043F(in io.Reader, out io.Writer) {
	const mx = 300_001
	var n, v, g int
	Fscan(in, &n)
	var hasG, has [mx]bool
	for range n {
		Fscan(in, &v)
		g = gcd43(g, v)
		hasG[v] = true
		has[v] = true
	}
	if g > 1 {
		Fprint(out, -1)
		return
	}

	f := [mx]int{}
	for ans := 1; ; ans++ {
		if hasG[1] {
			Fprint(out, ans)
			return
		}
		for i := mx - 1; i > 0; i-- {
			cntG, cnt := 0, 0
			for j := i; j < mx; j += i {
				if hasG[j] {
					cntG++
				}
				if has[j] {
					cnt++
				}
			}
			f[i] = cntG * cnt
			for j := i * 2; j < mx; j += i {
				f[i] -= f[j]
			}
		}
		for i := 1; i < mx; i++ {
			hasG[i] = f[i] > 0
		}
	}
}

//func main() { cf1043F(bufio.NewReader(os.Stdin), os.Stdout) }
func gcd43(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
