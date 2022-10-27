package models

type Condition struct {
	ID    int
	Water int `json:"water"`
	Wind  int `json:"wind"`
}
