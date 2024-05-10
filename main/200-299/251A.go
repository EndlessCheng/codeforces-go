package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf251A(in io.Reader, out io.Writer) {
	var n, d, l, ans int
	Fscan(in, &n, &d)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		for a[l] < a[i]-d {
			l++
		}
		ans += (i - l) * (i - l - 1) / 2
	}
	Fprint(out, ans)
}

//func main() { cf251A(bufio.NewReader(os.Stdin), os.Stdout) }
