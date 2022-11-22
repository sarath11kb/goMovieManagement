package Models

type Movie struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Genre    string `json:"genre"`
	Rating   int    `json:"rating"`
	Plot     string `json:"plot"`
	Released bool   `json:"released"`
}
