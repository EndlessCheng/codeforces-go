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

	var T, n, k, v int
	for Fscan(in, &T); T > 0; T-- {
		ans := 100
		for Fscan(in, &n, &k); n > 0; n-- {
			if Fscan(in, &v); k%v == 0 && k/v < ans {
				ans = k / v
			}
		}
		Fprintln(out, ans)
	}
}

func main() { run(os.Stdin, os.Stdout) }
