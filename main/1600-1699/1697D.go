package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
	"strconv"
)

// https://space.bilibili.com/206214
type interaction97 interface {
	readInitData() initData97
	query(request97) response97
	printAnswer(answer97)
}

type stdIO97 struct {
	in  *bufio.Reader
	out *bufio.Writer
}

type (
	initData97 struct{ n int }
	request97  struct{ q []int }
	response97 struct{ res string }
	answer97   struct{ ans string }
)

func (io stdIO97) readInitData() initData97 {
	in := io.in

	var n int
	Fscan(in, &n)

	return initData97{n}
}

func (io stdIO97) query(q request97) (resp response97) {
	in, out := io.in, io.out

	Fprint(out, "?")
	for _, v := range q.q {
		Fprint(out, " ", v)
	}
	Fprintln(out)

	out.Flush()

	Fscan(in, &resp.res)

	if resp.res == "0" {
		panic(-1)
	}
	return
}

func (io stdIO97) printAnswer(a answer97) {
	out := io.out

	Fprintln(out, "!", a.ans)

	out.Flush()
}

func doInteraction97(it interaction97) {
	dt := it.readInitData()
	n := dt.n

	getChar := func(i int) byte {
		return it.query(request97{[]int{1, i + 1}}).res[0]
	}
	getDiff := func(l, r int) int {
		v, _ := strconv.Atoi(it.query(request97{[]int{2, l + 1, r + 1}}).res)
		return v
	}

	ans := make([]byte, n)
	defer func() { it.printAnswer(answer97{string(ans)}) }()

	type pair struct {
		i int
		b byte
	}
	pos := []pair{}
	for i := range ans {
		j := sort.Search(len(pos), func(j int) bool { return getDiff(pos[j].i, i) > len(pos)-j }) - 1
		if j < 0 {
			ans[i] = getChar(i)
		} else {
			ans[i] = pos[j].b
			pos = append(pos[:j], pos[j+1:]...)
		}
		pos = append(pos, pair{i, ans[i]})
	}
}

func run97() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	T := 1
	for ; T > 0; T-- {
		doInteraction97(stdIO97{in, out})
	}
}

//func main() { run97() }
