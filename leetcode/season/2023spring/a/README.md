## 本题视频讲解

见[【力扣杯2023春·个人赛】](https://www.bilibili.com/video/BV1dg4y1j78A/)。

## 方法一：模拟

为方便描述，下文简称 $\textit{supplies}$ 为 $a$。

不断循环，每次找到相邻和最小的 $a[i-1]$ 和 $a[i]$，然后把 $a[i-1]$ 增加 $a[i]$，并去掉 $a[i]$。

```py [sol1-Python3]
class Solution:
    def supplyWagon(self, a: List[int]) -> List[int]:
        m = len(a) // 2
        while len(a) > m:
            idx = 1
            for i in range(1, len(a)):
                if a[i - 1] + a[i] < a[idx - 1] + a[idx]:
                    idx = i
            a[idx - 1] += a[idx]
            a.pop(idx)
        return a
```

```java [sol1-Java]
class Solution {
    public int[] supplyWagon(int[] a) {
        int m = a.length / 2;
        for (int n = a.length; n > m; --n) {
            int j = 1;
            for (int i = 1; i < n; ++i)
                if (a[i] + a[i - 1] < a[j] + a[j - 1])
                    j = i;
            a[j - 1] += a[j];
            System.arraycopy(a, j + 1, a, j, n - 1 - j);
        }
        return Arrays.copyOf(a, m);
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    vector<int> supplyWagon(vector<int> &a) {
        int m = a.size() / 2;
        while (a.size() > m) {
            int j = 1;
            for (int i = 1; i < a.size(); ++i)
                if (a[i] + a[i - 1] < a[j] + a[j - 1])
                    j = i;
            a[j - 1] += a[j];
            a.erase(a.begin() + j);
        }
        return a;
    }
};
```

```go [sol1-Go]
func supplyWagon(a []int) []int {
	m := len(a) / 2
	for len(a) > m {
		j := 1
		for i := 1; i < len(a); i++ {
			if a[i]+a[i-1] < a[j]+a[j-1] {
				j = i
			}
		}
		a[j-1] += a[j]
		a = append(a[:j], a[j+1:]...)
	}
	return a
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $a$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

## 方法二：数组模拟双向链表+最小堆


```py [sol2-Python3]

```

```java [sol2-Java]

```

```cpp [sol2-C++]

```

```go [sol2-Go]

```
