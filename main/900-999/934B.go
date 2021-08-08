package main

import (
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF934B(in io.Reader, out io.Writer) {
	var k int
	if Fscan(in, &k); k > 36 {
		Fprint(out, -1)
	} else {
		Fprint(out, strings.Repeat("8", k/2), strings.Repeat("4", k&1))
	}
}

//func main() { CF934B(os.Stdin, os.Stdout) }
