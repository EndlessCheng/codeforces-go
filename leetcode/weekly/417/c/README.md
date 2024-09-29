问题等价于如下两个问题：

- 每个元音字母至少出现一次，并且**至少**包含 $k$ 个辅音字母的子串个数。记作 $f_k$。
- 每个元音字母至少出现一次，并且**至少**包含 $k+1$ 个辅音字母的子串个数。记作 $f_{k+1}$。

二者相减，所表达的含义就是**恰好**包含 $k$ 个辅音字母了，所以答案为 $f_k - f_{k+1}$。

对于每个问题，由于子串越长，越满足要求，有单调性，所以可以用**滑动窗口**解决。如果你不了解滑动窗口，可以看视频[【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

如果你之前没有做过统计子串/子数组个数的滑动窗口，推荐先完成 [2962. 统计最大元素出现至少 K 次的子数组](https://leetcode.cn/problems/count-subarrays-where-max-element-appears-at-least-k-times/)（[我的题解](https://leetcode.cn/problems/count-subarrays-where-max-element-appears-at-least-k-times/solutions/2560940/hua-dong-chuang-kou-fu-ti-dan-pythonjava-xvwg/)），这也是一道至少+统计个数的问题，且比本题要简单许多。

## 答疑

**问**：能不能把 $f_k$ 定义成「至多」？

**答**：至多和前面的「每个元音字母**至少**出现一次」相克，「至少」要求子串越长越好，而「至多」要求子串越短越好，这样必须分开求解（总共要计算四个滑动窗口），相比下面代码的直接求解要麻烦许多。

**问**：代码中的 `ans += left` 是什么意思？

**答**：滑动窗口的内层循环结束时，右端点**固定**在 $\textit{right}$，左端点在 $0,1,2,\cdots,\textit{left}-1$ 的所有子串都是合法的，这一共有 $\textit{left}$ 个。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1TqxCeZEmb/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def f(self, word: str, k: int) -> int:
        cnt1 = defaultdict(int)  # 元音
        ans = cnt2 = left = 0  # cnt2 维护辅音
        for b in word:
            if b in "aeiou":
                cnt1[b] += 1
            else:
                cnt2 += 1
            while len(cnt1) == 5 and cnt2 >= k:
                out = word[left]
                if out in "aeiou":
                    cnt1[out] -= 1
                    if cnt1[out] == 0:
                        del cnt1[out]
                else:
                    cnt2 -= 1
                left += 1
            ans += left
        return ans

    def countOfSubstrings(self, word: str, k: int) -> int:
        return self.f(word, k) - self.f(word, k + 1)
```

```java [sol-Java]
class Solution {
    public long countOfSubstrings(String word, int k) {
        char[] s = word.toCharArray();
        return f(s, k) - f(s, k + 1);
    }

    private long f(char[] word, int k) {
        long ans = 0;
        // 这里用哈希表实现，替换成数组会更快
        HashMap<Character, Integer> cnt1 = new HashMap<>(); // 元音
        int cnt2 = 0; // 辅音
        int left = 0;
        for (char b : word) {
            if ("aeiou".indexOf(b) >= 0) {
                cnt1.merge(b, 1, Integer::sum); // ++cnt1[b]
            } else {
                cnt2++;
            }
            while (cnt1.size() == 5 && cnt2 >= k) {
                char out = word[left];
                if ("aeiou".indexOf(out) >= 0) {
                    if (cnt1.merge(out, -1, Integer::sum) == 0) { // --cnt1[out] == 0
                        cnt1.remove(out);
                    }
                } else {
                    cnt2--;
                }
                left++;
            }
            ans += left;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
    const string VOWEL = "aeiou";

    long long f(string& word, int k) {
        long long ans = 0;
        // 这里用哈希表实现，替换成数组会更快
        unordered_map<char, int> cnt1; // 元音
        int cnt2 = 0; // 辅音
        int left = 0;
        for (char b : word) {
            if (VOWEL.find(b) != string::npos) {
                cnt1[b]++;
            } else {
                cnt2++;
            }
            while (cnt1.size() == 5 && cnt2 >= k) {
                char out = word[left];
                if (VOWEL.find(out) != string::npos) {
                    if (--cnt1[out] == 0) {
                        cnt1.erase(out);
                    }
                } else {
                    cnt2--;
                }
                left++;
            }
            ans += left;
        }
        return ans;
    }

public:
    long long countOfSubstrings(string word, int k) {
        return f(word, k) - f(word, k + 1);
    }
};
```

```go [sol-Go]
func f(word string, k int) (ans int64) {
	// 这里用哈希表实现，替换成数组会更快
	cnt1 := map[byte]int{} // 元音
	cnt2 := 0 // 辅音
	left := 0
	for _, b := range word {
		if strings.ContainsRune("aeiou", b) {
			cnt1[byte(b)]++
		} else {
			cnt2++
		}
		for len(cnt1) == 5 && cnt2 >= k {
			out := word[left]
			if strings.ContainsRune("aeiou", rune(out)) {
				cnt1[out]--
				if cnt1[out] == 0 {
					delete(cnt1, out)
				}
			} else {
				cnt2--
			}
			left++
		}
		ans += int64(left)
	}
	return
}

func countOfSubstrings(word string, k int) int64 {
	return f(word, k) - f(word, k+1)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{word}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

更多相似题目，见下面滑动窗口题单中的「**§5.1 三指针滑动窗口**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
