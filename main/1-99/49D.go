package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF49D(in io.Reader, out io.Writer) {
	var n, ans int
	var s string
	Fscan(bufio.NewReader(in), &n, &s)
	for i, b := range s {
		ans += (int(b)^i)&1 ^ 1
	}
	if ans*2 > n {
		ans = n - ans
	}
	Fprint(out, ans)
}

//func main() { CF49D(os.Stdin, os.Stdout) }
