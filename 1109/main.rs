impl Solution {
    pub fn corp_flight_bookings(bookings: Vec<Vec<i32>>, n: i32) -> Vec<i32> {
        let mut result = vec![0; n as usize + 1];
        for booking in bookings {
            let (start, end, seats) = (booking[0] as usize - 1, booking[1] as usize - 1, booking[2]);
            // seats = the amount of change in the sum
            result[end + 1] -= seats;
            result[start] += seats;
        }
        let mut accumulated = 0;
        for i in 0..result.len() - 1 {
            accumulated += result[i];
            result[i] = accumulated;
        }
        result.pop();
        result
    }
}