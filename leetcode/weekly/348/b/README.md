[视频讲解](https://www.bilibili.com/video/BV1do4y1K7Wq/) 第二题。

设 $1$ 在数组中的下标为 $p$，$n$ 在数组中的下标为 $q$。

分类讨论：

- 如果 $p<q$，那么 $1$ 和 $n$ 井水不犯河水，分别移动到数组的两端，操作次数为 $p + (n-1-q)$。
- 否则 $p>q$（注意 $n>2$），那么 $1$ 和 $n$ 分别移动到数组的两端的过程中会相遇，相当于只操作了一次就让两个数都向左向右移动了一步，所以操作次数比上面的情况要少 $1$，即 $p + (n-1-q) - 1$。

```py [sol-Python3]
class Solution:
    def semiOrderedPermutation(self, nums: List[int]) -> int:
        n = len(nums)
        p = nums.index(1)
        q = nums.index(n)
        return p + n - 1 - q - (p > q)
```

```java [sol-Java]
class Solution {
    public int semiOrderedPermutation(int[] nums) {
        int n = nums.length, p = 0, q = 0;
        for (int i = 0; i < n; i++) {
            if (nums[i] == 1) p = i;
            else if (nums[i] == n) q = i; // 注意 n>=2
        }
        return p + n - 1 - q - (p > q ? 1 : 0);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int semiOrderedPermutation(vector<int> &nums) {
        auto [p, q] = minmax_element(nums.begin(), nums.end());
        return p - q + nums.size() - 1 - (p > q);
    }
};
```

```go [sol-Go]
func semiOrderedPermutation(nums []int) int {
	n := len(nums)
	var p, q int
	for i, v := range nums {
		if v == 1 {
			p = i
		} else if v == n { // 注意 n>=2
			q = i
		}
	}
	if p < q {
		return p + n - 1 - q
	}
	return p + n - 2 - q // 1 向左移动的时候和 n 交换了一次
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
