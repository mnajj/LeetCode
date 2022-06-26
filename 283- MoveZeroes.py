class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        n = len(nums)
        tail = n - 1
        head = 0
        temp = [None] * n
        for i in range(0, n):
            if nums[i] == 0:
                temp[tail] = 0
                tail -= 1
            else:
                temp[head] = nums[i]
                head += 1
        for i in range(0, n):
            nums[i] = temp[i]
        return None
            
# Effective
class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        n = len(nums)
        i = h = 0
        while i < n:
            if nums[i] != 0:
                nums[h] = nums[i]
                h += 1
            i += 1
        while h < n:
            nums[h] = 0
            h += 1
        return None
