class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        unique_chars = list(set(ransomNote))
        count_1 = [ransomNote.count(c) for c in unique_chars]
        count_2 = [magazine.count(c) for c in unique_chars]
        for i in range(len(unique_chars)):
            if count_1[i] > count_2[i]:
                return False
        return True
