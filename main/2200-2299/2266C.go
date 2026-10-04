package main

import (
	. "fmt"
	"io"
	"strings"
)

// https://github.com/EndlessCheng
func cf2266C(in io.Reader, out io.Writer) {
	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		cnt := strings.Count(s, "0")
		if s[0] == '1' {
			Fprintln(out, cnt)
			continue
		}
		ans := cnt
		for _, b := range s {
			if b == '1' {
				cnt++
			} else {
				cnt--
			}
			ans = min(ans, cnt)
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2266C(bufio.NewReader(os.Stdin), os.Stdout) }
