请看 [视频讲解](https://www.bilibili.com/video/BV1um4y1M7Rv/)。

忽略本题非常特殊的字符串长度，通用做法是，既然可以交换相距为 $2$ 的字符，那么相距为 $4$ 的字符可以通过多次交换实现。例如 $x-y-z$ 变成 $y-x-z$ 变成 $y-z-x$ 变成 $z-y-x$。

依此类推，所有相距为偶数的字符都可以随意交换。

所以只需要看下标为偶数的字符个数是否都一样，以及下标为奇数的字符个数是否都一样。

```py [sol-Python3]
class Solution:
    def canBeEqual(self, s1: str, s2: str) -> bool:
        return Counter(s1[::2]) == Counter(s2[::2]) and Counter(s1[1::2]) == Counter(s2[1::2])
```

```java [sol-Java]
class Solution {
    public boolean canBeEqual(String s1, String s2) {
        var cnt1 = new int[2][26];
        var cnt2 = new int[2][26];
        for (int i = 0; i < s1.length(); i++) {
            cnt1[i % 2][s1.charAt(i) - 'a']++;
            cnt2[i % 2][s2.charAt(i) - 'a']++;
        }
        return Arrays.deepEquals(cnt1, cnt2);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool canBeEqual(string s1, string s2) {
        int cnt1[2][26]{}, cnt2[2][26]{};
        for (int i = 0; i < s1.length(); i++) {
            cnt1[i % 2][s1[i] - 'a']++;
            cnt2[i % 2][s2[i] - 'a']++;
        }
        return memcmp(cnt1, cnt2, sizeof(cnt1)) == 0;
    }
};
```

```go [sol-Go]
func canBeEqual(s1, s2 string) bool {
	var cnt1, cnt2 [2][26]int
	for i, c := range s1 {
		cnt1[i%2][c-'a']++
		cnt2[i%2][s2[i]-'a']++
	}
	return cnt1 == cnt2
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+|\Sigma|)$，其中 $n$ 为 $s_1$ 的长度。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$，其中 $|\Sigma|$ 为字符集合的大小，本题中字符均为小写字母，所以 $|\Sigma|=26$。

## 思考题

改成 $j-i=3$ 要怎么做？
