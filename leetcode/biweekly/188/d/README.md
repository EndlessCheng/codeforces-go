

## 核心思路

1. 求出最多能服务的车辆数 $\textit{maxCars}$。
2. 二分最大等待时间 $m$。
3. 判断在最大等待时间为 $m$ 的情况下，能否服务 $\textit{maxCars}$ 辆车。

## 最多能服务的车辆数

写一个记忆化搜索，定义 $\textit{dfs}(i,\textit{fuel}_0,\textit{fuel}_1)$ 表示从车 $i$ 开始服务，两个加油机的剩余燃料量分别为 $\textit{fuel}_0$ 和 $\textit{fuel}_1$ 的前提下，最多能服务的车辆数。

枚举选哪个：

- 车 $i$ 去加油机 0，前提是 $\textit{demand}[i]\le \textit{fuel}_0$。递归到 $\textit{dfs}(i+1,\textit{fuel}_0-\textit{demand}[i],\textit{fuel}_1)$。
- 车 $i$ 去加油机 1，前提是 $\textit{demand}[i]\le \textit{fuel}_1$。递归到 $\textit{dfs}(i+1,\textit{fuel}_0,\textit{fuel}_1-\textit{demand}[i])$。

两种情况取最大值，再加上 $1$（车 $i$），就得到了 $\textit{dfs}(i,\textit{fuel}_0,\textit{fuel}_1)$。

递归边界：$\textit{dfs}(n,\textit{fuel}_0,\textit{fuel}_1) = 0$。没有可服务的车。

递归入口：$\textit{dfs}(0,\textit{fuel}[0],\textit{fuel}[1])$，其返回值记作 $\textit{maxCars}$。

如果 $\textit{maxCars} = 0$，返回 $-1$。

## 二分答案

设车辆等待时间的上限为 $m$。

如果在 $m$ 的约束下可以服务 $\textit{maxCars}$ 辆车，那么对于更大的 $m$，约束更宽松，也可以服务 $\textit{maxCars}$ 辆车。

如果在 $m$ 的约束下无法服务 $\textit{maxCars}$ 辆车，那么对于更小的 $m$，约束更苛刻，也无法服务 $\textit{maxCars}$ 辆车。

据此，可以**二分猜答案**。关于二分算法的原理，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

现在问题转化成一个判定性问题：

- 给定车辆等待时间的上限 $m$，判断能否在 $m$ 的约束下服务 $\textit{maxCars}$ 辆车。

判断方法见下一小节，这里先讨论二分范围。

下面代码采用开区间二分。使用闭区间或者半闭半开区间也是可以的，喜欢哪种写法就用哪种。

- 开区间左端点初始值：$-1$。无法满足要求。
- 开区间右端点初始值：$\max(\textit{demand})$。即使只有一个加油机，最多也只需等 $\max(\textit{demand})$ 秒。

> 对于开区间写法，简单来说 `check(mid) == true` 时更新的是谁，最后就返回谁。相比其他二分写法，开区间写法不需要思考加一减一等细节，更简单。推荐使用开区间写二分。

## 能否在 m 的约束下服务 maxCars 辆车

在前一个 $\textit{dfs}$ 的基础上，额外增加两个参数：

- 加油机 0 在 $\textit{wait}_0$ 秒后空闲。
- 加油机 1 在 $\textit{wait}_1$ 秒后空闲。

设 $d = \textit{demand}[i]$，分类讨论：

- 车 $i$ 选择加油机 0。
    - 车 $i$ 要等 $\textit{wait}_0$ 秒才能开始加油，所以要满足 $\textit{wait}_0\le m$（以及 $d \le \textit{fuel}_0$）。
    - 因为时间流逝了 $\textit{wait}_0$ 秒，加油机 1 变成在 $\max(\textit{wait}_1 - \textit{wait}_0,0)$ 秒后空闲。
    - 从状态 $(i,\textit{wait}_0,\textit{wait}_1,\textit{fuel}_0,\textit{fuel}_1)$ 移动到状态 $(i,d,\max(\textit{wait}_1 - \textit{wait}_0,0),\textit{fuel}_0-d,\textit{fuel}_1)$。
- 车 $i$ 选择加油机 0。
    - 车 $i$ 要等 $\textit{wait}_1$ 秒才能开始加油，所以要满足 $\textit{wait}_1\le m$（以及 $\textit{demand}[i]\le \textit{fuel}_1$）。
    - 因为时间流逝了 $\textit{wait}_1$ 秒，加油机 0 变成在 $\max(\textit{wait}_0 - \textit{wait}_1,0)$ 秒后空闲。
    - 从状态 $(i,\textit{wait}_0,\textit{wait}_1,\textit{fuel}_0,\textit{fuel}_1)$ 移动到状态 $(i,\max(\textit{wait}_0 - \textit{wait}_1,0),d,\textit{fuel}_0,\textit{fuel}_1-d)$。

如果能走到 $i=\textit{maxCars}$ 的状态，返回 $\texttt{true}$。

初始状态 $(0,0,0,\textit{fuel}[0],\textit{fuel}[1])$。

[本题视频讲解](https://www.bilibili.com/video/BV1Y73R6zEwG/?t=18m13s)，基于原始题面讲解，思路是一样的。

```py [sol-Python3]
class Solution:
    def minMaxWaitingTime(self, demand: list[int], fuel: list[int]) -> int:
        # 1. 求出最多能服务的车辆数 max_cars
        @cache
        def calc_max_cars(i: int, fuel0: int, fuel1: int) -> int:
            if i == len(demand):
                return 0
            res = 0
            d = demand[i]
            if d <= fuel0:
                res = max(res, calc_max_cars(i + 1, fuel0 - d, fuel1) + 1)
            if d <= fuel1:
                res = max(res, calc_max_cars(i + 1, fuel0, fuel1 - d) + 1)
            return res

        max_cars = calc_max_cars(0, fuel[0], fuel[1])
        if max_cars == 0:
            return -1

        # 3. 判断在最大等待时间 <= max_waiting_time 的约束下，能否服务 max_cars 辆车
        def check(max_waiting_time: int) -> bool:
            # 加油机 0 在 wait0 秒后空闲，剩余燃料量 fuel0
            # 加油机 1 在 wait1 秒后空闲，剩余燃料量 fuel1
            @cache  # 这里的 @cache 相当于 vis set
            def dfs(i: int, wait0: int, wait1: int, fuel0: int, fuel1: int) -> bool:
                if i == max_cars:
                    return True

                d = demand[i]

                # 选择加油机 0，等 wait0 秒开始加油，加油机 1 的等待时间减少 wait0 秒
                if wait0 <= max_waiting_time and d <= fuel0 and \
                        dfs(i + 1, d, max(wait1 - wait0, 0), fuel0 - d, fuel1):
                    return True

                # 选择加油机 1，等 wait1 秒开始加油，加油机 0 的等待时间减少 wait1 秒
                if wait1 <= max_waiting_time and d <= fuel1 and \
                        dfs(i + 1, max(wait0 - wait1, 0), d, fuel0, fuel1 - d):
                    return True

                return False

            return dfs(0, 0, 0, fuel[0], fuel[1])

        # 2. 二分最大等待时间
        return bisect_left(range(max(demand)), True, key=check)
```

```go [sol-Go]
func minMaxWaitingTime(demand []int, fuel []int) int {
	// 1. 求出最多能服务的车辆数 maxCars
	type fuelArgs struct{ i, fuel0, fuel1 int }
	memo := map[fuelArgs]int{}

	var calcMaxCars func(int, int, int) int
	calcMaxCars = func(i, fuel0, fuel1 int) (res int) {
		if i == len(demand) {
			return
		}

		args := fuelArgs{i, fuel0, fuel1}
		if v, ok := memo[args]; ok {
			return v
		}

		d := demand[i]
		if d <= fuel0 {
			res = calcMaxCars(i+1, fuel0-d, fuel1) + 1
		}
		if d <= fuel1 {
			res = max(res, calcMaxCars(i+1, fuel0, fuel1-d)+1)
		}

		memo[args] = res
		return res
	}

	maxCars := calcMaxCars(0, fuel[0], fuel[1])
	if maxCars == 0 {
		return -1
	}

	// 2. 二分最大等待时间
	ans := sort.Search(slices.Max(demand), func(maxWaitingTime int) bool {
		// 3. 判断在最大等待时间 <= maxWaitingTime 的约束下，能否服务 maxCars 辆车
		type state struct{ i, wait0, wait1, fuel0, fuel1 int }
		vis := map[state]bool{}

		// 加油机 0 在 wait0 秒后空闲，剩余燃料量 fuel0
		// 加油机 1 在 wait1 秒后空闲，剩余燃料量 fuel1
		var dfs func(int, int, int, int, int) bool
		dfs = func(i, wait0, wait1, fuel0, fuel1 int) bool {
			if i == maxCars {
				return true
			}

			st := state{i, wait0, wait1, fuel0, fuel1}
			if vis[st] {
				return false
			}
			vis[st] = true

			d := demand[i]

			// 选择加油机 0，等 wait0 秒开始加油，加油机 1 的等待时间减少 wait0 秒
			if wait0 <= maxWaitingTime && d <= fuel0 &&
				dfs(i+1, d, max(wait1-wait0, 0), fuel0-d, fuel1) {
				return true
			}

			// 选择加油机 1，等 wait1 秒开始加油，加油机 0 的等待时间减少 wait1 秒
			if wait1 <= maxWaitingTime && d <= fuel1 &&
				dfs(i+1, max(wait0-wait1, 0), d, fuel0, fuel1-d) {
				return true
			}

			return false
		}

		return dfs(0, 0, 0, fuel[0], fuel[1])
	})
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nU^2F\log U)$，其中 $n$ 是 $\textit{demand}$ 的长度，$U=\max(\textit{demand})$，$F = \min(\textit{fuel})$。由于车 $[0,i-1]$ 消耗的燃料量是固定的，知道 $i$ 和 $\textit{fuel}_0$ 可以直接算出 $\textit{fuel}_1$，所以实际上第二个递归函数是四维的。
- 空间复杂度：$\mathcal{O}(nU^2F)$。

## 状态优化

由于其中一个 $\textit{wait}_{i}$ 始终是 $d$，所以只需要保留另一个 $\textit{wait}_i$。

这个优化技巧类似 [1320. 二指输入的的最小距离](https://leetcode.cn/problems/minimum-distance-to-type-a-word-using-two-fingers/)，推荐先把那题做了，并阅读 [我的题解](https://leetcode.cn/problems/minimum-distance-to-type-a-word-using-two-fingers/solutions/3946229/jiao-ni-yi-bu-bu-si-kao-dpji-yi-hua-sou-d9vls/)。

```py [sol-Python3]
class Solution:
    def minMaxWaitingTime(self, demand: list[int], fuel: list[int]) -> int:
        # 1. 求出最多能服务的车辆数 max_cars
        @cache
        def calc_max_cars(i: int, fuel0: int, fuel1: int) -> int:
            if i == len(demand):
                return 0
            res = 0
            d = demand[i]
            if d <= fuel0:
                res = max(res, calc_max_cars(i + 1, fuel0 - d, fuel1) + 1)
            if d <= fuel1:
                res = max(res, calc_max_cars(i + 1, fuel0, fuel1 - d) + 1)
            return res

        max_cars = calc_max_cars(0, fuel[0], fuel[1])
        if max_cars == 0:
            return -1

        # 3. 判断在最大等待时间 <= max_waiting_time 的约束下，能否服务 max_cars 辆车
        def check(max_waiting_time: int) -> bool:
            @cache  # 这里的 @cache 相当于 vis set
            def dfs(i: int, wait1: int, fuel0: int, fuel1: int) -> bool:
                if i == max_cars:
                    return True

                wait0 = demand[i - 1] if i else 0
                d = demand[i]

                # 跟在车 i-1 后面加油
                if wait0 <= max_waiting_time and d <= fuel0 and \
                        dfs(i + 1, max(wait1 - wait0, 0), fuel0 - d, fuel1):
                    return True

                # 不跟在车 i-1 后面加油
                if wait1 <= max_waiting_time and d <= fuel1 and \
                        dfs(i + 1, max(wait0 - wait1, 0), fuel1 - d, fuel0):  # 注意这里交换了 fuel0 和 fuel1
                    return True

                return False

            return dfs(0, 0, fuel[0], fuel[1])

        # 2. 二分最大等待时间
        return bisect_left(range(max(demand)), True, key=check)
```

```go [sol-Go]
func minMaxWaitingTime(demand []int, fuel []int) int {
	// 1. 求出最多能服务的车辆数 maxCars
	type fuelArgs struct{ i, fuel0, fuel1 int }
	memo := map[fuelArgs]int{}

	var calcMaxCars func(int, int, int) int
	calcMaxCars = func(i, fuel0, fuel1 int) (res int) {
		if i == len(demand) {
			return
		}

		args := fuelArgs{i, fuel0, fuel1}
		if v, ok := memo[args]; ok {
			return v
		}

		d := demand[i]
		if d <= fuel0 {
			res = calcMaxCars(i+1, fuel0-d, fuel1) + 1
		}
		if d <= fuel1 {
			res = max(res, calcMaxCars(i+1, fuel0, fuel1-d)+1)
		}

		memo[args] = res
		return res
	}

	maxCars := calcMaxCars(0, fuel[0], fuel[1])
	if maxCars == 0 {
		return -1
	}

	// 2. 二分最大等待时间
	ans := sort.Search(slices.Max(demand), func(maxWaitingTime int) bool {
		// 3. 判断在最大等待时间 <= maxWaitingTime 的约束下，能否服务 maxCars 辆车
		type state struct{ i, wait1, fuel0, fuel1 int }
		vis := map[state]bool{}

		var dfs func(int, int, int, int) bool
		dfs = func(i, wait1, fuel0, fuel1 int) bool {
			if i == maxCars {
				return true
			}

			st := state{i, wait1, fuel0, fuel1}
			if vis[st] {
				return false
			}
			vis[st] = true

			wait0 := 0
			if i > 0 {
				wait0 = demand[i-1]
			}
			d := demand[i]

			// 跟在车 i-1 后面加油
			if wait0 <= maxWaitingTime && d <= fuel0 &&
				dfs(i+1, max(wait1-wait0, 0), fuel0-d, fuel1) {
				return true
			}

			// 不跟在车 i-1 后面加油
			if wait1 <= maxWaitingTime && d <= fuel1 &&
				dfs(i+1, max(wait0-wait1, 0), fuel1-d, fuel0) { // 注意这里交换了 fuel0 和 fuel1
				return true
			}

			return false
		}

		return dfs(0, 0, fuel[0], fuel[1])
	})
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nUF\log U)$，其中 $n$ 是 $\textit{demand}$ 的长度，$U=\max(\textit{demand})$，$F = \max(\textit{fuel})$。由于车 $[0,i-1]$ 消耗的燃料量是固定的，知道 $i$ 和 $\textit{fuel}_0$ 可以直接算出 $\textit{fuel}_1$，所以实际上第二个递归函数是三维的。
- 空间复杂度：$\mathcal{O}(nUF)$。

## 专题训练

见下面动态规划题单的「**§7.6 多维 DP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
