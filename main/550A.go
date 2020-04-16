package main

import (
	. "fmt"
	"io"
	. "strings"
)

// github.com/EndlessCheng/codeforces-go
func CF550A(r io.Reader, w io.Writer) {
	var s string
	Fscan(r, &s)
	l1, l2, r1, r2 := Index(s, "AB"), Index(s, "BA"), LastIndex(s, "AB"), LastIndex(s, "BA")
	if l1 != -1 && l1+1 < r2 || l2 != -1 && l2+1 < r1 {
		Fprint(w, "YES")
	} else {
		Fprint(w, "NO")
	}
}

//func main() { CF550A(os.Stdin, os.Stdout) }
