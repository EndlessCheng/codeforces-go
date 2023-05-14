下午两点[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，记得关注哦~

---

遍历每个字符串 $s$，如果 $s[11]$ 和 $s[12]$ 组成的数字大于 $60$，则答案加一。

```py [sol1-Python3]
class Solution:
    def countSeniors(self, details: List[str]) -> int:
        return sum(int(s[11:13]) > 60 for s in details)
```

```java [sol1-Java]
class Solution {
    public int countSeniors(String[] details) {
        int ans = 0;
        for (var s : details)
            if ((s.charAt(11) - '0') * 10 + s.charAt(12) - '0' > 60)
                ans++;
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int countSeniors(vector<string> &details) {
        int ans = 0;
        for (auto &s: details)
            ans += (s[11] - '0') * 10 + s[12] - '0' > 60;
        return ans;
    }
};
```

```go [sol1-Go]
func countSeniors(details []string) (ans int) {
	for _, s := range details {
		// 对于数字字符，&15 等价于 -'0'，但是不需要加括号
		if s[11]&15*10+s[12]&15 > 60 {
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{details}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
