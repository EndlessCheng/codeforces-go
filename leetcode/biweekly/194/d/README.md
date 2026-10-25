三变量问题，枚举中间的 $j$。

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

那么答案为

$$
\sum_{i=0}^{n-2}\sum_{j=1}^{26} \textit{pre}[i][j] \cdot \textit{suf}[i+1][j]
$$

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

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
