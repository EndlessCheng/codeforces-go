package main

import (
	"bufio"
	. "fmt"
	"io"
)

func pow484A(n int64) int64 {
	if n == 0 {
		return 0
	}
	msb := uint(0)
	for n >>= 1; n != 0; n >>= 1 {
		msb++
	}
	return int64(1) << msb
}

func f484A(l, r int64, i int) int64 {
	if l == r || (r+1)&r == 0 {
		return r
	}
	if common := l & r; common < r&^common {
		return pow484A(r) - 1
	}
	for ; i >= 0; i-- {
		pow2 := int64(1) << uint(i)
		if l&pow2 > 0 && r&pow2 > 0 {
			return pow2 | f484A(l-pow2, r-pow2, i-1)
		}
	}
	panic("err")
}

func Sol484A(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	for Fscan(in, &n); n > 0; n-- {
		var l, r int64
		Fscan(in, &l, &r)
		Fprintln(out, f484A(l, r, 60))
	}
}

//func main() {
//	Sol484A(os.Stdin, os.Stdout)
//}
