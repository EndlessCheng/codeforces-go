请看 [视频讲解](https://www.bilibili.com/video/BV1MG411X7zR/)。

把名字相同的员工对应的访问时间（转成分钟数）分到同一组中。

对于每一组的访问时间 $a$，排序后，判断是否有 $a[i]-a[i-2] < 60$，如果有，那么把这一组的员工名字加到答案中。

```py [sol-Python3]
class Solution:
    def findHighAccessEmployees(self, access_times: List[List[str]]) -> List[str]:
        name2times = defaultdict(list)
        for name, s in access_times:
            t = int(s[:2]) * 60 + int(s[2:])
            name2times[name].append(t)

        ans = []
        for name, a in name2times.items():
            a.sort()
            if any(a[i] - a[i - 2] < 60 for i in range(2, len(a))):
                ans.append(name)
        return ans
```

```java [sol-Java]
class Solution {
    public List<String> findHighAccessEmployees(List<List<String>> accessTimes) {
        Map<String, List<Integer>> groups = new HashMap<>();
        for (var entry : accessTimes) {
            String name = entry.get(0), s = entry.get(1);
            int t = Integer.parseInt(s.substring(0, 2)) * 60 + Integer.parseInt(s.substring(2));
            groups.computeIfAbsent(name, k -> new ArrayList<>()).add(t);
        }

        List<String> ans = new ArrayList<>();
        for (var entry : groups.entrySet()) {
            List<Integer> a = entry.getValue();
            Collections.sort(a);
            for (int i = 2; i < a.size(); i++) {
                if (a.get(i) - a.get(i - 2) < 60) {
                    ans.add(entry.getKey());
                    break;
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<string> findHighAccessEmployees(vector<vector<string>> &access_times) {
        map<string, vector<int>> groups;
        for (auto &p: access_times) {
            string name = p[0], s = p[1];
            int t = stoi(s.substr(0, 2)) * 60 + stoi(s.substr(2));
            groups[name].push_back(t);
        }

        vector<string> ans;
        for (auto &[name, a]: groups) {
            sort(a.begin(), a.end());
            for (int i = 2; i < a.size(); i++) {
                if (a[i] - a[i - 2] < 60) {
                    ans.emplace_back(name);
                    break;
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func findHighAccessEmployees(accessTimes [][]string) (ans []string) {
	groups := map[string][]int{}
	for _, p := range accessTimes {
		name, s := p[0], p[1]
		t := int(s[0]&15*10+s[1]&15)*60 + int(s[2]&15*10+s[3]&15)
		groups[name] = append(groups[name], t)
	}

	for name, a := range groups {
		slices.Sort(a)
		for i := 2; i < len(a); i++ {
			if a[i]-a[i-2] < 60 {
				ans = append(ans, name)
				break
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(Ln + n\log n)$，其中 $n$ 为 $\textit{accessTimes}$ 的长度，$L$ 为员工姓名的最大长度，本题不超过 $10$。
- 空间复杂度：$\mathcal{O}(Ln)$。
