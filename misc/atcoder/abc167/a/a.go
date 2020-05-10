package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s, t string
	Fscan(in, &s, &t)
	if strings.HasPrefix(t, s) {
		Fprintln(out, "Yes")
	} else {
		Fprintln(out, "No")
	}
}

func main() { run(os.Stdin, os.Stdout) }
