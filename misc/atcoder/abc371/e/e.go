package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, v, ans, sumG int
	Fscan(in, &n)
	last := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		sumG += i - last[v]
		ans += sumG
		last[v] = i
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
