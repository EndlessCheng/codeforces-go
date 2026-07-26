package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1848E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var x, q, mod int
	Fscan(in, &x, &q, &mod)

	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n >>= 1 {
			if n&1 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	s := 1
	t := 0

	cal := func(x int, add bool) {
		y := 0
		for x%mod == 0 {
			x /= mod
			y++
		}
		if add {
			t += y
			s = s * x % mod
		} else {
			t -= y
			s = s * pow(x, mod-2) % mod
		}
	}

	cnt := map[int]int{}
	factor := func(x int) {
		for x&1 == 0 {
			x >>= 1
		}
		for i := 3; i*i <= x; i += 2 {
			if x%i == 0 {
				cal(cnt[i]+1, false)
				for x%i == 0 {
					x /= i
					cnt[i]++
				}
				cal(cnt[i]+1, true)
			}
		}
		if x > 1 {
			cal(cnt[x]+1, false)
			cnt[x]++
			cal(cnt[x]+1, true)
		}
	}

	for x&1 == 0 {
		x >>= 1
	}
	for i := 3; i*i <= x; i += 2 {
		if x%i == 0 {
			cal(cnt[i]+1, false)
			for x%i == 0 {
				x /= i
				cnt[i]++
			}
			cal(cnt[i]+1, true)
		}
	}
	if x > 1 {
		cnt[x]++
		s = s * 2 % mod
	}

	for range q {
		Fscan(in, &x)
		factor(x)
		if t > 0 {
			Fprintln(out, 0)
		} else {
			Fprintln(out, s)
		}
	}
}

//func main() { cf1848E(bufio.NewReader(os.Stdin), os.Stdout) }
