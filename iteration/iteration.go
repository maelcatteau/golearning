package iteration

import "strings"

func Repeter(caractere string, nombre int) string {
	var repete strings.Builder
	for i := 0; i < nombre; i++ {
		repete.WriteString(caractere)
	}
	return repete.String()
}
