use std::collections::HashMap;

impl Solution {
    pub fn most_common_word(paragraph: String, banned: Vec<String>) -> String {
        let paragraph = paragraph
            .chars()
            .map(|c| {
                if c.is_alphabetic() || c.is_whitespace() {
                    c
                } else {
                    ' '
                }
            })
            .collect::<String>()
            .to_lowercase();
        let mut map: HashMap<&str, i32> = HashMap::new();
        for word in paragraph.split_whitespace() {
            if !banned.contains(&word.to_owned()) {
                let count = map.entry(word).or_insert(0);
                *count += 1;
            }
        }
        let (mut max_key, mut max_value) = ("", std::i32::MIN);
        for (key, &value) in &map {
            if value > max_value {
                max_key = *key;
                max_value = value;
            }
        }
        max_key.to_string()
    }
}
