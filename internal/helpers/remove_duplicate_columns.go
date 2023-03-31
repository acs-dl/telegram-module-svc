package helpers

import "strings"

func RemoveDuplicateColumn(arr []string) []string {
	allKeys := make(map[string]bool)
	var list []string

	for i := range arr {
		splittedColumName := strings.Split(arr[i], ".")
		if len(splittedColumName) != 2 {
			continue
		}

		columnName := splittedColumName[1] // [0] is table name

		if _, value := allKeys[columnName]; !value {
			allKeys[columnName] = true
			list = append(list, arr[i])
		}
	}

	return list
}
