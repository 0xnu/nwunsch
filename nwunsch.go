package nwunsch

func NeedlemanWunsch(seq1, seq2 string, match, mismatch, gap int) (int, string, string) {
	m, n := len(seq1), len(seq2)
	dp := generateDPTable(m, n, gap)

	fillDPTable(seq1, seq2, dp, match, mismatch, gap)

	align1, align2 := traceback(seq1, seq2, dp, match, mismatch, gap)

	return dp[m][n], align1, align2
}

func generateDPTable(m, n, gap int) [][]int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i * gap
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j * gap
	}
	return dp
}

func fillDPTable(seq1, seq2 string, dp [][]int, match, mismatch, gap int) {
	for i := 1; i <= len(seq1); i++ {
		for j := 1; j <= len(seq2); j++ {
			matchScore := dp[i-1][j-1]
			score := mismatch
			if seq1[i-1] == seq2[j-1] {
				score = match
			}
			dp[i][j] = max(matchScore+score, dp[i-1][j]+gap, dp[i][j-1]+gap)
		}
	}
}

func traceback(seq1, seq2 string, dp [][]int, match, mismatch, gap int) (string, string) {
	align1, align2 := "", ""
	i, j := len(seq1), len(seq2)
	for i > 0 && j > 0 {
		score := mismatch
		if seq1[i-1] == seq2[j-1] {
			score = match
		}
		switch {
		case dp[i][j] == dp[i-1][j-1]+score:
			align1, align2 = string(seq1[i-1])+align1, string(seq2[j-1])+align2
			i, j = i-1, j-1
		case dp[i][j] == dp[i-1][j]+gap:
			align1, align2 = string(seq1[i-1])+align1, "-"+align2
			i--
		default:
			align1, align2 = "-"+align1, string(seq2[j-1])+align2
			j--
		}
	}

	for ; i > 0; i-- {
		align1, align2 = string(seq1[i-1])+align1, "-"+align2
	}
	for ; j > 0; j-- {
		align1, align2 = "-"+align1, string(seq2[j-1])+align2
	}

	return align1, align2
}

func max(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
		return c
	}
	if b > c {
		return b
	}
	return c
}
