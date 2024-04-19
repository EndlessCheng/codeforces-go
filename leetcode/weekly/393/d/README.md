本题是标准的**划分型 DP**，见 [DP 题单](https://leetcode.cn/circle/discuss/tXLS3i/) 中的「§6.3 约束划分个数」。

定义 $\textit{dfs}(i,j,\textit{and})$ 表示当前考虑到 $\textit{nums}[i]$，已经划分了 $j$ 段，且当前待划分的这一段已经参与 AND 运算的结果为 $\textit{and}$，在这种情况下，继续向后划分，可以得到的最小和。

首先把 $\textit{and}$ 与 $\textit{nums}[i]$ 计算 AND。

用「选或不选」来分类讨论：

- 不划分：继续向后递归 $\textit{dfs}(i+1,j,\textit{and})$。
- 划分：如果 $\textit{and}=\textit{andValues}[j]$，那么划分，即 $\textit{dfs}(i+1,j+1,-1) + \textit{nums}[i]$。
- 这两种情况取最小值。

注：因为 $-1$ 的二进制全为 $1$，与任何数 $x$ 的 AND 都是 $x$，适合初始化。

递归边界：

- 如果 $m-j>n-i$，那么剩余元素无法划分，返回 $\infty$。
- 如果 $j=m$ 且 $i<n$，还有元素没有划分，返回 $\infty$。
- 如果 $j=m$ 且 $i=n$，划分成功，返回 $0$。

递归入口：$\textit{dfs}(0,0,-1)$，即答案。如果答案是 $\infty$ 则返回 $-1$。

请看 [视频讲解](https://www.bilibili.com/video/BV1dJ4m1V7hK/) 第四题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def minimumValueSum(self, nums: List[int], andValues: List[int]) -> int:
        n, m = len(nums), len(andValues)
        @cache
        def dfs(i: int, j: int, and_: int) -> int:
            if m - j > n - i:  # 剩余元素不足
                return inf
            if j == m:  # 分了 m 段
                return 0 if i == n else inf
            and_ &= nums[i]
            if and_ < andValues[j]:  # 剪枝：无法等于 andValues[j]
                return inf
            res = dfs(i + 1, j, and_)  # 不划分
            if and_ == andValues[j]:  # 划分，nums[i] 是这一段的最后一个数
                res = min(res, dfs(i + 1, j + 1, -1) + nums[i])
            return res
        ans = dfs(0, 0, -1)
        return ans if ans < inf else -1
```

```java [sol-Java]
class Solution {
    public int minimumValueSum(int[] nums, int[] andValues) {
        Map<Long, Integer> memo = new HashMap<>();
        int ans = dfs(0, 0, -1, nums, andValues, memo);
        return ans < Integer.MAX_VALUE / 2 ? ans : -1;
    }

    private int dfs(int i, int j, int and, int[] nums, int[] andValues, Map<Long, Integer> memo) {
        int n = nums.length;
        int m = andValues.length;
        if (m - j > n - i) { // 剩余元素不足
            return Integer.MAX_VALUE / 2;
        }
        if (j == m) { // 分了 m 段
            return i == n ? 0 : Integer.MAX_VALUE / 2;
        }
        and &= nums[i];
        if (and < andValues[j]) { // 剪枝：无法等于 andValues[j]
            return Integer.MAX_VALUE / 2;
        }
        long mask = (long) i << 36 | (long) j << 32 | and; // 三个状态压缩成一个 long
        if (memo.containsKey(mask)) {
            return memo.get(mask);
        }
        int res = dfs(i + 1, j, and, nums, andValues, memo); // 不划分
        if (and == andValues[j]) { // 划分，nums[i] 是这一段的最后一个数
            res = Math.min(res, dfs(i + 1, j + 1, -1, nums, andValues, memo) + nums[i]);
        }
        memo.put(mask, res);
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    unordered_map<long long, int> memo;

    int dfs(int i, int j, int and_, vector<int>& nums, vector<int>& andValues) {
        int n = nums.size(), m = andValues.size();
        if (m - j > n - i) { // 剩余元素不足
            return INT_MAX / 2;
        }
        if (j == m) { // 分了 m 段
            return i == n ? 0 : INT_MAX / 2;
        }
        and_ &= nums[i];
        if (and_ < andValues[j]) { // 剪枝：无法等于 andValues[j]
            return INT_MAX / 2;
        }
        long long mask = (long long) i << 36 | (long long) j << 32 | and_; // 三个状态压缩成一个 long long
        if (memo.contains(mask)) {
            return memo[mask];
        }
        int res = dfs(i + 1, j, and_, nums, andValues); // 不划分
        if (and_ == andValues[j]) { // 划分，nums[i] 是这一段的最后一个数
            res = min(res, dfs(i + 1, j + 1, -1, nums, andValues) + nums[i]);
        }
        return memo[mask] = res;
    }

public:
    int minimumValueSum(vector<int>& nums, vector<int>& andValues) {
        int ans = dfs(0, 0, -1, nums, andValues);
        return ans < INT_MAX / 2 ? ans : -1;
    }
};
```

```go [sol-Go]
func minimumValueSum(nums, andValues []int) int {
	n, m := len(nums), len(andValues)
	type args struct{ i, j, and int }
	memo := map[args]int{}
	var dfs func(int, int, int) int
	dfs = func(i, j, and int) int {
		if m-j > n-i { // 剩余元素不足
			return math.MaxInt / 2
		}
		if j == m { // 分了 m 段
			if i == n {
				return 0
			}
			return math.MaxInt / 2
		}
		and &= nums[i]
		if and < andValues[j] { // 剪枝：无法等于 andValues[j]
			return math.MaxInt / 2
		}
		p := args{i, j, and}
		if res, ok := memo[p]; ok {
			return res
		}
		res := dfs(i+1, j, and)  // 不划分
		if and == andValues[j] { // 划分，nums[i] 是这一段的最后一个数
			res = min(res, dfs(i+1, j+1, -1)+nums[i])
		}
		memo[p] = res
		return res
	}
	ans := dfs(0, 0, -1)
	if ans == math.MaxInt/2 {
		return -1
	}
	return ans
}
```

#### 复杂度分析

有多少个状态？

AND 的**性质**：AND 的数越多，AND 的结果就越小。

**定理**：总共有 $\mathcal{O}(n\log U)$ 个不同的子数组 AND 值，其中 $U=\max(\textit{nums})$。

**证明**：考察子数组右端点固定为 $i$ 的情况。我们从 $i$ 开始，向左枚举子数组的左端点 $j$。随着 AND 的数字越来越多，AND 的结果，要么不变，要么减小，且减少一定是某个比特 $1$ 变成 $0$。由于 $\textit{nums}[i]$ 至多有 $\mathcal{O}(\log U)$ 个比特 $1$，所以在右端点固定为 $i$ 的前提下，向左不断 AND $\textit{nums}[j]$ 只能得到 $\mathcal{O}(\log U)$ 个不同的 AND 结果。那么对于所有的子数组，一共有 $\mathcal{O}(n\log U)$ 个不同的子数组 AND 值。

所以对于 $\textit{dfs}$ 中的一个特定的 $i$，只有 $\mathcal{O}(\log U)$ 个不同的 $\textit{and}$。

- 时间复杂度：$\mathcal{O}(nm\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$m$ 为 $\textit{andValues}$ 的长度，$U=\max(\textit{nums})$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(nm\log U)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以动态规划的时间复杂度为 $\mathcal{O}(nm\log U)$。
- 空间复杂度：$\mathcal{O}(nm\log U)$。

### 附：单调队列优化

对于一个固定的 $\textit{andValues}[j]$，当子数组右端点 $i$ 变大时，符合要求（子数组 AND 等于 $\textit{andValues}[j]$）的子数组左端点的范围区间也在右移，所以计算 DP 的转移来源，类似计算 [239. 滑动窗口最大值](https://leetcode.cn/problems/sliding-window-maximum/)（本题是滑动窗口最小值）。

原理请看 [单调队列【基础算法精讲 27】](https://www.bilibili.com/video/BV1bM411X72E/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minimumValueSum(self, nums: List[int], andValues: List[int]) -> int:
        n = len(nums)
        f = [0] + [inf] * n
        new_f = [0] * (n + 1)
        for target in andValues:
            new_f[0] = inf
            a = []
            q = deque()
            qi = 0
            for i, x in enumerate(nums):
                for p in a:
                    p[0] &= x
                a.append([x, i])

                # 原地去重
                j = 1
                for ap, aq in pairwise(a):
                    if ap[0] != aq[0]:
                        a[j] = aq
                        j += 1
                del a[j:]

                # 去掉无用数据
                while a and a[0][0] < target:
                    a.pop(0)

                if a and a[0][0] == target:
                    r = (a[1][1] - 1) if len(a) > 1 else i
                    while qi <= r:
                        while q and f[qi] <= f[q[-1]]:
                            q.pop()
                        q.append(qi)
                        qi += 1
                    while q[0] < a[0][1]:
                        q.popleft()
                    new_f[i + 1] = f[q[0]] + x  # 队首就是最小值
                else:
                    new_f[i + 1] = inf
            f, new_f = new_f, f
        return f[n] if f[n] < inf else -1
```

```go [sol-Go]
func minimumValueSum(nums, andValues []int) int {
	const inf = math.MaxInt / 2
	n := len(nums)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = inf
	}
	newF := make([]int, n+1)
	for _, target := range andValues {
		newF[0] = inf
		type pair struct{ and, l int }
		a := []pair{}
		q := []int{}
		qi := 0
		for i, x := range nums {
			for j := range a {
				a[j].and &= x
			}
			a = append(a, pair{x, i})

			// 原地去重
			j := 1
			for k := 1; k < len(a); k++ {
				if a[k].and != a[k-1].and {
					a[j] = a[k]
					j++
				}
			}
			a = a[:j]

			// 去掉无用数据
			for len(a) > 0 && a[0].and < target {
				a = a[1:]
			}

			if len(a) > 0 && a[0].and == target {
				r := i
				if len(a) > 1 {
					r = a[1].l - 1
				}
				for ; qi <= r; qi++ {
					for len(q) > 0 && f[qi] <= f[q[len(q)-1]] {
						q = q[:len(q)-1]
					}
					q = append(q, qi)
				}
				for q[0] < a[0].l {
					q = q[1:]
				}
				newF[i+1] = f[q[0]] + x // 队首就是最小值
			} else {
				newF[i+1] = inf
			}
		}
		f, newF = newF, f
	}
	if f[n] < inf {
		return f[n]
	}
	return -1
}
```

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
