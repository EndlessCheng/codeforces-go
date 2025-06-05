package copypasta

import . "fmt"

/* 博弈论 Game Theory

另见 dp.go 中的「博弈类 DP」

https://en.wikipedia.org/wiki/Game_theory
定义必胜状态为先手必胜的状态，必败状态为先手必败的状态
定理 1：没有后继状态的状态是必败状态
定理 2：一个状态是必胜状态当且仅当存在至少一个必败状态为它的后继状态
定理 3：一个状态是必败状态当且仅当它的所有后继状态均为必胜状态
对于定理 1，如果游戏无法继续下去，那么这个玩家就输掉了游戏
对于定理 2，如果该状态至少有一个后继状态为必败状态，那么玩家可以通过操作到该必败状态；
          此时对手的状态为必败状态——对手必定是失败的，而相反地，自己就获得了胜利
对于定理 3，如果不存在一个后继状态为必败状态，那么无论如何，玩家只能操作到必胜状态；
          此时对手的状态为必胜状态——对手必定是胜利的，自己就输掉了游戏
推荐 https://blog.csdn.net/acm_cxlove/article/details/7854530
https://oi-wiki.org/math/game-theory/
博弈论学习笔记 https://www.luogu.com.cn/blog/368107/notes-of-Game-Theory
todo 取石子游戏总结 https://www.luogu.com.cn/blog/Wolfycz/qian-tan-suan-fa-bo-yi-lun-zong-ling-kai-shi-di-bo-yi-lun-post
 阶梯博弈 https://codeforces.com/blog/entry/44651
 巧克力博弈 https://en.wikipedia.org/wiki/Chomp

mex = minimum excluded
https://en.wikipedia.org/wiki/Mex_(mathematics)

https://codeforces.com/problemset/problem/197/A 1600 入门
https://codeforces.com/problemset/problem/1033/C 1600 三定理的模板题
https://codeforces.com/problemset/problem/1194/D 1700 1-2-K Game 我的题解 https://www.acwing.com/file_system/file/content/whole/index/content/3179098/
https://atcoder.jp/contests/dp/tasks/dp_k
https://codeforces.com/problemset/problem/78/C 2000 分类讨论
LC2868 https://leetcode.cn/problems/the-wording-game/
TODO: 题目推荐 https://blog.csdn.net/ACM_cxlove/article/details/7854526
https://codeforces.com/problemset/problem/936/B 2100 一道不错的有向图博弈
todo 威佐夫博弈 https://www.luogu.com.cn/problem/P2252
阶梯博弈 https://codeforces.com/problemset/problem/812/E 2300
todo 阶梯博弈 移动金币 https://www.luogu.com.cn/problem/P5363
todo poj 2484 2348 1704 2311 | 1082 2068 3688 1740 2975 3537 2315
todo https://codeforces.com/problemset/problem/138/D (注：这是挑战上推荐的题目)
对于有环图的博弈，可以从终点（确定的状态）来倒推 https://leetcode.cn/problems/cat-and-mouse-ii/solution/mao-he-lao-shu-ii-bu-xu-yao-xian-zhi-bu-d2yxn/
通过必败态去筛必胜态 https://ac.nowcoder.com/acm/contest/11166/A
两端取数问题 https://atcoder.jp/contests/dp/tasks/dp_l
- LC486 https://leetcode.cn/problems/predict-the-winner/
- LC877 https://leetcode.cn/problems/stone-game/

交互+博弈
https://codeforces.com/problemset/problem/1934/D2 2400
todo https://codeforces.com/problemset/problem/1903/E 

*/
func _() {
	{
		// 基础打表
		p, q := 3, 4

		// 如果两个人的规则一样，则可以去掉 who
		// 例题 https://www.lanqiao.cn/problems/8051/learning/?contest_id=146
		const mx int = 100
		memo := make([][2]int8, mx+1) // 负数表示败；正数表示胜
		var f func(int, uint8) int8
		f = func(i int, who uint8) (res int8) { // 0 为先手；1 为后手
			// 无法操作的情况
			if i == 0 {
				return -1
			}
			if who == 0 {
				// 检查边界
				if i <= p {
					return 1
				}
			} else {
				// 检查边界
				if i <= q {
					return 1
				}
			}

			dv := &memo[i][who]
			if *dv != 0 {
				return *dv
			}
			defer func() { *dv = res }()

			// 检查是否可以转移到必败态
			if who == 0 {
				for j := 1; j <= p; j++ {
					if f(i-j, who^1) < 0 {
						return 1
					}
				}
			} else {
				for j := 1; j <= q; j++ {
					if f(i-j, who^1) < 0 {
						return 1
					}
				}
			}
			return -1
		}
		for i := 1; i <= mx; i++ {
			res := f(i, 0)
			if res > 0 {
				Println(i)
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
	// 因为先手可以构造出一个异或和为 0 的局面（把某一堆 a[i] 减少至 a[i]^xorSum），然后就可以模仿对手了。
	// 具体来说，设 xorSum 的最高位为 p，找到第 p 为是 1 的数 a[i]（这样的数一定存在，否则 xorSum 第 p 位一定是 0），
	// 那么 a[i]^xorSum 必然比 a[i] 小，所以这个减少操作是存在的。
	// 为什么可以模仿对手？因为无论对手如何操作，都会得到一个异或和不为 0 的局面，我们可以用同样的方法，构造出一个异或和为 0 的局面。
	// https://blog.csdn.net/weixin_44023181/article/details/85619512
	// 模板题 https://www.luogu.com.cn/problem/P2197
	// https://codeforces.com/problemset/problem/15/C
	// https://atcoder.jp/contests/abc368/tasks/abc368_f
	// LC1908 https://leetcode.cn/problems/game-of-nim/
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
	// mex = minimum excluded https://en.wikipedia.org/wiki/Mex_(mathematics)
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
	// - [2005. 斐波那契树的移除子树游戏](https://leetcode.cn/problems/subtree-removal-game-with-fibonacci-tree/)（会员题）
	// 整数分拆博弈 https://codeforces.com/problemset/problem/87/C
	// 取石子变形
	// - https://codeforces.com/problemset/problem/850/C
	// - https://codeforces.com/problemset/problem/1823/E
	// todo https://www.luogu.com.cn/problem/P2148
	//  https://atcoder.jp/contests/arc151/tasks/arc151_c

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
