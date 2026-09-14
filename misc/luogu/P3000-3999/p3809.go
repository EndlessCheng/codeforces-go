package main

import (
	"bytes"
	"index/suffixarray"
	"io"
	"runtime/debug"
	"syscall"
	"unsafe"
)

// https://space.bilibili.com/206214
func init() { debug.SetGCPercent(-1) }

var t809 [6888896]byte

func p3809(in io.Reader, _w io.Writer) {
	n, _ := syscall.Read(syscall.Stdin, t809[:])
	sa := t809[:0]
	tmp := [7]byte{}
	for _, v := range (*struct{_[]byte;sa[]int32})(unsafe.Pointer(suffixarray.New(bytes.TrimSpace(t809[:n])))).sa {
		p := len(tmp)
		for v++; v > 0; v /= 10 {
			p--
			tmp[p] = '0' | byte(v%10)
		}
		sa = append(append(sa, tmp[p:]...), ' ')
	}
	syscall.Write(syscall.Stdout, sa)
}