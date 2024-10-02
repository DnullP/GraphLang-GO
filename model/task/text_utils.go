package task

import "strings"

func removeFirstAndLastLine(input string) string {
	lines := strings.Split(input, "\n")

	if len(lines) <= 2 {
		return ""
	}

	lines = lines[1 : len(lines)-1]
	result := strings.Join(lines, "\n")

	return result
}
