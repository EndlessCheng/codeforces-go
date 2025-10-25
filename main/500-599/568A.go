package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf568A(in io.Reader, out io.Writer) {
	const mx int = 2e6 + 1
	np := [mx]bool{true, true}
	for i := 2; i < mx; i++ {
		if !np[i] {
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
	isPal := func(x int) bool {
		if x%10 == 0 {
			return false
		}
		rev := 0
		for rev < x/10 {
			rev = rev*10 + x%10
			x /= 10
		}
		return rev == x || rev == x/10
	}

	var p, q, a, b, ans int
	Fscan(in, &p, &q)
	for i := 1; i < mx; i++ {
		if !np[i] {
			a++
		}
		if isPal(i) {
			b++
		}
		if a*q <= b*p {
			ans = i
		}
	}
	Fprint(out, ans)
}

//func main() { cf568A(os.Stdin, os.Stdout) }
