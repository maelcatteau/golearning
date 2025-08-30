package dependances

import (
	"bytes"
	"testing"
)

func TestSaluer(t *testing.T) {
	buffer := bytes.Buffer{}
	Saluer(&buffer, "Chris")

	resultat := buffer.String()
	attendu := "Bonjour, Chris"

	if resultat != attendu {
		t.Errorf("resultat '%s', attendu '%s'", resultat, attendu)
	}
}
