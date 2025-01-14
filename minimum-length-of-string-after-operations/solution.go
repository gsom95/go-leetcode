package minimumlengthofstringafteroperations

func minimumLength(s string) int {
    var (
        count = [26]int{}
        result = 0
    )

    for _, chr := range s {
        count[chr-'a']++
    }

    for i := 0; i < 26; i++ {
        if count[i] <= 2 {
            result += count[i]
            continue
        }

        if count[i] % 2 == 1 {
            result += 1
        } else {
            result += 2
        }
    }

    return result
}