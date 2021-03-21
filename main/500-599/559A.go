package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF559A(in io.Reader, out io.Writer) {
	var a, b, c, e int
	Fscan(in, &a, &b, &c, &e, &e)
	Fprint(out, (a+b+c)*(a+b+c)-a*a-c*c-e*e)
}

//func main() { CF559A(os.Stdin, os.Stdout) }
