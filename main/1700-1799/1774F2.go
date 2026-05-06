package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1774F2(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n, dec int
	Fscan(in, &n)
	op := make([]int, n)
	x := make([]int, n)
	for i := range op {
		Fscan(in, &op[i])
		if op[i] == 1 {
			Fscan(in, &x[i])
		} else if op[i] == 2 {
			Fscan(in, &x[i])
			dec += x[i]
		} else {
			x[i] = dec
			dec *= 2
		}
		dec = min(dec, 1e18)
	}

	ans := 0
	pow2 := 1
	s := 0
	a := []int{}
	for i := n - 1; i >= 0; i-- {
		if op[i] == 1 {
			left := x[i] - s
			for j, v := range a {
				if v < left {
					ans = (ans + 1<<(len(a)-1-j)%mod*pow2) % mod
					left -= v
				}
			}
			if left > 0 {
				ans = (ans + pow2) % mod
			}
		} else if op[i] == 2 {
			s += x[i]
		} else if x[i] == 0 {
			pow2 = pow2 * 2 % mod
		} else if x[i] != 1e18 {
			a = append(a, x[i])
		}
	}
	Fprint(out, ans)
}

//func main() { cf1774F2(bufio.NewReader(os.Stdin), os.Stdout) }
