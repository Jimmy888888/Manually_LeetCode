go forward only

can only hold at most one share of the stock at any time

trying to find the max profit

[7,1,5,3,6,4]

7: -6, -2, -4, -1, -3

1: 4, 2, 5, 3

5: -2, 1, -1

3: 3, 1

6: -2

it looks like not work

## only can go forward

using Greedy algorithm, partical best will be global best

ex: 
[3,6,1,4,10]

### way one: trying to find the best buy and sell point

buy: 3
sell: 6
profit: 3

buy: 6
sell: 6
profit: 0

buy: 1
sell: 10
profit: 9

total profit: 12

### way two: greedy algorithm

buy: 3
sell: 6
profit: 3

buy: 6
sell: 6
profit: 0

buy: 1
sell: 4
profit: 3

buy: 4
sell: 10
profit: 6

total profit: 12


## the result is the same

### the greedy algorithm strategy

at every point, if it is higher than the previous point, then sell it