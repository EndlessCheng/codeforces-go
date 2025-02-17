## 题意

把 $\textit{target}$ 划分成若干段，要求每一段都是某个 $\textit{words}[i]$ 的前缀。

返回**最小划分成多少段**。如果无法划分，返回 $-1$。

## 分析

示例 1 的 $\textit{words}=[\texttt{abc},\texttt{aaaaa},\texttt{bcdef}]$，$\textit{target}=\texttt{aabcdabc}$。

要让划分的段数尽量小，那么每一段的长度要尽量长？

虽然还不知道要怎么划分，但这启发我们考虑如下内容：

- 从 $\textit{target}[0]$ 开始的段，最长可以是多长？答案是 $2$，因为 $\texttt{aa}$ 是 $\textit{words}[1]=\texttt{aaaaa}$ 的前缀。
- 从 $\textit{target}[1]$ 开始的段，最长可以是多长？答案是 $3$，因为 $\texttt{abc}$ 是 $\textit{words}[0]=\texttt{abc}$ 的前缀。
- 从 $\textit{target}[2]$ 开始的段，最长可以是多长？答案是 $3$，因为 $\texttt{bcd}$ 是 $\textit{words}[2]=\texttt{bcdef}$ 的前缀。
- ……

注意 $\textit{words}[i]$ 前缀的前缀还是 $\textit{words}[i]$ 的前缀。$\texttt{bcd}$ 是 $\texttt{bcdef}$ 的前缀，意味着 $\texttt{b}$ 和 $\texttt{bc}$ 也是 $\texttt{bcdef}$ 的前缀。如果从 $\textit{target}[2]$ 开始的段，最长可以是 $3$，那么这个段的长度也可以是 $1$ 或者 $2$。

如果我们算出了上面这些最长长度 $2,3,3,\cdots$，那么问题就变成：

- 给你一个数组 $\textit{maxJumps} = [2,3,3,\cdots]$。你从 $0$ 开始向右跳。在下标 $i$ 处，可以跳到 $[i+1,i+ \textit{maxJumps}[i]]$ 中的任意下标。目标是到达 $n$，最小要跳多少次？即 [45. 跳跃游戏 II](https://leetcode.cn/problems/jump-game-ii/) 或 [1326. 灌溉花园的最少水龙头数目](https://leetcode.cn/problems/minimum-number-of-taps-to-open-to-water-a-garden/)。

示例 1 的 $\textit{maxJumps} = [2, 3, 3, 0, 0, 3, 2, 0]$，跳法是 $0\to 2\to 5\to 8$。

现在剩下的问题是，如何计算 $\textit{maxJumps}$ 数组？

## 方法一：Z 函数

对于字符串 $s$，定义 $z[i]$ 表示后缀 $s[i:]$ 与 $s$ 的 LCP（最长公共前缀）长度，其中 $s[i:]$ 表示从 $s[i]$ 到 $s[n-1]$ 的子串。

遍历 $\textit{words}$，对于 $\textit{word} =\textit{words}[i]$，构造字符串

$$
s = \textit{word} + \texttt{#} + \textit{target}
$$

中间插入 $\texttt{#}$ 目的是避免 $z[i]$ 超过 $\textit{word}$ 的长度。

计算 $s$ 的 $z$ 数组。设 $m$ 为 $\textit{word}$ 的长度加一，那么 $\textit{target}[i:]$ 与 $\textit{word}$ 的最长公共前缀，就是 $z[m+i]$。用 $z[m+i]$ 更新 $\textit{maxJumps}[i]$ 的最大值。

```py [sol-Python3]
class Solution:
    def calc_z(self, s: str) -> list[int]:
        n = len(s)
        z = [0] * n
        box_l = box_r = 0  # z-box 左右边界（闭区间）
        for i in range(1, n):
            if i <= box_r:
                z[i] = min(z[i - box_l], box_r - i + 1)
            while i + z[i] < n and s[z[i]] == s[i + z[i]]:
                box_l, box_r = i, i + z[i]
                z[i] += 1
        return z

    # 桥的概念，见我在 45 或 1326 题下的题解
    def jump(self, max_jumps: List[int]) -> int:
        ans = 0
        cur_r = 0  # 已建造的桥的右端点
        nxt_r = 0  # 下一座桥的右端点的最大值
        for i, max_jump in enumerate(max_jumps):  # 如果走到 n-1 时没有返回 -1，那么必然可以到达 n
            nxt_r = max(nxt_r, i + max_jump)
            if i == cur_r:  # 到达已建造的桥的右端点
                if i == nxt_r:  # 无论怎么造桥，都无法从 i 到 i+1
                    return -1
                cur_r = nxt_r  # 造一座桥
                ans += 1
        return ans

    def minValidStrings(self, words: List[str], target: str) -> int:
        n = len(target)
        max_jumps = [0] * n
        for word in words:
            z = self.calc_z(word + "#" + target)
            m = len(word) + 1
            for i in range(n):
                max_jumps[i] = max(max_jumps[i], z[m + i])
        return self.jump(max_jumps)
```

```py [sol-Python3 手动比大小]
class Solution:
    def calc_z(self, s: str) -> list[int]:
        n = len(s)
        z = [0] * n
        box_l = box_r = 0  # z-box 左右边界（闭区间）
        for i in range(1, n):
            if i <= box_r:
                # 手动 min，加快速度
                x = z[i - box_l]
                y = box_r - i + 1
                z[i] = x if x < y else y
            while i + z[i] < n and s[z[i]] == s[i + z[i]]:
                box_l, box_r = i, i + z[i]
                z[i] += 1
        return z

    # 桥的概念，见我在 45 或 1326 题下的题解
    def jump(self, max_jumps: List[int]) -> int:
        ans = 0
        cur_r = 0  # 已建造的桥的右端点
        nxt_r = 0  # 下一座桥的右端点的最大值
        for i, max_jump in enumerate(max_jumps):  # 如果走到 n-1 时没有返回 -1，那么必然可以到达 n
            nxt_r = max(nxt_r, i + max_jump)
            if i == cur_r:  # 到达已建造的桥的右端点
                if i == nxt_r:  # 无论怎么造桥，都无法从 i 到 i+1
                    return -1
                cur_r = nxt_r  # 造一座桥
                ans += 1
        return ans

    def minValidStrings(self, words: List[str], target: str) -> int:
        n = len(target)
        max_jumps = [0] * n
        for word in words:
            z = self.calc_z(word + "#" + target)
            m = len(word) + 1
            for i in range(n):
                # 手动 max，加快速度
                if z[m + i] > max_jumps[i]:
                    max_jumps[i] = z[m + i]
        return self.jump(max_jumps)
```

```java [sol-Java]
class Solution {
    public int minValidStrings(String[] words, String target) {
        int n = target.length();
        int[] maxJumps = new int[n];
        for (String word : words) {
            int[] z = calcZ(word + "#" + target);
            int m = word.length() + 1;
            for (int i = 0; i < n; i++) {
                maxJumps[i] = Math.max(maxJumps[i], z[m + i]);
            }
        }
        return jump(maxJumps);
    }

    private int[] calcZ(String S) {
        char[] s = S.toCharArray();
        int n = s.length;
        int[] z = new int[n];
        int boxL = 0;
        int boxR = 0; // z-box 左右边界（闭区间）
        for (int i = 1; i < n; i++) {
            if (i <= boxR) {
                z[i] = Math.min(z[i - boxL], boxR - i + 1);
            }
            while (i + z[i] < n && s[z[i]] == s[i + z[i]]) {
                boxL = i;
                boxR = i + z[i];
                z[i]++;
            }
        }
        return z;
    }

    // 桥的概念，见我在 45 或 1326 题下的题解
    private int jump(int[] maxJumps) {
        int ans = 0;
        int curR = 0; // 已建造的桥的右端点
        int nxtR = 0; // 下一座桥的右端点的最大值
        for (int i = 0; i < maxJumps.length; i++) {
            nxtR = Math.max(nxtR, i + maxJumps[i]);
            if (i == curR) { // 到达已建造的桥的右端点
                if (i == nxtR) { // 无论怎么造桥，都无法从 i 到 i+1
                    return -1;
                }
                curR = nxtR; // 造一座桥
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<int> calc_z(string s) {
        int n = s.length();
        vector<int> z(n);
        int box_l = 0, box_r = 0; // z-box 左右边界（闭区间）
        for (int i = 1; i < n; i++) {
            if (i <= box_r) {
                z[i] = min(z[i - box_l], box_r - i + 1);
            }
            while (i + z[i] < n && s[z[i]] == s[i + z[i]]) {
                box_l = i;
                box_r = i + z[i];
                z[i]++;
            }
        }
        return z;
    }

    // 桥的概念，见我在 45 或 1326 题下的题解
    int jump(vector<int>& max_jumps) {
        int ans = 0;
        int cur_r = 0; // 已建造的桥的右端点
        int nxt_r = 0; // 下一座桥的右端点的最大值
        for (int i = 0; i < max_jumps.size(); i++) {
            nxt_r = max(nxt_r, i + max_jumps[i]);
            if (i == cur_r) { // 到达已建造的桥的右端点
                if (i == nxt_r) { // 无论怎么造桥，都无法从 i 到 i+1
                    return -1;
                }
                cur_r = nxt_r; // 造一座桥
                ans++;
            }
        }
        return ans;
    }

public:
    int minValidStrings(vector<string>& words, string target) {
        int n = target.length();
        vector<int> max_jumps(n);
        for (auto& word : words) {
            vector<int> z = calc_z(word + "#" + target);
            int m = word.length() + 1;
            for (int i = 0; i < n; i++) {
                max_jumps[i] = max(max_jumps[i], z[m + i]);
            }
        }
        return jump(max_jumps);
    }
};
```

```go [sol-Go]
func calcZ(s string) []int {
	n := len(s)
	z := make([]int, n)
	boxL, boxR := 0, 0 // z-box 左右边界（闭区间）
	for i := 1; i < n; i++ {
		if i <= boxR {
			z[i] = min(z[i-boxL], boxR-i+1)
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			boxL, boxR = i, i+z[i]
			z[i]++
		}
	}
	return z
}

// 桥的概念，见我在 45 或 1326 题下的题解
func jump(maxJumps []int) (ans int) {
	curR := 0 // 已建造的桥的右端点
	nxtR := 0 // 下一座桥的右端点的最大值
	for i, maxJump := range maxJumps {
		nxtR = max(nxtR, i+maxJump)
		if i == curR { // 到达已建造的桥的右端点
			if i == nxtR { // 无论怎么造桥，都无法从 i 到 i+1
				return -1
			}
			curR = nxtR // 造一座桥
			ans++
		}
	}
	return
}

func minValidStrings(words []string, target string) int {
	maxJumps := make([]int, len(target))
	for _, word := range words {
		z := calcZ(word + "#" + target)
		for i, z := range z[len(word)+1:] {
			maxJumps[i] = max(maxJumps[i], z)
		}
	}
	return jump(maxJumps)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L+nm)$，其中 $n$ 是 $\textit{target}$ 的长度，$m$ 是 $\textit{words}$ 的长度，$L$ 是 $\textit{words}$ 中所有字符串的长度之和。
- 空间复杂度：$\mathcal{O}(l+n)$。其中 $l$ 是 $\textit{words}[i]$ 的长度。

## 方法二：字符串哈希

### 写法一：二分

预处理每个 $\textit{words}[i]$ 的每个前缀的字符串哈希值，**按照前缀长度分组**，保存到不同的集合中。每个集合保存的是相同前缀长度的哈希值。

由于 $\textit{words}$ 的长度至多为 $100$，所以每个集合至多保存 $100$ 个哈希值，根据生日攻击理论，单模哈希绰绰有余，碰撞概率很小。

然后对于每个 $i$，二分求出 $\textit{maxJumps}[i]$。

二分的 $\text{check}(\textit{mid})$ 函数怎么写？判断从 $\textit{target}[i]$ 开始的长为 $\textit{mid}$ 的子串，哈希值是否在集合中。

具体请看 [本题视频讲解](https://www.bilibili.com/video/BV1Qp4me2Emz/) 第四题，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minValidStrings(self, words: List[str], target: str) -> int:
        n = len(target)

        # 多项式字符串哈希（方便计算子串哈希值）
        # 哈希函数 hash(s) = s[0] * BASE^(n-1) + s[1] * BASE^(n-2) + ... + s[n-2] * BASE + s[n-1]
        MOD = 1_070_777_777
        BASE = randint(8 * 10 ** 8, 9 * 10 ** 8)  # 随机 BASE，防止 hack
        pow_base = [1] + [0] * n  # pow_base[i] = BASE^i
        pre_hash = [0] * (n + 1)  # 前缀哈希值 pre_hash[i] = hash(target[:i])
        for i, b in enumerate(target):
            pow_base[i + 1] = pow_base[i] * BASE % MOD
            pre_hash[i + 1] = (pre_hash[i] * BASE + ord(b)) % MOD  # 秦九韶算法计算多项式哈希

        # 计算子串 target[l:r] 的哈希值，注意这是左闭右开区间 [l,r)
        # 计算方法类似前缀和
        def sub_hash(l: int, r: int) -> int:
            return (pre_hash[r] - pre_hash[l] * pow_base[r - l]) % MOD

        # 保存每个 words[i] 的每个前缀的哈希值，按照长度分组
        max_len = max(map(len, words))
        sets = [set() for _ in range(max_len)]
        for w in words:
            h = 0
            for j, b in enumerate(w):
                h = (h * BASE + ord(b)) % MOD
                sets[j].add(h)  # 注意 j 从 0 开始

        ans = 0
        cur_r = 0  # 已建造的桥的右端点
        nxt_r = 0  # 下一座桥的右端点的最大值
        for i in range(n):
            check = lambda j: sub_hash(i, i + j + 1) not in sets[j]
            max_jump = bisect_left(range(min(n - i, max_len)), True, key=check)
            nxt_r = max(nxt_r, i + max_jump)
            if i == cur_r:  # 到达已建造的桥的右端点
                if i == nxt_r:  # 无论怎么造桥，都无法从 i 到 i+1
                    return -1
                cur_r = nxt_r  # 建造下一座桥
                ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_070_777_777;

    public int minValidStrings(String[] words, String target) {
        char[] t = target.toCharArray();
        int n = t.length;

        // 多项式字符串哈希（方便计算子串哈希值）
        // 哈希函数 hash(s) = s[0] * base^(n-1) + s[1] * base^(n-2) + ... + s[n-2] * base + s[n-1]
        final int BASE = (int) 8e8 + new Random().nextInt((int) 1e8); // 随机 base，防止 hack
        int[] powBase = new int[n + 1]; // powBase[i] = base^i
        int[] preHash = new int[n + 1]; // 前缀哈希值 preHash[i] = hash(target[0] 到 target[i-1])
        powBase[0] = 1;
        for (int i = 0; i < n; i++) {
            powBase[i + 1] = (int) ((long) powBase[i] * BASE % MOD);
            preHash[i + 1] = (int) (((long) preHash[i] * BASE + t[i]) % MOD); // 秦九韶算法计算多项式哈希
        }

        int maxLen = 0;
        for (String w : words) {
            maxLen = Math.max(maxLen, w.length());
        }
        Set<Integer>[] sets = new HashSet[maxLen];
        Arrays.setAll(sets, i -> new HashSet<>());
        for (String w : words) {
            long h = 0;
            for (int j = 0; j < w.length(); j++) {
                h = (h * BASE + w.charAt(j)) % MOD;
                sets[j].add((int) h); // 注意 j 从 0 开始
            }
        }

        int ans = 0;
        int curR = 0; // 已建造的桥的右端点
        int nxtR = 0; // 下一座桥的右端点的最大值
        for (int i = 0; i < n; i++) {
            int maxJump = calcMaxJump(i, preHash, powBase, sets);
            nxtR = Math.max(nxtR, i + maxJump);
            if (i == curR) { // 到达已建造的桥的右端点
                if (i == nxtR) { // 无论怎么造桥，都无法从 i 到 i+1
                    return -1;
                }
                curR = nxtR; // 造一座桥
                ans++;
            }
        }
        return ans;
    }

    private int calcMaxJump(int i, int[] preHash, int[] powBase, Set<Integer>[] sets) {
        // 开区间二分，left 一定满足要求，right 一定不满足要求
        int left = 0;
        int right = Math.min(preHash.length - 1 - i, sets.length) + 1;
        while (left + 1 < right) {
            int mid = (left + right) >>> 1;
            long subHash = (((long) preHash[i + mid] - (long) preHash[i] * powBase[mid]) % MOD + MOD) % MOD;
            if (sets[mid - 1].contains((int) subHash)) {
                left = mid;
            } else {
                right = mid;
            }
        }
        return left;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minValidStrings(vector<string>& words, string target) {
        int n = target.length();

        // 多项式字符串哈希（方便计算子串哈希值）
        // 哈希函数 hash(s) = s[0] * base^(n-1) + s[1] * base^(n-2) + ... + s[n-2] * base + s[n-1]
        const int MOD = 1'070'777'777;
        mt19937 rng(chrono::steady_clock::now().time_since_epoch().count());
        const int BASE = uniform_int_distribution<>(8e8, 9e8)(rng); // 随机 base，防止 hack
        vector<int> pow_base(n + 1); // pow_base[i] = base^i
        vector<int> pre_hash(n + 1); // 前缀哈希值 pre_hash[i] = hash(target[:i])
        pow_base[0] = 1;
        for (int i = 0; i < n; i++) {
            pow_base[i + 1] = (long long) pow_base[i] * BASE % MOD;
            pre_hash[i + 1] = ((long long) pre_hash[i] * BASE + target[i]) % MOD; // 秦九韶算法计算多项式哈希
        }
        // 计算 target[l] 到 target[r-1] 的哈希值
        auto sub_hash = [&](int l, int r) {
            return ((pre_hash[r] - (long long) pre_hash[l] * pow_base[r - l]) % MOD + MOD) % MOD;
        };

        int max_len = 0;
        for (auto& w : words) {
            max_len = max(max_len, (int) w.length());
        }
        vector<unordered_set<int>> sets(max_len);
        for (auto& w : words) {
            long long h = 0;
            for (int j = 0; j < w.size(); j++) {
                h = (h * BASE + w[j]) % MOD;
                sets[j].insert(h); // 注意 j 从 0 开始
            }
        }

        auto max_jump = [&](int i) -> int {
            // 开区间二分，left 一定满足要求，right 一定不满足要求
            int left = 0, right = min(n - i, max_len) + 1;
            while (left + 1 < right) {
                int mid = (left + right) / 2;
                (sets[mid - 1].contains(sub_hash(i, i + mid)) ? left : right) = mid;
            }
            return left;
        };

        int ans = 0;
        int cur_r = 0; // 已建造的桥的右端点
        int nxt_r = 0; // 下一座桥的右端点的最大值
        for (int i = 0; i < n; i++) {
            nxt_r = max(nxt_r, i + max_jump(i));
            if (i == cur_r) { // 到达已建造的桥的右端点
                if (i == nxt_r) { // 无论怎么造桥，都无法从 i 到 i+1
                    return -1;
                }
                cur_r = nxt_r; // 造一座桥
                ans++;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minValidStrings(words []string, target string) (ans int) {
	n := len(target)

	// 多项式字符串哈希（方便计算子串哈希值）
	// 哈希函数 hash(s) = s[0] * base^(n-1) + s[1] * base^(n-2) + ... + s[n-2] * base + s[n-1]
	const mod = 1_070_777_777
	base := 9e8 - rand.Intn(1e8) // 随机 base，防止 hack（注意 Go1.20 之后的版本，每次随机的数都不一样）
	powBase := make([]int, n+1)  // powBase[i] = base^i
	preHash := make([]int, n+1)  // 前缀哈希值 preHash[i] = hash(target[:i])
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

	maxLen := 0
	for _, w := range words {
		maxLen = max(maxLen, len(w))
	}
	sets := make([]map[int]bool, maxLen)
	for i := range sets {
		sets[i] = map[int]bool{}
	}
	for _, w := range words {
		h := 0
		for j, b := range w {
			h = (h*base + int(b)) % mod
			sets[j][h] = true // 注意 j 从 0 开始
		}
	}

	curR := 0 // 已建造的桥的右端点
	nxtR := 0 // 下一座桥的右端点的最大值
	for i := range target {
		maxJump := sort.Search(min(n-i, maxLen), func(j int) bool { return !sets[j][subHash(i, i+j+1)] })
		nxtR = max(nxtR, i+maxJump)
		if i == curR { // 到达已建造的桥的右端点
			if i == nxtR { // 无论怎么造桥，都无法从 i 到 i+1
				return -1
			}
			curR = nxtR // 建造下一座桥
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L + n\log n)$，其中 $n$ 是 $\textit{target}$ 的长度，$L$ 是 $\textit{words}$ 中所有字符串的长度之和。
- 空间复杂度：$\mathcal{O}(L + n)$。

### 写法二：双指针

用**双指针**更新代码中的 $\textit{nxtR}$：

- 外层循环枚举 $i$，内层循环右移 $\textit{nxtR}$。
- 对于 $\textit{target}$ 的下标从 $i$ 到 $\textit{nxtR}$ 的子串，如果其哈希值在哈希表中，那么把 $\textit{nxtR}$ 加一。

```py [sol-Python3]
class Solution:
    def minValidStrings(self, words: List[str], target: str) -> int:
        n = len(target)

        # 多项式字符串哈希（方便计算子串哈希值）
        # 哈希函数 hash(s) = s[0] * BASE^(n-1) + s[1] * BASE^(n-2) + ... + s[n-2] * BASE + s[n-1]
        MOD = 1_070_777_777
        BASE = randint(8 * 10 ** 8, 9 * 10 ** 8)  # 随机 BASE，防止 hack
        pow_base = [1] + [0] * n  # pow_base[i] = BASE^i
        pre_hash = [0] * (n + 1)  # 前缀哈希值 pre_hash[i] = hash(target[:i])
        for i, b in enumerate(target):
            pow_base[i + 1] = pow_base[i] * BASE % MOD
            pre_hash[i + 1] = (pre_hash[i] * BASE + ord(b)) % MOD  # 秦九韶算法计算多项式哈希

        # 计算子串 target[l:r] 的哈希值，注意这是左闭右开区间 [l,r)
        # 计算方法类似前缀和
        def sub_hash(l: int, r: int) -> int:
            return (pre_hash[r] - pre_hash[l] * pow_base[r - l]) % MOD

        # 保存每个 words[i] 的每个前缀的哈希值，按照长度分组
        max_len = max(map(len, words))
        sets = [set() for _ in range(max_len)]
        for w in words:
            h = 0
            for j, b in enumerate(w):
                h = (h * BASE + ord(b)) % MOD
                sets[j].add(h)  # 注意 j 从 0 开始

        ans = 0
        cur_r = 0  # 已建造的桥的右端点
        nxt_r = 0  # 下一座桥的右端点的最大值
        for i in range(n):
            while nxt_r < n and nxt_r - i < max_len and sub_hash(i, nxt_r + 1) in sets[nxt_r - i]:
                nxt_r += 1  # 尽量伸长
            if i == cur_r:  # 到达已建造的桥的右端点
                if i == nxt_r:  # 无论怎么造桥，都无法从 i 到 i+1
                    return -1
                cur_r = nxt_r  # 建造下一座桥
                ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_070_777_777;

    public int minValidStrings(String[] words, String target) {
        char[] t = target.toCharArray();
        int n = t.length;

        // 多项式字符串哈希（方便计算子串哈希值）
        // 哈希函数 hash(s) = s[0] * base^(n-1) + s[1] * base^(n-2) + ... + s[n-2] * base + s[n-1]
        final int BASE = (int) 8e8 + new Random().nextInt((int) 1e8); // 随机 base，防止 hack
        int[] powBase = new int[n + 1]; // powBase[i] = base^i
        int[] preHash = new int[n + 1]; // 前缀哈希值 preHash[i] = hash(target[0] 到 target[i-1])
        powBase[0] = 1;
        for (int i = 0; i < n; i++) {
            powBase[i + 1] = (int) ((long) powBase[i] * BASE % MOD);
            preHash[i + 1] = (int) (((long) preHash[i] * BASE + t[i]) % MOD); // 秦九韶算法计算多项式哈希
        }

        int maxLen = 0;
        for (String w : words) {
            maxLen = Math.max(maxLen, w.length());
        }
        Set<Integer>[] sets = new HashSet[maxLen];
        Arrays.setAll(sets, i -> new HashSet<>());
        for (String w : words) {
            long h = 0;
            for (int j = 0; j < w.length(); j++) {
                h = (h * BASE + w.charAt(j)) % MOD;
                sets[j].add((int) h); // 注意 j 从 0 开始
            }
        }

        int ans = 0;
        int curR = 0; // 已建造的桥的右端点
        int nxtR = 0; // 下一座桥的右端点的最大值
        for (int i = 0; i < n; i++) {
            while (nxtR < n && nxtR - i < maxLen && sets[nxtR - i].contains(subHash(i, nxtR + 1, powBase, preHash))) {
                nxtR++;
            }
            if (i == curR) { // 到达已建造的桥的右端点
                if (i == nxtR) { // 无论怎么造桥，都无法从 i 到 i+1
                    return -1;
                }
                curR = nxtR; // 造一座桥
                ans++;
            }
        }
        return ans;
    }

    private int subHash(int l, int r, int[] powBase, int[] preHash) {
        return (int) ((((long) preHash[r] - (long) preHash[l] * powBase[r - l]) % MOD + MOD) % MOD);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minValidStrings(vector<string>& words, string target) {
        int n = target.length();

        // 多项式字符串哈希（方便计算子串哈希值）
        // 哈希函数 hash(s) = s[0] * base^(n-1) + s[1] * base^(n-2) + ... + s[n-2] * base + s[n-1]
        const int MOD = 1'070'777'777;
        mt19937 rng(chrono::steady_clock::now().time_since_epoch().count());
        const int BASE = uniform_int_distribution<>(8e8, 9e8)(rng); // 随机 base，防止 hack
        vector<int> pow_base(n + 1); // pow_base[i] = base^i
        vector<int> pre_hash(n + 1); // 前缀哈希值 pre_hash[i] = hash(target[:i])
        pow_base[0] = 1;
        for (int i = 0; i < n; i++) {
            pow_base[i + 1] = (long long) pow_base[i] * BASE % MOD;
            pre_hash[i + 1] = ((long long) pre_hash[i] * BASE + target[i]) % MOD; // 秦九韶算法计算多项式哈希
        }
        // 计算 target[l] 到 target[r-1] 的哈希值
        auto sub_hash = [&](int l, int r) {
            return ((pre_hash[r] - (long long) pre_hash[l] * pow_base[r - l]) % MOD + MOD) % MOD;
        };

        int max_len = 0;
        for (auto& w : words) {
            max_len = max(max_len, (int) w.length());
        }
        vector<unordered_set<int>> sets(max_len);
        for (auto& w : words) {
            long long h = 0;
            for (int j = 0; j < w.size(); j++) {
                h = (h * BASE + w[j]) % MOD;
                sets[j].insert(h); // 注意 j 从 0 开始
            }
        }

        int ans = 0;
        int cur_r = 0; // 已建造的桥的右端点
        int nxt_r = 0; // 下一座桥的右端点的最大值
        for (int i = 0; i < n; i++) {
            while (nxt_r < n && nxt_r - i < max_len && sets[nxt_r - i].contains(sub_hash(i, nxt_r + 1))) {
                nxt_r++;
            }
            if (i == cur_r) { // 到达已建造的桥的右端点
                if (i == nxt_r) { // 无论怎么造桥，都无法从 i 到 i+1
                    return -1;
                }
                cur_r = nxt_r; // 造一座桥
                ans++;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minValidStrings(words []string, target string) (ans int) {
	n := len(target)

	// 多项式字符串哈希（方便计算子串哈希值）
	// 哈希函数 hash(s) = s[0] * base^(n-1) + s[1] * base^(n-2) + ... + s[n-2] * base + s[n-1]
	const mod = 1_070_777_777
	base := 9e8 - rand.Intn(1e8) // 随机 base，防止 hack（注意 Go1.20 之后的版本，每次随机的数都不一样）
	powBase := make([]int, n+1)  // powBase[i] = base^i
	preHash := make([]int, n+1)  // 前缀哈希值 preHash[i] = hash(target[:i])
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

	maxLen := 0
	for _, w := range words {
		maxLen = max(maxLen, len(w))
	}
	sets := make([]map[int]bool, maxLen)
	for i := range sets {
		sets[i] = map[int]bool{}
	}
	for _, w := range words {
		h := 0
		for j, b := range w {
			h = (h*base + int(b)) % mod
			sets[j][h] = true // 注意 j 从 0 开始
		}
	}

	curR := 0 // 已建造的桥的右端点
	nxtR := 0 // 下一座桥的右端点的最大值
	for i := range target {
		for nxtR < n && nxtR-i < maxLen && sets[nxtR-i][subHash(i, nxtR+1)] {
			nxtR++
		}
		if i == curR { // 到达已建造的桥的右端点
			if i == nxtR { // 无论怎么造桥，都无法从 i 到 i+1
				return -1
			}
			curR = nxtR // 建造下一座桥
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L + n)$，其中 $n$ 是 $\textit{target}$ 的长度，$L$ 是 $\textit{words}$ 中所有字符串的长度之和。
- 空间复杂度：$\mathcal{O}(L + n)$。

## 方法三：AC 自动机优化 DP

看示例 1，对比以下两个 $\textit{target}$ 的前缀：

- $\texttt{aabcd}$，需要连接 $2$ 个 $\textit{words}[i]$ 的前缀 $\texttt{aa}$ 和 $\texttt{bcd}$。
- $\texttt{aab}$，需要连接多少个前缀？可以肯定的是，答案一定不会比 $2$ 还大，因为我们把 $\texttt{aabcd}$ 末尾的 $\texttt{cd}$ 去掉就可以得到 $\texttt{aab}$。这仍然是 $2$ 个前缀的连接。

根据上述讨论，如果用 $f[i]$ 表示 $\textit{target}$ 的长为 $i$ 的前缀需要连接的最少字符串数量，那么 $f[i]\le f[i+1]$ 一定成立。

既然 $f$ 是有序数组，那么对于 $f[i]$，我们只需要知道最小的 $j$，满足从 $\textit{target}[j]$ 到 $\textit{target}[i-1]$ 是某个 $\textit{words}[i]$ 的前缀。

也就是说，匹配的 $\textit{words}[i]$ 的前缀要尽量长。这正是 **AC 自动机**的应用。原理见 [OI Wiki](https://oi-wiki.org/string/ac-automaton/)。学习之前推荐先看看我的 [KMP 原理讲解](https://www.zhihu.com/question/21923021/answer/37475572)。

算出了 $j$，那么有

$$
f[i] = f[j] + 1
$$

初始值 $f[0]=0$。

答案为 $f[n]$。

如果 AC 自动机没法匹配任何 $\textit{words}[i]$ 的非空前缀，返回 $-1$。

```py [sol-Python3]
# 从根到 node 的字符串是某个（某些）words[i] 的前缀
class Node:
    __slots__ = 'son', 'fail', 'len'

    def __init__(self, len=0):
        self.son = [None] * 26
        self.fail = None  # 当 cur.son[i] 不能匹配 target 中的某个字符时，cur.fail.son[i] 即为下一个待匹配节点（等于 root 则表示没有匹配）
        self.len = len  # 从根到 node 的字符串的长度，也是 node 在 trie 中的深度

class AhoCorasick:
    def __init__(self):
        self.root = Node()

    def put(self, s: str) -> None:
        cur = self.root
        for b in s:
            b = ord(b) - ord('a')
            if cur.son[b] is None:
                cur.son[b] = Node(cur.len + 1)
            cur = cur.son[b]

    def build_fail(self) -> None:
        self.root.fail = self.root
        q = deque()
        for i, son in enumerate(self.root.son):
            if son is None:
                self.root.son[i] = self.root
            else:
                son.fail = self.root  # 第一层的失配指针，都指向根节点 ∅
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
                q.append(son)

class Solution:
    def minValidStrings(self, words: List[str], target: str) -> int:
        ac = AhoCorasick()
        for w in words:
            ac.put(w)
        ac.build_fail()

        n = len(target)
        f = [0] * (n + 1)
        cur = root = ac.root
        for i, c in enumerate(target, 1):
            # 如果没有匹配相当于移动到 fail 的 son[c]
            cur = cur.son[ord(c) - ord('a')]
            # 没有任何字符串的前缀与 target[..i] 的后缀匹配
            if cur is root:
                return -1
            f[i] = f[i - cur.len] + 1
        return f[n]
```

```java [sol-Java]
// 从根到 node 的字符串是某个（某些）words[i] 的前缀
class Node {
    Node[] son = new Node[26];
    Node fail; // 当 cur.son[i] 不能匹配 target 中的某个字符时，cur.fail.son[i] 即为下一个待匹配节点（等于 root 则表示没有匹配）
    int len;

    Node(int len) {
        this.len = len;
    }
}

class AhoCorasick {
    Node root = new Node(0);

    void put(String s) {
        Node cur = root;
        for (char b : s.toCharArray()) {
            b -= 'a';
            if (cur.son[b] == null) {
                cur.son[b] = new Node(cur.len + 1);
            }
            cur = cur.son[b];
        }
    }

    void buildFail() {
        root.fail = root;
        Queue<Node> q = new ArrayDeque<>();
        for (int i = 0; i < root.son.length; i++) {
            Node son = root.son[i];
            if (son == null) {
                root.son[i] = root;
            } else {
                son.fail = root; // 第一层的失配指针，都指向根节点 ∅
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
                q.add(son);
            }
        }
    }
}

class Solution {
    public int minValidStrings(String[] words, String target) {
        AhoCorasick ac = new AhoCorasick();
        for (String w : words) {
            ac.put(w);
        }
        ac.buildFail();

        char[] t = target.toCharArray();
        int n = t.length;
        int[] f = new int[n + 1];
        Node cur = ac.root;
        for (int i = 0; i < n; i++) {
            // 如果没有匹配相当于移动到 fail 的 son[t[i]-'a']
            cur = cur.son[t[i] - 'a'];
            // 没有任何字符串的前缀与 target[..i] 的后缀匹配
            if (cur == ac.root) {
                return -1;
            }
            f[i + 1] = f[i + 1 - cur.len] + 1;
        }
        return f[n];
    }
}
```

```cpp [sol-C++]
// 从根到 node 的字符串是某个（某些）words[i] 的前缀
struct Node {
    Node* son[26]{};
    Node* fail; // 当 cur.son[i] 不能匹配 target 中的某个字符时，cur.fail.son[i] 即为下一个待匹配节点（等于 root 则表示没有匹配）
    int len; // 从根到 node 的字符串的长度，也是 node 在 trie 中的深度

    Node(int len) : len(len) {}
};

struct AhoCorasick {
    Node* root = new Node(0);

    void put(string& s) {
        auto cur = root;
        for (char b : s) {
            b -= 'a';
            if (cur->son[b] == nullptr) {
                cur->son[b] = new Node(cur->len + 1);
            }
            cur = cur->son[b];
        }
    }

    void build_fail() {
        root->fail = root;
        queue<Node*> q;
        for (auto& son : root->son) {
            if (son == nullptr) {
                son = root;
            } else {
                son->fail = root; // 第一层的失配指针，都指向根节点 ∅
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
                q.push(son);
            }
        }
    }
};

class Solution {
public:
    int minValidStrings(vector<string>& words, string target) {
        AhoCorasick ac;
        for (auto& w : words) {
            ac.put(w);
        }
        ac.build_fail();

        int n = target.length();
        vector<int> f(n + 1);
        auto cur = ac.root;
        for (int i = 0; i < n; i++) {
            // 如果没有匹配相当于移动到 fail 的 son[target[i]-'a']
            cur = cur->son[target[i] - 'a'];
            // 没有任何字符串的前缀与 target[..i] 的后缀匹配
            if (cur == ac.root) {
                return -1;
            }
            f[i + 1] = f[i + 1 - cur->len] + 1;
        }
        return f[n];
    }
};
```

```go [sol-Go]
// 从根到 node 的字符串是某个（某些）words[i] 的前缀
type node struct {
	son  [26]*node
	fail *node // 当 cur.son[i] 不能匹配 target 中的某个字符时，cur.fail.son[i] 即为下一个待匹配节点（等于 root 则表示没有匹配）
	len  int   // 从根到 node 的字符串的长度，也是 node 在 trie 中的深度
}

type acam struct {
	root *node
}

func (ac *acam) put(s string) {
	cur := ac.root
	for _, b := range s {
		b -= 'a'
		if cur.son[b] == nil {
			cur.son[b] = &node{len: cur.len + 1}
		}
		cur = cur.son[b]
	}
}

func (ac *acam) buildFail() {
	ac.root.fail = ac.root
	q := []*node{}
	for i, son := range ac.root.son[:] {
		if son == nil {
			ac.root.son[i] = ac.root
		} else {
			son.fail = ac.root // 第一层的失配指针，都指向根节点 ∅
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
			q = append(q, son)
		}
	}
}

func minValidStrings(words []string, target string) int {
	ac := &acam{root: &node{}}
	for _, w := range words {
		ac.put(w)
	}
	ac.buildFail()

	n := len(target)
	f := make([]int, n+1)
	cur := ac.root
	for i, b := range target {
		// 如果没有匹配相当于移动到 fail 的 son[b-'a']
		cur = cur.son[b-'a']
		// 没有任何字符串的前缀与 target[:i+1] 的后缀匹配
		if cur == ac.root {
			return -1
		}
		f[i+1] = f[i+1-cur.len] + 1
	}
	return f[n]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L|\Sigma| + n)$，其中 $n$ 是 $\textit{target}$ 的长度，$L$ 是 $\textit{words}$ 中所有字符串的长度之和，$|\Sigma|$ 是字符集合的大小，本题字符均为小写字母，所以 $|\Sigma|=26$。如果把数组替换成哈希表，可以做到 $\mathcal{O}(L+n)$ 的时间复杂度。
- 空间复杂度：$\mathcal{O}(L|\Sigma| + n)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. 【本题相关】[字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
