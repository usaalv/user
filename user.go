package user

// User ...
type User struct {
	Name  string
	Email string
}

// Create ...
func Create(name string, email string) *User {
	u := User{
		Name:  name,
		Email: email,
	}

	return &u

}
