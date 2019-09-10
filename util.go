package main

//UniqueValuesInList extract unique values in a list
func UniqueValuesInList(arraySlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range arraySlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

//IsKeyValueEqualInMap if value exists in array
func IsKeyValueEqualInMap(managerEmployeeMap map[string]string) string {
	for k, v := range managerEmployeeMap {
		if k == v {
			return k
		}
	}
	return ""
}
