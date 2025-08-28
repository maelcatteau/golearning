package pointeurs

import (
	"errors"
	"fmt"
)

var ErrFondsInsuffisants = errors.New("impossible de retirer, fonds insuffisants")

type Bitcoin int

type Portefeuille struct {
	solde Bitcoin
}

func (p *Portefeuille) Deposer(montant Bitcoin) {
	p.solde += montant
}
func (p *Portefeuille) Retirer(montant Bitcoin) error {
	if montant > p.solde {
		return ErrFondsInsuffisants
	}

	p.solde -= montant
	return nil
}

func (p *Portefeuille) Solde() Bitcoin {
	return p.solde
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
