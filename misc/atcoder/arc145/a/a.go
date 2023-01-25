package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	n, s := 0, ""
	Fscan(bufio.NewReader(in), &n, &s)
	if s == "BA" || s[0] == 'A' && s[n-1] == 'B' {
		Fprint(out, "No")
	} else {
		Fprint(out, "Yes")
	}
}

func main() { run(os.Stdin, os.Stdout) }
