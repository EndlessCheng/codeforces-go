package main

import (
	"bufio"
	. "fmt"
	"os"
)

type (
	input struct{ n int }
	req   struct{ i int }
	resp  struct{ v int }
	guess struct{ ans []int }
)

// github.com/EndlessCheng/codeforces-go
func run(in input, Q func(req) resp) (gs guess) {
	q := func(i int) int { return Q(req{i}).v }
	n := in.n
	ans := make([]int, n)
	defer func() { gs.ans = ans }()

	return
}

// TODO: 提交前运行下，检查输出格式是否正确
func ioq() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	Q := func(req req) (resp resp) {
		Fprintln(out, "?", req.i)
		out.Flush()
		Fscan(in, &resp.v)
		return
	}

	var t int
	for Fscan(in, &t); t > 0; t-- { // TODO: remove if not multi-cases
		d := input{}
		Fscan(in, &d.n)

		gs := run(d, Q)
		ans := gs.ans
		Fprint(out, "!")
		for _, v := range ans {
			Fprint(out, " ", v)
		}
		Fprintln(out)
		out.Flush()
		// some problems need to read an extra string like "Correct" or "Incorrect" after guessed the answer
		//var s string
		//Fscan(in, &s)
	}
}

func main() { ioq() }
