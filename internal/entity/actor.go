package entity

type Actor struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Gender    string   `json:"gender"`
	Birthdate string   `json:"birthdate"`
	Movies    []string `json:"movies"`
}

func NewActor() Actor {
	return Actor{}
}
