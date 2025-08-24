package main

import "fmt"

func Bonjour(nom string) string {
	return "Bonjour, " + nom
}

func main() {
	fmt.Println(Bonjour("Chris"))
}
