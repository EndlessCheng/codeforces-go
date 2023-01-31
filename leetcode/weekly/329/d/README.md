# 方法一：O(n^2) 划分型动态规划

### 如何思考

划分出**第一个**子数组，问题变成一个规模更小的子问题。

由于「划分出长为 $x$ 和 $y$ 的子数组」和「划分出长为 $y$ 和 $x$ 的子数组」之后，剩余的子问题是相同的，因此这题适合用动态规划解决。

附：[视频讲解](https://www.bilibili.com/video/BV1Gv4y1y753/)

### 具体算法

定义 $f[i+1]$ 表示划分 $\textit{nums}$ 的前 $i$ 个数的最小代价，从 $i$ 开始倒序枚举最后一个子数组的开始位置 $j$，同时用一个数组 $\textit{state}$ 维护每个元素的出现次数，用一个变量 $\textit{unique}$ 维护只出现一次的元素个数。

具体来说：

- $\textit{state}[x]=0$ 表示 $x$ 出现 $0$ 次；
- $\textit{state}[x]=1$ 表示 $x$ 出现 $1$ 次；
- $\textit{state}[x]=2$ 表示 $x$ 出现超过 $1$ 次。
- 如果 $x$ 首次遇到，那么 $\textit{unique}$ 加一，$\textit{state}[x]=1$；
- 如果 $x$ 第二次遇到，那么 $\textit{unique}$ 减一，$\textit{state}[x]=2$。

> 经测试，这种写法比直接计算特征要更快一些，尤其是 Python。

重要性为子数组的长度减去只出现一次的元素个数加 $k$，即

$$
i-j+1 - \textit{unique}_{j,i} + k
$$

这里 $\textit{unique}_{j,i}$ 表示枚举到 $j$ 时的 $\textit{unique}$ 值。

加上前面子数组的最小代价，所有结果取最小值，得

$$
\begin{aligned}
f[i+1] &= \min\limits_{j=0}^{i} f[j] + i-j+1 - \textit{unique}_{j,i} + k \\ 
       &= i+1+k+ \min\limits_{j=0}^{i} f[j] -j - \textit{unique}_{j,i}
\end{aligned}
$$

初始值 $f[0] = 0$，答案为 $f[n]$。

### 优化

注意到 $f[j]$ 每次都要减去 $j$，而 $f[i+1]$ 最后还要加上 $i+1$，如果定义 $f'[i] = f[i]-i$，则有

$$
f'[i+1] = k+\min\limits_{j=0}^{i} f'[j] - \textit{unique}_{j,i}
$$

由于 $f'[n] = f[n]-n$，所以最后答案为 $f'[n]+n$。

```py [sol1-Python3]
class Solution:
    def minCost(self, nums: List[int], k: int) -> int:
        n = len(nums)
        f = [0] * (n + 1)
        for i in range(n):
            state, unique, mn = [0] * n, 0, inf
            for j in range(i, -1, -1):
                x = nums[j]
                if state[x] == 0:  # 首次出现
                    state[x] = 1
                    unique += 1
                elif state[x] == 1:  # 不再唯一
                    state[x] = 2
                    unique -= 1
                mn = min(mn, f[j] - unique)
                # if f[j]-unique < mn: mn = f[j]-unique  # 手写 min 会快很多
            f[i + 1] = k + mn
        return f[n] + n
```

```java [sol1-Java]
class Solution {
    public int minCost(int[] nums, int k) {
        int n = nums.length;
        int[] f = new int[n + 1];
        byte[] state = new byte[n];
        for (int i = 0; i < n; ++i) {
            Arrays.fill(state, (byte) 0);
            int unique = 0, mn = Integer.MAX_VALUE;
            for (int j = i; j >= 0; --j) {
                int x = nums[j];
                if (state[x] == 0) { // 首次出现
                    state[x] = 1;
                    ++unique;
                } else if (state[x] == 1) { // 不再唯一
                    state[x] = 2;
                    --unique;
                }
                mn = Math.min(mn, f[j] - unique);
            }
            f[i + 1] = k + mn;
        }
        return f[n] + n;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int minCost(vector<int> &nums, int k) {
        int n = nums.size(), f[n + 1];
        f[0] = 0;
        int8_t state[n];
        for (int i = 0; i < n; ++i) {
            memset(state, 0, sizeof(state));
            int unique = 0, mn = INT_MAX;
            for (int j = i; j >= 0; --j) {
                int x = nums[j];
                if (state[x] == 0) state[x] = 1, ++unique; // 首次出现
                else if (state[x] == 1) state[x] = 2, --unique; // 不再唯一
                mn = min(mn, f[j] - unique);
            }
            f[i + 1] = k + mn;
        }
        return f[n] + n;
    }
};
```

```go [sol1-Go]
func minCost(nums []int, k int) int {
	n := len(nums)
	f := make([]int, n+1)
	for i := 0; i < n; i++ {
		state, unique, mn := make([]int8, n), 0, math.MaxInt
		for j := i; j >= 0; j-- {
			x := nums[j]
			if state[x] == 0 { // 首次出现
				state[x] = 1
				unique++
			} else if state[x] == 1 { // 不再唯一
				state[x] = 2
				unique--
			}
			mn = min(mn, f[j]-unique)
		}
		f[i+1] = mn + k
	}
	return f[n] + n
}

func min(a, b int) int { if a > b { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$O(n^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。

# 方法二：线段树优化

记 $x$ 上一次出现的位置是 $\textit{last}[x]$，上上一次出现的位置是 $\textit{last}_2[x]$。

如果从左到右枚举 $x=\textit{nums}[i]$，那么：

- 区间 $[\textit{last}[x]+1,i]$ 内的数的 $\textit{unique}$ 都加一；
- 区间 $[\textit{last}_2[x]+1,\textit{last}[x]]$ 内的数的 $\textit{unique}$ 都减一，相当于把之前的加一撤销掉（如果 $\textit{last}_2[x]$ 不存在则不更新）。

此外，我们求的是一段区间内的 $f[j]-\textit{unique}_{j,i}$ 的最小值，这可以用线段树优化（区间更新，区间查询）。注意 $\textit{unique}_{j,i}$ 前面是负号，所以上面的区间更新值要取反。

代码实现时，由于枚举 $\textit{nums}[i]$ 时，更新的是 $f[i+1]$，我们可以在上一轮循环把它记录下来，在下一轮循环去把它加到线段树中。

```py [sol2-Python3]
class Solution:
    def minCost(self, nums: List[int], k: int) -> int:
        # Lazy 线段树模板（区间加，查询区间最小）
        n = len(nums)
        mn = [0] * (4 * n)
        todo = [0] * (4 * n)

        def do(o: int, v: int) -> None:
            mn[o] += v
            todo[o] += v

        def spread(o: int) -> None:
            v = todo[o]
            if v:
                do(o * 2, v)
                do(o * 2 + 1, v)
                todo[o] = 0

        # 区间 [L,R] 内的数都加上 v   o,l,r=1,1,n
        def update(o: int, l: int, r: int, L: int, R: int, v: int) -> None:
            if L <= l and r <= R:
                do(o, v)
                return
            spread(o)
            m = (l + r) // 2
            if m >= L: update(o * 2, l, m, L, R, v)
            if m < R: update(o * 2 + 1, m + 1, r, L, R, v)
            mn[o] = min(mn[o * 2], mn[o * 2 + 1])

        # 查询区间 [L,R] 的最小值   o,l,r=1,1,n
        def query(o: int, l: int, r: int, L: int, R: int) -> int:
            if L <= l and r <= R:
                return mn[o]
            spread(o)
            m = (l + r) // 2
            if m >= R: return query(o * 2, l, m, L, R)
            if m < L: return query(o * 2 + 1, m + 1, r, L, R)
            return min(query(o * 2, l, m, L, R), query(o * 2 + 1, m + 1, r, L, R))

        ans = 0
        last = [0] * n
        last2 = [0] * n
        for i, x in enumerate(nums, 1):
            update(1, 1, n, i, i, ans)
            update(1, 1, n, last[x] + 1, i, -1)
            if last[x]: update(1, 1, n, last2[x] + 1, last[x], 1)
            ans = k + query(1, 1, n, 1, i)
            last2[x] = last[x]
            last[x] = i
        return ans + n
```

```java [sol2-Java]
class Solution {
    public int minCost(int[] nums, int k) {
        int n = nums.length, ans = 0;
        min = new int[n * 4];
        todo = new int[n * 4];
        int[] last = new int[n], last2 = new int[n];
        for (int i = 1; i <= n; ++i) {
            int x = nums[i - 1];
            update(1, 1, n, i, i, ans);
            update(1, 1, n, last[x] + 1, i, -1);
            if (last[x] > 0) update(1, 1, n, last2[x] + 1, last[x], 1);
            ans = k + query(1, 1, n, 1, i);
            last2[x] = last[x];
            last[x] = i;
        }
        return ans + n;
    }

    // Lazy 线段树模板（区间加，查询区间最小）
    private int[] min, todo;

    private void do_(int o, int v) {
        min[o] += v;
        todo[o] += v;
    }

    private void spread(int o) {
        int v = todo[o];
        if (v != 0) {
            do_(o * 2, v);
            do_(o * 2 + 1, v);
            todo[o] = 0;
        }
    }

    // 区间 [L,R] 内的数都加上 v   o,l,r=1,1,n
    private void update(int o, int l, int r, int L, int R, int v) {
        if (L <= l && r <= R) {
            do_(o, v);
            return;
        }
        spread(o);
        int m = (l + r) / 2;
        if (m >= L) update(o * 2, l, m, L, R, v);
        if (m < R) update(o * 2 + 1, m + 1, r, L, R, v);
        min[o] = Math.min(min[o * 2], min[o * 2 + 1]);
    }

    // 查询区间 [L,R] 的最小值   o,l,r=1,1,n
    private int query(int o, int l, int r, int L, int R) {
        if (L <= l && r <= R)
            return min[o];
        spread(o);
        int m = (l + r) / 2;
        if (m >= R) return query(o * 2, l, m, L, R);
        if (m < L) return query(o * 2 + 1, m + 1, r, L, R);
        return Math.min(query(o * 2, l, m, L, R), query(o * 2 + 1, m + 1, r, L, R));
    }
}
```

```cpp [sol2-C++]
class Solution {
    // Lazy 线段树模板（区间加，查询区间最小）
    int mn[4000], todo[4000], last[1000], last2[1000];

    void do_(int o, int v) {
        mn[o] += v;
        todo[o] += v;
    }

    void spread(int o) {
        int v = todo[o];
        if (v) {
            do_(o * 2, v);
            do_(o * 2 + 1, v);
            todo[o] = 0;
        }
    }

    // 区间 [L,R] 内的数都加上 v   o,l,r=1,1,n
    void update(int o, int l, int r, int L, int R, int v) {
        if (L <= l && r <= R) {
            do_(o, v);
            return;
        }
        spread(o);
        int m = (l + r) / 2;
        if (m >= L) update(o * 2, l, m, L, R, v);
        if (m < R) update(o * 2 + 1, m + 1, r, L, R, v);
        mn[o] = min(mn[o * 2], mn[o * 2 + 1]);
    }

    // 查询区间 [L,R] 的最小值   o,l,r=1,1,n
    int query(int o, int l, int r, int L, int R) {
        if (L <= l && r <= R)
            return mn[o];
        spread(o);
        int m = (l + r) / 2;
        if (m >= R) return query(o * 2, l, m, L, R);
        if (m < L) return query(o * 2 + 1, m + 1, r, L, R);
        return min(query(o * 2, l, m, L, R), query(o * 2 + 1, m + 1, r, L, R));
    }

public:
    int minCost(vector<int> &nums, int k) {
        int n = nums.size(), ans = 0;
        for (int i = 1; i <= n; ++i) {
            int x = nums[i - 1];
            update(1, 1, n, i, i, ans);
            update(1, 1, n, last[x] + 1, i, -1);
            if (last[x]) update(1, 1, n, last2[x] + 1, last[x], 1);
            ans = k + query(1, 1, n, 1, i);
            last2[x] = last[x];
            last[x] = i;
        }
        return ans + n;
    }
};
```

```go [sol2-Go]
// Lazy 线段树模板（区间加，查询区间最小）
type seg []struct{ l, r, min, todo int }

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg) do(o, v int) {
	t[o].min += v
	t[o].todo += v
}

func (t seg) spread(o int) {
	if v := t[o].todo; v != 0 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = 0
	}
}

// 区间 [l,r] 内的数都加上 v   o=1
func (t seg) update(o, l, r, v int) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o, v)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, v)
	}
	if m < r {
		t.update(o<<1|1, l, r, v)
	}
	t[o].min = min(t[o<<1].min, t[o<<1|1].min)
}

// 查询区间 [l,r] 的最小值   o=1
func (t seg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].min
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return min(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func minCost(nums []int, k int) (ans int) {
	n := len(nums)
	last := make([]int, n)
	last2 := make([]int, n)
	t := make(seg, n*4)
	t.build(1, 1, n)
	for i, x := range nums {
		i++ // 线段树区间从 1 开始
		t.update(1, i, i, ans)
		t.update(1, last[x]+1, i, -1)
		if last[x] > 0 {
			t.update(1, last2[x]+1, last[x], 1)
		}
		ans = k + t.query(1, 1, i)
		last2[x] = last[x]
		last[x] = i
	}
	return ans + n
}

func min(a, b int) int { if a > b { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。
