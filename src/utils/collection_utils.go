package utils

func CopyMap(input map[string]string) map[string]string {
	retv := make(map[string]string)

	for key, value := range input {
		retv[key] = value
	}
	return retv
}
