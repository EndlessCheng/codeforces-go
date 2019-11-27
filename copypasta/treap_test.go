package copypasta

import (
	"fmt"
	"testing"
)

func Test_treap(t_ *testing.T) {
	// 使用 https://www.luogu.org/problem/P3369 来测试

	t := newTreap()
	for i := 1; i <= 20; i++ {
		t.put(tpKeyType(i), 1)
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
