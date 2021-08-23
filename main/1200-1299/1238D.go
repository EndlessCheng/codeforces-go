package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1238D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	var s string
	Fscan(in, &n, &s)

	ans := int64(0)
	p := [2]int{-1, -1}
	for i, b := range s {
		b -= 'A'
		ans += int64(p[b] + 1)
		if i > 0 && p[b] == i-1 && p[b^1] >= 0 { // AB...BB 或 BA...AA
			ans--
		}
		p[b] = i
	}
	Fprint(out, ans)
}

//func main() { CF1238D(os.Stdin, os.Stdout) }
