package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var a,b int
	Fscan(in, &a,&b)
	Fprintf(out, "%.3f\n", float64(b)/float64(a))
}

func main() { run(os.Stdin, os.Stdout) }
