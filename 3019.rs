impl Solution {
    pub fn count_key_changes(s: String) -> i32 {
        let mut count = 0;
        let s = s.to_lowercase();
        if let Some(first_char) = s.chars().next() {
            let mut last_char = first_char;
            for c in s.chars() {
                if c.is_alphabetic() && c != last_char {
                    count += 1;
                    last_char = c;
                }
            }
        }
        count
    }
}
