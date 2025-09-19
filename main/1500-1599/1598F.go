package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1598F(in io.Reader, out io.Writer) {
	var n, ans int
	var s string
	Fscan(in, &n)
	cnt := make([][]int, n)
	bal := make([]int, n)
	minBal := make([]int, n)
	for i := range cnt {
		Fscan(in, &s)
		cnt[i] = make([]int, len(s)+1)
		b, mn := 0, 0
		for _, c := range s {
			b += int(1 - c%2*2)
			if b <= mn {
				cnt[i][-b]++
				mn = min(mn, b)
			}
		}
		bal[i] = b
		minBal[i] = mn
	}

	f := make([]int, 1<<n)
	for i := range f {
		f[i] = -1
	}
	f[0] = 0
	for s, fs := range f {
		if fs < 0 {
			continue
		}
		sum := 0
		for t := uint(s); t > 0; t &= t - 1 {
			sum += bal[bits.TrailingZeros(t)]
		}
		for t, lb := 1<<n-1^s, 0; t > 0; t ^= lb {
			lb = t & -t
			ns := s | lb
			i := bits.TrailingZeros(uint(lb))
			if sum < len(cnt[i]) {
				ans = max(ans, fs+cnt[i][sum])
			}
			if sum >= -minBal[i] {
				nf := fs
				if sum < len(cnt[i]) {
					nf += cnt[i][sum]
				}
				f[ns] = max(f[ns], nf)
			}
		}
	}
	Fprint(out, ans)
}

//func main() { cf1598F(bufio.NewReader(os.Stdin), os.Stdout) }
