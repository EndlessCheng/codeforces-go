[本题视频讲解](https://www.bilibili.com/video/BV1Fg4y1Q7wv/)

设 $\textit{nums}$ 的异或和为 $s$。

$s=k$ 等价于 $s\oplus k = 0$，其中 $\oplus$ 表示异或。 

设 $x = s\oplus k$，我们把 $\textit{nums}$ 中的任意数字的某个比特位翻转，那么 $x$ 的这个比特位也会翻转。要让 $x=0$，就必须把 $x$ 中的每个 $1$ 都翻转，所以 $x$ 中的 $1$ 的个数就是我们的操作次数。

```py [sol-Python3]
class Solution:
    def minOperations(self, nums: List[int], k: int) -> int:
        return (reduce(xor, nums) ^ k).bit_count()
```

```java [sol-Java]
class Solution {
    public int minOperations(int[] nums, int k) {
        for (int x : nums) {
            k ^= x;
        }
        return Integer.bitCount(k);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(vector<int> &nums, int k) {
        for (int x : nums) {
            k ^= x;
        }
        return __builtin_popcount(k);
    }
};
```

```go [sol-Go]
func minOperations(nums []int, k int) int {
	for _, x := range nums {
		k ^= x
	}
	return bits.OnesCount(uint(k))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

周赛总结更新啦！请看 [2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
