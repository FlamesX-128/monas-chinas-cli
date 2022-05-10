package tools

func FindIndex(arr []string, target string) uint16 {
	for i, e := range arr {
		if e == target {

			return uint16(i)
		}

	}

	return 0
}
