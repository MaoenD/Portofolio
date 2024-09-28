package GestionBDD

type Projet struct { // Declare the struct Projet. This struct holds the project information.
	Id_Projet   int
	Nom_Projet  string
	Description string
	Date_Debut  string
	Date_Fin    string
	Durée       string
}

type Formation struct { // Declare the struct Formation. This struct holds the formation information.
	Id_Formation  int
	Nom_Formation string
	Description   string
	Date_Debut    string
	Date_Fin      string
	Durée         string
}
