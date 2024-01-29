[视频讲解](https://www.bilibili.com/video/BV1we411J7Y8/) 第四题。

## 提示 1

从高到低考虑：答案在这一位能不能是 $0$？

## 提示 2

一次操作，用 AND 合并两个相邻数字。

多次操作，相当于把一段**连续子数组**合并成 $0$。

尝试从左到右合并（忽略低位和必须是 $1$ 的位），如果合并出 $0$，就开始合并下一段。

## 思路

以 $[5,2,3,6]$ 为例说明，这四个数的二进制表示如下：

$$
\begin{aligned}
&101\\
&010\\
&011\\
&110
\end{aligned}
$$

设 $k=2$。从高到低考虑：答案在这一位能不能是 $0$？

- 最高位有两个 $1$，合并掉这两个 $1$ 需要操作 $2$ 次（$\le k$），所以答案的最高位可以是 $0$。
- 对于次高位，我们需要通过一连串的合并，让合并结果的最高位和次高位都是 $0$。数组前两个数（只看最高位和次高位）可以合并成 $0$，操作 $1$ 次。数组后两个数（只看最高位和次高位）无法合并成 $0$，那么用前两个数合并出来的 $0$，与后两个数合并，操作 $2$ 次，得到 $0$。所以一共要操作 $3$ 次才能让最高位和次高位都是 $0$，无法做到，所以答案的次高位一定是 $1$。
- 对于最低位，我们需要通过一连串的合并，让合并结果的最高位和最低位都是 $0$。注意我们**无需考虑次高位**，因为前面已经确定答案这一位是 $1$ 了。数组前两个数（只看最高位和最低位）可以合并成 $0$，操作 $1$ 次。数组后两个数（只看最高位和最低位）也可以合并成 $0$，操作 $1$ 次。一共操作 $2$ 次，所以答案的最低位可以是 $0$。
- 综上所述，答案的二进制表示为 $010$，即十进制 $2$。

代码中用到了一些位运算技巧，不了解位运算的同学可以看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

注意，如果整个数组都无法合并成 $0$，那么代码计算出来的操作次数是 $n$，题目保证了这是大于 $k$ 的，无需特判这种情况。

```py [sol-Python3]
class Solution:
    def minOrAfterOperations(self, nums: List[int], k: int) -> int:
        ans = mask = 0
        for b in range(max(nums).bit_length() - 1, -1, -1):
            mask |= 1 << b
            cnt = 0  # 操作次数
            and_res = -1  # -1 的二进制全为 1
            for x in nums:
                and_res &= x & mask
                if and_res:
                    cnt += 1  # 合并 x，操作次数加一
                else:
                    and_res = -1  # 准备合并下一段
            if cnt > k:
                ans |= 1 << b  # 答案的这个比特位必须是 1
                mask ^= 1 << b  # 后面不考虑这个比特位
        return ans
```

```java [sol-Java]
class Solution {
    public int minOrAfterOperations(int[] nums, int k) {
        int ans = 0;
        int mask = 0;
        for (int b = 29; b >= 0; b--) {
            mask |= 1 << b;
            int cnt = 0; // 操作次数
            int and = -1; // -1 的二进制全为 1
            for (int x : nums) {
                and &= x & mask;
                if (and != 0) {
                    cnt++; // 合并 x，操作次数加一
                } else {
                    and = -1; // 准备合并下一段
                }
            }
            if (cnt > k) {
                ans |= 1 << b; // 答案的这个比特位必须是 1
                mask ^= 1 << b; // 后面不考虑这个比特位
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOrAfterOperations(vector<int> &nums, int k) {
        int ans = 0, mask = 0;
        for (int b = 29; b >= 0; b--) {
            mask |= 1 << b;
            int cnt = 0, and_res = -1; // -1 的二进制全为 1
            for (int x : nums) {
                and_res &= x & mask;
                if (and_res) {
                    cnt++; // 合并 x，操作次数加一
                } else {
                    and_res = -1; // 准备合并下一段
                }
            }
            if (cnt > k) {
                ans |= 1 << b; // 答案的这个比特位必须是 1
                mask ^= 1 << b; // 后面不考虑这个比特位
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minOrAfterOperations(nums []int, k int) (ans int) {
	mask := 0
	for b := 29; b >= 0; b-- {
		mask |= 1 << b
		cnt := 0  // 操作次数
		and := -1 // -1 的二进制全为 1
		for _, x := range nums {
			and &= x & mask
			if and != 0 {
				cnt++ // 合并 x，操作次数加一
			} else {
				and = -1 // 准备合并下一段
			}
		}
		if cnt > k {
			ans |= 1 << b  // 答案的这个比特位必须是 1
			mask ^= 1 << b // 后面不考虑这个比特位
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(1)$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
