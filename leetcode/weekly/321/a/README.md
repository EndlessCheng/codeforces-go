[视频讲解](https://www.bilibili.com/video/BV1sD4y1e7pr/) 已出炉，欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

$1$ 到 $x$ 的元素和为 $\dfrac{x(x+1)}{2}$，$x$ 到 $n$ 的元素和为 $1$ 到 $n$ 的元素和减去 $1$ 到 $x-1$ 的元素和，即 $\dfrac{n(n+1)-x(x-1)}{2}$。

两式相等，简化后即

$$
x = \sqrt{\dfrac{n(n+1)}{2}}
$$

如果 $x$ 不是整数则返回 $-1$。

```py [sol1-Python3]
class Solution:
    def pivotInteger(self, n: int) -> int:
        m = n * (n + 1) // 2
        x = isqrt(m)
        return x if x * x == m else -1
```

```go [sol1-Go]
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

- 时间复杂度：$O(1)$。计算平方根有专门的 CPU 指令，可以视作是 $O(1)$ 时间。
- 空间复杂度：$O(1)$，仅用到若干变量。

---

注意到第 $n$ 个三角形数同时也是完全平方数的情况应该是比较少的（见 [OEIS A001108](https://oeis.org/A001108)），在题目数据范围下只有

$$
1,8,49,288
$$

这四个，对应的答案（见 [OEIS A001109](https://oeis.org/A001109)）为

$$
1,6,35,204
$$

```py [sol2-Python3]
ANS = {1: 1, 8: 6, 49: 35, 288: 204}

class Solution:
    def pivotInteger(self, n: int) -> int:
        return ANS.get(n, -1)
```

```go [sol2-Go]
var m = map[int]int{1: 1, 8: 6, 49: 35, 288: 204}

func pivotInteger(n int) int {
	if ans, ok := m[n]; ok {
		return ans
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$O(1)$。
- 空间复杂度：$O(1)$，仅用到常数空间。
