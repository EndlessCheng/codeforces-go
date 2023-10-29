请看 [视频讲解](https://www.bilibili.com/video/BV1Tz4y1N7Wx/) 第二题。

美丽字符串等价于对于所有偶数下标 $i$，有

$$
s[i] = s[i+1]
$$

如果不满足上式，修改其中一个字母即可。

```py [sol-Python3]
class Solution:
    def minChanges(self, s: str) -> int:
        return sum(s[i] != s[i + 1] for i in range(0, len(s), 2))
```

```java [sol-Java]
class Solution {
    public int minChanges(String s) {
        int ans = 0;
        for (int i = 0; i < s.length(); i += 2) {
            if (s.charAt(i) != s.charAt(i + 1)) {
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minChanges(string s) {
        int ans = 0;
        for (int i = 0; i < s.length(); i += 2) {
            ans += s[i] != s[i + 1];
        }
        return ans;
    }
};
```

```go [sol-Go]
func minChanges(s string) (ans int) {
	for i := 0; i < len(s); i += 2 {
		if s[i] != s[i+1] {
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。
