package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf926E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, v int
	Fscan(in, &n)
	st := []int{}
	for range n {
		Fscan(in, &v)
		st = append(st, v)
		for len(st) > 1 && st[len(st)-2] == st[len(st)-1] {
			st[len(st)-2]++
			st = st[:len(st)-1]
		}
	}
	Fprintln(out, len(st))
	for _, v := range st {
		Fprint(out, v, " ")
	}
}

//func main() { cf926E(bufio.NewReader(os.Stdin), os.Stdout) }
