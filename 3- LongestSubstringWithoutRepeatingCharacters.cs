// straightforward
public class Solution
{
	public int LengthOfLongestSubstring(string s)
	{
		var dict = new Dictionary<int, bool>();
		var maxLen = 0;
		var loopLen = 0;

		for (int i = 0; i < s.Length; i++)
		{
			dict[s[i]] = true;
			loopLen++;
			for (int j = i + 1; j < s.Length; j++)
			{
				if (dict.ContainsKey(s[j]))
				{
					break;
				}

				dict[s[j]] = true;
				loopLen++;
			}

			if (maxLen < loopLen)
			{
				maxLen = loopLen;
			}

			loopLen = 0;
			dict.Clear();
		}

		return maxLen;
	}
}



public class Solution
{
	public int LengthOfLongestSubstring(string s)
	{
		var n = s.Length;
		var maxSub = 0;
		var charIdxDic = new Dictionary<int, int>();
		for (int i = 0; i < n; i++)
		{
			var c = s[i];
			if (charIdxDic.ContainsKey(c))
			{
				i = charIdxDic[c];
				charIdxDic.Clear();
			}
			else
			{
				charIdxDic.Add(c, i);
				maxSub = Math.Max(charIdxDic.Count, maxSub);
			}
		}
		return maxSub; 
	}
}



public class Solution
{
	public int LengthOfLongestSubstring(string s)
	{
		var n = s.Length;
		var maxSub = 0;
		var dic = new Dictionary<char, int>();
		var start = 0;

		for (int i = 0; i < n; i++)
		{
			var c = s[i];
			if (dic.ContainsKey(c))
			{
				start = Math.Max(start, dic[c] + 1);
			}

			maxSub = Math.Max(maxSub, i - start + 1);
			dic[c] = i;
		}

		return maxSub;
	}
}
