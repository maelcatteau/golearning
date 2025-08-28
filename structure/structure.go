package structure

import "math"

type Rectangle struct {
	Longueur float64
	Largeur  float64
}
type Cercle struct {
	Rayon float64
}
type Triangle struct {
	Base    float64
	Hauteur float64
}
type Forme interface {
	Aire() float64
}

func Perimetre(rectangle Rectangle) (perimetre float64) {
	perimetre = 2 * (rectangle.Longueur + rectangle.Largeur)
	return
}

func (c Cercle) Aire() (Aire float64) {
	Aire = (c.Rayon * c.Rayon) * math.Pi
	return
}

func (r Rectangle) Aire() (Aire float64) {
	Aire = r.Longueur * r.Largeur
	return
}
func (t Triangle) Aire() (Aire float64) {
	Aire = (t.Base * t.Hauteur) * 0.5
	return
}
