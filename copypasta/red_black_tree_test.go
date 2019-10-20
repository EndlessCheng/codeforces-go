package copypasta

import (
	"fmt"
	"testing"
)

func Test_rbTree(t_ *testing.T) {
	// 使用 https://www.luogu.org/problem/P3369 来测试

	t := newRBTree()
	for i := 1; i < 100; i++ {
		t.put(rbKeyType(i), 1)
		fmt.Println(t)
	}
}
