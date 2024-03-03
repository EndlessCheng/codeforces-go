按题意模拟即可。

```py [sol-Python3]
class Solution:
    def resultArray(self, nums: List[int]) -> List[int]:
        a = [nums[0]]
        b = [nums[1]]
        for x in nums[2:]:
            if a[-1] > b[-1]:
                a.append(x)
            else:
                b.append(x)
        return a + b
```

```java [sol-Java]
class Solution {
    public int[] resultArray(int[] nums) {
        int n = nums.length;
        List<Integer> a = new ArrayList<>();
        List<Integer> b = new ArrayList<>();
        a.add(nums[0]);
        b.add(nums[1]);
        for (int i = 2; i < n; i++) {
            if (a.get(a.size() - 1) > b.get(b.size() - 1)) {
                a.add(nums[i]);
            } else {
                b.add(nums[i]);
            }
        }
        a.addAll(b);
        for (int i = 0; i < n; i++) {
            nums[i] = a.get(i);
        }
        return nums;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> resultArray(vector<int> &nums) {
        vector<int> a{nums[0]}, b{nums[1]};
        for (int i = 2; i < nums.size(); i++) {
            (a.back() > b.back() ? a : b).push_back(nums[i]);
        }
        a.insert(a.end(), b.begin(), b.end());
        return a;
    }
};
```

```go [sol-Go]
func resultArray(nums []int) []int {
	a := nums[:1]
	b := []int{nums[1]}
	for _, x := range nums[2:] {
		if a[len(a)-1] > b[len(b)-1] {
			a = append(a, x)
		} else {
			b = append(b, x)
		}
	}
	return append(a, b...)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
