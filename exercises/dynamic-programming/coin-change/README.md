You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.



Example 1:

Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
Example 2:

Input: coins = [2], amount = 3
Output: -1
Example 3:

Input: coins = [1], amount = 0
Output: 0


Constraints:

1 <= coins.length <= 12
1 <= coins[i] <= 231 - 1
0 <= amount <= 104

Hints:
https://www.youtube.com/watch?v=YzIjgBONTJk&ab_channel=DEEPTITALESRA
https://www.youtube.com/watch?v=mBNrRy2_hVs&ab_channel=NeetCode
https://www.youtube.com/watch?v=dT6dvdbpChA&ab_channel=TECHDOSE
https://runestone.academy/ns/books/published/pythonds/Recursion/DynamicProgramming.html
http://www.cs.uni.edu/~fienup/cs270s04/lectures/lec6_1-29-04_coin_change_web.htm
https://runestone.academy/ns/books/published/pythonds/Recursion/DynamicProgramming.html#lst-change2
http://www.cs.uni.edu/~fienup/cs270s04/lectures/lec6_1-29-04_coin_change_web.htm