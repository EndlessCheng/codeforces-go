下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

把替换操作看成是 $\texttt{1}$ 同时往左移动（如果左侧为 $\texttt{0}$ 的话）。

定义 $f[i]$ 表示把 $s$ 的前 $i$ 个字符完成移动所需的秒数。

如果 $s[i]=\texttt{0}$，则 $f[i]=f[i-1]$。

如果 $s[i]=\texttt{1}$，记前 $i$ 个字符中 $\texttt{0}$ 的个数为 $\textit{pre}_0$，则 $f[i]$ 至少为 $\textit{pre}_0$。如果 $\texttt{1}$ 在移动的过程中被前面的 $\texttt{1}$ 堵住了，那么必须要等前面空出一个 $\texttt{0}$ 才能继续移动；另外，一旦前面空出了 $\texttt{0}$，说明前方「道路通畅」，后续移动不会出现堵车，因此 $f[i]=f[i-1]+1$。

这两者取最大值，即

$$
f[i] = \max(f[i-1]+1, \textit{pre}_0)
$$

答案为 $f[n-1]$。

代码实现时 $f$ 可以压缩成一个变量。

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干变量。

```py [sol1-Python3]
class Solution:
    def secondsToRemoveOccurrences(self, s: str) -> int:
        f = pre0 = 0
        for c in s:
            if c == '0': pre0 += 1
            elif pre0: f = max(f + 1, pre0)  # 前面有 0 的时候才会移动
        return f
```

```java [sol1-Java]
class Solution {
    public int secondsToRemoveOccurrences(String s) {
        int f = 0, pre0 = 0;
        for (var i = 0; i < s.length(); i++)
            if (s.charAt(i) == '0') ++pre0;
            else if (pre0 > 0) f = Math.max(f + 1, pre0); // 前面有 0 的时候才会移动
        return f;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int secondsToRemoveOccurrences(string &s) {
        int f = 0, pre0 = 0;
        for (char c : s)
            if (c == '0') ++pre0;
            else if (pre0) f = max(f + 1, pre0); // 前面有 0 的时候才会移动
        return f;
    }
};
```

```go [sol1-Go]
func secondsToRemoveOccurrences(s string) (f int) {
	pre0 := 0
	for _, c := range s {
		if c == '0' {
			pre0++
		} else if pre0 > 0 { // 前面有 0 的时候才会移动
			f = max(f+1, pre0)
		}
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```
