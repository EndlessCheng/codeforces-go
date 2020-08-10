package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, p int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &p)
		if p == 0 {
			Fprintln(out, 0)
			continue
		}
		for n--; ; n-- {
			if p == 1<<n-1 {
				Fprintln(out, 0)
				break
			}
			if p >= 1<<n {
				p -= 1 << n
				if p == 0 {
					Fprintln(out, 1<<n-1)
					break
				}
			}
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
