请看 [视频讲解](https://www.bilibili.com/video/BV1e84y117R9/) 第四题。

## 提示 1

对于同一个比特位，由于 AND 和 OR 不会改变都为 $0$ 和都为 $1$ 的情况，所以操作等价于：

把一个数的 $0$ 和另一个数的同一个比特位上的 $1$ **交换**。

## 提示 2

设交换前两个数是 $x$ 和 $y$，且 $x>y$。把小的数上的 $1$ 给大的数，假设交换后 $x$ 增加了 $d$，那么 $y$ 也减少了 $d$。

交换前：$x^2+y^2$。

交换后：$(x+d)^2+(y-d)^2 = x^2+y^2+2d(x-y)+2d^2 > x^2+y^2$。

这说明应该通过交换，让一个数越大越好。

相当于把 $1$ 都**聚集**在一个数中，比分散到不同的数更好。

## 提示 3

由于可以操作任意次，那么一定可以「组装」出尽量大的数，做法如下：

1. 对于每个比特位，统计 $\textit{nums}$ 在这个比特位上有多少个 $1$，记到一个长（至多）为 $30$ 的 $\textit{cnt}$ 数组中（因为 $10^9 < 2^{30}$）。
2. 循环 $k$ 次。 
3. 每次循环，「组装」一个数（记作 $x$）：遍历 $\textit{cnt}$，只要 $\textit{cnt}[i]>0$ 就将其减一，同时将 $2^i$ 加到 $x$ 中。这样相当于把 $1$ 尽量聚集在一个数中。
4. 把 $x^2$ 加到答案中。

有关位运算的技巧，请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

```py [sol-Python3]
class Solution:
    def maxSum(self, nums: List[int], k: int) -> int:
        m = max(nums).bit_length()
        cnt = [0] * m
        for x in nums:
            for i in range(m):
                cnt[i] += x >> i & 1
        ans = 0
        for _ in range(k):
            x = 0
            for i in range(m):
                if cnt[i]:
                    cnt[i] -= 1
                    x |= 1 << i
            ans += x * x
        return ans % (10 ** 9 + 7)
```

```java [sol-Java]
public class Solution {
    public int maxSum(List<Integer> nums, int k) {
        final long MOD = 1_000_000_007;
        int[] cnt = new int[30];
        for (int x : nums) {
            for (int i = 0; i < 30; i++) {
                cnt[i] += (x >> i) & 1;
            }
        }
        long ans = 0;
        while (k-- > 0) {
            int x = 0;
            for (int i = 0; i < 30; i++) {
                if (cnt[i] > 0) {
                    cnt[i]--;
                    x |= 1 << i;
                }
            }
            ans = (ans + (long) x * x) % MOD;
        }
        return (int) ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxSum(vector<int> &nums, int k) {
        const int MOD = 1'000'000'007;
        int cnt[30]{};
        for (int x: nums) {
            for (int i = 0; i < 30; i++) {
                cnt[i] += (x >> i) & 1;
            }
        }
        long long ans = 0;
        while (k--) {
            int x = 0;
            for (int i = 0; i < 30; i++) {
                if (cnt[i] > 0) {
                    cnt[i]--;
                    x |= 1 << i;
                }
            }
            ans = (ans + (long long) x * x) % MOD;
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxSum(nums []int, k int) (ans int) {
	const mod = 1_000_000_007
	cnt := [30]int{}
	for _, x := range nums {
		for i := range cnt {
			cnt[i] += x >> i & 1
		}
	}
	for ; k > 0; k-- {
		x := 0
		for i := range cnt {
			if cnt[i] > 0 {
				cnt[i]--
				x |= 1 << i
			}
		}
		ans = (ans + x*x) % mod
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(\log U)$。
