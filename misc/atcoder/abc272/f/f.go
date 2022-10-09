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

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	var n int32
	var c, ans int
	var s, t string
	Fscan(bufio.NewReader(in), &n, &s, &t)
	sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New([]byte(s + s + "#" + t + t + "|"))).Elem().FieldByName("sa").Field(0).UnsafeAddr()))
	for _, p := range sa {
		if p < n {
			c++
		} else if n*2 < p && p <= n*3 {
			ans += c
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
