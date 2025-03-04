## 方法一：暴力修改

首先说做法。下文把 $\textit{str}_1$ 简记为 $s$，把 $\textit{str}_2$ 简记为 $t$。

先**模拟**：处理 $s$ 中的 T，把字符串 $t$ 填入答案的对应位置，如果发现矛盾，就返回空串。没填的位置（待定位置）初始化为 $\texttt{a}$。

再**贪心**：从左到右检查 F 对应的答案子串，如果发现子串和 $t$ 相同，那么把子串的最后一个待定位置改成 $\texttt{b}$。

本题的贪心策略是简单的，难点在正确性上。考虑如下问题：

- 按照上述贪心策略，是否存在一种情况，当我们把待定位置改成 $\texttt{b}$ 后，前面的某个 F 对应子串反而变成和 $t$ 相同了？

### 情况一

$t$ 全为 $\texttt{a}$ 的情况。

这是容易证明的，因为把待定位置改成 $\texttt{b}$ 后，前面的受到影响的子串（包含这个 $\texttt{b}$ 的子串）一定不会等于 $t$，毕竟 $t$ 只有 $\texttt{a}$。

例如 $t=\texttt{aaa}$，现在 $\textit{ans}=\texttt{aaa?????aaa}$。其中 $\texttt{?}$ 表示待定位置，初始值为 $\texttt{a}$。

- 我们遇到的第一个待定位置就会改成 $\texttt{b}$，后续所有包含这个 $\texttt{b}$ 的子串必然不等于 $t$，所以仍然为默认值 $\texttt{a}$。
- 直到我们遇到下一个需要改成 $\texttt{b}$ 的待定位置。
- 最终 $\textit{ans} = \texttt{aaa}\underline{\texttt{baabb}}\texttt{aaa}$。请动手算算，特别注意最后一个 $\texttt{b}$ 是怎么改的。

### 情况二

下面讨论 $t$ 包含不等于 $\texttt{a}$ 的字母的情况。

**猜想**：$t$ 形如 $t' + \texttt{aa\ldots a} + t'$。例如 $\texttt{baab},\texttt{baaaaba},\texttt{abaaaba}$ 等。

例如 $t=\texttt{baaaaba}$，即 $\texttt{ba} + \texttt{aaa} + \texttt{ba}$。

设 $\textit{ans} = \texttt{baaaaba???baaaaba}$。中间的 $\texttt{???}$ 不能全为 $\texttt{a}$，改成 $\texttt{aab}$，得 $\texttt{baaaaba}\underline{\texttt{aab}}\texttt{baaaaba}$，这里产生的 $\texttt{baaab}$ 可以保证前面的 F 对应子串不会和 $t$ 相同。

这可以推广到一般情况。抛砖引玉，欢迎在评论区发表你的证明。

此外还有一个**性质**：只需要改最后一个待定位置，不会出现改倒数第二个待定位置的情况。也留给读者证明。

**推论**：如果把 $\textit{ans}[j]$ 改成 $\texttt{b}$，那么所有包含 $\textit{ans}[j]$ 的子串都无需检查。这个推论可以用于方法二。

```py [sol-Python3]
class Solution:
    def generateString(self, s: str, t: str) -> str:
        n, m = len(s), len(t)
        ans = ['?'] * (n + m - 1)  # ? 表示待定位置
        for i, b in enumerate(s):
            if b != 'T':
                continue
            # 子串必须等于 t
            for j, c in enumerate(t):
                v = ans[i + j]
                if v != '?' and v != c:
                    return ""
                ans[i + j] = c

        old_ans = ans
        ans = ['a' if c == '?' else c for c in ans]  # 待定位置的初始值为 a

        for i, b in enumerate(s):
            if b != 'F':
                continue
            # 子串必须不等于 t
            if ''.join(ans[i: i + m]) != t:
                continue
            # 找最后一个待定位置
            for j in range(i + m - 1, i - 1, -1):
                if old_ans[j] == '?':  # 之前填 a，现在改成 b
                    ans[j] = 'b'
                    break
            else:
                return ""

        return ''.join(ans)
```

```java [sol-Java]
class Solution {
    public String generateString(String S, String t) {
        char[] s = S.toCharArray();
        int n = s.length;
        int m = t.length();
        char[] ans = new char[n + m - 1];
        Arrays.fill(ans, '?'); // '?' 表示待定位置
        for (int i = 0; i < n; i++) {
            if (s[i] != 'T') {
                continue;
            }
            // 子串必须等于 t
            for (int j = 0; j < m; j++) {
                char v = ans[i + j];
                if (v != '?' && v != t.charAt(j)) {
                    return "";
                }
                ans[i + j] = t.charAt(j);
            }
        }

        char[] oldAns = ans.clone();
        for (int i = 0; i < ans.length; i++) {
            if (ans[i] == '?') {
                ans[i] = 'a'; // 待定位置的初始值为 'a'
            }
        }

        for (int i = 0; i < n; i++) {
            if (s[i] != 'F') {
                continue;
            }
            // 子串必须不等于 t
            if (!new String(ans, i, m).equals(t)) {
                continue;
            }
            // 找最后一个待定位置
            boolean ok = false;
            for (int j = i + m - 1; j >= i; j--) {
                if (oldAns[j] == '?') { // 之前填 'a'，现在改成 'b'
                    ans[j] = 'b';
                    ok = true;
                    break;
                }
            }
            if (!ok) {
                return "";
            }
        }

        return new String(ans);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string generateString(string s, string t) {
        int n = s.size(), m = t.size();
        string ans(n + m - 1, '?'); // ? 表示待定位置
        for (int i = 0; i < n; i++) {
            if (s[i] != 'T') {
                continue;
            }
            // 子串必须等于 t
            for (int j = 0; j < m; j++) {
                char v = ans[i + j];
                if (v != '?' && v != t[j]) {
                    return "";
                }
                ans[i + j] = t[j];
            }
        }

        string old_ans = ans;
        for (char& c : ans) {
            if (c == '?') {
                c = 'a'; // 待定位置的初始值为 a
            }
        }

        for (int i = 0; i < n; i++) {
            if (s[i] != 'F') {
                continue;
            }
            // 子串必须不等于 t
            if (string(ans.begin() + i, ans.begin() + i + m) != t) {
                continue;
            }
            // 找最后一个待定位置
            bool ok = false;
            for (int j = i + m - 1; j >= i; j--) {
                if (old_ans[j] == '?') { // 之前填 a，现在改成 b
                    ans[j] = 'b';
                    ok = true;
                    break;
                }
            }
            if (!ok) {
                return "";
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
func generateString(s, T string) string {
	n, m := len(s), len(T)
	t := []byte(T)
	ans := bytes.Repeat([]byte{'?'}, n+m-1) // ? 表示待定位置
	for i, b := range s {
		if b != 'T' {
			continue
		}
		// sub 必须等于 t
		sub := ans[i : i+m]
		for j, c := range sub {
			if c != '?' && c != t[j] {
				return ""
			}
			sub[j] = t[j]
		}
	}
	oldAns := ans
	ans = bytes.ReplaceAll(ans, []byte{'?'}, []byte{'a'}) // 待定位置的初始值为 a

next:
	for i, b := range s {
		if b != 'F' {
			continue
		}
		// sub 必须不等于 t 
		sub := ans[i : i+m]
		if !bytes.Equal(sub, t) {
			continue
		}
		// 找最后一个待定位置
		old := oldAns[i : i+m]
		for j := m - 1; j >= 0; j-- {
			if old[j] == '?' { // 之前填 a，现在改成 b
				sub[j] = 'b'
				continue next
			}
		}
		return ""
	}

	return string(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm)$，其中 $n$ 是 $s$ 的长度，$m$ 是 $t$ 的长度。
- 空间复杂度：$\mathcal{O}(n+m)$。如果不考虑切片和返回值的话是 $\mathcal{O}(1)$。

## 方法二：Z 函数

在模拟（处理 $s$ 中的 T）的过程中，如果两个 $t$ 重叠，我们需要判断 $t$ 的某个长度的前后缀是否相同，这可以用 Z 函数直接解决。

判断 $\textit{ans}$ 子串是否等于 $t$ 也可以用 Z 函数。计算 $t + \textit{ans}$ 的 Z 函数，如果 $z[i+m]<m$，就说明从 $i$ 开始的 $\textit{ans}$ 子串不等于 $t$。

如果子串等于 $t$，那么找一个小于 $i+m$ 的最近待定位置，改成 $\texttt{b}$。这可以用一个数组 $\textit{preQ}$ 预处理每个 $\le i$ 的最近待定位置。

```py [sol-Python3]
class Solution:
    def calc_z(self, s: str) -> List[int]:
        n = len(s)
        z = [0] * n
        box_l, box_r = 0, 0  # z-box 左右边界（闭区间）
        for i in range(1, n):
            if i <= box_r:
                z[i] = min(z[i - box_l], box_r - i + 1)
            while i + z[i] < n and s[z[i]] == s[i + z[i]]:
                box_l, box_r = i, i + z[i]
                z[i] += 1
        z[0] = n
        return z

    def generateString(self, s: str, t: str) -> str:
        z = self.calc_z(t)
        n, m = len(s), len(t)
        ans = ['?'] * (n + m - 1)
        pre = -m
        for i, b in enumerate(s):
            if b != 'T':
                continue
            size = max(pre + m - i, 0)
            # t 的长为 size 的前后缀必须相同
            if size > 0 and z[m - size] < size:
                return ""
            # size 后的内容都是 '?'，填入 t
            ans[i + size: i + m] = t[size:]
            pre = i

        # 计算 <= i 的最近待定位置
        pre_q = [-1] * len(ans)
        pre = -1
        for i, c in enumerate(ans):
            if c == '?':
                ans[i] = 'a'  # 待定位置的初始值为 a
                pre = i
            pre_q[i] = pre

        # 找 ans 中的等于 t 的位置，可以用 KMP 或者 Z 函数
        z = self.calc_z(t + "".join(ans))
        i = 0
        while i < n:
            if s[i] != 'F':
                i += 1
                continue
            # 子串必须不等于 t
            if z[m + i] < m:
                i += 1
                continue
            # 找最后一个待定位置
            j = pre_q[i + m - 1]
            if j < i:  # 没有
                return ""
            ans[j] = 'b'
            i = j + 1  # 直接跳过 j

        return ''.join(ans)
```

```java [sol-Java]
class Solution {
    public String generateString(String S, String t) {
        // 生成符合条件的字符串
        int[] z = calcZ(t);
        char[] s = S.toCharArray();
        int n = s.length;
        int m = t.length();
        char[] ans = new char[n + m - 1];
        Arrays.fill(ans, '?');
        int pre = -m;
        for (int i = 0; i < n; i++) {
            if (s[i] != 'T') {
                continue;
            }
            int size = Math.max(pre + m - i, 0);
            // t 的长为 size 的前后缀必须相同
            if (size > 0 && z[m - size] < size) {
                return "";
            }
            // size 后的内容都是 '?'，填入 t
            for (int j = size; j < m; j++) {
                ans[i + j] = t.charAt(j);
            }
            pre = i;
        }

        // 计算 <= i 的最近待定位置
        int[] preQ = new int[ans.length];
        pre = -1;
        for (int i = 0; i < ans.length; i++) {
            if (ans[i] == '?') {
                ans[i] = 'a'; // 待定位置的初始值为 a
                pre = i;
            }
            preQ[i] = pre;
        }

        // 找 ans 中的等于 t 的位置，可以用 KMP 或者 Z 函数
        z = calcZ(t + new String(ans));
        for (int i = 0; i < n; i++) {
            if (s[i] != 'F') {
                continue;
            }
            // 子串必须不等于 t
            if (z[m + i] < m) {
                continue;
            }
            // 找最后一个待定位置
            int j = preQ[i + m - 1];
            if (j < i) { // 没有
                return "";
            }
            ans[j] = 'b';
            i = j; // 直接跳到 j
        }

        return new String(ans);
    }

    private int[] calcZ(String S) {
        char[] s = S.toCharArray();
        int n = s.length;
        int[] z = new int[n];
        int boxL = 0; // z-box 左右边界（闭区间）
        int boxR = 0;
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
        z[0] = n;
        return z;
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<int> calc_z(const string& s) {
        int n = s.size();
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
        z[0] = n;
        return z;
    }

public:
    string generateString(string s, string t) {
        // 生成符合条件的字符串
        vector<int> z = calc_z(t);
        int n = s.size(), m = t.size();
        string ans(n + m - 1, '?');
        int pre = -m;
        for (int i = 0; i < n; i++) {
            if (s[i] != 'T') {
                continue;
            }
            int size = max(pre + m - i, 0);
            // t 的长为 size 的前后缀必须相同
            if (size > 0 && z[m - size] < size) {
                return "";
            }
            // size 后的内容都是 '?'，填入 t
            for (int j = size; j < m; j++) {
                ans[i + j] = t[j];
            }
            pre = i;
        }

        // 计算 <= i 的最近待定位置
        vector<int> pre_q(ans.size());
        pre = -1;
        for (int i = 0; i < ans.size(); i++) {
            if (ans[i] == '?') {
                ans[i] = 'a'; // 待定位置的初始值为 a
                pre = i;
            }
            pre_q[i] = pre;
        }

        // 找 ans 中的等于 t 的位置，可以用 KMP 或者 Z 函数
        z = calc_z(t + ans);
        for (int i = 0; i < n; i++) {
            if (s[i] != 'F') {
                continue;
            }
            // 子串必须不等于 t
            if (z[m + i] < m) {
                continue;
            }
            // 找最后一个待定位置
            int j = pre_q[i + m - 1];
            if (j < i) { // 没有
                return "";
            }
            ans[j] = 'b';
            i = j; // 直接跳到 j
        }

        return ans;
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
	z[0] = n
	return z
}

func generateString(s, t string) string {
	z := calcZ(t)
	n, m := len(s), len(t)
	ans := bytes.Repeat([]byte{'?'}, n+m-1)
	pre := -m
	for i, b := range s {
		if b != 'T' {
			continue
		}
		size := max(pre+m-i, 0)
		// t 的长为 size 的前后缀必须相同
		if size > 0 && z[m-size] < size {
			return ""
		}
		// size 后的内容都是 '?'，填入 t
		copy(ans[i+size:], t[size:])
		pre = i
	}

	// 计算 <= i 的最近待定位置
	preQ := make([]int, len(ans))
	pre = -1
	for i, c := range ans {
		if c == '?' {
			ans[i] = 'a' // 待定位置的初始值为 a
			pre = i
		}
		preQ[i] = pre
	}

	// 找 ans 中的等于 t 的位置，可以用 KMP 或者 Z 函数
	z = calcZ(t + string(ans))
	for i := 0; i < n; i++ {
		if s[i] != 'F' {
			continue
		}
		// 子串必须不等于 t 
		if z[m+i] < m {
			continue
		}
		// 找最后一个待定位置
		j := preQ[i+m-1]
		if j < i { // 没有
			return ""
		}
		ans[j] = 'b'
		i = j // 直接跳到 j
	}

	return string(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $n$ 是 $s$ 的长度，$m$ 是 $t$ 的长度。
- 空间复杂度：$\mathcal{O}(n+m)$。

更多相似题目，见下面贪心题单中的「**§3.1 字典序最小/最大**」和字符串题单中的「**二、Z 函数**」。

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
10. 【本题相关】[贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. 【本题相关】[字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
