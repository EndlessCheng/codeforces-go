package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1650B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, l, r, a int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &l, &r, &a)
		ans := r/a + r%a
		if r-r%a-1 >= l && r/a+a-2 > ans {
			ans = r/a + a - 2
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1650B(os.Stdin, os.Stdout) }
