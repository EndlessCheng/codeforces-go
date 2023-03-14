package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func Sol1156B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var T int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		a := [2][]byte{}
		for _, b := range s {
			a[b&1] = append(a[b&1], b)
		}
		x, y := a[0], a[1]
		if len(x) == 0 {
			Fprintf(out, "%s\n", y)
		} else if len(y) == 0 {
			Fprintf(out, "%s\n", x)
		} else if abs(int(x[len(x)-1])-int(y[0])) != 1 {
			Fprintf(out, "%s%s\n", x, y)
		} else if abs(int(y[len(y)-1])-int(x[0])) != 1 {
			Fprintf(out, "%s%s\n", y, x)
		} else {
			Fprintln(out, "No answer")
		}
	}
}

//func main() { Sol1156B(os.Stdin, os.Stdout) }
