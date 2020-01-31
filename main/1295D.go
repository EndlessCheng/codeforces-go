package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1295D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	calcGCD := func(a, b int64) int64 {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	calcPhi := func(n int64) int64 {
		ans := n
		for i := int64(2); i*i <= n; i++ {
			if n%i == 0 {
				ans = ans / i * (i - 1)
				for ; n%i == 0; n /= i {
				}
			}
		}
		if n > 1 {
			ans = ans / n * (n - 1)
		}
		return ans
	}

	var t int
	for Fscan(in, &t); t > 0; t-- {
		var a, m int64
		Fscan(in, &a, &m)
		Fprintln(out, calcPhi(m/calcGCD(a, m)))
	}
}

//func main() {
//	CF1295D(os.Stdin, os.Stdout)
//}
