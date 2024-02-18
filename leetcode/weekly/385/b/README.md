## 方法一：字符串

把 $\textit{arr}_1$ 的所有前缀丢到一个哈希集合中，然后遍历 $\textit{arr}_2$ 的所有前缀，统计在哈希集合中的最长长度。

```py [sol-Python3]
class Solution:
    def longestCommonPrefix(self, arr1: List[int], arr2: List[int]) -> int:
        st = set()
        for s in map(str, arr1):
            for i in range(1, len(s) + 1):
                st.add(s[:i])

        ans = 0
        for s in map(str, arr2):
            for i in range(1, len(s) + 1):
                if s[:i] not in st:
                    break
                ans = max(ans, i)
        return ans
```

```java [sol-Java]
class Solution {
    public int longestCommonPrefix(int[] arr1, int[] arr2) {
        Set<String> st = new HashSet<>();
        for (int x : arr1) {
            String s = Integer.toString(x);
            for (int i = 1; i <= s.length(); i++) {
                st.add(s.substring(0, i));
            }
        }

        int ans = 0;
        for (int x : arr2) {
            String s = Integer.toString(x);
            for (int i = 1; i <= s.length(); i++) {
                if (!st.contains(s.substring(0, i))) {
                    break;
                }
                ans = Math.max(ans, i);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestCommonPrefix(vector<int> &arr1, vector<int> &arr2) {
        unordered_set<string> st;
        for (int x : arr1) {
            string s = to_string(x);
            for (int i = 1; i <= s.length(); i++) {
                st.insert(s.substr(0, i));
            }
        }

        int ans = 0;
        for (int x : arr2) {
            string s = to_string(x);
            for (int i = 1; i <= s.length(); i++) {
                if (!st.contains(s.substr(0, i))) {
                    break;
                }
                ans = max(ans, i);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func longestCommonPrefix(arr1, arr2 []int) (ans int) {
	has := map[string]bool{}
	for _, x := range arr1 {
		s := strconv.Itoa(x)
		for i := 1; i <= len(s); i++ {
			has[s[:i]] = true
		}
	}

	for _, x := range arr2 {
		s := strconv.Itoa(x)
		for i := 1; i <= len(s) && has[s[:i]]; i++ {
			ans = max(ans, i)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+m)\log^2 U)$，其中 $n$ 为 $\textit{arr}_1$ 的长度，$m$ 为 $\textit{arr}_2$ 的长度，$U$ 为数组元素的最大值。
- 空间复杂度：$\mathcal{O}(n\log^2 U)$。

## 方法二：不用字符串

```py [sol-Python3]
class Solution:
    def longestCommonPrefix(self, arr1: List[int], arr2: List[int]) -> int:
        st = set()
        for x in arr1:
            while x:
                st.add(x)
                x //= 10

        ans = 0
        for x in arr2:
            while x and x not in st:
                x //= 10
            cnt = 0
            while x:
                cnt += 1
                x //= 10
            ans = max(ans, cnt)
        return ans
```

```java [sol-Java]
class Solution {
    public int longestCommonPrefix(int[] arr1, int[] arr2) {
        Set<Integer> st = new HashSet<>();
        for (int x : arr1) {
            for (; x > 0; x /= 10) {
                st.add(x);
            }
        }

        int ans = 0;
        for (int x : arr2) {
            for (; x > 0 && !st.contains(x); x /= 10) ;
            int cnt = 0;
            for (; x > 0; x /= 10) {
                cnt++;
            }
            ans = Math.max(ans, cnt);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestCommonPrefix(vector<int> &arr1, vector<int> &arr2) {
        unordered_set<int> st;
        for (int x : arr1) {
            for (; x; x /= 10) {
                st.insert(x);
            }
        }

        int ans = 0;
        for (int x : arr2) {
            for (; x && !st.contains(x); x /= 10);
            int cnt = 0;
            for (; x; x /= 10) {
                cnt++;
            }
            ans = max(ans, cnt);
        }
        return ans;
    }
};
```

```go [sol-Go]
func longestCommonPrefix(arr1, arr2 []int) (ans int) {
	has := map[int]bool{}
	for _, v := range arr1 {
		for ; v > 0; v /= 10 {
			has[v] = true
		}
	}

	for _, v := range arr2 {
		for ; v > 0 && !has[v]; v /= 10 {
		}
		cnt := 0
		for ; v > 0; v /= 10 {
			cnt++
		}
		ans = max(ans, cnt)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+m)\log U)$，其中 $n$ 为 $\textit{arr}_1$ 的长度，$m$ 为 $\textit{arr}_2$ 的长度，$U$ 为数组元素的最大值。
- 空间复杂度：$\mathcal{O}(n\log U)$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
