package structure

import "testing"

func TestPerimetre(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	resultat := Perimetre(rectangle)
	attendu := 40.0

	if resultat != attendu {
		t.Errorf("attendu %.2f mais reçu %.2f", resultat, attendu)
	}
}

func TestAire(t *testing.T) {

	verifierAire := func(t testing.TB, forme Forme, attendu float64) {
		t.Helper()
		resultat := forme.Aire()
		if resultat != attendu {
			t.Errorf("attendu %g mais reçu %g", resultat, attendu)
		}
	}
	t.Run("Test rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		verifierAire(t, rectangle, 72.0)
	})

	t.Run("Test cercle", func(t *testing.T) {
		cercle := Cercle{10.0}
		verifierAire(t, cercle, 314.1592653589793)
	})
}
