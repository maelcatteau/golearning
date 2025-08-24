package main

import "testing"

func TestBonjour(t *testing.T) {

	verifierMessageCorrect := func(t testing.TB, resultat, attendu string) {
		t.Helper()
		if resultat != attendu {
			t.Errorf("reçu %q attendu %q", resultat, attendu)
		}
	}
	t.Run("dire bonjour aux personnes", func(t *testing.T) {
		resultat := Bonjour("Chris", "Francais")
		attendu := "Bonjour, Chris"
		verifierMessageCorrect(t, resultat, attendu)
	})
	t.Run("dire bonjour Monde quand une chaine vide est fournie, verifie egalement la langue par défaut", func(t *testing.T) {
		resultat := Bonjour("", "")
		attendu := "Bonjour, Monde"
		verifierMessageCorrect(t, resultat, attendu)
	})

	t.Run("en espagnol", func(t *testing.T) {
		resultat := Bonjour("Elodie", "Espagnol")
		attendu := "Hola, Elodie"
		verifierMessageCorrect(t, resultat, attendu)
	})
	t.Run("en anglais", func(t *testing.T) {
		resultat := Bonjour("Elodie", "Anglais")
		attendu := "Hello, Elodie"
		verifierMessageCorrect(t, resultat, attendu)
	})

}
