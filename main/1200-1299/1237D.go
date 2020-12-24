package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type mqData37 struct {
	v   int
	del int
}
type mq37 struct {
	data []mqData37
	size int
}

func (q mq37) less(a, b mqData37) bool { return a.v >= b.v }
func (q *mq37) push(v int) {
	q.size++
	d := mqData37{v, 1}
	for len(q.data) > 0 && q.less(d, q.data[len(q.data)-1]) {
		d.del += q.data[len(q.data)-1].del
		q.data = q.data[:len(q.data)-1]
	}
	q.data = append(q.data, d)
}
func (q *mq37) pop() {
	q.size--
	if q.data[0].del > 1 {
		q.data[0].del--
	} else {
		q.data = q.data[1:]
	}
}
func (q mq37) top() (v int) { return q.data[0].v }

func CF1237D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n, 3*n)
	for i := range a {
		Fscan(in, &a[i])
	}
	a = append(append(a, a...), a...)
	q := mq37{}
	for i, j := 0, 0; i < n; i++ {
		for ; j < 3*n && (q.size == 0 || 2*a[j] >= q.top()); j++ {
			q.push(a[j])
		}
		sz := j - i
		if sz > 2*n {
			sz = -1
		}
		Fprint(out, sz, " ")
		if q.size == sz {
			q.pop()
		}
	}
}

//func main() { CF1237D(os.Stdin, os.Stdout) }
