package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()
	const mod int = 1e9 + 7

	var n int
	Fscan(in, &n)

	var t int
	Fscan(in, &t)
	for case_ := 0; case_ < t; case_++ {
		solveCase(in, out)
		//Fprintln(out, ifElseS(solveCase(in, out), "YES", "NO"))
	}
}

func solveCase(in *bufio.Reader, out *bufio.Writer) {
	var n int
	Fscan(in, &n)

}

func main() {
	solve(os.Stdin, os.Stdout)
	//r, _ := os.Open("input.txt")
	//w, _ := os.Create("output.txt")
	//solve(r, w)
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
