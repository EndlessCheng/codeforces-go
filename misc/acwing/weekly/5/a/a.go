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

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		has := 0
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &v)
			has |= 1 << (v & 1)
		}
		if has == 3 {
			Fprintln(out, "NO")
		} else {
			Fprintln(out, "YES")
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
