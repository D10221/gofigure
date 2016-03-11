package inutil

import "testing"

func Test_maps(t *testing.T) {
	ones := map[string]int{"1": 1 }
	if one := ones["1"]; one != 1 || len(ones) != one {
		t.Error("One?")
	} else {
		t.Log("Ones: ", ones)
	}
	if uno, exists := ones["uno"]; exists && uno == 1 {
		t.Failed()
	}

}