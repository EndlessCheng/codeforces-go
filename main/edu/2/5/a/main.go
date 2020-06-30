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
	var s []byte
	Fscan(bufio.NewReader(_r), &s)
	n := len(s)
	sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New(s)).Elem().FieldByName("sa").Field(0).UnsafeAddr()))
	rank := make([]int, n)
	for i := range rank {
		rank[sa[i]] = i
	}
	height := make([]int, n)
	ans, h := int64(n)*int64(n+1)/2, 0
	for i, ri := range rank {
		if h > 0 {
			h--
		}
		if ri > 0 {
			for j := int(sa[ri-1]); i+h < n && j+h < n && s[i+h] == s[j+h]; h++ {
			}
		}
		height[ri] = h
		ans -= int64(h)
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
