package main

import "testing"

func TestSoma(t *testing.T){
	total := soma(15,10)

	if (total != 30){
		t.Errorf("Resultado da soma %d é invalido, era esperado %d", total, 30)
	}

}