## 方法一：暴力修改

首先说做法。下文把 $\textit{str}_1$ 简记为 $s$，把 $\textit{str}_2$ 简记为 $t$。

先**模拟**：处理 $s$ 中的 $\texttt{T}$，把字符串 $t$ 填入答案的对应位置，如果发现矛盾，就返回空串。没填的位置（待定位置）初始化为 $\texttt{a}$。

再**贪心**：从左到右检查 $\texttt{F}$ 对应的答案子串，如果发现子串和 $t$ 相同，那么把子串的最后一个待定位置改成 $\texttt{b}$。

本题的贪心策略是简单的，难点在正确性上。考虑如下问题：

- 按照上述贪心策略，是否存在一种情况，当我们把待定位置改成 $\texttt{b}$ 后，前面的某个 $\texttt{F}$ 对应子串反而变成和 $t$ 相同了？

如下图，这只会发生在 $t = A + \texttt{a} + B + \texttt{b} + C$ 的情况。下图的 $\texttt{?}$ 原来是 $\texttt{a}$，为了让从 $i$ 开始的子串不等于 $t$，我们把 $\texttt{a}$ 改成 $\texttt{b}$，但这可能会导致之前的某个子串变成 $t$。不能修复了一个 bug，又产生了一个新的 bug 呀！

![lc3474-c.png](https://pic.leetcode.cn/1774925078-nOtrAY-lc3474-c.png)

具体地，定义：

- $t = A + \texttt{a} + B + \texttt{b} + C$。
- $s[i] = s[k] = \texttt{F}$。
- 修改前，下标 $j$ 处的字母是 $\texttt{a}$，从 $i$ 开始的长为 $m$ 的子串等于 $t$。
- $j$ 是在处理 $s[i]$ 时，我们倒着遍历找到的第一个待定位置。

在处理 $s[i]$ 时，由于从 $i$ 开始的长为 $m$ 的子串等于 $t$，要把 $j$ 处的字母改成 $\texttt{b}$。我们担心的是，修改后，是否会把更早的一个 $s[k] = \texttt{F}$ 对应的子串给改成 $t$ 了？

这是不会的，可以用**反证法**证明，也就是假设出现了上图的情况。

$j$ 是在处理 $s[i]$ 时，我们倒着遍历找到的第一个待定位置，这意味着什么？

这意味着从 $j+1$ 往右的 $B + \texttt{b} + C$，是我们在一开始处理 $\texttt{T}$ 时，已经确定的那些字母。所以 $s[j+1]$ 一定是 $\texttt{T}$，所以 $B + \texttt{b} + C$ 是 $t$ 的前缀。**这是最重要的结论，后面会反复用到这个结论**。

下面用记号 $|s|$ 表示字符串 $s$ 的长度。

分类讨论。

**情况一**：$|A| = |B|$。

$t = A + \texttt{a} + B + \texttt{b} + C$ 去掉前 $|A|$ 个字母后，剩下的第一个字母是 $\texttt{a}$。

$B + \texttt{b} + C$ 去掉前 $|B|$ 个字母后，剩下的第一个字母是 $\texttt{b}$。

由于 $B + \texttt{b} + C$ 是 $t = A + \texttt{a} + B + \texttt{b} + C$ 的前缀，在 $|A| = |B|$ 的情况下，$t$ 去掉前 $|A|$ 个字母后，剩下的第一个字母既是 $\texttt{a}$，又是 $\texttt{b}$，矛盾。

**情况二**：$|A| > |B|$。

由于 $B + \texttt{b} + C$ 是 $t = A + \texttt{a} + B + \texttt{b} + C$ 的前缀，在 $|A| > |B|$ 的情况下，$B + \texttt{b}$ 是 $A$ 的前缀，所以 $A$ 可以表示为 $A = B + \texttt{b} + D$。

此外，根据上文的图，$A$ 是 $A + \texttt{a} + B$ 的后缀。把这句话中的 $A$ 用 $B + \texttt{b} + D$ 替换，得到 $B + \texttt{b} + D$ 是 $B + \texttt{b} + D + \texttt{a} + B$ 的后缀。由于 $B + \texttt{b} + D + \texttt{a} + B$ 的长为 $|B|+1+|D|$ 的后缀是 $D + \texttt{a} + B$，所以 $B + \texttt{b} + D = D + \texttt{a} + B$。对比这两个字符串的每个字母的出现次数，发现左边多了一个 $\texttt{b}$，右边多了一个 $\texttt{a}$，所以这两个字符串不可能相等，矛盾。 

**情况三**：$|A| < |B|$。

根据上文的图，$A$ 是 $A + \texttt{a} + B$ 的后缀。在 $|A| < |B|$ 的情况下，$A$ 也是 $B$ 的后缀，所以 $B$ 可以表示为 $B = D + A$。

$B + \texttt{b} + C$ 是 $t = A + \texttt{a} + B + \texttt{b} + C$ 的前缀。把这句话中的 $B$ 用 $D + A$ 替换，得到 $D + A + \texttt{b} + C$ 是 $A + \texttt{a} + D + A + \texttt{b} + C$ 的前缀。由于 $D + A + \texttt{b} + C$ 和 $A + \texttt{a} + D + A + \texttt{b} + C$ 的长为 $|A|+|D| + 1$ 的前缀分别是 $D + A + \texttt{b}$ 和 $A + \texttt{a} + D$，所以 $D + A + \texttt{b} = A + \texttt{a} + D$。对比这两个字符串的每个字母的出现次数，发现左边多了一个 $\texttt{b}$，右边多了一个 $\texttt{a}$，所以这两个字符串不可能相等，矛盾。

综上所述，原命题成立，我们不会把更早的一个 $s[k] = \texttt{F}$ 对应的子串给改成 $t$。

```py [sol-Python3]
class Solution:
    def generateString(self, s: str, t: str) -> str:
        n, m = len(s), len(t)
        ans = ['?'] * (n + m - 1)  # ? 表示待定位置

        # 处理 T
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

        # 处理 F
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

        // 处理 T
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

        // 处理 F
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

        // 处理 T
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

        // 处理 F
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
    
    // 处理 T
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

    // 处理 F
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
        n, m = len(s), len(t)
        ans = ['?'] * (n + m - 1)

        # 处理 T
        z = self.calc_z(t)
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
        z = self.calc_z(t + ''.join(ans))

        # 处理 F
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
        char[] s = S.toCharArray();
        int n = s.length;
        int m = t.length();
        char[] ans = new char[n + m - 1];
        Arrays.fill(ans, '?');

        // 处理 T
        int[] z = calcZ(t);
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

        // 处理 F
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
        int n = s.size(), m = t.size();
        string ans(n + m - 1, '?');

        // 处理 T
        vector<int> z = calc_z(t);
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

        // 处理 F
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
    n, m := len(s), len(t)
    ans := bytes.Repeat([]byte{'?'}, n+m-1)

    // 处理 T
    pre := -m
    z := calcZ(t)
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

    // 处理 F
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
