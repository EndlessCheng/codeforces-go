package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1091D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int64
	Fscan(in, &n)
	ans := n
	for i := int64(2); i <= n; i++ {
		ans = (ans - 1) * i % 998244353
	}
	Fprint(out, ans)
}

//func main() {
//	CF1091D(os.Stdin, os.Stdout)
//}
