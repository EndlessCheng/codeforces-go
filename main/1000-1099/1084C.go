package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1084C(in io.Reader, out io.Writer) {
	const mod int = 1e9 + 7
	var s string
	Fscan(bufio.NewReader(in), &s)

	ans, pre := 0, 0
	for _, ch := range s {
		if ch == 'a' {
			ans = (ans + pre + 1) % mod
		} else if ch == 'b' {
			pre = ans
		}
	}
	Fprint(out, ans)
}

func CF1084C_(in io.Reader, out io.Writer) {
	const mod int = 1e9 + 7
	var s string
	Fscan(bufio.NewReader(in), &s)

	f := make([]int, len(s)+1)
	pre := 0
	for i, b := range s {
		f[i+1] = f[i]
		if b == 'a' {
			f[i+1] = (f[i+1] + 1 + f[pre]) % mod
		} else if b == 'b' {
			pre = i
		}
	}
	Fprint(out, f[len(s)])
}

//func main() { CF1084C(os.Stdin, os.Stdout) }
