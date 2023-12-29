遍历每个字符串 $s$，如果 $s[11]$ 和 $s[12]$ 组成的数字大于 $60$，则答案加一。

```py [sol-Python3]
class Solution:
    def countSeniors(self, details: List[str]) -> int:
        return sum(int(s[11:13]) > 60 for s in details)
```

```java [sol-Java]
class Solution {
    public int countSeniors(String[] details) {
        int ans = 0;
        for (String s : details)
            if ((s.charAt(11) - '0') * 10 + s.charAt(12) - '0' > 60)
                ans++;
        return ans;
    }
}
```

```cpp [sol-C++]
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

```go [sol-Go]
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

```js [sol-JavaScript]
var countSeniors = function(details) {
    let ans = 0;
    for (const s of details) {
        ans += parseInt(s.substring(11, 13)) > 60 ? 1 : 0;
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn count_seniors(details: Vec<String>) -> i32 {
        let mut ans = 0;
        for s in &details {
            if s[11..13].parse::<i32>().unwrap() > 60 {
                ans += 1;
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{details}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

更多精彩题解，请看 [往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
