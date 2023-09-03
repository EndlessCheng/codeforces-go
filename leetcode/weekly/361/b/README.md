请看 [视频讲解](https://www.bilibili.com/video/BV1Nj411178Z/) 第二题。

一个数能被 $25$ 整除，有如下五种情况：

- 这个数是 $0$。
- 这个数的末尾是 $00$。
- 这个数的末尾是 $25$。
- 这个数的末尾是 $50$。
- 这个数的末尾是 $75$。

设 $\textit{num}$ 的长度为 $n$。

首先，根据题目说的，我们可以把 $\textit{num}$ 中的所有数字都删除，得到 $0$，这需要删除 $n$ 次。

但如果 $\textit{num}$ 中有 $0$，那么删除 $n-1$ 也可以得到 $0$。

接下来，看示例 1。以 $50$ 为例说明：

1. 从右到左遍历 $\textit{num}$，找到第一个 $0$。
2. 继续向左遍历，找到第一个 $5$，设其下标为 $i$。
3. 删除这个 $5$ 右侧的所有非 $0$ 数字，这样就得到了一个以 $50$ 结尾的数字。
4. 删除次数为 $n-i-2$，例如示例 1 中 $5$ 的下标是 $i=3$，需要删除 $7-3-2=2$ 次。

其余 $00,25,75$ 的计算方式同理，取 $n-i-2$ 的最小值作为答案。

注意：如果没有找到要找的字符，则跳过。

#### 答疑

**问**：如果删除后只得到 $00$，产生了前导零，这就不合法了呀？

**答**：不可能在删除后只得到 $00$，因为题目保证 $\textit{num}$ 不含前导零，如果有多个 $0$，那么 $0$ 的左侧必然有非 $0$ 数字。

```py [sol-Python3]
class Solution:
    def minimumOperations(self, num: str) -> int:
        n = len(num)
        def f(tail: str) -> int:
            i = num.rfind(tail[1])
            if i < 0: return n
            i = num.rfind(tail[0], 0, i)  # 写成 num[:i].rfind(tail[0]) 会产生额外的切片空间
            if i < 0: return n
            return n - i - 2
        return min(n - ('0' in num), f("00"), f("25"), f("50"), f("75"))
```

```java [sol-Java]
public class Solution {
    private int ans;

    public int minimumOperations(String num) {
        ans = num.length();
        if (num.contains("0"))
            ans--;
        f(num, "00");
        f(num, "25");
        f(num, "50");
        f(num, "75");
        return ans;
    }

    private void f(String num, String tail) {
        int n = num.length();
        int i = num.lastIndexOf(tail.charAt(1));
        if (i < 0) return;
        i = num.lastIndexOf(tail.charAt(0), i - 1);
        if (i < 0) return;
        ans = Math.min(ans, n - i - 2);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumOperations(string num) {
        int n = num.length();
        auto f = [&](string tail) {
            int i = num.rfind(tail[1]);
            if (i == string::npos || i == 0) return n;
            i = num.rfind(tail[0], i - 1);
            if (i == string::npos) return n;
            return n - i - 2;
        };
        return min({n - (num.find('0') != string::npos), f("00"), f("25"), f("50"), f("75")});
    }
};
```

```go [sol-Go]
func minimumOperations(num string) int {
	ans := len(num)
	if strings.Contains(num, "0") {
		ans-- // 可以删除 len(num)-1 次得到 "0"
	}
	f := func(tail string) {
		i := strings.LastIndexByte(num, tail[1])
		if i < 0 {
			return
		}
		i = strings.LastIndexByte(num[:i], tail[0])
		if i < 0 {
			return
		}
		ans = min(ans, len(num)-i-2)
	}
	f("00")
	f("25")
	f("50")
	f("75")
	return ans
}

func min(a, b int) int { if b < a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{num}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
