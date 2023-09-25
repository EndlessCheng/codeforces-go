到达端点需要 $n-1$ 时间。

设 $t=\textit{time}\bmod (n-1)$，分类讨论：

- 如果 $\dfrac{\textit{time}}{n-1}$ 是偶数，说明正在从 $1$ 到 $n$，答案为 $1+t$；
- 如果 $\dfrac{\textit{time}}{n-1}$ 是奇数，说明正在从 $n$ 到 $1$，答案为 $n-t$。

```py [sol-Python3]
class Solution:
    def passThePillow(self, n: int, time: int) -> int:
        k, t = divmod(time, n - 1)
        return n - t if k % 2 else 1 + t
```

```java [sol-Java]
class Solution {
    public int passThePillow(int n, int time) {
        int t = time % (n - 1);
        return time / (n - 1) % 2 > 0 ? n - t : 1 + t;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int passThePillow(int n, int time) {
        int t = time % (n - 1);
        return time / (n - 1) % 2 ? n - t : 1 + t;
    }
};
```

```go [sol-Go]
func passThePillow(n, time int) int {
	t := time % (n - 1)
	if time/(n-1)%2 > 0 {
		return n - t
	}
	return 1 + t
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。
