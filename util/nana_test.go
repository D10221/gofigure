package util

import "testing"

func Test_Nada(t *testing.T)  {
	nada := Nada();
	if nada == "" && nada == "Nada"{
		t.Error("No Nada == Somtehing")
	}
}
