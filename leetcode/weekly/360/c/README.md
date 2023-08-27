## 前置知识

请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

## 思路

由于可以把一个数一分为二，所以整个数组可以全部变成 $1$。因此如果 $\textit{nums}$ 的元素和小于 $\textit{target}$，则无解，返回 $-1$。否则一定有解。

然后从低位到高位贪心：

- 如果 $\textit{target}$ 的第 $i$ 位是 $0$，跳过。
- 如果 $\textit{target}$ 的第 $i$ 位是 $1$，那么先看看所有 $\le 2^i$ 的元素和能否 $\ge \textit{target}\& \textit{mask}$，其中 $\textit{mask}=2^{i+1}-1$。如果能，那么必然可以合并出 $\textit{target}\&  \textit{mask}$，无需操作（见 [视频](https://www.bilibili.com/video/BV1Em4y1T7Bq/) 中的证明）。
- 如果不能，那么就需要把一个更大的数（设它是 $2^j$）不断地一分为二，直到分解出 $2^i$ 为止。
- 注意分解完后，$2^i,2^{i+1},2^{i+2},\cdots,2^{j-1}$ 这些 $2$ 的幂我们都有了。所以后面 $i+1,i+2,\cdots, j-1$ 这些比特位都无需判断了，可以直接从第 $j$ 个比特位开始判断。

```py [sol-Python3]
class Solution:
    def minOperations(self, nums: List[int], target: int) -> int:
        if sum(nums) < target:
            return -1
        cnt = Counter(nums)
        ans = s = i = 0
        while 1 << i <= target:
            s += cnt[1 << i] << i
            mask = (1 << (i + 1)) - 1
            i += 1
            if s >= target & mask:
                continue
            ans += 1  # 一定要找更大的数操作
            while cnt[1 << i] == 0:
                ans += 1  # 还没找到，继续找更大的数
                i += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int minOperations(List<Integer> nums, int target) {
        long s = 0;
        var cnt = new int[31];
        for (int x : nums) {
            s += x;
            cnt[Integer.numberOfTrailingZeros(x)]++;
        }
        if (s < target)
            return -1;
        int ans = 0, i = 0;
        s = 0;
        while ((1L << i) <= target) {
            s += (long) cnt[i] << i;
            long mask = (1L << ++i) - 1;
            if (s >= (target & mask))
                continue;
            ans++; // 一定要找更大的数操作
            for (; cnt[i] == 0; i++)
                ans++; // 还没找到，继续找更大的数
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(vector<int> &nums, int target) {
        if (accumulate(nums.begin(), nums.end(), 0LL) < target)
            return -1;
        int cnt[31]{};
        for (int x: nums)
            cnt[__builtin_ctz(x)]++;
        int ans = 0, i = 0;
        long long s = 0;
        while ((1LL << i) <= target) {
            s += (long long) cnt[i] << i;
            int mask = (1LL << ++i) - 1;
            if (s >= (target & mask))
                continue;
            ans++; // 一定要找更大的数操作
            for (; cnt[i] == 0; i++)
                ans++; // 还没找到，继续找更大的数
        }
        return ans;
    }
};
```

```go [sol-Go]
func minOperations(nums []int, target int) (ans int) {
	s := 0
	cnt := [31]int{}
	for _, v := range nums {
		s += v
		cnt[bits.TrailingZeros(uint(v))]++
	}
	if s < target {
		return -1
	}
	s = 0
	for i := 0; 1<<i <= target; {
		s += cnt[i] << i
		mask := 1<<(i+1) - 1
		if s >= target&mask {
			i++
			continue
		}
		ans++
		for i++; cnt[i] == 0; i++ {
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+\log \textit{target})$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(\log \textit{target})$。
