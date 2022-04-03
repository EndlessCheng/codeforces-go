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

	solve := func(curCase int) {
		var n int
		Fscan(in, &n)

	}

	var cases int
	Fscan(in, &cases)
	for curCase := 1; curCase <= cases; curCase++ {
		Fprintf(out, "Case #%d: ", curCase)
		solve(curCase)
	}

	leftData, _ := ioutil.ReadAll(in)
	if s := strings.TrimSpace(string(leftData)); s != "" {
		panic("有未读入的数据：\n" + s)
	}

}

func main() { run(os.Stdin, os.Stdout) }
