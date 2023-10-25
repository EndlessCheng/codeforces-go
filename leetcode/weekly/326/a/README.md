核心思路：逐个遍历 $\textit{num}$ 的每个数位，统计能整除 $\textit{num}$ 的数位个数。

代码实现时，可以不用转成字符串处理，而是不断取最低位（模 $10$），去掉最低位（除以 $10$），直到数字为 $0$。

例如 $\textit{num}=123$：

1. 初始化 $x=\textit{num}$。
2. 通过 $x\bmod 10$ 取到个位数 $3$，然后把 $x$ 除以 $10$（下取整），得到 $x=12$。
3. 再次 $x\bmod 10$ 取到十位数 $2$，然后把 $x$ 除以 $10$（下取整），得到 $x=1$。
4. 最后 $x\bmod 10$ 取到百位数 $1$，然后把 $x$ 除以 $10$（下取整），得到 $x=0$。此时完成了遍历 $\textit{num}$ 的每个数位，退出循环。
5. 在这个过程中，设取到的数位为 $d$，每次遇到 $\textit{num}\bmod d = 0$ 的情况，就把答案加一。

```py [sol-Python3]
class Solution:
    def countDigits(self, num: int) -> int:
        ans = 0
        x = num
        while x:
            ans += num % (x % 10) == 0
            x //= 10
        return ans
```

```py [sol-Python3 写法二]
class Solution:
    def countDigits(self, num: int) -> int:
        ans = 0
        x = num
        while x:
            x, d = divmod(x, 10)
            ans += num % d == 0
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

- 时间复杂度：$\mathcal{O}(\log \textit{num})$，即 $\textit{num}$ 的十进制长度。
- 空间复杂度：$\mathcal{O}(1)$，仅用到若干变量。

欢迎关注 [B站@灵茶山艾府](https://b23.tv/JMcHRRp)

[往期题解精选（按 tag 分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
