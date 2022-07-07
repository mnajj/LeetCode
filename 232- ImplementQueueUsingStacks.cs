
public class MyQueue
{
	private Stack<int> _stack1;
	private Stack<int> _stack2;

	public MyQueue()
	{
		_stack1 = new();
		_stack2 = new();
	}

	public void Push(int x)
	{
		while (_stack1.Any())
		{
			_stack2.Push(_stack1.Pop());
		}

		_stack1.Push(x);
		while (_stack2.Any())
		{
			_stack1.Push(_stack2.Pop());
		}
	}

	public int Pop()
	{
		return _stack1.Pop();
	}

	public int? Peek()
	{
		return _stack1.Peek();
	}

	public bool Empty()
	{
		return _stack1.Any();
	}
}
