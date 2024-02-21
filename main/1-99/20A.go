package main

import (
	. "fmt"
	"io"
	"strings"
)

// https://space.bilibili.com/206214
func cf20A(in io.Reader, out io.Writer) {
	s := ""
	Fscan(in, &s)
	sp := strings.FieldsFunc(s, func(r rune) bool { return r == '/' })
	Fprint(out, "/", strings.Join(sp, "/"))
}

//func main() { cf20A(os.Stdin, os.Stdout) }
