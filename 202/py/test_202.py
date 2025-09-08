import unittest
from . import coinChange

class TestCoinChange(unittest.TestCase):
    def test_coin_change(self):
        test_cases = [
            ([1, 2, 5], 11, 3),
            ([2], 3, -1),
            ([1], 0, 0),
            ([1, 3, 4, 5], 7, 2),
            ([2, 5, 10, 1], 27, 4),
        ]
        
        for coins, amount, expected in test_cases:
            with self.subTest(coins=coins, amount=amount):
                self.assertEqual(coinChange(coins, amount), expected)

if __name__ == '__main__':
    unittest.main()
