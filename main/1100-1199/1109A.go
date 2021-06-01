package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1109A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v int
	Fscan(in, &n)
	s := make([]int, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		s[i+1] = s[i] ^ v
	}
	ans := int64(0)
	f := func(s []int) {
		c := map[int]int{}
		for i := 0; i < len(s); i += 2 {
			ans += int64(c[s[i]])
			c[s[i]]++
		}
	}
	f(s)
	f(s[1:])
	Fprint(out, ans)
}

//func main() { CF1109A(os.Stdin, os.Stdout) }
