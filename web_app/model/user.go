package model

// User : user object
type User struct {
	FirstName string
	LastName  string
	Age       int
}

// UserAPIFormattedData :
type UserAPIFormattedData struct {
	FirstName     string
	LastName      string
	FormattedName string
	Age           int
}

// GetFormattedName :
func (u User) GetFormattedName() string {
	return u.FirstName + " " + u.LastName
}

// GetAPIFormattedData :
func (u User) GetAPIFormattedData() *UserAPIFormattedData {
	return &UserAPIFormattedData{
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		FormattedName: u.GetFormattedName(),
		Age:           u.Age,
	}
}
