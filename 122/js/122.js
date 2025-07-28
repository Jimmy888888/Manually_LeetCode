/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function(prices) {
    let buy = prices[0]
    let profit = 0

    for (const price of prices){
        if (price > buy){
            profit += price - buy
        }
        buy = price
    }

    return profit
};