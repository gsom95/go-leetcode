package exclusivetimeoffunctions

import (
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	times := make([]int, n)

	stack := make([]int, 0, len(logs))
	curPos := 0

	for _, log := range logs {
		splitLog := strings.Split(log, ":")
		id, _ := strconv.Atoi(splitLog[0])
		op := splitLog[1]
		ts, _ := strconv.Atoi(splitLog[2])
		if op == "start" {
			if len(stack) > 0 {
				last := stack[len(stack)-1]
				times[last] += ts - curPos
			}

			stack = append(stack, id)
			curPos = ts
			continue
		}

		last := stack[len(stack)-1]
		times[last] += ts - curPos + 1
		stack = stack[:len(stack)-1]
		curPos = ts + 1
	}

	return times
}
