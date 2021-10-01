package models

type Medico struct {
	IdMedico        int
	Nome            string
	Cognome         string
	CodiceFiscale   string
	CittaStudio     string
	CapStudio       string
	ViaStudio       string
	CivicoStudio    int
	Prefisso        string
	Telefono        string
	Email           string
	EmailConfermata bool
	Psw             string
	CodiceRegionale string
}
