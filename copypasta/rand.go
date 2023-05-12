package copypasta

import (
	"math"
	"math/rand"
	"time"
)

/* 随机化技巧
https://oi-wiki.org/misc/rand-technique/
随机梯度下降 SGD, Stochastic Gradient Descent https://en.wikipedia.org/wiki/Stochastic_gradient_descent
https://codeforces.com/problemset/problem/995/C
https://codeforces.com/problemset/problem/1314/D 推荐
https://codeforces.com/problemset/problem/1523/D
Kick Start 2021 Round C Binary Operator https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435c44/00000000007ec290
https://codeforces.com/problemset/problem/1689/D https://www.luogu.com.cn/blog/wangxiwen/solution-cf1689d
*/

/* 模拟退火 (Simulated Annealing, SA)
基于 Metropolis 准则
https://en.wikipedia.org/wiki/Simulated_annealing
https://en.wikipedia.org/wiki/Metropolis%E2%80%93Hastings_algorithm
https://oi-wiki.org/misc/simulated-annealing/
https://www.luogu.com.cn/blog/Darth-Che/mu-ni-tui-huo-xue-xi-bi-ji
https://zhuanlan.zhihu.com/p/47234502
https://www.cnblogs.com/ECJTUACM-873284962/p/8468831.html
技巧：可以在时限内重复跑 SA 取最优值，防止脸黑

Heuristic algorithm for Hamiltonian path in undirected graphs https://codeforces.com/blog/entry/90743

模板题 https://www.luogu.com.cn/problem/P1337
LC1515 https://leetcode-cn.com/problems/best-position-for-a-service-centre/ http://poj.org/problem?id=2420 UVa 10228 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=14&page=show_problem&problem=1169
todo 教学题 https://atcoder.jp/contests/intro-heuristics/tasks/intro_heuristics_a
 https://atcoder.jp/contests/ahc001/tasks/ahc001_a
 https://atcoder.jp/contests/ahc002/tasks/ahc002_a
*/
func simulatedAnnealing(f func(x float64) float64) float64 {
	// 例：最小值
	x := .0
	ans := f(x)
	for t := 1e5; t > 1e-8; t *= 0.99 {
		y := x + (2*rand.Float64()-1)*t
		v := f(y)
		if v < ans || math.Exp((ans-v)/t) > rand.Float64() { // 最小直接取，或者以一定概率接受较大的值
			ans = v
			x = y
		}
	}
	return ans
}

// 另一种写法（利用时限）
// 此时 alpha 可以设大点，例如 0.999
func simulatedAnnealingWithinTimeLimit(f func(x float64) float64) float64 {
	const timeLimit = 2 - 0.1
	t0 := time.Now()
	// 例：最小值
	x := .0
	ans := f(x)
	for t := 1e5; time.Since(t0).Seconds() < timeLimit; {
		y := x + (2*rand.Float64()-1)*t
		v := f(y)
		if v < ans || math.Exp((ans-v)/t) > rand.Float64() { // 最小直接取，或者以一定概率接受较大的值
			ans = v
			x = y
		}
		t *= 0.999 // 置于末尾，方便在 roll 到不合适的数据时直接 continue，同时也保证不会因为 roll 不到合适的数据而超时
	}
	return ans
}

/* 爬山算法 (Hill Climbing, HC)
https://en.wikipedia.org/wiki/Hill_climbing
https://oi-wiki.org/misc/hill-climbing/

https://en.wikipedia.org/wiki/Geometric_median
LC1515 https://leetcode.cn/problems/best-position-for-a-service-centre/
https://leetcode.cn/problems/best-position-for-a-service-centre/solution/fu-wu-zhong-xin-de-zui-jia-wei-zhi-by-leetcode-sol/
*/
