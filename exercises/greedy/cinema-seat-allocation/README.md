Cinema Seat Allocation

A cinema has n rows of seats, numbered from 1 to n and there are ten seats in each row, labelled from 1 to 10 as shown in the figure in this [link](https://leetcode.com/problems/cinema-seat-allocation/).

Given the array reservedSeats containing the numbers of seats already reserved, for example, reservedSeats[i] = [3,8] 
means the seat located in row 3 and labelled with 8 is already reserved.

Return the maximum number of four-person groups you can assign on the cinema seats. A four-person group occupies four 
adjacent seats in one single row. Seats across an aisle (such as [3,3] and [3,4]) are not considered to be adjacent, but 
there is an exceptional case on which an aisle split a four-person group, in that case, the aisle split a four-person 
group in the middle, which means to have two people on each side.

Hint 1
Note you can allocate at most two families in one row.

Hint 2
Greedily check if you can allocate seats for two families, one family or none.

Hint 3
Process only rows that appear in the input, for other rows you can always allocate seats for two families.

Solutions:
//https://github.com/letientai299/leetcode/blob/61ee5835463dea796760386a9574d33dfd1e9b8f/go/1386.cinema-seat-allocation.go
//https://github.com/atidat/leetcode-diary/blob/0869950fa841cb9ea1d7f705b203b6bbb33829fe/leetcode/1386-M-cinema-seat-allocation/main.go
//https://github.com/yuzixun/algorithm_exercise/blob/2087972d330fcff5801c7ea55260e1cf600031e3/22/2.go
//https://github.com/Crshi/LeetCode-go/blob/1cc62d2feafe2160fd518e09dc721562f52386af/LeetCode/%E5%8F%8C%E5%91%A8%E8%B5%9B/%E5%AE%89%E6%8E%92%E7%94%B5%E5%BD%B1%E9%99%A2%E5%BA%A7%E4%BD%8D/main.go
//https://github.com/Lihe97/Leetcode-Training/blob/579c3030241f9e1e3d7870ffbd4fa7799b749804/Week_Contest/5349.%20%E5%AE%89%E6%8E%92%E7%94%B5%E5%BD%B1%E9%99%A2%E5%BA%A7%E4%BD%8D.go
//https://github.com/rhzx3519/leetcode/blob/fd4b0dd9eacf042c3ff0d41810cb45f870e20eba/go/src/leetcode/1386.%20Cinema%20Seat%20Allocation/1386.%20Cinema%20Seat%20Allocation.go
