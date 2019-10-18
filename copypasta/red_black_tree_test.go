package copypasta

import "testing"

func Test_rbt(t *testing.T) {
	rbt := NewWithIntComparator()
	rbt.Put(1, "a")
	rbt.Put(10, "b")
	rbt.Put(10, "bb")
	rbt.Put(100, "c")

	t.Log(rbt.lookup(10), rbt.lookup(10).Value)
	t.Log(rbt.Floor(9))
	t.Log(rbt.Floor(10))
	t.Log(rbt.Floor(11))

	o, _ := rbt.Floor(11)
	it := rbt.Iterator(o)
	it.Next()
	it.Next()
	t.Log("next:", it.Key())

	t.Log(rbt.Left())
	t.Log(rbt.Right())
}
