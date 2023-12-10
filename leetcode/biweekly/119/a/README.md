把数组转成哈希表，就可以 $\mathcal{O}(1)$ 判断元素是否在数组中了。

```py [sol-Python3]
class Solution:
    def findIntersectionValues(self, nums1: List[int], nums2: List[int]) -> List[int]:
        set1 = set(nums1)
        set2 = set(nums2)
        return [sum(x in set2 for x in nums1),
                sum(x in set1 for x in nums2)]
```

```java [sol-Java]
class Solution {
    public int[] findIntersectionValues(int[] nums1, int[] nums2) {
        HashSet<Integer> set1 = new HashSet<>();
        for (int x : nums1) {
            set1.add(x);
        }
        HashSet<Integer> set2 = new HashSet<>();
        for (int x : nums2) {
            set2.add(x);
        }

        int[] ans = new int[2];
        for (int x : nums1) {
            if (set2.contains(x)) {
                ans[0]++;
            }
        }
        for (int x : nums2) {
            if (set1.contains(x)) {
                ans[1]++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> findIntersectionValues(vector<int> &nums1, vector<int> &nums2) {
        unordered_set<int> set1(nums1.begin(), nums1.end());
        unordered_set<int> set2(nums2.begin(), nums2.end());
        vector<int> ans(2);
        for (int x: nums1) ans[0] += set2.count(x);
        for (int x: nums2) ans[1] += set1.count(x);
        return ans;
    }
};
```

```go [sol-Go]
func findIntersectionValues(nums1, nums2 []int) []int {
	set1 := map[int]int{}
	for _, x := range nums1 {
		set1[x] = 1
	}
	set2 := map[int]int{}
	for _, x := range nums2 {
		set2[x] = 1
	}
	
	ans := [2]int{}
	for _, x := range nums1 {
		ans[0] += set2[x]
	}
	for _, x := range nums2 {
		ans[1] += set1[x]
	}
	return ans[:]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $n$ 为 $\textit{nums}_1$ 的长度，$m$ 为 $\textit{nums}_2$ 的长度。
- 空间复杂度：$\mathcal{O}(n+m)$。
