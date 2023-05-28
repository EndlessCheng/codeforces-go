## 视频讲解

见[【周赛 347】](https://www.bilibili.com/video/BV1fo4y1T7MQ/) 第一题，欢迎点赞投币！

## 思路

由于输入保证是**正整数**，所以去掉所有的尾零 $0$ 即可。

```py [sol-Python3]
class Solution:
    def removeTrailingZeros(self, num: str) -> str:
        return num.rstrip('0')
```

```java [sol-Java]
class Solution {
    public String removeTrailingZeros(String num) {
        return num.replaceAll("0+$", ""); // 注：可能是 O(n^2)，推荐手写
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string removeTrailingZeros(string s) {
        s.erase(find_if(s.rbegin(), s.rend(), [](auto c) { return c != '0'; }).base(), s.end());
        return s;
    }
};
```

```go [sol-Go]
func removeTrailingZeros(num string) string {
	return strings.TrimRight(num, "0")
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
