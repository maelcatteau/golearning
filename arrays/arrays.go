package arrays

func Somme(nombres []int) (somme int) {
	for _, nombre := range nombres {
		somme += nombre
	}
	return
}
