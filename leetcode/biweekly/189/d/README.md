## 一、寻找子问题

暴力的想法是，枚举 $\textit{requests}$ 的全排列，但这显然太慢了。

设 $a,b,c$ 是 $\textit{requests}$ 中的三个数，且 $a < b < c$。在从楼层 $a$ 移动到楼层 $c$ 的过程中，一定会经过楼层 $b$，我们楼层 $b$ 的请求可以顺带处理。这意味着，我们可以把 $a\to c$ 的过程拆分成 $a\to b\to c$。相应地，惩罚值也可以拆分成 $a\to b$ 的惩罚值加上 $b\to c$ 的惩罚值。

设电梯到达过的最小楼层为 $p$，最大楼层为 $q$，那么电梯也一定到达过 $[p,q]$ 中的楼层，我们也顺带处理了相应的请求。根据上一段的分析，我们**只需考虑 $p$ 下面的最近请求楼层，以及 $q$ 上面的最近请求楼层，无需枚举所有请求楼层**。并且，电梯移动后，要么位于最小楼层，要么位于最大楼层。

例如，现在处理完了 $\textit{requests}$（已排序）的连续子数组 $[3,6]$，且电梯位于左端点 $\textit{requests}[3]$，那么下一步可以：

- 把电梯移到 $\textit{requests}[2]$，问题变成处理完了 $\textit{requests}$ 的连续子数组 $[2,6]$，且电梯位于左端点，处理剩余请求的最小总惩罚。
- 把电梯移到 $\textit{requests}[7]$，问题变成处理完了 $\textit{requests}$ 的连续子数组 $[3,7]$，且电梯位于右端点，处理剩余请求的最小总惩罚。

这些问题都是**和原问题相似的、规模更小的子问题**，可以用**递归**解决。

## 二、状态定义与状态转移方程

根据上面的讨论，定义 $\textit{dfs}(i,j,\textit{isRight})$ 表示处理完了 $\textit{requests}$ 的连续子数组 $[i,j]$，且电梯位于左（右）端点，处理剩余请求的最小总惩罚。如果 $\textit{isRight} = \texttt{true}$，则电梯位于 $\textit{requests}[j]$，否则位于 $\textit{requests}[i]$。

设 $\textit{requests}$ 的长度为 $m$，电梯现在位于楼层 $x$。枚举接下来电梯去哪：

- 去 $\textit{requests}[i-1]$，那么接下来要解决的问题是：处理完了 $\textit{requests}$ 的连续子数组 $[i-1,j]$，且电梯位于左端点，处理剩余请求的最小总惩罚，即 $\textit{dfs}(i-1,j,\texttt{false})$，再加上 $x\to \textit{requests}[i-1]$ 的惩罚值。有 $j-i+1$ 个已处理的请求，剩下 $m-(j-i+1)$ 个未处理的请求。经过了 $x - \textit{requests}[i-1]$ 秒，所以 $x\to \textit{requests}[i-1]$ 的惩罚值为 $(x-\textit{requests}[i-1])\cdot(m-(j-i+1))$。
- 去 $\textit{requests}[j+1]$，那么接下来要解决的问题是：处理完了 $\textit{requests}$ 的连续子数组 $[i,j+1]$，且电梯位于右端点，处理剩余请求的最小总惩罚，即 $\textit{dfs}(i,j+1,\texttt{true})$，再加上 $x\to \textit{requests}[j+1]$ 的惩罚值 $(\textit{requests}[j+1]-x)\cdot(m-(j-i+1))$。

这两种情况取最小值，就得到了 $\textit{dfs}(i,j,\textit{isRight})$。

**递归边界**：$\textit{dfs}(0,m-1)=0$。处理完所有请求。

**递归入口**：代码实现时，可以把 $\textit{start}$ 插入 $\textit{requests}$，设其下标为 $i$。那么答案为 $\textit{dfs}(i,i,\texttt{false})$。

进一步地，可以把 $-1$ 和 $n$ 作为哨兵插入 $\textit{requests}$，简化越界判断逻辑。

## 三、递归搜索 + 保存递归返回值 = 记忆化搜索

考虑到整个递归过程中有大量重复递归调用（递归入参相同）。由于递归函数没有副作用，同样的入参无论计算多少次，算出来的结果都是一样的，因此可以用**记忆化搜索**来优化：

- 如果一个状态（递归入参）是第一次遇到，那么可以在返回前，把状态及其结果记到一个 $\textit{memo}$ 数组中。
- 如果一个状态不是第一次遇到（$\textit{memo}$ 中保存的结果不等于 $\textit{memo}$ 的初始值），那么可以直接返回 $\textit{memo}$ 中保存的结果。

⚠**注意**：$\textit{memo}$ 数组的**初始值**一定不能等于要记忆化的值！例如初始值设置为 $0$，并且要记忆化的 $\textit{dfs}(i,j,\textit{isRight})$ 也等于 $0$，那就没法判断 $0$ 到底表示第一次遇到这个状态，还是表示之前遇到过了，从而导致记忆化失效。一般把初始值设置为 $-1$。

> Python 用户可以无视上面这段，直接用 `@cache` 装饰器。

[本题视频讲解](https://www.bilibili.com/video/BV1Q9bD6gEi1/?t=10m9s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def elevatorRequests(self, n: int, start: int, requests: list[int]) -> int:
        requests += [start, -1, n]  # 插入 start 和两个哨兵
        requests.sort()
        m = len(requests)

        # 已处理完 requests 的子数组 [i, j]
        # is_right=False 表示电梯在 requests[i]
        # is_right=True  表示电梯在 requests[j]
        @cache  # 缓存装饰器，避免重复计算 dfs（一行代码实现记忆化）
        def dfs(i: int, j: int, is_right: bool) -> int:
            if i == 0 or j == m - 1:  # 出界
                return inf
            if i == 1 and j == m - 2:  # 已处理完所有请求
                return 0
            x = requests[j if is_right else i]
            remain = m - 3 - j + i
            return min(dfs(i - 1, j, False) + (x - requests[i - 1]) * remain,  # 往左
                       dfs(i, j + 1, True) + (requests[j + 1] - x) * remain)  # 往右

        i = bisect_left(requests, start)
        ans = dfs(i, i, False)  # 这里 False 和 True 是一样的
        dfs.cache_clear()  # 避免超出内存限制
        return ans
```

```java [sol-Java]
class Solution {
    public long elevatorRequests(int n, int start, int[] requests) {
        int m = requests.length + 3;
        int[] a = Arrays.copyOf(requests, m);
        // 插入 start 和两个哨兵
        a[m - 3] = start;
        a[m - 2] = -1;
        a[m - 1] = n;
        Arrays.sort(a);

        long[][][] memo = new long[m - 1][m - 1][2];
        for (long[][] mat : memo) {
            for (long[] row : mat) {
                Arrays.fill(row, -1);  // -1 表示该状态没有计算过
            }
        }

        int i = Arrays.binarySearch(a, start);
        return dfs(i, i, 0, a, memo); // 这里 0 和 1 是一样的
    }

    // 已处理完 requests 的子数组 [i, j]
    // isRight = 0 表示电梯在 requests[i]
    // isRight = 1  表示电梯在 requests[j]
    private long dfs(int i, int j, int isRight, int[] a, long[][][] memo) {
        int m = a.length;
        if (i == 0 || j == m - 1) { // 出界
            return Long.MAX_VALUE / 2;
        }
        if (i == 1 && j == m - 2) { // 已处理完所有请求
            return 0;
        }

        long res = memo[i][j][isRight];
        if (res != -1) { // 之前计算过
            return res;
        }

        int x = a[isRight > 0 ? j : i];
        int remain = m - 3 - j + i;
        res = Math.min(dfs(i - 1, j, 0, a, memo) + (long) (x - a[i - 1]) * remain,  // 往左
                       dfs(i, j + 1, 1, a, memo) + (long) (a[j + 1] - x) * remain); // 往右
        memo[i][j][isRight] = res; // 记忆化
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long elevatorRequests(int n, int start, vector<int>& requests) {
        requests.insert(requests.end(), {start, -1, n}); // 插入 start 和两个哨兵
        ranges::sort(requests);
        int m = requests.size();
        vector memo(m - 1, vector(m - 1, array<long long, 2>{-1, -1})); // -1 表示该状态没有计算过

        // 已处理完 requests 的子数组 [i, j]
        // is_right=false 表示电梯在 requests[i]
        // is_right=true  表示电梯在 requests[j]
        auto dfs = [&](this auto&& dfs, int i, int j, bool is_right) -> long long {
            if (i == 0 || j == m - 1) { // 出界
                return LLONG_MAX / 2;
            }
            if (i == 1 && j == m - 2) { // 已处理完所有请求
                return 0;
            }

            long long& res = memo[i][j][is_right]; // 注意这里是引用
            if (res != -1) {
                return res;
            }

            int x = requests[is_right ? j : i];
            int remain = m - 3 - j + i;
            res = min(dfs(i - 1, j, false) + 1LL * (x - requests[i - 1]) * remain, // 往左
                      dfs(i, j + 1, true) + 1LL * (requests[j + 1] - x) * remain); // 往右
            return res;
        };

        int i = ranges::lower_bound(requests, start) - requests.begin();
        return dfs(i, i, false); // 这里 false 和 true 是一样的
    }
};
```

```go [sol-Go]
func elevatorRequests(n int, start int, requests []int) int64 {
	requests = append(requests, start, -1, n) // 插入 start 和两个哨兵
	slices.Sort(requests)
	m := len(requests)

	memo := make([][][2]int, m-1)
	for i := range memo {
		memo[i] = make([][2]int, m-1)
		for j := range memo[i] {
			memo[i][j] = [2]int{-1, -1} // -1 表示该状态没有计算过
		}
	}

	// 已处理完 requests 的子数组 [i, j]
	// isRight = 0 表示电梯在 requests[i]
	// isRight = 1 表示电梯在 requests[j]
	var dfs func(int, int, int) int
	dfs = func(i, j, isRight int) int {
		if i == 0 || j == m-1 { // 出界
			return math.MaxInt / 2
		}
		if i == 1 && j == m-2 { // 已处理完所有请求
			return 0
		}

		p := &memo[i][j][isRight]
		if *p != -1 { // 之前计算过
			return *p
		}

		var x int
		if isRight > 0 {
			x = requests[j]
		} else {
			x = requests[i]
		}
		remain := m - 3 - j + i
		*p = min(dfs(i-1, j, 0)+(x-requests[i-1])*remain, // 往左
				dfs(i, j+1, 1)+(requests[j+1]-x)*remain) // 往右
		return *p
	}

	i := sort.SearchInts(requests, start)
	return int64(dfs(i, i, 0)) // 这里 0 和 1 是一样的
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m^2)$，其中 $m$ 是 $\textit{requests}$ 的长度。由于每个状态只会计算一次，记忆化搜索的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数为 $\mathcal{O}(m^2)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以总的时间复杂度为 $\mathcal{O}(m^2)$。
- 空间复杂度：$\mathcal{O}(m^2)$。保存多少状态，就需要多少空间。

## 四、1:1 翻译成递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

具体来说，$f[i][j][\textit{isRight}]$ 的定义和 $\textit{dfs}(i,j,\textit{isRight})$ 的定义是一样的，都表示表示处理完了 $\textit{requests}$ 的连续子数组 $[i,j]$，且电梯位于左（右）端点，处理剩余请求的最小总惩罚。如果 $\textit{isRight} = 0$，则电梯位于 $\textit{requests}[i]$；如果 $\textit{isRight} = 1$，则电梯位于 $\textit{requests}[j]$。

#### 答疑

**问**：如何思考循环顺序？什么时候要正序枚举，什么时候要倒序枚举？

**答**：这里有一个通用的做法。看状态转移方程，为了计算 $f[i][j][\cdot]$，一方面，必须先把 $f[i-1][j][\cdot]$ 算出来，那么只有 $i$ 从小到大枚举才能做到；另一方面，必须先把 $f[i][j+1][\cdot]$ 算出来，那么只有 $j$ 从大到小枚举才能做到。

```py [sol-Python3]
class Solution:
    def elevatorRequests(self, n: int, start: int, requests: list[int]) -> int:
        requests += [start, -1, n]  # 插入 start 和两个哨兵
        requests.sort()
        m = len(requests)
        idx = bisect_left(requests, start)
        f = [[[inf, inf] for _ in range(m)] for _ in range(idx + 1)]
        for i in range(1, idx + 1):
            for j in range(m - 2, idx - 1, -1):
                if i == 1 and j == m - 2:
                    f[i][j][0] = f[i][j][1] = 0
                    continue
                remain = m - 3 - j + i
                f[i][j][0] = min(f[i - 1][j][0] + (requests[i] - requests[i - 1]) * remain,  # 往左
                                 f[i][j + 1][1] + (requests[j + 1] - requests[i]) * remain)  # 往右
                f[i][j][1] = min(f[i - 1][j][0] + (requests[j] - requests[i - 1]) * remain,  # 往左
                                 f[i][j + 1][1] + (requests[j + 1] - requests[j]) * remain)  # 往右
        return f[idx][idx][0]
```

```java [sol-Java]
class Solution {
    public long elevatorRequests(int n, int start, int[] requests) {
        int m = requests.length + 3;
        int[] a = Arrays.copyOf(requests, m);
        // 插入 start 和两个哨兵
        a[m - 3] = start;
        a[m - 2] = -1;
        a[m - 1] = n;
        Arrays.sort(a);

        int idx = Arrays.binarySearch(a, start);
        long[][][] f = new long[idx + 1][m][2];
        for (long[] row : f[0]) {
            Arrays.fill(row, Long.MAX_VALUE / 2);
        }

        for (int i = 1; i <= idx; i++) {
            f[i][m - 1][0] = f[i][m - 1][1] = Long.MAX_VALUE / 2;
            for (int j = m - 2; j >= idx; j--) {
                if (i == 1 && j == m - 2) {
                    f[i][j][0] = f[i][j][1] = 0;
                    continue;
                }
                int remain = m - 3 - j + i;
                f[i][j][0] = Math.min(f[i - 1][j][0] + (long) (a[i] - a[i - 1]) * remain,  // 往左
                                      f[i][j + 1][1] + (long) (a[j + 1] - a[i]) * remain); // 往右
                f[i][j][1] = Math.min(f[i - 1][j][0] + (long) (a[j] - a[i - 1]) * remain,  // 往左
                                      f[i][j + 1][1] + (long) (a[j + 1] - a[j]) * remain); // 往右
            }
        }
        return f[idx][idx][0];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long elevatorRequests(int n, int start, vector<int>& requests) {
        requests.insert(requests.end(), {start, -1, n}); // 插入 start 和两个哨兵
        ranges::sort(requests);
        int m = requests.size();
        int idx = ranges::lower_bound(requests, start) - requests.begin();
        vector f(idx + 1, vector(m, array<long long, 2>{LLONG_MAX / 2, LLONG_MAX / 2}));
        for (int i = 1; i <= idx; i++) {
            for (int j = m - 2; j >= idx; j--) {
                if (i == 1 && j == m - 2) {
                    f[i][j][0] = f[i][j][1] = 0;
                    continue;
                }
                int remain = m - 3 - j + i;
                f[i][j][0] = min(f[i - 1][j][0] + 1LL * (requests[i] - requests[i - 1]) * remain,  // 往左
                                 f[i][j + 1][1] + 1LL * (requests[j + 1] - requests[i]) * remain); // 往右
                f[i][j][1] = min(f[i - 1][j][0] + 1LL * (requests[j] - requests[i - 1]) * remain,  // 往左
                                 f[i][j + 1][1] + 1LL * (requests[j + 1] - requests[j]) * remain); // 往右
            }
        }
        return f[idx][idx][0];
    }
};
```

```go [sol-Go]
func elevatorRequests(n int, start int, requests []int) int64 {
	requests = append(requests, start, -1, n) // 插入 start 和两个哨兵
	slices.Sort(requests)
	m := len(requests)

	idx := sort.SearchInts(requests, start)
	f := make([][][2]int, idx+1)
	for i := range f {
		f[i] = make([][2]int, m)
	}
	for j := range f[0] {
		f[0][j] = [2]int{math.MaxInt / 2, math.MaxInt / 2}
	}

	for i := 1; i <= idx; i++ {
		f[i][m-1] = [2]int{math.MaxInt / 2, math.MaxInt / 2}
		for j := m - 2; j >= idx; j-- {
			if i == 1 && j == m-2 {
				f[i][j] = [2]int{}
				continue
			}
			remain := m - 3 - j + i
			f[i][j][0] = min(f[i-1][j][0]+(requests[i]-requests[i-1])*remain, // 往左
							f[i][j+1][1]+(requests[j+1]-requests[i])*remain) // 往右
			f[i][j][1] = min(f[i-1][j][0]+(requests[j]-requests[i-1])*remain, // 往左
							f[i][j+1][1]+(requests[j+1]-requests[j])*remain) // 往右
		}
	}
	return int64(f[idx][idx][0])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m^2)$，其中 $m$ 是 $\textit{requests}$ 的长度。
- 空间复杂度：$\mathcal{O}(m^2)$。

## 五、空间优化

观察上面的状态转移方程，在计算 $f[i]$ 时，只会用到 $f[i-1]$，不会用到比 $i$ 更早的状态。我们可以去掉 $f$ 的第一个维度，把 $f[i-1]$ 和 $f[i]$ 保存到同一个数组中。

```py [sol-Python3]
class Solution:
    def elevatorRequests(self, n: int, start: int, requests: list[int]) -> int:
        requests += [start, -1, n]  # 插入 start 和两个哨兵
        requests.sort()
        m = len(requests)
        idx = bisect_left(requests, start)
        f = [[inf, inf] for _ in range(m)]
        for i in range(1, idx + 1):
            for j in range(m - 2, idx - 1, -1):
                if i == 1 and j == m - 2:
                    f[j][0] = f[j][1] = 0
                    continue
                remain = m - 3 - j + i
                f[j][1] = min(f[j][0] + (requests[j] - requests[i - 1]) * remain,  # 往左
                              f[j + 1][1] + (requests[j + 1] - requests[j]) * remain)  # 往右
                f[j][0] = min(f[j][0] + (requests[i] - requests[i - 1]) * remain,  # 往左
                              f[j + 1][1] + (requests[j + 1] - requests[i]) * remain)  # 往右
        return f[idx][0]
```

```java [sol-Java]
class Solution {
    public long elevatorRequests(int n, int start, int[] requests) {
        int m = requests.length + 3;
        int[] a = Arrays.copyOf(requests, m);
        // 插入 start 和两个哨兵
        a[m - 3] = start;
        a[m - 2] = -1;
        a[m - 1] = n;
        Arrays.sort(a);

        int idx = Arrays.binarySearch(a, start);
        long[][] f = new long[m][2];
        for (long[] row : f) {
            Arrays.fill(row, Long.MAX_VALUE / 2);
        }

        for (int i = 1; i <= idx; i++) {
            for (int j = m - 2; j >= idx; j--) {
                if (i == 1 && j == m - 2) {
                    f[j][0] = f[j][1] = 0;
                    continue;
                }
                int remain = m - 3 - j + i;
                f[j][1] = Math.min(f[j][0] + (long) (a[j] - a[i - 1]) * remain, // 往左
                                   f[j + 1][1] + (long) (a[j + 1] - a[j]) * remain); // 往右
                f[j][0] = Math.min(f[j][0] + (long) (a[i] - a[i - 1]) * remain, // 往左
                                   f[j + 1][1] + (long) (a[j + 1] - a[i]) * remain); // 往右
            }
        }
        return f[idx][0];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long elevatorRequests(int n, int start, vector<int>& requests) {
        requests.insert(requests.end(), {start, -1, n}); // 插入 start 和两个哨兵
        ranges::sort(requests);
        int m = requests.size();
        int idx = ranges::lower_bound(requests, start) - requests.begin();
        vector f(m, array<long long, 2>{LLONG_MAX / 2, LLONG_MAX / 2});
        for (int i = 1; i <= idx; i++) {
            for (int j = m - 2; j >= idx; j--) {
                if (i == 1 && j == m - 2) {
                    f[j][0] = f[j][1] = 0;
                    continue;
                }
                int remain = m - 3 - j + i;
                f[j][1] = min(f[j][0] + 1LL * (requests[j] - requests[i - 1]) * remain, // 往左
                              f[j + 1][1] + 1LL * (requests[j + 1] - requests[j]) * remain); // 往右
                f[j][0] = min(f[j][0] + 1LL * (requests[i] - requests[i - 1]) * remain, // 往左
                              f[j + 1][1] + 1LL * (requests[j + 1] - requests[i]) * remain); // 往右
            }
        }
        return f[idx][0];
    }
};
```

```go [sol-Go]
func elevatorRequests(n int, start int, requests []int) int64 {
	requests = append(requests, start, -1, n) // 插入 start 和两个哨兵
	slices.Sort(requests)
	m := len(requests)

	idx := sort.SearchInts(requests, start)
	f := make([][2]int, m)
	for i := range f {
		f[i] = [2]int{math.MaxInt / 2, math.MaxInt / 2}
	}

	for i := 1; i <= idx; i++ {
		for j := m - 2; j >= idx; j-- {
			if i == 1 && j == m-2 {
				f[j] = [2]int{}
				continue
			}
			remain := m - 3 - j + i
			f[j][1] = min(f[j][0]+(requests[j]-requests[i-1])*remain, // 往左
						f[j+1][1]+(requests[j+1]-requests[j])*remain) // 往右
			f[j][0] = min(f[j][0]+(requests[i]-requests[i-1])*remain, // 往左
						f[j+1][1]+(requests[j+1]-requests[i])*remain) // 往右
		}
	}
	return int64(f[idx][0])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m^2)$，其中 $m$ 是 $\textit{requests}$ 的长度。
- 空间复杂度：$\mathcal{O}(m)$。

## 专题训练

见下面动态规划题单的「**八、区间 DP**」。

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
