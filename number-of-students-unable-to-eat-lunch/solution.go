package numberofstudentsunabletoeatlunch

func countStudents(students []int, sandwiches []int) int {
	attempt := 0
	for attempt < len(students) {
		student := students[0]
		sandwich := sandwiches[0]
		if student == sandwich {
			students = students[1:]
			sandwiches = sandwiches[1:]
			attempt = 0
			continue
		}
		students = append(students[1:], student)
		attempt++
	}

	return len(students)
}
