package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf578D(in io.Reader, out io.Writer) {
	var n, m, cnt int
	var s string
	Fscan(in, &n, &m, &s)
	s = "##" + s
	ans := 1
	for i := 1; i <= n; i++ {
		if s[i+1] == s[i-1] {
			cnt++
		} else {
			cnt = 0
		}
		if s[i+1] != s[i] {
			ans += n*(m-1) - cnt - 1
		}
	}
	Fprint(out, ans)
}

//func main() { cf578D(bufio.NewReader(os.Stdin), os.Stdout) }
