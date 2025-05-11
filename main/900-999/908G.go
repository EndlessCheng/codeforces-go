package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf908G(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var s string
	Fscan(in, &s)
	n := len(s)
	ans := 0
	type pair struct{ f, g int }
	memo := make([]pair, n)
	for tar := 1; tar < 10; tar++ {
		for i := range memo {
			memo[i].f = -1
		}
		var f func(int, bool) pair
		f = func(i int, isLimit bool) (res pair) {
			if i == n {
				return pair{0, 1}
			}
			if !isLimit {
				p := &memo[i]
				if p.f >= 0 {
					return *p
				}
				defer func() { *p = res }()
			}

			up := 9
			if isLimit {
				up = int(s[i] - '0')
			}

			for d := range up + 1 {
				p := f(i+1, isLimit && d == up)
				if d < tar {
					res.f += p.f
					res.g += p.g
				} else if d == tar {
					// 比如 s="23"，tar=1
					// p.f*10 的意思是，第一位填 1，会把第二位填的 1（排序后）变成 10 
					// p.g 的意思是，第一位填 1，那么第二位填 0~9 对第一位的 1 的影响（乘积系数）是 1,1,10,10,...,10，一共是 82，这个就是 p.g
					res.f += p.f*10 + p.g
					res.g += p.g
				} else {
					res.f += p.f * 10
					res.g += p.g * 10
				}
			}
			res.f %= mod
			res.g %= mod
			return
		}
		ans += f(0, true).f * tar
	}
	Fprint(out, ans%mod)
}

//func main() { cf908G(bufio.NewReader(os.Stdin), os.Stdout) }
