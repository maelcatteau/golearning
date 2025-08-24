package main

import "testing"

func TestBonjour(t *testing.T) {
	resultat := Bonjour("Chris")
	attendu := "Bonjour, Chris"

	if resultat != attendu {
		t.Errorf("reçu %q attendu %q", resultat, attendu)
	}
}
