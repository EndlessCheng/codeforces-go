package main

import (
	"bufio"
	. "fmt"
	"index/suffixarray"
	"io"
	"os"
	"reflect"
	"unsafe"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var s, t []byte
	Fscan(in, &s, &t)
	l := len(s)
	s = append(s, '#')
	s = append(s, t...)
	n := len(s)
	sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New(s)).Elem().FieldByName("sa").Field(0).UnsafeAddr()))
	rank := make([]int, n)
	for i := range rank {
		rank[sa[i]] = i
	}
	height := make([]int, n)
	var h, maxH, p int
	for i, ri := range rank {
		if h > 0 {
			h--
		}
		if ri > 0 {
			for j := int(sa[ri-1]); i+h < n && j+h < n && s[i+h] == s[j+h]; h++ {
			}
		}
		height[ri] = h
		if h > maxH && int(sa[ri]) < l != (int(sa[ri-1]) < l) {
			maxH, p = h, ri
		}
	}
	q := int(sa[p])
	Fprint(out, string(s[q:q+maxH]))
}

func main() { run(os.Stdin, os.Stdout) }
