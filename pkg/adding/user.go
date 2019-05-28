package adding

// User defines the information to be added to the
type User struct {
	PlatformID     string    `json:"platformID"`
	Platform       string    `json:"platform"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	Email          string    `json:"email"`
	ProfilePicture string    `json:"pictureURL"`
	Description    string    `json:"description"`
}
