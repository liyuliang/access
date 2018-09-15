package access

import "testing"

func TestSortMap(t *testing.T) {
	m := make(map[string]string)
	m["error"] = "xxxxxx"
	m["code"] = "xxxxxx"
	m["message"] = "xxxxxx"
	m["list"] = "xxxxxx"

	for _, v := range SortMap(m) {
		println(v.Key, v.Val)
	}
}
