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

    
class Solution:
    def reverseWords(self, s: 'str') -> 'str':
        return " ".join(i[::-1] for i in s.split())
