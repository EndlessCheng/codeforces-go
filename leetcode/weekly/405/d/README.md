## 方法一：字符串哈希 + DP

[视频讲解](https://www.bilibili.com/video/BV1Ry411q71f/) 讲了为什么字典树做法是 $\mathcal{O}(n^2)$ 的，以及字符串哈希的解题思路。

定义 $f[i]$ 表示使 $s$ 等于 $\textit{target}$ 的长为 $i$ 的前缀的最小成本。

用「枚举选哪个」思考，枚举 $\textit{target}$ 的从下标 $j$ 到下标 $i-1$ 的这段子串 $t$，如果 $t$ 等于 $\textit{words}[k]$，则有

$$
f[i] = f[j] + \textit{costs}[k]
$$

取所有转移来源的最小值，即为 $f[i]$。

但即使用了字符串哈希，上述做法仍然是 $\mathcal{O}(n^2)$ 的。关键在于，枚举的子串 $t$ 的**长度**，如果压根就不出现在 $\textit{words}$ 中，那么无需枚举这样的 $j$，或者说长度。

注意到，设 $L$ 是 $\textit{words}$ 中所有字符串的长度之和，那么 $\textit{words}$ 中至多有 $\mathcal{O}(\sqrt L)$ 个长度不同的字符串。（考虑长度和 $1+2+3+\cdots \le L$） 

所以我们只需要枚举这 $\mathcal{O}(\sqrt L)$ 个长度，而不是枚举 $\mathcal{O}(n)$ 个 $j$。

初始值：$f[0]=0$。

答案：$f[n]$。如果 $f[n]=\infty$ 则返回 $-1$。

细节：$\textit{words}$ 中可能有相同字符串，这些字符串对应的成本应当取最小的。

### 写法一（单模哈希）

根据**生日攻击**，设 $M$ 为模数，仅仅计算大约 $1.18\cdot \sqrt {M}$ 个不同字符串的哈希值，就有 $50\%$ 的概率会发生哈希碰撞，也就是有两个不同的字符串，哈希值是一样的。

所以单模哈希通常来说无法保证其正确性，本题只要字符串的个数再大些，单模哈希就完全不行了。

更安全也更通用的做法是后面的写法二（双模哈希）。

```py [sol-Python3]
class Solution:
    def minimumCost(self, target: str, words: List[str], costs: List[int]) -> int:
        n = len(target)

        # 多项式字符串哈希（方便计算子串哈希值）
        # 哈希函数 hash(s) = s[0] * BASE^(n-1) + s[1] * BASE^(n-2) + ... + s[n-2] * BASE + s[n-1]
        MOD = 1_070_777_777
        BASE = randint(8 * 10 ** 8, 9 * 10 ** 8)  # 随机 BASE，防止 hack
        pow_base = [1] + [0] * n  # pow_base[i] = BASE^i
        pre_hash = [0] * (n + 1)  # 前缀哈希值 pre_hash[i] = hash(s[:i])
        for i, b in enumerate(target):
            pow_base[i + 1] = pow_base[i] * BASE % MOD
            pre_hash[i + 1] = (pre_hash[i] * BASE + ord(b)) % MOD  # 秦九韶算法计算多项式哈希

        # 每个 words[i] 的哈希值 -> 最小成本
        min_cost = defaultdict(lambda: inf)
        for w, c in zip(words, costs):
            h = 0
            for b in w:
                h = (h * BASE + ord(b)) % MOD
            min_cost[h] = min(min_cost[h], c)

        # 有 O(√L) 个不同的长度
        sorted_lens = sorted(set(map(len, words)))

        f = [0] + [inf] * n
        for i in range(1, n + 1):
            for sz in sorted_lens:
                if sz > i:
                    break
                # 计算子串 target[i-sz:i] 的哈希值（计算方法类似前缀和）
                sub_hash = (pre_hash[i] - pre_hash[i - sz] * pow_base[sz]) % MOD
                # 手写 min，避免超时
                tmp = f[i - sz] + min_cost[sub_hash]
                if tmp < f[i]:
                    f[i] = tmp
        return -1 if f[n] == inf else f[n]
```

```java [sol-Java]
class Solution {
    public int minimumCost(String target, String[] words, int[] costs) {
        char[] t = target.toCharArray();
        int n = t.length;

        // 多项式字符串哈希（方便计算子串哈希值）
        // 哈希函数 hash(s) = s[0] * base^(n-1) + s[1] * base^(n-2) + ... + s[n-2] * base + s[n-1]
        final int MOD = 1_070_777_777;
        final int BASE = (int) 8e8 + new Random().nextInt((int) 1e8); // 随机 base，防止 hack
        int[] powBase = new int[n + 1]; // powBase[i] = base^i
        int[] preHash = new int[n + 1]; // 前缀哈希值 preHash[i] = hash(target[0] 到 target[i-1])
        powBase[0] = 1;
        for (int i = 0; i < n; i++) {
            powBase[i + 1] = (int) ((long) powBase[i] * BASE % MOD);
            preHash[i + 1] = (int) (((long) preHash[i] * BASE + t[i]) % MOD); // 秦九韶算法计算多项式哈希
        }

        Map<Integer, Map<Integer, Integer>> minCost = new TreeMap<>(); // 长度 -> 哈希值 -> 最小成本
        for (int i = 0; i < words.length; i++) {
            long h = 0;
            for (char b : words[i].toCharArray()) {
                h = (h * BASE + b) % MOD;
            }
            minCost.computeIfAbsent(words[i].length(), k -> new HashMap<>())
                   .merge((int) h, costs[i], Integer::min);
        }

        int[] f = new int[n + 1];
        Arrays.fill(f, Integer.MAX_VALUE / 2);
        f[0] = 0;
        for (int i = 1; i <= n; i++) {
            for (Map.Entry<Integer, Map<Integer, Integer>> e : minCost.entrySet()) {
                int len = e.getKey();
                if (len > i) {
                    break;
                }
                // 计算子串 target[i-sz] 到 target[i-1] 的哈希值（计算方法类似前缀和）
                int subHash = (int) (((preHash[i] - (long) preHash[i - len] * powBase[len]) % MOD + MOD) % MOD);
                f[i] = Math.min(f[i], f[i - len] + e.getValue().getOrDefault(subHash, Integer.MAX_VALUE / 2));
            }
        }
        return f[n] == Integer.MAX_VALUE / 2 ? -1 : f[n];
    }
}
```

```java [sol-Java 更快写法]
class Solution {
    public int minimumCost(String target, String[] words, int[] costs) {
        char[] t = target.toCharArray();
        int n = t.length;

        // 多项式字符串哈希（方便计算子串哈希值）
        // 哈希函数 hash(s) = s[0] * base^(n-1) + s[1] * base^(n-2) + ... + s[n-2] * base + s[n-1]
        final int MOD = 1_070_777_777;
        final int BASE = (int) 8e8 + new Random().nextInt((int) 1e8); // 随机 base，防止 hack
        int[] powBase = new int[n + 1]; // powBase[i] = base^i
        int[] preHash = new int[n + 1]; // 前缀哈希值 preHash[i] = hash(target[0] 到 target[i-1])
        powBase[0] = 1;
        for (int i = 0; i < n; i++) {
            powBase[i + 1] = (int) ((long) powBase[i] * BASE % MOD);
            preHash[i + 1] = (int) (((long) preHash[i] * BASE + t[i]) % MOD); // 秦九韶算法计算多项式哈希
        }

        Map<Integer, Integer> minCost = new HashMap<>(); // 哈希值 -> 最小成本
        for (int i = 0; i < words.length; i++) {
            long h = 0;
            for (char b : words[i].toCharArray()) {
                h = (h * BASE + b) % MOD;
            }
            minCost.merge((int) h, costs[i], Integer::min);
        }

        // 有 O(√L) 个不同的长度
        Set<Integer> lengthSet = new HashSet<>();
        for (String w : words) {
            lengthSet.add(w.length());
        }
        int[] sortedLens = new int[lengthSet.size()];
        int k = 0;
        for (int len : lengthSet) {
            sortedLens[k++] = len;
        }
        Arrays.sort(sortedLens);

        int[] f = new int[n + 1];
        Arrays.fill(f, Integer.MAX_VALUE / 2);
        f[0] = 0;
        for (int i = 1; i <= n; i++) {
            for (int len : sortedLens) {
                if (len > i) {
                    break;
                }
                // 计算子串 target[i-sz] 到 target[i-1] 的哈希值（计算方法类似前缀和）
                int subHash = (int) (((preHash[i] - (long) preHash[i - len] * powBase[len]) % MOD + MOD) % MOD);
                f[i] = Math.min(f[i], f[i - len] + minCost.getOrDefault(subHash, Integer.MAX_VALUE / 2));
            }
        }
        return f[n] == Integer.MAX_VALUE / 2 ? -1 : f[n];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumCost(string target, vector<string>& words, vector<int>& costs) {
        int n = target.length();

        // 多项式字符串哈希（方便计算子串哈希值）
        // 哈希函数 hash(s) = s[0] * base^(n-1) + s[1] * base^(n-2) + ... + s[n-2] * base + s[n-1]
        const int MOD = 1'070'777'777;
        mt19937 rng(chrono::steady_clock::now().time_since_epoch().count());
        const int BASE = uniform_int_distribution<>(8e8, 9e8)(rng); // 随机 base，防止 hack
        vector<int> pow_base(n + 1); // pow_base[i] = base^i
        vector<int> pre_hash(n + 1); // 前缀哈希值 pre_hash[i] = hash(s[:i])
        pow_base[0] = 1;
        for (int i = 0; i < n; i++) {
            pow_base[i + 1] = (long long) pow_base[i] * BASE % MOD;
            pre_hash[i + 1] = ((long long) pre_hash[i] * BASE + target[i]) % MOD; // 秦九韶算法计算多项式哈希
        }

        // 计算 target[l] 到 target[r-1] 的哈希值
        auto sub_hash = [&](int l, int r) {
            return ((pre_hash[r] - (long long) pre_hash[l] * pow_base[r - l]) % MOD + MOD) % MOD;
        };

        map<int, unordered_map<int, int>> min_cost; // 长度 -> 哈希值 -> 最小成本
        for (int i = 0; i < words.size(); i++) {
            long long h = 0;
            for (char b : words[i]) {
                h = (h * BASE + b) % MOD;
            }
            int m = words[i].length();
            if (!min_cost[m].contains(h)) {
                min_cost[m][h] = costs[i];
            } else {
                min_cost[m][h] = min(min_cost[m][h], costs[i]);
            }
        }

        vector<int> f(n + 1, INT_MAX / 2);
        f[0] = 0;
        for (int i = 1; i <= n; i++) {
            for (auto& [len, mc] : min_cost) {
                if (len > i) {
                    break;
                }
                auto it = mc.find(sub_hash(i - len, i));
                if (it != mc.end()) {
                    f[i] = min(f[i], f[i - len] + it->second);
                }
            }
        }
        return f[n] == INT_MAX / 2 ? -1 : f[n];
    }
};
```

```go [sol-Go]
func minimumCost(target string, words []string, costs []int) int {
	n := len(target)

	// 多项式字符串哈希（方便计算子串哈希值）
	// 哈希函数 hash(s) = s[0] * base^(n-1) + s[1] * base^(n-2) + ... + s[n-2] * base + s[n-1]
	const mod = 1_070_777_777
	base := 9e8 - rand.Intn(1e8) // 随机 base，防止 hack（注意 Go1.20 之后的版本，每次随机的数都不一样）
	powBase := make([]int, n+1) // powBase[i] = base^i
	preHash := make([]int, n+1) // 前缀哈希值 preHash[i] = hash(s[:i])
	powBase[0] = 1
	for i, b := range target {
		powBase[i+1] = powBase[i] * base % mod
		preHash[i+1] = (preHash[i]*base + int(b)) % mod // 秦九韶算法计算多项式哈希
	}
	// 计算子串 target[l:r] 的哈希值，注意这是左闭右开区间 [l,r)
	// 计算方法类似前缀和
	subHash := func(l, r int) int {
		return ((preHash[r]-preHash[l]*powBase[r-l])%mod + mod) % mod
	}

	minCost := map[int]int{} // words[i] 的哈希值 -> 最小成本
	lens := map[int]struct{}{} // 所有 words[i] 的长度集合
	for i, w := range words {
		lens[len(w)] = struct{}{}
		h := 0
		for _, b := range w {
			h = (h*base + int(b)) % mod
		}
		if minCost[h] == 0 {
			minCost[h] = costs[i]
		} else {
			minCost[h] = min(minCost[h], costs[i])
		}
	}

	// 有 O(√L) 个不同的长度
	sortedLens := make([]int, 0, len(lens))
	for l := range lens {
		sortedLens = append(sortedLens, l)
	}
	slices.Sort(sortedLens)

	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt / 2
		for _, sz := range sortedLens {
			if sz > i {
				break
			}
			if cost, ok := minCost[subHash(i-sz, i)]; ok {
				f[i] = min(f[i], f[i-sz]+cost)
			}
		}
	}
	if f[n] == math.MaxInt/2 {
		return -1
	}
	return f[n]
}
```

```go [sol-Go 更快写法]
func minimumCost(target string, words []string, costs []int) int {
	n := len(target)

	// 多项式字符串哈希（方便计算子串哈希值）
	// 哈希函数 hash(s) = s[0] * base^(n-1) + s[1] * base^(n-2) + ... + s[n-2] * base + s[n-1]
	const mod = 1_070_777_777
	base := 9e8 - rand.Intn(1e8) // 随机 base，防止 hack（注意 Go1.20 之后的版本，每次随机的数都不一样）
	powBase := make([]int, n+1) // powBase[i] = base^i
	preHash := make([]int, n+1) // 前缀哈希值 preHash[i] = hash(s[:i])
	powBase[0] = 1
	for i, b := range target {
		powBase[i+1] = powBase[i] * base % mod
		preHash[i+1] = (preHash[i]*base + int(b)) % mod // 秦九韶算法计算多项式哈希
	}
	// 计算子串 target[l:r] 的哈希值，注意这是左闭右开区间 [l,r)
	// 计算方法类似前缀和
	subHash := func(l, r int) int {
		return ((preHash[r]-preHash[l]*powBase[r-l])%mod + mod) % mod
	}

	minCost := make([]map[int]int, n+1) // [words[i] 的长度][words[i] 的哈希值] -> 最小成本
	lens := map[int]struct{}{} // 所有 words[i] 的长度集合
	for i, w := range words {
		m := len(w)
		lens[m] = struct{}{}
		// 计算 w 的哈希值
		h := 0
		for _, b := range w {
			h = (h*base + int(b)) % mod
		}
		if minCost[m] == nil {
			minCost[m] = map[int]int{}
		}
		if minCost[m][h] == 0 {
			minCost[m][h] = costs[i]
		} else {
			minCost[m][h] = min(minCost[m][h], costs[i])
		}
	}

	// 有 O(√L) 个不同的长度
	sortedLens := make([]int, 0, len(lens))
	for l := range lens {
		sortedLens = append(sortedLens, l)
	}
	slices.Sort(sortedLens)

	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt / 2
		for _, sz := range sortedLens {
			if sz > i {
				break
			}
			if cost, ok := minCost[sz][subHash(i-sz, i)]; ok {
				f[i] = min(f[i], f[i-sz]+cost)
			}
		}
	}
	if f[n] == math.MaxInt/2 {
		return -1
	}
	return f[n]
}
```

### 写法二（双模哈希）

用两对 $\textit{mod}$ 和 $\textit{base}$ 计算哈希值，也就是每个字符串都有两个哈希值。

设 $M_1$ 和 $M_2$ 为模数，此时要计算大约 $1.18\cdot \sqrt {M_1M_2}$ 个不同字符串的哈希，才有 $50\%$ 的概率会发生哈希碰撞。

注：对于 Python 来说，把 $\textit{mod}$ 和 $\textit{base}$ 改大就行。这里提供一个大质数可以作为模数：$10^{18} + 3$。

```py [sol-Python3]
# 由于乘法超过 64 位整数范围，需要用到 bigint，所以效率不如写法一
class Solution:
    def minimumCost(self, target: str, words: List[str], costs: List[int]) -> int:
        n = len(target)

        # 多项式字符串哈希（方便计算子串哈希值）
        # 哈希函数 hash(s) = s[0] * BASE^(n-1) + s[1] * BASE^(n-2) + ... + s[n-2] * BASE + s[n-1]
        MOD = 10 ** 18 + 3
        BASE = randint(8 * 10 ** 17, 9 * 10 ** 17)  # 随机 BASE，防止 hack
        pow_base = [1] + [0] * n  # pow_base[i] = BASE^i
        pre_hash = [0] * (n + 1)  # 前缀哈希值 pre_hash[i] = hash(s[:i])
        for i, b in enumerate(target):
            pow_base[i + 1] = pow_base[i] * BASE % MOD
            pre_hash[i + 1] = (pre_hash[i] * BASE + ord(b)) % MOD  # 秦九韶算法计算多项式哈希

        # 每个 words[i] 的哈希值 -> 最小成本
        min_cost = defaultdict(lambda: inf)
        for w, c in zip(words, costs):
            h = 0
            for b in w:
                h = (h * BASE + ord(b)) % MOD
            min_cost[h] = min(min_cost[h], c)

        # 有 O(√L) 个不同的长度
        sorted_lens = sorted(set(map(len, words)))

        f = [0] + [inf] * n
        for i in range(1, n + 1):
            for sz in sorted_lens:
                if sz > i:
                    break
                # 计算子串 target[i-sz:i] 的哈希值（计算方法类似前缀和）
                sub_hash = (pre_hash[i] - pre_hash[i - sz] * pow_base[sz]) % MOD
                # 手写 min，避免超时
                tmp = f[i - sz] + min_cost[sub_hash]
                if tmp < f[i]:
                    f[i] = tmp
        return -1 if f[n] == inf else f[n]
```

```java [sol-Java]
// 基于上面的「更快写法」
class Solution {
    public int minimumCost(String target, String[] words, int[] costs) {
        char[] t = target.toCharArray();
        int n = t.length;
        final int MOD1 = 1_070_777_777;
        final int MOD2 = 1_000_000_007;
        final int BASE1 = (int) 8e8 + new Random().nextInt((int) 1e8);
        final int BASE2 = (int) 8e8 + new Random().nextInt((int) 1e8);
        int[] powBase1 = new int[n + 1];
        int[] powBase2 = new int[n + 1];
        int[] preHash1 = new int[n + 1];
        int[] preHash2 = new int[n + 1];
        powBase1[0] = powBase2[0] = 1;
        for (int i = 0; i < n; i++) {
            powBase1[i + 1] = (int) ((long) powBase1[i] * BASE1 % MOD1);
            powBase2[i + 1] = (int) ((long) powBase2[i] * BASE2 % MOD2);
            preHash1[i + 1] = (int) (((long) preHash1[i] * BASE1 + t[i]) % MOD1);
            preHash2[i + 1] = (int) (((long) preHash2[i] * BASE2 + t[i]) % MOD2);
        }

        Map<Long, Integer> minCost = new HashMap<>(); // 哈希值 -> 最小成本
        for (int i = 0; i < words.length; i++) {
            long h1 = 0;
            long h2 = 0;
            for (char b : words[i].toCharArray()) {
                h1 = (h1 * BASE1 + b) % MOD1;
                h2 = (h2 * BASE2 + b) % MOD2;
            }
            minCost.merge(h1 << 32 | h2, costs[i], Integer::min);
        }

        // 有 O(√L) 个不同的长度
        Set<Integer> lengthSet = new HashSet<>();
        for (String w : words) {
            lengthSet.add(w.length());
        }
        int[] sortedLens = new int[lengthSet.size()];
        int k = 0;
        for (int len : lengthSet) {
            sortedLens[k++] = len;
        }
        Arrays.sort(sortedLens);

        int[] f = new int[n + 1];
        Arrays.fill(f, Integer.MAX_VALUE / 2);
        f[0] = 0;
        for (int i = 1; i <= n; i++) {
            for (int len : sortedLens) {
                if (len > i) {
                    break;
                }
                // 计算子串 target[i-sz] 到 target[i-1] 的哈希值（计算方法类似前缀和）
                long subHash1 = ((preHash1[i] - (long) preHash1[i - len] * powBase1[len]) % MOD1 + MOD1) % MOD1;
                long subHash2 = ((preHash2[i] - (long) preHash2[i - len] * powBase2[len]) % MOD2 + MOD2) % MOD2;
                long subHash = subHash1 << 32 | subHash2;
                f[i] = Math.min(f[i], f[i - len] + minCost.getOrDefault(subHash, Integer.MAX_VALUE / 2));
            }
        }
        return f[n] == Integer.MAX_VALUE / 2 ? -1 : f[n];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumCost(string target, vector<string>& words, vector<int>& costs) {
        int n = target.length();
        const int MOD1 = 1'070'777'777;
        const int MOD2 = 1'000'000'007;
        mt19937 rng(chrono::steady_clock::now().time_since_epoch().count());
        const int BASE1 = uniform_int_distribution<>(8e8, 9e8)(rng);
        const int BASE2 = uniform_int_distribution<>(8e8, 9e8)(rng);
        vector<int> pow_base1(n + 1), pow_base2(n + 1);
        vector<int> pre_hash1(n + 1), pre_hash2(n + 1);
        pow_base1[0] = pow_base2[0] = 1;
        for (int i = 0; i < n; i++) {
            pow_base1[i + 1] = (long long) pow_base1[i] * BASE1 % MOD1;
            pow_base2[i + 1] = (long long) pow_base2[i] * BASE2 % MOD2;
            pre_hash1[i + 1] = ((long long) pre_hash1[i] * BASE1 + target[i]) % MOD1;
            pre_hash2[i + 1] = ((long long) pre_hash2[i] * BASE2 + target[i]) % MOD2;
        }

        // 计算 target[l] 到 target[r-1] 的哈希值
        auto sub_hash = [&](int l, int r) {
            long long h1 = ((pre_hash1[r] - (long long) pre_hash1[l] * pow_base1[r - l]) % MOD1 + MOD1) % MOD1;
            long long h2 = ((pre_hash2[r] - (long long) pre_hash2[l] * pow_base2[r - l]) % MOD2 + MOD2) % MOD2;
            return h1 << 32 | h2;
        };

        map<int, unordered_map<long long, int>> min_cost; // 长度 -> 哈希值 -> 最小成本
        for (int i = 0; i < words.size(); i++) {
            long long h1 = 0, h2 = 0;
            for (char b : words[i]) {
                h1 = (h1 * BASE1 + b) % MOD1;
                h2 = (h2 * BASE2 + b) % MOD2;
            }
            long long h = h1 << 32 | h2;
            int m = words[i].length();
            if (!min_cost[m].contains(h)) {
                min_cost[m][h] = costs[i];
            } else {
                min_cost[m][h] = min(min_cost[m][h], costs[i]);
            }
        }

        vector<int> f(n + 1, INT_MAX / 2);
        f[0] = 0;
        for (int i = 1; i <= n; i++) {
            for (auto& [len, mc] : min_cost) {
                if (len > i) {
                    break;
                }
                auto it = mc.find(sub_hash(i - len, i));
                if (it != mc.end()) {
                    f[i] = min(f[i], f[i - len] + it->second);
                }
            }
        }
        return f[n] == INT_MAX / 2 ? -1 : f[n];
    }
};
```

```go [sol-Go]
// 基于上面的「更快写法」
func minimumCost(target string, words []string, costs []int) int {
	n := len(target)
	const mod1 = 1_070_777_777
	const mod2 = 1_000_000_007
	base1 := 9e8 - rand.Intn(1e8)
	base2 := 9e8 - rand.Intn(1e8)

	type hPair struct{ h1, h2 int }
	powBase := make([]hPair, n+1)
	preHash := make([]hPair, n+1)
	powBase[0] = hPair{1, 1}
	for i, b := range target {
		powBase[i+1] = hPair{powBase[i].h1 * base1 % mod1, powBase[i].h2 * base2 % mod2}
		preHash[i+1] = hPair{(preHash[i].h1*base1 + int(b)) % mod1, (preHash[i].h2*base2 + int(b)) % mod2}
	}

	// 计算子串 target[l:r] 的哈希值
	// 空串的哈希值为 0
	subHash := func(l, r int) hPair {
		h1 := ((preHash[r].h1-preHash[l].h1*powBase[r-l].h1)%mod1 + mod1) % mod1
		h2 := ((preHash[r].h2-preHash[l].h2*powBase[r-l].h2)%mod2 + mod2) % mod2
		return hPair{h1, h2}
	}

	calcHash := func(t string) (p hPair) {
		for _, b := range t {
			p.h1 = (p.h1*base1 + int(b)) % mod1
			p.h2 = (p.h2*base2 + int(b)) % mod2
		}
		return
	}

	minCost := make([]map[hPair]int, n+1) // [words[i] 的长度][words[i] 的哈希值] -> 最小成本
	lens := map[int]struct{}{} // 所有 words[i] 的长度集合
	for i, w := range words {
		m := len(w)
		lens[m] = struct{}{}
		h := calcHash(w)
		if minCost[m] == nil {
			minCost[m] = map[hPair]int{}
		}
		if minCost[m][h] == 0 {
			minCost[m][h] = costs[i]
		} else {
			minCost[m][h] = min(minCost[m][h], costs[i])
		}
	}

	// 有 O(√L) 个不同的长度
	sortedLens := make([]int, 0, len(lens))
	for l := range lens {
		sortedLens = append(sortedLens, l)
	}
	slices.Sort(sortedLens)

	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt / 2
		for _, sz := range sortedLens {
			if sz > i {
				break
			}
			if cost, ok := minCost[sz][subHash(i-sz, i)]; ok {
				f[i] = min(f[i], f[i-sz]+cost)
			}
		}
	}
	if f[n] == math.MaxInt/2 {
		return -1
	}
	return f[n]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L + n\sqrt{L})$，其中 $n$ 是 $\textit{target}$ 的长度，$L$ 是 $\textit{words}$ 中所有字符串的长度之和。
- 空间复杂度：$\mathcal{O}(n + m)$。其中 $m$ 是 $\textit{words}$ 的长度。

## 方法二：后缀数组

利用**后缀数组**，可以快速计算出每个 $\textit{words}[i]$ 在 $\textit{target}$ 中的出现位置（匹配位置）。

一共有多少个匹配位置？考虑这样一个最坏情况：$\textit{target}$ 全是 $\texttt{a}$，$\textit{words}=[\texttt{a},\texttt{aa},\texttt{aaa},\cdots]$。设 $L$ 是 $\textit{words}$ 中所有字符串的长度之和，在这种情况下，$\textit{words}$ 中有 $\mathcal{O}(\sqrt L)$ 个字符串。每个字符串都会产生 $\mathcal{O}(n)$ 次匹配，所以一共有 $\mathcal{O}(n\sqrt L)$ 个匹配位置。在本题数据范围下，这是可以接受的。

如果 $\textit{words}[i]$ 与 $\textit{target}$ 的下标 $[l,r)$ 匹配，那么把二元组 $(l, \textit{costs}[i])$ 添加到 $\textit{from}[r]$ 中。

定义 $f[i]$ 表示使 $s$ 等于 $\textit{target}$ 的长为 $i$ 的前缀的最小成本。枚举 $\textit{from}[i]$ 中的二元组，假设我们匹配了 $\textit{target}$ 的下标 $[l,i)$ 这一段子串，那么我们需要解决的问题变成：使 $s$ 等于 $\textit{target}$ 的长为 $l$ 的前缀的最小成本。所以有

$$
f[i] = \min_j\{ f[\textit{from}[i][j].l] + \textit{from}[i][j].\textit{cost}  \}
$$

如果 $\textit{from}[i]$ 是空的，则 $f[i]=\infty$。

初始值：$f[0]=0$。

答案：$f[n]$。如果 $f[n]=\infty$ 则返回 $-1$。

细节：$\textit{words}$ 中可能有相同字符串，这些字符串对应的成本应当取最小的。

```go [sol-Go（自带后缀数组）]
func minimumCost(target string, words []string, costs []int) int {
	minCost := map[string]uint16{}
	for i, w := range words {
		c := uint16(costs[i])
		if minCost[w] == 0 {
			minCost[w] = c
		} else {
			minCost[w] = min(minCost[w], c)
		}
	}

	n := len(target)
	type pair struct{ l, cost uint16 }
	from := make([][]pair, n+1)
	sa := suffixarray.New([]byte(target))
	for w, c := range minCost {
		for _, l := range sa.Lookup([]byte(w), -1) {
			r := l + len(w)
			from[r] = append(from[r], pair{uint16(l), c})
		}
	}

	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt / 2
		for _, p := range from[i] {
			f[i] = min(f[i], f[p.l]+int(p.cost))
		}
	}
	if f[n] == math.MaxInt/2 {
		return -1
	}
	return f[n]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L + n\sqrt{L})$，其中 $n$ 是 $\textit{target}$ 的长度，$L$ 是 $\textit{words}$ 中所有字符串的长度之和。有多少个匹配，就有多少次状态转移。
- 空间复杂度：$\mathcal{O}(L + n\sqrt{L})$。

## 方法三：AC 自动机

> 考虑到字符串哈希不一定 100% 正确，后缀数组又需要太多空间，从某种程度上来说，AC 自动机是本题标准做法。

原理见 [OI Wiki](https://oi-wiki.org/string/ac-automaton/)。学习之前推荐先看看我的 [KMP 原理讲解](https://www.zhihu.com/question/21923021/answer/37475572)。

做法同后缀数组，改成在 AC 自动机上寻找匹配。

```py [sol-Python3]
class Node:
    __slots__ = 'son', 'fail', 'last', 'len', 'cost'

    def __init__(self):
        self.son = [None] * 26
        self.fail = None  # 当 cur.son[i] 不能匹配 target 中的某个字符时，cur.fail.son[i] 即为下一个待匹配节点（等于 root 则表示没有匹配）
        self.last = None  # 后缀链接（suffix link），用来快速跳到一定是某个 words[k] 的最后一个字母的节点（等于 root 则表示没有）
        self.len = 0
        self.cost = inf


class AhoCorasick:
    def __init__(self):
        self.root = Node()

    def put(self, s: str, cost: int) -> None:
        cur = self.root
        for b in s:
            b = ord(b) - ord('a')
            if cur.son[b] is None:
                cur.son[b] = Node()
            cur = cur.son[b]
        cur.len = len(s)
        cur.cost = min(cur.cost, cost)

    def build_fail(self) -> None:
        self.root.fail = self.root.last = self.root
        q = deque()
        for i, son in enumerate(self.root.son):
            if son is None:
                self.root.son[i] = self.root
            else:
                son.fail = son.last = self.root  # 第一层的失配指针，都指向根节点 ∅
                q.append(son)
        # BFS
        while q:
            cur = q.popleft()
            for i, son in enumerate(cur.son):
                if son is None:
                    # 虚拟子节点 cur.son[i]，和 cur.fail.son[i] 是同一个
                    # 方便失配时直接跳到下一个可能匹配的位置（但不一定是某个 words[k] 的最后一个字母）
                    cur.son[i] = cur.fail.son[i]
                    continue
                son.fail = cur.fail.son[i]  # 计算失配位置
                # 沿着 last 往上走，可以直接跳到一定是某个 words[k] 的最后一个字母的节点（如果跳到 root 表示没有匹配）
                son.last = son.fail if son.fail.len else son.fail.last
                q.append(son)


class Solution:
    def minimumCost(self, target: str, words: List[str], costs: List[int]) -> int:
        ac = AhoCorasick()
        for w, c in zip(words, costs):
            ac.put(w, c)
        ac.build_fail()

        n = len(target)
        f = [0] + [inf] * n
        cur = root = ac.root
        for i in range(1, n + 1):
            cur = cur.son[ord(target[i - 1]) - ord('a')]  # 如果没有匹配相当于移动到 fail 的 son[target[i-1]-'a']
            if cur.len:  # 匹配到了一个尽可能长的 words[k]
                f[i] = min(f[i], f[i - cur.len] + cur.cost)
            # 还可能匹配其余更短的 words[k]，要在 last 链上找
            match_node = cur.last
            while match_node != root:
                # 手写 min 更快
                tmp = f[i - match_node.len] + match_node.cost
                if tmp < f[i]:
                    f[i] = tmp
                match_node = match_node.last
        return -1 if f[n] == inf else f[n]
```

```java [sol-Java]
class Node {
    Node[] son = new Node[26];
    Node fail; // 当 cur.son[i] 不能匹配 target 中的某个字符时，cur.fail.son[i] 即为下一个待匹配节点（等于 root 则表示没有匹配）
    Node last; // 后缀链接（suffix link），用来快速跳到一定是某个 words[k] 的最后一个字母的节点（等于 root 则表示没有）
    int len;
    int cost = Integer.MAX_VALUE;
}

class AhoCorasick {
    Node root = new Node();

    void put(String s, int cost) {
        Node cur = root;
        for (char b : s.toCharArray()) {
            b -= 'a';
            if (cur.son[b] == null) {
                cur.son[b] = new Node();
            }
            cur = cur.son[b];
        }
        cur.len = s.length();
        cur.cost = Math.min(cur.cost, cost);
    }

    void buildFail() {
        root.fail = root.last = root;
        Queue<Node> q = new ArrayDeque<>();
        for (int i = 0; i < root.son.length; i++) {
            Node son = root.son[i];
            if (son == null) {
                root.son[i] = root;
            } else {
                son.fail = son.last = root; // 第一层的失配指针，都指向根节点 ∅
                q.add(son);
            }
        }
        // BFS
        while (!q.isEmpty()) {
            Node cur = q.poll();
            for (int i = 0; i < 26; i++) {
                Node son = cur.son[i];
                if (son == null) {
                    // 虚拟子节点 cur.son[i]，和 cur.fail.son[i] 是同一个
                    // 方便失配时直接跳到下一个可能匹配的位置（但不一定是某个 words[k] 的最后一个字母）
                    cur.son[i] = cur.fail.son[i];
                    continue;
                }
                son.fail = cur.fail.son[i]; // 计算失配位置
                // 沿着 last 往上走，可以直接跳到一定是某个 words[k] 的最后一个字母的节点（如果跳到 root 表示没有匹配）
                son.last = son.fail.len > 0 ? son.fail : son.fail.last;
                q.add(son);
            }
        }
    }
}

class Solution {
    public int minimumCost(String target, String[] words, int[] costs) {
        AhoCorasick ac = new AhoCorasick();
        for (int i = 0; i < words.length; i++) {
            ac.put(words[i], costs[i]);
        }
        ac.buildFail();

        char[] t = target.toCharArray();
        int n = t.length;
        int[] f = new int[n + 1];
        Arrays.fill(f, Integer.MAX_VALUE / 2);
        f[0] = 0;
        Node cur = ac.root;
        for (int i = 1; i <= n; i++) {
            cur = cur.son[t[i - 1] - 'a']; // 如果没有匹配相当于移动到 fail 的 son[t[i-1]-'a']
            if (cur.len > 0) { // 匹配到了一个尽可能长的 words[k]
                f[i] = Math.min(f[i], f[i - cur.len] + cur.cost);
            }
            // 还可能匹配其余更短的 words[k]，要在 last 链上找
            for (Node match = cur.last; match != ac.root; match = match.last) {
                f[i] = Math.min(f[i], f[i - match.len] + match.cost);
            }
        }
        return f[n] == Integer.MAX_VALUE / 2 ? -1 : f[n];
    }
}
```

```cpp [sol-C++]
struct Node {
    Node* son[26]{};
    Node* fail; // 当 cur.son[i] 不能匹配 target 中的某个字符时，cur.fail.son[i] 即为下一个待匹配节点（等于 root 则表示没有匹配）
    Node* last; // 后缀链接（suffix link），用来快速跳到一定是某个 words[k] 的最后一个字母的节点（等于 root 则表示没有）
    int len = 0;
    int cost = INT_MAX;
};

struct AhoCorasick {
    Node* root = new Node();

    void put(string& s, int cost) {
        auto cur = root;
        for (char b : s) {
            b -= 'a';
            if (cur->son[b] == nullptr) {
                cur->son[b] = new Node();
            }
            cur = cur->son[b];
        }
        cur->len = s.length();
        cur->cost = min(cur->cost, cost);
    }

    void build_fail() {
        root->fail = root->last = root;
        queue<Node*> q;
        for (auto& son : root->son) {
            if (son == nullptr) {
                son = root;
            } else {
                son->fail = son->last = root; // 第一层的失配指针，都指向根节点 ∅
                q.push(son);
            }
        }
        // BFS
        while (!q.empty()) {
            auto cur = q.front();
            q.pop();
            for (int i = 0; i < 26; i++) {
                auto& son = cur->son[i];
                if (son == nullptr) {
                    // 虚拟子节点 cur.son[i]，和 cur.fail.son[i] 是同一个
                    // 方便失配时直接跳到下一个可能匹配的位置（但不一定是某个 words[k] 的最后一个字母）
                    son = cur->fail->son[i];
                    continue;
                }
                son->fail = cur->fail->son[i]; // 计算失配位置
                // 沿着 last 往上走，可以直接跳到一定是某个 words[k] 的最后一个字母的节点（如果跳到 root 表示没有匹配）
                son->last = son->fail->len ? son->fail : son->fail->last;
                q.push(son);
            }
        }
    }
};

class Solution {
public:
    int minimumCost(string target, vector<string>& words, vector<int>& costs) {
        AhoCorasick ac;
        for (int i = 0; i < words.size(); i++) {
            ac.put(words[i], costs[i]);
        }
        ac.build_fail();

        int n = target.length();
        vector<int> f(n + 1, INT_MAX / 2);
        f[0] = 0;
        auto cur = ac.root;
        for (int i = 1; i <= n; i++) {
            cur = cur->son[target[i - 1] - 'a']; // 如果没有匹配相当于移动到 fail 的 son[target[i-1]-'a']
            if (cur->len) { // 匹配到了一个尽可能长的 words[k]
                f[i] = min(f[i], f[i - cur->len] + cur->cost);
            }
            // 还可能匹配其余更短的 words[k]，要在 last 链上找
            for (auto match = cur->last; match != ac.root; match = match->last) {
                f[i] = min(f[i], f[i - match->len] + match->cost);
            }
        }
        return f[n] == INT_MAX / 2 ? -1 : f[n];
    }
};
```

```go [sol-Go]
type node struct {
	son  [26]*node
	fail *node // 当 cur.son[i] 不能匹配 target 中的某个字符时，cur.fail.son[i] 即为下一个待匹配节点（等于 root 则表示没有匹配）
	last *node // 后缀链接（suffix link），用来快速跳到一定是某个 words[k] 的最后一个字母的节点（等于 root 则表示没有）
	len  int
	cost int
}

type acam struct {
	root *node
}

func (ac *acam) put(s string, cost int) {
	cur := ac.root
	for _, b := range s {
		b -= 'a'
		if cur.son[b] == nil {
			cur.son[b] = &node{cost: math.MaxInt}
		}
		cur = cur.son[b]
	}
	cur.len = len(s)
	cur.cost = min(cur.cost, cost)
}

func (ac *acam) buildFail() {
	ac.root.fail = ac.root
	ac.root.last = ac.root
	q := []*node{}
	for i, son := range ac.root.son[:] {
		if son == nil {
			ac.root.son[i] = ac.root
		} else {
			son.fail = ac.root // 第一层的失配指针，都指向根节点 ∅
			son.last = ac.root
			q = append(q, son)
		}
	}
	// BFS
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for i, son := range cur.son[:] {
			if son == nil {
				// 虚拟子节点 cur.son[i]，和 cur.fail.son[i] 是同一个
				// 方便失配时直接跳到下一个可能匹配的位置（但不一定是某个 words[k] 的最后一个字母）
				cur.son[i] = cur.fail.son[i]
				continue
			}
			son.fail = cur.fail.son[i] // 计算失配位置
			if son.fail.len > 0 {
				son.last = son.fail
			} else {
				// 沿着 last 往上走，可以直接跳到一定是某个 words[k] 的最后一个字母的节点（如果跳到 root 表示没有匹配）
				son.last = son.fail.last
			}
			q = append(q, son)
		}
	}
}

func minimumCost(target string, words []string, costs []int) int {
	ac := &acam{root: &node{}}
	for i, w := range words {
		ac.put(w, costs[i])
	}
	ac.buildFail()

	n := len(target)
	f := make([]int, n+1)
	cur := ac.root
	for i, b := range target {
		cur = cur.son[b-'a'] // 如果没有匹配相当于移动到 fail 的 son[b-'a']
		i++
		f[i] = math.MaxInt / 2
		if cur.len > 0 { // 匹配到了一个尽可能长的 words[k]
			f[i] = min(f[i], f[i-cur.len]+cur.cost)
		}
		// 还可能匹配其余更短的 words[k]，要在 last 链上找
		for match := cur.last; match != ac.root; match = match.last {
			f[i] = min(f[i], f[i-match.len]+match.cost)
		}
	}
	if f[n] == math.MaxInt/2 {
		return -1
	}
	return f[n]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L|\Sigma| + n\sqrt{L})$，其中 $n$ 是 $\textit{target}$ 的长度，$L$ 是 $\textit{words}$ 中所有字符串的长度之和，$|\Sigma|$ 是字符集合的大小，本题字符均为小写字母，所以 $|\Sigma|=26$。有多少个匹配，就有多少次状态转移。
- 空间复杂度：$\mathcal{O}(L|\Sigma| + n)$。

## 相似题目

- [3292. 形成目标字符串需要的最少字符串数 II](https://leetcode.cn/problems/minimum-number-of-valid-strings-to-form-target-ii/)

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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
