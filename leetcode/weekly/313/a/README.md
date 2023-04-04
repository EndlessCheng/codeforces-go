### 视频讲解

为什么只需要枚举到 $\sqrt{g}$？见[【周赛 313】](https://www.bilibili.com/video/BV1kd4y1q7fC)。

### 思路

枚举因子，挨个判断能否整除 $a$ 和 $b$。

改进方案是枚举 $a$ 和 $b$ 的最大公因数的因子。

```py [sol1-Python3]
class Solution:
    def commonFactors(self, a: int, b: int) -> int:
        g = gcd(a, b)
        ans, i = 0, 1
        while i * i <= g:
            if g % i == 0:
                ans += 1  # i 是公因子
                if i * i < g:
                    ans += 1  # g/i 是公因子
            i += 1
        return ans
```

```java [sol1-Java]
class Solution {
    public int commonFactors(int a, int b) {
        int ans = 0, g = gcd(a, b);
        for (int i = 1; i * i <= g; ++i)
            if (g % i == 0) {
                ++ans; // i 是公因子
                if (i * i < g)
                    ++ans; // g/i 是公因子
            }
        return ans;
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
    int commonFactors(int a, int b) {
        int ans = 0, g = gcd(a, b);
        for (int i = 1; i * i <= g; ++i)
            if (g % i == 0) {
                ++ans; // i 是公因子
                if (i * i < g)
                    ++ans; // g/i 是公因子
            }
        return ans;
    }
};
```

```go [sol1-Go]
func commonFactors(a, b int) (ans int) {
	g := gcd(a, b)
	for i := 1; i*i <= g; i++ {
		if g%i == 0 {
			ans++ // i 是公因子
			if i*i < g {
				ans++ // g/i 是公因子
			}
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$O(\sqrt{\min(a,b)})$。
- 空间复杂度：$O(1)$，仅用到若干变量。
