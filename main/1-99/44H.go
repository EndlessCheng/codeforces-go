package main

import (
	. "fmt"
	"io"
)

// todo 答案的上界是多少？如何构造一个让答案最大的输入？

// github.com/EndlessCheng/codeforces-go
func CF44H(in io.Reader, out io.Writer) {
	var s string
	Fscan(in, &s)
	f := [10]int64{}
	for i := range f {
		f[i] = 1
	}
	for _, c := range s[1:] {
		c := int(c & 15)
		g := [10]int64{}
		for i, v := range f {
			j := (i + c) / 2
			g[j] += v
			if (i+c)%2 > 0 {
				g[j+1] += v
			}
		}
		f = g
	}

	ans := int64(0)
	for _, v := range f {
		ans += v
	}
	for i := 1; i < len(s); i++ {
		if s[i] > s[i-1]+1 || s[i] < s[i-1]-1 {
			Fprint(out, ans)
			return
		}
	}
	Fprint(out, ans-1)
}

//func main() { CF44H(os.Stdin, os.Stdout) }
