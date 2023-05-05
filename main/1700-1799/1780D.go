package main

import (
	"bufio"
	. "fmt"
	"math/bits"
	"os"
)

// https://space.bilibili.com/206214
type interaction80 interface {
	readInitData() initData80
	query(request80) response80
	printAnswer(answer80)
}

type stdIO80 struct {
	in  *bufio.Reader
	out *bufio.Writer
}

type (
	initData80 struct{ n int }
	request80  struct{ q int }
	response80 struct{ res int }
	answer80   struct{ ans int }
)

func (io stdIO80) readInitData() initData80 {
	in := io.in

	var n int
	Fscan(in, &n)
	return initData80{n}
}

func (io stdIO80) query(q request80) (resp response80) {
	in, out := io.in, io.out

	Fprintln(out, "-", q.q)
	out.Flush()

	Fscan(in, &resp.res)

	if resp.res < 0 {
		panic(-1)
	}
	return
}

func (io stdIO80) printAnswer(a answer80) {
	out := io.out

	Fprintln(out, "!", a.ans)

	out.Flush()
}

func doInteraction80(it interaction80) {
	dt := it.readInitData()
	init1 := dt.n

	q := func(q int) int {
		return it.query(request80{q}).res
	}

	var ans int
	defer func() { it.printAnswer(answer80{ans}) }()

	x := 1
	pre := init1
	for i := 0; i < init1; i++ {
		res := q(x)
		if res >= pre {
			x = 1 << (res - pre + bits.Len(uint(x)))
		}
		ans |= x
		pre = res
	}
}

func run80() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	T := 1
	Fscan(in, &T)
	for ; T > 0; T-- {
		doInteraction80(stdIO80{in, out})
	}
}

//func main() { run80() }
