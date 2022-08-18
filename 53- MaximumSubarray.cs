public class Solution
{
	public int MaxSubArray(int[] nums)
	{
		var maxSub = nums.FirstOrDefault();
		int currSum = 0;
		for (int i = 1; i < nums.Length; i++)
		{
			if (currSum < 0) currSum = 0;
			currSum += nums[i];
			maxSub = Math.Max(maxSub, currSum);
		}
		return maxSub;
	}
}
