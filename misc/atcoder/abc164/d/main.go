package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	var s []byte
	Fscan(in, &s)
	cnts := [2019]int{0: 1}
	ans, v := 0, 0
	for _, b := range s {
		v = (v*10 + int(b-'0')) % 2019
		ans += cnts[v]
		cnts[v]++
	}
	Fprint(_w, ans)
}

func main() { run(os.Stdin, os.Stdout) }
