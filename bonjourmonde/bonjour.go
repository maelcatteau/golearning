package main

const anglais = "Anglais"
const francais = "Français"
const espagnol = "Espagnol"

const prefixeSalutAnglais = "Hello, "
const prefixeSalutFrancais = "Bonjour, "
const prefixeSalutEspagnol = "Hola, "

func Bonjour(nom string, langue string) string {
	if nom == "" {
		nom = "Monde"
	}
	return prefixeSalut(langue) + nom
}

func prefixeSalut(langue string) (prefixe string) {
	switch langue {
	case anglais:
		prefixe = prefixeSalutAnglais
	case espagnol:
		prefixe = prefixeSalutEspagnol
	default:
		prefixe = prefixeSalutFrancais
	}
	return
}
