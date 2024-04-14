## 方法一：判断质数

原理请看 [视频讲解](https://www.bilibili.com/video/BV1dJ4m1V7hK/) 第二题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def maximumPrimeDifference(self, nums: List[int]) -> int:
        is_prime = lambda n: n >= 2 and all(n % i for i in range(2, isqrt(n) + 1))
        i = 0
        while not is_prime(nums[i]):
            i += 1
        j = len(nums) - 1
        while not is_prime(nums[j]):
            j -= 1
        return j - i
```

```java [sol-Java]
class Solution {
    public int maximumPrimeDifference(int[] nums) {
        int i = 0;
        while (!isPrime(nums[i])) {
            i++;
        }
        int j = nums.length - 1;
        while (!isPrime(nums[j])) {
            j--;
        }
        return j - i;
    }

    private boolean isPrime(int n) {
        for (int i = 2; i * i <= n; i++) {
            if (n % i == 0) {
                return false;
            }
        }
        return n >= 2;
    }
}
```

```cpp [sol-C++]
class Solution {
    bool is_prime(int n) {
        for (int i = 2; i * i <= n; i++) {
            if (n % i == 0) {
                return false;
            }
        }
        return n >= 2;
    }

public:
    int maximumPrimeDifference(vector<int>& nums) {
        int i = 0;
        while (!is_prime(nums[i])) {
            i++;
        }
        int j = nums.size() - 1;
        while (!is_prime(nums[j])) {
            j--;
        }
        return j - i;
    }
};
```

```go [sol-Go]
func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n >= 2
}

func maximumPrimeDifference(nums []int) int {
	i := 0
	for !isPrime(nums[i]) {
		i++
	}
	j := len(nums)-1
	for !isPrime(nums[j]) {
		j--
	}
	return j - i
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\sqrt U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：预处理

原理可以看 [视频讲解](https://www.bilibili.com/video/BV1H8411E7hn/) 第四题方法一。

```py [sol-Python3]
not_prime = [False] * 101
not_prime[1] = True
for i in 2, 3, 5, 7:
    for j in range(i * i, 101, i):
        not_prime[j] = True

class Solution:
    def maximumPrimeDifference(self, nums: List[int]) -> int:
        i = 0
        while not_prime[nums[i]]:
            i += 1
        j = len(nums) - 1
        while not_prime[nums[j]]:
            j -= 1
        return j - i
```

```py [sol-Python3 通用写法]
MX = 101
not_prime = [True, True] + [False] * (MX - 2)
for i in range(2, isqrt(MX) + 1):
    if not_prime[i]: continue
    for j in range(i * i, MX, i):
        not_prime[j] = True  # j 是质数 i 的倍数

class Solution:
    def maximumPrimeDifference(self, nums: List[int]) -> int:
        i = 0
        while not_prime[nums[i]]:
            i += 1
        j = len(nums) - 1
        while not_prime[nums[j]]:
            j -= 1
        return j - i
```

```java [sol-Java]
class Solution {
    private static final int MX = 101;
    private static final boolean[] NOT_PRIME = new boolean[MX];

    static {
        NOT_PRIME[1] = true;
        for (int i = 2; i * i < MX; i++) {
            if (NOT_PRIME[i]) continue;
            for (int j = i * i; j < MX; j += i) {
                NOT_PRIME[j] = true; // j 是质数 i 的倍数
            }
        }
    }

    public int maximumPrimeDifference(int[] nums) {
        int i = 0;
        while (NOT_PRIME[nums[i]]) {
            i++;
        }
        int j = nums.length - 1;
        while (NOT_PRIME[nums[j]]) {
            j--;
        }
        return j - i;
    }
}
```

```cpp [sol-C++]
const int MX = 101;
bool not_prime[MX];

auto init = [] {
    not_prime[1] = true;
    for (int i = 2; i * i < MX; i++) {
        if (not_prime[i]) continue;
        for (int j = i * i; j < MX; j += i) {
            not_prime[j] = true; // j 是质数 i 的倍数
        }
    }
    return 0;
}();

class Solution {
public:
    int maximumPrimeDifference(vector<int>& nums) {
        int i = 0;
        while (not_prime[nums[i]]) {
            i++;
        }
        int j = nums.size() - 1;
        while (not_prime[nums[j]]) {
            j--;
        }
        return j - i;
    }
};
```

```go [sol-Go]
const mx = 101
var notPrime = [mx]bool{true, true}
func init() {
	for i := 2; i*i < mx; i++ {
		if !notPrime[i] {
			for j := i * i; j < mx; j += i {
				notPrime[j] = true // j 是质数 i 的倍数
			}
		}
	}
}

func maximumPrimeDifference(nums []int) int {
	i := 0
	for notPrime[nums[i]] {
		i++
	}
	j := len(nums) - 1
	for notPrime[nums[j]] {
		j--
	}
	return j - i
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。忽略预处理的时间和空间。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
