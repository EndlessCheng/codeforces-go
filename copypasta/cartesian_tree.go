package copypasta

/* 笛卡尔树 Cartesian tree
https://en.wikipedia.org/wiki/Cartesian_tree
https://oi-wiki.org/ds/cartesian-tree/
todo 一些题目 https://www.luogu.com.cn/blog/AAAbbb123/di-ka-er-shu-xue-xi-bi-ji
另见单调栈 monotone_stack.go

模板题 https://www.luogu.com.cn/problem/P5854 https://judge.yosupo.jp/problem/cartesian_tree
      https://www.luogu.com.cn/problem/P1377
https://codeforces.com/problemset/problem/1777/D
*/

type ctNode struct {
	lr      [2]*ctNode
	id, val int
}

func buildCartesianTree(a []int) *ctNode {
	if len(a) == 0 {
		return nil
	}
	s := []*ctNode{}
	for i, v := range a {
		o := &ctNode{id: i, val: v}
		for len(s) > 0 {
			top := s[len(s)-1]
			if top.val < v {
				top.lr[1] = o
				break
			}
			o.lr[0] = top
			s = s[:len(s)-1]
		}
		s = append(s, o)
	}
	return s[0]
}

// 非指针版，返回每个节点的左右儿子的编号
func buildCartesianTree2(a []int) [][2]int {
	n := len(a)
	lr := make([][2]int, n)
	for i := range lr {
		lr[i] = [2]int{-1, -1}
	}
	s := []int{}
	for i, v := range a {
		for len(s) > 0 {
			topI := s[len(s)-1]
			if a[topI] < v {
				lr[topI][1] = i
				break
			}
			lr[i][0] = topI
			s = s[:len(s)-1]
		}
		s = append(s, i)
	}
	return lr
}
