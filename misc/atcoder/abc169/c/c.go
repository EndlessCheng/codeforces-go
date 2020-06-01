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

	var a, b, c int
	Fscanf(in, "%d%d.%d", &a, &b, &c)
	Fprint(out, a*(100*b+c)/100)
}

func main() { run(os.Stdin, os.Stdout) }
