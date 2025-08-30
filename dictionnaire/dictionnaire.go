package dictionnaire

import "errors"

var (
	ErrMotInexistant   = errors.New("impossible de trouver le mot que vous recherchez")
	ErrMotExistant     = errors.New("le mot que vous cherchez à définir existe déjà")
	ErrMAJImpossible   = errors.New("maj impossible le mot est inexistant")
	ErrSupprImpossible = errors.New("suppression impossible le mot est inexistant")
)

type Dictionnaire map[string]string

func (d Dictionnaire) Recherche(mot string) (string, error) {
	definition, ok := d[mot]
	if !ok {
		return "", ErrMotInexistant
	}
	return definition, nil
}
func (d Dictionnaire) Ajouter(mot string, definition string) error {
	_, err := d.Recherche(mot)
	switch err {
	case ErrMotInexistant:
		d[mot] = definition
	case nil:
		return ErrMotExistant
	default:
		return err
	}
	return nil
}
func (d Dictionnaire) MettreAJour(mot string, nouvelleDef string) error {
	_, err := d.Recherche(mot)
	switch err {
	case ErrMotInexistant:
		return ErrMAJImpossible
	case nil:
		d[mot] = nouvelleDef
	default:
		return err
	}
	return nil
}
func (d Dictionnaire) Supprimer(mot string) error {
	_, err := d.Recherche(mot)
	switch err {
	case ErrMotInexistant:
		return ErrSupprImpossible
	case nil:
		delete(d, mot)
	default:
		return err
	}
	return nil
}
