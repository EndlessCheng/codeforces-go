package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://space.bilibili.com/206214
type interaction interface {
	readInitData() initData
	query(request) response
	printAnswer(answer)
}

type stdIO struct {
	in  *bufio.Reader
	out *bufio.Writer
}

type (
	initData struct{ n int }
	request  struct{ q int } //
	response struct{ res int }
	answer   struct{ ans int }
)

func (io stdIO) readInitData() initData {
	in := io.in

	var n int
	Fscan(in, &n)
	// TODO 初始输入格式？

	return initData{n}
}

func (io stdIO) query(q request) (resp response) {
	in, out := io.in, io.out

	Fprintln(out, "?", q.q)
	//Fprint(out, "?")
	//Fprint(out, " ", len(q.q)) // TODO 输出 query 长度？
	//for _, v := range q.q {
	//	Fprint(out, " ", v)
	//}
	//Fprintln(out)

	out.Flush()

	Fscan(in, &resp.res)

	// TODO: 题目定义了 incorrect request？
	if resp.res < 0 {
		panic(-1)
	}
	return
}

func (io stdIO) printAnswer(a answer) {
	out := io.out

	Fprintln(out, "!", a.ans)
	//Fprint(out, "!")
	//Fprint(out, " ", len(a.ans)) // TODO 输出答案长度？
	//for _, v := range a.ans {
	//	Fprint(out, " ", v)
	//}
	//Fprintln(out)

	out.Flush()

	// TODO judge 是否返回答案非法？（通常是 move on to the next test case）
	//var res int
	//if Fscan(io.in, &res); res < 0 {
	//	panic(res)
	//}
}

func doInteraction(it interaction) {
	// TODO 初始输入格式？
	dt := it.readInitData()
	n := dt.n
	_ = n

	// TODO query 格式？
	q := func(q int) int {
		//for i := range q {
		//	q[i]++
		//}
		return it.query(request{q}).res
	}
	_ = q

	// TODO 答案类型？
	var ans int
	//ans := make([]int, n)
	defer func() { it.printAnswer(answer{ans}) }()

	// TODO: 在这里实现

}

func run() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	T := 1
	// TODO：多测？
	Fscan(in, &T)
	for ; T > 0; T-- {
		doInteraction(stdIO{in, out})
	}
}

// TODO: 运行 & 测试！检查格式是否正确
func main() { run() }
