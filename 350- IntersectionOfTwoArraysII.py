class Solution:
    def intersect(self, nums1: List[int], nums2: List[int]) -> List[int]:
        dict = {}
        res = []
        for n in nums1:
            if n not in dict:
                dict[n] = 1
            else:
                dict[n] += 1
        for n in nums2:
            if n in dict and dict[n] != 0:
                res.append(n)
                dict[n] -= 1
        return res
