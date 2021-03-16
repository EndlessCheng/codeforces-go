package copypasta

import (
	"math"
	"math/rand"
)

/* 随机化技巧
https://oi-wiki.org/misc/rand-technique/
*/

/* 模拟退火 (Simulated Annealing, SA)
基于 Metropolis 准则
https://en.wikipedia.org/wiki/Simulated_annealing
https://en.wikipedia.org/wiki/Metropolis%E2%80%93Hastings_algorithm
https://oi-wiki.org/misc/simulated-annealing/
https://www.luogu.com.cn/blog/Darth-Che/mu-ni-tui-huo-xue-xi-bi-ji
https://zhuanlan.zhihu.com/p/47234502
技巧：可以在时限内重复跑 SA 取最优值，防止脸黑

模板题 https://www.luogu.com.cn/problem/P1337
LC/周赛197D https://leetcode-cn.com/contest/weekly-contest-197/problems/best-position-for-a-service-centre/ http://poj.org/problem?id=2420 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=14&page=show_problem&problem=1169
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

/* 爬山算法 (Hill Climbing, HC)
https://en.wikipedia.org/wiki/Hill_climbing
https://oi-wiki.org/misc/hill-climbing/
https://leetcode-cn.com/problems/best-position-for-a-service-centre/solution/fu-wu-zhong-xin-de-zui-jia-wei-zhi-by-leetcode-sol/
*/
