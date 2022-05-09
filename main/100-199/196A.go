package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF196A(in io.Reader, out io.Writer) {
	var s string
	Fscan(bufio.NewReader(in), &s)
	ans := []byte{}
	for i := byte('z'); i >= 'a' && s != ""; i-- {
		j := strings.LastIndexByte(s, i)
		if j >= 0 {
			ans = append(ans, strings.Repeat(string(i), strings.Count(s, string(i)))...)
			s = s[j+1:]
		}
	}
	Fprintf(out, "%s", ans)
}

//func main() { CF196A(os.Stdin, os.Stdout) }
