package ladoktypes

// InternalError type
type InternalError struct {
	Msg           string `json:"msg"`
	Type          string `json:"type"`
	Func          string `json:"func"`
	PreviousError string `json:"previous_error"`
}

// LadokError returns by Ladok
type LadokError struct {
	Detaljkod       string        `json:"Detaljkod"`
	DetaljkodText   string        `json:"DetaljkodText"`
	FelUID          string        `json:"FelUID"`
	Felgrupp        string        `json:"Felgrupp"`
	FelgruppText    string        `json:"FelgruppText"`
	Felkategori     string        `json:"Felkategori"`
	FelkategoriText string        `json:"FelkategoriText"`
	Meddelande      string        `json:"Meddelande"`
	Link            []interface{} `json:"link"`
}
