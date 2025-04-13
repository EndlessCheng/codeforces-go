由于 $s$ 是回文串，我们只需关心左半部分如何排列。

特别地，如果 $s$ 的长度是奇数，那么正中间的那个字母 $\textit{c}$ 恰好出现奇数次，在重排后的回文串中，字母 $\textit{c}$ 也出现奇数次，所以正中间的字母也必须是 $\textit{c}$。

所以无论 $s$ 长度是奇数还是偶数，都只需关心左半部分（奇数去掉回文中心）如何排列。

要求字典序最小，那么把左半部分排序即可。右半部分通过反转左半部分得到。

[视频讲解](https://www.bilibili.com/video/BV1e3dBYLEDz/?t=1m21s)

## 写法一

```py [sol-Python3]
class Solution:
    def smallestPalindrome(self, s: str) -> str:
        n = len(s)
        t = sorted(s[:n // 2])

        ans = ''.join(t)
        if n % 2:
            ans += s[n // 2]
        return ans + ''.join(reversed(t))
```

```java [sol-Java]
class Solution {
    public String smallestPalindrome(String s) {
        int n = s.length();
        int m = n / 2;
        char[] t = s.substring(0, m).toCharArray();
        Arrays.sort(t);

        StringBuilder ans = new StringBuilder(n); // 预分配空间
        ans.append(t);
        if (n % 2 > 0) {
            ans.append(s.charAt(m));
        }
        for (int i = m - 1; i >= 0; i--) {
            ans.append(t[i]);
        }
        return ans.toString();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string smallestPalindrome(string s) {
        int n = s.size();
        string t = s.substr(0, n / 2);
        ranges::sort(t);

        string ans = t;
        if (n % 2) {
            ans += s[n / 2];
        }
        ranges::reverse(t);
        return ans + t;
    }
};
```

```go [sol-Go]
func smallestPalindrome(s string) string {
	n := len(s)
	t := []byte(s[:n/2])
	slices.Sort(t)

	ans := string(t)
	if n%2 > 0 {
		ans += string(s[n/2])
	}
	slices.Reverse(t)
	return ans + string(t)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 写法二：计数排序

```py [sol-Python3]
class Solution:
    def smallestPalindrome(self, s: str) -> str:
        n = len(s)
        cnt = Counter(s[:n // 2])

        ans = []
        for ch in ascii_lowercase:
            ans.append(ch * cnt[ch])
        ans = ''.join(ans)

        t = ans
        if n % 2:
            ans += s[n // 2]
        return ans + t[::-1]
```

```java [sol-Java]
class Solution {
    public String smallestPalindrome(String s) {
        int n = s.length();
        int[] cnt = new int[26];
        for (int i = 0; i < n / 2; i++) {
            cnt[s.charAt(i) - 'a']++;
        }

        StringBuilder ans = new StringBuilder(n); // 预分配空间
        for (int i = 0; i < 26; i++) {
            ans.repeat('a' + i, cnt[i]);
        }

        StringBuilder t = new StringBuilder(ans);
        if (n % 2 > 0) {
            ans.append(s.charAt(n / 2));
        }
        ans.append(t.reverse());
        return ans.toString();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string smallestPalindrome(string s) {
        int n = s.size();
        int cnt[26]{};
        for (int i = 0; i < n / 2; i++) {
            cnt[s[i] - 'a']++;
        }

        string ans;
        for (int i = 0; i < 26; i++) {
            ans += string(cnt[i], 'a' + i);
        }

        string t = ans;
        if (n % 2) {
            ans += s[n / 2];
        }
        ranges::reverse(t);
        return ans + t;
    }
};
```

```go [sol-Go]
func smallestPalindrome(s string) string {
	n := len(s)
	cnt := [26]int{}
	for _, b := range s[:n/2] {
		cnt[b-'a']++
	}

	ans := make([]byte, 0, n) // 预分配空间
	for i, c := range cnt {
		ans = append(ans, bytes.Repeat([]byte{'a' + byte(i)}, c)...)
	}

	t := slices.Clone(ans)
	if n%2 > 0 {
		ans = append(ans, s[n/2])
	}
	slices.Reverse(t)
	return string(append(ans, t...))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + |\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(n + |\Sigma|)$ 或 $\mathcal{O}(n)$，取决于实现。

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
