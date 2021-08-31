impl Solution {
    pub fn num_decodings(s: String) -> i32 {
        let s = s.as_bytes();
        if s[0] == b'0' {
            return 0;
        }
        let mut last_but_1 = 0;
        let mut last = 1;
        let mut current = 0;
        for current_length in 1..=s.len() {
            if s[current_length - 1] != b'0' {
                current = last;
            }
            if current_length >= 2
                && s[current_length - 2] != b'0'
                && (s[current_length - 2] < b'2'
                    || (s[current_length - 2] == b'2' && s[current_length - 1] <= b'6'))
            {
                current += last_but_1;
            }
            last_but_1 = last;
            last = current;
            current = 0;
        }
        last
    }
}
