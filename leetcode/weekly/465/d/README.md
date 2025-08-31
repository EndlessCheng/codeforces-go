## 前置知识

1. **倍数容斥**。请看 [3312. 查询排序后的最大公约数](https://leetcode.cn/problems/sorted-gcd-pair-queries/)，[我的题解](https://leetcode.cn/problems/sorted-gcd-pair-queries/solutions/2940415/mei-ju-rong-chi-qian-zhui-he-er-fen-pyth-ujis/)。
2. **树状数组**。讲解：[带你发明树状数组！附数学证明](https://leetcode.cn/problems/range-sum-query-mutable/solution/dai-ni-fa-ming-shu-zhuang-shu-zu-fu-shu-lyfll/)。

## 转化

设 $U=\max(\textit{nums})$

本题需要对每个 $g=1,2,3,\dots,U$，求出有多少个严格递增子序列的 GCD **恰好**等于 $g$。

我们可以先求出有多少个严格递增子序列的 GCD 是 $g$ 的**倍数**，再利用倍数容斥，转化成「恰好」。

## 分组

对每个 $g=1,2,3,\dots,U$，计算 $\textit{nums}$ 的子序列 $b$，其中每个元素都是 $g$ 的倍数。

这可以在遍历 $\textit{nums}$ 的过程中完成：

- 枚举 $x = \textit{nums}[i]$ 的因子 $d$。
- 把 $x$ 加到一个二维列表 $\textit{groups}[d]$ 中。

## 计算严格递增子序列的个数

枚举 $g$，计算子序列 $b = \textit{groups}[g]$ 的严格递增子序列的个数。

定义 $f[i][x]$ 表示 $b[0]$ 到 $b[i]$ 中的，以 $x$ 结尾的严格递增子序列的个数。⚠**注意**：$x$ 是元素值，不是下标。

如果 $x\ne b[i]$，那么问题变成 $b[0]$ 到 $b[i-1]$ 中的，以 $x$ 结尾的严格递增子序列的个数，此时有 $f[i][x] = f[i-1][x]$。

如果 $x=b[i]$，我们需要知道在 $b[0]$ 到 $b[i-1]$ 中的，以 $y\ (y<x)$ 结尾的严格递增子序列的个数，再算上 $b[i]$ 单独组成一个子序列，得 

$$
f[i][x] = 1 + \sum_{y=1}^{x-1} f[i-1][y]
$$

这可以用值域树状数组优化。

代码实现时，$f$ 数组的第一个维度可以优化掉，并直接记录在树状数组中。

## 写法一

设 $U=\max(\textit{nums})\le 7\times 10^4$。如果每次计算子序列个数时，就创建一棵新的 $\mathcal{O}(U)$ 大小的树状数组，那么光是创建树状数组，就需要 $\mathcal{O}(U^2)$ 的时间，其中 

一个简单的想法是，把数组改成哈希表，这样每次初始化只需要 $\mathcal{O}(1)$ 时间。但哈希表常数太大。

注意到子序列中的元素都是 $g$ 的倍数，当 $g$ 较大时比较离散，可以考虑离散化。

但更简单的做法是，把每个元素除以 $g$。相应地，只需创建 $\mathcal{O}(m/g)$ 大小的树状数组。

注意取模，为什么可以在计算中取模？请看 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

[本题视频讲解](https://www.bilibili.com/video/BV1SMaGz7EXe/?t=26m58s)，欢迎点赞关注~

```py [sol-Python3]
MOD = 1_000_000_007

# 预处理每个数的因子
MX = 70_001
divisors = [[] for _ in range(MX)]
for i in range(1, MX):
    for j in range(i, MX, i):  # 枚举 i 的倍数 j
        divisors[j].append(i)  # i 是 j 的因子


# 完整模板见 https://leetcode.cn/circle/discuss/mOr1u6/
class FenwickTree:
    def __init__(self, n: int):
        self.tree = [0] * (n + 1)  # 使用下标 1 到 n

    # a[i] 增加 val
    # 1 <= i <= n
    # 时间复杂度 O(log n)
    def update(self, i: int, val: int) -> None:
        t = self.tree
        while i < len(t):
            t[i] += val
            i += i & -i

    # 计算前缀和 a[1] + ... + a[i]
    # 1 <= i <= n
    # 时间复杂度 O(log n)
    def pre(self, i: int) -> int:
        t = self.tree
        res = 0
        while i > 0:
            res += t[i]
            i &= i - 1
        return res % MOD


class Solution:
    def totalBeauty(self, nums: List[int]) -> int:
        m = max(nums)

        # 计算 b 的严格递增子序列的个数
        def count_increasing_subsequence(b: List[int], g: int) -> int:
            t = FenwickTree(m // g)
            res = 0
            for x in b:
                x //= g
                # cnt 表示以 x 结尾的严格递增子序列的个数
                cnt = t.pre(x - 1) + 1  # +1 是因为 x 可以一个数组成一个子序列
                res += cnt
                t.update(x, cnt)  # 更新以 x 结尾的严格递增子序列的个数
            return res

        groups = [[] for _ in range(m + 1)]
        for x in nums:
            for d in divisors[x]:
                groups[d].append(x)

        f = [0] * (m + 1)
        ans = 0
        for i in range(m, 0, -1):
            f[i] = count_increasing_subsequence(groups[i], i)
            # 倍数容斥
            for j in range(i * 2, m + 1, i):
                f[i] -= f[j]
            ans += f[i] * i
        return ans % MOD
```

```java [sol-Java]
// 完整模板见 https://leetcode.cn/circle/discuss/mOr1u6/
class FenwickTree {
    private final long[] tree;

    public FenwickTree(int n) {
        tree = new long[n + 1]; // 使用下标 1 到 n
    }

    // a[i] 增加 val
    // 1 <= i <= n
    // 时间复杂度 O(log n)
    public void update(int i, long val) {
        for (; i < tree.length; i += i & -i) {
            tree[i] += val;
        }
    }

    // 求前缀和 a[1] + ... + a[i]
    // 1 <= i <= n
    // 时间复杂度 O(log n)
    public long pre(int i) {
        long res = 0;
        for (; i > 0; i &= i - 1) {
            res += tree[i];
        }
        return res;
    }
}

class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int MX = 70_001;
    private static final List<Integer>[] divisors = new ArrayList[MX];
    private static boolean initialized = false;

    // 这样写比 static block 更快
    private void init() {
        if (initialized) {
            return;
        }
        initialized = true;

        // 预处理每个数的因子
        Arrays.setAll(divisors, _ -> new ArrayList<>());
        for (int i = 1; i < MX; i++) {
            for (int j = i; j < MX; j += i) { // 枚举 i 的倍数 j
                divisors[j].add(i); // i 是 j 的因子
            }
        }
    }

    public int totalBeauty(int[] nums) {
        init();
        int m = 0;
        for (int x : nums) {
            m = Math.max(m, x);
        }

        List<Integer>[] groups = new ArrayList[m + 1];
        Arrays.setAll(groups, _ -> new ArrayList<>());
        for (int x : nums) {
            for (int d : divisors[x]) {
                groups[d].add(x);
            }
        }

        int[] f = new int[m + 1];
        long ans = 0;
        for (int i = m; i > 0; i--) {
            long res = countIncreasingSubsequence(groups[i], i, m);
            // 倍数容斥
            for (int j = i * 2; j <= m; j += i) {
                res -= f[j];
            }
            res %= MOD;
            f[i] = (int) res;
            // m 个 MOD * m 相加，至多为 MOD * m * m，不会超过 64 位整数最大值
            ans += res * i;
        }
        // 保证结果非负
        return (int) ((ans % MOD + MOD) % MOD);
    }

    // 计算 b 的严格递增子序列的个数
    private long countIncreasingSubsequence(List<Integer> b, int g, int m) {
        FenwickTree t = new FenwickTree(m / g);
        long res = 0;
        for (int x : b) {
            x /= g;
            // cnt 表示以 x 结尾的严格递增子序列的个数
            long cnt = t.pre(x - 1) + 1; // +1 是因为 x 可以一个数组成一个子序列
            cnt %= MOD;
            res += cnt;
            t.update(x, cnt); // 更新以 x 结尾的严格递增子序列的个数
        }
        return res;
    }
}
```

```cpp [sol-C++]
const int MOD = 1'000'000'007;
const int MX = 70'001;
vector<int> divisors[MX];

int init = [] {
    // 预处理每个数的因子
    for (int i = 1; i < MX; i++) {
        for (int j = i; j < MX; j += i) { // 枚举 i 的倍数 j
            divisors[j].push_back(i); // i 是 j 的因子
        }
    }
    return 0;
}();

// 完整模板见 https://leetcode.cn/circle/discuss/mOr1u6/
// 根据题目用 FenwickTree<int> t(n) 或者 FenwickTree<long long> t(n) 初始化
template<typename T>
class FenwickTree {
    vector<T> tree;

public:
    // 使用下标 1 到 n
    FenwickTree(int n) : tree(n + 1) {}

    // a[i] 增加 val
    // 1 <= i <= n
    // 时间复杂度 O(log n)
    void update(int i, T val) {
        for (; i < tree.size(); i += i & -i) {
            tree[i] += val;
        }
    }

    // 求前缀和 a[1] + ... + a[i]
    // 1 <= i <= n
    // 时间复杂度 O(log n)
    T pre(int i) const {
        T res = 0;
        for (; i > 0; i &= i - 1) {
            res += tree[i];
        }
        return res;
    }
};

class Solution {
public:
    int totalBeauty(vector<int>& nums) {
        int m = ranges::max(nums);

        // 计算 b 的严格递增子序列的个数
        auto count_increasing_subsequence = [&](vector<int>& b, int g) -> long long {
            FenwickTree<long long> t(m / g);
            long long res = 0;
            for (int x : b) {
                x /= g;
                // cnt 表示以 x 结尾的严格递增子序列的个数
                long long cnt = t.pre(x - 1) + 1; // +1 是因为 x 可以一个数组成一个子序列
                cnt %= MOD;
                res += cnt;
                t.update(x, cnt); // 更新以 x 结尾的严格递增子序列的个数
            }
            return res;
        };

        vector<vector<int>> groups(m + 1);
        for (int x : nums) {
            for (int d : divisors[x]) {
                groups[d].push_back(x);
            }
        }

        vector<int> f(m + 1);
        long long ans = 0;
        for (int i = m; i > 0; i--) {
            long long res = count_increasing_subsequence(groups[i], i);
            // 倍数容斥
            for (int j = i * 2; j <= m; j += i) {
                res -= f[j];
            }
            f[i] = res % MOD;
            // m 个 MOD * m 相加，至多为 MOD * m * m，不会超过 64 位整数最大值
            ans += 1LL * f[i] * i;
        }
        // 保证结果非负
        return (ans % MOD + MOD) % MOD;
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007
const mx = 70_001
var divisors [mx][]int

func init() {
	// 预处理每个数的因子
	for i := 1; i < mx; i++ {
		for j := i; j < mx; j += i { // 枚举 i 的倍数 j
			divisors[j] = append(divisors[j], i) // i 是 j 的因子
		}
	}
}

// 完整模板见 https://leetcode.cn/circle/discuss/mOr1u6/
type fenwick []int

func newFenwickTree(n int) fenwick {
	return make(fenwick, n+1) // 使用下标 1 到 n
}

// a[i] 增加 val
// 1 <= i <= n
// 时间复杂度 O(log n)
func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] += val
	}
}

// 求前缀和 a[1] + ... + a[i]
// 1 <= i <= n
// 时间复杂度 O(log n)
func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return res % mod
}

func totalBeauty(nums []int) (ans int) {
	m := slices.Max(nums)

	// 计算 b 的严格递增子序列的个数
	countIncreasingSubsequence := func(b []int, g int) (res int) {
		t := newFenwickTree(m / g)
		for _, x := range b {
			x /= g
			// cnt 表示以 x 结尾的严格递增子序列的个数
			cnt := t.pre(x-1) + 1 // +1 是因为 x 可以一个数组成一个子序列
			res += cnt
			t.update(x, cnt) // 更新以 x 结尾的严格递增子序列的个数
		}
		return res % mod
	}

	groups := make([][]int, m+1)
	for _, x := range nums {
		for _, d := range divisors[x] {
			groups[d] = append(groups[d], x)
		}
	}

	f := make([]int, m+1)
	for i := m; i > 0; i-- {
		f[i] = countIncreasingSubsequence(groups[i], i)
		// 倍数容斥
		for j := i * 2; j <= m; j += i {
			f[i] -= f[j]
		}
		// 注意 |f[i]| * i < mod * (m / i) * i = mod * m
		// m 个 mod * m 相加，至多为 mod * m * m，不会超过 64 位整数最大值
		ans += f[i] * i
	}
	// 保证结果非负
	return (ans%mod + mod) % mod
}
```

## 写法二：用时间戳懒初始化

这个写法只需创建一次树状数组，**适用于更一般的场景**。

- 初始化时间戳 $\textit{now}=0$。
- 用 $\textit{now}$ 加一代替树状数组的初始化。
- 用 $\textit{time}[i]$ 表示 $i$ 这个位置的最新时间戳。
- 如果更新时发现 $\textit{time}[i] < \textit{now}$，那么这个时候再执行树状数组的初始化，即初始化 $\textit{tree}[i] = 0$，并更新 $\textit{time}[i] = \textit{now}$。
- 查询时只累加 $\textit{time}[i] = \textit{now}$ 的 $\textit{tree}[i]$。

```py [sol-Python3]
MOD = 1_000_000_007

# 预处理每个数的因子
MX = 70_001
divisors = [[] for _ in range(MX)]
for i in range(1, MX):
    for j in range(i, MX, i):  # 枚举 i 的倍数 j
        divisors[j].append(i)  # i 是 j 的因子

class Solution:
    def totalBeauty(self, nums: List[int]) -> int:
        m = max(nums)

        # 树状数组（时间戳优化）
        tree = [0] * (m + 1)
        time = [0] * (m + 1)
        now = 0

        def update(i: int, val: int) -> None:
            nonlocal now
            while i <= m:
                if time[i] < now:
                    time[i] = now
                    tree[i] = 0  # 懒重置
                tree[i] += val
                i += i & -i

        def pre(i: int) -> int:
            res = 0
            while i > 0:
                if time[i] == now:
                    res += tree[i]
                i &= i - 1
            return res % MOD

        # 计算 b 的严格递增子序列的个数
        def count_increasing_subsequence(b: List[int]) -> int:
            nonlocal now
            now += 1  # 重置树状数组（懒重置）
            res = 0
            for x in b:
                # cnt 表示以 x 结尾的严格递增子序列的个数
                cnt = pre(x - 1) + 1  # +1 是因为 x 可以一个数组成一个子序列
                res += cnt
                update(x, cnt)  # 更新以 x 结尾的严格递增子序列的个数
            return res

        groups = [[] for _ in range(m + 1)]
        for x in nums:
            for d in divisors[x]:
                groups[d].append(x)

        f = [0] * (m + 1)
        ans = 0
        for i in range(m, 0, -1):
            f[i] = count_increasing_subsequence(groups[i])
            # 倍数容斥
            for j in range(i * 2, m + 1, i):
                f[i] -= f[j]
            ans += f[i] * i
        return ans % MOD
```

```java [sol-Java]
class FenwickTree {
    private static final int MOD = 1_000_000_007;
    private final int[] tree;
    private final int[] time;
    private int now = 0;

    public FenwickTree(int size) {
        tree = new int[size + 1];
        time = new int[size + 1];
    }

    // 重置树状数组（懒重置）
    public void reset() {
        now++;
    }

    // 把位置 i 的数增加 val
    public void update(int i, int val) {
        while (i < tree.length) {
            if (time[i] < now) {
                time[i] = now;
                tree[i] = 0; // 懒重置
            }
            tree[i] = (tree[i] + val) % MOD;
            i += i & -i;
        }
    }

    // 计算 [1,i] 的元素和
    public int pre(int i) {
        long res = 0;
        while (i > 0) {
            if (time[i] == now) {
                res += tree[i];
            }
            i &= i - 1;
        }
        return (int) (res % MOD);
    }
}

class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int MX = 70_001;
    private static final List<Integer>[] divisors = new ArrayList[MX];
    private static boolean initialized = false;

    // 这样写比 static block 更快
    private void init() {
        if (initialized) {
            return;
        }
        initialized = true;

        // 预处理每个数的因子
        Arrays.setAll(divisors, _ -> new ArrayList<>());
        for (int i = 1; i < MX; i++) {
            for (int j = i; j < MX; j += i) { // 枚举 i 的倍数 j
                divisors[j].add(i); // i 是 j 的因子
            }
        }
    }

    public int totalBeauty(int[] nums) {
        init();
        int m = 0;
        for (int x : nums) {
            m = Math.max(m, x);
        }

        List<Integer>[] groups = new ArrayList[m + 1];
        Arrays.setAll(groups, _ -> new ArrayList<>());
        for (int x : nums) {
            for (int d : divisors[x]) {
                groups[d].add(x);
            }
        }

        FenwickTree t = new FenwickTree(m);
        int[] f = new int[m + 1];
        long ans = 0;
        for (int i = m; i > 0; i--) {
            long res = countIncreasingSubsequence(groups[i], t);
            // 倍数容斥
            for (int j = i * 2; j <= m; j += i) {
                res -= f[j];
            }
            res %= MOD;
            f[i] = (int) res;
            // m 个 MOD * m 相加，至多为 MOD * m * m，不会超过 64 位整数最大值
            ans += res * i;
        }
        // 保证结果非负
        return (int) ((ans % MOD + MOD) % MOD);
    }

    // 计算 b 的严格递增子序列的个数
    private long countIncreasingSubsequence(List<Integer> b, FenwickTree t) {
        t.reset();
        long res = 0;
        for (int x : b) {
            // cnt 表示以 x 结尾的严格递增子序列的个数
            int cnt = t.pre(x - 1) + 1; // +1 是因为 x 可以一个数组成一个子序列
            res += cnt;
            t.update(x, cnt); // 更新以 x 结尾的严格递增子序列的个数
        }
        return res;
    }
}
```

```cpp [sol-C++]
const int MX = 70'001;
vector<int> divisors[MX];

int init = [] {
    // 预处理每个数的因子
    for (int i = 1; i < MX; i++) {
        for (int j = i; j < MX; j += i) { // 枚举 i 的倍数 j
            divisors[j].push_back(i); // i 是 j 的因子
        }
    }
    return 0;
}();

class Solution {
public:
    int totalBeauty(vector<int>& nums) {
        const int MOD = 1'000'000'007;
        int m = ranges::max(nums);

        // 树状数组（时间戳优化）
        vector<int> tree(m + 1), time(m + 1);
        int now = 0;
        auto update = [&](int i, int val) -> void {
            while (i <= m) {
                if (time[i] < now) {
                    time[i] = now;
                    tree[i] = 0; // 懒重置
                }
                tree[i] = (tree[i] + val) % MOD;
                i += i & -i;
            }
        };
        auto pre = [&](int i) -> int {
            long long res = 0;
            while (i > 0) {
                if (time[i] == now) {
                    res += tree[i];
                }
                i &= i - 1;
            }
            return res % MOD;
        };

        // 计算 b 的严格递增子序列的个数
        auto count_increasing_subsequence = [&](vector<int>& b) -> long long {
            now++; // 重置树状数组（懒重置）
            long long res = 0;
            for (int x : b) {
                // cnt 表示以 x 结尾的严格递增子序列的个数
                int cnt = pre(x - 1) + 1; // +1 是因为 x 可以一个数组成一个子序列
                res += cnt;
                update(x, cnt); // 更新以 x 结尾的严格递增子序列的个数
            }
            return res;
        };

        vector<vector<int>> groups(m + 1);
        for (int x : nums) {
            for (int d : divisors[x]) {
                groups[d].push_back(x);
            }
        }

        vector<int> f(m + 1);
        long long ans = 0;
        for (int i = m; i > 0; i--) {
            long long res = count_increasing_subsequence(groups[i]);
            // 倍数容斥
            for (int j = i * 2; j <= m; j += i) {
                res -= f[j];
            }
            f[i] = res % MOD;
            // m 个 MOD * m 相加，至多为 MOD * m * m，不会超过 64 位整数最大值
            ans += 1LL * f[i] * i;
        }
        // 保证结果非负
        return (ans % MOD + MOD) % MOD;
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007
const mx = 70_001

var divisors [mx][]int

func init() {
	// 预处理每个数的因子
	for i := 1; i < mx; i++ {
		for j := i; j < mx; j += i { // 枚举 i 的倍数 j
			divisors[j] = append(divisors[j], i) // i 是 j 的因子
		}
	}
}

func totalBeauty(nums []int) (ans int) {
	m := slices.Max(nums)

	// 树状数组（时间戳优化）
	tree := make([]int, m+1)
	time := make([]int, m+1) // 避免反复初始化树状数组
	now := 0
	update := func(i, val int) {
		for ; i <= m; i += i & -i {
			if time[i] < now {
				time[i] = now
				tree[i] = 0 // 懒重置
			}
			tree[i] += val
		}
	}
	pre := func(i int) (res int) {
		for ; i > 0; i &= i - 1 {
			if time[i] == now {
				res += tree[i]
			}
		}
		return res % mod
	}

	// 计算 b 的严格递增子序列的个数
	countIncreasingSubsequence := func(b []int) (res int) {
		now++ // 重置树状数组（懒重置）
		for _, x := range b {
			// cnt 表示以 x 结尾的严格递增子序列的个数
			cnt := pre(x-1) + 1 // +1 是因为 x 可以一个数组成一个子序列
			res += cnt
			update(x, cnt) // 更新以 x 结尾的严格递增子序列的个数
		}
		return res % mod
	}

	groups := make([][]int, m+1)
	for _, x := range nums {
		for _, d := range divisors[x] {
			groups[d] = append(groups[d], x)
		}
	}

	f := make([]int, m+1)
	for i := m; i > 0; i-- {
		f[i] = countIncreasingSubsequence(groups[i])
		// 倍数容斥
		for j := i * 2; j <= m; j += i {
			f[i] -= f[j]
		}
		// 注意 |f[i]| * i < mod * (m / i) * i = mod * m
		// m 个 mod * m 相加，至多为 mod * m * m，不会超过 64 位整数最大值
		ans += f[i] * i
	}
	// 保证结果非负
	return (ans%mod + mod) % mod
}
```

#### 复杂度分析

预处理的时间和空间不计入。

- 时间复杂度：$\mathcal{O}(Dn\log U + U\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})\le 7\times 10^4$，$D\le 120$ 为单个数的最大因子个数。
- 空间复杂度：$\mathcal{O}(Dn + U)$。

## 相似题目

[3312. 查询排序后的最大公约数](https://leetcode.cn/problems/sorted-gcd-pair-queries/)

## 专题训练

1. 数学题单的「**§1.6 最大公约数（GCD）**」。
2. 数据结构题单的「**§8.1 树状数组**」。
3. 动态规划题单的「**§7.4 合法子序列 DP**」和「**§11.4 树状数组/线段树优化 DP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
