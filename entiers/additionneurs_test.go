package entiers

import (
	"fmt"
	"testing"
)

func TestAdditionneurs(t *testing.T) {
	somme := Additionner(2, 2)
	attendu := 4

	if somme != attendu {
		t.Errorf("attendu %d mais reçu %d", attendu, somme)
	}
}

func ExampleAdditionner() {
	somme := Additionner(1, 5)
	fmt.Println(somme)
	// Output: 6
}
