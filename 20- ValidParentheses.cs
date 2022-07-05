
public class Solution
{
	public bool IsValid(string s)
	{
		Stack<char> stack = new();

		foreach (var c in s)
		{
			if (c == '(' || c == '{' || c == '[')
			{
				stack.Push(c);
			}
			else if (stack.Any())
			{
				if (c == ')')
					if (stack.Peek() == '(') stack.Pop();
					else return false;

				if (c == '}')
					if (stack.Peek() == '{') stack.Pop();
					else return false;

				if (c == ']')
					if (stack.Peek() == '[') stack.Pop();
					else return false;
			}
			else
			{
				return false;
			}
		}

		if (stack.Any()) return false;
		return true;
	}
}
