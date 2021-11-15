package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1043C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s []byte
	Fscan(in, &s)
	// 另一种做法是在 s 末尾加上 'b' 从而只用判断相邻字符
	for i, b := range s {
		if i < len(s)-1 && b != s[i+1] || i == len(s)-1 && b == 'a' {
			Fprint(out, 1, " ")
		} else {
			Fprint(out, 0, " ")
		}
	}
}

//func main() { CF1043C(os.Stdin, os.Stdout) }
