package main

import (
	"bufio"
	. "fmt"
	"os"
)

type (
	input struct{ n int }
	guess struct{ ans int }
	qIn   struct{ q int }
	qOut  struct{ ok bool }
)

// github.com/EndlessCheng/codeforces-go
func run(in input, _Q func(qIn) qOut) (gs guess) {
	Q := func(q int) qOut { return _Q(qIn{q}) }
	n := in.n

	return
}

func ioq() {
	in := bufio.NewReader(os.Stdin)
	// if the number of output times is small, just use Println without bufio things
	out := bufio.NewWriter(os.Stdout)
	Q := func(qi qIn) (resp qOut) {
		Fprintln(out, "?", qi.q)
		out.Flush()
		// ... or read int and return it
		var s []byte
		Fscan(in, &s)
		resp.ok = s[0] == 'Y'
		return
	}
	var t int
	for Fscan(in, &t); t > 0; t-- {
		d := input{}
		Fscan(in, &d.n)
		gs := run(d, Q)
		Fprintln(out, "!", gs.ans)
		out.Flush()
		// some problems need to read an extra string like "Correct" or "Incorrect" after guessed the answer
		//var s []byte
		//Fscan(in, &s)
	}
}

func main() { ioq() }
