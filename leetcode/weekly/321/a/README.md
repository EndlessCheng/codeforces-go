### 视频讲解

见[【周赛 321】](https://www.bilibili.com/video/BV1sD4y1e7pr/)第一题。

### 思路

$1$ 到 $x$ 的元素和为 $\dfrac{x(x+1)}{2}$，$x$ 到 $n$ 的元素和为 $1$ 到 $n$ 的元素和减去 $1$ 到 $x-1$ 的元素和，即 $\dfrac{n(n+1)-x(x-1)}{2}$。

两式相等，简化后即

$$
x = \sqrt{\dfrac{n(n+1)}{2}}
$$

如果 $x$ 不是整数则返回 $-1$。

```py [sol-Python3]
class Solution:
    def pivotInteger(self, n: int) -> int:
        m = n * (n + 1) // 2
        x = isqrt(m)
        return x if x * x == m else -1
```

```java [sol-Java]
class Solution {
    public int pivotInteger(int n) {
        int m = n * (n + 1) / 2;
        int x = (int) Math.sqrt(m);
        return x * x == m ? x : -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int pivotInteger(int n) {
        int m = n * (n + 1) / 2;
        int x = sqrt(m);
        return x * x == m ? x : -1;
    }
};
```

```go [sol-Go]
func pivotInteger(n int) int {
	m := n * (n + 1) / 2
	x := int(math.Sqrt(float64(m)))
	if x*x == m {
		return x
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。计算平方根有专门的 CPU 指令，可以视作是 $\mathcal{O}(1)$ 时间。Python 的 `math.isqrt` 用的牛顿迭代法，这里也视作 $\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$，仅用到若干变量。

---

注意到 $\dfrac{n(n+1)}{2}$ 同时也是完全平方数的情况应该是比较少的（见 [OEIS A001108](https://oeis.org/A001108)）。在本题数据范围下，$n$ 只有

$$
1,8,49,288
$$

这四个，对应的答案（见 [OEIS A001109](https://oeis.org/A001109)）为

$$
1,6,35,204
$$

> 上述数据用程序枚举 $[1,1000]$ 内的 $n$，调用上面的代码，即可得到。

```py [sol-Python3]
ANS = {1: 1, 8: 6, 49: 35, 288: 204}

class Solution:
    def pivotInteger(self, n: int) -> int:
        return ANS.get(n, -1)
```

```java [sol-Java]
class Solution {
    private static final Map<Integer, Integer> m = Map.of(1, 1, 8, 6, 49, 35, 288, 204);

    public int pivotInteger(int n) {
        return m.getOrDefault(n, -1);
    }
}
```

```cpp [sol-C++]
class Solution {
    const unordered_map<int, int> m{{1, 1}, {8, 6}, {49, 35}, {288, 204}};
public:
    int pivotInteger(int n) {
        auto it = m.find(n);
        return it != m.end() ? it->second : -1;
    }
};
```

```go [sol-Go]
var m = map[int]int{1: 1, 8: 6, 49: 35, 288: 204}

func pivotInteger(n int) int {
	if ans, ok := m[n]; ok {
		return ans
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$，仅用到常数空间。
