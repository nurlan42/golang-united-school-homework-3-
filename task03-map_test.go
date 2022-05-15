package homework

import (
	"testing"
)

func TestSortMapValues(t *testing.T) {
	val := map[int]string{10: "aa", 0: "bb", 500: "cc"}
	want := []string{"bb", "aa", "cc"}

	got := sortMapValues(val)
	for i := 0; i < len(got); i++ {
		for j := 0; j < len(want); j++ {
			if got[i] == want[j] {
				break
			}
		}
	}
}
