下午两点[【b站@灵茶山艾府】](https://b23.tv/JMcHRRp)直播讲题，欢迎关注！

```py [sol-Python3]
class Solution:
    def lastVisitedIntegers(self, words: List[str]) -> List[int]:
        ans = []
        nums = []
        k = 0
        for s in words:
            if s[0] != 'p':  # 不是 prev
                nums.append(int(s))
                k = 0
            else:
                k += 1
                ans.append(-1 if k > len(nums) else nums[-k])  # 倒数第 k 个
        return ans
```

```java [sol-Java]
// https://space.bilibili.com/206214
class Solution {
    public List<Integer> lastVisitedIntegers(List<String> words) {
        List<Integer> ans = new ArrayList<>();
        List<Integer> a = new ArrayList<>();
        int k = 0;
        for (String s : words) {
            if (s.charAt(0) != 'p') { // 不是 prev
                a.add(Integer.parseInt(s));
                k = 0;
            } else {
                ans.add(++k > a.size() ? -1 : a.get(a.size() - k)); // 倒数第 k 个
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> lastVisitedIntegers(vector<string> &words) {
        vector<int> ans, a;
        int k = 0;
        for (auto &s: words) {
            if (s[0] != 'p') { // 不是 prev
                a.push_back(stoi(s));
                k = 0;
            } else {
                ans.push_back(++k > a.size() ? -1 : a[a.size() - k]); // 倒数第 k 个
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func lastVisitedIntegers(words []string) (ans []int) {
	nums := []int{}
	k := 0
	for _, s := range words {
		if s[0] != 'p' { // 不是 prev
			x, _ := strconv.Atoi(s)
			nums = append(nums, x)
			k = 0
		} else {
			k++
			if k > len(nums) {
				ans = append(ans, -1)
			} else {
				ans = append(ans, nums[len(nums)-k]) // 倒数第 k 个
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nL)$，其中 $n$ 为 $\textit{words}$ 的长度，$L$ 为数字字符串的长度，不超过 $3$。
- 空间复杂度：$\mathcal{O}(n)$。
