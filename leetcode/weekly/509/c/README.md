## 方法一：枚举质因子 + 最大子数组和

由于题目要求 $k>1$，我们先特判 $\textit{nums}$ 只包含 $1$ 的情况，此时最优解是只选一个 $1$，分数差为 $-1$，最小 $k$ 为 $2$（见示例 3）。

否则，选 $\textit{nums}[i]$ 的因子作为 $k$ 更好（否则分数差一定是负数）。我们可以先收集所有因子，去重，再枚举因子，这样可以避免重复计算。

进一步地，如果一个数是 $x$ 的倍数，那么也是 $x$ 的质因子的倍数，所以我们**只需枚举质因子**，这样倍数更多。即使倍数相同，质因子作为 $k$ 也更小，符合题目要求。

对于一个固定的 $k$，设 $x=\textit{nums}[i]$，如果 $x$ 不是 $k$ 的倍数，则视作 $-x$（减去 Bob 的分数）。问题变成 [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/)，请看 [我的题解](https://leetcode.cn/problems/maximum-subarray/solutions/2533977/qian-zhui-he-zuo-fa-ben-zhi-shi-mai-mai-abu71/)。

[本题视频讲解](https://www.bilibili.com/video/BV1ioTC6BECj/?t=13m11s)，欢迎点赞关注~

```py [sol-Python3]
# 预处理每个数的质因子
MX = 1_000_001
prime_divisors = [[] for _ in range(MX)]
for i in range(2, MX):
    if not prime_divisors[i]:  # i 是质数
        for j in range(i, MX, i):  # 枚举 i 的倍数 j
            prime_divisors[j].append(i)  # i 是 j 的质因子


class Solution:
    # 53. 最大子数组和（如果 nums[i] 不是 k 的倍数，则视作 -nums[i]）
    def maxSubArray(self, nums: list[int], k: int) -> int:
        ans = -inf
        f = 0
        for x in nums:
            f = max(f, 0) + (-x if x % k else x)
            ans = max(ans, f)
        return ans

    def divisibleGame(self, nums: list[int]) -> int:
        MOD = 1_000_000_007

        # 收集所有质因子
        all_prime_divisors = []
        for x in nums:
            all_prime_divisors += prime_divisors[x]

        if not all_prime_divisors:
            # 每个数都是 1
            # 最优是只选一个 1（分数差为 -1），最小 k 为 2（见示例 3）
            return MOD - 2

        # 排序去重
        all_prime_divisors = sorted(set(all_prime_divisors))

        max_diff = -inf
        best_k = 0
        # 枚举质因子作为 k，计算最大子数组和
        for k in all_prime_divisors:
            diff = self.maxSubArray(nums, k)
            if diff > max_diff:
                max_diff = diff
                best_k = k

        return max_diff * best_k % MOD
```

```java [sol-Java]
class Solution {
    public int divisibleGame(int[] nums) {
        final int MOD = 1_000_000_007;

        // 收集所有质因子
        // 预处理有些慢，改成不预处理的写法
        List<Integer> allPrimeDivisors = new ArrayList<>();
        for (int x : nums) {
            for (int p = 2; p * p <= x; p++) {
                if (x % p == 0) {
                    allPrimeDivisors.add(p);
                    do {
                        x /= p;
                    } while (x % p == 0);
                }
            }
            if (x > 1) {
                allPrimeDivisors.add(x);
            }
        }

        if (allPrimeDivisors.isEmpty()) {
            // 每个数都是 1
            // 最优是只选一个 1（分数差为 -1），最小 k 为 2（见示例 3）
            return MOD - 2;
        }

        Collections.sort(allPrimeDivisors);

        int maxDiff = Integer.MIN_VALUE;
        int bestK = 0;
        int preK = 0;
        // 枚举质因子作为 k，计算最大子数组和
        for (int k : allPrimeDivisors) {
            if (k == preK) {
                continue;
            }
            int diff = maxSubArray(nums, k);
            if (diff > maxDiff) {
                maxDiff = diff;
                bestK = k;
            }
            preK = k;
        }

        return (int) ((long) maxDiff * bestK % MOD);
    }

    // 53. 最大子数组和（如果 nums[i] 不是 k 的倍数，则视作 -nums[i]）
    private int maxSubArray(int[] nums, int k) {
        int ans = Integer.MIN_VALUE;
        int f = 0;
        for (int x : nums) {
            f = Math.max(f, 0) + (x % k == 0 ? x : -x);
            ans = Math.max(ans, f);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
constexpr int MX = 1'000'001;
vector<int> prime_divisors[MX];

// 预处理每个数的质因子
int init = [] {
    for (int i = 2; i < MX; i++) {
        if (prime_divisors[i].empty()) { // i 是质数
            for (int j = i; j < MX; j += i) { // 枚举 i 的倍数 j
                prime_divisors[j].push_back(i); // i 是 j 的因子
            }
        }
    }
    return 0;
}();

class Solution {
    // 53. 最大子数组和（如果 nums[i] 不是 k 的倍数，则视作 -nums[i]）
    int maxSubArray(vector<int>& nums, int k) {
        int ans = INT_MIN;
        int f = 0;
        for (int x : nums) {
            f = max(f, 0) + (x % k ? -x : x);
            ans = max(ans, f);
        }
        return ans;
    }

public:
    int divisibleGame(vector<int>& nums) {
        constexpr int MOD = 1'000'000'007;

        // 收集所有质因子
        vector<int> all_prime_divisors;
        for (int x : nums) {
            auto& pd = prime_divisors[x];
            all_prime_divisors.insert(all_prime_divisors.end(), pd.begin(), pd.end());
        }

        if (all_prime_divisors.empty()) {
            // 每个数都是 1
            // 最优是只选一个 1（分数差为 -1），最小 k 为 2（见示例 3）
            return MOD - 2;
        }

        // 排序去重
        ranges::sort(all_prime_divisors);
        all_prime_divisors.erase(ranges::unique(all_prime_divisors).begin(), all_prime_divisors.end());

        int max_diff = INT_MIN;
        int best_k = 0;
        // 枚举质因子作为 k，计算最大子数组和
        for (int k : all_prime_divisors) {
            int diff = maxSubArray(nums, k);
            if (diff > max_diff) {
                max_diff = diff;
                best_k = k;
            }
        }

        return 1LL * max_diff * best_k % MOD;
    }
};
```

```go [sol-Go]
const mx = 1_000_001
var primeDivisors [mx][]int32

// 预处理每个数的质因子
func init() {
	for i := int32(2); i < mx; i++ {
		if primeDivisors[i] == nil { // i 是质数
			for j := i; j < mx; j += i { // 枚举 i 的倍数 j
				primeDivisors[j] = append(primeDivisors[j], i) // i 是 j 的质因子
			}
		}
	}
}

// 53. 最大子数组和（如果 nums[i] 不是 k 的倍数，则视作 -nums[i]）
func maxSubArray(nums []int, k int) int {
	ans := math.MinInt
	f := 0
	for _, x := range nums {
		if x%k != 0 {
			x = -x
		}
		f = max(f, 0) + x
		ans = max(ans, f)
	}
	return ans
}

func divisibleGame(nums []int) (ans int) {
	const mod = 1_000_000_007

	// 收集所有质因子
	allPrimeDivisors := []int32{}
	for _, x := range nums {
		allPrimeDivisors = append(allPrimeDivisors, primeDivisors[x]...)
	}

	if len(allPrimeDivisors) == 0 {
		// 每个数都是 1
		// 最优是只选一个 1（分数差为 -1），最小 k 为 2（见示例 3）
		return mod - 2
	}

	// 排序去重
	slices.Sort(allPrimeDivisors)
	allPrimeDivisors = slices.Compact(allPrimeDivisors)

	maxDiff, bestK := math.MinInt, 0
	// 枚举质因子作为 k，计算最大子数组和
	for _, d := range allPrimeDivisors {
		k := int(d)
		diff := maxSubArray(nums, k)
		if diff > maxDiff {
			maxDiff, bestK = diff, k
		}
	}

	return maxDiff * bestK % mod
}
```

#### 复杂度分析

忽略预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(n^2\log\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。由于我们去重了，最坏情况是 $\textit{nums}$ 的所有元素互不相同。在 $n$ 个不同元素的情况下，平均每个数有 $\mathcal{O}(\log\log U)$ 个质因子（理由同埃式筛），一共有 $\mathcal{O}(n\log\log U)$ 个质因子。这是宽松的估计，重复统计了相同的质因子，去重后可能更小。
- 空间复杂度：$\mathcal{O}(n\log\log U)$。

## 方法二：线段树维护动态最大子数组和

这个方法可以解决 $n\le 10^5$ 的数据范围。

先把所有 $\textit{nums}[i]$ 都变成 $-\textit{nums}[i]$。

还是枚举质因子作为 $k$。对于一个固定的 $k$，我们需要把 $\textit{nums}$ 中是 $k$ 的倍数的那些数，从 $-\textit{nums}[i]$ 还原成 $\textit{nums}[i]$，然后计算 $\textit{nums}$ 的最大子数组和。

这是一个单点修改的**动态最大子数组和**问题，见 [P4513 小白逛公园](https://www.luogu.com.cn/problem/P4513)。

```py [sol-Python3]
# 预处理每个数的质因子
MX = 1_000_001
prime_divisors = [[] for _ in range(MX)]
for i in range(2, MX):
    if not prime_divisors[i]:  # i 是质数
        for j in range(i, MX, i):  # 枚举 i 的倍数 j
            prime_divisors[j].append(i)  # i 是 j 的质因子


class Data:
    __slots__ = "sum", "pre", "suf", "ans"

    def __init__(self, v=0):
        self.sum = self.pre = self.suf = self.ans = v


class SegmentTree:
    def __init__(self, a: list[int]) -> None:
        n = len(a)
        self.t = [Data() for _ in range(2 << (n - 1).bit_length())]
        self.build(a, 1, 0, n - 1)

    def set(self, node: int, v: int) -> None:
        self.t[node] = Data(v)

    def maintain(self, node: int) -> None:
        lo, ro = self.t[node * 2], self.t[node * 2 + 1]
        cur = self.t[node]
        cur.sum = lo.sum + ro.sum
        cur.pre = max(lo.pre, lo.sum + ro.pre)
        cur.suf = max(ro.suf, ro.sum + lo.suf)
        cur.ans = max(lo.ans, ro.ans, lo.suf + ro.pre)

    def build(self, a: list[int], node: int, l: int, r: int) -> None:
        if l == r:
            self.set(node, -a[l])
            return
        m = (l + r) // 2
        self.build(a, node * 2, l, m)
        self.build(a, node * 2 + 1, m + 1, r)
        self.maintain(node)

    def update(self, node: int, l: int, r: int, i: int, val: int) -> None:
        if l == r:
            self.set(node, val)
            return
        m = (l + r) // 2
        if i <= m:
            self.update(node * 2, l, m, i, val)
        else:
            self.update(node * 2 + 1, m + 1, r, i, val)
        self.maintain(node)


class Solution:
    def divisibleGame(self, nums: list[int]) -> int:
        MOD = 1_000_000_007
        prime_divisors_to_indices = defaultdict(list)
        for i, x in enumerate(nums):
            for d in prime_divisors[x]:
                prime_divisors_to_indices[d].append(i)

        if not prime_divisors_to_indices:
            # 每个数都是 1
            # 最优是只选一个 1（分数差为 -1），最小 k 为 2（见示例 3）
            return MOD - 2

        n = len(nums)
        t = SegmentTree(nums)
        max_diff = -inf
        best_k = 0

        # 枚举质因子作为 k，计算最大子数组和
        for k, indices in prime_divisors_to_indices.items():
            for i in indices:
                # nums[i] 是质因子 k 的倍数
                t.update(1, 0, n - 1, i, nums[i])

            diff = t.t[1].ans
            if diff > max_diff or diff == max_diff and k < best_k:
                max_diff = diff
                best_k = k

            for i in indices:
                t.update(1, 0, n - 1, i, -nums[i])

        return max_diff * best_k % MOD
```

```java [sol-Java]
class SegmentTree {
    private static class Data {
        int sum, pre, suf, ans;

        Data() {
        }

        Data(int v) {
            this.sum = v;
            this.pre = v;
            this.suf = v;
            this.ans = v;
        }
    }

    private final Data[] tree;

    public SegmentTree(int[] a) {
        int n = a.length;
        int size = 2 << (32 - Integer.numberOfLeadingZeros(n - 1));
        tree = new Data[size];
        Arrays.setAll(tree, _ -> new Data());
        build(a, 1, 0, n - 1);
    }

    private void maintain(int node) {
        Data cur = tree[node];
        Data lo = tree[node * 2];
        Data ro = tree[node * 2 + 1];
        cur.sum = lo.sum + ro.sum;
        cur.pre = Math.max(lo.pre, lo.sum + ro.pre);
        cur.suf = Math.max(ro.suf, ro.sum + lo.suf);
        cur.ans = Math.max(Math.max(lo.ans, ro.ans), lo.suf + ro.pre);
    }

    private void build(int[] a, int node, int l, int r) {
        if (l == r) {
            tree[node] = new Data(-a[l]);
            return;
        }
        int m = (l + r) >>> 1;
        build(a, node * 2, l, m);
        build(a, node * 2 + 1, m + 1, r);
        maintain(node);
    }

    public void update(int node, int l, int r, int i, int val) {
        if (l == r) {
            tree[node] = new Data(val);
            return;
        }
        int m = (l + r) >>> 1;
        if (i <= m) {
            update(node * 2, l, m, i, val);
        } else {
            update(node * 2 + 1, m + 1, r, i, val);
        }
        maintain(node);
    }

    public int query() {
        return tree[1].ans;
    }
}

class Solution {
    public int divisibleGame(int[] nums) {
        final int MOD = 1_000_000_007;

        int n = nums.length;
        Map<Integer, List<Integer>> primeToIndices = new HashMap<>();
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            for (int p = 2; p * p <= x; p++) {
                if (x % p == 0) {
                    primeToIndices.computeIfAbsent(p, _ -> new ArrayList<>()).add(i);
                    do {
                        x /= p;
                    } while (x % p == 0);
                }
            }
            if (x > 1) {
                primeToIndices.computeIfAbsent(x, _ -> new ArrayList<>()).add(i);
            }
        }

        if (primeToIndices.isEmpty()) {
            // 每个数都是 1
            // 最优是只选一个 1（分数差为 -1），最小 k 为 2（见示例 3）
            return MOD - 2;
        }

        SegmentTree t = new SegmentTree(nums);
        int maxDiff = Integer.MIN_VALUE;
        int bestK = 0;

        // 枚举质因子作为 k，计算最大子数组和
        for (Map.Entry<Integer, List<Integer>> e : primeToIndices.entrySet()) {
            int k = e.getKey();
            List<Integer> indices = e.getValue();

            for (int i : indices) {
                // nums[i] 是质因子 k 的倍数
                t.update(1, 0, n - 1, i, nums[i]);
            }

            int diff = t.query();
            if (diff > maxDiff || diff == maxDiff && k < bestK) {
                maxDiff = diff;
                bestK = k;
            }

            for (int i : indices) {
                t.update(1, 0, n - 1, i, -nums[i]);
            }
        }

        return (int) ((long) maxDiff * bestK % MOD);
    }
}
```

```cpp [sol-C++]
constexpr int MX = 1'000'001;
vector<int> prime_divisors[MX];

// 预处理每个数的质因子
int init = [] {
    for (int i = 2; i < MX; i++) {
        if (prime_divisors[i].empty()) { // i 是质数
            for (int j = i; j < MX; j += i) { // 枚举 i 的倍数 j
                prime_divisors[j].push_back(i); // i 是 j 的因子
            }
        }
    }
    return 0;
}();

class SegmentTree {
    struct Data {
        int sum, pre, suf, ans;

        Data(int v = 0) : sum(v), pre(v), suf(v), ans(v) {}
    };

    vector<Data> tree;

    void maintain(int node) {
        Data& cur = tree[node];
        Data& lo = tree[node * 2];
        Data& ro = tree[node * 2 + 1];
        cur.sum = lo.sum + ro.sum;
        cur.pre = max(lo.pre, lo.sum + ro.pre);
        cur.suf = max(ro.suf, ro.sum + lo.suf);
        cur.ans = max(max(lo.ans, ro.ans), lo.suf + ro.pre);
    }

public:
    SegmentTree(const vector<int>& a) : tree(2 << bit_width(a.size() - 1)) {
        build(a, 1, 0, a.size() - 1);
    }

    void build(const vector<int>& a, int node, int l, int r) {
        if (l == r) {
            tree[node] = Data(-a[l]);
            return;
        }
        int m = (l + r) >> 1;
        build(a, node * 2, l, m);
        build(a, node * 2 + 1, m + 1, r);
        maintain(node);
    }

    void update(int node, int l, int r, int i, int val) {
        if (l == r) {
            tree[node] = Data(val);
            return;
        }
        int m = (l + r) >> 1;
        if (i <= m) {
            update(node * 2, l, m, i, val);
        } else {
            update(node * 2 + 1, m + 1, r, i, val);
        }
        maintain(node);
    }

    int query() const {
        return tree[1].ans;
    }
};

class Solution {
public:
    int divisibleGame(vector<int>& nums) {
        constexpr int MOD = 1'000'000'007;

        int n = nums.size();
        unordered_map<int, vector<int>> prime_to_indices;
        for (int i = 0; i < n; i++) {
            for (int d : prime_divisors[nums[i]]) {
                prime_to_indices[d].push_back(i);
            }
        }

        if (prime_to_indices.empty()) {
            // 每个数都是 1
            // 最优是只选一个 1（分数差为 -1），最小 k 为 2（见示例 3）
            return MOD - 2;
        }

        SegmentTree t(nums);
        int max_diff = INT_MIN;
        int best_k = 0;

        // 枚举质因子作为 k，计算最大子数组和
        for (auto& [k, indices] : prime_to_indices) {
            for (int i : indices) {
                // nums[i] 是质因子 k 的倍数
                t.update(1, 0, n - 1, i, nums[i]);
            }

            int diff = t.query();
            if (diff > max_diff || diff == max_diff && k < best_k) {
                max_diff = diff;
                best_k = k;
            }

            for (int i : indices) {
                t.update(1, 0, n - 1, i, -nums[i]);
            }
        }

        return 1LL * max_diff * best_k % MOD;
    }
};
```

```go [sol-Go]
const mx = 1_000_001
var primeDivisors [mx][]int32

// 预处理每个数的质因子
func init() {
	for i := int32(2); i < mx; i++ {
		if primeDivisors[i] == nil { // i 是质数
			for j := i; j < mx; j += i { // 枚举 i 的倍数 j
				primeDivisors[j] = append(primeDivisors[j], i) // i 是 j 的质因子
			}
		}
	}
}

type data struct {
	sum, pre, suf, ans int
}

type seg []data

func (t seg) set(node, v int) {
	t[node] = data{v, v, v, v}
}

func (t seg) maintain(node int) {
	lo, ro := t[node*2], t[node*2+1]
	t[node].sum = lo.sum + ro.sum
	t[node].pre = max(lo.pre, lo.sum+ro.pre)
	t[node].suf = max(ro.suf, ro.sum+lo.suf)
	t[node].ans = max(lo.ans, ro.ans, lo.suf+ro.pre)
}

func (t seg) build(a []int, node, l, r int) {
	if l == r {
		t.set(node, -a[l])
		return
	}
	m := (l + r) >> 1
	t.build(a, node*2, l, m)
	t.build(a, node*2+1, m+1, r)
	t.maintain(node)
}

func (t seg) update(node, l, r, i, val int) {
	if l == r {
		t.set(node, val)
		return
	}
	m := (l + r) >> 1
	if i <= m {
		t.update(node*2, l, m, i, val)
	} else {
		t.update(node*2+1, m+1, r, i, val)
	}
	t.maintain(node)
}

func divisibleGame(nums []int) (ans int) {
	const mod = 1_000_000_007

	primeDivisorsToIndices := map[int32][]int{}
	for i, x := range nums {
		for _, d := range primeDivisors[x] {
			primeDivisorsToIndices[d] = append(primeDivisorsToIndices[d], i)
		}
	}

	if len(primeDivisorsToIndices) == 0 {
		// 每个数都是 1
		// 最优是只选一个 1（分数差为 -1），最小 k 为 2
		return mod - 2
	}

	n := len(nums)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(nums, 1, 0, n-1)
	maxDiff, bestK := math.MinInt, int32(0)

	// 枚举质因子作为 k，计算最大子数组和
	for k, indices := range primeDivisorsToIndices {
		for _, i := range indices {
			// nums[i] 是质因子 k 的倍数
			t.update(1, 0, n-1, i, nums[i])
		}

		diff := t[1].ans
		if diff > maxDiff || diff == maxDiff && k < bestK {
			maxDiff, bestK = diff, k
		}

		for _, i := range indices {
			t.update(1, 0, n-1, i, -nums[i])
		}
	}

	return maxDiff * int(bestK) % mod
}
```

#### 复杂度分析

忽略预处理的时间和空间。

- 时间复杂度：$\mathcal{O}\left(\dfrac{n\log n\log U}{\log\log U} \right)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。最坏情况是所有元素都相同，每个数至多有 $\mathcal{O}\left(\dfrac{\log U}{\log\log U}\right)$ 个不同的质因子。
- 空间复杂度：$\mathcal{O}\left(\dfrac{n\log U}{\log\log U}\right)$。

## 专题训练

1. 数学题单的「**§1.3 质因数分解**」和「**§1.5 因子**」。
2. 动态规划题单的「**§1.3 最大子数组和**」。
3. 数据结构题单的「**§8.3 线段树**」。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
