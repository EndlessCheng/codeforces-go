[视频讲解](https://www.bilibili.com/video/BV1Ld4y1r71H) 已出炉，重点分析了记忆化搜索与递推的优缺点，以及为什么空间压缩要倒序遍历。

欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

## 方法一：记忆化搜索

用**邻项交换法**可以证明，对机器人和工厂按照位置从小到大排序，那么每个工厂修复的机器人就是连续的一段了。

定义 $f(i,j)$ 表示用第 $i$ 个及其右侧的工厂，修理第 $j$ 个及其右侧的机器人，机器人移动的最小总距离。

枚举第 $i$ 个工厂修了 $k$ 个机器人，则有 $f(i,j) = \min\limits_{k}^{}{f(i+1,j+k) + \text{cost}(i,j,k)}$。

这里 $\text{cost}(i,j,k)$ 表示第 $i$ 个工厂修复从 $j$ 到 $j+k-1$ 的机器人，移动距离就是这些机器人到第 $i$ 个工厂的距离之和。

注意 $k\le\textit{limit}[i]$。

```py [sol-Python3]
class Solution:
    def minimumTotalDistance(self, robot: List[int], factory: List[List[int]]) -> int:
        n, m = len(factory), len(robot)
        factory.sort(key=lambda f: f[0])
        robot.sort()

        @cache
        def f(i: int, j: int) -> int:
            if j == m:
                return 0
            if i == n - 1:
                if m - j > factory[i][1]:
                    return inf
                return sum(abs(x - factory[i][0]) for x in robot[j:])
            res = f(i + 1, j)
            s, k = 0, 1
            while k <= factory[i][1] and j + k - 1 < m:
                s += abs(robot[j + k - 1] - factory[i][0])
                res = min(res, s + f(i + 1, j + k))
                k += 1
            return res
        return f(0, 0)
```

```go [sol-Go]
func minimumTotalDistance(robot []int, factory [][]int) int64 {
	slices.SortFunc(factory, func(a, b []int) int { return a[0] - b[0] })
	slices.Sort(robot)
	n, m := len(factory), len(robot)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int
	f = func(i, j int) (res int) {
		if j >= m {
			return
		}
		dv := &dp[i][j]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		if i == n-1 {
			if m-j > factory[i][1] {
				return 1e18
			}
			for _, x := range robot[j:] {
				res += abs(x - factory[i][0])
			}
			return
		}
		res = f(i+1, j)
		for s, k := 0, 1; k <= factory[i][1] && j+k-1 < m; k++ {
			s += abs(robot[j+k-1] - factory[i][0])
			res = min(res, s+f(i+1, j+k))
		}
		return
	}
	return int64(f(0, 0))
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm^2)$，其中 $n$ 为 $\textit{factory}$ 的长度，$m$ 为 $\textit{robot}$ 的长度。
- 空间复杂度：$\mathcal{O}(nm)$。

## 方法二：递推

根据方法一，排序后，定义 $f[i][j]$ 表示前 $i$ 个工厂修复前 $j$ 个机器人的最小移动总距离。

枚举第 $i$ 个工厂修了 $k$ 个机器人，则有

$$
f[i][j] = \min\limits_{k=0}^{\min(j, \textit{limit}[i])} f[i-1][j-k] + \sum_{p=j-k+1}^{j} |\textit{robot}[p]-\textit{position}[i]|
$$

代码实现时，第一个维度可以像 01 背包那样优化掉。

```py [sol-Python3]
class Solution:
    def minimumTotalDistance(self, robot: List[int], factory: List[List[int]]) -> int:
        m = len(robot)
        factory.sort(key=lambda f: f[0])
        robot.sort()

        f = [0] + [inf] * m
        for pos, limit in factory:
            for j in range(m, 0, -1):
                cost = 0
                for k in range(1, min(j, limit) + 1):
                    cost += abs(robot[j - k] - pos)
                    f[j] = min(f[j], f[j - k] + cost)
        return f[m]
```

```java [sol-Java]
class Solution {
    public long minimumTotalDistance(List<Integer> robot, int[][] factory) {
        Arrays.sort(factory, (a, b) -> a[0] - b[0]);
        Integer[] r = robot.toArray(Integer[]::new);
        Arrays.sort(r);
        int m = r.length;

        long[] f = new long[m + 1];
        Arrays.fill(f, Long.MAX_VALUE / 2);
        f[0] = 0;
        for (int[] fa : factory) {
            for (int j = m; j > 0; j--) {
                long cost = 0;
                for (int k = 1; k <= Math.min(j, fa[1]); k++) {
                    cost += Math.abs(r[j - k] - fa[0]);
                    f[j] = Math.min(f[j], f[j - k] + cost);
                }
            }
        }
        return f[m];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minimumTotalDistance(vector<int>& robot, vector<vector<int>>& factory) {
        ranges::sort(factory, {}, [](auto& a) { return a[0]; });
        ranges::sort(robot);
        int m = robot.size();

        vector<long long> f(m + 1, LLONG_MAX / 2);
        f[0] = 0;
        for (auto& fa: factory) {
            for (int j = m; j > 0; j--) {
                long long cost = 0;
                for (int k = 1; k <= min(j, fa[1]); ++k) {
                    cost += abs(robot[j - k] - fa[0]);
                    f[j] = min(f[j], f[j - k] + cost);
                }
            }
        }
        return f[m];
    }
};
```

```go [sol-Go]
func minimumTotalDistance(robot []int, factory [][]int) int64 {
	slices.SortFunc(factory, func(a, b []int) int { return a[0] - b[0] })
	slices.Sort(robot)
	m := len(robot)
	f := make([]int, m+1)
	for i := range f {
		f[i] = math.MaxInt / 2
	}
	f[0] = 0
	for _, fa := range factory {
		for j := m; j > 0; j-- {
			for k, cost := 1, 0; k <= min(j, fa[1]); k++ {
				cost += abs(robot[j-k] - fa[0])
				f[j] = min(f[j], f[j-k]+cost)
			}
		}
	}
	return int64(f[m])
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm^2)$，其中 $n$ 为 $\textit{factory}$ 的长度，$m$ 为 $\textit{robot}$ 的长度。
- 空间复杂度：$\mathcal{O}(m)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
