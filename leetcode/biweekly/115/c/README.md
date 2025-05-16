## 子序列 DP 的常用套路

- 子序列 + 相邻元素无关：**选或不选**。代表题目：[494. 目标和（0-1 背包）](https://leetcode.cn/problems/target-sum/)。
- 子序列 + 相邻元素相关：**枚举选哪个**。代表题目：[300. 最长递增子序列](https://leetcode.cn/problems/longest-increasing-subsequence/)。
- 值域 DP。具体在方法二中介绍。

## 方法一：枚举选哪个

本题属于「子序列 + 相邻元素相关」，用**枚举选哪个**解决。

类似最长递增子序列，定义 $f[i]$ 表示在子序列第一个字符串是 $\textit{words}[i]$ 的前提下，从后缀 $[i,n-1]$ 中能选出的最长子序列的长度。为什么定义成后缀？方便后面输出具体方案。

枚举子序列的第二个字符串 $\textit{words}[j]$，如果 $\textit{groups}[j] \ne \textit{groups}[i]$ 且 $\textit{words}[j]$ 和 $\textit{words}[i]$ 满足题目要求（长度相等且汉明距离为 $1$），那么问题变成在子序列第一个字符串是 $\textit{words}[j]$ 的前提下，从后缀 $[j,n-1]$ 中能选出的最长子序列的长度，即 $f[j]$。在 $f[j]$ 子序列的前面加上 $\textit{words}[i]$，得 $f[i] = f[j]+1$。

取最大值，得

$$
f[i] = 1 + \max_{j=i+1}^{n-1} f[j]
$$

上式中的 $j$ 需要满足前文说的要求。

计算 $\max$ 的过程中，记录最优转移来源 $\textit{from}[i] = j$，方便输出具体方案。

初始值：$f[i]=1$，表示 $\textit{words}[i]$ 单独组成一个长为 $1$ 的子序列。

答案：$\max(f)$。

### 如何输出方案

设 $\textit{maxI}$ 是 $\max(f)$ 在 $f$ 中的下标，即 $f[\textit{maxI}]=\max(f)$。如果有多个这样的下标，随便取哪个都行。

从 $i=\textit{maxI}$ 开始循环，每次把 $\textit{words}[i]$ 加入答案，然后更新

$$
i = \textit{from}[i]
$$

表示顺着转移来源往右走，找子序列的下一个字符串。

找到 $\max(f)$ 个字符串时，退出循环。

```py [sol-Python3]
class Solution:
    def getWordsInLongestSubsequence(self, words: List[str], groups: List[int]) -> List[str]:
        def check(s: str, t: str) -> bool:
            return len(s) == len(t) and sum(x != y for x, y in zip(s, t)) == 1

        n = len(words)
        f = [0] * n
        from_ = [0] * n
        max_i = n - 1
        for i in range(n - 1, -1, -1):
            for j in range(i + 1, n):
                # 提前比较 f[j] 与 f[i] 的大小，如果 f[j] <= f[i]，就不用执行更耗时的 check 了
                if f[j] > f[i] and groups[j] != groups[i] and check(words[i], words[j]):
                    f[i] = f[j]
                    from_[i] = j
            f[i] += 1  # 加一写在这里
            if f[i] > f[max_i]:
                max_i = i

        i = max_i
        ans = [''] * f[i]
        for k in range(f[i]):
            ans[k] = words[i]
            i = from_[i]
        return ans
```

```java [sol-Java]
class Solution {
    public List<String> getWordsInLongestSubsequence(String[] words, int[] groups) {
        int n = words.length;
        int[] f = new int[n];
        int[] from = new int[n];
        int maxI = n - 1;
        for (int i = n - 1; i >= 0; i--) {
            for (int j = i + 1; j < n; j++) {
                // 提前比较 f[j] 与 f[i] 的大小，如果 f[j] <= f[i]，就不用执行更耗时的 check 了
                if (f[j] > f[i] && groups[j] != groups[i] && check(words[i], words[j])) {
                    f[i] = f[j];
                    from[i] = j;
                }
            }
            f[i]++; // 加一写在这里
            if (f[i] > f[maxI]) {
                maxI = i;
            }
        }

        int i = maxI;
        int m = f[i];
        List<String> ans = new ArrayList<>(m); // 预分配空间
        for (int k = 0; k < m; k++) {
            ans.add(words[i]);
            i = from[i];
        }
        return ans;
    }

    private boolean check(String s, String t) {
        if (s.length() != t.length()) {
            return false;
        }
        boolean diff = false;
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) != t.charAt(i)) {
                if (diff) { // 汉明距离大于 1
                    return false;
                }
                diff = true;
            }
        }
        return diff;
    }
}
```

```cpp [sol-C++]
class Solution {
    bool ok(string& s, string& t) {
        if (s.size() != t.size()) {
            return false;
        }
        bool diff = false;
        for (int i = 0; i < s.size(); i++) {
            if (s[i] != t[i]) {
                if (diff) { // 汉明距离大于 1
                    return false;
                }
                diff = true;
            }
        }
        return diff;
    }

public:
    vector<string> getWordsInLongestSubsequence(vector<string>& words, vector<int>& groups) {
        int n = words.size();
        vector<int> f(n), from(n);
        int max_i = n - 1;
        for (int i = n - 1; i >= 0; i--) {
            for (int j = i + 1; j < n; j++) {
                // 提前比较 f[j] 与 f[i] 的大小，如果 f[j] <= f[i]，就不用执行更耗时的 check 了
                if (f[j] > f[i] && groups[j] != groups[i] && ok(words[i], words[j])) {
                    f[i] = f[j];
                    from[i] = j;
                }
            }
            f[i]++; // 加一写在这里
            if (f[i] > f[max_i]) {
                max_i = i;
            }
        }

        int i = max_i;
        int m = f[i];
        vector<string> ans(m);
        for (int k = 0; k < m; k++) {
            ans[k] = words[i];
            i = from[i];
        }
        return ans;
    }
};
```

```go [sol-Go]
func ok(s, t string) (diff bool) {
	if len(s) != len(t) {
		return
	}
	for i := range s {
		if s[i] != t[i] {
			if diff { // 汉明距离大于 1
				return false
			}
			diff = true
		}
	}
	return
}

func getWordsInLongestSubsequence(words []string, groups []int) []string {
	n := len(words)
	f := make([]int, n)
	from := make([]int, n)
	maxI := n - 1
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			// 提前比较 f[j] 与 f[i] 的大小，如果 f[j] <= f[i]，就不用执行更耗时的 check 了
			if f[j] > f[i] && groups[j] != groups[i] && ok(words[i], words[j]) {
				f[i] = f[j]
				from[i] = j
			}
		}
		f[i]++ // 加一写在这里
		if f[i] > f[maxI] {
			maxI = i
		}
	}

	i := maxI
	ans := make([]string, f[i])
	for k := range ans {
		ans[k] = words[i]
		i = from[i]
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2l)$，其中 $n$ 是 $\textit{words}$ 的长度，$l\le 10$ 为 $\textit{words}[i]$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：值域 DP

### 1

如果 $n=10^5$，方法一就超时了。怎么办？

在方法一中，我们其实没有充分利用这个性质：汉明距离只有 $1$，也就是只有一个字母不同。如何利用这个性质？

方法一是枚举字符串，计算不同字母个数。横看成岭侧成峰，反过来，枚举哪个字母不同。

字符串 $w=\textit{words}[i]$ 的长度至多为 $10$，与 $w$ 只有一个字母不同的字符串，本质只有 $10$ 种：

- 与 $w[0]$ 不同的字符串。
- 与 $w[1]$ 不同的字符串。
- ……
- 与 $w[9]$ 不同的字符串。

**关键思路**：我们只需枚举这 $10$ 种字符串，而不是 $\mathcal{O}(n)$ 个字符串！

### 2

首先考虑一个简单情况，所有 $\textit{groups}[i]$ 互不相同。此时只需保证子序列相邻字符串的汉明距离为 $1$。

设子序列中的一对相邻字符串为 $s$ 和 $t$。

比如 $s=\texttt{bab}$，那么 $t$ 需要与 $s$ 恰好有一个字符不同，例如 $t=\texttt{cab}$。

我们可以把 $s$ 视作三种字符串：$\texttt{?ab},\ \texttt{b?b},\ \texttt{ba?}$，其中 $\texttt{?}$ 表示通配符，可以匹配任意单个字符。

同样地，$t=\texttt{cab}$ 也可以视作 $\texttt{?ab}$，这样就与 $s=\texttt{?ab}$ **相等**了。

⚠**注意**：题目保证所有字符串互不相同。用「通配符」转化后，两个相等的字符串的汉明距离一定恰好等于 $1$。

### 3

子序列 DP 的另一个套路是，把元素值作为 DP 的状态。（方法一是把下标作为状态）

本题元素是字符串，可以用哈希表，把字符串作为哈希表的 key，子序列长度作为哈希表的 value。

具体地，定义 $\textit{fMap}[s]$ 表示以字符串 $s$ 开头的合法子序列的最长长度，从子序列的第二个字符串转移过来。其中 $s$ 包含一个 $\texttt{?}$ 通配符。

同方法一，倒着遍历 $\textit{words}$，设 $w=\textit{words}[i]$，设把 $w[k]$ 改成 $\texttt{?}$ 后的字符串为 $w'$。

1. 计算 $\textit{fMap}[w']$ 的最大值。例如 $w=\texttt{bab}$，计算的是
    $$
    \max(\textit{fMap}[\texttt{?ab}],\textit{fMap}[\texttt{b?b}],\textit{fMap}[\texttt{ba?}])
    $$
2. 把上式加一，就是方法一费了九牛二虎之力 $\mathcal{O}(nl)$ 算出的 $f[i]$。方法二只需要 $\mathcal{O}(l)$ 的时间。
3. 用上式加一的结果，再去更新 $\textit{fMap}[w']$ 的最大值。接着上面的例子，就是更新 $\textit{fMap}[\texttt{?ab}],\textit{fMap}[\texttt{b?b}],\textit{fMap}[\texttt{ba?}]$ 的最大值。

例如 $\textit{words}=[\texttt{aab},\texttt{aaa},\texttt{baa}]$ 的答案是 $3$，读者可以动手算算这个例子，体会 $\texttt{aaa}$ 是如何作为「桥梁」连接 $\texttt{aab}$ 和 $\texttt{baa}$ 的。

### 4

把 $\textit{groups}$ 的约束加进来。

在上述方法的基础上，额外维护**次大值**，并保证最大值和次大值对应的 $\textit{group}$ 值不同，从而保证最大值和次大值中一定有一个可以转移到当前状态。

具体来说，$\textit{fMap}[s]$ 保存 $4$ 个值：

- 以 $s$ 为首的最大子序列长度 $\textit{maxF}$，以及这个 $s$ 的下标 $j$。
- 以 $s$ 为首的次大子序列长度 $\textit{maxF}_2$，以及这个 $s$ 的下标 $j_2$。
- 要求：$\textit{groups}[j]\ne \textit{groups}[j_2]$。

最终答案就是 $\textit{fMap}[s].\textit{maxF}$ 的最大值。

### 5

代码实现时，由于字符串长度 $\le 10$，可以把字符串压缩成一个长度 $\le 50$ 的二进制数（完美哈希），作为哈希表的 key：

- 每个字母用 $5$ 个比特存储。
- 字母 $\texttt{a}$ 到 $\texttt{z}$ 对应数字 $1$ 到 $26$。计算时，可以直接取 ASCII 值的低 $5$ 位。
- $\texttt{?}$ 对应数字 $31$，也就是二进制数 $11111$。
- 把字母改成 $\texttt{?}$，可以把对应比特位 OR 上 $11111$。形象地理解为，用记号笔把这个字母涂黑。

```py [sol-Python3]
class Solution:
    def getWordsInLongestSubsequence(self, words: List[str], groups: List[int]) -> List[str]:
        n = len(words)
        f_map = {}  # 哈希值 -> (max_f, j, max_f2, j2)
        from_ = [0] * n
        global_max_f = max_i = 0
        for i in range(n - 1, -1, -1):
            w, g = words[i], groups[i]

            # 计算 w 的哈希值
            hash_val = sum((ord(ch) & 31) << (k * 5) for k, ch in enumerate(w))

            f = 0  # 方法一中的 f[i]
            for k in range(len(w)):
                h = hash_val | (31 << (k * 5))  # 用记号笔把 w[k] 涂黑（置为 11111）
                max_f, j, max_f2, j2 = f_map.get(h, (0, 0, 0, 0))
                if g != groups[j]:  # 可以从最大值转移过来
                    if max_f > f:
                        f = max_f
                        from_[i] = j
                else:  # 只能从次大值转移过来
                    if max_f2 > f:
                        f = max_f2
                        from_[i] = j2

            f += 1
            if f > global_max_f:
                global_max_f, max_i = f, i

            # 用 f 更新 f_map[h] 的最大次大
            # 注意要保证最大次大的 group 值不同
            for k in range(len(w)):
                h = hash_val | (31 << (k * 5))
                max_f, j, max_f2, j2 = f_map.get(h, (0, 0, 0, 0))
                if f > max_f:  # 最大值需要更新
                    if g != groups[j]:
                        max_f2, j2 = max_f, j  # 旧最大值变成次大值
                    max_f, j = f, i
                elif f > max_f2 and g != groups[j]:  # 次大值需要更新
                    max_f2, j2 = f, i
                f_map[h] = (max_f, j, max_f2, j2)

        ans = [''] * global_max_f
        i = max_i
        for k in range(global_max_f):
            ans[k] = words[i]
            i = from_[i]
        return ans
```

```java [sol-Java]
class Solution {
    private record Info(int maxF, int j, int maxF2, int j2) {
    }

    public List<String> getWordsInLongestSubsequence(String[] words, int[] groups) {
        int n = words.length;
        Map<Long, Info> fMap = new HashMap<>();
        int[] from = new int[n];
        int globalMaxF = 0;
        int maxI = 0;

        for (int i = n - 1; i >= 0; i--) {
            char[] w = words[i].toCharArray();
            int g = groups[i];

            // 计算 w 的哈希值
            long hash = 0;
            for (char c : w) {
                hash = (hash << 5) | (c & 31);
            }

            int f = 0; // 方法一中的 f[i]
            for (int k = 0; k < w.length; k++) {
                long h = hash | (31L << (k * 5)); // 用记号笔把 w[k] 涂黑（置为 11111）
                Info t = fMap.get(h);
                if (t == null) {
                    continue;
                }
                if (g != groups[t.j]) { // 可以从最大值转移过来
                    if (t.maxF > f) {
                        f = t.maxF;
                        from[i] = t.j;
                    }
                } else { // 只能从次大值转移过来
                    if (t.maxF2 > f) {
                        f = t.maxF2;
                        from[i] = t.j2;
                    }
                }
            }

            f++;
            if (f > globalMaxF) {
                globalMaxF = f;
                maxI = i;
            }

            // 用 f 更新 fMap[h] 的最大次大
            // 注意要保证最大次大的 group 值不同
            for (int k = 0; k < w.length; k++) {
                long h = hash | (31L << (k * 5));
                Info t = fMap.getOrDefault(h, new Info(0, 0, 0, 0));
                int maxF = t.maxF, j = t.j, maxF2 = t.maxF2, j2 = t.j2;
                if (f > maxF) { // 最大值需要更新
                    if (g != groups[j]) {
                        maxF2 = maxF; // 旧最大值变成次大值
                        j2 = j;
                    }
                    maxF = f;
                    j = i;
                } else if (f > maxF2 && g != groups[j]) { // 次大值需要更新
                    maxF2 = f;
                    j2 = i;
                }
                fMap.put(h, new Info(maxF, j, maxF2, j2));
            }
        }

        List<String> ans = new ArrayList<>(globalMaxF); // 预分配空间
        int i = maxI;
        for (int k = 0; k < globalMaxF; k++) {
            ans.add(words[i]);
            i = from[i];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<string> getWordsInLongestSubsequence(vector<string>& words, vector<int>& groups) {
        int n = words.size();
        unordered_map<long long, tuple<int, int, int, int>> f_map; // 哈希值 -> (max_f, j, max_f2, j2)
        vector<int> from(n);
        int global_max_f = 0, max_i = 0;
        for (int i = n - 1; i >= 0; i--) {
            string& w = words[i];
            int g = groups[i];

            // 计算 w 的哈希值
            long long hash = 0;
            for (char ch : w) {
                hash = (hash << 5) | (ch & 31);
            }

            int f = 0; // 方法一中的 f[i]
            for (int k = 0; k < w.size(); k++) {
                long long h = hash | (31LL << (k * 5)); // 用记号笔把 w[k] 涂黑（置为 11111）
                auto& [max_f, j, max_f2, j2] = f_map[h];
                if (g != groups[j]) { // 可以从最大值转移过来
                    if (max_f > f) {
                        f = max_f;
                        from[i] = j;
                    }
                } else { // 只能从次大值转移过来
                    if (max_f2 > f) {
                        f = max_f2;
                        from[i] = j2;
                    }
                }
            }

            f++;
            if (f > global_max_f) {
                global_max_f = f;
                max_i = i;
            }

            // 用 f 更新 f_map[h] 的最大次大
            // 注意要保证最大次大的 group 值不同
            for (int k = 0; k < w.size(); k++) {
                long long h = hash | (31LL << (k * 5));
                auto& [max_f, j, max_f2, j2] = f_map[h];
                if (f > max_f) { // 最大值需要更新
                    if (g != groups[j]) {
                        max_f2 = max_f; // 旧最大值变成次大值
                        j2 = j;
                    }
                    max_f = f;
                    j = i;
                } else if (f > max_f2 && g != groups[j]) { // 次大值需要更新
                    max_f2 = f;
                    j2 = i;
                }
            }
        }

        vector<string> ans(global_max_f);
        int i = max_i;
        for (int k = 0; k < global_max_f; k++) {
            ans[k] = words[i];
            i = from[i];
        }
        return ans;
    }
};
```

```go [sol-Go]
func getWordsInLongestSubsequence(words []string, groups []int) []string {
	n := len(words)
	fMap := map[int]struct{ maxF, j, maxF2, j2 int }{}
	from := make([]int, n)
	maxF, maxI := 0, 0
	for i := n - 1; i >= 0; i-- {
		w, g := words[i], groups[i]

		// 计算 w 的哈希值
		hash := 0
		for _, ch := range w {
			hash = hash<<5 | int(ch&31)
		}

		f := 0 // 方法一中的 f[i]
		for j := range w {
			h := hash | 31<<(j*5) // 用记号笔把 w[k] 涂黑（置为 11111）
			t := fMap[h]
			if g != groups[t.j] { // 可以从最大值转移过来
				if t.maxF > f {
					f = t.maxF
					from[i] = t.j
				}
			} else { // 只能从次大值转移过来
				if t.maxF2 > f {
					f = t.maxF2
					from[i] = t.j2
				}
			}
		}

		f++
		if f > maxF {
			maxF, maxI = f, i
		}

		// 用 f 更新 fMap[h] 的最大次大
		// 注意要保证最大次大的 group 值不同
		for j := range w {
			h := hash | 31<<(j*5)
			t := fMap[h]
			if f > t.maxF { // 最大值需要更新
				if g != groups[t.j] {
					t.maxF2 = t.maxF // 旧最大值变成次大值
					t.j2 = t.j
				}
				t.maxF = f
				t.j = i
			} else if f > t.maxF2 && g != groups[t.j] { // 次大值需要更新
				t.maxF2 = f
				t.j2 = i
			}
			fMap[h] = t
		}
	}

	ans := make([]string, maxF)
	i := maxI
	for k := range ans {
		ans[k] = words[i]
		i = from[i]
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nl)$，其中 $n$ 是 $\textit{words}$ 的长度，$l\le 10$ 为 $\textit{words}[i]$ 的长度。这是线性时间复杂度，与输入量成正比。
- 空间复杂度：$\mathcal{O}(nl)$。

更多相似题目，见动态规划题单的「**§4.2 最长递增子序列**」和「**§7.4 合法子序列 DP**」。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
