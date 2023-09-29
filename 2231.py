class Solution:
    def largestInteger(self, num: int) -> int:
        l = list(map(int, str(num)))
        for i in range(len(l)):
            is_even = l[i] % 2 == 0
            for j in range(i+1, len(l)):
                if is_even == (l[j] % 2 == 0):
                    if l[i] < l[j]:
                        l[i], l[j] = l[j], l[i]
        return int(''.join(map(str, l)))
