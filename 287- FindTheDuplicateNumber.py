class Solution:
    def findDuplicate(self, nums: List[int]) -> int:
        dict = {}
        for i in range(len(nums)):
            if nums[i] in dict:
                return nums[i]
            else:
                dict[nums[i]] = True
