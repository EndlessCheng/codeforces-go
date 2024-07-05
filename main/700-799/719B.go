package main

import (
	. "fmt"
	"io"
)

func cf719B(in io.Reader, out io.Writer) {
	n, s := 0, ""
	Fscan(in, &n, &s)
	cnt := [2][2]int{}
	for i, b := range s {
		cnt[b>>4&1][i%2]++
	}
	Fprint(out, min(max(cnt[0][0], cnt[1][1]), max(cnt[0][1], cnt[1][0])))
}

//func main() { cf719B(bufio.NewReader(os.Stdin), os.Stdout) }
