package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var n, a, b, j, ans int
	var c, r int64
	var s []byte
	Fscan(bufio.NewReader(in), &n, &c, &s)
	for i, ch := range s {
		if ch == 'a' {
			a++
		} else if ch == 'b' {
			b++
			r += int64(a)
		}
		for r > c {
			if s[j] == 'a' {
				r -= int64(b)
				a--
			} else if s[j] == 'b' {
				b--
			}
			j++
		}
		if i-j+1 > ans {
			ans = i - j + 1
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
