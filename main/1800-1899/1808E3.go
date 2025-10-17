package main

import (
	. "fmt"
	"io"
)

// https://chatgpt.com/c/68e8f447-7ec8-8320-b8ea-7dfbe496a959
// https://www.luogu.com.cn/article/muox5g44

// https://github.com/EndlessCheng
func cf1808E3(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	var n, k, M, ans int
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % M
			}
			x = x * x % M
		}
		return res
	}
	Fscan(in, &n, &k, &M)
	if n == 1 {
		Fprint(out, 1)
		return
	}
	if k%2 > 0 {
		ans = pow(k, n) - pow(k-1, n) - pow(-1, n)*(gcd(n-2, k)-1)
	} else {
		inv2 := pow(2, M-2)
		ans = pow(k, n)*inv2 - pow(2, n-1)*(pow(k/2-1, n)+pow(-1, n)*(gcd(n-2, k/2)-1))
	}
	Fprint(out, (ans%M+M)%M)
}

//func main() { cf1808E3(os.Stdin, os.Stdout) }
