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
	var s, p []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &p)
		var id []interface{}
	o:
		for i := len(p); i <= len(s); i++ {
			for j, b := range s[i-len(p) : i] {
				if p[j] != '?' && p[j] != b {
					continue o
				}
			}
			id = append(id, i-len(p))
		}
		Fprintln(out, len(id))
		Fprintln(out, id...)
	}
}

func main() { run(os.Stdin, os.Stdout) }
