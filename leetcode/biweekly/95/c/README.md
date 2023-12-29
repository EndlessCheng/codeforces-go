### 视频讲解

见[【双周赛 95】](https://www.bilibili.com/video/BV1i24y1e7E7/)第三题。视频额外讲了一种使用对称性的做法。

### 前言

我知道有的同学猜个结论过了，但假如我把问题中的「异或」换成「求和」，阁下又该如何应对？

### 思路

位运算经典技巧：由于每个比特位互不相干，所以拆分成每个比特位分别计算。

由于只有 $0$ 和 $1$，这样就好算了。

对异或有影响的是 $1$，所以只需要统计 $(a|b)\&c=1$ 的情况。

那么 $c$ 必须是 $1$，$a$ 和 $b$ 不能都是 $0$。

设有 $y$ 个 $1$，那么就有 $x=n-y$ 个 $0$。

那么 $c$ 有 $y$ 个，$a|b$ 有 $n^2-x^2$ 个（任意选是 $n^2$，减去两个都是 $0$ 的 $x^2$ 个）。

根据乘法原理，一共可以产生

$$
\textit{ones} = (n^2-x^2)y = (n^2-(n-y)^2)y = (2ny-y^2)y
$$

个 $1$。

由于异或只在乎 $\textit{ones}$ 的奇偶性，所以 $2ny$ 可以去掉，那么就变成看 $y^3$ 的奇偶性，也就是 $y$ 的奇偶性。

如果 $y$ 是奇数，那么这个比特位的异或值就是 $1$。

这实际上就是看每个比特位的异或值是否为 $1$。

那么把 $\textit{nums}$ 的每个数异或起来，就是答案。

```py [sol1-Python3]
class Solution:
    def xorBeauty(self, nums: List[int]) -> int:
        return reduce(xor, nums)
```

```java [sol1-Java]
class Solution {
    public int xorBeauty(int[] nums) {
        int ans = 0;
        for (int x : nums) ans ^= x;
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int xorBeauty(vector<int>& nums) {
        int ans = 0;
        for (int x : nums) ans ^= x;
        return ans;
    }
};
```

```go [sol1-Go]
func xorBeauty(nums []int) (ans int) {
	for _, x := range nums {
		ans ^= x
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干额外变量。
