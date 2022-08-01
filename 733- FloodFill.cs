public class Solution
{
	private int[][] _grid;
	private int _from;
	private int _to;

	public int[][] FloodFill(int[][] image, int sr, int sc, int color)
	{
		_grid = image;
		_from = image[sr][sc];
		_to = color;
		if (_from != _to)
		{
			Dfs(sr, sc);
		}
		return image;
	}

	private void Dfs(int i, int j)
	{
		if (i >= 0 &&
			j >= 0 &&
			i < _grid.Length &&
			j < _grid[0].Length &&
			_grid[i][j] == _from)
		{
			_grid[i][j] = _to;
			Dfs(i + 1, j);
			Dfs(i - 1, j);
			Dfs(i, j + 1);
			Dfs(i, j - 1);
		}
	}
}
