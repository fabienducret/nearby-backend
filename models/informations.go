package models

type Informations struct {
	Weather Weather `json:"weather"`
	News    []News  `json:"news"`
}
