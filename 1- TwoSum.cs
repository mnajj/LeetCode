public class Solution
{
	public int[] TwoSum(int[] nums, int target)
	{
		var dict = new Dictionary<int, int>();
		for (int i = 0; i < nums.Length; i++)
		{
			int num = target - nums[i];
			if (dict.ContainsKey(num) && dict[num] != i)
			{
				return new int[]
				{
					dict[num], i
				};
			}
			dict[nums[i]] = i;
		}
		return null;
	}
}
