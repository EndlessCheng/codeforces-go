### 本题视频讲解

见[【周赛 338】](https://www.bilibili.com/video/BV11o4y1p7Ci/?t=3m)，从 3:00 开始。

### 前置知识：二分查找

见[【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

### 前置知识：筛质数

见[【周赛 326】](https://www.bilibili.com/video/BV1H8411E7hn/?t=9m58s)，从 9:58 开始。

### 思路

设 $\textit{pre}$ 是上一个减完后的数字，$x=\textit{nums}[i]$ 为当前数字。

设 $p$ 是满足 $x-p>\textit{pre}$ 的最大质数，换言之，$p$ 是小于 $x-\textit{pre}$ 的最大质数，这可以预处理质数列表后，用二分查找得到。

```py [sol1-Python3]
MX = 1000
P = [0]  # 哨兵，避免二分越界
is_prime = [True] * MX
for i in range(2, MX):
    if is_prime[i]:
        P.append(i)  # 预处理质数列表
        for j in range(i * i, MX, i):
            is_prime[j] = False

class Solution:
    def primeSubOperation(self, nums: List[int]) -> bool:
        pre = 0
        for x in nums:
            if x <= pre: return False
            pre = x - P[bisect_left(P, x - pre) - 1]  # 减去 < x-pre 的最大质数
        return True
```

```java [sol1-Java]
class Solution {
    private final static int MX = 1000;
    private final static int[] primes = new int[169];

    static {
        var np = new boolean[MX + 1];
        int pi = 1; // primes[0] = 0 避免二分越界
        for (int i = 2; i <= MX; ++i)
            if (!np[i]) {
                primes[pi++] = i; // 预处理质数列表
                for (int j = i; j <= MX / i; ++j)
                    np[i * j] = true;
            }
    }

    public boolean primeSubOperation(int[] nums) {
        int pre = 0;
        for (int x : nums) {
            if (x <= pre) return false;
            int j = lowerBound(primes, x - pre);
            pre = x - primes[j - 1]; // 减去 < x-pre 的最大质数
        }
        return true;
    }

    // 见 https://www.bilibili.com/video/BV1AP41137w7/
    private int lowerBound(int[] nums, int target) {
        int left = -1, right = nums.length; // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // nums[left] < target
            // nums[right] >= target
            int mid = (left + right) >>> 1;
            if (nums[mid] < target)
                left = mid; // 范围缩小到 (mid, right)
            else
                right = mid; // 范围缩小到 (left, mid)
        }
        return right;
    }
}
```

```cpp [sol1-C++]
const int MX = 1000;
vector<int> primes{0}; // 哨兵，避免二分越界

int init = []() {
    bool np[MX]{};
    for (int i = 2; i < MX; ++i)
        if (!np[i]) {
            primes.push_back(i); // 预处理质数列表
            for (int j = i; j < MX / i; ++j)
                np[i * j] = true;
        }
    return 0;
}();

class Solution {
public:
    bool primeSubOperation(vector<int> &nums) {
        int pre = 0;
        for (int x: nums) {
            if (x <= pre) return false;
            pre = x - *--lower_bound(primes.begin(), primes.end(), x - pre); // 减去 < x-pre 的最大质数
        }
        return true;
    }
};
```

```go [sol1-Go]
var p = []int{0} // 哨兵，避免二分越界

func init() {
	const mx = 1000
	np := [mx]bool{}
	for i := 2; i < mx; i++ {
		if !np[i] {
			p = append(p, i) // 预处理质数列表
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
}

func primeSubOperation(nums []int) bool {
	pre := 0
	for _, x := range nums {
		if x <= pre {
			return false
		}
		pre = x - p[sort.SearchInts(p, x-pre)-1] // 减去 < x-pre 的最大质数
	}
	return true
}
```

### 复杂度分析

- 时间复杂度：$O(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U$ 为 $1000$ 以内的质数个数。
- 空间复杂度：$O(1)$。忽略预处理的空间，仅用到若干额外变量。

### 相似题目

- [2523. 范围内最接近的两个质数](https://leetcode.cn/problems/closest-prime-numbers-in-range/)
