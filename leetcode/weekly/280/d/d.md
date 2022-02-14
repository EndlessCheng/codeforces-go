#### 提示 1

由于每个篮子至多可以放 $2$ 个整数，我们可以视作有 $2\cdot\textit{numSlots}$ 个篮子。

#### 提示 2

数据范围很小，考虑状压 DP。

#### 提示 3

以谁作为状态定义的对象呢？注意篮子编号是不能变的，而 $\textit{nums}$ 中元素的位置信息是不重要的。

---

由于每个篮子至多可以放 $2$ 个整数，我们可以视作有 $2\cdot\textit{numSlots}$ 个篮子。由于篮子个数很少，我们可以用二进制数 $x$ 表示这 $2\cdot\textit{numSlots}$ 个篮子中放了数字的篮子集合，其中 $x$ 从低到高的第 $i$ 位为 $1$ 表示第 $i$ 个篮子放了数字，为 $0$ 表示第 $i$ 个篮子为空。

设 $i$ 的二进制中的 $1$ 的个数为 $c$，定义 $f[i]$ 表示将 $\textit{nums}$ 的前 $c$ 个数字放到篮子中，且放了数字的篮子集合为 $i$ 时的最大与和。初始值 $f[0]=0$。

考虑将 $\textit{nums}[c]$ 放到一个空篮子时的状态转移方程（下标从 $0$ 开始，此时 $\textit{nums}[c]$ 还没被放入篮中），我们可以枚举 $i$ 中的 $0$，即空篮子的位置 $j$，该空篮子对应的编号为 $\dfrac{j}{2}+1$，则有

$$
f[i+2^j] = \max(f[i+2^j],\ f[i] + (\dfrac{j}{2}+1)\&\textit{nums}[c])
$$

设 $\textit{nums}$ 的长度为 $n$，最后答案为 $\max_{c=n}(f)$。

代码实现时需要注意，若 $c\ge n$ 则 $f[i]$ 无法转移，需要跳过。

相似题目：

- [1879. 两个数组最小的异或值之和](https://leetcode-cn.com/problems/minimum-xor-sum-of-two-arrays/)

```go [sol1-Go]
func maximumANDSum(nums []int, numSlots int) (ans int) {
	f := make([]int, 1<<(numSlots*2))
	for i, fi := range f {
		c := bits.OnesCount(uint(i))
		if c >= len(nums) {
			continue
		}
		for j := 0; j < numSlots*2; j++ {
			if i>>j&1 == 0 {
				s := i | 1<<j
				f[s] = max(f[s], fi+(j/2+1)&nums[c])
				ans = max(ans, f[s])
			}
		}
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

```C++ [sol1-C++]
class Solution {
public:
    int maximumANDSum(vector<int> &nums, int numSlots) {
        int ans = 0;
        vector<int> f(1 << (numSlots * 2));
        for (int i = 0; i < f.size(); ++i) {
            int c = __builtin_popcount(i);
            if (c >= nums.size()) continue;
            for (int j = 0; j < numSlots * 2; ++j) {
                if ((i & (1 << j)) == 0) {
                    int s = i | (1 << j);
                    f[s] = max(f[s], f[i] + ((j / 2 + 1) & nums[c]));
                    ans = max(ans, f[s]);
                }
            }
        }
        return ans;
    }
};
```

```Python [sol1-Python3]
class Solution:
    def maximumANDSum(self, nums: List[int], numSlots: int) -> int:
        ans = 0
        f = [0] * (1 << (numSlots * 2))
        for i, fi in enumerate(f):
            c = i.bit_count()
            if c >= len(nums):
                continue
            for j in range(numSlots * 2):
                if (i & (1 << j)) == 0:
                    s = i | (1 << j)
                    f[s] = max(f[s], fi + ((j // 2 + 1) & nums[c]))
                    ans = max(ans, f[s])
        return ans
```

```java [sol1-Java]
class Solution {
    public int maximumANDSum(int[] nums, int numSlots) {
        var ans = 0;
        var f = new int[1 << (numSlots * 2)];
        for (var i = 0; i < f.length; i++) {
            var c = Integer.bitCount(i);
            if (c >= nums.length) continue;
            for (var j = 0; j < numSlots * 2; ++j) {
                if ((i & (1 << j)) == 0) {
                    var s = i | (1 << j);
                    f[s] = Math.max(f[s], f[i] + ((j / 2 + 1) & nums[c]));
                    ans = Math.max(ans, f[s]);
                }
            }
        }
        return ans;
    }
}
```
