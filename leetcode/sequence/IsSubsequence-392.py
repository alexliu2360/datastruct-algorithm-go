# 思路一：库函数

class Solution1:
    def isSubsequence(self, s: str, t: str) -> bool:
        loc = -1
        for a in s:
            loc = t.find(a, loc + 1)
            if loc == -1:
                return False
        return True


# 思路二：生成迭代器

class Solution2:
    def isSubsequence(self, s: str, t: str) -> bool:
        t = iter(t)
        return all(c in t for c in s)


# 思路三：双指针

class Solution3:
    def isSubsequence(self, s: str, t: str) -> bool:
        i = 0
        j = 0
        while i < len(s) and j < len(t):
            # print(i, j)
            if s[i] == t[j]:
                i += 1
                j += 1
            else:
                j += 1
        return i == len(s)


# 思路四：二分法

class Solution4:
    def isSubsequence(self, s: str, t: str) -> bool:
        from collections import defaultdict
        import bisect
        lookup = defaultdict(list)
        for idx, val in enumerate(t):
            lookup[val].append(idx)
        # print(lookup)
        loc = -1
        for a in s:
            j = bisect.bisect_left(lookup[a], loc + 1)
            if j >= len(lookup[a]): return False
            loc = lookup[a][j]
        return True


if __name__ == "__main__":
    s = Solution2()
    r = s.isSubsequence("abc", "axxxxbxsssxqc")
    print(r)
