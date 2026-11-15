本题实际上是两个独立的问题：

1. 如何判断子数组每个数的频率（出现次数）都是偶数？
2. 如何判断子数组恰好有 $k$ 个不同元素？

## 子数组每个数的频率都是偶数

把 $\textit{nums}[i]$ 映射成一个 $[0,2^{64}-1]$ 内的随机整数。由于同一个数异或两次等于 $0$，异或偶数次也等于 $0$，所以「子数组每个数的频率都是偶数」等价于「**子数组的异或和为** $0$」。这个技巧叫做**异或哈希**（XOR hashing）。

计算任意子数组的异或和可以用 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)。

> **注**：在 $q=10^5$ 的情况下，至少有一个询问算错的概率约为 $5.4\times 10^{-15}$，足以通过本题。详见文末的分析。

## 子数组恰好有 k 个不同元素

下面介绍两种方法。

### 方法一：滑动窗口（在线回答询问）

本题 $k$ 是固定的，我们可以在枚举子数组右端点 $r$ 的同时，维护子数组左端点的范围 $[\ell_1, \ell_2)$，使得子数组内恰好有 $k$ 个不同元素。

可以用**滑动窗口**计算 $\ell_1$ 和 $\ell_2$，原理请看视频[【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

可以记录每个 $r$ 对应的 $\ell_1$ 和 $\ell_2$，保存在数组中。

对于询问 $[\ell,r]$，判断 $\ell$ 是否在区间 $[\ell_1[r], \ell_2[r])$ 中。

[本题视频讲解](https://www.bilibili.com/video/BV18p846TEwX/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def validSubarrays(self, nums: list[int], k: int, queries: list[list[int]]) -> list[bool]:
        n = len(nums)
        s = [0] * (n + 1)
        # 把 nums[i] 映射成随机数
        # 不要在 [0, 2**L) 内随机，见 https://codeforces.com/blog/entry/153335
        mp = defaultdict(lambda: randrange(10 ** 18))
        for i, x in enumerate(nums):
            s[i + 1] = s[i] ^ mp[x]

        def calc_left(k: int) -> list[int]:
            lefts = [0] * n
            cnt = defaultdict(int)
            l = 0
            for i, x in enumerate(nums):
                cnt[x] += 1
                while len(cnt) >= k:
                    v = nums[l]
                    if cnt[v] > 1:
                        cnt[v] -= 1
                    else:
                        del cnt[v]  # 保证 len(cnt) 是窗口内的不同元素个数
                    l += 1
                lefts[i] = l
            return lefts

        l1 = calc_left(k + 1)
        l2 = calc_left(k)

        return [s[r + 1] == s[l] and l1[r] <= l < l2[r] for l, r in queries]
```

```java [sol-Java]
class Solution {
    private static final Random random = new Random();

    public boolean[] validSubarrays(int[] nums, int k, int[][] queries) {
        int n = nums.length;
        long[] sum = new long[n + 1];
        Map<Integer, Long> hash = new HashMap<>();
        for (int i = 0; i < n; i++) {
            // 把 nums[i] 映射成一个随机的 long
            long randVal = hash.computeIfAbsent(nums[i], _ -> random.nextLong());
            sum[i + 1] = sum[i] ^ randVal;
        }

        int[] l1 = calcLeft(nums, k + 1);
        int[] l2 = calcLeft(nums, k);

        boolean[] ans = new boolean[queries.length];
        for (int i = 0; i < queries.length; i++) {
            int l = queries[i][0];
            int r = queries[i][1];
            ans[i] = sum[r + 1] == sum[l] && l1[r] <= l && l < l2[r];
        }
        return ans;
    }

    private int[] calcLeft(int[] nums, int k) {
        int n = nums.length;
        int[] lefts = new int[n];
        Map<Integer, Integer> cnt = new HashMap<>();
        int l = 0;
        for (int i = 0; i < n; i++) {
            cnt.merge(nums[i], 1, Integer::sum); // ++cnt[nums[i]]
            while (cnt.size() >= k) {
                int c = cnt.merge(nums[l], -1, Integer::sum); // c = --cnt[nums[l]]
                if (c == 0) {
                    cnt.remove(nums[l]); // 保证 cnt.size() 是窗口内的不同元素个数
                }
                l++;
            }
            lefts[i] = l;
        }
        return lefts;
    }
}
```

```cpp [sol-C++]
class Solution {
    static inline mt19937_64 rng = mt19937_64(chrono::steady_clock::now().time_since_epoch().count());

public:
    vector<bool> validSubarrays(vector<int>& nums, int k, vector<vector<int>>& queries) {
        int n = nums.size();
        vector<uint64_t> sum(n + 1);
        unordered_map<int, uint64_t> hash;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            // 把 nums[i] 映射成一个随机的 uint64_t
            if (!hash.contains(x)) {
                hash[x] = rng() / 3; // https://codeforces.com/blog/entry/153335
            }
            sum[i + 1] = sum[i] ^ hash[x];
        }

        auto calc_left = [&](int k) -> vector<int> {
            vector<int> lefts(n);
            unordered_map<int, int> cnt;
            int l = 0;
            for (int i = 0; i < n; i++) {
                cnt[nums[i]]++;
                while (cnt.size() >= k) {
                    auto it = cnt.find(nums[l]);
                    if (--it->second == 0) {
                        cnt.erase(it); // 保证 cnt.size() 是窗口内的不同元素个数
                    }
                    l++;
                }
                lefts[i] = l;
            }
            return lefts;
        };

        auto l1 = calc_left(k + 1);
        auto l2 = calc_left(k);

        vector<bool> ans(queries.size());
        for (int i = 0; i < queries.size(); i++) {
            auto& q = queries[i];
            int l = q[0], r = q[1];
            ans[i] = sum[r + 1] == sum[l] && l1[r] <= l && l < l2[r];
        }
        return ans;
    }
};
```

```go [sol-Go]
func validSubarrays(nums []int, k int, queries [][]int) []bool {
	n := len(nums)
	sum := make([]uint64, n+1)
	hash := map[int]uint64{}
	for i, x := range nums {
		// 把 nums[i] 映射成一个随机的 uint64
		if _, ok := hash[x]; !ok {
			hash[x] = rand.Uint64()
		}
		sum[i+1] = sum[i] ^ hash[x]
	}

	calcLeft := func(k int) []int {
		lefts := make([]int, n)
		cnt := map[int]int{}
		l := 0
		for i, x := range nums {
			cnt[x]++
			for len(cnt) >= k {
				v := nums[l]
				if cnt[v] > 1 {
					cnt[v]--
				} else {
					delete(cnt, v) // 保证 len(cnt) 是窗口内的不同元素个数
				}
				l++
			}
			lefts[i] = l
		}
		return lefts
	}

	l1 := calcLeft(k + 1)
	l2 := calcLeft(k)

	ans := make([]bool, len(queries))
	for i, p := range queries {
		l, r := p[0], p[1]
		ans[i] = sum[r+1] == sum[l] && l1[r] <= l && l < l2[r]
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+q)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。对于滑动窗口，虽然写了个二重循环，但是内层循环中对 $\ell$ 加一的**总**执行次数不会超过 $n$ 次，所以二重循环的循环次数为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

### 方法二：树状数组（离线回答询问）

更通用的做法是，直接求出子数组有多少个不同元素。即使题目把 $k$ 放在询问中，即 $\textit{queries}[i] = [l_i, r_i, k_i]$，这个做法也适用。

做法同 [HH 的项链](https://www.luogu.com.cn/problem/P1972)，原理如下：

从左到右遍历 $\textit{nums}$，同时维护元素 $x$ 最近一次出现的下标 $\textit{last}[x]$。如果固定子数组右端点为 $i$，要想让子数组包含某个元素 $x$，左端点必须 $\le \textit{last}[x]$。把最近一次出现的 $x$ 视作 $1$，其余更靠左的 $x$ 视作 $0$，那么「子数组的不同元素个数」等价于「**子数组的和**」。由于数组元素会动态变化，需要用 [树状数组](https://leetcode.cn/problems/range-sum-query-mutable/solution/dai-ni-fa-ming-shu-zhuang-shu-zu-fu-shu-lyfll/) 维护。

以 $\textit{nums}=[1,1,2,2,3,2]$ 为例：

- 遍历到 $i=0$，把 $\textit{nums}$ 视作 $[1,*,*,*,*,*]$。其中 $*$ 表示尚未遍历到的元素，不考虑。
- 遍历到 $i=1$，把 $\textit{nums}$ 视作 $[0,1,*,*,*,*]$。
- 遍历到 $i=2$，把 $\textit{nums}$ 视作 $[0,1,1,*,*,*]$。
- 遍历到 $i=3$，把 $\textit{nums}$ 视作 $[0,1,0,1,*,*]$。
- 遍历到 $i=4$，把 $\textit{nums}$ 视作 $[0,1,0,1,1,*]$。
- 遍历到 $i=5$，把 $\textit{nums}$ 视作 $[0,1,0,0,1,1]$。

为了能在遍历到 $\textit{nums}[i]$ 时，回答 $r=i$ 的询问，我们需要提前把询问按照右端点分组，然后回答这一组内的询问。这个技巧叫做**离线询问**。

```py [sol-Python3]
# 模板来源 https://leetcode.cn/discuss/post/3583665/
class FenwickTree:
    def __init__(self, n: int):
        self.tree = [0] * (n + 1)  # 使用下标 1 到 n

    # a[i] 增加 val
    # 时间复杂度 O(log n)
    def update(self, i: int, val: int) -> None:
        t = self.tree
        i += 1
        while i < len(t):
            t[i] += val
            i += i & -i

    # 计算前缀和 a[1] + ... + a[i]
    # 时间复杂度 O(log n)
    def pre(self, i: int) -> int:
        t = self.tree
        res = 0
        i += 1
        while i > 0:
            res += t[i]
            i &= i - 1
        return res

    # 计算区间和 a[l] + ... + a[r]
    # 时间复杂度 O(log n)
    def query(self, l: int, r: int) -> int:
        return self.pre(r) - self.pre(l - 1)


class Solution:
    def validSubarrays(self, nums: list[int], k: int, queries: list[list[int]]) -> list[bool]:
        n = len(nums)
        s = [0] * (n + 1)
        # 把 nums[i] 映射成随机数
        # 不要在 [0, 2**L) 内随机，见 https://codeforces.com/blog/entry/153335
        mp = defaultdict(lambda: randrange(10 ** 18))
        for i, x in enumerate(nums):
            s[i + 1] = s[i] ^ mp[x]

        # 离线询问：按照右端点分组
        groups = [[] for _ in range(n)]
        for i, (l, r) in enumerate(queries):
            groups[r].append((l, i))

        t = FenwickTree(n)
        last = {}
        ans = [False] * len(queries)
        for r, x in enumerate(nums):
            if x in last:
                t.update(last[x], -1)  # 撤销
            last[x] = r
            t.update(r, 1)
            for l, qid in groups[r]:
                ans[qid] = s[r + 1] == s[l] and t.query(l, r) == k
        return ans
```

```java [sol-Java]
// 模板来源 https://leetcode.cn/discuss/post/3583665/
class FenwickTree {
    private final int[] tree;

    public FenwickTree(int n) {
        tree = new int[n + 1]; // 使用下标 1 到 n
    }

    // a[i] 增加 val
    // 时间复杂度 O(log n)
    public void update(int i, int val) {
        for (i++; i < tree.length; i += i & -i) {
            tree[i] += val;
        }
    }

    // 求前缀和 a[1] + ... + a[i]
    // 时间复杂度 O(log n)
    public int pre(int i) {
        int res = 0;
        for (i++; i > 0; i &= i - 1) {
            res += tree[i];
        }
        return res;
    }

    // 求区间和 a[l] + ... + a[r]
    // 时间复杂度 O(log n)
    public int query(int l, int r) {
        return pre(r) - pre(l - 1);
    }
}

class Solution {
    private static final Random random = new Random();

    public boolean[] validSubarrays(int[] nums, int k, int[][] queries) {
        int n = nums.length;
        long[] sum = new long[n + 1];
        Map<Integer, Long> hash = new HashMap<>();
        for (int i = 0; i < n; i++) {
            // 把 nums[i] 映射成一个随机的 long
            long randVal = hash.computeIfAbsent(nums[i], _ -> random.nextLong());
            sum[i + 1] = sum[i] ^ randVal;
        }

        // 离线询问：按照右端点分组
        List<int[]>[] groups = new ArrayList[n];
        Arrays.setAll(groups, _ -> new ArrayList<>());
        for (int i = 0; i < queries.length; i++) {
            int[] q = queries[i];
            groups[q[1]].add(new int[]{q[0], i});
        }

        FenwickTree t = new FenwickTree(n);
        Map<Integer, Integer> last = HashMap.newHashMap(hash.size()); // 预分配空间
        boolean[] ans = new boolean[queries.length];
        for (int r = 0; r < n; r++) {
            int x = nums[r];
            if (last.containsKey(x)) {
                t.update(last.get(x), -1); // 撤销
            }
            last.put(x, r);
            t.update(r, 1);
            for (int[] p : groups[r]) {
                int l = p[0];
                ans[p[1]] = sum[r + 1] == sum[l] && t.query(l, r) == k;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
// 模板来源 https://leetcode.cn/discuss/post/3583665/
// 根据题目用 FenwickTree<int> t(n) 或者 FenwickTree<long long> t(n) 初始化
template<typename T>
class FenwickTree {
    vector<T> tree;

public:
    // 使用下标 1 到 n
    FenwickTree(int n) : tree(n + 1) {}

    // a[i] 增加 val
    // 时间复杂度 O(log n)
    void update(int i, T val) {
        for (i++; i < tree.size(); i += i & -i) {
            tree[i] += val;
        }
    }

    // 求前缀和 a[1] + ... + a[i]
    // 时间复杂度 O(log n)
    T pre(int i) const {
        T res{};
        for (i++; i > 0; i &= i - 1) {
            res += tree[i];
        }
        return res;
    }

    // 求区间和 a[l] + ... + a[r]
    // 时间复杂度 O(log n)
    T query(int l, int r) const {
        return pre(r) - pre(l - 1);
    }
};

class Solution {
    static inline mt19937_64 rng = mt19937_64(chrono::steady_clock::now().time_since_epoch().count());

public:
    vector<bool> validSubarrays(vector<int>& nums, int k, vector<vector<int>>& queries) {
        int n = nums.size();
        vector<uint64_t> sum(n + 1);
        unordered_map<int, uint64_t> hash;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            // 把 nums[i] 映射成一个随机的 uint64_t
            if (!hash.contains(x)) {
                hash[x] = rng() / 3; // https://codeforces.com/blog/entry/153335
            }
            sum[i + 1] = sum[i] ^ hash[x];
        }

        // 离线询问：按照右端点分组
        vector<vector<pair<int, int>>> groups(n);
        for (int i = 0; i < queries.size(); i++) {
            auto& q = queries[i];
            groups[q[1]].emplace_back(q[0], i);
        }

        FenwickTree<int> t(n);
        unordered_map<int, int> last;
        last.reserve(hash.size()); // 预分配空间
        vector<bool> ans(queries.size());
        for (int r = 0; r < n; r++) {
            int x = nums[r];
            if (last.contains(x)) {
                t.update(last[x], -1); // 撤销
            }
            last[x] = r;
            t.update(r, 1);
            for (auto& [l, qid] : groups[r]) {
                ans[qid] = sum[r + 1] == sum[l] && t.query(l, r) == k;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
// 模板来源 https://leetcode.cn/discuss/post/3583665/
type fenwick []int

func newFenwickTree(n int) fenwick {
	return make(fenwick, n+1) // 使用下标 1 到 n
}

// a[i] 增加 val
// 时间复杂度 O(log n)
func (f fenwick) update(i, val int) {
	for i++; i < len(f); i += i & -i {
		f[i] += val
	}
}

// 求前缀和 a[1] + ... + a[i]
// 时间复杂度 O(log n)
func (f fenwick) pre(i int) (res int) {
	for i++; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

// 求区间和 a[l] + ... + a[r]
// 时间复杂度 O(log n)
func (f fenwick) query(l, r int) int {
	return f.pre(r) - f.pre(l-1)
}

func validSubarrays(nums []int, k int, queries [][]int) []bool {
	n := len(nums)
	sum := make([]uint64, n+1)
	hash := map[int]uint64{}
	for i, x := range nums {
		// 把 nums[i] 映射成一个随机的 uint64
		if _, ok := hash[x]; !ok {
			hash[x] = rand.Uint64()
		}
		sum[i+1] = sum[i] ^ hash[x]
	}

	// 离线询问：按照右端点分组
	type pair struct{ l, qid int }
	groups := make([][]pair, n)
	for i, q := range queries {
		groups[q[1]] = append(groups[q[1]], pair{q[0], i})
	}

	t := newFenwickTree(n)
	last := make(map[int]int, len(hash)) // 预分配空间
	ans := make([]bool, len(queries))
	for r, x := range nums {
		if i, ok := last[x]; ok {
			t.update(i, -1) // 撤销
		}
		last[x] = r
		t.update(r, 1)
		for _, p := range groups[r] {
			ans[p.qid] = sum[r+1] == sum[p.l] && t.query(p.l, r) == k
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + q\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n + q)$。

## 附：算错的概率是多少？

使用异或哈希，算错的概率是多少？即子数组有元素出现奇数次，但异或和为 $0$。

设 $m$ 是子数组的出现奇数次的元素种数。在随机情况下，每一位是互相独立的。单独讨论某个比特位。$m$ 个 $\{0,1\}$ 内的随机整数异或，有 $2^m$ 种情况，这可以与 $m$ 个元素的子集一一对应（选表示 $1$，不选表示 $0$）。由于从 $m$ 个元素中选奇数个元素的方案数是 $2^{m-1}$，所以异或和为 $1$ 的概率为 $\dfrac{2^{m-1}}{2^m} = \dfrac{1}{2}$，异或和为 $0$ 的概率也为 $\dfrac{1}{2}$。证明见 [从 n 个数中选奇数个数的方案数](https://zhuanlan.zhihu.com/p/1909852852114948837)。所以异或结果的每个比特位都是 $\{0,1\}$ 内的随机整数。总的来看，$m$ 个 $[0,2^{64}-1]$ 内的随机整数的异或和仍然可以视作一个 $[0,2^{64}-1]$ 内的随机整数。

如果 $m = 0$，没有出现奇数次的元素，那么异或和一定是 $0$，一定算对。

如果 $m\ge 1$，那么（单次询问）算错的概率为

$$
P(m\ 个数的异或和 = 0) = \dfrac{1}{2^{64}}
$$

在有 $q$ 个询问的情况下，即使所有询问都互相独立，我们有

$$
P(至少一个询问算错) = 1 - P(所有询问都算对) = 1 - (1 - 2^{-64})^q \le 1 - (1 - q\cdot 2^{-64}) = q\cdot 2^{-64}
$$

这里用到了伯努利不等式 $(1+x)^q\ge 1+qx$。

在 $q=10^5$ 的情况下，至少有一个询问算错的概率约为 $5.4\times 10^{-15}$，足以通过本题。

## 思考题

把题干中的「偶数」改成「奇数」，怎么做？

解答见 [本题视频讲解](https://www.bilibili.com/video/BV18p846TEwX/)。

## 专题训练

1. 滑动窗口题单的「**二、不定长滑动窗口**」。
2. 数据结构题单的「**§8.1 树状数组**」和「**专题：离线算法**」。

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
