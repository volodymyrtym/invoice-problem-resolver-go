package login

type createCommand struct {
	Email    string `json:"Email"`
	Password string `json:"HashedPassword"`
}
