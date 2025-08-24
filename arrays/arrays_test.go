package arrays

import (
	"reflect"
	"testing"
)

var verifierSommes = func(t testing.TB, resultat, attendu []int) {
	t.Helper()
	if !reflect.DeepEqual(resultat, attendu) {
		t.Errorf("reçu %v attendu %v", resultat, attendu)
	}
}

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

func TestSommeTout(t *testing.T) {
	t.Run("Test de Somme tout de base", func(t *testing.T) {

		resultat := SommeTout([]int{1, 2}, []int{0, 9})
		attendu := []int{3, 9}

		verifierSommes(t, resultat, attendu)

	})
}

func TestSommeToutQueue(t *testing.T) {
	t.Run("Test de SommeToutQueue", func(t *testing.T) {

		resultat := SommeToutQueue([]int{1, 2, 3}, []int{5, 1, 1})
		attendu := []int{5, 2}

		verifierSommes(t, resultat, attendu)
	})

	t.Run("Test de sommage de slice vide", func(t *testing.T) {
		resultat := SommeToutQueue([]int{}, []int{1, 2, 3})
		attendu := []int{0, 5}

		verifierSommes(t, resultat, attendu)
	})
}
