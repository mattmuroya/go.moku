package common

func IndexOf(input string, data []string) int {
	for i, e := range data {
		if e == input {
			return i
		}
	}
	return -1
}
