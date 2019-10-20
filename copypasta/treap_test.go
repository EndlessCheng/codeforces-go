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

	fmt.Println("multi put")
	t.put(3,1)
	fmt.Println(t)

	// must have 3
	fmt.Println("first delete")
	t.put(3,-1)

	fmt.Println(t)
	fmt.Println("second delete")
	t.put(3,-1)
	fmt.Println(t)

}
