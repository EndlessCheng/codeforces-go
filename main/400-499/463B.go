package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF463B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, ans int
	for Fscan(in, &n); n > 0; n-- {
		if Fscan(in, &v); v > ans {
			ans = v
		}
	}
	Fprint(out, ans)
}

//func main() { CF463B(os.Stdin, os.Stdout) }
