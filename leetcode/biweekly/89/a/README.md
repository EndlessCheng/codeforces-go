### 视频讲解

见[【双周赛 89】](https://www.bilibili.com/video/BV1cV4y157BY)。

### 思路

分别计算小时和分钟的有效个数，根据乘法原理，答案为这两个个数的乘积。

对于小时，枚举 $[0,23]$ 内的每个整数 $i$，需同时满足下面两个条件：

- $\textit{time}[0]$ 是 $\texttt{?}$ 或者等于 $i$ 的十位数字（$i<10$ 时为 $0$）。
- $\textit{time}[1]$ 是 $\texttt{?}$ 或者等于 $i$ 的个位数字。

对于分钟，枚举 $[0,59]$ 内的每个整数 $i$，需同时满足下面两个条件：

- $\textit{time}[3]$ 是 $\texttt{?}$ 或者等于 $i$ 的十位数字（$i<10$ 时为 $0$）。
- $\textit{time}[4]$ 是 $\texttt{?}$ 或者等于 $i$ 的个位数字。

代码实现时，可以把这两个判断逻辑合并至一个函数 $\texttt{count}$ 中。

```py [sol1-Python3]
def count(t: str, period: int) -> int:
    ans = 0
    for i in range(period):
        if (t[0] == '?' or i // 10 == int(t[0])) and \
           (t[1] == '?' or i % 10 == int(t[1])):
            ans += 1
    return ans

class Solution:
    def countTime(self, time: str) -> int:
        return count(time[:2], 24) * count(time[3:], 60)
```

```java [sol1-Java]
class Solution {
    public int countTime(String time) {
        return count(time.substring(0, 2), 24) * count(time.substring(3), 60);
    }

    private int count(String time, int period) {
        var t = time.toCharArray();
        int ans = 0;
        for (int i = 0; i < period; i++)
            if ((t[0] == '?' || i / 10 == t[0] - '0') &&
                (t[1] == '?' || i % 10 == t[1] - '0'))
                ans++;
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
    int count(string t, int period) {
        int ans = 0;
        for (int i = 0; i < period; i++)
            if ((t[0] == '?' || i / 10 == t[0] - '0') &&
                (t[1] == '?' || i % 10 == t[1] - '0'))
                ans++;
        return ans;
    }
public:
    int countTime(string time) {
        return count(time.substr(0, 2), 24) * count(time.substr(3), 60);
    }
};
```

```go [sol1-Go]
func count(t string, period int) (ans int) {
    for i := 0; i < period; i++ {
        if (t[0] == '?' || i/10 == int(t[0]-'0')) &&
           (t[1] == '?' || i%10 == int(t[1]-'0')) {
            ans++
        }
    }
    return
}

func countTime(time string) int {
    return count(time[:2], 24) * count(time[3:], 60)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(24+60)=\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$，仅用到若干额外变量。

---

[往期每日一题题解](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注[ biIibiIi@灵茶山艾府](https://space.bilibili.com/206214)，高质量算法教学，持续输出中~
