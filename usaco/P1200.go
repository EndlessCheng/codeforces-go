package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func SolP1200(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var s, t string
	Fscan(in, &s, &t)
	ms, mt := 1, 1
	for _, c := range s {
		ms *= int(c - 'A' + 1)
	}
	for _, c := range t {
		mt *= int(c - 'A' + 1)
	}
	if (ms-mt)%47 == 0 {
		Fprintln(out, "GO")
	} else {
		Fprintln(out, "STAY")
	}
}

//func main() {
//	SolP1200(os.Stdin, os.Stdout)
//}
