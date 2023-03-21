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

func (io io) readInitData() initData {
	var n int
	Fscan(io.in, &n)
	// TODO 初始输入

	return initData{n}
}

func (io io) query(q request) (resp response) {
	Fprint(io.out, "?")
	//Fprint(io.out, " ", len(q.q)) // TODO 询问是否需要输出长度？
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
	//Fprint(io.out, " ", len(a.ans)) // TODO 输出最终答案，是否需要输出其长度？
	for _, v := range a.ans {
		Fprint(io.out, " ", v)
	}
	Fprintln(io.out)
	io.out.Flush()

	// TODO 可选，如果题目还提供返回值的话
	//var res int
	//if Fscan(io.in, &res); res < 0 {
	//	panic(res)
	//}
}

func doInteraction(it interaction) {
	// 初始输入
	dt := it.readInitData()
	n := dt.n

	q := func(q ...int) int {
		//for i := range q {
		//	q[i]++
		//}
		return it.query(request{q}).v
	}

	//var ans int
	var ans []int
	ans = make([]int, n)
	defer func() { it.printAnswer(answer{ans}) }()

}

func run() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	T := 1
	Fscan(in, &T) //
	for ; T > 0; T-- {
		doInteraction(io{in, out})
	}
}

// TODO: 运行一下，检查格式是否正确
func main() { run() }
