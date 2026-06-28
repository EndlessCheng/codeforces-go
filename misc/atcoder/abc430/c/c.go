package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, a, b, l1, l2, ans int
	var s string
	Fscan(in, &n, &a, &b, &s)
	cnt := [2]int{}
	for _, ch := range s {
		cnt[ch-'a']++
		for cnt[0] >= a {
			if s[l1] == 'a' {
				cnt[0]--
			}
			l1++
		}
		for cnt[1] >= b {
			if s[l2] == 'b' {
				cnt[1]--
			}
			l2++
		}
		ans += max(l1-l2, 0)
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
