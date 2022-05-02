package tools

func FindIndex(arr []string, target string) uint {
	for i, e := range arr {
		if e == target {

			return uint(i)
		}

	}

	return 0
}
