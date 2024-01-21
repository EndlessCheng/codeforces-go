[视频讲解](https://www.bilibili.com/video/BV1oV411D7gB/)

题目意思：把数组分成三段，每一段取第一个数再求和，问和的最小值是多少。

第一段的第一个数是确定的，即 $\textit{nums}[0]$。

如果知道了第二段的第一个数的位置，和第三段的第一个数的位置，那么这个划分方案也就确定了。

这两个下标可以在 $[1,n-1]$ 中随意取。

所以问题变成求下标在 $[1,n-1]$ 中的两个最小的数。

## 方法一：直接排序

```py [sol-Python3]
class Solution:
    def minimumCost(self, nums: List[int]) -> int:
        return nums[0] + sum(sorted(nums[1:])[:2])
```

```java [sol-Java]
class Solution {
    public int minimumCost(int[] nums) {
        Arrays.sort(nums, 1, nums.length);
        return nums[0] + nums[1] + nums[2];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumCost(vector<int> &nums) {
        sort(nums.begin() + 1, nums.end());
        return accumulate(nums.begin(), nums.begin() + 3, 0);
    }
};
```

```go [sol-Go]
func minimumCost(nums []int) int {
	slices.Sort(nums[1:])
	return nums[0] + nums[1] + nums[2]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 方法二：维护最小值和次小值

```py [sol-Python3]
class Solution:
    def minimumCost(self, nums: List[int]) -> int:
        return nums[0] + sum(nsmallest(2, nums[1:]))
```

```java [sol-Java]
class Solution {
    public int minimumCost(int[] nums) {
        int fi = Integer.MAX_VALUE, se = Integer.MAX_VALUE;
        for (int i = 1; i < nums.length; i++) {
            int x = nums[i];
            if (x < fi) {
                se = fi;
                fi = x;
            } else if (x < se) {
                se = x;
            }
        }
        return nums[0] + fi + se;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumCost(vector<int> &nums) {
        int fi = INT_MAX, se = INT_MAX;
        for (int i = 1; i < nums.size(); i++) {
            int x = nums[i];
            if (x < fi) {
                se = fi;
                fi = x;
            } else if (x < se) {
                se = x;
            }
        }
        return nums[0] + fi + se;
    }
};
```

```go [sol-Go]
func minimumCost(nums []int) int {
	fi, se := math.MaxInt, math.MaxInt
	for _, x := range nums[1:] {
		if x < fi {
			se = fi
			fi = x
		} else if x < se {
			se = x
		}
	}
	return nums[0] + fi + se
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
