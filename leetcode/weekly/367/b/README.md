[本题视频讲解](https://www.bilibili.com/video/BV1aC4y1G7dB/)

## 方法一：枚举

首先，如果 $s$ 中 $1$ 的个数不足 $k$，直接返回空串。

否则一定有解。

从 $k$ 开始枚举答案的长度 $\textit{size}$，然后在 $s$ 中枚举所有长为 $\textit{size}$ 的子串，同时维护字典序最小的子串。如果存在一个子串，其中 $1$ 的个数等于 $k$，则返回字典序最小的子串。

```py [sol-Python3]
class Solution:
    def shortestBeautifulSubstring(self, s: str, k: int) -> str:
        # 1 的个数不足 k
        if s.count('1') < k:
            return ''
        # 否则一定有解
        for size in count(k):  # 从 k 开始枚举
            ans = ''
            for i in range(size, len(s) + 1):
                t = s[i - size: i]
                if (ans == '' or t < ans) and t.count('1') == k:
                    ans = t
            if ans: return ans
```

```java [sol-Java]
class Solution {
    public String shortestBeautifulSubstring(String s, int k) {
        // 1 的个数不足 k
        if (s.replace("0", "").length() < k) {
            return "";
        }
        // 否则一定有解
        for (int size = k; ; size++) {
            String ans = "";
            for (int i = size; i <= s.length(); i++) {
                String t = s.substring(i - size, i);
                if ((ans.isEmpty() || t.compareTo(ans) < 0) && t.replace("0", "").length() == k) {
                    ans = t;
                }
            }
            if (!ans.isEmpty()) {
                return ans;
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string shortestBeautifulSubstring(string s, int k) {
        // 1 的个数不足 k
        if (count(s.begin(), s.end(), '1') < k) {
            return "";
        }
        // 否则一定有解
        for (int size = k;; size++) {
            string ans = "";
            for (int i = size; i <= s.length(); i++) {
                string t = s.substr(i - size, size);
                if ((ans == "" || t < ans) && count(t.begin(), t.end(), '1') == k) {
                    ans = t;
                }
            }
            if (!ans.empty()) {
                return ans;
            }
        }
    }
};
```

```go [sol-Go]
func shortestBeautifulSubstring(s string, k int) string {
	// 1 的个数不足 k
	if strings.Count(s, "1") < k {
		return ""
	}
	// 否则一定有解
	for size := k; ; size++ {
		ans := ""
		for i := size; i <= len(s); i++ {
			t := s[i-size : i]
			if (ans == "" || t < ans) && strings.Count(t, "1") == k {
				ans = t
			}
		}
		if ans != "" {
			return ans
		}
	}
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^3)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$。字符串切片需要 $\mathcal{O}(n)$ 的空间，Go 除外。

## 方法二：滑动窗口

原理请看 [滑动窗口【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)

由于答案中恰好有 $k$ 个 $1$，我们也可以用滑动窗口找最短的答案。

如果窗口内的 $1$ 的个数超过 $k$，或者窗口端点是 $0$，就可以缩小窗口。

> 注：利用字符串哈希（或者后缀数组等），可以把比较字典序的时间降至 $\mathcal{O}(n\log n)$，这样可以做到 $\mathcal{O}(n\log n)$ 的时间复杂度。

```py [sol-Python3]
class Solution:
    def shortestBeautifulSubstring(self, s: str, k: int) -> str:
        if s.count('1') < k:
            return ''
        ans = s
        cnt1 = left = 0
        for right, c in enumerate(s):
            cnt1 += int(c)
            while cnt1 > k or s[left] == '0':
                cnt1 -= int(s[left])
                left += 1
            if cnt1 == k:
                t = s[left: right + 1]
                if len(t) < len(ans) or len(t) == len(ans) and t < ans:
                    ans = t
        return ans
```

```java [sol-Java]
class Solution {
    public String shortestBeautifulSubstring(String S, int k) {
        if (S.replace("0", "").length() < k) {
            return "";
        }
        char[] s = S.toCharArray();
        String ans = S;
        int cnt1 = 0, left = 0;
        for (int right = 0; right < s.length; right++) {
            cnt1 += s[right] - '0';
            while (cnt1 > k || s[left] == '0') {
                cnt1 -= s[left++] - '0';
            }
            if (cnt1 == k) {
                String t = S.substring(left, right + 1);
                if (t.length() < ans.length() || t.length() == ans.length() && t.compareTo(ans) < 0) {
                    ans = t;
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string shortestBeautifulSubstring(string s, int k) {
        if (count(s.begin(), s.end(), '1') < k) {
            return "";
        }
        string ans = s;
        int cnt1 = 0, left = 0;
        for (int right = 0; right < s.length(); right++) {
            cnt1 += s[right] - '0';
            while (cnt1 > k || s[left] == '0') {
                cnt1 -= s[left++] - '0';
            }
            if (cnt1 == k) {
                string t = s.substr(left, right - left + 1);
                if (t.length() < ans.length() || t.length() == ans.length() && t < ans) {
                    ans = move(t);
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func shortestBeautifulSubstring(s string, k int) string {
	if strings.Count(s, "1") < k {
		return ""
	}
	ans := s
	cnt1 := 0
	left := 0
	for right, b := range s {
		cnt1 += int(b & 1)
		for cnt1 > k || s[left] == '0' {
			cnt1 -= int(s[left] & 1)
			left++
		}
		if cnt1 == k {
			t := s[left : right+1]
			if len(t) < len(ans) || len(t) == len(ans) && t < ans {
				ans = t
			}
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$。字符串切片需要 $\mathcal{O}(n)$ 的空间，Go 除外。
