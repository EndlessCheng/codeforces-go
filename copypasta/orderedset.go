package copypasta

import "container/list"

// 哈希表：支持插入和删除，并能够按照插入顺序遍历哈希表中的元素
// 做 https://codeforces.com/contest/962/problem/D 时有感而发
type OrderedSet struct {
	set map[interface{}]*list.Element
	lst *list.List
}

func NewOrderedSet() OrderedSet {
	return NewOrderedSetWithSpace(0)
}

func NewOrderedSetWithSpace(space int) OrderedSet {
	return OrderedSet{make(map[interface{}]*list.Element, space), list.New()}
}

func (s *OrderedSet) Len() int {
	return len(s.set)
}

// 添加元素 v
func (s *OrderedSet) Store(v interface{}) {
	s.set[v] = s.lst.PushBack(v)
}

// 删除元素 v
func (s *OrderedSet) Delete(v interface{}) {
	s.lst.Remove(s.set[v])
	delete(s.set, v)
}

// 判断 v 是否存在
func (s *OrderedSet) Contains(v interface{}) bool {
	_, ok := s.set[v]
	return ok
}

// 按照插入顺序遍历哈希表中的元素
// 当 f 返回 false 时停止遍历
func (s *OrderedSet) Range(f func(key interface{}) bool) {
	for node := s.lst.Front(); node != nil; node = node.Next() {
		if !f(node.Value) {
			break
		}
	}
}
