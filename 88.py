class Solution(object):
    def merge(self, nums1, m, nums2, n):
      if (n == 0):
        return
      if (m == 0):
        for i in range(len(nums2)):
          nums1[i] = nums2[i]
        return
      merged = []
      itr1 = itr2 = 0
      while itr1 < m and itr2 < n:
        if nums1[itr1] <= nums2[itr2]:
          merged.append(nums1[itr1])
          itr1 += 1
        else:
          merged.append(nums2[itr2])
          itr2 += 1
      if itr2 < n:
        for i in range(itr2, n):
          merged.append(nums2[i])
      if itr1 < m:
        for i in range(itr1, m):
          merged.append(nums1[i])      
      for i in range(len(merged)):
        nums1[i] = merged[i]
