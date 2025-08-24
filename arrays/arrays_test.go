package arrays

import "testing"

func TestSomme(t *testing.T) {
	t.Run("Test avec un tableau de 5 nombres", func(t *testing.T) {
		nombres := []int{1, 2, 3, 4, 5}
		resultat := Somme(nombres)
		attendu := 15
		if resultat != attendu {
			t.Errorf("attendu %d mais reçu %d, %v", attendu, resultat, nombres)
		}
	})
}
