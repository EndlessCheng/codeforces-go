package copypasta

import (
	"fmt"
	"testing"
)

func Test_treap(t_ *testing.T) {
	//assert := assert.New(t_)

	seed = uint32(1)
	t := newTreap()
	for i := 1; i < 20; i++ {
		t.put(tKeyType(i), 1)
		fmt.Println(t)
	}
	t.delete(3)
	fmt.Println(t)

	//for i := 1; i < 20; i++ {
	//	t.delete(tKeyType(i))
	//	fmt.Println(t)
	//}
}
