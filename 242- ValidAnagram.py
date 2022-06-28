class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        dict = {}
        for c in s:
            if c not in dict:
                dict[c] = 1
            else:
                dict[c] += 1
        for c in t:
            if c in dict:
                dict[c] -= 1
                if dict[c] < 0:
                    return False
            else:
                return False
        return True
