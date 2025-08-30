package dictionnaire

import "testing"

func TestRecherche(t *testing.T) {
	dictionnaire := Dictionnaire{"test": "ceci est juste un test"}

	t.Run("mot connu", func(t *testing.T) {
		resultat, _ := dictionnaire.Recherche("test")
		attendu := "ceci est juste un test"

		assertChaineEgal(t, resultat, attendu)
	})
	t.Run("mot inconnu", func(t *testing.T) {
		_, err := dictionnaire.Recherche("inconnu")

		assertErreur(t, err, ErrMotInexistant)
	})
}
func TestAjouter(t *testing.T) {
	t.Run("nouveau mot", func(t *testing.T) {
		dictionnaire := Dictionnaire{}
		mot := "test"
		definition := "ceci est juste un test"
		err := dictionnaire.Ajouter(mot, definition)

		assertDefinition(t, dictionnaire, mot, definition)
		assertErreur(t, err, nil)
	})
	t.Run("Mot existant", func(t *testing.T) {
		mot := "test"
		definition := "ceci est juste un test"
		dictionnaire := Dictionnaire{mot: definition}

		err := dictionnaire.Ajouter(mot, "ceci est juste un test d'ecrasement")
		assertErreur(t, err, ErrMotExistant)
		assertDefinition(t, dictionnaire, mot, definition)
	})

}
func TestMettreAJour(t *testing.T) {

	t.Run("mot existant", func(t *testing.T) {
		mot := "test"
		definition := "ceci est juste un test"
		dictionnaire := Dictionnaire{mot: definition}
		nouvelleDefinition := "nouvelle definition"

		err := dictionnaire.MettreAJour(mot, nouvelleDefinition)
		assertDefinition(t, dictionnaire, mot, nouvelleDefinition)
		assertErreur(t, err, nil)

	})
	t.Run("mot inexistant", func(t *testing.T) {
		dictionnaire := Dictionnaire{}
		mot := "test"
		nouvelleDefinition := "nouvelle definition"
		err := dictionnaire.MettreAJour(mot, nouvelleDefinition)
		assertErreur(t, err, ErrMAJImpossible)
	})

}
func TestSupprimer(t *testing.T) {
	t.Run("Mot existant", func(t *testing.T) {
		mot := "test"
		definition := "ceci est juste un test"
		dictionnaire := Dictionnaire{mot: definition}
		dictionnaire.Supprimer(mot)
		_, err := dictionnaire.Recherche(mot)
		assertErreur(t, err, ErrMotInexistant)
	})
	t.Run("Mot inexistant", func(t *testing.T) {
		mot := "test"
		dictionnaire := Dictionnaire{}
		err := dictionnaire.Supprimer(mot)
		assertErreur(t, err, ErrSupprImpossible)
	})

}

func assertChaineEgal(t testing.TB, resultat string, attendu string) {
	if resultat != attendu {
		t.Errorf("attendu '%s' mais reçu '%s'", attendu, resultat)
	}
}
func assertErreur(t testing.TB, recu, attendu error) {
	if recu != attendu {
		t.Errorf("erreur attendue : '%s', mais recu '%s'", attendu, recu)
	}
}
func assertDefinition(t testing.TB, d Dictionnaire, mot, definition string) {
	t.Helper()
	resultat, err := d.Recherche(mot)
	if err != nil {
		t.Fatal("ne devrait pas retourner une erreur", err)
	}
	if definition != resultat {
		t.Errorf("attendu '%s' mais reçu '%s'", definition, resultat)
	}
}
