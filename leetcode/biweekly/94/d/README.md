[视频讲解](https://www.bilibili.com/video/BV1Dd4y1h72z/) 已出炉，欢迎点赞三连，在评论区分享你对这场双周赛的看法~

---

首先，每个单词自己是互相独立的，因此分别计算每个单词的同位异构字符串的数目，再用乘法原理相乘。

对于一个长为 $n$ 的单词，其全排列的个数为 $n!$，但由于相同的字母不做区分，所以如果有 $x$ 个字母 $\text{a}$，还需要除以这些 $\text{a}$ 的全排列的个数，即 $x!$，对于其余字母同理。

代码实现时，分子分母可以分别计算，最后再用 [费马小定理](https://oi-wiki.org/math/number-theory/fermat/) 相除。

```py [sol1-Python3]
class Solution:
    def countAnagrams(self, s: str) -> int:
        MOD = 10 ** 9 + 7
        ans = mul = 1
        for s in s.split():
            cnt = Counter()
            for i, c in enumerate(s, 1):
                cnt[c] += 1
                mul = mul * cnt[c] % MOD
                ans = ans * i % MOD
        return ans * pow(mul, -1, MOD) % MOD
```

```java [sol1-Java]
class Solution {
    private static final int MOD = (int) 1e9 + 7;

    public int countAnagrams(String S) {
        var s = S.toCharArray();
        long ans = 1L, mul = 1L;
        var cnt = new int[26];
        for (int i = 0, j = 0; i < s.length; ++i) {
            if (s[i] == ' ') {
                Arrays.fill(cnt, 0);
                j = 0;
            } else {
                mul = mul * ++cnt[s[i] - 'a'] % MOD;
                ans = ans * ++j % MOD;
            }
        }
        return (int) (ans * pow(mul, MOD - 2) % MOD);
    }

    private long pow(long x, int n) {
        var res = 1L;
        for (; n > 0; n /= 2) {
            if (n % 2 > 0) res = res * x % MOD;
            x = x * x % MOD;
        }
        return res;
    }
}
```

```cpp [sol1-C++]
class Solution {
    const int MOD = 1e9 + 7;

    long pow(long x, int n) {
        long res = 1L;
        for (; n; n /= 2) {
            if (n % 2) res = res * x % MOD;
            x = x * x % MOD;
        }
        return res;
    }

public:
    int countAnagrams(string &s) {
        long ans = 1L, mul = 1L;
        int cnt[26]{};
        for (int i = 0, j = 0; i < s.length(); ++i) {
            if (s[i] == ' ') {
                memset(cnt, 0, sizeof(cnt));
                j = 0;
            } else {
                mul = mul * ++cnt[s[i] - 'a'] % MOD;
                ans = ans * ++j % MOD;
            }
        }
        return ans * pow(mul, MOD - 2) % MOD;
    }
};
```

```go [sol1-Go]
const mod int = 1e9 + 7

func countAnagrams(s string) int {
	ans, mul := 1, 1
	for _, s := range strings.Split(s, " ") {
		cnt := [26]int{}
		for i, c := range s {
			cnt[c-'a']++
			mul = mul * cnt[c-'a'] % mod
			ans = ans * (i + 1) % mod
		}
	}
	return ans * pow(mul, mod-2) % mod
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n >>= 1 {
		if n&1 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
```

#### 复杂度分析

- 时间复杂度：$O(n + \log M)$，其中 $n$ 为 $s$ 的长度，$M=10^9+7$。
- 空间复杂度：$O(|\Sigma|)$，其中 $|\Sigma|$ 为字符集合的大小，本题中字符均为小写字母，所以 $|\Sigma|=26$。为方便计算，代码中用到的库函数会带来额外空间消耗，这里忽略不计。
