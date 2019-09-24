package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

func Sol494A(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var s string
	Fscan(in, &s)
	cnt := strings.Count(s, "#")
	s = strings.Replace(s, "#", ")", cnt-1)
	need := strings.Count(s, "(") - strings.Count(s, ")")
	if need < 1 {
		need = 1
	}
	s = strings.Replace(s, "#", strings.Repeat(")", need), 1)
	beauty := 0
	for _, c := range s {
		if c == '(' {
			beauty++
		} else {
			beauty--
			if beauty < 0 {
				Fprintln(out, -1)
				return
			}
		}
	}
	Fprint(out, strings.Repeat("1\n", cnt-1))
	Fprintln(out, need)
}

//func main() {
//	Sol494A(os.Stdin, os.Stdout)
//}
