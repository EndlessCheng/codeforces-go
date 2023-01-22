下午两点【B站@灵茶山艾府】直播讲题，关注UP不迷路~

---

#### 比赛时的想法

下文用 $g$ 表示最大公约数。

- 前两个移动很像辗转相除法（这个套路在 Codeforces 上已经出烂了），$g$ 不变；
- 后两个移动可以让 $g$ 乘上 $2^k$。

而一开始 $1,1$ 的 $g=1$，那么最终 $\textit{targetX},\textit{targetY}$ 的 $g$ 只能是 $2^k$。

根据 [231. 2 的幂](https://leetcode.cn/problems/power-of-two/)，用 $g\&(g-1)$ 是否为 $0$ 来判断 $g$ 是否为 $2^k$。

#### 如何构造出具体的移动方案？

从终点倒着走，那么 $(x,y)$ 可以移动到 $(x,x+y),(x+y,y),(x/2,y),(x,y/2)$ 这些位置，后两个移动只有偶数才能除以 $2$。

不断循环直到 $x=y$：

- 只要有偶数就除以 $2$。
- 如果都为奇数，比如 $x<y$，那么走两步可以得到 $(x,(x+y)/2)$，这里修改 $y$ 是因为 $x<(x+y)/2<y$。

那么总是可以让 $x$ 和 $y$ 不断变小。循环结束时如果 $x=1$，则说明可以做到。

```py [sol1-Python3]
class Solution:
    def isReachable(self, targetX: int, targetY: int) -> bool:
        g = gcd(targetX, targetY)
        return (g & (g - 1)) == 0
```

```java [sol1-Java]
class Solution {
    public boolean isReachable(int targetX, int targetY) {
        int g = gcd(targetX, targetY);
        return (g & (g - 1)) == 0;
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    bool isReachable(int targetX, int targetY) {
        int g = gcd(targetX, targetY);
        return (g & (g - 1)) == 0;
    }
};
```

```go [sol1-Go]
func isReachable(targetX, targetY int) bool {
	g := gcd(targetX, targetY)
	return g&(g-1) == 0
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$O(\log(\min(\textit{targetX}, \textit{targetY})))$。
- 空间复杂度：$O(1)$，仅用到若干变量。
