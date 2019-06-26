package access

import (
	"fmt"
	"testing"
)

func TestSortMap(t *testing.T) {
	m := make(map[string]string)
	m["error"] = "xxxxxx"
	m["code"] = "xxxxxx"
	m["message"] = "xxxxxx"
	m["list"] = "xxxxxx"

	newMap := SortMap(m)
	if newMap[0].Key != "code" || newMap[0].Val != "xxxxxx" {
		e := fmt.Sprintf("Sort Map Wrong, which should be %s, but get %s", "code", newMap[0].Key)
		t.Error(e)
	}
	if newMap[1].Key != "error" || newMap[1].Val != "xxxxxx" {
		e := fmt.Sprintf("Sort Map Wrong, which should be %s, but get %s", "error", newMap[0].Key)
		t.Error(e)
	}
	if newMap[2].Key != "list" || newMap[2].Val != "xxxxxx" {
		e := fmt.Sprintf("Sort Map Wrong, which should be %s, but get %s", "list", newMap[0].Key)
		t.Error(e)
	}
	if newMap[3].Key != "message" || newMap[3].Val != "xxxxxx" {
		e := fmt.Sprintf("Sort Map Wrong, which should be %s, but get %s", "message", newMap[0].Key)
		t.Error(e)
	}
}
