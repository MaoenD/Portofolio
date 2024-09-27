package GestionBDD

type Projet struct {
	Id_Projet   int
	Nom_Projet  string
	Description string
	Date_Debut  string
	Date_Fin    string
	Durée       string
}

type Formation struct {
	Id_Formation  int
	Nom_Formation string
	Description   string
	Date_Debut    string
	Date_Fin      string
	Durée         string
}
