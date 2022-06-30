# Stack
class Solution:
    def reverseWords(self, s: str) -> str:
        stack = []
        rev = ""
        for c in s:
            if c != ' ':
                stack.append(c)
            else:
                rev += self.free_stack(stack)
        while stack:
            rev += stack.pop()
        return rev

    def free_stack(self, stack: List) -> str:
        s = ""
        while stack:
            s += stack.pop()
        s += " "
        return s

# Pythonic    
class Solution:
    def reverseWords(self, s: 'str') -> 'str':
        return " ".join(i[::-1] for i in s.split())


# Two pointers
class Solution:
    def reverseWords(self, s: str) -> str:
        l = r = 0
        rev = ''
        while r < len(s):
            if s[r] != ' ':
                r += 1
                continue
            rev += self.reverse(s, l, r - 1)
            l = r + 1
            r += 1
        rev += self.reverse(s, l, r - 1)
        rev = rev.rstrip()
        return rev

    def reverse(self, s, l, r):
        rev = ''
        while r >= l:
            rev += s[r]
            r -= 1
        rev += ' '
        return rev    
