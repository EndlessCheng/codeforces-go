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
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		t := 0
		for i := 0; i < n; i++ {
			Fscan(in, &l, &r)
			if t < l {
				t = l
			}
			if t > r { // 前面还有人
				Fprint(out, "0 ")
			} else {
				Fprint(out, t, " ")
				t++ // 打饭需要花费一分钟时间
			}
		}
		Fprintln(out)
	}
}

func main() { run(os.Stdin, os.Stdout) }
