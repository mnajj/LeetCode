class Solution(object):
    def maxSubArray(self, nums):
        curr_sum = 0
        max_sum = -sys.maxsize - 1
        for num in nums:
            curr_sum += num
            if (max_sum < curr_sum):
                max_sum = curr_sum
            if(curr_sum < 0):
                curr_sum = 0
        return max_sum
