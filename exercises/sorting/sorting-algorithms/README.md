
Bubble Sort:

    Time Complexity: O(N2)
    Auxiliary Space: O(1)

    The worst-case condition for bubble sort occurs when elements of the array are arranged in decreasing order.
    In the worst case, the total number of iterations or passes required to sort a given array is (n-1). where ‘n’ is a number 
    of elements present in the array.


Insertion Sort:

    Time Complexity: O(N^2)
    Auxiliary Space: O(1)
    
    When is the Insertion Sort algorithm used?
    Insertion sort is used when number of elements is small. It can also be useful when input array is almost sorted, only 
    few elements are misplaced in complete big array.


Merge Sort
    Time Complexity: O(nLogn)
    Auxiliary Space: O(n)

    Worst Case Complexity - It occurs when the array elements are required to be sorted in reverse order. That means suppose 
    you have to sort the array elements in ascending order, but its elements are in descending order. 

    Important Characteristics of Merge Sort:
        Merge Sort is useful for sorting linked lists.
        Merge Sort is a stable sort which means that the same element in an array maintain their original positions with
        respect to each other.
        The space complexity of Merge sort is O(n). This means that this algorithm takes a lot of space and may slower 
        down operations for the last data sets.

    Documentation: https://www.youtube.com/watch?v=jlHkDBEumP0&ab_channel=Jenny%27slecturesCS%2FITNET%26JRF


Quick Sort
    Time Complexity: O(N^2)
    Auxiliary Space: O(n*logn)

    Worst Case Complexity - This happens when we encounter the most unbalanced partitions possible, then the original 
    call takes n time, the recursive call on n-1 elements will take (n-1) time, the recursive call on (n-2) elements will 
    take (n-2) time, and so on. The worst case time complexity of Quick Sort would be O(n2).

    What is Randomised Quick Sort? Why is it used?
    
    Sometimes, it happens that by choosing the rightmost element at all times might result in the worst case scenario.
    In such cases, choosing a random element as your pivot at each step will reduce the probability of triggering the worst case behavior. 
    We will be more likely choosing pivots closer to the center of the array, and when this happens, the recursion branches more evenly and 
    thus the algorithm terminates a lot faster. The runtime complexity is expected to be O(n log n) as the selected random pivots are supposed 
    to avoid the worst case behavior.

    Why Quick Sort is better than Merge Sort?

    - Auxiliary Space : Quick sort is an in-place sorting algorithm whereas Merge sort uses extra space. In-place sorting 
    means no additional storage space is used to perform sorting (except recursion stack). Merge sort requires a new temporary 
    array to merge the sorted arrays thereby making Quick sort the better option.
    - Worst Cases : The worst case runtime of quick sort is O(n2) can be avoided by using randomized quicksort as explained 
    in the previous point. Obtaining average case behavior by choosing random pivot element improves the performance and becomes 
    as efficient as merge sort.
    - Cache Friendly: Quick Sort is also a cache friendly sorting algorithm as it has good locality of reference when used 
    for arrays. 
     
  Documentation: 
  - https://www.youtube.com/watch?v=jlHkDBEumP0&ab_channel=Jenny%27slecturesCS%2FITNET%26JRF
  - https://www.interviewbit.com/tutorial/quicksort-algorithm/#h705ske1hs7p7v4lg5h1cde5xlmtz0ns
  - https://www.youtube.com/watch?v=-qOVVRIZzao&ab_channel=AbdulBari

Heap Sort
    
    Documentation:
    - https://www.youtube.com/watch?v=NEtwJASLU8Q&ab_channel=Jenny%27slecturesCS%2FITNET%26JRF
    - https://www.youtube.com/watch?v=Q_eia3jC9Ts&ab_channel=Jenny%27slecturesCS%2FITNET%26JRF
