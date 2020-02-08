package main

import (
	"bufio"
	. "fmt"
	"io"
)

func run(_r io.Reader, _w io.Writer) {
	type int int32
	type uint uint32
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
	readNeg := func() (x int) {
		in.Scan()
		data := in.Bytes()
		if data[0] == '-' {
			for _, b := range data {
				x = x*10 + int(b-'0')
			}
			return -x
		}
		for _, b := range data {
			x = x*10 + int(b-'0')
		}
		return x
	}

	n := read()

	Fprintln(out, n, readNeg())
}

//func main() { run(os.Stdin, os.Stdout) }
