[视频讲解](https://www.bilibili.com/video/BV1we411J7Y8/)

由于同一字母的大小写，ASCII 码的低 $5$ 位是相同的，我们只需要统计

$$
s[i-1] \& 31 \ne s[i]\& 31
$$

的个数。

其中 $\& 31$ 表示取二进制的低 $5$ 位。

扩展阅读：[从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

```cpp [sol-C++]
class Solution {
public:
    int countKeyChanges(string s) {
        int ans = 0;
        for (int i = 1; i < s.length(); i++) {
            ans += (s[i - 1] & 31) != (s[i] & 31);
        }
        return ans;
    }
};
```

```java [sol-Java]
class Solution {
    public int countKeyChanges(String s) {
        int ans = 0;
        for (int i = 1; i < s.length(); i++) {
            if ((s.charAt(i - 1) & 31) != (s.charAt(i) & 31)) {
                ans++;
            }
        }
        return ans;
    }
}
```

```go [sol-Go]
func countKeyChanges(s string) (ans int) {
	for i := 1; i < len(s); i++ {
		if s[i-1]&31 != s[i]&31 {
			ans++
		}
	}
	return
}
```

```py [sol-Python3]
class Solution:
    def countKeyChanges(self, s: str) -> int:
        return sum(x != y for x, y in pairwise(s.lower()))
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
