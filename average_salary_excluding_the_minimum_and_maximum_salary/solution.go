package averagesalaryexcludingtheminimumandmaximumsalary

import "sort"

func average(salary []int) float64 {
	sort.Ints(salary)
	sum := 0
	for _, s := range salary[1 : len(salary)-1] {
		sum += s
	}
	return float64(sum) / float64(len(salary)-2)
}
