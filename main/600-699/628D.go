package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF628D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod int = 1e9 + 7
	var m, D int
	var lower, upper string
	Fscan(in, &m, &D, &lower, &upper)

	calc := func(s string) int {
		const lowerC, upperC byte = '0', '9'
		n := len(s)
		dp := make([][]int, n)
		for i := range dp {
			dp[i] = make([]int, m)
			for j := range dp[i] {
				dp[i][j] = -1
			}
		}
		var f func(p, val int, limitUp bool) int
		f = func(p, val int, limitUp bool) (res int) {
			if p == n {
				if val == 0 {
					return 1
				}
				return
			}
			if !limitUp {
				dv := &dp[p][val]
				if *dv >= 0 {
					return *dv
				}
				defer func() { *dv = res }()
			}
			up := upperC
			if limitUp {
				up = s[p]
			}
			for d := lowerC; d <= up; d++ {
				if p&1 > 0 == (int(d&15) == D) {
					cnt := f(p+1, (val*10+int(d&15))%m, limitUp && d == up)
					res = (res + cnt) % mod
				}
			}
			return
		}
		return f(0, 0, true)
	}
	ans := calc(upper) - calc(lower)
	val := 0
	for i, b := range lower {
		if i&1 > 0 == (int(b&15) != D) {
			goto end
		}
		val = (val*10 + int(b&15)) % m
	}
	if val == 0 {
		ans++
	}
end:
	Fprint(out, (ans%mod+mod)%mod)
}

//func main() { CF628D(os.Stdin, os.Stdout) }
