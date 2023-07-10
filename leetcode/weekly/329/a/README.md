[视频讲解](https://www.bilibili.com/video/BV1Gv4y1y753/)

最简单的做法是把 $n$ 变成字符串，然后按照题目要求计算。

能否不转成字符串，直接计算呢？

我们可以从最低位开始，通过 $n \bmod 10$ 得到个位数，再把 $n$ 除以 $10$，重复前面的过程，得到十位数、百位数、……

最后得到最高位，如果发现它取的是负号，则把答案取反。

```py [sol-Python3]
class Solution:
    def alternateDigitSum(self, n: int) -> int:
        ans, sign = 0, 1
        while n:
            ans += n % 10 * sign
            sign = -sign
            n //= 10
        return ans * -sign
```

```java [sol-Java]
class Solution {
    public int alternateDigitSum(int n) {
        int ans = 0, sign = 1;
        for (; n > 0; n /= 10) {
            ans += n % 10 * sign;
            sign = -sign;
        }
        return ans * -sign;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int alternateDigitSum(int n) {
        int ans = 0, sign = 1;
        for (; n; n /= 10) {
            ans += n % 10 * sign;
            sign = -sign;
        }
        return ans * -sign;
    }
};
```

```go [sol-Go]
func alternateDigitSum(n int) (ans int) {
	sign := 1
	for ; n > 0; n /= 10 {
		ans += n % 10 * sign
		sign = -sign
	}
	return ans * -sign
}
```

```js [sol-JavaScript]
var alternateDigitSum = function (n) {
    let ans = 0, sign = 1;
    while (n) {
        ans += n % 10 * sign;
        sign = -sign;
        n = Math.floor(n / 10);
    }
    return ans * -sign;
};
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log n)$。
- 空间复杂度：$\mathcal{O}(1)$，仅用到若干变量。

[往期每日一题题解（按 tag 分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
