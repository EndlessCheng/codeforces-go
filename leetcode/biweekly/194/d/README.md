三变量问题，枚举中间的 $j$ 更方便。

问题变成：

1. 计算右端点为 $j$ 的子串的字母种数。
2. 计算左端点为 $j+1$ 的子串的字母种数。

## 方法一：恰好型滑动窗口 

本题可以拆分成如下 $26$ 个滑动窗口问题：

- 计算恰好有 $k$ 种字母的子串个数。其中 $k=1,2,\ldots,26$。

关于**恰好型滑动窗口**，见 [930 题的题解](https://leetcode.cn/problems/binary-subarrays-with-sum/solutions/3986900/qia-hao-xing-hua-dong-chuang-kou-pythonj-3i6o/)。

考察滑动窗口的过程，我们实际上对每个右端点 $i=0,1,2,\ldots,n-1$，都算出了有多少个合法子串。

所以在从左到右滑动窗口的过程中，我们可以得到 $\textit{pre}[i][j]$，表示子串右端点为 $i$ 时，有多少个子串恰好有 $j$ 种字母。

同理，倒着滑窗，可以得到 $\textit{suf}[i][j]$，表示子串左端点为 $i$ 时，有多少个子串恰好有 $j$ 种字母。

以 $(i,i+1)$ 为分割线，左右子串互相独立，根据乘法原理，分割线为 $(i,i+1)$ 且字母种数恰好为 $j$ 的合法子串对的个数为

$$
\textit{pre}[i][j] \cdot \textit{suf}[i+1][j]
$$

枚举分割线以及 $j$，答案为

$$
\sum_{i=0}^{n-2}\sum_{j=1}^{26} \textit{pre}[i][j] \cdot \textit{suf}[i+1][j]
$$

[本题视频讲解](https://www.bilibili.com/video/BV15ZH36UELN/?t=4m47s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def sliding_window(self, s: list[int]) -> list[list[int]]:
        res = [[0] * 26 for _ in s]

        # 计算至少有 k 种字母的子串个数
        for k in range(1, 27):
            cnt = [0] * 26
            kinds = 0
            left = 0
            for i, ch in enumerate(s):
                if cnt[ch] == 0:
                    kinds += 1
                cnt[ch] += 1
                while kinds >= k:
                    cnt[s[left]] -= 1
                    if cnt[s[left]] == 0:
                        kinds -= 1
                    left += 1
                res[i][k - 1] = left  # 右端点为 i 的合法子串个数
            if left == 0:  # 没有 k 种字母，也没有 > k 种字母，不用再算
                break

        for cnt in res:
            # (至少 j 种) - (至少 j+1 种) = (恰好 j 种)
            for j in range(25):
                if cnt[j + 1] == 0:
                    break
                cnt[j] -= cnt[j + 1]

        return res

    def countMagicalPairs(self, s: str) -> int:
        s = [ord(ch) - ord('a') for ch in s]
        pre = self.sliding_window(s)
        s.reverse()
        suf = self.sliding_window(s)  # 相当于倒着滑窗
        suf.reverse()

        ans = 0
        for i in range(len(pre) - 1):
            for p, s in zip(pre[i], suf[i + 1]):
                if p == 0 or s == 0:
                    break
                ans += p * s
        return ans
```

```java [sol-Java]
class Solution {
    public long countMagicalPairs(String S) {
        char[] s = S.toCharArray();
        int[][] pre = slidingWindow(s);

        int n = s.length;
        // 反转 s
        for (int i = 0; i < n / 2; i++) {
            char tmp = s[i];
            s[i] = s[n - 1 - i];
            s[n - 1 - i] = tmp;
        }
        int[][] suf = slidingWindow(s); // 相当于倒着滑窗
        // 反转 suf
        for (int i = 0; i < n / 2; i++) {
            int[] tmp = suf[i];
            suf[i] = suf[n - 1 - i];
            suf[n - 1 - i] = tmp;
        }

        long ans = 0;
        for (int i = 0; i < n - 1; i++) {
            for (int j = 0; j < 26 && pre[i][j] > 0 && suf[i + 1][j] > 0; j++) {
                ans += (long) pre[i][j] * suf[i + 1][j];
            }
        }
        return ans;
    }

    private int[][] slidingWindow(char[] s) {
        int n = s.length;
        int[][] res = new int[n][26];

        // 计算至少有 k 种字母的子串个数
        for (int k = 1; k <= 26; k++) {
            int[] cnt = new int[26];
            int kinds = 0;
            int left = 0;
            for (int i = 0; i < n; i++) {
                if (cnt[s[i] - 'a']++ == 0) {
                    kinds++;
                }
                while (kinds >= k) {
                    if (--cnt[s[left] - 'a'] == 0) {
                        kinds--;
                    }
                    left++;
                }
                res[i][k - 1] = left; // 右端点为 i 的合法子串个数
            }
            if (left == 0) { // 没有 k 种字母，也没有 > k 种字母，不用再算
                break;
            }
        }

        for (int[] cnt : res) {
            // (至少 j 种) - (至少 j+1 种) = (恰好 j 种)
            for (int j = 0; j < 25; j++) {
                cnt[j] -= cnt[j + 1];
            }
        }

        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<array<int, 26>> sliding_window(const string& s) {
        int n = s.size();
        vector<array<int, 26>> res(n);

        // 计算至少有 k 种字母的子串个数
        for (int k = 1; k <= 26; k++) {
            int cnt[26]{};
            int kinds = 0;
            int left = 0;
            for (int i = 0; i < n; i++) {
                if (cnt[s[i] - 'a']++ == 0) {
                    kinds++;
                }
                while (kinds >= k) {
                    if (--cnt[s[left] - 'a'] == 0) {
                        kinds--;
                    }
                    left++;
                }
                res[i][k - 1] = left; // 右端点为 i 的合法子串个数
            }
            if (left == 0) { // 没有 k 种字母，也没有 > k 种字母，不用再算
                break;
            }
        }

        for (auto& cnt : res) {
            // (至少 j 种) - (至少 j+1 种) = (恰好 j 种)
            for (int j = 0; j < 25; j++) {
                cnt[j] -= cnt[j + 1];
            }
        }

        return res;
    }

public:
    long long countMagicalPairs(string s) {
        auto pre = sliding_window(s);
        ranges::reverse(s);
        auto suf = sliding_window(s); // 相当于倒着滑窗
        ranges::reverse(suf);

        int n = s.size();
        long long ans = 0;
        for (int i = 0; i < n - 1; i++) {
            for (int j = 0; j < 26 && pre[i][j] > 0 && suf[i + 1][j] > 0; j++) {
                ans += 1LL * pre[i][j] * suf[i + 1][j];
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func slidingWindow(s []byte) [][26]int {
	res := make([][26]int, len(s))

	// 计算至少有 k 种字母的子串个数
	for k := 1; k <= 26; k++ {
		cnt := [26]int{}
		kinds := 0
		left := 0
		for i, ch := range s {
			if cnt[ch-'a'] == 0 {
				kinds++
			}
			cnt[ch-'a']++
			for kinds >= k {
				cnt[s[left]-'a']--
				if cnt[s[left]-'a'] == 0 {
					kinds--
				}
				left++
			}
			res[i][k-1] = left // 右端点为 i 的合法子串个数
		}
		if left == 0 { // 没有 k 种字母，也没有 > k 种字母，不用再算
			break
		}
	}

	for i := range res {
		// (至少 j 种) - (至少 j+1 种) = (恰好 j 种)
		for j := range 25 {
			res[i][j] -= res[i][j+1]
		}
	}

	return res
}

func countMagicalPairs(S string) (ans int64) {
	s := []byte(S)
	pre := slidingWindow(s)
	slices.Reverse(s)
	suf := slidingWindow(s) // 相当于倒着滑窗
	slices.Reverse(suf)

	for i := range len(pre) - 1 {
		for j := 0; j < 26 && pre[i][j] > 0 && suf[i+1][j] > 0; j++ {
			ans += int64(pre[i][j]) * int64(suf[i+1][j])
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(n|\Sigma|)$。

## 方法二：OR LogTrick

**前置知识**：

1. [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/discuss/post/3571304/cong-ji-he-lun-dao-wei-yun-suan-chang-ji-enve/)。
2. [LogTrick 入门教程](https://zhuanlan.zhihu.com/p/1933215367158830792)。

把 $s[i]$ 视作 `1 << (s[i] - 'a')`，那么子串中的字母种数等于：

- 子串的按位或（OR）的二进制中的 $1$ 的个数。

所以本题是一个子串（子数组）OR 问题，套 LogTrick 模板即可。

```py [sol-Python3]
class Solution:
    # 模板来源：https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/
    def log_trick(self, nums: list[int]) -> list[list[int]]:
        res = [[0] * 26 for _ in nums]
        or_left = []  # (子数组或值，最小左端点)

        for i, x in enumerate(nums):
            # 计算右端点为 i 的子数组或值
            for p in or_left:
                p[0] |= x
            # x 单独一个数作为子数组
            or_left.append([x, i])

            # 原地去重（相同或值只保留最左边的）
            idx = 1
            for j in range(1, len(or_left)):
                if or_left[j][0] != or_left[j - 1][0]:
                    or_left[idx] = or_left[j]
                    idx += 1
            del or_left[idx:]

            for k, (_, left) in enumerate(or_left):
                right = or_left[k + 1][1] - 1 if k < len(or_left) - 1 else i
                # 对于左端点在 [left, right]，右端点为 i 的子数组，OR 值都相同
                res[i][len(or_left) - 1 - k] = right - left + 1

        return res

    def countMagicalPairs(self, s: str) -> int:
        a = [1 << (ord(ch) - ord('a')) for ch in s]
        pre = self.log_trick(a)
        a.reverse()
        suf = self.log_trick(a)  # 相当于倒着遍历 a
        suf.reverse()

        ans = 0
        for i in range(len(pre) - 1):
            for p, s in zip(pre[i], suf[i + 1]):
                if p == 0 or s == 0:
                    break
                ans += p * s
        return ans
```

```java [sol-Java]
class Solution {
    public long countMagicalPairs(String S) {
        char[] s = S.toCharArray();
        int[][] pre = logTrick(s);

        int n = s.length;
        // 反转 s
        for (int i = 0; i < n / 2; i++) {
            char tmp = s[i];
            s[i] = s[n - 1 - i];
            s[n - 1 - i] = tmp;
        }
        int[][] suf = logTrick(s); // 相当于倒着遍历 s
        // 反转 suf
        for (int i = 0; i < n / 2; i++) {
            int[] tmp = suf[i];
            suf[i] = suf[n - 1 - i];
            suf[n - 1 - i] = tmp;
        }

        long ans = 0;
        for (int i = 0; i < n - 1; i++) {
            for (int j = 0; j < 26 && pre[i][j] > 0 && suf[i + 1][j] > 0; j++) {
                ans += (long) pre[i][j] * suf[i + 1][j];
            }
        }
        return ans;
    }

    // 模板来源：https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/
    private int[][] logTrick(char[] s) {
        int n = s.length;
        int[][] res = new int[n][26];
        List<int[]> orLeft = new ArrayList<>(); // (子串或值，最小左端点)

        for (int i = 0; i < n; i++) {
            int bit = 1 << (s[i] - 'a');
            // 计算右端点为 i 的子串或值
            for (int[] p : orLeft) {
                p[0] |= bit;
            }
            // bit 单独作为子串
            orLeft.add(new int[]{bit, i});

            // 原地去重（相同或值只保留最左边的）
            int m = 1;
            for (int j = 1; j < orLeft.size(); j++) {
                if (orLeft.get(j)[0] != orLeft.get(j - 1)[0]) {
                    orLeft.set(m++, orLeft.get(j));
                }
            }
            orLeft.subList(m, orLeft.size()).clear();

            for (int k = 0; k < m; k++) {
                int left = orLeft.get(k)[1];
                int right = k < m - 1 ? orLeft.get(k + 1)[1] - 1 : i;
                // 对于左端点在 [left, right]，右端点为 i 的子串，OR 值都相同
                res[i][m - 1 - k] = right - left + 1;
            }
        }

        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    // 模板来源：https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/
    vector<array<int, 26>> log_trick(const string& s) {
        int n = s.size();
        vector<array<int, 26>> res(n);
        vector<pair<int, int>> or_left; // (子串或值，最小左端点)

        for (int i = 0; i < s.size(); i++) {
            int bit = 1 << (s[i] - 'a');
            // 计算右端点为 i 的子串或值
            for (auto& [or_val, _] : or_left) {
                or_val |= bit;
            }
            // s[i] 单独作为子串
            or_left.emplace_back(bit, i);

            // 原地去重（相同或值只保留最左边的）
            int m = 1;
            for (int j = 1; j < or_left.size(); j++) {
                if (or_left[j].first != or_left[j - 1].first) {
                    or_left[m++] = or_left[j];
                }
            }
            or_left.resize(m);

            for (int k = 0; k < m; k++) {
                int left = or_left[k].second;
                int right = k + 1 < m ? or_left[k + 1].second - 1 : i;
                // 对于左端点在 [left, right]，右端点为 i 的子串，OR 值都相同
                res[i][m - 1 - k] = right - left + 1;
            }
        }

        return res;
    }

public:
    long long countMagicalPairs(string s) {
        auto pre = log_trick(s);
        ranges::reverse(s);
        auto suf = log_trick(s); // 相当于倒着遍历 s
        ranges::reverse(suf);

        int n = s.size();
        long long ans = 0;
        for (int i = 0; i < n - 1; i++) {
            for (int j = 0; j < 26 && pre[i][j] > 0 && suf[i + 1][j] > 0; j++) {
                ans += 1LL * pre[i][j] * suf[i + 1][j];
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
// 模板来源：https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/
func logTrick(s []byte) [][26]int {
	// 子串包含的字母集合，最小左端点
	type pair struct{ mask, left int }
	data := []pair{}
	res := make([][26]int, len(s))

	for i, ch := range s {
		bit := 1 << (ch - 'a')
		// 计算右端点为 i 的子串的 mask
		for j := range data {
			data[j].mask |= bit
		}
		// ch 单独作为子串
		data = append(data, pair{bit, i})

		// 原地去重（相同 mask 只保留最左边的）
		idx := 1
		for j := 1; j < len(data); j++ {
			if data[j].mask != data[j-1].mask {
				data[idx] = data[j]
				idx++
			}
		}
		data = data[:idx]

		for k, p := range data {
			right := i
			if k < len(data)-1 {
				right = data[k+1].left - 1
			}
			// 对于左端点在 [p.left, right]，右端点为 i 的子串，mask 都相同
			res[i][len(data)-1-k] = right - p.left + 1
		}
	}

	return res
}

func countMagicalPairs(S string) (ans int64) {
	s := []byte(S)
	pre := logTrick(s)
	slices.Reverse(s)
	suf := logTrick(s) // 相当于倒着遍历 s
	slices.Reverse(suf)

	for i := range len(pre) - 1 {
		for j := 0; j < 26 && pre[i][j] > 0 && suf[i+1][j] > 0; j++ {
			ans += int64(pre[i][j]) * int64(suf[i+1][j])
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(n|\Sigma|)$。

## 方法三：维护左右最近不同字母的位置

举个例子。假设 $s[7]$ 及其左侧最近的字母为 $\texttt{a},\texttt{b},\texttt{c}$，下标分别为 $1,5,7$，那么：

- 右端点为 $7$ 且恰好包含 $1$ 种字母的子串个数为 $7-5=2$。
- 右端点为 $7$ 且恰好包含 $2$ 种字母的子串个数为 $5-1=4$。
- 右端点为 $7$ 且恰好包含 $3$ 种字母的子串个数为 $1-(-1)=2$。

假设 $s[8]$ 及其右侧最近的字母为 $\texttt{d},\texttt{e}$，下标分别为 $8,10$，那么：

- 左端点为 $8$ 且恰好包含 $1$ 种字母的子串个数为 $10-8=2$。
- 左端点为 $8$ 且恰好包含 $2$ 种字母的子串个数为 $n-10$。

一般地，枚举分割线 $(i,i+1)$，同时维护 $i$ 及其左侧最近字母的位置，$i+1$ 及其右侧最近字母的位置，就可以用方法一的公式计算了。

```py [sol-Python3]
class Solution:
    def countMagicalPairs(self, s: str) -> int:
        n = len(s)
        # next_idx[i] 表示下一个 s[i] 的位置
        next_idx = [n] * n
        last = [n] * 26
        for i in range(n - 1, -1, -1):
            ch = ord(s[i]) - ord('a')
            next_idx[i] = last[ch]
            last[ch] = i

        # 右侧最近不同字母的位置
        right = sorted(i for i in last if i < n)
        right.append(n)  # 哨兵

        # 左侧最近不同字母的位置
        left = [-1]  # 哨兵
        last = [-1] * 26
        ans = 0

        # 枚举分割线 (i,i+1)
        for i in range(n - 1):
            ch = ord(s[i]) - ord('a')
            # 更新左侧最近不同字母的位置
            idx = last[ch]
            if idx >= 0:
                left.pop(bisect_left(left, idx))  # 删除旧的
            left.append(i)  # 添加新的
            last[ch] = i

            # 更新右侧最近不同字母的位置
            right.pop(0)  # 删除旧的
            idx = next_idx[i]
            if idx < n:
                insort(right, idx)  # 添加新的

            for k in range(min(len(left), len(right)) - 1):
                ans += (left[-1 - k] - left[-2 - k]) * (right[k + 1] - right[k])

        return ans
```

```java [sol-Java]
class Solution {
    public long countMagicalPairs(String S) {
        char[] s = S.toCharArray();
        int n = s.length;
        // nextIdx[i] 表示下一个 s[i] 的位置
        int[] nextIdx = new int[n];
        int[] last = new int[26];
        Arrays.fill(last, n);
        for (int i = n - 1; i >= 0; i--) {
            int ch = s[i] - 'a';
            nextIdx[i] = last[ch];
            last[ch] = i;
        }

        // 右侧最近不同字母的位置
        List<Integer> right = new ArrayList<>(); // 注：换成数组更快，见另一份代码【Java 数组】
        for (int ch = 0; ch < 26; ch++) {
            int p = last[ch];
            if (p < n) {
                right.add(p);
                last[ch] = -1;
            }
        }
        Collections.sort(right);
        right.add(n); // 哨兵

        // 左侧最近不同字母的位置
        List<Integer> left = new ArrayList<>(); // 注：换成数组更快，见另一份代码【Java 数组】
        left.add(-1); // 哨兵
        long ans = 0;

        // 枚举分割线 (i,i+1)
        for (int i = 0; i < n - 1; i++) {
            int ch = s[i] - 'a';
            // 更新左侧最近不同字母的位置
            int idx = last[ch];
            if (idx >= 0) {
                int j = Collections.binarySearch(left, idx);
                left.remove(j); // 删除旧的
            }
            left.add(i); // 添加新的
            last[ch] = i;

            // 更新右侧最近不同字母的位置
            right.remove(0); // 删除旧的
            idx = nextIdx[i];
            if (idx < n) {
                int j = ~Collections.binarySearch(right, idx);
                right.add(j, idx); // 添加新的
            }

            int nl = left.size();
            int nr = right.size();
            for (int k = 1; k < Math.min(nl, nr); k++) {
                ans += (long) (left.get(nl - k) - left.get(nl - k - 1)) * (right.get(k) - right.get(k - 1));
            }
        }

        return ans;
    }
}
```

```java [sol-Java 数组]
class Solution {
    public long countMagicalPairs(String S) {
        char[] s = S.toCharArray();
        int n = s.length;
        // nextIdx[i] 表示下一个 s[i] 的位置
        int[] nextIdx = new int[n];
        int[] last = new int[26];
        Arrays.fill(last, n);
        for (int i = n - 1; i >= 0; i--) {
            int ch = s[i] - 'a';
            nextIdx[i] = last[ch];
            last[ch] = i;
        }

        // 右侧最近不同字母的位置
        int[] right = new int[27];
        int nr = 0;
        for (int ch = 0; ch < 26; ch++) {
            int p = last[ch];
            if (p < n) {
                right[nr++] = p;
                last[ch] = -1;
            }
        }
        Arrays.sort(right, 0, nr);
        right[nr++] = n; // 哨兵

        // 左侧最近不同字母的位置
        int[] left = new int[27];
        int nl = 1;
        left[0] = -1; // 哨兵
        long ans = 0;

        // 枚举分割线 (i,i+1)
        for (int i = 0; i < n - 1; i++) {
            int ch = s[i] - 'a';
            // 更新左侧最近不同字母的位置
            int idx = last[ch];
            if (idx >= 0) {
                int j = Arrays.binarySearch(left, 0, nl, idx);
                nl--;
                System.arraycopy(left, j + 1, left, j, nl - j); // 删除旧的
            }
            left[nl++] = i; // 添加新的
            last[ch] = i;

            // 更新右侧最近不同字母的位置
            nr--;
            System.arraycopy(right, 1, right, 0, nr); // 删除旧的
            idx = nextIdx[i];
            if (idx < n) {
                int j = ~Arrays.binarySearch(right, 0, nr, idx);
                System.arraycopy(right, j, right, j + 1, nr - j);
                nr++;
                right[j] = idx; // 添加新的
            }

            for (int k = 1; k < Math.min(nl, nr); k++) {
                ans += (long) (left[nl - k] - left[nl - k - 1]) * (right[k] - right[k - 1]);
            }
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countMagicalPairs(string s) {
        int n = s.size();
        // next_idx[i] 表示下一个 s[i] 的位置
        vector<int> next_idx(n);
        int last[26];
        ranges::fill(last, n);
        for (int i = n - 1; i >= 0; i--) {
            int ch = s[i] - 'a';
            next_idx[i] = last[ch];
            last[ch] = i;
        }

        // 右侧最近不同字母的位置
        vector<int> right;
        for (int& i : last) {
            if (i < n) {
                right.push_back(i);
                i = -1;
            }
        }
        ranges::sort(right);
        right.push_back(n); // 哨兵

        // 左侧最近不同字母的位置
        vector<int> left = {-1}; // 哨兵
        long long ans = 0;

        // 枚举分割线 (i,i+1)
        for (int i = 0; i < n - 1; i++) {
            int ch = s[i] - 'a';
            // 更新左侧最近不同字母的位置
            int idx = last[ch];
            if (idx >= 0) {
                auto it = ranges::lower_bound(left, idx);
                left.erase(it); // 删除旧的
            }
            left.push_back(i); // 添加新的
            last[ch] = i;

            // 更新右侧最近不同字母的位置
            right.erase(right.begin()); // 删除旧的
            idx = next_idx[i];
            if (idx < n) {
                auto it = ranges::lower_bound(right, idx);
                right.insert(it, idx); // 添加新的
            }

            int nl = left.size(), nr = right.size();
            for (int k = 1; k < min(nl, nr); k++) {
                ans += 1LL * (left[nl - k] - left[nl - k - 1]) * (right[k] - right[k - 1]);
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
func countMagicalPairs(s string) (ans int64) {
	n := len(s)
	// nextIdx[i] 表示下一个 s[i] 的位置
	nextIdx := make([]int, n)
	last := [26]int{}
	for i := range last {
		last[i] = n
	}
	for i := n - 1; i >= 0; i-- {
		ch := s[i] - 'a'
		nextIdx[i] = last[ch]
		last[ch] = i
	}

	// 右侧最近不同字母的位置
	right := []int{}
	for ch, p := range last {
		if p < n {
			right = append(right, p)
			last[ch] = -1
		}
	}
	slices.Sort(right)
	right = append(right, n) // 哨兵

	// 左侧最近不同字母的位置
	left := []int{-1} // 哨兵

	// 枚举分割线 (i,i+1)
	for i := range n - 1 {
		ch := s[i] - 'a'
		// 更新左侧最近不同字母的位置
		idx := last[ch]
		if idx >= 0 {
			j := sort.SearchInts(left, idx)
			left = append(left[:j], left[j+1:]...) // 删除旧的
		}
		left = append(left, i) // 添加新的
		last[ch] = i

		// 更新右侧最近不同字母的位置
		right = right[1:] // 删除旧的
		idx = nextIdx[i]
		if idx < n {
			j := sort.SearchInts(right, idx)
			right = slices.Insert(right, j, idx) // 添加新的
		}

		nl := len(left)
		for k := range min(nl, len(right)) - 1 {
			ans += int64(left[nl-1-k]-left[nl-2-k]) * int64(right[k+1]-right[k])
		}
	}

	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(n + |\Sigma|)$。

## 专题训练

1. 滑动窗口题单的「**§2.3.3 恰好型滑动窗口**」。
2. 位运算题单的「**AND/OR LogTrick**」。
3. 动态规划题单的「**专题：前后缀分解**」。

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
