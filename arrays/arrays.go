package arrays

func Somme(nombres []int) (somme int) {
	for _, nombre := range nombres {
		somme += nombre
	}
	return
}

func SommeTout(nombresASommer ...[]int) (sommes []int) {

	for _, nombres := range nombresASommer {
		sommes = append(sommes, Somme(nombres))
	}
	return

}

func SommeToutQueue(queueASommer ...[]int) (sommes []int) {

	for _, nombres := range queueASommer {
		if len(nombres) == 0 {
			sommes = append(sommes, 0)
		} else {
			queue := nombres[1:]
			sommes = append(sommes, Somme(queue))
		}
	}
	return
}
