```py [sol-Python3]
class Solution:
    def triangleType(self, nums: List[int]) -> str:
        nums.sort()
        x, y, z = nums
        if x + y <= z:  # 排序后，只需比较 x+y 和 z
            return "none"
        if x == z:  # 排序了，说明 y 也和 x z 相等
            return "equilateral"
        if x == y or y == z:
            return "isosceles"
        return "scalene"
```

```java [sol-Java]
class Solution {
    public String triangleType(int[] nums) {
        Arrays.sort(nums);
        int x = nums[0];
        int y = nums[1];
        int z = nums[2];
        if (x + y <= z) { // 排序后，只需比较 x+y 和 z
            return "none";
        }
        if (x == z) { // 排序了，说明 y 也和 x z 相等
            return "equilateral";
        }
        if (x == y || y == z) {
            return "isosceles";
        }
        return "scalene";
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string triangleType(vector<int> &nums) {
        sort(nums.begin(), nums.end());
        int x = nums[0], y = nums[1], z = nums[2];
        if (x + y <= z) { // 排序后，只需比较 x+y 和 z
            return "none";
        }
        if (x == z) { // 排序了，说明 y 也和 x z 相等
            return "equilateral";
        }
        if (x == y || y == z) {
            return "isosceles";
        }
        return "scalene";
    }
};
```

```go [sol-Go]
func triangleType(nums []int) string {
	slices.Sort(nums)
	x, y, z := nums[0], nums[1], nums[2]
	if x+y <= z { // 排序后，只需比较 x+y 和 z
		return "none"
	}
	if x == z { // 排序了，说明 y 也和 x z 相等
		return "equilateral"
	}
	if x == y || y == z {
		return "isosceles"
	}
	return "scalene"
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
