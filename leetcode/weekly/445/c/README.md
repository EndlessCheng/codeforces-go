## 初步分析

由于 $s$ 是回文串，我们只需关心左半部分如何排列。

特别地，如果 $s$ 的长度是奇数，那么正中间的那个字母 $\textit{c}$ 恰好出现奇数次，在重排后的回文串中，字母 $\textit{c}$ 也出现奇数次，所以正中间的字母也必须是 $\textit{c}$。

所以无论 $s$ 长度是奇数还是偶数，都只需关心左半部分（奇数去掉回文中心）如何排列。

## 思路

统计左半部分每个字母的出现次数。

然后用**试填法**构造答案：

- 最左边能不能是字母 $\texttt{a}$？如果不能，试试字母 $\texttt{b},\texttt{c},\ldots,\textit{z}$。
- 怎么判断能不能？假设最左边填字母 $\texttt{a}$，问题变成计算剩余位置的字符串的排列个数 $p$，如果 $p<k$，说明 $k$ 太大，继续尝试填字母 $\texttt{b}$；如果 $p\ge k$，说明右边足以容纳至少 $k$ 个排列，最左边就是字母 $\texttt{a}$。

如何计算字符串的排列个数？

#### 方法一（不适用）

设剩余长度为 $\textit{sz}$，如果随便排，有 $\textit{sz}!$ 种方案。

当然，这里面有重复的，例如 $\texttt{aabbb}$，其中两个 $\texttt{a}$ 和三个 $\texttt{b}$ 的排列就是重复的，由于这两个 $\texttt{a}$ **无法区分**，三个 $\texttt{b}$ **无法区分**，方案数要除以 $2!3!$。

排列个数为

$$
\dfrac{\textit{sz}!}{\prod\limits_{i=\texttt{a}}^{\texttt{z}}\textit{cnt}_i!}
$$

但问题是，这样算会溢出（或者要用高精度，但那样太慢）。我们需要找到一个算法，能够**在计算的过程中**判断方案数是否 $\ge k$，及时结束，避免溢出。

#### 方法二（中途退出循环）

假设字母 $\texttt{a}$ 有 $2$ 个，字母 $\texttt{b}$ 有 $3$ 个，其余字母个数略。

我们可以先从 $\textit{sz}$ 个位置中，选 $2$ 个位置填字母 $\texttt{a}$，方案数为 $\binom {\textit{sz}} 2$。然后再从剩余 $\textit{sz}-2$ 个位置中，选 $3$ 个位置填字母 $\texttt{b}$，方案数为 $\binom {\textit{sz}-2} 3$。以此类推。

根据乘法原理，排列数为这 $26$ 个组合数的乘积。

我们可以用普通的循环计算组合数，原理见 [62. 不同路径的题解](https://leetcode.cn/problems/unique-paths/solutions/3062432/liang-chong-fang-fa-dong-tai-gui-hua-zu-o5k32/) 的「六、另一种方法：组合数学」。

在循环的过程中，如果方案数 $\ge k$，就立刻退出循环。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1e3dBYLEDz/?t=6m56s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def smallestPalindrome(self, s: str, k: int) -> str:
        n = len(s)
        m = n // 2
        cnt = [0] * 26
        for b in s[:m]:
            cnt[ord(b) - ord('a')] += 1

        # 为什么这样做是对的？见 62. 不同路径 我的题解
        def comb(n: int, m: int) -> int:
            m = min(m, n - m)
            res = 1
            for i in range(1, m + 1):
                res = res * (n + 1 - i) // i
                if res >= k:  # 太大了
                    return k
            return res

        # 计算长度为 sz 的字符串的排列个数
        def perm(sz: int) -> int:
            res = 1
            for c in cnt:
                if c == 0:
                    continue
                # 先从 sz 个里面选 c 个位置填当前字母
                res *= comb(sz, c)
                if res >= k:  # 太大了
                    return k
                # 从剩余位置中选位置填下一个字母
                sz -= c
            return res

        # k 太大
        if perm(m) < k:
            return ""

        # 构造回文串的左半部分
        left_s = [''] * m
        for i in range(m):
            for j in range(26):
                if cnt[j] == 0:
                    continue
                cnt[j] -= 1  # 假设填字母 j，看是否有足够的排列
                p = perm(m - i - 1)  # 剩余位置的排列个数
                if p >= k:  # 有足够的排列
                    left_s[i] = ascii_lowercase[j]
                    break
                k -= p  # k 太大，要填更大的字母
                cnt[j] += 1

        ans = ''.join(left_s)
        if n % 2:
            ans += s[n // 2]
        return ans + ''.join(reversed(left_s))
```

```java [sol-Java]
class Solution {
    public String smallestPalindrome(String s, int k) {
        int n = s.length();
        int m = n / 2;
        int[] cnt = new int[26];
        for (int i = 0; i < m; i++) {
            cnt[s.charAt(i) - 'a']++;
        }

        if (perm(m, cnt, k) < k) { // k 太大
            return "";
        }

        char[] leftS = new char[m];
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < 26; j++) {
                if (cnt[j] == 0) {
                    continue;
                }
                cnt[j]--; // 假设填字母 j，看是否有足够的排列
                int p = perm(m - i - 1, cnt, k); // 剩余位置的排列个数
                if (p >= k) { // 有足够的排列
                    leftS[i] = (char) ('a' + j);
                    break;
                }
                k -= p; // k 太大，要填更大的字母
                cnt[j]++;
            }
        }

        StringBuilder ans = new StringBuilder(n);
        ans.append(leftS);
        if (n % 2 > 0) {
            ans.append(s.charAt(n / 2));
        }
        for (int i = m - 1; i >= 0; i--) {
            ans.append(leftS[i]);
        }
        return ans.toString();
    }

    // 为什么这样做是对的？见 62. 不同路径 我的题解
    private int comb(int n, int m, int k) {
        m = Math.min(m, n - m);
        long res = 1;
        for (int i = 1; i <= m; i++) {
            res = res * (n + 1 - i) / i;
            if (res >= k) { // 太大了
                return k;
            }
        }
        return (int) res;
    }

    // 计算长度为 sz 的字符串的排列个数
    private int perm(int sz, int[] cnt, int k) {
        long res = 1;
        for (int c : cnt) {
            if (c == 0) {
                continue;
            }
            res *= comb(sz, c, k); // 先从 sz 个里面选 c 个位置填当前字母
            if (res >= k) { // 太大了
                return k;
            }
            sz -= c; // 从剩余位置中选位置填下一个字母
        }
        return (int) res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string smallestPalindrome(string s, int k) {
        int n = s.size();
        int m = n / 2;
        int cnt[26]{};
        for (int i = 0; i < m; i++) {
            cnt[s[i] - 'a']++;
        }

        // 为什么这样做是对的？见 62. 不同路径 我的题解
        auto comb = [&](int n, int m) -> int {
            m = min(m, n - m);
            long long res = 1;
            for (int i = 1; i <= m; i++) {
                res = res * (n + 1 - i) / i;
                if (res >= k) { // 太大了
                    return k;
                }
            }
            return res;
        };

        // 计算长度为 sz 的字符串的排列个数
        auto perm = [&](int sz) -> int {
            long long res = 1;
            for (int c : cnt) {
                if (c == 0) {
                    continue;
                }
                // 先从 sz 个里面选 c 个位置填当前字母
                res *= comb(sz, c);
                if (res >= k) { // 太大了
                    return k;
                }
                // 从剩余位置中选位置填下一个字母
                sz -= c;
            }
            return res;
        };

        // k 太大
        if (perm(m) < k) {
            return "";
        }

        // 构造回文串的左半部分
        string left_s(m, 0);
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < 26; j++) {
                if (cnt[j] == 0) {
                    continue;
                }
                cnt[j]--; // 假设填字母 j，看是否有足够的排列
                int p = perm(m - i - 1); // 剩余位置的排列个数
                if (p >= k) { // 有足够的排列
                    left_s[i] = 'a' + j;
                    break;
                }
                k -= p; // k 太大，要填更大的字母
                cnt[j]++;
            }
        }

        string ans = left_s;
        if (n % 2) {
            ans += s[n / 2];
        }
        ranges::reverse(left_s);
        return ans + left_s;
    }
};
```

```go [sol-Go]
func smallestPalindrome(s string, k int) string {
	n := len(s)
	m := n / 2
	cnt := make([]int, 26)
	for _, b := range s[:m] {
		cnt[b-'a']++
	}

	// 为什么这样做是对的？见 62. 不同路径 我的题解
	comb := func(n, m int) int {
		m = min(m, n-m)
		res := 1
		for i := 1; i <= m; i++ {
			res = res * (n + 1 - i) / i
			if res >= k { // 太大了
				return k
			}
		}
		return res
	}

	// 计算长为 sz 的字符串的排列个数
	perm := func(sz int) int {
		res := 1
		for _, c := range cnt {
			if c == 0 {
				continue
			}
			// 先从 sz 个里面选 c 个位置填当前字母
			res *= comb(sz, c)
			if res >= k { // 太大了
				return k
			}
			// 从剩余位置中选位置填下一个字母
			sz -= c
		}
		return res
	}

	// k 太大
	if perm(m) < k {
		return ""
	}

	// 构造回文串的左半部分
	leftS := make([]byte, m)
	for i := range leftS {
		for j := range cnt {
			if cnt[j] == 0 {
				continue
			}
			cnt[j]-- // 假设填字母 j，看是否有足够的排列
			p := perm(m - i - 1) // 剩余位置的排列个数
			if p >= k { // 有足够的排列
				leftS[i] = 'a' + byte(j)
				break
			}
			k -= p // k 太大，要填更大的字母
			cnt[j]++
		}
	}

	ans := string(leftS)
	if n%2 > 0 {
		ans += string(s[n/2])
	}
	slices.Reverse(leftS)
	return ans + string(leftS)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n|\Sigma|(|\Sigma|+\log k))$，其中 $n$ 是 $\textit{nums}$ 的长度，$|\Sigma|=26$ 是字符集合的大小。计算 $\texttt{perm}$ 需要 $\mathcal{O}(|\Sigma| + \log k)$ 的时间。注意是加法不是乘法，因为内层循环的 $\textit{res}$ 至多乘以 $\mathcal{O}(\log k)$ 个大于 $1$ 的数。
- 空间复杂度：$\mathcal{O}(n + |\Sigma|)$。

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
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
