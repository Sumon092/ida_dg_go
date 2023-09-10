package user

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	PhoneNo string `json:"phone_no"`
	Address string `json:"address"`
}
