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
	ans := make([]int, n) // 0
	defer func() { gs.ans = ans }()



	return
}

// TODO: 提交前运行下，检查输出格式是否正确
func ioq() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	// Interaction
	Q := func(req req) (resp resp) {
		Fprintln(out, "?", req.i)
		out.Flush()
		Fscan(in, &resp.v)
		return
	}

	var T int
	for Fscan(in, &T); T > 0; T-- { // TODO: remove if not multi-cases
		// Input
		d := input{}
		Fscan(in, &d.n)

		// Output
		gs := run(d, Q)
		ans := gs.ans
		Fprint(out, "! ")
		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out)
		out.Flush()

		// Optional
		var res int
		if Fscan(in, &res); res < 0 {
			panic(-1)
		}
	}
}

func main() { ioq() }
