class Solution {
public:
		int scoreOfParentheses(string s) {
			stack<int> st;
			st.push(0);
			for (char c: s) {
				if (c == '(') {
					st.push(0);
				} else {
					int last = st.top();
					st.pop();
					if (!last) last = 1;
					else last *= 2;
					int parent = st.top() + last;
					st.pop();
					st.push(parent);
				}
			}
			return st.top();
		}
};
