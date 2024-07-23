package main

import (
	. "fmt"
	"io"
)

func cf1037C(in io.Reader, out io.Writer) {
	var s, t []byte
	Fscan(in, &s, &s, &t)
	ans := 0
	preI := -2
	for i, b := range s {
		if b == t[i] {
			continue
		}
		if preI == i-1 && b != s[preI] {
			preI = -2
		} else {
			ans++
			preI = i
		}
	}
	Fprint(out, ans)
}

//func main() { cf1037C(bufio.NewReader(os.Stdin), os.Stdout) }
