package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
type mqData struct {
	v   int64
	del int
}
type mq struct {
	data []mqData
	size int
}

func (q mq) less(a, b mqData) bool { return a.v >= b.v }
func (q *mq) push(v int64) {
	q.size++
	d := mqData{v, 1}
	for len(q.data) > 0 && q.less(d, q.data[len(q.data)-1]) {
		d.del += q.data[len(q.data)-1].del
		q.data = q.data[:len(q.data)-1]
	}
	q.data = append(q.data, d)
}
func (q *mq) pop() {
	q.size--
	if q.data[0].del > 1 {
		q.data[0].del--
	} else {
		q.data = q.data[1:]
	}
}
func (q mq) top() (v int64) { return q.data[0].v }

type mq2 struct {
	data []mqData
	size int
}

func (q mq2) less(a, b mqData) bool { return a.v <= b.v }
func (q *mq2) push(v int64) {
	q.size++
	d := mqData{v, 1}
	for len(q.data) > 0 && q.less(d, q.data[len(q.data)-1]) {
		d.del += q.data[len(q.data)-1].del
		q.data = q.data[:len(q.data)-1]
	}
	q.data = append(q.data, d)
}
func (q *mq2) pop() {
	q.size--
	if q.data[0].del > 1 {
		q.data[0].del--
	} else {
		q.data = q.data[1:]
	}
}
func (q mq2) top() (v int64) { return q.data[0].v }

func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, ans int64
	Fscan(in, &n, &k)
	a := make([]int64, n)
	mx, mi := mq{}, mq2{}
	for i := range a {
		Fscan(in, &a[i])
		mx.push(a[i])
		mi.push(a[i])
		for mx.top()-mi.top() > k {
			mx.pop()
			mi.pop()
		}
		ans += int64(mx.size)
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
