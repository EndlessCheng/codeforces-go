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

	var T int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		for k := len(s) - 1; k >= 0; k-- {
			i, j := 0, k
			for ; i < j && s[i] == s[j]; i++ {
				j--
			}
			if i >= j {
				Fprintln(out, k+1)
				break
			}
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
