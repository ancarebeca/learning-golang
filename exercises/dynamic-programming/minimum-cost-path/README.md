Minimum Cost Path Problem

    Given a square grid of size N, each cell of which contains integer cost which represents a cost to traverse through that cell, 
    we need to find a path from top left cell to bottom right cell by which the total cost incurred is minimum.
    From the cell (i,j) we can go (i,j-1), (i, j+1), (i-1, j), (i+1, j).
    
    Note: It is assumed that negative cost cycles do not exist in the input matrix.

    Example 1:

    Input: grid = {{9,4,9,9},{6,7,6,4},
    {8,3,3,7},{7,4,9,10}}
    Output: 43
    Explanation: The grid is-
    9 4 9 9
    6 7 6 4
    8 3 3 7
    7 4 9 10
    The minimum cost is-
    9 + 4 + 7 + 3 + 3 + 7 + 10 = 43.

    Your Task:
    You don't need to read or print anything. Your task is to complete the function minimumCostPath() which takes grid as 
    input parameter and returns the minimum cost to react at bottom right cell from top left cell.
    
    Expected Time Compelxity: O(n2*log(n))
    Expected Auxiliary Space: O(n2)

    Constraints:
    1 ≤ n ≤ 500
    1 ≤ cost of cells ≤ 1000