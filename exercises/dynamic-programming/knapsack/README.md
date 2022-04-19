Given the weights and profits of ‘N’ items, we are asked to put these items in a knapsack that has a capacity ‘C’. 
The goal is to get the maximum profit from the items in the knapsack. Each item can only be selected once, 
as we don’t have multiple quantities of any item.

Let’s take Merry’s example, who wants to carry some fruits in the knapsack to get maximum profit. Here are the weights 
and profits of the fruits:

Items: { Apple, Orange, Banana, Melon }
Weights: { 2, 3, 1, 4 }
Profits: { 4, 5, 3, 7 }
Knapsack capacity: 5

Let’s try to put different combinations of fruits in the knapsack, such that their total weight is not more than 5:

Apple + Orange (total weight 5) => 9 profit
Apple + Banana (total weight 3) => 7 profit
Orange + Banana (total weight 4) => 8 profit
Banana + Melon (total weight 5) => 10 profit

This shows that Banana + Melon is the best combination, as it gives us the maximum profit and the total weight 
does not exceed the capacity

Documentation:
- https://www.educative.io/courses/grokking-dynamic-programming-patterns-for-coding-interviews/RM1BDv71V60
- https://www.youtube.com/watch?v=PfkBS9qIMRE&ab_channel=Jenny%27slecturesCS%2FITNET%26JRF
- https://www.youtube.com/watch?v=dT6dvdbpChA&ab_channel=TECHDOSE