package dependances

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Saluer(writer io.Writer, nom string) {
	fmt.Fprintf(writer, "Bonjour, %s", nom)
}

func HandlerMonSalut(w http.ResponseWriter, r *http.Request) {
	Saluer(w, "Monde")
}
func main() {
	Saluer(os.Stdout, "Elodie")

	logger := log.New(os.Stdout, "INFO :", log.LstdFlags)
	Saluer(logger.Writer(), "Elodie")
}
