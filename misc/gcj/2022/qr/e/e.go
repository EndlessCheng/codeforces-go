package main

import (
	"bufio"
	. "fmt"
	"math"
	"math/rand"
	"os"
)

// github.com/EndlessCheng/codeforces-go
type (
	initData struct{ n int }
	request  struct {
		s string
		q int
	}
	response struct{ v, e int }
	answer   struct{ ans int }
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
	Fscan(io.in, &d.n, new(int), new(int), new(int))
	if d.n < 0 {
		panic(-1)
	}
	return
}

func (io io) query(q request) (resp response) {
	Fprint(io.out, q.s)
	if q.s == "T" {
		Fprint(io.out, " ", q.q)
	}
	Fprintln(io.out)
	io.out.Flush()

	Fscan(io.in, &resp.v, &resp.e)
	if resp.v < 0 {
		panic(-1)
	}
	return
}

func (io io) printAnswer(a answer) {
	Fprintln(io.out, "E", a.ans)
	io.out.Flush()
}

const K = 8000

func doInteraction(it interaction) {
	dt := it.readInitData()
	n := dt.n

	q := func(s string, q int) (int, int) {
		resp := it.query(request{s, q})
		return resp.v, resp.e
	}

	var ans int
	defer func() { it.printAnswer(answer{ans}) }()

	if n <= K {
		for i := 1; i <= n; i++ {
			_, e := q("T", i)
			ans += e
		}
		ans /= 2
		return
	}

	es := map[int]int{}
	sumT, sumTW := 0, 0
	for i := 0; i < K/2; i++ {
		v, e := q("T", rand.Intn(n)+1)
		sumT += e
		sumTW += e - es[v]
		es[v] = e
		v, e = q("W", 0)
		sumTW += e - es[v]
		es[v] = e
	}
	avg := float64(sumT) / float64(K/2)
	ans = int(math.Round(float64(sumTW)+avg*float64(n-len(es))) / 2)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	var T int
	for Fscan(in, &T); T > 0; T-- {
		doInteraction(io{in, out})
	}
}
