package main

import (
	"bufio"
	. "fmt"
	"os"
)

// github.com/EndlessCheng/codeforces-go
type (
	initData struct{ n int }
	request  struct{ q []int }
	response struct{ v int }
	answer   struct{ ans []int }
)

type interaction interface {
	readInitData() initData
	query(request) response
	printAnswer(answer)
}

type io struct {
	in  *bufio.Reader
	out *bufio.Writer
}

func (io io) readInitData() (d initData) {
	Fscan(io.in, &d.n)
	return
}

func (io io) query(q request) (resp response) {
	Fprint(io.out, "?")
	for _, v := range q.q {
		Fprint(io.out, " ", v)
	}
	Fprintln(io.out)
	io.out.Flush()

	Fscan(io.in, &resp.v)
	if resp.v < 0 {
		panic(-1)
	}
	return
}

func (io io) printAnswer(a answer) {
	Fprint(io.out, "!")
	//Fprint(io.out, " ", len(a.ans))
	for _, v := range a.ans {
		Fprint(io.out, " ", v)
	}
	Fprintln(io.out)
	io.out.Flush()

	// TODO: optional, panic if we got -1
	var res int
	if Fscan(io.in, &res); res < 0 {
		panic(res)
	}
}

func doInteraction(it interaction) {
	dt := it.readInitData()
	n := dt.n

	q := func(q ...int) int {
		//for i := range q {
		//	q[i]++
		//}
		return it.query(request{q}).v
	}

	var ans []int
	ans = make([]int, n) //
	defer func() { it.printAnswer(answer{ans}) }()


}

// TODO: check output format before submit
func run() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	T := 1
	Fscan(in, &T) //
	for ; T > 0; T-- {
		doInteraction(io{in, out})
	}
}

func main() { run() }
