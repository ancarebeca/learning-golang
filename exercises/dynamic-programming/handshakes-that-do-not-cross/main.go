package main

func countHandShakes(numPeople int) int {
	dp := make([]int, numPeople+1)
	dp[0] = 1

	for n := 2; n <= numPeople; n = +2 {
		for i := 0; i <= n-2; i = +2 {
			dp[n] += dp[i] * dp[n-2+i]
		}
	}
	return dp[numPeople]

}

// N= 8
//dp[0] = 1
//
//n=2  <= numPeople = 8
//i=0  until i <= 0
//dp[2] += dp[0] * dp[0]         => dp[2] = 1
//
//n=4  <= numPeople = 8
//i=0 until i <= 2
//dp[4] += dp[0] * dp[2]
//
//i=2 until i <= 2
//dp[4] += dp[2] * dp[2]       => dp[4] = 2
//
//n=6  <= numPeople = 8
//i=0 until i <= 4
//dp[6] += dp[0] * dp[4]
//i=2 until i <= 4
//dp[6] += dp[2] * dp[2]
//i=4 until i <= 4
//dp[6] += dp[4] * dp[0]        => dp[6]  = 5
//
//n=8  <= numPeople = 8
//i=0 until i <= 6
//dp[8] += dp[0] * dp[6]
//i=2 until i <= 6
//dp[8] += dp[2] * dp[4]
//i=4 until i <= 6
//dp[8] += dp[4] * dp[2]
//i=6 until i <= 6
//dp[8] += dp[6] * dp[0]           => dp[8] =  5 + 2 + 2  + 5 =14
