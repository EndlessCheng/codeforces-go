package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// NOTE: 洛谷的机器是 64 位的，有时候为了避免不必要的内存浪费以及加快运算速度，请使用 typedef
type int int32

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewScanner(_r)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	read := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b-'0')
		}
		return
	}

	n := read()

	Fprintln(out, n)
}

func main() { run(os.Stdin, os.Stdout) }
