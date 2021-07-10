package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	. "strings"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		for i := 1; ; i++ {
			if k <= i {
				Fprintln(out, Repeat("a", n-1-i)+"b"+Repeat("a", i-k)+"b"+Repeat("a", k-1))
				break
			}
			k -= i
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
