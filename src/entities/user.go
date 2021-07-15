package entities

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"` // こうすると、jsonに変換した際にomitされる
}
