package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1119E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, left int
	ans := int64(0)
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		if left > 0 {
			if v > 2*left {
				ans += int64(left)
				v -= 2 * left
				left = 0
			} else {
				ans += int64(v / 2)
				left -= v / 2
				v %= 2
			}
		}
		ans += int64(v / 3)
		left += v % 3
	}
	Fprint(out, ans)
}

//func main() { CF1119E(os.Stdin, os.Stdout) }
