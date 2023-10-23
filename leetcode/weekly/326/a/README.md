[视频讲解](https://www.bilibili.com/video/BV1H8411E7hn)

遍历每个数位，判断能否整除 $\textit{num}$。

不用转成字符串处理，而是不断取最低位（模 $10$），去掉最低位（除以 $10$），直到数字为 $0$。

```py [sol-Python3]
class Solution:
    def countDigits(self, num: int) -> int:
        ans, x = 0, num
        while x:
            ans += num % (x % 10) == 0
            x //= 10
        return ans
```

```java [sol-Java]
public class Solution {
    public int countDigits(int num) {
        int ans = 0;
        for (int x = num; x != 0; x /= 10) {
            ans += num % (x % 10) == 0 ? 1 : 0;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countDigits(int num) {
        int ans = 0;
        for (int x = num; x; x /= 10) {
            ans += num % (x % 10) == 0;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countDigits(num int) (ans int) {
	for x := num; x > 0; x /= 10 {
		if num%(x%10) == 0 {
			ans++
		}
	}
	return
}
```

```js [sol-JavaScript]
var countDigits = function(num) {
    let ans = 0;
    for (let x = num; x != 0; x = Math.floor(x / 10)) {
        ans += num % (x % 10) === 0 ? 1 : 0;
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn count_digits(num: i32) -> i32 {
        let mut ans = 0;
        let mut x = num;
        while x != 0 {
            if num % (x % 10) == 0 {
                ans += 1;
            }
            x /= 10;
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log \textit{num})$。即数字的十进制长度。
- 空间复杂度：$\mathcal{O}(1)$。

欢迎关注 [B站@灵茶山艾府](https://b23.tv/JMcHRRp)

[往期题解精选（按 tag 分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
