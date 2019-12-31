package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol727C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)

	var n, s12, s13, s23 int
	Fscan(in, &n)
	Fprintln(out, "? 1 2")
	out.Flush()
	Fscan(in, &s12)
	Fprintln(out, "? 1 3")
	out.Flush()
	Fscan(in, &s13)
	Fprintln(out, "? 2 3")
	out.Flush()
	Fscan(in, &s23)

	ans := make([]int, n)
	ans[0] = (s12 + s13 - s23) / 2
	ans[1] = s12 - ans[0]
	ans[2] = s13 - ans[0]
	for i := 3; i < n; i++ {
		Fprintln(out, Sprintf("? 1 %d", i+1))
		out.Flush()
		var sum int
		Fscan(in, &sum)
		ans[i] = sum - ans[0]
	}
	Fprint(out, "!")
	for _, v := range ans {
		Fprint(out, " ", v)
	}
	out.Flush()
}

//func main() {
//	Sol727C(os.Stdin, os.Stdout)
//}
