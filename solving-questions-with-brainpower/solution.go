package solvingquestionswithbrainpower

func mostPoints(questions [][]int) int64 {
	n := int64(len(questions))
	dp := make([]int64, n+1)
	for i := n - 1; i >= 0; i-- {
		points, brainpower := int64(questions[i][0]), int64(questions[i][1])
		if i+brainpower+1 < n {
			dp[i] = max(dp[i+1], points+dp[i+brainpower+1])
		} else {
			dp[i] = max(dp[i+1], points)
		}
	}
	return dp[0]
}
