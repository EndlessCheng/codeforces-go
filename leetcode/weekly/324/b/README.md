[视频讲解](https://www.bilibili.com/video/BV1LW4y1T7if/) 已出炉，欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

不断循环，计算 $n$ 的质因数之和 $s$。如果 $s=n$ 说明无法再继续减小 $n$ 了，返回 $n$；否则更新 $n$ 为 $s$，继续循环。

```py [sol1-Python3]
class Solution:
    def smallestValue(self, n: int) -> int:
        while True:
            x, s, i = n, 0, 2
            while i * i <= x:
                while x % i == 0:
                    s += i
                    x //= i
                i += 1
            if x > 1: s += x
            if s == n: return n
            n = s
```

```go [sol1-Go]
func smallestValue(n int) int {
	for {
		x, s := n, 0
		for i := 2; i*i <= x; i++ {
			for ; x%i == 0; x /= i {
				s += i
			}
		}
		if x > 1 {
			s += x
		}
		if s == n {
			return n
		}
		n = s
	}
}
```

#### 复杂度分析

- 时间复杂度：$O(\sqrt n)$。最坏情况下每次循环 $n$ 更新为 $2+n/2$，近似看成是 $n$ 减半，那么时间复杂度为 $O\left(\sqrt n + \sqrt\dfrac{n}{2} + \sqrt\dfrac{n}{4} + \cdots \right)=O\left(\sqrt n\cdot \left(1 +  \sqrt\dfrac{1}{2} + \sqrt\dfrac{1}{4} + \cdots\right)\right) = O(\sqrt n)$。
- 空间复杂度：$O(1)$，仅用到若干变量。
