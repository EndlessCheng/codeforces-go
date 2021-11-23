package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF496B(in io.Reader, out io.Writer) {
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n int
	var base []byte
	Fscan(bufio.NewReader(in), &n, &base)

	ans := ""
	for loop := 10; loop > 0; loop-- {
		s := string(base) + string(base)
		i := 0
		for j := 1; j < n; {
			k := 0
			for ; k < n && s[i+k] == s[j+k]; k++ {
			}
			if k >= n {
				break
			}
			if s[i+k] < s[j+k] {
				j += k + 1
			} else {
				i, j = j, max(j, i+k)+1
			}
		}
		s = s[i : i+n]
		if ans == "" || s < ans {
			ans = s
		}
		for i, b := range base {
			if b == '9' {
				base[i] = '0'
			} else {
				base[i]++
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF496B(os.Stdin, os.Stdout) }
