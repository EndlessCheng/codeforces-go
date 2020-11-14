package copypasta

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func buildSplay(l, r int) *spNode {
	if l == r {
		return nil
	}
	m := (l + r) >> 1
	o := &spNode{key: spKeyType(m)}
	o.lr[0] = buildSplay(l, m)
	o.lr[1] = buildSplay(m+1, r)
	o.maintain()
	return o
}

func Test_splay(t *testing.T) {
	//root := buildSplay(1,9)
	//root = root.splay(1)
	rand.Seed(time.Now().UnixNano())
	n := 30
	a := make([]int, n)
	for i := range a {
		a[i] = i + 1
	}
	rand.Shuffle(n, func(i, j int) { a[i], a[j] = a[j], a[i] })
	s := newSplay()
	for i, v := range a {
		s.put(spKeyType(v), 1)
		fmt.Println(i+1, s)
	}
	rand.Shuffle(n, func(i, j int) { a[i], a[j] = a[j], a[i] })
	for i, v := range a {
		s.delete(spKeyType(v))
		fmt.Println(i+1, s)
	}
}
