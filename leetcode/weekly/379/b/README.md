[本题视频讲解](https://www.bilibili.com/video/BV1ae411e7fn/)

分类讨论：

- 如果车能直接攻击到皇后，答案是 $1$。
- 如果象能直接攻击到皇后，答案是 $1$。
- 如果车被象挡住，那么移走象，车就可以攻击到皇后，答案是 $2$。小知识：这在国际象棋中称作「闪击」。
- 如果象被车挡住，那么移走车，象就可以攻击到皇后，答案是 $2$。
- 如果车不能直接攻击到皇后，那么车可以水平移动或者垂直移动，其中一种方式必定不会被象挡住，可以攻击到皇后，答案是 $2$。

对于车，如果和皇后在同一水平线或者同一竖直线，且中间没有象，那么就可以直接攻击到皇后。

对于象，如果和皇后在同一斜线，且中间没有车，那么就可以直接攻击到皇后。判断的技巧我在[【基础算法精讲 16】](https://www.bilibili.com/video/BV1mY411D7f6/) 中有讲到，欢迎收看。

```py [sol-Python3]
class Solution:
    def minMovesToCaptureTheQueen(self, a: int, b: int, c: int, d: int, e: int, f: int) -> int:
        def ok(l: int, m: int, r: int) -> bool:
            return not min(l, r) < m < max(l, r)

        if a == e and (c != e or ok(b, d, f)) or \
           b == f and (d != f or ok(a, c, e)) or \
           c + d == e + f and (a + b != e + f or ok(c, a, e)) or \
           c - d == e - f and (a - b != e - f or ok(c, a, e)):
            return 1
        return 2
```

```java [sol-Java]
public class Solution {
    public int minMovesToCaptureTheQueen(int a, int b, int c, int d, int e, int f) {
        if (a == e && (c != e || ok(b, d, f)) ||
            b == f && (d != f || ok(a, c, e)) ||
            c + d == e + f && (a + b != e + f || ok(c, a, e)) ||
            c - d == e - f && (a - b != e - f || ok(c, a, e))) {
            return 1;
        }
        return 2;
    }

    private boolean ok(int l, int m, int r) {
        return m < Math.min(l, r) || m > Math.max(l, r);
    }
}
```

```cpp [sol-C++]
class Solution {
    bool ok(int l, int m, int r) {
        return m < min(l, r) || m > max(l, r);
    }

public:
    int minMovesToCaptureTheQueen(int a, int b, int c, int d, int e, int f) {
        if (a == e && (c != e || ok(b, d, f)) ||
            b == f && (d != f || ok(a, c, e)) ||
            c + d == e + f && (a + b != e + f || ok(c, a, e)) ||
            c - d == e - f && (a - b != e - f || ok(c, a, e))) {
            return 1;
        }
        return 2;
    }
};
```

```go [sol-Go]
func minMovesToCaptureTheQueen(a, b, c, d, e, f int) int {
	if a == e && (c != e || ok(b, d, f)) ||
		b == f && (d != f || ok(a, c, e)) ||
		c+d == e+f && (a+b != e+f || ok(c, a, e)) ||
		c-d == e-f && (a-b != e-f || ok(c, a, e)) {
		return 1
	}
	return 2
}

func ok(l, m, r int) bool {
	return m < min(l, r) || m > max(l, r)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

如果要输出移动方案呢？

附：周赛总结更新啦！请看 [2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
