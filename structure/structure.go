package structure

import "math"

type Rectangle struct {
	longueur float64
	largeur  float64
}
type Cercle struct {
	rayon float64
}
type Forme interface {
	Aire() float64
}

func Perimetre(rectangle Rectangle) (perimetre float64) {
	perimetre = 2 * (rectangle.longueur + rectangle.largeur)
	return
}

func (c Cercle) Aire() (aire float64) {
	aire = (c.rayon * c.rayon) * math.Pi
	return
}

func (r Rectangle) Aire() (aire float64) {
	aire = r.longueur * r.largeur
	return
}
