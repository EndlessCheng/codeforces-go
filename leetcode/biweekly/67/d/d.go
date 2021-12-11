package main

import "github.com/emirpasic/gods/trees/redblacktree"

/* 巧妙利用查询的特殊性

为了高效插入数据，我们可以用平衡树来维护所有景点。

注意到每次 `get` 后仅会将查询次数加一，我们可以直接用一个迭代器（指针）$\textit{cur}$ 指向下次查询需要返回的元素。

- 对于添加操作，如果添加的景点排在当前 $\textit{cur}$ 前面，那么移动 $\textit{cur}$ 至其前一个元素，否则不移动 $\textit{cur}$；
- 对于查询操作，每次查询结束后将 $\textit{cur}$ 移至其下一个元素。

代码实现时，可以在初始时插入一个哨兵元素，从而简化判断逻辑。

*/

// github.com/EndlessCheng/codeforces-go
type pair struct {
	score int
	name  string
}

func compare(x, y interface{}) int {
	a, b := x.(pair), y.(pair)
	if a.score > b.score || a.score == b.score && a.name < b.name {
		return -1
	}
	return 1
}

type SORTracker struct {
	*redblacktree.Tree
	cur redblacktree.Iterator
}

func Constructor() SORTracker {
	root := redblacktree.NewWith(compare)
	root.Put(pair{}, nil) // 哨兵
	return SORTracker{root, root.IteratorAt(root.Left())}
}

func (t *SORTracker) Add(name string, score int) {
	p := pair{score, name}
	t.Put(p, nil)
	if compare(p, t.cur.Key()) < 0 {
		t.cur.Prev() // 移动至前一个元素
	}
}

func (t *SORTracker) Get() string {
	name := t.cur.Key().(pair).name
	t.cur.Next() // 移动至下一个元素
	return name
}
