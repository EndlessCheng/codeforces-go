package copypasta

import . "fmt"

/* 博弈论 Game Theory
https://en.wikipedia.org/wiki/Game_theory
定义必胜状态为先手必胜的状态，必败状态为先手必败的状态
定理 1：没有后继状态的状态是必败状态
定理 2：一个状态是必胜状态当且仅当存在至少一个必败状态为它的后继状态
定理 3：一个状态是必败状态当且仅当它的所有后继状态均为必胜状态
对于定理 1，如果游戏进行不下去了，那么这个玩家就输掉了游戏
对于定理 2，如果该状态至少有一个后继状态为必败状态，那么玩家可以通过操作到该必败状态；
          此时对手的状态为必败状态——对手必定是失败的，而相反地，自己就获得了胜利
对于定理 3，如果不存在一个后继状态为必败状态，那么无论如何，玩家只能操作到必胜状态；
          此时对手的状态为必胜状态——对手必定是胜利的，自己就输掉了游戏
推荐 https://blog.csdn.net/acm_cxlove/article/details/7854530
https://oi-wiki.org/math/game-theory/
博弈论学习笔记 https://www.luogu.com.cn/blog/368107/notes-of-Game-Theory
todo 取石子游戏总结 https://www.luogu.com.cn/blog/Wolfycz/qian-tan-suan-fa-bo-yi-lun-zong-ling-kai-shi-di-bo-yi-lun-post
todo 阶梯博弈 https://codeforces.com/blog/entry/44651

入门分类讨论 https://codeforces.com/problemset/problem/78/C
三定理的模板题 https://codeforces.com/problemset/problem/1033/C
             https://atcoder.jp/contests/dp/tasks/dp_k
1-2-K Game https://codeforces.com/problemset/problem/1194/D 我的题解 https://www.acwing.com/file_system/file/content/whole/index/content/3179098/
TODO: 题目推荐 https://blog.csdn.net/ACM_cxlove/article/details/7854526
一道不错的有向图博弈 https://codeforces.com/problemset/problem/936/B
todo 威佐夫博弈 https://www.luogu.com.cn/problem/P2252
阶梯博弈 https://codeforces.com/problemset/problem/812/E
todo 阶梯博弈 移动金币 https://www.luogu.com.cn/problem/P5363
todo poj 2484 2348 1704 2311 | 1082 2068 3688 1740 2975 3537 2315
todo https://codeforces.com/problemset/problem/138/D (注：这是挑战上推荐的题目)
对于有环图的博弈，可以从终点（确定的状态）来倒推 https://leetcode-cn.com/problems/cat-and-mouse-ii/solution/mao-he-lao-shu-ii-bu-xu-yao-xian-zhi-bu-d2yxn/
通过必败态去筛必胜态 https://ac.nowcoder.com/acm/contest/11166/A
两端取数问题 https://atcoder.jp/contests/dp/tasks/dp_l LC486 https://leetcode-cn.com/problems/predict-the-winner/ LC877 https://leetcode-cn.com/problems/stone-game/
*/
func _() {
	{
		// 基础打表
		p, q := 3, 4

		const mx int = 100
		win := [mx + 1][2]int{} // -1 表示败；1 表示胜
		var f func(int, int) int
		f = func(n, who int) (res int) { // 0 为先手；1 为后手
			// 无法操作的情况
			if n == 0 {
				return -1
			}
			if who == 0 {
				// 检查边界
				if n <= p {
					return 1
				}
			} else {
				// 检查边界
				if n <= q {
					return 1
				}
			}
			dv := &win[n][who]
			if *dv != 0 {
				return *dv
			}
			defer func() { *dv = res }()
			// 检查是否可以转移到必败态
			if who == 0 {
				for i := 1; i <= p; i++ {
					if f(n-i, who^1) == -1 {
						return 1
					}
				}
			} else {
				for i := 1; i <= q; i++ {
					if f(n-i, who^1) == -1 {
						return 1
					}
				}
			}
			return -1
		}
		for i := 1; i <= mx; i++ {
			res := f(i, 0)
			if res == 1 {
				Println(i, res)
			}
		}
	}

	{
		// CF 1194D 打表
		// 上面三定理的基础题目
		const mx = 1000
		const k = 4
		win := [mx]bool{}
		win[1] = true
		win[2] = true
		for i := 3; i < k; i++ {
			win[i] = !win[i-1] || !win[i-2]
		}
		win[k] = true
		for i := k + 1; i < mx; i++ {
			win[i] = !win[i-1] || !win[i-2] || !win[i-k]
		}
		for i := 0; i < mx; i++ {
			Println(i, win[i])
		}
	}

	// 异或和不为零则先手必胜
	// https://blog.csdn.net/weixin_44023181/article/details/85619512
	// 模板题 https://www.luogu.com.cn/problem/P2197
	nim := func(a []int) (firstWin bool) {
		sum := 0
		for _, v := range a {
			sum ^= v
		}
		return sum > 0
	}

	// SG 函数 Sprague-Grundy theorem
	// 有向图游戏的某个局面必胜 <=> 该局面对应节点的 SG 函数值 > 0
	// 有向图游戏的某个局面必败 <=> 该局面对应节点的 SG 函数值 = 0
	// 推荐资料 Competitive Programmer’s Handbook Ch.25
	// https://oi-wiki.org/math/game-theory/#sg
	// https://en.wikipedia.org/wiki/Sprague%E2%80%93Grundy_theorem
	// https://cp-algorithms.com/game_theory/sprague-grundy-nim.html
	// todo https://zhuanlan.zhihu.com/p/257013159
	// todo 推荐论文《组合游戏略述——浅谈 SG 游戏的若干拓展及变形》
	// todo Anti-SG
	//
	// 参考《福州大学 ACMICPC 集训队资料》9.4
	// 常见限制条件下的 SG 值：
	// 最多取 m 个：SG(n) = n%(m+1)
	// 只能取奇数个：SG(n) = n%2
	// 只能取 2^i 个：SG(n) = n%3
	// 只能取 p^i 个（p 为奇素数）：SG(n) = n%2
	//
	// 整数分拆博弈 https://codeforces.com/problemset/problem/87/C
	// 取石子变形
	// - https://codeforces.com/problemset/problem/850/C
	// - https://codeforces.com/problemset/problem/1823/E
	// todo https://www.luogu.com.cn/problem/P2148

	// 剪纸博弈
	// https://www.acwing.com/problem/content/description/221/ http://poj.org/problem?id=2311
	// 要求 n >= 2, m >= 2
	cutPaperGame := func(n, m int) bool {
		_sg := make([][]int, n+5) // 简单地 +5，保证下面设置初始局面时不会越界
		for i := range _sg {
			_sg[i] = make([]int, m+5)
			for j := range _sg[i] {
				_sg[i][j] = -1
			}
		}
		var sg func(int, int) int
		sg = func(x, y int) (mex int) {
			ptr := &_sg[x][y]
			if *ptr != -1 {
				return *ptr
			}
			defer func() { *ptr = mex }()
			has := map[int]bool{} // 若能确定 mex 上限可以用 bool 数组
			for i := 2; i <= x-i; i++ {
				has[sg(i, y)^sg(x-i, y)] = true
			}
			for i := 2; i <= y-i; i++ {
				has[sg(x, i)^sg(x, y-i)] = true
			}
			for ; has[mex]; mex++ {
			}
			return
		}

		// 设定一些初始必败局面
		_sg[2][2] = 0
		_sg[2][3] = 0
		_sg[3][2] = 0
		// 计算有向图游戏的 SG 函数值
		return sg(n, m) > 0
	}

	_ = []interface{}{nim, cutPaperGame}
}
