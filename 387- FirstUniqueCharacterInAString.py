# Lookup
class Solution:
    def firstUniqChar(self, s: str) -> int:
        dict = {}
        for i in range(0, len(s)):
            if s[i] in dict:
                dict[s[i]][1] += 1
            else:
                dict[s[i]] = [i, 1]
        for k in dict:
            if dict[k][1] == 1:
                return dict[k][0]
        return -1

# Faster   
class Solution:
    def firstUniqChar(self, s: str) -> int:
        letters = 'abcdefghijklmnopqrstuvwxyz'
        index = [s.index(l) for l in letters if s.count(l) == 1]
        return min(index) if len(index) > 0 else -1      
