package main

import (
	"bufio"
	. "fmt"
	"io"
)

func do559B(s string) string {
	if len(s)&1 == 1 {
		return s
	}
	s1 := do559B(s[:len(s)>>1])
	s2 := do559B(s[len(s)>>1:])
	if s1 < s2 {
		return s1 + s2
	}
	return s2 + s1
}

// github.com/EndlessCheng/codeforces-go
func Sol559B(reader io.Reader, writer io.Writer) {
	ifElse := func(cond bool, r1, r2 string) string {
		if cond {
			return r1
		}
		return r2
	}

	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var s, t string
	Fscan(in, &s, &t)
	Fprint(out, ifElse(do559B(s) == do559B(t), "YES", "NO"))
}

//func main() {
//	Sol559B(os.Stdin, os.Stdout)
//}
