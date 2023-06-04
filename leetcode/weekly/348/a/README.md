下午两点直播讲题，记得关注哦~（见个人主页）

---

根据题意，只要有重复字母，就可以继续操作。

所以最后每种出现过的字母，只会出现一次。所以答案就是 $s$ 中不同字母的个数。

代码实现时，可以用哈希表，也可以用一个二进制数记录每个字母是否出现过。原理下午直播会讲。

```py [sol-Python3]
class Solution:
    def minimizedStringLength(self, s: str) -> int:
        return len(set(s))
```

```java [sol-Java]
class Solution {
    public int minimizedStringLength(String s) {
        int mask = 0;
        for (var c : s.toCharArray())
            mask |= 1 << (c - 'a');
        return Integer.bitCount(mask);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimizedStringLength(string s) {
        int mask = 0;
        for (char c: s) mask |= 1 << (c - 'a');
        return __builtin_popcount(mask);
    }
};
```

```go [sol-Go]
func minimizedStringLength(s string) int {
	mask := uint(0)
	for _, c := range s {
		mask |= 1 << (c - 'a')
	}
	return bits.OnesCount(mask)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{s}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。用二进制实现，仅用到若干额外变量。
