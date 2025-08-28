package pointeurs

import "testing"

func TestPortefeuille(t *testing.T) {
	t.Run("Déposer", func(t *testing.T) {
		portefeuille := Portefeuille{}
		portefeuille.Deposer(Bitcoin(10))
		verifieSolde(t, portefeuille, Bitcoin(10))
	})
	t.Run("Retirer avec fonds suffisants", func(t *testing.T) {
		portefeuille := Portefeuille{solde: Bitcoin(20)}
		erreur := portefeuille.Retirer(Bitcoin(10))
		verifieSolde(t, portefeuille, Bitcoin(10))
		verifieAucuneErreur(t, erreur)
	})

	t.Run("Retirer avec des fonds insuffisants", func(t *testing.T) {
		soldeInitial := Bitcoin(10)
		portefeuille := Portefeuille{soldeInitial}
		erreur := portefeuille.Retirer(Bitcoin(20))
		verifieSolde(t, portefeuille, soldeInitial)
		verifieErreur(t, erreur, ErrFondsInsuffisants)
	})
}
func verifieSolde(t testing.TB, portefeuille Portefeuille, attendu Bitcoin) {
	t.Helper()
	solde := portefeuille.Solde()
	if solde != attendu {
		t.Errorf("solde %s, attendu %s", solde, attendu)
	}
}
func verifieErreur(t testing.TB, erreur error, attendu error) {
	t.Helper()
	if erreur == nil {
		t.Fatal("n'a pas reçu d'erreur mais en voulait une")
	}
	if erreur != attendu {
		t.Errorf("erreur attendu %s, mais recu %s", attendu, erreur)
	}
}
func verifieAucuneErreur(t testing.TB, erreur error) {
	t.Helper()
	if erreur != nil {
		t.Fatal("a reçu une erreur mais n'en attendait aucune")
	}
}
