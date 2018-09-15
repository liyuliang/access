package access

import "sort"

func SortMap(m map[string]string) (result map[string]string) {

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	result = make(map[string]string)
	for _, k := range keys {
		result[k] = m[k]
	}
	return result
}
