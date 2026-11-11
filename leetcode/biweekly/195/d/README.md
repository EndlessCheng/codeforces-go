设 $A = (a * x),\ B=(b * y)$。注：比 $(a * x)$ 更常用的记号是 $a^x$。

暴力匹配的想法是，枚举 $B$ 在 $A$ 中的匹配起点 $\textit{start}=0,1,2,\ldots,|a|-1$，判断 $B$ 是否等于 $A$ 的从 $\textit{start}$ 开始的长为 $|B|$ 的子串。

但 $|B|$ 太大了，这样做会超时。

## 方法一：比较 lcm(|a|,|b|) 个字母

例如 $|a|=4,\ |b|=6,\ \textit{start}=0$，我们从 $i=0$ 开始，判断 $a[i\bmod 4] = b[i\bmod 6]$ 是否成立。如果前 $12$ 个字母都匹配，那么继续匹配相当于又从 $a[12\bmod 4]=a[0]$ 和 $b[12\bmod 6]=b[0]$ 开始匹配了，这两个字母（以及后续字母）我们已经知道它们是匹配的。所以**只要前 $12$ 个字母都匹配，那么后面所有字母都匹配**。我们需要找到一个数，即是 $4$ 的倍数，又是 $6$ 的倍数，即公倍数。$12$ 是 $4$ 和 $6$ 的**最小公倍数**（LCM）。

一般地，只需比较 $A[\textit{start}:]$ 和 $B$ 的前

$$
\text{lcm}(|a|,|b|)
$$

个字母是否匹配。

代码实现时：

1. 如果 $B$ 比 $A$ 还长，那么 $B$ 一定不是 $A$ 的子串。
2. 枚举 $\textit{start}$ 还需要保证 $|A[\textit{start}:]| \ge |B|$，即 $|A| - \textit{start}\ge |B|$，即 $\textit{start} \le |A| - |B|$。

[本题视频讲解](https://www.bilibili.com/video/BV152W76MEX3/?t=7m55s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def repeatedStringMatch(self, a: str, b: str, x: int, y: int) -> bool:
        n, m = len(a), len(b)
        len_a, len_b = n * x, m * y
        if len_b > len_a:
            return False

        firsts = min(lcm(n, m), len_b)
        for start in range(min(n, len_a - len_b + 1)):
            # 判断 B 是否为 A[start:] 的前缀
            # 只需比较前 firsts 个字母
            if all(a[(start + i) % n] == b[i % m] for i in range(firsts)):
                return True
        return False
```

```java [sol-Java]
class Solution {
    public boolean repeatedStringMatch(String a, String b, int x, int y) {
        char[] s = a.toCharArray();
        char[] t = b.toCharArray();
        int n = s.length;
        int m = t.length;
        long lenA = (long) n * x;
        long lenB = (long) m * y;
        if (lenB > lenA) {
            return false;
        }

        int firsts = (int) Math.min(lcm(n, m), lenB);
        next:
        for (int start = 0; start < Math.min(n, lenA - lenB + 1); start++) {
            // 判断 B 是否为 A[start,lenA) 的前缀
            for (int i = 0; i < firsts; i++) { // 只需比较前 firsts 个字母
                if (s[(start + i) % n] != t[i % m]) {
                    continue next;
                }
            }
            return true;
        }
        return false;
    }

    private int lcm(int a, int b) {
        return a / gcd(a, b) * b;
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool repeatedStringMatch(string a, string b, int x, int y) {
        int n = a.size(), m = b.size();
        long long len_a = 1LL * n * x, len_b = 1LL * m * y;
        if (len_b > len_a) {
            return false;
        }

        int firsts = min(1LL * lcm(n, m), len_b);
        int mx = min(1LL * n, len_a - len_b + 1);
        for (int start = 0; start < mx; start++) {
            // 判断 B 是否为 A[start,len_a) 的前缀
            bool ok = true;
            for (int i = 0; i < firsts; i++) { // 只需比较前 firsts 个字母
                if (a[(start + i) % n] != b[i % m]) {
                    ok = false;
                    break;
                }
            }
            if (ok) {
                return true;
            }
        }
        return false;
    }
};
```

```go [sol-Go]
func repeatedStringMatch(a, b string, x, y int) bool {
	n, m := len(a), len(b)
	lenA, lenB := n*x, m*y
	if lenB > lenA {
		return false
	}

	firsts := min(lcm(n, m), lenB)
next:
	for start := range min(n, lenA-lenB+1) {
		// 判断 B 是否为 A[start:] 的前缀
		for i := range firsts { // 只需比较前 firsts 个字母
			if a[(start+i)%n] != b[i%m] {
				continue next
			}
		}
		return true
	}
	return false
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2m)$，其中 $n$ 是 $a$ 的长度，$m$ 是 $\textit{b}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：比较 |a| + |b| 个字母

设 $n = |a|,\ m = |b|$。

对于 $A$ 和 $B$，如果它们的前 $n+m$ 个字母相同，我们还需要比较后面的字母吗？

由于 $B$ 的周期为 $m$，所以有

$$
B_i = B_{i+m}
$$

由于前 $n+m$ 个字母相同，所以当 $0\le i < n$ 时，有 $B_i = A_i$ 且 $B_{i+m} = A_{i+m}$，替换上式，得

$$
A_i = A_{i+m}
$$

由于 $A$ 的周期为 $n$，所以有

$$
A_i = A_{(i+m)\bmod n}
$$

想象一个长为 $n$ 的环，从环长任意一点出发，走 $m$ 步，遇到的字母不变。所以再走 $m$ 步，遇到的字母仍然不变。

所以上式也可以写成

$$
A_i = A_{(i+pm)\bmod n}
$$

根据**贝祖定理**，存在整数 $p$ 和 $q$ 使得 $pm-qn=\gcd(n,m)=g$，代入上式，得

$$
A_i = A_{(i+qn+g)\bmod n} = A_{(i+g)\bmod n}
$$

所以 $g=\gcd(n,m)$ 也是 $A$ 的周期。

$A$ 和 $B$ 的前 $m$ 个字母相同，前 $m$ 个字母的周期是 $g$，$B$ 又由前 $m$ 个字母重复得到，所以 $B$ 由前 $g$ 个字母重复得到，所以 $g$ 也是 $B$ 的周期。

综上所述，我们证明了如下结论：

- 如果 $A$ 和 $B$ 的前 $n+m$ 个字母相同，那么 $A$ 和 $B$ 的周期都是 $g = \gcd(n,m)$。
- 由于 $n+m = k_1g+k_2g = (k_1+k_2)g$ 是 $g$ 的倍数，所以继续比较相当于又从 $A[(n+m)\bmod g]=A[0]$ 和 $B[(n+m)\bmod g]=B[0]$ 开始了，这两个字母（以及后续字母）我们已经知道它们是匹配的，所以无需比较后续字母。

对于 $A[\textit{start}:]$ 和 $B$，由于 $A[\textit{start}:]$ 的周期也是 $n$，所以结论和上面是一样的。例如 $A = \texttt{abcabcabc}\cdots$ 是由 $\texttt{abc}$ 重复得到，那么 $A[1:] = \texttt{bcabcabca}\cdots$ 是由 $\texttt{bca}$ 重复得到，仍然是周期字符串且周期不变。

所以只需比前 $n+m$ 个字母。

### 写法一

```py [sol-Python3]
class Solution:
    def repeatedStringMatch(self, a: str, b: str, x: int, y: int) -> bool:
        n, m = len(a), len(b)
        len_a, len_b = n * x, m * y
        if len_b > len_a:
            return False

        firsts = min(n + m, len_b)
        for start in range(min(n, len_a - len_b + 1)):
            # 判断 B 是否为 A[start:] 的前缀
            # 只需比较前 firsts 个字母
            if all(a[(start + i) % n] == b[i % m] for i in range(firsts)):
                return True
        return False
```

```java [sol-Java]
class Solution {
    public boolean repeatedStringMatch(String a, String b, int x, int y) {
        char[] s = a.toCharArray();
        char[] t = b.toCharArray();
        int n = s.length;
        int m = t.length;
        long lenA = (long) n * x;
        long lenB = (long) m * y;
        if (lenB > lenA) {
            return false;
        }

        int firsts = (int) Math.min(n + m, lenB);
        next:
        for (int start = 0; start < Math.min(n, lenA - lenB + 1); start++) {
            // 判断 B 是否为 A[start,lenA) 的前缀
            for (int i = 0; i < firsts; i++) { // 只需比较前 firsts 个字母
                if (s[(start + i) % n] != t[i % m]) {
                    continue next;
                }
            }
            return true;
        }
        return false;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool repeatedStringMatch(string a, string b, int x, int y) {
        int n = a.size(), m = b.size();
        long long len_a = 1LL * n * x, len_b = 1LL * m * y;
        if (len_b > len_a) {
            return false;
        }

        int firsts = min(1LL * (n + m), len_b);
        int mx = min(1LL * n, len_a - len_b + 1);
        for (int start = 0; start < mx; start++) {
            // 判断 B 是否为 A[start,len_a) 的前缀
            bool ok = true;
            for (int i = 0; i < firsts; i++) { // 只需比较前 firsts 个字母
                if (a[(start + i) % n] != b[i % m]) {
                    ok = false;
                    break;
                }
            }
            if (ok) {
                return true;
            }
        }
        return false;
    }
};
```

```go [sol-Go]
func repeatedStringMatch(a, b string, x, y int) bool {
	n, m := len(a), len(b)
	lenA, lenB := n*x, m*y
	if lenB > lenA {
		return false
	}

	firsts := min(n+m, lenB)
next:
	for start := range min(n, lenA-lenB+1) {
		// 判断 B 是否为 A[start:] 的前缀
		for i := range firsts { // 只需比较前 firsts 个字母
			if a[(start+i)%n] != b[i%m] {
				continue next
			}
		}
		return true
	}
	return false
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n(n+m))$，其中 $n$ 是 $a$ 的长度，$m$ 是 $\textit{b}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

### 写法二

既然 $B$ 只需比较前 $\textit{firsts} = \min(n+m,|B|)$ 个字母，我们可以把 $b$ 重复 $\left\lceil\dfrac{\textit{firsts}}{m}\right\rceil$ 次，得到字符串 $B'$（多比一点也没关系，前 $\textit{firsts}$ 一样，后面也一样）。

$A$ 中的起点 $ < \min(|a|, |A|-|B|+1)$，所以终点 $< \min(|a|, |A|-|B|+1) + \textit{firsts} - 1$，记作 $k$，我们可以把 $a$ 重复 $\left\lceil\dfrac{k}{|a|}\right\rceil$ 次，得到字符串 $A'$。

问题变成判断 $B'$ 是否在 $A'$ 中。

关于上取整的计算，当 $p$ 和 $q$ 均为正整数时，我们有

$$
\left\lceil\dfrac{p}{q}\right\rceil = \left\lfloor\dfrac{p-1}{q}\right\rfloor + 1
$$

证明见 [上取整下取整转换公式的证明](https://zhuanlan.zhihu.com/p/1890356682149838951)。

```py [sol-Python3]
class Solution:
    def repeatedStringMatch(self, a: str, b: str, x: int, y: int) -> bool:
        n, m = len(a), len(b)
        len_a, len_b = n * x, m * y
        if len_b > len_a:
            return False

        mx = min(n, len_a - len_b + 1)
        firsts = min(n + m, len_b)
        A = a * ((mx + firsts - 2) // n + 1)
        B = b * ((firsts - 1) // m + 1)
        return B in A
```

```java [sol-Java]
class Solution {
    public boolean repeatedStringMatch(String a, String b, int x, int y) {
        int n = a.length();
        int m = b.length();
        long lenA = (long) n * x;
        long lenB = (long) m * y;
        if (lenB > lenA) {
            return false;
        }

        int mx = (int) Math.min(n, lenA - lenB + 1);
        int firsts = (int) Math.min(n + m, lenB);
        String A = a.repeat((mx + firsts - 2) / n + 1);
        String B = b.repeat((firsts - 1) / m + 1);
        return A.contains(B);
    }
}
```

```cpp [sol-C++]
class Solution {
    string repeat(const string& s, int k) {
        string res;
        while (k--) {
            res += s;
        }
        return res;
    }

public:
    bool repeatedStringMatch(string a, string b, int x, int y) {
        int n = a.size(), m = b.size();
        long long len_a = 1LL * n * x, len_b = 1LL * m * y;
        if (len_b > len_a) {
            return false;
        }

        int mx = min(1LL * n, len_a - len_b + 1);
        int firsts = min(1LL * (n + m), len_b);
        string A = repeat(a, (mx + firsts - 2) / n + 1);
        string B = repeat(b, (firsts - 1) / m + 1);
        return A.contains(B);
    }
};
```

```go [sol-Go]
func repeatedStringMatch(a, b string, x, y int) bool {
	n, m := len(a), len(b)
	lenA, lenB := n*x, m*y
	if lenB > lenA {
		return false
	}

	mx := min(n, lenA-lenB+1)
	firsts := min(n+m, lenB)
	A := strings.Repeat(a, (mx+firsts-2)/n+1)
	B := strings.Repeat(b, (firsts-1)/m+1)
	return strings.Contains(A, B)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n(n+m))$，其中 $n$ 是 $a$ 的长度，$m$ 是 $\textit{b}$ 的长度。
- 空间复杂度：$\mathcal{O}(n+m)$。

### 优化

用 [KMP 算法](https://www.zhihu.com/question/21923021/answer/37475572) 加速判断 $B'$ 是否在 $A'$ 中。

```py [sol-Python3]
# 在文本串 text 中查找模式串 pattern
def kmp_search(text: str, pattern: str) -> bool:
    m = len(pattern)
    pi = [0] * m
    cnt = 0
    for i in range(1, m):
        b = pattern[i]
        while cnt and pattern[cnt] != b:
            cnt = pi[cnt - 1]
        if pattern[cnt] == b:
            cnt += 1
        pi[i] = cnt

    cnt = 0
    for b in text:
        while cnt and pattern[cnt] != b:
            cnt = pi[cnt - 1]
        if pattern[cnt] == b:
            cnt += 1
        if cnt == len(pattern):
            return True
    return False

class Solution:
    def repeatedStringMatch(self, a: str, b: str, x: int, y: int) -> bool:
        n, m = len(a), len(b)
        len_a, len_b = n * x, m * y
        if len_b > len_a:
            return False

        mx = min(n, len_a - len_b + 1)
        firsts = min(n + m, len_b)
        A = a * ((mx + firsts - 2) // n + 1)
        B = b * ((firsts - 1) // m + 1)
        return kmp_search(A, B)
```

```java [sol-Java]
class Solution {
    public boolean repeatedStringMatch(String a, String b, int x, int y) {
        int n = a.length();
        int m = b.length();
        long lenA = (long) n * x;
        long lenB = (long) m * y;
        if (lenB > lenA) {
            return false;
        }

        int mx = (int) Math.min(n, lenA - lenB + 1);
        int firsts = (int) Math.min(n + m, lenB);
        String A = a.repeat((mx + firsts - 2) / n + 1);
        String B = b.repeat((firsts - 1) / m + 1);
        return kmpSearch(A.toCharArray(), B.toCharArray());
    }

    // 在文本串 text 中查找模式串 pattern
    private boolean kmpSearch(char[] text, char[] pattern) {
        int m = pattern.length;
        int[] pi = new int[m];
        int cnt = 0;
        for (int i = 1; i < m; i++) {
            char b = pattern[i];
            while (cnt > 0 && pattern[cnt] != b) {
                cnt = pi[cnt - 1];
            }
            if (pattern[cnt] == b) {
                cnt++;
            }
            pi[i] = cnt;
        }

        cnt = 0;
        for (char b : text) {
            while (cnt > 0 && pattern[cnt] != b) {
                cnt = pi[cnt - 1];
            }
            if (pattern[cnt] == b) {
                cnt++;
            }
            if (cnt == m) {
                return true;
            }
        }
        return false;
    }
}
```

```cpp [sol-C++]
class Solution {
    string repeat(const string& s, int k) {
        string res;
        while (k--) {
            res += s;
        }
        return res;
    }

    // 在文本串 text 中查找模式串 pattern
    bool kmp_search(const string& text, const string& pattern) {
        int m = pattern.size();
        vector<int> pi(m);
        int cnt = 0;
        for (int i = 1; i < m; i++) {
            char b = pattern[i];
            while (cnt && pattern[cnt] != b) {
                cnt = pi[cnt - 1];
            }
            if (pattern[cnt] == b) {
                cnt++;
            }
            pi[i] = cnt;
        }

        cnt = 0;
        for (char b : text) {
            while (cnt && pattern[cnt] != b) {
                cnt = pi[cnt - 1];
            }
            if (pattern[cnt] == b) {
                cnt++;
            }
            if (cnt == m) {
                return true;
            }
        }
        return false;
    }

public:
    bool repeatedStringMatch(string a, string b, int x, int y) {
        int n = a.size(), m = b.size();
        long long len_a = 1LL * n * x, len_b = 1LL * m * y;
        if (len_b > len_a) {
            return false;
        }

        int mx = min(1LL * n, len_a - len_b + 1);
        int firsts = min(1LL * (n + m), len_b);
        string A = repeat(a, (mx + firsts - 2) / n + 1);
        string B = repeat(b, (firsts - 1) / m + 1);
        return kmp_search(A, B);
    }
};
```

```go [sol-Go]
// 在文本串 text 中查找模式串 pattern
func kmpSearch(text, pattern string) bool {
	m := len(pattern)
	pi := make([]int, m)
	cnt := 0
	for i := 1; i < m; i++ {
		b := pattern[i]
		for cnt > 0 && pattern[cnt] != b {
			cnt = pi[cnt-1]
		}
		if pattern[cnt] == b {
			cnt++
		}
		pi[i] = cnt
	}

	cnt = 0
	for _, b := range text {
		for cnt > 0 && pattern[cnt] != byte(b) {
			cnt = pi[cnt-1]
		}
		if pattern[cnt] == byte(b) {
			cnt++
		}
		if cnt == m {
			return true
		}
	}
	return false
}

func repeatedStringMatch(a, b string, x, y int) bool {
	n, m := len(a), len(b)
	lenA, lenB := n*x, m*y
	if lenB > lenA {
		return false
	}

	mx := min(n, lenA-lenB+1)
	firsts := min(n+m, lenB)
	A := strings.Repeat(a, (mx+firsts-2)/n+1)
	B := strings.Repeat(b, (firsts-1)/m+1)
	return kmpSearch(A, B)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $n$ 是 $a$ 的长度，$m$ 是 $\textit{b}$ 的长度。
- 空间复杂度：$\mathcal{O}(n+m)$。

## 相似题目

- [686. 重复叠加字符串匹配](https://leetcode.cn/problems/repeated-string-match/)
- [466. 统计重复个数](https://leetcode.cn/problems/count-the-repetitions/)

## 专题训练

见下面字符串题单的「**一、KMP**」。

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
