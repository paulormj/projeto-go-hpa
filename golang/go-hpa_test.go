package main
import (
  "testing"
)

func TestGoHpa(t *testing.T){
	got:= executa_calculo()
	want:= 2499999414254808.5

	if got != want {
		t.Errorf("Calculo errado")
	}
}