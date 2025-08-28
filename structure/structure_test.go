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

	testsAire := []struct {
		nom   string
		forme Forme
		aAire float64
	}{
		{nom: "Rectangle", forme: Rectangle{Largeur: 12.0, Longueur: 6.0}, aAire: 72.0},
		{nom: "Cercle", forme: Cercle{Rayon: 10.0}, aAire: 314.1592653589793},
		{nom: "Triangle", forme: Triangle{Base: 12.0, Hauteur: 6.0}, aAire: 36.0},
	}
	for _, tt := range testsAire {
		t.Run(tt.nom, func(t *testing.T) {
			resultat := tt.forme.Aire()
			if resultat != tt.aAire {
				t.Errorf("%#v, reçu %g mais attendu %g", tt.forme, resultat, tt.aAire)
			}
		})
	}

}
