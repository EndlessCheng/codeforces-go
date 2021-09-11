package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		low := 0
		for _, b := range s {
			if unicode.IsLower(b) {
				low++
			}
		}
		if low*2 < len(s) {
			Fprintln(out, strings.ToUpper(s))
		} else {
			Fprintln(out, strings.ToLower(s))
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
