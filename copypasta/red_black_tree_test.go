package copypasta

import (
	"fmt"
	"testing"
)

func Test_rbTree(t_ *testing.T) {
	// 使用 https://www.luogu.org/problem/P3369 来测试

	t := newRBTree()
	n := 20
	for i := 1; i <= n; i++ {
		t.put(rbKeyType(i), 1)
		fmt.Println(t)
	}

	fmt.Println("put another 3")
	t.put(3, 1)
	fmt.Println(t)

	fmt.Println("delete 3")
	t.delete(3)
	fmt.Println(t)

	fmt.Println("delete 3")
	t.delete(3)
	fmt.Println(t)
}
