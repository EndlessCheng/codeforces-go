package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf159D(in io.Reader, out io.Writer) {
	var s string
	Fscan(in, &s)

	n := len(s)
	sum := make([]int, n)
	for i := range 2*n - 1 {
		l, r := i/2, (i+1)/2
		for l >= 0 && r < n && s[l] == s[r] {
			sum[r]++
			l--
			r++
		}
	}

	for i := 1; i < n; i++ {
		sum[i] += sum[i-1]
	}

	ans := 0
	for i := range 2*n - 1 {
		l, r := i/2, (i+1)/2
		for l > 0 && r < n && s[l] == s[r] {
			ans += sum[l-1]
			l--
			r++
		}
	}
	Fprint(out, ans)
}

//func main() { cf159D(os.Stdin, os.Stdout) }
