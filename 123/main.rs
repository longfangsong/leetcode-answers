use std::cmp;

struct Solution;

impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        let mut current_profit_trade_at_least_once_still_hold = i32::min_value();
        let mut current_profit_trade_at_least_once_not_hold = 0;
        let mut current_profit_trade_at_least_twice_still_hold = i32::min_value();
        let mut current_profit_trade_at_least_twice_not_hold = 0;
        for price in prices {
          current_profit_trade_at_least_twice_not_hold = cmp::max(
            current_profit_trade_at_least_twice_not_hold,
            current_profit_trade_at_least_twice_still_hold + price
          );
          current_profit_trade_at_least_twice_still_hold = cmp::max(
            current_profit_trade_at_least_twice_still_hold, 
            current_profit_trade_at_least_once_not_hold - price);

          current_profit_trade_at_least_once_not_hold = cmp::max(
            current_profit_trade_at_least_once_not_hold,
            current_profit_trade_at_least_once_still_hold + price
          );
          current_profit_trade_at_least_once_still_hold = cmp::max(
            current_profit_trade_at_least_once_still_hold, 
            -price);
        }
      current_profit_trade_at_least_twice_not_hold
    }
}