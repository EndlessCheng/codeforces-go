package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1325D(in io.Reader, out io.Writer) {
	var xor, sum int64
	Fscan(in, &xor, &sum)
	if xor > sum || (sum-xor)&1 > 0 {
		Fprint(out, -1)
	} else if xor == sum {
		if xor == 0 {
			Fprint(out, 0)
		} else {
			Fprint(out, "1\n", xor)
		}
	} else if and := (sum - xor) / 2; and&xor > 0 { // p&q=1 和 p^q=1 不能同时成立
		Fprint(out, "3\n", xor, and, and)
	} else {
		Fprint(out, "2\n", xor|and, and) // and 可以直接加到 xor 里面
	}
}

//func main() { CF1325D(os.Stdin, os.Stdout) }
