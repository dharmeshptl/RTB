package helper

func DoesStringArrayContain(needle string, haystack []string) bool {
	for _, ele := range haystack {
		if ele == needle {
			return true
		}
	}

	return false
}

func DoesIntArrayContain(needle int, haystack []int) bool {
	for _, ele := range haystack {
		if ele == needle {
			return true
		}
	}

	return false
}
