package main

import (
	. "fmt"
	"io"
	"strings"
)

func cf115B(in io.Reader, out io.Writer) {
	var n, m, ans, cur, down int
	var s string
	Fscan(in, &n, &m)
	for i := 0; i < n; i++ {
		Fscan(in, &s)
		l := strings.IndexByte(s, 'W')
		if l < 0 {
			continue
		}
		down = i
		r := strings.LastIndexByte(s, 'W')
		if i%2 == 0 {
			if cur <= l {
				ans += r - cur
			} else {
				ans += cur - l + r - l
			}
			cur = r
		} else {
			if cur >= r {
				ans += cur - l
			} else {
				ans += r - cur + r - l
			}
			cur = l
		}
	}
	Fprint(out, ans+down)
}

//func main() { cf115B(bufio.NewReader(os.Stdin), os.Stdout) }
