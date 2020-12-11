package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	defer func() {
		leftData, _ := ioutil.ReadAll(in)
		s := strings.TrimSpace(string(leftData))
		if s != "" {
			panic("有未读入的数据：\n" + s)
		}
	}()

	solve := func(Case int) {
		var n int
		Fscan(in, &n)

	}

	var t int
	Fscan(in, &t)
	for Case := 1; Case <= t; Case++ {
		Fprintf(out, "Case #%d: ", Case)
		solve(Case)
	}
}

func main() { run(os.Stdin, os.Stdout) }
