请看 [视频讲解](https://www.bilibili.com/video/BV1U34y1N7Pe/) 第二题。

由于可以往 $8$ 个方向走，那么最快可以在

$$
\max\{|sx-fx|, |sy-fy|\}
$$

秒后到达终点（先斜着走再直走）。

上式只要小于等于 $t$ 就能恰好到达终点。因为我们可以在到达终点附近时，在终点周围不断「绕圈」消耗时间，这样可以直到最后一秒才走到终点。

**特殊情况**：如果起点和终点重合，那么 $t=1$ 的情况是无法回到起点的；如果 $t\ne 1$，我们可以同样地在起点不断「绕圈」，最后回到起点。

```py [sol-Python3]
class Solution:
    def isReachableAtTime(self, sx: int, sy: int, fx: int, fy: int, t: int) -> bool:
        if sx == fx and sy == fy:
            return t != 1
        return max(abs(sx - fx), abs(sy - fy)) <= t
```

```java [sol-Java]
public class Solution {
    public boolean isReachableAtTime(int sx, int sy, int fx, int fy, int t) {
        if (sx == fx && sy == fy) {
            return t != 1;
        }
        return Math.max(Math.abs(sx - fx), Math.abs(sy - fy)) <= t;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool isReachableAtTime(int sx, int sy, int fx, int fy, int t) {
        if (sx == fx && sy == fy)
            return t != 1;
        return max(abs(sx - fx), abs(sy - fy)) <= t;
    }
};
```

```go [sol-Go]
func isReachableAtTime(sx, sy, fx, fy, t int) bool {
	if sx == fx && sy == fy {
		return t != 1
	}
	return max(abs(sx-fx), abs(sy-fy)) <= t
}

func abs(x int) int { if x < 0 { return -x }; return x }
func max(a, b int) int { if b > a { return b }; return a }
```

```js [sol-JavaScript]
var isReachableAtTime = function(sx, sy, fx, fy, t) {
    if (sx === fx && sy === fy) {
        return t !== 1;
    }
    return Math.max(Math.abs(sx - fx), Math.abs(sy - fy)) <= t;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

如果改成只能往 $4$ 个方向走呢？
