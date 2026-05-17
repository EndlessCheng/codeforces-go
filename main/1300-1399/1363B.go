package main

import (
	. "fmt"
	"io"
	"strings"
)

// https://github.com/EndlessCheng
func cf1363B(in io.Reader, out io.Writer) {
	T, s := 0, ""
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		n := len(s)
		suf1 := strings.Count(s, "1")
		pre1 := 0
		ans := n
		for i, b := range s {
			suf1 -= int(b - '0')
			pre1 += int(b - '0')
			ans = min(ans, i+1-pre1+suf1, pre1+n-i-1-suf1)
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1363B(bufio.NewReader(os.Stdin), os.Stdout) }
