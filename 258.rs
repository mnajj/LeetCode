impl Solution {
    pub fn add_digits(num: i32) -> i32 {
        return match num {
            0 => 0,
            _ => match num % 9 {
                0 => 9,
                r => r,
            },
        };
    }
}
