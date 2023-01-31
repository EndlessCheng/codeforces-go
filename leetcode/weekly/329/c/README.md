如果字符串中有 $1$，那么：

- 选 $1$ 和 $0$ 可以把 $0$ 变成 $1$；
- 选 $1$ 和 $1$ 可以把 $1$ 变成 $0$。

而如果只有 $0$，是无法得到 $1$ 的。

因此，只要两个字符串中都有 $1$ 或者都没有 $1$，就可以互相转换。

附：[视频讲解](https://www.bilibili.com/video/BV1Gv4y1y753/)

```py [sol1-Python3]
class Solution:
    def makeStringsEqual(self, s: str, target: str) -> bool:
        return ('1' in s) == ('1' in target)
```

```java [sol1-Java]
class Solution {
    public boolean makeStringsEqual(String s, String target) {
        return s.contains("1") == target.contains("1");
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    bool makeStringsEqual(string s, string target) {
        return (s.find('1') == string::npos) == (target.find('1') == string::npos);
    }
};
```

```go [sol1-Go]
func makeStringsEqual(s, target string) bool {
	return strings.Contains(s, "1") == strings.Contains(target, "1")
}
```

### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干额外变量。
