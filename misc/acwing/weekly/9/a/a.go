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

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fprintln(out, n)
		for i := 1; i <= n; i++ {
			Fprint(out, i, " ") // 将其他所有元素都增加 i 等价于将该元素减小 i
		}
		Fprintln(out)
	}
}

func main() { run(os.Stdin, os.Stdout) }
