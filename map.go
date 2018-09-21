package access

import "sort"

type pair struct {
	Key string
	Val string
}

func SortMap(m map[string]string) (pairs []pair) {

	keys := []string{}
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := m[k]
		if k != "" && v != "" {
			pairs = append(pairs, pair{
				Key: k,
				Val: v,
			})
		}
	}

	return pairs
}
