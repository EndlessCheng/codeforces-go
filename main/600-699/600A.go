package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
	"unicode"
)

// https://space.bilibili.com/206214
func isDigit(s string) int {
	if len(s) == 0 || len(s) > 1 && s[0] == '0' {
		return 1
	}
	for _, b := range s {
		if !unicode.IsDigit(b) {
			return 1
		}
	}
	return 0
}

func CF600A(in io.Reader, out io.Writer) {
	var s string
	Fscan(bufio.NewReader(in), &s)
	a := [2][]string{}
	for _, s := range strings.Split(strings.ReplaceAll(s, ";", ","), ",") {
		v := isDigit(s)
		a[v] = append(a[v], s)
	}
	for _, b := range a {
		if b == nil {
			Fprintln(out, "-")
		} else {
			Fprintf(out, "\"%s\"\n", strings.Join(b, ","))
		}
	}
}

//func main() { CF600A(os.Stdin, os.Stdout) }
