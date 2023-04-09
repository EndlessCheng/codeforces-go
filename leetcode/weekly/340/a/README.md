### 本题视频讲解

如何高效地判断一个数是否为质数？见[【周赛 340】](https://www.bilibili.com/video/BV1iN411w7my/)。

### 思路

遍历两条对角线上的元素，如果是质数则更新答案的最大值。

注意 $1$ 不是质数。

```py [sol1-Python3]
def is_prime(n: int) -> bool:
    i = 2
    while i * i <= n:
        if n % i == 0:
            return False
        i += 1
    return n >= 2  # 1 不是质数

class Solution:
    def diagonalPrime(self, nums: List[List[int]]) -> int:
        ans = 0
        for i, row in enumerate(nums):
            for x in row[i], row[-1 - i]:
                if x > ans and is_prime(x):
                    ans = x
        return ans
```

```java [sol1-Java]
class Solution {
    public int diagonalPrime(int[][] nums) {
        int n = nums.length, ans = 0;
        for (int i = 0; i < n; ++i) {
            int x = nums[i][i];
            if (x > ans && isPrime(x))
                ans = x;
            x = nums[i][n - 1 - i];
            if (x > ans && isPrime(x))
                ans = x;
        }
        return ans;
    }

    private boolean isPrime(int n) {
        for (int i = 2; i * i <= n; ++i)
            if (n % i == 0)
                return false;
        return n >= 2; // 1 不是质数
    }
}
```

```cpp [sol1-C++]
class Solution {
    bool is_prime(int n) {
        for (int i = 2; i * i <= n; ++i)
            if (n % i == 0)
                return false;
        return n >= 2; // 1 不是质数
    }

public:
    int diagonalPrime(vector<vector<int>> &nums) {
        int n = nums.size(), ans = 0;
        for (int i = 0; i < n; ++i) {
            if (int x = nums[i][i]; x > ans && is_prime(x))
                ans = x;
            if (int x = nums[i][n - 1 - i]; x > ans && is_prime(x))
                ans = x;
        }
        return ans;
    }
};
```

```go [sol1-Go]
func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n >= 2 // 1 不是质数
}

func diagonalPrime(nums [][]int) (ans int) {
	for i, row := range nums {
		if x := row[i]; x > ans && isPrime(x) {
			ans = x
		}
		if x := row[len(nums)-1-i]; x > ans && isPrime(x) {
			ans = x
		}
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$O(n\sqrt{U})$，其中 $n$ 为 $\textit{nums}$ 的长度，$U$ 为两条对角线上的最大值。
- 空间复杂度：$O(1)$。仅用到若干额外变量。
