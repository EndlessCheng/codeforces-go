package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(reader io.Reader, writer io.Writer) {
	// 别忘记取模！
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)

	var t int
	Fscan(in, &t)
	for case_ := 0; case_ < t; case_++ {
		solveCase(in, out)
	}
}

func solveCase(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)

}

func main() {
	solve(os.Stdin, os.Stdout)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func ifElseI(cond bool, a, b int) int {
	if cond {
		return a
	}
	return b
}
func ifElseS(cond bool, a, b string) string {
	if cond {
		return a
	}
	return b
}
