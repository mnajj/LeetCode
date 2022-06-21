# straightforward
class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        m = len(matrix)
        n = len(matrix[0])
        idx = []
        for i in range(0, m):
            for j in range(0, n):
                if matrix[i][j] == 0:
                    idx.append([i, j])
        for i in idx:
            r = i[0]
            c = i[1]
            for i in range(0, n):
                matrix[r][i] = 0
            for i in range(0, m):
                matrix[i][c] = 0
                
