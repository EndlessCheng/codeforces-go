另外下午两点在B站讲这场周赛的题目，会讲到本题的证明，[欢迎关注](https://space.bilibili.com/206214)~

---

# 方法一：记忆化搜索

用**邻项交换法**可以证明，对机器人和工厂按照位置从小到大排序，那么每个工厂修复的机器人就是连续的一段了。

定义 $f(i,j)$ 表示用第 $i$ 个及其右侧的工厂，修理第 $j$ 个及其右侧的机器人，机器人移动的最小总距离。

则有 $f(i,j) = \min\limits_{k}^{}{f(i+1,k+1) + \text{cost}(i,j,k)}$。

这里 $\text{cost}(i,j,k)$ 表示第 $i$ 个工厂修复从 $j$ 到 $k$ 的机器人，移动距离就是这些机器人到第 $i$ 个工厂的距离之和。

注意 $k-j+1\le\textit{limit}[i]$。

```py [sol1-Python3]
class Solution:
    def minimumTotalDistance(self, robot: List[int], factory: List[List[int]]) -> int:
        factory.sort(key=lambda f: f[0])
        robot.sort()
        n, m = len(factory), len(robot)
        @cache
        def f(i: int, j: int) -> int:
            if j == m: return 0
            if i == n - 1:
                if m - j > factory[i][1]: return inf
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

```go [sol1-Go]
func minimumTotalDistance(robot []int, factory [][]int) int64 {
	sort.Slice(factory, func(i, j int) bool { return factory[i][0] < factory[j][0] })
	sort.Ints(robot)
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
```

#### 复杂度分析

- 时间复杂度：$O(nm^2)$，其中 $n$ 为 $\textit{factory}$ 的长度，$m$ 为 $\textit{robot}$ 的长度。
- 空间复杂度：$O(nm)$。

# 方法二：递推

根据方法一，排序后，定义 $f[i][j]$ 表示前 $i$ 个工厂修复前 $j$ 个机器人的最小移动总距离。

枚举第 $i$ 个工厂修了 $k$ 个机器人，则有

$$
f[i][j] = \min\limits_{k=0}^{\min(j, \textit{limit}[i])} f[i-1][j-k] + \sum_{p=j-k+1}^{j} |\textit{robot}[p]-\textit{position}[i]|
$$

代码实现时，第一个维度可以像 01 背包那样优化掉。

```py [sol2-Python3]
class Solution:
    def minimumTotalDistance(self, robot: List[int], factory: List[List[int]]) -> int:
        factory.sort(key=lambda f: f[0])
        robot.sort()
        m = len(robot)
        f = [0] + [inf] * m
        for pos, limit in factory:
            for j in range(m, 0, -1):
                cost = 0
                for k in range(1, min(j, limit) + 1):
                    cost += abs(robot[j - k] - pos)
                    f[j] = min(f[j], f[j - k] + cost)
        return f[m]
```

```java [sol2-Java]
class Solution {
    public long minimumTotalDistance(List<Integer> robot, int[][] factory) {
        Arrays.sort(factory, (a, b) -> a[0] - b[0]);
        var r = robot.stream().mapToInt(i -> i).toArray();
        Arrays.sort(r);
        var m = r.length;
        var f = new long[m + 1];
        Arrays.fill(f, (long) 1e18);
        f[0] = 0;
        for (var fa : factory)
            for (var j = m; j > 0; j--) {
                var cost = 0L;
                for (var k = 1; k <= Math.min(j, fa[1]); ++k) {
                    cost += Math.abs(r[j - k] - fa[0]);
                    f[j] = Math.min(f[j], f[j - k] + cost);
                }
            }
        return f[m];
    }
}
```

```cpp [sol2-C++]
class Solution {
public:
    long long minimumTotalDistance(vector<int> &robot, vector<vector<int>> &factory) {
        sort(factory.begin(), factory.end(), [](auto &a, auto &b) { return a[0] < b[0]; });
        sort(robot.begin(), robot.end());
        int m = robot.size();
        long f[m + 1]; memset(f, 0x3f, sizeof(f));
        f[0] = 0;
        for (auto &fa: factory)
            for (int j = m; j > 0; j--) {
                long cost = 0L;
                for (int k = 1; k <= min(j, fa[1]); ++k) {
                    cost += abs(robot[j - k] - fa[0]);
                    f[j] = min(f[j], f[j - k] + cost);
                }
            }
        return f[m];
    }
};
```

```go [sol2-Go]
func minimumTotalDistance(robot []int, factory [][]int) int64 {
	sort.Slice(factory, func(i, j int) bool { return factory[i][0] < factory[j][0] })
	sort.Ints(robot)
	m := len(robot)
	f := make([]int, m+1)
	for i := range f {
		f[i] = 1e18
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
```

#### 复杂度分析

- 时间复杂度：$O(nm^2)$，其中 $n$ 为 $\textit{factory}$ 的长度，$m$ 为 $\textit{robot}$ 的长度。
- 空间复杂度：$O(m)$。

#### 相似题目

- [1478. 安排邮筒](https://leetcode.cn/problems/allocate-mailboxes/)

