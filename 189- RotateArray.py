# Pythonic
class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        i = 0
        n = len(nums) - 1
        while i < k:
            nums.insert(0, nums[n])
            del nums[n + 1]
            i += 1
        return None

# Effective
class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        n = len(nums)
        temp = [None] * n
        for i in range(0, n):
            temp[(i + k) % n] = nums[i]
        for i in range(0, n):
            nums[i] = temp[i]
        return None
