package main

import (
	. "fmt"
	"io"
	"strconv"
)

// https://github.com/EndlessCheng
func cf204A(in io.Reader, out io.Writer) {
	f := func(n int) int {
		if n < 10 {
			return n
		}
		res := 9 + n/10
		if int(strconv.Itoa(n)[0]-'0') > n%10 {
			res--
		}
		return res
	}
	var l, r int
	Fscan(in, &l, &r)
	Fprint(out, f(r)-f(l-1))
}

//func main() { cf204A(os.Stdin, os.Stdout) }
