划分型 DP **固定套路**，见 [动态规划题单](https://leetcode.cn/circle/discuss/tXLS3i/) 中的「§6.2 计算划分个数」。

注意本题答案是存在的，因为单个字母是平衡的，我们一定可以划分成 $n$ 个子串。

## 方法一：记忆化搜索

定义 $\textit{dfs}(i)$ 表示划分前缀 $s[0]$ 到 $s[i]$ 的最小划分个数，则有

$$
\textit{dfs}(i) = \min_{j=0}^{i} \textit{dfs}(j-1) + 1
$$

其中 $s[j]$ 到 $s[i]$ 是平衡子串。

我们可以在倒序枚举 $j$ 的同时，用一个哈希表（或者数组）统计每个字符的出现次数。如果子串中每个字母的出现次数都相等，那么子串是平衡的。

**优化**：设子串中有 $k$ 种字母，字母出现次数的最大值为 $\textit{maxCnt}$。子串是平衡的，当且仅当子串长度等于 $k\cdot \textit{maxCnt}$。

递归边界：$\textit{dfs}(-1) = 0$。

递归入口：$\textit{dfs}(n-1)$，即答案。

考虑到整个递归过程中有大量重复递归调用（递归入参相同）。由于递归函数没有副作用，同样的入参无论计算多少次，算出来的结果都是一样的，因此可以用**记忆化搜索**来优化：

- 如果一个状态（递归入参）是第一次遇到，那么可以在返回前，把状态及其结果记到一个 $\textit{memo}$ 数组中。
- 如果一个状态不是第一次遇到（$\textit{memo}$ 中保存的结果不等于 $\textit{memo}$ 的初始值），那么可以直接返回 $\textit{memo}$ 中保存的结果。

**注意**：$\textit{memo}$ 数组的**初始值**一定不能等于要记忆化的值！例如初始值设置为 $0$，并且要记忆化的 $\textit{dfs}(i)$ 也等于 $0$，那就没法判断 $0$ 到底表示第一次遇到这个状态，还是表示之前遇到过了，从而导致记忆化失效。一般把初始值设置为 $-1$。本题由于 $\textit{dfs}(i) > 0$，可以初始化成 $0$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1cz421m786/) 第三题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def minimumSubstringsInPartition(self, s: str) -> int:
        @cache
        def dfs(i: int) -> int:
            if i < 0:
                return 0
            res = inf
            cnt = defaultdict(int)
            max_cnt = 0
            for j in range(i, -1, -1):
                cnt[s[j]] += 1
                max_cnt = max(max_cnt, cnt[s[j]])
                if i - j + 1 == len(cnt) * max_cnt:
                    res = min(res, dfs(j - 1) + 1)
            return res
        return dfs(len(s) - 1)
```

```py [sol-Python3 写法二]
class Solution:
    def minimumSubstringsInPartition(self, s: str) -> int:
        @cache
        def dfs(i: int) -> int:
            if i < 0:
                return 0
            res = inf
            cnt = defaultdict(int)
            max_cnt = 0
            for j in range(i, -1, -1):
                cnt[s[j]] += 1
                # 手动计算 max 和 min
                if cnt[s[j]] > max_cnt:
                    max_cnt = cnt[s[j]]
                if i - j + 1 == len(cnt) * max_cnt:
                    r = dfs(j - 1) + 1
                    if r < res:
                        res = r
            return res
        return dfs(len(s) - 1)
```

```java [sol-Java]
class Solution {
    public int minimumSubstringsInPartition(String S) {
        char[] s = S.toCharArray();
        int n = s.length;
        int[] memo = new int[n];
        Arrays.fill(memo, -1); // -1 表示没有计算过
        return dfs(n - 1, s, memo);
    }

    private int dfs(int i, char[] s, int[] memo) {
        if (i < 0) {
            return 0;
        }
        if (memo[i] != -1) { // 之前计算过
            return memo[i];
        }
        int res = Integer.MAX_VALUE;
        int[] cnt = new int[26];
        int k = 0, maxCnt = 0;
        for (int j = i; j >= 0; j--) {
            k += cnt[s[j] - 'a']++ == 0 ? 1 : 0;
            maxCnt = Math.max(maxCnt, cnt[s[j] - 'a']);
            if (i - j + 1 == k * maxCnt) {
                res = Math.min(res, dfs(j - 1, s, memo) + 1);
            }
        }
        memo[i] = res; // 记忆化
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumSubstringsInPartition(string s) {
        int n = s.length();
        vector<int> memo(n, -1); // -1 表示没有计算过
        function<int(int)> dfs = [&](int i) -> int {
            if (i < 0) {
                return 0;
            }
            int& res = memo[i]; // 注意这里是引用
            if (res != -1) { // 之前计算过
                return res;
            }
            res = INT_MAX;
            int cnt[26]{}, k = 0, max_cnt = 0;
            for (int j = i; j >= 0; j--) {
                k += cnt[s[j] - 'a']++ == 0;
                max_cnt = max(max_cnt, cnt[s[j] - 'a']);
                if (i - j + 1 == k * max_cnt) {
                    res = min(res, dfs(j - 1) + 1);
                }
            }
            return res;
        };
        return dfs(n - 1);
    }
};
```

```go [sol-Go]
func minimumSubstringsInPartition(s string) int {
	n := len(s)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1 // -1 表示没有计算过
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		p := &memo[i]
		if *p != -1 { // 之前计算过
			return *p
		}
		res := math.MaxInt
		cnt := [26]int{}
		k, maxCnt := 0, 0
		for j := i; j >= 0; j-- {
			b := s[j] - 'a'
			if cnt[b] == 0 {
				k++
			}
			cnt[b]++
			maxCnt = max(maxCnt, cnt[b])
			if i-j+1 == k*maxCnt {
				res = min(res, dfs(j-1)+1)
			}
		}
		*p = res // 记忆化
		return res
	}
	return dfs(n - 1)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $s$ 的长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(n)$，单个状态的计算时间为 $\mathcal{O}(n)$，所以动态规划的时间复杂度为 $\mathcal{O}(n^2)$。
- 空间复杂度：$\mathcal{O}(n|\Sigma|)$。其中 $|\Sigma|$ 为字符集合的大小，本题字符均为小写字母，所以 $|\Sigma|=26$。注意递归中至多会创建 $n$ 个长为 $|\Sigma|$ 的 $\textit{cnt}$ 数组。

## 方法二：递推（1:1 翻译）

定义 $f[i+1]$ 表示划分前缀 $s[0]$ 到 $s[i]$ 的最小划分个数，则有

$$
f[i+1] = \min_{j=0}^{i}f[j] + 1
$$

其中 $s[j]$ 到 $s[i]$ 是平衡子串。

初始值 $f[0]= 0$，翻译自递归边界 $\textit{dfs}(-1) = 0$。

答案为 $f[n]$，翻译自递归入口 $\textit{dfs}(n-1)$。

```py [sol-Python3]
class Solution:
    def minimumSubstringsInPartition(self, s: str) -> int:
        n = len(s)
        f = [0] + [inf] * n
        for i in range(n):
            cnt = defaultdict(int)
            max_cnt = 0
            for j in range(i, -1, -1):
                cnt[s[j]] += 1
                max_cnt = max(max_cnt, cnt[s[j]])
                if i - j + 1 == len(cnt) * max_cnt:
                    f[i + 1] = min(f[i + 1], f[j] + 1)
        return f[n]
```

```py [sol-Python3 写法二]
class Solution:
    def minimumSubstringsInPartition(self, s: str) -> int:
        n = len(s)
        f = [0] + [inf] * n
        for i in range(n):
            cnt = defaultdict(int)
            max_cnt = 0
            for j in range(i, -1, -1):
                cnt[s[j]] += 1
                if cnt[s[j]] > max_cnt:
                    max_cnt = cnt[s[j]]
                if i - j + 1 == len(cnt) * max_cnt and f[j] + 1 < f[i + 1]:
                    f[i + 1] = f[j] + 1
        return f[n]
```

```java [sol-Java]
class Solution {
    public int minimumSubstringsInPartition(String S) {
        char[] s = S.toCharArray();
        int n = s.length;
        int[] f = new int[n + 1];
        Arrays.fill(f, Integer.MAX_VALUE);
        f[0] = 0;
        int[] cnt = new int[26];
        for (int i = 0; i < n; i++) {
            Arrays.fill(cnt, 0);
            int k = 0, maxCnt = 0;
            for (int j = i; j >= 0; j--) {
                k += cnt[s[j] - 'a']++ == 0 ? 1 : 0;
                maxCnt = Math.max(maxCnt, cnt[s[j] - 'a']);
                if (i - j + 1 == k * maxCnt) {
                    f[i + 1] = Math.min(f[i + 1], f[j] + 1);
                }
            }
        }
        return f[n];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumSubstringsInPartition(string s) {
        int n = s.length();
        vector<int> f(n + 1, INT_MAX);
        f[0] = 0;
        for (int i = 0; i < n; i++) {
            int cnt[26]{}, k = 0, max_cnt = 0;
            for (int j = i; j >= 0; j--) {
                k += cnt[s[j] - 'a']++ == 0;
                max_cnt = max(max_cnt, cnt[s[j] - 'a']);
                if (i - j + 1 == k * max_cnt) {
                    f[i + 1] = min(f[i + 1], f[j] + 1);
                }
            }
        }
        return f[n];
    }
};
```

```go [sol-Go]
func minimumSubstringsInPartition(s string) int {
	n := len(s)
	f := make([]int, n+1)
	for i := range s {
		f[i+1] = math.MaxInt
		cnt := [26]int{}
		k, maxCnt := 0, 0
		for j := i; j >= 0; j-- {
			b := s[j] - 'a'
			if cnt[b] == 0 {
				k++
			}
			cnt[b]++
			maxCnt = max(maxCnt, cnt[b])
			if i-j+1 == k*maxCnt {
				f[i+1] = min(f[i+1], f[j]+1)
			}
		}
	}
	return f[n]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n + |\Sigma|)$。其中 $|\Sigma|$ 为字符集合的大小，本题字符均为小写字母，所以 $|\Sigma|=26$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
