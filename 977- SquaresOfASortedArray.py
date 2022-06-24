class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        n = len(nums)
        res = [0] * n
        l = 0
        r = n - 1
        idx = r
        while l <= r:
            if nums[l] ** 2 > nums[r] ** 2:
                res[idx] = nums[l] ** 2
                idx -= 1
                l += 1
            else:
                res[idx] = nums[r] ** 2
                idx -= 1
                r -= 1
        return res
