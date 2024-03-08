把一个 $1$ 放末尾，其余全部放在开头。

```py [sol-Python3]
class Solution:
    def maximumOddBinaryNumber(self, s: str) -> str:
        cnt1 = s.count('1')
        return '1' * (cnt1 - 1) + '0' * (len(s) - cnt1) + '1'
```

```java [sol-Java]
public class Solution {
    public String maximumOddBinaryNumber(String s) {
        int cnt1 = (int) s.chars().filter(c -> c == '1').count();
        return "1".repeat(cnt1 - 1) + "0".repeat(s.length() - cnt1) + "1";
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string maximumOddBinaryNumber(string s) {
        int cnt1 = ranges::count(s, '1');
        return string(cnt1 - 1, '1') + string(s.length() - cnt1, '0') + '1';
    }
};
```

```go [sol-Go]
func maximumOddBinaryNumber(s string) string {
	cnt1 := strings.Count(s, "1")
	return strings.Repeat("1", cnt1-1) + strings.Repeat("0", len(s)-cnt1) + "1"
}
```

```js [sol-JavaScript]
var maximumOddBinaryNumber = function (s) {
    const cnt1 = _.sumBy(s, c => c === '1')
    return '1'.repeat(cnt1 - 1) + '0'.repeat(s.length - cnt1) + '1';
};
```

```rust [sol-Rust]
impl Solution {
    pub fn maximum_odd_binary_number(s: String) -> String {
        let cnt1 = s.chars().filter(|&c| c == '1').count();
        "1".repeat(cnt1 - 1) + &*"0".repeat(s.len() - cnt1) + "1"
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
