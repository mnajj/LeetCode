class Solution:
    def firstBadVersion(self, n: int) -> int:
        low = 1
        high = n
        ans = -1
        while(low <= high):
            mid = (low+high) // 2
            if isBadVersion(mid):
                high = mid - 1
                ans = mid
            elif not isBadVersion(mid):
                low = mid + 1
        return ans
