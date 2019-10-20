package copypasta

import (
	"fmt"
	"testing"
)

func Test_treap(t_ *testing.T) {
	// 使用 https://www.luogu.org/problem/P3369 来测试

	seed = uint(1)
	t := newTreap()
	for i := 1; i < 20; i++ {
		t.put(tKeyType(i), 1)
		fmt.Println(t)
	}
}
