class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        m = len(matrix)
        n = len(matrix[0])
        l = 0
        h = m * n - 1
        while l <= h:
            mid = (h + l) // 2
            r = mid // n
            c = mid % n
            val = matrix[r][c]
            if val > target:
                h = mid - 1
            elif val < target:
                l = mid + 1
            else:
                return True
        return False
