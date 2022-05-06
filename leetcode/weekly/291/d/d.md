#### 提示 1-1

将**所有子串**按照其末尾字符的下标分组。

#### 提示 1-2

考虑两组**相邻**的子串：以 $s[i-1]$ 结尾的子串、以 $s[i]$ 结尾的子串。

#### 提示 1-3

以 $s[i]$ 结尾的子串，可以看成是以 $s[i-1]$ 结尾的子串，在末尾添加上 $s[i]$ 组成。

**上面这一串提示是思考子串统计类问题的通用技巧之一。**

#### 提示 2-1

从左往右遍历 $s$，考虑将 $s[i]$ 添加到以 $s[i-1]$ 结尾的子串的末尾。添加后，这些子串的引力值会增加多少？

#### 提示 2-2

分类讨论：

- 如果 $s[i]$ 之前没有遇到过，那么这些子串的引力值都会增加 $1$，这些子串的引力值之和会增加 $i$，再加上 $1$，即 $s[i]$ 单独组成的子串的引力值；
- 如果 $s[i]$ 之前遇到过，设其上次出现的下标为 $j$，那么向子串 $s[0..i-1],\ s[1..i-1],\ s[2..i-1],\cdots,s[j..i-1]$ 的末尾添加 $s[i]$ 后，这些子串的引力值是不会变化的，因为 $s[i]$ 已经在 $s[j]$ 处出现过了；而子串 $s[j+1..i-1],\ s[j+2..i-1],\cdots,s[i-1..i-1]$ 由于不包含字符 $s[i]$，这些子串的引力值都会增加 $1$，因此有 $i-j-1$ 个子串的引力值会增加 $1$，这些子串的引力值之和会增加 $i-j-1$，再加上 $1$，即 $s[i]$ 单独组成的子串的引力值。

---

模拟上述过程，遍历 $s$ 的过程中用一个变量 $\textit{sumG}$ 维护以 $s[i]$ 结尾的子串的引力值之和，同时用一个数组 $\textit{last}$ 记录每个字符上次出现的下标。

累加遍历中的 $\textit{sumG}$ 即为答案。（提示 1-1）

- 时间复杂度：$O(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$O(|\Sigma|)$，其中 $|\Sigma|$ 为字符集合的大小，本题中字符均为小写字母，所以 $|\Sigma|=26$。

```Python [sol1-Python3]
class Solution:
    def appealSum(self, s: str) -> int:
        ans, sum_g, last = 0, 0, [-1] * 26
        for i, c in enumerate(s):
            c = ord(c) - ord('a')
            sum_g += i - last[c]
            ans += sum_g
            last[c] = i
        return ans
```

```java [sol1-Java]
class Solution {
    public long appealSum(String s) {
        var ans = 0L;
        var last = new int[26];
        Arrays.fill(last, -1);
        for (int i = 0, sumG = 0; i < s.length(); i++) {
            var c = s.charAt(i) - 'a';
            sumG += i - last[c];
            ans += sumG;
            last[c] = i;
        }
        return ans;
    }
}
```

```C++ [sol1-C++]
class Solution {
public:
    long long appealSum(string &s) {
        long ans = 0L;
        vector<int> last(26, -1);
        for (int i = 0, sum_g = 0; i < s.length(); ++i) {
            char c = s[i] - 'a';
            sum_g += i - last[c];
            ans += sum_g;
            last[c] = i;
        }
        return ans;
    }
};
```

```go [sol1-Go]
func appealSum(s string) (ans int64) {
	sumG, last := 0, [26]int{}
	for i := range last { last[i] = -1 }
	for i, c := range s {
		c -= 'a'
		sumG += i - last[c]
		ans += int64(sumG)
		last[c] = i
	}
	return
}
```
