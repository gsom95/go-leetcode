package takingmaximumenergyfromthemysticdungeon

func maximumEnergy(energy []int, k int) int {
    lene := len(energy)
    dp := make([]int, lene)

    for i := lene-1; i >= 0; i-- {
        dp[i] = energy[i]
        if i+k < lene {
            dp[i] += dp[i+k]
        }
    }
    
    maxE := -1001
    for _, p := range dp {
        maxE = max(maxE, p)
    }


    return maxE
}