public class Solution
{
	public int[] TwoSum(int[] nums, int target)
	{
		int[] ans = new int[2];
		var dict = new Dictionary<int, int>();
		for (int i = 0; i < nums.Length; i++)
		{
			dict[nums[i]] = i;
		}

		for (int i = 0; i < nums.Length; i++)
		{
			int num = target - nums[i];
			if (dict.ContainsKey(num) && dict[num] != i)
			{
				ans[0] = i;
				ans[1] = dict[num];
				break;
			}
		}
		return ans;
	}
}
