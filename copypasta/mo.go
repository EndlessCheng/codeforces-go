package copypasta

import (
	. "fmt"
	"io"
	"math"
	"slices"
	"sort"
)

// 有关根号分解的说明另见 sqrt_decomposition.go

// 莫队知乎：https://www.zhihu.com/people/mythly/answers

// 莫队算法：对询问分块
// todo 各类莫队综述
//    https://www.luogu.com.cn/blog/DPair2005/talk-talk-bruh-forces-in-oi
//    https://www.cnblogs.com/WAMonster/p/10118934.html
//    https://ouuan.github.io/post/%E8%8E%AB%E9%98%9F%E5%B8%A6%E4%BF%AE%E8%8E%AB%E9%98%9F%E6%A0%91%E4%B8%8A%E8%8E%AB%E9%98%9F%E8%AF%A6%E8%A7%A3/
//    https://blog.csdn.net/weixin_43914593/article/details/108485396
// todo【推荐】文章及题单 https://www.luogu.com.cn/training/2914
// todo 题单 https://www.luogu.com.cn/training/73984
// https://cp-algorithms.com/data_structures/sqrt_decomposition.html#toc-tgt-8

// 普通莫队（没有修改操作）
// 视频讲解 https://www.bilibili.com/video/BV1p3h3zYEbc/
// 文字讲解 https://leetcode.cn/problems/threshold-majority-queries/solutions/3740919/mo-dui-suan-fa-hui-gun-mo-dui-pythonjava-x7yw/
// 本质是通过巧妙地改变回答询问的顺序，使区间左右端点移动的次数由 O(nm) 降至 O(n√m)，其中 n 是数组长度，m 是询问个数
// 对于每个块，右端点在 [1,n] 中一直向右或者一直向左，而左端点只在块内「抖动」
// 对于每个块，右端点的最坏移动次数是 n，平均移动次数是 n/2，总移动次数是 O(n * 块个数) = O(n^2 / 块大小)，系数为 1/2 ~ 1
// 对于每个询问，左端点移动次数是 O(块大小)，总移动次数是 O(m * 块大小)
// n^2 / 块大小 = m * 块大小 => 块大小取 n/√m 时，总移动次数最优，为 O(n√m)。
// 在随机数据下，块大小取 n/√(2m) 更好。见 https://zhuanlan.zhihu.com/p/1920472309522740969
// 注 1：如果块大小取 √n，那么移动次数约为 (n+m)√n >= 2n√m，当且仅当 n=m 时取等号（基本不等式），其中不等式右侧为块大小取 n/√m 时的移动次数
// 注 2：为防止块大小为 0，代码中要取 ceil(n/√m)
//
// https://oi-wiki.org/misc/mo-algo/
// 模板题 https://www.luogu.com.cn/problem/P1494
// https://www.luogu.com.cn/problem/P2709
// https://www.luogu.com.cn/problem/P4462
// 恰好出现两次 https://www.luogu.com.cn/problem/P7764
// https://www.luogu.com.cn/problem/P5673
// https://ac.nowcoder.com/acm/problem/25458
// 至少出现两次 https://ac.nowcoder.com/acm/problem/20545
// 至少出现 k 次 https://codeforces.com/problemset/problem/375/D
// 至少出现 k 次 https://www.codechef.com/problems/KCHIPS
// https://codeforces.com/contest/220/problem/B
// https://atcoder.jp/contests/abc242/tasks/abc242_g
// https://atcoder.jp/contests/abc293/tasks/abc293_g
// 区间 mex https://blog.csdn.net/includelhc/article/details/79593496
//     反向构造题 https://www.luogu.com.cn/problem/P6852
// https://codeforces.com/contest/86/problem/D
// https://codeforces.com/problemset/problem/617/E 2200
// https://codeforces.com/contest/877/problem/F
// https://atcoder.jp/contests/abc405/tasks/abc405_g 莫队 + 分块
// - 相似题目 https://leetcode.cn/problems/kth-smallest-path-xor-sum/
// https://www.codechef.com/problems/QCHEF
func normalMo(a []int, queries [][]int) []int {
	n := len(a)
	nq := len(queries)
	// 注：由于每个块右端点平均总移动 n^/(2B) 次，左端点总移动次数为 2qB，两个 2 抵消了
	blockSize := int(math.Ceil(float64(n) / math.Sqrt(float64(nq))))
	type moQuery struct{ bid, l, r, qid int } // [l,r)
	qs := make([]moQuery, nq)
	for i, q := range queries {
		// 输入是从 1 开始的
		l, r := q[0], q[1] // read...
		qs[i] = moQuery{l / blockSize, l, r + 1, i}
	}
	slices.SortFunc(qs, func(a, b moQuery) int {
		if a.bid != b.bid {
			return a.bid - b.bid
		}
		// 奇偶化排序
		if a.bid%2 == 0 {
			return a.r - b.r
		}
		return b.r - a.r
	})

	cnt := 0
	l, r := 1, 1 // 区间从 1 开始，方便 debug
	move := func(idx, delta int) {
		// NOTE: 有些题目在 delta 为 1 和 -1 时逻辑的顺序是严格对称的
		// v := a[idx-1]
		// ...
		// cnt += delta
		if delta > 0 {
			cnt++
		} else {
			cnt--
		}
	}
	getAns := func(q moQuery) int {
		// 提醒：q.r 是加一后的，计算时需要注意
		// sz := q.r - q.l
		// ...
		return cnt
	}
	ans := make([]int, len(qs))
	for _, q := range qs {
		for ; l < q.l; l++ {
			move(l, -1)
		}
		for l > q.l {
			l--
			move(l, 1)
		}
		for ; r < q.r; r++ {
			move(r, 1)
		}
		for r > q.r {
			r--
			move(r, -1)
		}
		ans[q.qid] = getAns(q)
	}
	return ans
}

// 带修莫队（支持单点修改）
// https://oi-wiki.org/misc/modifiable-mo-algo/
// https://codeforces.com/blog/entry/72690
// 模板题 数颜色 https://www.luogu.com.cn/problem/P1903
// https://codeforces.com/problemset/problem/940/F
// https://codeforces.com/problemset/problem/1476/G
// todo https://www.codechef.com/FEB17/problems/DISTNUM3
func moWithUpdate(in io.Reader) []int {
	var n, q int
	Fscan(in, &n, &q)
	a := make([]int, n+1) // 从 1 开始，方便 debug
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	blockSize := int(math.Round(math.Pow(float64(n), 2.0/3)))
	type query struct{ lb, rb, l, r, t, qid int }
	type modify struct{ pos, val int }
	qs := []query{}
	ms := []modify{}
	for ; q > 0; q-- {
		var op string
		if Fscan(in, &op); op[0] == 'Q' {
			var l, r int
			Fscan(in, &l, &r)
			// 改成左闭右开
			qs = append(qs, query{l / blockSize, (r + 1) / blockSize, l, r + 1, len(ms), len(qs)})
		} else {
			var pos, val int
			Fscan(in, &pos, &val)
			ms = append(ms, modify{pos, val})
		}
	}
	sort.Slice(qs, func(i, j int) bool {
		a, b := qs[i], qs[j]
		if a.lb != b.lb {
			return a.lb < b.lb
		}
		if a.rb != b.rb {
			if a.lb&1 == 0 {
				return a.rb < b.rb
			}
			return a.rb > b.rb
		}
		if a.rb&1 == 0 {
			return a.t < b.t
		}
		return a.t > b.t
	})

	const mx int = 1e6 // TODO
	cnt, cc := [mx + 1]int{}, 0
	l, r, now := 1, 1, 0
	add := func(val int) {
		if cnt[val] == 0 {
			cc++
		}
		cnt[val]++
	}
	del := func(val int) {
		cnt[val]--
		if cnt[val] == 0 {
			cc--
		}
	}
	// 注：由于函数套函数不会内联，直接写到主流程的 for now 循环中会快不少
	timeSlip := func(l, r int) {
		m := ms[now]
		p, v := m.pos, m.val
		if l <= p && p < r {
			del(a[p])
			add(v)
		}
		a[p], ms[now].val = v, a[p]
	}
	getAns := func(q query) int {
		// 提醒：q.r 是加一后的，计算时需要注意
		// sz := q.r - q.l
		// ...
		return cc
	}
	ans := make([]int, len(qs))
	for _, q := range qs {
		for ; r < q.r; r++ {
			add(a[r])
		}
		for ; l < q.l; l++ {
			del(a[l])
		}
		for l > q.l {
			l--
			add(a[l])
		}
		for r > q.r {
			r--
			del(a[r])
		}
		for ; now < q.t; now++ {
			timeSlip(q.l, q.r)
		}
		for now > q.t {
			now--
			timeSlip(q.l, q.r)
		}
		ans[q.qid] = getAns(q)
	}
	return ans
}

// 回滚莫队
// 复杂度同普通莫队
// 视频讲解 https://www.bilibili.com/video/BV1p3h3zYEbc/
// 文字讲解 https://leetcode.cn/problems/threshold-majority-queries/solutions/3740919/mo-dui-suan-fa-hui-gun-mo-dui-pythonjava-x7yw/
// https://oi-wiki.org/misc/rollback-mo-algo/
// 浅谈回滚莫队 https://www.luogu.com.cn/blog/bfqaq/qian-tan-hui-gun-mu-dui
// todo 回滚莫队及其简单运用 https://www.cnblogs.com/Parsnip/p/10969989.html
// 模板题 https://leetcode.cn/problems/threshold-majority-queries/
// - https://ac.nowcoder.com/acm/contest/103151/F
// 历史研究 https://www.luogu.com.cn/problem/AT1219 https://atcoder.jp/contests/joisc2014/tasks/joisc2014_c
// https://loj.ac/p/6285
// todo https://www.luogu.com.cn/problem/P5906
// todo https://www.luogu.com.cn/problem/P5386
// todo https://www.luogu.com.cn/problem/P6072
func moWithRollback(in io.Reader) []int {
	var n, q int
	Fscan(in, &n, &q)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}

	ans := make([]int, q)
	blockSize := int(math.Ceil(float64(n) / math.Sqrt(float64(q))))
	type query struct{ bid, l, r, qid int } // [l,r)
	qs := []query{}
	cnt := make([]int, n+1)
	for i := range ans {
		var l, r int
		Fscan(in, &l, &r)
		r++ // 左闭右开
		// 大区间离线
		if r-l > blockSize {
			qs = append(qs, query{l / blockSize, l, r, i})
			continue
		}
		// 小区间暴力
		res := 0
		for _, v := range a[l:r] {
			cnt[v]++
			// ...
		}
		ans[i] = res
		// 重置数据 ...
		for _, v := range a[l:r] {
			cnt[v] = 0
		}
	}
	slices.SortFunc(qs, func(a, b query) int {
		if a.bid != b.bid {
			return a.bid - b.bid
		}
		return a.r - b.r
	})

	var res int
	add := func(i int) {
		v := a[i]
		cnt[v]++
		// ...
	}
	getAns := func(q query) int {
		// ...
		return res
	}

	l, r := 0, 0
	for i, q := range qs {
		l0 := (q.bid + 1) * blockSize
		if i == 0 || q.bid > qs[i-1].bid {
			l, r = l0, l0
			// 重置数据 ...
			res = 0
			clear(cnt)
		}

		for ; r < q.r; r++ {
			add(r)
		}
		tmp := res // 其他的如最值等，也在这里 tmp 记录一下
		// 由于下面回滚了，每次 l 都是从 l0 开始的
		// 遍历区间 [q.l, l0)
		for l > q.l {
			l--
			add(l)
		}
		ans[q.qid] = getAns(q)
		res = tmp // 同理，恢复其他记录的值

		// 回滚，始终保持 l 在 l0 的位置
		for ; l < l0; l++ {
			// 回滚 ...
			cnt[a[l]]--
		}
	}
	return ans
}

// 树上莫队
// 通过 DFS 序转化成序列上的查询
// https://oi-wiki.org/misc/mo-algo-on-tree/
// 有关树分块的内容见 graph_tree.go 中的 limitSizeDecomposition
// NOTE: 对于带修莫队，去掉 timeSlip 中的参数，且 if l <= p && p < r 替换成 if vis[p] https://www.luogu.com.cn/record/46714923
// 模板题 糖果公园 https://www.luogu.com.cn/problem/P4074
//       https://www.acwing.com/problem/content/2536/ https://www.luogu.com.cn/problem/SP10707
// https://leetcode.cn/problems/minimum-edge-weight-equilibrium-queries-in-a-tree/
func moOnTree(n, root, q int, g [][]int, vals []int) []int {
	vs := make([]int, 0, 2*n)
	tin := make([]int, n)
	tout := make([]int, n)
	var initTime func(v, fa int)
	initTime = func(v, fa int) {
		tin[v] = len(vs)
		vs = append(vs, v)
		for _, w := range g[v] {
			if w != fa {
				initTime(w, v)
			}
		}
		tout[v] = len(vs)
		vs = append(vs, v)
	}
	initTime(root, -1)

	// initTime 的逻辑可以合并到求 pa dep 的 DFS 中
	var getLCA func(int, int) int // 见 graph_tree.go 中的 lcaBinarySearch

	blockSize := int(math.Ceil(float64(2*n) / math.Sqrt(float64(q)))) // int(math.Round(math.Pow(float64(2*n), 2.0/3)))
	type query struct{ lb, l, r, lca, qid int }
	qs := make([]query, q)
	for i := range qs {
		var v, w int
		//Fscan(in, &v, &w)
		v--
		w--
		if tin[v] > tin[w] {
			v, w = w, v
		}
		if lca := getLCA(v, w); lca != v {
			qs[i] = query{tout[v] / blockSize, tout[v], tin[w] + 1, lca, i}
		} else {
			qs[i] = query{tin[v] / blockSize, tin[v], tin[w] + 1, -1, i}
		}
	}
	sort.Slice(qs, func(i, j int) bool {
		a, b := qs[i], qs[j]
		if a.lb != b.lb {
			return a.lb < b.lb
		}
		if a.lb&1 == 0 {
			return a.r < b.r
		}
		return a.r > b.r
	})

	var k int // vals 不同元素个数
	cnt := make([]int, k+1)
	cc := 0
	l, r := 0, 0
	vis := make([]bool, n)
	move := func(v int) {
		x := vals[v]
		vis[v] = !vis[v]
		if vis[v] {
			if cnt[x] == 0 {
				cc++
			}
			cnt[x]++
		} else {
			cnt[x]--
			if cnt[x] == 0 {
				cc--
			}
		}
	}
	getAns := func(q query) int {
		return cc
	}
	ans := make([]int, q)
	for _, q := range qs {
		for ; r < q.r; r++ {
			move(vs[r])
		}
		for ; l < q.l; l++ {
			move(vs[l])
		}
		for l > q.l {
			l--
			move(vs[l])
		}
		for r > q.r {
			r--
			move(vs[r])
		}
		if q.lca >= 0 {
			move(q.lca)
		}
		ans[q.qid] = getAns(q)
		if q.lca >= 0 {
			move(q.lca)
		}
	}
	return ans
}

// 二次离线莫队
// 1. 用莫队把询问分块（第一次离线）
// 2. 把莫队左右指针的移动记录下来（第二次离线）
// 3. 用（另一种）离线算法，更高效地处理左右指针的移动
// https://oi-wiki.org/misc/mo-algo-secondary-offline/
// https://www.cnblogs.com/Nero-Claudius/p/MoQueue1.html
// todo https://www.luogu.com.cn/blog/gxy001/mu-dui-er-ci-li-xian
//  https://kewth.github.io/2019/10/16/%E8%8E%AB%E9%98%9F%E4%BA%8C%E6%AC%A1%E7%A6%BB%E7%BA%BF/
//  静态区间逆序对 https://www.luogu.com.cn/problem/P5047
//  https://www.luogu.com.cn/problem/P4887
//  https://www.luogu.com.cn/problem/P5501
//  https://www.luogu.com.cn/problem/P5398
