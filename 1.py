class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        dict = {}
        for i in range(0, len(nums)):
            num = target - nums[i]
            if num in dict and dict[num] != i:
                return [dict[num], i]
            dict[nums[i]] = i
        return None
	
