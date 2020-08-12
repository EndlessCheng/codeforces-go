package main

import (
	"bufio"
	. "fmt"
	"io"
	"strconv"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF747E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s string
	Fscan(bufio.NewReader(in), &s)
	sp := strings.Split(s, ",")
	ss := [4e5][]interface{}{}
	maxD, p := 0, 0
	var f func(d int)
	f = func(d int) {
		if d > maxD {
			maxD = d
		}
		ss[d] = append(ss[d], sp[p])
		ch, _ := strconv.Atoi(sp[p+1])
		p += 2
		for ; ch > 0; ch-- {
			f(d + 1)
		}
	}
	for p < len(sp) {
		f(0)
	}
	Fprintln(out, maxD+1)
	for _, a := range ss[:maxD+1] {
		Fprintln(out, a...)
	}
}

//func main() { CF747E(os.Stdin, os.Stdout) }
