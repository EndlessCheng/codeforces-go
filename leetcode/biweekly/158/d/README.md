## 初步分析

题目要求子集中 $0$ 到 $9$ 每个数字在所有数的数位中最多出现一次。

把 $\textit{vals}[i]$ 的数位保存到集合 $V_i$ 中（不考虑有重复数位的数字），比如数字 $130$ 对应集合 $\{0,1,3\}$。

题目要求转化成：

- 从子树中选择若干没有交集的集合，对应的 $\textit{vals}[i]$ 之和越大越好。

## 方法一：枚举子集

定义 $f_x[S]$ 表示从 $x$ 子树中选择若干没有交集的集合，这些集合的并集为 $S$ 的情况下，对应的 $\textit{vals}[i]$ 之和的最大值。

枚举 $S$ 的非空真子集 $T$，把集合 $S$ 视作 $T$ 和 $\complement_ST$ 的并集（这两个集合是两个规模更小的子问题），那么：

- 集合的并集为 $T$ 的情况下，对应的 $\textit{vals}[i]$ 之和的最大值，即 $f_x[T]$。
- 集合的并集为 $\complement_ST$ 的情况下，对应的 $\textit{vals}[i]$ 之和的最大值，即 $f_x[\complement_ST]$。

二者相加，更新 $f_x[S]$ 的最大值，即

$$
f_x[S] = \max_{T\subseteq S} f_x[T] + f_x[\complement_ST]
$$

初始值：

- 选 $\textit{vals}[x]$，初始化 $f_x[V_x] = \textit{vals}[x]$。
- 枚举 $x$ 的儿子 $y$，由于同一个集合至多选一个（否则就有交集了），所以取最大值得 $f_x[S] = \max\limits_y f_y[S]$。

$\max(f_x)$ 就是题目要求的 $\textit{maxScore}[x]$，加到答案中。

**代码实现时，用二进制表示集合，用位运算实现集合操作，具体请看** [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/circle/discuss/CaOJ45/)。

⚠**注意**：这个方法很慢，更快的做法见方法二。

```py [sol-Python3]
# 超时了！请看方法二！
class Solution:
    def goodSubtreeSum(self, vals: List[int], par: List[int]) -> int:
        D = 10
        n = len(par)
        g = [[] for _ in range(n)]
        for i in range(1, n):
            g[par[i]].append(i)

        ans = 0

        def dfs(x: int) -> List[int]:
            max_score = [0] * (1 << D)

            # 计算 vals[x] 的数位集合 mask
            mask = 0
            v = vals[x]
            while v:
                v, d = divmod(v, D)
                if mask >> d & 1:
                    break
                mask |= 1 << d
            else:  # 没有中途 break
                max_score[mask] = vals[x]

            # 同一个 mask 至多选一个，直接取 max
            for y in g[x]:
                fy = dfs(y)
                for i, s in enumerate(fy):
                    max_score[i] = max(max_score[i], s)

            for i in range(3, 1 << D):
                # 枚举 i 的非空真子集 sub
                sub = (i - 1) & i
                while sub:
                    max_score[i] = max(max_score[i], max_score[i ^ sub] + max_score[sub])
                    sub = (sub - 1) & i

            nonlocal ans
            ans += max(max_score)
            return max_score

        dfs(0)
        return ans % 1_000_000_007
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int D = 10;
    private long ans = 0;

    public int goodSubtreeSum(int[] vals, int[] par) {
        int n = par.length;
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int i = 1; i < n; i++) {
            g[par[i]].add(i);
        }
        dfs(0, vals, g);
        return (int) (ans % MOD);
    }

    private int[] dfs(int x, int[] vals, List<Integer>[] g) {
        int[] f = new int[1 << D];

        // 计算 vals[x] 的数位集合 mask
        int mask = 0;
        for (int v = vals[x]; v > 0; v /= D) {
            int d = v % D;
            if ((mask >> d & 1) > 0) { // d 在集合 mask 中
                mask = 0; // 不符合要求
                break;
            }
            mask |= 1 << d; // 把 d 加到集合 mask 中
        }

        if (mask > 0) {
            f[mask] = vals[x];
        }

        // 同一个集合 i 至多选一个，直接取 max
        for (int y : g[x]) {
            int[] fy = dfs(y, vals, g);
            for (int i = 0; i < f.length; i++) {
                f[i] = Math.max(f[i], fy[i]);
            }
        }

        int mx = 0;
        for (int i = 3; i < f.length; i++) {
            // 枚举集合 i 的非空真子集 sub
            for (int sub = i & (i - 1); sub > 0; sub = (sub - 1) & i) {
                f[i] = Math.max(f[i], f[sub] + f[i ^ sub]);
            }
            mx = Math.max(mx, f[i]);
        }
        ans += mx;

        return f;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int goodSubtreeSum(vector<int>& vals, vector<int>& par) {
        const int MOD = 1'000'000'007;
        const int D = 10;
        int n = par.size();
        vector<vector<int>> g(n);
        for (int i = 1; i < n; i++) {
            g[par[i]].push_back(i);
        }

        long long ans = 0;

        auto dfs = [&](this auto&& dfs, int x) -> array<int, 1 << D> {
            array<int, 1 << D> f{};

            // 计算 vals[x] 的数位集合 mask
            int mask = 0;
            for (int v = vals[x]; v > 0; v /= D) {
                int d = v % D;
                if (mask >> d & 1) { // d 在集合 mask 中
                    mask = 0; // 不符合要求
                    break;
                }
                mask |= 1 << d; // 把 d 加到集合 mask 中
            }

            if (mask > 0) {
                f[mask] = vals[x];
            }

            // 同一个集合 i 至多选一个，直接取 max
            for (int y : g[x]) {
                auto fy = dfs(y);
                for (int i = 0; i < (1 << D); i++) {
                    f[i] = max(f[i], fy[i]);
                }
            }

            for (int i = 3; i < (1 << D); i++) {
                // 枚举集合 i 的非空真子集 sub
                for (int sub = i & (i - 1); sub > 0; sub = (sub - 1) & i) {
                    f[i] = max(f[i], f[sub] + f[i ^ sub]);
                }
            }

            ans += ranges::max(f);
            return f;
        };

        dfs(0);
        return ans % MOD;
    }
};
```

```go [sol-Go]
func goodSubtreeSum(vals, par []int) (ans int) {
	const mod = 1_000_000_007
	const D = 10
	n := len(par)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		p := par[i]
		g[p] = append(g[p], i)
	}

	var dfs func(int) [1 << D]int
	dfs = func(x int) (f [1 << D]int) {
		// 计算 vals[x] 的数位集合 mask
		mask := 0
		for v := vals[x]; v > 0; v /= D {
			d := v % D
			if mask>>d&1 > 0 { // d 在集合 mask 中
				mask = 0 // 不符合要求
				break
			}
			mask |= 1 << d // 把 d 加到集合 mask 中
		}

		if mask > 0 {
			f[mask] = vals[x]
		}

		// 同一个集合 i 至多选一个，直接取 max
		for _, y := range g[x] {
			fy := dfs(y)
			for i, sum := range fy {
				f[i] = max(f[i], sum)
			}
		}

		for i := range f {
			// 枚举集合 i 的非空真子集 sub
			for sub := i & (i - 1); sub > 0; sub = (sub - 1) & i {
				f[i] = max(f[i], f[sub]+f[i^sub])
			}
		}

		ans += slices.Max(f[:])
		return
	}
	dfs(0)
	return ans % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\cdot 3^D)$，其中 $n$ 是 $\textit{vals}$ 的长度，$D=10$。大小为 $D$ 的集合的大小为 $m$ 的子集有 $\binom D m$ 个，子集的子集有 $2^m$ 个，根据二项式定理，$\sum\limits_{m=0}^D \binom D m 2^m = (2+1)^D = 3^D$，所以二重循环的时间复杂度为 $\mathcal{O}(3^D)$。
- 空间复杂度：$\mathcal{O}(n\cdot 2^D)$。

## 方法二：合并

类似 [3562. 折扣价交易股票的最大利润](https://leetcode.cn/problems/maximum-profit-from-trading-stocks-with-discounts/)，把 $f_y$ 合并到 $f_x$ 中：

- 枚举 $f_y$ 中的集合 $S$ 和 $f_x$ 中的集合 $T$。
- 如果 $S$ 和 $T$ 交集为空，那么（刷表法）用 $f_y[S] + f_x[T]$ 更新 $f_x[S\cup T]$ 的最大值。

此外用 $f_y[S]$ 更新 $f_x[T]$ 的最大值。

用哈希表实现，相比数组可以避免遍历很多无效状态。

```py [sol-Python3]
class Solution:
    def goodSubtreeSum(self, vals: List[int], par: List[int]) -> int:
        n = len(par)
        g = [[] for _ in range(n)]
        for i in range(1, n):
            g[par[i]].append(i)

        ans = 0

        def dfs(x: int) -> Dict[int, int]:
            f = defaultdict(int)

            # 计算 vals[x] 的数位集合 mask
            mask = 0
            v = vals[x]
            while v:
                v, d = divmod(v, 10)
                if mask >> d & 1:
                    break
                mask |= 1 << d
            else:  # 没有中途 break
                f[mask] = vals[x]

            for y in g[x]:
                fy = dfs(y)
                nf = f.copy()
                for msk, s in fy.items():
                    if s <= nf[msk]:  # 重要优化！无法让最大值变大
                        continue
                    nf[msk] = s
                    # 求两个 mask 的并集，刷表转移
                    for msk2, s2 in f.items():
                        if msk & msk2 == 0:
                            nf[msk | msk2] = max(nf[msk | msk2], s + s2)
                f = nf

            if f:
                nonlocal ans
                ans += max(f.values())
            return f

        dfs(0)
        return ans % 1_000_000_007
```

```java [sol-Java]
class Solution {
    private long ans = 0;

    public int goodSubtreeSum(int[] vals, int[] par) {
        int n = par.length;
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int i = 1; i < n; i++) {
            g[par[i]].add(i);
        }

        dfs(0, vals, g);
        return (int) (ans % 1_000_000_007);
    }

    // 返回每种 mask 下的最大子树和
    private Map<Integer, Integer> dfs(int x, int[] vals, List<Integer>[] g) {
        Map<Integer, Integer> f = new HashMap<>();

        // 计算 vals[x] 的数位集合 mask
        int mask = 0;
        for (int v = vals[x]; v > 0; v /= 10) {
            int d = v % 10;
            if ((mask >> d & 1) > 0) { // d 在集合 mask 中
                mask = 0; // 不符合要求
                break;
            }
            mask |= 1 << d; // 把 d 加到集合 mask 中
        }

        if (mask > 0) {
            f.put(mask, vals[x]);
        }

        for (int y : g[x]) {
            Map<Integer, Integer> fy = dfs(y, vals, g);
            Map<Integer, Integer> nf = new HashMap<>(f);
            for (Map.Entry<Integer, Integer> e : fy.entrySet()) {
                int msk = e.getKey();
                int s = e.getValue();
                if (s <= nf.getOrDefault(msk, 0)) { // 重要优化！无法让最大值变大
                    continue;
                }
                nf.put(msk, s);
                // 求两个 mask 的并集，刷表转移
                for (Map.Entry<Integer, Integer> e2 : f.entrySet()) {
                    int msk2 = e2.getKey();
                    if ((msk & msk2) == 0) {
                        nf.merge(msk | msk2, s + e2.getValue(), Math::max);
                    }
                }
            }
            f = nf;
        }

        if (!f.isEmpty()) {
            ans += Collections.max(f.values());
        }

        return f;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int goodSubtreeSum(vector<int>& vals, vector<int>& par) {
        const int MOD = 1'000'000'007;
        const int D = 10;
        int n = par.size();
        vector<vector<int>> g(n);
        for (int i = 1; i < n; i++) {
            g[par[i]].push_back(i);
        }

        long long ans = 0;

        auto dfs = [&](this auto&& dfs, int x) -> unordered_map<int, int> {
            unordered_map<int, int> f;

            // 计算 vals[x] 的数位集合 mask
            int mask = 0;
            for (int v = vals[x]; v > 0; v /= D) {
                int d = v % D;
                if (mask >> d & 1) { // d 在集合 mask 中
                    mask = 0; // 不符合要求
                    break;
                }
                mask |= 1 << d; // 把 d 加到集合 mask 中
            }

            if (mask > 0) {
                f[mask] = vals[x];
            }

            for (int y : g[x]) {
                auto fy = dfs(y);
                auto nf = f;
                for (auto& [msk, s] : fy) {
                    if (s <= nf[msk]) { // 重要优化！无法让最大值变大
                        continue;
                    }
                    nf[msk] = s;
                    // 求两个 mask 的并集，刷表转移
                    for (auto& [msk2, s2] : f) {
                        if ((msk & msk2) == 0) {
                            nf[msk | msk2] = max(nf[msk | msk2], s + s2);
                        }
                    }
                }
                f = move(nf);
            }

            int mx = 0;
            for (auto& [_, s] : f) {
                mx = max(mx, s);
            }
            ans += mx;

            return f;
        };

        dfs(0);
        return ans % MOD;
    }
};
```

```go [sol-Go]
func goodSubtreeSum(vals, par []int) (ans int) {
	const mod = 1_000_000_007
	const D = 10
	n := len(par)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		p := par[i]
		g[p] = append(g[p], i)
	}

	var dfs func(int) map[int]int
	dfs = func(x int) map[int]int {
		f := map[int]int{}

		// 计算 vals[x] 的数位集合 mask
		mask := 0
		for v := vals[x]; v > 0; v /= D {
			d := v % D
			if mask>>d&1 > 0 { // d 在集合 mask 中
				mask = 0 // 不符合要求
				break
			}
			mask |= 1 << d // 把 d 加到集合 mask 中
		}

		if mask > 0 {
			f[mask] = vals[x]
		}

		for _, y := range g[x] {
			fy := dfs(y)
			nf := maps.Clone(f)
			for msk, s := range fy {
				if s <= nf[msk] { // 重要优化！无法让最大值变大
					continue
				}
				nf[msk] = s
				// 求两个 mask 的并集，刷表转移
				for msk2, s2 := range f {
					if msk&msk2 == 0 {
						nf[msk|msk2] = max(nf[msk|msk2], s+s2)
					}
				}
			}
			f = nf
		}

		mx := 0
		for _, s := range f {
			mx = max(mx, s)
		}
		ans += mx

		return f
	}
	dfs(0)
	return ans % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\cdot 4^D)$，其中 $n$ 是 $\textit{vals}$ 的长度，$D=10$。有 $n-1$ 条边，每条边都有一次 $\mathcal{O}(4^D)$ 的状态转移。
- 空间复杂度：$\mathcal{O}(n\cdot 2^D)$。

## 方法三：树上启发式合并

方法二是把子树 $y$ 的 $f_y$ 的每个集合 $S$，合并到 $f_x$ 中。每次需要遍历 $\mathcal{O}(2^D)$ 个集合 $S$。

我们也可以把子树 $y$ 中的每个 $V_i$，一个一个地添加到 $f_x$ 中。其中集合 $V_i$ 是 $\textit{vals}[i]$ 的数位集合。利用**树上启发式合并**，均摊地说，每次只需要遍历 $\mathcal{O}(\log n)$ 个集合。

树上启发式合并有两种实现，一种是按照 [轻重儿子](https://oi-wiki.org/graph/dsu-on-tree/) 合并，另一种是按照子树大小合并，后者更简单。

> **注**：从运行时间上看，可能读者觉得方法二和方法三没有区别，这是因为方法二中的「重要优化」起到很大作用。读者可以去掉方法二和方法三中的「重要优化」，同时把 `nf[msk] = s` 改成 `nf[msk] = max(nf[msk], s)`，这样就能看出这两个方法的区别。但对于 C++ 来说，这个写法似乎是个负优化。

```py [sol-Python3]
class Solution:
    def goodSubtreeSum(self, vals: List[int], par: List[int]) -> int:
        n = len(par)
        g = [[] for _ in range(n)]
        for i in range(1, n):
            g[par[i]].append(i)

        ans = 0

        def dfs(x: int) -> (Dict[int, int], List[Tuple[int, int]]):
            f = defaultdict(int)
            single = []

            # 计算 vals[x] 的数位集合 mask
            mask = 0
            v = val = vals[x]
            while v:
                v, d = divmod(v, 10)
                if mask >> d & 1:
                    break
                mask |= 1 << d
            else:  # 没有中途 break
                f[mask] = val
                single.append((mask, val))

            for y in g[x]:
                fy, single_y = dfs(y)

                # 启发式合并
                if len(single_y) > len(single):
                    single, single_y = single_y, single
                    f, fy = fy, f

                single += single_y

                # 把子树 y 中的 mask 和 val 一个一个地加到 f 中
                for msk, v in single_y:
                    if v <= f[msk]:
                        continue
                    nf = f.copy()
                    nf[msk] = v
                    for msk2, s2 in f.items():
                        if msk & msk2 == 0:
                            nf[msk | msk2] = max(nf[msk | msk2], v + s2)
                    f = nf

            if f:
                nonlocal ans
                ans += max(f.values())
            return f, single

        dfs(0)
        return ans % 1_000_000_007
```

```java [sol-Java]
class Solution {
    private long ans = 0;

    public int goodSubtreeSum(int[] vals, int[] par) {
        int n = par.length;
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int i = 1; i < n; i++) {
            g[par[i]].add(i);
        }

        dfs(0, vals, g);
        return (int) (ans % 1_000_000_007);
    }

    private record Pair(Map<Integer, Integer> f, List<int[]> single) {
    }

    private Pair dfs(int x, int[] vals, List<Integer>[] g) {
        Map<Integer, Integer> f = new HashMap<>();
        List<int[]> single = new ArrayList<>();

        // 计算 vals[x] 的数位集合 mask
        int val = vals[x];
        int mask = 0;
        for (int v = val; v > 0; v /= 10) {
            int d = v % 10;
            if ((mask >> d & 1) > 0) {
                mask = 0;
                break;
            }
            mask |= 1 << d;
        }

        if (mask > 0) {
            f.put(mask, val);
            single.add(new int[]{mask, val});
        }

        for (int y : g[x]) {
            Pair resY = dfs(y, vals, g);
            Map<Integer, Integer> fy = resY.f;
            List<int[]> singleY = resY.single;

            // 启发式合并
            if (singleY.size() > single.size()) {
                List<int[]> tmpList = single;
                single = singleY;
                singleY = tmpList;

                Map<Integer, Integer> tmpMap = f;
                f = fy;
                fy = tmpMap;
            }

            single.addAll(singleY);

			// 把子树 y 中的 mask 和 val 一个一个地加到 f 中
            for (int[] p : singleY) {
                int msk = p[0];
                int v = p[1];
                if (v <= f.getOrDefault(msk, 0)) {
                    continue;
                }
                Map<Integer, Integer> nf = new HashMap<>(f);
                nf.put(msk, v);
                for (Map.Entry<Integer, Integer> e2 : f.entrySet()) {
                    int msk2 = e2.getKey();
                    if ((msk & msk2) == 0) {
                        nf.merge(msk | msk2, v + e2.getValue(), Math::max);
                    }
                }
                f = nf;
            }
        }

        if (!f.isEmpty()) {
            ans += Collections.max(f.values());
        }

        return new Pair(f, single);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int goodSubtreeSum(vector<int>& vals, vector<int>& par) {
        const int MOD = 1'000'000'007;
        const int D = 10;
        int n = par.size();
        vector<vector<int>> g(n);
        for (int i = 1; i < n; i++) {
            g[par[i]].push_back(i);
        }

        long long ans = 0;

        auto dfs = [&](this auto&& dfs, int x) -> pair<unordered_map<int, int>, vector<pair<int, int>>> {
            unordered_map<int, int> f;
            vector<pair<int, int>> single;

            // 计算 vals[x] 的数位集合 mask
            int val = vals[x];
            int mask = 0;
            for (int v = val; v > 0; v /= D) {
                int d = v % D;
                if (mask >> d & 1) {
                    mask = 0;
                    break;
                }
                mask |= 1 << d;
            }

            if (mask > 0) {
                f[mask] = val;
                single.emplace_back(mask, val);
            }

            for (int y : g[x]) {
                auto [fy, single_y] = dfs(y);
                
                // 启发式合并
                if (single_y.size() > single.size()) {
                    swap(single, single_y);
                    swap(f, fy);
                }

                single.insert(single.end(), single_y.begin(), single_y.end());

                // 把子树 y 中的 mask 和 val 一个一个地加到 f 中
                for (auto& [msk, v] : single_y) {
                    if (v <= f[msk]) {
                        continue;
                    }
                    auto nf = f;
                    nf[msk] = v;
                    for (auto& [msk2, s2] : f) {
                        if ((msk & msk2) == 0) {
                            nf[msk | msk2] = max(nf[msk | msk2], v + s2);
                        }
                    }
                    f = move(nf);
                }
            }

            int mx = 0;
            for (auto& [_, s] : f) {
                mx = max(mx, s);
            }
            ans += mx;

            return {f, single};
        };

        dfs(0);
        return ans % MOD;
    }
};
```

```go [sol-Go]
func goodSubtreeSum(vals, par []int) (ans int) {
	const mod = 1_000_000_007
	const D = 10
	n := len(par)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		p := par[i]
		g[p] = append(g[p], i)
	}

	type pair struct{ mask, val int }
	var dfs func(int) (map[int]int, []pair)
	dfs = func(x int) (f map[int]int, single []pair) {
		f = map[int]int{}

		// 计算 val 的数位集合 mask
		val := vals[x]
		mask := 0
		for v := val; v > 0; v /= D {
			d := v % D
			if mask>>d&1 > 0 {
				mask = 0
				break
			}
			mask |= 1 << d
		}

		if mask > 0 {
			f[mask] = val
			single = append(single, pair{mask, val})
		}

		for _, y := range g[x] {
			fy, singleY := dfs(y)

			// 启发式合并
			if len(singleY) > len(single) {
				single, singleY = singleY, single
				f, fy = fy, f
			}

			single = append(single, singleY...)

			// 把子树 y 中的 mask 和 val 一个一个地加到 f 中
			for _, p := range singleY {
				msk, v := p.mask, p.val
				if v <= f[msk] {
					continue
				}
				nf := maps.Clone(f)
				nf[msk] = v
				for msk2, s2 := range f {
					if msk&msk2 == 0 {
						nf[msk|msk2] = max(nf[msk|msk2], v+s2)
					}
				}
				f = nf
			}
		}

		mx := 0
		for _, s := range f {
			mx = max(mx, s)
		}
		ans += mx

		return
	}
	dfs(0)
	return ans % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n\cdot 2^D)$，其中 $n$ 是 $\textit{vals}$ 的长度，$D=10$。一个节点会多次「合并到更大子树中」，每次合并时，这个「更大子树」的大小至少比之前大一倍，所以每个节点的合并次数是 $\mathcal{O}(\log n)$ 的。
- 空间复杂度：$\mathcal{O}(n\cdot 2^D)$。

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
