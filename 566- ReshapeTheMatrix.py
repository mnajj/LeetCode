class Solution:
    def matrixReshape(self, mat: List[List[int]], r: int, c: int) -> List[List[int]]:
        m = len(mat)
        n = len(mat[0])

        if m*n != r*c:
            return mat

        flat_mat = []
        for i in range(0, m):
            for j in range(0, n):
                flat_mat.append(mat[i][j])
        reshaped_mat = []
        row = []
        counter = 0
        for i in range(0, r):
            for j in range(0, c):
                row.append(flat_mat[counter])
                counter += 1
            reshaped_mat.append(row)
            row = []
        return reshaped_mat
