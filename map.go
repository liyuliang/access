package access

import "sort"

type pair struct {
	Key string
	Val string
}

func SortMap(m map[string]string) (pairs []pair) {

	keys := make([]string, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	pairs = make([]pair, len(m))

	for _, k := range keys {
		pairs = append(pairs, pair{
			Key: k,
			Val: m[k],
		})
	}

	return pairs
}
