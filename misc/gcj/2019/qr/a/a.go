package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t int
	var s []byte
	Fscan(in, &t)
	for i := 1; i <= t; i++ {
		Fscan(in, &s)
		ans := make([]byte, len(s))
		pos := -1
		for i, b := range s {
			if b == '4' {
				s[i] = '2'
				ans[i] = '2'
				if pos == -1 {
					pos = i
				}
			} else if pos != -1 {
				ans[i] = '0'
			}
		}
		Fprintf(out, "Case #%d: %s %s\n", i, s, ans[pos:])
	}
}

func main() { run(os.Stdin, os.Stdout) }
