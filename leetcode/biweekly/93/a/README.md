## 视频讲解

见[【双周赛 93】](https://www.bilibili.com/video/BV1kR4y1r7Df/)第一题。

```py [sol-Python3]
class Solution:
    def maximumValue(self, strs: List[str]) -> int:
        return max(int(s) if s.isdigit() else len(s)
                   for s in strs)
```

```java [sol-Java]
class Solution {
    public int maximumValue(String[] strs) {
        int ans = 0, x;
        for (var s : strs) {
            try {
                x = Integer.parseInt(s);
            } catch (Exception e) {
                x = s.length();
            }
            ans = Math.max(ans, x);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumValue(vector<string> &strs) {
        int ans = 0, x;
        for (auto &s: strs) {
            try {
                size_t sz = 0;
                x = stoi(s, &sz);
                if (sz < s.length())
                    x = s.length();
            } catch (exception) {
                x = s.length();
            }
            ans = max(ans, x);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumValue(strs []string) (ans int) {
	for _, s := range strs {
		x, err := strconv.Atoi(s)
		if err != nil {
			x = len(s)
		}
		ans = max(ans, x)
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L)$，其中 $L$ 为所有字符串的长度之和。
- 空间复杂度：$\mathcal{O}(1)$，仅用到若干额外变量。
